package data

import "github.com/curio-research/keystone/state"

func AnyAvailablePosition(w state.IWorld) (state.Pos, bool) {
	entities := Tile.Filter(w, TileSchema{
		Terrain: Ground,
	}, []string{"Terrain"})

	for entity := range entities {
		groundTile := Tile.Get(w, entity)
		if p := Player.Filter(w, PlayerSchema{Position: groundTile.Position}, []string{"Position"}); len(p) != 0 {
			continue
		}
		if a := Animal.Filter(w, AnimalSchema{Position: groundTile.Position}, []string{"Position"}); len(a) != 0 {
			continue
		}
		return groundTile.Position, true
	}

	return state.Pos{}, false
}
