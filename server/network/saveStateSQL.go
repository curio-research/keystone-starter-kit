package network

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/helper"
	"github.com/curio-research/keystone/game/startup"

	"github.com/curio-research/keystone/game/tables"
	"github.com/curio-research/keystone/logging"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/gin-gonic/gin"
	"github.com/golang-collections/collections/stack"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLSaveStateHandler struct {
	DBConnection *gorm.DB
}

// initialize connection mySQL
func (handler *MySQLSaveStateHandler) InitializeDBConnection(dialector gorm.Dialector) error {
	if handler.DBConnection != nil {
		logging.Log().Error("db connection already exists")
		return fmt.Errorf("db connection already exists")
	}

	db, err := gorm.Open(dialector, &defaultGormOpts)
	if err != nil {
		return err
	}

	handler.DBConnection = db
	return nil
}

// initialize mySQL tables for saving state updates
func (handler *MySQLSaveStateHandler) InitializeDBTables() error {
	db := handler.DBConnection

	// all tables that need to be created
	allSchemas := []any{}
	for schema := range tables.TableSchemasToAccessors {
		allSchemas = append(allSchemas, schema)
	}

	// fetch a list of all existing tables
	var tableNames []string
	result := db.Raw("SHOW TABLES").Scan(&tableNames)
	if result.Error != nil {
		return result.Error
	}

	// loop through each table by name and drop them. this resets the entire database
	for _, tableName := range tableNames {
		if err := db.Migrator().DropTable(tableName); err != nil {
			return err
		}
	}

	fmt.Println("-> Existing tables have been removed")

	// create fresh tables
	err := db.AutoMigrate(allSchemas...)
	if err != nil {
		return err
	}

	fmt.Println("-> All tables have been created")
	return nil
	// tables are now ready to be used for inserting state updates
}

// save state updates to mySQL database
func (handler *MySQLSaveStateHandler) SaveState(tableUpdates []state.TableUpdate) error {
	// process table updates
	tableUpdateOperationsByTable, tableRemovalOperationsByTable := processUpdatesForUpload(tableUpdates)

	// update operations
	for table, updates := range tableUpdateOperationsByTable {
		arr := castToSchemaArray(table, updates)
		if arr != nil {
			tx := handler.DBConnection.Table(table).Save(arr)
			if tx.Error != nil {
				return tx.Error
			}
		}
	}

	// removal operations
	for table, removals := range tableRemovalOperationsByTable {
		arr := castToSchemaArray(table, removals)
		tx := handler.DBConnection.Table(table).Delete(arr)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

// restore state updates from mySQL database
func (handler *MySQLSaveStateHandler) RestoreState(ctx *server.EngineCtx, gameId string) error {
	gw := ctx.World
	for schema, tableAccessor := range tables.TableSchemasToAccessors {
		rows, err := handler.DBConnection.Table(tableAccessor.Name()).Rows()
		if err != nil {
			return err
		}

		for rows.Next() {
			obj, id, err := convertSQLRowToSchema(rows, schema)
			if err != nil {
				panic(err)
			}

			tableAccessor.Set(gw, id, obj)
		}
	}

	return nil
}

// only initialize tables, game tick, and request mapping
func newGameEngineForRestoreStateFromTxs(randSeed int) *server.EngineCtx {
	ctx := &server.EngineCtx{}
	tableBasedWorld := state.NewWorld()
	helper.InitGame(tableBasedWorld, randSeed)

	gameTick := server.NewGameTick(constants.TickRate)
	startup.AddSystems(gameTick)

	ctx.World = tableBasedWorld
	ctx.GameTick = gameTick

	return ctx
}

func convertSQLRowToSchema(rows *sql.Rows, schema interface{}) (interface{}, int, error) {
	// Validate that schema is a pointer to a struct
	v := reflect.ValueOf(schema)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil, -1, fmt.Errorf("schema must be a pointer to a struct")
	}

	t := v.Elem().Type()
	schemaStruct := reflect.New(t).Elem()
	fieldPointers := make([]interface{}, 0)

	s := stack.New()
	for i := t.NumField() - 1; i >= 0; i-- {
		s.Push(schemaStruct.Field(i))
	}

	idIndex := -1
	for s.Len() != 0 {
		val := s.Pop().(reflect.Value)
		if val.Kind() == reflect.Struct {
			for j := val.NumField() - 1; j >= 0; j-- {
				s.Push(val.Field(j))
			}
		} else {
			fieldPointers = append(fieldPointers, val.Addr().Interface())
		}
	}

	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		if name == "Id" {
			idIndex = i
		}
	}

	err := rows.Scan(fieldPointers...)
	if err != nil {
		return nil, -1, err
	}

	id := schemaStruct.Field(idIndex).Int()

	// Return the populated struct and its primary key
	return schemaStruct.Interface(), int(id), nil
}

