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
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/currencyhistory"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
)

// CurrencyHistoryQuery is the builder for querying CurrencyHistory entities.
type CurrencyHistoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CurrencyHistory
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CurrencyHistoryQuery builder.
func (chq *CurrencyHistoryQuery) Where(ps ...predicate.CurrencyHistory) *CurrencyHistoryQuery {
	chq.predicates = append(chq.predicates, ps...)
	return chq
}

// Limit adds a limit step to the query.
func (chq *CurrencyHistoryQuery) Limit(limit int) *CurrencyHistoryQuery {
	chq.limit = &limit
	return chq
}

// Offset adds an offset step to the query.
func (chq *CurrencyHistoryQuery) Offset(offset int) *CurrencyHistoryQuery {
	chq.offset = &offset
	return chq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (chq *CurrencyHistoryQuery) Unique(unique bool) *CurrencyHistoryQuery {
	chq.unique = &unique
	return chq
}

// Order adds an order step to the query.
func (chq *CurrencyHistoryQuery) Order(o ...OrderFunc) *CurrencyHistoryQuery {
	chq.order = append(chq.order, o...)
	return chq
}

// First returns the first CurrencyHistory entity from the query.
// Returns a *NotFoundError when no CurrencyHistory was found.
func (chq *CurrencyHistoryQuery) First(ctx context.Context) (*CurrencyHistory, error) {
	nodes, err := chq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{currencyhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (chq *CurrencyHistoryQuery) FirstX(ctx context.Context) *CurrencyHistory {
	node, err := chq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CurrencyHistory ID from the query.
// Returns a *NotFoundError when no CurrencyHistory ID was found.
func (chq *CurrencyHistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = chq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{currencyhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (chq *CurrencyHistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := chq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CurrencyHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CurrencyHistory entity is found.
// Returns a *NotFoundError when no CurrencyHistory entities are found.
func (chq *CurrencyHistoryQuery) Only(ctx context.Context) (*CurrencyHistory, error) {
	nodes, err := chq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{currencyhistory.Label}
	default:
		return nil, &NotSingularError{currencyhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (chq *CurrencyHistoryQuery) OnlyX(ctx context.Context) *CurrencyHistory {
	node, err := chq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CurrencyHistory ID in the query.
// Returns a *NotSingularError when more than one CurrencyHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (chq *CurrencyHistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = chq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{currencyhistory.Label}
	default:
		err = &NotSingularError{currencyhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (chq *CurrencyHistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := chq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CurrencyHistories.
func (chq *CurrencyHistoryQuery) All(ctx context.Context) ([]*CurrencyHistory, error) {
	if err := chq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return chq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (chq *CurrencyHistoryQuery) AllX(ctx context.Context) []*CurrencyHistory {
	nodes, err := chq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CurrencyHistory IDs.
func (chq *CurrencyHistoryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := chq.Select(currencyhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (chq *CurrencyHistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := chq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (chq *CurrencyHistoryQuery) Count(ctx context.Context) (int, error) {
	if err := chq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return chq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (chq *CurrencyHistoryQuery) CountX(ctx context.Context) int {
	count, err := chq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (chq *CurrencyHistoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := chq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return chq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (chq *CurrencyHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := chq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CurrencyHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (chq *CurrencyHistoryQuery) Clone() *CurrencyHistoryQuery {
	if chq == nil {
		return nil
	}
	return &CurrencyHistoryQuery{
		config:     chq.config,
		limit:      chq.limit,
		offset:     chq.offset,
		order:      append([]OrderFunc{}, chq.order...),
		predicates: append([]predicate.CurrencyHistory{}, chq.predicates...),
		// clone intermediate query.
		sql:    chq.sql.Clone(),
		path:   chq.path,
		unique: chq.unique,
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
//	client.CurrencyHistory.Query().
//		GroupBy(currencyhistory.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (chq *CurrencyHistoryQuery) GroupBy(field string, fields ...string) *CurrencyHistoryGroupBy {
	grbuild := &CurrencyHistoryGroupBy{config: chq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := chq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return chq.sqlQuery(ctx), nil
	}
	grbuild.label = currencyhistory.Label
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
//	client.CurrencyHistory.Query().
//		Select(currencyhistory.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (chq *CurrencyHistoryQuery) Select(fields ...string) *CurrencyHistorySelect {
	chq.fields = append(chq.fields, fields...)
	selbuild := &CurrencyHistorySelect{CurrencyHistoryQuery: chq}
	selbuild.label = currencyhistory.Label
	selbuild.flds, selbuild.scan = &chq.fields, selbuild.Scan
	return selbuild
}

func (chq *CurrencyHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range chq.fields {
		if !currencyhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if chq.path != nil {
		prev, err := chq.path(ctx)
		if err != nil {
			return err
		}
		chq.sql = prev
	}
	if currencyhistory.Policy == nil {
		return errors.New("ent: uninitialized currencyhistory.Policy (forgotten import ent/runtime?)")
	}
	if err := currencyhistory.Policy.EvalQuery(ctx, chq); err != nil {
		return err
	}
	return nil
}

func (chq *CurrencyHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CurrencyHistory, error) {
	var (
		nodes = []*CurrencyHistory{}
		_spec = chq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CurrencyHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CurrencyHistory{config: chq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(chq.modifiers) > 0 {
		_spec.Modifiers = chq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, chq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (chq *CurrencyHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := chq.querySpec()
	if len(chq.modifiers) > 0 {
		_spec.Modifiers = chq.modifiers
	}
	_spec.Node.Columns = chq.fields
	if len(chq.fields) > 0 {
		_spec.Unique = chq.unique != nil && *chq.unique
	}
	return sqlgraph.CountNodes(ctx, chq.driver, _spec)
}

func (chq *CurrencyHistoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := chq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (chq *CurrencyHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   currencyhistory.Table,
			Columns: currencyhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: currencyhistory.FieldID,
			},
		},
		From:   chq.sql,
		Unique: true,
	}
	if unique := chq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := chq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, currencyhistory.FieldID)
		for i := range fields {
			if fields[i] != currencyhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := chq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := chq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := chq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := chq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (chq *CurrencyHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(chq.driver.Dialect())
	t1 := builder.Table(currencyhistory.Table)
	columns := chq.fields
	if len(columns) == 0 {
		columns = currencyhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if chq.sql != nil {
		selector = chq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if chq.unique != nil && *chq.unique {
		selector.Distinct()
	}
	for _, m := range chq.modifiers {
		m(selector)
	}
	for _, p := range chq.predicates {
		p(selector)
	}
	for _, p := range chq.order {
		p(selector)
	}
	if offset := chq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := chq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (chq *CurrencyHistoryQuery) ForUpdate(opts ...sql.LockOption) *CurrencyHistoryQuery {
	if chq.driver.Dialect() == dialect.Postgres {
		chq.Unique(false)
	}
	chq.modifiers = append(chq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return chq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (chq *CurrencyHistoryQuery) ForShare(opts ...sql.LockOption) *CurrencyHistoryQuery {
	if chq.driver.Dialect() == dialect.Postgres {
		chq.Unique(false)
	}
	chq.modifiers = append(chq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return chq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (chq *CurrencyHistoryQuery) Modify(modifiers ...func(s *sql.Selector)) *CurrencyHistorySelect {
	chq.modifiers = append(chq.modifiers, modifiers...)
	return chq.Select()
}

// CurrencyHistoryGroupBy is the group-by builder for CurrencyHistory entities.
type CurrencyHistoryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (chgb *CurrencyHistoryGroupBy) Aggregate(fns ...AggregateFunc) *CurrencyHistoryGroupBy {
	chgb.fns = append(chgb.fns, fns...)
	return chgb
}

// Scan applies the group-by query and scans the result into the given value.
func (chgb *CurrencyHistoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := chgb.path(ctx)
	if err != nil {
		return err
	}
	chgb.sql = query
	return chgb.sqlScan(ctx, v)
}

func (chgb *CurrencyHistoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range chgb.fields {
		if !currencyhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := chgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := chgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (chgb *CurrencyHistoryGroupBy) sqlQuery() *sql.Selector {
	selector := chgb.sql.Select()
	aggregation := make([]string, 0, len(chgb.fns))
	for _, fn := range chgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(chgb.fields)+len(chgb.fns))
		for _, f := range chgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(chgb.fields...)...)
}

// CurrencyHistorySelect is the builder for selecting fields of CurrencyHistory entities.
type CurrencyHistorySelect struct {
	*CurrencyHistoryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (chs *CurrencyHistorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := chs.prepareQuery(ctx); err != nil {
		return err
	}
	chs.sql = chs.CurrencyHistoryQuery.sqlQuery(ctx)
	return chs.sqlScan(ctx, v)
}

func (chs *CurrencyHistorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := chs.sql.Query()
	if err := chs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (chs *CurrencyHistorySelect) Modify(modifiers ...func(s *sql.Selector)) *CurrencyHistorySelect {
	chs.modifiers = append(chs.modifiers, modifiers...)
	return chs
}
