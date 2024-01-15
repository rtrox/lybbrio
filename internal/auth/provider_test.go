package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func testProviderHS512(t *testing.T, expiration ...time.Duration) *JWTProvider {
	if len(expiration) == 0 {
		expiration = []time.Duration{10 * time.Second}
	}
	kc, err := NewHS512KeyContainer("some_secret")
	require.NoError(t, err)
	provider, err := NewJWTProvider(
		kc,
		"some_issuer",
		expiration[0],
	)
	require.NoError(t, err)
	return provider
}

func testProviderRS512(t *testing.T, expiration ...time.Duration) *JWTProvider {
	if len(expiration) == 0 {
		expiration = []time.Duration{10 * time.Second}
	}
	kc, err := NewRS512KeyContainer(
		"test/jwtRS512.key",
		"test/jwtRS512.key.pub",
	)
	require.NoError(t, err)
	provider, err := NewJWTProvider(
		kc,
		"some_issuer",
		expiration[0],
	)
	require.NoError(t, err)
	return provider
}

type providerTestCase struct {
	name     string
	provider *JWTProvider
}

func testEachProvider(t *testing.T, expiration ...time.Duration) []providerTestCase {
	return []providerTestCase{
		{
			name:     "HS512",
			provider: testProviderHS512(t, expiration...),
		},
		{
			name:     "RS512",
			provider: testProviderRS512(t, expiration...),
		},
	}
}

func Test_CreateToken(t *testing.T) {
	for _, tt := range testEachProvider(t) {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			signedToken, err := tt.provider.CreateToken("some_user_id", "some_user_name", []string{"some_permission"})
			require.NoError(err)

			claims := signedToken.Claims()
			require.Equal("some_user_id", claims.UserID)
			require.Equal("some_user_name", claims.UserName)
			require.Equal("some_issuer", claims.Issuer)
			require.Equal("some_user_id", claims.Subject)
			require.Equal("some_issuer", claims.Audience[0])
			require.Contains(claims.Permissions, "some_permission")

			claims2, err := tt.provider.ParseToken(signedToken.String())
			require.NoError(err)
			require.Equal(claims, *claims2)
		})
	}
}

func Test_ExpiredToken(t *testing.T) {
	for _, tt := range testEachProvider(t, 1*time.Nanosecond) {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			token, err := tt.provider.CreateToken(
				"some_user_id",
				"some_user_name",
				[]string{"some_permission"},
			)
			require.NoError(err)

			time.Sleep(20 * time.Nanosecond)
			_, err = tt.provider.ParseToken(token.String())
			require.Error(err)
		})
	}
}

func Test_BadToken(t *testing.T) {
	for _, tt := range testEachProvider(t) {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			_, err := tt.provider.ParseToken("some_bad_token")
			require.Error(err)
		})
	}
}
