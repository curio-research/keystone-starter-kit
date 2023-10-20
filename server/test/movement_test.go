package test

import (
	"fmt"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/game/startup"
	"github.com/curio-research/keystone/game/systems"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/curio-research/keystone/utils"
	"github.com/go-playground/assert/v2"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var terrainKind = map[rune]data.Terrain{
	'.': data.Ground,
	'X': data.Obstacle,
}

var playerRegex, _ = regexp.Compile("[0-9]")

func TestMovement(t *testing.T) {
	ctx := worldWithPath(t, `
....X.......
....1X......
............
............
............
`)
	playerID := 1
	// doesn't move up => obstacle above
	server.QueueTxFromExternal(ctx, systems.MovementRequest{
		Direction: systems.Up,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 100)

	w := ctx.World
	player := getPlayers(w, playerID)[0]
	assert.Equal(t, state.Pos{X: 4, Y: 3}, player.Position)

	// move down
	server.QueueTxFromExternal(ctx, systems.MovementRequest{
		Direction: systems.Down,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 100)

	player = getPlayers(w, playerID)[0]
	assert.Equal(t, state.Pos{X: 4, Y: 2}, player.Position)

	// move right
	server.QueueTxFromExternal(ctx, systems.MovementRequest{
		Direction: systems.Right,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 100)

	player = getPlayers(w, playerID)[0]
	assert.Equal(t, state.Pos{X: 5, Y: 2}, player.Position)

	// can't move up
	server.QueueTxFromExternal(ctx, systems.MovementRequest{
		Direction: systems.Up,
		PlayerId:  1,
	}, "")
	utils.TickWorldForward(ctx, 100)

	player = getPlayers(w, playerID)[0]
	assert.Equal(t, state.Pos{X: 5, Y: 2}, player.Position)
}

func worldWithPath(t *testing.T, input string) *server.EngineCtx {
	w := state.NewWorld()
	startup.RegisterTablesToWorld(w)
	parseIntoWorld(t, w, input)

	ctx := NewTestEngine(w, systems.MovementSystem)

	return ctx
}

func parseIntoWorld(t *testing.T, w *state.GameWorld, input string) {
	rows := strings.Split(strings.TrimSpace(input), "\n")
	rowCount := len(rows)

	for y, row := range rows {
		for x, elem := range row {
			terrainKind, ok := terrainKind[elem]
			pos := state.Pos{X: x, Y: rowCount - y - 1}
			if ok {
				data.Tile.Add(w, data.TileSchema{
					Position: pos,
					Terrain:  terrainKind,
				})
			} else {
				data.Tile.Add(w, data.TileSchema{
					Position: pos,
					Terrain:  data.Ground,
				})

				symbol := string(elem)
				if symbol == "A" {
					data.Animal.Add(w, data.AnimalSchema{
						Position: pos,
					})
				} else if playerRegex.Match([]byte(symbol)) {
					p, _ := strconv.Atoi(symbol)
					data.Player.Add(w, data.PlayerSchema{
						Position:  pos,
						Resources: 10,
						PlayerID:  p,
					})
				} else {
					t.Fatal(fmt.Sprintf("character %s does not match any known symbol", symbol))
				}
			}
		}
	}
}

func getPlayers(w *state.GameWorld, playerID int) []data.PlayerSchema {
	playerEntity := data.Player.Filter(w, data.PlayerSchema{
		PlayerID: playerID,
	}, []string{"PlayerID"})

	if len(playerEntity) == 0 {
		return nil
	}

	var players []data.PlayerSchema
	for _, p := range playerEntity {
		players = append(players, data.Player.Get(w, p))
	}

	return players
}
