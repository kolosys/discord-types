// Package rest provides Discord REST API types and utilities.
//
// This file contains query parameter builders and URL utilities
// for Discord REST API endpoints.
package rest

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/kolosys/discord-types/discord"
)

// GetMessagesQuery represents query parameters for getting channel messages.
type GetMessagesQuery struct {
	// Around gets messages around this message ID.
	Around *discord.Snowflake `url:"around,omitempty"`
	// Before gets messages before this message ID.
	Before *discord.Snowflake `url:"before,omitempty"`
	// After gets messages after this message ID.
	After *discord.Snowflake `url:"after,omitempty"`
	// Limit is the max number of messages to return (1-100, default 50).
	Limit *int `url:"limit,omitempty"`
}

// Note: GetGuildMembersQuery and SearchGuildMembersQuery are now defined in guild_types.go

// GetReactionsQuery represents query parameters for getting message reactions.
type GetReactionsQuery struct {
	// After gets users after this user ID.
	After *discord.Snowflake `url:"after,omitempty"`
	// Limit is the max number of users to return (1-100, default 25).
	Limit *int `url:"limit,omitempty"`
}

// GetAuditLogQuery represents query parameters for getting guild audit logs.
type GetAuditLogQuery struct {
	// UserID filters entries by user that made the changes.
	UserID *discord.Snowflake `url:"user_id,omitempty"`
	// ActionType filters entries by audit log event type.
	ActionType *int `url:"action_type,omitempty"`
	// Before gets entries before this entry ID.
	Before *discord.Snowflake `url:"before,omitempty"`
	// After gets entries after this entry ID.
	After *discord.Snowflake `url:"after,omitempty"`
	// Limit is the max number of entries to return (1-100, default 50).
	Limit *int `url:"limit,omitempty"`
}

// GetBansQuery represents query parameters for getting guild bans.
type GetBansQuery struct {
	// Limit is the max number of bans to return (1-1000, default 1000).
	Limit *int `url:"limit,omitempty"`
	// Before gets bans before this user ID.
	Before *discord.Snowflake `url:"before,omitempty"`
	// After gets bans after this user ID.
	After *discord.Snowflake `url:"after,omitempty"`
}

// GetInvitesQuery represents query parameters for getting invites.
type GetInvitesQuery struct {
	// WithCounts includes approximate member counts in the returned invite object.
	WithCounts *bool `url:"with_counts,omitempty"`
	// WithExpiration includes the expiration date in the returned invite object.
	WithExpiration *bool `url:"with_expiration,omitempty"`
	// GuildScheduledEventID includes the guild scheduled event in the invite.
	GuildScheduledEventID *discord.Snowflake `url:"guild_scheduled_event_id,omitempty"`
}

// GetThreadsQuery represents query parameters for getting threads.
type GetThreadsQuery struct {
	// Before gets threads before this timestamp (ISO8601 timestamp).
	Before *string `url:"before,omitempty"`
	// Limit is the max number of threads to return (1-100, default 100).
	Limit *int `url:"limit,omitempty"`
}

// GetScheduledEventUsersQuery represents query parameters for getting scheduled event users.
type GetScheduledEventUsersQuery struct {
	// Limit is the max number of users to return (1-100, default 100).
	Limit *int `url:"limit,omitempty"`
	// WithMember includes guild member data if it exists.
	WithMember *bool `url:"with_member,omitempty"`
	// Before gets users before this user ID.
	Before *discord.Snowflake `url:"before,omitempty"`
	// After gets users after this user ID.
	After *discord.Snowflake `url:"after,omitempty"`
}

// GetGuildsQuery represents query parameters for getting current user's guilds.
type GetGuildsQuery struct {
	// Before gets guilds before this guild ID.
	Before *discord.Snowflake `url:"before,omitempty"`
	// After gets guilds after this guild ID.
	After *discord.Snowflake `url:"after,omitempty"`
	// Limit is the max number of guilds to return (1-200, default 200).
	Limit *int `url:"limit,omitempty"`
	// WithCounts includes approximate member and presence counts for the guilds.
	WithCounts *bool `url:"with_counts,omitempty"`
}

