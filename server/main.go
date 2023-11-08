package main

import (
	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone-starter-kit/startup"
	gamedb "github.com/curio-research/keystone/db"
	startKeystone "github.com/curio-research/keystone/server/startup"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// Initialize new game engine
	ctx := startKeystone.NewGameEngine()

	ctx.SetTickRate(constants.TickRate)
	ctx.SetStreamRate(50)

	// Add systems
	startup.AddSystems(ctx)

	// Setup HTTP routes
	startup.SetupRoutes(ctx)

	// Register tables schemas to world
	ctx.AddTables(data.SchemaMapping)

	// Provision local SQLite
	gormDB, err := gorm.Open(sqlite.Open("local.db"))
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	SQLiteSaveStateHandler, SQLiteSaveTxHandler, err := gamedb.SQLHandlersFromDialector(gormDB.Dialector, ctx.GameId, data.SchemaMapping)
	if err != nil {
		panic("failed to create sqlite save handlers: " + err.Error())
	}

	startKeystone.RegisterRewindEndpoint(ctx)

	ctx.SetSaveStateHandler(SQLiteSaveStateHandler, 0)
	ctx.SetSaveTxHandler(SQLiteSaveTxHandler, 0)

	ctx.SetSaveState(false)
	ctx.SetSaveTx(false)

	// Initialize game map
	startup.InitWorld(ctx)

	// Start game server!
	ctx.Start()

}
