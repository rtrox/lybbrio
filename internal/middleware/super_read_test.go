package middleware

import (
	"lybbrio/internal/viewer"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SuperRead(t *testing.T) {
	tests := []struct {
		name    string
		reqFunc func() *http.Request
		want    bool
	}{
		{
			name: "superRead with Admin",
			reqFunc: func() *http.Request {
				req := httptest.NewRequest("GET", "/?superRead=true", nil)
				req = req.WithContext(
					viewer.NewSystemAdminContext(req.Context()),
				)
				return req
			},
			want: true,
		},
		{
			name: "superRead without Admin",
			reqFunc: func() *http.Request {
				req := httptest.NewRequest("GET", "/?superRead=true", nil)
				req = req.WithContext(
					viewer.NewContext(req.Context(), "", nil),
				)
				return req
			},
			want: false,
		},
		{
			name: "no superRead with Admin",
			reqFunc: func() *http.Request {
				req := httptest.NewRequest("GET", "/", nil)
				req = req.WithContext(
					viewer.NewSystemAdminContext(req.Context()),
				)
				return req
			},
			want: false,
		},
		{
			name: "no superRead without Admin",
			reqFunc: func() *http.Request {
				req := httptest.NewRequest("GET", "/", nil)
				req = req.WithContext(
					viewer.NewContext(req.Context(), "", nil),
				)
				return req
			},
			want: false,
		},
		{
			name: "no super read no viewer context",
			reqFunc: func() *http.Request {
				req := httptest.NewRequest("GET", "/", nil)
				return req
			},
			want: false,
		},
		{
			name: "super read no viewer context",
			reqFunc: func() *http.Request {
				req := httptest.NewRequest("GET", "/?superRead=true", nil)
				return req
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			var got bool
			handler := SuperRead(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				got = SuperReadFromCtx(r.Context())
				w.WriteHeader(http.StatusOK)
			}))

			req := tt.reqFunc()
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			require.Equal(tt.want, got)
		})
	}
}