// GetEntitlementsQuery represents query parameters for getting application entitlements.
type GetEntitlementsQuery struct {
	// UserID filters entitlements by user ID.
	UserID *discord.Snowflake `url:"user_id,omitempty"`
	// SKUIDS filters entitlements by SKU IDs.
	SKUIDS []discord.Snowflake `url:"sku_ids,omitempty"`
	// Before gets entitlements before this entitlement ID.
	Before *discord.Snowflake `url:"before,omitempty"`
	// After gets entitlements after this entitlement ID.
	After *discord.Snowflake `url:"after,omitempty"`
	// Limit is the max number of entitlements to return (1-100, default 100).
	Limit *int `url:"limit,omitempty"`
	// GuildID filters entitlements by guild ID.
	GuildID *discord.Snowflake `url:"guild_id,omitempty"`
	// ExcludeEnded excludes entitlements that have ended.
	ExcludeEnded *bool `url:"exclude_ended,omitempty"`
}

// BuildQueryString converts a query struct to a URL query string.
func BuildQueryString(query interface{}) string {
	values := url.Values{}

	switch q := query.(type) {
	case GetMessagesQuery:
		if q.Around != nil {
			values.Set("around", q.Around.String())
		}
		if q.Before != nil {
			values.Set("before", q.Before.String())
		}
		if q.After != nil {
			values.Set("after", q.After.String())
		}
		if q.Limit != nil {
			values.Set("limit", strconv.Itoa(*q.Limit))
		}
	case GetGuildMembersQuery:
		if q.Limit != nil {
			values.Set("limit", strconv.Itoa(*q.Limit))
		}
		if q.After != nil {
			values.Set("after", q.After.String())
		}
	case SearchGuildMembersQuery:
		values.Set("query", q.Query)
		if q.Limit != nil {
			values.Set("limit", strconv.Itoa(*q.Limit))
		}
	case GetReactionsQuery:
		if q.After != nil {
			values.Set("after", q.After.String())
		}
		if q.Limit != nil {
			values.Set("limit", strconv.Itoa(*q.Limit))
		}
	case GetAuditLogQuery:
		if q.UserID != nil {
			values.Set("user_id", q.UserID.String())
		}
		if q.ActionType != nil {
			values.Set("action_type", strconv.Itoa(*q.ActionType))
		}
		if q.Before != nil {
			values.Set("before", q.Before.String())
		}
		if q.After != nil {
			values.Set("after", q.After.String())
		}
		if q.Limit != nil {
			values.Set("limit", strconv.Itoa(*q.Limit))
		}
	case GetBansQuery:
		if q.Limit != nil {
			values.Set("limit", strconv.Itoa(*q.Limit))
		}
		if q.Before != nil {
			values.Set("before", q.Before.String())
		}
		if q.After != nil {
			values.Set("after", q.After.String())
		}
	case GetInvitesQuery:
		if q.WithCounts != nil {
			values.Set("with_counts", strconv.FormatBool(*q.WithCounts))
		}
		if q.WithExpiration != nil {
			values.Set("with_expiration", strconv.FormatBool(*q.WithExpiration))
		}
		if q.GuildScheduledEventID != nil {
			values.Set("guild_scheduled_event_id", q.GuildScheduledEventID.String())
		}
	case GetThreadsQuery:
		if q.Before != nil {
			values.Set("before", *q.Before)
		}
		if q.Limit != nil {
			values.Set("limit", strconv.Itoa(*q.Limit))
		}
	case GetScheduledEventUsersQuery:
		if q.Limit != nil {
			values.Set("limit", strconv.Itoa(*q.Limit))
		}
		if q.WithMember != nil {
			values.Set("with_member", strconv.FormatBool(*q.WithMember))
		}
		if q.Before != nil {
			values.Set("before", q.Before.String())
		}
		if q.After != nil {
			values.Set("after", q.After.String())
		}
	case GetGuildsQuery:
		if q.Before != nil {
			values.Set("before", q.Before.String())
		}
		if q.After != nil {
			values.Set("after", q.After.String())
		}
		if q.Limit != nil {
			values.Set("limit", strconv.Itoa(*q.Limit))
		}
		if q.WithCounts != nil {
			values.Set("with_counts", strconv.FormatBool(*q.WithCounts))
		}
	case GetEntitlementsQuery:
		if q.UserID != nil {
			values.Set("user_id", q.UserID.String())
		}
		if len(q.SKUIDS) > 0 {
			var skuStrs []string
			for _, sku := range q.SKUIDS {
				skuStrs = append(skuStrs, sku.String())
			}
			values.Set("sku_ids", strings.Join(skuStrs, ","))
		}
		if q.Before != nil {
			values.Set("before", q.Before.String())
		}
		if q.After != nil {
			values.Set("after", q.After.String())
		}
		if q.Limit != nil {
			values.Set("limit", strconv.Itoa(*q.Limit))
		}
		if q.GuildID != nil {
			values.Set("guild_id", q.GuildID.String())
		}
		if q.ExcludeEnded != nil {
			values.Set("exclude_ended", strconv.FormatBool(*q.ExcludeEnded))
		}
	case CDNQuery:
		if q.Size != nil {
			values.Set("size", strconv.Itoa(int(*q.Size)))
		}
	}

	if len(values) == 0 {
		return ""
	}
	return "?" + values.Encode()
}

