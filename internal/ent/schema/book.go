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
		BaseMixin{},
		CalibreMixin{},
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
			Annotations(entgql.OrderField("NAME")),
		field.Time("published_date").
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
		field.Float("series_index").
			Optional(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("authors", Author.Type).
			Ref("books"),
		edge.From("publisher", Publisher.Type).
			Ref("books"),
		edge.From("series", Series.Type).
			Ref("books"),
		edge.From("identifiers", Identifier.Type).
			Ref("book"),
		edge.From("tags", Tag.Type).
			Ref("books"),
		edge.From("language", Language.Type).
			Ref("books"),
		edge.From("shelf", Shelf.Type).
			Ref("books"), // TODO: will need privacy on this edge.
	}
}

func (Book) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title"),
		index.Fields("sort"),
		index.Fields("published_date"),
		index.Fields("isbn"),
	}
}
