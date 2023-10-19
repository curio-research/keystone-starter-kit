package systems

import (
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/server"
)

type EstablishPlayerRequest struct {
	PlayerID int `json:"playerId"`
}

var EstablishPlayerSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[EstablishPlayerRequest]) {
	w := ctx.W
	req := ctx.Req

	player := data.Player.Filter(w,
		data.PlayerSchema{
			PlayerID: req.PlayerID,
		}, []string{"PlayerID"})
	if len(player) != 0 {
		ctx.EmitError("already created a player", []int{req.PlayerID})
		return
	}

	availablePos, ok := data.RandomAvailablePosition(w)
	if !ok {
		ctx.EmitError("this is awkward... there is no more space for a new player :(", []int{req.PlayerID})
		return
	}

	data.Player.Add(w, data.PlayerSchema{
		Position:  availablePos,
		Resources: 10,
		PlayerID:  req.PlayerID,
	})
})
