package handler

import (
	"lybbrio/internal/auth"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/argon2id"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/viewer"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const (
	ACCESS_TOKEN_HEADER            = "X-JWT"
	ACCESS_TOKEN_EXPIRATION_HEADER = "X-JWT-Expires"
)

type PasswordRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User        *ent.User        `json:"user"`
	AccessToken auth.SignedToken `json:"accessToken"`
}

func accessTokenClaimsFromUser(user *ent.User) *auth.AccessTokenClaims {
	return &auth.AccessTokenClaims{
		UserID:      user.ID.String(),
		UserName:    user.Username,
		Email:       user.Email,
		Permissions: permissions.From(user.Edges.UserPermissions).StringSlice(),
	}
}

func AuthRoutes(client *ent.Client, jwt *auth.JWTProvider, conf argon2id.Config) http.Handler {
	r := chi.NewRouter()
	r.Post("/password", PasswordAuth(client, jwt, conf))
	r.Get("/refresh", RefreshAuth(client, jwt))
	return r
}

func PasswordAuth(client *ent.Client, jwt *auth.JWTProvider, conf argon2id.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := &PasswordRequest{}
		if err := render.DecodeJSON(r.Body, data); err != nil {
			statusCodeResponse(w, r, http.StatusBadRequest)
			return
		}

		if data.Username == "" || data.Password == "" {
			statusCodeResponse(w, r, http.StatusBadRequest)
			return
		}

		adminCtx := viewer.NewSystemAdminContext(ctx)
		user, err := client.User.Query().
			Where(user.Username(data.Username)).
			WithUserPermissions().
			Only(adminCtx)
		if err != nil {
			if ent.IsNotFound(err) {
				statusCodeResponse(w, r, http.StatusUnauthorized)
				return
			}
			statusCodeResponse(w, r, http.StatusInternalServerError)
		}

		if err := user.PasswordHash.Verify([]byte(data.Password)); err != nil {
			statusCodeResponse(w, r, http.StatusUnauthorized)
			return
		}

		refreshClaims := &auth.AccessTokenClaims{
			UserID: user.ID.String(),
		}
		accessClaims := accessTokenClaimsFromUser(user)

		refreshToken, err := jwt.CreateToken(refreshClaims)
		if err != nil {
			statusCodeResponse(w, r, http.StatusInternalServerError)
			return
		}
		accessToken, err := jwt.CreateToken(accessClaims)
		if err != nil {
			statusCodeResponse(w, r, http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "refreshToken",
			Value:    refreshToken.Token,
			Expires:  refreshToken.ExpiresAt,
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		})

		w.Header().Add(ACCESS_TOKEN_HEADER, accessToken.Token)
		w.Header().Add(ACCESS_TOKEN_EXPIRATION_HEADER, accessToken.ExpiresAt.Format(time.RFC3339))
		render.Status(r, http.StatusOK)
		render.JSON(w, r, AuthResponse{
			AccessToken: accessToken,
			User:        user,
		})
	}
}

func RefreshAuth(client *ent.Client, jwt *auth.JWTProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var token string
		if tokenCookie, err := r.Cookie("refreshToken"); err != nil || tokenCookie == nil {
			token = r.Header.Get("X-Refresh-Token")
		} else {
			token = tokenCookie.Value
		}

		if token == "" {
			statusCodeResponse(w, r, http.StatusUnauthorized)
			return
		}

		claims := &auth.RefreshTokenClaims{}
		if err := jwt.ParseToken(token, claims); err != nil {
			statusCodeResponse(w, r, http.StatusUnauthorized)
			return
		}

		adminCtx := viewer.NewSystemAdminContext(ctx)
		user, err := client.User.Query().
			Where(user.ID(ksuid.ID(claims.UserID))).
			WithUserPermissions().
			Only(adminCtx)

		if err != nil {
			statusCodeResponse(w, r, http.StatusUnauthorized)
			return
		}

		accessClaims := accessTokenClaimsFromUser(user)
		accessToken, err := jwt.CreateToken(accessClaims)
		if err != nil {
			statusCodeResponse(w, r, http.StatusInternalServerError)
			return
		}

		w.Header().Add(
			ACCESS_TOKEN_HEADER,
			accessToken.Token,
		)
		w.Header().Add(
			ACCESS_TOKEN_EXPIRATION_HEADER,
			accessToken.ExpiresAt.Format(time.RFC3339),
		)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, AuthResponse{
			AccessToken: accessToken,
			User:        user,
		})
	}
}
