// Package rest provides Discord REST API types and utilities.
//
// This file contains comprehensive specialized REST API request and response types.
// Includes: Emoji, Sticker, Soundboard, Poll, User, Voice, Invite, Application, OAuth2
package rest

import (
	"github.com/kolosys/discord-types/discord"
	"github.com/kolosys/discord-types/payloads"
)

// ====================
// Emoji Types
// ====================

// GetGuildEmojisResult represents the response from GET /guilds/{guild.id}/emojis
type GetGuildEmojisResult = []payloads.Emoji

// GetGuildEmojiResult represents the response from GET /guilds/{guild.id}/emojis/{emoji.id}
type GetGuildEmojiResult = payloads.Emoji

// PostGuildEmojiJSONBody represents the request body for POST /guilds/{guild.id}/emojis
type PostGuildEmojiJSONBody struct {
	// Name of the emoji
	Name string `json:"name"`
	// The 128x128 emoji image (base64 data URI)
	Image string `json:"image"`
	// Roles for which this emoji will be whitelisted
	Roles []discord.Snowflake `json:"roles,omitempty"`
}

// PostGuildEmojiResult represents the response from POST /guilds/{guild.id}/emojis
type PostGuildEmojiResult = payloads.Emoji

// PatchGuildEmojiJSONBody represents the request body for PATCH /guilds/{guild.id}/emojis/{emoji.id}
type PatchGuildEmojiJSONBody struct {
	// Name of the emoji
	Name *string `json:"name,omitempty"`
	// Roles for which this emoji will be whitelisted
	Roles []discord.Snowflake `json:"roles,omitempty"`
}

// PatchGuildEmojiResult represents the response from PATCH /guilds/{guild.id}/emojis/{emoji.id}
type PatchGuildEmojiResult = payloads.Emoji

// DeleteGuildEmojiResult represents the response from DELETE /guilds/{guild.id}/emojis/{emoji.id}
type DeleteGuildEmojiResult struct{} // Empty response

// GetApplicationEmojisResult represents the response from GET /applications/{application.id}/emojis
type GetApplicationEmojisResult struct {
	Items []payloads.Emoji `json:"items"`
}

// GetApplicationEmojiResult represents the response from GET /applications/{application.id}/emojis/{emoji.id}
type GetApplicationEmojiResult = payloads.Emoji

// PostApplicationEmojiJSONBody represents the request body for POST /applications/{application.id}/emojis
type PostApplicationEmojiJSONBody struct {
	// Name of the emoji
	Name string `json:"name"`
	// The 128x128 emoji image (base64 data URI)
	Image string `json:"image"`
}

// PostApplicationEmojiResult represents the response from POST /applications/{application.id}/emojis
type PostApplicationEmojiResult = payloads.Emoji

// PatchApplicationEmojiJSONBody represents the request body for PATCH /applications/{application.id}/emojis/{emoji.id}
type PatchApplicationEmojiJSONBody struct {
	// Name of the emoji
	Name *string `json:"name,omitempty"`
}

// PatchApplicationEmojiResult represents the response from PATCH /applications/{application.id}/emojis/{emoji.id}
type PatchApplicationEmojiResult = payloads.Emoji

// DeleteApplicationEmojiResult represents the response from DELETE /applications/{application.id}/emojis/{emoji.id}
type DeleteApplicationEmojiResult struct{} // Empty response

// ====================
// Sticker Types
// ====================

// GetStickerResult represents the response from GET /stickers/{sticker.id}
type GetStickerResult = payloads.Sticker

// GetStickerPacksResult represents the response from GET /sticker-packs
type GetStickerPacksResult struct {
	StickerPacks []payloads.StickerPack `json:"sticker_packs"`
}

// GetStickerPackResult represents the response from GET /sticker-packs/{pack.id}
type GetStickerPackResult = payloads.StickerPack

// GetGuildStickersResult represents the response from GET /guilds/{guild.id}/stickers
type GetGuildStickersResult = []payloads.Sticker

// GetGuildStickerResult represents the response from GET /guilds/{guild.id}/stickers/{sticker.id}
type GetGuildStickerResult = payloads.Sticker

// PostGuildStickerJSONBody represents the request body for POST /guilds/{guild.id}/stickers
type PostGuildStickerJSONBody struct {
	// Name of the sticker (2-30 characters)
	Name string `json:"name"`
	// Description of the sticker (empty or 2-100 characters)
	Description string `json:"description"`
	// Autocomplete/suggestion tags for the sticker (max 200 characters)
	Tags string `json:"tags"`
	// The sticker file to upload (PNG, APNG, GIF, JSON)
	File []byte `json:"file,omitempty"`
}

// PostGuildStickerResult represents the response from POST /guilds/{guild.id}/stickers
type PostGuildStickerResult = payloads.Sticker

// PatchGuildStickerJSONBody represents the request body for PATCH /guilds/{guild.id}/stickers/{sticker.id}
type PatchGuildStickerJSONBody struct {
	// Name of the sticker (2-30 characters)
	Name *string `json:"name,omitempty"`
	// Description of the sticker (2-100 characters)
	Description *string `json:"description,omitempty"`
	// Autocomplete/suggestion tags for the sticker (max 200 characters)
	Tags *string `json:"tags,omitempty"`
}

