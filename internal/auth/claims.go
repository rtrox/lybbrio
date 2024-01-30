package auth

import "github.com/golang-jwt/jwt/v5"

type Claims interface {
	Subject() string
	SetRegisteredClaims(claims jwt.RegisteredClaims)
	jwt.Claims
}

type RefreshTokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewRefreshTokenClaims(userID string) *RefreshTokenClaims {
	return &RefreshTokenClaims{
		UserID: userID,
	}
}

func (c RefreshTokenClaims) Subject() string {
	return c.UserID
}

func (c *RefreshTokenClaims) SetRegisteredClaims(claims jwt.RegisteredClaims) {
	c.RegisteredClaims = claims
}

type AccessTokenClaims struct {
	UserID      string   `json:"user_id"`
	UserName    string   `json:"user_name"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

func NewAccessTokenClaims(userID, userName, email string, permissions []string) *AccessTokenClaims {
	return &AccessTokenClaims{
		UserID:      userID,
		UserName:    userName,
		Email:       email,
		Permissions: permissions,
	}
}

func (c AccessTokenClaims) Subject() string {
	return c.UserID
}

func (c *AccessTokenClaims) SetRegisteredClaims(claims jwt.RegisteredClaims) {
	c.RegisteredClaims = claims
}
