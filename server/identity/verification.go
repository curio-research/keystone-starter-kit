package identity

import (
	"errors"
	"os"

	pb_base "github.com/curio-research/keystone/game/proto/output/pb.base"
	"github.com/golang-jwt/jwt"

	"github.com/curio-research/keystone/server"
)

// jwt verification payload
type JwtPayload struct {
	// universal player ID
	PlayerId int

	jwt.StandardClaims
}

func DecodeJwtPayload(tokenString string) (JwtPayload, error) {
	payload := &JwtPayload{}

	jwtTokenString := os.Getenv("JWT_AUTH")

	// Parse the JWT token and store the claims in the 'claims' variable
	token, err := jwt.ParseWithClaims(tokenString, payload, func(token *jwt.Token) (interface{}, error) {
		// Replace with your secret key used to sign the JWT token
		return []byte(jwtTokenString), nil
	})

	if err != nil {
		return JwtPayload{}, err
	} else if token.Valid {
		return *payload, nil
	}

	return JwtPayload{}, errors.New("Invalid token")
}

// returns: playerId, isVerified
func VerifyIdentity(ctx *server.EngineCtx, payload *pb_base.IdentityPayload) (int, bool) {
	if ctx.Mode == "prod" {
		// playerId derived from the identity payload after signature verification
		jwtPayload, err := DecodeJwtPayload(payload.JwtToken)
		if err != nil {
			return 0, false
		} else {
			return jwtPayload.PlayerId, true
		}

	}

	// in dev mode, bypass the verify identity and return playerID directly
	return int(payload.PlayerId), true
}
