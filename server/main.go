package main

import (
	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone-starter-kit/network"
	"github.com/curio-research/keystone-starter-kit/startup"
	gamedb "github.com/curio-research/keystone/db"
	ks "github.com/curio-research/keystone/server/startup"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	godotenv.Load()

	// Initialize new game engine
	ctx := ks.NewGameEngine()

	ctx.SetTickRate(constants.TickRate)

	ctx.SetEmitErrorHandler(&network.ProtoBasedErrorHandler{})
	ctx.SetEmitEventHandler(&network.ProtoBasedBroadcastHandler{})

	// Add systems
	startup.AddSystems(ctx)

	// Setup HTTP routes
	startup.SetupRoutes(ctx)

	// TODO: kevin: make this cleaner imo
	// Register tables schemas to world
	startup.RegisterTablesToWorld(ctx.World)

	// provision SQLite
	gormDB, err := gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	SQLiteSaveStateHandler, SQLiteSaveTxHandler, err := gamedb.SQLHandlersFromDialector(gormDB.Dialector, ctx.GameId, data.TableSchemasToAccessors)

	ctx.SetSaveStateHandler(SQLiteSaveStateHandler, 0)
	ctx.SetSaveTxHandler(SQLiteSaveTxHandler, 0)

	// Initialize game map
	startup.InitWorld(ctx)

	// Start game server!
	ctx.Start()

}
