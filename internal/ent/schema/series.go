package schema

import (
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Series holds the schema definition for the Series entity.
type Series struct {
	ent.Schema
}

func (Series) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

func (Series) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ksuid.MixinWithPrefix("srs"),
	}
}

// Fields of the Series.
func (Series) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("sort").NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
	}
}

// Edges of the Series.
func (Series) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type).
			Through("series_books", SeriesBook.Type).
			Annotations(entgql.RelayConnection()),
	}
}

func (Series) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}
