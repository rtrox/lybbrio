package schema

import (
	"lybbrio/internal/ent/schema/ksuid"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

func (Book) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}

func (Book) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ksuid.MixinWithPrefix("bok"),
	}
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.Text("title").
			NotEmpty().
			Annotations(entgql.OrderField("TITLE")),
		field.Text("sort").
			Annotations(entgql.OrderField("SORT")),
		field.Time("addedAt").
			Default(time.Now).
			Annotations(entgql.OrderField("ADDED_AT")),
		field.Time("pubDate").
			Optional().
			Annotations(entgql.OrderField("PUB_DATE")),
		field.Text("path").
			NotEmpty(),
		// Unique(),
		field.Text("isbn").
			Optional().
			Annotations(entgql.OrderField("ISBN")),
		field.Text("description").
			Optional(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("authors", Author.Type).
			Ref("books"),
	}
}
