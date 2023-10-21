package systems

import (
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/data"
	systems "github.com/curio-research/keystone/game/helpers"
	"github.com/curio-research/keystone/server"
	"math/rand"
)

var UpdateAnimalSystem = server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
	w := ctx.W

	animalEntities := data.Animal.Entities(w)
	for _, animalEntity := range animalEntities {
		if systems.WeightedBoolean(constants.AnimalUpdateRatio) {
			animal := data.Animal.Get(w, animalEntity)
			availablePositions, found := availablePositionsToTravel(w, animal.Position)
			if found {
				animal.Position = availablePositions[rand.Intn(len(availablePositions))]
				data.Animal.Set(w, animalEntity, animal)
			}
		}
	}
})
