package schema

import (
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SeriesBook holds the schema definition for the SeriesBook entity.
type SeriesBook struct {
	ent.Schema
}

func (SeriesBook) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the SeriesBook.
func (SeriesBook) Fields() []ent.Field {
	return []ent.Field{
		field.Float("series_index").Positive().Optional(),
		field.String("series_id").GoType(ksuid.ID("")).NotEmpty(),
		field.String("book_id").GoType(ksuid.ID("")).NotEmpty(),
	}
}

// Edges of the SeriesBook.
func (SeriesBook) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("series", Series.Type).Unique().Required().Field("series_id"),
		edge.To("book", Book.Type).Unique().Required().Field("book_id"),
	}
}
