package test

import (
	"testing"

	"github.com/curio-research/keystone-starter-kit/server/helper"
	"github.com/curio-research/keystone-starter-kit/server/systems"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
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
	for _, direction := range []helper.Direction{helper.Up, helper.Right, helper.Left} {
		req := systems.UpdatePlayerRequest{
			Direction: direction,
			PlayerId:  1,
		}
		server.QueueTxFromExternal(ctx, server.NewKeystoneTx(req, testECDSAAuthHeader(t, req)), "")
		server.TickWorldForward(ctx, 100)

		player, found := getPlayer(w, playerID)
		require.True(t, found)
		assert.Equal(t, state.Pos{X: 4, Y: 3}, player.Position)
	}

	// move down
	req := systems.UpdatePlayerRequest{
		Direction: helper.Down,
		PlayerId:  1,
	}
	server.QueueTxFromExternal(ctx, server.NewKeystoneTx(req, testECDSAAuthHeader(t, req)), "")
	server.TickWorldForward(ctx, 100)

	player, found := getPlayer(w, playerID)
	require.True(t, found)
	assert.Equal(t, state.Pos{X: 4, Y: 2}, player.Position)

	// move right
	req = systems.UpdatePlayerRequest{
		Direction: helper.Right,
		PlayerId:  1,
	}
	server.QueueTxFromExternal(ctx, server.NewKeystoneTx(req, testECDSAAuthHeader(t, req)), "")
	server.TickWorldForward(ctx, 100)

	player, found = getPlayer(w, playerID)
	require.True(t, found)
	assert.Equal(t, state.Pos{X: 5, Y: 2}, player.Position)

	// can't move up
	req = systems.UpdatePlayerRequest{
		Direction: helper.Up,
		PlayerId:  1,
	}
	server.QueueTxFromExternal(ctx, server.NewKeystoneTx(req, testECDSAAuthHeader(t, req)), "")
	server.TickWorldForward(ctx, 100)

	player, found = getPlayer(w, playerID)
	require.True(t, found)
	assert.Equal(t, state.Pos{X: 5, Y: 2}, player.Position)
}
