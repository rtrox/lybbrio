package handler

import (
	"context"
	"encoding/json"
	"lybbrio/internal/auth"
	"lybbrio/internal/db"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/argon2id"
	"lybbrio/internal/viewer"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type authHandlerTestContext struct {
	client       *ent.Client
	jwt          *auth.JWTProvider
	conf         argon2id.Config
	user         *ent.User
	perms        *ent.UserPermissions
	refreshToken auth.SignedToken
	teardown     func()
}

func (h authHandlerTestContext) Teardown() {
	h.client.Close()
	h.teardown()
}

func setupAuthHandlerTest(t *testing.T, testName string, teardown ...func()) authHandlerTestContext {
	var ret authHandlerTestContext
	var err error
	kc, err := auth.NewHS512KeyContainer("testkey")
	require.NoError(t, err)

	ret.client = db.OpenTest(t, testName)

	ret.jwt, err = auth.NewJWTProvider(
		kc,
		"testissuer",
		10*time.Second,
		30*time.Second,
	)
	require.NoError(t, err)

	ret.conf = argon2id.Config{
		Memory:      64,
		Iterations:  1,
		Parallelism: 1,
		SaltLen:     16,
		KeyLen:      32,
	}

	hash, err := argon2id.NewArgon2idHashFromPassword([]byte("notasafepassword"), ret.conf)
	require.NoError(t, err)
	adminCtx := viewer.NewSystemAdminContext(context.Background())
	ret.perms = ret.client.UserPermissions.Create().SetAdmin(true).SaveX(adminCtx)
	ret.user = ret.client.User.Create().
		SetUsername(testName).
		SetEmail(testName + "@notarealemail.com").
		SetPasswordHash(*hash).
		SetUserPermissions(ret.perms).
		SaveX(adminCtx)

	refreshClaims := auth.NewRefreshTokenClaims(ret.user.ID.String())
	ret.refreshToken, err = ret.jwt.CreateToken(refreshClaims)
	require.NoError(t, err)

	if len(teardown) > 0 {
		ret.teardown = teardown[0]
	} else {
		ret.teardown = func() {}
	}
	return ret
}

func makeRequestCookie(c *http.Cookie) string {
	w := httptest.NewRecorder()
	http.SetCookie(w, c)
	return w.Header().Get("Set-Cookie")
}

func Test_PasswordAuth(t *testing.T) {
	tests := []struct {
		name             string
		username         string
		password         string
		setEmptyUsername bool
		wantCode         int
	}{
		{
			name:     "Valid",
			password: "notasafepassword",
			wantCode: http.StatusOK,
		},
		{
			name:     "Invalid Password",
			password: "notasafepassword2",
			wantCode: http.StatusUnauthorized,
		},
		{
			name:     "Invalid Username",
			username: "notarealusername",
			password: "notasafepassword",
			wantCode: http.StatusUnauthorized,
		},
		{
			name:             "Empty Username",
			setEmptyUsername: true,
			password:         "notasafepassword",
			wantCode:         http.StatusBadRequest,
		},
		{
			name:     "Empty Password",
			password: "",
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			tc := setupAuthHandlerTest(t, tt.name)
			defer tc.Teardown()

			username := tt.name
			if tt.username != "" || tt.setEmptyUsername {
				username = tt.username
			}
			data := PasswordRequest{
				Username: username,
				Password: tt.password,
			}
			jsonData, err := json.Marshal(data)
			require.NoError(err)
			jsonReader := strings.NewReader(string(jsonData))

			req := httptest.NewRequest(http.MethodPost, "/", jsonReader)
			w := httptest.NewRecorder()

			handler := PasswordAuth(tc.client, tc.jwt, tc.conf)
			handler.ServeHTTP(w, req)
			res := w.Result()
			defer res.Body.Close()

			require.Equal(tt.wantCode, res.StatusCode)

			requireFunc := require.Empty
			if tt.wantCode == http.StatusOK {
				requireFunc = require.NotEmpty
			}
			requireFunc(res.Header.Get(ACCESS_TOKEN_HEADER))
			requireFunc(res.Header.Get(ACCESS_TOKEN_EXPIRATION_HEADER))
			requireFunc(res.Header.Get("Set-Cookie"))
		})
	}
}

func Test_PasswordAuth_BadRequest(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	tc := setupAuthHandlerTest(t, "Test_PasswordAuth_BadRequest")
	defer tc.Teardown()

	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	w := httptest.NewRecorder()

	handler := PasswordAuth(tc.client, tc.jwt, tc.conf)
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	require.Equal(http.StatusBadRequest, res.StatusCode)
}

func Test_RefreshToken(t *testing.T) {
	tests := []struct {
		name         string
		reqFunc      func(authHandlerTestContext) *http.Request
		expectedCode int
	}{
		{
			name: "Cookie",
			reqFunc: func(tc authHandlerTestContext) *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/", nil)
				req.AddCookie(&http.Cookie{
					Name:     "refreshToken",
					Value:    tc.refreshToken.Token,
					Expires:  tc.refreshToken.ExpiresAt,
					HttpOnly: true,
					Secure:   true,
					Path:     "/",
				})
				return req
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "Header",
			reqFunc: func(tc authHandlerTestContext) *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/", nil)
				req.Header.Set("X-Refresh-Token", tc.refreshToken.Token)
				return req
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "No Token",
			reqFunc: func(tc authHandlerTestContext) *http.Request {
				return httptest.NewRequest(http.MethodPost, "/", nil)
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Invalid Token",
			reqFunc: func(tc authHandlerTestContext) *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/", nil)
				req.Header.Set("X-Refresh-Token", "invalidtoken")
				return req
			},
			expectedCode: http.StatusUnauthorized,
		},
		{
			name: "Access Token as Refresh Token",
			reqFunc: func(tc authHandlerTestContext) *http.Request {
				claims := auth.NewAccessTokenClaims(tc.user.ID.String(), tc.user.Username, tc.user.Email, []string{"admin"})
				token, err := tc.jwt.CreateToken(claims)
				require.NoError(t, err)
				req := httptest.NewRequest(http.MethodPost, "/", nil)
				req.Header.Set("X-Refresh-Token", token.Token)
				return req
			},
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			tc := setupAuthHandlerTest(t, tt.name)
			defer tc.Teardown()

			w := httptest.NewRecorder()

			handler := RefreshAuth(tc.client, tc.jwt)
			handler.ServeHTTP(w, tt.reqFunc(tc))
			res := w.Result()
			defer res.Body.Close()

			require.Equal(tt.expectedCode, res.StatusCode)

			requireFunc := require.Empty
			if tt.expectedCode == http.StatusOK {
				requireFunc = require.NotEmpty
			}
			if tt.expectedCode == http.StatusOK {
				requireFunc(res.Header.Get(ACCESS_TOKEN_HEADER))
				requireFunc(res.Header.Get(ACCESS_TOKEN_EXPIRATION_HEADER))
			}
		})
	}
}

func Test_DeletedUser_CantRefresh(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	tc := setupAuthHandlerTest(t, "Test_DeletedUser_CantRefresh")
	defer tc.Teardown()

	adminCtx := viewer.NewSystemAdminContext(context.Background())
	require.NoError(tc.client.User.DeleteOneID(tc.user.ID).Exec(adminCtx))

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("X-Refresh-Token", tc.refreshToken.Token)

	w := httptest.NewRecorder()

	handler := RefreshAuth(tc.client, tc.jwt)
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	require.Equal(http.StatusUnauthorized, res.StatusCode)
}
