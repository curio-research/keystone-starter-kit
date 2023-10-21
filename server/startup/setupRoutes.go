package startup

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"github.com/curio-research/keystone-starter-kit/systems"
	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
)

// TODO WS routes for error + updates

// each route should be a call to a system
// setup routes that should be called

func SetupRoutes(router *gin.Engine, gameCtx *server.EngineCtx) {

	// Setup any http requests here
	router.POST("/establishPlayer", func(ctx *gin.Context) {
		pushUpdateToQueue[systems.CreatePlayerRequest](ctx, gameCtx)
	})
	router.POST("/move", func(ctx *gin.Context) {
		pushUpdateToQueue[systems.UpdatePlayerRequest](ctx, gameCtx)
	})

	// get game state
	router.POST("/getState", DownloadStateHandler(gameCtx))
}

func pushUpdateToQueue[T any](ctx *gin.Context, engine *server.EngineCtx) {
	request := ctx.Request
	writer := ctx.Writer

	var t T
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
