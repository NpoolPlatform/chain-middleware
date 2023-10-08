// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coindescription"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CoinDescriptionUpdate is the builder for updating CoinDescription entities.
type CoinDescriptionUpdate struct {
	config
	hooks     []Hook
	mutation  *CoinDescriptionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CoinDescriptionUpdate builder.
func (cdu *CoinDescriptionUpdate) Where(ps ...predicate.CoinDescription) *CoinDescriptionUpdate {
	cdu.mutation.Where(ps...)
	return cdu
}

// SetCreatedAt sets the "created_at" field.
func (cdu *CoinDescriptionUpdate) SetCreatedAt(u uint32) *CoinDescriptionUpdate {
	cdu.mutation.ResetCreatedAt()
	cdu.mutation.SetCreatedAt(u)
	return cdu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cdu *CoinDescriptionUpdate) SetNillableCreatedAt(u *uint32) *CoinDescriptionUpdate {
	if u != nil {
		cdu.SetCreatedAt(*u)
	}
	return cdu
}

// AddCreatedAt adds u to the "created_at" field.
func (cdu *CoinDescriptionUpdate) AddCreatedAt(u int32) *CoinDescriptionUpdate {
	cdu.mutation.AddCreatedAt(u)
	return cdu
}

// SetUpdatedAt sets the "updated_at" field.
func (cdu *CoinDescriptionUpdate) SetUpdatedAt(u uint32) *CoinDescriptionUpdate {
	cdu.mutation.ResetUpdatedAt()
	cdu.mutation.SetUpdatedAt(u)
	return cdu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cdu *CoinDescriptionUpdate) AddUpdatedAt(u int32) *CoinDescriptionUpdate {
	cdu.mutation.AddUpdatedAt(u)
	return cdu
}

// SetDeletedAt sets the "deleted_at" field.
func (cdu *CoinDescriptionUpdate) SetDeletedAt(u uint32) *CoinDescriptionUpdate {
	cdu.mutation.ResetDeletedAt()
	cdu.mutation.SetDeletedAt(u)
	return cdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cdu *CoinDescriptionUpdate) SetNillableDeletedAt(u *uint32) *CoinDescriptionUpdate {
	if u != nil {
		cdu.SetDeletedAt(*u)
	}
	return cdu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cdu *CoinDescriptionUpdate) AddDeletedAt(u int32) *CoinDescriptionUpdate {
	cdu.mutation.AddDeletedAt(u)
	return cdu
}

// SetEntID sets the "ent_id" field.
func (cdu *CoinDescriptionUpdate) SetEntID(u uuid.UUID) *CoinDescriptionUpdate {
	cdu.mutation.SetEntID(u)
	return cdu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cdu *CoinDescriptionUpdate) SetNillableEntID(u *uuid.UUID) *CoinDescriptionUpdate {
	if u != nil {
		cdu.SetEntID(*u)
	}
	return cdu
}

// SetAppID sets the "app_id" field.
func (cdu *CoinDescriptionUpdate) SetAppID(u uuid.UUID) *CoinDescriptionUpdate {
	cdu.mutation.SetAppID(u)
	return cdu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (cdu *CoinDescriptionUpdate) SetNillableAppID(u *uuid.UUID) *CoinDescriptionUpdate {
	if u != nil {
		cdu.SetAppID(*u)
	}
	return cdu
}

// ClearAppID clears the value of the "app_id" field.
func (cdu *CoinDescriptionUpdate) ClearAppID() *CoinDescriptionUpdate {
	cdu.mutation.ClearAppID()
	return cdu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cdu *CoinDescriptionUpdate) SetCoinTypeID(u uuid.UUID) *CoinDescriptionUpdate {
	cdu.mutation.SetCoinTypeID(u)
	return cdu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cdu *CoinDescriptionUpdate) SetNillableCoinTypeID(u *uuid.UUID) *CoinDescriptionUpdate {
	if u != nil {
		cdu.SetCoinTypeID(*u)
	}
	return cdu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (cdu *CoinDescriptionUpdate) ClearCoinTypeID() *CoinDescriptionUpdate {
	cdu.mutation.ClearCoinTypeID()
	return cdu
}

// SetUsedFor sets the "used_for" field.
func (cdu *CoinDescriptionUpdate) SetUsedFor(s string) *CoinDescriptionUpdate {
	cdu.mutation.SetUsedFor(s)
	return cdu
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (cdu *CoinDescriptionUpdate) SetNillableUsedFor(s *string) *CoinDescriptionUpdate {
	if s != nil {
		cdu.SetUsedFor(*s)
	}
	return cdu
}

// ClearUsedFor clears the value of the "used_for" field.
func (cdu *CoinDescriptionUpdate) ClearUsedFor() *CoinDescriptionUpdate {
	cdu.mutation.ClearUsedFor()
	return cdu
}

// SetTitle sets the "title" field.
func (cdu *CoinDescriptionUpdate) SetTitle(s string) *CoinDescriptionUpdate {
	cdu.mutation.SetTitle(s)
	return cdu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cdu *CoinDescriptionUpdate) SetNillableTitle(s *string) *CoinDescriptionUpdate {
	if s != nil {
		cdu.SetTitle(*s)
	}
	return cdu
}

// ClearTitle clears the value of the "title" field.
func (cdu *CoinDescriptionUpdate) ClearTitle() *CoinDescriptionUpdate {
	cdu.mutation.ClearTitle()
	return cdu
}

// SetMessage sets the "message" field.
func (cdu *CoinDescriptionUpdate) SetMessage(s string) *CoinDescriptionUpdate {
	cdu.mutation.SetMessage(s)
	return cdu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cdu *CoinDescriptionUpdate) SetNillableMessage(s *string) *CoinDescriptionUpdate {
	if s != nil {
		cdu.SetMessage(*s)
	}
	return cdu
}

// ClearMessage clears the value of the "message" field.
func (cdu *CoinDescriptionUpdate) ClearMessage() *CoinDescriptionUpdate {
	cdu.mutation.ClearMessage()
	return cdu
}

// Mutation returns the CoinDescriptionMutation object of the builder.
func (cdu *CoinDescriptionUpdate) Mutation() *CoinDescriptionMutation {
	return cdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cdu *CoinDescriptionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cdu.defaults(); err != nil {
		return 0, err
	}
	if len(cdu.hooks) == 0 {
		affected, err = cdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cdu.mutation = mutation
			affected, err = cdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cdu.hooks) - 1; i >= 0; i-- {
			if cdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cdu *CoinDescriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := cdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cdu *CoinDescriptionUpdate) Exec(ctx context.Context) error {
	_, err := cdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdu *CoinDescriptionUpdate) ExecX(ctx context.Context) {
	if err := cdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cdu *CoinDescriptionUpdate) defaults() error {
	if _, ok := cdu.mutation.UpdatedAt(); !ok {
		if coindescription.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coindescription.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coindescription.UpdateDefaultUpdatedAt()
		cdu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cdu *CoinDescriptionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinDescriptionUpdate {
	cdu.modifiers = append(cdu.modifiers, modifiers...)
	return cdu
}

func (cdu *CoinDescriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coindescription.Table,
			Columns: coindescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coindescription.FieldID,
			},
		},
	}
	if ps := cdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cdu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldCreatedAt,
		})
	}
	if value, ok := cdu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldCreatedAt,
		})
	}
	if value, ok := cdu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldUpdatedAt,
		})
	}
	if value, ok := cdu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldUpdatedAt,
		})
	}
	if value, ok := cdu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldDeletedAt,
		})
	}
	if value, ok := cdu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldDeletedAt,
		})
	}
	if value, ok := cdu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coindescription.FieldEntID,
		})
	}
	if value, ok := cdu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coindescription.FieldAppID,
		})
	}
	if cdu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coindescription.FieldAppID,
		})
	}
	if value, ok := cdu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coindescription.FieldCoinTypeID,
		})
	}
	if cdu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coindescription.FieldCoinTypeID,
		})
	}
	if value, ok := cdu.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coindescription.FieldUsedFor,
		})
	}
	if cdu.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coindescription.FieldUsedFor,
		})
	}
	if value, ok := cdu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coindescription.FieldTitle,
		})
	}
	if cdu.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coindescription.FieldTitle,
		})
	}
	if value, ok := cdu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coindescription.FieldMessage,
		})
	}
	if cdu.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coindescription.FieldMessage,
		})
	}
	_spec.Modifiers = cdu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coindescription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CoinDescriptionUpdateOne is the builder for updating a single CoinDescription entity.
type CoinDescriptionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CoinDescriptionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cduo *CoinDescriptionUpdateOne) SetCreatedAt(u uint32) *CoinDescriptionUpdateOne {
	cduo.mutation.ResetCreatedAt()
	cduo.mutation.SetCreatedAt(u)
	return cduo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cduo *CoinDescriptionUpdateOne) SetNillableCreatedAt(u *uint32) *CoinDescriptionUpdateOne {
	if u != nil {
		cduo.SetCreatedAt(*u)
	}
	return cduo
}

// AddCreatedAt adds u to the "created_at" field.
func (cduo *CoinDescriptionUpdateOne) AddCreatedAt(u int32) *CoinDescriptionUpdateOne {
	cduo.mutation.AddCreatedAt(u)
	return cduo
}

// SetUpdatedAt sets the "updated_at" field.
func (cduo *CoinDescriptionUpdateOne) SetUpdatedAt(u uint32) *CoinDescriptionUpdateOne {
	cduo.mutation.ResetUpdatedAt()
	cduo.mutation.SetUpdatedAt(u)
	return cduo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cduo *CoinDescriptionUpdateOne) AddUpdatedAt(u int32) *CoinDescriptionUpdateOne {
	cduo.mutation.AddUpdatedAt(u)
	return cduo
}

// SetDeletedAt sets the "deleted_at" field.
func (cduo *CoinDescriptionUpdateOne) SetDeletedAt(u uint32) *CoinDescriptionUpdateOne {
	cduo.mutation.ResetDeletedAt()
	cduo.mutation.SetDeletedAt(u)
	return cduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cduo *CoinDescriptionUpdateOne) SetNillableDeletedAt(u *uint32) *CoinDescriptionUpdateOne {
	if u != nil {
		cduo.SetDeletedAt(*u)
	}
	return cduo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cduo *CoinDescriptionUpdateOne) AddDeletedAt(u int32) *CoinDescriptionUpdateOne {
	cduo.mutation.AddDeletedAt(u)
	return cduo
}

