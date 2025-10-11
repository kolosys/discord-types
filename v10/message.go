package v10

import (
	"time"
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
	ID                   Snowflake                `json:"id"`
	ChannelID            Snowflake                `json:"channel_id"`
	GuildID              *Snowflake               `json:"guild_id,omitempty"`
	Author               User                     `json:"author"`
	Member               *GuildMember             `json:"member,omitempty"`
	Content              string                   `json:"content"`
	Timestamp            time.Time                `json:"timestamp"`
	EditedTimestamp      *time.Time               `json:"edited_timestamp"`
	TTS                  bool                     `json:"tts"`
	MentionEveryone      bool                     `json:"mention_everyone"`
	Mentions             []User                   `json:"mentions"`
	MentionRoles         []Snowflake              `json:"mention_roles"`
	MentionChannels      []ChannelMention         `json:"mention_channels,omitempty"`
	Attachments          []Attachment             `json:"attachments"`
	Embeds               []Embed                  `json:"embeds"`
	Reactions            []Reaction               `json:"reactions,omitempty"`
	Nonce                any                      `json:"nonce,omitempty"` // string or number
	Pinned               bool                     `json:"pinned"`
	WebhookID            *Snowflake               `json:"webhook_id,omitempty"`
	Type                 MessageType              `json:"type"`
	Activity             *MessageActivity         `json:"activity,omitempty"`
	Application          *Application             `json:"application,omitempty"`
	ApplicationID        *Snowflake               `json:"application_id,omitempty"`
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
	MessageID       *Snowflake `json:"message_id,omitempty"`
	ChannelID       *Snowflake `json:"channel_id,omitempty"`
	GuildID         *Snowflake `json:"guild_id,omitempty"`
	FailIfNotExists *bool      `json:"fail_if_not_exists,omitempty"`
}

// AllowedMentions represents allowed mentions structure
type AllowedMentions struct {
	Parse       []AllowedMentionType `json:"parse"`
	Roles       []Snowflake          `json:"roles,omitempty"`
	Users       []Snowflake          `json:"users,omitempty"`
	RepliedUser *bool                `json:"replied_user,omitempty"`
}

// AllowedMentionType represents the type of allowed mention
type AllowedMentionType string

const (
	AllowedMentionTypeRoles    AllowedMentionType = "roles"
	AllowedMentionTypeUsers    AllowedMentionType = "users"
	AllowedMentionTypeEveryone AllowedMentionType = "everyone"
)

// RoleSubscriptionData represents role subscription data.
type RoleSubscriptionData struct {
	// RoleSubscriptionListingID is the id of the role subscription listing.
	RoleSubscriptionListingID Snowflake `json:"role_subscription_listing_id"`

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
	Participants []Snowflake `json:"participants"`
}

// Reaction represents a message reaction
// https://discord.com/developers/docs/resources/channel#reaction-object-reaction-structure
type Reaction struct {
	Count        int                  `json:"count"`
	CountDetails ReactionCountDetails `json:"count_details"`
	Me           bool                 `json:"me"`
	MeBurst      bool                 `json:"me_burst"`
	Emoji        PartialEmoji         `json:"emoji"`
	BurstColors  []string             `json:"burst_colors,omitempty"`
}

// ReactionCountDetails provides detailed reaction count information
// https://discord.com/developers/docs/resources/channel#reaction-object-reaction-count-details-structure
type ReactionCountDetails struct {
	Burst  int `json:"burst"`
	Normal int `json:"normal"`
}

// ReactionType represents the type of reaction
type ReactionType int

const (
	ReactionTypeNormal ReactionType = iota
	ReactionTypeBurst
)
