package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ErrInvalidToken struct{}

func (e ErrInvalidToken) Error() string {
	return "invalid token"
}

type ErrInvalidAlgorithm struct{}

func (e ErrInvalidAlgorithm) Error() string {
	return "invalid algorithm"
}

type Claims struct {
	UserID      string   `json:"user_id"`
	UserName    string   `json:"user_name"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

type SignedToken struct {
	token  string
	claims Claims
}

func (s SignedToken) String() string {
	return s.token
}

func (s SignedToken) Claims() Claims {
	return s.claims
}

type JWTProvider struct {
	keyContainer KeyContainer
	issuer       string
	expiry       time.Duration
}

func NewJWTProvider(keyContainer KeyContainer, issuer string, expiry time.Duration) (*JWTProvider, error) {
	return &JWTProvider{
		keyContainer: keyContainer,
		issuer:       issuer,
		expiry:       expiry,
	}, nil
}

func (p *JWTProvider) CreateToken(userID, userName string, permissions []string) (SignedToken, error) {
	claims := Claims{
		UserID:      userID,
		UserName:    userName,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    p.issuer,
			Subject:   userID,
			Audience:  jwt.ClaimStrings{p.issuer},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(p.expiry)),
		},
	}
	signed, err := p.keyContainer.SignedToken(claims)
	if err != nil {
		return SignedToken{}, err
	}
	return SignedToken{
		token:  signed,
		claims: claims,
	}, nil
}

func (p *JWTProvider) ParseToken(tokenString string) (*Claims, error) {
	c := Claims{}
	_, err := jwt.ParseWithClaims(tokenString, &c, p.keyContainer.VerificationKeyFunc)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
