package v10

// Interaction represents an interaction with a Discord application.
// Reference: https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-structure
type Interaction struct {
	// ID of the interaction
	ID Snowflake `json:"id"`

	// ID of the application that the interaction was sent from
	ApplicationID Snowflake `json:"application_id"`

	// Type of interaction
	Type InteractionType `json:"type"`

	// This would embed the data for the interaction
	// Can be PING(1), APPLICATION_COMMAND(2), MESSAGE_COMPONENT(3), APPLICATION_COMMAND_AUTOCOMPLETE(4), MODAL_SUBMIT(5)
	// This is always present on application command, message component, and modal submit interaction types. It is optional for future-proofing against new interaction types
	Data any `json:"data"`

	// Guild that the interaction was sent from
	Guild *PartialGuild `json:"guild,omitempty"`

	// ID of the guild that the interaction was sent from
	GuildID *Snowflake `json:"guild_id,omitempty"`

	// Channel that the interaction was sent from
	Channel *PartialChannel `json:"channel,omitempty"`

	// ID of the channel that the interaction was sent to
	ChannelID *Snowflake `json:"channel_id,omitempty"`

	// Guild member data for the invoking user, including permissions
	// Member is sent when the interaction is invoked in a guild, and user is sent when invoked in a DM
	Member *GuildMember `json:"member,omitempty"`

	// User object for the invoking user, if invoked in a DM
	User *User `json:"user,omitempty"`

	// Continuation token for responding to the interaction
	Token string `json:"token"`

	// Read-only property, always 1
	Version int `json:"version" default:"1"`

	// For components or modals triggered by components, the message they were attached to
	Message *Message `json:"message,omitempty"`

	// Bitwise set of permissions the app has in the source location of the interaction
	AppPermissions Permissions `json:"app_permissions"`

	// Selected language of the invoking user
	Locale Locale `json:"locale"`
	// Guild's preferred locale, if invoked in a guild
	GuildLocale *Locale `json:"guild_locale,omitempty"`

	// For monetized apps, any entitlements for the invoking user, representing access to premium SKUs
	Entitlements []Entitlement `json:"entitlements"`

	// Mapping of installation contexts that the interaction was authorized for to related user or guild IDs.
	// See Authorizing Integration Owners Object for details
	AuthorizingIntegrationOwners AuthorizingIntegrationOwnersMap `json:"authorizing_integration_owners"`

	//Context where the interaction was triggered from
	Context *InteractionContextType `json:"context,omitempty"`

	// Attachment size limit in bytes
	AttachmentSizeLimit int `json:"attachment_size_limit"`
}

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

// String implements fmt.Stringer for debugging
func (t InteractionType) String() string {
	switch t {
	case InteractionTypePing:
		return "PING"
	case InteractionTypeApplicationCommand:
		return "APPLICATION_COMMAND"
	case InteractionTypeMessageComponent:
		return "MESSAGE_COMPONENT"
	case InteractionTypeApplicationCommandAutocomplete:
		return "APPLICATION_COMMAND_AUTOCOMPLETE"
	case InteractionTypeModalSubmit:
		return "MODAL_SUBMIT"
	default:
		return "UNKNOWN"
	}
}

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

// AuthorizingIntegrationOwnersMap maps integration types to their owner IDs
type AuthorizingIntegrationOwnersMap map[ApplicationIntegrationType]Snowflake

// MessageInteraction represents message interaction data
// https://discord.com/developers/docs/interactions/receiving-and-responding#message-interaction-object
type MessageInteraction struct {
	ID     Snowflake       `json:"id"`
	Type   InteractionType `json:"type"`
	Name   string          `json:"name"`
	User   User            `json:"user"`
	Member *GuildMember    `json:"member,omitempty"`
}

// ApplicationCommandInteraction represents an application command interaction
type ApplicationCommandInteraction struct {
	Interaction
	Data ApplicationCommandInteractionData `json:"data"`
}

// ApplicationCommandInteractionData represents the data for application command interactions
type ApplicationCommandInteractionData struct {
	ID       Snowflake                                 `json:"id"`
	Name     string                                    `json:"name"`
	Type     ApplicationCommandType                    `json:"type"`
	Resolved *InteractionDataResolved                  `json:"resolved,omitempty"`
	Options  []ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	GuildID  *Snowflake                                `json:"guild_id,omitempty"`
	TargetID *Snowflake                                `json:"target_id,omitempty"`
}

