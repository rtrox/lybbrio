// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/series"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SeriesUpdate is the builder for updating Series entities.
type SeriesUpdate struct {
	config
	hooks    []Hook
	mutation *SeriesMutation
}

// Where appends a list predicates to the SeriesUpdate builder.
func (su *SeriesUpdate) Where(ps ...predicate.Series) *SeriesUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *SeriesUpdate) SetUpdateTime(t time.Time) *SeriesUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetCalibreID sets the "calibre_id" field.
func (su *SeriesUpdate) SetCalibreID(i int64) *SeriesUpdate {
	su.mutation.ResetCalibreID()
	su.mutation.SetCalibreID(i)
	return su
}

// SetNillableCalibreID sets the "calibre_id" field if the given value is not nil.
func (su *SeriesUpdate) SetNillableCalibreID(i *int64) *SeriesUpdate {
	if i != nil {
		su.SetCalibreID(*i)
	}
	return su
}

// AddCalibreID adds i to the "calibre_id" field.
func (su *SeriesUpdate) AddCalibreID(i int64) *SeriesUpdate {
	su.mutation.AddCalibreID(i)
	return su
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (su *SeriesUpdate) ClearCalibreID() *SeriesUpdate {
	su.mutation.ClearCalibreID()
	return su
}

// SetName sets the "name" field.
func (su *SeriesUpdate) SetName(s string) *SeriesUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SeriesUpdate) SetNillableName(s *string) *SeriesUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetSort sets the "sort" field.
func (su *SeriesUpdate) SetSort(s string) *SeriesUpdate {
	su.mutation.SetSort(s)
	return su
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (su *SeriesUpdate) SetNillableSort(s *string) *SeriesUpdate {
	if s != nil {
		su.SetSort(*s)
	}
	return su
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (su *SeriesUpdate) AddBookIDs(ids ...ksuid.ID) *SeriesUpdate {
	su.mutation.AddBookIDs(ids...)
	return su
}

// AddBooks adds the "books" edges to the Book entity.
func (su *SeriesUpdate) AddBooks(b ...*Book) *SeriesUpdate {
	ids := make([]ksuid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return su.AddBookIDs(ids...)
}

// Mutation returns the SeriesMutation object of the builder.
func (su *SeriesUpdate) Mutation() *SeriesMutation {
	return su.mutation
}

// ClearBooks clears all "books" edges to the Book entity.
func (su *SeriesUpdate) ClearBooks() *SeriesUpdate {
	su.mutation.ClearBooks()
	return su
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (su *SeriesUpdate) RemoveBookIDs(ids ...ksuid.ID) *SeriesUpdate {
	su.mutation.RemoveBookIDs(ids...)
	return su
}

// RemoveBooks removes "books" edges to Book entities.
func (su *SeriesUpdate) RemoveBooks(b ...*Book) *SeriesUpdate {
	ids := make([]ksuid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return su.RemoveBookIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SeriesUpdate) Save(ctx context.Context) (int, error) {
	if err := su.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SeriesUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SeriesUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SeriesUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SeriesUpdate) defaults() error {
	if _, ok := su.mutation.UpdateTime(); !ok {
		if series.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized series.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := series.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (su *SeriesUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := series.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Series.name": %w`, err)}
		}
	}
	if v, ok := su.mutation.Sort(); ok {
		if err := series.SortValidator(v); err != nil {
			return &ValidationError{Name: "sort", err: fmt.Errorf(`ent: validator failed for field "Series.sort": %w`, err)}
		}
	}
	return nil
}

func (su *SeriesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(series.Table, series.Columns, sqlgraph.NewFieldSpec(series.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(series.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.CalibreID(); ok {
		_spec.SetField(series.FieldCalibreID, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedCalibreID(); ok {
		_spec.AddField(series.FieldCalibreID, field.TypeInt64, value)
	}
	if su.mutation.CalibreIDCleared() {
		_spec.ClearField(series.FieldCalibreID, field.TypeInt64)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(series.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Sort(); ok {
		_spec.SetField(series.FieldSort, field.TypeString, value)
	}
	if su.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   series.BooksTable,
			Columns: series.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedBooksIDs(); len(nodes) > 0 && !su.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   series.BooksTable,
			Columns: series.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   series.BooksTable,
			Columns: series.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{series.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SeriesUpdateOne is the builder for updating a single Series entity.
type SeriesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SeriesMutation
}

// SetUpdateTime sets the "update_time" field.
func (suo *SeriesUpdateOne) SetUpdateTime(t time.Time) *SeriesUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetCalibreID sets the "calibre_id" field.
func (suo *SeriesUpdateOne) SetCalibreID(i int64) *SeriesUpdateOne {
	suo.mutation.ResetCalibreID()
	suo.mutation.SetCalibreID(i)
	return suo
}

// SetNillableCalibreID sets the "calibre_id" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillableCalibreID(i *int64) *SeriesUpdateOne {
	if i != nil {
		suo.SetCalibreID(*i)
	}
	return suo
}

// AddCalibreID adds i to the "calibre_id" field.
func (suo *SeriesUpdateOne) AddCalibreID(i int64) *SeriesUpdateOne {
	suo.mutation.AddCalibreID(i)
	return suo
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (suo *SeriesUpdateOne) ClearCalibreID() *SeriesUpdateOne {
	suo.mutation.ClearCalibreID()
	return suo
}

// SetName sets the "name" field.
func (suo *SeriesUpdateOne) SetName(s string) *SeriesUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillableName(s *string) *SeriesUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetSort sets the "sort" field.
func (suo *SeriesUpdateOne) SetSort(s string) *SeriesUpdateOne {
	suo.mutation.SetSort(s)
	return suo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillableSort(s *string) *SeriesUpdateOne {
	if s != nil {
		suo.SetSort(*s)
	}
	return suo
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (suo *SeriesUpdateOne) AddBookIDs(ids ...ksuid.ID) *SeriesUpdateOne {
	suo.mutation.AddBookIDs(ids...)
	return suo
}

// AddBooks adds the "books" edges to the Book entity.
func (suo *SeriesUpdateOne) AddBooks(b ...*Book) *SeriesUpdateOne {
	ids := make([]ksuid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return suo.AddBookIDs(ids...)
}

// Mutation returns the SeriesMutation object of the builder.
func (suo *SeriesUpdateOne) Mutation() *SeriesMutation {
	return suo.mutation
}

// ClearBooks clears all "books" edges to the Book entity.
func (suo *SeriesUpdateOne) ClearBooks() *SeriesUpdateOne {
	suo.mutation.ClearBooks()
	return suo
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (suo *SeriesUpdateOne) RemoveBookIDs(ids ...ksuid.ID) *SeriesUpdateOne {
	suo.mutation.RemoveBookIDs(ids...)
	return suo
}

// RemoveBooks removes "books" edges to Book entities.
func (suo *SeriesUpdateOne) RemoveBooks(b ...*Book) *SeriesUpdateOne {
	ids := make([]ksuid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return suo.RemoveBookIDs(ids...)
}

// Where appends a list predicates to the SeriesUpdate builder.
func (suo *SeriesUpdateOne) Where(ps ...predicate.Series) *SeriesUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SeriesUpdateOne) Select(field string, fields ...string) *SeriesUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Series entity.
func (suo *SeriesUpdateOne) Save(ctx context.Context) (*Series, error) {
	if err := suo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SeriesUpdateOne) SaveX(ctx context.Context) *Series {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SeriesUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SeriesUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SeriesUpdateOne) defaults() error {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		if series.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized series.UpdateDefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := series.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (suo *SeriesUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := series.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Series.name": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Sort(); ok {
		if err := series.SortValidator(v); err != nil {
			return &ValidationError{Name: "sort", err: fmt.Errorf(`ent: validator failed for field "Series.sort": %w`, err)}
		}
	}
	return nil
}

func (suo *SeriesUpdateOne) sqlSave(ctx context.Context) (_node *Series, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(series.Table, series.Columns, sqlgraph.NewFieldSpec(series.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Series.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, series.FieldID)
		for _, f := range fields {
			if !series.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != series.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(series.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.CalibreID(); ok {
		_spec.SetField(series.FieldCalibreID, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedCalibreID(); ok {
		_spec.AddField(series.FieldCalibreID, field.TypeInt64, value)
	}
	if suo.mutation.CalibreIDCleared() {
		_spec.ClearField(series.FieldCalibreID, field.TypeInt64)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(series.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Sort(); ok {
		_spec.SetField(series.FieldSort, field.TypeString, value)
	}
	if suo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   series.BooksTable,
			Columns: series.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedBooksIDs(); len(nodes) > 0 && !suo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   series.BooksTable,
			Columns: series.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   series.BooksTable,
			Columns: series.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Series{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{series.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
