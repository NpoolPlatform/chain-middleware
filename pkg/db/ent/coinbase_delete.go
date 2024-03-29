// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/coinbase"
	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent/predicate"
)

// CoinBaseDelete is the builder for deleting a CoinBase entity.
type CoinBaseDelete struct {
	config
	hooks    []Hook
	mutation *CoinBaseMutation
}

// Where appends a list predicates to the CoinBaseDelete builder.
func (cbd *CoinBaseDelete) Where(ps ...predicate.CoinBase) *CoinBaseDelete {
	cbd.mutation.Where(ps...)
	return cbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cbd *CoinBaseDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cbd.hooks) == 0 {
		affected, err = cbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoinBaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cbd.mutation = mutation
			affected, err = cbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cbd.hooks) - 1; i >= 0; i-- {
			if cbd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbd *CoinBaseDelete) ExecX(ctx context.Context) int {
	n, err := cbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cbd *CoinBaseDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: coinbase.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coinbase.FieldID,
			},
		},
	}
	if ps := cbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cbd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// CoinBaseDeleteOne is the builder for deleting a single CoinBase entity.
type CoinBaseDeleteOne struct {
	cbd *CoinBaseDelete
}

// Exec executes the deletion query.
func (cbdo *CoinBaseDeleteOne) Exec(ctx context.Context) error {
	n, err := cbdo.cbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coinbase.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cbdo *CoinBaseDeleteOne) ExecX(ctx context.Context) {
	cbdo.cbd.ExecX(ctx)
}
