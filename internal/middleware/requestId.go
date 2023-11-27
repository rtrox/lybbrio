package middleware

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/ksuid"
)

type requestIdCtxKeyType string

const requestIdCtxKey requestIdCtxKeyType = "requestId"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If a request ID is already present in the request, use it.
		reqId := r.Header.Get("X-Request-Id")
		if reqId == "" {
			reqId = ksuid.New().String()
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, requestIdCtxKey, reqId)
		l := log.Ctx(ctx)
		l.UpdateContext(func(c zerolog.Context) zerolog.Context {
			return c.Str("requestId", reqId)
		})
		ctx = l.WithContext(ctx)

		w.Header().Set("X-Request-Id", reqId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequestIDFromCtx(ctx context.Context) string {
	return ctx.Value(requestIdCtxKey).(string)
}
