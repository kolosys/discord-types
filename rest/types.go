// Package rest provides Discord REST API types and utilities.
//
// This file contains request and response payload types for Discord REST API endpoints.
// These types provide complete type safety for REST API interactions.
package rest

import (
	"time"

	"github.com/kolosys/discord-types/discord"
	"github.com/kolosys/discord-types/payloads"
)

// ====================
// Channel Types
// ====================

// GetChannelResponse represents the response from GET /channels/{channel.id}
type GetChannelResponse = payloads.GuildTextChannel

// ModifyChannelRequest represents the request body for PATCH /channels/{channel.id}
type ModifyChannelRequest struct {
	Name                          *string                             `json:"name,omitempty"`
	Type                          *payloads.ChannelType               `json:"type,omitempty"`
	Position                      *int                                `json:"position,omitempty"`
	Topic                         *string                             `json:"topic,omitempty"`
	NSFW                          *bool                               `json:"nsfw,omitempty"`
	RateLimitPerUser              *int                                `json:"rate_limit_per_user,omitempty"`
	Bitrate                       *int                                `json:"bitrate,omitempty"`
	UserLimit                     *int                                `json:"user_limit,omitempty"`
	PermissionOverwrites          []payloads.Overwrite                `json:"permission_overwrites,omitempty"`
	ParentID                      *discord.Snowflake                  `json:"parent_id,omitempty"`
	RTCRegion                     *string                             `json:"rtc_region,omitempty"`
	VideoQualityMode              *payloads.VideoQualityMode          `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration    *payloads.ThreadAutoArchiveDuration `json:"default_auto_archive_duration,omitempty"`
	Flags                         *payloads.ChannelFlags              `json:"flags,omitempty"`
	AvailableTags                 []payloads.ForumTag                 `json:"available_tags,omitempty"`
	DefaultReactionEmoji          *payloads.DefaultReaction           `json:"default_reaction_emoji,omitempty"`
	DefaultThreadRateLimitPerUser *int                                `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultSortOrder              *payloads.SortOrderType             `json:"default_sort_order,omitempty"`
	DefaultForumLayout            *payloads.ForumLayoutType           `json:"default_forum_layout,omitempty"`
}

// GetChannelMessagesResponse represents the response from GET /channels/{channel.id}/messages
type GetChannelMessagesResponse = []payloads.Message

// CreateMessageRequest represents the request body for POST /channels/{channel.id}/messages
type CreateMessageRequest struct {
	Content          *string                       `json:"content,omitempty"`
	Nonce            *string                       `json:"nonce,omitempty"`
	TTS              *bool                         `json:"tts,omitempty"`
	Embeds           []payloads.Embed              `json:"embeds,omitempty"`
	AllowedMentions  *payloads.AllowedMentions     `json:"allowed_mentions,omitempty"`
	MessageReference *payloads.MessageReference    `json:"message_reference,omitempty"`
	Components       []payloads.ActionRowComponent `json:"components,omitempty"`
	StickerIDs       []discord.Snowflake           `json:"sticker_ids,omitempty"`
	Attachments      []payloads.PartialAttachment  `json:"attachments,omitempty"`
	Flags            *payloads.MessageFlags        `json:"flags,omitempty"`
	Poll             *payloads.Poll                `json:"poll,omitempty"`
}

// CreateMessageResponse represents the response from POST /channels/{channel.id}/messages
type CreateMessageResponse = payloads.Message

// GetMessageResponse represents the response from GET /channels/{channel.id}/messages/{message.id}
type GetMessageResponse = payloads.Message

// EditMessageRequest represents the request body for PATCH /channels/{channel.id}/messages/{message.id}
type EditMessageRequest struct {
	Content         *string                       `json:"content,omitempty"`
	Embeds          []payloads.Embed              `json:"embeds,omitempty"`
	Flags           *payloads.MessageFlags        `json:"flags,omitempty"`
	AllowedMentions *payloads.AllowedMentions     `json:"allowed_mentions,omitempty"`
	Components      []payloads.ActionRowComponent `json:"components,omitempty"`
	Attachments     []payloads.PartialAttachment  `json:"attachments,omitempty"`
}

