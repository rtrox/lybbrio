// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"lybbrio/internal/ent/language"
	"lybbrio/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LanguageDelete is the builder for deleting a Language entity.
type LanguageDelete struct {
	config
	hooks    []Hook
	mutation *LanguageMutation
}

// Where appends a list predicates to the LanguageDelete builder.
func (ld *LanguageDelete) Where(ps ...predicate.Language) *LanguageDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LanguageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ld.sqlExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LanguageDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LanguageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(language.Table, sqlgraph.NewFieldSpec(language.FieldID, field.TypeString))
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ld.mutation.done = true
	return affected, err
}

// LanguageDeleteOne is the builder for deleting a single Language entity.
type LanguageDeleteOne struct {
	ld *LanguageDelete
}

// Where appends a list predicates to the LanguageDelete builder.
func (ldo *LanguageDeleteOne) Where(ps ...predicate.Language) *LanguageDeleteOne {
	ldo.ld.mutation.Where(ps...)
	return ldo
}

// Exec executes the deletion query.
func (ldo *LanguageDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{language.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LanguageDeleteOne) ExecX(ctx context.Context) {
	if err := ldo.Exec(ctx); err != nil {
		panic(err)
	}
}
