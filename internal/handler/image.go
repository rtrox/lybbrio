package handler

import (
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/bookcover"
	"lybbrio/internal/ent/schema/ksuid"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func ImageRoutes(client *ent.Client) http.Handler {
	r := chi.NewRouter()
	r.Get("/{bookID}/{width}/{height}/{fileName}", CoverImage(client))
	return r
}

func CoverImage(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log := log.Ctx(ctx)
		fileName := chi.URLParam(r, "fileName")
		bookID := chi.URLParam(r, "bookID")
		widthStr := chi.URLParam(r, "width")
		heightStr := chi.URLParam(r, "height")
		if bookID == "" || fileName == "" || widthStr == "" || heightStr == "" {
			statusCodeResponse(w, r, http.StatusBadRequest)
			return
		}
		width, err := strconv.Atoi(widthStr)
		if err != nil {
			log.Debug().Err(err).Msg("Failed to parse width")
			statusCodeResponse(w, r, http.StatusBadRequest)
			return
		}
		height, err := strconv.Atoi(heightStr)
		if err != nil {
			log.Debug().Err(err).Msg("Failed to parse height")
			statusCodeResponse(w, r, http.StatusBadRequest)
			return
		}
		ext := filepath.Ext(fileName)
		if strings.TrimSuffix(fileName, ext) != "cover" {
			log.Debug().
				Str("ext", ext).
				Str("requestedFileName", fileName).
				Msg("Invalid file name")
			statusCodeResponse(w, r, http.StatusBadRequest)
			return
		}
		contentType := mime.TypeByExtension(ext)
		if contentType == "" {
			log.Debug().
				Str("ext", ext).
				Str("requestedFileName", fileName).
				Msg("Unknown content type")
			statusCodeResponse(w, r, http.StatusBadRequest)
			return
		}
		cover, err := client.BookCover.Query().
			Where(
				bookcover.And(
					bookcover.HasBookWith(book.ID(ksuid.ID(bookID))),
					bookcover.WidthLTE(width),
					bookcover.HeightLTE(height),
					bookcover.ContentType(contentType),
				),
			).
			Order(bookcover.BySize(sql.OrderAsc())).
			First(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				statusCodeResponse(w, r, http.StatusNotFound)
				return
			}
			log.Error().Err(err).Msg("Failed to query book cover")
			statusCodeResponse(w, r, http.StatusInternalServerError)
		}
		rr, err := os.Open(cover.Path)
		if err != nil {
			log.Error().Err(err).Msg("Failed to open file")
			statusCodeResponse(w, r, http.StatusInternalServerError)
			return
		}
		defer rr.Close()
		http.ServeContent(w, r, fileName, cover.UpdateTime, rr)
	}
}
