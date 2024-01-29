package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func Test_CreateToken(t *testing.T) {
	tests := []struct {
		name   string
		claims Claims
	}{
		{
			name: "AccessTokenClaims",
			claims: &AccessTokenClaims{
				ClaimsType:  AccessTokenClaimsType,
				UserID:      "user_id",
				UserName:    "user_name",
				Email:       "email",
				Permissions: []string{"permission1", "permission2"},
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "issuer",
					Subject:   "subject",
					Audience:  []string{"audience"},
					IssuedAt:  jwt.NewNumericDate(time.Time{}),
					ExpiresAt: jwt.NewNumericDate(time.Time{}),
				},
			},
		},
		{
			name: "RefreshTokenClaims",
			claims: &RefreshTokenClaims{
				ClaimsType: RefreshTokenClaimsType,
				UserID:     "user_id",
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "issuer",
					Subject:   "subject",
					Audience:  []string{"audience"},
					IssuedAt:  jwt.NewNumericDate(time.Time{}),
					ExpiresAt: jwt.NewNumericDate(time.Time{}),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(*testing.T) {
			require := require.New(t)
			kc := testKCHS512(t)
			p, err := NewJWTProvider(kc, "an_issuer", 10*time.Second, 24*time.Hour)
			require.NoError(err)
			token, err := p.CreateToken(tt.claims)
			require.NoError(err)
			require.NotEqual(time.Time{}, token.ExpiresAt)
			require.NotEqual(time.Time{}, token.IssuedAt)
			require.NotEmpty(token.Token)

			var claims Claims
			switch tt.claims.(type) {
			case *AccessTokenClaims:
				claims = &AccessTokenClaims{}
			case *RefreshTokenClaims:
				claims = &RefreshTokenClaims{}
			}
			err = p.ParseToken(token.Token, claims)
			require.NoError(err)

			expiresAt, err := claims.GetExpirationTime()
			require.NoError(err)
			require.Equal(jwt.NewNumericDate(token.ExpiresAt), expiresAt)
			fmt.Println(expiresAt)
			fmt.Println(jwt.NewNumericDate(token.ExpiresAt))
			fmt.Println(time.Now())

			issuer, err := claims.GetIssuer()
			require.NoError(err)
			require.Equal("an_issuer", issuer)

			subject, err := claims.GetSubject()
			require.NoError(err)
			require.Equal("user_id", subject)
			audience, err := claims.GetAudience()
			require.NoError(err)
			require.Equal(jwt.ClaimStrings{"an_issuer"}, audience)

			switch tt.claims.(type) {
			case *AccessTokenClaims:
				require.Equal("user_id", claims.(*AccessTokenClaims).UserID)
				require.Equal("user_name", claims.(*AccessTokenClaims).UserName)
				require.Equal("email", claims.(*AccessTokenClaims).Email)
				require.Equal([]string{"permission1", "permission2"}, claims.(*AccessTokenClaims).Permissions)
				require.Less(
					token.ExpiresAt.Sub(token.IssuedAt),
					30*time.Second,
				)
			case *RefreshTokenClaims:
				require.Equal("user_id", claims.(*RefreshTokenClaims).UserID)
				require.Greater(
					token.ExpiresAt.Sub(token.IssuedAt),
					12*time.Hour,
				)
			}
		})
	}
}
