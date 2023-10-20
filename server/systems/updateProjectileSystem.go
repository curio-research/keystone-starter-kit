package systems

import (
	"encoding/json"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"reflect"
)

type UpdateProjectileRequest struct {
	NewPosition  state.Pos
	Direction    Direction
	ProjectileID int
	PlayerID     int
}

var UpdateProjectileSystem = server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
	w := ctx.W
	tickNumber := ctx.GameCtx.GameTick.TickNumber

	projectileJobType := reflect.TypeOf(UpdateProjectileRequest{}).String()
	projectileUpdatesAtTick := server.TransactionTable.Filter(w, server.TransactionSchema{
		Type:       projectileJobType,
		TickNumber: tickNumber,
	}, []string{"Type", "TickNumber"})

	for _, projectileUpdateEntity := range projectileUpdatesAtTick {
		projectileUpdateJob := server.TransactionTable.Get(w, projectileUpdateEntity)

		var projectileReq UpdateProjectileRequest
		json.Unmarshal([]byte(projectileUpdateJob.Data), &projectileReq)

		// check collisions
		collision := updateWorldForCollision(w, projectileReq.NewPosition)
		if collision {
			// if collided, remove the projectile
			data.Projectile.RemoveEntity(w, projectileReq.ProjectileID)

			// TODO have better query methods
			// remove future jobs for the projectile
			projectileJobs := server.TransactionTable.Filter(w, server.TransactionSchema{
				Type: projectileJobType,
			}, []string{"Type"})

			for _, projectileJobEntity := range projectileJobs {
				projectileTx := server.TransactionTable.Get(w, projectileJobEntity)

				var futureProjectileReq UpdateProjectileRequest
				json.Unmarshal([]byte(projectileTx.Data), &futureProjectileReq)

				if futureProjectileReq.ProjectileID == projectileReq.ProjectileID {
					server.TransactionTable.RemoveEntity(w, projectileJobEntity)
				}
			}
		} else {
			// update the position of the projectile
			projectile := data.Projectile.Get(w, projectileReq.ProjectileID)
			projectile.Position = projectileReq.NewPosition
			data.Projectile.Set(w, projectileReq.ProjectileID, projectile)
		}
	}

})

func updateWorldForCollision(w state.IWorld, position state.Pos) (collision bool) {
	players := playersAtLocation(w, position)
	if len(players) != 0 {
		collision = true
		for _, player := range players {
			data.Player.RemoveEntity(w, player)
		}
	}

	animals := animalsAtLocation(w, position)
	if len(animals) != 0 {
		collision = true
		for _, animal := range animals {
			data.Animal.RemoveEntity(w, animal)
		}
	}

	if isObstacleTile(w, position) {
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

func isObstacleTile(w state.IWorld, pos state.Pos) bool {
	ids := data.Tile.Filter(w, data.TileSchema{
		Position: pos,
		Terrain:  data.Obstacle,
	}, []string{"Position", "Terrain"})

	return len(ids) != 0
}
