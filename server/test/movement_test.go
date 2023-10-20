package test

import (
	"github.com/curio-research/keystone/game/systems"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/curio-research/keystone/utils"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestMovement(t *testing.T) {
	ctx := worldWithPath(t, `
....X.......
...21A......
............
............
............
`, systems.UpdatePlayerSystem)

	w := ctx.World
	playerID := 1

	// doesn't move up/right/left because of obstacles
	for _, direction := range []systems.Direction{systems.Up, systems.Right, systems.Left} {
		server.QueueTxFromExternal(ctx, systems.PlayerMovementRequest{
			Direction: direction,
			PlayerId:  1,
		}, "")
		utils.TickWorldForward(ctx, 100)

		player := getPlayers(w, playerID)[0]
		assert.Equal(t, state.Pos{X: 4, Y: 3}, player.Position)
	}

	// move down
	server.QueueTxFromExternal(ctx, systems.PlayerMovementRequest{
		Direction: systems.Down,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 100)

	player := getPlayers(w, playerID)[0]
	assert.Equal(t, state.Pos{X: 4, Y: 2}, player.Position)

	// move right
	server.QueueTxFromExternal(ctx, systems.PlayerMovementRequest{
		Direction: systems.Right,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 100)

	player = getPlayers(w, playerID)[0]
	assert.Equal(t, state.Pos{X: 5, Y: 2}, player.Position)

	// can't move up
	server.QueueTxFromExternal(ctx, systems.PlayerMovementRequest{
		Direction: systems.Up,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 100)

	player = getPlayers(w, playerID)[0]
	assert.Equal(t, state.Pos{X: 5, Y: 2}, player.Position)
}
