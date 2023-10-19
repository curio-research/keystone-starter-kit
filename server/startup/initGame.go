package startup

import (
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"math/rand"
)

func InitGame(w *state.GameWorld) {
	// initialize game storage tables
	RegisterTablesToWorld(w)

	// initialize game data into those tables
	InitWorld(w)
}

// register tables to the world
func RegisterTablesToWorld(w *state.GameWorld) {
	server.RegisterDefaultTables(w)

	var tableInterfacesToAdd []state.ITable
	for _, accessor := range data.TableSchemasToAccessors {
		tableInterfacesToAdd = append(tableInterfacesToAdd, accessor)
	}

	w.AddTables(tableInterfacesToAdd...)
}

func InitWorld(w *state.GameWorld) {
	data.Game.AddSpecific(w, constants.GameEntity, data.GameSchema{
		Weather: data.Sunny,
	})

	largeTileId := 1
	for i := 0; i < constants.WorldHeight; i++ {
		for j := 0; j < constants.WorldWidth; j++ {
			terrain := data.Terrain(weightedBoolean(constants.FreeTilesRatio))
			pos := state.Pos{
				X: j,
				Y: i,
			}

			data.Tile.AddSpecific(w, largeTileId, data.TileSchema{
				Position: pos,
				Terrain:  terrain,
			})

			if terrain == data.Ground {
				initializeNPC := weightedBoolean(constants.AnimalsToFreeTilesRatio)
				if initializeNPC {
					data.Animal.Add(w, data.AnimalSchema{
						Position: pos,
					})
				}
			}

			largeTileId++
		}
	}
}

func weightedBoolean(trueWeight float64) bool {
	if trueWeight > 1 || trueWeight < 0 {
		panic("boolean weight cannot be more than 1 or less than 0")
	}

	return rand.Float64() < trueWeight
}
