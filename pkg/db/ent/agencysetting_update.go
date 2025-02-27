// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/agencysetting"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AgencySettingUpdate is the builder for updating AgencySetting entities.
type AgencySettingUpdate struct {
	config
	hooks    []Hook
	mutation *AgencySettingMutation
}

// Where appends a list predicates to the AgencySettingUpdate builder.
func (asu *AgencySettingUpdate) Where(ps ...predicate.AgencySetting) *AgencySettingUpdate {
	asu.mutation.Where(ps...)
	return asu
}

// SetAppID sets the "app_id" field.
func (asu *AgencySettingUpdate) SetAppID(u uuid.UUID) *AgencySettingUpdate {
	asu.mutation.SetAppID(u)
	return asu
}

// SetRegistrationRewardThreshold sets the "registration_reward_threshold" field.
func (asu *AgencySettingUpdate) SetRegistrationRewardThreshold(i int32) *AgencySettingUpdate {
	asu.mutation.ResetRegistrationRewardThreshold()
	asu.mutation.SetRegistrationRewardThreshold(i)
	return asu
}

// AddRegistrationRewardThreshold adds i to the "registration_reward_threshold" field.
func (asu *AgencySettingUpdate) AddRegistrationRewardThreshold(i int32) *AgencySettingUpdate {
	asu.mutation.AddRegistrationRewardThreshold(i)
	return asu
}

// SetRegistrationCouponID sets the "registration_coupon_id" field.
func (asu *AgencySettingUpdate) SetRegistrationCouponID(u uuid.UUID) *AgencySettingUpdate {
	asu.mutation.SetRegistrationCouponID(u)
	return asu
}

// SetKycRewardThreshold sets the "kyc_reward_threshold" field.
func (asu *AgencySettingUpdate) SetKycRewardThreshold(i int32) *AgencySettingUpdate {
	asu.mutation.ResetKycRewardThreshold()
	asu.mutation.SetKycRewardThreshold(i)
	return asu
}

// AddKycRewardThreshold adds i to the "kyc_reward_threshold" field.
func (asu *AgencySettingUpdate) AddKycRewardThreshold(i int32) *AgencySettingUpdate {
	asu.mutation.AddKycRewardThreshold(i)
	return asu
}

// SetKycCouponID sets the "kyc_coupon_id" field.
func (asu *AgencySettingUpdate) SetKycCouponID(u uuid.UUID) *AgencySettingUpdate {
	asu.mutation.SetKycCouponID(u)
	return asu
}

// SetTotalPurchaseRewardPercent sets the "total_purchase_reward_percent" field.
func (asu *AgencySettingUpdate) SetTotalPurchaseRewardPercent(i int32) *AgencySettingUpdate {
	asu.mutation.ResetTotalPurchaseRewardPercent()
	asu.mutation.SetTotalPurchaseRewardPercent(i)
	return asu
}

// AddTotalPurchaseRewardPercent adds i to the "total_purchase_reward_percent" field.
func (asu *AgencySettingUpdate) AddTotalPurchaseRewardPercent(i int32) *AgencySettingUpdate {
	asu.mutation.AddTotalPurchaseRewardPercent(i)
	return asu
}

// SetPurchaseRewardChainLevels sets the "purchase_reward_chain_levels" field.
func (asu *AgencySettingUpdate) SetPurchaseRewardChainLevels(i int32) *AgencySettingUpdate {
	asu.mutation.ResetPurchaseRewardChainLevels()
	asu.mutation.SetPurchaseRewardChainLevels(i)
	return asu
}

// AddPurchaseRewardChainLevels adds i to the "purchase_reward_chain_levels" field.
func (asu *AgencySettingUpdate) AddPurchaseRewardChainLevels(i int32) *AgencySettingUpdate {
	asu.mutation.AddPurchaseRewardChainLevels(i)
	return asu
}

