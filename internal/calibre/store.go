package calibre

import (
	stdlog "log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type Calibre interface {
	GetAuthor(id int64) (*Author, error)
	GetAuthors() ([]*Author, error)
	GetAuthorBooks(id int64) ([]*Book, error)
	GetAuthorSeries(id int64) ([]*Series, error)

	GetBook(id int64) (*Book, error)
	GetBooks() ([]*Book, error)

	GetTag(id int64) (*Tag, error)
	GetTags() ([]*Tag, error)
	GetTagBooks(id int64) ([]*Book, error)

	GetPublisher(id int64) (*Publisher, error)
	GetPublishers() ([]*Publisher, error)
	GetPublisherBooks(id int64) ([]*Book, error)

	GetLanguage(id int64) (*Language, error)
	GetLanguages() ([]*Language, error)
	GetLanguageBooks(id int64) ([]*Book, error)

	GetIdentifier(id int64) (*Identifier, error)
	GetIdentifierBook(id int64) (*Book, error)

	GetSeriesList() ([]*Series, error)
	GetSeries(id int64) (*Series, error)
	GetSeriesBooks(id int64) ([]*Book, error)

	WithPagination(page, pageSize int) Calibre
}

type CalibreSQLite struct {
	db *gorm.DB
}

func NewCalibreSQLite(dbPath string) (*CalibreSQLite, error) {
	newLogger := logger.New(
		stdlog.New(os.Stdout, "\r\n", stdlog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}
	return &CalibreSQLite{db: db}, nil
}

func (c *CalibreSQLite) WithPagination(page, pageSize int) Calibre {
	return &CalibreSQLite{db: c.db.Offset((page - 1) * pageSize).Limit(pageSize)}
}

func (s *CalibreSQLite) GetAuthor(id int64) (*Author, error) {
	var author Author
	err := s.db.First(&author, id).Error
	return &author, err
}

func (s *CalibreSQLite) GetAuthors() ([]*Author, error) {
	var authors []*Author
	err := s.db.Find(&authors).Error
	return authors, err
}

func (s *CalibreSQLite) GetAuthorBooks(id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.Model(&Author{ID: id}).Association("Books").Find(&books)
	return books, err
}

func (s *CalibreSQLite) GetAuthorSeries(id int64) ([]*Series, error) {
	var series []*Series
	err := s.db.
		Model(&Series{}).
		Select("series.*, COUNT(DISTINCT books_authors_link.book) AS book_count").
		Joins("JOIN books_series_link ON books_series_link.series = series.id").
		Joins("JOIN books_authors_link ON books_authors_link.book = books_series_link.book").
		Where("books_authors_link.author = ?", id).
		Group("series.id").
		Scan(&series).
		Error
	return series, err
}

func (s *CalibreSQLite) GetBook(id int64) (*Book, error) {
	var book Book
	err := s.db.Preload(clause.Associations).First(&book, id).Error
	return &book, err
}

func (s *CalibreSQLite) GetBooks() ([]*Book, error) {
	var books []*Book
	err := s.db.
		Preload("Authors").
		Preload("Tags").
		Preload("Comments").
		Find(&books).
		Error
	return books, err
}

func (s *CalibreSQLite) GetTag(id int64) (*Tag, error) {
	var tag Tag
	err := s.db.First(&tag, id).Error
	return &tag, err
}

func (s *CalibreSQLite) GetTags() ([]*Tag, error) {
	var tags []*Tag
	err := s.db.Find(&tags).Error
	return tags, err
}

func (s *CalibreSQLite) GetTagBooks(id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.Model(&Tag{ID: id}).
		Preload("Book.Authors").
		Preload("Book.Tags").
		Preload("Book.Comments").
		Association("Books").
		Find(&books)
	return books, err
}

func (s *CalibreSQLite) GetPublisher(id int64) (*Publisher, error) {
	var publisher Publisher
	err := s.db.First(&publisher, id).Error
	return &publisher, err
}

func (s *CalibreSQLite) GetPublishers() ([]*Publisher, error) {
	var publishers []*Publisher
	err := s.db.Find(&publishers).Error
	return publishers, err
}

func (s *CalibreSQLite) GetPublisherBooks(id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.Model(&Publisher{ID: id}).
		Preload("Book.Authors").
		Preload("Book.Tags").
		Preload("Book.Comments").
		Association("Books").
		Find(&books)
	return books, err
}

func (s *CalibreSQLite) GetLanguage(id int64) (*Language, error) {
	var language Language
	err := s.db.First(&language, id).Error
	return &language, err
}

func (s *CalibreSQLite) GetLanguages() ([]*Language, error) {
	var languages []*Language
	err := s.db.Find(&languages).Error
	return languages, err
}

func (s *CalibreSQLite) GetLanguageBooks(id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.Model(&Language{ID: id}).Association("Books").Find(&books)
	return books, err
}

func (s *CalibreSQLite) GetIdentifier(id int64) (*Identifier, error) {
	var identifier Identifier
	err := s.db.First(&identifier, id).Error
	return &identifier, err
}

func (s *CalibreSQLite) GetIdentifierBook(id int64) (*Book, error) {
	var book Book
	err := s.db.Model(&Identifier{ID: id}).
		Preload("Book.Authors").
		Preload("Book.Tags").
		Preload("Book.Comments").
		Association("Book").
		Find(&book)
	return &book, err
}

func (s *CalibreSQLite) GetSeries(id int64) (*Series, error) {
	var series Series
	err := s.db.
		Model(&Series{}).
		Select("series.*, COUNT(DISTINCT books_series_link.book) AS book_count").
		Joins("JOIN books_series_link ON books_series_link.series = series.id").
		Where("series.id = ?", id).
		Group("series.id").
		Scan(&series).
		Error
	return &series, err
}

func (s *CalibreSQLite) GetSeriesList() ([]*Series, error) {
	var series []*Series
	err := s.db.
		Model(&Series{}).
		Select("series.*, COUNT(DISTINCT books_series_link.book) AS book_count").
		Joins("JOIN books_series_link ON books_series_link.series = series.id").
		Group("series.id").
		Scan(&series).
		Error
	return series, err
}

func (s *CalibreSQLite) GetSeriesBooks(id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.Model(&Series{ID: id}).
		Preload("Book.Authors").
		Preload("Book.Tags").
		Preload("Book.Comments").
		Association("Books").
		Find(&books)
	return books, err
}
