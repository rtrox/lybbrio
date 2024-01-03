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
	"lybbrio/internal/ent/publisher"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/series"
	"lybbrio/internal/ent/shelf"
	"lybbrio/internal/ent/tag"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookCreate is the builder for creating a Book entity.
type BookCreate struct {
	config
	mutation *BookMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCalibreID sets the "calibre_id" field.
func (bc *BookCreate) SetCalibreID(i int64) *BookCreate {
	bc.mutation.SetCalibreID(i)
	return bc
}

// SetNillableCalibreID sets the "calibre_id" field if the given value is not nil.
func (bc *BookCreate) SetNillableCalibreID(i *int64) *BookCreate {
	if i != nil {
		bc.SetCalibreID(*i)
	}
	return bc
}

// SetTitle sets the "title" field.
func (bc *BookCreate) SetTitle(s string) *BookCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetSort sets the "sort" field.
func (bc *BookCreate) SetSort(s string) *BookCreate {
	bc.mutation.SetSort(s)
	return bc
}

// SetPublishedDate sets the "published_date" field.
func (bc *BookCreate) SetPublishedDate(t time.Time) *BookCreate {
	bc.mutation.SetPublishedDate(t)
	return bc
}

// SetNillablePublishedDate sets the "published_date" field if the given value is not nil.
func (bc *BookCreate) SetNillablePublishedDate(t *time.Time) *BookCreate {
	if t != nil {
		bc.SetPublishedDate(*t)
	}
	return bc
}

// SetPath sets the "path" field.
func (bc *BookCreate) SetPath(s string) *BookCreate {
	bc.mutation.SetPath(s)
	return bc
}

// SetIsbn sets the "isbn" field.
func (bc *BookCreate) SetIsbn(s string) *BookCreate {
	bc.mutation.SetIsbn(s)
	return bc
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (bc *BookCreate) SetNillableIsbn(s *string) *BookCreate {
	if s != nil {
		bc.SetIsbn(*s)
	}
	return bc
}

// SetDescription sets the "description" field.
func (bc *BookCreate) SetDescription(s string) *BookCreate {
	bc.mutation.SetDescription(s)
	return bc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bc *BookCreate) SetNillableDescription(s *string) *BookCreate {
	if s != nil {
		bc.SetDescription(*s)
	}
	return bc
}

// SetSeriesIndex sets the "series_index" field.
func (bc *BookCreate) SetSeriesIndex(f float64) *BookCreate {
	bc.mutation.SetSeriesIndex(f)
	return bc
}

// SetNillableSeriesIndex sets the "series_index" field if the given value is not nil.
func (bc *BookCreate) SetNillableSeriesIndex(f *float64) *BookCreate {
	if f != nil {
		bc.SetSeriesIndex(*f)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BookCreate) SetID(k ksuid.ID) *BookCreate {
	bc.mutation.SetID(k)
	return bc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bc *BookCreate) SetNillableID(k *ksuid.ID) *BookCreate {
	if k != nil {
		bc.SetID(*k)
	}
	return bc
}

// AddAuthorIDs adds the "authors" edge to the Author entity by IDs.
func (bc *BookCreate) AddAuthorIDs(ids ...ksuid.ID) *BookCreate {
	bc.mutation.AddAuthorIDs(ids...)
	return bc
}

// AddAuthors adds the "authors" edges to the Author entity.
func (bc *BookCreate) AddAuthors(a ...*Author) *BookCreate {
	ids := make([]ksuid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bc.AddAuthorIDs(ids...)
}

// AddPublisherIDs adds the "publisher" edge to the Publisher entity by IDs.
func (bc *BookCreate) AddPublisherIDs(ids ...ksuid.ID) *BookCreate {
	bc.mutation.AddPublisherIDs(ids...)
	return bc
}

// AddPublisher adds the "publisher" edges to the Publisher entity.
func (bc *BookCreate) AddPublisher(p ...*Publisher) *BookCreate {
	ids := make([]ksuid.ID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bc.AddPublisherIDs(ids...)
}

// AddSeriesIDs adds the "series" edge to the Series entity by IDs.
func (bc *BookCreate) AddSeriesIDs(ids ...ksuid.ID) *BookCreate {
	bc.mutation.AddSeriesIDs(ids...)
	return bc
}

// AddSeries adds the "series" edges to the Series entity.
func (bc *BookCreate) AddSeries(s ...*Series) *BookCreate {
	ids := make([]ksuid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bc.AddSeriesIDs(ids...)
}

// AddIdentifierIDs adds the "identifiers" edge to the Identifier entity by IDs.
func (bc *BookCreate) AddIdentifierIDs(ids ...ksuid.ID) *BookCreate {
	bc.mutation.AddIdentifierIDs(ids...)
	return bc
}

// AddIdentifiers adds the "identifiers" edges to the Identifier entity.
func (bc *BookCreate) AddIdentifiers(i ...*Identifier) *BookCreate {
	ids := make([]ksuid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return bc.AddIdentifierIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (bc *BookCreate) AddTagIDs(ids ...ksuid.ID) *BookCreate {
	bc.mutation.AddTagIDs(ids...)
	return bc
}

// AddTags adds the "tags" edges to the Tag entity.
func (bc *BookCreate) AddTags(t ...*Tag) *BookCreate {
	ids := make([]ksuid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddTagIDs(ids...)
}

// AddLanguageIDs adds the "language" edge to the Language entity by IDs.
func (bc *BookCreate) AddLanguageIDs(ids ...ksuid.ID) *BookCreate {
	bc.mutation.AddLanguageIDs(ids...)
	return bc
}

// AddLanguage adds the "language" edges to the Language entity.
func (bc *BookCreate) AddLanguage(l ...*Language) *BookCreate {
	ids := make([]ksuid.ID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return bc.AddLanguageIDs(ids...)
}

// AddShelfIDs adds the "shelf" edge to the Shelf entity by IDs.
func (bc *BookCreate) AddShelfIDs(ids ...ksuid.ID) *BookCreate {
	bc.mutation.AddShelfIDs(ids...)
	return bc
}

// AddShelf adds the "shelf" edges to the Shelf entity.
func (bc *BookCreate) AddShelf(s ...*Shelf) *BookCreate {
	ids := make([]ksuid.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bc.AddShelfIDs(ids...)
}

// Mutation returns the BookMutation object of the builder.
func (bc *BookCreate) Mutation() *BookMutation {
	return bc.mutation
}

// Save creates the Book in the database.
func (bc *BookCreate) Save(ctx context.Context) (*Book, error) {
	if err := bc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookCreate) SaveX(ctx context.Context) *Book {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookCreate) defaults() error {
	if _, ok := bc.mutation.ID(); !ok {
		if book.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized book.DefaultID (forgotten import ent/runtime?)")
		}
		v := book.DefaultID()
		bc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookCreate) check() error {
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Book.title"`)}
	}
	if v, ok := bc.mutation.Title(); ok {
		if err := book.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Book.title": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Book.sort"`)}
	}
	if _, ok := bc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Book.path"`)}
	}
	if v, ok := bc.mutation.Path(); ok {
		if err := book.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Book.path": %w`, err)}
		}
	}
	return nil
}

func (bc *BookCreate) sqlSave(ctx context.Context) (*Book, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(ksuid.ID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Book.ID type: %T", _spec.ID.Value)
		}
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BookCreate) createSpec() (*Book, *sqlgraph.CreateSpec) {
	var (
		_node = &Book{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(book.Table, sqlgraph.NewFieldSpec(book.FieldID, field.TypeString))
	)
	_spec.OnConflict = bc.conflict
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.CalibreID(); ok {
		_spec.SetField(book.FieldCalibreID, field.TypeInt64, value)
		_node.CalibreID = value
	}
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.Sort(); ok {
		_spec.SetField(book.FieldSort, field.TypeString, value)
		_node.Sort = value
	}
	if value, ok := bc.mutation.PublishedDate(); ok {
		_spec.SetField(book.FieldPublishedDate, field.TypeTime, value)
		_node.PublishedDate = value
	}
	if value, ok := bc.mutation.Path(); ok {
		_spec.SetField(book.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := bc.mutation.Isbn(); ok {
		_spec.SetField(book.FieldIsbn, field.TypeString, value)
		_node.Isbn = value
	}
	if value, ok := bc.mutation.Description(); ok {
		_spec.SetField(book.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := bc.mutation.SeriesIndex(); ok {
		_spec.SetField(book.FieldSeriesIndex, field.TypeFloat64, value)
		_node.SeriesIndex = value
	}
	if nodes := bc.mutation.AuthorsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.PublisherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.PublisherTable,
			Columns: book.PublisherPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(publisher.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.SeriesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.IdentifiersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   book.IdentifiersTable,
			Columns: []string{book.IdentifiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(identifier.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.TagsTable,
			Columns: book.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.LanguageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   book.LanguageTable,
			Columns: book.LanguagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.ShelfIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Book.Create().
//		SetCalibreID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BookUpsert) {
//			SetCalibreID(v+v).
//		}).
//		Exec(ctx)
func (bc *BookCreate) OnConflict(opts ...sql.ConflictOption) *BookUpsertOne {
	bc.conflict = opts
	return &BookUpsertOne{
		create: bc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Book.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bc *BookCreate) OnConflictColumns(columns ...string) *BookUpsertOne {
	bc.conflict = append(bc.conflict, sql.ConflictColumns(columns...))
	return &BookUpsertOne{
		create: bc,
	}
}

type (
	// BookUpsertOne is the builder for "upsert"-ing
	//  one Book node.
	BookUpsertOne struct {
		create *BookCreate
	}

	// BookUpsert is the "OnConflict" setter.
	BookUpsert struct {
		*sql.UpdateSet
	}
)

// SetCalibreID sets the "calibre_id" field.
func (u *BookUpsert) SetCalibreID(v int64) *BookUpsert {
	u.Set(book.FieldCalibreID, v)
	return u
}

// UpdateCalibreID sets the "calibre_id" field to the value that was provided on create.
func (u *BookUpsert) UpdateCalibreID() *BookUpsert {
	u.SetExcluded(book.FieldCalibreID)
	return u
}

// AddCalibreID adds v to the "calibre_id" field.
func (u *BookUpsert) AddCalibreID(v int64) *BookUpsert {
	u.Add(book.FieldCalibreID, v)
	return u
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (u *BookUpsert) ClearCalibreID() *BookUpsert {
	u.SetNull(book.FieldCalibreID)
	return u
}

// SetTitle sets the "title" field.
func (u *BookUpsert) SetTitle(v string) *BookUpsert {
	u.Set(book.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *BookUpsert) UpdateTitle() *BookUpsert {
	u.SetExcluded(book.FieldTitle)
	return u
}

// SetSort sets the "sort" field.
func (u *BookUpsert) SetSort(v string) *BookUpsert {
	u.Set(book.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *BookUpsert) UpdateSort() *BookUpsert {
	u.SetExcluded(book.FieldSort)
	return u
}

// SetPublishedDate sets the "published_date" field.
func (u *BookUpsert) SetPublishedDate(v time.Time) *BookUpsert {
	u.Set(book.FieldPublishedDate, v)
	return u
}

// UpdatePublishedDate sets the "published_date" field to the value that was provided on create.
func (u *BookUpsert) UpdatePublishedDate() *BookUpsert {
	u.SetExcluded(book.FieldPublishedDate)
	return u
}

// ClearPublishedDate clears the value of the "published_date" field.
func (u *BookUpsert) ClearPublishedDate() *BookUpsert {
	u.SetNull(book.FieldPublishedDate)
	return u
}

// SetPath sets the "path" field.
func (u *BookUpsert) SetPath(v string) *BookUpsert {
	u.Set(book.FieldPath, v)
	return u
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *BookUpsert) UpdatePath() *BookUpsert {
	u.SetExcluded(book.FieldPath)
	return u
}

// SetIsbn sets the "isbn" field.
func (u *BookUpsert) SetIsbn(v string) *BookUpsert {
	u.Set(book.FieldIsbn, v)
	return u
}

// UpdateIsbn sets the "isbn" field to the value that was provided on create.
func (u *BookUpsert) UpdateIsbn() *BookUpsert {
	u.SetExcluded(book.FieldIsbn)
	return u
}

// ClearIsbn clears the value of the "isbn" field.
func (u *BookUpsert) ClearIsbn() *BookUpsert {
	u.SetNull(book.FieldIsbn)
	return u
}

// SetDescription sets the "description" field.
func (u *BookUpsert) SetDescription(v string) *BookUpsert {
	u.Set(book.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *BookUpsert) UpdateDescription() *BookUpsert {
	u.SetExcluded(book.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *BookUpsert) ClearDescription() *BookUpsert {
	u.SetNull(book.FieldDescription)
	return u
}

// SetSeriesIndex sets the "series_index" field.
func (u *BookUpsert) SetSeriesIndex(v float64) *BookUpsert {
	u.Set(book.FieldSeriesIndex, v)
	return u
}

// UpdateSeriesIndex sets the "series_index" field to the value that was provided on create.
func (u *BookUpsert) UpdateSeriesIndex() *BookUpsert {
	u.SetExcluded(book.FieldSeriesIndex)
	return u
}

// AddSeriesIndex adds v to the "series_index" field.
func (u *BookUpsert) AddSeriesIndex(v float64) *BookUpsert {
	u.Add(book.FieldSeriesIndex, v)
	return u
}

// ClearSeriesIndex clears the value of the "series_index" field.
func (u *BookUpsert) ClearSeriesIndex() *BookUpsert {
	u.SetNull(book.FieldSeriesIndex)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Book.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(book.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BookUpsertOne) UpdateNewValues() *BookUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(book.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Book.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BookUpsertOne) Ignore() *BookUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BookUpsertOne) DoNothing() *BookUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BookCreate.OnConflict
// documentation for more info.
func (u *BookUpsertOne) Update(set func(*BookUpsert)) *BookUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BookUpsert{UpdateSet: update})
	}))
	return u
}

// SetCalibreID sets the "calibre_id" field.
func (u *BookUpsertOne) SetCalibreID(v int64) *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.SetCalibreID(v)
	})
}

// AddCalibreID adds v to the "calibre_id" field.
func (u *BookUpsertOne) AddCalibreID(v int64) *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.AddCalibreID(v)
	})
}

// UpdateCalibreID sets the "calibre_id" field to the value that was provided on create.
func (u *BookUpsertOne) UpdateCalibreID() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.UpdateCalibreID()
	})
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (u *BookUpsertOne) ClearCalibreID() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.ClearCalibreID()
	})
}

// SetTitle sets the "title" field.
func (u *BookUpsertOne) SetTitle(v string) *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *BookUpsertOne) UpdateTitle() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.UpdateTitle()
	})
}

// SetSort sets the "sort" field.
func (u *BookUpsertOne) SetSort(v string) *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.SetSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *BookUpsertOne) UpdateSort() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.UpdateSort()
	})
}

