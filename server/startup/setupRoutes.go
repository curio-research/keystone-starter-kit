package startup

import (
	"github.com/gin-gonic/gin"
)

// each route should be a call to a system
// setup routes that be be called
func SetupRoutes(engine *gin.Engine) {
	engine.GET("/")

}
