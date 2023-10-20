package test

import (
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/network"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"sync"
)

func NewTestEngine(gameWorld *state.GameWorld, systems ...server.TickSystemFunction) *server.EngineCtx {
	gameTick := server.NewGameTick(constants.TickRate)

	// initiate an empty tick schedule
	tickSchedule := server.NewTickSchedule()
	gameTick.Schedule = tickSchedule
	for _, system := range systems {
		tickSchedule.AddTickSystem(constants.TickRate, system)
	}

	gameCtx := &server.EngineCtx{
		GameId:                 "prototype-game",
		IsLive:                 true,
		World:                  gameWorld,
		GameTick:               gameTick,
		TransactionsToSaveLock: sync.Mutex{},
		ShouldRecordError:      true,
		ErrorLog:               []server.ErrorLog{},
		Mode:                   "dev",
		SystemErrorHandler:     &network.ProtoBasedErrorHandler{},
		SystemBroadcastHandler: &network.ProtoBasedBroadcastHandler{},
	}

	return gameCtx
}
