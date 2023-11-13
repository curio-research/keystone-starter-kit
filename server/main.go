package main

import (
	"time"

	"github.com/curio-research/keystone-starter-kit/server/constants"
	"github.com/curio-research/keystone-starter-kit/server/data"
	"github.com/curio-research/keystone-starter-kit/server/startup"
	gamedb "github.com/curio-research/keystone/db"
	startKeystone "github.com/curio-research/keystone/server/startup"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// GEVM
	gevm "github.com/daweth/gevm/core"
	"github.com/daweth/gevm/examples"
	vm "github.com/daweth/gevm/vm"
)

func main() {

	// Initialize new game engine
	ctx := startKeystone.NewGameEngine()

	ctx.SetTickRate(constants.TickRate)

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

	startKeystone.RegisterRewindEndpoint(ctx)

	ctx.SetSaveStateHandler(SQLiteSaveStateHandler, 0)
	ctx.SetSaveTxHandler(SQLiteSaveTxHandler, 0)

	ctx.SetSaveState(false)
	ctx.SetSaveTx(false)

	// Initialize game map
	startup.InitWorld(ctx)

	node := gevm.Default()
	vm.InitializeEngine(ctx)

	examples.RunPrecompile(&node)
	// Start game server!
	go ctx.Start()

	time.Sleep(30 * time.Second)
	examples.RunPrecompile(&node)
}
