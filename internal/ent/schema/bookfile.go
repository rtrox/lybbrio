package schema

import (
	"lybbrio/internal/ent/schema/filetype"
	"lybbrio/internal/ent/schema/ksuid"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BookFile holds the schema definition for the BookFile entity.
type BookFile struct {
	ent.Schema
}

func (BookFile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

func (BookFile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ksuid.MixinWithPrefix("fil"),
	}
}

// Fields of the BookFile.
func (BookFile) Fields() []ent.Field {
	ret := []ent.Field{
		field.Text("path").
			NotEmpty(),
		field.Int64("size").
			Positive().
			Comment("Size in bytes"),
	}
	values := []string{}
	for _, format := range filetype.All() {
		values = append(values, format.String())
	}
	return append(ret, field.Enum("format").
		Values(values...))
}

// Edges of the BookFile.
func (BookFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("book", Book.Type).
			Unique().
			Required(),
	}
}
