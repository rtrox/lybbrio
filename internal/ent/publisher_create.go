// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/publisher"
	"lybbrio/internal/ent/schema/ksuid"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PublisherCreate is the builder for creating a Publisher entity.
type PublisherCreate struct {
	config
	mutation *PublisherMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (pc *PublisherCreate) SetCreateTime(t time.Time) *PublisherCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableCreateTime(t *time.Time) *PublisherCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PublisherCreate) SetUpdateTime(t time.Time) *PublisherCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableUpdateTime(t *time.Time) *PublisherCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetCalibreID sets the "calibre_id" field.
func (pc *PublisherCreate) SetCalibreID(i int64) *PublisherCreate {
	pc.mutation.SetCalibreID(i)
	return pc
}

// SetNillableCalibreID sets the "calibre_id" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableCalibreID(i *int64) *PublisherCreate {
	if i != nil {
		pc.SetCalibreID(*i)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PublisherCreate) SetName(s string) *PublisherCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PublisherCreate) SetID(k ksuid.ID) *PublisherCreate {
	pc.mutation.SetID(k)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PublisherCreate) SetNillableID(k *ksuid.ID) *PublisherCreate {
	if k != nil {
		pc.SetID(*k)
	}
	return pc
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (pc *PublisherCreate) AddBookIDs(ids ...ksuid.ID) *PublisherCreate {
	pc.mutation.AddBookIDs(ids...)
	return pc
}

// AddBooks adds the "books" edges to the Book entity.
func (pc *PublisherCreate) AddBooks(b ...*Book) *PublisherCreate {
	ids := make([]ksuid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddBookIDs(ids...)
}

// Mutation returns the PublisherMutation object of the builder.
func (pc *PublisherCreate) Mutation() *PublisherMutation {
	return pc.mutation
}

// Save creates the Publisher in the database.
func (pc *PublisherCreate) Save(ctx context.Context) (*Publisher, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PublisherCreate) SaveX(ctx context.Context) *Publisher {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PublisherCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PublisherCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PublisherCreate) defaults() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		if publisher.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized publisher.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := publisher.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		if publisher.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized publisher.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := publisher.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		if publisher.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized publisher.DefaultID (forgotten import ent/runtime?)")
		}
		v := publisher.DefaultID()
		pc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PublisherCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Publisher.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Publisher.update_time"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Publisher.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := publisher.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Publisher.name": %w`, err)}
		}
	}
	return nil
}

func (pc *PublisherCreate) sqlSave(ctx context.Context) (*Publisher, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(ksuid.ID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Publisher.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PublisherCreate) createSpec() (*Publisher, *sqlgraph.CreateSpec) {
	var (
		_node = &Publisher{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(publisher.Table, sqlgraph.NewFieldSpec(publisher.FieldID, field.TypeString))
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(publisher.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(publisher.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.CalibreID(); ok {
		_spec.SetField(publisher.FieldCalibreID, field.TypeInt64, value)
		_node.CalibreID = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(publisher.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := pc.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   publisher.BooksTable,
			Columns: publisher.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeString),
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
//	client.Publisher.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PublisherUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pc *PublisherCreate) OnConflict(opts ...sql.ConflictOption) *PublisherUpsertOne {
	pc.conflict = opts
	return &PublisherUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Publisher.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PublisherCreate) OnConflictColumns(columns ...string) *PublisherUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PublisherUpsertOne{
		create: pc,
	}
}

type (
	// PublisherUpsertOne is the builder for "upsert"-ing
	//  one Publisher node.
	PublisherUpsertOne struct {
		create *PublisherCreate
	}

	// PublisherUpsert is the "OnConflict" setter.
	PublisherUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *PublisherUpsert) SetUpdateTime(v time.Time) *PublisherUpsert {
	u.Set(publisher.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PublisherUpsert) UpdateUpdateTime() *PublisherUpsert {
	u.SetExcluded(publisher.FieldUpdateTime)
	return u
}

// SetCalibreID sets the "calibre_id" field.
func (u *PublisherUpsert) SetCalibreID(v int64) *PublisherUpsert {
	u.Set(publisher.FieldCalibreID, v)
	return u
}

// UpdateCalibreID sets the "calibre_id" field to the value that was provided on create.
func (u *PublisherUpsert) UpdateCalibreID() *PublisherUpsert {
	u.SetExcluded(publisher.FieldCalibreID)
	return u
}

// AddCalibreID adds v to the "calibre_id" field.
func (u *PublisherUpsert) AddCalibreID(v int64) *PublisherUpsert {
	u.Add(publisher.FieldCalibreID, v)
	return u
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (u *PublisherUpsert) ClearCalibreID() *PublisherUpsert {
	u.SetNull(publisher.FieldCalibreID)
	return u
}

// SetName sets the "name" field.
func (u *PublisherUpsert) SetName(v string) *PublisherUpsert {
	u.Set(publisher.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PublisherUpsert) UpdateName() *PublisherUpsert {
	u.SetExcluded(publisher.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Publisher.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(publisher.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PublisherUpsertOne) UpdateNewValues() *PublisherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(publisher.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(publisher.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Publisher.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PublisherUpsertOne) Ignore() *PublisherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PublisherUpsertOne) DoNothing() *PublisherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PublisherCreate.OnConflict
// documentation for more info.
func (u *PublisherUpsertOne) Update(set func(*PublisherUpsert)) *PublisherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PublisherUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *PublisherUpsertOne) SetUpdateTime(v time.Time) *PublisherUpsertOne {
	return u.Update(func(s *PublisherUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PublisherUpsertOne) UpdateUpdateTime() *PublisherUpsertOne {
	return u.Update(func(s *PublisherUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetCalibreID sets the "calibre_id" field.
func (u *PublisherUpsertOne) SetCalibreID(v int64) *PublisherUpsertOne {
	return u.Update(func(s *PublisherUpsert) {
		s.SetCalibreID(v)
	})
}

// AddCalibreID adds v to the "calibre_id" field.
func (u *PublisherUpsertOne) AddCalibreID(v int64) *PublisherUpsertOne {
	return u.Update(func(s *PublisherUpsert) {
		s.AddCalibreID(v)
	})
}

// UpdateCalibreID sets the "calibre_id" field to the value that was provided on create.
func (u *PublisherUpsertOne) UpdateCalibreID() *PublisherUpsertOne {
	return u.Update(func(s *PublisherUpsert) {
		s.UpdateCalibreID()
	})
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (u *PublisherUpsertOne) ClearCalibreID() *PublisherUpsertOne {
	return u.Update(func(s *PublisherUpsert) {
		s.ClearCalibreID()
	})
}

// SetName sets the "name" field.
func (u *PublisherUpsertOne) SetName(v string) *PublisherUpsertOne {
	return u.Update(func(s *PublisherUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PublisherUpsertOne) UpdateName() *PublisherUpsertOne {
	return u.Update(func(s *PublisherUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *PublisherUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PublisherCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PublisherUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PublisherUpsertOne) ID(ctx context.Context) (id ksuid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PublisherUpsertOne.ID is not supported by MySQL driver. Use PublisherUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PublisherUpsertOne) IDX(ctx context.Context) ksuid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PublisherCreateBulk is the builder for creating many Publisher entities in bulk.
type PublisherCreateBulk struct {
	config
	err      error
	builders []*PublisherCreate
	conflict []sql.ConflictOption
}

// Save creates the Publisher entities in the database.
func (pcb *PublisherCreateBulk) Save(ctx context.Context) ([]*Publisher, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Publisher, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PublisherMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PublisherCreateBulk) SaveX(ctx context.Context) []*Publisher {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PublisherCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PublisherCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Publisher.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PublisherUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (pcb *PublisherCreateBulk) OnConflict(opts ...sql.ConflictOption) *PublisherUpsertBulk {
	pcb.conflict = opts
	return &PublisherUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Publisher.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PublisherCreateBulk) OnConflictColumns(columns ...string) *PublisherUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PublisherUpsertBulk{
		create: pcb,
	}
}

// PublisherUpsertBulk is the builder for "upsert"-ing
// a bulk of Publisher nodes.
type PublisherUpsertBulk struct {
	create *PublisherCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Publisher.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(publisher.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PublisherUpsertBulk) UpdateNewValues() *PublisherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(publisher.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(publisher.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Publisher.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PublisherUpsertBulk) Ignore() *PublisherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PublisherUpsertBulk) DoNothing() *PublisherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PublisherCreateBulk.OnConflict
// documentation for more info.
func (u *PublisherUpsertBulk) Update(set func(*PublisherUpsert)) *PublisherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PublisherUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *PublisherUpsertBulk) SetUpdateTime(v time.Time) *PublisherUpsertBulk {
	return u.Update(func(s *PublisherUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *PublisherUpsertBulk) UpdateUpdateTime() *PublisherUpsertBulk {
	return u.Update(func(s *PublisherUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetCalibreID sets the "calibre_id" field.
func (u *PublisherUpsertBulk) SetCalibreID(v int64) *PublisherUpsertBulk {
	return u.Update(func(s *PublisherUpsert) {
		s.SetCalibreID(v)
	})
}

// AddCalibreID adds v to the "calibre_id" field.
func (u *PublisherUpsertBulk) AddCalibreID(v int64) *PublisherUpsertBulk {
	return u.Update(func(s *PublisherUpsert) {
		s.AddCalibreID(v)
	})
}

// UpdateCalibreID sets the "calibre_id" field to the value that was provided on create.
func (u *PublisherUpsertBulk) UpdateCalibreID() *PublisherUpsertBulk {
	return u.Update(func(s *PublisherUpsert) {
		s.UpdateCalibreID()
	})
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (u *PublisherUpsertBulk) ClearCalibreID() *PublisherUpsertBulk {
	return u.Update(func(s *PublisherUpsert) {
		s.ClearCalibreID()
	})
}

// SetName sets the "name" field.
func (u *PublisherUpsertBulk) SetName(v string) *PublisherUpsertBulk {
	return u.Update(func(s *PublisherUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PublisherUpsertBulk) UpdateName() *PublisherUpsertBulk {
	return u.Update(func(s *PublisherUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *PublisherUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PublisherCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PublisherCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PublisherUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