// EditMessageResponse represents the response from PATCH /channels/{channel.id}/messages/{message.id}
type EditMessageResponse = payloads.Message

// BulkDeleteMessagesRequest represents the request body for POST /channels/{channel.id}/messages/bulk-delete
type BulkDeleteMessagesRequest struct {
	Messages []discord.Snowflake `json:"messages"`
}

// EditChannelPermissionsRequest represents the request body for PUT /channels/{channel.id}/permissions/{overwrite.id}
type EditChannelPermissionsRequest struct {
	Allow *discord.Permissions   `json:"allow,omitempty"`
	Deny  *discord.Permissions   `json:"deny,omitempty"`
	Type  payloads.OverwriteType `json:"type"`
}

// InviteTargetType represents the type of invite target
type InviteTargetType int

const (
	InviteTargetTypeStream              InviteTargetType = 1
	InviteTargetTypeEmbeddedApplication InviteTargetType = 2
)

// Invite represents a Discord invite
type Invite struct {
	Code                     string                        `json:"code"`
	Guild                    *payloads.PartialGuild        `json:"guild,omitempty"`
	Channel                  *payloads.PartialChannel      `json:"channel"`
	Inviter                  *payloads.User                `json:"inviter,omitempty"`
	TargetType               *InviteTargetType             `json:"target_type,omitempty"`
	TargetUser               *payloads.User                `json:"target_user,omitempty"`
	TargetApplication        *payloads.Application         `json:"target_application,omitempty"`
	ApproximatePresenceCount *int                          `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount   *int                          `json:"approximate_member_count,omitempty"`
	ExpiresAt                *time.Time                    `json:"expires_at,omitempty"`
	GuildScheduledEvent      *payloads.GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
}

// GetChannelInvitesResponse represents the response from GET /channels/{channel.id}/invites
type GetChannelInvitesResponse = []Invite

// CreateChannelInviteRequest represents the request body for POST /channels/{channel.id}/invites
type CreateChannelInviteRequest struct {
	MaxAge              *int               `json:"max_age,omitempty"`
	MaxUses             *int               `json:"max_uses,omitempty"`
	Temporary           *bool              `json:"temporary,omitempty"`
	Unique              *bool              `json:"unique,omitempty"`
	TargetType          *InviteTargetType  `json:"target_type,omitempty"`
	TargetUserID        *discord.Snowflake `json:"target_user_id,omitempty"`
	TargetApplicationID *discord.Snowflake `json:"target_application_id,omitempty"`
}

// CreateChannelInviteResponse represents the response from POST /channels/{channel.id}/invites
type CreateChannelInviteResponse = Invite

// GetPinnedMessagesResponse represents the response from GET /channels/{channel.id}/messages/pins
type GetPinnedMessagesResponse = []payloads.Message

// ====================
// Guild Types
// ====================

// GetGuildResponse represents the response from GET /guilds/{guild.id}
type GetGuildResponse = payloads.Guild

// GuildPreview represents a preview of a Discord guild
type GuildPreview struct {
	ID                       discord.Snowflake       `json:"id"`
	Name                     string                  `json:"name"`
	Icon                     *string                 `json:"icon"`
	Splash                   *string                 `json:"splash"`
	DiscoverySplash          *string                 `json:"discovery_splash"`
	Emojis                   []payloads.Emoji        `json:"emojis"`
	Features                 []payloads.GuildFeature `json:"features"`
	ApproximateMemberCount   int                     `json:"approximate_member_count"`
	ApproximatePresenceCount int                     `json:"approximate_presence_count"`
	Description              *string                 `json:"description"`
	Stickers                 []payloads.Sticker      `json:"stickers"`
}

// GetGuildPreviewResponse represents the response from GET /guilds/{guild.id}/preview
type GetGuildPreviewResponse = GuildPreview

// ModifyGuildRequest represents the request body for PATCH /guilds/{guild.id}
type ModifyGuildRequest struct {
	Name                        *string                                    `json:"name,omitempty"`
	VerificationLevel           *payloads.GuildVerificationLevel           `json:"verification_level,omitempty"`
	DefaultMessageNotifications *payloads.GuildDefaultMessageNotifications `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter       *payloads.GuildExplicitContentFilter       `json:"explicit_content_filter,omitempty"`
	AFKChannelID                *discord.Snowflake                         `json:"afk_channel_id,omitempty"`
	AFKTimeout                  *int                                       `json:"afk_timeout,omitempty"`
	Icon                        *string                                    `json:"icon,omitempty"`
	OwnerID                     *discord.Snowflake                         `json:"owner_id,omitempty"`
	Splash                      *string                                    `json:"splash,omitempty"`
	DiscoverySplash             *string                                    `json:"discovery_splash,omitempty"`
	Banner                      *string                                    `json:"banner,omitempty"`
	SystemChannelID             *discord.Snowflake                         `json:"system_channel_id,omitempty"`
	SystemChannelFlags          *payloads.GuildSystemChannelFlags          `json:"system_channel_flags,omitempty"`
	RulesChannelID              *discord.Snowflake                         `json:"rules_channel_id,omitempty"`
	PublicUpdatesChannelID      *discord.Snowflake                         `json:"public_updates_channel_id,omitempty"`
	PreferredLocale             *payloads.Locale                           `json:"preferred_locale,omitempty"`
	Features                    []payloads.GuildFeature                    `json:"features,omitempty"`
	Description                 *string                                    `json:"description,omitempty"`
	PremiumProgressBarEnabled   *bool                                      `json:"premium_progress_bar_enabled,omitempty"`
	SafetyAlertsChannelID       *discord.Snowflake                         `json:"safety_alerts_channel_id,omitempty"`
}

