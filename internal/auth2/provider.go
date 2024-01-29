package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignedToken struct {
	Token     string    `json:"token"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type JWTProvider struct {
	keyContainer  KeyContainer
	issuer        string
	expiry        time.Duration
	refreshExpiry time.Duration
}

func NewJWTProvider(keyContainer KeyContainer, issuer string, expiry time.Duration, refreshExpiry time.Duration) (*JWTProvider, error) {
	return &JWTProvider{
		keyContainer:  keyContainer,
		issuer:        issuer,
		expiry:        expiry,
		refreshExpiry: refreshExpiry,
	}, nil
}

func (p *JWTProvider) CreateToken(claims Claims) (SignedToken, error) {
	expiry := p.ExpiryFromClaims(claims)

	reg := jwt.RegisteredClaims{
		Issuer:    p.issuer,
		Subject:   claims.Subject(),
		Audience:  jwt.ClaimStrings{p.issuer},
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
	}
	claims.SetRegisteredClaims(reg)
	token := jwt.NewWithClaims(p.keyContainer.SigningMethod(), claims)
	signedToken, err := token.SignedString(p.keyContainer.SigningKey())
	if err != nil {
		return SignedToken{}, err
	}
	return SignedToken{
		Token:     signedToken,
		IssuedAt:  reg.IssuedAt.Time,
		ExpiresAt: reg.ExpiresAt.Time,
	}, nil
}

func (p *JWTProvider) ExpiryFromClaims(claims Claims) time.Duration {
	switch claims.(type) {
	case *RefreshTokenClaims:
		return p.refreshExpiry
	case *AccessTokenClaims:
		return p.expiry
	}
	return time.Duration(0) // ensure any invalid claims type is expired
}

func (p JWTProvider) ParseToken(tokenString string, dest jwt.Claims) error {
	_, err := jwt.ParseWithClaims(tokenString, dest, p.keyContainer.VerificationKey)
	return err
}
