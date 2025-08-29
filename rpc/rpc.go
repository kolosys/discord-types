package rpc

import "github.com/kolosys/discord-types/discord"

// RPCVersion represents the RPC version.
const RPCVersion = "1"

// =============================================================================
// RPC COMMANDS
// =============================================================================

// RPCCommands represents RPC commands.
//
// See: https://discord.com/developers/docs/topics/rpc#commands-and-events-rpc-commands
type RPCCommands string

// RPC command constants
const (
	// RPCCommandAcceptActivityInvite accepts an activity invite.
	RPCCommandAcceptActivityInvite RPCCommands = "ACCEPT_ACTIVITY_INVITE"

	// RPCCommandActivityInviteUser invites a user to an activity.
	RPCCommandActivityInviteUser RPCCommands = "ACTIVITY_INVITE_USER"

	// RPCCommandAuthenticate is used to authenticate an existing client with your app.
	RPCCommandAuthenticate RPCCommands = "AUTHENTICATE"

	// RPCCommandAuthorize is used to authorize a new client with your app.
	RPCCommandAuthorize RPCCommands = "AUTHORIZE"

	// RPCCommandBraintreePopupBridgeCallback handles Braintree popup bridge callback.
	RPCCommandBraintreePopupBridgeCallback RPCCommands = "BRAINTREE_POPUP_BRIDGE_CALLBACK"

	// RPCCommandBrowserHandoff handles browser handoff.
	RPCCommandBrowserHandoff RPCCommands = "BROWSER_HANDOFF"

	// RPCCommandCaptureShortcut captures a shortcut.
	RPCCommandCaptureShortcut RPCCommands = "CAPTURE_SHORTCUT"

	// RPCCommandCloseActivityJoinRequest is used to reject a Rich Presence Ask to Join request.
	RPCCommandCloseActivityJoinRequest RPCCommands = "CLOSE_ACTIVITY_JOIN_REQUEST"

	// RPCCommandConnectionsCallback handles connections callback.
	RPCCommandConnectionsCallback RPCCommands = "CONNECTIONS_CALLBACK"

	// RPCCommandCreateChannelInvite creates a channel invite.
	RPCCommandCreateChannelInvite RPCCommands = "CREATE_CHANNEL_INVITE"

	// RPCCommandDeepLink handles deep linking.
	RPCCommandDeepLink RPCCommands = "DEEP_LINK"

	// RPCCommandDispatch represents event dispatch.
	RPCCommandDispatch RPCCommands = "DISPATCH"

	// RPCCommandGetApplicationTicket gets an application ticket.
	RPCCommandGetApplicationTicket RPCCommands = "GET_APPLICATION_TICKET"

	// RPCCommandGetChannel is used to retrieve channel information from the client.
	RPCCommandGetChannel RPCCommands = "GET_CHANNEL"

	// RPCCommandGetChannels is used to retrieve a list of channels for a guild from the client.
	RPCCommandGetChannels RPCCommands = "GET_CHANNELS"

	// RPCCommandGetEntitlementTicket gets an entitlement ticket.
	RPCCommandGetEntitlementTicket RPCCommands = "GET_ENTITLEMENT_TICKET"

	// RPCCommandGetEntitlements gets entitlements.
	RPCCommandGetEntitlements RPCCommands = "GET_ENTITLEMENTS"

	// RPCCommandGetGuild is used to retrieve guild information from the client.
	RPCCommandGetGuild RPCCommands = "GET_GUILD"

	// RPCCommandGetGuilds is used to retrieve a list of guilds from the client.
	RPCCommandGetGuilds RPCCommands = "GET_GUILDS"

	// RPCCommandGetImage gets an image.
	RPCCommandGetImage RPCCommands = "GET_IMAGE"

	// RPCCommandGetPlatformBehaviors gets platform behaviors.
	RPCCommandGetPlatformBehaviors RPCCommands = "GET_PLATFORM_BEHAVIORS"

	// RPCCommandGetRelationships gets relationships.
	RPCCommandGetRelationships RPCCommands = "GET_RELATIONSHIPS"

	// RPCCommandGetSelectedVoiceChannel is used to retrieve voice channel information from the client.
	RPCCommandGetSelectedVoiceChannel RPCCommands = "GET_SELECTED_VOICE_CHANNEL"

	// RPCCommandGetSkus gets SKUs.
	RPCCommandGetSkus RPCCommands = "GET_SKUS"

	// RPCCommandGetThermalState gets thermal state.
	RPCCommandGetThermalState RPCCommands = "GET_THERMAL_STATE"

	// RPCCommandGetUser gets user information.
	RPCCommandGetUser RPCCommands = "GET_USER"

	// RPCCommandGetUserAchievements gets user achievements.
	RPCCommandGetUserAchievements RPCCommands = "GET_USER_ACHIEVEMENTS"

	// RPCCommandGetVoiceSettings is used to retrieve the client's voice settings.
	RPCCommandGetVoiceSettings RPCCommands = "GET_VOICE_SETTINGS"

	// RPCCommandInitiateImageUpload initiates image upload.
	RPCCommandInitiateImageUpload RPCCommands = "INITIATE_IMAGE_UPLOAD"

	// RPCCommandOpenExternalLink opens an external link.
	RPCCommandOpenExternalLink RPCCommands = "OPEN_EXTERNAL_LINK"

	// RPCCommandOpenInviteDialog opens invite dialog.
	RPCCommandOpenInviteDialog RPCCommands = "OPEN_INVITE_DIALOG"

	// RPCCommandOpenOverlayActivityInvite opens overlay activity invite.
	RPCCommandOpenOverlayActivityInvite RPCCommands = "OPEN_OVERLAY_ACTIVITY_INVITE"

	// RPCCommandOpenOverlayGuildInvite opens overlay guild invite.
	RPCCommandOpenOverlayGuildInvite RPCCommands = "OPEN_OVERLAY_GUILD_INVITE"

	// RPCCommandOpenOverlayVoiceSettings opens overlay voice settings.
	RPCCommandOpenOverlayVoiceSettings RPCCommands = "OPEN_OVERLAY_VOICE_SETTINGS"

	// RPCCommandOverlay handles overlay operations.
	RPCCommandOverlay RPCCommands = "OVERLAY"

	// RPCCommandSelectTextChannel is used to join or leave a text channel, group dm, or dm.
	RPCCommandSelectTextChannel RPCCommands = "SELECT_TEXT_CHANNEL"

	// RPCCommandSelectVoiceChannel is used to join or leave a voice channel, group dm, or dm.
	RPCCommandSelectVoiceChannel RPCCommands = "SELECT_VOICE_CHANNEL"

	// RPCCommandSendActivityJoinInvite is used to consent to a Rich Presence Ask to Join request.
	RPCCommandSendActivityJoinInvite RPCCommands = "SEND_ACTIVITY_JOIN_INVITE"

	// RPCCommandSetActivity is used to update a user's Rich Presence.
	RPCCommandSetActivity RPCCommands = "SET_ACTIVITY"

	// RPCCommandSetCertifiedDevices is used to send info about certified hardware devices.
	RPCCommandSetCertifiedDevices RPCCommands = "SET_CERTIFIED_DEVICES"

	// RPCCommandSetOverlayLocked sets overlay locked state.
	RPCCommandSetOverlayLocked RPCCommands = "SET_OVERLAY_LOCKED"

	// RPCCommandSetUserVoiceSettings is used to change voice settings of users in voice channels.
	RPCCommandSetUserVoiceSettings RPCCommands = "SET_USER_VOICE_SETTINGS"

	// RPCCommandSetUserVoiceSettings2 is the v2 version of SET_USER_VOICE_SETTINGS.
	RPCCommandSetUserVoiceSettings2 RPCCommands = "SET_USER_VOICE_SETTINGS_2"

	// RPCCommandSetVoiceSettings is used to set the client's voice settings.
	RPCCommandSetVoiceSettings RPCCommands = "SET_VOICE_SETTINGS"

	// RPCCommandSetVoiceSettings2 is the v2 version of SET_VOICE_SETTINGS.
	RPCCommandSetVoiceSettings2 RPCCommands = "SET_VOICE_SETTINGS_2"

	// RPCCommandStartPurchase starts a purchase.
	RPCCommandStartPurchase RPCCommands = "START_PURCHASE"

	// RPCCommandSubscribe subscribes to an event.
	RPCCommandSubscribe RPCCommands = "SUBSCRIBE"

	// RPCCommandUnsubscribe unsubscribes from an event.
	RPCCommandUnsubscribe RPCCommands = "UNSUBSCRIBE"

	// RPCCommandUserSettingsGetLocale gets user locale settings.
	RPCCommandUserSettingsGetLocale RPCCommands = "USER_SETTINGS_GET_LOCALE"
)