// SetEntID sets the "ent_id" field.
func (cduo *CoinDescriptionUpdateOne) SetEntID(u uuid.UUID) *CoinDescriptionUpdateOne {
	cduo.mutation.SetEntID(u)
	return cduo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (cduo *CoinDescriptionUpdateOne) SetNillableEntID(u *uuid.UUID) *CoinDescriptionUpdateOne {
	if u != nil {
		cduo.SetEntID(*u)
	}
	return cduo
}

// SetAppID sets the "app_id" field.
func (cduo *CoinDescriptionUpdateOne) SetAppID(u uuid.UUID) *CoinDescriptionUpdateOne {
	cduo.mutation.SetAppID(u)
	return cduo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (cduo *CoinDescriptionUpdateOne) SetNillableAppID(u *uuid.UUID) *CoinDescriptionUpdateOne {
	if u != nil {
		cduo.SetAppID(*u)
	}
	return cduo
}

// ClearAppID clears the value of the "app_id" field.
func (cduo *CoinDescriptionUpdateOne) ClearAppID() *CoinDescriptionUpdateOne {
	cduo.mutation.ClearAppID()
	return cduo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (cduo *CoinDescriptionUpdateOne) SetCoinTypeID(u uuid.UUID) *CoinDescriptionUpdateOne {
	cduo.mutation.SetCoinTypeID(u)
	return cduo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (cduo *CoinDescriptionUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *CoinDescriptionUpdateOne {
	if u != nil {
		cduo.SetCoinTypeID(*u)
	}
	return cduo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (cduo *CoinDescriptionUpdateOne) ClearCoinTypeID() *CoinDescriptionUpdateOne {
	cduo.mutation.ClearCoinTypeID()
	return cduo
}

// SetUsedFor sets the "used_for" field.
func (cduo *CoinDescriptionUpdateOne) SetUsedFor(s string) *CoinDescriptionUpdateOne {
	cduo.mutation.SetUsedFor(s)
	return cduo
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (cduo *CoinDescriptionUpdateOne) SetNillableUsedFor(s *string) *CoinDescriptionUpdateOne {
	if s != nil {
		cduo.SetUsedFor(*s)
	}
	return cduo
}

// ClearUsedFor clears the value of the "used_for" field.
func (cduo *CoinDescriptionUpdateOne) ClearUsedFor() *CoinDescriptionUpdateOne {
	cduo.mutation.ClearUsedFor()
	return cduo
}

// SetTitle sets the "title" field.
func (cduo *CoinDescriptionUpdateOne) SetTitle(s string) *CoinDescriptionUpdateOne {
	cduo.mutation.SetTitle(s)
	return cduo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cduo *CoinDescriptionUpdateOne) SetNillableTitle(s *string) *CoinDescriptionUpdateOne {
	if s != nil {
		cduo.SetTitle(*s)
	}
	return cduo
}

// ClearTitle clears the value of the "title" field.
func (cduo *CoinDescriptionUpdateOne) ClearTitle() *CoinDescriptionUpdateOne {
	cduo.mutation.ClearTitle()
	return cduo
}

// SetMessage sets the "message" field.
func (cduo *CoinDescriptionUpdateOne) SetMessage(s string) *CoinDescriptionUpdateOne {
	cduo.mutation.SetMessage(s)
	return cduo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cduo *CoinDescriptionUpdateOne) SetNillableMessage(s *string) *CoinDescriptionUpdateOne {
	if s != nil {
		cduo.SetMessage(*s)
	}
	return cduo
}

// ClearMessage clears the value of the "message" field.
func (cduo *CoinDescriptionUpdateOne) ClearMessage() *CoinDescriptionUpdateOne {
	cduo.mutation.ClearMessage()
	return cduo
}

// Mutation returns the CoinDescriptionMutation object of the builder.
func (cduo *CoinDescriptionUpdateOne) Mutation() *CoinDescriptionMutation {
	return cduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cduo *CoinDescriptionUpdateOne) Select(field string, fields ...string) *CoinDescriptionUpdateOne {
	cduo.fields = append([]string{field}, fields...)
	return cduo
}

// Save executes the query and returns the updated CoinDescription entity.
func (cduo *CoinDescriptionUpdateOne) Save(ctx context.Context) (*CoinDescription, error) {
	var (
		err  error
		node *CoinDescription
	)
	if err := cduo.defaults(); err != nil {
		return nil, err
	}
	if len(cduo.hooks) == 0 {
		node, err = cduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cduo.mutation = mutation
			node, err = cduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cduo.hooks) - 1; i >= 0; i-- {
			if cduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CoinDescription)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CoinDescriptionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cduo *CoinDescriptionUpdateOne) SaveX(ctx context.Context) *CoinDescription {
	node, err := cduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cduo *CoinDescriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := cduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cduo *CoinDescriptionUpdateOne) ExecX(ctx context.Context) {
	if err := cduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cduo *CoinDescriptionUpdateOne) defaults() error {
	if _, ok := cduo.mutation.UpdatedAt(); !ok {
		if coindescription.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized coindescription.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := coindescription.UpdateDefaultUpdatedAt()
		cduo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cduo *CoinDescriptionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CoinDescriptionUpdateOne {
	cduo.modifiers = append(cduo.modifiers, modifiers...)
	return cduo
}

func (cduo *CoinDescriptionUpdateOne) sqlSave(ctx context.Context) (_node *CoinDescription, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coindescription.Table,
			Columns: coindescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coindescription.FieldID,
			},
		},
	}
	id, ok := cduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoinDescription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coindescription.FieldID)
		for _, f := range fields {
			if !coindescription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coindescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cduo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldCreatedAt,
		})
	}
	if value, ok := cduo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldCreatedAt,
		})
	}
	if value, ok := cduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldUpdatedAt,
		})
	}
	if value, ok := cduo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldUpdatedAt,
		})
	}
	if value, ok := cduo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldDeletedAt,
		})
	}
	if value, ok := cduo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: coindescription.FieldDeletedAt,
		})
	}
	if value, ok := cduo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coindescription.FieldEntID,
		})
	}
	if value, ok := cduo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coindescription.FieldAppID,
		})
	}
	if cduo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coindescription.FieldAppID,
		})
	}
	if value, ok := cduo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: coindescription.FieldCoinTypeID,
		})
	}
	if cduo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: coindescription.FieldCoinTypeID,
		})
	}
	if value, ok := cduo.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coindescription.FieldUsedFor,
		})
	}
	if cduo.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coindescription.FieldUsedFor,
		})
	}
	if value, ok := cduo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coindescription.FieldTitle,
		})
	}
	if cduo.mutation.TitleCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coindescription.FieldTitle,
		})
	}
	if value, ok := cduo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coindescription.FieldMessage,
		})
	}
	if cduo.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coindescription.FieldMessage,
		})
	}
	_spec.Modifiers = cduo.modifiers
	_node = &CoinDescription{config: cduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coindescription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