// ModifyGuildResponse represents the response from PATCH /guilds/{guild.id}
type ModifyGuildResponse = payloads.Guild

// GetGuildChannelsResponse represents the response from GET /guilds/{guild.id}/channels
type GetGuildChannelsResponse = []payloads.GuildTextChannel

// CreateGuildChannelRequest represents the request body for POST /guilds/{guild.id}/channels
type CreateGuildChannelRequest struct {
	Name                          string                              `json:"name"`
	Type                          *payloads.ChannelType               `json:"type,omitempty"`
	Topic                         *string                             `json:"topic,omitempty"`
	Bitrate                       *int                                `json:"bitrate,omitempty"`
	UserLimit                     *int                                `json:"user_limit,omitempty"`
	RateLimitPerUser              *int                                `json:"rate_limit_per_user,omitempty"`
	Position                      *int                                `json:"position,omitempty"`
	PermissionOverwrites          []payloads.Overwrite                `json:"permission_overwrites,omitempty"`
	ParentID                      *discord.Snowflake                  `json:"parent_id,omitempty"`
	NSFW                          *bool                               `json:"nsfw,omitempty"`
	RTCRegion                     *string                             `json:"rtc_region,omitempty"`
	VideoQualityMode              *payloads.VideoQualityMode          `json:"video_quality_mode,omitempty"`
	DefaultAutoArchiveDuration    *payloads.ThreadAutoArchiveDuration `json:"default_auto_archive_duration,omitempty"`
	DefaultReactionEmoji          *payloads.DefaultReaction           `json:"default_reaction_emoji,omitempty"`
	AvailableTags                 []payloads.ForumTag                 `json:"available_tags,omitempty"`
	DefaultSortOrder              *payloads.SortOrderType             `json:"default_sort_order,omitempty"`
	DefaultForumLayout            *payloads.ForumLayoutType           `json:"default_forum_layout,omitempty"`
	DefaultThreadRateLimitPerUser *int                                `json:"default_thread_rate_limit_per_user,omitempty"`
}

// CreateGuildChannelResponse represents the response from POST /guilds/{guild.id}/channels
type CreateGuildChannelResponse = payloads.GuildTextChannel

