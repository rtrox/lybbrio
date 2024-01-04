package calibre

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetAuthor(t *testing.T) {
	tests := []struct {
		name     string
		expected *Author
	}{
		{
			name: "Pierce Brown",
			expected: &Author{
				ID:   1,
				Name: "Pierce Brown",
				Sort: "Brown, Pierce",
			},
		},
		{
			name: "James S.A. Corey",
			expected: &Author{
				ID:   3,
				Name: "James S.A. Corey",
				Sort: "Corey, James S.A.",
			},
		},
		{
			name: "Simon R. Green",
			expected: &Author{
				ID:   4,
				Name: "Simon R. Green",
				Sort: "Green, Simon R.",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			author, err := db.GetAuthor(context.Background(), tt.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expected, author)
		})
	}
}

func TestGetAuthors(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	db, err := NewCalibreSQLite("test_fixtures/metadata.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	authors, err := db.GetAuthors(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(3, len(authors))
	require.Equal(int64(1), authors[0].ID)
	require.Equal("Pierce Brown", authors[0].Name)
}

func TestGetAuthorBooks(t *testing.T) {
	tests := []struct {
		name           string
		id             int64
		expected_count int
	}{
		{
			name:           "Pierce Brown",
			id:             1,
			expected_count: 4,
		},
		{
			name:           "James S.A. Corey",
			id:             3,
			expected_count: 5,
		},
		{
			name:           "Simon R. Green",
			id:             4,
			expected_count: 2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			books, err := db.GetAuthorBooks(context.Background(), tt.id)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expected_count, len(books))
		})
	}
}

