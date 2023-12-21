package schema

import (
	"lybbrio/internal/ent/privacy"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/rule"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin for all schemas in the graph.
type BaseMixin struct {
	mixin.Schema
}

// Policy defines the privacy policy of the BaseMixin.
func (BaseMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			// Deny any operation in case there is no "viewer context".
			rule.DenyIfNoViewer(),
			// Allow admins to query any information.
			rule.AllowIfAdmin(),
		},
		Mutation: privacy.MutationPolicy{
			// Deny any operation in case there is no "viewer context".
			rule.DenyIfNoViewer(),
		},
	}
}

type UserScopedMixin struct {
	mixin.Schema
}

func (UserScopedMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			GoType(ksuid.ID("")).
			Immutable(),
	}
}

func (UserScopedMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required().
			Immutable(),
	}
}

func (UserScopedMixin) Policy() ent.Policy {
	return rule.FilterUserRule()
}
