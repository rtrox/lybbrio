// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"lybbrio/internal/ent/publisher"
	"lybbrio/internal/ent/schema/ksuid"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Publisher is the model entity for the Publisher schema.
type Publisher struct {
	config `json:"-"`
	// ID of the ent.
	ID ksuid.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// CalibreID holds the value of the "calibre_id" field.
	CalibreID int64 `json:"calibre_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PublisherQuery when eager-loading is set.
	Edges        PublisherEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PublisherEdges holds the relations/edges for other nodes in the graph.
type PublisherEdges struct {
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
func (e PublisherEdges) BooksOrErr() ([]*Book, error) {
	if e.loadedTypes[0] {
		return e.Books, nil
	}
	return nil, &NotLoadedError{edge: "books"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Publisher) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case publisher.FieldCalibreID:
			values[i] = new(sql.NullInt64)
		case publisher.FieldID, publisher.FieldName:
			values[i] = new(sql.NullString)
		case publisher.FieldCreateTime, publisher.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Publisher fields.
func (pu *Publisher) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case publisher.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pu.ID = ksuid.ID(value.String)
			}
		case publisher.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pu.CreateTime = value.Time
			}
		case publisher.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pu.UpdateTime = value.Time
			}
		case publisher.FieldCalibreID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field calibre_id", values[i])
			} else if value.Valid {
				pu.CalibreID = value.Int64
			}
		case publisher.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pu.Name = value.String
			}
		default:
			pu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Publisher.
// This includes values selected through modifiers, order, etc.
func (pu *Publisher) Value(name string) (ent.Value, error) {
	return pu.selectValues.Get(name)
}

// QueryBooks queries the "books" edge of the Publisher entity.
func (pu *Publisher) QueryBooks() *BookQuery {
	return NewPublisherClient(pu.config).QueryBooks(pu)
}

// Update returns a builder for updating this Publisher.
// Note that you need to call Publisher.Unwrap() before calling this method if this Publisher
// was returned from a transaction, and the transaction was committed or rolled back.
func (pu *Publisher) Update() *PublisherUpdateOne {
	return NewPublisherClient(pu.config).UpdateOne(pu)
}

// Unwrap unwraps the Publisher entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pu *Publisher) Unwrap() *Publisher {
	_tx, ok := pu.config.driver.(*txDriver)
	if !ok {
		panic("ent: Publisher is not a transactional entity")
	}
	pu.config.driver = _tx.drv
	return pu
}

// String implements the fmt.Stringer.
func (pu *Publisher) String() string {
	var builder strings.Builder
	builder.WriteString("Publisher(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pu.ID))
	builder.WriteString("create_time=")
	builder.WriteString(pu.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(pu.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("calibre_id=")
	builder.WriteString(fmt.Sprintf("%v", pu.CalibreID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pu.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedBooks returns the Books named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pu *Publisher) NamedBooks(name string) ([]*Book, error) {
	if pu.Edges.namedBooks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pu.Edges.namedBooks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pu *Publisher) appendNamedBooks(name string, edges ...*Book) {
	if pu.Edges.namedBooks == nil {
		pu.Edges.namedBooks = make(map[string][]*Book)
	}
	if len(edges) == 0 {
		pu.Edges.namedBooks[name] = []*Book{}
	} else {
		pu.Edges.namedBooks[name] = append(pu.Edges.namedBooks[name], edges...)
	}
}

// Publishers is a parsable slice of Publisher.
type Publishers []*Publisher
