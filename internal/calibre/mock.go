package calibre

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockCalibre struct {
	mock.Mock
}

func (m *MockCalibre) GetAuthor(ctx context.Context, id int64) (*Author, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Author), args.Error(1)
}

func (m *MockCalibre) GetAuthors(ctx context.Context) ([]*Author, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Author), args.Error(1)
}

func (m *MockCalibre) GetAuthorBooks(ctx context.Context, id int64) ([]*Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]*Book), args.Error(1)
}

func (m *MockCalibre) GetAuthorSeries(ctx context.Context, id int64) ([]*Series, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]*Series), args.Error(1)
}

func (m *MockCalibre) GetBook(ctx context.Context, id int64) (*Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Book), args.Error(1)
}

func (m *MockCalibre) GetBooks(ctx context.Context) ([]*Book, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Book), args.Error(1)
}

func (m *MockCalibre) GetBookByIdentifier(ctx context.Context, val string) (*Book, error) {
	args := m.Called(ctx, val)
	return args.Get(0).(*Book), args.Error(1)
}

func (m *MockCalibre) GetTag(ctx context.Context, id int64) (*Tag, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Tag), args.Error(1)
}

func (m *MockCalibre) GetTags(ctx context.Context) ([]*Tag, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Tag), args.Error(1)
}

func (m *MockCalibre) GetTagBooks(ctx context.Context, id int64) ([]*Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]*Book), args.Error(1)
}

func (m *MockCalibre) GetPublisher(ctx context.Context, id int64) (*Publisher, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Publisher), args.Error(1)
}

func (m *MockCalibre) GetPublishers(ctx context.Context) ([]*Publisher, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Publisher), args.Error(1)
}

func (m *MockCalibre) GetPublisherBooks(ctx context.Context, id int64) ([]*Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]*Book), args.Error(1)
}

func (m *MockCalibre) GetLanguage(ctx context.Context, id int64) (*Language, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Language), args.Error(1)
}

func (m *MockCalibre) GetLanguages(ctx context.Context) ([]*Language, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Language), args.Error(1)
}

func (m *MockCalibre) GetLanguageBooks(ctx context.Context, id int64) ([]*Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]*Book), args.Error(1)
}

func (m *MockCalibre) GetSeriesList(ctx context.Context) ([]*Series, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*Series), args.Error(1)
}

func (m *MockCalibre) GetSeries(ctx context.Context, id int64) (*Series, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Series), args.Error(1)
}

func (m *MockCalibre) GetSeriesBooks(ctx context.Context, id int64) ([]*Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]*Book), args.Error(1)
}

func (m *MockCalibre) WithPagination(page int, pageSize int) Calibre {
	args := m.Called(page, pageSize)
	return args.Get(0).(Calibre)
}
