package systems

import (
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/tables"
	"github.com/curio-research/keystone/server"
)

// weather automatically changes
func weatherSystem(ctx *server.TransactionCtx[any]) {
	game := tables.Game.Get(ctx.W, constants.GameEntity)

	// change weather
	if game.Weather == tables.Sunny {
		game.Weather = tables.Windy
	} else {
		game.Weather = tables.Sunny
	}

	tables.Game.Set(ctx.W, constants.GameEntity, game)
}

var WeatherSystem = server.CreateGeneralSystem(weatherSystem)
