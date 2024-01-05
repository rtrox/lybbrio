// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/ent/userpermissions"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPermissionsUpdate is the builder for updating UserPermissions entities.
type UserPermissionsUpdate struct {
	config
	hooks    []Hook
	mutation *UserPermissionsMutation
}

// Where appends a list predicates to the UserPermissionsUpdate builder.
func (upu *UserPermissionsUpdate) Where(ps ...predicate.UserPermissions) *UserPermissionsUpdate {
	upu.mutation.Where(ps...)
	return upu
}

// SetUserID sets the "user_id" field.
func (upu *UserPermissionsUpdate) SetUserID(k ksuid.ID) *UserPermissionsUpdate {
	upu.mutation.SetUserID(k)
	return upu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (upu *UserPermissionsUpdate) SetNillableUserID(k *ksuid.ID) *UserPermissionsUpdate {
	if k != nil {
		upu.SetUserID(*k)
	}
	return upu
}

// ClearUserID clears the value of the "user_id" field.
func (upu *UserPermissionsUpdate) ClearUserID() *UserPermissionsUpdate {
	upu.mutation.ClearUserID()
	return upu
}

// SetAdmin sets the "Admin" field.
func (upu *UserPermissionsUpdate) SetAdmin(b bool) *UserPermissionsUpdate {
	upu.mutation.SetAdmin(b)
	return upu
}

// SetNillableAdmin sets the "Admin" field if the given value is not nil.
func (upu *UserPermissionsUpdate) SetNillableAdmin(b *bool) *UserPermissionsUpdate {
	if b != nil {
		upu.SetAdmin(*b)
	}
	return upu
}

// SetCanCreatePublic sets the "CanCreatePublic" field.
func (upu *UserPermissionsUpdate) SetCanCreatePublic(b bool) *UserPermissionsUpdate {
	upu.mutation.SetCanCreatePublic(b)
	return upu
}

// SetNillableCanCreatePublic sets the "CanCreatePublic" field if the given value is not nil.
func (upu *UserPermissionsUpdate) SetNillableCanCreatePublic(b *bool) *UserPermissionsUpdate {
	if b != nil {
		upu.SetCanCreatePublic(*b)
	}
	return upu
}

// SetCanEdit sets the "CanEdit" field.
func (upu *UserPermissionsUpdate) SetCanEdit(b bool) *UserPermissionsUpdate {
	upu.mutation.SetCanEdit(b)
	return upu
}

// SetNillableCanEdit sets the "CanEdit" field if the given value is not nil.
func (upu *UserPermissionsUpdate) SetNillableCanEdit(b *bool) *UserPermissionsUpdate {
	if b != nil {
		upu.SetCanEdit(*b)
	}
	return upu
}

// SetUser sets the "user" edge to the User entity.
func (upu *UserPermissionsUpdate) SetUser(u *User) *UserPermissionsUpdate {
	return upu.SetUserID(u.ID)
}

// Mutation returns the UserPermissionsMutation object of the builder.
func (upu *UserPermissionsUpdate) Mutation() *UserPermissionsMutation {
	return upu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (upu *UserPermissionsUpdate) ClearUser() *UserPermissionsUpdate {
	upu.mutation.ClearUser()
	return upu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upu *UserPermissionsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, upu.sqlSave, upu.mutation, upu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upu *UserPermissionsUpdate) SaveX(ctx context.Context) int {
	affected, err := upu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upu *UserPermissionsUpdate) Exec(ctx context.Context) error {
	_, err := upu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upu *UserPermissionsUpdate) ExecX(ctx context.Context) {
	if err := upu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (upu *UserPermissionsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userpermissions.Table, userpermissions.Columns, sqlgraph.NewFieldSpec(userpermissions.FieldID, field.TypeString))
	if ps := upu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upu.mutation.Admin(); ok {
		_spec.SetField(userpermissions.FieldAdmin, field.TypeBool, value)
	}
	if value, ok := upu.mutation.CanCreatePublic(); ok {
		_spec.SetField(userpermissions.FieldCanCreatePublic, field.TypeBool, value)
	}
	if value, ok := upu.mutation.CanEdit(); ok {
		_spec.SetField(userpermissions.FieldCanEdit, field.TypeBool, value)
	}
	if upu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userpermissions.UserTable,
			Columns: []string{userpermissions.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userpermissions.UserTable,
			Columns: []string{userpermissions.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, upu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userpermissions.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	upu.mutation.done = true
	return n, nil
}

// UserPermissionsUpdateOne is the builder for updating a single UserPermissions entity.
type UserPermissionsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserPermissionsMutation
}

// SetUserID sets the "user_id" field.
func (upuo *UserPermissionsUpdateOne) SetUserID(k ksuid.ID) *UserPermissionsUpdateOne {
	upuo.mutation.SetUserID(k)
	return upuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (upuo *UserPermissionsUpdateOne) SetNillableUserID(k *ksuid.ID) *UserPermissionsUpdateOne {
	if k != nil {
		upuo.SetUserID(*k)
	}
	return upuo
}

// ClearUserID clears the value of the "user_id" field.
func (upuo *UserPermissionsUpdateOne) ClearUserID() *UserPermissionsUpdateOne {
	upuo.mutation.ClearUserID()
	return upuo
}

// SetAdmin sets the "Admin" field.
func (upuo *UserPermissionsUpdateOne) SetAdmin(b bool) *UserPermissionsUpdateOne {
	upuo.mutation.SetAdmin(b)
	return upuo
}

// SetNillableAdmin sets the "Admin" field if the given value is not nil.
func (upuo *UserPermissionsUpdateOne) SetNillableAdmin(b *bool) *UserPermissionsUpdateOne {
	if b != nil {
		upuo.SetAdmin(*b)
	}
	return upuo
}

// SetCanCreatePublic sets the "CanCreatePublic" field.
func (upuo *UserPermissionsUpdateOne) SetCanCreatePublic(b bool) *UserPermissionsUpdateOne {
	upuo.mutation.SetCanCreatePublic(b)
	return upuo
}

// SetNillableCanCreatePublic sets the "CanCreatePublic" field if the given value is not nil.
func (upuo *UserPermissionsUpdateOne) SetNillableCanCreatePublic(b *bool) *UserPermissionsUpdateOne {
	if b != nil {
		upuo.SetCanCreatePublic(*b)
	}
	return upuo
}

// SetCanEdit sets the "CanEdit" field.
func (upuo *UserPermissionsUpdateOne) SetCanEdit(b bool) *UserPermissionsUpdateOne {
	upuo.mutation.SetCanEdit(b)
	return upuo
}

// SetNillableCanEdit sets the "CanEdit" field if the given value is not nil.
func (upuo *UserPermissionsUpdateOne) SetNillableCanEdit(b *bool) *UserPermissionsUpdateOne {
	if b != nil {
		upuo.SetCanEdit(*b)
	}
	return upuo
}

// SetUser sets the "user" edge to the User entity.
func (upuo *UserPermissionsUpdateOne) SetUser(u *User) *UserPermissionsUpdateOne {
	return upuo.SetUserID(u.ID)
}

// Mutation returns the UserPermissionsMutation object of the builder.
func (upuo *UserPermissionsUpdateOne) Mutation() *UserPermissionsMutation {
	return upuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (upuo *UserPermissionsUpdateOne) ClearUser() *UserPermissionsUpdateOne {
	upuo.mutation.ClearUser()
	return upuo
}

// Where appends a list predicates to the UserPermissionsUpdate builder.
func (upuo *UserPermissionsUpdateOne) Where(ps ...predicate.UserPermissions) *UserPermissionsUpdateOne {
	upuo.mutation.Where(ps...)
	return upuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (upuo *UserPermissionsUpdateOne) Select(field string, fields ...string) *UserPermissionsUpdateOne {
	upuo.fields = append([]string{field}, fields...)
	return upuo
}

// Save executes the query and returns the updated UserPermissions entity.
func (upuo *UserPermissionsUpdateOne) Save(ctx context.Context) (*UserPermissions, error) {
	return withHooks(ctx, upuo.sqlSave, upuo.mutation, upuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upuo *UserPermissionsUpdateOne) SaveX(ctx context.Context) *UserPermissions {
	node, err := upuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upuo *UserPermissionsUpdateOne) Exec(ctx context.Context) error {
	_, err := upuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upuo *UserPermissionsUpdateOne) ExecX(ctx context.Context) {
	if err := upuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (upuo *UserPermissionsUpdateOne) sqlSave(ctx context.Context) (_node *UserPermissions, err error) {
	_spec := sqlgraph.NewUpdateSpec(userpermissions.Table, userpermissions.Columns, sqlgraph.NewFieldSpec(userpermissions.FieldID, field.TypeString))
	id, ok := upuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserPermissions.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := upuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userpermissions.FieldID)
		for _, f := range fields {
			if !userpermissions.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userpermissions.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := upuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upuo.mutation.Admin(); ok {
		_spec.SetField(userpermissions.FieldAdmin, field.TypeBool, value)
	}
	if value, ok := upuo.mutation.CanCreatePublic(); ok {
		_spec.SetField(userpermissions.FieldCanCreatePublic, field.TypeBool, value)
	}
	if value, ok := upuo.mutation.CanEdit(); ok {
		_spec.SetField(userpermissions.FieldCanEdit, field.TypeBool, value)
	}
	if upuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userpermissions.UserTable,
			Columns: []string{userpermissions.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := upuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userpermissions.UserTable,
			Columns: []string{userpermissions.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserPermissions{config: upuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userpermissions.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	upuo.mutation.done = true
	return _node, nil
}
