package middleware

import (
	"encoding/json"
	"github.com/curio-research/keystone-starter-kit/systems"
	"github.com/curio-research/keystone/server"
)

const PlayerIDHeader = "playerIDHeader"

func VerifyIdentity[T any]() server.IMiddleware[T] {
	return func(ctx *server.TransactionCtx[T]) bool {
		req := ctx.Req
		headers := req.Headers

		if headers == nil {
			return false
		}

		publicKeyAuth := headers[server.ECDSAPublicKeyAuthHeader]
		if publicKeyAuth == nil {
			return false
		}
		var e server.ECDSAPublicKeyAuth
		err := json.Unmarshal(publicKeyAuth, &e)
		if err != nil {
			return false
		}

		var p int
		playerID := headers[PlayerIDHeader]
		if playerID == nil {
			return false
		}
		err = json.Unmarshal(playerID, &p)
		if err != nil {
			return false
		}

		matching := matchPublicKey(ctx, e, p)
		if !matching {
			return false
		}

		return e.Verify()
	}
}

func matchPublicKey[T any](ctx *server.TransactionCtx[T], e server.ECDSAPublicKeyAuth, playerID int) bool {
	player, found := systems.PlayerWithID(ctx.W, playerID)
	if !found {
		return false
	}

	if player.Base64PublicKey != e.Base64PublicKey {
		return false
	}
	return true
}
