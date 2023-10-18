package network

import (
	"strconv"

	pb_test "github.com/curio-research/keystone/game/proto/output/pb.test"

	pb_battle "github.com/curio-research/keystone/game/proto/output/pb.battle"
	pb_dict "github.com/curio-research/keystone/game/proto/output/pb.dict"
	pb_game "github.com/curio-research/keystone/game/proto/output/pb.game"
	pb_round "github.com/curio-research/keystone/game/proto/output/pb.round"

	"github.com/curio-research/keystone/server"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

// there are currently two ways users can submit inputs that will be queued as transactions
// for the production game, users send inputs via websockets per Unity team's convention

// the websocket router routes incoming requests based on protobuf types
func SocketRequestRouter(ctx *server.EngineCtx, requestMsg *server.NetworkMessage, socketConnection *websocket.Conn) {

	// data received through websocket from game clients
	requestType := pb_dict.CMD(requestMsg.GetCommand())

	// route incoming data based on command routes
	switch requestType {

	// Game
	case pb_dict.CMD_pb_game_C2S_EstablishPlayer: // assign player to the connection
		req := queueTxIntoSystems[*pb_game.C2S_EstablishPlayer](ctx, requestMsg, &pb_game.C2S_EstablishPlayer{})
		ctx.Stream.SetPlayerIdToConnection(int(req.PlayerId), socketConnection) // assign playerId to a connection. this allows us to broadcast events to specific players
	case pb_dict.CMD_pb_game_C2S_GameState:
		queueTxIntoSystems[*pb_game.C2S_GameState](ctx, requestMsg, &pb_game.C2S_GameState{})
	case pb_dict.CMD_pb_game_C2S_Ping:
		queueTxIntoSystems[*pb_game.C2S_Ping](ctx, requestMsg, &pb_game.C2S_Ping{})
	case pb_dict.CMD_pb_game_C2S_ClaimDealer:
		queueTxIntoSystems[*pb_game.C2S_ClaimDealer](ctx, requestMsg, &pb_game.C2S_ClaimDealer{})
	case pb_dict.CMD_pb_game_C2S_PlayerReady:
		queueTxIntoSystems[*pb_game.C2S_PlayerReady](ctx, requestMsg, &pb_game.C2S_PlayerReady{})
	case pb_dict.CMD_pb_game_C2S_PreparationState:
		queueTxIntoSystems[*pb_game.C2S_PreparationState](ctx, requestMsg, &pb_game.C2S_PreparationState{})

	// Round
	case pb_dict.CMD_pb_round_C2S_DiscardCards:
		queueTxIntoSystems[*pb_round.C2S_DiscardCards](ctx, requestMsg, &pb_round.C2S_DiscardCards{})
	case pb_dict.CMD_pb_round_C2S_TurnEnd:
		queueTxIntoSystems[*pb_round.C2S_TurnEnd](ctx, requestMsg, &pb_round.C2S_TurnEnd{})

	// Battle
	case pb_dict.CMD_pb_battle_C2S_Produce:
		queueTxIntoSystems[*pb_battle.C2S_Produce](ctx, requestMsg, &pb_battle.C2S_Produce{})
	case pb_dict.CMD_pb_battle_C2S_ProduceBuilding:
		queueTxIntoSystems[*pb_battle.C2S_ProduceBuilding](ctx, requestMsg, &pb_battle.C2S_ProduceBuilding{})
	case pb_dict.CMD_pb_battle_C2S_MoveTroops:
		queueTxIntoSystems[*pb_battle.C2S_MoveTroops](ctx, requestMsg, &pb_battle.C2S_MoveTroops{})
	case pb_dict.CMD_pb_battle_C2S_Attack:
		queueTxIntoSystems[*pb_battle.C2S_Attack](ctx, requestMsg, &pb_battle.C2S_Attack{})
	case pb_dict.CMD_pb_battle_C2S_ToggleTankGuardMode:
		queueTxIntoSystems[*pb_battle.C2S_ToggleTankGuardMode](ctx, requestMsg, &pb_battle.C2S_ToggleTankGuardMode{})
	case pb_dict.CMD_pb_battle_C2S_PlaneLoad:
		queueTxIntoSystems[*pb_battle.C2S_PlaneLoad](ctx, requestMsg, &pb_battle.C2S_PlaneLoad{})
	case pb_dict.CMD_pb_battle_C2S_PlaneUnload:
		queueTxIntoSystems[*pb_battle.C2S_PlaneUnload](ctx, requestMsg, &pb_battle.C2S_PlaneUnload{})
	case pb_dict.CMD_pb_battle_C2S_UpgradeBuilding:
		queueTxIntoSystems[*pb_battle.C2S_UpgradeBuilding](ctx, requestMsg, &pb_battle.C2S_UpgradeBuilding{})
	case pb_dict.CMD_pb_battle_C2S_UpgradeCapital:
		queueTxIntoSystems[*pb_battle.C2S_UpgradeCapital](ctx, requestMsg, &pb_battle.C2S_UpgradeCapital{})
	case pb_dict.CMD_pb_battle_C2S_ProduceBlueprint:
		queueTxIntoSystems[*pb_battle.C2S_ProduceBlueprint](ctx, requestMsg, &pb_battle.C2S_ProduceBlueprint{})

	// Test
	case pb_dict.CMD_pb_test_C2S_ResetGameState:
		queueTxIntoSystems[*pb_test.C2S_ResetGameState](ctx, requestMsg, &pb_test.C2S_ResetGameState{})
	case pb_dict.CMD_pb_test_C2S_Test: // No-op, only used in integration tests
		queueTxIntoSystems[*pb_test.C2S_Test](ctx, requestMsg, &pb_test.C2S_Test{})
	}
}

// queue transactions for systems from the outside
func queueTxIntoSystems[T proto.Message](ctx *server.EngineCtx, requestMsg *server.NetworkMessage, req T) T {
	requestMsg.GetProtoMessage(req)
	requestId := requestMsg.Param()

	server.QueueTxFromExternal(ctx, req, strconv.Itoa(int(requestId)))
	return req
}
