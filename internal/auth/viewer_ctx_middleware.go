package auth

import (
	"context"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/viewer"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func viewerCtxFromClaims(ctx context.Context, claims *Claims) context.Context {
	perms := permissions.NewPermissions()
	for _, perm := range claims.Permissions {
		perms.Add(permissions.FromString(perm))
	}
	return viewer.NewContext(ctx, ksuid.ID(claims.UserID), perms)
}

func ViewerContextMiddleware(prov *JWTProvider) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := log.Ctx(ctx)

			if r.URL.Query().Get("anonymous") == "true" {
				ctx = viewer.NewAnonymousContext(ctx)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
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
				log.Error().Err(err).Str("token", token).Msg("Failed to Parse Token")
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, map[string]string{"error": "Unauthorized"})
				return
			}

			ctx = viewerCtxFromClaims(ctx, claims)
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("user_id", claims.UserID)
			})
			ctx = log.WithContext(ctx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
