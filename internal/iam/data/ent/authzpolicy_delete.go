// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/authzpolicy"
	"github.com/yangchnet/skeleton/internal/iam/data/ent/predicate"
)

// AuthzPolicyDelete is the builder for deleting a AuthzPolicy entity.
type AuthzPolicyDelete struct {
	config
	hooks    []Hook
	mutation *AuthzPolicyMutation
}

// Where appends a list predicates to the AuthzPolicyDelete builder.
func (apd *AuthzPolicyDelete) Where(ps ...predicate.AuthzPolicy) *AuthzPolicyDelete {
	apd.mutation.Where(ps...)
	return apd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (apd *AuthzPolicyDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(apd.hooks) == 0 {
		affected, err = apd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthzPolicyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			apd.mutation = mutation
			affected, err = apd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(apd.hooks) - 1; i >= 0; i-- {
			if apd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = apd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, apd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (apd *AuthzPolicyDelete) ExecX(ctx context.Context) int {
	n, err := apd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (apd *AuthzPolicyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: authzpolicy.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authzpolicy.FieldID,
			},
		},
	}
	if ps := apd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, apd.driver, _spec)
}

// AuthzPolicyDeleteOne is the builder for deleting a single AuthzPolicy entity.
type AuthzPolicyDeleteOne struct {
	apd *AuthzPolicyDelete
}

// Exec executes the deletion query.
func (apdo *AuthzPolicyDeleteOne) Exec(ctx context.Context) error {
	n, err := apdo.apd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{authzpolicy.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (apdo *AuthzPolicyDeleteOne) ExecX(ctx context.Context) {
	apdo.apd.ExecX(ctx)
}