// PatchGuildStickerResult represents the response from PATCH /guilds/{guild.id}/stickers/{sticker.id}
type PatchGuildStickerResult = payloads.Sticker

// DeleteGuildStickerResult represents the response from DELETE /guilds/{guild.id}/stickers/{sticker.id}
type DeleteGuildStickerResult struct{} // Empty response

// ====================
// Soundboard Types
// ====================

// GetGuildSoundboardSoundsResult represents the response from GET /guilds/{guild.id}/soundboard-sounds
type GetGuildSoundboardSoundsResult = []payloads.SoundboardSound

// GetGuildSoundboardSoundResult represents the response from GET /guilds/{guild.id}/soundboard-sounds/{sound.id}
type GetGuildSoundboardSoundResult = payloads.SoundboardSound

// PostGuildSoundboardSoundJSONBody represents the request body for POST /guilds/{guild.id}/soundboard-sounds
type PostGuildSoundboardSoundJSONBody struct {
	// Name of the soundboard sound (2-32 characters)
	Name string `json:"name"`
	// The sound file to upload (MP3, OGG)
	Sound []byte `json:"sound"`
	// The volume of the soundboard sound (0.0 to 1.0)
	Volume *float64 `json:"volume,omitempty"`
	// Emoji for the soundboard sound
	EmojiID *discord.Snowflake `json:"emoji_id,omitempty"`
	// Unicode emoji for the soundboard sound
	EmojiName *string `json:"emoji_name,omitempty"`
}

// PostGuildSoundboardSoundResult represents the response from POST /guilds/{guild.id}/soundboard-sounds
type PostGuildSoundboardSoundResult = payloads.SoundboardSound

// PatchGuildSoundboardSoundJSONBody represents the request body for PATCH /guilds/{guild.id}/soundboard-sounds/{sound.id}
type PatchGuildSoundboardSoundJSONBody struct {
	// Name of the soundboard sound (2-32 characters)
	Name *string `json:"name,omitempty"`
	// The volume of the soundboard sound (0.0 to 1.0)
	Volume *float64 `json:"volume,omitempty"`
	// Emoji for the soundboard sound
	EmojiID *discord.Snowflake `json:"emoji_id,omitempty"`
	// Unicode emoji for the soundboard sound
	EmojiName *string `json:"emoji_name,omitempty"`
}

// PatchGuildSoundboardSoundResult represents the response from PATCH /guilds/{guild.id}/soundboard-sounds/{sound.id}
type PatchGuildSoundboardSoundResult = payloads.SoundboardSound

// DeleteGuildSoundboardSoundResult represents the response from DELETE /guilds/{guild.id}/soundboard-sounds/{sound.id}
type DeleteGuildSoundboardSoundResult struct{} // Empty response

// PostSendSoundboardSoundJSONBody represents the request body for POST /channels/{channel.id}/send-soundboard-sound
type PostSendSoundboardSoundJSONBody struct {
	// The id of the soundboard sound to play
	SoundID discord.Snowflake `json:"sound_id"`
	// The id of the guild the soundboard sound is from
	SourceGuildID *discord.Snowflake `json:"source_guild_id,omitempty"`
}

// PostSendSoundboardSoundResult represents the response from POST /channels/{channel.id}/send-soundboard-sound
type PostSendSoundboardSoundResult struct{} // Empty response

// ====================
// Poll Types
// ====================

// GetPollAnswerVotersQuery represents query parameters for GET /channels/{channel.id}/polls/{message.id}/answers/{answer.id}
type GetPollAnswerVotersQuery struct {
	// Get users after this user ID
	After *discord.Snowflake `json:"after,omitempty"`
	// Max number of users to return (1-100, default 25)
	Limit *int `json:"limit,omitempty"`
}

// GetPollAnswerVotersResult represents the response from GET /channels/{channel.id}/polls/{message.id}/answers/{answer.id}
type GetPollAnswerVotersResult struct {
	Users []payloads.User `json:"users"`
}

// PostEndPollResult represents the response from POST /channels/{channel.id}/polls/{message.id}/expire
type PostEndPollResult = payloads.Message

// ====================
// User Types
// ====================

// GetCurrentUserResult represents the response from GET /users/@me
type GetCurrentUserResult = payloads.User

// GetUserResult represents the response from GET /users/{user.id}
type GetUserResult = payloads.User

// PatchCurrentUserJSONBody represents the request body for PATCH /users/@me
type PatchCurrentUserJSONBody struct {
	// User's username, if changed may cause the user's discriminator to be randomized
	Username *string `json:"username,omitempty"`
	// If passed, modifies the user's avatar
	Avatar *string `json:"avatar,omitempty"`
	// If passed, modifies the user's banner
	Banner *string `json:"banner,omitempty"`
}

// PatchCurrentUserResult represents the response from PATCH /users/@me
type PatchCurrentUserResult = payloads.User

