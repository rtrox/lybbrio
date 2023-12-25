package middleware

import (
	"context"
	"lybbrio/internal/viewer"
	"net/http"
)

type SuperReadCtxKeyType string

const SuperReadCtxKey SuperReadCtxKeyType = "superRead"

func SuperRead(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		view := viewer.FromContext(ctx)
		if view.IsAdmin() && r.URL.Query().Get("superRead") == "true" {
			ctx = WithSuperRead(ctx)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SuperReadFromCtx(ctx context.Context) bool {
	sr, ok := ctx.Value(SuperReadCtxKey).(bool)
	if !ok {
		return false
	}
	return sr
}

func WithSuperRead(ctx context.Context) context.Context {
	return context.WithValue(ctx, SuperReadCtxKey, true)
}
