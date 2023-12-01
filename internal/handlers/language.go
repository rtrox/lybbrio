package handlers

import (
	"context"
	"lybbrio/internal/calibre"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LanguageRoutes(cal calibre.Calibre) *chi.Mux {
	r := chi.NewRouter()

	r.With(PaginationCtx).Get("/", GetLanguages(cal))

	r.Route("/{languageId}", func(r chi.Router) {
		r.Use(LanguageCtx(cal))
		r.Get("/", GetLanguage())
		r.With(PaginationCtx).Get("/books", GetLanguageBooks(cal))
	})

	return r
}

type languageCtxKey string

const languageCtxKeyKey languageCtxKey = "language"

func languageFromContext(ctx context.Context) *calibre.Language {
	return ctx.Value(languageCtxKeyKey).(*calibre.Language)
}

func LanguageCtx(cal calibre.Calibre) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			s := chi.URLParam(r, "languageId")
			if s == "" {
				render.Render(w, r, ErrNotFound) //nolint:errcheck
				return
			}
			log := log.Ctx(r.Context())
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("languageId", s)
			})
			ctx := log.WithContext(r.Context())
			languageId, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				render.Render(w, r, ErrBadRequest(err)) //nolint:errcheck
				return
			}

			language, err := cal.GetLanguage(ctx, languageId)
			if err != nil {
				render.Render(w, r, ErrNotFound) //nolint:errcheck
				return
			}
			ctx = context.WithValue(ctx, languageCtxKeyKey, language)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetLanguage godoc
// @Summary Get a language
// @Description Get a language by ID
// @Tags Languages
// @Accept  json
// @Produce  json
// @Param languageId path int true "Language ID"
// @Success 200 {object} calibre.Language
// @Router /languages/{languageId} [get]
func GetLanguage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		language := languageFromContext(r.Context())
		render.JSON(w, r, language)
	}
}

// GetLanguageBooks godoc
// @Summary Get books for a language
// @Description Get books for a language by ID
// @Tags Languages
// @Accept  json
// @Produce  json
// @Param languageId path int true "Language ID"
// @Success 200 {object} BookListResponse
// @Router /languages/{languageId}/books [get]
func GetLanguageBooks(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		language := languageFromContext(ctx)
		pagination := PaginationFromCtx(ctx)
		books, err := Paginate(cal, pagination.Token).GetLanguageBooks(ctx, language.ID)
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrLanguageBooksDB, err.Error()})) //nolint:errcheck
			return
		}
		render.Render(w, r, BookListResponse{Items: books, Page: &pagination.Response}) //nolint:errcheck
	}
}

// GetLanguages godoc
// @Summary Get all languages
// @Description Get all languages
// @Tags Languages
// @Accept  json
// @Produce  json
// @Success 200 {object} LanguageListResponse
// @Router /languages [get]
func GetLanguages(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pagination := PaginationFromCtx(ctx)
		languages, err := Paginate(cal, pagination.Token).GetLanguages(ctx)
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrLanguagesDB, err.Error()})) //nolint:errcheck
			return
		}
		render.Render(w, r, LanguageListResponse{Items: languages, Page: &pagination.Response}) //nolint:errcheck
	}
}