// SetPublishedDate sets the "published_date" field.
func (u *BookUpsertOne) SetPublishedDate(v time.Time) *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.SetPublishedDate(v)
	})
}

// UpdatePublishedDate sets the "published_date" field to the value that was provided on create.
func (u *BookUpsertOne) UpdatePublishedDate() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.UpdatePublishedDate()
	})
}

// ClearPublishedDate clears the value of the "published_date" field.
func (u *BookUpsertOne) ClearPublishedDate() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.ClearPublishedDate()
	})
}

// SetPath sets the "path" field.
func (u *BookUpsertOne) SetPath(v string) *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *BookUpsertOne) UpdatePath() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.UpdatePath()
	})
}

// SetIsbn sets the "isbn" field.
func (u *BookUpsertOne) SetIsbn(v string) *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.SetIsbn(v)
	})
}

// UpdateIsbn sets the "isbn" field to the value that was provided on create.
func (u *BookUpsertOne) UpdateIsbn() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.UpdateIsbn()
	})
}

// ClearIsbn clears the value of the "isbn" field.
func (u *BookUpsertOne) ClearIsbn() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.ClearIsbn()
	})
}

// SetDescription sets the "description" field.
func (u *BookUpsertOne) SetDescription(v string) *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *BookUpsertOne) UpdateDescription() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *BookUpsertOne) ClearDescription() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.ClearDescription()
	})
}

