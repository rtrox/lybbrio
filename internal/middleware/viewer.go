package middleware

import (
	"lybbrio/internal/auth"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/viewer"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

// Trying to unit test this middleware directly results in an import cycle (auth <-> viewer, or the standard ent cycle).
// To combat this, this file is tested via integration tests, at internal/tests/viewer_context_integ_test.go

func ViewerContextMiddleware(client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			claims := auth.ClaimsFromCtx(ctx)
			adminViewerCtx := viewer.NewSystemAdminContext(ctx)
			user, err := client.User.Query().
				Where(user.ID(ksuid.ID(claims.UserID))).
				First(adminViewerCtx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get user from claims")
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, map[string]string{"error": "Unauthorized"})
				return
			}
			perms, err := user.QueryUserPermissions().First(adminViewerCtx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get user permissions")
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, map[string]string{"error": "Unauthorized"})
				return
			}
			viewerCtx := viewer.NewContext(ctx, user.ID, permissions.From(perms))
			next.ServeHTTP(w, r.WithContext(viewerCtx))
		})
	}
}
