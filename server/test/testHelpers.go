package test

import (
	"fmt"
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/game/network"
	"github.com/curio-research/keystone/game/startup"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"testing"
)

var terrainKind = map[rune]data.Terrain{
	'.': data.Ground,
	'X': data.Obstacle,
}

var playerRegex, _ = regexp.Compile("[0-9]")

func newTestEngine(gameWorld *state.GameWorld, systems ...server.TickSystemFunction) *server.EngineCtx {
	gameTick := server.NewGameTick(constants.TickRate)

	// initiate an empty tick schedule
	tickSchedule := server.NewTickSchedule()
	gameTick.Schedule = tickSchedule
	for _, system := range systems {
		tickSchedule.AddTickSystem(constants.TickRate, system)
	}

	gameCtx := &server.EngineCtx{
		GameId:                 "prototype-game",
		IsLive:                 true,
		World:                  gameWorld,
		GameTick:               gameTick,
		TransactionsToSaveLock: sync.Mutex{},
		ShouldRecordError:      true,
		ErrorLog:               []server.ErrorLog{},
		Mode:                   "dev",
		SystemErrorHandler:     &network.ProtoBasedErrorHandler{},
		SystemBroadcastHandler: &network.ProtoBasedBroadcastHandler{},
	}

	return gameCtx
}

func worldWithPath(t *testing.T, input string, systems ...server.TickSystemFunction) *server.EngineCtx {
	w := state.NewWorld()
	startup.RegisterTablesToWorld(w)
	parseIntoWorld(t, w, input)

	ctx := newTestEngine(w, systems...)

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
						PlayerId:  p,
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
		PlayerId: playerID,
	}, []string{"PlayerId"})

	if len(playerEntity) == 0 {
		return nil
	}

	var players []data.PlayerSchema
	for _, p := range playerEntity {
		players = append(players, data.Player.Get(w, p))
	}

	return players
}
