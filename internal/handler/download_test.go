package handler

import (
	"context"
	"lybbrio/internal/db"
	"lybbrio/internal/ent"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"

	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/viewer"
)

type downloadTestContext struct {
	client   *ent.Client
	teardown func()
	adminCtx context.Context
}

func (t downloadTestContext) Teardown() {
	t.client.Close()
	t.teardown()
}

func setupDownloadHandlerTest(t *testing.T, testName string, teardown ...func()) downloadTestContext {
	var ret downloadTestContext

	ret.client = db.OpenTest(t, testName)

	ret.teardown = func() {
		ret.client.Close()
	}

	ret.adminCtx = viewer.NewSystemAdminContext(context.Background())

	return ret
}

func addContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := viewer.NewContext(r.Context(), "usr_asdf", permissions.NewPermissions())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TestDownload(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	tc := setupDownloadHandlerTest(t, "TestDownload")
	defer tc.Teardown()

	r := chi.NewRouter()
	r.Use(addContextMiddleware)
	r.Mount("/", DownloadRoutes(tc.client))
	server := httptest.NewServer(r)
	defer server.Close()

	book, err := tc.client.Book.Create().
		SetTitle("Test Book").
		SetSort("Test Book").
		SetPath("/testdata/").Save(tc.adminCtx)
	require.NoError(err)

	bookFile, err := tc.client.BookFile.Create().
		SetBook(book).
		SetName("Twenty Thousand Leagues under t - Jules Verne").
		SetPath("testdata/Twenty Thousand Leagues under t - Jules Verne.epub").
		SetSize(369814).
		SetFormat("EPUB").
		Save(tc.adminCtx)
	require.NoError(err)

	url := server.URL + "/" + book.ID.String() + "/" + bookFile.Format.String()
	resp, err := server.Client().Get(url)
	require.NoError(err)
	require.Equal(200, resp.StatusCode)
	require.Equal("application/epub+zip", resp.Header.Get("Content-Type"))
	require.Equal(
		"attachment; filename=Twenty+Thousand+Leagues+under+t+-+Jules+Verne.epub; "+
			"filename*=UTF-8''Twenty+Thousand+Leagues+under+t+-+Jules+Verne.epub",
		resp.Header.Get("Content-Disposition"),
	)
	require.Equal("369814", resp.Header.Get("Content-Length"))
}

func TestDownloadErrors(t *testing.T) {
	tests := []struct {
		name       string
		bookID     string
		bookFormat string
		wantCode   int
	}{
		{
			name:       "Missing bookID",
			bookFormat: "EPUB",
			wantCode:   http.StatusBadRequest,
		},
		{
			name:     "Missing book_format",
			bookID:   "asdf",
			wantCode: http.StatusNotFound,
		},
		{
			name:       "Nonexistent bookID",
			bookID:     "asdf",
			bookFormat: "EPUB",
			wantCode:   http.StatusNotFound,
		},
		{
			name:       "Invalid book_format",
			bookID:     "asdf",
			bookFormat: "asdf",
			wantCode:   http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tc := setupDownloadHandlerTest(t, tt.name)
			defer tc.Teardown()

			r := chi.NewRouter()
			r.Use(addContextMiddleware)
			r.Mount("/", DownloadRoutes(tc.client))
			server := httptest.NewServer(r)
			defer server.Close()

			url := server.URL + "/" + tt.bookID + "/" + tt.bookFormat
			resp, err := server.Client().Get(url)
			require.NoError(t, err)
			require.Equal(t, tt.wantCode, resp.StatusCode)
		})
	}
}
