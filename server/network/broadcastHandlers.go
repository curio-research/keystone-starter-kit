package network

import (
	"strings"

	pb_dict "github.com/curio-research/keystone/game/proto/output/pb.dict"
	pb_game "github.com/curio-research/keystone/game/proto/output/pb.game"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

// TODO: refurther refactor to separate proto and error handler

// error handler for emitting errors from systems
// protobuf based implementation for our game

type ProtoBasedErrorHandler struct {
}

// format message into protobuf
func (h *ProtoBasedErrorHandler) FormatMessage(transactionUuidIdentifier int, errorMessage string) *server.NetworkMessage {
	msg, _ := server.NewMessage(0, uint32(pb_dict.CMD_pb_game_S2C_ServerMessage), uint32(transactionUuidIdentifier), &pb_game.S2C_ServerMessage{
		Content: errorMessage,
	})

	return msg
}

type ProtoBasedBroadcastHandler struct {
}

func (h *ProtoBasedBroadcastHandler) BroadcastMessage(ctx *server.EngineCtx, clientEvents []server.ClientEvent) {
	stateChanges := filterTableUpdatesWithoutLocal(ctx.World.TableUpdates)

	if ctx.ShouldRecordError {
		// for error logging purposes during testing, log them

		// loop through all client client events and see which one is a server message
		for _, clientEvent := range clientEvents {
			if clientEvent.NetworkMessage.GetCommand() == uint32(pb_dict.CMD_pb_game_S2C_ServerMessage) {

				// decode the error message string from proto and log it
				data := &pb_game.S2C_ServerMessage{}
				clientEvent.NetworkMessage.GetProtoMessage(data)

				ctx.ErrorLog = append(ctx.ErrorLog, server.ErrorLog{
					Tick:    ctx.GameTick.TickNumber,
					Message: data.Content,
				})
			}
		}
	}

	if len(stateChanges) == 0 && clientEvents == nil {
		return
	}

	ctx.Stream.PublishStateChanges(stateChanges, clientEvents)

}

func filterTableUpdatesWithoutLocal(tableUpdates state.TableUpdateArray) state.TableUpdateArray {
	// if the table name starts with the word local, then filter it out and not broadcast it

	filteredUpdates := state.TableUpdateArray{}

	for _, tableUpdate := range tableUpdates {
		if !strings.HasPrefix(tableUpdate.Table, "local") {
			filteredUpdates = append(filteredUpdates, tableUpdate)
		}
	}

	return filteredUpdates
}
