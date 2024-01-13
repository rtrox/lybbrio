package schema

import (
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Publisher holds the schema definition for the Publisher entity.
type Publisher struct {
	ent.Schema
}

func (Publisher) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

func (Publisher) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
		CalibreMixin{},
		ksuid.MixinWithPrefix("pub"),
	}
}

// Fields of the Publisher.
func (Publisher) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
	}
}

// Edges of the Publisher.
func (Publisher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type).
			Annotations(
				entgql.RelayConnection(),
				entgql.OrderField("BOOKS_COUNT"),
			),
	}
}

// Indexes of the Publisher.
func (Publisher) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}