// GetGuildMemberResponse represents the response from GET /guilds/{guild.id}/members/{user.id}
type GetGuildMemberResponse = payloads.GuildMember

// GetGuildMembersResponse represents the response from GET /guilds/{guild.id}/members
type GetGuildMembersResponse = []payloads.GuildMember

// SearchGuildMembersResponse represents the response from GET /guilds/{guild.id}/members/search
type SearchGuildMembersResponse = []payloads.GuildMember

// AddGuildMemberRequest represents the request body for PUT /guilds/{guild.id}/members/{user.id}
type AddGuildMemberRequest struct {
	AccessToken *string             `json:"access_token"`
	Nick        *string             `json:"nick,omitempty"`
	Roles       []discord.Snowflake `json:"roles,omitempty"`
	Mute        *bool               `json:"mute,omitempty"`
	Deaf        *bool               `json:"deaf,omitempty"`
}

// AddGuildMemberResponse represents the response from PUT /guilds/{guild.id}/members/{user.id}
type AddGuildMemberResponse = payloads.GuildMember

// ModifyGuildMemberRequest represents the request body for PATCH /guilds/{guild.id}/members/{user.id}
type ModifyGuildMemberRequest struct {
	Nick                       *string                    `json:"nick,omitempty"`
	Roles                      []discord.Snowflake        `json:"roles,omitempty"`
	Mute                       *bool                      `json:"mute,omitempty"`
	Deaf                       *bool                      `json:"deaf,omitempty"`
	ChannelID                  *discord.Snowflake         `json:"channel_id,omitempty"`
	CommunicationDisabledUntil *time.Time                 `json:"communication_disabled_until,omitempty"`
	Flags                      *payloads.GuildMemberFlags `json:"flags,omitempty"`
}

// ModifyGuildMemberResponse represents the response from PATCH /guilds/{guild.id}/members/{user.id}
type ModifyGuildMemberResponse = payloads.GuildMember

// Ban represents a Discord guild ban
type Ban struct {
	Reason *string       `json:"reason"`
	User   payloads.User `json:"user"`
}

// GetGuildBansResponse represents the response from GET /guilds/{guild.id}/bans
type GetGuildBansResponse = []Ban

// GetGuildBanResponse represents the response from GET /guilds/{guild.id}/bans/{user.id}
type GetGuildBanResponse = Ban

// CreateGuildBanRequest represents the request body for PUT /guilds/{guild.id}/bans/{user.id}
type CreateGuildBanRequest struct {
	DeleteMessageSeconds *int    `json:"delete_message_seconds,omitempty"`
	Reason               *string `json:"reason,omitempty"`
}

// BulkGuildBanRequest represents the request body for POST /guilds/{guild.id}/bulk-ban
type BulkGuildBanRequest struct {
	UserIDs              []discord.Snowflake `json:"user_ids"`
	DeleteMessageSeconds *int                `json:"delete_message_seconds,omitempty"`
}

// BulkGuildBanResponse represents the response from POST /guilds/{guild.id}/bulk-ban
type BulkGuildBanResponse struct {
	BannedUsers []discord.Snowflake `json:"banned_users"`
	FailedUsers []discord.Snowflake `json:"failed_users"`
}

// GetGuildRolesResponse represents the response from GET /guilds/{guild.id}/roles
type GetGuildRolesResponse = []payloads.Role

// CreateGuildRoleRequest represents the request body for POST /guilds/{guild.id}/roles
type CreateGuildRoleRequest struct {
	Name         *string              `json:"name,omitempty"`
	Permissions  *discord.Permissions `json:"permissions,omitempty"`
	Color        *int                 `json:"color,omitempty"`
	Hoist        *bool                `json:"hoist,omitempty"`
	Icon         *string              `json:"icon,omitempty"`
	UnicodeEmoji *string              `json:"unicode_emoji,omitempty"`
	Mentionable  *bool                `json:"mentionable,omitempty"`
}

// CreateGuildRoleResponse represents the response from POST /guilds/{guild.id}/roles
type CreateGuildRoleResponse = payloads.Role

