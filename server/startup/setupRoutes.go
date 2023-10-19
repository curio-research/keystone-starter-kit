package startup

import (
	"encoding/json"
	"fmt"
	"github.com/curio-research/keystone/game/systems"
	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
	"io"
	"reflect"
)

// each route should be a call to a system
// setup routes that should be called

func SetupRoutes(router *gin.Engine, engine *server.EngineCtx) {
	// Setup any http requests here
	router.POST("/establishPlayer", func(ctx *gin.Context) {
		pushUpdateToQueue[systems.EstablishPlayerRequest](ctx, engine)
	})
	router.POST("/move", func(ctx *gin.Context) {
		pushUpdateToQueue[systems.MovementRequest](ctx, engine)
	})
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
		writer.Write([]byte(fmt.Sprintf("error unmarshalling request to type of %v", reflect.TypeOf(t).String())))
		return
	}

	err = server.QueueTxFromExternal[T](engine, t, "")
	if err != nil {
		writer.Write([]byte("error queuing transaction: " + err.Error()))
	}
}
