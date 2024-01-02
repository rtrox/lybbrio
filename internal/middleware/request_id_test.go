package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_RequestID(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	var idFromCtx string
	var idFromHeader string
	handler := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idFromCtx = RequestIDFromCtx(r.Context())
		w.WriteHeader(http.StatusOK)
	}))
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	idFromHeader = w.Header().Get("X-Request-Id")
	require.NotEmpty(idFromCtx)
	require.NotEmpty(idFromHeader)
	require.Equal(idFromCtx, idFromHeader)
}

func Test_RequestID_ExistingHeader(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	var idFromCtx string
	var idFromHeader string
	handler := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idFromCtx = RequestIDFromCtx(r.Context())
		w.WriteHeader(http.StatusOK)
	}))
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-Request-Id", "test")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	idFromHeader = w.Header().Get("X-Request-Id")
	require.Equal("test", idFromHeader)
	require.Equal("test", idFromCtx)
}
