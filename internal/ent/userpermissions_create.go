// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/ent/userpermissions"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPermissionsCreate is the builder for creating a UserPermissions entity.
type UserPermissionsCreate struct {
	config
	mutation *UserPermissionsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserID sets the "user_id" field.
func (upc *UserPermissionsCreate) SetUserID(k ksuid.ID) *UserPermissionsCreate {
	upc.mutation.SetUserID(k)
	return upc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (upc *UserPermissionsCreate) SetNillableUserID(k *ksuid.ID) *UserPermissionsCreate {
	if k != nil {
		upc.SetUserID(*k)
	}
	return upc
}

// SetCanEdit sets the "CanEdit" field.
func (upc *UserPermissionsCreate) SetCanEdit(b bool) *UserPermissionsCreate {
	upc.mutation.SetCanEdit(b)
	return upc
}

// SetNillableCanEdit sets the "CanEdit" field if the given value is not nil.
func (upc *UserPermissionsCreate) SetNillableCanEdit(b *bool) *UserPermissionsCreate {
	if b != nil {
		upc.SetCanEdit(*b)
	}
	return upc
}

// SetAdmin sets the "Admin" field.
func (upc *UserPermissionsCreate) SetAdmin(b bool) *UserPermissionsCreate {
	upc.mutation.SetAdmin(b)
	return upc
}

// SetNillableAdmin sets the "Admin" field if the given value is not nil.
func (upc *UserPermissionsCreate) SetNillableAdmin(b *bool) *UserPermissionsCreate {
	if b != nil {
		upc.SetAdmin(*b)
	}
	return upc
}

// SetCanCreatePublic sets the "CanCreatePublic" field.
func (upc *UserPermissionsCreate) SetCanCreatePublic(b bool) *UserPermissionsCreate {
	upc.mutation.SetCanCreatePublic(b)
	return upc
}

// SetNillableCanCreatePublic sets the "CanCreatePublic" field if the given value is not nil.
func (upc *UserPermissionsCreate) SetNillableCanCreatePublic(b *bool) *UserPermissionsCreate {
	if b != nil {
		upc.SetCanCreatePublic(*b)
	}
	return upc
}

// SetID sets the "id" field.
func (upc *UserPermissionsCreate) SetID(k ksuid.ID) *UserPermissionsCreate {
	upc.mutation.SetID(k)
	return upc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (upc *UserPermissionsCreate) SetNillableID(k *ksuid.ID) *UserPermissionsCreate {
	if k != nil {
		upc.SetID(*k)
	}
	return upc
}

// SetUser sets the "user" edge to the User entity.
func (upc *UserPermissionsCreate) SetUser(u *User) *UserPermissionsCreate {
	return upc.SetUserID(u.ID)
}

// Mutation returns the UserPermissionsMutation object of the builder.
func (upc *UserPermissionsCreate) Mutation() *UserPermissionsMutation {
	return upc.mutation
}

// Save creates the UserPermissions in the database.
func (upc *UserPermissionsCreate) Save(ctx context.Context) (*UserPermissions, error) {
	if err := upc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, upc.sqlSave, upc.mutation, upc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (upc *UserPermissionsCreate) SaveX(ctx context.Context) *UserPermissions {
	v, err := upc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upc *UserPermissionsCreate) Exec(ctx context.Context) error {
	_, err := upc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upc *UserPermissionsCreate) ExecX(ctx context.Context) {
	if err := upc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upc *UserPermissionsCreate) defaults() error {
	if _, ok := upc.mutation.CanEdit(); !ok {
		v := userpermissions.DefaultCanEdit
		upc.mutation.SetCanEdit(v)
	}
	if _, ok := upc.mutation.Admin(); !ok {
		v := userpermissions.DefaultAdmin
		upc.mutation.SetAdmin(v)
	}
	if _, ok := upc.mutation.CanCreatePublic(); !ok {
		v := userpermissions.DefaultCanCreatePublic
		upc.mutation.SetCanCreatePublic(v)
	}
	if _, ok := upc.mutation.ID(); !ok {
		if userpermissions.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized userpermissions.DefaultID (forgotten import ent/runtime?)")
		}
		v := userpermissions.DefaultID()
		upc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (upc *UserPermissionsCreate) check() error {
	if _, ok := upc.mutation.CanEdit(); !ok {
		return &ValidationError{Name: "CanEdit", err: errors.New(`ent: missing required field "UserPermissions.CanEdit"`)}
	}
	if _, ok := upc.mutation.Admin(); !ok {
		return &ValidationError{Name: "Admin", err: errors.New(`ent: missing required field "UserPermissions.Admin"`)}
	}
	if _, ok := upc.mutation.CanCreatePublic(); !ok {
		return &ValidationError{Name: "CanCreatePublic", err: errors.New(`ent: missing required field "UserPermissions.CanCreatePublic"`)}
	}
	return nil
}

func (upc *UserPermissionsCreate) sqlSave(ctx context.Context) (*UserPermissions, error) {
	if err := upc.check(); err != nil {
		return nil, err
	}
	_node, _spec := upc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(ksuid.ID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserPermissions.ID type: %T", _spec.ID.Value)
		}
	}
	upc.mutation.id = &_node.ID
	upc.mutation.done = true
	return _node, nil
}

func (upc *UserPermissionsCreate) createSpec() (*UserPermissions, *sqlgraph.CreateSpec) {
	var (
		_node = &UserPermissions{config: upc.config}
		_spec = sqlgraph.NewCreateSpec(userpermissions.Table, sqlgraph.NewFieldSpec(userpermissions.FieldID, field.TypeString))
	)
	_spec.OnConflict = upc.conflict
	if id, ok := upc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := upc.mutation.CanEdit(); ok {
		_spec.SetField(userpermissions.FieldCanEdit, field.TypeBool, value)
		_node.CanEdit = value
	}
	if value, ok := upc.mutation.Admin(); ok {
		_spec.SetField(userpermissions.FieldAdmin, field.TypeBool, value)
		_node.Admin = value
	}
	if value, ok := upc.mutation.CanCreatePublic(); ok {
		_spec.SetField(userpermissions.FieldCanCreatePublic, field.TypeBool, value)
		_node.CanCreatePublic = value
	}
	if nodes := upc.mutation.UserIDs(); len(nodes) > 0 {
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserPermissions.Create().
//		SetUserID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserPermissionsUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (upc *UserPermissionsCreate) OnConflict(opts ...sql.ConflictOption) *UserPermissionsUpsertOne {
	upc.conflict = opts
	return &UserPermissionsUpsertOne{
		create: upc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserPermissions.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (upc *UserPermissionsCreate) OnConflictColumns(columns ...string) *UserPermissionsUpsertOne {
	upc.conflict = append(upc.conflict, sql.ConflictColumns(columns...))
	return &UserPermissionsUpsertOne{
		create: upc,
	}
}

type (
	// UserPermissionsUpsertOne is the builder for "upsert"-ing
	//  one UserPermissions node.
	UserPermissionsUpsertOne struct {
		create *UserPermissionsCreate
	}

	// UserPermissionsUpsert is the "OnConflict" setter.
	UserPermissionsUpsert struct {
		*sql.UpdateSet
	}
)

// SetUserID sets the "user_id" field.
func (u *UserPermissionsUpsert) SetUserID(v ksuid.ID) *UserPermissionsUpsert {
	u.Set(userpermissions.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserPermissionsUpsert) UpdateUserID() *UserPermissionsUpsert {
	u.SetExcluded(userpermissions.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserPermissionsUpsert) ClearUserID() *UserPermissionsUpsert {
	u.SetNull(userpermissions.FieldUserID)
	return u
}

// SetCanEdit sets the "CanEdit" field.
func (u *UserPermissionsUpsert) SetCanEdit(v bool) *UserPermissionsUpsert {
	u.Set(userpermissions.FieldCanEdit, v)
	return u
}

// UpdateCanEdit sets the "CanEdit" field to the value that was provided on create.
func (u *UserPermissionsUpsert) UpdateCanEdit() *UserPermissionsUpsert {
	u.SetExcluded(userpermissions.FieldCanEdit)
	return u
}

// SetAdmin sets the "Admin" field.
func (u *UserPermissionsUpsert) SetAdmin(v bool) *UserPermissionsUpsert {
	u.Set(userpermissions.FieldAdmin, v)
	return u
}

// UpdateAdmin sets the "Admin" field to the value that was provided on create.
func (u *UserPermissionsUpsert) UpdateAdmin() *UserPermissionsUpsert {
	u.SetExcluded(userpermissions.FieldAdmin)
	return u
}

// SetCanCreatePublic sets the "CanCreatePublic" field.
func (u *UserPermissionsUpsert) SetCanCreatePublic(v bool) *UserPermissionsUpsert {
	u.Set(userpermissions.FieldCanCreatePublic, v)
	return u
}

// UpdateCanCreatePublic sets the "CanCreatePublic" field to the value that was provided on create.
func (u *UserPermissionsUpsert) UpdateCanCreatePublic() *UserPermissionsUpsert {
	u.SetExcluded(userpermissions.FieldCanCreatePublic)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UserPermissions.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userpermissions.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserPermissionsUpsertOne) UpdateNewValues() *UserPermissionsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(userpermissions.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserPermissions.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserPermissionsUpsertOne) Ignore() *UserPermissionsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserPermissionsUpsertOne) DoNothing() *UserPermissionsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserPermissionsCreate.OnConflict
// documentation for more info.
func (u *UserPermissionsUpsertOne) Update(set func(*UserPermissionsUpsert)) *UserPermissionsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserPermissionsUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserPermissionsUpsertOne) SetUserID(v ksuid.ID) *UserPermissionsUpsertOne {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserPermissionsUpsertOne) UpdateUserID() *UserPermissionsUpsertOne {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserPermissionsUpsertOne) ClearUserID() *UserPermissionsUpsertOne {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.ClearUserID()
	})
}

// SetCanEdit sets the "CanEdit" field.
func (u *UserPermissionsUpsertOne) SetCanEdit(v bool) *UserPermissionsUpsertOne {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.SetCanEdit(v)
	})
}

// UpdateCanEdit sets the "CanEdit" field to the value that was provided on create.
func (u *UserPermissionsUpsertOne) UpdateCanEdit() *UserPermissionsUpsertOne {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.UpdateCanEdit()
	})
}

// SetAdmin sets the "Admin" field.
func (u *UserPermissionsUpsertOne) SetAdmin(v bool) *UserPermissionsUpsertOne {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.SetAdmin(v)
	})
}

