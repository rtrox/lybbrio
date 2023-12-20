// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"lybbrio/internal/ent/book"
	"lybbrio/internal/ent/predicate"
	"lybbrio/internal/ent/publisher"
	"lybbrio/internal/ent/schema/ksuid"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PublisherQuery is the builder for querying Publisher entities.
type PublisherQuery struct {
	config
	ctx            *QueryContext
	order          []publisher.OrderOption
	inters         []Interceptor
	predicates     []predicate.Publisher
	withBooks      *BookQuery
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*Publisher) error
	withNamedBooks map[string]*BookQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PublisherQuery builder.
func (pq *PublisherQuery) Where(ps ...predicate.Publisher) *PublisherQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PublisherQuery) Limit(limit int) *PublisherQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PublisherQuery) Offset(offset int) *PublisherQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PublisherQuery) Unique(unique bool) *PublisherQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PublisherQuery) Order(o ...publisher.OrderOption) *PublisherQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryBooks chains the current query on the "books" edge.
func (pq *PublisherQuery) QueryBooks() *BookQuery {
	query := (&BookClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(publisher.Table, publisher.FieldID, selector),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, publisher.BooksTable, publisher.BooksColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Publisher entity from the query.
// Returns a *NotFoundError when no Publisher was found.
func (pq *PublisherQuery) First(ctx context.Context) (*Publisher, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{publisher.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PublisherQuery) FirstX(ctx context.Context) *Publisher {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Publisher ID from the query.
// Returns a *NotFoundError when no Publisher ID was found.
func (pq *PublisherQuery) FirstID(ctx context.Context) (id ksuid.ID, err error) {
	var ids []ksuid.ID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{publisher.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PublisherQuery) FirstIDX(ctx context.Context) ksuid.ID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Publisher entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Publisher entity is found.
// Returns a *NotFoundError when no Publisher entities are found.
func (pq *PublisherQuery) Only(ctx context.Context) (*Publisher, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{publisher.Label}
	default:
		return nil, &NotSingularError{publisher.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PublisherQuery) OnlyX(ctx context.Context) *Publisher {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Publisher ID in the query.
// Returns a *NotSingularError when more than one Publisher ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PublisherQuery) OnlyID(ctx context.Context) (id ksuid.ID, err error) {
	var ids []ksuid.ID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{publisher.Label}
	default:
		err = &NotSingularError{publisher.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PublisherQuery) OnlyIDX(ctx context.Context) ksuid.ID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Publishers.
func (pq *PublisherQuery) All(ctx context.Context) ([]*Publisher, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Publisher, *PublisherQuery]()
	return withInterceptors[[]*Publisher](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PublisherQuery) AllX(ctx context.Context) []*Publisher {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Publisher IDs.
func (pq *PublisherQuery) IDs(ctx context.Context) (ids []ksuid.ID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(publisher.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PublisherQuery) IDsX(ctx context.Context) []ksuid.ID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PublisherQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PublisherQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PublisherQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PublisherQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PublisherQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PublisherQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PublisherQuery) Clone() *PublisherQuery {
	if pq == nil {
		return nil
	}
	return &PublisherQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]publisher.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Publisher{}, pq.predicates...),
		withBooks:  pq.withBooks.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithBooks tells the query-builder to eager-load the nodes that are connected to
// the "books" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PublisherQuery) WithBooks(opts ...func(*BookQuery)) *PublisherQuery {
	query := (&BookClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withBooks = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Publisher.Query().
//		GroupBy(publisher.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PublisherQuery) GroupBy(field string, fields ...string) *PublisherGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PublisherGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = publisher.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Publisher.Query().
//		Select(publisher.FieldName).
//		Scan(ctx, &v)
func (pq *PublisherQuery) Select(fields ...string) *PublisherSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PublisherSelect{PublisherQuery: pq}
	sbuild.label = publisher.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PublisherSelect configured with the given aggregations.
func (pq *PublisherQuery) Aggregate(fns ...AggregateFunc) *PublisherSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PublisherQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !publisher.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	if publisher.Policy == nil {
		return errors.New("ent: uninitialized publisher.Policy (forgotten import ent/runtime?)")
	}
	if err := publisher.Policy.EvalQuery(ctx, pq); err != nil {
		return err
	}
	return nil
}

func (pq *PublisherQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Publisher, error) {
	var (
		nodes       = []*Publisher{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withBooks != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Publisher).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Publisher{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withBooks; query != nil {
		if err := pq.loadBooks(ctx, query, nodes,
			func(n *Publisher) { n.Edges.Books = []*Book{} },
			func(n *Publisher, e *Book) { n.Edges.Books = append(n.Edges.Books, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedBooks {
		if err := pq.loadBooks(ctx, query, nodes,
			func(n *Publisher) { n.appendNamedBooks(name) },
			func(n *Publisher, e *Book) { n.appendNamedBooks(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PublisherQuery) loadBooks(ctx context.Context, query *BookQuery, nodes []*Publisher, init func(*Publisher), assign func(*Publisher, *Book)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[ksuid.ID]*Publisher)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Book(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(publisher.BooksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.publisher_books
		if fk == nil {
			return fmt.Errorf(`foreign-key "publisher_books" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "publisher_books" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PublisherQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PublisherQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(publisher.Table, publisher.Columns, sqlgraph.NewFieldSpec(publisher.FieldID, field.TypeString))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, publisher.FieldID)
		for i := range fields {
			if fields[i] != publisher.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PublisherQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(publisher.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = publisher.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedBooks tells the query-builder to eager-load the nodes that are connected to the "books"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *PublisherQuery) WithNamedBooks(name string, opts ...func(*BookQuery)) *PublisherQuery {
	query := (&BookClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedBooks == nil {
		pq.withNamedBooks = make(map[string]*BookQuery)
	}
	pq.withNamedBooks[name] = query
	return pq
}

// PublisherGroupBy is the group-by builder for Publisher entities.
type PublisherGroupBy struct {
	selector
	build *PublisherQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PublisherGroupBy) Aggregate(fns ...AggregateFunc) *PublisherGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PublisherGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PublisherQuery, *PublisherGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PublisherGroupBy) sqlScan(ctx context.Context, root *PublisherQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PublisherSelect is the builder for selecting fields of Publisher entities.
type PublisherSelect struct {
	*PublisherQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PublisherSelect) Aggregate(fns ...AggregateFunc) *PublisherSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PublisherSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PublisherQuery, *PublisherSelect](ctx, ps.PublisherQuery, ps, ps.inters, v)
}

func (ps *PublisherSelect) sqlScan(ctx context.Context, root *PublisherQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
