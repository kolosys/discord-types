package payloads

import "github.com/kolosys/discord-types/discord"

// Entitlement represents a Discord entitlement.
//
// See: https://discord.com/developers/docs/monetization/entitlements#entitlement-object-entitlement-structure
type Entitlement struct {
	// ID is the ID of the entitlement.
	ID discord.Snowflake `json:"id"`

	// SkuID is the ID of the SKU.
	SkuID discord.Snowflake `json:"sku_id"`

	// UserID is the ID of the user that is granted access to the entitlement's sku.
	UserID *discord.Snowflake `json:"user_id,omitempty"`

	// GuildID is the ID of the guild that is granted access to the entitlement's sku.
	GuildID *discord.Snowflake `json:"guild_id,omitempty"`

	// ApplicationID is the ID of the parent application.
	ApplicationID discord.Snowflake `json:"application_id"`

	// Type is the type of entitlement.
	Type EntitlementType `json:"type"`

	// Deleted indicates whether the entitlement was deleted.
	Deleted bool `json:"deleted"`

	// StartsAt is the start date at which the entitlement is valid.
	StartsAt *string `json:"starts_at"`

	// EndsAt is the date at which the entitlement is no longer valid.
	EndsAt *string `json:"ends_at"`

	// Consumed indicates whether the entitlement has been consumed (for consumable items).
	Consumed *bool `json:"consumed,omitempty"`
}

// EntitlementType represents the type of entitlement.
//
// See: https://discord.com/developers/docs/monetization/entitlements#entitlement-object-entitlement-types
type EntitlementType int

// Entitlement type constants
const (
	// EntitlementTypePurchase indicates the entitlement was purchased by user.
	EntitlementTypePurchase EntitlementType = iota + 1

	// EntitlementTypePremiumSubscription indicates the entitlement for Discord Nitro subscription.
	EntitlementTypePremiumSubscription

	// EntitlementTypeDeveloperGift indicates the entitlement was gifted by developer.
	EntitlementTypeDeveloperGift

	// EntitlementTypeTestModePurchase indicates the entitlement was purchased by a dev in application test mode.
	EntitlementTypeTestModePurchase

	// EntitlementTypeFreePurchase indicates the entitlement was granted when the SKU was free.
	EntitlementTypeFreePurchase

	// EntitlementTypeUserGift indicates the entitlement was gifted by another user.
	EntitlementTypeUserGift

	// EntitlementTypePremiumPurchase indicates the entitlement was claimed by user for free as a Nitro Subscriber.
	EntitlementTypePremiumPurchase

	// EntitlementTypeApplicationSubscription indicates the entitlement was purchased as an app subscription.
	EntitlementTypeApplicationSubscription
)

// SKU represents a Discord SKU (Stock Keeping Unit).
//
// See: https://discord.com/developers/docs/monetization/skus#sku-object-sku-structure
type SKU struct {
	// ID is the ID of SKU.
	ID discord.Snowflake `json:"id"`

	// Type is the type of SKU.
	Type SKUType `json:"type"`

	// ApplicationID is the ID of the parent application.
	ApplicationID discord.Snowflake `json:"application_id"`

	// Name is the customer-facing name of your premium offering.
	Name string `json:"name"`

	// Slug is the system-generated URL slug based on the SKU's name.
	Slug string `json:"slug"`

	// Flags are the SKU flags combined as a bitfield.
	Flags SKUFlags `json:"flags"`
}

// SKUFlags represents SKU flags.
//
// See: https://discord.com/developers/docs/monetization/skus#sku-object-sku-flags
type SKUFlags int

// SKU flag constants
const (
	// SKUFlagAvailable indicates the SKU is available for purchase.
	SKUFlagAvailable SKUFlags = 1 << 2

	// SKUFlagGuildSubscription indicates recurring SKU that can be purchased by a user and applied to a single server.
	// Grants access to every user in that server.
	SKUFlagGuildSubscription SKUFlags = 1 << 7

	// SKUFlagUserSubscription indicates recurring SKU purchased by a user for themselves.
	// Grants access to the purchasing user in every server.
	SKUFlagUserSubscription SKUFlags = 1 << 8
)

// SKUType represents the type of SKU.
//
// See: https://discord.com/developers/docs/resources/sku#sku-object-sku-types
type SKUType int

// SKU type constants
const (
	// SKUTypeDurable represents a durable one-time purchase.
	SKUTypeDurable SKUType = iota + 2

	// SKUTypeConsumable represents a consumable one-time purchase.
	SKUTypeConsumable

	_ // Skip 4

	// SKUTypeSubscription represents a recurring subscription.
	SKUTypeSubscription

	// SKUTypeSubscriptionGroup represents a system-generated group for each Subscription SKU created.
	SKUTypeSubscriptionGroup
)

// Subscription represents a Discord subscription.
//
// See: https://discord.com/developers/docs/resources/subscription#subscription-object
type Subscription struct {
	// ID is the ID of the subscription.
	ID discord.Snowflake `json:"id"`

	// UserID is the ID of the user who is subscribed.
	UserID discord.Snowflake `json:"user_id"`

	// SkuIDs is the list of SKUs subscribed to.
	SkuIDs []discord.Snowflake `json:"sku_ids"`

	// EntitlementIDs is the list of entitlements granted for this subscription.
	EntitlementIDs []discord.Snowflake `json:"entitlement_ids"`

	// RenewalSkuIDs is the list of SKUs that this user will be subscribed to at renewal.
	RenewalSkuIDs []discord.Snowflake `json:"renewal_sku_ids"`

	// CurrentPeriodStart is the start of the current subscription period.
	CurrentPeriodStart string `json:"current_period_start"`

	// CurrentPeriodEnd is the end of the current subscription period.
	CurrentPeriodEnd string `json:"current_period_end"`

	// Status is the current status of the subscription.
	Status SubscriptionStatus `json:"status"`

	// CanceledAt is when the subscription was canceled.
	CanceledAt *string `json:"canceled_at"`

	// Country is the ISO3166-1 alpha-2 country code of the payment source used to purchase the subscription.
	// Missing unless queried with a private OAuth scope.
	Country *string `json:"country,omitempty"`
}

// SubscriptionStatus represents the status of a subscription.
//
// See: https://discord.com/developers/docs/resources/subscription#subscription-statuses
type SubscriptionStatus int

// Subscription status constants
const (
	// SubscriptionStatusActive indicates the subscription is active and scheduled to renew.
	SubscriptionStatusActive SubscriptionStatus = iota

	// SubscriptionStatusEnding indicates the subscription is active but will not renew.
	SubscriptionStatusEnding

	// SubscriptionStatusInactive indicates the subscription is inactive and not being charged.
	SubscriptionStatusInactive
)