func TestGetAuthorSeries(t *testing.T) {
	tests := []struct {
		name           string
		id             int64
		expected_count int
		expected       []*Series
	}{
		{
			name:           "Pierce Brown",
			id:             1,
			expected_count: 1,
			expected: []*Series{
				{
					ID:        1,
					Name:      "Red Rising Saga",
					Sort:      "Red Rising Saga",
					BookCount: 4,
				},
			},
		},
		{
			name:           "James S.A. Corey",
			id:             3,
			expected_count: 1,
			expected: []*Series{
				{
					ID:        2,
					Name:      "The Expanse",
					Sort:      "Expanse, The",
					BookCount: 5,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			series, err := db.GetAuthorSeries(context.Background(), tt.id)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(len(series), tt.expected_count)
			if len(tt.expected) > 0 {
				require.Equal(len(tt.expected), len(series))
				if len(tt.expected) != len(series) {
					return
				}
				for i, s := range series {
					require.Equal(tt.expected[i].ID, s.ID)
					require.Equal(tt.expected[i].Name, s.Name)
					require.Equal(tt.expected[i].Sort, s.Sort)
					require.Equal(tt.expected[i].BookCount, s.BookCount)
				}
			}
		})
	}
}

func TestGetBook(t *testing.T) {
	sid := float64(2.0)
	ts := time.Date(2023, time.November, 19, 0, 3, 5, 25617000, time.UTC)
	pd := time.Date(2015, time.January, 6, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name     string
		expected *Book
	}{
		{
			name: "Golden Son",
			expected: &Book{
				ID:           452,
				Title:        "Golden Son",
				Sort:         "Golden Son",
				Timestamp:    &ts,
				PubDate:      &pd,
				SeriesIndex:  &sid,
				AuthorSort:   "Brown, Pierce",
				ISBN:         "",
				LCCN:         "",
				Path:         "Pierce Brown/Golden Son (452)",
				Flags:        1,
				UUID:         "16b5596d-f940-4081-85d7-5eee5d15d737",
				HasCover:     true,
				LastModified: time.Date(2023, time.November, 19, 11, 17, 5, 98700000, time.UTC),
				Authors: []Author{
					{
						ID:   1,
						Name: "Pierce Brown",
						Sort: "Brown, Pierce",
					},
				},
				Tags: []Tag{
					{
						ID:   1,
						Name: "Science Fiction",
					},
					{
						ID:   2,
						Name: "Fantasy",
					},
					{
						ID:   3,
						Name: "Fiction",
					},
					{
						ID:   4,
						Name: "Dystopia",
					},
					{
						ID:   5,
						Name: "Young Adult",
					},
					{
						ID:   6,
						Name: "Audiobook",
					},
					{
						ID:   7,
						Name: "Adult",
					},
					{
						ID:   8,
						Name: "Science Fiction Fantasy",
					},
					{
						ID:   10,
						Name: "Space",
					},
					{
						ID:   15,
						Name: "Adventure",
					},
				},
				Identifiers: []Identifier{
					{
						ID:     731,
						BookID: 452,
						Type:   "goodreads",
						Val:    "18966819",
					},
					{
						ID:     732,
						BookID: 452,
						Type:   "isbn",
						Val:    "9780345539823",
					},
				},
				Publisher: []Publisher{
					{
						ID:   2,
						Name: "Del Rey",
					},
				},
				Comments: Comment{
					ID:   582,
					Book: 452,
					Text: "As a Red, Darrow grew up working the mines deep beneath the surface of Mars, enduring backbreaking labor while dreaming of the better future he was building for his descendants. But the Society he faithfully served was built on lies. Darrow’s kind have been betrayed and denied by their elitist masters, the Golds—and their only path to liberation is revolution. And so Darrow sacrifices himself in the name of the greater good for which Eo, his true love and inspiration, laid down her own life. He becomes a Gold, infiltrating their privileged realm so that he can destroy it from within. \r\n\r\nA lamb among wolves in a cruel world, Darrow finds friendship, respect, and even love—but also the wrath of powerful rivals. To wage and win the war that will change humankind’s destiny, Darrow must confront the treachery arrayed against him, overcome his all-too-human desire for retribution—and strive not for violent revolt but a hopeful rebirth. Though the road ahead is fraught with danger and deceit, Darrow must choose to follow Eo’s principles of love and justice to free his people. \r\n\r\nHe must live for more.",
				},
				Languages: []Language{
					{
						ID:       1,
						LangCode: "eng",
					},
				},
				Series: []Series{
					{
						ID:   1,
						Name: "Red Rising Saga",
						Sort: "Red Rising Saga",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			book, err := db.GetBook(context.Background(), tt.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expected, book)
		})
	}
}

func TestGetBookByIdentifier(t *testing.T) {
	tests := []struct {
		name             string
		identifier       string
		expectedBookId   int64
		expectedBookName string
	}{
		{
			name:             "goodreads_53356770",
			identifier:       "53356770",
			expectedBookId:   107,
			expectedBookName: "Red Rising",
		},
		{
			name:             "asin_B009SQ018I",
			identifier:       "B009SQ018I",
			expectedBookId:   121,
			expectedBookName: "Abaddon's Gate",
		},
		{
			name:             "goodreads_29217027",
			identifier:       "29217027",
			expectedBookId:   375,
			expectedBookName: "Iron Gold",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			book, err := db.GetBookByIdentifier(context.Background(), tt.identifier)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expectedBookId, book.ID)
			require.Equal(tt.expectedBookName, book.Title)
		})
	}
}

func TestGetBooks(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	db, err := NewCalibreSQLite("test_fixtures/metadata.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	books, err := db.GetBooks(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(11, len(books))
	require.Equal(int64(105), books[0].ID)
	require.Equal("Cibola Burn", books[0].Title)
}

func TestGetBooks_WithPagination(t *testing.T) {
	tests := []struct {
		name        string
		page        int
		pageSize    int
		expectedIds []int64
	}{
		{
			name:        "page_1_pageSize_3",
			page:        1,
			pageSize:    3,
			expectedIds: []int64{105, 107, 109},
		},
		{
			name:        "page_2_pageSize_3",
			page:        2,
			pageSize:    3,
			expectedIds: []int64{121, 122, 375},
		},
		{
			name:        "page_3_pageSize_3",
			page:        3,
			pageSize:    3,
			expectedIds: []int64{376, 452, 471},
		},
		{
			name:        "page_1_pageSize_10",
			page:        1,
			pageSize:    10,
			expectedIds: []int64{105, 107, 109, 121, 122, 375, 376, 452, 471, 492},
		},
		{
			name:        "page_2_pageSize_10",
			page:        2,
			pageSize:    1,
			expectedIds: []int64{107},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			books, err := db.WithPagination(tt.page, tt.pageSize).GetBooks(context.Background())
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.pageSize, len(books))
			ids := make([]int64, len(books))
			for i, b := range books {
				ids[i] = b.ID
			}
			if len(tt.expectedIds) == len(books) {
				for i, id := range tt.expectedIds {
					require.Equal(id, books[i].ID)
				}
			}
		})
	}
}
func TestGetTag(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tests := []struct {
		name     string
		expected *Tag
	}{
		{
			name: "Science Fiction",
			expected: &Tag{
				ID:   1,
				Name: "Science Fiction",
			},
		},
		{
			name: "Fantasy",
			expected: &Tag{
				ID:   2,
				Name: "Fantasy",
			},
		},
		{
			name: "Fiction",
			expected: &Tag{
				ID:   3,
				Name: "Fiction",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			tag, err := db.GetTag(context.Background(), tt.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tag, tt.expected)
		})
	}
}

