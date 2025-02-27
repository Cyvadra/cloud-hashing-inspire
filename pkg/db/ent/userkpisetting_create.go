// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/userkpisetting"
	"github.com/google/uuid"
)

// UserKpiSettingCreate is the builder for creating a UserKpiSetting entity.
type UserKpiSettingCreate struct {
	config
	mutation *UserKpiSettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAmount sets the "amount" field.
func (uksc *UserKpiSettingCreate) SetAmount(u uint64) *UserKpiSettingCreate {
	uksc.mutation.SetAmount(u)
	return uksc
}

// SetPercent sets the "percent" field.
func (uksc *UserKpiSettingCreate) SetPercent(i int32) *UserKpiSettingCreate {
	uksc.mutation.SetPercent(i)
	return uksc
}

// SetAppID sets the "app_id" field.
func (uksc *UserKpiSettingCreate) SetAppID(u uuid.UUID) *UserKpiSettingCreate {
	uksc.mutation.SetAppID(u)
	return uksc
}

// SetGoodID sets the "good_id" field.
func (uksc *UserKpiSettingCreate) SetGoodID(u uuid.UUID) *UserKpiSettingCreate {
	uksc.mutation.SetGoodID(u)
	return uksc
}

// SetUserID sets the "user_id" field.
func (uksc *UserKpiSettingCreate) SetUserID(u uuid.UUID) *UserKpiSettingCreate {
	uksc.mutation.SetUserID(u)
	return uksc
}

// SetCreateAt sets the "create_at" field.
func (uksc *UserKpiSettingCreate) SetCreateAt(u uint32) *UserKpiSettingCreate {
	uksc.mutation.SetCreateAt(u)
	return uksc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (uksc *UserKpiSettingCreate) SetNillableCreateAt(u *uint32) *UserKpiSettingCreate {
	if u != nil {
		uksc.SetCreateAt(*u)
	}
	return uksc
}

// SetUpdateAt sets the "update_at" field.
func (uksc *UserKpiSettingCreate) SetUpdateAt(u uint32) *UserKpiSettingCreate {
	uksc.mutation.SetUpdateAt(u)
	return uksc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (uksc *UserKpiSettingCreate) SetNillableUpdateAt(u *uint32) *UserKpiSettingCreate {
	if u != nil {
		uksc.SetUpdateAt(*u)
	}
	return uksc
}

// SetDeleteAt sets the "delete_at" field.
func (uksc *UserKpiSettingCreate) SetDeleteAt(u uint32) *UserKpiSettingCreate {
	uksc.mutation.SetDeleteAt(u)
	return uksc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (uksc *UserKpiSettingCreate) SetNillableDeleteAt(u *uint32) *UserKpiSettingCreate {
	if u != nil {
		uksc.SetDeleteAt(*u)
	}
	return uksc
}

// SetID sets the "id" field.
func (uksc *UserKpiSettingCreate) SetID(u uuid.UUID) *UserKpiSettingCreate {
	uksc.mutation.SetID(u)
	return uksc
}

// Mutation returns the UserKpiSettingMutation object of the builder.
func (uksc *UserKpiSettingCreate) Mutation() *UserKpiSettingMutation {
	return uksc.mutation
}

// Save creates the UserKpiSetting in the database.
func (uksc *UserKpiSettingCreate) Save(ctx context.Context) (*UserKpiSetting, error) {
	var (
		err  error
		node *UserKpiSetting
	)
	uksc.defaults()
	if len(uksc.hooks) == 0 {
		if err = uksc.check(); err != nil {
			return nil, err
		}
		node, err = uksc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserKpiSettingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uksc.check(); err != nil {
				return nil, err
			}
			uksc.mutation = mutation
			if node, err = uksc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uksc.hooks) - 1; i >= 0; i-- {
			if uksc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uksc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uksc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uksc *UserKpiSettingCreate) SaveX(ctx context.Context) *UserKpiSetting {
	v, err := uksc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uksc *UserKpiSettingCreate) Exec(ctx context.Context) error {
	_, err := uksc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uksc *UserKpiSettingCreate) ExecX(ctx context.Context) {
	if err := uksc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uksc *UserKpiSettingCreate) defaults() {
	if _, ok := uksc.mutation.CreateAt(); !ok {
		v := userkpisetting.DefaultCreateAt()
		uksc.mutation.SetCreateAt(v)
	}
	if _, ok := uksc.mutation.UpdateAt(); !ok {
		v := userkpisetting.DefaultUpdateAt()
		uksc.mutation.SetUpdateAt(v)
	}
	if _, ok := uksc.mutation.DeleteAt(); !ok {
		v := userkpisetting.DefaultDeleteAt()
		uksc.mutation.SetDeleteAt(v)
	}
	if _, ok := uksc.mutation.ID(); !ok {
		v := userkpisetting.DefaultID()
		uksc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uksc *UserKpiSettingCreate) check() error {
	if _, ok := uksc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "amount"`)}
	}
	if _, ok := uksc.mutation.Percent(); !ok {
		return &ValidationError{Name: "percent", err: errors.New(`ent: missing required field "percent"`)}
	}
	if _, ok := uksc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "app_id"`)}
	}
	if _, ok := uksc.mutation.GoodID(); !ok {
		return &ValidationError{Name: "good_id", err: errors.New(`ent: missing required field "good_id"`)}
	}
	if _, ok := uksc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if _, ok := uksc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := uksc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := uksc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (uksc *UserKpiSettingCreate) sqlSave(ctx context.Context) (*UserKpiSetting, error) {
	_node, _spec := uksc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uksc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (uksc *UserKpiSettingCreate) createSpec() (*UserKpiSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &UserKpiSetting{config: uksc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userkpisetting.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: userkpisetting.FieldID,
			},
		}
	)
	_spec.OnConflict = uksc.conflict
	if id, ok := uksc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uksc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: userkpisetting.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := uksc.mutation.Percent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: userkpisetting.FieldPercent,
		})
		_node.Percent = value
	}
	if value, ok := uksc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userkpisetting.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := uksc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userkpisetting.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := uksc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: userkpisetting.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := uksc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userkpisetting.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := uksc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userkpisetting.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := uksc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: userkpisetting.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserKpiSetting.Create().
//		SetAmount(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserKpiSettingUpsert) {
//			SetAmount(v+v).
//		}).
//		Exec(ctx)
//
func (uksc *UserKpiSettingCreate) OnConflict(opts ...sql.ConflictOption) *UserKpiSettingUpsertOne {
	uksc.conflict = opts
	return &UserKpiSettingUpsertOne{
		create: uksc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserKpiSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (uksc *UserKpiSettingCreate) OnConflictColumns(columns ...string) *UserKpiSettingUpsertOne {
	uksc.conflict = append(uksc.conflict, sql.ConflictColumns(columns...))
	return &UserKpiSettingUpsertOne{
		create: uksc,
	}
}

type (
	// UserKpiSettingUpsertOne is the builder for "upsert"-ing
	//  one UserKpiSetting node.
	UserKpiSettingUpsertOne struct {
		create *UserKpiSettingCreate
	}

	// UserKpiSettingUpsert is the "OnConflict" setter.
	UserKpiSettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetAmount sets the "amount" field.
func (u *UserKpiSettingUpsert) SetAmount(v uint64) *UserKpiSettingUpsert {
	u.Set(userkpisetting.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserKpiSettingUpsert) UpdateAmount() *UserKpiSettingUpsert {
	u.SetExcluded(userkpisetting.FieldAmount)
	return u
}

// SetPercent sets the "percent" field.
func (u *UserKpiSettingUpsert) SetPercent(v int32) *UserKpiSettingUpsert {
	u.Set(userkpisetting.FieldPercent, v)
	return u
}

// UpdatePercent sets the "percent" field to the value that was provided on create.
func (u *UserKpiSettingUpsert) UpdatePercent() *UserKpiSettingUpsert {
	u.SetExcluded(userkpisetting.FieldPercent)
	return u
}

// SetAppID sets the "app_id" field.
func (u *UserKpiSettingUpsert) SetAppID(v uuid.UUID) *UserKpiSettingUpsert {
	u.Set(userkpisetting.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserKpiSettingUpsert) UpdateAppID() *UserKpiSettingUpsert {
	u.SetExcluded(userkpisetting.FieldAppID)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *UserKpiSettingUpsert) SetGoodID(v uuid.UUID) *UserKpiSettingUpsert {
	u.Set(userkpisetting.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *UserKpiSettingUpsert) UpdateGoodID() *UserKpiSettingUpsert {
	u.SetExcluded(userkpisetting.FieldGoodID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserKpiSettingUpsert) SetUserID(v uuid.UUID) *UserKpiSettingUpsert {
	u.Set(userkpisetting.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserKpiSettingUpsert) UpdateUserID() *UserKpiSettingUpsert {
	u.SetExcluded(userkpisetting.FieldUserID)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *UserKpiSettingUpsert) SetCreateAt(v uint32) *UserKpiSettingUpsert {
	u.Set(userkpisetting.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserKpiSettingUpsert) UpdateCreateAt() *UserKpiSettingUpsert {
	u.SetExcluded(userkpisetting.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *UserKpiSettingUpsert) SetUpdateAt(v uint32) *UserKpiSettingUpsert {
	u.Set(userkpisetting.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserKpiSettingUpsert) UpdateUpdateAt() *UserKpiSettingUpsert {
	u.SetExcluded(userkpisetting.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserKpiSettingUpsert) SetDeleteAt(v uint32) *UserKpiSettingUpsert {
	u.Set(userkpisetting.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserKpiSettingUpsert) UpdateDeleteAt() *UserKpiSettingUpsert {
	u.SetExcluded(userkpisetting.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.UserKpiSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userkpisetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserKpiSettingUpsertOne) UpdateNewValues() *UserKpiSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(userkpisetting.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.UserKpiSetting.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *UserKpiSettingUpsertOne) Ignore() *UserKpiSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserKpiSettingUpsertOne) DoNothing() *UserKpiSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserKpiSettingCreate.OnConflict
// documentation for more info.
func (u *UserKpiSettingUpsertOne) Update(set func(*UserKpiSettingUpsert)) *UserKpiSettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserKpiSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetAmount sets the "amount" field.
func (u *UserKpiSettingUpsertOne) SetAmount(v uint64) *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserKpiSettingUpsertOne) UpdateAmount() *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateAmount()
	})
}

// SetPercent sets the "percent" field.
func (u *UserKpiSettingUpsertOne) SetPercent(v int32) *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetPercent(v)
	})
}

// UpdatePercent sets the "percent" field to the value that was provided on create.
func (u *UserKpiSettingUpsertOne) UpdatePercent() *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdatePercent()
	})
}

// SetAppID sets the "app_id" field.
func (u *UserKpiSettingUpsertOne) SetAppID(v uuid.UUID) *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserKpiSettingUpsertOne) UpdateAppID() *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateAppID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *UserKpiSettingUpsertOne) SetGoodID(v uuid.UUID) *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *UserKpiSettingUpsertOne) UpdateGoodID() *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateGoodID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserKpiSettingUpsertOne) SetUserID(v uuid.UUID) *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserKpiSettingUpsertOne) UpdateUserID() *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateUserID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *UserKpiSettingUpsertOne) SetCreateAt(v uint32) *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserKpiSettingUpsertOne) UpdateCreateAt() *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *UserKpiSettingUpsertOne) SetUpdateAt(v uint32) *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserKpiSettingUpsertOne) UpdateUpdateAt() *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserKpiSettingUpsertOne) SetDeleteAt(v uint32) *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserKpiSettingUpsertOne) UpdateDeleteAt() *UserKpiSettingUpsertOne {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *UserKpiSettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserKpiSettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserKpiSettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserKpiSettingUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserKpiSettingUpsertOne.ID is not supported by MySQL driver. Use UserKpiSettingUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserKpiSettingUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserKpiSettingCreateBulk is the builder for creating many UserKpiSetting entities in bulk.
type UserKpiSettingCreateBulk struct {
	config
	builders []*UserKpiSettingCreate
	conflict []sql.ConflictOption
}

// Save creates the UserKpiSetting entities in the database.
func (ukscb *UserKpiSettingCreateBulk) Save(ctx context.Context) ([]*UserKpiSetting, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ukscb.builders))
	nodes := make([]*UserKpiSetting, len(ukscb.builders))
	mutators := make([]Mutator, len(ukscb.builders))
	for i := range ukscb.builders {
		func(i int, root context.Context) {
			builder := ukscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserKpiSettingMutation)
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
					_, err = mutators[i+1].Mutate(root, ukscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ukscb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ukscb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ukscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ukscb *UserKpiSettingCreateBulk) SaveX(ctx context.Context) []*UserKpiSetting {
	v, err := ukscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ukscb *UserKpiSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := ukscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ukscb *UserKpiSettingCreateBulk) ExecX(ctx context.Context) {
	if err := ukscb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserKpiSetting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserKpiSettingUpsert) {
//			SetAmount(v+v).
//		}).
//		Exec(ctx)
//
func (ukscb *UserKpiSettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserKpiSettingUpsertBulk {
	ukscb.conflict = opts
	return &UserKpiSettingUpsertBulk{
		create: ukscb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserKpiSetting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ukscb *UserKpiSettingCreateBulk) OnConflictColumns(columns ...string) *UserKpiSettingUpsertBulk {
	ukscb.conflict = append(ukscb.conflict, sql.ConflictColumns(columns...))
	return &UserKpiSettingUpsertBulk{
		create: ukscb,
	}
}

// UserKpiSettingUpsertBulk is the builder for "upsert"-ing
// a bulk of UserKpiSetting nodes.
type UserKpiSettingUpsertBulk struct {
	create *UserKpiSettingCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserKpiSetting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(userkpisetting.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *UserKpiSettingUpsertBulk) UpdateNewValues() *UserKpiSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(userkpisetting.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserKpiSetting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *UserKpiSettingUpsertBulk) Ignore() *UserKpiSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserKpiSettingUpsertBulk) DoNothing() *UserKpiSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserKpiSettingCreateBulk.OnConflict
// documentation for more info.
func (u *UserKpiSettingUpsertBulk) Update(set func(*UserKpiSettingUpsert)) *UserKpiSettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserKpiSettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetAmount sets the "amount" field.
func (u *UserKpiSettingUpsertBulk) SetAmount(v uint64) *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *UserKpiSettingUpsertBulk) UpdateAmount() *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateAmount()
	})
}