// SetLevelPurchaseRewardPercent sets the "level_purchase_reward_percent" field.
func (asu *AgencySettingUpdate) SetLevelPurchaseRewardPercent(i int32) *AgencySettingUpdate {
	asu.mutation.ResetLevelPurchaseRewardPercent()
	asu.mutation.SetLevelPurchaseRewardPercent(i)
	return asu
}

// AddLevelPurchaseRewardPercent adds i to the "level_purchase_reward_percent" field.
func (asu *AgencySettingUpdate) AddLevelPurchaseRewardPercent(i int32) *AgencySettingUpdate {
	asu.mutation.AddLevelPurchaseRewardPercent(i)
	return asu
}

// SetCreateAt sets the "create_at" field.
func (asu *AgencySettingUpdate) SetCreateAt(u uint32) *AgencySettingUpdate {
	asu.mutation.ResetCreateAt()
	asu.mutation.SetCreateAt(u)
	return asu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (asu *AgencySettingUpdate) SetNillableCreateAt(u *uint32) *AgencySettingUpdate {
	if u != nil {
		asu.SetCreateAt(*u)
	}
	return asu
}

// AddCreateAt adds u to the "create_at" field.
func (asu *AgencySettingUpdate) AddCreateAt(u uint32) *AgencySettingUpdate {
	asu.mutation.AddCreateAt(u)
	return asu
}

// SetUpdateAt sets the "update_at" field.
func (asu *AgencySettingUpdate) SetUpdateAt(u uint32) *AgencySettingUpdate {
	asu.mutation.ResetUpdateAt()
	asu.mutation.SetUpdateAt(u)
	return asu
}

// AddUpdateAt adds u to the "update_at" field.
func (asu *AgencySettingUpdate) AddUpdateAt(u uint32) *AgencySettingUpdate {
	asu.mutation.AddUpdateAt(u)
	return asu
}

// SetDeleteAt sets the "delete_at" field.
func (asu *AgencySettingUpdate) SetDeleteAt(u uint32) *AgencySettingUpdate {
	asu.mutation.ResetDeleteAt()
	asu.mutation.SetDeleteAt(u)
	return asu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (asu *AgencySettingUpdate) SetNillableDeleteAt(u *uint32) *AgencySettingUpdate {
	if u != nil {
		asu.SetDeleteAt(*u)
	}
	return asu
}

// AddDeleteAt adds u to the "delete_at" field.
func (asu *AgencySettingUpdate) AddDeleteAt(u uint32) *AgencySettingUpdate {
	asu.mutation.AddDeleteAt(u)
	return asu
}

// Mutation returns the AgencySettingMutation object of the builder.
func (asu *AgencySettingUpdate) Mutation() *AgencySettingMutation {
	return asu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (asu *AgencySettingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	asu.defaults()
	if len(asu.hooks) == 0 {
		affected, err = asu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgencySettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			asu.mutation = mutation
			affected, err = asu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(asu.hooks) - 1; i >= 0; i-- {
			if asu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (asu *AgencySettingUpdate) SaveX(ctx context.Context) int {
	affected, err := asu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (asu *AgencySettingUpdate) Exec(ctx context.Context) error {
	_, err := asu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asu *AgencySettingUpdate) ExecX(ctx context.Context) {
	if err := asu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asu *AgencySettingUpdate) defaults() {
	if _, ok := asu.mutation.UpdateAt(); !ok {
		v := agencysetting.UpdateDefaultUpdateAt()
		asu.mutation.SetUpdateAt(v)
	}
}

func (asu *AgencySettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agencysetting.Table,
			Columns: agencysetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: agencysetting.FieldID,
			},
		},
	}
	if ps := asu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: agencysetting.FieldAppID,
		})
	}
	if value, ok := asu.mutation.RegistrationRewardThreshold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldRegistrationRewardThreshold,
		})
	}
	if value, ok := asu.mutation.AddedRegistrationRewardThreshold(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldRegistrationRewardThreshold,
		})
	}
	if value, ok := asu.mutation.RegistrationCouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: agencysetting.FieldRegistrationCouponID,
		})
	}
	if value, ok := asu.mutation.KycRewardThreshold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldKycRewardThreshold,
		})
	}
	if value, ok := asu.mutation.AddedKycRewardThreshold(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldKycRewardThreshold,
		})
	}
	if value, ok := asu.mutation.KycCouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: agencysetting.FieldKycCouponID,
		})
	}
	if value, ok := asu.mutation.TotalPurchaseRewardPercent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldTotalPurchaseRewardPercent,
		})
	}
	if value, ok := asu.mutation.AddedTotalPurchaseRewardPercent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldTotalPurchaseRewardPercent,
		})
	}
	if value, ok := asu.mutation.PurchaseRewardChainLevels(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldPurchaseRewardChainLevels,
		})
	}
	if value, ok := asu.mutation.AddedPurchaseRewardChainLevels(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldPurchaseRewardChainLevels,
		})
	}
	if value, ok := asu.mutation.LevelPurchaseRewardPercent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldLevelPurchaseRewardPercent,
		})
	}
	if value, ok := asu.mutation.AddedLevelPurchaseRewardPercent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldLevelPurchaseRewardPercent,
		})
	}
	if value, ok := asu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldCreateAt,
		})
	}
	if value, ok := asu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldCreateAt,
		})
	}
	if value, ok := asu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldUpdateAt,
		})
	}
	if value, ok := asu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldUpdateAt,
		})
	}
	if value, ok := asu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldDeleteAt,
		})
	}
	if value, ok := asu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, asu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agencysetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AgencySettingUpdateOne is the builder for updating a single AgencySetting entity.
type AgencySettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AgencySettingMutation
}

// SetAppID sets the "app_id" field.
func (asuo *AgencySettingUpdateOne) SetAppID(u uuid.UUID) *AgencySettingUpdateOne {
	asuo.mutation.SetAppID(u)
	return asuo
}

// SetRegistrationRewardThreshold sets the "registration_reward_threshold" field.
func (asuo *AgencySettingUpdateOne) SetRegistrationRewardThreshold(i int32) *AgencySettingUpdateOne {
	asuo.mutation.ResetRegistrationRewardThreshold()
	asuo.mutation.SetRegistrationRewardThreshold(i)
	return asuo
}

// AddRegistrationRewardThreshold adds i to the "registration_reward_threshold" field.
func (asuo *AgencySettingUpdateOne) AddRegistrationRewardThreshold(i int32) *AgencySettingUpdateOne {
	asuo.mutation.AddRegistrationRewardThreshold(i)
	return asuo
}

// SetRegistrationCouponID sets the "registration_coupon_id" field.
func (asuo *AgencySettingUpdateOne) SetRegistrationCouponID(u uuid.UUID) *AgencySettingUpdateOne {
	asuo.mutation.SetRegistrationCouponID(u)
	return asuo
}

// SetKycRewardThreshold sets the "kyc_reward_threshold" field.
func (asuo *AgencySettingUpdateOne) SetKycRewardThreshold(i int32) *AgencySettingUpdateOne {
	asuo.mutation.ResetKycRewardThreshold()
	asuo.mutation.SetKycRewardThreshold(i)
	return asuo
}

// AddKycRewardThreshold adds i to the "kyc_reward_threshold" field.
func (asuo *AgencySettingUpdateOne) AddKycRewardThreshold(i int32) *AgencySettingUpdateOne {
	asuo.mutation.AddKycRewardThreshold(i)
	return asuo
}

// SetKycCouponID sets the "kyc_coupon_id" field.
func (asuo *AgencySettingUpdateOne) SetKycCouponID(u uuid.UUID) *AgencySettingUpdateOne {
	asuo.mutation.SetKycCouponID(u)
	return asuo
}

