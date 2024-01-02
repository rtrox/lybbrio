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

	signedToken, err := provider.CreateToken("some_user_id", "some_user_name")
	require.NoError(err)

	claims := signedToken.Claims()
	require.Equal("some_user_id", claims.UserID)
	require.Equal("some_user_name", claims.UserName)
	require.Equal("some_issuer", claims.Issuer)
	require.Equal("some_user_id", claims.Subject)
	require.Equal("some_issuer", claims.Audience[0])

	claims2, err := provider.ParseToken(signedToken.String())
	require.NoError(err)
	require.Equal(claims, *claims2)
}
