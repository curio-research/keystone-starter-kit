package systems

import (
	"fmt"
	"github.com/curio-research/keystone/game/constants"
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

type MovementRequest struct {
	Direction `json:"direction"`
	PlayerId  int `json:"playerId"`
}

func movementSystem(ctx *server.TransactionCtx[MovementRequest]) {
	w := ctx.W
	req := ctx.Req

	playerRes := data.Player.Filter(w,
		data.PlayerSchema{
			PlayerID: req.PlayerId,
		}, []string{"PlayerID"})
	if len(playerRes) == 0 {
		ctx.EmitError("you have not created a player yet", []int{req.PlayerId})
		return
	}

	player := data.Player.Get(w, playerRes[0])
	targetPos := targetTile(player.Position, req.Direction)
	validTileToMove := validateTile(w, targetPos)

	if validTileToMove {
		player.Position = targetPos
		data.Player.Set(w, player.Id, player)
		fmt.Println("player ", player.PlayerID, " new position ", targetPos)
	}
}

func targetTile(position state.Pos, direction Direction) state.Pos {
	switch direction {
	case Up:
		position.Y += 1
	case Down:
		position.Y -= 1
	case Left:
		position.X -= 1
	case Right:
		position.X += 1
	}

	return position
}

// TODO can we have a position entity?
func validateTile(w state.IWorld, pos state.Pos) bool {
	if (pos.X >= constants.WorldWidth || pos.X < 0) || (pos.Y >= constants.WorldHeight || pos.Y < 0) {
		return false
	}

	if players := data.Player.Filter(w, data.PlayerSchema{
		Position: pos,
	}, []string{"Position"}); len(players) != 0 {
		return false
	}

	if animals := data.Animal.Filter(w, data.AnimalSchema{
		Position: pos,
	}, []string{"Position"}); len(animals) != 0 {
		return false
	}

	return true
}

var MovementSystem = server.CreateSystemFromRequestHandler(movementSystem)
