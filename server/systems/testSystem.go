package systems

import (
	"math/rand"

	"github.com/curio-research/keystone-starter-kit/data"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

// * NOT used in production. Disable when necessary

func testSystem(ctx *server.TransactionCtx[any]) {
	// get random number
	randNumber1 := randomInt()
	randNumber2 := randomInt()

	data.Tile.Add(ctx.W, data.TileSchema{
		Position: state.Pos{
			X: randNumber1,
			Y: randNumber2,
		},
	})
}

func randomInt() int {
	return rand.Intn(100) + 1
}

var TestSystem = server.CreateGeneralSystem(testSystem)
