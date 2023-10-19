package server

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/startup"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/tjarratt/babble"
)

func MainServer(randSeedNumber int) (*gin.Engine, error) {
	color.HiYellow("")
	color.HiYellow("---- 🗝  Powered by Keystone 🗿 ----")
	fmt.Println()

	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.Use(server.CORSMiddleware())

	// initialize gameCtx (should be refactored to keystone)
	gameId := babble.NewBabbler().Babble()
	gameCtx := setupWorld(gameId)

	// initialize in-game world
	startup.InitGame(gameCtx.World)
	color.HiWhite("Tick rate:         " + strconv.Itoa(gameCtx.GameTick.TickRateMs) + "ms")

	// setup server routes
	startup.SetupRoutes(s, gameCtx)

	return s, nil
}

// TODO this function should be in keystone
func setupWorld(gameId string) *server.EngineCtx {
	gameWorld := state.NewWorld()
	gameTick := server.NewGameTick(constants.TickRate)

	// add systems for game
	gameTick.Schedule = server.NewTickSchedule() // TODO tick schedule should be initialized in `newGameTick`
	startup.AddSystems(gameTick)

	// this is the master game context being passed around, containing pointers to everything
	gameCtx := &server.EngineCtx{ // TODO create a constructor for this
		GameId:                 gameId,
		IsLive:                 true,
		World:                  gameWorld,
		GameTick:               gameTick,
		TransactionsToSaveLock: sync.Mutex{},
	}

	// initialize a websocket streaming server for both incoming and outgoing requests
	gameTick.Setup(gameCtx, gameTick.Schedule) // TODO should just be a call on gameCtx
	return gameCtx
}