// =============================================================================
// RPC EVENTS
// =============================================================================

// RPCEvents represents RPC events.
//
// See: https://discord.com/developers/docs/topics/rpc#commands-and-events-rpc-events
type RPCEvents string

// RPC event constants
const (
	// RPCEventReady fires when the client is ready.
	RPCEventReady RPCEvents = "READY"

	// RPCEventError fires when an error occurs.
	RPCEventError RPCEvents = "ERROR"

	// RPCEventGuildStatus fires when guild status changes.
	RPCEventGuildStatus RPCEvents = "GUILD_STATUS"

	// RPCEventGuildCreate fires when a guild is created or becomes available.
	RPCEventGuildCreate RPCEvents = "GUILD_CREATE"

	// RPCEventChannelCreate fires when a channel is created.
	RPCEventChannelCreate RPCEvents = "CHANNEL_CREATE"

	// RPCEventRelationshipUpdate fires when a relationship is updated.
	RPCEventRelationshipUpdate RPCEvents = "RELATIONSHIP_UPDATE"

	// RPCEventVoiceChannelSelect fires when a voice channel is selected.
	RPCEventVoiceChannelSelect RPCEvents = "VOICE_CHANNEL_SELECT"

	// RPCEventVoiceStateCreate fires when a voice state is created.
	RPCEventVoiceStateCreate RPCEvents = "VOICE_STATE_CREATE"

	// RPCEventVoiceStateUpdate fires when a voice state is updated.
	RPCEventVoiceStateUpdate RPCEvents = "VOICE_STATE_UPDATE"

	// RPCEventVoiceStateDelete fires when a voice state is deleted.
	RPCEventVoiceStateDelete RPCEvents = "VOICE_STATE_DELETE"

	// RPCEventVoiceSettingsUpdate fires when voice settings are updated.
	RPCEventVoiceSettingsUpdate RPCEvents = "VOICE_SETTINGS_UPDATE"

	// RPCEventVoiceSettingsUpdate2 fires when voice settings are updated (v2).
	RPCEventVoiceSettingsUpdate2 RPCEvents = "VOICE_SETTINGS_UPDATE_2"

	// RPCEventVoiceConnectionStatus fires when voice connection status changes.
	RPCEventVoiceConnectionStatus RPCEvents = "VOICE_CONNECTION_STATUS"

	// RPCEventSpeakingStart fires when speaking starts.
	RPCEventSpeakingStart RPCEvents = "SPEAKING_START"

	// RPCEventSpeakingStop fires when speaking stops.
	RPCEventSpeakingStop RPCEvents = "SPEAKING_STOP"

	// RPCEventGameJoin fires when joining a game.
	RPCEventGameJoin RPCEvents = "GAME_JOIN"

	// RPCEventGameSpectate fires when spectating a game.
	RPCEventGameSpectate RPCEvents = "GAME_SPECTATE"

	// RPCEventActivityJoin fires when joining an activity.
	RPCEventActivityJoin RPCEvents = "ACTIVITY_JOIN"

	// RPCEventActivityJoinRequest fires when an activity join is requested.
	RPCEventActivityJoinRequest RPCEvents = "ACTIVITY_JOIN_REQUEST"

	// RPCEventActivitySpectate fires when spectating an activity.
	RPCEventActivitySpectate RPCEvents = "ACTIVITY_SPECTATE"

	// RPCEventActivityInvite fires when invited to an activity.
	RPCEventActivityInvite RPCEvents = "ACTIVITY_INVITE"

	// RPCEventNotificationCreate fires when a notification is created.
	RPCEventNotificationCreate RPCEvents = "NOTIFICATION_CREATE"

	// RPCEventMessageCreate fires when a message is created.
	RPCEventMessageCreate RPCEvents = "MESSAGE_CREATE"

	// RPCEventMessageUpdate fires when a message is updated.
	RPCEventMessageUpdate RPCEvents = "MESSAGE_UPDATE"

	// RPCEventMessageDelete fires when a message is deleted.
	RPCEventMessageDelete RPCEvents = "MESSAGE_DELETE"

	// RPCEventThermalStateUpdate fires when thermal state is updated.
	RPCEventThermalStateUpdate RPCEvents = "THERMAL_STATE_UPDATE"

	// RPCEventCurrentUserUpdate fires when current user is updated.
	RPCEventCurrentUserUpdate RPCEvents = "CURRENT_USER_UPDATE"

	// RPCEventEntitlementCreate fires when an entitlement is created.
	RPCEventEntitlementCreate RPCEvents = "ENTITLEMENT_CREATE"

	// RPCEventEntitlementDelete fires when an entitlement is deleted.
	RPCEventEntitlementDelete RPCEvents = "ENTITLEMENT_DELETE"
)

