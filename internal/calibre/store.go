package calibre

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Pagination Notes:
// Order: 		https://gorm.io/docs/query.html#Order
// Limit: 		https://gorm.io/docs/query.html#Limit
// Offset: 		https://gorm.io/docs/query.html#Offset
// Pagination: 	https://gorm.io/docs/scopes.html#pagination

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

	GetSeries(id int64) (*Series, error)
	GetSeriesBooks(id int64) ([]*Book, error)

	WithPagination(page, pageSize int) Calibre
}

type CalibreSQLite struct {
	db *gorm.DB
}

func NewCalibreSQLite(db *gorm.DB) *CalibreSQLite {
	return &CalibreSQLite{db: db}
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
	err := s.db.Model(&Author{ID: id}).Association("Series").Find(&series)
	return series, err
}

func (s *CalibreSQLite) GetBook(id int64) (*Book, error) {
	var book Book
	err := s.db.Preload("Authors").Preload("Tags").Preload("Identifiers").First(&book, id).Error
	return &book, err
}

func (s *CalibreSQLite) GetBooks() ([]*Book, error) {
	var books []*Book
	err := s.db.Preload(clause.Associations).Find(&books).Error
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
	err := s.db.Model(&Tag{ID: id}).Association("Books").Find(&books)
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
	err := s.db.Model(&Publisher{ID: id}).Association("Books").Find(&books)
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
	err := s.db.Model(&Identifier{ID: id}).Association("Book").Find(&book)
	return &book, err
}

func (s *CalibreSQLite) GetSeries(id int64) (*Series, error) {
	var series Series
	err := s.db.First(&series, id).Error
	return &series, err
}

func (s *CalibreSQLite) GetSeriesBooks(id int64) ([]*Book, error) {
	var books []*Book
	err := s.db.Model(&Series{ID: id}).Association("Books").Find(&books)
	return books, err
}
