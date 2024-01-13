package schema

import (
	"lybbrio/internal/ent/privacy"
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Shelf holds the schema definition for the Shelf entity.
type Shelf struct {
	ent.Schema
}

func (Shelf) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

func (Shelf) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
		PublicableUserScopedMixin{},
		ksuid.MixinWithPrefix("shf"),
	}
}

// Fields of the Shelf.
func (Shelf) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("description").
			Optional(),
	}
}

// Edges of the Shelf.
func (Shelf) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type).
			Annotations(
				entgql.OrderField("BOOKS_COUNT"),
				entgql.RelayConnection(),
			),
	}
}

func (Shelf) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Shelf) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}
