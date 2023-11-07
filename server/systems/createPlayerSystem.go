package systems

import (
	"fmt"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone-starter-kit/helper"
	"github.com/curio-research/keystone/server"
)

type CreatePlayerRequest struct {
	PlayerID        int    `json:"playerID"`
	Base64PublicKey string `json:"publicKey"`
}

var CreatePlayerSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[CreatePlayerRequest]) {
	w := ctx.W
	req := ctx.Req.Data

	playerID := req.PlayerID
	_, found := PlayerWithID(w, req.PlayerID)
	if found {
		ctx.EmitError(fmt.Sprintf("already created a player with player ID %v", playerID), []int{playerID})
		return
	}

	publicKey := req.Base64PublicKey
	_, found = PlayerWithPublicKey(w, publicKey)
	if found {
		ctx.EmitError(fmt.Sprintf("already created a player with public key %s", publicKey), []int{playerID})
		return
	}

	availablePos, ok := helper.RandomAvailablePosition(w)
	if !ok {
		ctx.EmitError("this is awkward... there is no more space for a new player :(", []int{playerID})
		return
	}

	data.Player.Add(w, data.PlayerSchema{
		Position:        availablePos,
		PlayerId:        playerID,
		Base64PublicKey: req.Base64PublicKey,
	})
})
