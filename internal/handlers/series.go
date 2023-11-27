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

func SeriesRoutes(cal calibre.Calibre) *chi.Mux {
	r := chi.NewRouter()

	r.With(PaginationCtx).Get("/", GetSeriesList(cal))

	r.Route("/{seriesId}", func(r chi.Router) {
		r.Use(SeriesCtx(cal))
		r.Get("/", GetSeries())
		r.With(PaginationCtx).Get("/books", GetSeriesBooks(cal))
	})

	return r
}

type seriesCtxKeyType string

const seriesCtxKey seriesCtxKeyType = "series"

func seriesFromContext(ctx context.Context) *calibre.Series {
	return ctx.Value(seriesCtxKey).(*calibre.Series)
}

func SeriesCtx(cal calibre.Calibre) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			s := chi.URLParam(r, "seriesId")
			if s == "" {
				render.Render(w, r, ErrNotFound)
				return
			}
			log := log.Ctx(r.Context())
			log.UpdateContext(func(c zerolog.Context) zerolog.Context {
				return c.Str("seriesId", s)
			})
			ctx := log.WithContext(r.Context())
			seriesId, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				render.Render(w, r, ErrBadRequest(err))
				return
			}

			series, err := cal.GetSeries(seriesId)
			if err != nil {
				render.Render(w, r, ErrNotFound)
				return
			}

			ctx = context.WithValue(ctx, seriesCtxKey, series)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetSeries godoc
// @Summary Get a series
// @Description Get a series by ID
// @Tags series
// @Produce json
// @Param seriesId path int true "Series ID"
// @Success 200 {object} calibre.Series
// @Router /series/{seriesId} [get]
func GetSeries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		series := seriesFromContext(r.Context())
		render.JSON(w, r, series)
	}
}

// GetSeriesBooks godoc
// @Summary Get a series' books
// @Description Get a series' books by ID
// @Tags series
// @Produce json
// @Param seriesId path int true "Series ID"
// @Success 200 {array} Book
// @Router /series/{seriesId}/books [get]
func GetSeriesBooks(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		series := seriesFromContext(r.Context())
		books, err := cal.GetSeriesBooks(series.ID)
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrSeriesBooksDB, err.Error()}))
		}
		render.JSON(w, r, BookListResponse{Books: books, Page: PaginationCtxFromRequest(r).Response})
	}
}

type SeriesListResponse struct {
	Series []*calibre.Series   `json:"series"`
	Page   *PaginationResponse `json:"page,omitempty"`
}

// GetSerieses godoc
// @Summary Get all series
// @Description Get all series
// @Tags series
// @Produce json
// @Success 200 {array} Series
// @Router /series [get]
func GetSeriesList(cal calibre.Calibre) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pagination := PaginationCtxFromRequest(r)
		series, err := Paginate(cal, pagination.Token).GetSeriesList()
		if err != nil {
			render.Render(w, r, ErrInternalError(AppError{ErrSeriesDB, err.Error()}))
			return
		}
		ret := SeriesListResponse{
			Series: series,
		}
		if len(series) == pagination.Token.PageSize {
			ret.Page = &pagination.Response
		}
		render.JSON(w, r, ret)
	}
}
