package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_CreateToken(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	provider, err := NewJWTProvider(
		"some_secret",
		"some_issuer",
		10*time.Second,
	)
	require.NoError(err)

	signedToken, err := provider.CreateToken("some_user_id", "some_user_name", []string{"some_permission"})
	require.NoError(err)

	claims := signedToken.Claims()
	require.Equal("some_user_id", claims.UserID)
	require.Equal("some_user_name", claims.UserName)
	require.Equal("some_issuer", claims.Issuer)
	require.Equal("some_user_id", claims.Subject)
	require.Equal("some_issuer", claims.Audience[0])
	require.Contains(claims.Permissions, "some_permission")

	claims2, err := provider.ParseToken(signedToken.String())
	require.NoError(err)
	require.Equal(claims, *claims2)
}

func Test_InvalidSigningKey(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	_, err := NewJWTProvider(
		"",
		"some_issuer",
		10*time.Second,
	)
	require.Error(err)
	_, ok := err.(ErrInvalidSigningKey)
	require.True(ok)
}

func Test_ExpiredToken(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	prov, err := NewJWTProvider(
		"some_secret",
		"some_issuer",
		1*time.Nanosecond,
	)
	require.NoError(err)

	token, err := prov.CreateToken(
		"some_user_id",
		"some_user_name",
		[]string{"some_permission"},
	)
	require.NoError(err)

	time.Sleep(2 * time.Nanosecond)
	_, err = prov.ParseToken(token.String())
	require.Error(err)
}

func Test_BadToken(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	prov, err := NewJWTProvider(
		"some_secret",
		"some_issuer",
		10*time.Second,
	)
	require.NoError(err)

	_, err = prov.ParseToken("some_bad_token")
	require.Error(err)
}