// SetPercent sets the "percent" field.
func (u *UserKpiSettingUpsertBulk) SetPercent(v int32) *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetPercent(v)
	})
}

// UpdatePercent sets the "percent" field to the value that was provided on create.
func (u *UserKpiSettingUpsertBulk) UpdatePercent() *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdatePercent()
	})
}

// SetAppID sets the "app_id" field.
func (u *UserKpiSettingUpsertBulk) SetAppID(v uuid.UUID) *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *UserKpiSettingUpsertBulk) UpdateAppID() *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateAppID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *UserKpiSettingUpsertBulk) SetGoodID(v uuid.UUID) *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *UserKpiSettingUpsertBulk) UpdateGoodID() *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateGoodID()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserKpiSettingUpsertBulk) SetUserID(v uuid.UUID) *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserKpiSettingUpsertBulk) UpdateUserID() *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateUserID()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *UserKpiSettingUpsertBulk) SetCreateAt(v uint32) *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *UserKpiSettingUpsertBulk) UpdateCreateAt() *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *UserKpiSettingUpsertBulk) SetUpdateAt(v uint32) *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *UserKpiSettingUpsertBulk) UpdateUpdateAt() *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *UserKpiSettingUpsertBulk) SetDeleteAt(v uint32) *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *UserKpiSettingUpsertBulk) UpdateDeleteAt() *UserKpiSettingUpsertBulk {
	return u.Update(func(s *UserKpiSettingUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *UserKpiSettingUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserKpiSettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserKpiSettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserKpiSettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
