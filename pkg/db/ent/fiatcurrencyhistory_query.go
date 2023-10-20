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
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrencyhistory"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
)

// FiatCurrencyHistoryQuery is the builder for querying FiatCurrencyHistory entities.
type FiatCurrencyHistoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FiatCurrencyHistory
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FiatCurrencyHistoryQuery builder.
func (fchq *FiatCurrencyHistoryQuery) Where(ps ...predicate.FiatCurrencyHistory) *FiatCurrencyHistoryQuery {
	fchq.predicates = append(fchq.predicates, ps...)
	return fchq
}

// Limit adds a limit step to the query.
func (fchq *FiatCurrencyHistoryQuery) Limit(limit int) *FiatCurrencyHistoryQuery {
	fchq.limit = &limit
	return fchq
}

// Offset adds an offset step to the query.
func (fchq *FiatCurrencyHistoryQuery) Offset(offset int) *FiatCurrencyHistoryQuery {
	fchq.offset = &offset
	return fchq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fchq *FiatCurrencyHistoryQuery) Unique(unique bool) *FiatCurrencyHistoryQuery {
	fchq.unique = &unique
	return fchq
}

// Order adds an order step to the query.
func (fchq *FiatCurrencyHistoryQuery) Order(o ...OrderFunc) *FiatCurrencyHistoryQuery {
	fchq.order = append(fchq.order, o...)
	return fchq
}

