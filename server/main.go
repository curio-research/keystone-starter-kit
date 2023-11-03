package main

import (
	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone-starter-kit/network"
	"github.com/curio-research/keystone-starter-kit/startup"
	gamedb "github.com/curio-research/keystone/db"
	ks "github.com/curio-research/keystone/server/startup"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// Initialize new game engine
	ctx := ks.NewGameEngine()

	ctx.SetTickRate(constants.TickRate)

	ctx.SetEmitErrorHandler(&network.ProtoBasedErrorHandler{})
	ctx.SetEmitEventHandler(&network.ProtoBasedBroadcastHandler{})

	// Add systems
	startup.AddSystems(ctx)

	// Setup HTTP routes
	startup.SetupRoutes(ctx)

	// Register tables schemas to world
	ctx.AddTables(data.TableSchemasToAccessors)

	// Provision local SQLite
	gormDB, err := gorm.Open(sqlite.Open("local.db"))
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	SQLiteSaveStateHandler, SQLiteSaveTxHandler, err := gamedb.SQLHandlersFromDialector(gormDB.Dialector, ctx.GameId, data.TableSchemasToAccessors)
	if err != nil {
		panic("failed to create sqlite save handlers: " + err.Error())
	}

	ctx.SetSaveStateHandler(SQLiteSaveStateHandler, 0)
	ctx.SetSaveTxHandler(SQLiteSaveTxHandler, 0)

	ctx.SetSaveState(false)
	ctx.SetSaveTx(false)

	// Initialize game map
	startup.InitWorld(ctx)

	// Start game server!
	ctx.Start()

}
