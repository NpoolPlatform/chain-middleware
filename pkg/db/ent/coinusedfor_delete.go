// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinusedfor"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
)

// CoinUsedForDelete is the builder for deleting a CoinUsedFor entity.
type CoinUsedForDelete struct {
	config
	hooks    []Hook
	mutation *CoinUsedForMutation
}

// Where appends a list predicates to the CoinUsedForDelete builder.
func (cufd *CoinUsedForDelete) Where(ps ...predicate.CoinUsedFor) *CoinUsedForDelete {
	cufd.mutation.Where(ps...)
	return cufd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cufd *CoinUsedForDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cufd.hooks) == 0 {
		affected, err = cufd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinUsedForMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cufd.mutation = mutation
			affected, err = cufd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cufd.hooks) - 1; i >= 0; i-- {
			if cufd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cufd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cufd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cufd *CoinUsedForDelete) ExecX(ctx context.Context) int {
	n, err := cufd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cufd *CoinUsedForDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: coinusedfor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coinusedfor.FieldID,
			},
		},
	}
	if ps := cufd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cufd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CoinUsedForDeleteOne is the builder for deleting a single CoinUsedFor entity.
type CoinUsedForDeleteOne struct {
	cufd *CoinUsedForDelete
}

// Exec executes the deletion query.
func (cufdo *CoinUsedForDeleteOne) Exec(ctx context.Context) error {
	n, err := cufdo.cufd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coinusedfor.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cufdo *CoinUsedForDeleteOne) ExecX(ctx context.Context) {
	cufdo.cufd.ExecX(ctx)
}
