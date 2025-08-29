// Package rest provides Discord REST API types and utilities.
//
// This file contains comprehensive guild-related REST API request and response types.
package rest

import (
	"github.com/kolosys/discord-types/discord"
	"github.com/kolosys/discord-types/payloads"
)

// ====================
// Guild Types - Complete Implementation
// ====================

// GuildCreateOverwrite represents a permission overwrite for guild creation
type GuildCreateOverwrite struct {
	ID    interface{}            `json:"id"` // Can be number or string placeholder
	Type  payloads.OverwriteType `json:"type"`
	Allow *discord.Permissions   `json:"allow,omitempty"`
	Deny  *discord.Permissions   `json:"deny,omitempty"`
}

// GuildCreatePartialChannel represents a partial channel for guild creation
type GuildCreatePartialChannel struct {
	Name                          string                              `json:"name"`
	ID                            *interface{}                        `json:"id,omitempty"` // Can be number or string placeholder
	Type                          *payloads.ChannelType               `json:"type,omitempty"`
	Topic                         *string                             `json:"topic,omitempty"`
	Bitrate                       *int                                `json:"bitrate,omitempty"`
	UserLimit                     *int                                `json:"user_limit,omitempty"`
	RateLimitPerUser              *int                                `json:"rate_limit_per_user,omitempty"`
	Position                      *int                                `json:"position,omitempty"`
	PermissionOverwrites          []GuildCreateOverwrite              `json:"permission_overwrites,omitempty"`
	ParentID                      *interface{}                        `json:"parent_id,omitempty"` // Can be number or string placeholder
	NSFW                          *bool                               `json:"nsfw,omitempty"`
	RTCRegion                     *string                             `json:"rtc_region,omitempty"`
	VideoQualityMode              *payloads.VideoQualityMode          `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration    *payloads.ThreadAutoArchiveDuration `json:"default_auto_archive_duration,omitempty"`
	DefaultReactionEmoji          *payloads.DefaultReaction           `json:"default_reaction_emoji,omitempty"`
	AvailableTags                 []payloads.ForumTag                 `json:"available_tags,omitempty"`
	DefaultSortOrder              *payloads.SortOrderType             `json:"default_sort_order,omitempty"`
	DefaultForumLayout            *payloads.ForumLayoutType           `json:"default_forum_layout,omitempty"`
	DefaultThreadRateLimitPerUser *int                                `json:"default_thread_rate_limit_per_user,omitempty"`
	Flags                         *payloads.ChannelFlags              `json:"flags,omitempty"`
}

// GuildCreateRole represents a role for guild creation
type GuildCreateRole struct {
	ID           interface{}          `json:"id"` // Integer placeholder
	Name         *string              `json:"name,omitempty"`
	Color        *int                 `json:"color,omitempty"`
	Hoist        *bool                `json:"hoist,omitempty"`
	Mentionable  *bool                `json:"mentionable,omitempty"`
	Permissions  *discord.Permissions `json:"permissions,omitempty"`
	Icon         *string              `json:"icon,omitempty"`
	UnicodeEmoji *string              `json:"unicode_emoji,omitempty"`
}

// PostGuildJSONBody represents the request body for POST /guilds (DEPRECATED)
type PostGuildJSONBody struct {
	// Name of the guild (2-100 characters)
	Name string `json:"name"`
	// Voice region id
	Region *string `json:"region,omitempty"`
	// base64 1024x1024 png/jpeg image for the guild icon
	Icon *string `json:"icon,omitempty"`
	// Verification level
	VerificationLevel *payloads.GuildVerificationLevel `json:"verification_level,omitempty"`
	// Default message notification level
	DefaultMessageNotifications *payloads.GuildDefaultMessageNotifications `json:"default_message_notifications,omitempty"`
	// Explicit content filter level
	ExplicitContentFilter *payloads.GuildExplicitContentFilter `json:"explicit_content_filter,omitempty"`
	// New guild roles
	Roles []GuildCreateRole `json:"roles,omitempty"`
	// New guild's channels
	Channels []GuildCreatePartialChannel `json:"channels,omitempty"`
	// ID for afk channel
	AFKChannelID *discord.Snowflake `json:"afk_channel_id,omitempty"`
	// afk timeout in seconds (60, 300, 900, 1800, 3600)
	AFKTimeout *int `json:"afk_timeout,omitempty"`
	// System channel id
	SystemChannelID *discord.Snowflake `json:"system_channel_id,omitempty"`
	// System channel flags
	SystemChannelFlags *payloads.GuildSystemChannelFlags `json:"system_channel_flags,omitempty"`
	// Whether the boosts progress bar should be enabled
	PremiumProgressBarEnabled *bool `json:"premium_progress_bar_enabled,omitempty"`
}

// PostGuildResult represents the response from POST /guilds (DEPRECATED)
type PostGuildResult = payloads.Guild

// PostGuildMFAJSONBody represents the request body for POST /guilds/{guild.id}/mfa (DEPRECATED)
type PostGuildMFAJSONBody struct {
	// MFA level
	Level payloads.GuildMFALevel `json:"level"`
}

// PostGuildMFAResult represents the response from POST /guilds/{guild.id}/mfa (DEPRECATED)
type PostGuildMFAResult = PostGuildMFAJSONBody

// GetGuildQuery represents query parameters for GET /guilds/{guild.id}
type GetGuildQuery struct {
	// When true, will return approximate member and presence counts for the guild
	WithCounts *bool `json:"with_counts,omitempty"`
}

// GetGuildResult represents the response from GET /guilds/{guild.id}
type GetGuildResult = payloads.Guild

// GetGuildPreviewResult represents the response from GET /guilds/{guild.id}/preview
type GetGuildPreviewResult = GuildPreview

// PatchGuildJSONBody represents the request body for PATCH /guilds/{guild.id}
type PatchGuildJSONBody struct {
	// New name for the guild (2-100 characters)
	Name *string `json:"name,omitempty"`
	// Voice region id
	Region *string `json:"region,omitempty"`
	// Verification level
	VerificationLevel *payloads.GuildVerificationLevel `json:"verification_level,omitempty"`
	// Default message notification level
	DefaultMessageNotifications *payloads.GuildDefaultMessageNotifications `json:"default_message_notifications,omitempty"`
	// Explicit content filter level
	ExplicitContentFilter *payloads.GuildExplicitContentFilter `json:"explicit_content_filter,omitempty"`
	// ID for afk channel
	AFKChannelID *discord.Snowflake `json:"afk_channel_id,omitempty"`
	// afk timeout in seconds
	AFKTimeout *int `json:"afk_timeout,omitempty"`
	// base64 1024x1024 png/jpeg image for the guild icon
	Icon *string `json:"icon,omitempty"`
	// User id to transfer guild ownership to (must be owner)
	OwnerID *discord.Snowflake `json:"owner_id,omitempty"`
	// base64 16:9 png/jpeg image for the guild splash
	Splash *string `json:"splash,omitempty"`
	// base64 16:9 png/jpeg image for the guild discovery splash
	DiscoverySplash *string `json:"discovery_splash,omitempty"`
	// base64 png/jpeg image for the guild banner
	Banner *string `json:"banner,omitempty"`
	// System channel id
	SystemChannelID *discord.Snowflake `json:"system_channel_id,omitempty"`
	// System channel flags
	SystemChannelFlags *payloads.GuildSystemChannelFlags `json:"system_channel_flags,omitempty"`
	// The id of the channel where Community guilds display rules and/or guidelines
	RulesChannelID *discord.Snowflake `json:"rules_channel_id,omitempty"`
	// The id of the channel where admins and moderators receive notices from Discord
	PublicUpdatesChannelID *discord.Snowflake `json:"public_updates_channel_id,omitempty"`
	// The preferred locale of a Community guild
	PreferredLocale *string `json:"preferred_locale,omitempty"`
	// Guild features
	Features []payloads.GuildFeature `json:"features,omitempty"`
	// The description for the guild
	Description *string `json:"description,omitempty"`
	// Whether the boosts progress bar should be enabled
	PremiumProgressBarEnabled *bool `json:"premium_progress_bar_enabled,omitempty"`
	// The id of the channel where admins and moderators receive safety alerts from Discord
	SafetyAlertsChannelID *discord.Snowflake `json:"safety_alerts_channel_id,omitempty"`
}

// PatchGuildResult represents the response from PATCH /guilds/{guild.id}
type PatchGuildResult = payloads.Guild

// DeleteGuildResult represents the response from DELETE /guilds/{guild.id}
type DeleteGuildResult struct{} // Empty response

// GetGuildChannelsResult represents the response from GET /guilds/{guild.id}/channels
type GetGuildChannelsResult = []payloads.GuildTextChannel

// PostGuildChannelJSONBody represents the request body for POST /guilds/{guild.id}/channels
type PostGuildChannelJSONBody struct {
	// Channel name (1-100 characters)
	Name string `json:"name"`
	// The type of channel
	Type *payloads.ChannelType `json:"type,omitempty"`
	// Channel topic (0-1024 characters)
	Topic *string `json:"topic,omitempty"`
	// The bitrate (in bits) of the voice or stage channel
	Bitrate *int `json:"bitrate,omitempty"`
	// The user limit of the voice channel
	UserLimit *int `json:"user_limit,omitempty"`
	// Amount of seconds a user has to wait before sending another message
	RateLimitPerUser *int `json:"rate_limit_per_user,omitempty"`
	// Sorting position of the channel
	Position *int `json:"position,omitempty"`
	// The channel's permission overwrites
	PermissionOverwrites []ChannelPatchOverwrite `json:"permission_overwrites,omitempty"`
	// ID of the parent category for a channel
	ParentID *discord.Snowflake `json:"parent_id,omitempty"`
	// Whether the channel is nsfw
	NSFW *bool `json:"nsfw,omitempty"`
	// Voice region id for the voice or stage channel
	RTCRegion *string `json:"rtc_region,omitempty"`
	// The camera video quality mode of the voice channel
	VideoQualityMode *payloads.VideoQualityMode `json:"video_quality_mode,omitempty"`
	// The default auto-archive duration for threads
	DefaultAutoArchiveDuration *payloads.ThreadAutoArchiveDuration `json:"default_auto_archive_duration,omitempty"`
	// Channel flags combined as a bitfield
	Flags *payloads.ChannelFlags `json:"flags,omitempty"`
	// The set of tags that can be used in a forum channel
	AvailableTags []payloads.ForumTag `json:"available_tags,omitempty"`
	// The emoji to show in the add reaction button
	DefaultReactionEmoji *payloads.DefaultReaction `json:"default_reaction_emoji,omitempty"`
	// The initial rate_limit_per_user to set on newly created threads
	DefaultThreadRateLimitPerUser *int `json:"default_thread_rate_limit_per_user,omitempty"`
	// The default sort order type used to order posts in forum channels
	DefaultSortOrder *payloads.SortOrderType `json:"default_sort_order,omitempty"`
	// The default forum layout view used to display posts
	DefaultForumLayout *payloads.ForumLayoutType `json:"default_forum_layout,omitempty"`
}

// PostGuildChannelResult represents the response from POST /guilds/{guild.id}/channels
type PostGuildChannelResult = payloads.GuildTextChannel

// PatchGuildChannelPositionsJSONBody represents the request body for PATCH /guilds/{guild.id}/channels
type PatchGuildChannelPositionsJSONBody = []PatchGuildChannelPositionJSONBody

// PatchGuildChannelPositionJSONBody represents a single channel position update
type PatchGuildChannelPositionJSONBody struct {
	// Channel id
	ID discord.Snowflake `json:"id"`
	// Sorting position of the channel
	Position *int `json:"position,omitempty"`
	// Syncs the permission overwrites with the new parent, if moving to a new category
	LockPermissions *bool `json:"lock_permissions,omitempty"`
	// The new parent ID for the channel that is moved
	ParentID *discord.Snowflake `json:"parent_id,omitempty"`
}

// GetGuildActiveThreadsResult represents the response from GET /guilds/{guild.id}/threads/active
type GetGuildActiveThreadsResult = ThreadList

// GetGuildMemberQuery represents query parameters for GET /guilds/{guild.id}/members/{user.id}
type GetGuildMemberQuery struct {
	// No query parameters for this endpoint
}

// GetGuildMemberResult represents the response from GET /guilds/{guild.id}/members/{user.id}
type GetGuildMemberResult = payloads.GuildMember

// GetGuildMembersQuery represents query parameters for GET /guilds/{guild.id}/members
type GetGuildMembersQuery struct {
	// Max number of members to return (1-1000, default 1)
	Limit *int `json:"limit,omitempty"`
	// The highest user id in the previous page
	After *discord.Snowflake `json:"after,omitempty"`
}

// GetGuildMembersResult represents the response from GET /guilds/{guild.id}/members
type GetGuildMembersResult = []payloads.GuildMember

// SearchGuildMembersQuery represents query parameters for GET /guilds/{guild.id}/members/search
type SearchGuildMembersQuery struct {
	// Query string to match username(s) and nickname(s) against
	Query string `json:"query"`
	// Max number of members to return (1-1000, default 1)
	Limit *int `json:"limit,omitempty"`
}

// SearchGuildMembersResult represents the response from GET /guilds/{guild.id}/members/search
type SearchGuildMembersResult = []payloads.GuildMember

// PutGuildMemberJSONBody represents the request body for PUT /guilds/{guild.id}/members/{user.id}
type PutGuildMemberJSONBody struct {
	// An oauth2 access token granted with the guilds.join scope
	AccessToken string `json:"access_token"`
	// Value to set users nickname to
	Nick *string `json:"nick,omitempty"`
	// Array of role ids the member is assigned
	Roles []discord.Snowflake `json:"roles,omitempty"`
	// Whether the user is muted in voice channels
	Mute *bool `json:"mute,omitempty"`
	// Whether the user is deafened in voice channels
	Deaf *bool `json:"deaf,omitempty"`
}

// PutGuildMemberResult represents the response from PUT /guilds/{guild.id}/members/{user.id}
type PutGuildMemberResult = payloads.GuildMember

// PatchGuildMemberJSONBody represents the request body for PATCH /guilds/{guild.id}/members/{user.id}
type PatchGuildMemberJSONBody struct {
	// Value to set users nickname to
	Nick *string `json:"nick,omitempty"`
	// Array of role ids the member is assigned
	Roles []discord.Snowflake `json:"roles,omitempty"`
	// Whether the user is muted in voice channels
	Mute *bool `json:"mute,omitempty"`
	// Whether the user is deafened in voice channels
	Deaf *bool `json:"deaf,omitempty"`
	// ID of channel to move user to (if they are connected to voice)
	ChannelID *discord.Snowflake `json:"channel_id,omitempty"`
	// Timestamp when the time out will be removed
	CommunicationDisabledUntil *string `json:"communication_disabled_until,omitempty"`
	// Guild member flags
	Flags *payloads.GuildMemberFlags `json:"flags,omitempty"`
}

// PatchGuildMemberResult represents the response from PATCH /guilds/{guild.id}/members/{user.id}
type PatchGuildMemberResult = payloads.GuildMember

// PatchCurrentGuildMemberJSONBody represents the request body for PATCH /guilds/{guild.id}/members/@me
type PatchCurrentGuildMemberJSONBody struct {
	// Value to set users nickname to
	Nick *string `json:"nick,omitempty"`
}

// PatchCurrentGuildMemberResult represents the response from PATCH /guilds/{guild.id}/members/@me
type PatchCurrentGuildMemberResult = payloads.GuildMember

// PutGuildMemberRoleResult represents the response from PUT /guilds/{guild.id}/members/{user.id}/roles/{role.id}
type PutGuildMemberRoleResult struct{} // Empty response

// DeleteGuildMemberRoleResult represents the response from DELETE /guilds/{guild.id}/members/{user.id}/roles/{role.id}
type DeleteGuildMemberRoleResult struct{} // Empty response

// DeleteGuildMemberResult represents the response from DELETE /guilds/{guild.id}/members/{user.id}
type DeleteGuildMemberResult struct{} // Empty response

// GetGuildBansQuery represents query parameters for GET /guilds/{guild.id}/bans
type GetGuildBansQuery struct {
	// Consider only users before given user id
	Before *discord.Snowflake `json:"before,omitempty"`
	// Consider only users after given user id
	After *discord.Snowflake `json:"after,omitempty"`
	// Number of users to return (1-1000, default 1000)
	Limit *int `json:"limit,omitempty"`
}

// GetGuildBansResult represents the response from GET /guilds/{guild.id}/bans
type GetGuildBansResult = []Ban

// GetGuildBanResult represents the response from GET /guilds/{guild.id}/bans/{user.id}
type GetGuildBanResult = Ban

// PutGuildBanJSONBody represents the request body for PUT /guilds/{guild.id}/bans/{user.id}
type PutGuildBanJSONBody struct {
	// Number of seconds to delete messages for, between 0 and 604800 (7 days)
	DeleteMessageSeconds *int `json:"delete_message_seconds,omitempty"`
}

// PutGuildBanResult represents the response from PUT /guilds/{guild.id}/bans/{user.id}
type PutGuildBanResult struct{} // Empty response

// DeleteGuildBanResult represents the response from DELETE /guilds/{guild.id}/bans/{user.id}
type DeleteGuildBanResult struct{} // Empty response

// PostGuildBulkBanJSONBody represents the request body for POST /guilds/{guild.id}/bulk-ban
type PostGuildBulkBanJSONBody struct {
	// List of user ids to ban (max 200)
	UserIDs []discord.Snowflake `json:"user_ids"`
	// Number of seconds to delete messages for, between 0 and 604800 (7 days)
	DeleteMessageSeconds *int `json:"delete_message_seconds,omitempty"`
}

// BulkBanResponse represents the response from POST /guilds/{guild.id}/bulk-ban
type BulkBanResponse struct {
	// List of user ids that were successfully banned
	BannedUsers []discord.Snowflake `json:"banned_users"`
	// List of user ids that failed to be banned
	FailedUsers []discord.Snowflake `json:"failed_users"`
}

// PostGuildBulkBanResult represents the response from POST /guilds/{guild.id}/bulk-ban
type PostGuildBulkBanResult = BulkBanResponse

// GetGuildRolesResult represents the response from GET /guilds/{guild.id}/roles
type GetGuildRolesResult = []payloads.Role

// PostGuildRoleJSONBody represents the request body for POST /guilds/{guild.id}/roles
type PostGuildRoleJSONBody struct {
	// Name of the role, max 100 characters (default "new role")
	Name *string `json:"name,omitempty"`
	// Bitwise value of the enabled/disabled permissions (default @everyone permissions)
	Permissions *discord.Permissions `json:"permissions,omitempty"`
	// RGB color value (default 0)
	Color *int `json:"color,omitempty"`
	// Whether the role should be displayed separately in the sidebar (default false)
	Hoist *bool `json:"hoist,omitempty"`
	// Role icon hash
	Icon *string `json:"icon,omitempty"`
	// Role unicode emoji
	UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
	// Whether the role should be mentionable (default false)
	Mentionable *bool `json:"mentionable,omitempty"`
}

// PostGuildRoleResult represents the response from POST /guilds/{guild.id}/roles
type PostGuildRoleResult = payloads.Role

// PatchGuildRolePositionsJSONBody represents the request body for PATCH /guilds/{guild.id}/roles
type PatchGuildRolePositionsJSONBody = []PatchGuildRolePositionJSONBody

// PatchGuildRolePositionJSONBody represents a single role position update
type PatchGuildRolePositionJSONBody struct {
	// Role id
	ID discord.Snowflake `json:"id"`
	// Sorting position of the role
	Position *int `json:"position,omitempty"`
}

// PatchGuildRolePositionsResult represents the response from PATCH /guilds/{guild.id}/roles
type PatchGuildRolePositionsResult = []payloads.Role

// PatchGuildRoleJSONBody represents the request body for PATCH /guilds/{guild.id}/roles/{role.id}
type PatchGuildRoleJSONBody struct {
	// Name of the role, max 100 characters
	Name *string `json:"name,omitempty"`
	// Bitwise value of the enabled/disabled permissions
	Permissions *discord.Permissions `json:"permissions,omitempty"`
	// RGB color value
	Color *int `json:"color,omitempty"`
	// Whether the role should be displayed separately in the sidebar
	Hoist *bool `json:"hoist,omitempty"`
	// Role icon hash
	Icon *string `json:"icon,omitempty"`
	// Role unicode emoji
	UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
	// Whether the role should be mentionable
	Mentionable *bool `json:"mentionable,omitempty"`
}

// PatchGuildRoleResult represents the response from PATCH /guilds/{guild.id}/roles/{role.id}
type PatchGuildRoleResult = payloads.Role

// DeleteGuildRoleResult represents the response from DELETE /guilds/{guild.id}/roles/{role.id}
type DeleteGuildRoleResult struct{} // Empty response

// GetGuildPruneCountQuery represents query parameters for GET /guilds/{guild.id}/prune
type GetGuildPruneCountQuery struct {
	// Number of days to prune (1-30, default 7)
	Days *int `json:"days,omitempty"`
	// Role(s) to include
	IncludeRoles []discord.Snowflake `json:"include_roles,omitempty"`
}

// GuildPruneCountResponse represents the response from GET /guilds/{guild.id}/prune
type GuildPruneCountResponse struct {
	Pruned *int `json:"pruned"`
}

// GetGuildPruneCountResult represents the response from GET /guilds/{guild.id}/prune
type GetGuildPruneCountResult = GuildPruneCountResponse

// PostGuildPruneJSONBody represents the request body for POST /guilds/{guild.id}/prune
type PostGuildPruneJSONBody struct {
	// Number of days to prune (1-30, default 7)
	Days *int `json:"days,omitempty"`
	// Whether pruned is returned (discouraged for large guilds, default true)
	ComputePruneCount *bool `json:"compute_prune_count,omitempty"`
	// Role(s) to include
	IncludeRoles []discord.Snowflake `json:"include_roles,omitempty"`
	// Reason for the prune
	Reason *string `json:"reason,omitempty"`
}

// PostGuildPruneResult represents the response from POST /guilds/{guild.id}/prune
type PostGuildPruneResult = GuildPruneCountResponse

// GetGuildVoiceRegionsResult represents the response from GET /guilds/{guild.id}/regions
type GetGuildVoiceRegionsResult = []payloads.VoiceRegion

// GetGuildInvitesResult represents the response from GET /guilds/{guild.id}/invites
type GetGuildInvitesResult = []Invite

// GuildIntegration represents a simplified guild integration
type GuildIntegration struct {
	ID                discord.Snowflake  `json:"id"`
	Name              string             `json:"name"`
	Type              string             `json:"type"`
	Enabled           bool               `json:"enabled"`
	Syncing           *bool              `json:"syncing,omitempty"`
	RoleID            *discord.Snowflake `json:"role_id,omitempty"`
	EnableEmoticons   *bool              `json:"enable_emoticons,omitempty"`
	ExpireBehavior    *int               `json:"expire_behavior,omitempty"`
	ExpireGracePeriod *int               `json:"expire_grace_period,omitempty"`
	User              *payloads.User     `json:"user,omitempty"`
	Account           interface{}        `json:"account,omitempty"`
	SyncedAt          *string            `json:"synced_at,omitempty"`
	SubscriberCount   *int               `json:"subscriber_count,omitempty"`
	Revoked           *bool              `json:"revoked,omitempty"`
	Application       interface{}        `json:"application,omitempty"`
	Scopes            []string           `json:"scopes,omitempty"`
}

// GetGuildIntegrationsResult represents the response from GET /guilds/{guild.id}/integrations
type GetGuildIntegrationsResult = []GuildIntegration

// DeleteGuildIntegrationResult represents the response from DELETE /guilds/{guild.id}/integrations/{integration.id}
type DeleteGuildIntegrationResult struct{} // Empty response

// GuildWidgetSettings represents guild widget settings
type GuildWidgetSettings struct {
	Enabled   bool               `json:"enabled"`
	ChannelID *discord.Snowflake `json:"channel_id"`
}

// GuildWidget represents a guild widget
type GuildWidget struct {
	ID            discord.Snowflake    `json:"id"`
	Name          string               `json:"name"`
	InstantInvite *string              `json:"instant_invite"`
	Channels      []GuildWidgetChannel `json:"channels"`
	Members       []GuildWidgetMember  `json:"members"`
	PresenceCount int                  `json:"presence_count"`
}

// GuildWidgetChannel represents a channel in guild widget
type GuildWidgetChannel struct {
	ID       discord.Snowflake `json:"id"`
	Name     string            `json:"name"`
	Position int               `json:"position"`
}

// GuildWidgetMember represents a member in guild widget
type GuildWidgetMember struct {
	ID            discord.Snowflake `json:"id"`
	Username      string            `json:"username"`
	Discriminator string            `json:"discriminator"`
	Avatar        *string           `json:"avatar"`
	Status        string            `json:"status"`
	AvatarURL     string            `json:"avatar_url"`
}

// GetGuildWidgetSettingsResult represents the response from GET /guilds/{guild.id}/widget
type GetGuildWidgetSettingsResult = GuildWidgetSettings

// PatchGuildWidgetSettingsJSONBody represents the request body for PATCH /guilds/{guild.id}/widget
type PatchGuildWidgetSettingsJSONBody struct {
	// Whether the widget is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The widget channel id
	ChannelID *discord.Snowflake `json:"channel_id,omitempty"`
}

// PatchGuildWidgetSettingsResult represents the response from PATCH /guilds/{guild.id}/widget
type PatchGuildWidgetSettingsResult = GuildWidgetSettings

// GetGuildWidgetResult represents the response from GET /guilds/{guild.id}/widget.json
type GetGuildWidgetResult = GuildWidget

// GetGuildVanityURLResult represents the response from GET /guilds/{guild.id}/vanity-url
type GetGuildVanityURLResult = Invite

// GuildWidgetStyle represents different widget image styles
type GuildWidgetStyle string

const (
	GuildWidgetStyleShield1 GuildWidgetStyle = "shield"
	GuildWidgetStyleBanner1 GuildWidgetStyle = "banner1"
	GuildWidgetStyleBanner2 GuildWidgetStyle = "banner2"
	GuildWidgetStyleBanner3 GuildWidgetStyle = "banner3"
	GuildWidgetStyleBanner4 GuildWidgetStyle = "banner4"
)

// GetGuildWidgetImageQuery represents query parameters for GET /guilds/{guild.id}/widget.png
type GetGuildWidgetImageQuery struct {
	// Style of the widget image returned
	Style *GuildWidgetStyle `json:"style,omitempty"`
}

// GuildWelcomeScreen represents a guild welcome screen
type GuildWelcomeScreen struct {
	Description     *string                     `json:"description"`
	WelcomeChannels []GuildWelcomeScreenChannel `json:"welcome_channels"`
}

// GuildWelcomeScreenChannel represents a welcome screen channel
type GuildWelcomeScreenChannel struct {
	ChannelID   discord.Snowflake  `json:"channel_id"`
	Description string             `json:"description"`
	EmojiID     *discord.Snowflake `json:"emoji_id"`
	EmojiName   *string            `json:"emoji_name"`
}

// GuildOnboarding represents guild onboarding settings
type GuildOnboarding struct {
	GuildID           discord.Snowflake       `json:"guild_id"`
	Prompts           []GuildOnboardingPrompt `json:"prompts"`
	DefaultChannelIDs []discord.Snowflake     `json:"default_channel_ids"`
	Enabled           bool                    `json:"enabled"`
	Mode              GuildOnboardingMode     `json:"mode"`
}

// GuildOnboardingPrompt represents an onboarding prompt
type GuildOnboardingPrompt struct {
	ID           discord.Snowflake             `json:"id"`
	Type         GuildOnboardingPromptType     `json:"type"`
	Options      []GuildOnboardingPromptOption `json:"options"`
	Title        string                        `json:"title"`
	SingleSelect bool                          `json:"single_select"`
	Required     bool                          `json:"required"`
	InOnboarding bool                          `json:"in_onboarding"`
}

// GuildOnboardingPromptOption represents an onboarding prompt option
type GuildOnboardingPromptOption struct {
	ID          discord.Snowflake   `json:"id"`
	ChannelIDs  []discord.Snowflake `json:"channel_ids"`
	RoleIDs     []discord.Snowflake `json:"role_ids"`
	EmojiID     *discord.Snowflake  `json:"emoji_id"`
	EmojiName   *string             `json:"emoji_name"`
	Title       string              `json:"title"`
	Description *string             `json:"description"`
}

// GuildOnboardingMode represents onboarding mode
type GuildOnboardingMode int

const (
	GuildOnboardingModeDefault  GuildOnboardingMode = 0
	GuildOnboardingModeAdvanced GuildOnboardingMode = 1
)

// GuildOnboardingPromptType represents onboarding prompt type
type GuildOnboardingPromptType int

const (
	GuildOnboardingPromptTypeMultipleChoice GuildOnboardingPromptType = 0
	GuildOnboardingPromptTypeDropdown       GuildOnboardingPromptType = 1
)

// GetGuildWelcomeScreenResult represents the response from GET /guilds/{guild.id}/welcome-screen
type GetGuildWelcomeScreenResult = GuildWelcomeScreen

// PatchGuildWelcomeScreenJSONBody represents the request body for PATCH /guilds/{guild.id}/welcome-screen
type PatchGuildWelcomeScreenJSONBody struct {
	// Whether the welcome screen is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Channels linked in the welcome screen and their display options
	WelcomeChannels []GuildWelcomeScreenChannel `json:"welcome_channels,omitempty"`
	// The server description to show in the welcome screen
	Description *string `json:"description,omitempty"`
}

// PatchGuildWelcomeScreenResult represents the response from PATCH /guilds/{guild.id}/welcome-screen
type PatchGuildWelcomeScreenResult = GuildWelcomeScreen

// GetGuildOnboardingResult represents the response from GET /guilds/{guild.id}/onboarding
type GetGuildOnboardingResult = GuildOnboarding

// PutGuildOnboardingJSONBody represents the request body for PUT /guilds/{guild.id}/onboarding
type PutGuildOnboardingJSONBody struct {
	// Prompts shown during onboarding and in customize community
	Prompts []GuildOnboardingPrompt `json:"prompts"`
	// Channel IDs that members get opted into automatically
	DefaultChannelIDs []discord.Snowflake `json:"default_channel_ids"`
	// Whether onboarding is enabled in the guild
	Enabled bool `json:"enabled"`
	// Current mode of onboarding
	Mode GuildOnboardingMode `json:"mode"`
}

// PutGuildOnboardingResult represents the response from PUT /guilds/{guild.id}/onboarding
type PutGuildOnboardingResult = GuildOnboarding