// First returns the first FiatCurrencyHistory entity from the query.
// Returns a *NotFoundError when no FiatCurrencyHistory was found.
func (fchq *FiatCurrencyHistoryQuery) First(ctx context.Context) (*FiatCurrencyHistory, error) {
	nodes, err := fchq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fiatcurrencyhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fchq *FiatCurrencyHistoryQuery) FirstX(ctx context.Context) *FiatCurrencyHistory {
	node, err := fchq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FiatCurrencyHistory ID from the query.
// Returns a *NotFoundError when no FiatCurrencyHistory ID was found.
func (fchq *FiatCurrencyHistoryQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fchq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fiatcurrencyhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fchq *FiatCurrencyHistoryQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := fchq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FiatCurrencyHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FiatCurrencyHistory entity is found.
// Returns a *NotFoundError when no FiatCurrencyHistory entities are found.
func (fchq *FiatCurrencyHistoryQuery) Only(ctx context.Context) (*FiatCurrencyHistory, error) {
	nodes, err := fchq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fiatcurrencyhistory.Label}
	default:
		return nil, &NotSingularError{fiatcurrencyhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fchq *FiatCurrencyHistoryQuery) OnlyX(ctx context.Context) *FiatCurrencyHistory {
	node, err := fchq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FiatCurrencyHistory ID in the query.
// Returns a *NotSingularError when more than one FiatCurrencyHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (fchq *FiatCurrencyHistoryQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = fchq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fiatcurrencyhistory.Label}
	default:
		err = &NotSingularError{fiatcurrencyhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fchq *FiatCurrencyHistoryQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := fchq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FiatCurrencyHistories.
func (fchq *FiatCurrencyHistoryQuery) All(ctx context.Context) ([]*FiatCurrencyHistory, error) {
	if err := fchq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fchq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fchq *FiatCurrencyHistoryQuery) AllX(ctx context.Context) []*FiatCurrencyHistory {
	nodes, err := fchq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FiatCurrencyHistory IDs.
func (fchq *FiatCurrencyHistoryQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := fchq.Select(fiatcurrencyhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fchq *FiatCurrencyHistoryQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := fchq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fchq *FiatCurrencyHistoryQuery) Count(ctx context.Context) (int, error) {
	if err := fchq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fchq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fchq *FiatCurrencyHistoryQuery) CountX(ctx context.Context) int {
	count, err := fchq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fchq *FiatCurrencyHistoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := fchq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fchq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fchq *FiatCurrencyHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := fchq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FiatCurrencyHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fchq *FiatCurrencyHistoryQuery) Clone() *FiatCurrencyHistoryQuery {
	if fchq == nil {
		return nil
	}
	return &FiatCurrencyHistoryQuery{
		config:     fchq.config,
		limit:      fchq.limit,
		offset:     fchq.offset,
		order:      append([]OrderFunc{}, fchq.order...),
		predicates: append([]predicate.FiatCurrencyHistory{}, fchq.predicates...),
		// clone intermediate query.
		sql:    fchq.sql.Clone(),
		path:   fchq.path,
		unique: fchq.unique,
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
//	client.FiatCurrencyHistory.Query().
//		GroupBy(fiatcurrencyhistory.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fchq *FiatCurrencyHistoryQuery) GroupBy(field string, fields ...string) *FiatCurrencyHistoryGroupBy {
	grbuild := &FiatCurrencyHistoryGroupBy{config: fchq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fchq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fchq.sqlQuery(ctx), nil
	}
	grbuild.label = fiatcurrencyhistory.Label
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
//	client.FiatCurrencyHistory.Query().
//		Select(fiatcurrencyhistory.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (fchq *FiatCurrencyHistoryQuery) Select(fields ...string) *FiatCurrencyHistorySelect {
	fchq.fields = append(fchq.fields, fields...)
	selbuild := &FiatCurrencyHistorySelect{FiatCurrencyHistoryQuery: fchq}
	selbuild.label = fiatcurrencyhistory.Label
	selbuild.flds, selbuild.scan = &fchq.fields, selbuild.Scan
	return selbuild
}

func (fchq *FiatCurrencyHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fchq.fields {
		if !fiatcurrencyhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fchq.path != nil {
		prev, err := fchq.path(ctx)
		if err != nil {
			return err
		}
		fchq.sql = prev
	}
	if fiatcurrencyhistory.Policy == nil {
		return errors.New("ent: uninitialized fiatcurrencyhistory.Policy (forgotten import ent/runtime?)")
	}
	if err := fiatcurrencyhistory.Policy.EvalQuery(ctx, fchq); err != nil {
		return err
	}
	return nil
}

func (fchq *FiatCurrencyHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FiatCurrencyHistory, error) {
	var (
		nodes = []*FiatCurrencyHistory{}
		_spec = fchq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*FiatCurrencyHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &FiatCurrencyHistory{config: fchq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(fchq.modifiers) > 0 {
		_spec.Modifiers = fchq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fchq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fchq *FiatCurrencyHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fchq.querySpec()
	if len(fchq.modifiers) > 0 {
		_spec.Modifiers = fchq.modifiers
	}
	_spec.Node.Columns = fchq.fields
	if len(fchq.fields) > 0 {
		_spec.Unique = fchq.unique != nil && *fchq.unique
	}
	return sqlgraph.CountNodes(ctx, fchq.driver, _spec)
}

func (fchq *FiatCurrencyHistoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fchq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fchq *FiatCurrencyHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fiatcurrencyhistory.Table,
			Columns: fiatcurrencyhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fiatcurrencyhistory.FieldID,
			},
		},
		From:   fchq.sql,
		Unique: true,
	}
	if unique := fchq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fchq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fiatcurrencyhistory.FieldID)
		for i := range fields {
			if fields[i] != fiatcurrencyhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fchq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fchq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fchq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fchq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fchq *FiatCurrencyHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fchq.driver.Dialect())
	t1 := builder.Table(fiatcurrencyhistory.Table)
	columns := fchq.fields
	if len(columns) == 0 {
		columns = fiatcurrencyhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fchq.sql != nil {
		selector = fchq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fchq.unique != nil && *fchq.unique {
		selector.Distinct()
	}
	for _, m := range fchq.modifiers {
		m(selector)
	}
	for _, p := range fchq.predicates {
		p(selector)
	}
	for _, p := range fchq.order {
		p(selector)
	}
	if offset := fchq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fchq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (fchq *FiatCurrencyHistoryQuery) ForUpdate(opts ...sql.LockOption) *FiatCurrencyHistoryQuery {
	if fchq.driver.Dialect() == dialect.Postgres {
		fchq.Unique(false)
	}
	fchq.modifiers = append(fchq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return fchq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (fchq *FiatCurrencyHistoryQuery) ForShare(opts ...sql.LockOption) *FiatCurrencyHistoryQuery {
	if fchq.driver.Dialect() == dialect.Postgres {
		fchq.Unique(false)
	}
	fchq.modifiers = append(fchq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return fchq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fchq *FiatCurrencyHistoryQuery) Modify(modifiers ...func(s *sql.Selector)) *FiatCurrencyHistorySelect {
	fchq.modifiers = append(fchq.modifiers, modifiers...)
	return fchq.Select()
}

// FiatCurrencyHistoryGroupBy is the group-by builder for FiatCurrencyHistory entities.
type FiatCurrencyHistoryGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fchgb *FiatCurrencyHistoryGroupBy) Aggregate(fns ...AggregateFunc) *FiatCurrencyHistoryGroupBy {
	fchgb.fns = append(fchgb.fns, fns...)
	return fchgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fchgb *FiatCurrencyHistoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fchgb.path(ctx)
	if err != nil {
		return err
	}
	fchgb.sql = query
	return fchgb.sqlScan(ctx, v)
}

func (fchgb *FiatCurrencyHistoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fchgb.fields {
		if !fiatcurrencyhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fchgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fchgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fchgb *FiatCurrencyHistoryGroupBy) sqlQuery() *sql.Selector {
	selector := fchgb.sql.Select()
	aggregation := make([]string, 0, len(fchgb.fns))
	for _, fn := range fchgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fchgb.fields)+len(fchgb.fns))
		for _, f := range fchgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fchgb.fields...)...)
}

// FiatCurrencyHistorySelect is the builder for selecting fields of FiatCurrencyHistory entities.
type FiatCurrencyHistorySelect struct {
	*FiatCurrencyHistoryQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fchs *FiatCurrencyHistorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := fchs.prepareQuery(ctx); err != nil {
		return err
	}
	fchs.sql = fchs.FiatCurrencyHistoryQuery.sqlQuery(ctx)
	return fchs.sqlScan(ctx, v)
}

func (fchs *FiatCurrencyHistorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fchs.sql.Query()
	if err := fchs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (fchs *FiatCurrencyHistorySelect) Modify(modifiers ...func(s *sql.Selector)) *FiatCurrencyHistorySelect {
	fchs.modifiers = append(fchs.modifiers, modifiers...)
	return fchs
}
