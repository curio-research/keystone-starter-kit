package startup

import (
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/systems"
	"github.com/curio-research/keystone/server"
)

// different events can happen on different ticks
// some operations require the game to tick slower
// ex: health only regenerates every 20 seconds, but you attack every 5 seconds

func AddSystems(gameTick *server.GameTick) {
	// different events can happen on different ticks
	// some operations require the game to tick slower
	// ex: health only regenerates every 20 seconds, but you attack every 5 seconds

	tickSchedule := gameTick.Schedule

	// ---------------------
	// External Systems
	// External systems require external input to run
	// ---------------------
	tickSchedule.AddTickSystem(constants.TickRate, systems.EstablishPlayerSystem)
	tickSchedule.AddTickSystem(constants.TickRate, systems.MovementSystem)

	// ---------------------
	// Internal Systems
	// internal systems run by themselves without external input
	// ---------------------
	tickSchedule.AddTickSystem(constants.WeatherChangeIntervalMs, systems.WeatherSystem)
	tickSchedule.AddTickSystem(1_000, systems.TestSystem)

	gameTick.Schedule = tickSchedule
}
