package auth

import "errors"

var (
	ErrInvalidToken           = errors.New("invalid token")
	ErrInvalidSigningKey      = errors.New("invalid signing key")
	ErrInvalidVerificationKey = errors.New("invalid verification key")
	ErrInvalidAlgorithm       = errors.New("invalid algorithm")
	ErrInvalidClaimsType      = errors.New("invalid claim type")
)
