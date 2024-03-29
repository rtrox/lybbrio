// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/series"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SeriesCreate is the builder for creating a Series entity.
type SeriesCreate struct {
	config
	mutation *SeriesMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (sc *SeriesCreate) SetCreateTime(t time.Time) *SeriesCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableCreateTime(t *time.Time) *SeriesCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *SeriesCreate) SetUpdateTime(t time.Time) *SeriesCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableUpdateTime(t *time.Time) *SeriesCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetCalibreID sets the "calibre_id" field.
func (sc *SeriesCreate) SetCalibreID(i int64) *SeriesCreate {
	sc.mutation.SetCalibreID(i)
	return sc
}

// SetNillableCalibreID sets the "calibre_id" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableCalibreID(i *int64) *SeriesCreate {
	if i != nil {
		sc.SetCalibreID(*i)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SeriesCreate) SetName(s string) *SeriesCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetSort sets the "sort" field.
func (sc *SeriesCreate) SetSort(s string) *SeriesCreate {
	sc.mutation.SetSort(s)
	return sc
}

// SetID sets the "id" field.
func (sc *SeriesCreate) SetID(k ksuid.ID) *SeriesCreate {
	sc.mutation.SetID(k)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableID(k *ksuid.ID) *SeriesCreate {
	if k != nil {
		sc.SetID(*k)
	}
	return sc
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (sc *SeriesCreate) AddBookIDs(ids ...ksuid.ID) *SeriesCreate {
	sc.mutation.AddBookIDs(ids...)
	return sc
}

// AddBooks adds the "books" edges to the Book entity.
func (sc *SeriesCreate) AddBooks(b ...*Book) *SeriesCreate {
	ids := make([]ksuid.ID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return sc.AddBookIDs(ids...)
}

// Mutation returns the SeriesMutation object of the builder.
func (sc *SeriesCreate) Mutation() *SeriesMutation {
	return sc.mutation
}

// Save creates the Series in the database.
func (sc *SeriesCreate) Save(ctx context.Context) (*Series, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SeriesCreate) SaveX(ctx context.Context) *Series {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SeriesCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SeriesCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SeriesCreate) defaults() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		if series.DefaultCreateTime == nil {
			return fmt.Errorf("ent: uninitialized series.DefaultCreateTime (forgotten import ent/runtime?)")
		}
		v := series.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		if series.DefaultUpdateTime == nil {
			return fmt.Errorf("ent: uninitialized series.DefaultUpdateTime (forgotten import ent/runtime?)")
		}
		v := series.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		if series.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized series.DefaultID (forgotten import ent/runtime?)")
		}
		v := series.DefaultID()
		sc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SeriesCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Series.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Series.update_time"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Series.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := series.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Series.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Series.sort"`)}
	}
	if v, ok := sc.mutation.Sort(); ok {
		if err := series.SortValidator(v); err != nil {
			return &ValidationError{Name: "sort", err: fmt.Errorf(`ent: validator failed for field "Series.sort": %w`, err)}
		}
	}
	return nil
}

func (sc *SeriesCreate) sqlSave(ctx context.Context) (*Series, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(ksuid.ID); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Series.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SeriesCreate) createSpec() (*Series, *sqlgraph.CreateSpec) {
	var (
		_node = &Series{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(series.Table, sqlgraph.NewFieldSpec(series.FieldID, field.TypeString))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(series.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.SetField(series.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.CalibreID(); ok {
		_spec.SetField(series.FieldCalibreID, field.TypeInt64, value)
		_node.CalibreID = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(series.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Sort(); ok {
		_spec.SetField(series.FieldSort, field.TypeString, value)
		_node.Sort = value
	}
	if nodes := sc.mutation.BooksIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Series.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SeriesUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (sc *SeriesCreate) OnConflict(opts ...sql.ConflictOption) *SeriesUpsertOne {
	sc.conflict = opts
	return &SeriesUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Series.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SeriesCreate) OnConflictColumns(columns ...string) *SeriesUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SeriesUpsertOne{
		create: sc,
	}
}

type (
	// SeriesUpsertOne is the builder for "upsert"-ing
	//  one Series node.
	SeriesUpsertOne struct {
		create *SeriesCreate
	}

	// SeriesUpsert is the "OnConflict" setter.
	SeriesUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *SeriesUpsert) SetUpdateTime(v time.Time) *SeriesUpsert {
	u.Set(series.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SeriesUpsert) UpdateUpdateTime() *SeriesUpsert {
	u.SetExcluded(series.FieldUpdateTime)
	return u
}

// SetCalibreID sets the "calibre_id" field.
func (u *SeriesUpsert) SetCalibreID(v int64) *SeriesUpsert {
	u.Set(series.FieldCalibreID, v)
	return u
}

// UpdateCalibreID sets the "calibre_id" field to the value that was provided on create.
func (u *SeriesUpsert) UpdateCalibreID() *SeriesUpsert {
	u.SetExcluded(series.FieldCalibreID)
	return u
}

// AddCalibreID adds v to the "calibre_id" field.
func (u *SeriesUpsert) AddCalibreID(v int64) *SeriesUpsert {
	u.Add(series.FieldCalibreID, v)
	return u
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (u *SeriesUpsert) ClearCalibreID() *SeriesUpsert {
	u.SetNull(series.FieldCalibreID)
	return u
}

// SetName sets the "name" field.
func (u *SeriesUpsert) SetName(v string) *SeriesUpsert {
	u.Set(series.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SeriesUpsert) UpdateName() *SeriesUpsert {
	u.SetExcluded(series.FieldName)
	return u
}

// SetSort sets the "sort" field.
func (u *SeriesUpsert) SetSort(v string) *SeriesUpsert {
	u.Set(series.FieldSort, v)
	return u
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SeriesUpsert) UpdateSort() *SeriesUpsert {
	u.SetExcluded(series.FieldSort)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Series.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(series.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SeriesUpsertOne) UpdateNewValues() *SeriesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(series.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(series.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Series.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SeriesUpsertOne) Ignore() *SeriesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SeriesUpsertOne) DoNothing() *SeriesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SeriesCreate.OnConflict
// documentation for more info.
func (u *SeriesUpsertOne) Update(set func(*SeriesUpsert)) *SeriesUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SeriesUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *SeriesUpsertOne) SetUpdateTime(v time.Time) *SeriesUpsertOne {
	return u.Update(func(s *SeriesUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SeriesUpsertOne) UpdateUpdateTime() *SeriesUpsertOne {
	return u.Update(func(s *SeriesUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetCalibreID sets the "calibre_id" field.
func (u *SeriesUpsertOne) SetCalibreID(v int64) *SeriesUpsertOne {
	return u.Update(func(s *SeriesUpsert) {
		s.SetCalibreID(v)
	})
}

// AddCalibreID adds v to the "calibre_id" field.
func (u *SeriesUpsertOne) AddCalibreID(v int64) *SeriesUpsertOne {
	return u.Update(func(s *SeriesUpsert) {
		s.AddCalibreID(v)
	})
}

// UpdateCalibreID sets the "calibre_id" field to the value that was provided on create.
func (u *SeriesUpsertOne) UpdateCalibreID() *SeriesUpsertOne {
	return u.Update(func(s *SeriesUpsert) {
		s.UpdateCalibreID()
	})
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (u *SeriesUpsertOne) ClearCalibreID() *SeriesUpsertOne {
	return u.Update(func(s *SeriesUpsert) {
		s.ClearCalibreID()
	})
}

// SetName sets the "name" field.
func (u *SeriesUpsertOne) SetName(v string) *SeriesUpsertOne {
	return u.Update(func(s *SeriesUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SeriesUpsertOne) UpdateName() *SeriesUpsertOne {
	return u.Update(func(s *SeriesUpsert) {
		s.UpdateName()
	})
}

// SetSort sets the "sort" field.
func (u *SeriesUpsertOne) SetSort(v string) *SeriesUpsertOne {
	return u.Update(func(s *SeriesUpsert) {
		s.SetSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SeriesUpsertOne) UpdateSort() *SeriesUpsertOne {
	return u.Update(func(s *SeriesUpsert) {
		s.UpdateSort()
	})
}

// Exec executes the query.
func (u *SeriesUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SeriesCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SeriesUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SeriesUpsertOne) ID(ctx context.Context) (id ksuid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SeriesUpsertOne.ID is not supported by MySQL driver. Use SeriesUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SeriesUpsertOne) IDX(ctx context.Context) ksuid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SeriesCreateBulk is the builder for creating many Series entities in bulk.
type SeriesCreateBulk struct {
	config
	err      error
	builders []*SeriesCreate
	conflict []sql.ConflictOption
}

// Save creates the Series entities in the database.
func (scb *SeriesCreateBulk) Save(ctx context.Context) ([]*Series, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Series, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SeriesMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SeriesCreateBulk) SaveX(ctx context.Context) []*Series {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SeriesCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SeriesCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Series.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SeriesUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (scb *SeriesCreateBulk) OnConflict(opts ...sql.ConflictOption) *SeriesUpsertBulk {
	scb.conflict = opts
	return &SeriesUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Series.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SeriesCreateBulk) OnConflictColumns(columns ...string) *SeriesUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SeriesUpsertBulk{
		create: scb,
	}
}

// SeriesUpsertBulk is the builder for "upsert"-ing
// a bulk of Series nodes.
type SeriesUpsertBulk struct {
	create *SeriesCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Series.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(series.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SeriesUpsertBulk) UpdateNewValues() *SeriesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(series.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(series.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Series.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SeriesUpsertBulk) Ignore() *SeriesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SeriesUpsertBulk) DoNothing() *SeriesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SeriesCreateBulk.OnConflict
// documentation for more info.
func (u *SeriesUpsertBulk) Update(set func(*SeriesUpsert)) *SeriesUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SeriesUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *SeriesUpsertBulk) SetUpdateTime(v time.Time) *SeriesUpsertBulk {
	return u.Update(func(s *SeriesUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SeriesUpsertBulk) UpdateUpdateTime() *SeriesUpsertBulk {
	return u.Update(func(s *SeriesUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetCalibreID sets the "calibre_id" field.
func (u *SeriesUpsertBulk) SetCalibreID(v int64) *SeriesUpsertBulk {
	return u.Update(func(s *SeriesUpsert) {
		s.SetCalibreID(v)
	})
}

// AddCalibreID adds v to the "calibre_id" field.
func (u *SeriesUpsertBulk) AddCalibreID(v int64) *SeriesUpsertBulk {
	return u.Update(func(s *SeriesUpsert) {
		s.AddCalibreID(v)
	})
}

// UpdateCalibreID sets the "calibre_id" field to the value that was provided on create.
func (u *SeriesUpsertBulk) UpdateCalibreID() *SeriesUpsertBulk {
	return u.Update(func(s *SeriesUpsert) {
		s.UpdateCalibreID()
	})
}

// ClearCalibreID clears the value of the "calibre_id" field.
func (u *SeriesUpsertBulk) ClearCalibreID() *SeriesUpsertBulk {
	return u.Update(func(s *SeriesUpsert) {
		s.ClearCalibreID()
	})
}

// SetName sets the "name" field.
func (u *SeriesUpsertBulk) SetName(v string) *SeriesUpsertBulk {
	return u.Update(func(s *SeriesUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SeriesUpsertBulk) UpdateName() *SeriesUpsertBulk {
	return u.Update(func(s *SeriesUpsert) {
		s.UpdateName()
	})
}

// SetSort sets the "sort" field.
func (u *SeriesUpsertBulk) SetSort(v string) *SeriesUpsertBulk {
	return u.Update(func(s *SeriesUpsert) {
		s.SetSort(v)
	})
}

// UpdateSort sets the "sort" field to the value that was provided on create.
func (u *SeriesUpsertBulk) UpdateSort() *SeriesUpsertBulk {
	return u.Update(func(s *SeriesUpsert) {
		s.UpdateSort()
	})
}

// Exec executes the query.
func (u *SeriesUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SeriesCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SeriesCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SeriesUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
