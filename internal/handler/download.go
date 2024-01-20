package handler

import (
	"errors"
	"fmt"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/bookfile"
	"lybbrio/internal/ent/schema/ksuid"
	"mime"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func DownloadRoutes(client *ent.Client) http.Handler {
	r := chi.NewRouter()
	r.Get("/{bookID}/{book_format}", DownloadHandler(client))
	return r
}

func DownloadHandler(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		bookID := chi.URLParam(r, "bookID")
		bookFormat := chi.URLParam(r, "book_format")
		if bookID == "" || bookFormat == "" {
			status := http.StatusBadRequest
			w.WriteHeader(status)
			render.JSON(w, r, map[string]string{
				"error": http.StatusText(status),
			})
			return
		}

		bookFile, err := client.BookFile.Query().
			Where(
				bookfile.And(
					bookfile.HasBookWith(book.ID(ksuid.ID(bookID))),
					bookfile.FormatEQ(bookfile.Format(strings.ToUpper(bookFormat))),
				),
			).
			First(ctx)
		if err != nil {
			// TODO: Convert if Format isn't available?
			status := http.StatusNotFound
			if !errors.Is(err, &ent.NotFoundError{}) {
				log.Error().Err(err).Msg("Failed to query book file")
				status = http.StatusInternalServerError
			}
			w.WriteHeader(status)
			render.JSON(w, r, map[string]string{
				"error": http.StatusText(status),
			})
			return
		}
		mtype := mime.TypeByExtension("." + strings.ToLower(bookFile.Format.String()))
		if mtype == "" {
			mtype = "application/octet-stream"
		}
		w.Header().Set("Content-Type", mtype)
		dispo := fmt.Sprintf("attachment; filename=%s; filename*=UTF-8''%s",
			url.QueryEscape(bookFile.Name),
			url.QueryEscape(bookFile.Name),
		)
		w.Header().Set("Content-Disposition", dispo)
		http.ServeFile(w, r, bookFile.Path)
	}
}
