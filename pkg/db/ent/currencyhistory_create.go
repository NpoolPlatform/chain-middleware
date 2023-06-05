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
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/currencyhistory"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CurrencyHistoryCreate is the builder for creating a CurrencyHistory entity.
type CurrencyHistoryCreate struct {
	config
	mutation *CurrencyHistoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (chc *CurrencyHistoryCreate) SetCreatedAt(u uint32) *CurrencyHistoryCreate {
	chc.mutation.SetCreatedAt(u)
	return chc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (chc *CurrencyHistoryCreate) SetNillableCreatedAt(u *uint32) *CurrencyHistoryCreate {
	if u != nil {
		chc.SetCreatedAt(*u)
	}
	return chc
}

// SetUpdatedAt sets the "updated_at" field.
func (chc *CurrencyHistoryCreate) SetUpdatedAt(u uint32) *CurrencyHistoryCreate {
	chc.mutation.SetUpdatedAt(u)
	return chc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (chc *CurrencyHistoryCreate) SetNillableUpdatedAt(u *uint32) *CurrencyHistoryCreate {
	if u != nil {
		chc.SetUpdatedAt(*u)
	}
	return chc
}

// SetDeletedAt sets the "deleted_at" field.
func (chc *CurrencyHistoryCreate) SetDeletedAt(u uint32) *CurrencyHistoryCreate {
	chc.mutation.SetDeletedAt(u)
	return chc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (chc *CurrencyHistoryCreate) SetNillableDeletedAt(u *uint32) *CurrencyHistoryCreate {
	if u != nil {
		chc.SetDeletedAt(*u)
	}
	return chc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (chc *CurrencyHistoryCreate) SetCoinTypeID(u uuid.UUID) *CurrencyHistoryCreate {
	chc.mutation.SetCoinTypeID(u)
	return chc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (chc *CurrencyHistoryCreate) SetNillableCoinTypeID(u *uuid.UUID) *CurrencyHistoryCreate {
	if u != nil {
		chc.SetCoinTypeID(*u)
	}
	return chc
}

// SetFeedType sets the "feed_type" field.
func (chc *CurrencyHistoryCreate) SetFeedType(s string) *CurrencyHistoryCreate {
	chc.mutation.SetFeedType(s)
	return chc
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (chc *CurrencyHistoryCreate) SetNillableFeedType(s *string) *CurrencyHistoryCreate {
	if s != nil {
		chc.SetFeedType(*s)
	}
	return chc
}

// SetMarketValueHigh sets the "market_value_high" field.
func (chc *CurrencyHistoryCreate) SetMarketValueHigh(d decimal.Decimal) *CurrencyHistoryCreate {
	chc.mutation.SetMarketValueHigh(d)
	return chc
}

// SetNillableMarketValueHigh sets the "market_value_high" field if the given value is not nil.
func (chc *CurrencyHistoryCreate) SetNillableMarketValueHigh(d *decimal.Decimal) *CurrencyHistoryCreate {
	if d != nil {
		chc.SetMarketValueHigh(*d)
	}
	return chc
}

// SetMarketValueLow sets the "market_value_low" field.
func (chc *CurrencyHistoryCreate) SetMarketValueLow(d decimal.Decimal) *CurrencyHistoryCreate {
	chc.mutation.SetMarketValueLow(d)
	return chc
}

// SetNillableMarketValueLow sets the "market_value_low" field if the given value is not nil.
func (chc *CurrencyHistoryCreate) SetNillableMarketValueLow(d *decimal.Decimal) *CurrencyHistoryCreate {
	if d != nil {
		chc.SetMarketValueLow(*d)
	}
	return chc
}

// SetID sets the "id" field.
func (chc *CurrencyHistoryCreate) SetID(u uuid.UUID) *CurrencyHistoryCreate {
	chc.mutation.SetID(u)
	return chc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (chc *CurrencyHistoryCreate) SetNillableID(u *uuid.UUID) *CurrencyHistoryCreate {
	if u != nil {
		chc.SetID(*u)
	}
	return chc
}

// Mutation returns the CurrencyHistoryMutation object of the builder.
func (chc *CurrencyHistoryCreate) Mutation() *CurrencyHistoryMutation {
	return chc.mutation
}

// Save creates the CurrencyHistory in the database.
func (chc *CurrencyHistoryCreate) Save(ctx context.Context) (*CurrencyHistory, error) {
	var (
		err  error
		node *CurrencyHistory
	)
	if err := chc.defaults(); err != nil {
		return nil, err
	}
	if len(chc.hooks) == 0 {
		if err = chc.check(); err != nil {
			return nil, err
		}
		node, err = chc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CurrencyHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = chc.check(); err != nil {
				return nil, err
			}
			chc.mutation = mutation
			if node, err = chc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(chc.hooks) - 1; i >= 0; i-- {
			if chc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = chc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, chc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CurrencyHistory)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CurrencyHistoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (chc *CurrencyHistoryCreate) SaveX(ctx context.Context) *CurrencyHistory {
	v, err := chc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (chc *CurrencyHistoryCreate) Exec(ctx context.Context) error {
	_, err := chc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chc *CurrencyHistoryCreate) ExecX(ctx context.Context) {
	if err := chc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (chc *CurrencyHistoryCreate) defaults() error {
	if _, ok := chc.mutation.CreatedAt(); !ok {
		if currencyhistory.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyhistory.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := currencyhistory.DefaultCreatedAt()
		chc.mutation.SetCreatedAt(v)
	}
	if _, ok := chc.mutation.UpdatedAt(); !ok {
		if currencyhistory.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyhistory.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := currencyhistory.DefaultUpdatedAt()
		chc.mutation.SetUpdatedAt(v)
	}
	if _, ok := chc.mutation.DeletedAt(); !ok {
		if currencyhistory.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyhistory.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := currencyhistory.DefaultDeletedAt()
		chc.mutation.SetDeletedAt(v)
	}
	if _, ok := chc.mutation.CoinTypeID(); !ok {
		if currencyhistory.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized currencyhistory.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := currencyhistory.DefaultCoinTypeID()
		chc.mutation.SetCoinTypeID(v)
	}
	if _, ok := chc.mutation.FeedType(); !ok {
		v := currencyhistory.DefaultFeedType
		chc.mutation.SetFeedType(v)
	}
	if _, ok := chc.mutation.MarketValueHigh(); !ok {
		v := currencyhistory.DefaultMarketValueHigh
		chc.mutation.SetMarketValueHigh(v)
	}
	if _, ok := chc.mutation.MarketValueLow(); !ok {
		v := currencyhistory.DefaultMarketValueLow
		chc.mutation.SetMarketValueLow(v)
	}
	if _, ok := chc.mutation.ID(); !ok {
		if currencyhistory.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized currencyhistory.DefaultID (forgotten import ent/runtime?)")
		}
		v := currencyhistory.DefaultID()
		chc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (chc *CurrencyHistoryCreate) check() error {
	if _, ok := chc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CurrencyHistory.created_at"`)}
	}
	if _, ok := chc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CurrencyHistory.updated_at"`)}
	}
	if _, ok := chc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "CurrencyHistory.deleted_at"`)}
	}
	return nil
}

func (chc *CurrencyHistoryCreate) sqlSave(ctx context.Context) (*CurrencyHistory, error) {
	_node, _spec := chc.createSpec()
	if err := sqlgraph.CreateNode(ctx, chc.driver, _spec); err != nil {
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

func (chc *CurrencyHistoryCreate) createSpec() (*CurrencyHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &CurrencyHistory{config: chc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: currencyhistory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: currencyhistory.FieldID,
			},
		}
	)
	_spec.OnConflict = chc.conflict
	if id, ok := chc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := chc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyhistory.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := chc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyhistory.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := chc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyhistory.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := chc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: currencyhistory.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := chc.mutation.FeedType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: currencyhistory.FieldFeedType,
		})
		_node.FeedType = value
	}
	if value, ok := chc.mutation.MarketValueHigh(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: currencyhistory.FieldMarketValueHigh,
		})
		_node.MarketValueHigh = value
	}
	if value, ok := chc.mutation.MarketValueLow(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: currencyhistory.FieldMarketValueLow,
		})
		_node.MarketValueLow = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CurrencyHistory.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CurrencyHistoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (chc *CurrencyHistoryCreate) OnConflict(opts ...sql.ConflictOption) *CurrencyHistoryUpsertOne {
	chc.conflict = opts
	return &CurrencyHistoryUpsertOne{
		create: chc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CurrencyHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (chc *CurrencyHistoryCreate) OnConflictColumns(columns ...string) *CurrencyHistoryUpsertOne {
	chc.conflict = append(chc.conflict, sql.ConflictColumns(columns...))
	return &CurrencyHistoryUpsertOne{
		create: chc,
	}
}

type (
	// CurrencyHistoryUpsertOne is the builder for "upsert"-ing
	//  one CurrencyHistory node.
	CurrencyHistoryUpsertOne struct {
		create *CurrencyHistoryCreate
	}

	// CurrencyHistoryUpsert is the "OnConflict" setter.
	CurrencyHistoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyHistoryUpsert) SetCreatedAt(v uint32) *CurrencyHistoryUpsert {
	u.Set(currencyhistory.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyHistoryUpsert) UpdateCreatedAt() *CurrencyHistoryUpsert {
	u.SetExcluded(currencyhistory.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyHistoryUpsert) AddCreatedAt(v uint32) *CurrencyHistoryUpsert {
	u.Add(currencyhistory.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyHistoryUpsert) SetUpdatedAt(v uint32) *CurrencyHistoryUpsert {
	u.Set(currencyhistory.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyHistoryUpsert) UpdateUpdatedAt() *CurrencyHistoryUpsert {
	u.SetExcluded(currencyhistory.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyHistoryUpsert) AddUpdatedAt(v uint32) *CurrencyHistoryUpsert {
	u.Add(currencyhistory.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyHistoryUpsert) SetDeletedAt(v uint32) *CurrencyHistoryUpsert {
	u.Set(currencyhistory.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyHistoryUpsert) UpdateDeletedAt() *CurrencyHistoryUpsert {
	u.SetExcluded(currencyhistory.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyHistoryUpsert) AddDeletedAt(v uint32) *CurrencyHistoryUpsert {
	u.Add(currencyhistory.FieldDeletedAt, v)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyHistoryUpsert) SetCoinTypeID(v uuid.UUID) *CurrencyHistoryUpsert {
	u.Set(currencyhistory.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyHistoryUpsert) UpdateCoinTypeID() *CurrencyHistoryUpsert {
	u.SetExcluded(currencyhistory.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CurrencyHistoryUpsert) ClearCoinTypeID() *CurrencyHistoryUpsert {
	u.SetNull(currencyhistory.FieldCoinTypeID)
	return u
}

// SetFeedType sets the "feed_type" field.
func (u *CurrencyHistoryUpsert) SetFeedType(v string) *CurrencyHistoryUpsert {
	u.Set(currencyhistory.FieldFeedType, v)
	return u
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *CurrencyHistoryUpsert) UpdateFeedType() *CurrencyHistoryUpsert {
	u.SetExcluded(currencyhistory.FieldFeedType)
	return u
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *CurrencyHistoryUpsert) ClearFeedType() *CurrencyHistoryUpsert {
	u.SetNull(currencyhistory.FieldFeedType)
	return u
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *CurrencyHistoryUpsert) SetMarketValueHigh(v decimal.Decimal) *CurrencyHistoryUpsert {
	u.Set(currencyhistory.FieldMarketValueHigh, v)
	return u
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *CurrencyHistoryUpsert) UpdateMarketValueHigh() *CurrencyHistoryUpsert {
	u.SetExcluded(currencyhistory.FieldMarketValueHigh)
	return u
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *CurrencyHistoryUpsert) ClearMarketValueHigh() *CurrencyHistoryUpsert {
	u.SetNull(currencyhistory.FieldMarketValueHigh)
	return u
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *CurrencyHistoryUpsert) SetMarketValueLow(v decimal.Decimal) *CurrencyHistoryUpsert {
	u.Set(currencyhistory.FieldMarketValueLow, v)
	return u
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *CurrencyHistoryUpsert) UpdateMarketValueLow() *CurrencyHistoryUpsert {
	u.SetExcluded(currencyhistory.FieldMarketValueLow)
	return u
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *CurrencyHistoryUpsert) ClearMarketValueLow() *CurrencyHistoryUpsert {
	u.SetNull(currencyhistory.FieldMarketValueLow)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CurrencyHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(currencyhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CurrencyHistoryUpsertOne) UpdateNewValues() *CurrencyHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(currencyhistory.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.CurrencyHistory.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CurrencyHistoryUpsertOne) Ignore() *CurrencyHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CurrencyHistoryUpsertOne) DoNothing() *CurrencyHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CurrencyHistoryCreate.OnConflict
// documentation for more info.
func (u *CurrencyHistoryUpsertOne) Update(set func(*CurrencyHistoryUpsert)) *CurrencyHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CurrencyHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyHistoryUpsertOne) SetCreatedAt(v uint32) *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyHistoryUpsertOne) AddCreatedAt(v uint32) *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertOne) UpdateCreatedAt() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyHistoryUpsertOne) SetUpdatedAt(v uint32) *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyHistoryUpsertOne) AddUpdatedAt(v uint32) *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertOne) UpdateUpdatedAt() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyHistoryUpsertOne) SetDeletedAt(v uint32) *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyHistoryUpsertOne) AddDeletedAt(v uint32) *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertOne) UpdateDeletedAt() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyHistoryUpsertOne) SetCoinTypeID(v uuid.UUID) *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertOne) UpdateCoinTypeID() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CurrencyHistoryUpsertOne) ClearCoinTypeID() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetFeedType sets the "feed_type" field.
func (u *CurrencyHistoryUpsertOne) SetFeedType(v string) *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetFeedType(v)
	})
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertOne) UpdateFeedType() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateFeedType()
	})
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *CurrencyHistoryUpsertOne) ClearFeedType() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.ClearFeedType()
	})
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *CurrencyHistoryUpsertOne) SetMarketValueHigh(v decimal.Decimal) *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetMarketValueHigh(v)
	})
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertOne) UpdateMarketValueHigh() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateMarketValueHigh()
	})
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *CurrencyHistoryUpsertOne) ClearMarketValueHigh() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.ClearMarketValueHigh()
	})
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *CurrencyHistoryUpsertOne) SetMarketValueLow(v decimal.Decimal) *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetMarketValueLow(v)
	})
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertOne) UpdateMarketValueLow() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateMarketValueLow()
	})
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *CurrencyHistoryUpsertOne) ClearMarketValueLow() *CurrencyHistoryUpsertOne {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.ClearMarketValueLow()
	})
}

// Exec executes the query.
func (u *CurrencyHistoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CurrencyHistoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CurrencyHistoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CurrencyHistoryUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CurrencyHistoryUpsertOne.ID is not supported by MySQL driver. Use CurrencyHistoryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CurrencyHistoryUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CurrencyHistoryCreateBulk is the builder for creating many CurrencyHistory entities in bulk.
type CurrencyHistoryCreateBulk struct {
	config
	builders []*CurrencyHistoryCreate
	conflict []sql.ConflictOption
}

// Save creates the CurrencyHistory entities in the database.
func (chcb *CurrencyHistoryCreateBulk) Save(ctx context.Context) ([]*CurrencyHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(chcb.builders))
	nodes := make([]*CurrencyHistory, len(chcb.builders))
	mutators := make([]Mutator, len(chcb.builders))
	for i := range chcb.builders {
		func(i int, root context.Context) {
			builder := chcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CurrencyHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, chcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = chcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, chcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, chcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (chcb *CurrencyHistoryCreateBulk) SaveX(ctx context.Context) []*CurrencyHistory {
	v, err := chcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (chcb *CurrencyHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := chcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chcb *CurrencyHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := chcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CurrencyHistory.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CurrencyHistoryUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (chcb *CurrencyHistoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *CurrencyHistoryUpsertBulk {
	chcb.conflict = opts
	return &CurrencyHistoryUpsertBulk{
		create: chcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CurrencyHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (chcb *CurrencyHistoryCreateBulk) OnConflictColumns(columns ...string) *CurrencyHistoryUpsertBulk {
	chcb.conflict = append(chcb.conflict, sql.ConflictColumns(columns...))
	return &CurrencyHistoryUpsertBulk{
		create: chcb,
	}
}

// CurrencyHistoryUpsertBulk is the builder for "upsert"-ing
// a bulk of CurrencyHistory nodes.
type CurrencyHistoryUpsertBulk struct {
	create *CurrencyHistoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CurrencyHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(currencyhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CurrencyHistoryUpsertBulk) UpdateNewValues() *CurrencyHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(currencyhistory.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CurrencyHistory.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CurrencyHistoryUpsertBulk) Ignore() *CurrencyHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CurrencyHistoryUpsertBulk) DoNothing() *CurrencyHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CurrencyHistoryCreateBulk.OnConflict
// documentation for more info.
func (u *CurrencyHistoryUpsertBulk) Update(set func(*CurrencyHistoryUpsert)) *CurrencyHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CurrencyHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CurrencyHistoryUpsertBulk) SetCreatedAt(v uint32) *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CurrencyHistoryUpsertBulk) AddCreatedAt(v uint32) *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertBulk) UpdateCreatedAt() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CurrencyHistoryUpsertBulk) SetUpdatedAt(v uint32) *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CurrencyHistoryUpsertBulk) AddUpdatedAt(v uint32) *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertBulk) UpdateUpdatedAt() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CurrencyHistoryUpsertBulk) SetDeletedAt(v uint32) *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CurrencyHistoryUpsertBulk) AddDeletedAt(v uint32) *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertBulk) UpdateDeletedAt() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *CurrencyHistoryUpsertBulk) SetCoinTypeID(v uuid.UUID) *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertBulk) UpdateCoinTypeID() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *CurrencyHistoryUpsertBulk) ClearCoinTypeID() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetFeedType sets the "feed_type" field.
func (u *CurrencyHistoryUpsertBulk) SetFeedType(v string) *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetFeedType(v)
	})
}

// UpdateFeedType sets the "feed_type" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertBulk) UpdateFeedType() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateFeedType()
	})
}

// ClearFeedType clears the value of the "feed_type" field.
func (u *CurrencyHistoryUpsertBulk) ClearFeedType() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.ClearFeedType()
	})
}

// SetMarketValueHigh sets the "market_value_high" field.
func (u *CurrencyHistoryUpsertBulk) SetMarketValueHigh(v decimal.Decimal) *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetMarketValueHigh(v)
	})
}

// UpdateMarketValueHigh sets the "market_value_high" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertBulk) UpdateMarketValueHigh() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateMarketValueHigh()
	})
}

// ClearMarketValueHigh clears the value of the "market_value_high" field.
func (u *CurrencyHistoryUpsertBulk) ClearMarketValueHigh() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.ClearMarketValueHigh()
	})
}

// SetMarketValueLow sets the "market_value_low" field.
func (u *CurrencyHistoryUpsertBulk) SetMarketValueLow(v decimal.Decimal) *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.SetMarketValueLow(v)
	})
}

// UpdateMarketValueLow sets the "market_value_low" field to the value that was provided on create.
func (u *CurrencyHistoryUpsertBulk) UpdateMarketValueLow() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.UpdateMarketValueLow()
	})
}

// ClearMarketValueLow clears the value of the "market_value_low" field.
func (u *CurrencyHistoryUpsertBulk) ClearMarketValueLow() *CurrencyHistoryUpsertBulk {
	return u.Update(func(s *CurrencyHistoryUpsert) {
		s.ClearMarketValueLow()
	})
}

// Exec executes the query.
func (u *CurrencyHistoryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CurrencyHistoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CurrencyHistoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CurrencyHistoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
