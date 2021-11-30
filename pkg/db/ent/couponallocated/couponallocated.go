// Code generated by entc, DO NOT EDIT.

package couponallocated

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the couponallocated type in the database.
	Label = "coupon_allocated"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCouponID holds the string denoting the coupon_id field in the database.
	FieldCouponID = "coupon_id"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the couponallocated in the database.
	Table = "coupon_allocateds"
)

// Columns holds all SQL columns for couponallocated fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldAppID,
	FieldType,
	FieldCouponID,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeDiscount Type = "discount"
	TypeCoupon   Type = "coupon"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeDiscount, TypeCoupon:
		return nil
	default:
		return fmt.Errorf("couponallocated: invalid enum value for type field: %q", _type)
	}
}
