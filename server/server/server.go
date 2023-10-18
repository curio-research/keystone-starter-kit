package server

import (
	"fmt"
	"log"
	"net/http"
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

func StartMainServer(mode string, websocketPort int, randSeedNumber int) (*gin.Engine, *server.EngineCtx, error) {
	// for debugging using profiler
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	color.HiYellow("")
	color.HiYellow("---- 🗝  Powered by Keystone 🗿 ----")
	fmt.Println()

	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.Use(server.CORSMiddleware())

	// initialize gameCtx (should be refactored to keystone)
	gameId := babble.NewBabbler().Babble()
	gameCtx := setupWorld(mode, gameId, randSeedNumber)

	// initialize in-game world
	startup.InitGame(gameCtx.World, randSeedNumber)

	color.HiWhite("Tick rate:         " + strconv.Itoa(gameCtx.GameTick.TickRateMs) + "ms")

	// add systems for game
	startup.AddSystems(gameCtx)

	// setup server routes
	startup.SetupRoutes(s)

	return s, gameCtx, nil
}

func setupWorld(mode string, gameId string, randSeedNumber int) *server.EngineCtx {
	gameWorld := state.NewWorld()
	gameTick := server.NewGameTick(constants.TickRate)

	// this is the master game context being passed around, containing pointers to everything
	gameCtx := &server.EngineCtx{ // TODO create a constructor for this
		GameId:                 gameId,
		IsLive:                 true,
		World:                  gameWorld,
		GameTick:               gameTick,
		TransactionsToSaveLock: sync.Mutex{},
		Mode:                   mode,
		RandSeed:               randSeedNumber,
	}

	// initialize a websocket streaming server for both incoming and outgoing requests
	gameTick.Setup(gameCtx, gameTick.Schedule) // TODO should just be a call on gameCtx
	return gameCtx
}
