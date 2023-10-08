// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinbase"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// CoinBaseUpdate is the builder for updating CoinBase entities.
type CoinBaseUpdate struct {
	config
	hooks     []Hook
	mutation  *CoinBaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CoinBaseUpdate builder.
func (cbu *CoinBaseUpdate) Where(ps ...predicate.CoinBase) *CoinBaseUpdate {
	cbu.mutation.Where(ps...)
	return cbu
}

// SetCreatedAt sets the "created_at" field.
func (cbu *CoinBaseUpdate) SetCreatedAt(u uint32) *CoinBaseUpdate {
	cbu.mutation.ResetCreatedAt()
	cbu.mutation.SetCreatedAt(u)
	return cbu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillableCreatedAt(u *uint32) *CoinBaseUpdate {
	if u != nil {
		cbu.SetCreatedAt(*u)
	}
	return cbu
}

// AddCreatedAt adds u to the "created_at" field.
func (cbu *CoinBaseUpdate) AddCreatedAt(u int32) *CoinBaseUpdate {
	cbu.mutation.AddCreatedAt(u)
	return cbu
}

// SetUpdatedAt sets the "updated_at" field.
func (cbu *CoinBaseUpdate) SetUpdatedAt(u uint32) *CoinBaseUpdate {
	cbu.mutation.ResetUpdatedAt()
	cbu.mutation.SetUpdatedAt(u)
	return cbu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cbu *CoinBaseUpdate) AddUpdatedAt(u int32) *CoinBaseUpdate {
	cbu.mutation.AddUpdatedAt(u)
	return cbu
}

// SetDeletedAt sets the "deleted_at" field.
func (cbu *CoinBaseUpdate) SetDeletedAt(u uint32) *CoinBaseUpdate {
	cbu.mutation.ResetDeletedAt()
	cbu.mutation.SetDeletedAt(u)
	return cbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillableDeletedAt(u *uint32) *CoinBaseUpdate {
	if u != nil {
		cbu.SetDeletedAt(*u)
	}
	return cbu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cbu *CoinBaseUpdate) AddDeletedAt(u int32) *CoinBaseUpdate {
	cbu.mutation.AddDeletedAt(u)
	return cbu
}

// SetEntID sets the "ent_id" field.
func (cbu *CoinBaseUpdate) SetEntID(u uuid.UUID) *CoinBaseUpdate {
	cbu.mutation.SetEntID(u)
	return cbu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillableEntID(u *uuid.UUID) *CoinBaseUpdate {
	if u != nil {
		cbu.SetEntID(*u)
	}
	return cbu
}

// SetName sets the "name" field.
func (cbu *CoinBaseUpdate) SetName(s string) *CoinBaseUpdate {
	cbu.mutation.SetName(s)
	return cbu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillableName(s *string) *CoinBaseUpdate {
	if s != nil {
		cbu.SetName(*s)
	}
	return cbu
}

// ClearName clears the value of the "name" field.
func (cbu *CoinBaseUpdate) ClearName() *CoinBaseUpdate {
	cbu.mutation.ClearName()
	return cbu
}

// SetLogo sets the "logo" field.
func (cbu *CoinBaseUpdate) SetLogo(s string) *CoinBaseUpdate {
	cbu.mutation.SetLogo(s)
	return cbu
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillableLogo(s *string) *CoinBaseUpdate {
	if s != nil {
		cbu.SetLogo(*s)
	}
	return cbu
}

// ClearLogo clears the value of the "logo" field.
func (cbu *CoinBaseUpdate) ClearLogo() *CoinBaseUpdate {
	cbu.mutation.ClearLogo()
	return cbu
}

// SetPresale sets the "presale" field.
func (cbu *CoinBaseUpdate) SetPresale(b bool) *CoinBaseUpdate {
	cbu.mutation.SetPresale(b)
	return cbu
}

// SetNillablePresale sets the "presale" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillablePresale(b *bool) *CoinBaseUpdate {
	if b != nil {
		cbu.SetPresale(*b)
	}
	return cbu
}

// ClearPresale clears the value of the "presale" field.
func (cbu *CoinBaseUpdate) ClearPresale() *CoinBaseUpdate {
	cbu.mutation.ClearPresale()
	return cbu
}

// SetUnit sets the "unit" field.
func (cbu *CoinBaseUpdate) SetUnit(s string) *CoinBaseUpdate {
	cbu.mutation.SetUnit(s)
	return cbu
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillableUnit(s *string) *CoinBaseUpdate {
	if s != nil {
		cbu.SetUnit(*s)
	}
	return cbu
}

// ClearUnit clears the value of the "unit" field.
func (cbu *CoinBaseUpdate) ClearUnit() *CoinBaseUpdate {
	cbu.mutation.ClearUnit()
	return cbu
}

// SetEnv sets the "env" field.
func (cbu *CoinBaseUpdate) SetEnv(s string) *CoinBaseUpdate {
	cbu.mutation.SetEnv(s)
	return cbu
}

// SetNillableEnv sets the "env" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillableEnv(s *string) *CoinBaseUpdate {
	if s != nil {
		cbu.SetEnv(*s)
	}
	return cbu
}

// ClearEnv clears the value of the "env" field.
func (cbu *CoinBaseUpdate) ClearEnv() *CoinBaseUpdate {
	cbu.mutation.ClearEnv()
	return cbu
}

// SetReservedAmount sets the "reserved_amount" field.
func (cbu *CoinBaseUpdate) SetReservedAmount(d decimal.Decimal) *CoinBaseUpdate {
	cbu.mutation.SetReservedAmount(d)
	return cbu
}

// SetNillableReservedAmount sets the "reserved_amount" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillableReservedAmount(d *decimal.Decimal) *CoinBaseUpdate {
	if d != nil {
		cbu.SetReservedAmount(*d)
	}
	return cbu
}

// ClearReservedAmount clears the value of the "reserved_amount" field.
func (cbu *CoinBaseUpdate) ClearReservedAmount() *CoinBaseUpdate {
	cbu.mutation.ClearReservedAmount()
	return cbu
}

// SetForPay sets the "for_pay" field.
func (cbu *CoinBaseUpdate) SetForPay(b bool) *CoinBaseUpdate {
	cbu.mutation.SetForPay(b)
	return cbu
}

// SetNillableForPay sets the "for_pay" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillableForPay(b *bool) *CoinBaseUpdate {
	if b != nil {
		cbu.SetForPay(*b)
	}
	return cbu
}

// ClearForPay clears the value of the "for_pay" field.
func (cbu *CoinBaseUpdate) ClearForPay() *CoinBaseUpdate {
	cbu.mutation.ClearForPay()
	return cbu
}

// SetDisabled sets the "disabled" field.
func (cbu *CoinBaseUpdate) SetDisabled(b bool) *CoinBaseUpdate {
	cbu.mutation.SetDisabled(b)
	return cbu
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (cbu *CoinBaseUpdate) SetNillableDisabled(b *bool) *CoinBaseUpdate {
	if b != nil {
		cbu.SetDisabled(*b)
	}
	return cbu
}

// ClearDisabled clears the value of the "disabled" field.
func (cbu *CoinBaseUpdate) ClearDisabled() *CoinBaseUpdate {
	cbu.mutation.ClearDisabled()
	return cbu
}

// Mutation returns the CoinBaseMutation object of the builder.
func (cbu *CoinBaseUpdate) Mutation() *CoinBaseMutation {
	return cbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cbu *CoinBaseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cbu.defaults(); err != nil {
		return 0, err
	}
	if len(cbu.hooks) == 0 {
		affected, err = cbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinBaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cbu.mutation = mutation
			affected, err = cbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cbu.hooks) - 1; i >= 0; i-- {
			if cbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cbu *CoinBaseUpdate) SaveX(ctx context.Context) int {
	affected, err := cbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cbu *CoinBaseUpdate) Exec(ctx context.Context) error {
	_, err := cbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbu *CoinBaseUpdate) ExecX(ctx context.Context) {
	if err := cbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cbu *CoinBaseUpdate) defaults() error {
	if _, ok := cbu.mutation.UpdatedAt(); !ok {
		if coinbase.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinbase.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinbase.UpdateDefaultUpdatedAt()
		cbu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cbu *CoinBaseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinBaseUpdate {
	cbu.modifiers = append(cbu.modifiers, modifiers...)
	return cbu
}

func (cbu *CoinBaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinbase.Table,
			Columns: coinbase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coinbase.FieldID,
			},
		},
	}
	if ps := cbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldCreatedAt,
		})
	}
	if value, ok := cbu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldCreatedAt,
		})
	}
	if value, ok := cbu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldUpdatedAt,
		})
	}
	if value, ok := cbu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldUpdatedAt,
		})
	}
	if value, ok := cbu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldDeletedAt,
		})
	}
	if value, ok := cbu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldDeletedAt,
		})
	}
	if value, ok := cbu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinbase.FieldEntID,
		})
	}
	if value, ok := cbu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinbase.FieldName,
		})
	}
	if cbu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinbase.FieldName,
		})
	}
	if value, ok := cbu.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinbase.FieldLogo,
		})
	}
	if cbu.mutation.LogoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinbase.FieldLogo,
		})
	}
	if value, ok := cbu.mutation.Presale(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coinbase.FieldPresale,
		})
	}
	if cbu.mutation.PresaleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: coinbase.FieldPresale,
		})
	}
	if value, ok := cbu.mutation.Unit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinbase.FieldUnit,
		})
	}
	if cbu.mutation.UnitCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinbase.FieldUnit,
		})
	}
	if value, ok := cbu.mutation.Env(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinbase.FieldEnv,
		})
	}
	if cbu.mutation.EnvCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinbase.FieldEnv,
		})
	}
	if value, ok := cbu.mutation.ReservedAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: coinbase.FieldReservedAmount,
		})
	}
	if cbu.mutation.ReservedAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: coinbase.FieldReservedAmount,
		})
	}
	if value, ok := cbu.mutation.ForPay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coinbase.FieldForPay,
		})
	}
	if cbu.mutation.ForPayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: coinbase.FieldForPay,
		})
	}
	if value, ok := cbu.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coinbase.FieldDisabled,
		})
	}
	if cbu.mutation.DisabledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: coinbase.FieldDisabled,
		})
	}
	_spec.Modifiers = cbu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinbase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CoinBaseUpdateOne is the builder for updating a single CoinBase entity.
type CoinBaseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CoinBaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cbuo *CoinBaseUpdateOne) SetCreatedAt(u uint32) *CoinBaseUpdateOne {
	cbuo.mutation.ResetCreatedAt()
	cbuo.mutation.SetCreatedAt(u)
	return cbuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillableCreatedAt(u *uint32) *CoinBaseUpdateOne {
	if u != nil {
		cbuo.SetCreatedAt(*u)
	}
	return cbuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cbuo *CoinBaseUpdateOne) AddCreatedAt(u int32) *CoinBaseUpdateOne {
	cbuo.mutation.AddCreatedAt(u)
	return cbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cbuo *CoinBaseUpdateOne) SetUpdatedAt(u uint32) *CoinBaseUpdateOne {
	cbuo.mutation.ResetUpdatedAt()
	cbuo.mutation.SetUpdatedAt(u)
	return cbuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cbuo *CoinBaseUpdateOne) AddUpdatedAt(u int32) *CoinBaseUpdateOne {
	cbuo.mutation.AddUpdatedAt(u)
	return cbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cbuo *CoinBaseUpdateOne) SetDeletedAt(u uint32) *CoinBaseUpdateOne {
	cbuo.mutation.ResetDeletedAt()
	cbuo.mutation.SetDeletedAt(u)
	return cbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillableDeletedAt(u *uint32) *CoinBaseUpdateOne {
	if u != nil {
		cbuo.SetDeletedAt(*u)
	}
	return cbuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cbuo *CoinBaseUpdateOne) AddDeletedAt(u int32) *CoinBaseUpdateOne {
	cbuo.mutation.AddDeletedAt(u)
	return cbuo
}

// SetEntID sets the "ent_id" field.
func (cbuo *CoinBaseUpdateOne) SetEntID(u uuid.UUID) *CoinBaseUpdateOne {
	cbuo.mutation.SetEntID(u)
	return cbuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillableEntID(u *uuid.UUID) *CoinBaseUpdateOne {
	if u != nil {
		cbuo.SetEntID(*u)
	}
	return cbuo
}

// SetName sets the "name" field.
func (cbuo *CoinBaseUpdateOne) SetName(s string) *CoinBaseUpdateOne {
	cbuo.mutation.SetName(s)
	return cbuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillableName(s *string) *CoinBaseUpdateOne {
	if s != nil {
		cbuo.SetName(*s)
	}
	return cbuo
}

// ClearName clears the value of the "name" field.
func (cbuo *CoinBaseUpdateOne) ClearName() *CoinBaseUpdateOne {
	cbuo.mutation.ClearName()
	return cbuo
}

// SetLogo sets the "logo" field.
func (cbuo *CoinBaseUpdateOne) SetLogo(s string) *CoinBaseUpdateOne {
	cbuo.mutation.SetLogo(s)
	return cbuo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillableLogo(s *string) *CoinBaseUpdateOne {
	if s != nil {
		cbuo.SetLogo(*s)
	}
	return cbuo
}

// ClearLogo clears the value of the "logo" field.
func (cbuo *CoinBaseUpdateOne) ClearLogo() *CoinBaseUpdateOne {
	cbuo.mutation.ClearLogo()
	return cbuo
}

// SetPresale sets the "presale" field.
func (cbuo *CoinBaseUpdateOne) SetPresale(b bool) *CoinBaseUpdateOne {
	cbuo.mutation.SetPresale(b)
	return cbuo
}

// SetNillablePresale sets the "presale" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillablePresale(b *bool) *CoinBaseUpdateOne {
	if b != nil {
		cbuo.SetPresale(*b)
	}
	return cbuo
}

// ClearPresale clears the value of the "presale" field.
func (cbuo *CoinBaseUpdateOne) ClearPresale() *CoinBaseUpdateOne {
	cbuo.mutation.ClearPresale()
	return cbuo
}

// SetUnit sets the "unit" field.
func (cbuo *CoinBaseUpdateOne) SetUnit(s string) *CoinBaseUpdateOne {
	cbuo.mutation.SetUnit(s)
	return cbuo
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillableUnit(s *string) *CoinBaseUpdateOne {
	if s != nil {
		cbuo.SetUnit(*s)
	}
	return cbuo
}

// ClearUnit clears the value of the "unit" field.
func (cbuo *CoinBaseUpdateOne) ClearUnit() *CoinBaseUpdateOne {
	cbuo.mutation.ClearUnit()
	return cbuo
}

// SetEnv sets the "env" field.
func (cbuo *CoinBaseUpdateOne) SetEnv(s string) *CoinBaseUpdateOne {
	cbuo.mutation.SetEnv(s)
	return cbuo
}

// SetNillableEnv sets the "env" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillableEnv(s *string) *CoinBaseUpdateOne {
	if s != nil {
		cbuo.SetEnv(*s)
	}
	return cbuo
}

// ClearEnv clears the value of the "env" field.
func (cbuo *CoinBaseUpdateOne) ClearEnv() *CoinBaseUpdateOne {
	cbuo.mutation.ClearEnv()
	return cbuo
}

// SetReservedAmount sets the "reserved_amount" field.
func (cbuo *CoinBaseUpdateOne) SetReservedAmount(d decimal.Decimal) *CoinBaseUpdateOne {
	cbuo.mutation.SetReservedAmount(d)
	return cbuo
}

// SetNillableReservedAmount sets the "reserved_amount" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillableReservedAmount(d *decimal.Decimal) *CoinBaseUpdateOne {
	if d != nil {
		cbuo.SetReservedAmount(*d)
	}
	return cbuo
}

// ClearReservedAmount clears the value of the "reserved_amount" field.
func (cbuo *CoinBaseUpdateOne) ClearReservedAmount() *CoinBaseUpdateOne {
	cbuo.mutation.ClearReservedAmount()
	return cbuo
}

// SetForPay sets the "for_pay" field.
func (cbuo *CoinBaseUpdateOne) SetForPay(b bool) *CoinBaseUpdateOne {
	cbuo.mutation.SetForPay(b)
	return cbuo
}

// SetNillableForPay sets the "for_pay" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillableForPay(b *bool) *CoinBaseUpdateOne {
	if b != nil {
		cbuo.SetForPay(*b)
	}
	return cbuo
}

// ClearForPay clears the value of the "for_pay" field.
func (cbuo *CoinBaseUpdateOne) ClearForPay() *CoinBaseUpdateOne {
	cbuo.mutation.ClearForPay()
	return cbuo
}

// SetDisabled sets the "disabled" field.
func (cbuo *CoinBaseUpdateOne) SetDisabled(b bool) *CoinBaseUpdateOne {
	cbuo.mutation.SetDisabled(b)
	return cbuo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (cbuo *CoinBaseUpdateOne) SetNillableDisabled(b *bool) *CoinBaseUpdateOne {
	if b != nil {
		cbuo.SetDisabled(*b)
	}
	return cbuo
}

// ClearDisabled clears the value of the "disabled" field.
func (cbuo *CoinBaseUpdateOne) ClearDisabled() *CoinBaseUpdateOne {
	cbuo.mutation.ClearDisabled()
	return cbuo
}

// Mutation returns the CoinBaseMutation object of the builder.
func (cbuo *CoinBaseUpdateOne) Mutation() *CoinBaseMutation {
	return cbuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cbuo *CoinBaseUpdateOne) Select(field string, fields ...string) *CoinBaseUpdateOne {
	cbuo.fields = append([]string{field}, fields...)
	return cbuo
}

// Save executes the query and returns the updated CoinBase entity.
func (cbuo *CoinBaseUpdateOne) Save(ctx context.Context) (*CoinBase, error) {
	var (
		err  error
		node *CoinBase
	)
	if err := cbuo.defaults(); err != nil {
		return nil, err
	}
	if len(cbuo.hooks) == 0 {
		node, err = cbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinBaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cbuo.mutation = mutation
			node, err = cbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cbuo.hooks) - 1; i >= 0; i-- {
			if cbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cbuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CoinBase)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CoinBaseMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cbuo *CoinBaseUpdateOne) SaveX(ctx context.Context) *CoinBase {
	node, err := cbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cbuo *CoinBaseUpdateOne) Exec(ctx context.Context) error {
	_, err := cbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbuo *CoinBaseUpdateOne) ExecX(ctx context.Context) {
	if err := cbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cbuo *CoinBaseUpdateOne) defaults() error {
	if _, ok := cbuo.mutation.UpdatedAt(); !ok {
		if coinbase.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coinbase.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coinbase.UpdateDefaultUpdatedAt()
		cbuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cbuo *CoinBaseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinBaseUpdateOne {
	cbuo.modifiers = append(cbuo.modifiers, modifiers...)
	return cbuo
}

func (cbuo *CoinBaseUpdateOne) sqlSave(ctx context.Context) (_node *CoinBase, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coinbase.Table,
			Columns: coinbase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: coinbase.FieldID,
			},
		},
	}
	id, ok := cbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinBase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coinbase.FieldID)
		for _, f := range fields {
			if !coinbase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coinbase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldCreatedAt,
		})
	}
	if value, ok := cbuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldCreatedAt,
		})
	}
	if value, ok := cbuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldUpdatedAt,
		})
	}
	if value, ok := cbuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldUpdatedAt,
		})
	}
	if value, ok := cbuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldDeletedAt,
		})
	}
	if value, ok := cbuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coinbase.FieldDeletedAt,
		})
	}
	if value, ok := cbuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coinbase.FieldEntID,
		})
	}
	if value, ok := cbuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinbase.FieldName,
		})
	}
	if cbuo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinbase.FieldName,
		})
	}
	if value, ok := cbuo.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinbase.FieldLogo,
		})
	}
	if cbuo.mutation.LogoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinbase.FieldLogo,
		})
	}
	if value, ok := cbuo.mutation.Presale(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coinbase.FieldPresale,
		})
	}
	if cbuo.mutation.PresaleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: coinbase.FieldPresale,
		})
	}
	if value, ok := cbuo.mutation.Unit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinbase.FieldUnit,
		})
	}
	if cbuo.mutation.UnitCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinbase.FieldUnit,
		})
	}
	if value, ok := cbuo.mutation.Env(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coinbase.FieldEnv,
		})
	}
	if cbuo.mutation.EnvCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coinbase.FieldEnv,
		})
	}
	if value, ok := cbuo.mutation.ReservedAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: coinbase.FieldReservedAmount,
		})
	}
	if cbuo.mutation.ReservedAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: coinbase.FieldReservedAmount,
		})
	}
	if value, ok := cbuo.mutation.ForPay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coinbase.FieldForPay,
		})
	}
	if cbuo.mutation.ForPayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: coinbase.FieldForPay,
		})
	}
	if value, ok := cbuo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: coinbase.FieldDisabled,
		})
	}
	if cbuo.mutation.DisabledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: coinbase.FieldDisabled,
		})
	}
	_spec.Modifiers = cbuo.modifiers
	_node = &CoinBase{config: cbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coinbase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