// SetSeriesIndex sets the "series_index" field.
func (u *BookUpsertOne) SetSeriesIndex(v float64) *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.SetSeriesIndex(v)
	})
}

// AddSeriesIndex adds v to the "series_index" field.
func (u *BookUpsertOne) AddSeriesIndex(v float64) *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.AddSeriesIndex(v)
	})
}

// UpdateSeriesIndex sets the "series_index" field to the value that was provided on create.
func (u *BookUpsertOne) UpdateSeriesIndex() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.UpdateSeriesIndex()
	})
}

// ClearSeriesIndex clears the value of the "series_index" field.
func (u *BookUpsertOne) ClearSeriesIndex() *BookUpsertOne {
	return u.Update(func(s *BookUpsert) {
		s.ClearSeriesIndex()
	})
}

// Exec executes the query.
func (u *BookUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BookCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BookUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BookUpsertOne) ID(ctx context.Context) (id ksuid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: BookUpsertOne.ID is not supported by MySQL driver. Use BookUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BookUpsertOne) IDX(ctx context.Context) ksuid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BookCreateBulk is the builder for creating many Book entities in bulk.
type BookCreateBulk struct {
	config
	err      error
	builders []*BookCreate
	conflict []sql.ConflictOption
}

// Save creates the Book entities in the database.
func (bcb *BookCreateBulk) Save(ctx context.Context) ([]*Book, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Book, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookCreateBulk) SaveX(ctx context.Context) []*Book {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Book.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BookUpsert) {
//			SetCalibreID(v+v).
//		}).
//		Exec(ctx)
func (bcb *BookCreateBulk) OnConflict(opts ...sql.ConflictOption) *BookUpsertBulk {
	bcb.conflict = opts
	return &BookUpsertBulk{
		create: bcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Book.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bcb *BookCreateBulk) OnConflictColumns(columns ...string) *BookUpsertBulk {
	bcb.conflict = append(bcb.conflict, sql.ConflictColumns(columns...))
	return &BookUpsertBulk{
		create: bcb,
	}
}

// BookUpsertBulk is the builder for "upsert"-ing
// a bulk of Book nodes.
type BookUpsertBulk struct {
	create *BookCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Book.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(book.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BookUpsertBulk) UpdateNewValues() *BookUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(book.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Book.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BookUpsertBulk) Ignore() *BookUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BookUpsertBulk) DoNothing() *BookUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BookCreateBulk.OnConflict
// documentation for more info.
func (u *BookUpsertBulk) Update(set func(*BookUpsert)) *BookUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BookUpsert{UpdateSet: update})
	}))
	return u
}

