// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lybbrio/internal/ent/author"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/identifier"
	"lybbrio/internal/ent/language"
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/series"
	"lybbrio/internal/ent/shelf"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookUpdate is the builder for updating Book entities.
type BookUpdate struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// Where appends a list predicates to the BookUpdate builder.
func (bu *BookUpdate) Where(ps ...predicate.Book) *BookUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BookUpdate) SetTitle(s string) *BookUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bu *BookUpdate) SetNillableTitle(s *string) *BookUpdate {
	if s != nil {
		bu.SetTitle(*s)
	}
	return bu
}

// SetSort sets the "sort" field.
func (bu *BookUpdate) SetSort(s string) *BookUpdate {
	bu.mutation.SetSort(s)
	return bu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (bu *BookUpdate) SetNillableSort(s *string) *BookUpdate {
	if s != nil {
		bu.SetSort(*s)
	}
	return bu
}

// SetAddedAt sets the "added_at" field.
func (bu *BookUpdate) SetAddedAt(t time.Time) *BookUpdate {
	bu.mutation.SetAddedAt(t)
	return bu
}

// SetNillableAddedAt sets the "added_at" field if the given value is not nil.
func (bu *BookUpdate) SetNillableAddedAt(t *time.Time) *BookUpdate {
	if t != nil {
		bu.SetAddedAt(*t)
	}
	return bu
}

// SetPubDate sets the "pub_date" field.
func (bu *BookUpdate) SetPubDate(t time.Time) *BookUpdate {
	bu.mutation.SetPubDate(t)
	return bu
}

// SetNillablePubDate sets the "pub_date" field if the given value is not nil.
func (bu *BookUpdate) SetNillablePubDate(t *time.Time) *BookUpdate {
	if t != nil {
		bu.SetPubDate(*t)
	}
	return bu
}

// ClearPubDate clears the value of the "pub_date" field.
func (bu *BookUpdate) ClearPubDate() *BookUpdate {
	bu.mutation.ClearPubDate()
	return bu
}

// SetPath sets the "path" field.
func (bu *BookUpdate) SetPath(s string) *BookUpdate {
	bu.mutation.SetPath(s)
	return bu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (bu *BookUpdate) SetNillablePath(s *string) *BookUpdate {
	if s != nil {
		bu.SetPath(*s)
	}
	return bu
}

// SetIsbn sets the "isbn" field.
func (bu *BookUpdate) SetIsbn(s string) *BookUpdate {
	bu.mutation.SetIsbn(s)
	return bu
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (bu *BookUpdate) SetNillableIsbn(s *string) *BookUpdate {
	if s != nil {
		bu.SetIsbn(*s)
	}
	return bu
}

// ClearIsbn clears the value of the "isbn" field.
func (bu *BookUpdate) ClearIsbn() *BookUpdate {
	bu.mutation.ClearIsbn()
	return bu
}

// SetDescription sets the "description" field.
func (bu *BookUpdate) SetDescription(s string) *BookUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bu *BookUpdate) SetNillableDescription(s *string) *BookUpdate {
	if s != nil {
		bu.SetDescription(*s)
	}
	return bu
}

// ClearDescription clears the value of the "description" field.
func (bu *BookUpdate) ClearDescription() *BookUpdate {
	bu.mutation.ClearDescription()
	return bu
}

// AddAuthorIDs adds the "authors" edge to the Author entity by IDs.
func (bu *BookUpdate) AddAuthorIDs(ids ...ksuid.ID) *BookUpdate {
	bu.mutation.AddAuthorIDs(ids...)
	return bu
}

// AddAuthors adds the "authors" edges to the Author entity.
func (bu *BookUpdate) AddAuthors(a ...*Author) *BookUpdate {
	ids := make([]ksuid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bu.AddAuthorIDs(ids...)
}

// AddSeriesIDs adds the "series" edge to the Series entity by IDs.
func (bu *BookUpdate) AddSeriesIDs(ids ...ksuid.ID) *BookUpdate {
	bu.mutation.AddSeriesIDs(ids...)
	return bu
}

// AddSeries adds the "series" edges to the Series entity.
func (bu *BookUpdate) AddSeries(s ...*Series) *BookUpdate {
	ids := make([]ksuid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bu.AddSeriesIDs(ids...)
}

// AddIdentifierIDs adds the "identifier" edge to the Identifier entity by IDs.
func (bu *BookUpdate) AddIdentifierIDs(ids ...ksuid.ID) *BookUpdate {
	bu.mutation.AddIdentifierIDs(ids...)
	return bu
}

// AddIdentifier adds the "identifier" edges to the Identifier entity.
func (bu *BookUpdate) AddIdentifier(i ...*Identifier) *BookUpdate {
	ids := make([]ksuid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return bu.AddIdentifierIDs(ids...)
}

// SetLanguageID sets the "language" edge to the Language entity by ID.
func (bu *BookUpdate) SetLanguageID(id ksuid.ID) *BookUpdate {
	bu.mutation.SetLanguageID(id)
	return bu
}

// SetNillableLanguageID sets the "language" edge to the Language entity by ID if the given value is not nil.
func (bu *BookUpdate) SetNillableLanguageID(id *ksuid.ID) *BookUpdate {
	if id != nil {
		bu = bu.SetLanguageID(*id)
	}
	return bu
}

// SetLanguage sets the "language" edge to the Language entity.
func (bu *BookUpdate) SetLanguage(l *Language) *BookUpdate {
	return bu.SetLanguageID(l.ID)
}

// AddShelfIDs adds the "shelf" edge to the Shelf entity by IDs.
func (bu *BookUpdate) AddShelfIDs(ids ...ksuid.ID) *BookUpdate {
	bu.mutation.AddShelfIDs(ids...)
	return bu
}

// AddShelf adds the "shelf" edges to the Shelf entity.
func (bu *BookUpdate) AddShelf(s ...*Shelf) *BookUpdate {
	ids := make([]ksuid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bu.AddShelfIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (bu *BookUpdate) Mutation() *BookMutation {
	return bu.mutation
}

// ClearAuthors clears all "authors" edges to the Author entity.
func (bu *BookUpdate) ClearAuthors() *BookUpdate {
	bu.mutation.ClearAuthors()
	return bu
}

// RemoveAuthorIDs removes the "authors" edge to Author entities by IDs.
func (bu *BookUpdate) RemoveAuthorIDs(ids ...ksuid.ID) *BookUpdate {
	bu.mutation.RemoveAuthorIDs(ids...)
	return bu
}

// RemoveAuthors removes "authors" edges to Author entities.
func (bu *BookUpdate) RemoveAuthors(a ...*Author) *BookUpdate {
	ids := make([]ksuid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bu.RemoveAuthorIDs(ids...)
}

// ClearSeries clears all "series" edges to the Series entity.
func (bu *BookUpdate) ClearSeries() *BookUpdate {
	bu.mutation.ClearSeries()
	return bu
}

// RemoveSeriesIDs removes the "series" edge to Series entities by IDs.
func (bu *BookUpdate) RemoveSeriesIDs(ids ...ksuid.ID) *BookUpdate {
	bu.mutation.RemoveSeriesIDs(ids...)
	return bu
}

// RemoveSeries removes "series" edges to Series entities.
func (bu *BookUpdate) RemoveSeries(s ...*Series) *BookUpdate {
	ids := make([]ksuid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bu.RemoveSeriesIDs(ids...)
}

// ClearIdentifier clears all "identifier" edges to the Identifier entity.
func (bu *BookUpdate) ClearIdentifier() *BookUpdate {
	bu.mutation.ClearIdentifier()
	return bu
}

// RemoveIdentifierIDs removes the "identifier" edge to Identifier entities by IDs.
func (bu *BookUpdate) RemoveIdentifierIDs(ids ...ksuid.ID) *BookUpdate {
	bu.mutation.RemoveIdentifierIDs(ids...)
	return bu
}

// RemoveIdentifier removes "identifier" edges to Identifier entities.
func (bu *BookUpdate) RemoveIdentifier(i ...*Identifier) *BookUpdate {
	ids := make([]ksuid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return bu.RemoveIdentifierIDs(ids...)
}

// ClearLanguage clears the "language" edge to the Language entity.
func (bu *BookUpdate) ClearLanguage() *BookUpdate {
	bu.mutation.ClearLanguage()
	return bu
}

// ClearShelf clears all "shelf" edges to the Shelf entity.
func (bu *BookUpdate) ClearShelf() *BookUpdate {
	bu.mutation.ClearShelf()
	return bu
}

// RemoveShelfIDs removes the "shelf" edge to Shelf entities by IDs.
func (bu *BookUpdate) RemoveShelfIDs(ids ...ksuid.ID) *BookUpdate {
	bu.mutation.RemoveShelfIDs(ids...)
	return bu
}

// RemoveShelf removes "shelf" edges to Shelf entities.
func (bu *BookUpdate) RemoveShelf(s ...*Shelf) *BookUpdate {
	ids := make([]ksuid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bu.RemoveShelfIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookUpdate) check() error {
	if v, ok := bu.mutation.Title(); ok {
		if err := book.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Book.title": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Path(); ok {
		if err := book.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Book.path": %w`, err)}
		}
	}
	return nil
}

func (bu *BookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeString))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Sort(); ok {
		_spec.SetField(book.FieldSort, field.TypeString, value)
	}
	if value, ok := bu.mutation.AddedAt(); ok {
		_spec.SetField(book.FieldAddedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.PubDate(); ok {
		_spec.SetField(book.FieldPubDate, field.TypeTime, value)
	}
	if bu.mutation.PubDateCleared() {
		_spec.ClearField(book.FieldPubDate, field.TypeTime)
	}
	if value, ok := bu.mutation.Path(); ok {
		_spec.SetField(book.FieldPath, field.TypeString, value)
	}
	if value, ok := bu.mutation.Isbn(); ok {
		_spec.SetField(book.FieldIsbn, field.TypeString, value)
	}
	if bu.mutation.IsbnCleared() {
		_spec.ClearField(book.FieldIsbn, field.TypeString)
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.SetField(book.FieldDescription, field.TypeString, value)
	}
	if bu.mutation.DescriptionCleared() {
		_spec.ClearField(book.FieldDescription, field.TypeString)
	}
	if bu.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedAuthorsIDs(); len(nodes) > 0 && !bu.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.SeriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.SeriesTable,
			Columns: book.SeriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(series.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedSeriesIDs(); len(nodes) > 0 && !bu.mutation.SeriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.SeriesTable,
			Columns: book.SeriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(series.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.SeriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.SeriesTable,
			Columns: book.SeriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(series.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.IdentifierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   book.IdentifierTable,
			Columns: []string{book.IdentifierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(identifier.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedIdentifierIDs(); len(nodes) > 0 && !bu.mutation.IdentifierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   book.IdentifierTable,
			Columns: []string{book.IdentifierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(identifier.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.IdentifierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   book.IdentifierTable,
			Columns: []string{book.IdentifierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(identifier.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.LanguageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.LanguageTable,
			Columns: []string{book.LanguageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.LanguageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.LanguageTable,
			Columns: []string{book.LanguageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.ShelfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.ShelfTable,
			Columns: book.ShelfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shelf.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedShelfIDs(); len(nodes) > 0 && !bu.mutation.ShelfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.ShelfTable,
			Columns: book.ShelfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shelf.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ShelfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.ShelfTable,
			Columns: book.ShelfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shelf.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookUpdateOne is the builder for updating a single Book entity.
type BookUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookMutation
}

// SetTitle sets the "title" field.
func (buo *BookUpdateOne) SetTitle(s string) *BookUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableTitle(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetTitle(*s)
	}
	return buo
}

// SetSort sets the "sort" field.
func (buo *BookUpdateOne) SetSort(s string) *BookUpdateOne {
	buo.mutation.SetSort(s)
	return buo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableSort(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetSort(*s)
	}
	return buo
}

// SetAddedAt sets the "added_at" field.
func (buo *BookUpdateOne) SetAddedAt(t time.Time) *BookUpdateOne {
	buo.mutation.SetAddedAt(t)
	return buo
}

// SetNillableAddedAt sets the "added_at" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableAddedAt(t *time.Time) *BookUpdateOne {
	if t != nil {
		buo.SetAddedAt(*t)
	}
	return buo
}

// SetPubDate sets the "pub_date" field.
func (buo *BookUpdateOne) SetPubDate(t time.Time) *BookUpdateOne {
	buo.mutation.SetPubDate(t)
	return buo
}

// SetNillablePubDate sets the "pub_date" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillablePubDate(t *time.Time) *BookUpdateOne {
	if t != nil {
		buo.SetPubDate(*t)
	}
	return buo
}

// ClearPubDate clears the value of the "pub_date" field.
func (buo *BookUpdateOne) ClearPubDate() *BookUpdateOne {
	buo.mutation.ClearPubDate()
	return buo
}

// SetPath sets the "path" field.
func (buo *BookUpdateOne) SetPath(s string) *BookUpdateOne {
	buo.mutation.SetPath(s)
	return buo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillablePath(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetPath(*s)
	}
	return buo
}

// SetIsbn sets the "isbn" field.
func (buo *BookUpdateOne) SetIsbn(s string) *BookUpdateOne {
	buo.mutation.SetIsbn(s)
	return buo
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableIsbn(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetIsbn(*s)
	}
	return buo
}

// ClearIsbn clears the value of the "isbn" field.
func (buo *BookUpdateOne) ClearIsbn() *BookUpdateOne {
	buo.mutation.ClearIsbn()
	return buo
}

// SetDescription sets the "description" field.
func (buo *BookUpdateOne) SetDescription(s string) *BookUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableDescription(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetDescription(*s)
	}
	return buo
}

// ClearDescription clears the value of the "description" field.
func (buo *BookUpdateOne) ClearDescription() *BookUpdateOne {
	buo.mutation.ClearDescription()
	return buo
}

// AddAuthorIDs adds the "authors" edge to the Author entity by IDs.
func (buo *BookUpdateOne) AddAuthorIDs(ids ...ksuid.ID) *BookUpdateOne {
	buo.mutation.AddAuthorIDs(ids...)
	return buo
}

// AddAuthors adds the "authors" edges to the Author entity.
func (buo *BookUpdateOne) AddAuthors(a ...*Author) *BookUpdateOne {
	ids := make([]ksuid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return buo.AddAuthorIDs(ids...)
}

// AddSeriesIDs adds the "series" edge to the Series entity by IDs.
func (buo *BookUpdateOne) AddSeriesIDs(ids ...ksuid.ID) *BookUpdateOne {
	buo.mutation.AddSeriesIDs(ids...)
	return buo
}

// AddSeries adds the "series" edges to the Series entity.
func (buo *BookUpdateOne) AddSeries(s ...*Series) *BookUpdateOne {
	ids := make([]ksuid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return buo.AddSeriesIDs(ids...)
}

// AddIdentifierIDs adds the "identifier" edge to the Identifier entity by IDs.
func (buo *BookUpdateOne) AddIdentifierIDs(ids ...ksuid.ID) *BookUpdateOne {
	buo.mutation.AddIdentifierIDs(ids...)
	return buo
}

// AddIdentifier adds the "identifier" edges to the Identifier entity.
func (buo *BookUpdateOne) AddIdentifier(i ...*Identifier) *BookUpdateOne {
	ids := make([]ksuid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return buo.AddIdentifierIDs(ids...)
}

// SetLanguageID sets the "language" edge to the Language entity by ID.
func (buo *BookUpdateOne) SetLanguageID(id ksuid.ID) *BookUpdateOne {
	buo.mutation.SetLanguageID(id)
	return buo
}

// SetNillableLanguageID sets the "language" edge to the Language entity by ID if the given value is not nil.
func (buo *BookUpdateOne) SetNillableLanguageID(id *ksuid.ID) *BookUpdateOne {
	if id != nil {
		buo = buo.SetLanguageID(*id)
	}
	return buo
}

// SetLanguage sets the "language" edge to the Language entity.
func (buo *BookUpdateOne) SetLanguage(l *Language) *BookUpdateOne {
	return buo.SetLanguageID(l.ID)
}

// AddShelfIDs adds the "shelf" edge to the Shelf entity by IDs.
func (buo *BookUpdateOne) AddShelfIDs(ids ...ksuid.ID) *BookUpdateOne {
	buo.mutation.AddShelfIDs(ids...)
	return buo
}

// AddShelf adds the "shelf" edges to the Shelf entity.
func (buo *BookUpdateOne) AddShelf(s ...*Shelf) *BookUpdateOne {
	ids := make([]ksuid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return buo.AddShelfIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (buo *BookUpdateOne) Mutation() *BookMutation {
	return buo.mutation
}

// ClearAuthors clears all "authors" edges to the Author entity.
func (buo *BookUpdateOne) ClearAuthors() *BookUpdateOne {
	buo.mutation.ClearAuthors()
	return buo
}

// RemoveAuthorIDs removes the "authors" edge to Author entities by IDs.
func (buo *BookUpdateOne) RemoveAuthorIDs(ids ...ksuid.ID) *BookUpdateOne {
	buo.mutation.RemoveAuthorIDs(ids...)
	return buo
}

// RemoveAuthors removes "authors" edges to Author entities.
func (buo *BookUpdateOne) RemoveAuthors(a ...*Author) *BookUpdateOne {
	ids := make([]ksuid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return buo.RemoveAuthorIDs(ids...)
}

// ClearSeries clears all "series" edges to the Series entity.
func (buo *BookUpdateOne) ClearSeries() *BookUpdateOne {
	buo.mutation.ClearSeries()
	return buo
}

// RemoveSeriesIDs removes the "series" edge to Series entities by IDs.
func (buo *BookUpdateOne) RemoveSeriesIDs(ids ...ksuid.ID) *BookUpdateOne {
	buo.mutation.RemoveSeriesIDs(ids...)
	return buo
}

// RemoveSeries removes "series" edges to Series entities.
func (buo *BookUpdateOne) RemoveSeries(s ...*Series) *BookUpdateOne {
	ids := make([]ksuid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return buo.RemoveSeriesIDs(ids...)
}

// ClearIdentifier clears all "identifier" edges to the Identifier entity.
func (buo *BookUpdateOne) ClearIdentifier() *BookUpdateOne {
	buo.mutation.ClearIdentifier()
	return buo
}

// RemoveIdentifierIDs removes the "identifier" edge to Identifier entities by IDs.
func (buo *BookUpdateOne) RemoveIdentifierIDs(ids ...ksuid.ID) *BookUpdateOne {
	buo.mutation.RemoveIdentifierIDs(ids...)
	return buo
}

// RemoveIdentifier removes "identifier" edges to Identifier entities.
func (buo *BookUpdateOne) RemoveIdentifier(i ...*Identifier) *BookUpdateOne {
	ids := make([]ksuid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return buo.RemoveIdentifierIDs(ids...)
}

// ClearLanguage clears the "language" edge to the Language entity.
func (buo *BookUpdateOne) ClearLanguage() *BookUpdateOne {
	buo.mutation.ClearLanguage()
	return buo
}

// ClearShelf clears all "shelf" edges to the Shelf entity.
func (buo *BookUpdateOne) ClearShelf() *BookUpdateOne {
	buo.mutation.ClearShelf()
	return buo
}

// RemoveShelfIDs removes the "shelf" edge to Shelf entities by IDs.
func (buo *BookUpdateOne) RemoveShelfIDs(ids ...ksuid.ID) *BookUpdateOne {
	buo.mutation.RemoveShelfIDs(ids...)
	return buo
}

// RemoveShelf removes "shelf" edges to Shelf entities.
func (buo *BookUpdateOne) RemoveShelf(s ...*Shelf) *BookUpdateOne {
	ids := make([]ksuid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return buo.RemoveShelfIDs(ids...)
}

// Where appends a list predicates to the BookUpdate builder.
func (buo *BookUpdateOne) Where(ps ...predicate.Book) *BookUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookUpdateOne) Select(field string, fields ...string) *BookUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Book entity.
func (buo *BookUpdateOne) Save(ctx context.Context) (*Book, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookUpdateOne) SaveX(ctx context.Context) *Book {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookUpdateOne) check() error {
	if v, ok := buo.mutation.Title(); ok {
		if err := book.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Book.title": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Path(); ok {
		if err := book.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Book.path": %w`, err)}
		}
	}
	return nil
}

func (buo *BookUpdateOne) sqlSave(ctx context.Context) (_node *Book, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeString))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Book.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, book.FieldID)
		for _, f := range fields {
			if !book.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != book.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Sort(); ok {
		_spec.SetField(book.FieldSort, field.TypeString, value)
	}
	if value, ok := buo.mutation.AddedAt(); ok {
		_spec.SetField(book.FieldAddedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.PubDate(); ok {
		_spec.SetField(book.FieldPubDate, field.TypeTime, value)
	}
	if buo.mutation.PubDateCleared() {
		_spec.ClearField(book.FieldPubDate, field.TypeTime)
	}
	if value, ok := buo.mutation.Path(); ok {
		_spec.SetField(book.FieldPath, field.TypeString, value)
	}
	if value, ok := buo.mutation.Isbn(); ok {
		_spec.SetField(book.FieldIsbn, field.TypeString, value)
	}
	if buo.mutation.IsbnCleared() {
		_spec.ClearField(book.FieldIsbn, field.TypeString)
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.SetField(book.FieldDescription, field.TypeString, value)
	}
	if buo.mutation.DescriptionCleared() {
		_spec.ClearField(book.FieldDescription, field.TypeString)
	}
	if buo.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedAuthorsIDs(); len(nodes) > 0 && !buo.mutation.AuthorsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.AuthorsTable,
			Columns: book.AuthorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(author.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.SeriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.SeriesTable,
			Columns: book.SeriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(series.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedSeriesIDs(); len(nodes) > 0 && !buo.mutation.SeriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.SeriesTable,
			Columns: book.SeriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(series.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.SeriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.SeriesTable,
			Columns: book.SeriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(series.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.IdentifierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   book.IdentifierTable,
			Columns: []string{book.IdentifierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(identifier.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedIdentifierIDs(); len(nodes) > 0 && !buo.mutation.IdentifierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   book.IdentifierTable,
			Columns: []string{book.IdentifierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(identifier.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.IdentifierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   book.IdentifierTable,
			Columns: []string{book.IdentifierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(identifier.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.LanguageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.LanguageTable,
			Columns: []string{book.LanguageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.LanguageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   book.LanguageTable,
			Columns: []string{book.LanguageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.ShelfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.ShelfTable,
			Columns: book.ShelfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shelf.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedShelfIDs(); len(nodes) > 0 && !buo.mutation.ShelfCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.ShelfTable,
			Columns: book.ShelfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shelf.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ShelfIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.ShelfTable,
			Columns: book.ShelfPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shelf.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Book{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}