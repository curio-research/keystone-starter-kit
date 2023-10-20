package startup

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"

	"github.com/curio-research/keystone/game/systems"
	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
)

// TODO WS routes for error + updates

// each route should be a call to a system
// setup routes that should be called

func SetupRoutes(router *gin.Engine, gameCtx *server.EngineCtx) {

	// Setup any http requests here
	router.POST("/establishPlayer", func(ctx *gin.Context) {
		pushUpdateToQueue[systems.CreatePlayerRequest](ctx, engine)
	})
	router.POST("/move", func(ctx *gin.Context) {
		pushUpdateToQueue[systems.UpdatePlayerRequest](ctx, engine)
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

// download state

type DownloadStateRequest struct {
	Tables []string `json:"tables"`
}

// response
type GameStateResponse struct {
	Tables []TableData `json:"tables"`
}

type TableData struct {
	Name   string             `json:"name"`
	Values []EntityValuePairs `json:"entityValuePairs"`
}

type EntityValuePairs struct {
	Entity int `json:"entity"`
	Value  any `json:"value"`
}

func DownloadStateHandler(ctx *server.EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := server.DecodeRequestBody[DownloadStateRequest](c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		gameStateResponse := &GameStateResponse{
			Tables: make([]TableData, 0),
		}

		// loop through all world tables
		for tableName, table := range ctx.World.Tables {

			// if table array is empty, return all tables
			if len(req.Tables) == 0 || ContainsString(req.Tables, tableName) {

				tableData := TableData{
					Name:   tableName,
					Values: make([]EntityValuePairs, 0),
				}

				for entity, value := range table.EntityToValue {
					tableData.Values = append(tableData.Values, EntityValuePairs{
						Entity: entity,
						Value:  value,
					})
				}
				gameStateResponse.Tables = append(gameStateResponse.Tables, tableData)

			}
		}

		c.JSON(http.StatusOK, gameStateResponse)
	}
}

func ContainsString(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}
