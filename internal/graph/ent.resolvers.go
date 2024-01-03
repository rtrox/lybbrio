package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/graph/generated"

	"entgo.io/contrib/entgql"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id ksuid.ID) (ent.Noder, error) {
	return r.client.Noder(ctx, id, ent.WithNodeType(ent.IDToType))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []ksuid.ID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids, ent.WithNodeType(ent.IDToType))
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.AuthorOrder, where *ent.AuthorWhereInput) (*ent.AuthorConnection, error) {
	return r.client.Author.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithAuthorOrder(orderBy),
			ent.WithAuthorFilter(where.Filter),
		)
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.BookOrder, where *ent.BookWhereInput) (*ent.BookConnection, error) {
	return r.client.Book.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithBookOrder(orderBy),
			ent.WithBookFilter(where.Filter),
		)
}

// Identifiers is the resolver for the identifiers field.
func (r *queryResolver) Identifiers(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.IdentifierOrder, where *ent.IdentifierWhereInput) (*ent.IdentifierConnection, error) {
	return r.client.Identifier.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithIdentifierOrder(orderBy),
			ent.WithIdentifierFilter(where.Filter),
		)
}

// Languages is the resolver for the languages field.
func (r *queryResolver) Languages(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.LanguageOrder, where *ent.LanguageWhereInput) (*ent.LanguageConnection, error) {
	return r.client.Language.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithLanguageOrder(orderBy),
			ent.WithLanguageFilter(where.Filter),
		)
}

// Publishers is the resolver for the publishers field.
func (r *queryResolver) Publishers(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.PublisherOrder, where *ent.PublisherWhereInput) (*ent.PublisherConnection, error) {
	return r.client.Publisher.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithPublisherOrder(orderBy),
			ent.WithPublisherFilter(where.Filter),
		)
}

// SeriesSlice is the resolver for the seriesSlice field.
func (r *queryResolver) SeriesSlice(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.SeriesOrder, where *ent.SeriesWhereInput) (*ent.SeriesConnection, error) {
	return r.client.Series.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithSeriesOrder(orderBy),
			ent.WithSeriesFilter(where.Filter),
		)
}

// Shelves is the resolver for the shelves field.
func (r *queryResolver) Shelves(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.ShelfOrder, where *ent.ShelfWhereInput) (*ent.ShelfConnection, error) {
	return r.client.Shelf.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithShelfOrder(orderBy),
			ent.WithShelfFilter(where.Filter),
		)
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.TagOrder, where *ent.TagWhereInput) (*ent.TagConnection, error) {
	return r.client.Tag.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTagOrder(orderBy),
			ent.WithTagFilter(where.Filter),
		)
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.TaskOrder, where *ent.TaskWhereInput) (*ent.TaskConnection, error) {
	return r.client.Task.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTaskOrder(orderBy),
			ent.WithTaskFilter(where.Filter),
		)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
