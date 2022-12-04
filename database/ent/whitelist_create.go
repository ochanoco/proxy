// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ochanoco/database/ent/serviceprovider"
	"github.com/ochanoco/database/ent/whitelist"
)

// WhiteListCreate is the builder for creating a WhiteList entity.
type WhiteListCreate struct {
	config
	mutation *WhiteListMutation
	hooks    []Hook
}

// SetPath sets the "path" field.
func (wlc *WhiteListCreate) SetPath(s string) *WhiteListCreate {
	wlc.mutation.SetPath(s)
	return wlc
}

// SetOwnerID sets the "owner" edge to the ServiceProvider entity by ID.
func (wlc *WhiteListCreate) SetOwnerID(id int) *WhiteListCreate {
	wlc.mutation.SetOwnerID(id)
	return wlc
}

// SetNillableOwnerID sets the "owner" edge to the ServiceProvider entity by ID if the given value is not nil.
func (wlc *WhiteListCreate) SetNillableOwnerID(id *int) *WhiteListCreate {
	if id != nil {
		wlc = wlc.SetOwnerID(*id)
	}
	return wlc
}

// SetOwner sets the "owner" edge to the ServiceProvider entity.
func (wlc *WhiteListCreate) SetOwner(s *ServiceProvider) *WhiteListCreate {
	return wlc.SetOwnerID(s.ID)
}

// Mutation returns the WhiteListMutation object of the builder.
func (wlc *WhiteListCreate) Mutation() *WhiteListMutation {
	return wlc.mutation
}

// Save creates the WhiteList in the database.
func (wlc *WhiteListCreate) Save(ctx context.Context) (*WhiteList, error) {
	var (
		err  error
		node *WhiteList
	)
	if len(wlc.hooks) == 0 {
		if err = wlc.check(); err != nil {
			return nil, err
		}
		node, err = wlc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WhiteListMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wlc.check(); err != nil {
				return nil, err
			}
			wlc.mutation = mutation
			if node, err = wlc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wlc.hooks) - 1; i >= 0; i-- {
			if wlc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wlc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wlc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*WhiteList)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WhiteListMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wlc *WhiteListCreate) SaveX(ctx context.Context) *WhiteList {
	v, err := wlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wlc *WhiteListCreate) Exec(ctx context.Context) error {
	_, err := wlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wlc *WhiteListCreate) ExecX(ctx context.Context) {
	if err := wlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wlc *WhiteListCreate) check() error {
	if _, ok := wlc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "WhiteList.path"`)}
	}
	return nil
}

func (wlc *WhiteListCreate) sqlSave(ctx context.Context) (*WhiteList, error) {
	_node, _spec := wlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (wlc *WhiteListCreate) createSpec() (*WhiteList, *sqlgraph.CreateSpec) {
	var (
		_node = &WhiteList{config: wlc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: whitelist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: whitelist.FieldID,
			},
		}
	)
	if value, ok := wlc.mutation.Path(); ok {
		_spec.SetField(whitelist.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if nodes := wlc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   whitelist.OwnerTable,
			Columns: []string{whitelist.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: serviceprovider.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.service_provider_whitelists = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WhiteListCreateBulk is the builder for creating many WhiteList entities in bulk.
type WhiteListCreateBulk struct {
	config
	builders []*WhiteListCreate
}

// Save creates the WhiteList entities in the database.
func (wlcb *WhiteListCreateBulk) Save(ctx context.Context) ([]*WhiteList, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wlcb.builders))
	nodes := make([]*WhiteList, len(wlcb.builders))
	mutators := make([]Mutator, len(wlcb.builders))
	for i := range wlcb.builders {
		func(i int, root context.Context) {
			builder := wlcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WhiteListMutation)
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
					_, err = mutators[i+1].Mutate(root, wlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, wlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wlcb *WhiteListCreateBulk) SaveX(ctx context.Context) []*WhiteList {
	v, err := wlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wlcb *WhiteListCreateBulk) Exec(ctx context.Context) error {
	_, err := wlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wlcb *WhiteListCreateBulk) ExecX(ctx context.Context) {
	if err := wlcb.Exec(ctx); err != nil {
		panic(err)
	}
}
