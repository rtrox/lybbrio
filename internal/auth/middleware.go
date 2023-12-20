package auth

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type claimCtxKeyType string

const claimCtxKey claimCtxKeyType = "claims"

func withClaims(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, claimCtxKey, claims)
}

func ClaimsFromCtx(ctx context.Context) *Claims {
	claims, ok := ctx.Value(claimCtxKey).(*Claims)
	if !ok {
		return nil
	}
	return claims
}

func Middleware(prov *JWTProvider) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := log.Ctx(ctx)

			var token string
			if tokenCookie, err := r.Cookie("token"); err != nil {
				token = r.Header.Get("X-Api-Token")
			} else {
				token = tokenCookie.Value
			}

			if token == "" {
				log.Info().Str("token", token).Msg("Empty Token")
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, map[string]string{"error": "Unauthorized"})
				return
			}

			claims, err := prov.ParseToken(token)
			if err != nil {
				log.Error().Err(err).Msg("Failed to Parse Token")
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, map[string]string{"error": "Unauthorized"})
				return
			}

			ctx = withClaims(ctx, claims)
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("user_id", claims.UserID)
			})
			ctx = log.WithContext(ctx)
			log.Info().Msg("Authenticated")
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
