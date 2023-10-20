package systems

import (
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone/server"
)

type CreatePlayerRequest struct {
	PlayerID int `json:"playerId"`
}

var CreatePlayerSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[CreatePlayerRequest]) {
	w := ctx.W
	req := ctx.Req

	player := data.Player.Filter(w,
		data.PlayerSchema{
			PlayerId: req.PlayerID,
		}, []string{"PlayerId"})
	if len(player) != 0 {
		ctx.EmitError("already created a player", []int{req.PlayerID})
		return
	}

	availablePos, ok := randomAvailablePosition(w)
	if !ok {
		ctx.EmitError("this is awkward... there is no more space for a new player :(", []int{req.PlayerID})
		return
	}

	data.Player.Add(w, data.PlayerSchema{
		Position: availablePos,
		PlayerId: req.PlayerID,
	})
})
