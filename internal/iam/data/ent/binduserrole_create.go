// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/binduserrole"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/role"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/user"
)

// BindUserRoleCreate is the builder for creating a BindUserRole entity.
type BindUserRoleCreate struct {
	config
	mutation *BindUserRoleMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (burc *BindUserRoleCreate) SetUsername(s string) *BindUserRoleCreate {
	burc.mutation.SetUsername(s)
	return burc
}

// SetRolename sets the "rolename" field.
func (burc *BindUserRoleCreate) SetRolename(s string) *BindUserRoleCreate {
	burc.mutation.SetRolename(s)
	return burc
}

// SetCreateTime sets the "create_time" field.
func (burc *BindUserRoleCreate) SetCreateTime(t time.Time) *BindUserRoleCreate {
	burc.mutation.SetCreateTime(t)
	return burc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (burc *BindUserRoleCreate) SetNillableCreateTime(t *time.Time) *BindUserRoleCreate {
	if t != nil {
		burc.SetCreateTime(*t)
	}
	return burc
}

// SetUpdateTime sets the "update_time" field.
func (burc *BindUserRoleCreate) SetUpdateTime(t time.Time) *BindUserRoleCreate {
	burc.mutation.SetUpdateTime(t)
	return burc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (burc *BindUserRoleCreate) SetNillableUpdateTime(t *time.Time) *BindUserRoleCreate {
	if t != nil {
		burc.SetUpdateTime(*t)
	}
	return burc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (burc *BindUserRoleCreate) SetUserID(id int) *BindUserRoleCreate {
	burc.mutation.SetUserID(id)
	return burc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (burc *BindUserRoleCreate) SetNillableUserID(id *int) *BindUserRoleCreate {
	if id != nil {
		burc = burc.SetUserID(*id)
	}
	return burc
}

// SetUser sets the "user" edge to the User entity.
func (burc *BindUserRoleCreate) SetUser(u *User) *BindUserRoleCreate {
	return burc.SetUserID(u.ID)
}

// SetRoleID sets the "role" edge to the Role entity by ID.
func (burc *BindUserRoleCreate) SetRoleID(id int) *BindUserRoleCreate {
	burc.mutation.SetRoleID(id)
	return burc
}

// SetNillableRoleID sets the "role" edge to the Role entity by ID if the given value is not nil.
func (burc *BindUserRoleCreate) SetNillableRoleID(id *int) *BindUserRoleCreate {
	if id != nil {
		burc = burc.SetRoleID(*id)
	}
	return burc
}

// SetRole sets the "role" edge to the Role entity.
func (burc *BindUserRoleCreate) SetRole(r *Role) *BindUserRoleCreate {
	return burc.SetRoleID(r.ID)
}

// Mutation returns the BindUserRoleMutation object of the builder.
func (burc *BindUserRoleCreate) Mutation() *BindUserRoleMutation {
	return burc.mutation
}

// Save creates the BindUserRole in the database.
func (burc *BindUserRoleCreate) Save(ctx context.Context) (*BindUserRole, error) {
	var (
		err  error
		node *BindUserRole
	)
	burc.defaults()
	if len(burc.hooks) == 0 {
		if err = burc.check(); err != nil {
			return nil, err
		}
		node, err = burc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BindUserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = burc.check(); err != nil {
				return nil, err
			}
			burc.mutation = mutation
			if node, err = burc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(burc.hooks) - 1; i >= 0; i-- {
			if burc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = burc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, burc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (burc *BindUserRoleCreate) SaveX(ctx context.Context) *BindUserRole {
	v, err := burc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (burc *BindUserRoleCreate) Exec(ctx context.Context) error {
	_, err := burc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (burc *BindUserRoleCreate) ExecX(ctx context.Context) {
	if err := burc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (burc *BindUserRoleCreate) defaults() {
	if _, ok := burc.mutation.CreateTime(); !ok {
		v := binduserrole.DefaultCreateTime()
		burc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (burc *BindUserRoleCreate) check() error {
	if _, ok := burc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "BindUserRole.username"`)}
	}
	if v, ok := burc.mutation.Username(); ok {
		if err := binduserrole.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "BindUserRole.username": %w`, err)}
		}
	}
	if _, ok := burc.mutation.Rolename(); !ok {
		return &ValidationError{Name: "rolename", err: errors.New(`ent: missing required field "BindUserRole.rolename"`)}
	}
	if v, ok := burc.mutation.Rolename(); ok {
		if err := binduserrole.RolenameValidator(v); err != nil {
			return &ValidationError{Name: "rolename", err: fmt.Errorf(`ent: validator failed for field "BindUserRole.rolename": %w`, err)}
		}
	}
	if _, ok := burc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "BindUserRole.create_time"`)}
	}
	return nil
}

func (burc *BindUserRoleCreate) sqlSave(ctx context.Context) (*BindUserRole, error) {
	_node, _spec := burc.createSpec()
	if err := sqlgraph.CreateNode(ctx, burc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (burc *BindUserRoleCreate) createSpec() (*BindUserRole, *sqlgraph.CreateSpec) {
	var (
		_node = &BindUserRole{config: burc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: binduserrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: binduserrole.FieldID,
			},
		}
	)
	if value, ok := burc.mutation.Username(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: binduserrole.FieldUsername,
		})
		_node.Username = value
	}
	if value, ok := burc.mutation.Rolename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: binduserrole.FieldRolename,
		})
		_node.Rolename = value
	}
	if value, ok := burc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: binduserrole.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := burc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: binduserrole.FieldUpdateTime,
		})
		_node.UpdateTime = &value
	}
	if nodes := burc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   binduserrole.UserTable,
			Columns: []string{binduserrole.UserColumn},
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
		_node.user_bindings = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := burc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   binduserrole.RoleTable,
			Columns: []string{binduserrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.role_bindings = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BindUserRoleCreateBulk is the builder for creating many BindUserRole entities in bulk.
type BindUserRoleCreateBulk struct {
	config
	builders []*BindUserRoleCreate
}

// Save creates the BindUserRole entities in the database.
func (burcb *BindUserRoleCreateBulk) Save(ctx context.Context) ([]*BindUserRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(burcb.builders))
	nodes := make([]*BindUserRole, len(burcb.builders))
	mutators := make([]Mutator, len(burcb.builders))
	for i := range burcb.builders {
		func(i int, root context.Context) {
			builder := burcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BindUserRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, burcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, burcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, burcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (burcb *BindUserRoleCreateBulk) SaveX(ctx context.Context) []*BindUserRole {
	v, err := burcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (burcb *BindUserRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := burcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (burcb *BindUserRoleCreateBulk) ExecX(ctx context.Context) {
	if err := burcb.Exec(ctx); err != nil {
		panic(err)
	}
}