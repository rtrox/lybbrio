// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/bookcover"
	"lybbrio/internal/ent/schema/ksuid"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// BookCover is the model entity for the BookCover schema.
type BookCover struct {
	config `json:"-"`
	// ID of the ent.
	ID ksuid.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Size in bytes
	Size int64 `json:"size,omitempty"`
	// Width in pixels
	Width int `json:"width,omitempty"`
	// Height in pixels
	Height int `json:"height,omitempty"`
	// URL to the image
	URL string `json:"url,omitempty"`
	// MIME type
	ContentType string `json:"contentType,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookCoverQuery when eager-loading is set.
	Edges           BookCoverEdges `json:"edges"`
	book_cover_book *ksuid.ID
	selectValues    sql.SelectValues
}

// BookCoverEdges holds the relations/edges for other nodes in the graph.
type BookCoverEdges struct {
	// Book holds the value of the book edge.
	Book *Book `json:"book,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// BookOrErr returns the Book value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookCoverEdges) BookOrErr() (*Book, error) {
	if e.loadedTypes[0] {
		if e.Book == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: book.Label}
		}
		return e.Book, nil
	}
	return nil, &NotLoadedError{edge: "book"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BookCover) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bookcover.FieldSize, bookcover.FieldWidth, bookcover.FieldHeight:
			values[i] = new(sql.NullInt64)
		case bookcover.FieldID, bookcover.FieldPath, bookcover.FieldURL, bookcover.FieldContentType:
			values[i] = new(sql.NullString)
		case bookcover.FieldCreateTime, bookcover.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case bookcover.ForeignKeys[0]: // book_cover_book
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BookCover fields.
func (bc *BookCover) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bookcover.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				bc.ID = ksuid.ID(value.String)
			}
		case bookcover.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				bc.CreateTime = value.Time
			}
		case bookcover.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				bc.UpdateTime = value.Time
			}
		case bookcover.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				bc.Path = value.String
			}
		case bookcover.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				bc.Size = value.Int64
			}
		case bookcover.FieldWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field width", values[i])
			} else if value.Valid {
				bc.Width = int(value.Int64)
			}
		case bookcover.FieldHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				bc.Height = int(value.Int64)
			}
		case bookcover.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				bc.URL = value.String
			}
		case bookcover.FieldContentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contentType", values[i])
			} else if value.Valid {
				bc.ContentType = value.String
			}
		case bookcover.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field book_cover_book", values[i])
			} else if value.Valid {
				bc.book_cover_book = new(ksuid.ID)
				*bc.book_cover_book = ksuid.ID(value.String)
			}
		default:
			bc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BookCover.
// This includes values selected through modifiers, order, etc.
func (bc *BookCover) Value(name string) (ent.Value, error) {
	return bc.selectValues.Get(name)
}

// QueryBook queries the "book" edge of the BookCover entity.
func (bc *BookCover) QueryBook() *BookQuery {
	return NewBookCoverClient(bc.config).QueryBook(bc)
}

// Update returns a builder for updating this BookCover.
// Note that you need to call BookCover.Unwrap() before calling this method if this BookCover
// was returned from a transaction, and the transaction was committed or rolled back.
func (bc *BookCover) Update() *BookCoverUpdateOne {
	return NewBookCoverClient(bc.config).UpdateOne(bc)
}

// Unwrap unwraps the BookCover entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bc *BookCover) Unwrap() *BookCover {
	_tx, ok := bc.config.driver.(*txDriver)
	if !ok {
		panic("ent: BookCover is not a transactional entity")
	}
	bc.config.driver = _tx.drv
	return bc
}

// String implements the fmt.Stringer.
func (bc *BookCover) String() string {
	var builder strings.Builder
	builder.WriteString("BookCover(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bc.ID))
	builder.WriteString("create_time=")
	builder.WriteString(bc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(bc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(bc.Path)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", bc.Size))
	builder.WriteString(", ")
	builder.WriteString("width=")
	builder.WriteString(fmt.Sprintf("%v", bc.Width))
	builder.WriteString(", ")
	builder.WriteString("height=")
	builder.WriteString(fmt.Sprintf("%v", bc.Height))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(bc.URL)
	builder.WriteString(", ")
	builder.WriteString("contentType=")
	builder.WriteString(bc.ContentType)
	builder.WriteByte(')')
	return builder.String()
}

// BookCovers is a parsable slice of BookCover.
type BookCovers []*BookCover