func InitializeSQLHandlers(g *server.EngineCtx, mySQLDSN string) error {
	saveStateHandler := &MySQLSaveStateHandler{}
	dialector := mysql.Open(mySQLDSN)
	err := saveStateHandler.InitializeDBConnection(dialector)
	if err != nil {
		return err
	}

	err = saveStateHandler.InitializeDBTables()
	if err != nil {
		return err
	}

	txHandler, err := NewMySQLSaveTransactionHandler(dialector, g.RandSeed)
	if err != nil {
		return err
	}

	g.SaveStateHandler = saveStateHandler
	g.SaveTransactionsHandler = txHandler

	return nil
}

func processUpdatesForUpload(tableUpdates []state.TableUpdate) (TableToUpdatesMap, TableToUpdatesMap) {

	// parse the array backwards and store the table updates that are the "latest"
	// ex: if i updated a table row but then deleted it, only the deletion matters
	seenUpdateEntities := make(map[int]bool)
	updates := []state.TableUpdate{}

	for i := len(tableUpdates) - 1; i >= 0; i-- {
		update := tableUpdates[i]
		if !seenUpdateEntities[update.Entity] {
			updates = append(updates, update)
			seenUpdateEntities[update.Entity] = true
		}
	}

	return categorizeTableUpdatesBySchema(updates)

}

// returns: table name -> []value updates

type TableToUpdatesMap map[string][]any

func categorizeTableUpdatesBySchema(updates []state.TableUpdate) (TableToUpdatesMap, TableToUpdatesMap) {
	tableUpdateOperationsByTable := make(TableToUpdatesMap)
	tableRemovalOperationsByTable := make(TableToUpdatesMap)

	for _, update := range updates {
		table := update.Table

		if update.OP == state.UpdateOP {
			tableUpdateOperationsByTable[table] = append(tableUpdateOperationsByTable[table], update.Value)
		} else if update.OP == state.RemovalOP {
			tableRemovalOperationsByTable[table] = append(tableRemovalOperationsByTable[table], update.Value)
		}
	}

	return tableUpdateOperationsByTable, tableRemovalOperationsByTable
}

// given a schema type, use the mapping from tables to cast to an array of that type
func castToSchemaArray(schemaType string, val []interface{}) interface{} {
	var accessor *state.TableBaseAccessor[any]
	for _, schemaAccessor := range tables.TableSchemasToAccessors {
		if strings.Contains(schemaAccessor.Name(), schemaType) {
			accessor = schemaAccessor
			break
		}
	}
	if accessor == nil {
		return nil
	}

	schema := accessor.Type()

	// Use reflection to cast val to the appropriate schema type.
	arrayType := reflect.SliceOf(schema)
	castedValue := reflect.MakeSlice(arrayType, len(val), len(val))
	for i, v := range val {
		castedValue.Index(i).Set(reflect.ValueOf(v))
	}

	return castedValue.Interface()
}

// save state HTTP handler
type RestoreStateRequest struct {
	GameId string `json:"gameId"`
}

func RestoreStateHandler(ctx *server.EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {

		err := ctx.SaveStateHandler.RestoreState(ctx, ctx.GameId)
		if err != nil {
			logging.Log().Errorw("restore state failed", "err", err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, "success")
		}
	}
}
