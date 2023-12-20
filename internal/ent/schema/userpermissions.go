package schema

import (
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserPermissions holds the schema definition for the UserPermissions entity.
type UserPermissions struct {
	ent.Schema
}

func (UserPermissions) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ksuid.MixinWithPrefix("prm"),
	}
}

// Fields of the UserPermissions.
func (UserPermissions) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("admin").Default(false),
	}
}

// Edges of the UserPermissions.
func (UserPermissions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("userPermissions").Unique(),
	}
}
