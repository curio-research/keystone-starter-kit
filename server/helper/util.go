package helper

import (
	"math/rand"

	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone/state"
)

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

func WeightedBoolean(trueWeight float64) bool {
	if trueWeight > 1 || trueWeight < 0 {
		panic("boolean weight cannot be more than 1 or less than 0")
	}

	return rand.Float64() < trueWeight
}

func AvailablePositionsToTravel(w state.IWorld, pos state.Pos) ([]state.Pos, bool) {
	var availablePositions []state.Pos
	for _, direction := range []Direction{Up, Down, Left, Right} {
		targetPos := TargetTile(pos, direction)
		if ValidateTileToMoveTo(w, targetPos) {
			availablePositions = append(availablePositions, targetPos)
		}
	}

	if len(availablePositions) == 0 {
		return nil, false
	}
	return availablePositions, true
}

func TargetTile(position state.Pos, direction Direction) state.Pos {
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
func ValidateTileToMoveTo(w state.IWorld, pos state.Pos) bool {
	if !WithinBoardBoundary(pos) {
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

	return !IsObstacleTile(w, pos)
}

func RandomAvailablePosition(w state.IWorld) (state.Pos, bool) {
	entities := data.Tile.Filter(w, data.TileSchema{
		Terrain: data.Ground,
	}, []string{"Terrain"})

	availablePositions := []state.Pos{}
	for _, entity := range entities {
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

// if position is within world boundaries
func WithinBoardBoundary(pos state.Pos) bool {
	if (pos.X >= constants.WorldWidth || pos.X < 0) || (pos.Y >= constants.WorldHeight || pos.Y < 0) {
		return false
	}
	return true
}

func WithinBoardBoundaryWithExtraLayer(pos state.Pos) bool {
	if (pos.X > constants.WorldWidth || pos.X < -1) || (pos.Y > constants.WorldHeight || pos.Y < -1) {
		return false
	}
	return true
}

func IsObstacleTile(w state.IWorld, pos state.Pos) bool {
	ids := data.Tile.Filter(w, data.TileSchema{
		Position: pos,
		Terrain:  data.Obstacle,
	}, []string{"Position", "Terrain"})

	return len(ids) != 0
}
