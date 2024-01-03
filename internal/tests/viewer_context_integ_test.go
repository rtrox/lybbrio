package test

import (
	"context"
	"lybbrio/internal/auth"
	"lybbrio/internal/db"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/middleware"
	"lybbrio/internal/viewer"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func Test_ViewerContextGetsSet(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	client := db.OpenTest(t, "Test_ViewerContextGetsSet")
	defer client.Close()

	adminCtx := viewer.NewSystemAdminContext(context.Background())
	perms := client.UserPermissions.Create().
		SetAdmin(true).
		SaveX(adminCtx)
	user := client.User.Create().
		SetEmail("asdf@asdf.com").
		SetUsername("asdf").
		SetUserPermissions(perms).
		SaveX(adminCtx)

	jwtProvider, err := auth.NewJWTProvider(
		"some_secret",
		"some_issuer",
		10*time.Second,
	)
	require.NoError(err)

	token, err := jwtProvider.CreateToken(
		user.ID.String(),
		user.Username,
		permissions.From(perms).StringSlice())
	require.NoError(err)

	r := chi.NewRouter()
	r.Use(auth.Middleware(jwtProvider))
	r.Use(middleware.ViewerContextMiddleware(client))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		viewerCtx := viewer.FromContext(r.Context())
		require.NotNil(viewerCtx)

		viewerUserID, ok := viewerCtx.UserID()
		require.True(ok)

		require.Equal(user.ID, viewerUserID)
		require.True(viewerCtx.Has(permissions.Admin))
		require.True(viewerCtx.IsAdmin())
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	require.NoError(err)

	req.Header.Set("X-Api-Token", token.String())

	_, err = http.DefaultClient.Do(req)
	require.NoError(err)

}
