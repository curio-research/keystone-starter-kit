package test

import (
	"testing"

	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone-starter-kit/helper"
	"github.com/curio-research/keystone-starter-kit/systems"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/utils"
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
`, systems.UpdatePlayerSystem, systems.CreateProjectileSystem, systems.UpdateProjectileSystem)
	w := ctx.World
	playerID := 3

	assert.Len(t, data.Animal.Entities(w), 1)

	player, found := getPlayer(w, 3)
	require.True(t, found)
	require.Equal(t, 0, player.Resources)

	server.QueueTxFromExternal(ctx, systems.CreateProjectileRequest{
		Direction: helper.Right,
		PlayerId:  playerID,
	}, "")
	utils.TickWorldForward(ctx, 50) // create projectile + queue projectile update jobs

	assert.Len(t, data.Animal.Entities(w), 0)

	for i := 0; i < 4; i++ {
		server.QueueTxFromExternal(ctx, systems.UpdatePlayerRequest{
			Direction: helper.Right,
			PlayerId:  playerID,
		}, "")
		utils.TickWorldForward(ctx, 100)
	}

	player, found = getPlayer(w, 3)
	require.True(t, found)
	assert.Equal(t, constants.AnimalGold, player.Resources)

}
