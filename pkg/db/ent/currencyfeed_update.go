// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/currencyfeed"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CurrencyFeedUpdate is the builder for updating CurrencyFeed entities.
type CurrencyFeedUpdate struct {
	config
	hooks     []Hook
	mutation  *CurrencyFeedMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CurrencyFeedUpdate builder.
func (cfu *CurrencyFeedUpdate) Where(ps ...predicate.CurrencyFeed) *CurrencyFeedUpdate {
	cfu.mutation.Where(ps...)
	return cfu
}

// SetCreatedAt sets the "created_at" field.
func (cfu *CurrencyFeedUpdate) SetCreatedAt(u uint32) *CurrencyFeedUpdate {
	cfu.mutation.ResetCreatedAt()
	cfu.mutation.SetCreatedAt(u)
	return cfu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cfu *CurrencyFeedUpdate) SetNillableCreatedAt(u *uint32) *CurrencyFeedUpdate {
	if u != nil {
		cfu.SetCreatedAt(*u)
	}
	return cfu
}

// AddCreatedAt adds u to the "created_at" field.
func (cfu *CurrencyFeedUpdate) AddCreatedAt(u int32) *CurrencyFeedUpdate {
	cfu.mutation.AddCreatedAt(u)
	return cfu
}

// SetUpdatedAt sets the "updated_at" field.
func (cfu *CurrencyFeedUpdate) SetUpdatedAt(u uint32) *CurrencyFeedUpdate {
	cfu.mutation.ResetUpdatedAt()
	cfu.mutation.SetUpdatedAt(u)
	return cfu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cfu *CurrencyFeedUpdate) AddUpdatedAt(u int32) *CurrencyFeedUpdate {
	cfu.mutation.AddUpdatedAt(u)
	return cfu
}

// SetDeletedAt sets the "deleted_at" field.
func (cfu *CurrencyFeedUpdate) SetDeletedAt(u uint32) *CurrencyFeedUpdate {
	cfu.mutation.ResetDeletedAt()
	cfu.mutation.SetDeletedAt(u)
	return cfu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cfu *CurrencyFeedUpdate) SetNillableDeletedAt(u *uint32) *CurrencyFeedUpdate {
	if u != nil {
		cfu.SetDeletedAt(*u)
	}
	return cfu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cfu *CurrencyFeedUpdate) AddDeletedAt(u int32) *CurrencyFeedUpdate {
	cfu.mutation.AddDeletedAt(u)
	return cfu
}

// SetEntID sets the "ent_id" field.
func (cfu *CurrencyFeedUpdate) SetEntID(u uuid.UUID) *CurrencyFeedUpdate {
	cfu.mutation.SetEntID(u)
	return cfu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cfu *CurrencyFeedUpdate) SetNillableEntID(u *uuid.UUID) *CurrencyFeedUpdate {
	if u != nil {
		cfu.SetEntID(*u)
	}
	return cfu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cfu *CurrencyFeedUpdate) SetCoinTypeID(u uuid.UUID) *CurrencyFeedUpdate {
	cfu.mutation.SetCoinTypeID(u)
	return cfu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cfu *CurrencyFeedUpdate) SetNillableCoinTypeID(u *uuid.UUID) *CurrencyFeedUpdate {
	if u != nil {
		cfu.SetCoinTypeID(*u)
	}
	return cfu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (cfu *CurrencyFeedUpdate) ClearCoinTypeID() *CurrencyFeedUpdate {
	cfu.mutation.ClearCoinTypeID()
	return cfu
}

// SetFeedType sets the "feed_type" field.
func (cfu *CurrencyFeedUpdate) SetFeedType(s string) *CurrencyFeedUpdate {
	cfu.mutation.SetFeedType(s)
	return cfu
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (cfu *CurrencyFeedUpdate) SetNillableFeedType(s *string) *CurrencyFeedUpdate {
	if s != nil {
		cfu.SetFeedType(*s)
	}
	return cfu
}

// ClearFeedType clears the value of the "feed_type" field.
func (cfu *CurrencyFeedUpdate) ClearFeedType() *CurrencyFeedUpdate {
	cfu.mutation.ClearFeedType()
	return cfu
}

// SetFeedCoinName sets the "feed_coin_name" field.
func (cfu *CurrencyFeedUpdate) SetFeedCoinName(s string) *CurrencyFeedUpdate {
	cfu.mutation.SetFeedCoinName(s)
	return cfu
}

// SetNillableFeedCoinName sets the "feed_coin_name" field if the given value is not nil.
func (cfu *CurrencyFeedUpdate) SetNillableFeedCoinName(s *string) *CurrencyFeedUpdate {
	if s != nil {
		cfu.SetFeedCoinName(*s)
	}
	return cfu
}

// ClearFeedCoinName clears the value of the "feed_coin_name" field.
func (cfu *CurrencyFeedUpdate) ClearFeedCoinName() *CurrencyFeedUpdate {
	cfu.mutation.ClearFeedCoinName()
	return cfu
}

// SetDisabled sets the "disabled" field.
func (cfu *CurrencyFeedUpdate) SetDisabled(b bool) *CurrencyFeedUpdate {
	cfu.mutation.SetDisabled(b)
	return cfu
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (cfu *CurrencyFeedUpdate) SetNillableDisabled(b *bool) *CurrencyFeedUpdate {
	if b != nil {
		cfu.SetDisabled(*b)
	}
	return cfu
}

// ClearDisabled clears the value of the "disabled" field.
func (cfu *CurrencyFeedUpdate) ClearDisabled() *CurrencyFeedUpdate {
	cfu.mutation.ClearDisabled()
	return cfu
}

// Mutation returns the CurrencyFeedMutation object of the builder.
func (cfu *CurrencyFeedUpdate) Mutation() *CurrencyFeedMutation {
	return cfu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cfu *CurrencyFeedUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cfu.defaults(); err != nil {
		return 0, err
	}
	if len(cfu.hooks) == 0 {
		affected, err = cfu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CurrencyFeedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cfu.mutation = mutation
			affected, err = cfu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cfu.hooks) - 1; i >= 0; i-- {
			if cfu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cfu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cfu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cfu *CurrencyFeedUpdate) SaveX(ctx context.Context) int {
	affected, err := cfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cfu *CurrencyFeedUpdate) Exec(ctx context.Context) error {
	_, err := cfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cfu *CurrencyFeedUpdate) ExecX(ctx context.Context) {
	if err := cfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cfu *CurrencyFeedUpdate) defaults() error {
	if _, ok := cfu.mutation.UpdatedAt(); !ok {
		if currencyfeed.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyfeed.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := currencyfeed.UpdateDefaultUpdatedAt()
		cfu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cfu *CurrencyFeedUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CurrencyFeedUpdate {
	cfu.modifiers = append(cfu.modifiers, modifiers...)
	return cfu
}

func (cfu *CurrencyFeedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   currencyfeed.Table,
			Columns: currencyfeed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: currencyfeed.FieldID,
			},
		},
	}
	if ps := cfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cfu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldCreatedAt,
		})
	}
	if value, ok := cfu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldCreatedAt,
		})
	}
	if value, ok := cfu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldUpdatedAt,
		})
	}
	if value, ok := cfu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldUpdatedAt,
		})
	}
	if value, ok := cfu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldDeletedAt,
		})
	}
	if value, ok := cfu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldDeletedAt,
		})
	}
	if value, ok := cfu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: currencyfeed.FieldEntID,
		})
	}
	if value, ok := cfu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: currencyfeed.FieldCoinTypeID,
		})
	}
	if cfu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: currencyfeed.FieldCoinTypeID,
		})
	}
	if value, ok := cfu.mutation.FeedType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: currencyfeed.FieldFeedType,
		})
	}
	if cfu.mutation.FeedTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: currencyfeed.FieldFeedType,
		})
	}
	if value, ok := cfu.mutation.FeedCoinName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: currencyfeed.FieldFeedCoinName,
		})
	}
	if cfu.mutation.FeedCoinNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: currencyfeed.FieldFeedCoinName,
		})
	}
	if value, ok := cfu.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: currencyfeed.FieldDisabled,
		})
	}
	if cfu.mutation.DisabledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: currencyfeed.FieldDisabled,
		})
	}
	_spec.Modifiers = cfu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{currencyfeed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CurrencyFeedUpdateOne is the builder for updating a single CurrencyFeed entity.
type CurrencyFeedUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CurrencyFeedMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cfuo *CurrencyFeedUpdateOne) SetCreatedAt(u uint32) *CurrencyFeedUpdateOne {
	cfuo.mutation.ResetCreatedAt()
	cfuo.mutation.SetCreatedAt(u)
	return cfuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cfuo *CurrencyFeedUpdateOne) SetNillableCreatedAt(u *uint32) *CurrencyFeedUpdateOne {
	if u != nil {
		cfuo.SetCreatedAt(*u)
	}
	return cfuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cfuo *CurrencyFeedUpdateOne) AddCreatedAt(u int32) *CurrencyFeedUpdateOne {
	cfuo.mutation.AddCreatedAt(u)
	return cfuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cfuo *CurrencyFeedUpdateOne) SetUpdatedAt(u uint32) *CurrencyFeedUpdateOne {
	cfuo.mutation.ResetUpdatedAt()
	cfuo.mutation.SetUpdatedAt(u)
	return cfuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cfuo *CurrencyFeedUpdateOne) AddUpdatedAt(u int32) *CurrencyFeedUpdateOne {
	cfuo.mutation.AddUpdatedAt(u)
	return cfuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cfuo *CurrencyFeedUpdateOne) SetDeletedAt(u uint32) *CurrencyFeedUpdateOne {
	cfuo.mutation.ResetDeletedAt()
	cfuo.mutation.SetDeletedAt(u)
	return cfuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cfuo *CurrencyFeedUpdateOne) SetNillableDeletedAt(u *uint32) *CurrencyFeedUpdateOne {
	if u != nil {
		cfuo.SetDeletedAt(*u)
	}
	return cfuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cfuo *CurrencyFeedUpdateOne) AddDeletedAt(u int32) *CurrencyFeedUpdateOne {
	cfuo.mutation.AddDeletedAt(u)
	return cfuo
}

