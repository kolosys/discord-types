package v10

import (
	"time"
)

// Channel represents a unified channel object that can handle all channel types
//
// # This struct includes all possible fields from the Discord Channel object
//
// https://discord.com/developers/docs/resources/channel#channel-object
type Channel struct {
	// Base fields (present in all channel types)
	ID    Snowflake     `json:"id"`
	Type  ChannelType   `json:"type"`
	Flags *ChannelFlags `json:"flags,omitempty"`

	// Guild channel fields
	GuildID              *Snowflake  `json:"guild_id,omitempty"`
	Position             *int        `json:"position,omitempty"`
	PermissionOverwrites []Overwrite `json:"permission_overwrites,omitempty"`
	Name                 *string     `json:"name,omitempty"`
	Topic                *string     `json:"topic,omitempty"`
	NSFW                 *bool       `json:"nsfw,omitempty"`
	ParentID             *Snowflake  `json:"parent_id,omitempty"`

	// Text-based channel fields
	LastMessageID    *Snowflake `json:"last_message_id,omitempty"`
	RateLimitPerUser *int       `json:"rate_limit_per_user,omitempty"`
	LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"`

	// Voice channel fields
	Bitrate          *int              `json:"bitrate,omitempty"`
	UserLimit        *int              `json:"user_limit,omitempty"`
	RTCRegion        *string           `json:"rtc_region,omitempty"`
	VideoQualityMode *VideoQualityMode `json:"video_quality_mode,omitempty"`

	// DM channel fields
	Recipients    []User     `json:"recipients,omitempty"`
	Icon          *string    `json:"icon,omitempty"`
	OwnerID       *Snowflake `json:"owner_id,omitempty"`
	ApplicationID *Snowflake `json:"application_id,omitempty"`
	Managed       *bool      `json:"managed,omitempty"`

	// Thread channel fields
	MessageCount     *int            `json:"message_count,omitempty"`
	MemberCount      *int            `json:"member_count,omitempty"`
	ThreadMetadata   *ThreadMetadata `json:"thread_metadata,omitempty"`
	Member           *ThreadMember   `json:"member,omitempty"`
	TotalMessageSent *int            `json:"total_message_sent,omitempty"`

	// Forum/Media channel fields
	DefaultAutoArchiveDuration    *ThreadAutoArchiveDuration `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser *int                       `json:"default_thread_rate_limit_per_user,omitempty"`
	AvailableTags                 []ForumTag                 `json:"available_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction           `json:"default_reaction_emoji,omitempty"`
	DefaultForumLayout            *ForumLayoutType           `json:"default_forum_layout,omitempty"`
	DefaultSortOrder              *SortOrderType             `json:"default_sort_order,omitempty"`
	AppliedTags                   []Snowflake                `json:"applied_tags,omitempty"`

	// Computed fields
	Permissions *Permissions `json:"permissions,omitempty"`
}

// ChannelType represents the type of channel
//
// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
type ChannelType int

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildAnnouncement

	// Thread and forum channels
	ChannelTypeAnnouncementThread ChannelType = 10
	ChannelTypePublicThread
	ChannelTypePrivateThread
	ChannelTypeGuildStageVoice
	ChannelTypeGuildDirectory
	ChannelTypeGuildForum
	ChannelTypeGuildMedia

	// Deprecated names
	ChannelTypeGuildNews          = ChannelTypeGuildAnnouncement  // Deprecated v11
	ChannelTypeGuildNewsThread    = ChannelTypeAnnouncementThread // Deprecated
	ChannelTypeGuildPublicThread  = ChannelTypePublicThread       // Deprecated
	ChannelTypeGuildPrivateThread = ChannelTypePrivateThread      // Deprecated
)

// VideoQualityMode represents the video quality mode for voice channels
type VideoQualityMode int

const (
	VideoQualityModeAuto VideoQualityMode = iota + 1
	VideoQualityModeFull
)

// ChannelFlags represents channel-specific flags
type ChannelFlags int

const (
	ChannelFlagPinned                   ChannelFlags = 1 << 1  // 2 - this channel is pinned in the guild channel list
	ChannelFlagRequireTag               ChannelFlags = 1 << 4  // 16 - whether a tag is required to be specified when creating a thread
	ChannelFlagHideMediaDownloadOptions ChannelFlags = 1 << 15 // 32768 - when set, hides the embedded media download options
)

// ThreadAutoArchiveDuration represents auto-archive duration for threads in minutes
type ThreadAutoArchiveDuration int

const (
	ThreadAutoArchiveDuration60    ThreadAutoArchiveDuration = 60
	ThreadAutoArchiveDuration1440  ThreadAutoArchiveDuration = 1440  // 1 day
	ThreadAutoArchiveDuration4320  ThreadAutoArchiveDuration = 4320  // 3 days
	ThreadAutoArchiveDuration10080 ThreadAutoArchiveDuration = 10080 // 1 week
)

