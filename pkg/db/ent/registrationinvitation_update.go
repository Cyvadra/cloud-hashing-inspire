// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/registrationinvitation"
	"github.com/google/uuid"
)

// RegistrationInvitationUpdate is the builder for updating RegistrationInvitation entities.
type RegistrationInvitationUpdate struct {
	config
	hooks    []Hook
	mutation *RegistrationInvitationMutation
}

// Where appends a list predicates to the RegistrationInvitationUpdate builder.
func (riu *RegistrationInvitationUpdate) Where(ps ...predicate.RegistrationInvitation) *RegistrationInvitationUpdate {
	riu.mutation.Where(ps...)
	return riu
}

// SetCreateAt sets the "create_at" field.
func (riu *RegistrationInvitationUpdate) SetCreateAt(u uint32) *RegistrationInvitationUpdate {
	riu.mutation.ResetCreateAt()
	riu.mutation.SetCreateAt(u)
	return riu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (riu *RegistrationInvitationUpdate) SetNillableCreateAt(u *uint32) *RegistrationInvitationUpdate {
	if u != nil {
		riu.SetCreateAt(*u)
	}
	return riu
}

// AddCreateAt adds u to the "create_at" field.
func (riu *RegistrationInvitationUpdate) AddCreateAt(u uint32) *RegistrationInvitationUpdate {
	riu.mutation.AddCreateAt(u)
	return riu
}

// SetUpdateAt sets the "update_at" field.
func (riu *RegistrationInvitationUpdate) SetUpdateAt(u uint32) *RegistrationInvitationUpdate {
	riu.mutation.ResetUpdateAt()
	riu.mutation.SetUpdateAt(u)
	return riu
}

// AddUpdateAt adds u to the "update_at" field.
func (riu *RegistrationInvitationUpdate) AddUpdateAt(u uint32) *RegistrationInvitationUpdate {
	riu.mutation.AddUpdateAt(u)
	return riu
}

// SetDeleteAt sets the "delete_at" field.
func (riu *RegistrationInvitationUpdate) SetDeleteAt(u uint32) *RegistrationInvitationUpdate {
	riu.mutation.ResetDeleteAt()
	riu.mutation.SetDeleteAt(u)
	return riu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (riu *RegistrationInvitationUpdate) SetNillableDeleteAt(u *uint32) *RegistrationInvitationUpdate {
	if u != nil {
		riu.SetDeleteAt(*u)
	}
	return riu
}

// AddDeleteAt adds u to the "delete_at" field.
func (riu *RegistrationInvitationUpdate) AddDeleteAt(u uint32) *RegistrationInvitationUpdate {
	riu.mutation.AddDeleteAt(u)
	return riu
}

// SetInviterID sets the "inviter_id" field.
func (riu *RegistrationInvitationUpdate) SetInviterID(u uuid.UUID) *RegistrationInvitationUpdate {
	riu.mutation.SetInviterID(u)
	return riu
}

// SetInviteeID sets the "invitee_id" field.
func (riu *RegistrationInvitationUpdate) SetInviteeID(u uuid.UUID) *RegistrationInvitationUpdate {
	riu.mutation.SetInviteeID(u)
	return riu
}

// SetAppID sets the "app_id" field.
func (riu *RegistrationInvitationUpdate) SetAppID(u uuid.UUID) *RegistrationInvitationUpdate {
	riu.mutation.SetAppID(u)
	return riu
}

// SetFulfilled sets the "fulfilled" field.
func (riu *RegistrationInvitationUpdate) SetFulfilled(b bool) *RegistrationInvitationUpdate {
	riu.mutation.SetFulfilled(b)
	return riu
}

// SetNillableFulfilled sets the "fulfilled" field if the given value is not nil.
func (riu *RegistrationInvitationUpdate) SetNillableFulfilled(b *bool) *RegistrationInvitationUpdate {
	if b != nil {
		riu.SetFulfilled(*b)
	}
	return riu
}

// Mutation returns the RegistrationInvitationMutation object of the builder.
func (riu *RegistrationInvitationUpdate) Mutation() *RegistrationInvitationMutation {
	return riu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (riu *RegistrationInvitationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	riu.defaults()
	if len(riu.hooks) == 0 {
		affected, err = riu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegistrationInvitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			riu.mutation = mutation
			affected, err = riu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(riu.hooks) - 1; i >= 0; i-- {
			if riu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = riu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, riu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (riu *RegistrationInvitationUpdate) SaveX(ctx context.Context) int {
	affected, err := riu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (riu *RegistrationInvitationUpdate) Exec(ctx context.Context) error {
	_, err := riu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riu *RegistrationInvitationUpdate) ExecX(ctx context.Context) {
	if err := riu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (riu *RegistrationInvitationUpdate) defaults() {
	if _, ok := riu.mutation.UpdateAt(); !ok {
		v := registrationinvitation.UpdateDefaultUpdateAt()
		riu.mutation.SetUpdateAt(v)
	}
}

func (riu *RegistrationInvitationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   registrationinvitation.Table,
			Columns: registrationinvitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: registrationinvitation.FieldID,
			},
		},
	}
	if ps := riu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := riu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldCreateAt,
		})
	}
	if value, ok := riu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldCreateAt,
		})
	}
	if value, ok := riu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldUpdateAt,
		})
	}
	if value, ok := riu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldUpdateAt,
		})
	}
	if value, ok := riu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldDeleteAt,
		})
	}
	if value, ok := riu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldDeleteAt,
		})
	}
	if value, ok := riu.mutation.InviterID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: registrationinvitation.FieldInviterID,
		})
	}
	if value, ok := riu.mutation.InviteeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: registrationinvitation.FieldInviteeID,
		})
	}
	if value, ok := riu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: registrationinvitation.FieldAppID,
		})
	}
	if value, ok := riu.mutation.Fulfilled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: registrationinvitation.FieldFulfilled,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, riu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{registrationinvitation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RegistrationInvitationUpdateOne is the builder for updating a single RegistrationInvitation entity.
type RegistrationInvitationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RegistrationInvitationMutation
}

