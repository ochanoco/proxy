// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ochanoco/proxy/ent/hashchain"
	"github.com/ochanoco/proxy/ent/predicate"
)

// HashChainDelete is the builder for deleting a HashChain entity.
type HashChainDelete struct {
	config
	hooks    []Hook
	mutation *HashChainMutation
}

// Where appends a list predicates to the HashChainDelete builder.
func (hcd *HashChainDelete) Where(ps ...predicate.HashChain) *HashChainDelete {
	hcd.mutation.Where(ps...)
	return hcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hcd *HashChainDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hcd.hooks) == 0 {
		affected, err = hcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HashChainMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hcd.mutation = mutation
			affected, err = hcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hcd.hooks) - 1; i >= 0; i-- {
			if hcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcd *HashChainDelete) ExecX(ctx context.Context) int {
	n, err := hcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hcd *HashChainDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: hashchain.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: hashchain.FieldID,
			},
		},
	}
	if ps := hcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// HashChainDeleteOne is the builder for deleting a single HashChain entity.
type HashChainDeleteOne struct {
	hcd *HashChainDelete
}

// Exec executes the deletion query.
func (hcdo *HashChainDeleteOne) Exec(ctx context.Context) error {
	n, err := hcdo.hcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hashchain.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hcdo *HashChainDeleteOne) ExecX(ctx context.Context) {
	hcdo.hcd.ExecX(ctx)
}
