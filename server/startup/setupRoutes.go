package startup

import (
	"github.com/curio-research/keystone/state"
	"github.com/gin-gonic/gin"
)

// each route should be a call to a system
// setup routes that be be called

type EstablishPlayerRequest struct {
	PlayerID int `json:"playerID"`
}

type EstablishPlayerResponse struct {
	PlayerID int `json:"playerID"`
}

func SetupRoutes(engine *gin.Engine, w *state.GameWorld) {
	// Setup any http requests here
}
