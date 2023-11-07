package test

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/curio-research/keystone-starter-kit/constants"
	"github.com/curio-research/keystone-starter-kit/data"
	"github.com/curio-research/keystone/server"
	ks "github.com/curio-research/keystone/server/startup"
	"github.com/curio-research/keystone/state"
)

var terrainKind = map[rune]data.Terrain{
	'.': data.Ground,
	'X': data.Obstacle,
}

var playerRegex, _ = regexp.Compile("[0-9]")

func newTestEngine(gameWorld *state.GameWorld, systems ...server.TickSystemFunction) *server.EngineCtx {

	ctx := ks.NewGameEngine()
	for _, system := range systems {
		ctx.AddSystem(constants.TickRate, system)
	}

	return ctx
}

func worldWithPath(t *testing.T, input string, systems ...server.TickSystemFunction) *server.EngineCtx {
	w := state.NewWorld()

	ctx := newTestEngine(w, systems...)

	// Register tables
	server.RegisterDefaultTables(ctx.World)
	ctx.AddTables(data.SchemaMapping)

	parseIntoWorld(t, ctx.World, input)

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
						Position: pos,
						PlayerId: p,
					})
				} else {
					t.Fatal(fmt.Sprintf("character %s does not match any known symbol", symbol))
				}
			}
		}
	}
}

func getPlayer(w *state.GameWorld, playerID int) (data.PlayerSchema, bool) {
	playerEntity := data.Player.Filter(w, data.PlayerSchema{
		PlayerId: playerID,
	}, []string{"PlayerId"})

	if len(playerEntity) == 0 {
		return data.PlayerSchema{}, false
	}

	return data.Player.Get(w, playerEntity[0]), true
}

func testECDSAAuthHeader[T any](t *testing.T, req T) map[server.HeaderField]json.RawMessage {
	privateKey, err := crypto.GenerateKey()
	require.Nil(t, err)

	auth, err := server.NewECDSAPublicKeyAuth(privateKey, req)
	require.Nil(t, err)

	authBytes, err := json.Marshal(auth)
	require.Nil(t, err)

	return map[server.HeaderField]json.RawMessage{
		server.ECDSAPublicKeyAuthHeader: authBytes,
	}
}
