package test

import (
	"testing"

	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone-starter-kit/helper"
	"github.com/curio-research/keystone-starter-kit/systems"

	"github.com/curio-research/keystone/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPickUpGold(t *testing.T) {
	ctx := worldWithPath(t, `
............
............
....3...A...
............
............
`, systems.CreatePlayerSystem, systems.MovePlayerSystem, systems.FireProjectileSystem, systems.UpdateProjectileSystem)
	w := ctx.World
	initialPlayerID := 3
	addedPlayerID := 2

	createPlayer := systems.CreatePlayerRequest{
		PlayerID:           addedPlayerID,
		EthBase64PublicKey: base64PublicKey(t, addedPlayerID),
	}
	server.QueueTxFromExternal(ctx, server.NewKeystoneTx(createPlayer, nil), "")
	server.TickWorldForward(ctx, 100)

	assert.Len(t, data.Animal.Entities(w), 1)

	player, found := systems.PlayerWithID(w, initialPlayerID)
	require.True(t, found)
	assert.Equal(t, 0, player.Resources)

	player2, found := systems.PlayerWithID(w, addedPlayerID)
	require.True(t, found)
	assert.Equal(t, 0, player.Resources)

	assert.NotEqual(t, player.Position, player2.Position)

	player2PublicKey := base64PublicKey(t, addedPlayerID)
	assert.Equal(t, player2PublicKey, player2.EthBase64PublicKey)

	req := systems.FireProjectileRequest{
		Direction: helper.Right,
		PlayerId:  initialPlayerID,
	}
	server.QueueTxFromExternal(ctx, server.NewKeystoneTx(req, testEthWalletAuthHeader(t, req, initialPlayerID)), "")
	server.TickWorldForward(ctx, 50) // create projectile + queue projectile update jobs

	assert.Len(t, data.Animal.Entities(w), 0)

	for i := 0; i < 4; i++ {
		req := systems.MovePlayerRequest{
			Direction: helper.Right,
			PlayerId:  initialPlayerID,
		}
		server.QueueTxFromExternal(ctx, server.NewKeystoneTx(req, testEthWalletAuthHeader(t, req, initialPlayerID)), "")
		server.TickWorldForward(ctx, 100)
	}

	player, found = systems.PlayerWithID(w, initialPlayerID)
	require.True(t, found)
	assert.Equal(t, constants.AnimalGold, player.Resources)

}
