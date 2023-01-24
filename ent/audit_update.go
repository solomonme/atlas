// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"playground/ronen-bootcamp/ent/audit"
	"playground/ronen-bootcamp/ent/organization"
	"playground/ronen-bootcamp/ent/predicate"
	"playground/ronen-bootcamp/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuditUpdate is the builder for updating Audit entities.
type AuditUpdate struct {
	config
	hooks    []Hook
	mutation *AuditMutation
}

// Where appends a list predicates to the AuditUpdate builder.
func (au *AuditUpdate) Where(ps ...predicate.Audit) *AuditUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetIdentity sets the "identity" field.
func (au *AuditUpdate) SetIdentity(s string) *AuditUpdate {
	au.mutation.SetIdentity(s)
	return au
}

// SetTimestamp sets the "timestamp" field.
func (au *AuditUpdate) SetTimestamp(t time.Time) *AuditUpdate {
	au.mutation.SetTimestamp(t)
	return au
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (au *AuditUpdate) SetNillableTimestamp(t *time.Time) *AuditUpdate {
	if t != nil {
		au.SetTimestamp(*t)
	}
	return au
}

// SetBalance sets the "balance" field.
func (au *AuditUpdate) SetBalance(f float64) *AuditUpdate {
	au.mutation.ResetBalance()
	au.mutation.SetBalance(f)
	return au
}

// AddBalance adds f to the "balance" field.
func (au *AuditUpdate) AddBalance(f float64) *AuditUpdate {
	au.mutation.AddBalance(f)
	return au
}

// SetDescription sets the "description" field.
func (au *AuditUpdate) SetDescription(s string) *AuditUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetOrgID sets the "org_id" field.
func (au *AuditUpdate) SetOrgID(i int) *AuditUpdate {
	au.mutation.SetOrgID(i)
	return au
}

// SetUserID sets the "user" edge to the User entity by ID.
func (au *AuditUpdate) SetUserID(id int) *AuditUpdate {
	au.mutation.SetUserID(id)
	return au
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (au *AuditUpdate) SetNillableUserID(id *int) *AuditUpdate {
	if id != nil {
		au = au.SetUserID(*id)
	}
	return au
}

// SetUser sets the "user" edge to the User entity.
func (au *AuditUpdate) SetUser(u *User) *AuditUpdate {
	return au.SetUserID(u.ID)
}

// SetOrganizationsID sets the "organizations" edge to the Organization entity by ID.
func (au *AuditUpdate) SetOrganizationsID(id int) *AuditUpdate {
	au.mutation.SetOrganizationsID(id)
	return au
}

// SetOrganizations sets the "organizations" edge to the Organization entity.
func (au *AuditUpdate) SetOrganizations(o *Organization) *AuditUpdate {
	return au.SetOrganizationsID(o.ID)
}

// Mutation returns the AuditMutation object of the builder.
func (au *AuditUpdate) Mutation() *AuditMutation {
	return au.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (au *AuditUpdate) ClearUser() *AuditUpdate {
	au.mutation.ClearUser()
	return au
}

// ClearOrganizations clears the "organizations" edge to the Organization entity.
func (au *AuditUpdate) ClearOrganizations() *AuditUpdate {
	au.mutation.ClearOrganizations()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AuditUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuditMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AuditUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AuditUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AuditUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AuditUpdate) check() error {
	if v, ok := au.mutation.Description(); ok {
		if err := audit.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Audit.description": %w`, err)}
		}
	}
	if _, ok := au.mutation.OrganizationsID(); au.mutation.OrganizationsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Audit.organizations"`)
	}
	return nil
}

func (au *AuditUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   audit.Table,
			Columns: audit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: audit.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Identity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: audit.FieldIdentity,
		})
	}
	if value, ok := au.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: audit.FieldTimestamp,
		})
	}
	if value, ok := au.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: audit.FieldBalance,
		})
	}
	if value, ok := au.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: audit.FieldBalance,
		})
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: audit.FieldDescription,
		})
	}
	if au.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audit.UserTable,
			Columns: []string{audit.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audit.UserTable,
			Columns: []string{audit.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audit.OrganizationsTable,
			Columns: []string{audit.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.OrganizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audit.OrganizationsTable,
			Columns: []string{audit.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{audit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AuditUpdateOne is the builder for updating a single Audit entity.
type AuditUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuditMutation
}

// SetIdentity sets the "identity" field.
func (auo *AuditUpdateOne) SetIdentity(s string) *AuditUpdateOne {
	auo.mutation.SetIdentity(s)
	return auo
}

// SetTimestamp sets the "timestamp" field.
func (auo *AuditUpdateOne) SetTimestamp(t time.Time) *AuditUpdateOne {
	auo.mutation.SetTimestamp(t)
	return auo
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (auo *AuditUpdateOne) SetNillableTimestamp(t *time.Time) *AuditUpdateOne {
	if t != nil {
		auo.SetTimestamp(*t)
	}
	return auo
}

// SetBalance sets the "balance" field.
func (auo *AuditUpdateOne) SetBalance(f float64) *AuditUpdateOne {
	auo.mutation.ResetBalance()
	auo.mutation.SetBalance(f)
	return auo
}

// AddBalance adds f to the "balance" field.
func (auo *AuditUpdateOne) AddBalance(f float64) *AuditUpdateOne {
	auo.mutation.AddBalance(f)
	return auo
}

// SetDescription sets the "description" field.
func (auo *AuditUpdateOne) SetDescription(s string) *AuditUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetOrgID sets the "org_id" field.
func (auo *AuditUpdateOne) SetOrgID(i int) *AuditUpdateOne {
	auo.mutation.SetOrgID(i)
	return auo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (auo *AuditUpdateOne) SetUserID(id int) *AuditUpdateOne {
	auo.mutation.SetUserID(id)
	return auo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (auo *AuditUpdateOne) SetNillableUserID(id *int) *AuditUpdateOne {
	if id != nil {
		auo = auo.SetUserID(*id)
	}
	return auo
}

// SetUser sets the "user" edge to the User entity.
func (auo *AuditUpdateOne) SetUser(u *User) *AuditUpdateOne {
	return auo.SetUserID(u.ID)
}

// SetOrganizationsID sets the "organizations" edge to the Organization entity by ID.
func (auo *AuditUpdateOne) SetOrganizationsID(id int) *AuditUpdateOne {
	auo.mutation.SetOrganizationsID(id)
	return auo
}

// SetOrganizations sets the "organizations" edge to the Organization entity.
func (auo *AuditUpdateOne) SetOrganizations(o *Organization) *AuditUpdateOne {
	return auo.SetOrganizationsID(o.ID)
}

// Mutation returns the AuditMutation object of the builder.
func (auo *AuditUpdateOne) Mutation() *AuditMutation {
	return auo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (auo *AuditUpdateOne) ClearUser() *AuditUpdateOne {
	auo.mutation.ClearUser()
	return auo
}

// ClearOrganizations clears the "organizations" edge to the Organization entity.
func (auo *AuditUpdateOne) ClearOrganizations() *AuditUpdateOne {
	auo.mutation.ClearOrganizations()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AuditUpdateOne) Select(field string, fields ...string) *AuditUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Audit entity.
func (auo *AuditUpdateOne) Save(ctx context.Context) (*Audit, error) {
	var (
		err  error
		node *Audit
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuditMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, auo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Audit)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AuditMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AuditUpdateOne) SaveX(ctx context.Context) *Audit {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AuditUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AuditUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AuditUpdateOne) check() error {
	if v, ok := auo.mutation.Description(); ok {
		if err := audit.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Audit.description": %w`, err)}
		}
	}
	if _, ok := auo.mutation.OrganizationsID(); auo.mutation.OrganizationsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Audit.organizations"`)
	}
	return nil
}

func (auo *AuditUpdateOne) sqlSave(ctx context.Context) (_node *Audit, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   audit.Table,
			Columns: audit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: audit.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Audit.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, audit.FieldID)
		for _, f := range fields {
			if !audit.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != audit.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Identity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: audit.FieldIdentity,
		})
	}
	if value, ok := auo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: audit.FieldTimestamp,
		})
	}
	if value, ok := auo.mutation.Balance(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: audit.FieldBalance,
		})
	}
	if value, ok := auo.mutation.AddedBalance(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: audit.FieldBalance,
		})
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: audit.FieldDescription,
		})
	}
	if auo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audit.UserTable,
			Columns: []string{audit.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audit.UserTable,
			Columns: []string{audit.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audit.OrganizationsTable,
			Columns: []string{audit.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.OrganizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audit.OrganizationsTable,
			Columns: []string{audit.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Audit{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{audit.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}