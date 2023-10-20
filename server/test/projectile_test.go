package test

import (
	"github.com/curio-research/keystone/game/systems"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Projectile(t *testing.T) {
	ctx := worldWithPath(t, `
............
....1.......
............
............
....2.......
`, systems.CreateProjectileSystem, systems.UpdateProjectileSystem)

	w := ctx.World
	server.QueueTxFromExternal(ctx, systems.CreateProjectileRequest{
		Direction: systems.Down,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 10)

	player2 := getPlayers(w, 2)
	require.Len(t, player2, 1)

	utils.TickWorldForward(ctx, 30)
	player2 = getPlayers(w, 2)
	assert.Len(t, player2, 0)
}

func Test_Projectile_SavedByObstacle(t *testing.T) {
	ctx := worldWithPath(t, `
............
....1.......
............
....X.......
....2.......
`, systems.CreateProjectileSystem, systems.UpdateProjectileSystem)

	w := ctx.World
	server.QueueTxFromExternal(ctx, systems.CreateProjectileRequest{
		Direction: systems.Down,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 40)

	player2 := getPlayers(w, 2)
	assert.Len(t, player2, 1)
}
