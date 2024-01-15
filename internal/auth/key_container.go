package auth

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type KeyContainer interface {
	SignedToken(Claims) (string, error)
	VerificationKeyFunc(token *jwt.Token) (interface{}, error)
}

type ErrInvalidSigningKey struct{}

func (e ErrInvalidSigningKey) Error() string {
	return "invalid signing key"
}

type ErrInvalidVerificationKey struct{}

func (e ErrInvalidVerificationKey) Error() string {
	return "invalid verification key"
}

type HS512KeyContainer struct {
	signingKey []byte
}

func NewHS512KeyContainer(signingKey string) (*HS512KeyContainer, error) {
	if signingKey == "" {
		return nil, ErrInvalidSigningKey{}
	}
	return &HS512KeyContainer{
		signingKey: []byte(signingKey),
	}, nil
}

func (k *HS512KeyContainer) SignedToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(k.signingKey)
}

func (k *HS512KeyContainer) VerificationKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, ErrInvalidAlgorithm{}
	}
	return k.signingKey, nil
}

// Generating Keys:
//
//	ssh-keygen -t rsa -b 4096 -m PEM -E SHA512 -f jwtRS512.key -N ""
//	# Don't add passphrase
//	openssl rsa -in jwtRS512.key -pubout -outform PEM -out jwtRS512.key.pub
type RS512KeyContainer struct {
	signingKey      *rsa.PrivateKey
	verificationKey *rsa.PublicKey
}

func NewRS512KeyContainer(signingKeyPath string, verificationKeyPath string) (*RS512KeyContainer, error) {
	signBytes, err := os.ReadFile(signingKeyPath)
	if err != nil {
		return nil, err
	}

	verifyBytes, err := os.ReadFile(verificationKeyPath)
	if err != nil {
		return nil, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return &RS512KeyContainer{
		signingKey:      signKey,
		verificationKey: verifyKey,
	}, nil
}

func (k *RS512KeyContainer) SignedToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	return token.SignedString(k.signingKey)
}

func (k *RS512KeyContainer) VerificationKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, ErrInvalidAlgorithm{}
	}
	return k.verificationKey, nil
}