// ModifyGuildRoleRequest represents the request body for PATCH /guilds/{guild.id}/roles/{role.id}
type ModifyGuildRoleRequest struct {
	Name         *string              `json:"name,omitempty"`
	Permissions  *discord.Permissions `json:"permissions,omitempty"`
	Color        *int                 `json:"color,omitempty"`
	Hoist        *bool                `json:"hoist,omitempty"`
	Icon         *string              `json:"icon,omitempty"`
	UnicodeEmoji *string              `json:"unicode_emoji,omitempty"`
	Mentionable  *bool                `json:"mentionable,omitempty"`
}

// ModifyGuildRoleResponse represents the response from PATCH /guilds/{guild.id}/roles/{role.id}
type ModifyGuildRoleResponse = payloads.Role

// ModifyGuildRolePositionsRequest represents the request body for PATCH /guilds/{guild.id}/roles
type ModifyGuildRolePositionsRequest = []ModifyGuildRolePositionRequest

// ModifyGuildRolePositionRequest represents a single role position update
type ModifyGuildRolePositionRequest struct {
	ID       discord.Snowflake `json:"id"`
	Position *int              `json:"position,omitempty"`
}

// ModifyGuildRolePositionsResponse represents the response from PATCH /guilds/{guild.id}/roles
type ModifyGuildRolePositionsResponse = []payloads.Role

// ====================
// User Types
// ====================

// GetCurrentUserResponse represents the response from GET /users/@me
type GetCurrentUserResponse = payloads.User

// GetUserResponse represents the response from GET /users/{user.id}
type GetUserResponse = payloads.User

// ModifyCurrentUserRequest represents the request body for PATCH /users/@me
type ModifyCurrentUserRequest struct {
	Username    *string `json:"username,omitempty"`
	Avatar      *string `json:"avatar,omitempty"`
	Banner      *string `json:"banner,omitempty"`
	AccentColor *int    `json:"accent_color,omitempty"`
}

// ModifyCurrentUserResponse represents the response from PATCH /users/@me
type ModifyCurrentUserResponse = payloads.User

// GetCurrentUserGuildsResponse represents the response from GET /users/@me/guilds
type GetCurrentUserGuildsResponse = []payloads.PartialGuild

// GetCurrentUserGuildMemberResponse represents the response from GET /users/@me/guilds/{guild.id}/member
type GetCurrentUserGuildMemberResponse = payloads.GuildMember

// ====================
// Interaction Types
// ====================

// CreateInteractionResponseRequest represents the request body for POST /interactions/{interaction.id}/{interaction.token}/callback
type CreateInteractionResponseRequest struct {
	Type payloads.InteractionResponseType `json:"type"`
	Data interface{}                      `json:"data,omitempty"`
}

// GetOriginalInteractionResponseResponse represents the response from GET /webhooks/{application.id}/{interaction.token}/messages/@original
type GetOriginalInteractionResponseResponse = payloads.Message

// EditOriginalInteractionResponseRequest represents the request body for PATCH /webhooks/{application.id}/{interaction.token}/messages/@original
type EditOriginalInteractionResponseRequest = EditMessageRequest

// EditOriginalInteractionResponseResponse represents the response from PATCH /webhooks/{application.id}/{interaction.token}/messages/@original
type EditOriginalInteractionResponseResponse = payloads.Message

// CreateFollowupMessageRequest represents the request body for POST /webhooks/{application.id}/{interaction.token}
type CreateFollowupMessageRequest = CreateMessageRequest

// CreateFollowupMessageResponse represents the response from POST /webhooks/{application.id}/{interaction.token}
type CreateFollowupMessageResponse = payloads.Message

// ====================
// Application Command Types
// ====================

