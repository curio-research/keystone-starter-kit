package helper

import (
	"math/rand"
	"time"

	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/tables"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

var playerIds = []int{1000, 1001, 1002, 1003, 1004}

func InitGame(w *state.GameWorld, randSeedNumber int) {
	// initialize game storage tables
	RegisterTablesToWorld(w)

	// add random seed
	tables.LocalRandomSeed.AddSpecific(w, constants.RandomnessEntity, tables.LocalRandSeedSchema{
		RandValue: randSeedNumber,
	})

	// add game as an object
	tables.Game.AddSpecific(w, constants.GameEntity, tables.GameSchema{
		Weather: tables.Sunny,
	})

	AddTilesToWorld(w)

}

// register tables to the world
func RegisterTablesToWorld(w *state.GameWorld) {
	server.RegisterDefaultTables(w)

	var tableInterfacesToAdd []state.ITable
	for _, accessor := range tables.TableSchemasToAccessors {
		tableInterfacesToAdd = append(tableInterfacesToAdd, accessor)
	}

	w.AddTables(tableInterfacesToAdd...)
}

func AddTilesToWorld(w state.IWorld) {
	largeTileId := 1
	for i := 0; i < constants.WorldHeight; i++ {
		for j := 0; j < constants.WorldWidth; j++ {

			tables.Tile.AddSpecific(w, largeTileId, tables.TileSchema{
				Position: state.Pos{
					X: j,
					Y: i,
				},
				Terrain: tables.Grass,
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
