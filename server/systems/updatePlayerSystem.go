package systems

import (
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

// TODO: add movement system

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

type UpdatePlayerRequest struct {
	Direction `json:"direction"`
	PlayerId  int `json:"playerId"`
}

var UpdatePlayerSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[UpdatePlayerRequest]) {
	w := ctx.W
	req := ctx.Req

	playerRes := data.Player.Filter(w,
		data.PlayerSchema{
			PlayerId: req.PlayerId,
		}, []string{"PlayerId"})
	if len(playerRes) == 0 {
		ctx.EmitError("you have not created a player yet", []int{req.PlayerId})
		return
	}

	player := data.Player.Get(w, playerRes[0])
	targetPos := targetTile(player.Position, req.Direction)
	validTileToMove := validateTileToMoveTo(w, targetPos)

	if validTileToMove {
		player.Position = targetPos

		// add any resources the player gained at the position
		resource, found := resourceAtPosition(w, targetPos)
		if found {
			data.Resource.RemoveEntity(w, resource.Id)
			player.Resources += resource.Amount
		}

		data.Player.Set(w, player.Id, player)
	}
})

func resourceAtPosition(w state.IWorld, position state.Pos) (data.ResourceSchema, bool) {
	resource := data.Resource.Filter(w, data.ResourceSchema{
		Position: position,
	}, []string{"Position"})

	if len(resource) == 0 {
		return data.ResourceSchema{}, false
	}
	return data.Resource.Get(w, resource[0]), true
}
