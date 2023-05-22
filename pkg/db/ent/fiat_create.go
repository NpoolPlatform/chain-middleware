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
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/fiat"
	"github.com/google/uuid"
)

// FiatCreate is the builder for creating a Fiat entity.
type FiatCreate struct {
	config
	mutation *FiatMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (fc *FiatCreate) SetCreatedAt(u uint32) *FiatCreate {
	fc.mutation.SetCreatedAt(u)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FiatCreate) SetNillableCreatedAt(u *uint32) *FiatCreate {
	if u != nil {
		fc.SetCreatedAt(*u)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FiatCreate) SetUpdatedAt(u uint32) *FiatCreate {
	fc.mutation.SetUpdatedAt(u)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FiatCreate) SetNillableUpdatedAt(u *uint32) *FiatCreate {
	if u != nil {
		fc.SetUpdatedAt(*u)
	}
	return fc
}

// SetDeletedAt sets the "deleted_at" field.
func (fc *FiatCreate) SetDeletedAt(u uint32) *FiatCreate {
	fc.mutation.SetDeletedAt(u)
	return fc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fc *FiatCreate) SetNillableDeletedAt(u *uint32) *FiatCreate {
	if u != nil {
		fc.SetDeletedAt(*u)
	}
	return fc
}

// SetName sets the "name" field.
func (fc *FiatCreate) SetName(s string) *FiatCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fc *FiatCreate) SetNillableName(s *string) *FiatCreate {
	if s != nil {
		fc.SetName(*s)
	}
	return fc
}

// SetLogo sets the "logo" field.
func (fc *FiatCreate) SetLogo(s string) *FiatCreate {
	fc.mutation.SetLogo(s)
	return fc
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (fc *FiatCreate) SetNillableLogo(s *string) *FiatCreate {
	if s != nil {
		fc.SetLogo(*s)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FiatCreate) SetID(u uuid.UUID) *FiatCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FiatCreate) SetNillableID(u *uuid.UUID) *FiatCreate {
	if u != nil {
		fc.SetID(*u)
	}
	return fc
}

// Mutation returns the FiatMutation object of the builder.
func (fc *FiatCreate) Mutation() *FiatMutation {
	return fc.mutation
}

// Save creates the Fiat in the database.
func (fc *FiatCreate) Save(ctx context.Context) (*Fiat, error) {
	var (
		err  error
		node *Fiat
	)
	if err := fc.defaults(); err != nil {
		return nil, err
	}
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FiatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			if node, err = fc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			if fc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Fiat)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FiatMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FiatCreate) SaveX(ctx context.Context) *Fiat {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FiatCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FiatCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FiatCreate) defaults() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		if fiat.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized fiat.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := fiat.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		if fiat.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized fiat.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := fiat.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.DeletedAt(); !ok {
		if fiat.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized fiat.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := fiat.DefaultDeletedAt()
		fc.mutation.SetDeletedAt(v)
	}
	if _, ok := fc.mutation.Name(); !ok {
		v := fiat.DefaultName
		fc.mutation.SetName(v)
	}
	if _, ok := fc.mutation.Logo(); !ok {
		v := fiat.DefaultLogo
		fc.mutation.SetLogo(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		if fiat.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized fiat.DefaultID (forgotten import ent/runtime?)")
		}
		v := fiat.DefaultID()
		fc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (fc *FiatCreate) check() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Fiat.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Fiat.updated_at"`)}
	}
	if _, ok := fc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Fiat.deleted_at"`)}
	}
	return nil
}

func (fc *FiatCreate) sqlSave(ctx context.Context) (*Fiat, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
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

func (fc *FiatCreate) createSpec() (*Fiat, *sqlgraph.CreateSpec) {
	var (
		_node = &Fiat{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fiat.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fiat.FieldID,
			},
		}
	)
	_spec.OnConflict = fc.conflict
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiat.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiat.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: fiat.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fiat.FieldName,
		})
		_node.Name = value
	}
	if value, ok := fc.mutation.Logo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fiat.FieldLogo,
		})
		_node.Logo = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Fiat.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FiatUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (fc *FiatCreate) OnConflict(opts ...sql.ConflictOption) *FiatUpsertOne {
	fc.conflict = opts
	return &FiatUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Fiat.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fc *FiatCreate) OnConflictColumns(columns ...string) *FiatUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FiatUpsertOne{
		create: fc,
	}
}