// BasePartialChannel represents the base partial channel structure
type BasePartialChannel struct {
	ID   Snowflake   `json:"id"`
	Type ChannelType `json:"type"`
}

// NameableChannel represents a channel that can have a name
type NameableChannel struct {
	Name *string `json:"name,omitempty"`
}

// PartialChannel represents a partial channel (id, name, type only)
type PartialChannel struct {
	BasePartialChannel
	NameableChannel
}

// InviteChannel represents a channel obtained from fetching an invite
type InviteChannel struct {
	PartialChannel
	Icon       *string       `json:"icon,omitempty"`
	Recipients []PartialUser `json:"recipients,omitempty"` // Only usernames
}

// ChannelBase represents the base channel structure with type parameter concept
type ChannelBase struct {
	BasePartialChannel
	Flags *ChannelFlags `json:"flags,omitempty"`
}

// SlowmodeChannel represents a channel with slowmode capabilities
type SlowmodeChannel struct {
	ChannelBase
	RateLimitPerUser *int `json:"rate_limit_per_user,omitempty"`
}

// SortableChannel represents a channel that can be sorted
type SortableChannel struct {
	Position int `json:"position"`
}

// TextBasedChannel represents a text-based channel
type TextBasedChannel struct {
	SlowmodeChannel
	LastMessageID *Snowflake `json:"last_message_id,omitempty"`
}

// PinChannel represents a channel that supports pinned messages
type PinChannel struct {
	ChannelBase
	LastPinTimestamp *time.Time `json:"last_pin_timestamp,omitempty"`
}

// GuildChannel represents a guild channel
type GuildChannel struct {
	ChannelBase
	Name                 string      `json:"name"`
	GuildID              *Snowflake  `json:"guild_id,omitempty"`
	PermissionOverwrites []Overwrite `json:"permission_overwrites,omitempty"`
	ParentID             *Snowflake  `json:"parent_id,omitempty"`
	NSFW                 *bool       `json:"nsfw,omitempty"`
}

