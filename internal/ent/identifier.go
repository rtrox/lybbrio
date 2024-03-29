// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/identifier"
	"lybbrio/internal/ent/schema/ksuid"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Identifier is the model entity for the Identifier schema.
type Identifier struct {
	config `json:"-"`
	// ID of the ent.
	ID ksuid.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// CalibreID holds the value of the "calibre_id" field.
	CalibreID int64 `json:"calibre_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the IdentifierQuery when eager-loading is set.
	Edges           IdentifierEdges `json:"edges"`
	identifier_book *ksuid.ID
	selectValues    sql.SelectValues
}

// IdentifierEdges holds the relations/edges for other nodes in the graph.
type IdentifierEdges struct {
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
func (e IdentifierEdges) BookOrErr() (*Book, error) {
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
func (*Identifier) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case identifier.FieldCalibreID:
			values[i] = new(sql.NullInt64)
		case identifier.FieldID, identifier.FieldType, identifier.FieldValue:
			values[i] = new(sql.NullString)
		case identifier.FieldCreateTime, identifier.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case identifier.ForeignKeys[0]: // identifier_book
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Identifier fields.
func (i *Identifier) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case identifier.FieldID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value.Valid {
				i.ID = ksuid.ID(value.String)
			}
		case identifier.FieldCreateTime:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[j])
			} else if value.Valid {
				i.CreateTime = value.Time
			}
		case identifier.FieldUpdateTime:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[j])
			} else if value.Valid {
				i.UpdateTime = value.Time
			}
		case identifier.FieldCalibreID:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field calibre_id", values[j])
			} else if value.Valid {
				i.CalibreID = value.Int64
			}
		case identifier.FieldType:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[j])
			} else if value.Valid {
				i.Type = value.String
			}
		case identifier.FieldValue:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[j])
			} else if value.Valid {
				i.Value = value.String
			}
		case identifier.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identifier_book", values[j])
			} else if value.Valid {
				i.identifier_book = new(ksuid.ID)
				*i.identifier_book = ksuid.ID(value.String)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Identifier.
// This includes values selected through modifiers, order, etc.
func (i *Identifier) GetValue(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryBook queries the "book" edge of the Identifier entity.
func (i *Identifier) QueryBook() *BookQuery {
	return NewIdentifierClient(i.config).QueryBook(i)
}

// Update returns a builder for updating this Identifier.
// Note that you need to call Identifier.Unwrap() before calling this method if this Identifier
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Identifier) Update() *IdentifierUpdateOne {
	return NewIdentifierClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Identifier entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Identifier) Unwrap() *Identifier {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Identifier is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Identifier) String() string {
	var builder strings.Builder
	builder.WriteString("Identifier(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("create_time=")
	builder.WriteString(i.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(i.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("calibre_id=")
	builder.WriteString(fmt.Sprintf("%v", i.CalibreID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(i.Type)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(i.Value)
	builder.WriteByte(')')
	return builder.String()
}

// Identifiers is a parsable slice of Identifier.
type Identifiers []*Identifier
