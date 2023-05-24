// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrencyfeed"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// FiatCurrencyFeedQuery is the builder for querying FiatCurrencyFeed entities.
type FiatCurrencyFeedQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FiatCurrencyFeed
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FiatCurrencyFeedQuery builder.
func (fcfq *FiatCurrencyFeedQuery) Where(ps ...predicate.FiatCurrencyFeed) *FiatCurrencyFeedQuery {
	fcfq.predicates = append(fcfq.predicates, ps...)
	return fcfq
}

// Limit adds a limit step to the query.
func (fcfq *FiatCurrencyFeedQuery) Limit(limit int) *FiatCurrencyFeedQuery {
	fcfq.limit = &limit
	return fcfq
}

// Offset adds an offset step to the query.
func (fcfq *FiatCurrencyFeedQuery) Offset(offset int) *FiatCurrencyFeedQuery {
	fcfq.offset = &offset
	return fcfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fcfq *FiatCurrencyFeedQuery) Unique(unique bool) *FiatCurrencyFeedQuery {
	fcfq.unique = &unique
	return fcfq
}

// Order adds an order step to the query.
func (fcfq *FiatCurrencyFeedQuery) Order(o ...OrderFunc) *FiatCurrencyFeedQuery {
	fcfq.order = append(fcfq.order, o...)
	return fcfq
}

