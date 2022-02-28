// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/binduserrole"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/predicate"
)

// BindUserRoleDelete is the builder for deleting a BindUserRole entity.
type BindUserRoleDelete struct {
	config
	hooks    []Hook
	mutation *BindUserRoleMutation
}

// Where appends a list predicates to the BindUserRoleDelete builder.
func (burd *BindUserRoleDelete) Where(ps ...predicate.BindUserRole) *BindUserRoleDelete {
	burd.mutation.Where(ps...)
	return burd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (burd *BindUserRoleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(burd.hooks) == 0 {
		affected, err = burd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BindUserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			burd.mutation = mutation
			affected, err = burd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(burd.hooks) - 1; i >= 0; i-- {
			if burd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = burd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, burd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (burd *BindUserRoleDelete) ExecX(ctx context.Context) int {
	n, err := burd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (burd *BindUserRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: binduserrole.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: binduserrole.FieldID,
			},
		},
	}
	if ps := burd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, burd.driver, _spec)
}

// BindUserRoleDeleteOne is the builder for deleting a single BindUserRole entity.
type BindUserRoleDeleteOne struct {
	burd *BindUserRoleDelete
}

// Exec executes the deletion query.
func (burdo *BindUserRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := burdo.burd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{binduserrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (burdo *BindUserRoleDeleteOne) ExecX(ctx context.Context) {
	burdo.burd.ExecX(ctx)
}