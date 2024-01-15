package auth

import (
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/viewer"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Middleware(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	provider := testProviderHS512(t)

	handler := ViewerContextMiddleware(provider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		view := viewer.FromContext(r.Context())
		uid, ok := view.UserID()
		require.True(ok)
		require.Equal("some_user_id", uid.String())
		require.True(view.Has(permissions.Admin))

	}))

	token, err := provider.CreateToken(
		"some_user_id",
		"some_user_name",
		[]string{"Admin"},
	)
	require.NoError(err)

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(err)
	req.Header.Add("X-Api-Token", token.String())

	handler.ServeHTTP(nil, req)

}

func Test_Middleware_BadToken(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	provider := testProviderHS512(t)

	wrongKC, err := NewHS512KeyContainer("some_wrong_secret")
	require.NoError(err)
	wrong_provider, err := NewJWTProvider(
		wrongKC,
		"some_issuer",
		10*time.Second,
	)
	require.NoError(err)
	token, err := wrong_provider.CreateToken(
		"some_user_id",
		"some_user_name",
		[]string{"Admin"},
	)
	require.NoError(err)

	handler := ViewerContextMiddleware(provider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req, err := http.NewRequest("GET", "/", nil)
	wr := httptest.NewRecorder()
	require.NoError(err)
	req.Header.Add("X-Api-Token", token.String())

	handler.ServeHTTP(wr, req)

	require.Equal(http.StatusUnauthorized, wr.Code)

	req, err = http.NewRequest("GET", "/", nil)
	require.NoError(err)
	wr = httptest.NewRecorder()
	req.Header.Add("X-Api-Token", "some_invalid_token")

	handler.ServeHTTP(wr, req)

	require.Equal(http.StatusUnauthorized, wr.Code)
}

func Test_Middleware_EmptyToken(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	provider := testProviderHS512(t)

	handler := ViewerContextMiddleware(provider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req, err := http.NewRequest("GET", "/", nil)
	wr := httptest.NewRecorder()
	require.NoError(err)

	handler.ServeHTTP(wr, req)

	require.Equal(http.StatusUnauthorized, wr.Code)
}