// UpdateAdmin sets the "Admin" field to the value that was provided on create.
func (u *UserPermissionsUpsertOne) UpdateAdmin() *UserPermissionsUpsertOne {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.UpdateAdmin()
	})
}

// SetCanCreatePublic sets the "CanCreatePublic" field.
func (u *UserPermissionsUpsertOne) SetCanCreatePublic(v bool) *UserPermissionsUpsertOne {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.SetCanCreatePublic(v)
	})
}

// UpdateCanCreatePublic sets the "CanCreatePublic" field to the value that was provided on create.
func (u *UserPermissionsUpsertOne) UpdateCanCreatePublic() *UserPermissionsUpsertOne {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.UpdateCanCreatePublic()
	})
}

// Exec executes the query.
func (u *UserPermissionsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserPermissionsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserPermissionsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserPermissionsUpsertOne) ID(ctx context.Context) (id ksuid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserPermissionsUpsertOne.ID is not supported by MySQL driver. Use UserPermissionsUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserPermissionsUpsertOne) IDX(ctx context.Context) ksuid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserPermissionsCreateBulk is the builder for creating many UserPermissions entities in bulk.
type UserPermissionsCreateBulk struct {
	config
	err      error
	builders []*UserPermissionsCreate
	conflict []sql.ConflictOption
}

