package main

import (
	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/network"
	"github.com/curio-research/keystone-starter-kit/startup"
	ks "github.com/curio-research/keystone/server/startup"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// Initialize new game engine
	ctx := ks.NewGameEngine()

	ctx.SetTickRate(constants.TickRate)

	// ctx.SetPort(9000)
	// ctx.SetWebsocketPort(9001)

	ctx.SetEmitErrorHandler(&network.ProtoBasedErrorHandler{})
	ctx.SetEmitEventHandler(&network.ProtoBasedBroadcastHandler{})

	// Add systems
	startup.AddSystems(ctx)

	// Setup HTTP routes
	startup.SetupRoutes(ctx)

	// TODO: kevin: make this cleaner imo
	// Register tables schemas to world
	startup.RegisterTablesToWorld(ctx.World)

	// Initialize game map
	startup.InitWorld(ctx)

	// Start game server!
	ctx.Start()

}
