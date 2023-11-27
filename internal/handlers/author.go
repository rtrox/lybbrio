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

func AuthorRoutes(cal calibre.Calibre) *chi.Mux {
	r := chi.NewRouter()

	r.With(PaginationCtx).Get("/", GetAuthors(cal))
	r.Route("/{authorId}", func(r chi.Router) {
		r.Use(AuthorCtx(cal))
		r.Get("/", GetAuthor())
		r.With(PaginationCtx).Get("/books", GetAuthorBooks(cal))
		r.With(PaginationCtx).Get("/series", GetAuthorSeries(cal))
	})

	return r
}

type authorCtxKeyType string

const authorCtxKey authorCtxKeyType = "author"

func authorFromContext(ctx context.Context) *calibre.Author {
	return ctx.Value(authorCtxKey).(*calibre.Author)
}

func AuthorCtx(cal calibre.Calibre) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			s := chi.URLParam(r, "authorId")
			if s == "" {
				render.Render(w, r, ErrNotFound)
				return
			}
			log := log.Ctx(r.Context())
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("authorId", s)
			})
			ctx := log.WithContext(r.Context())
			authorId, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				render.Render(w, r, ErrBadRequest(err))
				return
			}

			author, err := cal.GetAuthor(authorId)
			if err != nil {
				render.Render(w, r, ErrNotFound)
				return
			}
			ctx = context.WithValue(ctx, authorCtxKey, author)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetAuthor godoc
// @Summary Get an author
// @Description Get an author by ID
// @Tags authors
// @Accept json
// @Produce json
// @Param authorId path int true "Author ID"
// @Success 200 {object} calibre.Author
// @Failure 400 {object} ErrResponse
// @Failure 401 {object} ErrResponse
// @Failure 403 {object} ErrResponse
// @Failure 404 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /authors/{authorId} [get]
func GetAuthor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := authorFromContext(r.Context())
		render.JSON(w, r, author)
	}
}

// GetAuthorBooks godoc
// @Summary Get an author's books
// @Description Get an author's books by ID
// @Tags authors
// @Accept json
// @Produce json
// @Param authorId path int true "Author ID"
// @Success 200 {array} Book
// @Failure 400 {object} ErrResponse
// @Failure 401 {object} ErrResponse
// @Failure 403 {object} ErrResponse
// @Failure 404 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /authors/{authorId}/books [get]
func GetAuthorBooks(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := authorFromContext(r.Context())
		pagination := PaginationCtxFromRequest(r)
		books, err := Paginate(cal, pagination.Token).GetAuthorBooks(author.ID)
		if err != nil {
			render.Render(w, r, ErrInternalError(
				AppError{ErrAuthorBooksDB, err.Error()},
			))
			return
		}
		render.JSON(w, r, BookListResponse{
			Books: books,
			Page:  pagination.Response,
		})
	}
}

// TODO: move to series handler once created
type SeriesListResponse struct {
	Series []*calibre.Series `json:"series"`
	Page   PaginationResponse
}

func GetAuthorSeries(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := authorFromContext(r.Context())
		pagination := PaginationCtxFromRequest(r)
		series, err := Paginate(cal, pagination.Token).GetAuthorSeries(author.ID)
		if err != nil {
			render.Render(w, r, ErrInternalError(
				AppError{ErrAuthorBooksDB, err.Error()},
			))
			return
		}
		render.JSON(w, r, SeriesListResponse{
			Series: series,
			Page:   pagination.Response,
		})
	}
}

type AuthorListResponse struct {
	Authors []*calibre.Author `json:"authors"`
	Page    PaginationResponse
}

// GetAuthors godoc
// @Summary List Authors
// @Description List Authors
// @Tags authors
// @Accept json
// @Produce json
// @Param cursor query string false "Pagination cursor"
// @Success 200 {object} AuthorListResponse
// @Failure 400 {object} ErrResponse
// @Failure 401 {object} ErrResponse
// @Failure 403 {object} ErrResponse
// @Failure 404 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /authors [get]
func GetAuthors(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination := PaginationCtxFromRequest(r)
		authors, err := Paginate(cal, pagination.Token).GetAuthors()
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrRender, err.Error()}))
			return
		}

		render.JSON(w, r, AuthorListResponse{
			Authors: authors,
			Page:    pagination.Response,
		})
	}
}
