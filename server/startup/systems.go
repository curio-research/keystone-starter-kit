package startup

import (
	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/systems"
	"github.com/curio-research/keystone/server"
)

// different systems can be triggered on different ticks
// some operations require the game to tick slower
// ex: health only regenerates every 20 seconds, but you attack every 5 seconds

func AddSystems(ctx *server.EngineCtx) {

	// ---------------------
	// External Systems
	// External systems require external input to run
	// ---------------------
	ctx.AddSystem(constants.TickRate, systems.CreatePlayerSystem)
	ctx.AddSystem(constants.TickRate, systems.UpdatePlayerSystem)
	ctx.AddSystem(constants.TickRate, systems.CreateProjectileSystem)

	// ---------------------
	// Internal Systems
	// internal systems run by themselves without external input
	// ---------------------
	ctx.AddSystem(constants.WeatherChangeIntervalMs, systems.WeatherSystem)
	ctx.AddSystem(constants.TickRate, systems.UpdateProjectileSystem)
	ctx.AddSystem(constants.AnimalCreationRate, systems.CreateAnimalSystem)
	ctx.AddSystem(constants.AnimalUpdateRate, systems.UpdateAnimalSystem)

	// TODO: remove in prod
	// tickSchedule.AddTickSystem(1_000, systems.TestSystem)

}
