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
`, systems.CreatePlayerSystem, systems.FireProjectileSystem, systems.UpdateProjectileSystem)

	w := ctx.World
	attackingPlayer := 1
	victimPlayer := 2
	req := systems.FireProjectileRequest{
		Direction: helper.Down,
		PlayerId:  attackingPlayer,
	}
	server.QueueTxFromExternal(ctx, server.NewKeystoneTx(req, testECDSAAuthHeader(t, req, attackingPlayer)), "")
	server.TickWorldForward(ctx, 10)

	_, found := systems.PlayerWithID(w, victimPlayer)
	assert.True(t, found)

	server.TickWorldForward(ctx, 10)
	_, found = systems.PlayerWithID(w, victimPlayer)
	assert.False(t, found)
}

func Test_Projectile_SavedByObstacle(t *testing.T) {
	ctx := worldWithPath(t, `
............
....1.......
............
....X.......
....2.......
`, systems.FireProjectileSystem, systems.UpdateProjectileSystem)

	w := ctx.World
	attackingPlayer := 1
	victimPlayer := 2
	req := systems.FireProjectileRequest{
		Direction: helper.Down,
		PlayerId:  1,
	}
	server.QueueTxFromExternal(ctx, server.NewKeystoneTx(req, testECDSAAuthHeader(t, req, attackingPlayer)), "")
	server.TickWorldForward(ctx, 40)

	_, found := systems.PlayerWithID(w, victimPlayer)
	assert.True(t, found)
}
