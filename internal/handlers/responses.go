package handlers

import (
	"lybbrio/internal/calibre"
	"net/http"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

type AuthorListResponse = ListResponse[*calibre.Author]
type BookListResponse = ListResponse[*calibre.Book]
type SeriesListResponse = ListResponse[*calibre.Series]
type TagListResponse = ListResponse[*calibre.Tag]
type PublisherListResponse = ListResponse[*calibre.Publisher]
type LanguageListResponse = ListResponse[*calibre.Language]

type ErrStopRender struct{}

func (e ErrStopRender) Error() string {
	return "Placeholder error to prevent struct recursion in Render()."
}

type ListResponse[T any] struct {
	Items []T                 `json:"items"`
	Page  *PaginationResponse `json:"page,omitempty"`
}

func (l ListResponse[T]) Render(w http.ResponseWriter, r *http.Request) error {
	log.Info().Int("len", len(l.Items)).Int("pageSize", l.Page.currentToken.PageSize).Msg("rendering list")
	if len(l.Items) < l.Page.currentToken.PageSize {
		l.Page = nil
	}
	log.Info().Interface("page", l.Page).Msg("rendering list")
	render.JSON(w, r, l)
	return ErrStopRender{}
}