type (
	// FiatUpsertOne is the builder for "upsert"-ing
	//  one Fiat node.
	FiatUpsertOne struct {
		create *FiatCreate
	}

	// FiatUpsert is the "OnConflict" setter.
	FiatUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *FiatUpsert) SetCreatedAt(v uint32) *FiatUpsert {
	u.Set(fiat.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatUpsert) UpdateCreatedAt() *FiatUpsert {
	u.SetExcluded(fiat.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatUpsert) AddCreatedAt(v uint32) *FiatUpsert {
	u.Add(fiat.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatUpsert) SetUpdatedAt(v uint32) *FiatUpsert {
	u.Set(fiat.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatUpsert) UpdateUpdatedAt() *FiatUpsert {
	u.SetExcluded(fiat.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatUpsert) AddUpdatedAt(v uint32) *FiatUpsert {
	u.Add(fiat.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatUpsert) SetDeletedAt(v uint32) *FiatUpsert {
	u.Set(fiat.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatUpsert) UpdateDeletedAt() *FiatUpsert {
	u.SetExcluded(fiat.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatUpsert) AddDeletedAt(v uint32) *FiatUpsert {
	u.Add(fiat.FieldDeletedAt, v)
	return u
}

// SetName sets the "name" field.
func (u *FiatUpsert) SetName(v string) *FiatUpsert {
	u.Set(fiat.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FiatUpsert) UpdateName() *FiatUpsert {
	u.SetExcluded(fiat.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *FiatUpsert) ClearName() *FiatUpsert {
	u.SetNull(fiat.FieldName)
	return u
}

// SetLogo sets the "logo" field.
func (u *FiatUpsert) SetLogo(v string) *FiatUpsert {
	u.Set(fiat.FieldLogo, v)
	return u
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *FiatUpsert) UpdateLogo() *FiatUpsert {
	u.SetExcluded(fiat.FieldLogo)
	return u
}

// ClearLogo clears the value of the "logo" field.
func (u *FiatUpsert) ClearLogo() *FiatUpsert {
	u.SetNull(fiat.FieldLogo)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Fiat.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(fiat.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *FiatUpsertOne) UpdateNewValues() *FiatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(fiat.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Fiat.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *FiatUpsertOne) Ignore() *FiatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FiatUpsertOne) DoNothing() *FiatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FiatCreate.OnConflict
// documentation for more info.
func (u *FiatUpsertOne) Update(set func(*FiatUpsert)) *FiatUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FiatUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FiatUpsertOne) SetCreatedAt(v uint32) *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatUpsertOne) AddCreatedAt(v uint32) *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatUpsertOne) UpdateCreatedAt() *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatUpsertOne) SetUpdatedAt(v uint32) *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatUpsertOne) AddUpdatedAt(v uint32) *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatUpsertOne) UpdateUpdatedAt() *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatUpsertOne) SetDeletedAt(v uint32) *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatUpsertOne) AddDeletedAt(v uint32) *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatUpsertOne) UpdateDeletedAt() *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *FiatUpsertOne) SetName(v string) *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FiatUpsertOne) UpdateName() *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *FiatUpsertOne) ClearName() *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.ClearName()
	})
}

// SetLogo sets the "logo" field.
func (u *FiatUpsertOne) SetLogo(v string) *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *FiatUpsertOne) UpdateLogo() *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.UpdateLogo()
	})
}

// ClearLogo clears the value of the "logo" field.
func (u *FiatUpsertOne) ClearLogo() *FiatUpsertOne {
	return u.Update(func(s *FiatUpsert) {
		s.ClearLogo()
	})
}

// Exec executes the query.
func (u *FiatUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FiatCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FiatUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FiatUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: FiatUpsertOne.ID is not supported by MySQL driver. Use FiatUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FiatUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FiatCreateBulk is the builder for creating many Fiat entities in bulk.
type FiatCreateBulk struct {
	config
	builders []*FiatCreate
	conflict []sql.ConflictOption
}

// Save creates the Fiat entities in the database.
func (fcb *FiatCreateBulk) Save(ctx context.Context) ([]*Fiat, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Fiat, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FiatMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FiatCreateBulk) SaveX(ctx context.Context) []*Fiat {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FiatCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FiatCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Fiat.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FiatUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (fcb *FiatCreateBulk) OnConflict(opts ...sql.ConflictOption) *FiatUpsertBulk {
	fcb.conflict = opts
	return &FiatUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Fiat.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (fcb *FiatCreateBulk) OnConflictColumns(columns ...string) *FiatUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FiatUpsertBulk{
		create: fcb,
	}
}

// FiatUpsertBulk is the builder for "upsert"-ing
// a bulk of Fiat nodes.
type FiatUpsertBulk struct {
	create *FiatCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Fiat.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(fiat.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *FiatUpsertBulk) UpdateNewValues() *FiatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(fiat.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Fiat.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *FiatUpsertBulk) Ignore() *FiatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FiatUpsertBulk) DoNothing() *FiatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FiatCreateBulk.OnConflict
// documentation for more info.
func (u *FiatUpsertBulk) Update(set func(*FiatUpsert)) *FiatUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FiatUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *FiatUpsertBulk) SetCreatedAt(v uint32) *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *FiatUpsertBulk) AddCreatedAt(v uint32) *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *FiatUpsertBulk) UpdateCreatedAt() *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FiatUpsertBulk) SetUpdatedAt(v uint32) *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *FiatUpsertBulk) AddUpdatedAt(v uint32) *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FiatUpsertBulk) UpdateUpdatedAt() *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FiatUpsertBulk) SetDeletedAt(v uint32) *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *FiatUpsertBulk) AddDeletedAt(v uint32) *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FiatUpsertBulk) UpdateDeletedAt() *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetName sets the "name" field.
func (u *FiatUpsertBulk) SetName(v string) *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FiatUpsertBulk) UpdateName() *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *FiatUpsertBulk) ClearName() *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.ClearName()
	})
}

// SetLogo sets the "logo" field.
func (u *FiatUpsertBulk) SetLogo(v string) *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *FiatUpsertBulk) UpdateLogo() *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.UpdateLogo()
	})
}

// ClearLogo clears the value of the "logo" field.
func (u *FiatUpsertBulk) ClearLogo() *FiatUpsertBulk {
	return u.Update(func(s *FiatUpsert) {
		s.ClearLogo()
	})
}

// Exec executes the query.
func (u *FiatUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FiatCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FiatCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FiatUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
