package auth

import "github.com/golang-jwt/jwt/v5"

type Claims interface {
	Subject() string
	SetRegisteredClaims(claims jwt.RegisteredClaims)
	ValidateType() error
	jwt.Claims
}

type ClaimsType int

const (
	Access ClaimsType = iota + 1
	Refresh
)

type RefreshTokenClaims struct {
	Type   ClaimsType
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func NewRefreshTokenClaims(userID string) *RefreshTokenClaims {
	return &RefreshTokenClaims{
		UserID: userID,
		Type:   Refresh,
	}
}

func (c RefreshTokenClaims) Subject() string {
	return c.UserID
}

func (c *RefreshTokenClaims) SetRegisteredClaims(claims jwt.RegisteredClaims) {
	c.RegisteredClaims = claims
}

func (c RefreshTokenClaims) ValidateType() error {
	if c.Type != Refresh {
		return ErrInvalidClaimsType
	}
	return nil
}

type AccessTokenClaims struct {
	UserID      string   `json:"user_id"`
	UserName    string   `json:"user_name"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions"`
	Type        ClaimsType
	jwt.RegisteredClaims
}

func NewAccessTokenClaims(userID, userName, email string, permissions []string) *AccessTokenClaims {
	return &AccessTokenClaims{
		UserID:      userID,
		UserName:    userName,
		Email:       email,
		Permissions: permissions,
		Type:        Access,
	}
}

func (c AccessTokenClaims) Subject() string {
	return c.UserID
}

func (c *AccessTokenClaims) SetRegisteredClaims(claims jwt.RegisteredClaims) {
	c.RegisteredClaims = claims
}

func (c AccessTokenClaims) ValidateType() error {
	if c.Type != Access {
		return ErrInvalidClaimsType
	}
	return nil
}
