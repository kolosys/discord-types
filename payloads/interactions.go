package payloads

import (
	"time"

	"github.com/kolosys/discord-types/discord"
)

// InteractionType represents the type of interaction
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
type InteractionType int

const (
	InteractionTypePing InteractionType = iota + 1
	InteractionTypeApplicationCommand
	InteractionTypeMessageComponent
	InteractionTypeApplicationCommandAutocomplete
	InteractionTypeModalSubmit
)

// InteractionResponseType represents the type of interaction response
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
type InteractionResponseType int

const (
	InteractionResponseTypePong InteractionResponseType = iota + 1
	_                                                   // Skip 2 and 3
	_
	InteractionResponseTypeChannelMessageWithSource
	InteractionResponseTypeDeferredChannelMessageWithSource
	InteractionResponseTypeDeferredMessageUpdate
	InteractionResponseTypeUpdateMessage
	InteractionResponseTypeApplicationCommandAutocompleteResult
	InteractionResponseTypeModal
	InteractionResponseTypePremiumRequired
	InteractionResponseTypeLaunchActivity
)

// ApplicationIntegrationType is defined in shared.go

// InteractionContextType represents the context where an interaction was triggered
type InteractionContextType int

const (
	InteractionContextTypeGuild InteractionContextType = iota
	InteractionContextTypeBotDM
	InteractionContextTypePrivateChannel
)

// BaseInteraction represents the base structure for all interactions
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object
type BaseInteraction struct {
	ID                           discord.Snowflake               `json:"id"`
	ApplicationID                discord.Snowflake               `json:"application_id"`
	Type                         InteractionType                 `json:"type"`
	Guild                        *PartialInteractionGuild        `json:"guild,omitempty"`
	GuildID                      *discord.Snowflake              `json:"guild_id,omitempty"`
	Channel                      *PartialChannel                 `json:"channel,omitempty"`
	ChannelID                    *discord.Snowflake              `json:"channel_id,omitempty"` // Deprecated
	Member                       *InteractionGuildMember         `json:"member,omitempty"`
	User                         *User                           `json:"user,omitempty"`
	Token                        string                          `json:"token"`
	Version                      int                             `json:"version"` // Always 1
	Message                      *Message                        `json:"message,omitempty"`
	AppPermissions               discord.Permissions             `json:"app_permissions"`
	Locale                       Locale                          `json:"locale"`
	GuildLocale                  *Locale                         `json:"guild_locale,omitempty"`
	Entitlements                 []Entitlement                   `json:"entitlements"`
	AuthorizingIntegrationOwners AuthorizingIntegrationOwnersMap `json:"authorizing_integration_owners"`
	Context                      *InteractionContextType         `json:"context,omitempty"`
	AttachmentSizeLimit          int                             `json:"attachment_size_limit"`
}

// AuthorizingIntegrationOwnersMap maps integration types to their owner IDs
type AuthorizingIntegrationOwnersMap map[ApplicationIntegrationType]discord.Snowflake

// InteractionGuildMember represents a guild member in an interaction context
// https://discord.com/developers/docs/resources/guild#guild-member-object
type InteractionGuildMember struct {
	GuildMember
	Permissions discord.Permissions `json:"permissions"`
	User        User                `json:"user"`
}

// PartialInteractionGuild represents a partial guild in interaction context
type PartialInteractionGuild struct {
	ID       discord.Snowflake `json:"id"`
	Locale   Locale            `json:"locale"`
	Features []string          `json:"features"`
}

// MessageInteraction represents message interaction data
// https://discord.com/developers/docs/interactions/receiving-and-responding#message-interaction-object
type MessageInteraction struct {
	ID     discord.Snowflake                     `json:"id"`
	Type   InteractionType                       `json:"type"`
	Name   string                                `json:"name"`
	User   User                                  `json:"user"`
	Member *PartialMessageInteractionGuildMember `json:"member,omitempty"`
}

// PartialMessageInteractionGuildMember represents a partial guild member for message interactions
type PartialMessageInteractionGuildMember struct {
	Avatar                     *string             `json:"avatar,omitempty"`
	CommunicationDisabledUntil *time.Time          `json:"communication_disabled_until,omitempty"`
	Deaf                       *bool               `json:"deaf,omitempty"`
	JoinedAt                   time.Time           `json:"joined_at"`
	Mute                       *bool               `json:"mute,omitempty"`
	Nick                       *string             `json:"nick,omitempty"`
	Pending                    *bool               `json:"pending,omitempty"`
	PremiumSince               *time.Time          `json:"premium_since,omitempty"`
	Roles                      []discord.Snowflake `json:"roles"`
}

// ApplicationCommandInteraction represents an application command interaction
type ApplicationCommandInteraction struct {
	BaseInteraction
	Data ApplicationCommandInteractionData `json:"data"`
}

