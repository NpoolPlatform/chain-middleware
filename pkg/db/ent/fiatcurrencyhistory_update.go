// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrencyhistory"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FiatCurrencyHistoryUpdate is the builder for updating FiatCurrencyHistory entities.
type FiatCurrencyHistoryUpdate struct {
	config
	hooks     []Hook
	mutation  *FiatCurrencyHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FiatCurrencyHistoryUpdate builder.
func (fchu *FiatCurrencyHistoryUpdate) Where(ps ...predicate.FiatCurrencyHistory) *FiatCurrencyHistoryUpdate {
	fchu.mutation.Where(ps...)
	return fchu
}

// SetCreatedAt sets the "created_at" field.
func (fchu *FiatCurrencyHistoryUpdate) SetCreatedAt(u uint32) *FiatCurrencyHistoryUpdate {
	fchu.mutation.ResetCreatedAt()
	fchu.mutation.SetCreatedAt(u)
	return fchu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fchu *FiatCurrencyHistoryUpdate) SetNillableCreatedAt(u *uint32) *FiatCurrencyHistoryUpdate {
	if u != nil {
		fchu.SetCreatedAt(*u)
	}
	return fchu
}

// AddCreatedAt adds u to the "created_at" field.
func (fchu *FiatCurrencyHistoryUpdate) AddCreatedAt(u int32) *FiatCurrencyHistoryUpdate {
	fchu.mutation.AddCreatedAt(u)
	return fchu
}

// SetUpdatedAt sets the "updated_at" field.
func (fchu *FiatCurrencyHistoryUpdate) SetUpdatedAt(u uint32) *FiatCurrencyHistoryUpdate {
	fchu.mutation.ResetUpdatedAt()
	fchu.mutation.SetUpdatedAt(u)
	return fchu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fchu *FiatCurrencyHistoryUpdate) AddUpdatedAt(u int32) *FiatCurrencyHistoryUpdate {
	fchu.mutation.AddUpdatedAt(u)
	return fchu
}

// SetDeletedAt sets the "deleted_at" field.
func (fchu *FiatCurrencyHistoryUpdate) SetDeletedAt(u uint32) *FiatCurrencyHistoryUpdate {
	fchu.mutation.ResetDeletedAt()
	fchu.mutation.SetDeletedAt(u)
	return fchu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fchu *FiatCurrencyHistoryUpdate) SetNillableDeletedAt(u *uint32) *FiatCurrencyHistoryUpdate {
	if u != nil {
		fchu.SetDeletedAt(*u)
	}
	return fchu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fchu *FiatCurrencyHistoryUpdate) AddDeletedAt(u int32) *FiatCurrencyHistoryUpdate {
	fchu.mutation.AddDeletedAt(u)
	return fchu
}

// SetEntID sets the "ent_id" field.
func (fchu *FiatCurrencyHistoryUpdate) SetEntID(u uuid.UUID) *FiatCurrencyHistoryUpdate {
	fchu.mutation.SetEntID(u)
	return fchu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (fchu *FiatCurrencyHistoryUpdate) SetNillableEntID(u *uuid.UUID) *FiatCurrencyHistoryUpdate {
	if u != nil {
		fchu.SetEntID(*u)
	}
	return fchu
}

// SetFiatID sets the "fiat_id" field.
func (fchu *FiatCurrencyHistoryUpdate) SetFiatID(u uuid.UUID) *FiatCurrencyHistoryUpdate {
	fchu.mutation.SetFiatID(u)
	return fchu
}

// SetNillableFiatID sets the "fiat_id" field if the given value is not nil.
func (fchu *FiatCurrencyHistoryUpdate) SetNillableFiatID(u *uuid.UUID) *FiatCurrencyHistoryUpdate {
	if u != nil {
		fchu.SetFiatID(*u)
	}
	return fchu
}

// ClearFiatID clears the value of the "fiat_id" field.
func (fchu *FiatCurrencyHistoryUpdate) ClearFiatID() *FiatCurrencyHistoryUpdate {
	fchu.mutation.ClearFiatID()
	return fchu
}

// SetFeedType sets the "feed_type" field.
func (fchu *FiatCurrencyHistoryUpdate) SetFeedType(s string) *FiatCurrencyHistoryUpdate {
	fchu.mutation.SetFeedType(s)
	return fchu
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (fchu *FiatCurrencyHistoryUpdate) SetNillableFeedType(s *string) *FiatCurrencyHistoryUpdate {
	if s != nil {
		fchu.SetFeedType(*s)
	}
	return fchu
}

// ClearFeedType clears the value of the "feed_type" field.
func (fchu *FiatCurrencyHistoryUpdate) ClearFeedType() *FiatCurrencyHistoryUpdate {
	fchu.mutation.ClearFeedType()
	return fchu
}

// SetMarketValueLow sets the "market_value_low" field.
func (fchu *FiatCurrencyHistoryUpdate) SetMarketValueLow(d decimal.Decimal) *FiatCurrencyHistoryUpdate {
	fchu.mutation.SetMarketValueLow(d)
	return fchu
}

// SetNillableMarketValueLow sets the "market_value_low" field if the given value is not nil.
func (fchu *FiatCurrencyHistoryUpdate) SetNillableMarketValueLow(d *decimal.Decimal) *FiatCurrencyHistoryUpdate {
	if d != nil {
		fchu.SetMarketValueLow(*d)
	}
	return fchu
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (fchu *FiatCurrencyHistoryUpdate) ClearMarketValueLow() *FiatCurrencyHistoryUpdate {
	fchu.mutation.ClearMarketValueLow()
	return fchu
}

// SetMarketValueHigh sets the "market_value_high" field.
func (fchu *FiatCurrencyHistoryUpdate) SetMarketValueHigh(d decimal.Decimal) *FiatCurrencyHistoryUpdate {
	fchu.mutation.SetMarketValueHigh(d)
	return fchu
}

// SetNillableMarketValueHigh sets the "market_value_high" field if the given value is not nil.
func (fchu *FiatCurrencyHistoryUpdate) SetNillableMarketValueHigh(d *decimal.Decimal) *FiatCurrencyHistoryUpdate {
	if d != nil {
		fchu.SetMarketValueHigh(*d)
	}
	return fchu
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (fchu *FiatCurrencyHistoryUpdate) ClearMarketValueHigh() *FiatCurrencyHistoryUpdate {
	fchu.mutation.ClearMarketValueHigh()
	return fchu
}

// Mutation returns the FiatCurrencyHistoryMutation object of the builder.
func (fchu *FiatCurrencyHistoryUpdate) Mutation() *FiatCurrencyHistoryMutation {
	return fchu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fchu *FiatCurrencyHistoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := fchu.defaults(); err != nil {
		return 0, err
	}
	if len(fchu.hooks) == 0 {
		affected, err = fchu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FiatCurrencyHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fchu.mutation = mutation
			affected, err = fchu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fchu.hooks) - 1; i >= 0; i-- {
			if fchu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fchu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fchu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fchu *FiatCurrencyHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := fchu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fchu *FiatCurrencyHistoryUpdate) Exec(ctx context.Context) error {
	_, err := fchu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fchu *FiatCurrencyHistoryUpdate) ExecX(ctx context.Context) {
	if err := fchu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fchu *FiatCurrencyHistoryUpdate) defaults() error {
	if _, ok := fchu.mutation.UpdatedAt(); !ok {
		if fiatcurrencyhistory.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrencyhistory.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fiatcurrencyhistory.UpdateDefaultUpdatedAt()
		fchu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fchu *FiatCurrencyHistoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FiatCurrencyHistoryUpdate {
	fchu.modifiers = append(fchu.modifiers, modifiers...)
	return fchu
}

func (fchu *FiatCurrencyHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fiatcurrencyhistory.Table,
			Columns: fiatcurrencyhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fiatcurrencyhistory.FieldID,
			},
		},
	}
	if ps := fchu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fchu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldCreatedAt,
		})
	}
	if value, ok := fchu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldCreatedAt,
		})
	}
	if value, ok := fchu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldUpdatedAt,
		})
	}
	if value, ok := fchu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldUpdatedAt,
		})
	}
	if value, ok := fchu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldDeletedAt,
		})
	}
	if value, ok := fchu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldDeletedAt,
		})
	}
	if value, ok := fchu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fiatcurrencyhistory.FieldEntID,
		})
	}
	if value, ok := fchu.mutation.FiatID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fiatcurrencyhistory.FieldFiatID,
		})
	}
	if fchu.mutation.FiatIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fiatcurrencyhistory.FieldFiatID,
		})
	}
	if value, ok := fchu.mutation.FeedType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fiatcurrencyhistory.FieldFeedType,
		})
	}
	if fchu.mutation.FeedTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fiatcurrencyhistory.FieldFeedType,
		})
	}
	if value, ok := fchu.mutation.MarketValueLow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrencyhistory.FieldMarketValueLow,
		})
	}
	if fchu.mutation.MarketValueLowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fiatcurrencyhistory.FieldMarketValueLow,
		})
	}
	if value, ok := fchu.mutation.MarketValueHigh(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrencyhistory.FieldMarketValueHigh,
		})
	}
	if fchu.mutation.MarketValueHighCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fiatcurrencyhistory.FieldMarketValueHigh,
		})
	}
	_spec.Modifiers = fchu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, fchu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fiatcurrencyhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FiatCurrencyHistoryUpdateOne is the builder for updating a single FiatCurrencyHistory entity.
type FiatCurrencyHistoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FiatCurrencyHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetCreatedAt(u uint32) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.ResetCreatedAt()
	fchuo.mutation.SetCreatedAt(u)
	return fchuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetNillableCreatedAt(u *uint32) *FiatCurrencyHistoryUpdateOne {
	if u != nil {
		fchuo.SetCreatedAt(*u)
	}
	return fchuo
}

// AddCreatedAt adds u to the "created_at" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) AddCreatedAt(u int32) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.AddCreatedAt(u)
	return fchuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetUpdatedAt(u uint32) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.ResetUpdatedAt()
	fchuo.mutation.SetUpdatedAt(u)
	return fchuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) AddUpdatedAt(u int32) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.AddUpdatedAt(u)
	return fchuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetDeletedAt(u uint32) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.ResetDeletedAt()
	fchuo.mutation.SetDeletedAt(u)
	return fchuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetNillableDeletedAt(u *uint32) *FiatCurrencyHistoryUpdateOne {
	if u != nil {
		fchuo.SetDeletedAt(*u)
	}
	return fchuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) AddDeletedAt(u int32) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.AddDeletedAt(u)
	return fchuo
}

// SetEntID sets the "ent_id" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetEntID(u uuid.UUID) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.SetEntID(u)
	return fchuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetNillableEntID(u *uuid.UUID) *FiatCurrencyHistoryUpdateOne {
	if u != nil {
		fchuo.SetEntID(*u)
	}
	return fchuo
}

