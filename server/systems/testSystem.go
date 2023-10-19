package systems

import (
	"math/rand"
	"time"

	"github.com/curio-research/keystone/game/tables"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

// * NOT used in production. Disable when necessary

func testSystem(ctx *server.TransactionCtx[any]) {

	// get random number
	randNumber1 := randomInt()
	randNumber2 := randomInt()

	tables.Tile.Add(ctx.W, tables.TileSchema{
		Position: state.Pos{
			X: randNumber1,
			Y: randNumber2,
		},
	})
}

func randomInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}

var TestSystem = server.CreateGeneralSystem(testSystem)
