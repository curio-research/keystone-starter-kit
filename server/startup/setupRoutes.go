package startup

import (
	"encoding/json"
	"github.com/curio-research/keystone/game/data"
	"github.com/curio-research/keystone/state"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
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
	engine.POST("/establishPlayer", func(context *gin.Context) {
		request := context.Request
		writer := context.Writer

		b, err := io.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
			return
		}

		var req EstablishPlayerRequest
		err = json.Unmarshal(b, &req)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("error unmarshalling establish player request: " + err.Error()))
			return
		}

		availablePos, ok := data.AnyAvailablePosition(w)
		if !ok {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("this is awkward... there is no more space for a new player :("))
			return
		}

		data.Player.Add(w, data.PlayerSchema{
			Position:  availablePos,
			Resources: 10,
			PlayerID:  req.PlayerID,
		})

		res := EstablishPlayerResponse{PlayerID: req.PlayerID}
		b, _ = json.Marshal(res)
		writer.Write(b)
	})

}
