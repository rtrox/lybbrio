// Code generated by ent, DO NOT EDIT.

package series

import (
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the series type in the database.
	Label = "series"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// EdgeBooks holds the string denoting the books edge name in mutations.
	EdgeBooks = "books"
	// EdgeSeriesBooks holds the string denoting the series_books edge name in mutations.
	EdgeSeriesBooks = "series_books"
	// Table holds the table name of the series in the database.
	Table = "series"
	// BooksTable is the table that holds the books relation/edge. The primary key declared below.
	BooksTable = "series_books"
	// BooksInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	BooksInverseTable = "books"
	// SeriesBooksTable is the table that holds the series_books relation/edge.
	SeriesBooksTable = "series_books"
	// SeriesBooksInverseTable is the table name for the SeriesBook entity.
	// It exists in this package in order to avoid circular dependency with the "seriesbook" package.
	SeriesBooksInverseTable = "series_books"
	// SeriesBooksColumn is the table column denoting the series_books relation/edge.
	SeriesBooksColumn = "series_id"
)

// Columns holds all SQL columns for series fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSort,
}

var (
	// BooksPrimaryKey and BooksColumn2 are the table columns denoting the
	// primary key for the books relation (M2M).
	BooksPrimaryKey = []string{"series_id", "book_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "lybbrio/internal/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// SortValidator is a validator for the "sort" field. It is called by the builders before save.
	SortValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ksuid.ID
)

// OrderOption defines the ordering options for the Series queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySort orders the results by the sort field.
func BySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSort, opts...).ToFunc()
}

// ByBooksCount orders the results by books count.
func ByBooksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBooksStep(), opts...)
	}
}

// ByBooks orders the results by books terms.
func ByBooks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBooksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySeriesBooksCount orders the results by series_books count.
func BySeriesBooksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSeriesBooksStep(), opts...)
	}
}

// BySeriesBooks orders the results by series_books terms.
func BySeriesBooks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSeriesBooksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBooksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BooksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, BooksTable, BooksPrimaryKey...),
	)
}
func newSeriesBooksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SeriesBooksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, SeriesBooksTable, SeriesBooksColumn),
	)
}
