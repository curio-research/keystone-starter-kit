package test

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/curio-research/keystone-starter-kit/systems"
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

var keyPair1, _ = crypto.GenerateKey()
var keyPair2, _ = crypto.GenerateKey()
var keyPair3, _ = crypto.GenerateKey()

var playerIDToPrivateKey = map[int]*ecdsa.PrivateKey{
	1: keyPair1,
	2: keyPair2,
	3: keyPair3,
}

var playerRegex, _ = regexp.Compile("[0-9]")

func newTestEngine(systems ...server.TickSystemFunction) *server.EngineCtx {
	ctx := ks.NewGameEngine()
	for _, system := range systems {
		ctx.AddSystem(constants.TickRate, system)
	}

	return ctx
}

func worldWithPath(t *testing.T, input string, systems ...server.TickSystemFunction) *server.EngineCtx {
	ctx := newTestEngine(systems...)

	// Register tables
	server.RegisterDefaultTables(ctx.World)
	ctx.AddTables(data.TableSchemasToAccessors)

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
						Position:        pos,
						PlayerId:        p,
						Base64PublicKey: base64PublicKey(t, p),
					})
				} else {
					t.Fatal(fmt.Sprintf("character %s does not match any known symbol", symbol))
				}
			}
		}
	}
}

func testECDSAAuthHeader[T any](t *testing.T, req T, playerID int) map[server.HeaderField]json.RawMessage {
	privateKey, ok := playerIDToPrivateKey[playerID]
	require.Truef(t, ok, "playerID must be in `playerIDToPrivateKey` map")

	auth, err := server.NewECDSAPublicKeyAuth(privateKey, req)
	require.Nil(t, err)

	authBytes, err := json.Marshal(auth)
	require.Nil(t, err)

	return map[server.HeaderField]json.RawMessage{
		server.ECDSAPublicKeyAuthHeader: authBytes,
		systems.PlayerIDHeader:          json.RawMessage(strconv.Itoa(playerID)),
	}
}

func base64PublicKey(t *testing.T, playerID int) string {
	privateKey, ok := playerIDToPrivateKey[playerID]
	require.Truef(t, ok, "playerID must be in `playerIDToPrivateKey` map")

	auth, err := server.NewECDSAPublicKeyAuth(privateKey, struct{}{})
	require.Nil(t, err)

	return auth.Base64PublicKey
}
