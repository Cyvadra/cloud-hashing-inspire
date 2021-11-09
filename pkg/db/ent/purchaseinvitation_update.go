// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/purchaseinvitation"
	"github.com/google/uuid"
)

// PurchaseInvitationUpdate is the builder for updating PurchaseInvitation entities.
type PurchaseInvitationUpdate struct {
	config
	hooks    []Hook
	mutation *PurchaseInvitationMutation
}

// Where appends a list predicates to the PurchaseInvitationUpdate builder.
func (piu *PurchaseInvitationUpdate) Where(ps ...predicate.PurchaseInvitation) *PurchaseInvitationUpdate {
	piu.mutation.Where(ps...)
	return piu
}

// SetAppID sets the "app_id" field.
func (piu *PurchaseInvitationUpdate) SetAppID(u uuid.UUID) *PurchaseInvitationUpdate {
	piu.mutation.SetAppID(u)
	return piu
}

// SetOrderID sets the "order_id" field.
func (piu *PurchaseInvitationUpdate) SetOrderID(u uuid.UUID) *PurchaseInvitationUpdate {
	piu.mutation.SetOrderID(u)
	return piu
}

// SetInvitationCodeID sets the "invitation_code_id" field.
func (piu *PurchaseInvitationUpdate) SetInvitationCodeID(u uuid.UUID) *PurchaseInvitationUpdate {
	piu.mutation.SetInvitationCodeID(u)
	return piu
}

// SetFulfilled sets the "fulfilled" field.
func (piu *PurchaseInvitationUpdate) SetFulfilled(b bool) *PurchaseInvitationUpdate {
	piu.mutation.SetFulfilled(b)
	return piu
}

// SetNillableFulfilled sets the "fulfilled" field if the given value is not nil.
func (piu *PurchaseInvitationUpdate) SetNillableFulfilled(b *bool) *PurchaseInvitationUpdate {
	if b != nil {
		piu.SetFulfilled(*b)
	}
	return piu
}

// SetCreateAt sets the "create_at" field.
func (piu *PurchaseInvitationUpdate) SetCreateAt(u uint32) *PurchaseInvitationUpdate {
	piu.mutation.ResetCreateAt()
	piu.mutation.SetCreateAt(u)
	return piu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (piu *PurchaseInvitationUpdate) SetNillableCreateAt(u *uint32) *PurchaseInvitationUpdate {
	if u != nil {
		piu.SetCreateAt(*u)
	}
	return piu
}

// AddCreateAt adds u to the "create_at" field.
func (piu *PurchaseInvitationUpdate) AddCreateAt(u uint32) *PurchaseInvitationUpdate {
	piu.mutation.AddCreateAt(u)
	return piu
}

// SetUpdateAt sets the "update_at" field.
func (piu *PurchaseInvitationUpdate) SetUpdateAt(u uint32) *PurchaseInvitationUpdate {
	piu.mutation.ResetUpdateAt()
	piu.mutation.SetUpdateAt(u)
	return piu
}

// AddUpdateAt adds u to the "update_at" field.
func (piu *PurchaseInvitationUpdate) AddUpdateAt(u uint32) *PurchaseInvitationUpdate {
	piu.mutation.AddUpdateAt(u)
	return piu
}

// SetDeleteAt sets the "delete_at" field.
func (piu *PurchaseInvitationUpdate) SetDeleteAt(u uint32) *PurchaseInvitationUpdate {
	piu.mutation.ResetDeleteAt()
	piu.mutation.SetDeleteAt(u)
	return piu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (piu *PurchaseInvitationUpdate) SetNillableDeleteAt(u *uint32) *PurchaseInvitationUpdate {
	if u != nil {
		piu.SetDeleteAt(*u)
	}
	return piu
}

// AddDeleteAt adds u to the "delete_at" field.
func (piu *PurchaseInvitationUpdate) AddDeleteAt(u uint32) *PurchaseInvitationUpdate {
	piu.mutation.AddDeleteAt(u)
	return piu
}

// Mutation returns the PurchaseInvitationMutation object of the builder.
func (piu *PurchaseInvitationUpdate) Mutation() *PurchaseInvitationMutation {
	return piu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *PurchaseInvitationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	piu.defaults()
	if len(piu.hooks) == 0 {
		affected, err = piu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PurchaseInvitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			piu.mutation = mutation
			affected, err = piu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(piu.hooks) - 1; i >= 0; i-- {
			if piu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = piu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, piu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (piu *PurchaseInvitationUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *PurchaseInvitationUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *PurchaseInvitationUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (piu *PurchaseInvitationUpdate) defaults() {
	if _, ok := piu.mutation.UpdateAt(); !ok {
		v := purchaseinvitation.UpdateDefaultUpdateAt()
		piu.mutation.SetUpdateAt(v)
	}
}

func (piu *PurchaseInvitationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   purchaseinvitation.Table,
			Columns: purchaseinvitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: purchaseinvitation.FieldID,
			},
		},
	}
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: purchaseinvitation.FieldAppID,
		})
	}
	if value, ok := piu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: purchaseinvitation.FieldOrderID,
		})
	}
	if value, ok := piu.mutation.InvitationCodeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: purchaseinvitation.FieldInvitationCodeID,
		})
	}
	if value, ok := piu.mutation.Fulfilled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: purchaseinvitation.FieldFulfilled,
		})
	}
	if value, ok := piu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldCreateAt,
		})
	}
	if value, ok := piu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldCreateAt,
		})
	}
	if value, ok := piu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldUpdateAt,
		})
	}
	if value, ok := piu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldUpdateAt,
		})
	}
	if value, ok := piu.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldDeleteAt,
		})
	}
	if value, ok := piu.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{purchaseinvitation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PurchaseInvitationUpdateOne is the builder for updating a single PurchaseInvitation entity.
type PurchaseInvitationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PurchaseInvitationMutation
}

