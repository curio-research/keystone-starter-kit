package network

import (
	"fmt"
	serverpb "github.com/curio-research/keystone/game/serverpb/output"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"strings"
)

// TODO: refurther refactor to separate serverpb and error handler

// error handler for emitting errors from systems
// protobuf based implementation for our game

type ProtoBasedErrorHandler struct {
}

// format message into protobuf
func (h *ProtoBasedErrorHandler) FormatMessage(transactionUuidIdentifier int, errorMessage string) *server.NetworkMessage {
	fmt.Println(errorMessage)
	msg, _ := server.NewMessage(0, uint32(serverpb.CMD_S2C_Error), uint32(transactionUuidIdentifier), &serverpb.S2C_ErrorMessage{
		Content: errorMessage,
	})

	return msg
}

type ProtoBasedBroadcastHandler struct {
}

func (h *ProtoBasedBroadcastHandler) BroadcastMessage(ctx *server.EngineCtx, clientEvents []server.ClientEvent) {
	stateChanges := filterTableUpdatesWithoutLocal(ctx.World.TableUpdates)

	if ctx.ShouldRecordError {
		// loop through all client client events and see which one is an error
		for _, clientEvent := range clientEvents {
			if clientEvent.NetworkMessage.GetCommand() == uint32(serverpb.CMD_S2C_Error) {

				// decode the error message string from serverpb and log it
				data := &serverpb.S2C_ErrorMessage{}
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
