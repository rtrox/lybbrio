// Code generated by ent, DO NOT EDIT.

package bookfile

import (
	"fmt"
	"io"
	"lybbrio/internal/ent/schema/ksuid"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the bookfile type in the database.
	Label = "book_file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldFormat holds the string denoting the format field in the database.
	FieldFormat = "format"
	// EdgeBook holds the string denoting the book edge name in mutations.
	EdgeBook = "book"
	// Table holds the table name of the bookfile in the database.
	Table = "book_files"
	// BookTable is the table that holds the book relation/edge.
	BookTable = "book_files"
	// BookInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	BookInverseTable = "books"
	// BookColumn is the table column denoting the book relation/edge.
	BookColumn = "book_file_book"
)

// Columns holds all SQL columns for bookfile fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldPath,
	FieldSize,
	FieldFormat,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "book_files"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"book_file_book",
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// SizeValidator is a validator for the "size" field. It is called by the builders before save.
	SizeValidator func(int64) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ksuid.ID
)

// Format defines the type for the "format" enum field.
type Format string

// Format values.
const (
	FormatAZW3  Format = "AZW3"
	FormatEPUB  Format = "EPUB"
	FormatKEPUB Format = "KEPUB"
	FormatPDF   Format = "PDF"
	FormatCBC   Format = "CBC"
	FormatCBR   Format = "CBR"
	FormatCB7   Format = "CB7"
	FormatCBZ   Format = "CBZ"
	FormatCBT   Format = "CBT"
)

func (f Format) String() string {
	return string(f)
}

// FormatValidator is a validator for the "format" field enum values. It is called by the builders before save.
func FormatValidator(f Format) error {
	switch f {
	case FormatAZW3, FormatEPUB, FormatKEPUB, FormatPDF, FormatCBC, FormatCBR, FormatCB7, FormatCBZ, FormatCBT:
		return nil
	default:
		return fmt.Errorf("bookfile: invalid enum value for format field: %q", f)
	}
}

// OrderOption defines the ordering options for the BookFile queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByFormat orders the results by the format field.
func ByFormat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFormat, opts...).ToFunc()
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

// MarshalGQL implements graphql.Marshaler interface.
func (e Format) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Format) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Format(str)
	if err := FormatValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Format", str)
	}
	return nil
}
