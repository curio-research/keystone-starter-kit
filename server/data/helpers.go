package data

import (
	"github.com/curio-research/keystone/state"
	"math/rand"
)

func RandomAvailablePosition(w state.IWorld) (state.Pos, bool) {
	entities := Tile.Filter(w, TileSchema{
		Terrain: Ground,
	}, []string{"Terrain"})

	availablePositions := []state.Pos{}
	for entity := range entities {
		groundTile := Tile.Get(w, entity)
		if p := Player.Filter(w, PlayerSchema{Position: groundTile.Position}, []string{"Position"}); len(p) != 0 {
			continue
		}
		if a := Animal.Filter(w, AnimalSchema{Position: groundTile.Position}, []string{"Position"}); len(a) != 0 {
			continue
		}
		availablePositions = append(availablePositions, groundTile.Position)
	}

	if len(availablePositions) != 0 {
		return availablePositions[rand.Intn(len(availablePositions))], true
	}
	return state.Pos{}, false
}
