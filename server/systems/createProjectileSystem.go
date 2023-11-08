package systems

import (
	"github.com/curio-research/keystone-starter-kit/server/constants"
	"github.com/curio-research/keystone-starter-kit/server/data"
	"github.com/curio-research/keystone-starter-kit/server/helper"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

type CreateProjectileRequest struct {
	Direction helper.Direction `json:"direction"`
	PlayerId  int              `json:"playerId"`
}

var CreateProjectileSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[CreateProjectileRequest]) {
	req := ctx.Req.Data
	w := ctx.W

	direction := req.Direction
	initialPosition, found := locationOfPlayer(w, req.PlayerId)
	if !found {
		return
	}

	projectileID := data.Projectile.Add(w, data.ProjectileSchema{
		Position: initialPosition,
	})

	tickNumber := ctx.GameCtx.GameTick.TickNumber + constants.BulletSpeed

	server.QueueTxFromInternal[UpdateProjectileRequest](w, tickNumber, server.NewKeystoneTx(UpdateProjectileRequest{
		Direction:    direction,
		ProjectileID: projectileID,
		PlayerID:     req.PlayerId,
	}, nil), "")
}, server.VerifyECDSAPublicKeyAuth[CreateProjectileRequest]())

func locationOfPlayer(w state.IWorld, playerId int) (state.Pos, bool) {
	playerEntity := data.Player.Filter(w, data.PlayerSchema{PlayerId: playerId}, []string{"PlayerId"})
	if len(playerEntity) == 0 {
		return state.Pos{}, false
	}

	player := data.Player.Get(w, playerEntity[0])
	return player.Position, true
}
