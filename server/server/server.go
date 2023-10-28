package server

import (
	"fmt"
	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/network"
	"github.com/curio-research/keystone-starter-kit/startup"
	"github.com/curio-research/keystone/server"
	ks "github.com/curio-research/keystone/startup"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/tjarratt/babble"
	"strconv"
)

func MainServer(websocketPort int) (*gin.Engine, error) {
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
	startup.InitGame(gameCtx)
	color.HiWhite("Tick rate:         " + strconv.Itoa(gameCtx.GameTick.TickRateMs) + "ms")

	// setting up websocket requests to receive state updates (create router to handle getting WS requests in game)
	err := ks.RegisterWSRoutes(gameCtx, s, nil, websocketPort)
	if err != nil {
		return nil, err
	}

	// setup HTTP routes
	startup.SetupRoutes(s, gameCtx)

	return s, nil
}

// TODO this function should be in keystone
func setupWorld(gameId string) *server.EngineCtx {
	ctx := ks.NewGameEngine(gameId, constants.TickRate, 0)
	ks.RegisterErrorHandler(ctx, &network.ProtoBasedErrorHandler{})
	ks.RegisterBroadcastHandler(ctx, &network.ProtoBasedBroadcastHandler{})

	// add systems for game
	startup.AddSystems(ctx.GameTick)

	// initialize a websocket streaming server for both incoming and outgoing requests
	ks.Start(ctx)
	return ctx
}
