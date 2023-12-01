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

func TagRoutes(cal calibre.Calibre) *chi.Mux {
	r := chi.NewRouter()

	r.With(PaginationCtx).Get("/", GetTags(cal))

	r.Route("/{tagId}", func(r chi.Router) {
		r.Use(TagCtx(cal))
		r.Get("/", GetTag())
		r.With(PaginationCtx).Get("/books", GetTagBooks(cal))
	})

	return r
}

type tagCtxKeyType string

const tagCtxKey tagCtxKeyType = "tag"

func tagFromContext(ctx context.Context) *calibre.Tag {
	return ctx.Value(tagCtxKey).(*calibre.Tag)
}

func TagCtx(cal calibre.Calibre) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			s := chi.URLParam(r, "tagId")
			if s == "" {
				render.Render(w, r, ErrNotFound) //nolint:errcheck
				return
			}
			log := log.Ctx(r.Context())
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("tagId", s)
			})
			ctx := log.WithContext(r.Context())
			tagId, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				render.Render(w, r, ErrBadRequest(err)) //nolint:errcheck
				return
			}

			tag, err := cal.GetTag(ctx, tagId)
			if err != nil {
				render.Render(w, r, ErrNotFound) //nolint:errcheck
				return
			}

			ctx = context.WithValue(ctx, tagCtxKey, tag)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetTag godoc
// @Summary Get a tag
// @Description Get a tag by ID
// @Tags tags
// @Produce json
// @Param tagId path int true "Tag ID"
// @Success 200 {object} calibre.Tag
// @Router /tags/{tagId} [get]
func GetTag() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tag := tagFromContext(r.Context())
		render.JSON(w, r, tag)
	}
}

// GetTagBooks godoc
// @Summary Get books for a tag
// @Description Get books for a tag by ID
// @Tags tags
// @Produce json
// @Param tagId path int true "Tag ID"
// @Success 200 {object} BookListResponse
// @Router /tags/{tagId}/books [get]
func GetTagBooks(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tag := tagFromContext(ctx)
		pagination := PaginationFromCtx(ctx)
		books, err := Paginate(cal, pagination.Token).GetTagBooks(ctx, tag.ID)
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrPaginationToken, err.Error()})) //nolint:errcheck
			return
		}
		render.Render(w, r, BookListResponse{Items: books, Page: &pagination.Response}) //nolint:errcheck
	}
}

// GetTags godoc
// @Summary Get all tags
// @Description Get all tags
// @Tags tags
// @Produce json
// @Success 200 {object} TagListResponse
// @Router /tags [get]
func GetTags(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		pagination := PaginationFromCtx(ctx)
		tags, err := Paginate(cal, pagination.Token).GetTags(ctx)
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrPaginationToken, err.Error()})) //nolint:errcheck
			return
		}
		render.Render(w, r, TagListResponse{Items: tags, Page: &pagination.Response}) //nolint:errcheck
	}
}