// SetAppID sets the "app_id" field.
func (piuo *PurchaseInvitationUpdateOne) SetAppID(u uuid.UUID) *PurchaseInvitationUpdateOne {
	piuo.mutation.SetAppID(u)
	return piuo
}

// SetOrderID sets the "order_id" field.
func (piuo *PurchaseInvitationUpdateOne) SetOrderID(u uuid.UUID) *PurchaseInvitationUpdateOne {
	piuo.mutation.SetOrderID(u)
	return piuo
}

// SetInvitationCodeID sets the "invitation_code_id" field.
func (piuo *PurchaseInvitationUpdateOne) SetInvitationCodeID(u uuid.UUID) *PurchaseInvitationUpdateOne {
	piuo.mutation.SetInvitationCodeID(u)
	return piuo
}

// SetFulfilled sets the "fulfilled" field.
func (piuo *PurchaseInvitationUpdateOne) SetFulfilled(b bool) *PurchaseInvitationUpdateOne {
	piuo.mutation.SetFulfilled(b)
	return piuo
}

// SetNillableFulfilled sets the "fulfilled" field if the given value is not nil.
func (piuo *PurchaseInvitationUpdateOne) SetNillableFulfilled(b *bool) *PurchaseInvitationUpdateOne {
	if b != nil {
		piuo.SetFulfilled(*b)
	}
	return piuo
}

