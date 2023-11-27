package handlers

import (
	"lybbrio/internal/calibre"
	"net/http"

	"github.com/go-chi/render"
)

type AuthorListResponse = ListResponse[*calibre.Author]
type BookListResponse = ListResponse[*calibre.Book]
type SeriesListResponse = ListResponse[*calibre.Series]
type TagListResponse = ListResponse[*calibre.Tag]
type PublisherListResponse = ListResponse[*calibre.Publisher]

type ListResponse[T any] struct {
	Items []T                 `json:"items"`
	Page  *PaginationResponse `json:"page,omitempty"`
}

func (l ListResponse[T]) Render(w http.ResponseWriter, r *http.Request) error {
	if len(l.Items) < l.Page.currentToken.PageSize {
		l.Page = nil
	}
	render.JSON(w, r, l)
	return nil
}
