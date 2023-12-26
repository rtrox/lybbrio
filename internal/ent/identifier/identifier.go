// Code generated by ent, DO NOT EDIT.

package identifier

import (
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the identifier type in the database.
	Label = "identifier"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCalibreID holds the string denoting the calibre_id field in the database.
	FieldCalibreID = "calibre_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgeBook holds the string denoting the book edge name in mutations.
	EdgeBook = "book"
	// Table holds the table name of the identifier in the database.
	Table = "identifiers"
	// BookTable is the table that holds the book relation/edge.
	BookTable = "identifiers"
	// BookInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	BookInverseTable = "books"
	// BookColumn is the table column denoting the book relation/edge.
	BookColumn = "identifier_book"
)

// Columns holds all SQL columns for identifier fields.
var Columns = []string{
	FieldID,
	FieldCalibreID,
	FieldType,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "identifiers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"identifier_book",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ksuid.ID
)

// OrderOption defines the ordering options for the Identifier queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCalibreID orders the results by the calibre_id field.
func ByCalibreID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCalibreID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByBookField orders the results by book field.
func ByBookField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookStep(), sql.OrderByField(field, opts...))
	}
}
func newBookStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BookTable, BookColumn),
	)
}
