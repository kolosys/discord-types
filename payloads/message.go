package payloads

import (
	"time"

	"github.com/kolosys/discord-types/discord"
)

// MessageType represents the type of message
type MessageType int

const (
	MessageTypeDefault MessageType = iota
	MessageTypeRecipientAdd
	MessageTypeRecipientRemove
	MessageTypeCall
	MessageTypeChannelNameChange
	MessageTypeChannelIconChange
	MessageTypeChannelPinnedMessage
	MessageTypeUserJoin
	MessageTypeGuildBoost
	MessageTypeGuildBoostTier1
	MessageTypeGuildBoostTier2
	MessageTypeGuildBoostTier3
	MessageTypeChannelFollowAdd
	_
	MessageTypeGuildDiscoveryDisqualified
	MessageTypeGuildDiscoveryRequalified
	MessageTypeGuildDiscoveryGracePeriodInitialWarning
	MessageTypeGuildDiscoveryGracePeriodFinalWarning
	MessageTypeThreadCreated
	MessageTypeReply
	MessageTypeChatInputCommand
	MessageTypeThreadStarterMessage
	MessageTypeGuildInviteReminder
	MessageTypeContextMenuCommand
	MessageTypeAutoModerationAction
	MessageTypeRoleSubscriptionPurchase
	MessageTypeInteractionPremiumUpsell
	MessageTypeStageStart
	MessageTypeStageEnd
	MessageTypeStageSpeaker
	MessageTypeStageTopic
	MessageTypeGuildApplicationPremiumSubscription
)

// Message represents a Discord message
// https://discord.com/developers/docs/resources/channel#message-object-message-structure
type Message struct {
	ID                   discord.Snowflake        `json:"id"`
	ChannelID            discord.Snowflake        `json:"channel_id"`
	GuildID              *discord.Snowflake       `json:"guild_id,omitempty"`
	Author               User                     `json:"author"`
	Member               *GuildMember             `json:"member,omitempty"`
	Content              string                   `json:"content"`
	Timestamp            time.Time                `json:"timestamp"`
	EditedTimestamp      *time.Time               `json:"edited_timestamp"`
	TTS                  bool                     `json:"tts"`
	MentionEveryone      bool                     `json:"mention_everyone"`
	Mentions             []User                   `json:"mentions"`
	MentionRoles         []discord.Snowflake      `json:"mention_roles"`
	MentionChannels      []ChannelMention         `json:"mention_channels,omitempty"`
	Attachments          []Attachment             `json:"attachments"`
	Embeds               []Embed                  `json:"embeds"`
	Reactions            []Reaction               `json:"reactions,omitempty"`
	Nonce                interface{}              `json:"nonce,omitempty"` // string or number
	Pinned               bool                     `json:"pinned"`
	WebhookID            *discord.Snowflake       `json:"webhook_id,omitempty"`
	Type                 MessageType              `json:"type"`
	Activity             *MessageActivity         `json:"activity,omitempty"`
	Application          *Application             `json:"application,omitempty"`
	ApplicationID        *discord.Snowflake       `json:"application_id,omitempty"`
	MessageReference     *MessageReference        `json:"message_reference,omitempty"`
	Flags                *MessageFlags            `json:"flags,omitempty"`
	ReferencedMessage    *Message                 `json:"referenced_message,omitempty"`
	Interaction          *MessageInteraction      `json:"interaction,omitempty"`
	Thread               *ThreadChannel           `json:"thread,omitempty"`
	Components           []ActionRowComponent     `json:"components,omitempty"`
	StickerItems         []StickerItem            `json:"sticker_items,omitempty"`
	Stickers             []Sticker                `json:"stickers,omitempty"` // Deprecated
	Position             *int                     `json:"position,omitempty"`
	RoleSubscriptionData *RoleSubscriptionData    `json:"role_subscription_data,omitempty"`
	Resolved             *InteractionDataResolved `json:"resolved,omitempty"`
	Poll                 *Poll                    `json:"poll,omitempty"`
	Call                 *MessageCall             `json:"call,omitempty"`
}

// ChannelMention represents a channel mention in a message
type ChannelMention struct {
	ID      discord.Snowflake `json:"id"`
	GuildID discord.Snowflake `json:"guild_id"`
	Type    ChannelType       `json:"type"`
	Name    string            `json:"name"`
}

// Attachment represents a message attachment
// https://discord.com/developers/docs/resources/channel#attachment-object-attachment-structure
type Attachment struct {
	ID           discord.Snowflake `json:"id"`
	Filename     string            `json:"filename"`
	Description  *string           `json:"description,omitempty"`
	ContentType  *string           `json:"content_type,omitempty"`
	Size         int               `json:"size"`
	URL          string            `json:"url"`
	ProxyURL     string            `json:"proxy_url"`
	Height       *int              `json:"height,omitempty"`
	Width        *int              `json:"width,omitempty"`
	Ephemeral    *bool             `json:"ephemeral,omitempty"`
	DurationSecs *float64          `json:"duration_secs,omitempty"`
	Waveform     *string           `json:"waveform,omitempty"`
	Flags        *AttachmentFlags  `json:"flags,omitempty"`
}