// SetFiatID sets the "fiat_id" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetFiatID(u uuid.UUID) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.SetFiatID(u)
	return fchuo
}

// SetNillableFiatID sets the "fiat_id" field if the given value is not nil.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetNillableFiatID(u *uuid.UUID) *FiatCurrencyHistoryUpdateOne {
	if u != nil {
		fchuo.SetFiatID(*u)
	}
	return fchuo
}

// ClearFiatID clears the value of the "fiat_id" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) ClearFiatID() *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.ClearFiatID()
	return fchuo
}

// SetFeedType sets the "feed_type" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetFeedType(s string) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.SetFeedType(s)
	return fchuo
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetNillableFeedType(s *string) *FiatCurrencyHistoryUpdateOne {
	if s != nil {
		fchuo.SetFeedType(*s)
	}
	return fchuo
}

// ClearFeedType clears the value of the "feed_type" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) ClearFeedType() *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.ClearFeedType()
	return fchuo
}

// SetMarketValueLow sets the "market_value_low" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetMarketValueLow(d decimal.Decimal) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.SetMarketValueLow(d)
	return fchuo
}

// SetNillableMarketValueLow sets the "market_value_low" field if the given value is not nil.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetNillableMarketValueLow(d *decimal.Decimal) *FiatCurrencyHistoryUpdateOne {
	if d != nil {
		fchuo.SetMarketValueLow(*d)
	}
	return fchuo
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) ClearMarketValueLow() *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.ClearMarketValueLow()
	return fchuo
}

// SetMarketValueHigh sets the "market_value_high" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetMarketValueHigh(d decimal.Decimal) *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.SetMarketValueHigh(d)
	return fchuo
}

