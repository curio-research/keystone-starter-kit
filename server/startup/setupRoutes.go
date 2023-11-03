package startup

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"github.com/curio-research/keystone-starter-kit/systems"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/server/routes"
	"github.com/gin-gonic/gin"
)

// TODO WS routes for error + updates

// each route should be a call to a system
// setup routes that should be called

func SetupRoutes(ctx *server.EngineCtx) {

	// Setup any http requests here
	ctx.GinHttpEngine.POST("/createPlayer", func(ginCtx *gin.Context) {
		pushUpdateToQueue[systems.CreatePlayerRequest](ginCtx, ctx)
	})
	ctx.GinHttpEngine.POST("/move", func(ginCtx *gin.Context) {
		pushUpdateToQueue[systems.UpdatePlayerRequest](ginCtx, ctx)
	})
	ctx.GinHttpEngine.POST("/fire", func(ginCtx *gin.Context) {
		pushUpdateToQueue[systems.CreateProjectileRequest](ginCtx, ctx)
	})

	// get game state
	ctx.GinHttpEngine.POST("/getState", routes.GetStateRouteHandler(ctx))
}

func pushUpdateToQueue[T any](ctx *gin.Context, engine *server.EngineCtx) {
	request := ctx.Request
	writer := ctx.Writer

	var t server.KeystoneTx[T]
	b, err := io.ReadAll(request.Body)
	if err != nil {
		writer.Write([]byte("error reading request: " + err.Error()))
		return
	}

	err = json.Unmarshal(b, &t)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf("error unmarshalling request to type of %s: %s", reflect.TypeOf(t).String(), err.Error())))
		return
	}

	err = server.QueueTxFromExternal[T](engine, t, "")
	if err != nil {
		writer.Write([]byte("error queuing transaction: " + err.Error()))
	}
}
