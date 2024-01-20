package handler

import (
	"fmt"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/bookfile"
	"lybbrio/internal/ent/schema/filetype"
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
	r.Get("/{bookID}/{book_format}", Download(client))
	return r
}

func Download(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("path", r.URL.Path).Msg("Download")
		ctx := r.Context()
		bookID := chi.URLParam(r, "bookID")
		bookFormat := chi.URLParam(r, "book_format")
		log.Info().
			Str("bookID", bookID).
			Str("book_format", bookFormat).
			Msg("Params")
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
			if !ent.IsNotFound(err) {
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
		dispo := fmt.Sprintf("attachment; filename=%s%s; filename*=UTF-8''%s%s",
			url.QueryEscape(bookFile.Name),
			filetype.FromString(bookFile.Format.String()).Extension(),
			url.QueryEscape(bookFile.Name),
			filetype.FromString(bookFile.Format.String()).Extension(),
		)
		w.Header().Set("Content-Disposition", dispo)
		http.ServeFile(w, r, bookFile.Path)
	}
}