// ApplicationCommandInteractionData represents the data for application command interactions
type ApplicationCommandInteractionData struct {
	ID       discord.Snowflake                         `json:"id"`
	Name     string                                    `json:"name"`
	Type     ApplicationCommandType                    `json:"type"`
	Resolved *InteractionDataResolved                  `json:"resolved,omitempty"`
	Options  []ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	GuildID  *discord.Snowflake                        `json:"guild_id,omitempty"`
	TargetID *discord.Snowflake                        `json:"target_id,omitempty"`
}

// ApplicationCommandType represents the type of application command
type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = iota + 1
	ApplicationCommandTypeUser
	ApplicationCommandTypeMessage
)

// ApplicationCommandInteractionDataOption represents an option in command interaction data
type ApplicationCommandInteractionDataOption struct {
	Name    string                                    `json:"name"`
	Type    ApplicationCommandOptionType              `json:"type"`
	Value   interface{}                               `json:"value,omitempty"`
	Options []ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	Focused *bool                                     `json:"focused,omitempty"`
}

// ApplicationCommandOptionType represents the type of application command option
type ApplicationCommandOptionType int

const (
	ApplicationCommandOptionTypeSubCommand ApplicationCommandOptionType = iota + 1
	ApplicationCommandOptionTypeSubCommandGroup
	ApplicationCommandOptionTypeString
	ApplicationCommandOptionTypeInteger
	ApplicationCommandOptionTypeBoolean
	ApplicationCommandOptionTypeUser
	ApplicationCommandOptionTypeChannel
	ApplicationCommandOptionTypeRole
	ApplicationCommandOptionTypeMentionable
	ApplicationCommandOptionTypeNumber
	ApplicationCommandOptionTypeAttachment
)

// MessageComponentInteraction represents a message component interaction
type MessageComponentInteraction struct {
	BaseInteraction
	Data MessageComponentInteractionData `json:"data"`
}

// MessageComponentInteractionData represents the data for message component interactions
type MessageComponentInteractionData struct {
	CustomID      string                   `json:"custom_id"`
	ComponentType ComponentType            `json:"component_type"`
	Values        []string                 `json:"values,omitempty"`
	Resolved      *InteractionDataResolved `json:"resolved,omitempty"`
}

// ComponentType represents the type of message component
type ComponentType int

const (
	ComponentTypeActionRow ComponentType = iota + 1
	ComponentTypeButton
	ComponentTypeStringSelect
	ComponentTypeTextInput
	ComponentTypeUserSelect
	ComponentTypeRoleSelect
	ComponentTypeMentionableSelect
	ComponentTypeChannelSelect
)

// ModalSubmitInteraction represents a modal submit interaction
type ModalSubmitInteraction struct {
	BaseInteraction
	Data ModalSubmitInteractionData `json:"data"`
}

// ModalSubmitInteractionData represents the data for modal submit interactions
type ModalSubmitInteractionData struct {
	CustomID   string                          `json:"custom_id"`
	Components []ModalSubmitActionRowComponent `json:"components"`
}

// ModalSubmitActionRowComponent represents an action row component in modal submission
type ModalSubmitActionRowComponent struct {
	Type       ComponentType                   `json:"type"`
	Components []ModalSubmitTextInputComponent `json:"components"`
}

// ModalSubmitTextInputComponent represents a text input component in modal submission
type ModalSubmitTextInputComponent struct {
	Type     ComponentType `json:"type"`
	CustomID string        `json:"custom_id"`
	Value    string        `json:"value"`
}

// AutocompleteInteraction represents an autocomplete interaction
type AutocompleteInteraction struct {
	BaseInteraction
	Data ApplicationCommandInteractionData `json:"data"`
}

// PingInteraction represents a ping interaction
type PingInteraction struct {
	BaseInteraction
}

// InteractionDataResolved represents resolved data for interactions
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-resolved-data-structure
type InteractionDataResolved struct {
	Users       map[discord.Snowflake]User                               `json:"users,omitempty"`
	Members     map[discord.Snowflake]InteractionDataResolvedGuildMember `json:"members,omitempty"`
	Roles       map[discord.Snowflake]Role                               `json:"roles,omitempty"`
	Channels    map[discord.Snowflake]InteractionDataResolvedChannel     `json:"channels,omitempty"`
	Attachments map[discord.Snowflake]Attachment                         `json:"attachments,omitempty"`
}

// InteractionDataResolvedGuildMember represents a resolved guild member in interaction data
type InteractionDataResolvedGuildMember struct {
	BaseGuildMember
	Permissions discord.Permissions `json:"permissions"`
}

// InteractionDataResolvedChannel represents a resolved channel in interaction data
type InteractionDataResolvedChannel struct {
	PartialChannel
	Permissions    discord.Permissions `json:"permissions"`
	ParentID       *discord.Snowflake  `json:"parent_id,omitempty"`
	ThreadMetadata *ThreadMetadata     `json:"thread_metadata,omitempty"`
}

// Interaction response types

// InteractionResponse represents an interaction response
type InteractionResponse interface {
	GetType() InteractionResponseType
}

// InteractionResponsePong represents a pong response
type InteractionResponsePong struct {
	Type InteractionResponseType `json:"type"`
}

