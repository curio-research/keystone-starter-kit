package systems

import (
	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone/server"
)

// weather automatically changes
func weatherSystem(ctx *server.TransactionCtx[any]) {
	game := data.Game.Get(ctx.W, constants.GameEntity)

	// change weather
	if game.Weather == data.Sunny {
		game.Weather = data.Windy
	} else {
		game.Weather = data.Sunny
	}

	data.Game.Set(ctx.W, constants.GameEntity, game)
}

var WeatherSystem = server.CreateGeneralSystem(weatherSystem)
