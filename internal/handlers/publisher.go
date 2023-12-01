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

func PublisherRoutes(cal calibre.Calibre) *chi.Mux {
	r := chi.NewRouter()

	r.With(PaginationCtx).Get("/", GetPublishers(cal))

	r.Route("/{publisherId}", func(r chi.Router) {
		r.Use(PublisherCtx(cal))
		r.Get("/", GetPublisher())
		r.With(PaginationCtx).Get("/books", GetPublisherBooks(cal))
	})

	return r
}

type publisherCtxKeyType string

const publisherCtxKey publisherCtxKeyType = "publisher"

func publisherFromContext(ctx context.Context) *calibre.Publisher {
	return ctx.Value(publisherCtxKey).(*calibre.Publisher)
}

func PublisherCtx(cal calibre.Calibre) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			s := chi.URLParam(r, "publisherId")
			if s == "" {
				render.Render(w, r, ErrNotFound) //nolint:errcheck
				return
			}
			log := log.Ctx(r.Context())
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("publisherId", s)
			})
			ctx := log.WithContext(r.Context())
			publisherId, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				render.Render(w, r, ErrBadRequest(err)) //nolint:errcheck
				return
			}

			publisher, err := cal.GetPublisher(ctx, publisherId)
			if err != nil {
				render.Render(w, r, ErrNotFound) //nolint:errcheck
				return
			}

			ctx = context.WithValue(ctx, publisherCtxKey, publisher)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetPublisher godoc
// @Summary Get a publisher
// @Description Get a publisher by ID
// @Tags publishers
// @Produce json
// @Param publisherId path int true "Publisher ID"
// @Success 200 {object} calibre.Publisher
// @Router /publishers/{publisherId} [get]
func GetPublisher() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		publisher := publisherFromContext(r.Context())
		render.JSON(w, r, publisher)
	}
}

// GetPublisherBooks godoc
// @Summary Get a publisher's books
// @Description Get a publisher's books by ID
// @Tags publishers
// @Produce json
// @Param publisherId path int true "Publisher ID"
// @Success 200 {object} BookListResponse
// @Router /publishers/{publisherId}/books [get]
func GetPublisherBooks(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pagination := PaginationFromCtx(ctx)
		publisher := publisherFromContext(ctx)
		books, err := Paginate(cal, pagination.Token).GetPublisherBooks(ctx, publisher.ID)
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrPublisherBooksDB, err.Error()})) //nolint:errcheck
			return
		}
		render.JSON(w, r, books)
	}
}

// GetPublishers godoc
// @Summary Get all publishers
// @Description Get all publishers
// @Tags publishers
// @Produce json
// @Success 200 {object} PublisherListResponse
// @Router /publishers [get]
func GetPublishers(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pagination := PaginationFromCtx(ctx)
		publishers, err := Paginate(cal, pagination.Token).GetPublishers(ctx)
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrPublishersDB, err.Error()})) //nolint:errcheck
			return
		}

		render.Render(w, r, PublisherListResponse{ //nolint:errcheck
			Items: publishers,
			Page:  &pagination.Response,
		})
	}
}