// SetTotalPurchaseRewardPercent sets the "total_purchase_reward_percent" field.
func (asuo *AgencySettingUpdateOne) SetTotalPurchaseRewardPercent(i int32) *AgencySettingUpdateOne {
	asuo.mutation.ResetTotalPurchaseRewardPercent()
	asuo.mutation.SetTotalPurchaseRewardPercent(i)
	return asuo
}

// AddTotalPurchaseRewardPercent adds i to the "total_purchase_reward_percent" field.
func (asuo *AgencySettingUpdateOne) AddTotalPurchaseRewardPercent(i int32) *AgencySettingUpdateOne {
	asuo.mutation.AddTotalPurchaseRewardPercent(i)
	return asuo
}

// SetPurchaseRewardChainLevels sets the "purchase_reward_chain_levels" field.
func (asuo *AgencySettingUpdateOne) SetPurchaseRewardChainLevels(i int32) *AgencySettingUpdateOne {
	asuo.mutation.ResetPurchaseRewardChainLevels()
	asuo.mutation.SetPurchaseRewardChainLevels(i)
	return asuo
}

// AddPurchaseRewardChainLevels adds i to the "purchase_reward_chain_levels" field.
func (asuo *AgencySettingUpdateOne) AddPurchaseRewardChainLevels(i int32) *AgencySettingUpdateOne {
	asuo.mutation.AddPurchaseRewardChainLevels(i)
	return asuo
}

// SetLevelPurchaseRewardPercent sets the "level_purchase_reward_percent" field.
func (asuo *AgencySettingUpdateOne) SetLevelPurchaseRewardPercent(i int32) *AgencySettingUpdateOne {
	asuo.mutation.ResetLevelPurchaseRewardPercent()
	asuo.mutation.SetLevelPurchaseRewardPercent(i)
	return asuo
}

// AddLevelPurchaseRewardPercent adds i to the "level_purchase_reward_percent" field.
func (asuo *AgencySettingUpdateOne) AddLevelPurchaseRewardPercent(i int32) *AgencySettingUpdateOne {
	asuo.mutation.AddLevelPurchaseRewardPercent(i)
	return asuo
}

