package routes

import (
	"fmt"
	"net/http"

	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
)

type ToggleStateRequest struct {
	IsLive bool `json:"isLive"`
}

// toggle if the game is paused or not
func ToggleLiveStateHandler(ctx *server.EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := server.DecodeRequestBody[ToggleStateRequest](c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		ctx.IsLive = req.IsLive

		fmt.Println("Game state toggled to: ", req.IsLive)

		c.JSON(http.StatusOK, "success")
	}
}