// AttachmentFlags represents attachment flags
type AttachmentFlags int

const (
	AttachmentFlagIsRemix AttachmentFlags = 1 << 2
)

// PartialAttachment represents a partial attachment for uploading
type PartialAttachment struct {
	ID          discord.Snowflake `json:"id"`
	Filename    *string           `json:"filename,omitempty"`
	Description *string           `json:"description,omitempty"`
}

// Embed represents a message embed
// https://discord.com/developers/docs/resources/channel#embed-object-embed-structure
type Embed struct {
	Title       *string         `json:"title,omitempty"`
	Type        *EmbedType      `json:"type,omitempty"`
	Description *string         `json:"description,omitempty"`
	URL         *string         `json:"url,omitempty"`
	Timestamp   *time.Time      `json:"timestamp,omitempty"`
	Color       *int            `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []EmbedField    `json:"fields,omitempty"`
}

// EmbedType represents the type of embed
type EmbedType string

const (
	EmbedTypeRich    EmbedType = "rich"
	EmbedTypeImage   EmbedType = "image"
	EmbedTypeVideo   EmbedType = "video"
	EmbedTypeGifv    EmbedType = "gifv"
	EmbedTypeArticle EmbedType = "article"
	EmbedTypeLink    EmbedType = "link"
)

// EmbedFooter represents an embed footer
type EmbedFooter struct {
	Text         string  `json:"text"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

// EmbedImage represents an embed image
type EmbedImage struct {
	URL      string  `json:"url"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

// EmbedThumbnail represents an embed thumbnail
type EmbedThumbnail struct {
	URL      string  `json:"url"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

// EmbedVideo represents an embed video
type EmbedVideo struct {
	URL      *string `json:"url,omitempty"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

// EmbedProvider represents an embed provider
type EmbedProvider struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// EmbedAuthor represents an embed author
type EmbedAuthor struct {
	Name         string  `json:"name"`
	URL          *string `json:"url,omitempty"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

// EmbedField represents an embed field
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline *bool  `json:"inline,omitempty"`
}

// MessageActivity represents message activity
type MessageActivity struct {
	Type    MessageActivityType `json:"type"`
	PartyID *string             `json:"party_id,omitempty"`
}

// MessageActivityType represents the type of message activity
type MessageActivityType int

const (
	MessageActivityTypeJoin MessageActivityType = iota + 1
	MessageActivityTypeSpectate
	MessageActivityTypeListen
	MessageActivityTypeJoinRequest
)

// MessageReference represents a message reference
type MessageReference struct {
	MessageID       *discord.Snowflake `json:"message_id,omitempty"`
	ChannelID       *discord.Snowflake `json:"channel_id,omitempty"`
	GuildID         *discord.Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists *bool              `json:"fail_if_not_exists,omitempty"`
}

// AllowedMentions represents allowed mentions structure
type AllowedMentions struct {
	Parse       []AllowedMentionType `json:"parse"`
	Roles       []discord.Snowflake  `json:"roles,omitempty"`
	Users       []discord.Snowflake  `json:"users,omitempty"`
	RepliedUser *bool                `json:"replied_user,omitempty"`
}

// AllowedMentionType represents the type of allowed mention
type AllowedMentionType string

const (
	AllowedMentionTypeRoles    AllowedMentionType = "roles"
	AllowedMentionTypeUsers    AllowedMentionType = "users"
	AllowedMentionTypeEveryone AllowedMentionType = "everyone"
)

// Forward declarations for types defined in other files
// GuildMember is defined in guild_member.go
// ActionRowComponent is defined in components.go
// RoleSubscriptionData represents role subscription data.
type RoleSubscriptionData struct {
	// RoleSubscriptionListingID is the id of the role subscription listing.
	RoleSubscriptionListingID discord.Snowflake `json:"role_subscription_listing_id"`

	// TierName is the name of the tier.
	TierName string `json:"tier_name"`

	// TotalMonthsSubscribed is the total months subscribed.
	TotalMonthsSubscribed int `json:"total_months_subscribed"`

	// IsRenewal is whether this is a renewal.
	IsRenewal bool `json:"is_renewal"`

	// HasAccess is whether the user has access to the role.
	HasAccess bool `json:"has_access"`
}

// Poll is defined in poll.go
// MessageCall represents a message call.
type MessageCall struct {
	// EndedTimestamp is when the call ended.
	EndedTimestamp *string `json:"ended_timestamp,omitempty"`

	// Participants are the participants in the call.
	Participants []discord.Snowflake `json:"participants"`
}