// =============================================================================
// RELATIONSHIP
// =============================================================================

// Relationship represents a user relationship.
type Relationship struct {
	// ID is the id of the user.
	ID discord.Snowflake `json:"id"`

	// Type is the relationship type.
	Type RelationshipType `json:"type"`

	// User is the user object.
	User interface{} `json:"user"` // APIUser
}

// =============================================================================
// RPC MESSAGE
// =============================================================================

// RPCAPIMessage represents an RPC message.
type RPCAPIMessage struct {
	// ID is the message ID.
	ID discord.Snowflake `json:"id"`

	// ChannelID is the channel ID.
	ChannelID discord.Snowflake `json:"channel_id"`

	// GuildID is the guild ID.
	GuildID *discord.Snowflake `json:"guild_id,omitempty"`

	// Author is the message author.
	Author interface{} `json:"author"` // APIUser

	// Member is the guild member who sent the message.
	Member interface{} `json:"member,omitempty"` // APIGuildMember

	// Content is the message content.
	Content string `json:"content"`

	// Timestamp is when the message was sent.
	Timestamp string `json:"timestamp"`

	// EditedTimestamp is when the message was edited.
	EditedTimestamp *string `json:"edited_timestamp"`

	// TTS indicates if this is a TTS message.
	TTS bool `json:"tts"`

	// MentionEveryone indicates if this message mentions everyone.
	MentionEveryone bool `json:"mention_everyone"`

	// Mentions are users mentioned in the message.
	Mentions []interface{} `json:"mentions"` // APIUser array

	// MentionRoles are roles mentioned in the message.
	MentionRoles []discord.Snowflake `json:"mention_roles"`

	// MentionChannels are channels mentioned in the message.
	MentionChannels []interface{} `json:"mention_channels,omitempty"` // APIChannelMention array

	// Attachments are message attachments.
	Attachments []interface{} `json:"attachments"` // APIAttachment array

	// Embeds are message embeds.
	Embeds []interface{} `json:"embeds"` // APIEmbed array

	// Reactions are message reactions.
	Reactions []interface{} `json:"reactions,omitempty"` // APIReaction array

	// Nonce is the message nonce.
	Nonce interface{} `json:"nonce,omitempty"` // string or number

	// Pinned indicates if the message is pinned.
	Pinned bool `json:"pinned"`

	// WebhookID is the webhook ID if sent by webhook.
	WebhookID *discord.Snowflake `json:"webhook_id,omitempty"`

	// Type is the message type.
	Type int `json:"type"`

	// Activity is the message activity.
	Activity interface{} `json:"activity,omitempty"` // APIMessageActivity

	// Application is the message application.
	Application interface{} `json:"application,omitempty"` // APIMessageApplication

	// ApplicationID is the application ID.
	ApplicationID *discord.Snowflake `json:"application_id,omitempty"`

	// MessageReference is the message reference.
	MessageReference interface{} `json:"message_reference,omitempty"` // APIMessageReference

	// Flags are message flags.
	Flags *int `json:"flags,omitempty"`

	// ReferencedMessage is the referenced message.
	ReferencedMessage interface{} `json:"referenced_message,omitempty"` // APIMessage

	// Interaction is the message interaction.
	Interaction interface{} `json:"interaction,omitempty"` // APIMessageInteraction

	// Thread is the thread started from this message.
	Thread interface{} `json:"thread,omitempty"` // APIThreadChannel

	// Components are message components.
	Components []interface{} `json:"components,omitempty"` // APIActionRowComponent array

	// StickerItems are sticker items.
	StickerItems []interface{} `json:"sticker_items,omitempty"` // APIStickerItem array

	// Stickers are deprecated stickers.
	Stickers []interface{} `json:"stickers,omitempty"` // APISticker array

	// Position is the message position.
	Position *int `json:"position,omitempty"`

	// RoleSubscriptionData is role subscription data.
	RoleSubscriptionData interface{} `json:"role_subscription_data,omitempty"` // APIRoleSubscriptionData

	// Resolved is resolved data.
	Resolved interface{} `json:"resolved,omitempty"` // APIInteractionDataResolved

	// Poll is the poll data.
	Poll interface{} `json:"poll,omitempty"` // APIPoll

	// Call is the call data.
	Call interface{} `json:"call,omitempty"` // APIMessageCall

	// Nick is the nickname of the user who sent the message.
	Nick *string `json:"nick,omitempty"`

	// AuthorColor is the color of the author's name.
	AuthorColor *int `json:"author_color,omitempty"`

	// ContentParsed is the content of the message parsed into an array.
	ContentParsed []interface{} `json:"content_parsed"` // (RPCAPIMessageParsedContentMention | RPCAPIMessageParsedContentText)[]
}

