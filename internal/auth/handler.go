package auth

import (
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/user"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func Routes(client *ent.Client, jwt *JWTProvider) http.Handler {
	r := chi.NewRouter()
	r.Get("/testAuthDONOTUSE", TestAuth(client, jwt))
	return r
}

func TestAuth(client *ent.Client, jwt *JWTProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var staticUser *ent.User
		var err error
		staticUser, err = client.User.Query().Where(user.Username("test")).First(ctx)
		if err != nil {
			staticUser, err = client.User.Create().
				SetUsername("test").
				SetPasswordHash("asdf").
				SetEmail("not@arealemail.com").
				Save(ctx)
			if err != nil {
				log.Error().Err(err).Msg("Failed to create static user")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		token, err := jwt.CreateToken(string(staticUser.ID), staticUser.Username)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create token")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token.String(),
			Expires: token.Claims().ExpiresAt.Time,
		})
		w.Header().Add("X-Api-Token", token.String())
		w.Header().Add("X-Api-Expires", token.Claims().ExpiresAt.Format(time.RFC3339))
		render.Status(r, http.StatusOK)
		render.JSON(w, r, staticUser)
	}
}