// ApplicationCommandOption represents an application command option
type ApplicationCommandOption struct {
	Type                     payloads.ApplicationCommandOptionType     `json:"type"`
	Name                     string                                    `json:"name"`
	NameLocalizations        payloads.LocalizationMap                  `json:"name_localizations,omitempty"`
	Description              string                                    `json:"description"`
	DescriptionLocalizations payloads.LocalizationMap                  `json:"description_localizations,omitempty"`
	Required                 *bool                                     `json:"required,omitempty"`
	Choices                  []payloads.ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options                  []ApplicationCommandOption                `json:"options,omitempty"`
	ChannelTypes             []payloads.ChannelType                    `json:"channel_types,omitempty"`
	MinValue                 *float64                                  `json:"min_value,omitempty"`
	MaxValue                 *float64                                  `json:"max_value,omitempty"`
	MinLength                *int                                      `json:"min_length,omitempty"`
	MaxLength                *int                                      `json:"max_length,omitempty"`
	Autocomplete             *bool                                     `json:"autocomplete,omitempty"`
}

// ApplicationCommand represents a Discord application command
type ApplicationCommand struct {
	ID                       discord.Snowflake                     `json:"id"`
	Type                     payloads.ApplicationCommandType       `json:"type"`
	ApplicationID            discord.Snowflake                     `json:"application_id"`
	GuildID                  *discord.Snowflake                    `json:"guild_id,omitempty"`
	Name                     string                                `json:"name"`
	NameLocalizations        payloads.LocalizationMap              `json:"name_localizations,omitempty"`
	Description              string                                `json:"description"`
	DescriptionLocalizations payloads.LocalizationMap              `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption            `json:"options,omitempty"`
	DefaultMemberPermissions *discord.Permissions                  `json:"default_member_permissions,omitempty"`
	DMPermission             *bool                                 `json:"dm_permission,omitempty"`
	DefaultPermission        *bool                                 `json:"default_permission,omitempty"`
	NSFW                     *bool                                 `json:"nsfw,omitempty"`
	Contexts                 []payloads.InteractionContextType     `json:"contexts,omitempty"`
	IntegrationTypes         []payloads.ApplicationIntegrationType `json:"integration_types,omitempty"`
	Version                  discord.Snowflake                     `json:"version"`
}

// GetGlobalApplicationCommandsResponse represents the response from GET /applications/{application.id}/commands
type GetGlobalApplicationCommandsResponse = []ApplicationCommand

// CreateGlobalApplicationCommandRequest represents the request body for POST /applications/{application.id}/commands
type CreateGlobalApplicationCommandRequest struct {
	Name                     string                                `json:"name"`
	NameLocalizations        payloads.LocalizationMap              `json:"name_localizations,omitempty"`
	Description              *string                               `json:"description,omitempty"`
	DescriptionLocalizations payloads.LocalizationMap              `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption            `json:"options,omitempty"`
	DefaultMemberPermissions *discord.Permissions                  `json:"default_member_permissions,omitempty"`
	DMPermission             *bool                                 `json:"dm_permission,omitempty"`
	DefaultPermission        *bool                                 `json:"default_permission,omitempty"`
	Type                     *payloads.ApplicationCommandType      `json:"type,omitempty"`
	NSFW                     *bool                                 `json:"nsfw,omitempty"`
	Contexts                 []payloads.InteractionContextType     `json:"contexts,omitempty"`
	IntegrationTypes         []payloads.ApplicationIntegrationType `json:"integration_types,omitempty"`
}

// CreateGlobalApplicationCommandResponse represents the response from POST /applications/{application.id}/commands
type CreateGlobalApplicationCommandResponse = ApplicationCommand

// GetGlobalApplicationCommandResponse represents the response from GET /applications/{application.id}/commands/{command.id}
type GetGlobalApplicationCommandResponse = ApplicationCommand

// EditGlobalApplicationCommandRequest represents the request body for PATCH /applications/{application.id}/commands/{command.id}
type EditGlobalApplicationCommandRequest = CreateGlobalApplicationCommandRequest

// EditGlobalApplicationCommandResponse represents the response from PATCH /applications/{application.id}/commands/{command.id}
type EditGlobalApplicationCommandResponse = ApplicationCommand

// BulkOverwriteGlobalApplicationCommandsRequest represents the request body for PUT /applications/{application.id}/commands
type BulkOverwriteGlobalApplicationCommandsRequest = []CreateGlobalApplicationCommandRequest

