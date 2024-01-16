package auth

import (
	"fmt"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/argon2id"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/viewer"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func Routes(client *ent.Client, jwt *JWTProvider, conf argon2id.Config) http.Handler {
	r := chi.NewRouter()
	r.Get("/testAuthDONOTUSE/{username}", TestAuthDONOTUSE(client, jwt))
	r.Post("/login", PasswordAuth(client, jwt, conf))
	return r
}

type PasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func cookieFromToken(token SignedToken) *http.Cookie {
	return &http.Cookie{
		Name:    "token",
		Value:   token.String(),
		Expires: token.Claims().ExpiresAt.Time,
	}
}

func TestAuthDONOTUSE(client *ent.Client, jwt *JWTProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var staticUser *ent.User
		var err error
		username := chi.URLParam(r, "username")
		adminViewerCtx := viewer.NewSystemAdminContext(ctx)
		staticUser, err = client.User.Query().Where(user.Username(username)).First(adminViewerCtx)
		if err != nil {
			permsCreate := client.UserPermissions.Create()
			if r.URL.Query().Get("admin") == "true" {
				permsCreate.SetAdmin(true)
			}
			perms, err := permsCreate.Save(adminViewerCtx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to create static user permissions")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			staticUser, err = client.User.Create().
				SetUsername(username).
				SetEmail(fmt.Sprintf("%s@notarealemail.com", username)).
				SetUserPermissions(perms).
				Save(adminViewerCtx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to create static user")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		perms := staticUser.QueryUserPermissions().FirstX(adminViewerCtx)
		token, err := jwt.CreateToken(
			staticUser.ID.String(),
			staticUser.Username,
			permissions.From(perms).StringSlice(),
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create token")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, cookieFromToken(token))
		w.Header().Add("X-Api-Token", token.String())
		w.Header().Add("X-Api-Expires", token.Claims().ExpiresAt.Format(time.RFC3339))
		render.Status(r, http.StatusOK)
		render.JSON(w, r, staticUser)
	}
}

func PasswordAuth(client *ent.Client, jwt *JWTProvider, conf argon2id.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := &PasswordRequest{}
		if err := render.DecodeJSON(r.Body, data); err != nil {
			s := http.StatusBadRequest
			render.Status(r, s)
			render.DefaultResponder(w, r, render.M{"error": http.StatusText(s)})
			return
		}
		if data.Username == "" || data.Password == "" {
			s := http.StatusBadRequest
			render.Status(r, s)
			render.DefaultResponder(w, r, render.M{"error": http.StatusText(s)})
			return
		}
		adminCtx := viewer.NewSystemAdminContext(ctx)
		user, err := client.User.Query().
			Where(user.Username(data.Username)).
			First(adminCtx)
		if err != nil {
			s := http.StatusUnauthorized
			render.Status(r, s)
			render.DefaultResponder(w, r, render.M{"error": http.StatusText(s)})
			return
		}
		if err := user.PasswordHash.Verify([]byte(data.Password)); err != nil {
			s := http.StatusUnauthorized
			render.Status(r, s)
			render.DefaultResponder(w, r, render.M{"error": http.StatusText(s)})
			return
		}
		perms, err := user.QueryUserPermissions().First(adminCtx)
		if err != nil {
			s := http.StatusInternalServerError
			render.Status(r, s)
			render.DefaultResponder(w, r, render.M{"error": http.StatusText(s)})
			return
		}
		token, err := jwt.CreateToken(
			user.ID.String(),
			user.Username,
			permissions.From(perms).StringSlice(),
		)
		if err != nil {
			s := http.StatusInternalServerError
			render.Status(r, s)
			render.DefaultResponder(w, r, render.M{"error": http.StatusText(s)})
			return
		}
		http.SetCookie(w, cookieFromToken(token))
		w.Header().Add("X-Api-Token", token.String())
		w.Header().Add("X-Api-Expires", token.Claims().ExpiresAt.Format(time.RFC3339))
		render.Status(r, http.StatusOK)
		render.JSON(w, r, user)
	}
}
