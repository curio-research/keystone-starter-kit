package startup

import (
	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone-starter-kit/helper"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

// Initialize world map and content
func InitWorld(ctx *server.EngineCtx) {
	data.Game.AddSpecific(ctx.World, constants.GameEntity, data.GameSchema{
		Weather: data.Sunny,
		GameId:  ctx.GameId,
	})

	largeTileId := 1
	for i := 0; i < constants.WorldHeight; i++ {
		for j := 0; j < constants.WorldWidth; j++ {
			terrain := data.Terrain(helper.WeightedBoolean(constants.FreeTilesRatio))
			pos := state.Pos{
				X: j,
				Y: i,
			}

			data.Tile.AddSpecific(ctx.World, largeTileId, data.TileSchema{
				Position: pos,
				Terrain:  terrain,
			})

			largeTileId++
		}
	}
}
