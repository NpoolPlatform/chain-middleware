// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiatcurrencyhistory"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// FiatCurrencyHistoryCreate is the builder for creating a FiatCurrencyHistory entity.
type FiatCurrencyHistoryCreate struct {
	config
	mutation *FiatCurrencyHistoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (fchc *FiatCurrencyHistoryCreate) SetCreatedAt(u uint32) *FiatCurrencyHistoryCreate {
	fchc.mutation.SetCreatedAt(u)
	return fchc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fchc *FiatCurrencyHistoryCreate) SetNillableCreatedAt(u *uint32) *FiatCurrencyHistoryCreate {
	if u != nil {
		fchc.SetCreatedAt(*u)
	}
	return fchc
}

// SetUpdatedAt sets the "updated_at" field.
func (fchc *FiatCurrencyHistoryCreate) SetUpdatedAt(u uint32) *FiatCurrencyHistoryCreate {
	fchc.mutation.SetUpdatedAt(u)
	return fchc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fchc *FiatCurrencyHistoryCreate) SetNillableUpdatedAt(u *uint32) *FiatCurrencyHistoryCreate {
	if u != nil {
		fchc.SetUpdatedAt(*u)
	}
	return fchc
}

// SetDeletedAt sets the "deleted_at" field.
func (fchc *FiatCurrencyHistoryCreate) SetDeletedAt(u uint32) *FiatCurrencyHistoryCreate {
	fchc.mutation.SetDeletedAt(u)
	return fchc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fchc *FiatCurrencyHistoryCreate) SetNillableDeletedAt(u *uint32) *FiatCurrencyHistoryCreate {
	if u != nil {
		fchc.SetDeletedAt(*u)
	}
	return fchc
}

// SetFiatID sets the "fiat_id" field.
func (fchc *FiatCurrencyHistoryCreate) SetFiatID(u uuid.UUID) *FiatCurrencyHistoryCreate {
	fchc.mutation.SetFiatID(u)
	return fchc
}

// SetNillableFiatID sets the "fiat_id" field if the given value is not nil.
func (fchc *FiatCurrencyHistoryCreate) SetNillableFiatID(u *uuid.UUID) *FiatCurrencyHistoryCreate {
	if u != nil {
		fchc.SetFiatID(*u)
	}
	return fchc
}

// SetFeedType sets the "feed_type" field.
func (fchc *FiatCurrencyHistoryCreate) SetFeedType(s string) *FiatCurrencyHistoryCreate {
	fchc.mutation.SetFeedType(s)
	return fchc
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (fchc *FiatCurrencyHistoryCreate) SetNillableFeedType(s *string) *FiatCurrencyHistoryCreate {
	if s != nil {
		fchc.SetFeedType(*s)
	}
	return fchc
}

// SetMarketValueLow sets the "market_value_low" field.
func (fchc *FiatCurrencyHistoryCreate) SetMarketValueLow(d decimal.Decimal) *FiatCurrencyHistoryCreate {
	fchc.mutation.SetMarketValueLow(d)
	return fchc
}

// SetNillableMarketValueLow sets the "market_value_low" field if the given value is not nil.
func (fchc *FiatCurrencyHistoryCreate) SetNillableMarketValueLow(d *decimal.Decimal) *FiatCurrencyHistoryCreate {
	if d != nil {
		fchc.SetMarketValueLow(*d)
	}
	return fchc
}

// SetMarketValueHigh sets the "market_value_high" field.
func (fchc *FiatCurrencyHistoryCreate) SetMarketValueHigh(d decimal.Decimal) *FiatCurrencyHistoryCreate {
	fchc.mutation.SetMarketValueHigh(d)
	return fchc
}

// SetNillableMarketValueHigh sets the "market_value_high" field if the given value is not nil.
func (fchc *FiatCurrencyHistoryCreate) SetNillableMarketValueHigh(d *decimal.Decimal) *FiatCurrencyHistoryCreate {
	if d != nil {
		fchc.SetMarketValueHigh(*d)
	}
	return fchc
}

// SetID sets the "id" field.
func (fchc *FiatCurrencyHistoryCreate) SetID(u uuid.UUID) *FiatCurrencyHistoryCreate {
	fchc.mutation.SetID(u)
	return fchc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fchc *FiatCurrencyHistoryCreate) SetNillableID(u *uuid.UUID) *FiatCurrencyHistoryCreate {
	if u != nil {
		fchc.SetID(*u)
	}
	return fchc
}

// Mutation returns the FiatCurrencyHistoryMutation object of the builder.
func (fchc *FiatCurrencyHistoryCreate) Mutation() *FiatCurrencyHistoryMutation {
	return fchc.mutation
}

// Save creates the FiatCurrencyHistory in the database.
func (fchc *FiatCurrencyHistoryCreate) Save(ctx context.Context) (*FiatCurrencyHistory, error) {
	var (
		err  error
		node *FiatCurrencyHistory
	)
	if err := fchc.defaults(); err != nil {
		return nil, err
	}
	if len(fchc.hooks) == 0 {
		if err = fchc.check(); err != nil {
			return nil, err
		}
		node, err = fchc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FiatCurrencyHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fchc.check(); err != nil {
				return nil, err
			}
			fchc.mutation = mutation
			if node, err = fchc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fchc.hooks) - 1; i >= 0; i-- {
			if fchc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fchc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fchc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (fchc *FiatCurrencyHistoryCreate) SaveX(ctx context.Context) *FiatCurrencyHistory {
	v, err := fchc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fchc *FiatCurrencyHistoryCreate) Exec(ctx context.Context) error {
	_, err := fchc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fchc *FiatCurrencyHistoryCreate) ExecX(ctx context.Context) {
	if err := fchc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fchc *FiatCurrencyHistoryCreate) defaults() error {
	if _, ok := fchc.mutation.CreatedAt(); !ok {
		if fiatcurrencyhistory.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrencyhistory.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := fiatcurrencyhistory.DefaultCreatedAt()
		fchc.mutation.SetCreatedAt(v)
	}
	if _, ok := fchc.mutation.UpdatedAt(); !ok {
		if fiatcurrencyhistory.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrencyhistory.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fiatcurrencyhistory.DefaultUpdatedAt()
		fchc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fchc.mutation.DeletedAt(); !ok {
		if fiatcurrencyhistory.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrencyhistory.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := fiatcurrencyhistory.DefaultDeletedAt()
		fchc.mutation.SetDeletedAt(v)
	}
	if _, ok := fchc.mutation.FiatID(); !ok {
		if fiatcurrencyhistory.DefaultFiatID == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrencyhistory.DefaultFiatID (forgotten import ent/runtime?)")
		}
		v := fiatcurrencyhistory.DefaultFiatID()
		fchc.mutation.SetFiatID(v)
	}
	if _, ok := fchc.mutation.FeedType(); !ok {
		v := fiatcurrencyhistory.DefaultFeedType
		fchc.mutation.SetFeedType(v)
	}
	if _, ok := fchc.mutation.MarketValueLow(); !ok {
		v := fiatcurrencyhistory.DefaultMarketValueLow
		fchc.mutation.SetMarketValueLow(v)
	}
	if _, ok := fchc.mutation.MarketValueHigh(); !ok {
		v := fiatcurrencyhistory.DefaultMarketValueHigh
		fchc.mutation.SetMarketValueHigh(v)
	}
	if _, ok := fchc.mutation.ID(); !ok {
		if fiatcurrencyhistory.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized fiatcurrencyhistory.DefaultID (forgotten import ent/runtime?)")
		}
		v := fiatcurrencyhistory.DefaultID()
		fchc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (fchc *FiatCurrencyHistoryCreate) check() error {
	if _, ok := fchc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FiatCurrencyHistory.created_at"`)}
	}
	if _, ok := fchc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FiatCurrencyHistory.updated_at"`)}
	}
	if _, ok := fchc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "FiatCurrencyHistory.deleted_at"`)}
	}
	return nil
}

func (fchc *FiatCurrencyHistoryCreate) sqlSave(ctx context.Context) (*FiatCurrencyHistory, error) {
	_node, _spec := fchc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fchc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (fchc *FiatCurrencyHistoryCreate) createSpec() (*FiatCurrencyHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &FiatCurrencyHistory{config: fchc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fiatcurrencyhistory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fiatcurrencyhistory.FieldID,
			},
		}
	)
	_spec.OnConflict = fchc.conflict
	if id, ok := fchc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fchc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := fchc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := fchc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiatcurrencyhistory.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := fchc.mutation.FiatID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: fiatcurrencyhistory.FieldFiatID,
		})
		_node.FiatID = value
	}
	if value, ok := fchc.mutation.FeedType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fiatcurrencyhistory.FieldFeedType,
		})
		_node.FeedType = value
	}
	if value, ok := fchc.mutation.MarketValueLow(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrencyhistory.FieldMarketValueLow,
		})
		_node.MarketValueLow = value
	}
	if value, ok := fchc.mutation.MarketValueHigh(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: fiatcurrencyhistory.FieldMarketValueHigh,
		})
		_node.MarketValueHigh = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FiatCurrencyHistory.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FiatCurrencyHistoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (fchc *FiatCurrencyHistoryCreate) OnConflict(opts ...sql.ConflictOption) *FiatCurrencyHistoryUpsertOne {
	fchc.conflict = opts
	return &FiatCurrencyHistoryUpsertOne{
		create: fchc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FiatCurrencyHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fchc *FiatCurrencyHistoryCreate) OnConflictColumns(columns ...string) *FiatCurrencyHistoryUpsertOne {
	fchc.conflict = append(fchc.conflict, sql.ConflictColumns(columns...))
	return &FiatCurrencyHistoryUpsertOne{
		create: fchc,
	}
}

type (
	// FiatCurrencyHistoryUpsertOne is the builder for "upsert"-ing
	//  one FiatCurrencyHistory node.
	FiatCurrencyHistoryUpsertOne struct {
		create *FiatCurrencyHistoryCreate
	}

	// FiatCurrencyHistoryUpsert is the "OnConflict" setter.
	FiatCurrencyHistoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *FiatCurrencyHistoryUpsert) SetCreatedAt(v uint32) *FiatCurrencyHistoryUpsert {
	u.Set(fiatcurrencyhistory.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsert) UpdateCreatedAt() *FiatCurrencyHistoryUpsert {
	u.SetExcluded(fiatcurrencyhistory.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatCurrencyHistoryUpsert) AddCreatedAt(v uint32) *FiatCurrencyHistoryUpsert {
	u.Add(fiatcurrencyhistory.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatCurrencyHistoryUpsert) SetUpdatedAt(v uint32) *FiatCurrencyHistoryUpsert {
	u.Set(fiatcurrencyhistory.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsert) UpdateUpdatedAt() *FiatCurrencyHistoryUpsert {
	u.SetExcluded(fiatcurrencyhistory.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatCurrencyHistoryUpsert) AddUpdatedAt(v uint32) *FiatCurrencyHistoryUpsert {
	u.Add(fiatcurrencyhistory.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatCurrencyHistoryUpsert) SetDeletedAt(v uint32) *FiatCurrencyHistoryUpsert {
	u.Set(fiatcurrencyhistory.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsert) UpdateDeletedAt() *FiatCurrencyHistoryUpsert {
	u.SetExcluded(fiatcurrencyhistory.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatCurrencyHistoryUpsert) AddDeletedAt(v uint32) *FiatCurrencyHistoryUpsert {
	u.Add(fiatcurrencyhistory.FieldDeletedAt, v)
	return u
}

// SetFiatID sets the "fiat_id" field.
func (u *FiatCurrencyHistoryUpsert) SetFiatID(v uuid.UUID) *FiatCurrencyHistoryUpsert {
	u.Set(fiatcurrencyhistory.FieldFiatID, v)
	return u
}

// UpdateFiatID sets the "fiat_id" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsert) UpdateFiatID() *FiatCurrencyHistoryUpsert {
	u.SetExcluded(fiatcurrencyhistory.FieldFiatID)
	return u
}

// ClearFiatID clears the value of the "fiat_id" field.
func (u *FiatCurrencyHistoryUpsert) ClearFiatID() *FiatCurrencyHistoryUpsert {
	u.SetNull(fiatcurrencyhistory.FieldFiatID)
	return u
}

// SetFeedType sets the "feed_type" field.
func (u *FiatCurrencyHistoryUpsert) SetFeedType(v string) *FiatCurrencyHistoryUpsert {
	u.Set(fiatcurrencyhistory.FieldFeedType, v)
	return u
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsert) UpdateFeedType() *FiatCurrencyHistoryUpsert {
	u.SetExcluded(fiatcurrencyhistory.FieldFeedType)
	return u
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *FiatCurrencyHistoryUpsert) ClearFeedType() *FiatCurrencyHistoryUpsert {
	u.SetNull(fiatcurrencyhistory.FieldFeedType)
	return u
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *FiatCurrencyHistoryUpsert) SetMarketValueLow(v decimal.Decimal) *FiatCurrencyHistoryUpsert {
	u.Set(fiatcurrencyhistory.FieldMarketValueLow, v)
	return u
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsert) UpdateMarketValueLow() *FiatCurrencyHistoryUpsert {
	u.SetExcluded(fiatcurrencyhistory.FieldMarketValueLow)
	return u
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *FiatCurrencyHistoryUpsert) ClearMarketValueLow() *FiatCurrencyHistoryUpsert {
	u.SetNull(fiatcurrencyhistory.FieldMarketValueLow)
	return u
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *FiatCurrencyHistoryUpsert) SetMarketValueHigh(v decimal.Decimal) *FiatCurrencyHistoryUpsert {
	u.Set(fiatcurrencyhistory.FieldMarketValueHigh, v)
	return u
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsert) UpdateMarketValueHigh() *FiatCurrencyHistoryUpsert {
	u.SetExcluded(fiatcurrencyhistory.FieldMarketValueHigh)
	return u
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *FiatCurrencyHistoryUpsert) ClearMarketValueHigh() *FiatCurrencyHistoryUpsert {
	u.SetNull(fiatcurrencyhistory.FieldMarketValueHigh)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.FiatCurrencyHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(fiatcurrencyhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *FiatCurrencyHistoryUpsertOne) UpdateNewValues() *FiatCurrencyHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(fiatcurrencyhistory.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.FiatCurrencyHistory.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *FiatCurrencyHistoryUpsertOne) Ignore() *FiatCurrencyHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FiatCurrencyHistoryUpsertOne) DoNothing() *FiatCurrencyHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FiatCurrencyHistoryCreate.OnConflict
// documentation for more info.
func (u *FiatCurrencyHistoryUpsertOne) Update(set func(*FiatCurrencyHistoryUpsert)) *FiatCurrencyHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FiatCurrencyHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FiatCurrencyHistoryUpsertOne) SetCreatedAt(v uint32) *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatCurrencyHistoryUpsertOne) AddCreatedAt(v uint32) *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertOne) UpdateCreatedAt() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatCurrencyHistoryUpsertOne) SetUpdatedAt(v uint32) *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatCurrencyHistoryUpsertOne) AddUpdatedAt(v uint32) *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertOne) UpdateUpdatedAt() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatCurrencyHistoryUpsertOne) SetDeletedAt(v uint32) *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatCurrencyHistoryUpsertOne) AddDeletedAt(v uint32) *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertOne) UpdateDeletedAt() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetFiatID sets the "fiat_id" field.
func (u *FiatCurrencyHistoryUpsertOne) SetFiatID(v uuid.UUID) *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetFiatID(v)
	})
}

// UpdateFiatID sets the "fiat_id" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertOne) UpdateFiatID() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateFiatID()
	})
}

// ClearFiatID clears the value of the "fiat_id" field.
func (u *FiatCurrencyHistoryUpsertOne) ClearFiatID() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.ClearFiatID()
	})
}

// SetFeedType sets the "feed_type" field.
func (u *FiatCurrencyHistoryUpsertOne) SetFeedType(v string) *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetFeedType(v)
	})
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertOne) UpdateFeedType() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateFeedType()
	})
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *FiatCurrencyHistoryUpsertOne) ClearFeedType() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.ClearFeedType()
	})
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *FiatCurrencyHistoryUpsertOne) SetMarketValueLow(v decimal.Decimal) *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetMarketValueLow(v)
	})
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertOne) UpdateMarketValueLow() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateMarketValueLow()
	})
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *FiatCurrencyHistoryUpsertOne) ClearMarketValueLow() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.ClearMarketValueLow()
	})
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *FiatCurrencyHistoryUpsertOne) SetMarketValueHigh(v decimal.Decimal) *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetMarketValueHigh(v)
	})
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertOne) UpdateMarketValueHigh() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateMarketValueHigh()
	})
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *FiatCurrencyHistoryUpsertOne) ClearMarketValueHigh() *FiatCurrencyHistoryUpsertOne {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.ClearMarketValueHigh()
	})
}

// Exec executes the query.
func (u *FiatCurrencyHistoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FiatCurrencyHistoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FiatCurrencyHistoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FiatCurrencyHistoryUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: FiatCurrencyHistoryUpsertOne.ID is not supported by MySQL driver. Use FiatCurrencyHistoryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FiatCurrencyHistoryUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FiatCurrencyHistoryCreateBulk is the builder for creating many FiatCurrencyHistory entities in bulk.
type FiatCurrencyHistoryCreateBulk struct {
	config
	builders []*FiatCurrencyHistoryCreate
	conflict []sql.ConflictOption
}

// Save creates the FiatCurrencyHistory entities in the database.
func (fchcb *FiatCurrencyHistoryCreateBulk) Save(ctx context.Context) ([]*FiatCurrencyHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fchcb.builders))
	nodes := make([]*FiatCurrencyHistory, len(fchcb.builders))
	mutators := make([]Mutator, len(fchcb.builders))
	for i := range fchcb.builders {
		func(i int, root context.Context) {
			builder := fchcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FiatCurrencyHistoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fchcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fchcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fchcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fchcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fchcb *FiatCurrencyHistoryCreateBulk) SaveX(ctx context.Context) []*FiatCurrencyHistory {
	v, err := fchcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fchcb *FiatCurrencyHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := fchcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fchcb *FiatCurrencyHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := fchcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.FiatCurrencyHistory.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FiatCurrencyHistoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (fchcb *FiatCurrencyHistoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *FiatCurrencyHistoryUpsertBulk {
	fchcb.conflict = opts
	return &FiatCurrencyHistoryUpsertBulk{
		create: fchcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.FiatCurrencyHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fchcb *FiatCurrencyHistoryCreateBulk) OnConflictColumns(columns ...string) *FiatCurrencyHistoryUpsertBulk {
	fchcb.conflict = append(fchcb.conflict, sql.ConflictColumns(columns...))
	return &FiatCurrencyHistoryUpsertBulk{
		create: fchcb,
	}
}

// FiatCurrencyHistoryUpsertBulk is the builder for "upsert"-ing
// a bulk of FiatCurrencyHistory nodes.
type FiatCurrencyHistoryUpsertBulk struct {
	create *FiatCurrencyHistoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.FiatCurrencyHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(fiatcurrencyhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *FiatCurrencyHistoryUpsertBulk) UpdateNewValues() *FiatCurrencyHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(fiatcurrencyhistory.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.FiatCurrencyHistory.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *FiatCurrencyHistoryUpsertBulk) Ignore() *FiatCurrencyHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FiatCurrencyHistoryUpsertBulk) DoNothing() *FiatCurrencyHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FiatCurrencyHistoryCreateBulk.OnConflict
// documentation for more info.
func (u *FiatCurrencyHistoryUpsertBulk) Update(set func(*FiatCurrencyHistoryUpsert)) *FiatCurrencyHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FiatCurrencyHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FiatCurrencyHistoryUpsertBulk) SetCreatedAt(v uint32) *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatCurrencyHistoryUpsertBulk) AddCreatedAt(v uint32) *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertBulk) UpdateCreatedAt() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatCurrencyHistoryUpsertBulk) SetUpdatedAt(v uint32) *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatCurrencyHistoryUpsertBulk) AddUpdatedAt(v uint32) *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertBulk) UpdateUpdatedAt() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatCurrencyHistoryUpsertBulk) SetDeletedAt(v uint32) *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatCurrencyHistoryUpsertBulk) AddDeletedAt(v uint32) *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertBulk) UpdateDeletedAt() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetFiatID sets the "fiat_id" field.
func (u *FiatCurrencyHistoryUpsertBulk) SetFiatID(v uuid.UUID) *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetFiatID(v)
	})
}

// UpdateFiatID sets the "fiat_id" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertBulk) UpdateFiatID() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateFiatID()
	})
}

// ClearFiatID clears the value of the "fiat_id" field.
func (u *FiatCurrencyHistoryUpsertBulk) ClearFiatID() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.ClearFiatID()
	})
}

// SetFeedType sets the "feed_type" field.
func (u *FiatCurrencyHistoryUpsertBulk) SetFeedType(v string) *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetFeedType(v)
	})
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertBulk) UpdateFeedType() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateFeedType()
	})
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *FiatCurrencyHistoryUpsertBulk) ClearFeedType() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.ClearFeedType()
	})
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *FiatCurrencyHistoryUpsertBulk) SetMarketValueLow(v decimal.Decimal) *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetMarketValueLow(v)
	})
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertBulk) UpdateMarketValueLow() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateMarketValueLow()
	})
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *FiatCurrencyHistoryUpsertBulk) ClearMarketValueLow() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.ClearMarketValueLow()
	})
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *FiatCurrencyHistoryUpsertBulk) SetMarketValueHigh(v decimal.Decimal) *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.SetMarketValueHigh(v)
	})
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *FiatCurrencyHistoryUpsertBulk) UpdateMarketValueHigh() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.UpdateMarketValueHigh()
	})
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *FiatCurrencyHistoryUpsertBulk) ClearMarketValueHigh() *FiatCurrencyHistoryUpsertBulk {
	return u.Update(func(s *FiatCurrencyHistoryUpsert) {
		s.ClearMarketValueHigh()
	})
}

// Exec executes the query.
func (u *FiatCurrencyHistoryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FiatCurrencyHistoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FiatCurrencyHistoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FiatCurrencyHistoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}