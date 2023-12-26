// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/tag"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID ksuid.ID `json:"id,omitempty"`
	// CalibreID holds the value of the "calibre_id" field.
	CalibreID int64 `json:"calibre_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TagQuery when eager-loading is set.
	Edges        TagEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TagEdges holds the relations/edges for other nodes in the graph.
type TagEdges struct {
	// Books holds the value of the books edge.
	Books []*Book `json:"books,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedBooks map[string][]*Book
}

// BooksOrErr returns the Books value or an error if the edge
// was not loaded in eager-loading.
func (e TagEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[0] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldCalibreID:
			values[i] = new(sql.NullInt64)
		case tag.FieldID, tag.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				t.ID = ksuid.ID(value.String)
			}
		case tag.FieldCalibreID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field calibre_id", values[i])
			} else if value.Valid {
				t.CalibreID = value.Int64
			}
		case tag.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tag.
// This includes values selected through modifiers, order, etc.
func (t *Tag) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryBooks queries the "books" edge of the Tag entity.
func (t *Tag) QueryBooks() *BookQuery {
	return NewTagClient(t.config).QueryBooks(t)
}

// Update returns a builder for updating this Tag.
// Note that you need to call Tag.Unwrap() before calling this method if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return NewTagClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("calibre_id=")
	builder.WriteString(fmt.Sprintf("%v", t.CalibreID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (t *Tag) NamedBooks(name string) ([]*Book, error) {
	if t.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := t.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (t *Tag) appendNamedBooks(name string, edges ...*Book) {
	if t.Edges.namedBooks == nil {
		t.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		t.Edges.namedBooks[name] = []*Book{}
	} else {
		t.Edges.namedBooks[name] = append(t.Edges.namedBooks[name], edges...)
	}
}

// Tags is a parsable slice of Tag.
type Tags []*Tag