// BulkOverwriteGlobalApplicationCommandsResponse represents the response from PUT /applications/{application.id}/commands
type BulkOverwriteGlobalApplicationCommandsResponse = []ApplicationCommand

// ====================
// Webhook Types
// ====================

// CreateWebhookRequest represents the request body for POST /channels/{channel.id}/webhooks
type CreateWebhookRequest struct {
	Name   string  `json:"name"`
	Avatar *string `json:"avatar,omitempty"`
}

// CreateWebhookResponse represents the response from POST /channels/{channel.id}/webhooks
type CreateWebhookResponse = payloads.Webhook

// GetChannelWebhooksResponse represents the response from GET /channels/{channel.id}/webhooks
type GetChannelWebhooksResponse = []payloads.Webhook

// GetGuildWebhooksResponse represents the response from GET /guilds/{guild.id}/webhooks
type GetGuildWebhooksResponse = []payloads.Webhook

// GetWebhookResponse represents the response from GET /webhooks/{webhook.id}
type GetWebhookResponse = payloads.Webhook

// ModifyWebhookRequest represents the request body for PATCH /webhooks/{webhook.id}
type ModifyWebhookRequest struct {
	Name      *string            `json:"name,omitempty"`
	Avatar    *string            `json:"avatar,omitempty"`
	ChannelID *discord.Snowflake `json:"channel_id,omitempty"`
}

// ModifyWebhookResponse represents the response from PATCH /webhooks/{webhook.id}
type ModifyWebhookResponse = payloads.Webhook

// ExecuteWebhookRequest represents the request body for POST /webhooks/{webhook.id}/{webhook.token}
type ExecuteWebhookRequest struct {
	Content         *string                       `json:"content,omitempty"`
	Username        *string                       `json:"username,omitempty"`
	AvatarURL       *string                       `json:"avatar_url,omitempty"`
	TTS             *bool                         `json:"tts,omitempty"`
	Embeds          []payloads.Embed              `json:"embeds,omitempty"`
	AllowedMentions *payloads.AllowedMentions     `json:"allowed_mentions,omitempty"`
	Components      []payloads.ActionRowComponent `json:"components,omitempty"`
	Attachments     []payloads.PartialAttachment  `json:"attachments,omitempty"`
	Flags           *payloads.MessageFlags        `json:"flags,omitempty"`
	ThreadName      *string                       `json:"thread_name,omitempty"`
	AppliedTags     []discord.Snowflake           `json:"applied_tags,omitempty"`
	Poll            *payloads.Poll                `json:"poll,omitempty"`
}

// ExecuteWebhookResponse represents the response from POST /webhooks/{webhook.id}/{webhook.token}
type ExecuteWebhookResponse = payloads.Message

// ====================
// OAuth2 Types
// ====================

// GetCurrentBotApplicationInformationResponse represents the response from GET /oauth2/applications/@me
type GetCurrentBotApplicationInformationResponse = payloads.Application

// PartialApplication represents a partial Discord application
type PartialApplication struct {
	ID          discord.Snowflake `json:"id"`
	Name        string            `json:"name"`
	Icon        *string           `json:"icon"`
	Description string            `json:"description"`
	Summary     string            `json:"summary"`
	Verify      bool              `json:"verify_key"`
}

// GetCurrentAuthorizationInformationResponse represents the response from GET /oauth2/@me
type GetCurrentAuthorizationInformationResponse struct {
	Application PartialApplication `json:"application"`
	Scopes      []string           `json:"scopes"`
	Expires     time.Time          `json:"expires"`
	User        *payloads.User     `json:"user,omitempty"`
}

// ====================
// Error Types
// ====================

// RESTError represents a REST API error response
type RESTError struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors,omitempty"`
}

// Error implements the error interface
func (e RESTError) Error() string {
	return e.Message
}

// ====================
// Common Response Types
// ====================

// EmptyResponse represents endpoints that return no content (204 No Content)
type EmptyResponse struct{}

// GenericErrorResponse represents a generic error response structure
type GenericErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