// First returns the first FiatCurrencyFeed entity from the query.
// Returns a *NotFoundError when no FiatCurrencyFeed was found.
func (fcfq *FiatCurrencyFeedQuery) First(ctx context.Context) (*FiatCurrencyFeed, error) {
	nodes, err := fcfq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fiatcurrencyfeed.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) FirstX(ctx context.Context) *FiatCurrencyFeed {
	node, err := fcfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FiatCurrencyFeed ID from the query.
// Returns a *NotFoundError when no FiatCurrencyFeed ID was found.
func (fcfq *FiatCurrencyFeedQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fcfq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fiatcurrencyfeed.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := fcfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FiatCurrencyFeed entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FiatCurrencyFeed entity is found.
// Returns a *NotFoundError when no FiatCurrencyFeed entities are found.
func (fcfq *FiatCurrencyFeedQuery) Only(ctx context.Context) (*FiatCurrencyFeed, error) {
	nodes, err := fcfq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fiatcurrencyfeed.Label}
	default:
		return nil, &NotSingularError{fiatcurrencyfeed.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) OnlyX(ctx context.Context) *FiatCurrencyFeed {
	node, err := fcfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FiatCurrencyFeed ID in the query.
// Returns a *NotSingularError when more than one FiatCurrencyFeed ID is found.
// Returns a *NotFoundError when no entities are found.
func (fcfq *FiatCurrencyFeedQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = fcfq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fiatcurrencyfeed.Label}
	default:
		err = &NotSingularError{fiatcurrencyfeed.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := fcfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FiatCurrencyFeeds.
func (fcfq *FiatCurrencyFeedQuery) All(ctx context.Context) ([]*FiatCurrencyFeed, error) {
	if err := fcfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fcfq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) AllX(ctx context.Context) []*FiatCurrencyFeed {
	nodes, err := fcfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FiatCurrencyFeed IDs.
func (fcfq *FiatCurrencyFeedQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := fcfq.Select(fiatcurrencyfeed.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := fcfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fcfq *FiatCurrencyFeedQuery) Count(ctx context.Context) (int, error) {
	if err := fcfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fcfq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) CountX(ctx context.Context) int {
	count, err := fcfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fcfq *FiatCurrencyFeedQuery) Exist(ctx context.Context) (bool, error) {
	if err := fcfq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fcfq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fcfq *FiatCurrencyFeedQuery) ExistX(ctx context.Context) bool {
	exist, err := fcfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FiatCurrencyFeedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fcfq *FiatCurrencyFeedQuery) Clone() *FiatCurrencyFeedQuery {
	if fcfq == nil {
		return nil
	}
	return &FiatCurrencyFeedQuery{
		config:     fcfq.config,
		limit:      fcfq.limit,
		offset:     fcfq.offset,
		order:      append([]OrderFunc{}, fcfq.order...),
		predicates: append([]predicate.FiatCurrencyFeed{}, fcfq.predicates...),
		// clone intermediate query.
		sql:    fcfq.sql.Clone(),
		path:   fcfq.path,
		unique: fcfq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FiatCurrencyFeed.Query().
//		GroupBy(fiatcurrencyfeed.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fcfq *FiatCurrencyFeedQuery) GroupBy(field string, fields ...string) *FiatCurrencyFeedGroupBy {
	grbuild := &FiatCurrencyFeedGroupBy{config: fcfq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fcfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fcfq.sqlQuery(ctx), nil
	}
	grbuild.label = fiatcurrencyfeed.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.FiatCurrencyFeed.Query().
//		Select(fiatcurrencyfeed.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (fcfq *FiatCurrencyFeedQuery) Select(fields ...string) *FiatCurrencyFeedSelect {
	fcfq.fields = append(fcfq.fields, fields...)
	selbuild := &FiatCurrencyFeedSelect{FiatCurrencyFeedQuery: fcfq}
	selbuild.label = fiatcurrencyfeed.Label
	selbuild.flds, selbuild.scan = &fcfq.fields, selbuild.Scan
	return selbuild
}

func (fcfq *FiatCurrencyFeedQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fcfq.fields {
		if !fiatcurrencyfeed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fcfq.path != nil {
		prev, err := fcfq.path(ctx)
		if err != nil {
			return err
		}
		fcfq.sql = prev
	}
	if fiatcurrencyfeed.Policy == nil {
		return errors.New("ent: uninitialized fiatcurrencyfeed.Policy (forgotten import ent/runtime?)")
	}
	if err := fiatcurrencyfeed.Policy.EvalQuery(ctx, fcfq); err != nil {
		return err
	}
	return nil
}

func (fcfq *FiatCurrencyFeedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FiatCurrencyFeed, error) {
	var (
		nodes = []*FiatCurrencyFeed{}
		_spec = fcfq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FiatCurrencyFeed).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FiatCurrencyFeed{config: fcfq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(fcfq.modifiers) > 0 {
		_spec.Modifiers = fcfq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fcfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fcfq *FiatCurrencyFeedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fcfq.querySpec()
	if len(fcfq.modifiers) > 0 {
		_spec.Modifiers = fcfq.modifiers
	}
	_spec.Node.Columns = fcfq.fields
	if len(fcfq.fields) > 0 {
		_spec.Unique = fcfq.unique != nil && *fcfq.unique
	}
	return sqlgraph.CountNodes(ctx, fcfq.driver, _spec)
}

func (fcfq *FiatCurrencyFeedQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fcfq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fcfq *FiatCurrencyFeedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fiatcurrencyfeed.Table,
			Columns: fiatcurrencyfeed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fiatcurrencyfeed.FieldID,
			},
		},
		From:   fcfq.sql,
		Unique: true,
	}
	if unique := fcfq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fcfq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fiatcurrencyfeed.FieldID)
		for i := range fields {
			if fields[i] != fiatcurrencyfeed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fcfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fcfq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fcfq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fcfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fcfq *FiatCurrencyFeedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fcfq.driver.Dialect())
	t1 := builder.Table(fiatcurrencyfeed.Table)
	columns := fcfq.fields
	if len(columns) == 0 {
		columns = fiatcurrencyfeed.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fcfq.sql != nil {
		selector = fcfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fcfq.unique != nil && *fcfq.unique {
		selector.Distinct()
	}
	for _, m := range fcfq.modifiers {
		m(selector)
	}
	for _, p := range fcfq.predicates {
		p(selector)
	}
	for _, p := range fcfq.order {
		p(selector)
	}
	if offset := fcfq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fcfq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (fcfq *FiatCurrencyFeedQuery) ForUpdate(opts ...sql.LockOption) *FiatCurrencyFeedQuery {
	if fcfq.driver.Dialect() == dialect.Postgres {
		fcfq.Unique(false)
	}
	fcfq.modifiers = append(fcfq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return fcfq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (fcfq *FiatCurrencyFeedQuery) ForShare(opts ...sql.LockOption) *FiatCurrencyFeedQuery {
	if fcfq.driver.Dialect() == dialect.Postgres {
		fcfq.Unique(false)
	}
	fcfq.modifiers = append(fcfq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return fcfq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fcfq *FiatCurrencyFeedQuery) Modify(modifiers ...func(s *sql.Selector)) *FiatCurrencyFeedSelect {
	fcfq.modifiers = append(fcfq.modifiers, modifiers...)
	return fcfq.Select()
}

// FiatCurrencyFeedGroupBy is the group-by builder for FiatCurrencyFeed entities.
type FiatCurrencyFeedGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fcfgb *FiatCurrencyFeedGroupBy) Aggregate(fns ...AggregateFunc) *FiatCurrencyFeedGroupBy {
	fcfgb.fns = append(fcfgb.fns, fns...)
	return fcfgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fcfgb *FiatCurrencyFeedGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fcfgb.path(ctx)
	if err != nil {
		return err
	}
	fcfgb.sql = query
	return fcfgb.sqlScan(ctx, v)
}

func (fcfgb *FiatCurrencyFeedGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fcfgb.fields {
		if !fiatcurrencyfeed.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fcfgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcfgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fcfgb *FiatCurrencyFeedGroupBy) sqlQuery() *sql.Selector {
	selector := fcfgb.sql.Select()
	aggregation := make([]string, 0, len(fcfgb.fns))
	for _, fn := range fcfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fcfgb.fields)+len(fcfgb.fns))
		for _, f := range fcfgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fcfgb.fields...)...)
}

// FiatCurrencyFeedSelect is the builder for selecting fields of FiatCurrencyFeed entities.
type FiatCurrencyFeedSelect struct {
	*FiatCurrencyFeedQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fcfs *FiatCurrencyFeedSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fcfs.prepareQuery(ctx); err != nil {
		return err
	}
	fcfs.sql = fcfs.FiatCurrencyFeedQuery.sqlQuery(ctx)
	return fcfs.sqlScan(ctx, v)
}

func (fcfs *FiatCurrencyFeedSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fcfs.sql.Query()
	if err := fcfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fcfs *FiatCurrencyFeedSelect) Modify(modifiers ...func(s *sql.Selector)) *FiatCurrencyFeedSelect {
	fcfs.modifiers = append(fcfs.modifiers, modifiers...)
	return fcfs
}
