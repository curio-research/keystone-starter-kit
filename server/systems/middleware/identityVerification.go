package middleware

import (
	"github.com/curio-research/keystone/game/constants"
	"github.com/curio-research/keystone/game/identity"
	pb_base "github.com/curio-research/keystone/game/proto/output/pb.base"
	"github.com/curio-research/keystone/server"
)

// we assume that all incoming requests have an identity payload
type IVerifyIdentity interface {
	GetIdentityPayload() *pb_base.IdentityPayload
}

// verifies that a request has a valid jwt identity payload
func VerifyIdentity[T IVerifyIdentity]() server.IMiddleware[T] {
	return func(ctx *server.TransactionCtx[T], req T) bool {
		playerId, verified := identity.VerifyIdentity(ctx.GameCtx, req.GetIdentityPayload())
		if !verified {
			ctx.EmitError(constants.IdentityVerificationErrorString, []int{playerId})
			return false
		}

		return true
	}
}
