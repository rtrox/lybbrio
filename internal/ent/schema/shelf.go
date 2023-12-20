package schema

import (
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Shelf holds the schema definition for the Shelf entity.
type Shelf struct {
	ent.Schema
}

func (Shelf) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

func (Shelf) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
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
		field.Bool("public").
			Default(false),
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
		edge.From("user", User.Type).
			Unique().
			Ref("shelves").
			Required().
			Immutable().
			Annotations(entgql.OrderField("USER_USERNAME")),
	}
}
