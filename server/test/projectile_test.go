package test

import (
	"testing"

	"github.com/curio-research/keystone-starter-kit/helper"
	"github.com/curio-research/keystone-starter-kit/systems"
	"github.com/curio-research/keystone/server"
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
	server.TickWorldForward(ctx, 10)

	_, found := getPlayer(w, 2)
	assert.True(t, found)

	server.TickWorldForward(ctx, 10)
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
	server.TickWorldForward(ctx, 40)

	_, found := getPlayer(w, 2)
	assert.True(t, found)
}