// BuildURL constructs a complete URL from a base URL, route, and optional query parameters.
func BuildURL(base, route string, query interface{}) string {
	url := base + route
	if query != nil {
		queryStr := BuildQueryString(query)
		if queryStr != "" {
			url += queryStr
		}
	}
	return url
}

// Common helper functions for building full URLs

// BuildAPIURL constructs a full Discord API URL.
func BuildAPIURL(route string, query interface{}) string {
	return BuildURL(RouteBases.API, route, query)
}

// BuildCDNURL constructs a full Discord CDN URL.
func BuildCDNURL(route string, query interface{}) string {
	return BuildURL(RouteBases.CDN, route, query)
}

// Helper functions for common query patterns

// NewLimit creates a pointer to an int for limit parameters.
func NewLimit(limit int) *int {
	return &limit
}

// NewBool creates a pointer to a bool for optional boolean parameters.
func NewBool(b bool) *bool {
	return &b
}

// NewSnowflake creates a pointer to a Snowflake for optional ID parameters.
func NewSnowflake(id discord.Snowflake) *discord.Snowflake {
	return &id
}

// NewString creates a pointer to a string for optional string parameters.
func NewString(s string) *string {
	return &s
}

// ImageSize helpers for common sizes
var (
	// Common image sizes as pointers for easy use
	Size16   = &[]ImageSize{ImageSize16}[0]
	Size32   = &[]ImageSize{ImageSize32}[0]
	Size64   = &[]ImageSize{ImageSize64}[0]
	Size128  = &[]ImageSize{ImageSize128}[0]
	Size256  = &[]ImageSize{ImageSize256}[0]
	Size512  = &[]ImageSize{ImageSize512}[0]
	Size1024 = &[]ImageSize{ImageSize1024}[0]
	Size2048 = &[]ImageSize{ImageSize2048}[0]
	Size4096 = &[]ImageSize{ImageSize4096}[0]
)

// Default user avatar index calculation helpers

// GetDefaultAvatarIndex calculates the default avatar index for a user.
// For users on the new username system, uses (user.id >> 22) % 6.
// For users on the legacy username system, uses discriminator % 5.
func GetDefaultAvatarIndex(userID discord.Snowflake, discriminator *int) DefaultUserAvatarAssets {
	if discriminator != nil && *discriminator != 0 {
		// Legacy username system
		return DefaultUserAvatarAssets(*discriminator % 5)
	}
	// New username system
	id, err := userID.Int64()
	if err != nil {
		return DefaultUserAvatar0 // Fallback to default
	}
	return DefaultUserAvatarAssets((id >> 22) % 6)
}

// StickerPackApplicationID is the application ID for Discord's sticker packs.
const StickerPackApplicationID = "710982414301790216"