// SetCreateAt sets the "create_at" field.
func (riuo *RegistrationInvitationUpdateOne) SetCreateAt(u uint32) *RegistrationInvitationUpdateOne {
	riuo.mutation.ResetCreateAt()
	riuo.mutation.SetCreateAt(u)
	return riuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (riuo *RegistrationInvitationUpdateOne) SetNillableCreateAt(u *uint32) *RegistrationInvitationUpdateOne {
	if u != nil {
		riuo.SetCreateAt(*u)
	}
	return riuo
}

// AddCreateAt adds u to the "create_at" field.
func (riuo *RegistrationInvitationUpdateOne) AddCreateAt(u uint32) *RegistrationInvitationUpdateOne {
	riuo.mutation.AddCreateAt(u)
	return riuo
}

// SetUpdateAt sets the "update_at" field.
func (riuo *RegistrationInvitationUpdateOne) SetUpdateAt(u uint32) *RegistrationInvitationUpdateOne {
	riuo.mutation.ResetUpdateAt()
	riuo.mutation.SetUpdateAt(u)
	return riuo
}

// AddUpdateAt adds u to the "update_at" field.
func (riuo *RegistrationInvitationUpdateOne) AddUpdateAt(u uint32) *RegistrationInvitationUpdateOne {
	riuo.mutation.AddUpdateAt(u)
	return riuo
}

// SetDeleteAt sets the "delete_at" field.
func (riuo *RegistrationInvitationUpdateOne) SetDeleteAt(u uint32) *RegistrationInvitationUpdateOne {
	riuo.mutation.ResetDeleteAt()
	riuo.mutation.SetDeleteAt(u)
	return riuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (riuo *RegistrationInvitationUpdateOne) SetNillableDeleteAt(u *uint32) *RegistrationInvitationUpdateOne {
	if u != nil {
		riuo.SetDeleteAt(*u)
	}
	return riuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (riuo *RegistrationInvitationUpdateOne) AddDeleteAt(u uint32) *RegistrationInvitationUpdateOne {
	riuo.mutation.AddDeleteAt(u)
	return riuo
}

// SetInviterID sets the "inviter_id" field.
func (riuo *RegistrationInvitationUpdateOne) SetInviterID(u uuid.UUID) *RegistrationInvitationUpdateOne {
	riuo.mutation.SetInviterID(u)
	return riuo
}

// SetInviteeID sets the "invitee_id" field.
func (riuo *RegistrationInvitationUpdateOne) SetInviteeID(u uuid.UUID) *RegistrationInvitationUpdateOne {
	riuo.mutation.SetInviteeID(u)
	return riuo
}

// SetAppID sets the "app_id" field.
func (riuo *RegistrationInvitationUpdateOne) SetAppID(u uuid.UUID) *RegistrationInvitationUpdateOne {
	riuo.mutation.SetAppID(u)
	return riuo
}

// SetFulfilled sets the "fulfilled" field.
func (riuo *RegistrationInvitationUpdateOne) SetFulfilled(b bool) *RegistrationInvitationUpdateOne {
	riuo.mutation.SetFulfilled(b)
	return riuo
}

// SetNillableFulfilled sets the "fulfilled" field if the given value is not nil.
func (riuo *RegistrationInvitationUpdateOne) SetNillableFulfilled(b *bool) *RegistrationInvitationUpdateOne {
	if b != nil {
		riuo.SetFulfilled(*b)
	}
	return riuo
}

// Mutation returns the RegistrationInvitationMutation object of the builder.
func (riuo *RegistrationInvitationUpdateOne) Mutation() *RegistrationInvitationMutation {
	return riuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (riuo *RegistrationInvitationUpdateOne) Select(field string, fields ...string) *RegistrationInvitationUpdateOne {
	riuo.fields = append([]string{field}, fields...)
	return riuo
}

// Save executes the query and returns the updated RegistrationInvitation entity.
func (riuo *RegistrationInvitationUpdateOne) Save(ctx context.Context) (*RegistrationInvitation, error) {
	var (
		err  error
		node *RegistrationInvitation
	)
	riuo.defaults()
	if len(riuo.hooks) == 0 {
		node, err = riuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RegistrationInvitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			riuo.mutation = mutation
			node, err = riuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(riuo.hooks) - 1; i >= 0; i-- {
			if riuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = riuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, riuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (riuo *RegistrationInvitationUpdateOne) SaveX(ctx context.Context) *RegistrationInvitation {
	node, err := riuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (riuo *RegistrationInvitationUpdateOne) Exec(ctx context.Context) error {
	_, err := riuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (riuo *RegistrationInvitationUpdateOne) ExecX(ctx context.Context) {
	if err := riuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (riuo *RegistrationInvitationUpdateOne) defaults() {
	if _, ok := riuo.mutation.UpdateAt(); !ok {
		v := registrationinvitation.UpdateDefaultUpdateAt()
		riuo.mutation.SetUpdateAt(v)
	}
}

func (riuo *RegistrationInvitationUpdateOne) sqlSave(ctx context.Context) (_node *RegistrationInvitation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   registrationinvitation.Table,
			Columns: registrationinvitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: registrationinvitation.FieldID,
			},
		},
	}
	id, ok := riuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing RegistrationInvitation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := riuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, registrationinvitation.FieldID)
		for _, f := range fields {
			if !registrationinvitation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != registrationinvitation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := riuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := riuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldCreateAt,
		})
	}
	if value, ok := riuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldCreateAt,
		})
	}
	if value, ok := riuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldUpdateAt,
		})
	}
	if value, ok := riuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldUpdateAt,
		})
	}
	if value, ok := riuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldDeleteAt,
		})
	}
	if value, ok := riuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: registrationinvitation.FieldDeleteAt,
		})
	}
	if value, ok := riuo.mutation.InviterID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: registrationinvitation.FieldInviterID,
		})
	}
	if value, ok := riuo.mutation.InviteeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: registrationinvitation.FieldInviteeID,
		})
	}
	if value, ok := riuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: registrationinvitation.FieldAppID,
		})
	}
	if value, ok := riuo.mutation.Fulfilled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: registrationinvitation.FieldFulfilled,
		})
	}
	_node = &RegistrationInvitation{config: riuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, riuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{registrationinvitation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