// Save creates the UserPermissions entities in the database.
func (upcb *UserPermissionsCreateBulk) Save(ctx context.Context) ([]*UserPermissions, error) {
	if upcb.err != nil {
		return nil, upcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(upcb.builders))
	nodes := make([]*UserPermissions, len(upcb.builders))
	mutators := make([]Mutator, len(upcb.builders))
	for i := range upcb.builders {
		func(i int, root context.Context) {
			builder := upcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserPermissionsMutation)
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
					_, err = mutators[i+1].Mutate(root, upcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = upcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, upcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upcb *UserPermissionsCreateBulk) SaveX(ctx context.Context) []*UserPermissions {
	v, err := upcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upcb *UserPermissionsCreateBulk) Exec(ctx context.Context) error {
	_, err := upcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upcb *UserPermissionsCreateBulk) ExecX(ctx context.Context) {
	if err := upcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserPermissions.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserPermissionsUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (upcb *UserPermissionsCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserPermissionsUpsertBulk {
	upcb.conflict = opts
	return &UserPermissionsUpsertBulk{
		create: upcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserPermissions.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (upcb *UserPermissionsCreateBulk) OnConflictColumns(columns ...string) *UserPermissionsUpsertBulk {
	upcb.conflict = append(upcb.conflict, sql.ConflictColumns(columns...))
	return &UserPermissionsUpsertBulk{
		create: upcb,
	}
}

// UserPermissionsUpsertBulk is the builder for "upsert"-ing
// a bulk of UserPermissions nodes.
type UserPermissionsUpsertBulk struct {
	create *UserPermissionsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserPermissions.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userpermissions.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserPermissionsUpsertBulk) UpdateNewValues() *UserPermissionsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(userpermissions.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserPermissions.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserPermissionsUpsertBulk) Ignore() *UserPermissionsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserPermissionsUpsertBulk) DoNothing() *UserPermissionsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserPermissionsCreateBulk.OnConflict
// documentation for more info.
func (u *UserPermissionsUpsertBulk) Update(set func(*UserPermissionsUpsert)) *UserPermissionsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserPermissionsUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserPermissionsUpsertBulk) SetUserID(v ksuid.ID) *UserPermissionsUpsertBulk {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserPermissionsUpsertBulk) UpdateUserID() *UserPermissionsUpsertBulk {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *UserPermissionsUpsertBulk) ClearUserID() *UserPermissionsUpsertBulk {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.ClearUserID()
	})
}

// SetCanEdit sets the "CanEdit" field.
func (u *UserPermissionsUpsertBulk) SetCanEdit(v bool) *UserPermissionsUpsertBulk {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.SetCanEdit(v)
	})
}

