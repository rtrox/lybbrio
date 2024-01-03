package schema

import (
	"lybbrio/internal/ent/privacy"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/rule"

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
	fields := []ent.Field{
		field.String("user_id").
			GoType(ksuid.ID("")).
			Optional(),
	}
	for p := range permissions.All() {
		fields = append(fields,
			field.Bool(p.String()).
				Default(false),
		)
	}

	return fields
}

// Edges of the UserPermissions.
func (UserPermissions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Field("user_id").Ref("user_permissions").Unique(),
	}
}

func (UserPermissions) Policy() ent.Policy {
	// Users can query their own permissions, but not modify them.
	// Admins can query and modify any user's permissions.
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfAdmin(),
			rule.FilterUserRule(),
			privacy.AlwaysAllowRule(),
		},
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
