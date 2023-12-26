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

// Identifier holds the schema definition for the Identifier entity.
type Identifier struct {
	ent.Schema
}

func (Identifier) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

func (Identifier) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		CalibreMixin{},
		ksuid.MixinWithPrefix("idn"),
	}
}

// Fields of the Identifier.
func (Identifier) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").
			NotEmpty().
			Annotations(entgql.OrderField("TYPE")), // TODO: Enum Values?
		field.String("value").
			NotEmpty().
			Annotations(entgql.OrderField("VALUE")),
	}
}

// Edges of the Identifier.
func (Identifier) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("book", Book.Type).Unique().Required(),
	}
}

func (Identifier) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "value").Unique(),
	}
}
