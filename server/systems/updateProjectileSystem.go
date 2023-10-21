package systems

import (
	"encoding/json"
	"reflect"

	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone-starter-kit/helper"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

type UpdateProjectileRequest struct {
	NewPosition  state.Pos
	Direction    helper.Direction
	ProjectileID int
	PlayerID     int
}

var UpdateProjectileSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[UpdateProjectileRequest]) {
	w := ctx.W

	// check collisions
	collision := updateWorldForCollision(w, ctx.Req.NewPosition)
	if collision {
		// if collided, remove the projectile
		data.Projectile.RemoveEntity(w, ctx.Req.ProjectileID)

		// remove future jobs for the projectile
		projectileJobs := server.TransactionTable.Filter(w, server.TransactionSchema{
			Type: reflect.TypeOf(UpdateProjectileRequest{}).String(),
		}, []string{"Type"})

		for _, projectileJobEntity := range projectileJobs {
			projectileTx := server.TransactionTable.Get(w, projectileJobEntity)

			var futureProjectileReq UpdateProjectileRequest
			json.Unmarshal([]byte(projectileTx.Data), &futureProjectileReq)

			if futureProjectileReq.ProjectileID == ctx.Req.ProjectileID {
				server.TransactionTable.RemoveEntity(w, projectileJobEntity)
			}
		}
	} else {
		// update the position of the projectile
		projectile := data.Projectile.Get(w, ctx.Req.ProjectileID)
		projectile.Position = ctx.Req.NewPosition
		data.Projectile.Set(w, ctx.Req.ProjectileID, projectile)
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
