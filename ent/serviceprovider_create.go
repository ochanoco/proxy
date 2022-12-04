// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ochanoco/database/ent/authorizationcode"
	"github.com/ochanoco/database/ent/serviceprovider"
	"github.com/ochanoco/database/ent/whitelist"
)

// ServiceProviderCreate is the builder for creating a ServiceProvider entity.
type ServiceProviderCreate struct {
	config
	mutation *ServiceProviderMutation
	hooks    []Hook
}

// SetHost sets the "host" field.
func (spc *ServiceProviderCreate) SetHost(s string) *ServiceProviderCreate {
	spc.mutation.SetHost(s)
	return spc
}

// SetDestinationIP sets the "destination_ip" field.
func (spc *ServiceProviderCreate) SetDestinationIP(s string) *ServiceProviderCreate {
	spc.mutation.SetDestinationIP(s)
	return spc
}

// AddWhitelistIDs adds the "whitelists" edge to the WhiteList entity by IDs.
func (spc *ServiceProviderCreate) AddWhitelistIDs(ids ...int) *ServiceProviderCreate {
	spc.mutation.AddWhitelistIDs(ids...)
	return spc
}

// AddWhitelists adds the "whitelists" edges to the WhiteList entity.
func (spc *ServiceProviderCreate) AddWhitelists(w ...*WhiteList) *ServiceProviderCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return spc.AddWhitelistIDs(ids...)
}

// AddAuthorizationCodeIDs adds the "authorization_codes" edge to the AuthorizationCode entity by IDs.
func (spc *ServiceProviderCreate) AddAuthorizationCodeIDs(ids ...int) *ServiceProviderCreate {
	spc.mutation.AddAuthorizationCodeIDs(ids...)
	return spc
}

// AddAuthorizationCodes adds the "authorization_codes" edges to the AuthorizationCode entity.
func (spc *ServiceProviderCreate) AddAuthorizationCodes(a ...*AuthorizationCode) *ServiceProviderCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return spc.AddAuthorizationCodeIDs(ids...)
}

// Mutation returns the ServiceProviderMutation object of the builder.
func (spc *ServiceProviderCreate) Mutation() *ServiceProviderMutation {
	return spc.mutation
}

// Save creates the ServiceProvider in the database.
func (spc *ServiceProviderCreate) Save(ctx context.Context) (*ServiceProvider, error) {
	var (
		err  error
		node *ServiceProvider
	)
	if len(spc.hooks) == 0 {
		if err = spc.check(); err != nil {
			return nil, err
		}
		node, err = spc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = spc.check(); err != nil {
				return nil, err
			}
			spc.mutation = mutation
			if node, err = spc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(spc.hooks) - 1; i >= 0; i-- {
			if spc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, spc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ServiceProvider)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ServiceProviderMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (spc *ServiceProviderCreate) SaveX(ctx context.Context) *ServiceProvider {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *ServiceProviderCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *ServiceProviderCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *ServiceProviderCreate) check() error {
	if _, ok := spc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "ServiceProvider.host"`)}
	}
	if _, ok := spc.mutation.DestinationIP(); !ok {
		return &ValidationError{Name: "destination_ip", err: errors.New(`ent: missing required field "ServiceProvider.destination_ip"`)}
	}
	return nil
}

func (spc *ServiceProviderCreate) sqlSave(ctx context.Context) (*ServiceProvider, error) {
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (spc *ServiceProviderCreate) createSpec() (*ServiceProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &ServiceProvider{config: spc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: serviceprovider.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: serviceprovider.FieldID,
			},
		}
	)
	if value, ok := spc.mutation.Host(); ok {
		_spec.SetField(serviceprovider.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	if value, ok := spc.mutation.DestinationIP(); ok {
		_spec.SetField(serviceprovider.FieldDestinationIP, field.TypeString, value)
		_node.DestinationIP = value
	}
	if nodes := spc.mutation.WhitelistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceprovider.WhitelistsTable,
			Columns: []string{serviceprovider.WhitelistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: whitelist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := spc.mutation.AuthorizationCodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   serviceprovider.AuthorizationCodesTable,
			Columns: []string{serviceprovider.AuthorizationCodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: authorizationcode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServiceProviderCreateBulk is the builder for creating many ServiceProvider entities in bulk.
type ServiceProviderCreateBulk struct {
	config
	builders []*ServiceProviderCreate
}

// Save creates the ServiceProvider entities in the database.
func (spcb *ServiceProviderCreateBulk) Save(ctx context.Context) ([]*ServiceProvider, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*ServiceProvider, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceProviderMutation)
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
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *ServiceProviderCreateBulk) SaveX(ctx context.Context) []*ServiceProvider {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *ServiceProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *ServiceProviderCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}
