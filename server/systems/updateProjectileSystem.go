package systems

import (
	"github.com/curio-research/keystone-starter-kit/server/constants"
	"github.com/curio-research/keystone-starter-kit/server/data"
	"github.com/curio-research/keystone-starter-kit/server/helper"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

type UpdateProjectileRequest struct {
	Direction    helper.Direction
	ProjectileID int
	PlayerID     int
}

var UpdateProjectileSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[UpdateProjectileRequest]) {
	w := ctx.W
	req := ctx.Req.Data

	// get projectile's position
	projectile := data.Projectile.Get(w, req.ProjectileID)

	nextPosition := helper.TargetTile(projectile.Position, req.Direction)

	// check collisions
	collision := updateWorldForCollision(w, nextPosition)
	if collision {
		// if collided, remove the projectile
		data.Projectile.RemoveEntity(w, req.ProjectileID)

	} else {
		// update the position of the projectile
		projectile.Position = nextPosition
		data.Projectile.Set(w, req.ProjectileID, projectile)

		// queue the next projectile update
		tickNumber := ctx.GameCtx.GameTick.TickNumber + constants.BulletSpeed

		server.QueueTxFromInternal[UpdateProjectileRequest](w, tickNumber, server.NewKeystoneTx(UpdateProjectileRequest{
			Direction:    req.Direction,
			ProjectileID: req.ProjectileID,
			PlayerID:     req.PlayerID,
		}, nil), "")
	}

})

func updateWorldForCollision(w state.IWorld, position state.Pos) (collision bool) {

	// check if position is within world
	if !helper.WithinBoardBoundary(position) {
		return true
	}

	players := playersAtLocation(w, position)
	if len(players) != 0 {
		collision = true
		for _, player := range players {
			data.Player.RemoveEntity(w, player)
		}
		data.Resource.Add(w, data.ResourceSchema{
			Position: position,
			Amount:   constants.PlayerGold,
		})
	}

	animals := animalsAtLocation(w, position)
	if len(animals) != 0 {
		collision = true
		for _, animal := range animals {
			data.Animal.RemoveEntity(w, animal)
		}
		data.Resource.Add(w, data.ResourceSchema{
			Position: position,
			Amount:   constants.AnimalGold,
		})
	}

	if helper.IsObstacleTile(w, position) {
		collision = true
	}

	return collision
}

// would we ever have to handle a case where the bullet flies over more than one tile at once?
func playersAtLocation(w state.IWorld, pos state.Pos) []int {
	return data.Player.Filter(w, data.PlayerSchema{
		Position: pos,
	}, []string{"Position"})
}

func animalsAtLocation(w state.IWorld, pos state.Pos) []int {
	return data.Animal.Filter(w, data.AnimalSchema{
		Position: pos,
	}, []string{"Position"})
}
