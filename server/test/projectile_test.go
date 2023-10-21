package test

import (
	"testing"

	"github.com/curio-research/keystone-starter-kit/helper"
	"github.com/curio-research/keystone-starter-kit/systems"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/utils"
	"github.com/stretchr/testify/assert"
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
		Direction: helper.Down,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 10)

	_, found := getPlayer(w, 2)
	assert.True(t, found)

	utils.TickWorldForward(ctx, 30)
	_, found = getPlayer(w, 2)
	assert.False(t, found)
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
		Direction: helper.Down,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 40)

	_, found := getPlayer(w, 2)
	assert.True(t, found)
}