// UpdateCanEdit sets the "CanEdit" field to the value that was provided on create.
func (u *UserPermissionsUpsertBulk) UpdateCanEdit() *UserPermissionsUpsertBulk {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.UpdateCanEdit()
	})
}

// SetAdmin sets the "Admin" field.
func (u *UserPermissionsUpsertBulk) SetAdmin(v bool) *UserPermissionsUpsertBulk {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.SetAdmin(v)
	})
}

// UpdateAdmin sets the "Admin" field to the value that was provided on create.
func (u *UserPermissionsUpsertBulk) UpdateAdmin() *UserPermissionsUpsertBulk {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.UpdateAdmin()
	})
}

// SetCanCreatePublic sets the "CanCreatePublic" field.
func (u *UserPermissionsUpsertBulk) SetCanCreatePublic(v bool) *UserPermissionsUpsertBulk {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.SetCanCreatePublic(v)
	})
}

// UpdateCanCreatePublic sets the "CanCreatePublic" field to the value that was provided on create.
func (u *UserPermissionsUpsertBulk) UpdateCanCreatePublic() *UserPermissionsUpsertBulk {
	return u.Update(func(s *UserPermissionsUpsert) {
		s.UpdateCanCreatePublic()
	})
}

// Exec executes the query.
func (u *UserPermissionsUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserPermissionsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserPermissionsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserPermissionsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
