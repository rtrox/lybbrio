package auth

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type KeyContainer interface {
	SigningMethod() jwt.SigningMethod
	SigningKey() interface{}
	VerificationKey(token *jwt.Token) (interface{}, error)
}

type HS512KeyContainer struct {
	signingKey []byte
}

var _ KeyContainer = &HS512KeyContainer{}

func NewHS512KeyContainer(signingKey string) (*HS512KeyContainer, error) {
	if signingKey == "" {
		return nil, ErrInvalidSigningKey
	}
	return &HS512KeyContainer{
		signingKey: []byte(signingKey),
	}, nil
}

func (k *HS512KeyContainer) SigningMethod() jwt.SigningMethod {
	return jwt.SigningMethodHS512
}

func (k *HS512KeyContainer) SigningKey() interface{} {
	return k.signingKey
}

func (k *HS512KeyContainer) VerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, ErrInvalidAlgorithm
	}
	return k.signingKey, nil
}

type RS512KeyContainer struct {
	signingKey      *rsa.PrivateKey
	verificationKey *rsa.PublicKey
}

var _ KeyContainer = &RS512KeyContainer{}

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

func (k *RS512KeyContainer) SigningMethod() jwt.SigningMethod {
	return jwt.SigningMethodRS512
}

func (k *RS512KeyContainer) SigningKey() interface{} {
	return k.signingKey
}

func (k *RS512KeyContainer) VerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, ErrInvalidAlgorithm
	}
	return k.verificationKey, nil
}