// SetCalibreID sets the "calibre_id" field.
func (u *BookUpsertBulk) SetCalibreID(v int64) *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.SetCalibreID(v)
	})
}

// AddCalibreID adds v to the "calibre_id" field.
func (u *BookUpsertBulk) AddCalibreID(v int64) *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.AddCalibreID(v)
	})
}

// UpdateCalibreID sets the "calibre_id" field to the value that was provided on create.
func (u *BookUpsertBulk) UpdateCalibreID() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.UpdateCalibreID()
	})
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (u *BookUpsertBulk) ClearCalibreID() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.ClearCalibreID()
	})
}

// SetTitle sets the "title" field.
func (u *BookUpsertBulk) SetTitle(v string) *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *BookUpsertBulk) UpdateTitle() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.UpdateTitle()
	})
}

// SetSort sets the "sort" field.
func (u *BookUpsertBulk) SetSort(v string) *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.SetSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *BookUpsertBulk) UpdateSort() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.UpdateSort()
	})
}

// SetPublishedDate sets the "published_date" field.
func (u *BookUpsertBulk) SetPublishedDate(v time.Time) *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.SetPublishedDate(v)
	})
}

// UpdatePublishedDate sets the "published_date" field to the value that was provided on create.
func (u *BookUpsertBulk) UpdatePublishedDate() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.UpdatePublishedDate()
	})
}

