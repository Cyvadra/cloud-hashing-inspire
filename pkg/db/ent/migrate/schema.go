// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AgencySettingsColumns holds the columns for the "agency_settings" table.
	AgencySettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Unique: true},
		{Name: "registration_reward_threshold", Type: field.TypeInt32},
		{Name: "registration_coupon_id", Type: field.TypeUUID},
		{Name: "kyc_reward_threshold", Type: field.TypeInt32},
		{Name: "kyc_coupon_id", Type: field.TypeUUID},
		{Name: "total_purchase_reward_percent", Type: field.TypeInt32},
		{Name: "purchase_reward_chain_levels", Type: field.TypeInt32},
		{Name: "level_purchase_reward_percent", Type: field.TypeInt32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AgencySettingsTable holds the schema information for the "agency_settings" table.
	AgencySettingsTable = &schema.Table{
		Name:       "agency_settings",
		Columns:    AgencySettingsColumns,
		PrimaryKey: []*schema.Column{AgencySettingsColumns[0]},
	}
	// AppCouponSettingsColumns holds the columns for the "app_coupon_settings" table.
	AppCouponSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Unique: true},
		{Name: "domination_limit", Type: field.TypeUint64},
		{Name: "total_limit", Type: field.TypeInt32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppCouponSettingsTable holds the schema information for the "app_coupon_settings" table.
	AppCouponSettingsTable = &schema.Table{
		Name:       "app_coupon_settings",
		Columns:    AppCouponSettingsColumns,
		PrimaryKey: []*schema.Column{AppCouponSettingsColumns[0]},
	}
	// CouponAllocatedsColumns holds the columns for the "coupon_allocateds" table.
	CouponAllocatedsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"discount", "coupon"}},
		{Name: "coupon_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CouponAllocatedsTable holds the schema information for the "coupon_allocateds" table.
	CouponAllocatedsTable = &schema.Table{
		Name:       "coupon_allocateds",
		Columns:    CouponAllocatedsColumns,
		PrimaryKey: []*schema.Column{CouponAllocatedsColumns[0]},
	}
	// CouponPoolsColumns holds the columns for the "coupon_pools" table.
	CouponPoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "denomination", Type: field.TypeUint64},
		{Name: "circulation", Type: field.TypeInt32},
		{Name: "release_by_user_id", Type: field.TypeUUID},
		{Name: "start", Type: field.TypeUint32},
		{Name: "duration_days", Type: field.TypeInt32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString, Size: 512},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CouponPoolsTable holds the schema information for the "coupon_pools" table.
	CouponPoolsTable = &schema.Table{
		Name:       "coupon_pools",
		Columns:    CouponPoolsColumns,
		PrimaryKey: []*schema.Column{CouponPoolsColumns[0]},
	}
	// DefaultKpiSettingsColumns holds the columns for the "default_kpi_settings" table.
	DefaultKpiSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "percent", Type: field.TypeInt32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// DefaultKpiSettingsTable holds the schema information for the "default_kpi_settings" table.
	DefaultKpiSettingsTable = &schema.Table{
		Name:       "default_kpi_settings",
		Columns:    DefaultKpiSettingsColumns,
		PrimaryKey: []*schema.Column{DefaultKpiSettingsColumns[0]},
	}
	// DiscountPoolsColumns holds the columns for the "discount_pools" table.
	DiscountPoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "discount", Type: field.TypeUint32},
		{Name: "release_by_user_id", Type: field.TypeUUID},
		{Name: "start", Type: field.TypeUint32},
		{Name: "duration_days", Type: field.TypeInt32},
		{Name: "message", Type: field.TypeString, Size: 512},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// DiscountPoolsTable holds the schema information for the "discount_pools" table.
	DiscountPoolsTable = &schema.Table{
		Name:       "discount_pools",
		Columns:    DiscountPoolsColumns,
		PrimaryKey: []*schema.Column{DiscountPoolsColumns[0]},
	}
	// NewUserRewardSettingsColumns holds the columns for the "new_user_reward_settings" table.
	NewUserRewardSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Unique: true},
		{Name: "registration_coupon_id", Type: field.TypeUUID, Unique: true},
		{Name: "kyc_coupon_id", Type: field.TypeUUID, Unique: true},
		{Name: "auto_generate_invitation_code", Type: field.TypeBool, Default: true},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// NewUserRewardSettingsTable holds the schema information for the "new_user_reward_settings" table.
	NewUserRewardSettingsTable = &schema.Table{
		Name:       "new_user_reward_settings",
		Columns:    NewUserRewardSettingsColumns,
		PrimaryKey: []*schema.Column{NewUserRewardSettingsColumns[0]},
	}
	// PurchaseInvitationsColumns holds the columns for the "purchase_invitations" table.
	PurchaseInvitationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "invitation_code_id", Type: field.TypeUUID},
		{Name: "fulfilled", Type: field.TypeBool, Default: false},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// PurchaseInvitationsTable holds the schema information for the "purchase_invitations" table.
	PurchaseInvitationsTable = &schema.Table{
		Name:       "purchase_invitations",
		Columns:    PurchaseInvitationsColumns,
		PrimaryKey: []*schema.Column{PurchaseInvitationsColumns[0]},
	}
	// RegistrationInvitationsColumns holds the columns for the "registration_invitations" table.
	RegistrationInvitationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
		{Name: "inviter_id", Type: field.TypeUUID},
		{Name: "invitee_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "fulfilled", Type: field.TypeBool, Default: false},
	}
	// RegistrationInvitationsTable holds the schema information for the "registration_invitations" table.
	RegistrationInvitationsTable = &schema.Table{
		Name:       "registration_invitations",
		Columns:    RegistrationInvitationsColumns,
		PrimaryKey: []*schema.Column{RegistrationInvitationsColumns[0]},
	}
	// UserInvitationCodesColumns holds the columns for the "user_invitation_codes" table.
	UserInvitationCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "invitation_code", Type: field.TypeString, Unique: true},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UserInvitationCodesTable holds the schema information for the "user_invitation_codes" table.
	UserInvitationCodesTable = &schema.Table{
		Name:       "user_invitation_codes",
		Columns:    UserInvitationCodesColumns,
		PrimaryKey: []*schema.Column{UserInvitationCodesColumns[0]},
	}
	// UserKpiSettingsColumns holds the columns for the "user_kpi_settings" table.
	UserKpiSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "percent", Type: field.TypeInt32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UserKpiSettingsTable holds the schema information for the "user_kpi_settings" table.
	UserKpiSettingsTable = &schema.Table{
		Name:       "user_kpi_settings",
		Columns:    UserKpiSettingsColumns,
		PrimaryKey: []*schema.Column{UserKpiSettingsColumns[0]},
	}
	// UserSpecialReductionsColumns holds the columns for the "user_special_reductions" table.
	UserSpecialReductionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "release_by_user_id", Type: field.TypeUUID},
		{Name: "start", Type: field.TypeUint32},
		{Name: "duration_days", Type: field.TypeInt32},
		{Name: "message", Type: field.TypeString, Size: 512},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UserSpecialReductionsTable holds the schema information for the "user_special_reductions" table.
	UserSpecialReductionsTable = &schema.Table{
		Name:       "user_special_reductions",
		Columns:    UserSpecialReductionsColumns,
		PrimaryKey: []*schema.Column{UserSpecialReductionsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgencySettingsTable,
		AppCouponSettingsTable,
		CouponAllocatedsTable,
		CouponPoolsTable,
		DefaultKpiSettingsTable,
		DiscountPoolsTable,
		NewUserRewardSettingsTable,
		PurchaseInvitationsTable,
		RegistrationInvitationsTable,
		UserInvitationCodesTable,
		UserKpiSettingsTable,
		UserSpecialReductionsTable,
	}
)

func init() {
}
