package startup

import (
	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/systems"
	"github.com/curio-research/keystone/server"
)

// Different systems can be triggered on different ticks
// Some operations require the game to tick slower
// Ex: health only regenerates every 20 seconds, but you attack every 5 seconds

func AddSystems(ctx *server.EngineCtx) {

	// ---------------------
	// External Systems - require external transactions to trigger
	// ---------------------
	ctx.AddSystem(constants.TickRate, systems.CreatePlayerSystem)
	ctx.AddSystem(constants.TickRate, systems.UpdatePlayerSystem)
	ctx.AddSystem(constants.TickRate, systems.CreateProjectileSystem)

	// ---------------------
	// Internal Systems - run by themselves (like a cron job)
	// ---------------------
	ctx.AddSystem(constants.WeatherChangeIntervalMs, systems.WeatherSystem)
	ctx.AddSystem(constants.TickRate, systems.UpdateProjectileSystem)
	ctx.AddSystem(constants.AnimalCreationRate, systems.CreateAnimalSystem)
	ctx.AddSystem(constants.AnimalUpdateRate, systems.UpdateAnimalSystem)

}
