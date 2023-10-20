package systems

import (
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/state"
	"math/rand"
)

func availablePositionsToTravel(w state.IWorld, pos state.Pos) ([]state.Pos, bool) {
	var availablePositions []state.Pos
	for _, direction := range []Direction{Up, Down, Left, Right} {
		targetPos := targetTile(pos, direction)
		if validateTileToMoveTo(w, targetPos) {
			availablePositions = append(availablePositions, targetPos)
		}
	}

	if len(availablePositions) == 0 {
		return nil, false
	}
	return availablePositions, true
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

// getting position could be optimized by creating a separate table for position
func validateTileToMoveTo(w state.IWorld, pos state.Pos) bool {
	if !withinBoardBoundary(pos) {
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

	return !isObstacleTile(w, pos)
}

func randomAvailablePosition(w state.IWorld) (state.Pos, bool) {
	entities := data.Tile.Filter(w, data.TileSchema{
		Terrain: data.Ground,
	}, []string{"Terrain"})

	availablePositions := []state.Pos{}
	for entity := range entities {
		groundTile := data.Tile.Get(w, entity)
		if p := data.Player.Filter(w, data.PlayerSchema{Position: groundTile.Position}, []string{"Position"}); len(p) != 0 {
			continue
		}
		if a := data.Animal.Filter(w, data.AnimalSchema{Position: groundTile.Position}, []string{"Position"}); len(a) != 0 {
			continue
		}
		availablePositions = append(availablePositions, groundTile.Position)
	}

	if len(availablePositions) != 0 {
		return availablePositions[rand.Intn(len(availablePositions))], true
	}
	return state.Pos{}, false
}

func withinBoardBoundary(pos state.Pos) bool {
	if (pos.X >= constants.WorldWidth || pos.X < 0) || (pos.Y >= constants.WorldHeight || pos.Y < 0) {
		return false
	}
	return true
}
