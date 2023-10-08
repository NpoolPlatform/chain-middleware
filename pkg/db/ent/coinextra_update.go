// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinextra"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinExtraUpdate is the builder for updating CoinExtra entities.
type CoinExtraUpdate struct {
	config
	hooks     []Hook
	mutation  *CoinExtraMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CoinExtraUpdate builder.
func (ceu *CoinExtraUpdate) Where(ps ...predicate.CoinExtra) *CoinExtraUpdate {
	ceu.mutation.Where(ps...)
	return ceu
}

// SetCreatedAt sets the "created_at" field.
func (ceu *CoinExtraUpdate) SetCreatedAt(u uint32) *CoinExtraUpdate {
	ceu.mutation.ResetCreatedAt()
	ceu.mutation.SetCreatedAt(u)
	return ceu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ceu *CoinExtraUpdate) SetNillableCreatedAt(u *uint32) *CoinExtraUpdate {
	if u != nil {
		ceu.SetCreatedAt(*u)
	}
	return ceu
}

// AddCreatedAt adds u to the "created_at" field.
func (ceu *CoinExtraUpdate) AddCreatedAt(u int32) *CoinExtraUpdate {
	ceu.mutation.AddCreatedAt(u)
	return ceu
}

// SetUpdatedAt sets the "updated_at" field.
func (ceu *CoinExtraUpdate) SetUpdatedAt(u uint32) *CoinExtraUpdate {
	ceu.mutation.ResetUpdatedAt()
	ceu.mutation.SetUpdatedAt(u)
	return ceu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ceu *CoinExtraUpdate) AddUpdatedAt(u int32) *CoinExtraUpdate {
	ceu.mutation.AddUpdatedAt(u)
	return ceu
}

// SetDeletedAt sets the "deleted_at" field.
func (ceu *CoinExtraUpdate) SetDeletedAt(u uint32) *CoinExtraUpdate {
	ceu.mutation.ResetDeletedAt()
	ceu.mutation.SetDeletedAt(u)
	return ceu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ceu *CoinExtraUpdate) SetNillableDeletedAt(u *uint32) *CoinExtraUpdate {
	if u != nil {
		ceu.SetDeletedAt(*u)
	}
	return ceu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ceu *CoinExtraUpdate) AddDeletedAt(u int32) *CoinExtraUpdate {
	ceu.mutation.AddDeletedAt(u)
	return ceu
}

// SetEntID sets the "ent_id" field.
func (ceu *CoinExtraUpdate) SetEntID(u uuid.UUID) *CoinExtraUpdate {
	ceu.mutation.SetEntID(u)
	return ceu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ceu *CoinExtraUpdate) SetNillableEntID(u *uuid.UUID) *CoinExtraUpdate {
	if u != nil {
		ceu.SetEntID(*u)
	}
	return ceu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ceu *CoinExtraUpdate) SetCoinTypeID(u uuid.UUID) *CoinExtraUpdate {
	ceu.mutation.SetCoinTypeID(u)
	return ceu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (ceu *CoinExtraUpdate) SetNillableCoinTypeID(u *uuid.UUID) *CoinExtraUpdate {
	if u != nil {
		ceu.SetCoinTypeID(*u)
	}
	return ceu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (ceu *CoinExtraUpdate) ClearCoinTypeID() *CoinExtraUpdate {
	ceu.mutation.ClearCoinTypeID()
	return ceu
}

// SetHomePage sets the "home_page" field.
func (ceu *CoinExtraUpdate) SetHomePage(s string) *CoinExtraUpdate {
	ceu.mutation.SetHomePage(s)
	return ceu
}

// SetNillableHomePage sets the "home_page" field if the given value is not nil.
func (ceu *CoinExtraUpdate) SetNillableHomePage(s *string) *CoinExtraUpdate {
	if s != nil {
		ceu.SetHomePage(*s)
	}
	return ceu
}

// ClearHomePage clears the value of the "home_page" field.
func (ceu *CoinExtraUpdate) ClearHomePage() *CoinExtraUpdate {
	ceu.mutation.ClearHomePage()
	return ceu
}

// SetSpecs sets the "specs" field.
func (ceu *CoinExtraUpdate) SetSpecs(s string) *CoinExtraUpdate {
	ceu.mutation.SetSpecs(s)
	return ceu
}

// SetNillableSpecs sets the "specs" field if the given value is not nil.
func (ceu *CoinExtraUpdate) SetNillableSpecs(s *string) *CoinExtraUpdate {
	if s != nil {
		ceu.SetSpecs(*s)
	}
	return ceu
}

// ClearSpecs clears the value of the "specs" field.
func (ceu *CoinExtraUpdate) ClearSpecs() *CoinExtraUpdate {
	ceu.mutation.ClearSpecs()
	return ceu
}

// SetStableUsd sets the "stable_usd" field.
func (ceu *CoinExtraUpdate) SetStableUsd(b bool) *CoinExtraUpdate {
	ceu.mutation.SetStableUsd(b)
	return ceu
}

// SetNillableStableUsd sets the "stable_usd" field if the given value is not nil.
func (ceu *CoinExtraUpdate) SetNillableStableUsd(b *bool) *CoinExtraUpdate {
	if b != nil {
		ceu.SetStableUsd(*b)
	}
	return ceu
}

// ClearStableUsd clears the value of the "stable_usd" field.
func (ceu *CoinExtraUpdate) ClearStableUsd() *CoinExtraUpdate {
	ceu.mutation.ClearStableUsd()
	return ceu
}

// Mutation returns the CoinExtraMutation object of the builder.
func (ceu *CoinExtraUpdate) Mutation() *CoinExtraMutation {
	return ceu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ceu *CoinExtraUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ceu.defaults(); err != nil {
		return 0, err
	}
	if len(ceu.hooks) == 0 {
		affected, err = ceu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinExtraMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ceu.mutation = mutation
			affected, err = ceu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ceu.hooks) - 1; i >= 0; i-- {
			if ceu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ceu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ceu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ceu *CoinExtraUpdate) SaveX(ctx context.Context) int {
	affected, err := ceu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ceu *CoinExtraUpdate) Exec(ctx context.Context) error {
	_, err := ceu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceu *CoinExtraUpdate) ExecX(ctx context.Context) {
	if err := ceu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ceu *CoinExtraUpdate) defaults() error {
	if _, ok := ceu.mutation.UpdatedAt(); !ok {
		if coinextra.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinextra.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinextra.UpdateDefaultUpdatedAt()
		ceu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ceu *CoinExtraUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinExtraUpdate {
	ceu.modifiers = append(ceu.modifiers, modifiers...)
	return ceu
}

func (ceu *CoinExtraUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinextra.Table,
			Columns: coinextra.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coinextra.FieldID,
			},
		},
	}
	if ps := ceu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldCreatedAt,
		})
	}
	if value, ok := ceu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldCreatedAt,
		})
	}
	if value, ok := ceu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldUpdatedAt,
		})
	}
	if value, ok := ceu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldUpdatedAt,
		})
	}
	if value, ok := ceu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldDeletedAt,
		})
	}
	if value, ok := ceu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldDeletedAt,
		})
	}
	if value, ok := ceu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinextra.FieldEntID,
		})
	}
	if value, ok := ceu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinextra.FieldCoinTypeID,
		})
	}
	if ceu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinextra.FieldCoinTypeID,
		})
	}
	if value, ok := ceu.mutation.HomePage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinextra.FieldHomePage,
		})
	}
	if ceu.mutation.HomePageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinextra.FieldHomePage,
		})
	}
	if value, ok := ceu.mutation.Specs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinextra.FieldSpecs,
		})
	}
	if ceu.mutation.SpecsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinextra.FieldSpecs,
		})
	}
	if value, ok := ceu.mutation.StableUsd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coinextra.FieldStableUsd,
		})
	}
	if ceu.mutation.StableUsdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: coinextra.FieldStableUsd,
		})
	}
	_spec.Modifiers = ceu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ceu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinextra.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CoinExtraUpdateOne is the builder for updating a single CoinExtra entity.
type CoinExtraUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CoinExtraMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ceuo *CoinExtraUpdateOne) SetCreatedAt(u uint32) *CoinExtraUpdateOne {
	ceuo.mutation.ResetCreatedAt()
	ceuo.mutation.SetCreatedAt(u)
	return ceuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ceuo *CoinExtraUpdateOne) SetNillableCreatedAt(u *uint32) *CoinExtraUpdateOne {
	if u != nil {
		ceuo.SetCreatedAt(*u)
	}
	return ceuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ceuo *CoinExtraUpdateOne) AddCreatedAt(u int32) *CoinExtraUpdateOne {
	ceuo.mutation.AddCreatedAt(u)
	return ceuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ceuo *CoinExtraUpdateOne) SetUpdatedAt(u uint32) *CoinExtraUpdateOne {
	ceuo.mutation.ResetUpdatedAt()
	ceuo.mutation.SetUpdatedAt(u)
	return ceuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ceuo *CoinExtraUpdateOne) AddUpdatedAt(u int32) *CoinExtraUpdateOne {
	ceuo.mutation.AddUpdatedAt(u)
	return ceuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ceuo *CoinExtraUpdateOne) SetDeletedAt(u uint32) *CoinExtraUpdateOne {
	ceuo.mutation.ResetDeletedAt()
	ceuo.mutation.SetDeletedAt(u)
	return ceuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ceuo *CoinExtraUpdateOne) SetNillableDeletedAt(u *uint32) *CoinExtraUpdateOne {
	if u != nil {
		ceuo.SetDeletedAt(*u)
	}
	return ceuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ceuo *CoinExtraUpdateOne) AddDeletedAt(u int32) *CoinExtraUpdateOne {
	ceuo.mutation.AddDeletedAt(u)
	return ceuo
}