// SetEntID sets the "ent_id" field.
func (cfuo *CurrencyFeedUpdateOne) SetEntID(u uuid.UUID) *CurrencyFeedUpdateOne {
	cfuo.mutation.SetEntID(u)
	return cfuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cfuo *CurrencyFeedUpdateOne) SetNillableEntID(u *uuid.UUID) *CurrencyFeedUpdateOne {
	if u != nil {
		cfuo.SetEntID(*u)
	}
	return cfuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cfuo *CurrencyFeedUpdateOne) SetCoinTypeID(u uuid.UUID) *CurrencyFeedUpdateOne {
	cfuo.mutation.SetCoinTypeID(u)
	return cfuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cfuo *CurrencyFeedUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *CurrencyFeedUpdateOne {
	if u != nil {
		cfuo.SetCoinTypeID(*u)
	}
	return cfuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (cfuo *CurrencyFeedUpdateOne) ClearCoinTypeID() *CurrencyFeedUpdateOne {
	cfuo.mutation.ClearCoinTypeID()
	return cfuo
}

// SetFeedType sets the "feed_type" field.
func (cfuo *CurrencyFeedUpdateOne) SetFeedType(s string) *CurrencyFeedUpdateOne {
	cfuo.mutation.SetFeedType(s)
	return cfuo
}

// SetNillableFeedType sets the "feed_type" field if the given value is not nil.
func (cfuo *CurrencyFeedUpdateOne) SetNillableFeedType(s *string) *CurrencyFeedUpdateOne {
	if s != nil {
		cfuo.SetFeedType(*s)
	}
	return cfuo
}

// ClearFeedType clears the value of the "feed_type" field.
func (cfuo *CurrencyFeedUpdateOne) ClearFeedType() *CurrencyFeedUpdateOne {
	cfuo.mutation.ClearFeedType()
	return cfuo
}

// SetFeedCoinName sets the "feed_coin_name" field.
func (cfuo *CurrencyFeedUpdateOne) SetFeedCoinName(s string) *CurrencyFeedUpdateOne {
	cfuo.mutation.SetFeedCoinName(s)
	return cfuo
}

// SetNillableFeedCoinName sets the "feed_coin_name" field if the given value is not nil.
func (cfuo *CurrencyFeedUpdateOne) SetNillableFeedCoinName(s *string) *CurrencyFeedUpdateOne {
	if s != nil {
		cfuo.SetFeedCoinName(*s)
	}
	return cfuo
}

// ClearFeedCoinName clears the value of the "feed_coin_name" field.
func (cfuo *CurrencyFeedUpdateOne) ClearFeedCoinName() *CurrencyFeedUpdateOne {
	cfuo.mutation.ClearFeedCoinName()
	return cfuo
}

// SetDisabled sets the "disabled" field.
func (cfuo *CurrencyFeedUpdateOne) SetDisabled(b bool) *CurrencyFeedUpdateOne {
	cfuo.mutation.SetDisabled(b)
	return cfuo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (cfuo *CurrencyFeedUpdateOne) SetNillableDisabled(b *bool) *CurrencyFeedUpdateOne {
	if b != nil {
		cfuo.SetDisabled(*b)
	}
	return cfuo
}

// ClearDisabled clears the value of the "disabled" field.
func (cfuo *CurrencyFeedUpdateOne) ClearDisabled() *CurrencyFeedUpdateOne {
	cfuo.mutation.ClearDisabled()
	return cfuo
}

// Mutation returns the CurrencyFeedMutation object of the builder.
func (cfuo *CurrencyFeedUpdateOne) Mutation() *CurrencyFeedMutation {
	return cfuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cfuo *CurrencyFeedUpdateOne) Select(field string, fields ...string) *CurrencyFeedUpdateOne {
	cfuo.fields = append([]string{field}, fields...)
	return cfuo
}

// Save executes the query and returns the updated CurrencyFeed entity.
func (cfuo *CurrencyFeedUpdateOne) Save(ctx context.Context) (*CurrencyFeed, error) {
	var (
		err  error
		node *CurrencyFeed
	)
	if err := cfuo.defaults(); err != nil {
		return nil, err
	}
	if len(cfuo.hooks) == 0 {
		node, err = cfuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CurrencyFeedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cfuo.mutation = mutation
			node, err = cfuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cfuo.hooks) - 1; i >= 0; i-- {
			if cfuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cfuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cfuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CurrencyFeed)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CurrencyFeedMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cfuo *CurrencyFeedUpdateOne) SaveX(ctx context.Context) *CurrencyFeed {
	node, err := cfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cfuo *CurrencyFeedUpdateOne) Exec(ctx context.Context) error {
	_, err := cfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cfuo *CurrencyFeedUpdateOne) ExecX(ctx context.Context) {
	if err := cfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cfuo *CurrencyFeedUpdateOne) defaults() error {
	if _, ok := cfuo.mutation.UpdatedAt(); !ok {
		if currencyfeed.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized currencyfeed.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := currencyfeed.UpdateDefaultUpdatedAt()
		cfuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cfuo *CurrencyFeedUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CurrencyFeedUpdateOne {
	cfuo.modifiers = append(cfuo.modifiers, modifiers...)
	return cfuo
}

func (cfuo *CurrencyFeedUpdateOne) sqlSave(ctx context.Context) (_node *CurrencyFeed, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   currencyfeed.Table,
			Columns: currencyfeed.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: currencyfeed.FieldID,
			},
		},
	}
	id, ok := cfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CurrencyFeed.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, currencyfeed.FieldID)
		for _, f := range fields {
			if !currencyfeed.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != currencyfeed.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cfuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldCreatedAt,
		})
	}
	if value, ok := cfuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldCreatedAt,
		})
	}
	if value, ok := cfuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldUpdatedAt,
		})
	}
	if value, ok := cfuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldUpdatedAt,
		})
	}
	if value, ok := cfuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldDeletedAt,
		})
	}
	if value, ok := cfuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: currencyfeed.FieldDeletedAt,
		})
	}
	if value, ok := cfuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: currencyfeed.FieldEntID,
		})
	}
	if value, ok := cfuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: currencyfeed.FieldCoinTypeID,
		})
	}
	if cfuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: currencyfeed.FieldCoinTypeID,
		})
	}
	if value, ok := cfuo.mutation.FeedType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: currencyfeed.FieldFeedType,
		})
	}
	if cfuo.mutation.FeedTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: currencyfeed.FieldFeedType,
		})
	}
	if value, ok := cfuo.mutation.FeedCoinName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: currencyfeed.FieldFeedCoinName,
		})
	}
	if cfuo.mutation.FeedCoinNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: currencyfeed.FieldFeedCoinName,
		})
	}
	if value, ok := cfuo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: currencyfeed.FieldDisabled,
		})
	}
	if cfuo.mutation.DisabledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: currencyfeed.FieldDisabled,
		})
	}
	_spec.Modifiers = cfuo.modifiers
	_node = &CurrencyFeed{config: cfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{currencyfeed.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
