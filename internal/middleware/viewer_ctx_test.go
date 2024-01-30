package middleware

import (
	"lybbrio/internal/auth"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/viewer"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestViewerContextMiddleware(t *testing.T) {
	require := require.New(t)
	kc, err := auth.NewHS512KeyContainer("secret")
	require.NoError(err)
	provider, err := auth.NewJWTProvider(
		kc,
		"issuer",
		10*time.Second,
		30*time.Second,
	)
	require.NoError(err)

	testUID := "test-uid"
	tests := []struct {
		name       string
		reqFunc    func() *http.Request
		anonymous  bool
		wantStatus int
	}{
		{
			name: "no header",
			reqFunc: func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				return req
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "invalid token",
			reqFunc: func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				req.Header.Add("Authorization", "Bearer invalid-token")
				return req
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "missing token",
			reqFunc: func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				req.Header.Add("Authorization", "Bearer ")
				return req
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "empty header",
			reqFunc: func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				req.Header.Add("Authorization", "")
				return req
			},
			wantStatus: http.StatusUnauthorized,
		},
		{
			name: "valid token",
			reqFunc: func() *http.Request {
				req, _ := http.NewRequest("GET", "/", nil)
				claims := &auth.AccessTokenClaims{
					UserID:      testUID,
					Permissions: []string{"admin", "canedit"},
				}
				token, err := provider.CreateToken(claims)
				require.NoError(err)
				req.Header.Add("Authorization", "Bearer "+token.Token)
				return req
			},
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.NoError(err)

			handler := ViewerContextMiddleware(provider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				v := viewer.FromContext(ctx)
				uid, ok := v.UserID()
				require.True(ok)
				require.Equal(testUID, uid.String())
				require.True(v.Has(permissions.FromString("admin")))
				require.True(v.Has(permissions.FromString("canedit")))
			}))

			req := tt.reqFunc()
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)
			require.Equal(tt.wantStatus, w.Code)
		})
	}
}

func TestViewerContextMiddleware_Anonymous(t *testing.T) {
	require := require.New(t)
	provider, err := auth.NewJWTProvider(
		nil,
		"issuer",
		10*time.Second,
		30*time.Second,
	)
	require.NoError(err)

	handler := ViewerContextMiddleware(provider)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := viewer.FromContext(ctx)
		_, ok := v.UserID()
		require.False(ok)
		require.False(v.Has(permissions.FromString("canedit")))
		require.False(v.Has(permissions.Permission(12)))
		require.False(v.IsAdmin())

	}))

	req, _ := http.NewRequest("GET", "/?anonymous=true", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	require.Equal(http.StatusOK, w.Code)
}