func (r InteractionResponsePong) GetType() InteractionResponseType {
	return InteractionResponseTypePong
}

// InteractionResponseChannelMessageWithSource represents a message response
type InteractionResponseChannelMessageWithSource struct {
	Type InteractionResponseType         `json:"type"`
	Data InteractionResponseCallbackData `json:"data"`
}

func (r InteractionResponseChannelMessageWithSource) GetType() InteractionResponseType {
	return InteractionResponseTypeChannelMessageWithSource
}

// InteractionResponseCallbackData represents callback data for interaction responses
type InteractionResponseCallbackData struct {
	TTS             *bool                `json:"tts,omitempty"`
	Content         *string              `json:"content,omitempty"`
	Embeds          []Embed              `json:"embeds,omitempty"`
	AllowedMentions *AllowedMentions     `json:"allowed_mentions,omitempty"`
	Flags           *MessageFlags        `json:"flags,omitempty"`
	Components      []ActionRowComponent `json:"components,omitempty"`
	Attachments     []PartialAttachment  `json:"attachments,omitempty"`
}

// MessageFlags represents message flags
type MessageFlags int

const (
	MessageFlagCrossposted                      MessageFlags = 1 << 0
	MessageFlagIsCrosspost                      MessageFlags = 1 << 1
	MessageFlagSuppressEmbeds                   MessageFlags = 1 << 2
	MessageFlagSourceMessageDeleted             MessageFlags = 1 << 3
	MessageFlagUrgent                           MessageFlags = 1 << 4
	MessageFlagHasThread                        MessageFlags = 1 << 5
	MessageFlagEphemeral                        MessageFlags = 1 << 6
	MessageFlagLoading                          MessageFlags = 1 << 7
	MessageFlagFailedToMentionSomeRolesInThread MessageFlags = 1 << 8
	MessageFlagSuppressNotifications            MessageFlags = 1 << 12
	MessageFlagIsVoiceMessage                   MessageFlags = 1 << 13
)

// AutocompleteResponse represents an autocomplete response
type AutocompleteResponse struct {
	Type InteractionResponseType                            `json:"type"`
	Data CommandAutocompleteInteractionResponseCallbackData `json:"data"`
}

func (r AutocompleteResponse) GetType() InteractionResponseType {
	return InteractionResponseTypeApplicationCommandAutocompleteResult
}

// CommandAutocompleteInteractionResponseCallbackData represents autocomplete callback data
type CommandAutocompleteInteractionResponseCallbackData struct {
	Choices []ApplicationCommandOptionChoice `json:"choices"`
}

// ApplicationCommandOptionChoice represents a choice for application command options
type ApplicationCommandOptionChoice struct {
	Name              string          `json:"name"`
	NameLocalizations LocalizationMap `json:"name_localizations,omitempty"`
	Value             interface{}     `json:"value"`
}

// ModalResponse represents a modal response
type ModalResponse struct {
	Type InteractionResponseType              `json:"type"`
	Data ModalInteractionResponseCallbackData `json:"data"`
}

func (r ModalResponse) GetType() InteractionResponseType {
	return InteractionResponseTypeModal
}

// ModalInteractionResponseCallbackData represents modal callback data
type ModalInteractionResponseCallbackData struct {
	CustomID   string               `json:"custom_id"`
	Title      string               `json:"title"`
	Components []ActionRowComponent `json:"components"`
}

// Forward declarations for types defined in other files
// Message, Embed, AllowedMentions are defined in message.go
// ActionRowComponent will be defined in components.go
// PartialAttachment, Attachment are defined in message.go
// Entitlement is defined in monetization.go
// Locale is defined in common.go
// Guild member components are defined in guild_member.go

// Message interaction metadata types

// BaseInteractionMetadata represents base interaction metadata
type BaseInteractionMetadata struct {
	ID                           discord.Snowflake               `json:"id"`
	Type                         InteractionType                 `json:"type"`
	User                         User                            `json:"user"`
	AuthorizingIntegrationOwners AuthorizingIntegrationOwnersMap `json:"authorizing_integration_owners"`
	OriginalResponseMessageID    *discord.Snowflake              `json:"original_response_message_id,omitempty"`
}

// ApplicationCommandInteractionMetadata represents application command interaction metadata
type ApplicationCommandInteractionMetadata struct {
	BaseInteractionMetadata
	TargetUser      *User              `json:"target_user,omitempty"`
	TargetMessageID *discord.Snowflake `json:"target_message_id,omitempty"`
}

// MessageComponentInteractionMetadata represents message component interaction metadata
type MessageComponentInteractionMetadata struct {
	BaseInteractionMetadata
	InteractedMessageID discord.Snowflake `json:"interacted_message_id"`
}

// ModalSubmitInteractionMetadata represents modal submit interaction metadata
type ModalSubmitInteractionMetadata struct {
	BaseInteractionMetadata
	TriggeringInteractionMetadata interface{} `json:"triggering_interaction_metadata"` // Can be ApplicationCommand or MessageComponent
}
