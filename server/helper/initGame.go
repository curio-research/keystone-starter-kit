package helper

import (
	"math/rand"
	"time"

	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

var playerIds = []int{1000, 1001, 1002, 1003, 1004}

func InitGame(w *state.GameWorld, randSeedNumber int) {
	// initialize game storage tables
	RegisterTablesToWorld(w)

	// add random seed
	data.LocalRandomSeed.AddSpecific(w, constants.RandomnessEntity, data.LocalRandSeedSchema{
		RandValue: randSeedNumber,
	})

	// add game as an object
	data.Game.AddSpecific(w, constants.GameEntity, data.GameSchema{
		Weather: data.Sunny,
	})

	AddTilesToWorld(w)

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

func AddTilesToWorld(w state.IWorld) {
	largeTileId := 1
	for i := 0; i < constants.WorldHeight; i++ {
		for j := 0; j < constants.WorldWidth; j++ {

			data.Tile.AddSpecific(w, largeTileId, data.TileSchema{
				Position: state.Pos{
					X: j,
					Y: i,
				},
				Terrain: data.Grass,
			})

			largeTileId++

		}
	}
}

func ShuffleIntArr(arr []int) []int {
	// Create a new slice with the same length as the original array
	shuffled := make([]int, len(arr))
	copy(shuffled, arr)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}
