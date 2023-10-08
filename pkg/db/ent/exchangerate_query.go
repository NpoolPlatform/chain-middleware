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
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/exchangerate"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
)

// ExchangeRateQuery is the builder for querying ExchangeRate entities.
type ExchangeRateQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ExchangeRate
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExchangeRateQuery builder.
func (erq *ExchangeRateQuery) Where(ps ...predicate.ExchangeRate) *ExchangeRateQuery {
	erq.predicates = append(erq.predicates, ps...)
	return erq
}

// Limit adds a limit step to the query.
func (erq *ExchangeRateQuery) Limit(limit int) *ExchangeRateQuery {
	erq.limit = &limit
	return erq
}

// Offset adds an offset step to the query.
func (erq *ExchangeRateQuery) Offset(offset int) *ExchangeRateQuery {
	erq.offset = &offset
	return erq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (erq *ExchangeRateQuery) Unique(unique bool) *ExchangeRateQuery {
	erq.unique = &unique
	return erq
}

// Order adds an order step to the query.
func (erq *ExchangeRateQuery) Order(o ...OrderFunc) *ExchangeRateQuery {
	erq.order = append(erq.order, o...)
	return erq
}

// First returns the first ExchangeRate entity from the query.
// Returns a *NotFoundError when no ExchangeRate was found.
func (erq *ExchangeRateQuery) First(ctx context.Context) (*ExchangeRate, error) {
	nodes, err := erq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{exchangerate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (erq *ExchangeRateQuery) FirstX(ctx context.Context) *ExchangeRate {
	node, err := erq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExchangeRate ID from the query.
// Returns a *NotFoundError when no ExchangeRate ID was found.
func (erq *ExchangeRateQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = erq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{exchangerate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (erq *ExchangeRateQuery) FirstIDX(ctx context.Context) int {
	id, err := erq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExchangeRate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExchangeRate entity is found.
// Returns a *NotFoundError when no ExchangeRate entities are found.
func (erq *ExchangeRateQuery) Only(ctx context.Context) (*ExchangeRate, error) {
	nodes, err := erq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{exchangerate.Label}
	default:
		return nil, &NotSingularError{exchangerate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (erq *ExchangeRateQuery) OnlyX(ctx context.Context) *ExchangeRate {
	node, err := erq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExchangeRate ID in the query.
// Returns a *NotSingularError when more than one ExchangeRate ID is found.
// Returns a *NotFoundError when no entities are found.
func (erq *ExchangeRateQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = erq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{exchangerate.Label}
	default:
		err = &NotSingularError{exchangerate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (erq *ExchangeRateQuery) OnlyIDX(ctx context.Context) int {
	id, err := erq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExchangeRates.
func (erq *ExchangeRateQuery) All(ctx context.Context) ([]*ExchangeRate, error) {
	if err := erq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return erq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (erq *ExchangeRateQuery) AllX(ctx context.Context) []*ExchangeRate {
	nodes, err := erq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExchangeRate IDs.
func (erq *ExchangeRateQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := erq.Select(exchangerate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (erq *ExchangeRateQuery) IDsX(ctx context.Context) []int {
	ids, err := erq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (erq *ExchangeRateQuery) Count(ctx context.Context) (int, error) {
	if err := erq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return erq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (erq *ExchangeRateQuery) CountX(ctx context.Context) int {
	count, err := erq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (erq *ExchangeRateQuery) Exist(ctx context.Context) (bool, error) {
	if err := erq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return erq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (erq *ExchangeRateQuery) ExistX(ctx context.Context) bool {
	exist, err := erq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExchangeRateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (erq *ExchangeRateQuery) Clone() *ExchangeRateQuery {
	if erq == nil {
		return nil
	}
	return &ExchangeRateQuery{
		config:     erq.config,
		limit:      erq.limit,
		offset:     erq.offset,
		order:      append([]OrderFunc{}, erq.order...),
		predicates: append([]predicate.ExchangeRate{}, erq.predicates...),
		// clone intermediate query.
		sql:    erq.sql.Clone(),
		path:   erq.path,
		unique: erq.unique,
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
//	client.ExchangeRate.Query().
//		GroupBy(exchangerate.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (erq *ExchangeRateQuery) GroupBy(field string, fields ...string) *ExchangeRateGroupBy {
	grbuild := &ExchangeRateGroupBy{config: erq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := erq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return erq.sqlQuery(ctx), nil
	}
	grbuild.label = exchangerate.Label
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
//	client.ExchangeRate.Query().
//		Select(exchangerate.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (erq *ExchangeRateQuery) Select(fields ...string) *ExchangeRateSelect {
	erq.fields = append(erq.fields, fields...)
	selbuild := &ExchangeRateSelect{ExchangeRateQuery: erq}
	selbuild.label = exchangerate.Label
	selbuild.flds, selbuild.scan = &erq.fields, selbuild.Scan
	return selbuild
}

func (erq *ExchangeRateQuery) prepareQuery(ctx context.Context) error {
	for _, f := range erq.fields {
		if !exchangerate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if erq.path != nil {
		prev, err := erq.path(ctx)
		if err != nil {
			return err
		}
		erq.sql = prev
	}
	if exchangerate.Policy == nil {
		return errors.New("ent: uninitialized exchangerate.Policy (forgotten import ent/runtime?)")
	}
	if err := exchangerate.Policy.EvalQuery(ctx, erq); err != nil {
		return err
	}
	return nil
}

func (erq *ExchangeRateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExchangeRate, error) {
	var (
		nodes = []*ExchangeRate{}
		_spec = erq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ExchangeRate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ExchangeRate{config: erq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(erq.modifiers) > 0 {
		_spec.Modifiers = erq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, erq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (erq *ExchangeRateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := erq.querySpec()
	if len(erq.modifiers) > 0 {
		_spec.Modifiers = erq.modifiers
	}
	_spec.Node.Columns = erq.fields
	if len(erq.fields) > 0 {
		_spec.Unique = erq.unique != nil && *erq.unique
	}
	return sqlgraph.CountNodes(ctx, erq.driver, _spec)
}

func (erq *ExchangeRateQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := erq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (erq *ExchangeRateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   exchangerate.Table,
			Columns: exchangerate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: exchangerate.FieldID,
			},
		},
		From:   erq.sql,
		Unique: true,
	}
	if unique := erq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := erq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exchangerate.FieldID)
		for i := range fields {
			if fields[i] != exchangerate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := erq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := erq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := erq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := erq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (erq *ExchangeRateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(erq.driver.Dialect())
	t1 := builder.Table(exchangerate.Table)
	columns := erq.fields
	if len(columns) == 0 {
		columns = exchangerate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if erq.sql != nil {
		selector = erq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if erq.unique != nil && *erq.unique {
		selector.Distinct()
	}
	for _, m := range erq.modifiers {
		m(selector)
	}
	for _, p := range erq.predicates {
		p(selector)
	}
	for _, p := range erq.order {
		p(selector)
	}
	if offset := erq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := erq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (erq *ExchangeRateQuery) ForUpdate(opts ...sql.LockOption) *ExchangeRateQuery {
	if erq.driver.Dialect() == dialect.Postgres {
		erq.Unique(false)
	}
	erq.modifiers = append(erq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return erq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (erq *ExchangeRateQuery) ForShare(opts ...sql.LockOption) *ExchangeRateQuery {
	if erq.driver.Dialect() == dialect.Postgres {
		erq.Unique(false)
	}
	erq.modifiers = append(erq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return erq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (erq *ExchangeRateQuery) Modify(modifiers ...func(s *sql.Selector)) *ExchangeRateSelect {
	erq.modifiers = append(erq.modifiers, modifiers...)
	return erq.Select()
}

// ExchangeRateGroupBy is the group-by builder for ExchangeRate entities.
type ExchangeRateGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ergb *ExchangeRateGroupBy) Aggregate(fns ...AggregateFunc) *ExchangeRateGroupBy {
	ergb.fns = append(ergb.fns, fns...)
	return ergb
}

// Scan applies the group-by query and scans the result into the given value.
func (ergb *ExchangeRateGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ergb.path(ctx)
	if err != nil {
		return err
	}
	ergb.sql = query
	return ergb.sqlScan(ctx, v)
}

func (ergb *ExchangeRateGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ergb.fields {
		if !exchangerate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ergb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ergb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ergb *ExchangeRateGroupBy) sqlQuery() *sql.Selector {
	selector := ergb.sql.Select()
	aggregation := make([]string, 0, len(ergb.fns))
	for _, fn := range ergb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ergb.fields)+len(ergb.fns))
		for _, f := range ergb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ergb.fields...)...)
}

// ExchangeRateSelect is the builder for selecting fields of ExchangeRate entities.
type ExchangeRateSelect struct {
	*ExchangeRateQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ers *ExchangeRateSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ers.prepareQuery(ctx); err != nil {
		return err
	}
	ers.sql = ers.ExchangeRateQuery.sqlQuery(ctx)
	return ers.sqlScan(ctx, v)
}

func (ers *ExchangeRateSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ers.sql.Query()
	if err := ers.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ers *ExchangeRateSelect) Modify(modifiers ...func(s *sql.Selector)) *ExchangeRateSelect {
	ers.modifiers = append(ers.modifiers, modifiers...)
	return ers
}
