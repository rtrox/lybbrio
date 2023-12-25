package schema

import (
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/task"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
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
			GoType(task.TaskType("")).
			Default(string(task.TypeNoOp)).
			Annotations(
				entgql.OrderField("TYPE"),
			),
		field.Enum("status").
			GoType(task.Status("")).
			Default(string(task.StatusPending)).
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
		field.String("createdBy").
			GoType(ksuid.ID("")).
			Optional().
			Immutable().
			Comment("The user who created this task. Empty for System Task"),
		field.Bool("isSystemTask").
			Default(false).
			Comment("Whether this task is created by the system"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("creator", User.Type).
			Field("createdBy").
			Unique().
			Immutable(),
	}
}
