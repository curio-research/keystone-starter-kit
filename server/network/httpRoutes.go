package network

import (
	"github.com/curio-research/keystone/game/network/routes"
	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
)

// sets up routes that allow HTTP requests to queue transactions into the world, view state, etc
func SetupRoutes(s *gin.Engine, ctx *server.EngineCtx) {

	s.POST("/getState", routes.DownloadStateHandler(ctx)) // fetches entire game state

	s.POST("/toggleLiveState", routes.ToggleLiveStateHandler(ctx))

	s.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!!!",
		})
	})
}
