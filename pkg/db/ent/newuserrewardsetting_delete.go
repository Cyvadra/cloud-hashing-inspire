// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/newuserrewardsetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
)

// NewUserRewardSettingDelete is the builder for deleting a NewUserRewardSetting entity.
type NewUserRewardSettingDelete struct {
	config
	hooks    []Hook
	mutation *NewUserRewardSettingMutation
}

// Where appends a list predicates to the NewUserRewardSettingDelete builder.
func (nursd *NewUserRewardSettingDelete) Where(ps ...predicate.NewUserRewardSetting) *NewUserRewardSettingDelete {
	nursd.mutation.Where(ps...)
	return nursd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nursd *NewUserRewardSettingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nursd.hooks) == 0 {
		affected, err = nursd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewUserRewardSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nursd.mutation = mutation
			affected, err = nursd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nursd.hooks) - 1; i >= 0; i-- {
			if nursd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nursd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nursd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nursd *NewUserRewardSettingDelete) ExecX(ctx context.Context) int {
	n, err := nursd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nursd *NewUserRewardSettingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: newuserrewardsetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: newuserrewardsetting.FieldID,
			},
		},
	}
	if ps := nursd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, nursd.driver, _spec)
}

// NewUserRewardSettingDeleteOne is the builder for deleting a single NewUserRewardSetting entity.
type NewUserRewardSettingDeleteOne struct {
	nursd *NewUserRewardSettingDelete
}

// Exec executes the deletion query.
func (nursdo *NewUserRewardSettingDeleteOne) Exec(ctx context.Context) error {
	n, err := nursdo.nursd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{newuserrewardsetting.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nursdo *NewUserRewardSettingDeleteOne) ExecX(ctx context.Context) {
	nursdo.nursd.ExecX(ctx)
}
