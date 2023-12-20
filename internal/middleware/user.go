package middleware

import (
	"context"
	"lybbrio/internal/auth"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/user"
	"net/http"

	"github.com/go-chi/render"
)

type userCtxKeyType string

const userCtxKey userCtxKeyType = "user"

func UserMiddleware(client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			claims := auth.ClaimsFromCtx(ctx)
			user, err := client.User.Query().Where(user.ID(ksuid.ID(claims.UserID))).First(ctx)
			if err != nil {
				render.Status(r, http.StatusUnauthorized)
				render.JSON(w, r, map[string]string{"error": "Unauthorized"})
				return
			}
			ctx = context.WithValue(ctx, userCtxKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UserFromCtx(ctx context.Context) *ent.User {
	user, ok := ctx.Value(userCtxKey).(*ent.User)
	if !ok {
		return nil
	}
	return user
}
