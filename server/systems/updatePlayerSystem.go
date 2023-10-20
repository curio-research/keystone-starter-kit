package systems

import (
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/server"
)

// TODO: add movement system

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

type PlayerMovementRequest struct {
	Direction `json:"direction"`
	PlayerId  int `json:"playerId"`
}

var UpdatePlayerSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[PlayerMovementRequest]) {
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
		data.Player.Set(w, player.Id, player)
	}
})
