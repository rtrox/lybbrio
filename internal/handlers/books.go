package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"lybbrio/internal/calibre"
)

func BookRoutes(cal calibre.Calibre) *chi.Mux {
	r := chi.NewRouter()

	r.With(PaginationCtx).Get("/", GetBooks(cal))

	r.Route("/{bookId}", func(r chi.Router) {
		r.Use(BookCtx(cal))
		r.Get("/", GetBook())
	})

	return r
}

type bookCtxKeyType string

const bookCtxKey bookCtxKeyType = "book"

func bookFromContext(ctx context.Context) *calibre.Book {
	return ctx.Value(bookCtxKey).(*calibre.Book)
}

func BookCtx(cal calibre.Calibre) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			s := chi.URLParam(r, "bookId")
			if s == "" {
				render.Render(w, r, ErrNotFound)
				return
			}
			log := log.Ctx(r.Context())
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("bookId", s)
			})
			ctx := log.WithContext(r.Context())
			bookId, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				render.Render(w, r, ErrBadRequest(err))
				return
			}

			book, err := cal.GetBook(bookId)
			if err != nil {
				render.Render(w, r, ErrNotFound)
				return
			}

			ctx = context.WithValue(ctx, bookCtxKey, book)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetBook godoc
// @Summary Get Book by ID
// @Description Get Book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param bookId path int true "Book ID"
// @Success 200 {object} calibre.Book
// @Failure 400 {object} ErrResponse
// @Failure 401 {object} ErrResponse
// @Failure 403 {object} ErrResponse
// @Failure 404 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /books/{bookId} [get]
func GetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book := bookFromContext(r.Context())
		render.JSON(w, r, book)
	}
}

type BookListResponse struct {
	Books []*calibre.Book    `json:"books"`
	Page  PaginationResponse `json:"page"`
}

// GetBooks godoc
// @Summary List Books
// @Description List Books
// @Tags books
// @Accept json
// @Produce json
// @Param cursor query string false "Pagination cursor"
// @Success 200 {object} BookListResponse
// @Failure 400 {object} ErrResponse
// @Failure 401 {object} ErrResponse
// @Failure 403 {object} ErrResponse
// @Failure 404 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /books [get]
func GetBooks(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination := PaginationCtxFromRequest(r)
		books, err := Paginate(cal, pagination.Token).GetBooks()
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrRender, err.Error()}))
			return
		}

		render.JSON(w, r, BookListResponse{
			Books: books,
			Page:  pagination.Response,
		})
	}
}
