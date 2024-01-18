package schema

import (
	"lybbrio/internal/ent/privacy"
	"lybbrio/internal/ent/schema/argon2id"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/rule"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
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
		field.String("password_hash").
			GoType(argon2id.Argon2IDHash{}).
			Optional().
			Sensitive(),
		field.String("email").
			NotEmpty().
			Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shelves", Shelf.Type).
			Ref("user"),
		edge.To("user_permissions", UserPermissions.Type).
			Unique().
			Required().
			Immutable(),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Does Not use base policy mixin!! Tread with Care.
			rule.DenyIfNoViewer(),
			rule.AllowCreate(),
			rule.DenyIfAnonymousViewer(),
			rule.AllowIfAdmin(),
			rule.FilterSelfRule(),
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			// Does Not use base policy mixin!! Tread with Care.
			rule.DenyIfNoViewer(),
			rule.DenyIfAnonymousViewer(),
			privacy.AlwaysAllowRule(),
		},
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("password_hash"),
	}
}
