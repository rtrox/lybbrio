package middleware

import (
	"lybbrio/internal/auth"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/viewer"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func ViewerContextMiddleware(client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			claims := auth.ClaimsFromCtx(ctx)
			adminViewerCtx := viewer.NewSystemAdminContext(ctx)
			user, err := client.User.Query().Where(user.ID(ksuid.ID(claims.UserID))).First(adminViewerCtx)
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
			viewerCtx := viewer.NewContext(ctx, user, perms)
			next.ServeHTTP(w, r.WithContext(viewerCtx))
		})
	}
}