// GetCurrentUserGuildsQuery represents query parameters for GET /users/@me/guilds
type GetCurrentUserGuildsQuery struct {
	// Get guilds before this guild ID
	Before *discord.Snowflake `json:"before,omitempty"`
	// Get guilds after this guild ID
	After *discord.Snowflake `json:"after,omitempty"`
	// Max number of guilds to return (1-200, default 200)
	Limit *int `json:"limit,omitempty"`
	// Whether to include approximate member and presence counts in response
	WithCounts *bool `json:"with_counts,omitempty"`
}

// GetCurrentUserGuildsResult represents the response from GET /users/@me/guilds
type GetCurrentUserGuildsResult = []payloads.Guild

// GetCurrentUserGuildMemberResult represents the response from GET /users/@me/guilds/{guild.id}/member
type GetCurrentUserGuildMemberResult = payloads.GuildMember

// DeleteCurrentUserGuildResult represents the response from DELETE /users/@me/guilds/{guild.id}
type DeleteCurrentUserGuildResult struct{} // Empty response

// PostCreateDMJSONBody represents the request body for POST /users/@me/channels
type PostCreateDMJSONBody struct {
	// The recipient to open a DM channel with
	RecipientID discord.Snowflake `json:"recipient_id"`
}

// PostCreateDMResult represents the response from POST /users/@me/channels
type PostCreateDMResult = payloads.DMChannel

// GetCurrentUserConnectionsResult represents the response from GET /users/@me/connections
type GetCurrentUserConnectionsResult = []payloads.Connection

// GetCurrentUserApplicationRoleConnectionResult represents the response from GET /users/@me/applications/{application.id}/role-connection
type GetCurrentUserApplicationRoleConnectionResult = payloads.ApplicationRoleConnection

// PutCurrentUserApplicationRoleConnectionJSONBody represents the request body for PUT /users/@me/applications/{application.id}/role-connection
type PutCurrentUserApplicationRoleConnectionJSONBody struct {
	// The vanity name of the platform a bot has connected (max 50 characters)
	PlatformName *string `json:"platform_name,omitempty"`
	// The username on the platform a bot has connected (max 100 characters)
	PlatformUsername *string `json:"platform_username,omitempty"`
	// Object mapping application role connection metadata keys to their stringified value (max 100 characters) for the user on the platform a bot has connected
	Metadata map[string]string `json:"metadata,omitempty"`
}

// PutCurrentUserApplicationRoleConnectionResult represents the response from PUT /users/@me/applications/{application.id}/role-connection
type PutCurrentUserApplicationRoleConnectionResult = payloads.ApplicationRoleConnection

// ====================
// Voice Types
// ====================

// GetVoiceRegionsResult represents the response from GET /voice/regions
type GetVoiceRegionsResult = []payloads.VoiceRegion

// ====================
// Invite Types
// ====================

// GetInviteQuery represents query parameters for GET /invites/{invite.code}
type GetInviteQuery struct {
	// Whether the invite should contain approximate member counts
	WithCounts *bool `json:"with_counts,omitempty"`
	// Whether the invite should contain the expiration date
	WithExpiration *bool `json:"with_expiration,omitempty"`
	// The guild scheduled event to include with the invite
	GuildScheduledEventID *discord.Snowflake `json:"guild_scheduled_event_id,omitempty"`
}

// GetInviteResult represents the response from GET /invites/{invite.code}
type GetInviteResult = Invite

// DeleteInviteResult represents the response from DELETE /invites/{invite.code}
type DeleteInviteResult = Invite

// ====================
// Application Types
// ====================

// GetCurrentApplicationResult represents the response from GET /applications/@me
type GetCurrentApplicationResult = PartialApplication

// PatchCurrentApplicationJSONBody represents the request body for PATCH /applications/@me
type PatchCurrentApplicationJSONBody struct {
	// Custom authorization URL for the application, if enabled
	CustomInstallURL *string `json:"custom_install_url,omitempty"`
	// Description of the application
	Description *string `json:"description,omitempty"`
	// Role connection verification URL for the application
	RoleConnectionsVerificationURL *string `json:"role_connections_verification_url,omitempty"`
	// Settings for the application's default in-app authorization link, if enabled
	InstallParams *payloads.ApplicationInstallParams `json:"install_params,omitempty"`
	// List of supported installation contexts
	IntegrationTypesConfig map[payloads.ApplicationIntegrationType]payloads.ApplicationIntegrationTypeConfiguration `json:"integration_types_config,omitempty"`
	// Application's public flags
	Flags *payloads.ApplicationFlags `json:"flags,omitempty"`
	// Icon for the application
	Icon *string `json:"icon,omitempty"`
	// Default rich presence invite cover image for the application
	CoverImage *string `json:"cover_image,omitempty"`
	// Interactions endpoint URL for the application
	InteractionsEndpointURL *string `json:"interactions_endpoint_url,omitempty"`
	// List of tags describing the content and functionality of the application
	Tags []string `json:"tags,omitempty"`
}

// PatchCurrentApplicationResult represents the response from PATCH /applications/@me
type PatchCurrentApplicationResult = PartialApplication