func TestGetTags(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	db, err := NewCalibreSQLite("test_fixtures/metadata.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	tags, err := db.GetTags(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(19, len(tags))
	require.Equal(int64(1), tags[0].ID)
	require.Equal("Science Fiction", tags[0].Name)
}

func TestGetTagBooks(t *testing.T) {
	tests := []struct {
		name           string
		id             int64
		expected_count int
	}{
		{
			name:           "Science Fiction",
			id:             1,
			expected_count: 11,
		},
		{
			name:           "Fantasy",
			id:             2,
			expected_count: 10,
		},
		{
			name:           "Fiction",
			id:             3,
			expected_count: 11,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			books, err := db.GetTagBooks(context.Background(), tt.id)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expected_count, len(books))
		})
	}
}

func TestGetPublisher(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tests := []struct {
		name     string
		expected *Publisher
	}{
		{
			name: "Del Rey",
			expected: &Publisher{
				ID:   2,
				Name: "Del Rey",
			},
		},
		{
			name: "Tor Books",
			expected: &Publisher{
				ID:   3,
				Name: "Orbit",
			},
		},
		{
			name: "Ace",
			expected: &Publisher{
				ID:   6,
				Name: "Random House Publishing Group",
			},
		},
		{
			name: "Orbit",
			expected: &Publisher{
				ID:   8,
				Name: "Little, Brown Book Group",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			publisher, err := db.GetPublisher(context.Background(), tt.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expected, publisher)
		})
	}
}

func TestGetPublishers(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	db, err := NewCalibreSQLite("test_fixtures/metadata.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	publishers, err := db.GetPublishers(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(5, len(publishers))
	require.Equal(int64(2), publishers[0].ID)
	require.Equal("Del Rey", publishers[0].Name)
}

func TestGetPublisherBooks(t *testing.T) {
	tests := []struct {
		name           string
		id             int64
		expected_count int
	}{
		{
			name:           "Del Rey",
			id:             2,
			expected_count: 3,
		},
		{
			name:           "Orbit",
			id:             3,
			expected_count: 1,
		},
		{
			name:           "Random House Publishing Group",
			id:             6,
			expected_count: 1,
		},
		{
			name:           "Little, Brown Book Group",
			id:             8,
			expected_count: 1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			books, err := db.GetPublisherBooks(context.Background(), tt.id)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expected_count, len(books))
		})
	}
}

func TestGetLanguage(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tests := []struct {
		name     string
		expected *Language
	}{
		{
			name: "English",
			expected: &Language{
				ID:       1,
				LangCode: "eng",
			},
		},
		{
			name: "French",
			expected: &Language{
				ID:       2,
				LangCode: "dan",
			},
		},
		{
			name: "German",
			expected: &Language{
				ID:       3,
				LangCode: "por",
			},
		},
		{
			name: "Spanish",
			expected: &Language{
				ID:       4,
				LangCode: "swe",
			},
		},
		{
			name: "Italian",
			expected: &Language{
				ID:       5,
				LangCode: "rus",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			language, err := db.GetLanguage(context.Background(), tt.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expected, language)
		})
	}
}

func TestGetLanguages(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	db, err := NewCalibreSQLite("test_fixtures/metadata.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	languages, err := db.GetLanguages(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(5, len(languages))
	require.Equal(int64(1), languages[0].ID)
	require.Equal("eng", languages[0].LangCode)
}

func TestGetLanguageBooks(t *testing.T) {
	tests := []struct {
		name           string
		id             int64
		expected_count int
	}{
		{
			name:           "English",
			id:             1,
			expected_count: 11,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			books, err := db.GetLanguageBooks(context.Background(), tt.id)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expected_count, len(books))
		})
	}
}

func TestGetSeries(t *testing.T) {
	tests := []struct {
		name     string
		expected *Series
	}{
		{
			name: "Red Rising Saga",
			expected: &Series{
				ID:        1,
				Name:      "Red Rising Saga",
				Sort:      "Red Rising Saga",
				BookCount: 4,
			},
		},
		{
			name: "The Expanse",
			expected: &Series{
				ID:        2,
				Name:      "The Expanse",
				Sort:      "Expanse, The",
				BookCount: 5,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			series, err := db.GetSeries(context.Background(), tt.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expected, series)
		})
	}
}

func TestGetSeriesList(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	db, err := NewCalibreSQLite("test_fixtures/metadata.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	series, err := db.GetSeriesList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(3, len(series))
	require.Equal(int64(1), series[0].ID)
	require.Equal("Red Rising Saga", series[0].Name)
	require.Equal(int64(4), series[0].BookCount)
}

func TestGetSeriesBooks(t *testing.T) {
	tests := []struct {
		name           string
		id             int64
		expected_count int
	}{
		{
			name:           "Red Rising Saga",
			id:             1,
			expected_count: 4,
		},
		{
			name:           "The Expanse",
			id:             2,
			expected_count: 5,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)

			db, err := NewCalibreSQLite("test_fixtures/metadata.db")
			if err != nil {
				t.Fatal(err)
			}
			defer db.Close()

			books, err := db.GetSeriesBooks(context.Background(), tt.id)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(tt.expected_count, len(books))
		})
	}
}
