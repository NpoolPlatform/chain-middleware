// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrency"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FiatCurrencyUpdate is the builder for updating FiatCurrency entities.
type FiatCurrencyUpdate struct {
	config
	hooks     []Hook
	mutation  *FiatCurrencyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FiatCurrencyUpdate builder.
func (fcu *FiatCurrencyUpdate) Where(ps ...predicate.FiatCurrency) *FiatCurrencyUpdate {
	fcu.mutation.Where(ps...)
	return fcu
}

// SetCreatedAt sets the "created_at" field.
func (fcu *FiatCurrencyUpdate) SetCreatedAt(u uint32) *FiatCurrencyUpdate {
	fcu.mutation.ResetCreatedAt()
	fcu.mutation.SetCreatedAt(u)
	return fcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fcu *FiatCurrencyUpdate) SetNillableCreatedAt(u *uint32) *FiatCurrencyUpdate {
	if u != nil {
		fcu.SetCreatedAt(*u)
	}
	return fcu
}

// AddCreatedAt adds u to the "created_at" field.
func (fcu *FiatCurrencyUpdate) AddCreatedAt(u int32) *FiatCurrencyUpdate {
	fcu.mutation.AddCreatedAt(u)
	return fcu
}

// SetUpdatedAt sets the "updated_at" field.
func (fcu *FiatCurrencyUpdate) SetUpdatedAt(u uint32) *FiatCurrencyUpdate {
	fcu.mutation.ResetUpdatedAt()
	fcu.mutation.SetUpdatedAt(u)
	return fcu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fcu *FiatCurrencyUpdate) AddUpdatedAt(u int32) *FiatCurrencyUpdate {
	fcu.mutation.AddUpdatedAt(u)
	return fcu
}

// SetDeletedAt sets the "deleted_at" field.
func (fcu *FiatCurrencyUpdate) SetDeletedAt(u uint32) *FiatCurrencyUpdate {
	fcu.mutation.ResetDeletedAt()
	fcu.mutation.SetDeletedAt(u)
	return fcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fcu *FiatCurrencyUpdate) SetNillableDeletedAt(u *uint32) *FiatCurrencyUpdate {
	if u != nil {
		fcu.SetDeletedAt(*u)
	}
	return fcu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fcu *FiatCurrencyUpdate) AddDeletedAt(u int32) *FiatCurrencyUpdate {
	fcu.mutation.AddDeletedAt(u)
	return fcu
}

// SetFiatCurrencyTypeID sets the "fiat_currency_type_id" field.
func (fcu *FiatCurrencyUpdate) SetFiatCurrencyTypeID(u uuid.UUID) *FiatCurrencyUpdate {
	fcu.mutation.SetFiatCurrencyTypeID(u)
	return fcu
}

// SetNillableFiatCurrencyTypeID sets the "fiat_currency_type_id" field if the given value is not nil.
func (fcu *FiatCurrencyUpdate) SetNillableFiatCurrencyTypeID(u *uuid.UUID) *FiatCurrencyUpdate {
	if u != nil {
		fcu.SetFiatCurrencyTypeID(*u)
	}
	return fcu
}

// ClearFiatCurrencyTypeID clears the value of the "fiat_currency_type_id" field.
func (fcu *FiatCurrencyUpdate) ClearFiatCurrencyTypeID() *FiatCurrencyUpdate {
	fcu.mutation.ClearFiatCurrencyTypeID()
	return fcu
}

// SetFeedType sets the "feed_type" field.
func (fcu *FiatCurrencyUpdate) SetFeedType(s string) *FiatCurrencyUpdate {
	fcu.mutation.SetFeedType(s)
	return fcu
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (fcu *FiatCurrencyUpdate) SetNillableFeedType(s *string) *FiatCurrencyUpdate {
	if s != nil {
		fcu.SetFeedType(*s)
	}
	return fcu
}

// ClearFeedType clears the value of the "feed_type" field.
func (fcu *FiatCurrencyUpdate) ClearFeedType() *FiatCurrencyUpdate {
	fcu.mutation.ClearFeedType()
	return fcu
}

// SetMarketValueLow sets the "market_value_low" field.
func (fcu *FiatCurrencyUpdate) SetMarketValueLow(d decimal.Decimal) *FiatCurrencyUpdate {
	fcu.mutation.SetMarketValueLow(d)
	return fcu
}

// SetNillableMarketValueLow sets the "market_value_low" field if the given value is not nil.
func (fcu *FiatCurrencyUpdate) SetNillableMarketValueLow(d *decimal.Decimal) *FiatCurrencyUpdate {
	if d != nil {
		fcu.SetMarketValueLow(*d)
	}
	return fcu
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (fcu *FiatCurrencyUpdate) ClearMarketValueLow() *FiatCurrencyUpdate {
	fcu.mutation.ClearMarketValueLow()
	return fcu
}

// SetMarketValueHigh sets the "market_value_high" field.
func (fcu *FiatCurrencyUpdate) SetMarketValueHigh(d decimal.Decimal) *FiatCurrencyUpdate {
	fcu.mutation.SetMarketValueHigh(d)
	return fcu
}

// SetNillableMarketValueHigh sets the "market_value_high" field if the given value is not nil.
func (fcu *FiatCurrencyUpdate) SetNillableMarketValueHigh(d *decimal.Decimal) *FiatCurrencyUpdate {
	if d != nil {
		fcu.SetMarketValueHigh(*d)
	}
	return fcu
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (fcu *FiatCurrencyUpdate) ClearMarketValueHigh() *FiatCurrencyUpdate {
	fcu.mutation.ClearMarketValueHigh()
	return fcu
}

// Mutation returns the FiatCurrencyMutation object of the builder.
func (fcu *FiatCurrencyUpdate) Mutation() *FiatCurrencyMutation {
	return fcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fcu *FiatCurrencyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := fcu.defaults(); err != nil {
		return 0, err
	}
	if len(fcu.hooks) == 0 {
		affected, err = fcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FiatCurrencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fcu.mutation = mutation
			affected, err = fcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fcu.hooks) - 1; i >= 0; i-- {
			if fcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fcu *FiatCurrencyUpdate) SaveX(ctx context.Context) int {
	affected, err := fcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fcu *FiatCurrencyUpdate) Exec(ctx context.Context) error {
	_, err := fcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcu *FiatCurrencyUpdate) ExecX(ctx context.Context) {
	if err := fcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fcu *FiatCurrencyUpdate) defaults() error {
	if _, ok := fcu.mutation.UpdatedAt(); !ok {
		if fiatcurrency.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrency.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fiatcurrency.UpdateDefaultUpdatedAt()
		fcu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fcu *FiatCurrencyUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FiatCurrencyUpdate {
	fcu.modifiers = append(fcu.modifiers, modifiers...)
	return fcu
}

func (fcu *FiatCurrencyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fiatcurrency.Table,
			Columns: fiatcurrency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fiatcurrency.FieldID,
			},
		},
	}
	if ps := fcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fcu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldCreatedAt,
		})
	}
	if value, ok := fcu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldCreatedAt,
		})
	}
	if value, ok := fcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldUpdatedAt,
		})
	}
	if value, ok := fcu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldUpdatedAt,
		})
	}
	if value, ok := fcu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldDeletedAt,
		})
	}
	if value, ok := fcu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldDeletedAt,
		})
	}
	if value, ok := fcu.mutation.FiatCurrencyTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fiatcurrency.FieldFiatCurrencyTypeID,
		})
	}
	if fcu.mutation.FiatCurrencyTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fiatcurrency.FieldFiatCurrencyTypeID,
		})
	}
	if value, ok := fcu.mutation.FeedType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fiatcurrency.FieldFeedType,
		})
	}
	if fcu.mutation.FeedTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fiatcurrency.FieldFeedType,
		})
	}
	if value, ok := fcu.mutation.MarketValueLow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrency.FieldMarketValueLow,
		})
	}
	if fcu.mutation.MarketValueLowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fiatcurrency.FieldMarketValueLow,
		})
	}
	if value, ok := fcu.mutation.MarketValueHigh(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrency.FieldMarketValueHigh,
		})
	}
	if fcu.mutation.MarketValueHighCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fiatcurrency.FieldMarketValueHigh,
		})
	}
	_spec.Modifiers = fcu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, fcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fiatcurrency.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FiatCurrencyUpdateOne is the builder for updating a single FiatCurrency entity.
type FiatCurrencyUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FiatCurrencyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (fcuo *FiatCurrencyUpdateOne) SetCreatedAt(u uint32) *FiatCurrencyUpdateOne {
	fcuo.mutation.ResetCreatedAt()
	fcuo.mutation.SetCreatedAt(u)
	return fcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fcuo *FiatCurrencyUpdateOne) SetNillableCreatedAt(u *uint32) *FiatCurrencyUpdateOne {
	if u != nil {
		fcuo.SetCreatedAt(*u)
	}
	return fcuo
}

// AddCreatedAt adds u to the "created_at" field.
func (fcuo *FiatCurrencyUpdateOne) AddCreatedAt(u int32) *FiatCurrencyUpdateOne {
	fcuo.mutation.AddCreatedAt(u)
	return fcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fcuo *FiatCurrencyUpdateOne) SetUpdatedAt(u uint32) *FiatCurrencyUpdateOne {
	fcuo.mutation.ResetUpdatedAt()
	fcuo.mutation.SetUpdatedAt(u)
	return fcuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (fcuo *FiatCurrencyUpdateOne) AddUpdatedAt(u int32) *FiatCurrencyUpdateOne {
	fcuo.mutation.AddUpdatedAt(u)
	return fcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fcuo *FiatCurrencyUpdateOne) SetDeletedAt(u uint32) *FiatCurrencyUpdateOne {
	fcuo.mutation.ResetDeletedAt()
	fcuo.mutation.SetDeletedAt(u)
	return fcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fcuo *FiatCurrencyUpdateOne) SetNillableDeletedAt(u *uint32) *FiatCurrencyUpdateOne {
	if u != nil {
		fcuo.SetDeletedAt(*u)
	}
	return fcuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (fcuo *FiatCurrencyUpdateOne) AddDeletedAt(u int32) *FiatCurrencyUpdateOne {
	fcuo.mutation.AddDeletedAt(u)
	return fcuo
}

// SetFiatCurrencyTypeID sets the "fiat_currency_type_id" field.
func (fcuo *FiatCurrencyUpdateOne) SetFiatCurrencyTypeID(u uuid.UUID) *FiatCurrencyUpdateOne {
	fcuo.mutation.SetFiatCurrencyTypeID(u)
	return fcuo
}

// SetNillableFiatCurrencyTypeID sets the "fiat_currency_type_id" field if the given value is not nil.
func (fcuo *FiatCurrencyUpdateOne) SetNillableFiatCurrencyTypeID(u *uuid.UUID) *FiatCurrencyUpdateOne {
	if u != nil {
		fcuo.SetFiatCurrencyTypeID(*u)
	}
	return fcuo
}

// ClearFiatCurrencyTypeID clears the value of the "fiat_currency_type_id" field.
func (fcuo *FiatCurrencyUpdateOne) ClearFiatCurrencyTypeID() *FiatCurrencyUpdateOne {
	fcuo.mutation.ClearFiatCurrencyTypeID()
	return fcuo
}

// SetFeedType sets the "feed_type" field.
func (fcuo *FiatCurrencyUpdateOne) SetFeedType(s string) *FiatCurrencyUpdateOne {
	fcuo.mutation.SetFeedType(s)
	return fcuo
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (fcuo *FiatCurrencyUpdateOne) SetNillableFeedType(s *string) *FiatCurrencyUpdateOne {
	if s != nil {
		fcuo.SetFeedType(*s)
	}
	return fcuo
}

// ClearFeedType clears the value of the "feed_type" field.
func (fcuo *FiatCurrencyUpdateOne) ClearFeedType() *FiatCurrencyUpdateOne {
	fcuo.mutation.ClearFeedType()
	return fcuo
}

// SetMarketValueLow sets the "market_value_low" field.
func (fcuo *FiatCurrencyUpdateOne) SetMarketValueLow(d decimal.Decimal) *FiatCurrencyUpdateOne {
	fcuo.mutation.SetMarketValueLow(d)
	return fcuo
}

// SetNillableMarketValueLow sets the "market_value_low" field if the given value is not nil.
func (fcuo *FiatCurrencyUpdateOne) SetNillableMarketValueLow(d *decimal.Decimal) *FiatCurrencyUpdateOne {
	if d != nil {
		fcuo.SetMarketValueLow(*d)
	}
	return fcuo
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (fcuo *FiatCurrencyUpdateOne) ClearMarketValueLow() *FiatCurrencyUpdateOne {
	fcuo.mutation.ClearMarketValueLow()
	return fcuo
}

// SetMarketValueHigh sets the "market_value_high" field.
func (fcuo *FiatCurrencyUpdateOne) SetMarketValueHigh(d decimal.Decimal) *FiatCurrencyUpdateOne {
	fcuo.mutation.SetMarketValueHigh(d)
	return fcuo
}

// SetNillableMarketValueHigh sets the "market_value_high" field if the given value is not nil.
func (fcuo *FiatCurrencyUpdateOne) SetNillableMarketValueHigh(d *decimal.Decimal) *FiatCurrencyUpdateOne {
	if d != nil {
		fcuo.SetMarketValueHigh(*d)
	}
	return fcuo
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (fcuo *FiatCurrencyUpdateOne) ClearMarketValueHigh() *FiatCurrencyUpdateOne {
	fcuo.mutation.ClearMarketValueHigh()
	return fcuo
}

// Mutation returns the FiatCurrencyMutation object of the builder.
func (fcuo *FiatCurrencyUpdateOne) Mutation() *FiatCurrencyMutation {
	return fcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fcuo *FiatCurrencyUpdateOne) Select(field string, fields ...string) *FiatCurrencyUpdateOne {
	fcuo.fields = append([]string{field}, fields...)
	return fcuo
}

// Save executes the query and returns the updated FiatCurrency entity.
func (fcuo *FiatCurrencyUpdateOne) Save(ctx context.Context) (*FiatCurrency, error) {
	var (
		err  error
		node *FiatCurrency
	)
	if err := fcuo.defaults(); err != nil {
		return nil, err
	}
	if len(fcuo.hooks) == 0 {
		node, err = fcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FiatCurrencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fcuo.mutation = mutation
			node, err = fcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fcuo.hooks) - 1; i >= 0; i-- {
			if fcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*FiatCurrency)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FiatCurrencyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fcuo *FiatCurrencyUpdateOne) SaveX(ctx context.Context) *FiatCurrency {
	node, err := fcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fcuo *FiatCurrencyUpdateOne) Exec(ctx context.Context) error {
	_, err := fcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcuo *FiatCurrencyUpdateOne) ExecX(ctx context.Context) {
	if err := fcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fcuo *FiatCurrencyUpdateOne) defaults() error {
	if _, ok := fcuo.mutation.UpdatedAt(); !ok {
		if fiatcurrency.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrency.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fiatcurrency.UpdateDefaultUpdatedAt()
		fcuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fcuo *FiatCurrencyUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FiatCurrencyUpdateOne {
	fcuo.modifiers = append(fcuo.modifiers, modifiers...)
	return fcuo
}

func (fcuo *FiatCurrencyUpdateOne) sqlSave(ctx context.Context) (_node *FiatCurrency, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fiatcurrency.Table,
			Columns: fiatcurrency.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fiatcurrency.FieldID,
			},
		},
	}
	id, ok := fcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FiatCurrency.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fiatcurrency.FieldID)
		for _, f := range fields {
			if !fiatcurrency.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fiatcurrency.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fcuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldCreatedAt,
		})
	}
	if value, ok := fcuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldCreatedAt,
		})
	}
	if value, ok := fcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldUpdatedAt,
		})
	}
	if value, ok := fcuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldUpdatedAt,
		})
	}
	if value, ok := fcuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldDeletedAt,
		})
	}
	if value, ok := fcuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrency.FieldDeletedAt,
		})
	}
	if value, ok := fcuo.mutation.FiatCurrencyTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fiatcurrency.FieldFiatCurrencyTypeID,
		})
	}
	if fcuo.mutation.FiatCurrencyTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: fiatcurrency.FieldFiatCurrencyTypeID,
		})
	}
	if value, ok := fcuo.mutation.FeedType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fiatcurrency.FieldFeedType,
		})
	}
	if fcuo.mutation.FeedTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: fiatcurrency.FieldFeedType,
		})
	}
	if value, ok := fcuo.mutation.MarketValueLow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrency.FieldMarketValueLow,
		})
	}
	if fcuo.mutation.MarketValueLowCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fiatcurrency.FieldMarketValueLow,
		})
	}
	if value, ok := fcuo.mutation.MarketValueHigh(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrency.FieldMarketValueHigh,
		})
	}
	if fcuo.mutation.MarketValueHighCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: fiatcurrency.FieldMarketValueHigh,
		})
	}
	_spec.Modifiers = fcuo.modifiers
	_node = &FiatCurrency{config: fcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fiatcurrency.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
