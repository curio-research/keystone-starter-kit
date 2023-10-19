package systems

import (
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/server"
)

type EstablishPlayerRequest struct {
	OwnerID int `json:"playerID"`
}

var EstablishPlayerSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[EstablishPlayerRequest]) {
	w := ctx.W
	req := ctx.Req

	player := data.Player.Get(w, req.OwnerID)
	if player.Id != 0 {
		ctx.EmitError("already created a player", []int{req.OwnerID})
		return
	}

	availablePos, ok := data.AnyAvailablePosition(w)
	if !ok {
		ctx.EmitError("this is awkward... there is no more space for a new player :(", []int{req.OwnerID})
		return
	}

	data.Player.Add(w, data.PlayerSchema{
		Position:  availablePos,
		Resources: 10,
		PlayerID:  req.OwnerID,
	})
})
