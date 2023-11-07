package systems

import (
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone/state"
)

func PlayerWithID(w state.IWorld, playerId int) (data.PlayerSchema, bool) {
	playerEntity := data.Player.Filter(w, data.PlayerSchema{PlayerId: playerId}, []string{"PlayerId"})
	if len(playerEntity) == 0 {
		return data.PlayerSchema{}, false
	}

	player := data.Player.Get(w, playerEntity[0])
	return player, true
}
