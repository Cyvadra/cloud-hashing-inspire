// Code generated by entc, DO NOT EDIT.

package discountpool

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-inspire/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// Discount applies equality check predicate on the "discount" field. It's identical to DiscountEQ.
func Discount(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscount), v))
	})
}

// ReleaseByUserID applies equality check predicate on the "release_by_user_id" field. It's identical to ReleaseByUserIDEQ.
func ReleaseByUserID(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReleaseByUserID), v))
	})
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// DurationDays applies equality check predicate on the "duration_days" field. It's identical to DurationDaysEQ.
func DurationDays(v int32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDurationDays), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// DiscountEQ applies the EQ predicate on the "discount" field.
func DiscountEQ(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscount), v))
	})
}

// DiscountNEQ applies the NEQ predicate on the "discount" field.
func DiscountNEQ(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscount), v))
	})
}

// DiscountIn applies the In predicate on the "discount" field.
func DiscountIn(vs ...uint32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDiscount), v...))
	})
}

// DiscountNotIn applies the NotIn predicate on the "discount" field.
func DiscountNotIn(vs ...uint32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDiscount), v...))
	})
}

// DiscountGT applies the GT predicate on the "discount" field.
func DiscountGT(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscount), v))
	})
}

// DiscountGTE applies the GTE predicate on the "discount" field.
func DiscountGTE(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscount), v))
	})
}

// DiscountLT applies the LT predicate on the "discount" field.
func DiscountLT(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscount), v))
	})
}

// DiscountLTE applies the LTE predicate on the "discount" field.
func DiscountLTE(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscount), v))
	})
}

// ReleaseByUserIDEQ applies the EQ predicate on the "release_by_user_id" field.
func ReleaseByUserIDEQ(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReleaseByUserID), v))
	})
}

// ReleaseByUserIDNEQ applies the NEQ predicate on the "release_by_user_id" field.
func ReleaseByUserIDNEQ(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReleaseByUserID), v))
	})
}

// ReleaseByUserIDIn applies the In predicate on the "release_by_user_id" field.
func ReleaseByUserIDIn(vs ...uuid.UUID) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReleaseByUserID), v...))
	})
}

// ReleaseByUserIDNotIn applies the NotIn predicate on the "release_by_user_id" field.
func ReleaseByUserIDNotIn(vs ...uuid.UUID) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReleaseByUserID), v...))
	})
}

// ReleaseByUserIDGT applies the GT predicate on the "release_by_user_id" field.
func ReleaseByUserIDGT(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReleaseByUserID), v))
	})
}

// ReleaseByUserIDGTE applies the GTE predicate on the "release_by_user_id" field.
func ReleaseByUserIDGTE(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReleaseByUserID), v))
	})
}

// ReleaseByUserIDLT applies the LT predicate on the "release_by_user_id" field.
func ReleaseByUserIDLT(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReleaseByUserID), v))
	})
}

// ReleaseByUserIDLTE applies the LTE predicate on the "release_by_user_id" field.
func ReleaseByUserIDLTE(v uuid.UUID) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReleaseByUserID), v))
	})
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStart), v))
	})
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...uint32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStart), v...))
	})
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...uint32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStart), v...))
	})
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStart), v))
	})
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStart), v))
	})
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStart), v))
	})
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStart), v))
	})
}

// DurationDaysEQ applies the EQ predicate on the "duration_days" field.
func DurationDaysEQ(v int32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDurationDays), v))
	})
}

// DurationDaysNEQ applies the NEQ predicate on the "duration_days" field.
func DurationDaysNEQ(v int32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDurationDays), v))
	})
}

// DurationDaysIn applies the In predicate on the "duration_days" field.
func DurationDaysIn(vs ...int32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDurationDays), v...))
	})
}

// DurationDaysNotIn applies the NotIn predicate on the "duration_days" field.
func DurationDaysNotIn(vs ...int32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDurationDays), v...))
	})
}

// DurationDaysGT applies the GT predicate on the "duration_days" field.
func DurationDaysGT(v int32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDurationDays), v))
	})
}

// DurationDaysGTE applies the GTE predicate on the "duration_days" field.
func DurationDaysGTE(v int32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDurationDays), v))
	})
}

// DurationDaysLT applies the LT predicate on the "duration_days" field.
func DurationDaysLT(v int32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDurationDays), v))
	})
}

// DurationDaysLTE applies the LTE predicate on the "duration_days" field.
func DurationDaysLTE(v int32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDurationDays), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.DiscountPool {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscountPool(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DiscountPool) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DiscountPool) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DiscountPool) predicate.DiscountPool {
	return predicate.DiscountPool(func(s *sql.Selector) {
		p(s.Not())
	})
}