// SetCreateAt sets the "create_at" field.
func (asuo *AgencySettingUpdateOne) SetCreateAt(u uint32) *AgencySettingUpdateOne {
	asuo.mutation.ResetCreateAt()
	asuo.mutation.SetCreateAt(u)
	return asuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (asuo *AgencySettingUpdateOne) SetNillableCreateAt(u *uint32) *AgencySettingUpdateOne {
	if u != nil {
		asuo.SetCreateAt(*u)
	}
	return asuo
}

// AddCreateAt adds u to the "create_at" field.
func (asuo *AgencySettingUpdateOne) AddCreateAt(u uint32) *AgencySettingUpdateOne {
	asuo.mutation.AddCreateAt(u)
	return asuo
}

// SetUpdateAt sets the "update_at" field.
func (asuo *AgencySettingUpdateOne) SetUpdateAt(u uint32) *AgencySettingUpdateOne {
	asuo.mutation.ResetUpdateAt()
	asuo.mutation.SetUpdateAt(u)
	return asuo
}

// AddUpdateAt adds u to the "update_at" field.
func (asuo *AgencySettingUpdateOne) AddUpdateAt(u uint32) *AgencySettingUpdateOne {
	asuo.mutation.AddUpdateAt(u)
	return asuo
}

// SetDeleteAt sets the "delete_at" field.
func (asuo *AgencySettingUpdateOne) SetDeleteAt(u uint32) *AgencySettingUpdateOne {
	asuo.mutation.ResetDeleteAt()
	asuo.mutation.SetDeleteAt(u)
	return asuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (asuo *AgencySettingUpdateOne) SetNillableDeleteAt(u *uint32) *AgencySettingUpdateOne {
	if u != nil {
		asuo.SetDeleteAt(*u)
	}
	return asuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (asuo *AgencySettingUpdateOne) AddDeleteAt(u uint32) *AgencySettingUpdateOne {
	asuo.mutation.AddDeleteAt(u)
	return asuo
}

// Mutation returns the AgencySettingMutation object of the builder.
func (asuo *AgencySettingUpdateOne) Mutation() *AgencySettingMutation {
	return asuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (asuo *AgencySettingUpdateOne) Select(field string, fields ...string) *AgencySettingUpdateOne {
	asuo.fields = append([]string{field}, fields...)
	return asuo
}

// Save executes the query and returns the updated AgencySetting entity.
func (asuo *AgencySettingUpdateOne) Save(ctx context.Context) (*AgencySetting, error) {
	var (
		err  error
		node *AgencySetting
	)
	asuo.defaults()
	if len(asuo.hooks) == 0 {
		node, err = asuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgencySettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			asuo.mutation = mutation
			node, err = asuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(asuo.hooks) - 1; i >= 0; i-- {
			if asuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (asuo *AgencySettingUpdateOne) SaveX(ctx context.Context) *AgencySetting {
	node, err := asuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (asuo *AgencySettingUpdateOne) Exec(ctx context.Context) error {
	_, err := asuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asuo *AgencySettingUpdateOne) ExecX(ctx context.Context) {
	if err := asuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asuo *AgencySettingUpdateOne) defaults() {
	if _, ok := asuo.mutation.UpdateAt(); !ok {
		v := agencysetting.UpdateDefaultUpdateAt()
		asuo.mutation.SetUpdateAt(v)
	}
}

func (asuo *AgencySettingUpdateOne) sqlSave(ctx context.Context) (_node *AgencySetting, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agencysetting.Table,
			Columns: agencysetting.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: agencysetting.FieldID,
			},
		},
	}
	id, ok := asuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing AgencySetting.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := asuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, agencysetting.FieldID)
		for _, f := range fields {
			if !agencysetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != agencysetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := asuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: agencysetting.FieldAppID,
		})
	}
	if value, ok := asuo.mutation.RegistrationRewardThreshold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldRegistrationRewardThreshold,
		})
	}
	if value, ok := asuo.mutation.AddedRegistrationRewardThreshold(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldRegistrationRewardThreshold,
		})
	}
	if value, ok := asuo.mutation.RegistrationCouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: agencysetting.FieldRegistrationCouponID,
		})
	}
	if value, ok := asuo.mutation.KycRewardThreshold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldKycRewardThreshold,
		})
	}
	if value, ok := asuo.mutation.AddedKycRewardThreshold(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldKycRewardThreshold,
		})
	}
	if value, ok := asuo.mutation.KycCouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: agencysetting.FieldKycCouponID,
		})
	}
	if value, ok := asuo.mutation.TotalPurchaseRewardPercent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldTotalPurchaseRewardPercent,
		})
	}
	if value, ok := asuo.mutation.AddedTotalPurchaseRewardPercent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldTotalPurchaseRewardPercent,
		})
	}
	if value, ok := asuo.mutation.PurchaseRewardChainLevels(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldPurchaseRewardChainLevels,
		})
	}
	if value, ok := asuo.mutation.AddedPurchaseRewardChainLevels(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldPurchaseRewardChainLevels,
		})
	}
	if value, ok := asuo.mutation.LevelPurchaseRewardPercent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldLevelPurchaseRewardPercent,
		})
	}
	if value, ok := asuo.mutation.AddedLevelPurchaseRewardPercent(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: agencysetting.FieldLevelPurchaseRewardPercent,
		})
	}
	if value, ok := asuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldCreateAt,
		})
	}
	if value, ok := asuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldCreateAt,
		})
	}
	if value, ok := asuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldUpdateAt,
		})
	}
	if value, ok := asuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldUpdateAt,
		})
	}
	if value, ok := asuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldDeleteAt,
		})
	}
	if value, ok := asuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: agencysetting.FieldDeleteAt,
		})
	}
	_node = &AgencySetting{config: asuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, asuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agencysetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
