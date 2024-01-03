// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinusedfor"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinUsedForUpdate is the builder for updating CoinUsedFor entities.
type CoinUsedForUpdate struct {
	config
	hooks     []Hook
	mutation  *CoinUsedForMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CoinUsedForUpdate builder.
func (cufu *CoinUsedForUpdate) Where(ps ...predicate.CoinUsedFor) *CoinUsedForUpdate {
	cufu.mutation.Where(ps...)
	return cufu
}

// SetCreatedAt sets the "created_at" field.
func (cufu *CoinUsedForUpdate) SetCreatedAt(u uint32) *CoinUsedForUpdate {
	cufu.mutation.ResetCreatedAt()
	cufu.mutation.SetCreatedAt(u)
	return cufu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cufu *CoinUsedForUpdate) SetNillableCreatedAt(u *uint32) *CoinUsedForUpdate {
	if u != nil {
		cufu.SetCreatedAt(*u)
	}
	return cufu
}

// AddCreatedAt adds u to the "created_at" field.
func (cufu *CoinUsedForUpdate) AddCreatedAt(u int32) *CoinUsedForUpdate {
	cufu.mutation.AddCreatedAt(u)
	return cufu
}

// SetUpdatedAt sets the "updated_at" field.
func (cufu *CoinUsedForUpdate) SetUpdatedAt(u uint32) *CoinUsedForUpdate {
	cufu.mutation.ResetUpdatedAt()
	cufu.mutation.SetUpdatedAt(u)
	return cufu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cufu *CoinUsedForUpdate) AddUpdatedAt(u int32) *CoinUsedForUpdate {
	cufu.mutation.AddUpdatedAt(u)
	return cufu
}

// SetDeletedAt sets the "deleted_at" field.
func (cufu *CoinUsedForUpdate) SetDeletedAt(u uint32) *CoinUsedForUpdate {
	cufu.mutation.ResetDeletedAt()
	cufu.mutation.SetDeletedAt(u)
	return cufu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cufu *CoinUsedForUpdate) SetNillableDeletedAt(u *uint32) *CoinUsedForUpdate {
	if u != nil {
		cufu.SetDeletedAt(*u)
	}
	return cufu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cufu *CoinUsedForUpdate) AddDeletedAt(u int32) *CoinUsedForUpdate {
	cufu.mutation.AddDeletedAt(u)
	return cufu
}

// SetEntID sets the "ent_id" field.
func (cufu *CoinUsedForUpdate) SetEntID(u uuid.UUID) *CoinUsedForUpdate {
	cufu.mutation.SetEntID(u)
	return cufu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cufu *CoinUsedForUpdate) SetNillableEntID(u *uuid.UUID) *CoinUsedForUpdate {
	if u != nil {
		cufu.SetEntID(*u)
	}
	return cufu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cufu *CoinUsedForUpdate) SetCoinTypeID(u uuid.UUID) *CoinUsedForUpdate {
	cufu.mutation.SetCoinTypeID(u)
	return cufu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cufu *CoinUsedForUpdate) SetNillableCoinTypeID(u *uuid.UUID) *CoinUsedForUpdate {
	if u != nil {
		cufu.SetCoinTypeID(*u)
	}
	return cufu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (cufu *CoinUsedForUpdate) ClearCoinTypeID() *CoinUsedForUpdate {
	cufu.mutation.ClearCoinTypeID()
	return cufu
}

// SetUsedFor sets the "used_for" field.
func (cufu *CoinUsedForUpdate) SetUsedFor(s string) *CoinUsedForUpdate {
	cufu.mutation.SetUsedFor(s)
	return cufu
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (cufu *CoinUsedForUpdate) SetNillableUsedFor(s *string) *CoinUsedForUpdate {
	if s != nil {
		cufu.SetUsedFor(*s)
	}
	return cufu
}

// ClearUsedFor clears the value of the "used_for" field.
func (cufu *CoinUsedForUpdate) ClearUsedFor() *CoinUsedForUpdate {
	cufu.mutation.ClearUsedFor()
	return cufu
}

// Mutation returns the CoinUsedForMutation object of the builder.
func (cufu *CoinUsedForUpdate) Mutation() *CoinUsedForMutation {
	return cufu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cufu *CoinUsedForUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cufu.defaults(); err != nil {
		return 0, err
	}
	if len(cufu.hooks) == 0 {
		affected, err = cufu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinUsedForMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cufu.mutation = mutation
			affected, err = cufu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cufu.hooks) - 1; i >= 0; i-- {
			if cufu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cufu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cufu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cufu *CoinUsedForUpdate) SaveX(ctx context.Context) int {
	affected, err := cufu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cufu *CoinUsedForUpdate) Exec(ctx context.Context) error {
	_, err := cufu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cufu *CoinUsedForUpdate) ExecX(ctx context.Context) {
	if err := cufu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cufu *CoinUsedForUpdate) defaults() error {
	if _, ok := cufu.mutation.UpdatedAt(); !ok {
		if coinusedfor.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinusedfor.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinusedfor.UpdateDefaultUpdatedAt()
		cufu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cufu *CoinUsedForUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinUsedForUpdate {
	cufu.modifiers = append(cufu.modifiers, modifiers...)
	return cufu
}

func (cufu *CoinUsedForUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinusedfor.Table,
			Columns: coinusedfor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coinusedfor.FieldID,
			},
		},
	}
	if ps := cufu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cufu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldCreatedAt,
		})
	}
	if value, ok := cufu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldCreatedAt,
		})
	}
	if value, ok := cufu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldUpdatedAt,
		})
	}
	if value, ok := cufu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldUpdatedAt,
		})
	}
	if value, ok := cufu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldDeletedAt,
		})
	}
	if value, ok := cufu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldDeletedAt,
		})
	}
	if value, ok := cufu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinusedfor.FieldEntID,
		})
	}
	if value, ok := cufu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinusedfor.FieldCoinTypeID,
		})
	}
	if cufu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinusedfor.FieldCoinTypeID,
		})
	}
	if value, ok := cufu.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinusedfor.FieldUsedFor,
		})
	}
	if cufu.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinusedfor.FieldUsedFor,
		})
	}
	_spec.Modifiers = cufu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cufu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinusedfor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CoinUsedForUpdateOne is the builder for updating a single CoinUsedFor entity.
type CoinUsedForUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CoinUsedForMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cufuo *CoinUsedForUpdateOne) SetCreatedAt(u uint32) *CoinUsedForUpdateOne {
	cufuo.mutation.ResetCreatedAt()
	cufuo.mutation.SetCreatedAt(u)
	return cufuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cufuo *CoinUsedForUpdateOne) SetNillableCreatedAt(u *uint32) *CoinUsedForUpdateOne {
	if u != nil {
		cufuo.SetCreatedAt(*u)
	}
	return cufuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cufuo *CoinUsedForUpdateOne) AddCreatedAt(u int32) *CoinUsedForUpdateOne {
	cufuo.mutation.AddCreatedAt(u)
	return cufuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cufuo *CoinUsedForUpdateOne) SetUpdatedAt(u uint32) *CoinUsedForUpdateOne {
	cufuo.mutation.ResetUpdatedAt()
	cufuo.mutation.SetUpdatedAt(u)
	return cufuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cufuo *CoinUsedForUpdateOne) AddUpdatedAt(u int32) *CoinUsedForUpdateOne {
	cufuo.mutation.AddUpdatedAt(u)
	return cufuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cufuo *CoinUsedForUpdateOne) SetDeletedAt(u uint32) *CoinUsedForUpdateOne {
	cufuo.mutation.ResetDeletedAt()
	cufuo.mutation.SetDeletedAt(u)
	return cufuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cufuo *CoinUsedForUpdateOne) SetNillableDeletedAt(u *uint32) *CoinUsedForUpdateOne {
	if u != nil {
		cufuo.SetDeletedAt(*u)
	}
	return cufuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cufuo *CoinUsedForUpdateOne) AddDeletedAt(u int32) *CoinUsedForUpdateOne {
	cufuo.mutation.AddDeletedAt(u)
	return cufuo
}

