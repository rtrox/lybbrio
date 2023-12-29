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

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

func (Language) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

func (Language) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		CalibreMixin{},
		ksuid.MixinWithPrefix("lng"),
	}
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").NotEmpty(),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type).
			Annotations(
				entgql.RelayConnection(),
				entgql.OrderField("BOOKS_COUNT"),
			),
	}
}

func (Language) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
	}
}
