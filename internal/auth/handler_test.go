package auth

import (
	"context"
	"encoding/json"
	"lybbrio/internal/db"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/argon2id"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/viewer"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type handlerTestContext struct {
	client   *ent.Client
	jwt      *JWTProvider
	conf     argon2id.Config
	user     *ent.User
	perms    *ent.UserPermissions
	teardown func()
}

func (h handlerTestContext) Teardown() {
	h.client.Close()
	h.teardown()
}

func setupHandlerTest(t *testing.T, testName string, teardown ...func()) handlerTestContext {
	var ret handlerTestContext
	var err error
	kc, err := NewHS512KeyContainer("testkey")
	require.NoError(t, err)

	ret.client = db.OpenTest(t, testName)

	ret.jwt, err = NewJWTProvider(
		kc,
		"testissuer",
		10*time.Second,
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

	if len(teardown) > 0 {
		ret.teardown = teardown[0]
	} else {
		ret.teardown = func() {}
	}
	return ret
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
			tc := setupHandlerTest(t, tt.name)
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
			requireFunc(res.Header.Get("X-Api-Token"))
			requireFunc(res.Header.Get("X-Api-Expires"))
			requireFunc(res.Header.Get("Set-Cookie"))
		})
	}
}

func Test_PasswordAuth_BadRequest(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	tc := setupHandlerTest(t, "Test_PasswordAuth_BadRequest")
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
	t.Parallel()
	require := require.New(t)
	tc := setupHandlerTest(t, "Test_RefreshToken")
	defer tc.Teardown()

	testMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = viewer.NewContext(
				ctx,
				tc.user.ID,
				permissions.From(tc.perms),
			)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	req := httptest.NewRequest(http.MethodPost, "/", nil)

	w := httptest.NewRecorder()

	handler := testMiddleware(RefreshAuth(tc.client, tc.jwt))
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	require.Equal(http.StatusOK, res.StatusCode)

	require.NotEmpty(res.Header.Get("X-Api-Token"))
	require.NotEmpty(res.Header.Get("X-Api-Expires"))
	require.NotEmpty(res.Header.Get("Set-Cookie"))
}

func Test_AdminContext_CantRefresh(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	tc := setupHandlerTest(t, "Test_AdminContext_CantRefresh")
	defer tc.Teardown()

	testMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = viewer.NewSystemAdminContext(ctx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	req := httptest.NewRequest(http.MethodPost, "/", nil)

	w := httptest.NewRecorder()

	handler := testMiddleware(RefreshAuth(tc.client, tc.jwt))
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	require.Equal(http.StatusUnauthorized, res.StatusCode)
}

func Test_DeletedUser_CantRefresh(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	tc := setupHandlerTest(t, "Test_DeletedUser_CantRefresh")
	defer tc.Teardown()

	testMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = viewer.NewContext(
				ctx,
				tc.user.ID,
				permissions.From(tc.perms),
			)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	adminCtx := viewer.NewSystemAdminContext(context.Background())
	require.NoError(tc.client.User.DeleteOneID(tc.user.ID).Exec(adminCtx))

	req := httptest.NewRequest(http.MethodPost, "/", nil)

	w := httptest.NewRecorder()

	handler := testMiddleware(RefreshAuth(tc.client, tc.jwt))
	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	require.Equal(http.StatusUnauthorized, res.StatusCode)
}
