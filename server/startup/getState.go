package startup

import (
	"fmt"
	"net/http"
	"time"

	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
)

// download state

type DownloadStateRequest struct {
	Tables []string `json:"tables"`
}

// response
type GameStateResponse struct {
	Tick   int         `json:"tick"`
	Tables []TableData `json:"tables"`
}

type TableData struct {
	Name   string  `json:"name"`
	Values []Value `json:"values"`
}

type Value struct {
	Entity int `json:"entity"`
	Value  any `json:"value"`
}

func DownloadStateHandler(ctx *server.EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := server.DecodeRequestBody[DownloadStateRequest](c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		startTime := time.Now()

		gameStateResponse := &GameStateResponse{
			Tick:   ctx.GameTick.TickNumber,
			Tables: make([]TableData, 0),
		}

		// loop through all world tables
		for tableName, table := range ctx.World.Tables {

			// if table array is empty, return all tables
			if len(req.Tables) == 0 || ContainsString(req.Tables, tableName) {

				tableData := TableData{
					Name:   tableName,
					Values: make([]Value, 0),
				}

				for entity, value := range table.EntityToValue {
					tableData.Values = append(tableData.Values, Value{
						Entity: entity,
						Value:  value,
					})
				}
				gameStateResponse.Tables = append(gameStateResponse.Tables, tableData)

			}
		}

		fmt.Println("DownloadStateHandler took", time.Since(startTime))

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