// ClearPublishedDate clears the value of the "published_date" field.
func (u *BookUpsertBulk) ClearPublishedDate() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.ClearPublishedDate()
	})
}

// SetPath sets the "path" field.
func (u *BookUpsertBulk) SetPath(v string) *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *BookUpsertBulk) UpdatePath() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.UpdatePath()
	})
}

// SetIsbn sets the "isbn" field.
func (u *BookUpsertBulk) SetIsbn(v string) *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.SetIsbn(v)
	})
}

// UpdateIsbn sets the "isbn" field to the value that was provided on create.
func (u *BookUpsertBulk) UpdateIsbn() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.UpdateIsbn()
	})
}

// ClearIsbn clears the value of the "isbn" field.
func (u *BookUpsertBulk) ClearIsbn() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.ClearIsbn()
	})
}

// SetDescription sets the "description" field.
func (u *BookUpsertBulk) SetDescription(v string) *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *BookUpsertBulk) UpdateDescription() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *BookUpsertBulk) ClearDescription() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.ClearDescription()
	})
}

// SetSeriesIndex sets the "series_index" field.
func (u *BookUpsertBulk) SetSeriesIndex(v float64) *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.SetSeriesIndex(v)
	})
}

// AddSeriesIndex adds v to the "series_index" field.
func (u *BookUpsertBulk) AddSeriesIndex(v float64) *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.AddSeriesIndex(v)
	})
}

// UpdateSeriesIndex sets the "series_index" field to the value that was provided on create.
func (u *BookUpsertBulk) UpdateSeriesIndex() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.UpdateSeriesIndex()
	})
}

// ClearSeriesIndex clears the value of the "series_index" field.
func (u *BookUpsertBulk) ClearSeriesIndex() *BookUpsertBulk {
	return u.Update(func(s *BookUpsert) {
		s.ClearSeriesIndex()
	})
}

// Exec executes the query.
func (u *BookUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BookCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BookCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BookUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
