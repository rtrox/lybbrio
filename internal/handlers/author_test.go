package handlers

import (
	"encoding/json"
	"fmt"
	"lybbrio/internal/calibre"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAuthorCtx(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	mockCal := &calibre.MockCalibre{}
	mockCal.On("GetAuthor", mock.Anything, int64(1)).Return(&calibre.Author{
		ID:   1,
		Name: "Test Author",
		Sort: "Test Author",
		Link: "http://example.com",
	}, nil)

	r := chi.NewRouter()
	r.Route("/{authorId}", func(r chi.Router) {
		r.Use(AuthorCtx(mockCal))
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			author := authorFromContext(r.Context())
			require.NotNil(author)
			require.Equal(int64(1), author.ID)
			require.Equal("Test Author", author.Name)
			require.Equal("Test Author", author.Sort)
			require.Equal("http://example.com", author.Link)
		})
	})
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL+"/1", nil)
	require.NoError(err)
	resp, err := ts.Client().Do(req)
	require.NoError(err)
	require.Equal(http.StatusOK, resp.StatusCode)

	mockCal.AssertExpectations(t)
	mockCal.AssertCalled(t, "GetAuthor", mock.Anything, int64(1))
}

func TestGetAuthor(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	expected := &calibre.Author{
		ID:   1,
		Name: "Test Author",
		Sort: "Test Author",
		Link: "http://example.com",
	}
	mockCal := &calibre.MockCalibre{}
	mockCal.On("GetAuthor", mock.Anything, int64(1)).Return(expected, nil)

	ts := httptest.NewServer(AuthorRoutes(mockCal))
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/1")

	require.NoError(err)
	defer resp.Body.Close()
	require.Equal(http.StatusOK, resp.StatusCode)

	actual := &AuthorListResponse{}
	err = json.NewDecoder(resp.Body).Decode(actual)
	require.NoError(err)
	require.Equal(expected, actual)

	mockCal.AssertExpectations(t)
	mockCal.AssertCalled(t, "GetAuthor", mock.Anything, int64(1))
}

func TestGetAuthors(t *testing.T) {
	tests := []struct {
		name             string
		pageSize         int
		shouldReturnPage bool
	}{
		{
			name:             "pageSize_30",
			pageSize:         30,
			shouldReturnPage: false,
		},
		{
			name:             "pageSize_1",
			pageSize:         1,
			shouldReturnPage: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			t.Parallel()
			require := require.New(t)
			expected := []*calibre.Author{
				{
					ID:   1,
					Name: "Test Author",
					Sort: "Test Author",
					Link: "http://example.com",
				},
				{
					ID:   2,
					Name: "Test Author 2",
					Sort: "Test Author 2",
					Link: "http://example.com",
				},
			}
			mockCal := &calibre.MockCalibre{}
			mockCal.On("GetAuthors", mock.Anything).Return(expected, nil)
			mockCal.On("WithPagination", mock.Anything, mock.Anything).Return(mockCal, nil)
			ts := httptest.NewServer(AuthorRoutes(mockCal))
			defer ts.Close()

			resp, err := ts.Client().Get(ts.URL + "?pageSize=" + fmt.Sprintf("%d", tt.pageSize))

			require.NoError(err)
			defer resp.Body.Close()
			require.Equal(http.StatusOK, resp.StatusCode)

			actual := &AuthorListResponse{}
			err = json.NewDecoder(resp.Body).Decode(&actual)

			require.NoError(err)
			require.Equal(expected, actual.Items)
			if tt.shouldReturnPage {
				require.NotNil(actual.Page, "Page should not be nil when results >= pageSize")
			} else {
				require.Nil(actual.Page, "Should not return Page struct on final page")
			}

			mockCal.AssertExpectations(t)
			mockCal.AssertCalled(t, "GetAuthors", mock.Anything)
		})
	}
}
