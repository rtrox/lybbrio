package test

import (
	"context"
	"fmt"
	"lybbrio/internal/auth"
	"lybbrio/internal/db"
	"lybbrio/internal/ent/schema/argon2id"
	"lybbrio/internal/handler"
	"lybbrio/internal/middleware"
	"lybbrio/internal/viewer"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/publicsuffix"
)

func TestAuthWorkflow(t *testing.T) {
	require := require.New(t)
	testUsername := "testUser"
	testPassword := "notasafepassword"

	testClient := db.OpenTest(t, "TestAuthWorkflow")
	defer testClient.Close()

	kc, err := auth.NewHS512KeyContainer("testkey")
	require.NoError(err)

	testJWT, err := auth.NewJWTProvider(
		kc,
		"testissuer",
		10*time.Second,
		30*time.Second,
	)
	require.NoError(err)

	testConfig := argon2id.Config{
		Memory:      64,
		Iterations:  1,
		Parallelism: 1,
		SaltLen:     16,
		KeyLen:      32,
	}

	r := chi.NewRouter()
	r.Mount("/", handler.AuthRoutes(
		testClient,
		testJWT,
		testConfig,
	))
	r.With(middleware.ViewerContextMiddleware(testJWT)).Get("/test", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := viewer.FromContext(ctx)
		_, ok := v.UserID()
		require.True(ok)
		render.JSON(w, r, v)
	})

	hash, err := argon2id.NewArgon2idHashFromPassword([]byte(testPassword), testConfig)
	require.NoError(err)

	adminCtx := viewer.NewSystemAdminContext(context.Background())
	perms := testClient.UserPermissions.Create().SetAdmin(true).SaveX(adminCtx)
	testClient.User.Create().
		SetUsername(testUsername).
		SetEmail(testUsername + "@notarealemail.com").
		SetPasswordHash(*hash).
		SetUserPermissions(perms).
		SaveX(adminCtx)

	ts := httptest.NewTLSServer(r)
	defer ts.Close()

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	require.NoError(err)
	client := ts.Client()
	client.Jar = jar

	data := fmt.Sprintf("{\"username\": \"%s\", \"password\": \"%s\"}", testUsername, testPassword)
	req, err := http.NewRequest("POST", ts.URL+"/password", strings.NewReader(data))
	require.NoError(err)

	resp, err := client.Do(req)
	require.NoError(err)
	require.Equal(http.StatusOK, resp.StatusCode)

	var authResp handler.AuthResponse
	err = render.DecodeJSON(resp.Body, &authResp)
	require.NoError(err)
	require.NotNil(authResp.User)
	require.NotNil(authResp.AccessToken)

	auth.TimeFunc = func() time.Time {
		return time.Now().Add(11 * time.Second)
	}
	defer func() {
		auth.TimeFunc = time.Now
	}()

	refreshReq, err := http.NewRequest("GET", ts.URL+"/refresh", nil)
	require.NoError(err)

	refreshReq.Header.Add("Authorization", "Bearer "+authResp.AccessToken.Token)
	refreshResp, err := client.Do(refreshReq)
	require.NoError(err)
	require.Equal(http.StatusOK, refreshResp.StatusCode)

	var refreshAuthResp handler.AuthResponse
	err = render.DecodeJSON(refreshResp.Body, &refreshAuthResp)
	require.NoError(err)
	require.NotNil(refreshAuthResp.User)
	require.NotNil(refreshAuthResp.AccessToken)
	require.NotEqual(authResp.AccessToken.Token, refreshAuthResp.AccessToken.Token)

	testReq, err := http.NewRequest("GET", ts.URL+"/test", nil)
	require.NoError(err)
	testReq.Header.Add("Authorization", "Bearer "+refreshAuthResp.AccessToken.Token)

	testResp, err := client.Do(testReq)
	require.NoError(err)
	require.Equal(http.StatusOK, testResp.StatusCode)
}
