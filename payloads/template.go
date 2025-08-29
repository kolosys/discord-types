package payloads

import "github.com/kolosys/discord-types/discord"

// Template represents a Discord guild template.
//
// See: https://discord.com/developers/docs/resources/guild-template#guild-template-object
type Template struct {
	// Code is the template code (unique ID).
	Code string `json:"code"`

	// Name is the template name.
	Name string `json:"name"`

	// Description is the description for the template.
	Description *string `json:"description"`

	// UsageCount is the number of times this template has been used.
	UsageCount int `json:"usage_count"`

	// CreatorID is the ID of the user who created the template.
	CreatorID discord.Snowflake `json:"creator_id"`

	// Creator is the user who created the template.
	Creator User `json:"creator"`

	// CreatedAt is when this template was created.
	CreatedAt string `json:"created_at"`

	// UpdatedAt is when this template was last synced to the source guild.
	UpdatedAt string `json:"updated_at"`

	// SourceGuildID is the ID of the guild this template is based on.
	SourceGuildID discord.Snowflake `json:"source_guild_id"`

	// SerializedSourceGuild is the guild snapshot this template contains.
	SerializedSourceGuild TemplateSerializedGuild `json:"serialized_source_guild"`

	// IsDirty indicates whether the template has unsynced changes.
	IsDirty *bool `json:"is_dirty,omitempty"`
}

// TemplateSerializedGuild represents a serialized guild in a template.
type TemplateSerializedGuild struct {
	// ID is the guild id.
	ID discord.Snowflake `json:"id"`

	// Name is the guild name (2-100 characters).
	Name string `json:"name"`

	// Description is the description for the guild.
	Description *string `json:"description"`

	// Region is the voice region id for the guild (deprecated).
	Region string `json:"region"`

	// VerificationLevel is the verification level required for the guild.
	VerificationLevel GuildVerificationLevel `json:"verification_level"`

	// DefaultMessageNotifications is the default message notifications level.
	DefaultMessageNotifications GuildDefaultMessageNotifications `json:"default_message_notifications"`

	// ExplicitContentFilter is the explicit content filter level.
	ExplicitContentFilter GuildExplicitContentFilter `json:"explicit_content_filter"`

	// PreferredLocale is the preferred locale of a Community guild.
	PreferredLocale Locale `json:"preferred_locale"`

	// AFKTimeout is the afk timeout in seconds.
	AFKTimeout AFKTimeout `json:"afk_timeout"`

	// Roles are the roles in the guild.
	Roles []TemplateRole `json:"roles"`

	// Channels are the channels in the guild.
	Channels []TemplateChannel `json:"channels"`

	// AFKChannelID is the ID of afk channel.
	AFKChannelID *int `json:"afk_channel_id"`

	// SystemChannelID is the id of the channel where guild notices such as welcome messages and boost events are posted.
	SystemChannelID *int `json:"system_channel_id"`

	// SystemChannelFlags are the system channel flags.
	SystemChannelFlags GuildSystemChannelFlags `json:"system_channel_flags"`
}

// TemplateRole represents a role in a template.
type TemplateRole struct {
	// ID is the role id.
	ID int `json:"id"`

	// Name is the role name.
	Name string `json:"name"`

	// Color is the integer representation of hexadecimal color code.
	Color int `json:"color"`

	// Hoist indicates if this role is pinned in the user listing.
	Hoist bool `json:"hoist"`

	// Mentionable indicates whether this role is mentionable.
	Mentionable bool `json:"mentionable"`

	// Permissions are the permission bit set.
	Permissions discord.Permissions `json:"permissions"`
}

// TemplateChannel represents a channel in a template.
type TemplateChannel struct {
	// ID is the channel id.
	ID int `json:"id"`

	// Type is the type of channel.
	Type ChannelType `json:"type"`

	// Name is the channel name.
	Name string `json:"name"`

	// Position is the sorting position of the channel.
	Position int `json:"position"`

	// Topic is the channel topic.
	Topic *string `json:"topic,omitempty"`

	// Bitrate is the bitrate (in bits) of the voice or stage channel.
	Bitrate *int `json:"bitrate,omitempty"`

	// UserLimit is the user limit of the voice or stage channel.
	UserLimit *int `json:"user_limit,omitempty"`

	// RateLimitPerUser is the amount of seconds a user has to wait before sending another message.
	RateLimitPerUser *int `json:"rate_limit_per_user,omitempty"`

	// PermissionOverwrites are the explicit permission overwrites for members and roles.
	PermissionOverwrites []TemplateOverwrite `json:"permission_overwrites,omitempty"`

	// ParentID is the id of the parent category for a channel.
	ParentID *int `json:"parent_id,omitempty"`

	// NSFW indicates whether the channel is nsfw.
	NSFW *bool `json:"nsfw,omitempty"`
}

// TemplateOverwrite represents a permission overwrite in a template.
type TemplateOverwrite struct {
	// ID is the role or user id.
	ID int `json:"id"`

	// Type is the either 0 (role) or 1 (member).
	Type OverwriteType `json:"type"`

	// Allow are the permission bit set.
	Allow discord.Permissions `json:"allow"`

	// Deny are the permission bit set.
	Deny discord.Permissions `json:"deny"`
}
