package constants

import "errors"

var (
	IdentityVerificationErrorString = "identity verification failed"

	// errors
	IdentityVerificationError = errors.New(IdentityVerificationErrorString)
)
