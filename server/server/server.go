package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/helper"
	"github.com/curio-research/keystone/game/network"
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

	// initialize the in-memory world
	gameWorld := state.NewWorld()
	helper.InitGame(gameWorld, randSeedNumber)

	gameTick := server.NewGameTick(constants.TickRate)
	// this is where you setup the tick schedules for your game
	startup.AddSystems(gameTick)

	// randomly generate gameID
	gameId := babble.NewBabbler().Babble()

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
	streamServer, err := server.NewStreamServer(s, gameCtx, network.SocketRequestRouter, websocketPort)
	if err != nil {
		return nil, nil, err
	}
	gameCtx.Stream = streamServer
	gameTick.Setup(gameCtx, gameTick.Schedule) // TODO should just be a call on gameCtx

	color.HiWhite("Tick rate:         " + strconv.Itoa(gameTick.TickRateMs) + "ms")

	// setup server routes
	network.SetupRoutes(s, gameCtx)

	return s, gameCtx, nil
}
