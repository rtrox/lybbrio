package auth

import "github.com/golang-jwt/jwt/v5"

type ClaimsType int

const (
	RefreshTokenClaimsType ClaimsType = iota + 1
	AccessTokenClaimsType
)

type Claims interface {
	Type() ClaimsType
	Subject() string
	SetRegisteredClaims(claims jwt.RegisteredClaims)
	jwt.Claims
}

type GenericClaims struct {
	ClaimsType ClaimsType `json:"type"`
	jwt.RegisteredClaims
}

func (c GenericClaims) Subject() string {
	return c.RegisteredClaims.Subject
}

func (c *GenericClaims) SetRegisteredClaims(claims jwt.RegisteredClaims) {
	c.RegisteredClaims = claims
}

func (c GenericClaims) Type() ClaimsType {
	return c.ClaimsType
}

type RefreshTokenClaims struct {
	ClaimsType ClaimsType `json:"type"`
	UserID     string     `json:"user_id"`
	jwt.RegisteredClaims
}

func NewRefreshTokenClaims(userID string) *RefreshTokenClaims {
	return &RefreshTokenClaims{
		ClaimsType: RefreshTokenClaimsType,
		UserID:     userID,
	}
}

func (c RefreshTokenClaims) Subject() string {
	return c.UserID
}

func (c *RefreshTokenClaims) SetRegisteredClaims(claims jwt.RegisteredClaims) {
	c.RegisteredClaims = claims
}

func (c RefreshTokenClaims) Type() ClaimsType {
	return c.ClaimsType
}

type AccessTokenClaims struct {
	ClaimsType  ClaimsType `json:"type"`
	UserID      string     `json:"user_id"`
	UserName    string     `json:"user_name"`
	Email       string     `json:"email"`
	Permissions []string   `json:"permissions"`
	jwt.RegisteredClaims
}

func NewAccessTokenClaims(userID, userName, email string, permissions []string) *AccessTokenClaims {
	return &AccessTokenClaims{
		ClaimsType:  AccessTokenClaimsType,
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

func (c AccessTokenClaims) Type() ClaimsType {
	return c.ClaimsType
}
