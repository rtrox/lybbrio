package calibre

import (
	"context"

	gormLogger "github.com/mpalmer/gorm-zerolog"
	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func asdf() error {
	return nil
}

type Calibre interface {
	GetAuthor(ctx context.Context, id int64) (*Author, error)
	GetAuthors(ctx context.Context) ([]*Author, error)
	GetAuthorBooks(ctx context.Context, id int64) ([]*Book, error)
	GetAuthorSeries(ctx context.Context, id int64) ([]*Series, error)

	GetBook(ctx context.Context, id int64) (*Book, error)
	GetBooks(ctx context.Context) ([]*Book, error)
	GetBookByIdentifier(ctx context.Context, val string) (*Book, error)

	GetTag(ctx context.Context, id int64) (*Tag, error)
	GetTags(ctx context.Context) ([]*Tag, error)
	GetTagBooks(ctx context.Context, id int64) ([]*Book, error)

	GetPublisher(ctx context.Context, id int64) (*Publisher, error)
	GetPublishers(ctx context.Context) ([]*Publisher, error)
	GetPublisherBooks(ctx context.Context, id int64) ([]*Book, error)

	GetLanguage(ctx context.Context, id int64) (*Language, error)
	GetLanguages(ctx context.Context) ([]*Language, error)
	GetLanguageBooks(ctx context.Context, id int64) ([]*Book, error)

	GetSeriesList(ctx context.Context) ([]*Series, error)
	GetSeries(ctx context.Context, id int64) (*Series, error)
	GetSeriesBooks(ctx context.Context, id int64) ([]*Book, error)

	WithPagination(page, pageSize int) Calibre
}

type CalibreSQLite struct {
	db *gorm.DB
}

func NewCalibreSQLite(dbPath string) (*CalibreSQLite, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	asdf()
	if err != nil {
		return nil, err
	}
	return &CalibreSQLite{db: db}, nil
}

func (c *CalibreSQLite) Close() error {
	db, err := c.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (c *CalibreSQLite) WithLogger(logger *zerolog.Logger) Calibre {
	return &CalibreSQLite{
		db: c.db.Session(
			&gorm.Session{
				Logger: gormLogger.Logger{},
			},
		),
	}
}

func (c *CalibreSQLite) WithPagination(page, pageSize int) Calibre {
	return &CalibreSQLite{db: c.db.Offset((page - 1) * pageSize).Limit(pageSize)}
}

func (s *CalibreSQLite) GetAuthor(ctx context.Context, id int64) (*Author, error) {
	var author Author
	err := s.db.WithContext(ctx).
		First(&author, id).Error
	return &author, err
}

func (s *CalibreSQLite) GetAuthors(ctx context.Context) ([]*Author, error) {
	var authors []*Author
	err := s.db.WithContext(ctx).
		Find(&authors).Error
	return authors, err
}

func (s *CalibreSQLite) GetAuthorBooks(ctx context.Context, id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.WithContext(ctx).
		Model(&Author{ID: id}).
		Preload("Authors").
		Preload("Tags").
		Preload("Comments").
		Association("Books").
		Find(&books)
	return books, err
}

func (s *CalibreSQLite) GetAuthorSeries(ctx context.Context, id int64) ([]*Series, error) {
	var series []*Series
	err := s.db.WithContext(ctx).
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

func (s *CalibreSQLite) GetBook(ctx context.Context, id int64) (*Book, error) {
	var book Book
	err := s.db.WithContext(ctx).Preload(clause.Associations).First(&book, id).Error
	return &book, err
}

func (s *CalibreSQLite) GetBooks(ctx context.Context) ([]*Book, error) {
	var books []*Book
	err := s.db.WithContext(ctx).
		Preload(clause.Associations).
		Find(&books).
		Error
	return books, err
}

func (s *CalibreSQLite) GetBookByIdentifier(ctx context.Context, val string) (*Book, error) {
	var identifier Identifier
	err := s.db.WithContext(ctx).
		Model(&Identifier{}).
		Preload("Book").
		Preload("Book.Authors").
		Preload("Book.Tags").
		Preload("Book.Comments").
		Where("val = ?", val).
		Find(&identifier).Error
	return &identifier.Book, err
}

func (s *CalibreSQLite) GetTag(ctx context.Context, id int64) (*Tag, error) {
	var tag Tag
	err := s.db.WithContext(ctx).First(&tag, id).Error
	return &tag, err
}

func (s *CalibreSQLite) GetTags(ctx context.Context) ([]*Tag, error) {
	var tags []*Tag
	err := s.db.WithContext(ctx).Find(&tags).Error
	return tags, err
}

func (s *CalibreSQLite) GetTagBooks(ctx context.Context, id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.WithContext(ctx).
		Model(&Tag{ID: id}).
		Preload("Authors").
		Preload("Tags").
		Preload("Comments").
		Association("Books").
		Find(&books)
	return books, err
}

func (s *CalibreSQLite) GetPublisher(ctx context.Context, id int64) (*Publisher, error) {
	var publisher Publisher
	err := s.db.WithContext(ctx).First(&publisher, id).Error
	return &publisher, err
}

func (s *CalibreSQLite) GetPublishers(ctx context.Context) ([]*Publisher, error) {
	var publishers []*Publisher
	err := s.db.WithContext(ctx).Find(&publishers).Error
	return publishers, err
}

func (s *CalibreSQLite) GetPublisherBooks(ctx context.Context, id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.WithContext(ctx).
		Model(&Publisher{ID: id}).
		Preload("Authors").
		Preload("Tags").
		Preload("Comments").
		Association("Books").
		Find(&books)
	return books, err
}

func (s *CalibreSQLite) GetLanguage(ctx context.Context, id int64) (*Language, error) {
	var language Language
	err := s.db.WithContext(ctx).First(&language, id).Error
	return &language, err
}

func (s *CalibreSQLite) GetLanguages(ctx context.Context) ([]*Language, error) {
	var languages []*Language
	err := s.db.WithContext(ctx).Find(&languages).Error
	return languages, err
}

func (s *CalibreSQLite) GetLanguageBooks(ctx context.Context, id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.WithContext(ctx).Model(&Language{ID: id}).Association("Books").Find(&books)
	return books, err
}

func (s *CalibreSQLite) GetSeries(ctx context.Context, id int64) (*Series, error) {
	var series Series
	err := s.db.WithContext(ctx).
		Model(&Series{}).
		Select("series.*, COUNT(DISTINCT books_series_link.book) AS book_count").
		Joins("JOIN books_series_link ON books_series_link.series = series.id").
		Where("series.id = ?", id).
		Group("series.id").
		Scan(&series).
		Error
	return &series, err
}

func (s *CalibreSQLite) GetSeriesList(ctx context.Context) ([]*Series, error) {
	var series []*Series
	err := s.db.WithContext(ctx).
		Model(&Series{}).
		Select("series.*, COUNT(DISTINCT books_series_link.book) AS book_count").
		Joins("JOIN books_series_link ON books_series_link.series = series.id").
		Group("series.id").
		Scan(&series).
		Error
	return series, err
}

func (s *CalibreSQLite) GetSeriesBooks(ctx context.Context, id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.WithContext(ctx).
		Model(&Series{ID: id}).
		Preload("Authors").
		Preload("Tags").
		Preload("Comments").
		Association("Books").
		Find(&books)
	return books, err
}