// =============================================================================
// RPC DATA TYPES
// =============================================================================

// RPCGuildCreateEventData represents data for guild create events.
type RPCGuildCreateEventData struct {
	// ID is the guild ID.
	ID discord.Snowflake `json:"id"`

	// Name is the guild name.
	Name string `json:"name"`
}

// RPCChannelCreateEventData represents data for channel create events.
type RPCChannelCreateEventData struct {
	// ID is the channel ID.
	ID discord.Snowflake `json:"id"`

	// Name is the channel name.
	Name string `json:"name"`

	// Type is the channel type.
	Type int `json:"type"`
}

// RPCVoiceStateEventData represents data for voice state events.
type RPCVoiceStateEventData struct {
	// VoiceState is the voice state.
	VoiceState interface{} `json:"voice_state"` // APIVoiceState

	// User is the user.
	User interface{} `json:"user"` // APIUser

	// Nick is the user's nickname.
	Nick *string `json:"nick,omitempty"`

	// Volume is the user's volume.
	Volume float64 `json:"volume"`

	// Mute indicates if the user is muted.
	Mute bool `json:"mute"`

	// Pan is the user's pan.
	Pan *RPCVoicePan `json:"pan,omitempty"`
}

// RPCVoicePan represents voice panning.
type RPCVoicePan struct {
	// Left is the left pan value.
	Left float64 `json:"left"`

	// Right is the right pan value.
	Right float64 `json:"right"`
}

