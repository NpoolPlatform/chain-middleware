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
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrency"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
)

// FiatCurrencyQuery is the builder for querying FiatCurrency entities.
type FiatCurrencyQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FiatCurrency
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FiatCurrencyQuery builder.
func (fcq *FiatCurrencyQuery) Where(ps ...predicate.FiatCurrency) *FiatCurrencyQuery {
	fcq.predicates = append(fcq.predicates, ps...)
	return fcq
}

// Limit adds a limit step to the query.
func (fcq *FiatCurrencyQuery) Limit(limit int) *FiatCurrencyQuery {
	fcq.limit = &limit
	return fcq
}

// Offset adds an offset step to the query.
func (fcq *FiatCurrencyQuery) Offset(offset int) *FiatCurrencyQuery {
	fcq.offset = &offset
	return fcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fcq *FiatCurrencyQuery) Unique(unique bool) *FiatCurrencyQuery {
	fcq.unique = &unique
	return fcq
}

// Order adds an order step to the query.
func (fcq *FiatCurrencyQuery) Order(o ...OrderFunc) *FiatCurrencyQuery {
	fcq.order = append(fcq.order, o...)
	return fcq
}

// First returns the first FiatCurrency entity from the query.
// Returns a *NotFoundError when no FiatCurrency was found.
func (fcq *FiatCurrencyQuery) First(ctx context.Context) (*FiatCurrency, error) {
	nodes, err := fcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fiatcurrency.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fcq *FiatCurrencyQuery) FirstX(ctx context.Context) *FiatCurrency {
	node, err := fcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FiatCurrency ID from the query.
// Returns a *NotFoundError when no FiatCurrency ID was found.
func (fcq *FiatCurrencyQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fiatcurrency.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fcq *FiatCurrencyQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := fcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FiatCurrency entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FiatCurrency entity is found.
// Returns a *NotFoundError when no FiatCurrency entities are found.
func (fcq *FiatCurrencyQuery) Only(ctx context.Context) (*FiatCurrency, error) {
	nodes, err := fcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fiatcurrency.Label}
	default:
		return nil, &NotSingularError{fiatcurrency.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fcq *FiatCurrencyQuery) OnlyX(ctx context.Context) *FiatCurrency {
	node, err := fcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FiatCurrency ID in the query.
// Returns a *NotSingularError when more than one FiatCurrency ID is found.
// Returns a *NotFoundError when no entities are found.
func (fcq *FiatCurrencyQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fiatcurrency.Label}
	default:
		err = &NotSingularError{fiatcurrency.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fcq *FiatCurrencyQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := fcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FiatCurrencies.
func (fcq *FiatCurrencyQuery) All(ctx context.Context) ([]*FiatCurrency, error) {
	if err := fcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fcq *FiatCurrencyQuery) AllX(ctx context.Context) []*FiatCurrency {
	nodes, err := fcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FiatCurrency IDs.
func (fcq *FiatCurrencyQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := fcq.Select(fiatcurrency.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fcq *FiatCurrencyQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := fcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fcq *FiatCurrencyQuery) Count(ctx context.Context) (int, error) {
	if err := fcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fcq *FiatCurrencyQuery) CountX(ctx context.Context) int {
	count, err := fcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fcq *FiatCurrencyQuery) Exist(ctx context.Context) (bool, error) {
	if err := fcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fcq *FiatCurrencyQuery) ExistX(ctx context.Context) bool {
	exist, err := fcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FiatCurrencyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fcq *FiatCurrencyQuery) Clone() *FiatCurrencyQuery {
	if fcq == nil {
		return nil
	}
	return &FiatCurrencyQuery{
		config:     fcq.config,
		limit:      fcq.limit,
		offset:     fcq.offset,
		order:      append([]OrderFunc{}, fcq.order...),
		predicates: append([]predicate.FiatCurrency{}, fcq.predicates...),
		// clone intermediate query.
		sql:    fcq.sql.Clone(),
		path:   fcq.path,
		unique: fcq.unique,
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
//	client.FiatCurrency.Query().
//		GroupBy(fiatcurrency.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fcq *FiatCurrencyQuery) GroupBy(field string, fields ...string) *FiatCurrencyGroupBy {
	grbuild := &FiatCurrencyGroupBy{config: fcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fcq.sqlQuery(ctx), nil
	}
	grbuild.label = fiatcurrency.Label
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
//	client.FiatCurrency.Query().
//		Select(fiatcurrency.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (fcq *FiatCurrencyQuery) Select(fields ...string) *FiatCurrencySelect {
	fcq.fields = append(fcq.fields, fields...)
	selbuild := &FiatCurrencySelect{FiatCurrencyQuery: fcq}
	selbuild.label = fiatcurrency.Label
	selbuild.flds, selbuild.scan = &fcq.fields, selbuild.Scan
	return selbuild
}

func (fcq *FiatCurrencyQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fcq.fields {
		if !fiatcurrency.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fcq.path != nil {
		prev, err := fcq.path(ctx)
		if err != nil {
			return err
		}
		fcq.sql = prev
	}
	if fiatcurrency.Policy == nil {
		return errors.New("ent: uninitialized fiatcurrency.Policy (forgotten import ent/runtime?)")
	}
	if err := fiatcurrency.Policy.EvalQuery(ctx, fcq); err != nil {
		return err
	}
	return nil
}

func (fcq *FiatCurrencyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FiatCurrency, error) {
	var (
		nodes = []*FiatCurrency{}
		_spec = fcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FiatCurrency).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FiatCurrency{config: fcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(fcq.modifiers) > 0 {
		_spec.Modifiers = fcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fcq *FiatCurrencyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fcq.querySpec()
	if len(fcq.modifiers) > 0 {
		_spec.Modifiers = fcq.modifiers
	}
	_spec.Node.Columns = fcq.fields
	if len(fcq.fields) > 0 {
		_spec.Unique = fcq.unique != nil && *fcq.unique
	}
	return sqlgraph.CountNodes(ctx, fcq.driver, _spec)
}

func (fcq *FiatCurrencyQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fcq *FiatCurrencyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fiatcurrency.Table,
			Columns: fiatcurrency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fiatcurrency.FieldID,
			},
		},
		From:   fcq.sql,
		Unique: true,
	}
	if unique := fcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fiatcurrency.FieldID)
		for i := range fields {
			if fields[i] != fiatcurrency.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fcq *FiatCurrencyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fcq.driver.Dialect())
	t1 := builder.Table(fiatcurrency.Table)
	columns := fcq.fields
	if len(columns) == 0 {
		columns = fiatcurrency.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fcq.sql != nil {
		selector = fcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fcq.unique != nil && *fcq.unique {
		selector.Distinct()
	}
	for _, m := range fcq.modifiers {
		m(selector)
	}
	for _, p := range fcq.predicates {
		p(selector)
	}
	for _, p := range fcq.order {
		p(selector)
	}
	if offset := fcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (fcq *FiatCurrencyQuery) ForUpdate(opts ...sql.LockOption) *FiatCurrencyQuery {
	if fcq.driver.Dialect() == dialect.Postgres {
		fcq.Unique(false)
	}
	fcq.modifiers = append(fcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return fcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (fcq *FiatCurrencyQuery) ForShare(opts ...sql.LockOption) *FiatCurrencyQuery {
	if fcq.driver.Dialect() == dialect.Postgres {
		fcq.Unique(false)
	}
	fcq.modifiers = append(fcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return fcq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fcq *FiatCurrencyQuery) Modify(modifiers ...func(s *sql.Selector)) *FiatCurrencySelect {
	fcq.modifiers = append(fcq.modifiers, modifiers...)
	return fcq.Select()
}

// FiatCurrencyGroupBy is the group-by builder for FiatCurrency entities.
type FiatCurrencyGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fcgb *FiatCurrencyGroupBy) Aggregate(fns ...AggregateFunc) *FiatCurrencyGroupBy {
	fcgb.fns = append(fcgb.fns, fns...)
	return fcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fcgb *FiatCurrencyGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fcgb.path(ctx)
	if err != nil {
		return err
	}
	fcgb.sql = query
	return fcgb.sqlScan(ctx, v)
}

func (fcgb *FiatCurrencyGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fcgb.fields {
		if !fiatcurrency.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fcgb *FiatCurrencyGroupBy) sqlQuery() *sql.Selector {
	selector := fcgb.sql.Select()
	aggregation := make([]string, 0, len(fcgb.fns))
	for _, fn := range fcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fcgb.fields)+len(fcgb.fns))
		for _, f := range fcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fcgb.fields...)...)
}

// FiatCurrencySelect is the builder for selecting fields of FiatCurrency entities.
type FiatCurrencySelect struct {
	*FiatCurrencyQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fcs *FiatCurrencySelect) Scan(ctx context.Context, v interface{}) error {
	if err := fcs.prepareQuery(ctx); err != nil {
		return err
	}
	fcs.sql = fcs.FiatCurrencyQuery.sqlQuery(ctx)
	return fcs.sqlScan(ctx, v)
}

func (fcs *FiatCurrencySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fcs.sql.Query()
	if err := fcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fcs *FiatCurrencySelect) Modify(modifiers ...func(s *sql.Selector)) *FiatCurrencySelect {
	fcs.modifiers = append(fcs.modifiers, modifiers...)
	return fcs
}
