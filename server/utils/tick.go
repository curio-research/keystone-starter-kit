package utils

import "github.com/curio-research/keystone-starter-kit/constants"

// =========================
// Query
// =========================

func CalcTickFromMsInFuture(tickNumber int, timeInMs int) int {
	return tickNumber + (timeInMs / constants.TickRate)
}

func CalcTickFromSecInFuture(tickNumber int, timeInSeconds int) int {
	return CalcTickFromMsInFuture(tickNumber, timeInSeconds*1000)
}
