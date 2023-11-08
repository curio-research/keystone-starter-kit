package systems

import (
	"encoding/json"

	"github.com/curio-research/keystone/server"
)

const PlayerIDHeader = "playerIDHeader"

const PlayerIDTag = "playerID"
const PublicKeyTag = "publicKey"

func VerifyWalletAndIdentity[T any]() server.IMiddleware[T] {
	return func(ctx *server.TransactionCtx[T]) bool {
		req := ctx.Req
		headers := req.Headers

		if headers == nil {
			return false
		}

		publicKeyAuth := headers[server.EthereumWalletAuthHeader]
		if publicKeyAuth == nil {
			return false
		}
		var e server.EthereumWalletAuth
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

		if !e.Verify() {
			return false
		}

		ctx.Meta[PlayerIDTag] = playerID
		ctx.Meta[PublicKeyTag] = e.Base64PublicKey

		return true
	}
}

func matchPublicKey[T any](ctx *server.TransactionCtx[T], e server.EthereumWalletAuth, playerID int) bool {
	player, found := PlayerWithID(ctx.W, playerID)
	if !found {
		return false
	}

	if player.EthBase64PublicKey != e.Base64PublicKey {
		return false
	}
	return true
}
