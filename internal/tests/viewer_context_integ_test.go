package test

import (
	"context"
	"lybbrio/internal/auth"
	"lybbrio/internal/db"
	"lybbrio/internal/middleware"
	"lybbrio/internal/viewer"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
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

	token, err := jwtProvider.CreateToken(user.ID.String(), user.Username)
	require.NoError(err)

	r := chi.NewRouter()
	r.Use(auth.Middleware(jwtProvider))
	r.Use(middleware.ViewerContextMiddleware(client))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		viewerCtx := viewer.FromContext(r.Context())
		require.NotNil(viewerCtx)

		viewerUser, ok := viewerCtx.User()
		require.True(ok)
		viewerPerms, ok := viewerCtx.Permissions()
		require.True(ok)

		require.Equal(user.ID, viewerUser.ID)
		require.Equal(perms.ID, viewerPerms.ID)
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
