// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/shelf"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShelfUpdate is the builder for updating Shelf entities.
type ShelfUpdate struct {
	config
	hooks    []Hook
	mutation *ShelfMutation
}

// Where appends a list predicates to the ShelfUpdate builder.
func (su *ShelfUpdate) Where(ps ...predicate.Shelf) *ShelfUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetPublic sets the "public" field.
func (su *ShelfUpdate) SetPublic(b bool) *ShelfUpdate {
	su.mutation.SetPublic(b)
	return su
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (su *ShelfUpdate) SetNillablePublic(b *bool) *ShelfUpdate {
	if b != nil {
		su.SetPublic(*b)
	}
	return su
}

// SetName sets the "name" field.
func (su *ShelfUpdate) SetName(s string) *ShelfUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *ShelfUpdate) SetNillableName(s *string) *ShelfUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetDescription sets the "description" field.
func (su *ShelfUpdate) SetDescription(s string) *ShelfUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (su *ShelfUpdate) SetNillableDescription(s *string) *ShelfUpdate {
	if s != nil {
		su.SetDescription(*s)
	}
	return su
}

// ClearDescription clears the value of the "description" field.
func (su *ShelfUpdate) ClearDescription() *ShelfUpdate {
	su.mutation.ClearDescription()
	return su
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (su *ShelfUpdate) AddBookIDs(ids ...ksuid.ID) *ShelfUpdate {
	su.mutation.AddBookIDs(ids...)
	return su
}

// AddBooks adds the "books" edges to the Book entity.
func (su *ShelfUpdate) AddBooks(b ...*Book) *ShelfUpdate {
	ids := make([]ksuid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return su.AddBookIDs(ids...)
}

// Mutation returns the ShelfMutation object of the builder.
func (su *ShelfUpdate) Mutation() *ShelfMutation {
	return su.mutation
}

// ClearBooks clears all "books" edges to the Book entity.
func (su *ShelfUpdate) ClearBooks() *ShelfUpdate {
	su.mutation.ClearBooks()
	return su
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (su *ShelfUpdate) RemoveBookIDs(ids ...ksuid.ID) *ShelfUpdate {
	su.mutation.RemoveBookIDs(ids...)
	return su
}

// RemoveBooks removes "books" edges to Book entities.
func (su *ShelfUpdate) RemoveBooks(b ...*Book) *ShelfUpdate {
	ids := make([]ksuid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return su.RemoveBookIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *ShelfUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ShelfUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ShelfUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ShelfUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *ShelfUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := shelf.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Shelf.name": %w`, err)}
		}
	}
	if _, ok := su.mutation.UserID(); su.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Shelf.user"`)
	}
	return nil
}

func (su *ShelfUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(shelf.Table, shelf.Columns, sqlgraph.NewFieldSpec(shelf.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Public(); ok {
		_spec.SetField(shelf.FieldPublic, field.TypeBool, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(shelf.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.SetField(shelf.FieldDescription, field.TypeString, value)
	}
	if su.mutation.DescriptionCleared() {
		_spec.ClearField(shelf.FieldDescription, field.TypeString)
	}
	if su.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   shelf.BooksTable,
			Columns: shelf.BooksPrimaryKey,
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
			Table:   shelf.BooksTable,
			Columns: shelf.BooksPrimaryKey,
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
			Table:   shelf.BooksTable,
			Columns: shelf.BooksPrimaryKey,
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
			err = &NotFoundError{shelf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// ShelfUpdateOne is the builder for updating a single Shelf entity.
type ShelfUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShelfMutation
}

// SetPublic sets the "public" field.
func (suo *ShelfUpdateOne) SetPublic(b bool) *ShelfUpdateOne {
	suo.mutation.SetPublic(b)
	return suo
}

// SetNillablePublic sets the "public" field if the given value is not nil.
func (suo *ShelfUpdateOne) SetNillablePublic(b *bool) *ShelfUpdateOne {
	if b != nil {
		suo.SetPublic(*b)
	}
	return suo
}

// SetName sets the "name" field.
func (suo *ShelfUpdateOne) SetName(s string) *ShelfUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *ShelfUpdateOne) SetNillableName(s *string) *ShelfUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetDescription sets the "description" field.
func (suo *ShelfUpdateOne) SetDescription(s string) *ShelfUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (suo *ShelfUpdateOne) SetNillableDescription(s *string) *ShelfUpdateOne {
	if s != nil {
		suo.SetDescription(*s)
	}
	return suo
}

// ClearDescription clears the value of the "description" field.
func (suo *ShelfUpdateOne) ClearDescription() *ShelfUpdateOne {
	suo.mutation.ClearDescription()
	return suo
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (suo *ShelfUpdateOne) AddBookIDs(ids ...ksuid.ID) *ShelfUpdateOne {
	suo.mutation.AddBookIDs(ids...)
	return suo
}

// AddBooks adds the "books" edges to the Book entity.
func (suo *ShelfUpdateOne) AddBooks(b ...*Book) *ShelfUpdateOne {
	ids := make([]ksuid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return suo.AddBookIDs(ids...)
}

// Mutation returns the ShelfMutation object of the builder.
func (suo *ShelfUpdateOne) Mutation() *ShelfMutation {
	return suo.mutation
}

// ClearBooks clears all "books" edges to the Book entity.
func (suo *ShelfUpdateOne) ClearBooks() *ShelfUpdateOne {
	suo.mutation.ClearBooks()
	return suo
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (suo *ShelfUpdateOne) RemoveBookIDs(ids ...ksuid.ID) *ShelfUpdateOne {
	suo.mutation.RemoveBookIDs(ids...)
	return suo
}

// RemoveBooks removes "books" edges to Book entities.
func (suo *ShelfUpdateOne) RemoveBooks(b ...*Book) *ShelfUpdateOne {
	ids := make([]ksuid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return suo.RemoveBookIDs(ids...)
}

// Where appends a list predicates to the ShelfUpdate builder.
func (suo *ShelfUpdateOne) Where(ps ...predicate.Shelf) *ShelfUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *ShelfUpdateOne) Select(field string, fields ...string) *ShelfUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Shelf entity.
func (suo *ShelfUpdateOne) Save(ctx context.Context) (*Shelf, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ShelfUpdateOne) SaveX(ctx context.Context) *Shelf {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *ShelfUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ShelfUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *ShelfUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := shelf.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Shelf.name": %w`, err)}
		}
	}
	if _, ok := suo.mutation.UserID(); suo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Shelf.user"`)
	}
	return nil
}

func (suo *ShelfUpdateOne) sqlSave(ctx context.Context) (_node *Shelf, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(shelf.Table, shelf.Columns, sqlgraph.NewFieldSpec(shelf.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Shelf.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shelf.FieldID)
		for _, f := range fields {
			if !shelf.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != shelf.FieldID {
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
	if value, ok := suo.mutation.Public(); ok {
		_spec.SetField(shelf.FieldPublic, field.TypeBool, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(shelf.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.SetField(shelf.FieldDescription, field.TypeString, value)
	}
	if suo.mutation.DescriptionCleared() {
		_spec.ClearField(shelf.FieldDescription, field.TypeString)
	}
	if suo.mutation.BooksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   shelf.BooksTable,
			Columns: shelf.BooksPrimaryKey,
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
			Table:   shelf.BooksTable,
			Columns: shelf.BooksPrimaryKey,
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
			Table:   shelf.BooksTable,
			Columns: shelf.BooksPrimaryKey,
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
	_node = &Shelf{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shelf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
