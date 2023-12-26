// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/identifier"
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/schema/ksuid"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IdentifierQuery is the builder for querying Identifier entities.
type IdentifierQuery struct {
	config
	ctx        *QueryContext
	order      []identifier.OrderOption
	inters     []Interceptor
	predicates []predicate.Identifier
	withBook   *BookQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Identifier) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IdentifierQuery builder.
func (iq *IdentifierQuery) Where(ps ...predicate.Identifier) *IdentifierQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *IdentifierQuery) Limit(limit int) *IdentifierQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *IdentifierQuery) Offset(offset int) *IdentifierQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *IdentifierQuery) Unique(unique bool) *IdentifierQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *IdentifierQuery) Order(o ...identifier.OrderOption) *IdentifierQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryBook chains the current query on the "book" edge.
func (iq *IdentifierQuery) QueryBook() *BookQuery {
	query := (&BookClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(identifier.Table, identifier.FieldID, selector),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, identifier.BookTable, identifier.BookColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Identifier entity from the query.
// Returns a *NotFoundError when no Identifier was found.
func (iq *IdentifierQuery) First(ctx context.Context) (*Identifier, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{identifier.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *IdentifierQuery) FirstX(ctx context.Context) *Identifier {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Identifier ID from the query.
// Returns a *NotFoundError when no Identifier ID was found.
func (iq *IdentifierQuery) FirstID(ctx context.Context) (id ksuid.ID, err error) {
	var ids []ksuid.ID
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{identifier.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *IdentifierQuery) FirstIDX(ctx context.Context) ksuid.ID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Identifier entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Identifier entity is found.
// Returns a *NotFoundError when no Identifier entities are found.
func (iq *IdentifierQuery) Only(ctx context.Context) (*Identifier, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{identifier.Label}
	default:
		return nil, &NotSingularError{identifier.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *IdentifierQuery) OnlyX(ctx context.Context) *Identifier {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Identifier ID in the query.
// Returns a *NotSingularError when more than one Identifier ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *IdentifierQuery) OnlyID(ctx context.Context) (id ksuid.ID, err error) {
	var ids []ksuid.ID
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{identifier.Label}
	default:
		err = &NotSingularError{identifier.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *IdentifierQuery) OnlyIDX(ctx context.Context) ksuid.ID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Identifiers.
func (iq *IdentifierQuery) All(ctx context.Context) ([]*Identifier, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Identifier, *IdentifierQuery]()
	return withInterceptors[[]*Identifier](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *IdentifierQuery) AllX(ctx context.Context) []*Identifier {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Identifier IDs.
func (iq *IdentifierQuery) IDs(ctx context.Context) (ids []ksuid.ID, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(identifier.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *IdentifierQuery) IDsX(ctx context.Context) []ksuid.ID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *IdentifierQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*IdentifierQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *IdentifierQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *IdentifierQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *IdentifierQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IdentifierQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *IdentifierQuery) Clone() *IdentifierQuery {
	if iq == nil {
		return nil
	}
	return &IdentifierQuery{
		config:     iq.config,
		ctx:        iq.ctx.Clone(),
		order:      append([]identifier.OrderOption{}, iq.order...),
		inters:     append([]Interceptor{}, iq.inters...),
		predicates: append([]predicate.Identifier{}, iq.predicates...),
		withBook:   iq.withBook.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithBook tells the query-builder to eager-load the nodes that are connected to
// the "book" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *IdentifierQuery) WithBook(opts ...func(*BookQuery)) *IdentifierQuery {
	query := (&BookClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withBook = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CalibreID int64 `json:"calibre_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Identifier.Query().
//		GroupBy(identifier.FieldCalibreID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *IdentifierQuery) GroupBy(field string, fields ...string) *IdentifierGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IdentifierGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = identifier.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CalibreID int64 `json:"calibre_id,omitempty"`
//	}
//
//	client.Identifier.Query().
//		Select(identifier.FieldCalibreID).
//		Scan(ctx, &v)
func (iq *IdentifierQuery) Select(fields ...string) *IdentifierSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &IdentifierSelect{IdentifierQuery: iq}
	sbuild.label = identifier.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IdentifierSelect configured with the given aggregations.
func (iq *IdentifierQuery) Aggregate(fns ...AggregateFunc) *IdentifierSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *IdentifierQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !identifier.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	if identifier.Policy == nil {
		return errors.New("ent: uninitialized identifier.Policy (forgotten import ent/runtime?)")
	}
	if err := identifier.Policy.EvalQuery(ctx, iq); err != nil {
		return err
	}
	return nil
}

func (iq *IdentifierQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Identifier, error) {
	var (
		nodes       = []*Identifier{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [1]bool{
			iq.withBook != nil,
		}
	)
	if iq.withBook != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, identifier.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Identifier).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Identifier{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(iq.modifiers) > 0 {
		_spec.Modifiers = iq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withBook; query != nil {
		if err := iq.loadBook(ctx, query, nodes, nil,
			func(n *Identifier, e *Book) { n.Edges.Book = e }); err != nil {
			return nil, err
		}
	}
	for i := range iq.loadTotal {
		if err := iq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *IdentifierQuery) loadBook(ctx context.Context, query *BookQuery, nodes []*Identifier, init func(*Identifier), assign func(*Identifier, *Book)) error {
	ids := make([]ksuid.ID, 0, len(nodes))
	nodeids := make(map[ksuid.ID][]*Identifier)
	for i := range nodes {
		if nodes[i].identifier_book == nil {
			continue
		}
		fk := *nodes[i].identifier_book
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(book.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "identifier_book" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (iq *IdentifierQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	if len(iq.modifiers) > 0 {
		_spec.Modifiers = iq.modifiers
	}
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *IdentifierQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(identifier.Table, identifier.Columns, sqlgraph.NewFieldSpec(identifier.FieldID, field.TypeString))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, identifier.FieldID)
		for i := range fields {
			if fields[i] != identifier.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *IdentifierQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(identifier.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = identifier.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IdentifierGroupBy is the group-by builder for Identifier entities.
type IdentifierGroupBy struct {
	selector
	build *IdentifierQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *IdentifierGroupBy) Aggregate(fns ...AggregateFunc) *IdentifierGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *IdentifierGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IdentifierQuery, *IdentifierGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *IdentifierGroupBy) sqlScan(ctx context.Context, root *IdentifierQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IdentifierSelect is the builder for selecting fields of Identifier entities.
type IdentifierSelect struct {
	*IdentifierQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *IdentifierSelect) Aggregate(fns ...AggregateFunc) *IdentifierSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *IdentifierSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IdentifierQuery, *IdentifierSelect](ctx, is.IdentifierQuery, is, is.inters, v)
}

func (is *IdentifierSelect) sqlScan(ctx context.Context, root *IdentifierQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
