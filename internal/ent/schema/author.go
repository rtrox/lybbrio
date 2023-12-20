package schema

import (
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Author holds the schema definition for the Author entity.
type Author struct {
	ent.Schema
}

func (Author) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

func (Author) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ksuid.MixinWithPrefix("atr"),
	}
}

// Fields of the Author.
func (Author) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(entgql.OrderField("NAME")),
		field.String("sort").
			Annotations(entgql.OrderField("SORT")),
		field.String("link").
			Optional(),
	}
}

// Edges of the Author.
func (Author) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type).
			Annotations(entgql.RelayConnection()),
	}
}
