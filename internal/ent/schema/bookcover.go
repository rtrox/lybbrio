package schema

import (
	"context"
	"fmt"
	"lybbrio/internal/ent/hook"
	"lybbrio/internal/ent/privacy"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/rule"
	"mime"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BookCover holds the schema definition for the BookCover entity.
type BookCover struct {
	ent.Schema
}

func (BookCover) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}

func (BookCover) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		BaseMixin{},
		ksuid.MixinWithPrefix("cvr"),
	}
}

// Fields of the BookCover.
func (BookCover) Fields() []ent.Field {
	return []ent.Field{
		field.Text("path").
			NotEmpty().
			Immutable().
			Unique(),
		field.Int64("size").
			Positive().
			Comment("Size in bytes").
			Immutable().
			Annotations(entgql.OrderField("SIZE")),
		field.Int("width").
			Positive().
			Immutable().
			Comment("Width in pixels"),
		field.Int("height").
			Positive().
			Immutable().
			Comment("Height in pixels"),
		field.String("url").
			NotEmpty().
			Immutable().
			Comment("URL to the image"),
		field.String("contentType").
			NotEmpty().
			Immutable().
			Comment("MIME type"),
	}
}

// Edges of the BookCover.
func (BookCover) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("book", Book.Type).
			Unique().
			Immutable().
			Required(),
	}
}

func (BookCover) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Deny any operation in case there is no "viewer context".
			rule.DenyIfNoViewer(),
			rule.DenyIfAnonymousViewer(),
			rule.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}

func (BookCover) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				type CoverMutation interface {
					Width() (int, bool)
					Height() (int, bool)
					ContentType() (string, bool)
					BookID() (ksuid.ID, bool)
					SetURL(string)
				}
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					mx, ok := m.(CoverMutation)
					if !ok {
						return nil, fmt.Errorf("unexpected mutation type")
					}
					width, wok := mx.Width()
					height, hok := mx.Height()
					contentType, cok := mx.ContentType()
					bookID, bok := mx.BookID()
					if !wok || !hok || !cok || !bok {
						return nil, fmt.Errorf("missing required fields")
					}
					exts, err := mime.ExtensionsByType(contentType)
					if err != nil {
						return nil, fmt.Errorf("unknown content type")
					}
					ext := exts[len(exts)-1]
					mx.SetURL(fmt.Sprintf("/image/%s/%d/%d/cover%s", bookID.String(), width, height, ext))
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	}
}