// ApplicationCommandInteractionDataOption represents an option in command interaction data
type ApplicationCommandInteractionDataOption struct {
	Name    string                                    `json:"name"`
	Type    ApplicationCommandOptionType              `json:"type"`
	Value   any                                       `json:"value,omitempty"`
	Options []ApplicationCommandInteractionDataOption `json:"options,omitempty"`
	Focused *bool                                     `json:"focused,omitempty"`
}

// MessageComponentInteraction represents a message component interaction
type MessageComponentInteraction struct {
	Interaction
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
	Interaction
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
	Interaction
	Data ApplicationCommandInteractionData `json:"data"`
}

// PingInteraction represents a ping interaction
type PingInteraction = Interaction

// InteractionDataResolved represents resolved data for interactions
// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-resolved-data-structure
type InteractionDataResolved struct {
	Users       map[Snowflake]User                       `json:"users,omitempty"`
	Members     map[Snowflake]GuildMember                `json:"members,omitempty"`
	Roles       map[Snowflake]Role                       `json:"roles,omitempty"`
	Channels    map[Snowflake]InteractionResolvedChannel `json:"channels,omitempty"`
	Attachments map[Snowflake]Attachment                 `json:"attachments,omitempty"`
}

// InteractionResolvedChannel represents a resolved channel in interaction data
type InteractionResolvedChannel struct {
	PartialChannel
	Permissions    Permissions     `json:"permissions"`
	ParentID       *Snowflake      `json:"parent_id,omitempty"`
	ThreadMetadata *ThreadMetadata `json:"thread_metadata,omitempty"`
}

// InteractionResponse represents an interaction response
type InteractionResponse[T any] struct {
	Type InteractionResponseType `json:"type"`
	Data T                       `json:"data"`
}

// InteractionResponsePong represents a pong response
type InteractionResponsePong = InteractionResponse[any]

// InteractionResponseChannelMessageWithSource represents a message response
type InteractionResponseChannelMessageWithSource = InteractionResponse[InteractionResponseCallbackData]

// InteractionResponseDeferredChannelMessageWithSource represents a deferred message response
type InteractionResponseDeferredChannelMessageWithSource = InteractionResponse[*InteractionResponseCallbackData]

// InteractionResponseDeferredMessageUpdate represents a deferred message update response
type InteractionResponseDeferredMessageUpdate = InteractionResponse[any]

// InteractionResponseUpdateMessage represents a message update response
type InteractionResponseUpdateMessage = InteractionResponse[InteractionResponseCallbackData]

// InteractionResponsePremiumRequired represents a premium required response
type InteractionResponsePremiumRequired = InteractionResponse[any]

// InteractionResponseLaunchActivity represents a launch activity response
type InteractionResponseLaunchActivity = InteractionResponse[any]

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
type AutocompleteResponse = InteractionResponse[ApplicationCommandOptionChoices]

// CommandAutocompleteInteractionResponseCallbackData represents autocomplete callback data
type ApplicationCommandOptionChoices struct {
	Choices []ApplicationCommandOptionChoice `json:"choices"`
}

// ApplicationCommandOptionChoice represents a choice for application command options
type ApplicationCommandOptionChoice struct {
	Name              string          `json:"name"`
	NameLocalizations LocalizationMap `json:"name_localizations,omitempty"`
	Value             any             `json:"value"`
}

// ModalResponse represents a modal response
type ModalResponse = InteractionResponse[ModalInteractionResponseCallbackData]

// ModalInteractionResponseCallbackData represents modal callback data
type ModalInteractionResponseCallbackData struct {
	CustomID   string               `json:"custom_id"`
	Title      string               `json:"title"`
	Components []ActionRowComponent `json:"components"`
}

// InteractionMetadata represents interaction metadata
type InteractionMetadata struct {
	ID                            Snowflake                       `json:"id"`
	Type                          InteractionType                 `json:"type"`
	User                          User                            `json:"user"`
	AuthorizingIntegrationOwners  AuthorizingIntegrationOwnersMap `json:"authorizing_integration_owners"`
	OriginalResponseMessageID     *Snowflake                      `json:"original_response_message_id,omitempty"`
	TargetUser                    *User                           `json:"target_user,omitempty"`
	TargetMessageID               *Snowflake                      `json:"target_message_id,omitempty"`
	InteractedMessageID           Snowflake                       `json:"interacted_message_id"`
	TriggeringInteractionMetadata any                             `json:"triggering_interaction_metadata"` // Can be ApplicationCommand or MessageComponent
}
