package systems

import (
	"math/rand"

	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone-starter-kit/helper"
	"github.com/curio-research/keystone/server"
)

var UpdateAnimalSystem = server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
	w := ctx.W

	animalEntities := data.Animal.Entities(w)
	for _, animalEntity := range animalEntities {
		if helper.WeightedBoolean(constants.AnimalUpdateRatio) {
			animal := data.Animal.Get(w, animalEntity)
			availablePositions, found := helper.AvailablePositionsToTravel(w, animal.Position)
			if found {
				animal.Position = availablePositions[rand.Intn(len(availablePositions))]
				data.Animal.Set(w, animalEntity, animal)
			}
		}
	}
})