// SetNillableMarketValueHigh sets the "market_value_high" field if the given value is not nil.
func (fchuo *FiatCurrencyHistoryUpdateOne) SetNillableMarketValueHigh(d *decimal.Decimal) *FiatCurrencyHistoryUpdateOne {
	if d != nil {
		fchuo.SetMarketValueHigh(*d)
	}
	return fchuo
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (fchuo *FiatCurrencyHistoryUpdateOne) ClearMarketValueHigh() *FiatCurrencyHistoryUpdateOne {
	fchuo.mutation.ClearMarketValueHigh()
	return fchuo
}

// Mutation returns the FiatCurrencyHistoryMutation object of the builder.
func (fchuo *FiatCurrencyHistoryUpdateOne) Mutation() *FiatCurrencyHistoryMutation {
	return fchuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fchuo *FiatCurrencyHistoryUpdateOne) Select(field string, fields ...string) *FiatCurrencyHistoryUpdateOne {
	fchuo.fields = append([]string{field}, fields...)
	return fchuo
}

// Save executes the query and returns the updated FiatCurrencyHistory entity.
func (fchuo *FiatCurrencyHistoryUpdateOne) Save(ctx context.Context) (*FiatCurrencyHistory, error) {
	var (
		err  error
		node *FiatCurrencyHistory
	)
	if err := fchuo.defaults(); err != nil {
		return nil, err
	}
	if len(fchuo.hooks) == 0 {
		node, err = fchuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FiatCurrencyHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fchuo.mutation = mutation
			node, err = fchuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fchuo.hooks) - 1; i >= 0; i-- {
			if fchuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fchuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fchuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*FiatCurrencyHistory)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FiatCurrencyHistoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fchuo *FiatCurrencyHistoryUpdateOne) SaveX(ctx context.Context) *FiatCurrencyHistory {
	node, err := fchuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fchuo *FiatCurrencyHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := fchuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fchuo *FiatCurrencyHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := fchuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fchuo *FiatCurrencyHistoryUpdateOne) defaults() error {
	if _, ok := fchuo.mutation.UpdatedAt(); !ok {
		if fiatcurrencyhistory.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrencyhistory.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fiatcurrencyhistory.UpdateDefaultUpdatedAt()
		fchuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fchuo *FiatCurrencyHistoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FiatCurrencyHistoryUpdateOne {
	fchuo.modifiers = append(fchuo.modifiers, modifiers...)
	return fchuo
}

func (fchuo *FiatCurrencyHistoryUpdateOne) sqlSave(ctx context.Context) (_node *FiatCurrencyHistory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fiatcurrencyhistory.Table,
			Columns: fiatcurrencyhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fiatcurrencyhistory.FieldID,
			},
		},
	}
	id, ok := fchuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FiatCurrencyHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fchuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fiatcurrencyhistory.FieldID)
		for _, f := range fields {
			if !fiatcurrencyhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fiatcurrencyhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fchuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fchuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldCreatedAt,
		})
	}
	if value, ok := fchuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldCreatedAt,
		})
	}
	if value, ok := fchuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldUpdatedAt,
		})
	}
	if value, ok := fchuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldUpdatedAt,
		})
	}
	if value, ok := fchuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldDeletedAt,
		})
	}
	if value, ok := fchuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldDeletedAt,
		})
	}
	if value, ok := fchuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fiatcurrencyhistory.FieldEntID,
		})
	}
	if value, ok := fchuo.mutation.FiatID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fiatcurrencyhistory.FieldFiatID,
		})
	}
	if fchuo.mutation.FiatIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fiatcurrencyhistory.FieldFiatID,
		})
	}
	if value, ok := fchuo.mutation.FeedType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fiatcurrencyhistory.FieldFeedType,
		})
	}
	if fchuo.mutation.FeedTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fiatcurrencyhistory.FieldFeedType,
		})
	}
	if value, ok := fchuo.mutation.MarketValueLow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrencyhistory.FieldMarketValueLow,
		})
	}
	if fchuo.mutation.MarketValueLowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fiatcurrencyhistory.FieldMarketValueLow,
		})
	}
	if value, ok := fchuo.mutation.MarketValueHigh(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrencyhistory.FieldMarketValueHigh,
		})
	}
	if fchuo.mutation.MarketValueHighCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fiatcurrencyhistory.FieldMarketValueHigh,
		})
	}
	_spec.Modifiers = fchuo.modifiers
	_node = &FiatCurrencyHistory{config: fchuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fchuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fiatcurrencyhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
