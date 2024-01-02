package schema

import (
	"lybbrio/internal/ent/privacy"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/task_enums"
	"lybbrio/internal/rule"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
	}
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		mixin.Time{},
		ksuid.MixinWithPrefix("tsk"),
	}
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			GoType(task_enums.TaskType("")).
			Default(string(task_enums.TypeNoOp)).
			Immutable().
			Annotations(
				entgql.OrderField("TYPE"),
			),
		field.Enum("status").
			GoType(task_enums.Status("")).
			Default(string(task_enums.StatusPending)).
			Annotations(
				entgql.OrderField("STATUS"),
			),
		field.Float("progress").
			Default(0).
			Comment("Progress of the task. 0-1"),
		field.String("message").
			Optional().
			Comment("Message of the task"),
		field.String("error").
			Optional().
			Comment("Error message of the task"),
		field.String("user_id").
			GoType(ksuid.ID("")).
			Optional().
			Immutable().
			Comment("The user who created this task. Empty for System Task"),
		field.Bool("is_system_task").
			Default(false).
			Comment("Whether this task is created by the system"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Immutable(),
	}
}

// TODO: More detailed privacy rules
func (Task) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.AllowIfAdmin(),
			rule.FilterUserOrSystemRule(),
			privacy.AlwaysAllowRule(),
		},
		Mutation: privacy.MutationPolicy{
			rule.AllowIfAdmin(),
			rule.DenySystemTaskForNonAdmin(),
			rule.DenyMismatchedUserRule(),
			rule.AllowTasksOfType(task_enums.TypeNoOp),
			// Technically Unreachable, AllowTasksOfType will always allow or deny
			privacy.AlwaysDenyRule(),
		},
	}
}
