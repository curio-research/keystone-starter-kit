package network

import (
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/curio-research/keystone/utils"
	"gorm.io/gorm"
)

type MySQLSaveTransactionHandler struct {
	transactionTable *TransactionTable
	randSeed         int
}

func NewMySQLSaveTransactionHandler(dialector gorm.Dialector, randSeed int) (*MySQLSaveTransactionHandler, error) {
	handler := &MySQLSaveTransactionHandler{randSeed: randSeed}
	db, err := gorm.Open(dialector, &defaultGormOpts)

	if err != nil {
		panic(err)
	}

	txTable, err := NewTransactionTable(db)
	if err != nil {
		return nil, err
	}

	handler.transactionTable = txTable
	return handler, nil
}

func (h *MySQLSaveTransactionHandler) SaveTransactions(ctx *server.EngineCtx, transactions []server.TransactionSchema) error {
	updatesForSql := []TransactionSQLFormat{}
	for _, transaction := range transactions {
		updatesForSql = append(updatesForSql, TransactionSQLFormat{
			GameId:        ctx.GameId,
			UnixTimestamp: transaction.UnixTimestamp,
			Tick:          transaction.TickNumber,
			Data:          transaction.Data,
			Type:          transaction.Type,
		})

	}

	return h.transactionTable.AddEntries(updatesForSql...)
}

func (h *MySQLSaveTransactionHandler) RestoreStateFromTxs(tickNumber int, gameId string) (*state.GameWorld, error) {
	newCtx := newGameEngineForRestoreStateFromTxs(h.randSeed)
	err := h.RestoreStateWithTxsOntoContext(newCtx, tickNumber, gameId)
	if err != nil {
		return nil, err
	}

	return newCtx.World, nil
}

func (h *MySQLSaveTransactionHandler) RestoreStateWithTxsOntoContext(ctx *server.EngineCtx, tickNumber int, gameId string) error {
	entries, err := h.transactionTable.GetEntriesUntilTick(tickNumber)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		server.AddSystemTransaction(ctx.World, entry.Tick, entry.Type, entry.Data, "", false)
	}

	utils.TickWorldForward(ctx, tickNumber)

	return nil
}
