package systems

import (
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

type CreateProjectileRequest struct {
	Position  state.Pos `json:"position"`
	Direction Direction `json:"direction"`
	PlayerID  int       `json:"playerID"`
}

var CreateProjectileSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[CreateProjectileRequest]) {
	req := ctx.Req
	w := ctx.W

	direction := req.Direction
	position := req.Position
	if !validateTile(w, position) {
		return
	}

	var projectileID int
	tickNumber := ctx.GameCtx.GameTick.TickNumber + constants.BulletSpeed
	for validateTile(w, position) {
		if position == req.Position {
			position = targetTile(position, direction)
			projectileID = data.Projectile.Add(w, data.ProjectileSchema{ // starts one tile in the direction it was shot
				Position: position,
			})
		} else {
			server.QueueTxFromInternal[UpdateProjectileRequest](w, tickNumber, UpdateProjectileRequest{
				NewPosition:  position,
				Direction:    direction,
				ProjectileID: projectileID,
				PlayerID:     req.PlayerID,
			}, "")
			tickNumber += constants.BulletSpeed
		}

		position = targetTile(position, direction) // updates the position one step in the direction it was shot
	}
})
