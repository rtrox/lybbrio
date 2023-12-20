package schema

import (
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		// TODO: MutationCreate should not be allowed as-is through the graph long-term,
		// but it's useful for testing right now.
		entgql.Mutations(entgql.MutationUpdate(), entgql.MutationCreate()),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ksuid.MixinWithPrefix("usr"),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			NotEmpty().
			Unique().
			Annotations(
				entgql.OrderField("USERNAME"),
			),
		field.String("passwordHash").
			NotEmpty().
			Sensitive(),
		field.String("email").
			NotEmpty().
			Unique().
			Annotations(
				entgql.OrderField("EMAIL"),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("shelves", Shelf.Type).Annotations(entgql.RelayConnection()),
	}
}
