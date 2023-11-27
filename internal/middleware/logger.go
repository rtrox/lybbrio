package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	chi_middleware "github.com/go-chi/chi/middleware"
)

// Adapted from https://learninggolang.com/it5-gin-structured-logging.html

// DefaultStructuredLogger logs a gin HTTP request in JSON format. Uses the
// default logger from rs/zerolog.
func DefaultStructuredLogger(excluded_paths ...string) func(http.Handler) http.Handler {
	return StructuredLogger(&log.Logger, excluded_paths...)
}

func StructuredLogger(logger *zerolog.Logger, excluded_paths ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		excluded_path_map := make(map[string]struct{})
		for _, path := range excluded_paths {
			excluded_path_map[path] = struct{}{}
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now() // Start timer

			// attach logger to context for handler logging
			ctx := logger.WithContext(r.Context())

			ww := chi_middleware.NewWrapResponseWriter(w, r.ProtoMajor) // Wrap response writer
			defer func() {
				if _, ok := excluded_path_map[r.URL.Path]; ok {
					return
				}
				logger := log.Ctx(ctx)

				path := r.URL.Path
				raw := r.URL.RawQuery
				if raw != "" {
					path = path + "?" + raw
				}

				var logEvent *zerolog.Event
				if ww.Status() >= 500 {
					logEvent = logger.Error()
				} else {
					logEvent = logger.Info()
				}

				logEvent.
					Str("request_id", RequestIDFromCtx(ctx)).
					Str("client_id", r.RemoteAddr).
					Str("method", r.Method).
					Int("status_code", ww.Status()).
					Int("bytes_written", ww.BytesWritten()).
					Str("path", path).
					Str("latency", time.Since(start).String()).
					Msg("")
			}()

			next.ServeHTTP(ww, r.WithContext(ctx)) // Process request
		})
	}
}
