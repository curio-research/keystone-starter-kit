package routes

import (
	"net/http"

	"github.com/curio-research/keystone/game/helper"
	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
)

// request
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
			if len(req.Tables) == 0 || helper.ContainsString(req.Tables, tableName) {

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
