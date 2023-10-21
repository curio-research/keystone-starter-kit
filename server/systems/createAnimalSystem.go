package systems

import (
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/logging"
	"github.com/curio-research/keystone/server"
)

// npc movement (lose health if the npc walks on you!)

var CreateAnimalSystem = server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
	w := ctx.W

	entities := data.Animal.Entities(w)
	if len(entities) >= constants.MaxNPCInWorld {
		if len(entities) > constants.MaxNPCInWorld {
			logging.Log().Errorf("too many entities in the world")
		}
		return
	}

	pos, ok := randomAvailablePosition(w)
	if !ok {
		return
	}

	data.Animal.Add(w, data.AnimalSchema{
		Position: pos,
	})
})
