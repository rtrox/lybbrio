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
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPermissionsQuery is the builder for querying UserPermissions entities.
type UserPermissionsQuery struct {
	config
	ctx        *QueryContext
	order      []userpermissions.OrderOption
	inters     []Interceptor
	predicates []predicate.UserPermissions
	withUser   *UserQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*UserPermissions) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserPermissionsQuery builder.
func (upq *UserPermissionsQuery) Where(ps ...predicate.UserPermissions) *UserPermissionsQuery {
	upq.predicates = append(upq.predicates, ps...)
	return upq
}

// Limit the number of records to be returned by this query.
func (upq *UserPermissionsQuery) Limit(limit int) *UserPermissionsQuery {
	upq.ctx.Limit = &limit
	return upq
}

// Offset to start from.
func (upq *UserPermissionsQuery) Offset(offset int) *UserPermissionsQuery {
	upq.ctx.Offset = &offset
	return upq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (upq *UserPermissionsQuery) Unique(unique bool) *UserPermissionsQuery {
	upq.ctx.Unique = &unique
	return upq
}

// Order specifies how the records should be ordered.
func (upq *UserPermissionsQuery) Order(o ...userpermissions.OrderOption) *UserPermissionsQuery {
	upq.order = append(upq.order, o...)
	return upq
}

// QueryUser chains the current query on the "user" edge.
func (upq *UserPermissionsQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: upq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := upq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := upq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userpermissions.Table, userpermissions.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, userpermissions.UserTable, userpermissions.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(upq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserPermissions entity from the query.
// Returns a *NotFoundError when no UserPermissions was found.
func (upq *UserPermissionsQuery) First(ctx context.Context) (*UserPermissions, error) {
	nodes, err := upq.Limit(1).All(setContextOp(ctx, upq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userpermissions.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (upq *UserPermissionsQuery) FirstX(ctx context.Context) *UserPermissions {
	node, err := upq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserPermissions ID from the query.
// Returns a *NotFoundError when no UserPermissions ID was found.
func (upq *UserPermissionsQuery) FirstID(ctx context.Context) (id ksuid.ID, err error) {
	var ids []ksuid.ID
	if ids, err = upq.Limit(1).IDs(setContextOp(ctx, upq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userpermissions.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (upq *UserPermissionsQuery) FirstIDX(ctx context.Context) ksuid.ID {
	id, err := upq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserPermissions entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserPermissions entity is found.
// Returns a *NotFoundError when no UserPermissions entities are found.
func (upq *UserPermissionsQuery) Only(ctx context.Context) (*UserPermissions, error) {
	nodes, err := upq.Limit(2).All(setContextOp(ctx, upq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userpermissions.Label}
	default:
		return nil, &NotSingularError{userpermissions.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (upq *UserPermissionsQuery) OnlyX(ctx context.Context) *UserPermissions {
	node, err := upq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserPermissions ID in the query.
// Returns a *NotSingularError when more than one UserPermissions ID is found.
// Returns a *NotFoundError when no entities are found.
func (upq *UserPermissionsQuery) OnlyID(ctx context.Context) (id ksuid.ID, err error) {
	var ids []ksuid.ID
	if ids, err = upq.Limit(2).IDs(setContextOp(ctx, upq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userpermissions.Label}
	default:
		err = &NotSingularError{userpermissions.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (upq *UserPermissionsQuery) OnlyIDX(ctx context.Context) ksuid.ID {
	id, err := upq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserPermissionsSlice.
func (upq *UserPermissionsQuery) All(ctx context.Context) ([]*UserPermissions, error) {
	ctx = setContextOp(ctx, upq.ctx, "All")
	if err := upq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserPermissions, *UserPermissionsQuery]()
	return withInterceptors[[]*UserPermissions](ctx, upq, qr, upq.inters)
}

// AllX is like All, but panics if an error occurs.
func (upq *UserPermissionsQuery) AllX(ctx context.Context) []*UserPermissions {
	nodes, err := upq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserPermissions IDs.
func (upq *UserPermissionsQuery) IDs(ctx context.Context) (ids []ksuid.ID, err error) {
	if upq.ctx.Unique == nil && upq.path != nil {
		upq.Unique(true)
	}
	ctx = setContextOp(ctx, upq.ctx, "IDs")
	if err = upq.Select(userpermissions.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (upq *UserPermissionsQuery) IDsX(ctx context.Context) []ksuid.ID {
	ids, err := upq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (upq *UserPermissionsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, upq.ctx, "Count")
	if err := upq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, upq, querierCount[*UserPermissionsQuery](), upq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (upq *UserPermissionsQuery) CountX(ctx context.Context) int {
	count, err := upq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (upq *UserPermissionsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, upq.ctx, "Exist")
	switch _, err := upq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (upq *UserPermissionsQuery) ExistX(ctx context.Context) bool {
	exist, err := upq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserPermissionsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (upq *UserPermissionsQuery) Clone() *UserPermissionsQuery {
	if upq == nil {
		return nil
	}
	return &UserPermissionsQuery{
		config:     upq.config,
		ctx:        upq.ctx.Clone(),
		order:      append([]userpermissions.OrderOption{}, upq.order...),
		inters:     append([]Interceptor{}, upq.inters...),
		predicates: append([]predicate.UserPermissions{}, upq.predicates...),
		withUser:   upq.withUser.Clone(),
		// clone intermediate query.
		sql:  upq.sql.Clone(),
		path: upq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (upq *UserPermissionsQuery) WithUser(opts ...func(*UserQuery)) *UserPermissionsQuery {
	query := (&UserClient{config: upq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	upq.withUser = query
	return upq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserPermissions.Query().
//		GroupBy(userpermissions.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (upq *UserPermissionsQuery) GroupBy(field string, fields ...string) *UserPermissionsGroupBy {
	upq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserPermissionsGroupBy{build: upq}
	grbuild.flds = &upq.ctx.Fields
	grbuild.label = userpermissions.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.UserPermissions.Query().
//		Select(userpermissions.FieldCreateTime).
//		Scan(ctx, &v)
func (upq *UserPermissionsQuery) Select(fields ...string) *UserPermissionsSelect {
	upq.ctx.Fields = append(upq.ctx.Fields, fields...)
	sbuild := &UserPermissionsSelect{UserPermissionsQuery: upq}
	sbuild.label = userpermissions.Label
	sbuild.flds, sbuild.scan = &upq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserPermissionsSelect configured with the given aggregations.
func (upq *UserPermissionsQuery) Aggregate(fns ...AggregateFunc) *UserPermissionsSelect {
	return upq.Select().Aggregate(fns...)
}

func (upq *UserPermissionsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range upq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, upq); err != nil {
				return err
			}
		}
	}
	for _, f := range upq.ctx.Fields {
		if !userpermissions.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if upq.path != nil {
		prev, err := upq.path(ctx)
		if err != nil {
			return err
		}
		upq.sql = prev
	}
	if userpermissions.Policy == nil {
		return errors.New("ent: uninitialized userpermissions.Policy (forgotten import ent/runtime?)")
	}
	if err := userpermissions.Policy.EvalQuery(ctx, upq); err != nil {
		return err
	}
	return nil
}

func (upq *UserPermissionsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserPermissions, error) {
	var (
		nodes       = []*UserPermissions{}
		_spec       = upq.querySpec()
		loadedTypes = [1]bool{
			upq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserPermissions).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserPermissions{config: upq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(upq.modifiers) > 0 {
		_spec.Modifiers = upq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, upq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := upq.withUser; query != nil {
		if err := upq.loadUser(ctx, query, nodes, nil,
			func(n *UserPermissions, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	for i := range upq.loadTotal {
		if err := upq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (upq *UserPermissionsQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserPermissions, init func(*UserPermissions), assign func(*UserPermissions, *User)) error {
	ids := make([]ksuid.ID, 0, len(nodes))
	nodeids := make(map[ksuid.ID][]*UserPermissions)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (upq *UserPermissionsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := upq.querySpec()
	if len(upq.modifiers) > 0 {
		_spec.Modifiers = upq.modifiers
	}
	_spec.Node.Columns = upq.ctx.Fields
	if len(upq.ctx.Fields) > 0 {
		_spec.Unique = upq.ctx.Unique != nil && *upq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, upq.driver, _spec)
}

func (upq *UserPermissionsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userpermissions.Table, userpermissions.Columns, sqlgraph.NewFieldSpec(userpermissions.FieldID, field.TypeString))
	_spec.From = upq.sql
	if unique := upq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if upq.path != nil {
		_spec.Unique = true
	}
	if fields := upq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userpermissions.FieldID)
		for i := range fields {
			if fields[i] != userpermissions.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if upq.withUser != nil {
			_spec.Node.AddColumnOnce(userpermissions.FieldUserID)
		}
	}
	if ps := upq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := upq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := upq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := upq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (upq *UserPermissionsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(upq.driver.Dialect())
	t1 := builder.Table(userpermissions.Table)
	columns := upq.ctx.Fields
	if len(columns) == 0 {
		columns = userpermissions.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if upq.sql != nil {
		selector = upq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if upq.ctx.Unique != nil && *upq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range upq.predicates {
		p(selector)
	}
	for _, p := range upq.order {
		p(selector)
	}
	if offset := upq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := upq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserPermissionsGroupBy is the group-by builder for UserPermissions entities.
type UserPermissionsGroupBy struct {
	selector
	build *UserPermissionsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (upgb *UserPermissionsGroupBy) Aggregate(fns ...AggregateFunc) *UserPermissionsGroupBy {
	upgb.fns = append(upgb.fns, fns...)
	return upgb
}

// Scan applies the selector query and scans the result into the given value.
func (upgb *UserPermissionsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, upgb.build.ctx, "GroupBy")
	if err := upgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserPermissionsQuery, *UserPermissionsGroupBy](ctx, upgb.build, upgb, upgb.build.inters, v)
}

func (upgb *UserPermissionsGroupBy) sqlScan(ctx context.Context, root *UserPermissionsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(upgb.fns))
	for _, fn := range upgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*upgb.flds)+len(upgb.fns))
		for _, f := range *upgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*upgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := upgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserPermissionsSelect is the builder for selecting fields of UserPermissions entities.
type UserPermissionsSelect struct {
	*UserPermissionsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ups *UserPermissionsSelect) Aggregate(fns ...AggregateFunc) *UserPermissionsSelect {
	ups.fns = append(ups.fns, fns...)
	return ups
}

// Scan applies the selector query and scans the result into the given value.
func (ups *UserPermissionsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ups.ctx, "Select")
	if err := ups.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserPermissionsQuery, *UserPermissionsSelect](ctx, ups.UserPermissionsQuery, ups, ups.inters, v)
}

func (ups *UserPermissionsSelect) sqlScan(ctx context.Context, root *UserPermissionsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ups.fns))
	for _, fn := range ups.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ups.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ups.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