// GuildTextChannel represents a guild text channel
type GuildTextChannel struct {
	GuildChannel
	SortableChannel
	LastMessageID                 *Snowflake                 `json:"last_message_id,omitempty"`
	RateLimitPerUser              *int                       `json:"rate_limit_per_user,omitempty"`
	LastPinTimestamp              *time.Time                 `json:"last_pin_timestamp,omitempty"`
	DefaultAutoArchiveDuration    *ThreadAutoArchiveDuration `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser *int                       `json:"default_thread_rate_limit_per_user,omitempty"`
	Topic                         *string                    `json:"topic,omitempty"`
}

// TextChannel represents a guild text channel (consolidated with NewsChannel)
// Use GuildTextChannel directly and set Type field based on actual channel type
type TextChannel = GuildTextChannel
type NewsChannel = GuildTextChannel

// GuildCategoryChannel represents a guild category channel
type GuildCategoryChannel struct {
	GuildChannel
	SortableChannel
	Type ChannelType `json:"type"` // Should be ChannelTypeGuildCategory
}

// VoiceChannelBase represents base voice channel structure
type VoiceChannelBase struct {
	GuildChannel
	SortableChannel
	LastMessageID    *Snowflake        `json:"last_message_id,omitempty"`
	RateLimitPerUser *int              `json:"rate_limit_per_user,omitempty"`
	Bitrate          *int              `json:"bitrate,omitempty"`
	UserLimit        *int              `json:"user_limit,omitempty"`
	RTCRegion        *string           `json:"rtc_region,omitempty"`
	VideoQualityMode *VideoQualityMode `json:"video_quality_mode,omitempty"`
}

// VoiceChannel represents a guild voice channel (consolidated with StageChannel)
// Use VoiceChannelBase directly and set Type field based on actual channel type
type VoiceChannel = VoiceChannelBase
type StageChannel = VoiceChannelBase

// ForumChannel represents a guild forum channel
type ForumChannel struct {
	GuildChannel
	SortableChannel
	LastMessageID                 *Snowflake                 `json:"last_message_id,omitempty"`
	RateLimitPerUser              *int                       `json:"rate_limit_per_user,omitempty"`
	Topic                         *string                    `json:"topic,omitempty"`
	DefaultAutoArchiveDuration    *ThreadAutoArchiveDuration `json:"default_auto_archive_duration,omitempty"`
	DefaultThreadRateLimitPerUser *int                       `json:"default_thread_rate_limit_per_user,omitempty"`
	DefaultForumLayout            *ForumLayoutType           `json:"default_forum_layout,omitempty"`
	DefaultSortOrder              *SortOrderType             `json:"default_sort_order,omitempty"`
	AvailableTags                 []ForumTag                 `json:"available_tags,omitempty"`
	DefaultReactionEmoji          *DefaultReaction           `json:"default_reaction_emoji,omitempty"`
}

// MediaChannel represents a guild media channel (consolidated with ForumChannel)
// Use ForumChannel directly and set Type field based on actual channel type
type MediaChannel = ForumChannel

// DMChannel represents a direct message channel
type DMChannel struct {
	TextBasedChannel
	PinChannel
	Recipients []User      `json:"recipients"`
	Type       ChannelType `json:"type"` // Should be ChannelTypeDM
}

// GroupDMChannel represents a group direct message channel
type GroupDMChannel struct {
	TextBasedChannel
	PinChannel
	Name          *string     `json:"name,omitempty"`
	Icon          *string     `json:"icon,omitempty"`
	Recipients    []User      `json:"recipients"`
	OwnerID       Snowflake   `json:"owner_id"`
	ApplicationID *Snowflake  `json:"application_id,omitempty"`
	Managed       *bool       `json:"managed,omitempty"`
	Type          ChannelType `json:"type"` // Should be ChannelTypeGroupDM
}

// ThreadChannel represents a thread channel
type ThreadChannel struct {
	GuildChannel
	Member           *ThreadMember  `json:"member,omitempty"`
	ThreadMetadata   ThreadMetadata `json:"thread_metadata"`
	MessageCount     *int           `json:"message_count,omitempty"`
	MemberCount      *int           `json:"member_count,omitempty"`
	RateLimitPerUser *int           `json:"rate_limit_per_user,omitempty"`
	TotalMessageSent *int           `json:"total_message_sent,omitempty"`
	AppliedTags      []Snowflake    `json:"applied_tags,omitempty"`
}

// Thread channel types (consolidated)
// Use ThreadChannel directly and set Type field based on actual thread type
type PublicThreadChannel = ThreadChannel
type PrivateThreadChannel = ThreadChannel
type AnnouncementThreadChannel = ThreadChannel

// Overwrite represents a permission overwrite for a channel
// https://discord.com/developers/docs/resources/channel#overwrite-object
type Overwrite struct {
	ID    Snowflake     `json:"id"`
	Type  OverwriteType `json:"type"`
	Allow Permissions   `json:"allow"`
	Deny  Permissions   `json:"deny"`
}

// OverwriteType represents the type of permission overwrite
type OverwriteType int

const (
	OverwriteTypeRole OverwriteType = iota
	OverwriteTypeMember
)

// ThreadMetadata represents thread-specific metadata
type ThreadMetadata struct {
	Archived            bool       `json:"archived"`
	AutoArchiveDuration int        `json:"auto_archive_duration"`
	ArchiveTimestamp    time.Time  `json:"archive_timestamp"`
	Locked              *bool      `json:"locked,omitempty"`
	Invitable           *bool      `json:"invitable,omitempty"`
	CreateTimestamp     *time.Time `json:"create_timestamp,omitempty"`
}

// ThreadMember represents a thread member
type ThreadMember struct {
	ID            *Snowflake   `json:"id,omitempty"`
	UserID        *Snowflake   `json:"user_id,omitempty"`
	JoinTimestamp time.Time    `json:"join_timestamp"`
	Flags         int          `json:"flags"`
	Member        *GuildMember `json:"member,omitempty"`
}

// ForumTag represents a tag that can be applied to forum threads
type ForumTag struct {
	ID        Snowflake  `json:"id"`
	Name      string     `json:"name"`
	Moderated bool       `json:"moderated"`
	EmojiID   *Snowflake `json:"emoji_id,omitempty"`
	EmojiName *string    `json:"emoji_name,omitempty"`
}

// DefaultReaction represents the default reaction for forum posts
type DefaultReaction struct {
	EmojiID   *Snowflake `json:"emoji_id,omitempty"`
	EmojiName *string    `json:"emoji_name,omitempty"`
}

// ForumLayoutType represents the layout type for forum channels
type ForumLayoutType int

const (
	ForumLayoutTypeNotSet ForumLayoutType = iota
	ForumLayoutTypeListView
	ForumLayoutTypeGalleryView
)

// SortOrderType represents the sort order for forum channels
type SortOrderType int

const (
	SortOrderTypeLatestActivity SortOrderType = iota
	SortOrderTypeCreationDate
)

// FollowedChannel represents a followed channel
//
// https://discord.com/developers/docs/resources/channel#followed-channel-object
type FollowedChannel struct {
	ChannelID Snowflake `json:"channel_id"`
	WebhookID Snowflake `json:"webhook_id"`
}

// ChannelMention represents a channel mention in a message
//
// https://discord.com/developers/docs/resources/channel#channel-mention-object
type ChannelMention struct {
	ID      Snowflake   `json:"id"`
	GuildID Snowflake   `json:"guild_id"`
	Type    ChannelType `json:"type"`
	Name    string      `json:"name"`
}
