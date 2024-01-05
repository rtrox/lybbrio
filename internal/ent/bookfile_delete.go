// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"lybbrio/internal/ent/bookfile"
	"lybbrio/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookFileDelete is the builder for deleting a BookFile entity.
type BookFileDelete struct {
	config
	hooks    []Hook
	mutation *BookFileMutation
}

// Where appends a list predicates to the BookFileDelete builder.
func (bfd *BookFileDelete) Where(ps ...predicate.BookFile) *BookFileDelete {
	bfd.mutation.Where(ps...)
	return bfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bfd *BookFileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bfd.sqlExec, bfd.mutation, bfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bfd *BookFileDelete) ExecX(ctx context.Context) int {
	n, err := bfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bfd *BookFileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bookfile.Table, sqlgraph.NewFieldSpec(bookfile.FieldID, field.TypeString))
	if ps := bfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bfd.mutation.done = true
	return affected, err
}

// BookFileDeleteOne is the builder for deleting a single BookFile entity.
type BookFileDeleteOne struct {
	bfd *BookFileDelete
}

// Where appends a list predicates to the BookFileDelete builder.
func (bfdo *BookFileDeleteOne) Where(ps ...predicate.BookFile) *BookFileDeleteOne {
	bfdo.bfd.mutation.Where(ps...)
	return bfdo
}

// Exec executes the deletion query.
func (bfdo *BookFileDeleteOne) Exec(ctx context.Context) error {
	n, err := bfdo.bfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bookfile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bfdo *BookFileDeleteOne) ExecX(ctx context.Context) {
	if err := bfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