// RPCReadyEventData represents data for ready events.
type RPCReadyEventData struct {
	// V is the RPC version.
	V int `json:"v"`

	// Config is the configuration.
	Config RPCReadyConfig `json:"config"`

	// User is the current user.
	User interface{} `json:"user"` // APIUser
}

// RPCReadyConfig represents ready configuration.
type RPCReadyConfig struct {
	// CDNHost is the CDN host.
	CDNHost string `json:"cdn_host"`

	// APIEndpoint is the API endpoint.
	APIEndpoint string `json:"api_endpoint"`

	// Environment is the environment.
	Environment string `json:"environment"`
}

// =============================================================================
// RPC PAYLOAD INTERFACES
// =============================================================================

// RPCSendPayload represents payloads that can be sent via RPC.
type RPCSendPayload interface {
	isRPCSendPayload()
}

// RPCReceivePayload represents payloads that can be received via RPC.
type RPCReceivePayload interface {
	isRPCReceivePayload()
}

// RPCCommandPayload represents an RPC command payload.
type RPCCommandPayload struct {
	// Cmd is the command.
	Cmd RPCCommands `json:"cmd"`

	// Args are the command arguments.
	Args interface{} `json:"args,omitempty"`

	// Nonce is the command nonce.
	Nonce string `json:"nonce"`
}

func (p RPCCommandPayload) isRPCSendPayload() {}

// RPCEventPayload represents an RPC event payload.
type RPCEventPayload struct {
	// Cmd is the command (for events, this is "DISPATCH").
	Cmd RPCCommands `json:"cmd"`

	// Data is the event data.
	Data interface{} `json:"data"`

	// Evt is the event type.
	Evt RPCEvents `json:"evt"`

	// Nonce is the event nonce.
	Nonce *string `json:"nonce,omitempty"`
}

func (p RPCEventPayload) isRPCReceivePayload() {}

// RPCResponsePayload represents an RPC response payload.
type RPCResponsePayload struct {
	// Cmd is the command.
	Cmd RPCCommands `json:"cmd"`

	// Data is the response data.
	Data interface{} `json:"data,omitempty"`

	// Evt is the event type (null for responses).
	Evt interface{} `json:"evt"`

	// Nonce is the response nonce.
	Nonce string `json:"nonce"`
}

func (p RPCResponsePayload) isRPCReceivePayload() {}

// RPCErrorPayload represents an RPC error payload.
type RPCErrorPayload struct {
	// Cmd is the command.
	Cmd RPCCommands `json:"cmd"`

	// Data is the error data.
	Data RPCErrorData `json:"data"`

	// Evt is the event type (null for errors).
	Evt interface{} `json:"evt"`

	// Nonce is the error nonce.
	Nonce string `json:"nonce"`
}

func (p RPCErrorPayload) isRPCReceivePayload() {}

// RPCErrorData represents RPC error data.
type RPCErrorData struct {
	// Code is the error code.
	Code RPCErrorCodes `json:"code"`

	// Message is the error message.
	Message string `json:"message"`
}