// SetEntID sets the "ent_id" field.
func (cufuo *CoinUsedForUpdateOne) SetEntID(u uuid.UUID) *CoinUsedForUpdateOne {
	cufuo.mutation.SetEntID(u)
	return cufuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cufuo *CoinUsedForUpdateOne) SetNillableEntID(u *uuid.UUID) *CoinUsedForUpdateOne {
	if u != nil {
		cufuo.SetEntID(*u)
	}
	return cufuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cufuo *CoinUsedForUpdateOne) SetCoinTypeID(u uuid.UUID) *CoinUsedForUpdateOne {
	cufuo.mutation.SetCoinTypeID(u)
	return cufuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cufuo *CoinUsedForUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *CoinUsedForUpdateOne {
	if u != nil {
		cufuo.SetCoinTypeID(*u)
	}
	return cufuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (cufuo *CoinUsedForUpdateOne) ClearCoinTypeID() *CoinUsedForUpdateOne {
	cufuo.mutation.ClearCoinTypeID()
	return cufuo
}

// SetUsedFor sets the "used_for" field.
func (cufuo *CoinUsedForUpdateOne) SetUsedFor(s string) *CoinUsedForUpdateOne {
	cufuo.mutation.SetUsedFor(s)
	return cufuo
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (cufuo *CoinUsedForUpdateOne) SetNillableUsedFor(s *string) *CoinUsedForUpdateOne {
	if s != nil {
		cufuo.SetUsedFor(*s)
	}
	return cufuo
}

// ClearUsedFor clears the value of the "used_for" field.
func (cufuo *CoinUsedForUpdateOne) ClearUsedFor() *CoinUsedForUpdateOne {
	cufuo.mutation.ClearUsedFor()
	return cufuo
}

// Mutation returns the CoinUsedForMutation object of the builder.
func (cufuo *CoinUsedForUpdateOne) Mutation() *CoinUsedForMutation {
	return cufuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cufuo *CoinUsedForUpdateOne) Select(field string, fields ...string) *CoinUsedForUpdateOne {
	cufuo.fields = append([]string{field}, fields...)
	return cufuo
}

// Save executes the query and returns the updated CoinUsedFor entity.
func (cufuo *CoinUsedForUpdateOne) Save(ctx context.Context) (*CoinUsedFor, error) {
	var (
		err  error
		node *CoinUsedFor
	)
	if err := cufuo.defaults(); err != nil {
		return nil, err
	}
	if len(cufuo.hooks) == 0 {
		node, err = cufuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinUsedForMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cufuo.mutation = mutation
			node, err = cufuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cufuo.hooks) - 1; i >= 0; i-- {
			if cufuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cufuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cufuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CoinUsedFor)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CoinUsedForMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cufuo *CoinUsedForUpdateOne) SaveX(ctx context.Context) *CoinUsedFor {
	node, err := cufuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cufuo *CoinUsedForUpdateOne) Exec(ctx context.Context) error {
	_, err := cufuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cufuo *CoinUsedForUpdateOne) ExecX(ctx context.Context) {
	if err := cufuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cufuo *CoinUsedForUpdateOne) defaults() error {
	if _, ok := cufuo.mutation.UpdatedAt(); !ok {
		if coinusedfor.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinusedfor.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinusedfor.UpdateDefaultUpdatedAt()
		cufuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cufuo *CoinUsedForUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinUsedForUpdateOne {
	cufuo.modifiers = append(cufuo.modifiers, modifiers...)
	return cufuo
}

func (cufuo *CoinUsedForUpdateOne) sqlSave(ctx context.Context) (_node *CoinUsedFor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinusedfor.Table,
			Columns: coinusedfor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coinusedfor.FieldID,
			},
		},
	}
	id, ok := cufuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinUsedFor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cufuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinusedfor.FieldID)
		for _, f := range fields {
			if !coinusedfor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coinusedfor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cufuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cufuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldCreatedAt,
		})
	}
	if value, ok := cufuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldCreatedAt,
		})
	}
	if value, ok := cufuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldUpdatedAt,
		})
	}
	if value, ok := cufuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldUpdatedAt,
		})
	}
	if value, ok := cufuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldDeletedAt,
		})
	}
	if value, ok := cufuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinusedfor.FieldDeletedAt,
		})
	}
	if value, ok := cufuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinusedfor.FieldEntID,
		})
	}
	if value, ok := cufuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinusedfor.FieldCoinTypeID,
		})
	}
	if cufuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinusedfor.FieldCoinTypeID,
		})
	}
	if value, ok := cufuo.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinusedfor.FieldUsedFor,
		})
	}
	if cufuo.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinusedfor.FieldUsedFor,
		})
	}
	_spec.Modifiers = cufuo.modifiers
	_node = &CoinUsedFor{config: cufuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cufuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinusedfor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
