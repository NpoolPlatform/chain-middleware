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
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coindescription"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinDescriptionQuery is the builder for querying CoinDescription entities.
type CoinDescriptionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CoinDescription
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CoinDescriptionQuery builder.
func (cdq *CoinDescriptionQuery) Where(ps ...predicate.CoinDescription) *CoinDescriptionQuery {
	cdq.predicates = append(cdq.predicates, ps...)
	return cdq
}

// Limit adds a limit step to the query.
func (cdq *CoinDescriptionQuery) Limit(limit int) *CoinDescriptionQuery {
	cdq.limit = &limit
	return cdq
}

// Offset adds an offset step to the query.
func (cdq *CoinDescriptionQuery) Offset(offset int) *CoinDescriptionQuery {
	cdq.offset = &offset
	return cdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cdq *CoinDescriptionQuery) Unique(unique bool) *CoinDescriptionQuery {
	cdq.unique = &unique
	return cdq
}

// Order adds an order step to the query.
func (cdq *CoinDescriptionQuery) Order(o ...OrderFunc) *CoinDescriptionQuery {
	cdq.order = append(cdq.order, o...)
	return cdq
}

// First returns the first CoinDescription entity from the query.
// Returns a *NotFoundError when no CoinDescription was found.
func (cdq *CoinDescriptionQuery) First(ctx context.Context) (*CoinDescription, error) {
	nodes, err := cdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coindescription.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cdq *CoinDescriptionQuery) FirstX(ctx context.Context) *CoinDescription {
	node, err := cdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CoinDescription ID from the query.
// Returns a *NotFoundError when no CoinDescription ID was found.
func (cdq *CoinDescriptionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coindescription.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cdq *CoinDescriptionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CoinDescription entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CoinDescription entity is found.
// Returns a *NotFoundError when no CoinDescription entities are found.
func (cdq *CoinDescriptionQuery) Only(ctx context.Context) (*CoinDescription, error) {
	nodes, err := cdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coindescription.Label}
	default:
		return nil, &NotSingularError{coindescription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cdq *CoinDescriptionQuery) OnlyX(ctx context.Context) *CoinDescription {
	node, err := cdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CoinDescription ID in the query.
// Returns a *NotSingularError when more than one CoinDescription ID is found.
// Returns a *NotFoundError when no entities are found.
func (cdq *CoinDescriptionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coindescription.Label}
	default:
		err = &NotSingularError{coindescription.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cdq *CoinDescriptionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CoinDescriptions.
func (cdq *CoinDescriptionQuery) All(ctx context.Context) ([]*CoinDescription, error) {
	if err := cdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cdq *CoinDescriptionQuery) AllX(ctx context.Context) []*CoinDescription {
	nodes, err := cdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CoinDescription IDs.
func (cdq *CoinDescriptionQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := cdq.Select(coindescription.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cdq *CoinDescriptionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cdq *CoinDescriptionQuery) Count(ctx context.Context) (int, error) {
	if err := cdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cdq *CoinDescriptionQuery) CountX(ctx context.Context) int {
	count, err := cdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cdq *CoinDescriptionQuery) Exist(ctx context.Context) (bool, error) {
	if err := cdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cdq *CoinDescriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := cdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CoinDescriptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cdq *CoinDescriptionQuery) Clone() *CoinDescriptionQuery {
	if cdq == nil {
		return nil
	}
	return &CoinDescriptionQuery{
		config:     cdq.config,
		limit:      cdq.limit,
		offset:     cdq.offset,
		order:      append([]OrderFunc{}, cdq.order...),
		predicates: append([]predicate.CoinDescription{}, cdq.predicates...),
		// clone intermediate query.
		sql:    cdq.sql.Clone(),
		path:   cdq.path,
		unique: cdq.unique,
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
//	client.CoinDescription.Query().
//		GroupBy(coindescription.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cdq *CoinDescriptionQuery) GroupBy(field string, fields ...string) *CoinDescriptionGroupBy {
	grbuild := &CoinDescriptionGroupBy{config: cdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cdq.sqlQuery(ctx), nil
	}
	grbuild.label = coindescription.Label
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
//	client.CoinDescription.Query().
//		Select(coindescription.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (cdq *CoinDescriptionQuery) Select(fields ...string) *CoinDescriptionSelect {
	cdq.fields = append(cdq.fields, fields...)
	selbuild := &CoinDescriptionSelect{CoinDescriptionQuery: cdq}
	selbuild.label = coindescription.Label
	selbuild.flds, selbuild.scan = &cdq.fields, selbuild.Scan
	return selbuild
}

func (cdq *CoinDescriptionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cdq.fields {
		if !coindescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cdq.path != nil {
		prev, err := cdq.path(ctx)
		if err != nil {
			return err
		}
		cdq.sql = prev
	}
	if coindescription.Policy == nil {
		return errors.New("ent: uninitialized coindescription.Policy (forgotten import ent/runtime?)")
	}
	if err := coindescription.Policy.EvalQuery(ctx, cdq); err != nil {
		return err
	}
	return nil
}

func (cdq *CoinDescriptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CoinDescription, error) {
	var (
		nodes = []*CoinDescription{}
		_spec = cdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*CoinDescription).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &CoinDescription{config: cdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(cdq.modifiers) > 0 {
		_spec.Modifiers = cdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (cdq *CoinDescriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cdq.querySpec()
	if len(cdq.modifiers) > 0 {
		_spec.Modifiers = cdq.modifiers
	}
	_spec.Node.Columns = cdq.fields
	if len(cdq.fields) > 0 {
		_spec.Unique = cdq.unique != nil && *cdq.unique
	}
	return sqlgraph.CountNodes(ctx, cdq.driver, _spec)
}

func (cdq *CoinDescriptionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (cdq *CoinDescriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coindescription.Table,
			Columns: coindescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coindescription.FieldID,
			},
		},
		From:   cdq.sql,
		Unique: true,
	}
	if unique := cdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coindescription.FieldID)
		for i := range fields {
			if fields[i] != coindescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cdq *CoinDescriptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cdq.driver.Dialect())
	t1 := builder.Table(coindescription.Table)
	columns := cdq.fields
	if len(columns) == 0 {
		columns = coindescription.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cdq.sql != nil {
		selector = cdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cdq.unique != nil && *cdq.unique {
		selector.Distinct()
	}
	for _, m := range cdq.modifiers {
		m(selector)
	}
	for _, p := range cdq.predicates {
		p(selector)
	}
	for _, p := range cdq.order {
		p(selector)
	}
	if offset := cdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (cdq *CoinDescriptionQuery) ForUpdate(opts ...sql.LockOption) *CoinDescriptionQuery {
	if cdq.driver.Dialect() == dialect.Postgres {
		cdq.Unique(false)
	}
	cdq.modifiers = append(cdq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return cdq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (cdq *CoinDescriptionQuery) ForShare(opts ...sql.LockOption) *CoinDescriptionQuery {
	if cdq.driver.Dialect() == dialect.Postgres {
		cdq.Unique(false)
	}
	cdq.modifiers = append(cdq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return cdq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cdq *CoinDescriptionQuery) Modify(modifiers ...func(s *sql.Selector)) *CoinDescriptionSelect {
	cdq.modifiers = append(cdq.modifiers, modifiers...)
	return cdq.Select()
}

// CoinDescriptionGroupBy is the group-by builder for CoinDescription entities.
type CoinDescriptionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cdgb *CoinDescriptionGroupBy) Aggregate(fns ...AggregateFunc) *CoinDescriptionGroupBy {
	cdgb.fns = append(cdgb.fns, fns...)
	return cdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cdgb *CoinDescriptionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cdgb.path(ctx)
	if err != nil {
		return err
	}
	cdgb.sql = query
	return cdgb.sqlScan(ctx, v)
}

func (cdgb *CoinDescriptionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cdgb.fields {
		if !coindescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cdgb *CoinDescriptionGroupBy) sqlQuery() *sql.Selector {
	selector := cdgb.sql.Select()
	aggregation := make([]string, 0, len(cdgb.fns))
	for _, fn := range cdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cdgb.fields)+len(cdgb.fns))
		for _, f := range cdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cdgb.fields...)...)
}

// CoinDescriptionSelect is the builder for selecting fields of CoinDescription entities.
type CoinDescriptionSelect struct {
	*CoinDescriptionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cds *CoinDescriptionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cds.prepareQuery(ctx); err != nil {
		return err
	}
	cds.sql = cds.CoinDescriptionQuery.sqlQuery(ctx)
	return cds.sqlScan(ctx, v)
}

func (cds *CoinDescriptionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cds.sql.Query()
	if err := cds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cds *CoinDescriptionSelect) Modify(modifiers ...func(s *sql.Selector)) *CoinDescriptionSelect {
	cds.modifiers = append(cds.modifiers, modifiers...)
	return cds
}