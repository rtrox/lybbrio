// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"lybbrio/internal/ent/identifier"
	"lybbrio/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IdentifierDelete is the builder for deleting a Identifier entity.
type IdentifierDelete struct {
	config
	hooks    []Hook
	mutation *IdentifierMutation
}

// Where appends a list predicates to the IdentifierDelete builder.
func (id *IdentifierDelete) Where(ps ...predicate.Identifier) *IdentifierDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *IdentifierDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *IdentifierDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *IdentifierDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(identifier.Table, sqlgraph.NewFieldSpec(identifier.FieldID, field.TypeString))
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// IdentifierDeleteOne is the builder for deleting a single Identifier entity.
type IdentifierDeleteOne struct {
	id *IdentifierDelete
}

// Where appends a list predicates to the IdentifierDelete builder.
func (ido *IdentifierDeleteOne) Where(ps ...predicate.Identifier) *IdentifierDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *IdentifierDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{identifier.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *IdentifierDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}
