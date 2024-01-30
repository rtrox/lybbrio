package middleware

import (
	"context"
	"lybbrio/internal/auth"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/viewer"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func renderUnauthorized(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusUnauthorized)
	render.JSON(w, r, map[string]string{"error": http.StatusText(http.StatusUnauthorized)})
}

func viewerCtxFromClaims(ctx context.Context, claims *auth.AccessTokenClaims) context.Context {
	perms := permissions.NewPermissions()
	for _, perm := range claims.Permissions {
		perms.Add(permissions.FromString(perm))
	}
	return viewer.NewContext(ctx, ksuid.ID(claims.UserID), perms)
}

func ViewerContextMiddleware(prov *auth.JWTProvider) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := log.Ctx(ctx)

			if r.URL.Query().Get("anonymous") == "true" {
				ctx = viewer.NewAnonymousContext(ctx)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			authHeader := r.Header.Get("Authorization")
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				renderUnauthorized(w, r)
				return
			}
			token := parts[1]

			if token == "" {
				renderUnauthorized(w, r)
				return
			}

			claims := &auth.AccessTokenClaims{}
			err := prov.ParseToken(token, claims)
			if err != nil {
				renderUnauthorized(w, r)
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
