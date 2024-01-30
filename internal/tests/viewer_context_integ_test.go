package test

import (
	"context"
	"fmt"
	"lybbrio/internal/auth"
	"lybbrio/internal/db"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/middleware"
	"lybbrio/internal/viewer"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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

	kc, err := auth.NewHS512KeyContainer("some_secret")
	require.NoError(err)

	jwtProvider, err := auth.NewJWTProvider(
		kc,
		"some_issuer",
		10*time.Second,
		30*time.Second,
	)
	require.NoError(err)

	token, err := jwtProvider.CreateToken(
		auth.NewAccessTokenClaims(
			user.ID.String(),
			user.Username,
			user.Email,
			[]string{permissions.Admin.String()},
		),
	)
	require.NoError(err)

	executed := false
	ts := httptest.NewServer(middleware.ViewerContextMiddleware(jwtProvider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		viewerCtx := viewer.FromContext(r.Context())
		require.NotNil(viewerCtx)

		viewerUserID, ok := viewerCtx.UserID()
		require.True(ok)

		require.Equal(user.ID, viewerUserID)
		require.True(viewerCtx.Has(permissions.Admin))
		require.True(viewerCtx.IsAdmin())
		executed = true
	})))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	require.NoError(err)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	resp, err := ts.Client().Do(req)
	require.NoError(err)
	require.Equal(http.StatusOK, resp.StatusCode)
	require.True(executed)
}