// SetCreateAt sets the "create_at" field.
func (piuo *PurchaseInvitationUpdateOne) SetCreateAt(u uint32) *PurchaseInvitationUpdateOne {
	piuo.mutation.ResetCreateAt()
	piuo.mutation.SetCreateAt(u)
	return piuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (piuo *PurchaseInvitationUpdateOne) SetNillableCreateAt(u *uint32) *PurchaseInvitationUpdateOne {
	if u != nil {
		piuo.SetCreateAt(*u)
	}
	return piuo
}

// AddCreateAt adds u to the "create_at" field.
func (piuo *PurchaseInvitationUpdateOne) AddCreateAt(u uint32) *PurchaseInvitationUpdateOne {
	piuo.mutation.AddCreateAt(u)
	return piuo
}

// SetUpdateAt sets the "update_at" field.
func (piuo *PurchaseInvitationUpdateOne) SetUpdateAt(u uint32) *PurchaseInvitationUpdateOne {
	piuo.mutation.ResetUpdateAt()
	piuo.mutation.SetUpdateAt(u)
	return piuo
}

// AddUpdateAt adds u to the "update_at" field.
func (piuo *PurchaseInvitationUpdateOne) AddUpdateAt(u uint32) *PurchaseInvitationUpdateOne {
	piuo.mutation.AddUpdateAt(u)
	return piuo
}

// SetDeleteAt sets the "delete_at" field.
func (piuo *PurchaseInvitationUpdateOne) SetDeleteAt(u uint32) *PurchaseInvitationUpdateOne {
	piuo.mutation.ResetDeleteAt()
	piuo.mutation.SetDeleteAt(u)
	return piuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (piuo *PurchaseInvitationUpdateOne) SetNillableDeleteAt(u *uint32) *PurchaseInvitationUpdateOne {
	if u != nil {
		piuo.SetDeleteAt(*u)
	}
	return piuo
}

// AddDeleteAt adds u to the "delete_at" field.
func (piuo *PurchaseInvitationUpdateOne) AddDeleteAt(u uint32) *PurchaseInvitationUpdateOne {
	piuo.mutation.AddDeleteAt(u)
	return piuo
}

// Mutation returns the PurchaseInvitationMutation object of the builder.
func (piuo *PurchaseInvitationUpdateOne) Mutation() *PurchaseInvitationMutation {
	return piuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piuo *PurchaseInvitationUpdateOne) Select(field string, fields ...string) *PurchaseInvitationUpdateOne {
	piuo.fields = append([]string{field}, fields...)
	return piuo
}

// Save executes the query and returns the updated PurchaseInvitation entity.
func (piuo *PurchaseInvitationUpdateOne) Save(ctx context.Context) (*PurchaseInvitation, error) {
	var (
		err  error
		node *PurchaseInvitation
	)
	piuo.defaults()
	if len(piuo.hooks) == 0 {
		node, err = piuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PurchaseInvitationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			piuo.mutation = mutation
			node, err = piuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(piuo.hooks) - 1; i >= 0; i-- {
			if piuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = piuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, piuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *PurchaseInvitationUpdateOne) SaveX(ctx context.Context) *PurchaseInvitation {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *PurchaseInvitationUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *PurchaseInvitationUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (piuo *PurchaseInvitationUpdateOne) defaults() {
	if _, ok := piuo.mutation.UpdateAt(); !ok {
		v := purchaseinvitation.UpdateDefaultUpdateAt()
		piuo.mutation.SetUpdateAt(v)
	}
}

func (piuo *PurchaseInvitationUpdateOne) sqlSave(ctx context.Context) (_node *PurchaseInvitation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   purchaseinvitation.Table,
			Columns: purchaseinvitation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: purchaseinvitation.FieldID,
			},
		},
	}
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PurchaseInvitation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := piuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, purchaseinvitation.FieldID)
		for _, f := range fields {
			if !purchaseinvitation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != purchaseinvitation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: purchaseinvitation.FieldAppID,
		})
	}
	if value, ok := piuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: purchaseinvitation.FieldOrderID,
		})
	}
	if value, ok := piuo.mutation.InvitationCodeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: purchaseinvitation.FieldInvitationCodeID,
		})
	}
	if value, ok := piuo.mutation.Fulfilled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: purchaseinvitation.FieldFulfilled,
		})
	}
	if value, ok := piuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldCreateAt,
		})
	}
	if value, ok := piuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldCreateAt,
		})
	}
	if value, ok := piuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldUpdateAt,
		})
	}
	if value, ok := piuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldUpdateAt,
		})
	}
	if value, ok := piuo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldDeleteAt,
		})
	}
	if value, ok := piuo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: purchaseinvitation.FieldDeleteAt,
		})
	}
	_node = &PurchaseInvitation{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{purchaseinvitation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
