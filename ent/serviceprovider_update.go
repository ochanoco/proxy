// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ochanoco/database/ent/authorizationcode"
	"github.com/ochanoco/database/ent/predicate"
	"github.com/ochanoco/database/ent/serviceprovider"
	"github.com/ochanoco/database/ent/whitelist"
)

// ServiceProviderUpdate is the builder for updating ServiceProvider entities.
type ServiceProviderUpdate struct {
	config
	hooks    []Hook
	mutation *ServiceProviderMutation
}

// Where appends a list predicates to the ServiceProviderUpdate builder.
func (spu *ServiceProviderUpdate) Where(ps ...predicate.ServiceProvider) *ServiceProviderUpdate {
	spu.mutation.Where(ps...)
	return spu
}

// SetHost sets the "host" field.
func (spu *ServiceProviderUpdate) SetHost(s string) *ServiceProviderUpdate {
	spu.mutation.SetHost(s)
	return spu
}

// SetDestinationIP sets the "destination_ip" field.
func (spu *ServiceProviderUpdate) SetDestinationIP(s string) *ServiceProviderUpdate {
	spu.mutation.SetDestinationIP(s)
	return spu
}

// AddWhitelistIDs adds the "whitelists" edge to the WhiteList entity by IDs.
func (spu *ServiceProviderUpdate) AddWhitelistIDs(ids ...int) *ServiceProviderUpdate {
	spu.mutation.AddWhitelistIDs(ids...)
	return spu
}

// AddWhitelists adds the "whitelists" edges to the WhiteList entity.
func (spu *ServiceProviderUpdate) AddWhitelists(w ...*WhiteList) *ServiceProviderUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return spu.AddWhitelistIDs(ids...)
}

// AddAuthorizationCodeIDs adds the "authorization_codes" edge to the AuthorizationCode entity by IDs.
func (spu *ServiceProviderUpdate) AddAuthorizationCodeIDs(ids ...int) *ServiceProviderUpdate {
	spu.mutation.AddAuthorizationCodeIDs(ids...)
	return spu
}

// AddAuthorizationCodes adds the "authorization_codes" edges to the AuthorizationCode entity.
func (spu *ServiceProviderUpdate) AddAuthorizationCodes(a ...*AuthorizationCode) *ServiceProviderUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return spu.AddAuthorizationCodeIDs(ids...)
}

// Mutation returns the ServiceProviderMutation object of the builder.
func (spu *ServiceProviderUpdate) Mutation() *ServiceProviderMutation {
	return spu.mutation
}

// ClearWhitelists clears all "whitelists" edges to the WhiteList entity.
func (spu *ServiceProviderUpdate) ClearWhitelists() *ServiceProviderUpdate {
	spu.mutation.ClearWhitelists()
	return spu
}

// RemoveWhitelistIDs removes the "whitelists" edge to WhiteList entities by IDs.
func (spu *ServiceProviderUpdate) RemoveWhitelistIDs(ids ...int) *ServiceProviderUpdate {
	spu.mutation.RemoveWhitelistIDs(ids...)
	return spu
}

// RemoveWhitelists removes "whitelists" edges to WhiteList entities.
func (spu *ServiceProviderUpdate) RemoveWhitelists(w ...*WhiteList) *ServiceProviderUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return spu.RemoveWhitelistIDs(ids...)
}

// ClearAuthorizationCodes clears all "authorization_codes" edges to the AuthorizationCode entity.
func (spu *ServiceProviderUpdate) ClearAuthorizationCodes() *ServiceProviderUpdate {
	spu.mutation.ClearAuthorizationCodes()
	return spu
}

// RemoveAuthorizationCodeIDs removes the "authorization_codes" edge to AuthorizationCode entities by IDs.
func (spu *ServiceProviderUpdate) RemoveAuthorizationCodeIDs(ids ...int) *ServiceProviderUpdate {
	spu.mutation.RemoveAuthorizationCodeIDs(ids...)
	return spu
}

// RemoveAuthorizationCodes removes "authorization_codes" edges to AuthorizationCode entities.
func (spu *ServiceProviderUpdate) RemoveAuthorizationCodes(a ...*AuthorizationCode) *ServiceProviderUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return spu.RemoveAuthorizationCodeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (spu *ServiceProviderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(spu.hooks) == 0 {
		affected, err = spu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spu.mutation = mutation
			affected, err = spu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spu.hooks) - 1; i >= 0; i-- {
			if spu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (spu *ServiceProviderUpdate) SaveX(ctx context.Context) int {
	affected, err := spu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (spu *ServiceProviderUpdate) Exec(ctx context.Context) error {
	_, err := spu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spu *ServiceProviderUpdate) ExecX(ctx context.Context) {
	if err := spu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (spu *ServiceProviderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceprovider.Table,
			Columns: serviceprovider.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: serviceprovider.FieldID,
			},
		},
	}
	if ps := spu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spu.mutation.Host(); ok {
		_spec.SetField(serviceprovider.FieldHost, field.TypeString, value)
	}
	if value, ok := spu.mutation.DestinationIP(); ok {
		_spec.SetField(serviceprovider.FieldDestinationIP, field.TypeString, value)
	}
	if spu.mutation.WhitelistsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.RemovedWhitelistsIDs(); len(nodes) > 0 && !spu.mutation.WhitelistsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.WhitelistsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if spu.mutation.AuthorizationCodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.RemovedAuthorizationCodesIDs(); len(nodes) > 0 && !spu.mutation.AuthorizationCodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spu.mutation.AuthorizationCodesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, spu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ServiceProviderUpdateOne is the builder for updating a single ServiceProvider entity.
type ServiceProviderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ServiceProviderMutation
}

// SetHost sets the "host" field.
func (spuo *ServiceProviderUpdateOne) SetHost(s string) *ServiceProviderUpdateOne {
	spuo.mutation.SetHost(s)
	return spuo
}

// SetDestinationIP sets the "destination_ip" field.
func (spuo *ServiceProviderUpdateOne) SetDestinationIP(s string) *ServiceProviderUpdateOne {
	spuo.mutation.SetDestinationIP(s)
	return spuo
}

// AddWhitelistIDs adds the "whitelists" edge to the WhiteList entity by IDs.
func (spuo *ServiceProviderUpdateOne) AddWhitelistIDs(ids ...int) *ServiceProviderUpdateOne {
	spuo.mutation.AddWhitelistIDs(ids...)
	return spuo
}

// AddWhitelists adds the "whitelists" edges to the WhiteList entity.
func (spuo *ServiceProviderUpdateOne) AddWhitelists(w ...*WhiteList) *ServiceProviderUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return spuo.AddWhitelistIDs(ids...)
}