// SetEntID sets the "ent_id" field.
func (ceuo *CoinExtraUpdateOne) SetEntID(u uuid.UUID) *CoinExtraUpdateOne {
	ceuo.mutation.SetEntID(u)
	return ceuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ceuo *CoinExtraUpdateOne) SetNillableEntID(u *uuid.UUID) *CoinExtraUpdateOne {
	if u != nil {
		ceuo.SetEntID(*u)
	}
	return ceuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (ceuo *CoinExtraUpdateOne) SetCoinTypeID(u uuid.UUID) *CoinExtraUpdateOne {
	ceuo.mutation.SetCoinTypeID(u)
	return ceuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (ceuo *CoinExtraUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *CoinExtraUpdateOne {
	if u != nil {
		ceuo.SetCoinTypeID(*u)
	}
	return ceuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (ceuo *CoinExtraUpdateOne) ClearCoinTypeID() *CoinExtraUpdateOne {
	ceuo.mutation.ClearCoinTypeID()
	return ceuo
}

// SetHomePage sets the "home_page" field.
func (ceuo *CoinExtraUpdateOne) SetHomePage(s string) *CoinExtraUpdateOne {
	ceuo.mutation.SetHomePage(s)
	return ceuo
}

// SetNillableHomePage sets the "home_page" field if the given value is not nil.
func (ceuo *CoinExtraUpdateOne) SetNillableHomePage(s *string) *CoinExtraUpdateOne {
	if s != nil {
		ceuo.SetHomePage(*s)
	}
	return ceuo
}

// ClearHomePage clears the value of the "home_page" field.
func (ceuo *CoinExtraUpdateOne) ClearHomePage() *CoinExtraUpdateOne {
	ceuo.mutation.ClearHomePage()
	return ceuo
}

// SetSpecs sets the "specs" field.
func (ceuo *CoinExtraUpdateOne) SetSpecs(s string) *CoinExtraUpdateOne {
	ceuo.mutation.SetSpecs(s)
	return ceuo
}

// SetNillableSpecs sets the "specs" field if the given value is not nil.
func (ceuo *CoinExtraUpdateOne) SetNillableSpecs(s *string) *CoinExtraUpdateOne {
	if s != nil {
		ceuo.SetSpecs(*s)
	}
	return ceuo
}

// ClearSpecs clears the value of the "specs" field.
func (ceuo *CoinExtraUpdateOne) ClearSpecs() *CoinExtraUpdateOne {
	ceuo.mutation.ClearSpecs()
	return ceuo
}

// SetStableUsd sets the "stable_usd" field.
func (ceuo *CoinExtraUpdateOne) SetStableUsd(b bool) *CoinExtraUpdateOne {
	ceuo.mutation.SetStableUsd(b)
	return ceuo
}

// SetNillableStableUsd sets the "stable_usd" field if the given value is not nil.
func (ceuo *CoinExtraUpdateOne) SetNillableStableUsd(b *bool) *CoinExtraUpdateOne {
	if b != nil {
		ceuo.SetStableUsd(*b)
	}
	return ceuo
}

// ClearStableUsd clears the value of the "stable_usd" field.
func (ceuo *CoinExtraUpdateOne) ClearStableUsd() *CoinExtraUpdateOne {
	ceuo.mutation.ClearStableUsd()
	return ceuo
}

// Mutation returns the CoinExtraMutation object of the builder.
func (ceuo *CoinExtraUpdateOne) Mutation() *CoinExtraMutation {
	return ceuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ceuo *CoinExtraUpdateOne) Select(field string, fields ...string) *CoinExtraUpdateOne {
	ceuo.fields = append([]string{field}, fields...)
	return ceuo
}

// Save executes the query and returns the updated CoinExtra entity.
func (ceuo *CoinExtraUpdateOne) Save(ctx context.Context) (*CoinExtra, error) {
	var (
		err  error
		node *CoinExtra
	)
	if err := ceuo.defaults(); err != nil {
		return nil, err
	}
	if len(ceuo.hooks) == 0 {
		node, err = ceuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinExtraMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ceuo.mutation = mutation
			node, err = ceuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ceuo.hooks) - 1; i >= 0; i-- {
			if ceuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ceuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ceuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CoinExtra)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CoinExtraMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ceuo *CoinExtraUpdateOne) SaveX(ctx context.Context) *CoinExtra {
	node, err := ceuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ceuo *CoinExtraUpdateOne) Exec(ctx context.Context) error {
	_, err := ceuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ceuo *CoinExtraUpdateOne) ExecX(ctx context.Context) {
	if err := ceuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ceuo *CoinExtraUpdateOne) defaults() error {
	if _, ok := ceuo.mutation.UpdatedAt(); !ok {
		if coinextra.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinextra.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinextra.UpdateDefaultUpdatedAt()
		ceuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ceuo *CoinExtraUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinExtraUpdateOne {
	ceuo.modifiers = append(ceuo.modifiers, modifiers...)
	return ceuo
}

func (ceuo *CoinExtraUpdateOne) sqlSave(ctx context.Context) (_node *CoinExtra, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinextra.Table,
			Columns: coinextra.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coinextra.FieldID,
			},
		},
	}
	id, ok := ceuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinExtra.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ceuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinextra.FieldID)
		for _, f := range fields {
			if !coinextra.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coinextra.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ceuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ceuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldCreatedAt,
		})
	}
	if value, ok := ceuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldCreatedAt,
		})
	}
	if value, ok := ceuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldUpdatedAt,
		})
	}
	if value, ok := ceuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldUpdatedAt,
		})
	}
	if value, ok := ceuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldDeletedAt,
		})
	}
	if value, ok := ceuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinextra.FieldDeletedAt,
		})
	}
	if value, ok := ceuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinextra.FieldEntID,
		})
	}
	if value, ok := ceuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinextra.FieldCoinTypeID,
		})
	}
	if ceuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coinextra.FieldCoinTypeID,
		})
	}
	if value, ok := ceuo.mutation.HomePage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinextra.FieldHomePage,
		})
	}
	if ceuo.mutation.HomePageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinextra.FieldHomePage,
		})
	}
	if value, ok := ceuo.mutation.Specs(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinextra.FieldSpecs,
		})
	}
	if ceuo.mutation.SpecsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinextra.FieldSpecs,
		})
	}
	if value, ok := ceuo.mutation.StableUsd(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coinextra.FieldStableUsd,
		})
	}
	if ceuo.mutation.StableUsdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: coinextra.FieldStableUsd,
		})
	}
	_spec.Modifiers = ceuo.modifiers
	_node = &CoinExtra{config: ceuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ceuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinextra.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
