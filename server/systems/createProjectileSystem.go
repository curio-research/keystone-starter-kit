package systems

import (
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

type CreateProjectileRequest struct {
	Direction Direction `json:"direction"`
	PlayerId  int       `json:"playerId"`
}

var CreateProjectileSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[CreateProjectileRequest]) {
	req := ctx.Req
	w := ctx.W

	direction := req.Direction
	initialPosition, found := locationOfPlayer(w, req.PlayerId)
	if !found {
		return
	}

	projectileID := data.Projectile.Add(w, data.ProjectileSchema{
		Position: initialPosition,
	})
	position := targetTile(initialPosition, direction)
	tickNumber := ctx.GameCtx.GameTick.TickNumber + constants.BulletSpeed
	for withinBoardBoundary(position) {
		server.QueueTxFromInternal[UpdateProjectileRequest](w, tickNumber, UpdateProjectileRequest{
			NewPosition:  position,
			Direction:    direction,
			ProjectileID: projectileID,
			PlayerID:     req.PlayerId,
		}, "")
		tickNumber += constants.BulletSpeed
		position = targetTile(position, direction) // updates the position one step in the direction it was shot
	}

})

func locationOfPlayer(w state.IWorld, playerId int) (state.Pos, bool) {
	playerEntity := data.Player.Filter(w, data.PlayerSchema{PlayerId: playerId}, []string{"PlayerId"})
	if len(playerEntity) == 0 {
		return state.Pos{}, false
	}

	player := data.Player.Get(w, playerEntity[0])
	return player.Position, true
}