// AddAuthorizationCodeIDs adds the "authorization_codes" edge to the AuthorizationCode entity by IDs.
func (spuo *ServiceProviderUpdateOne) AddAuthorizationCodeIDs(ids ...int) *ServiceProviderUpdateOne {
	spuo.mutation.AddAuthorizationCodeIDs(ids...)
	return spuo
}

// AddAuthorizationCodes adds the "authorization_codes" edges to the AuthorizationCode entity.
func (spuo *ServiceProviderUpdateOne) AddAuthorizationCodes(a ...*AuthorizationCode) *ServiceProviderUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return spuo.AddAuthorizationCodeIDs(ids...)
}

// Mutation returns the ServiceProviderMutation object of the builder.
func (spuo *ServiceProviderUpdateOne) Mutation() *ServiceProviderMutation {
	return spuo.mutation
}

// ClearWhitelists clears all "whitelists" edges to the WhiteList entity.
func (spuo *ServiceProviderUpdateOne) ClearWhitelists() *ServiceProviderUpdateOne {
	spuo.mutation.ClearWhitelists()
	return spuo
}

// RemoveWhitelistIDs removes the "whitelists" edge to WhiteList entities by IDs.
func (spuo *ServiceProviderUpdateOne) RemoveWhitelistIDs(ids ...int) *ServiceProviderUpdateOne {
	spuo.mutation.RemoveWhitelistIDs(ids...)
	return spuo
}

// RemoveWhitelists removes "whitelists" edges to WhiteList entities.
func (spuo *ServiceProviderUpdateOne) RemoveWhitelists(w ...*WhiteList) *ServiceProviderUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return spuo.RemoveWhitelistIDs(ids...)
}

// ClearAuthorizationCodes clears all "authorization_codes" edges to the AuthorizationCode entity.
func (spuo *ServiceProviderUpdateOne) ClearAuthorizationCodes() *ServiceProviderUpdateOne {
	spuo.mutation.ClearAuthorizationCodes()
	return spuo
}

// RemoveAuthorizationCodeIDs removes the "authorization_codes" edge to AuthorizationCode entities by IDs.
func (spuo *ServiceProviderUpdateOne) RemoveAuthorizationCodeIDs(ids ...int) *ServiceProviderUpdateOne {
	spuo.mutation.RemoveAuthorizationCodeIDs(ids...)
	return spuo
}

// RemoveAuthorizationCodes removes "authorization_codes" edges to AuthorizationCode entities.
func (spuo *ServiceProviderUpdateOne) RemoveAuthorizationCodes(a ...*AuthorizationCode) *ServiceProviderUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return spuo.RemoveAuthorizationCodeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (spuo *ServiceProviderUpdateOne) Select(field string, fields ...string) *ServiceProviderUpdateOne {
	spuo.fields = append([]string{field}, fields...)
	return spuo
}

// Save executes the query and returns the updated ServiceProvider entity.
func (spuo *ServiceProviderUpdateOne) Save(ctx context.Context) (*ServiceProvider, error) {
	var (
		err  error
		node *ServiceProvider
	)
	if len(spuo.hooks) == 0 {
		node, err = spuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spuo.mutation = mutation
			node, err = spuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(spuo.hooks) - 1; i >= 0; i-- {
			if spuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, spuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (spuo *ServiceProviderUpdateOne) SaveX(ctx context.Context) *ServiceProvider {
	node, err := spuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (spuo *ServiceProviderUpdateOne) Exec(ctx context.Context) error {
	_, err := spuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spuo *ServiceProviderUpdateOne) ExecX(ctx context.Context) {
	if err := spuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (spuo *ServiceProviderUpdateOne) sqlSave(ctx context.Context) (_node *ServiceProvider, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   serviceprovider.Table,
			Columns: serviceprovider.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: serviceprovider.FieldID,
			},
		},
	}
	id, ok := spuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ServiceProvider.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := spuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serviceprovider.FieldID)
		for _, f := range fields {
			if !serviceprovider.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != serviceprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := spuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spuo.mutation.Host(); ok {
		_spec.SetField(serviceprovider.FieldHost, field.TypeString, value)
	}
	if value, ok := spuo.mutation.DestinationIP(); ok {
		_spec.SetField(serviceprovider.FieldDestinationIP, field.TypeString, value)
	}
	if spuo.mutation.WhitelistsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.RemovedWhitelistsIDs(); len(nodes) > 0 && !spuo.mutation.WhitelistsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.WhitelistsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if spuo.mutation.AuthorizationCodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.RemovedAuthorizationCodesIDs(); len(nodes) > 0 && !spuo.mutation.AuthorizationCodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := spuo.mutation.AuthorizationCodesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ServiceProvider{config: spuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, spuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceprovider.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
