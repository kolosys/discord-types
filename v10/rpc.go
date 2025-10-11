package v10

import (
	"encoding/json"
	"fmt"
	"time"
)

// RPCVersion represents the RPC version.
const RPCVersion = "1"

// =============================================================================
// RPC COMMANDS (Optimized with type grouping)
// =============================================================================

// RPCCommands represents RPC commands with improved type safety.
//
// Reference: https://discord.com/developers/docs/topics/rpc#commands-and-events-rpc-commands
type RPCCommands string

// Authentication and Authorization Commands
const (
	RPCCommandAuthenticate RPCCommands = "AUTHENTICATE"
	RPCCommandAuthorize    RPCCommands = "AUTHORIZE"
)

// Activity Commands
const (
	RPCCommandAcceptActivityInvite     RPCCommands = "ACCEPT_ACTIVITY_INVITE"
	RPCCommandActivityInviteUser       RPCCommands = "ACTIVITY_INVITE_USER"
	RPCCommandCloseActivityJoinRequest RPCCommands = "CLOSE_ACTIVITY_JOIN_REQUEST"
	RPCCommandSendActivityJoinInvite   RPCCommands = "SEND_ACTIVITY_JOIN_INVITE"
	RPCCommandSetActivity              RPCCommands = "SET_ACTIVITY"
)

// Channel Commands
const (
	RPCCommandCreateChannelInvite     RPCCommands = "CREATE_CHANNEL_INVITE"
	RPCCommandGetChannel              RPCCommands = "GET_CHANNEL"
	RPCCommandGetChannels             RPCCommands = "GET_CHANNELS"
	RPCCommandSelectTextChannel       RPCCommands = "SELECT_TEXT_CHANNEL"
	RPCCommandSelectVoiceChannel      RPCCommands = "SELECT_VOICE_CHANNEL"
	RPCCommandGetSelectedVoiceChannel RPCCommands = "GET_SELECTED_VOICE_CHANNEL"
)

// Voice Commands
const (
	RPCCommandGetVoiceSettings      RPCCommands = "GET_VOICE_SETTINGS"
	RPCCommandSetVoiceSettings      RPCCommands = "SET_VOICE_SETTINGS"
	RPCCommandSetVoiceSettings2     RPCCommands = "SET_VOICE_SETTINGS_2"
	RPCCommandSetUserVoiceSettings  RPCCommands = "SET_USER_VOICE_SETTINGS"
	RPCCommandSetUserVoiceSettings2 RPCCommands = "SET_USER_VOICE_SETTINGS_2"
	RPCCommandSetCertifiedDevices   RPCCommands = "SET_CERTIFIED_DEVICES"
)

// Guild Commands
const (
	RPCCommandGetGuild  RPCCommands = "GET_GUILD"
	RPCCommandGetGuilds RPCCommands = "GET_GUILDS"
)

// User Commands
const (
	RPCCommandGetUser             RPCCommands = "GET_USER"
	RPCCommandGetUserAchievements RPCCommands = "GET_USER_ACHIEVEMENTS"
	RPCCommandGetRelationships    RPCCommands = "GET_RELATIONSHIPS"
)

// Monetization Commands
const (
	RPCCommandGetEntitlements      RPCCommands = "GET_ENTITLEMENTS"
	RPCCommandGetEntitlementTicket RPCCommands = "GET_ENTITLEMENT_TICKET"
	RPCCommandGetSkus              RPCCommands = "GET_SKUS"
	RPCCommandStartPurchase        RPCCommands = "START_PURCHASE"
)

// System Commands
const (
	RPCCommandBraintreePopupBridgeCallback RPCCommands = "BRAINTREE_POPUP_BRIDGE_CALLBACK"
	RPCCommandBrowserHandoff               RPCCommands = "BROWSER_HANDOFF"
	RPCCommandCaptureShortcut              RPCCommands = "CAPTURE_SHORTCUT"
	RPCCommandConnectionsCallback          RPCCommands = "CONNECTIONS_CALLBACK"
	RPCCommandDeepLink                     RPCCommands = "DEEP_LINK"
	RPCCommandDispatch                     RPCCommands = "DISPATCH"
	RPCCommandGetApplicationTicket         RPCCommands = "GET_APPLICATION_TICKET"
	RPCCommandGetImage                     RPCCommands = "GET_IMAGE"
	RPCCommandGetPlatformBehaviors         RPCCommands = "GET_PLATFORM_BEHAVIORS"
	RPCCommandGetThermalState              RPCCommands = "GET_THERMAL_STATE"
	RPCCommandInitiateImageUpload          RPCCommands = "INITIATE_IMAGE_UPLOAD"
	RPCCommandOpenExternalLink             RPCCommands = "OPEN_EXTERNAL_LINK"
	RPCCommandOpenInviteDialog             RPCCommands = "OPEN_INVITE_DIALOG"
	RPCCommandOpenOverlayActivityInvite    RPCCommands = "OPEN_OVERLAY_ACTIVITY_INVITE"
	RPCCommandOpenOverlayGuildInvite       RPCCommands = "OPEN_OVERLAY_GUILD_INVITE"
	RPCCommandOpenOverlayVoiceSettings     RPCCommands = "OPEN_OVERLAY_VOICE_SETTINGS"
	RPCCommandOverlay                      RPCCommands = "OVERLAY"
	RPCCommandSetOverlayLocked             RPCCommands = "SET_OVERLAY_LOCKED"
	RPCCommandSubscribe                    RPCCommands = "SUBSCRIBE"
	RPCCommandUnsubscribe                  RPCCommands = "UNSUBSCRIBE"
	RPCCommandUserSettingsGetLocale        RPCCommands = "USER_SETTINGS_GET_LOCALE"
)

// =============================================================================
// RPC EVENTS (Optimized with type grouping)
// =============================================================================

// RPCEvents represents RPC events with improved categorization.
// Reference: https://discord.com/developers/docs/topics/rpc#commands-and-events-rpc-events
type RPCEvents string

// Core Events
const (
	RPCEventReady RPCEvents = "READY"
	RPCEventError RPCEvents = "ERROR"
)

// Guild Events
const (
	RPCEventGuildStatus RPCEvents = "GUILD_STATUS"
	RPCEventGuildCreate RPCEvents = "GUILD_CREATE"
)

// Channel Events
const (
	RPCEventChannelCreate RPCEvents = "CHANNEL_CREATE"
)

// Relationship Events
const (
	RPCEventRelationshipUpdate RPCEvents = "RELATIONSHIP_UPDATE"
)

// Voice Events
const (
	RPCEventVoiceChannelSelect    RPCEvents = "VOICE_CHANNEL_SELECT"
	RPCEventVoiceStateCreate      RPCEvents = "VOICE_STATE_CREATE"
	RPCEventVoiceStateUpdate      RPCEvents = "VOICE_STATE_UPDATE"
	RPCEventVoiceStateDelete      RPCEvents = "VOICE_STATE_DELETE"
	RPCEventVoiceSettingsUpdate   RPCEvents = "VOICE_SETTINGS_UPDATE"
	RPCEventVoiceSettingsUpdate2  RPCEvents = "VOICE_SETTINGS_UPDATE_2"
	RPCEventVoiceConnectionStatus RPCEvents = "VOICE_CONNECTION_STATUS"
	RPCEventSpeakingStart         RPCEvents = "SPEAKING_START"
	RPCEventSpeakingStop          RPCEvents = "SPEAKING_STOP"
)

// Activity Events
const (
	RPCEventActivityJoin        RPCEvents = "ACTIVITY_JOIN"
	RPCEventActivityJoinRequest RPCEvents = "ACTIVITY_JOIN_REQUEST"
	RPCEventActivitySpectate    RPCEvents = "ACTIVITY_SPECTATE"
	RPCEventActivityInvite      RPCEvents = "ACTIVITY_INVITE"
	// Legacy game events (deprecated but maintained for compatibility)
	RPCEventGameJoin     RPCEvents = "GAME_JOIN"
	RPCEventGameSpectate RPCEvents = "GAME_SPECTATE"
)

// Message Events
const (
	RPCEventMessageCreate RPCEvents = "MESSAGE_CREATE"
	RPCEventMessageUpdate RPCEvents = "MESSAGE_UPDATE"
	RPCEventMessageDelete RPCEvents = "MESSAGE_DELETE"
)

// System Events
const (
	RPCEventNotificationCreate RPCEvents = "NOTIFICATION_CREATE"
	RPCEventThermalStateUpdate RPCEvents = "THERMAL_STATE_UPDATE"
	RPCEventCurrentUserUpdate  RPCEvents = "CURRENT_USER_UPDATE"
	RPCEventEntitlementCreate  RPCEvents = "ENTITLEMENT_CREATE"
	RPCEventEntitlementDelete  RPCEvents = "ENTITLEMENT_DELETE"
)

// =============================================================================
// OPTIMIZED DATA STRUCTURES
// =============================================================================

// Relationship represents a user relationship with proper typing.
type Relationship struct {
	// ID is the user's ID.
	ID Snowflake `json:"id"`

	// Type is the relationship type.
	Type RelationshipType `json:"type"`

	// User is the user object.
	User *User `json:"user,omitempty"`
}

// =============================================================================
// OPTIMIZED RPC MESSAGE STRUCTURES
// =============================================================================

// RPCAPIMessage represents an RPC message with optimized memory layout and type safety.
type RPCAPIMessage struct {
	// Core identifiers (aligned for memory efficiency)
	ID        Snowflake  `json:"id"`
	ChannelID Snowflake  `json:"channel_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`

	// Author information (strongly typed)
	Author *User        `json:"author,omitempty"`
	Member *GuildMember `json:"member,omitempty"`

	// Content fields with proper time handling
	Content         string     `json:"content"`
	Timestamp       time.Time  `json:"timestamp"`
	EditedTimestamp *time.Time `json:"edited_timestamp,omitempty"`

	// Boolean flags (grouped for memory efficiency)
	TTS             bool `json:"tts"`
	MentionEveryone bool `json:"mention_everyone"`
	Pinned          bool `json:"pinned"`

	// Strongly typed arrays
	Mentions        []*User               `json:"mentions,omitempty"`
	MentionRoles    []Snowflake           `json:"mention_roles,omitempty"`
	MentionChannels []*ChannelMention     `json:"mention_channels,omitempty"`
	Attachments     []*Attachment         `json:"attachments,omitempty"`
	Embeds          []*Embed              `json:"embeds,omitempty"`
	Reactions       []*Reaction           `json:"reactions,omitempty"`
	Components      []*ActionRowComponent `json:"components,omitempty"`
	StickerItems    []*StickerItem        `json:"sticker_items,omitempty"`
	Stickers        []*Sticker            `json:"stickers,omitempty"`

	// Numeric fields
	Type        MessageType   `json:"type"`
	AuthorColor *int          `json:"author_color,omitempty"`
	Position    *int          `json:"position,omitempty"`
	Flags       *MessageFlags `json:"flags,omitempty"`

	// Optional fields
	WebhookID     *Snowflake `json:"webhook_id,omitempty"`
	ApplicationID *Snowflake `json:"application_id,omitempty"`
	Nick          *string    `json:"nick,omitempty"`

	// Flexible nonce field
	Nonce any `json:"nonce,omitempty"`

	// Parsed content with proper typing
	ContentParsed []RPCMessageParsedContent `json:"content_parsed,omitempty"`

	// Complex objects using json.RawMessage for performance
	Activity             json.RawMessage `json:"activity,omitempty"`
	Application          json.RawMessage `json:"application,omitempty"`
	MessageReference     json.RawMessage `json:"message_reference,omitempty"`
	ReferencedMessage    json.RawMessage `json:"referenced_message,omitempty"`
	Interaction          json.RawMessage `json:"interaction,omitempty"`
	Thread               json.RawMessage `json:"thread,omitempty"`
	RoleSubscriptionData json.RawMessage `json:"role_subscription_data,omitempty"`
	Resolved             json.RawMessage `json:"resolved,omitempty"`
	Poll                 json.RawMessage `json:"poll,omitempty"`
	Call                 json.RawMessage `json:"call,omitempty"`
}

// =============================================================================
// ENHANCED RPC DATA TYPES
// =============================================================================

// RPCGuildCreateEventData represents guild create event data with validation.
type RPCGuildCreateEventData struct {
	ID   Snowflake `json:"id"`
	Name string    `json:"name"`
}

// Validate checks if the guild create event data is valid.
func (g *RPCGuildCreateEventData) Validate() error {
	if g.ID == "" {
		return fmt.Errorf("guild ID cannot be empty")
	}
	if g.Name == "" {
		return fmt.Errorf("guild name cannot be empty")
	}
	return nil
}

// RPCChannelCreateEventData represents channel create event data with validation.
type RPCChannelCreateEventData struct {
	ID   Snowflake   `json:"id"`
	Name string      `json:"name"`
	Type ChannelType `json:"type"`
}

// Validate checks if the channel create event data is valid.
func (c *RPCChannelCreateEventData) Validate() error {
	if c.ID == "" {
		return fmt.Errorf("channel ID cannot be empty")
	}
	if c.Name == "" {
		return fmt.Errorf("channel name cannot be empty")
	}
	return nil
}

// RPCVoiceStateEventData represents voice state event data with optimized structure.
type RPCVoiceStateEventData struct {
	// Embedded voice state for direct access
	VoiceState *VoiceState `json:"voice_state,omitempty"`

	// User information
	User *User   `json:"user,omitempty"`
	Nick *string `json:"nick,omitempty"`

	// Audio settings (grouped for memory efficiency)
	Volume float64      `json:"volume"`
	Mute   bool         `json:"mute"`
	Pan    *RPCVoicePan `json:"pan,omitempty"`
}

// RPCVoicePan represents voice panning with validation.
type RPCVoicePan struct {
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

// Validate ensures pan values are within valid range [-1.0, 1.0].
func (p *RPCVoicePan) Validate() error {
	if p.Left < -1.0 || p.Left > 1.0 {
		return fmt.Errorf("left pan value must be between -1.0 and 1.0, got %f", p.Left)
	}
	if p.Right < -1.0 || p.Right > 1.0 {
		return fmt.Errorf("right pan value must be between -1.0 and 1.0, got %f", p.Right)
	}
	return nil
}

// RPCReadyEventData represents ready event data with enhanced structure.
type RPCReadyEventData struct {
	V      int            `json:"v"`
	Config RPCReadyConfig `json:"config"`
	User   *User          `json:"user,omitempty"`
}

// RPCReadyConfig represents ready configuration with validation.
type RPCReadyConfig struct {
	CDNHost     string `json:"cdn_host"`
	APIEndpoint string `json:"api_endpoint"`
	Environment string `json:"environment"`
}

// Validate checks if the ready configuration is valid.
func (c *RPCReadyConfig) Validate() error {
	if c.CDNHost == "" {
		return fmt.Errorf("CDN host cannot be empty")
	}
	if c.APIEndpoint == "" {
		return fmt.Errorf("API endpoint cannot be empty")
	}
	if c.Environment == "" {
		return fmt.Errorf("environment cannot be empty")
	}
	return nil
}

// =============================================================================
// ENHANCED PAYLOAD INTERFACES
// =============================================================================

// RPCPayload is the base interface for all RPC payloads.
type RPCPayload interface {
	// GetCommand returns the RPC command for this payload.
	GetCommand() RPCCommands

	// Validate ensures the payload is well-formed.
	Validate() error
}

// RPCSendPayload represents payloads that can be sent via RPC.
type RPCSendPayload interface {
	RPCPayload
	isRPCSendPayload()
}

// RPCReceivePayload represents payloads that can be received via RPC.
type RPCReceivePayload interface {
	RPCPayload
	isRPCReceivePayload()
}

// =============================================================================
// OPTIMIZED PAYLOAD IMPLEMENTATIONS
// =============================================================================

// RPCCommandPayload represents an RPC command payload with validation.
type RPCCommandPayload struct {
	Cmd   RPCCommands `json:"cmd"`
	Args  any         `json:"args,omitempty"`
	Nonce string      `json:"nonce"`
}

// GetCommand implements RPCPayload.
func (p *RPCCommandPayload) GetCommand() RPCCommands {
	return p.Cmd
}

// Validate implements RPCPayload.
func (p *RPCCommandPayload) Validate() error {
	if p.Cmd == "" {
		return fmt.Errorf("command cannot be empty")
	}
	if p.Nonce == "" {
		return fmt.Errorf("nonce cannot be empty")
	}
	return nil
}

func (p *RPCCommandPayload) isRPCSendPayload() {}

// RPCEventPayload represents an RPC event payload with validation.
type RPCEventPayload struct {
	Cmd   RPCCommands `json:"cmd"`
	Data  any         `json:"data,omitempty"`
	Evt   RPCEvents   `json:"evt"`
	Nonce *string     `json:"nonce,omitempty"`
}

// GetCommand implements RPCPayload.
func (p *RPCEventPayload) GetCommand() RPCCommands {
	return p.Cmd
}

// Validate implements RPCPayload.
func (p *RPCEventPayload) Validate() error {
	if p.Cmd == "" {
		return fmt.Errorf("command cannot be empty")
	}
	if p.Evt == "" {
		return fmt.Errorf("event type cannot be empty")
	}
	return nil
}

func (p *RPCEventPayload) isRPCReceivePayload() {}

// RPCResponsePayload represents an RPC response payload with validation.
type RPCResponsePayload struct {
	Cmd   RPCCommands `json:"cmd"`
	Data  any         `json:"data,omitempty"`
	Evt   any         `json:"evt"`
	Nonce string      `json:"nonce"`
}

// GetCommand implements RPCPayload.
func (p *RPCResponsePayload) GetCommand() RPCCommands {
	return p.Cmd
}

// Validate implements RPCPayload.
func (p *RPCResponsePayload) Validate() error {
	if p.Cmd == "" {
		return fmt.Errorf("command cannot be empty")
	}
	if p.Nonce == "" {
		return fmt.Errorf("nonce cannot be empty")
	}
	return nil
}

func (p *RPCResponsePayload) isRPCReceivePayload() {}

// RPCErrorPayload represents an RPC error payload with enhanced error handling.
type RPCErrorPayload struct {
	Cmd   RPCCommands  `json:"cmd"`
	Data  RPCErrorData `json:"data"`
	Evt   any          `json:"evt"`
	Nonce string       `json:"nonce"`
}

// GetCommand implements RPCPayload.
func (p *RPCErrorPayload) GetCommand() RPCCommands {
	return p.Cmd
}

// Validate implements RPCPayload.
func (p *RPCErrorPayload) Validate() error {
	if p.Cmd == "" {
		return fmt.Errorf("command cannot be empty")
	}
	if p.Nonce == "" {
		return fmt.Errorf("nonce cannot be empty")
	}
	return p.Data.Validate()
}

// Error implements the error interface for better error handling.
func (p *RPCErrorPayload) Error() string {
	return fmt.Sprintf("RPC Error %d: %s", p.Data.Code, p.Data.Message)
}

func (p *RPCErrorPayload) isRPCReceivePayload() {}

// RPCErrorData represents RPC error data with validation.
type RPCErrorData struct {
	Code    RPCErrorCodes `json:"code"`
	Message string        `json:"message"`
}

// Validate checks if the error data is well-formed.
func (e *RPCErrorData) Validate() error {
	if e.Code == 0 {
		return fmt.Errorf("error code cannot be zero")
	}
	if e.Message == "" {
		return fmt.Errorf("error message cannot be empty")
	}
	return nil
}

// Error implements the error interface.
func (e *RPCErrorData) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// =============================================================================
// OPTIMIZED ERROR CODES
// =============================================================================

// RPCErrorCodes represents RPC error codes with enhanced categorization.
// Reference: https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-error-codes
type RPCErrorCodes int

// Unknown Errors
const (
	RPCErrorCodeUnknownError RPCErrorCodes = 1000
)

// Client Errors (4000-4999)
const (
	RPCErrorCodeInvalidPayload RPCErrorCodes = 4000 + iota
	RPCErrorCodeInvalidCommand
	RPCErrorCodeInvalidGuild
	RPCErrorCodeInvalidEvent
	RPCErrorCodeInvalidChannel
	RPCErrorCodeInvalidPermissions
	RPCErrorCodeInvalidClientId
	RPCErrorCodeInvalidOrigin
	RPCErrorCodeInvalidToken
	RPCErrorCodeInvalidUser
	RPCErrorCodeInvalidInvite
	RPCErrorCodeInvalidActivityJoinRequest
	RPCErrorCodeInvalidEntitlement
	RPCErrorCodeInvalidGiftCode
)

// OAuth2 and Timeout Errors
const (
	RPCErrorCodeOAuth2Error                     RPCErrorCodes = 5000
	RPCErrorCodeSelectChannelTimedOut           RPCErrorCodes = 5001
	RPCErrorCodeGetGuildTimedOut                RPCErrorCodes = 5002
	RPCErrorCodeSelectVoiceForceRequired        RPCErrorCodes = 5003
	RPCErrorCodeCaptureShortcutAlreadyListening RPCErrorCodes = 5004
	RPCErrorCodeInvalidActivitySecret           RPCErrorCodes = 5005
	RPCErrorCodeNoEligibleActivity              RPCErrorCodes = 5006
	RPCErrorCodePurchaseCanceled                RPCErrorCodes = 5007
	RPCErrorCodePurchaseError                   RPCErrorCodes = 5008
	RPCErrorCodeUnauthorizedForAchievement      RPCErrorCodes = 5009
	RPCErrorCodeRateLimited                     RPCErrorCodes = 5010
)

// IsClientError returns true if the error code represents a client error (4000-4999).
func (c RPCErrorCodes) IsClientError() bool {
	return c >= 4000 && c < 5000
}

// IsServerError returns true if the error code represents a server error (5000+).
func (c RPCErrorCodes) IsServerError() bool {
	return c >= 5000
}

// String provides human-readable error descriptions.
func (c RPCErrorCodes) String() string {
	switch c {
	case RPCErrorCodeUnknownError:
		return "Unknown Error"
	case RPCErrorCodeInvalidPayload:
		return "Invalid Payload"
	case RPCErrorCodeInvalidCommand:
		return "Invalid Command"
	case RPCErrorCodeInvalidGuild:
		return "Invalid Guild"
	case RPCErrorCodeInvalidEvent:
		return "Invalid Event"
	case RPCErrorCodeInvalidChannel:
		return "Invalid Channel"
	case RPCErrorCodeInvalidPermissions:
		return "Invalid Permissions"
	case RPCErrorCodeInvalidClientId:
		return "Invalid Client ID"
	case RPCErrorCodeInvalidOrigin:
		return "Invalid Origin"
	case RPCErrorCodeInvalidToken:
		return "Invalid Token"
	case RPCErrorCodeInvalidUser:
		return "Invalid User"
	case RPCErrorCodeInvalidInvite:
		return "Invalid Invite"
	case RPCErrorCodeInvalidActivityJoinRequest:
		return "Invalid Activity Join Request"
	case RPCErrorCodeInvalidEntitlement:
		return "Invalid Entitlement"
	case RPCErrorCodeInvalidGiftCode:
		return "Invalid Gift Code"
	case RPCErrorCodeOAuth2Error:
		return "OAuth2 Error"
	case RPCErrorCodeSelectChannelTimedOut:
		return "Select Channel Timed Out"
	case RPCErrorCodeGetGuildTimedOut:
		return "Get Guild Timed Out"
	case RPCErrorCodeSelectVoiceForceRequired:
		return "Select Voice Force Required"
	case RPCErrorCodeCaptureShortcutAlreadyListening:
		return "Capture Shortcut Already Listening"
	case RPCErrorCodeInvalidActivitySecret:
		return "Invalid Activity Secret"
	case RPCErrorCodeNoEligibleActivity:
		return "No Eligible Activity"
	case RPCErrorCodePurchaseCanceled:
		return "Purchase Canceled"
	case RPCErrorCodePurchaseError:
		return "Purchase Error"
	case RPCErrorCodeUnauthorizedForAchievement:
		return "Unauthorized For Achievement"
	case RPCErrorCodeRateLimited:
		return "Rate Limited"
	default:
		return fmt.Sprintf("Unknown Error Code: %d", c)
	}
}

// =============================================================================
// RPC CLOSE EVENT CODES (Optimized)
// =============================================================================

// RPCCloseEventCodes represents RPC close event codes with enhanced descriptions.
// Reference: https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-close-event-codes
type RPCCloseEventCodes int

// Standard WebSocket Close Codes
const (
	RPCCloseEventCodeCloseNormal      RPCCloseEventCodes = 1000
	RPCCloseEventCodeCloseUnsupported RPCCloseEventCodes = 1003
	RPCCloseEventCodeCloseAbnormal    RPCCloseEventCodes = 1006
)

// Discord-specific Close Codes
const (
	RPCCloseEventCodeInvalidClientId RPCCloseEventCodes = 4000
	RPCCloseEventCodeInvalidOrigin   RPCCloseEventCodes = 4001
	RPCCloseEventCodeRateLimited     RPCCloseEventCodes = 4002
	RPCCloseEventCodeTokenRevoked    RPCCloseEventCodes = 4003
	RPCCloseEventCodeInvalidVersion  RPCCloseEventCodes = 4004
	RPCCloseEventCodeInvalidEncoding RPCCloseEventCodes = 4005
)

// String provides human-readable close event descriptions.
func (c RPCCloseEventCodes) String() string {
	switch c {
	case RPCCloseEventCodeCloseNormal:
		return "Normal Closure"
	case RPCCloseEventCodeCloseUnsupported:
		return "Unsupported Protocol"
	case RPCCloseEventCodeCloseAbnormal:
		return "Abnormal Closure"
	case RPCCloseEventCodeInvalidClientId:
		return "Invalid Client ID"
	case RPCCloseEventCodeInvalidOrigin:
		return "Invalid Origin"
	case RPCCloseEventCodeRateLimited:
		return "Rate Limited"
	case RPCCloseEventCodeTokenRevoked:
		return "Token Revoked"
	case RPCCloseEventCodeInvalidVersion:
		return "Invalid Version"
	case RPCCloseEventCodeInvalidEncoding:
		return "Invalid Encoding"
	default:
		return fmt.Sprintf("Unknown Close Code: %d", c)
	}
}

// =============================================================================
// ENHANCED RELATIONSHIP TYPES
// =============================================================================

// RelationshipType represents relationship types with improved methods.
type RelationshipType int

const (
	RelationshipTypeNone RelationshipType = iota
	RelationshipTypeFriend
	RelationshipTypeBlocked
	RelationshipTypePendingIncoming
	RelationshipTypePendingOutgoing
	RelationshipTypeImplicit
)

// String provides human-readable relationship type descriptions.
func (r RelationshipType) String() string {
	switch r {
	case RelationshipTypeNone:
		return "None"
	case RelationshipTypeFriend:
		return "Friend"
	case RelationshipTypeBlocked:
		return "Blocked"
	case RelationshipTypePendingIncoming:
		return "Pending Incoming"
	case RelationshipTypePendingOutgoing:
		return "Pending Outgoing"
	case RelationshipTypeImplicit:
		return "Implicit"
	default:
		return fmt.Sprintf("Unknown Relationship Type: %d", r)
	}
}

// IsPending returns true if the relationship is pending (incoming or outgoing).
func (r RelationshipType) IsPending() bool {
	return r == RelationshipTypePendingIncoming || r == RelationshipTypePendingOutgoing
}

// =============================================================================
// ENHANCED VOICE CONNECTION STATES
// =============================================================================

// VoiceConnectionStates represents voice connection states with validation.
type VoiceConnectionStates string

const (
	VoiceConnectionStateDisconnected      VoiceConnectionStates = "DISCONNECTED"
	VoiceConnectionStateAwaitingEndpoint  VoiceConnectionStates = "AWAITING_ENDPOINT"
	VoiceConnectionStateAuthenticating    VoiceConnectionStates = "AUTHENTICATING"
	VoiceConnectionStateConnecting        VoiceConnectionStates = "CONNECTING"
	VoiceConnectionStateConnected         VoiceConnectionStates = "CONNECTED"
	VoiceConnectionStateVoiceDisconnected VoiceConnectionStates = "VOICE_DISCONNECTED"
	VoiceConnectionStateVoiceConnecting   VoiceConnectionStates = "VOICE_CONNECTING"
	VoiceConnectionStateVoiceConnected    VoiceConnectionStates = "VOICE_CONNECTED"
	VoiceConnectionStateNoRoute           VoiceConnectionStates = "NO_ROUTE"
	VoiceConnectionStateIceChecking       VoiceConnectionStates = "ICE_CHECKING"
)

// IsConnected returns true if the connection state indicates an active connection.
func (s VoiceConnectionStates) IsConnected() bool {
	return s == VoiceConnectionStateConnected || s == VoiceConnectionStateVoiceConnected
}

// IsDisconnected returns true if the connection state indicates disconnection.
func (s VoiceConnectionStates) IsDisconnected() bool {
	return s == VoiceConnectionStateDisconnected || s == VoiceConnectionStateVoiceDisconnected
}

// =============================================================================
// ENHANCED VOICE SETTINGS
// =============================================================================

// RPCVoiceSettingsMode represents voice settings mode with validation.
type RPCVoiceSettingsMode string

const (
	RPCVoiceSettingsModeVoiceActivity RPCVoiceSettingsMode = "VOICE_ACTIVITY"
	RPCVoiceSettingsModePushToTalk    RPCVoiceSettingsMode = "PUSH_TO_TALK"
)

// Validate ensures the voice settings mode is valid.
func (m RPCVoiceSettingsMode) Validate() error {
	switch m {
	case RPCVoiceSettingsModeVoiceActivity, RPCVoiceSettingsModePushToTalk:
		return nil
	default:
		return fmt.Errorf("invalid voice settings mode: %s", m)
	}
}

// RPCVoiceSettingsInput represents voice settings input with validation.
type RPCVoiceSettingsInput struct {
	DeviceID         string      `json:"device_id"`
	Volume           float64     `json:"volume"`
	AvailableDevices []RPCDevice `json:"available_devices"`
}

// Validate ensures the input settings are valid.
func (i *RPCVoiceSettingsInput) Validate() error {
	if i.Volume < 0 || i.Volume > 100 {
		return fmt.Errorf("volume must be between 0 and 100, got %f", i.Volume)
	}
	return nil
}

// RPCVoiceSettingsOutput represents voice settings output with validation.
type RPCVoiceSettingsOutput struct {
	DeviceID         string      `json:"device_id"`
	Volume           float64     `json:"volume"`
	AvailableDevices []RPCDevice `json:"available_devices"`
}

// Validate ensures the output settings are valid.
func (o *RPCVoiceSettingsOutput) Validate() error {
	if o.Volume < 0 || o.Volume > 100 {
		return fmt.Errorf("volume must be between 0 and 100, got %f", o.Volume)
	}
	return nil
}

// =============================================================================
// ENHANCED DEVICE TYPES
// =============================================================================

// RPCDevice represents an RPC device with validation.
type RPCDevice struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Validate ensures the device information is valid.
func (d *RPCDevice) Validate() error {
	if d.ID == "" {
		return fmt.Errorf("device ID cannot be empty")
	}
	if d.Name == "" {
		return fmt.Errorf("device name cannot be empty")
	}
	return nil
}

// RPCDeviceVendor represents a device vendor with validation.
type RPCDeviceVendor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Validate ensures the vendor information is valid.
func (v *RPCDeviceVendor) Validate() error {
	if v.Name == "" {
		return fmt.Errorf("vendor name cannot be empty")
	}
	if v.URL == "" {
		return fmt.Errorf("vendor URL cannot be empty")
	}
	return nil
}

// RPCDeviceModel represents a device model with validation.
type RPCDeviceModel struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Validate ensures the model information is valid.
func (m *RPCDeviceModel) Validate() error {
	if m.Name == "" {
		return fmt.Errorf("model name cannot be empty")
	}
	if m.URL == "" {
		return fmt.Errorf("model URL cannot be empty")
	}
	return nil
}

// RPCCertifiedDevice represents a certified device with enhanced features.
type RPCCertifiedDevice struct {
	Type                 string          `json:"type"`
	ID                   string          `json:"id"`
	Vendor               RPCDeviceVendor `json:"vendor"`
	Model                RPCDeviceModel  `json:"model"`
	Related              []string        `json:"related"`
	EchoCancellation     *bool           `json:"echo_cancellation,omitempty"`
	NoiseSuppression     *bool           `json:"noise_suppression,omitempty"`
	AutomaticGainControl *bool           `json:"automatic_gain_control,omitempty"`
	HardwareMute         *bool           `json:"hardware_mute,omitempty"`
}

// Validate ensures the certified device information is valid.
func (d *RPCCertifiedDevice) Validate() error {
	if d.Type == "" {
		return fmt.Errorf("device type cannot be empty")
	}
	if d.ID == "" {
		return fmt.Errorf("device ID cannot be empty")
	}
	if err := d.Vendor.Validate(); err != nil {
		return fmt.Errorf("vendor validation failed: %w", err)
	}
	if err := d.Model.Validate(); err != nil {
		return fmt.Errorf("model validation failed: %w", err)
	}
	return nil
}

// HasAdvancedFeatures returns true if the device supports advanced audio features.
func (d *RPCCertifiedDevice) HasAdvancedFeatures() bool {
	return d.EchoCancellation != nil && *d.EchoCancellation ||
		d.NoiseSuppression != nil && *d.NoiseSuppression ||
		d.AutomaticGainControl != nil && *d.AutomaticGainControl
}

// =============================================================================
// ENHANCED OAUTH2 APPLICATION
// =============================================================================

// RPCOAuth2Application represents an OAuth2 application with validation.
// Reference: https://discord.com/developers/docs/topics/rpc#authenticate-oauth2-application-structure
type RPCOAuth2Application struct {
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	ID          Snowflake `json:"id"`
	RPCOrigins  []string  `json:"rpc_origins"`
	Name        string    `json:"name"`
}

// Validate ensures the OAuth2 application data is valid.
func (a *RPCOAuth2Application) Validate() error {
	if a.ID == "" {
		return fmt.Errorf("application ID cannot be empty")
	}
	if a.Name == "" {
		return fmt.Errorf("application name cannot be empty")
	}
	if len(a.RPCOrigins) == 0 {
		return fmt.Errorf("at least one RPC origin must be specified")
	}
	return nil
}

// =============================================================================
// ENHANCED VOICE CONNECTION STATUS
// =============================================================================

// RPCVoiceConnectionStatusPing represents ping data with validation.
type RPCVoiceConnectionStatusPing struct {
	Time  int64 `json:"time"`
	Value int   `json:"value"`
}

// Validate ensures the ping data is valid.
func (p *RPCVoiceConnectionStatusPing) Validate() error {
	if p.Time <= 0 {
		return fmt.Errorf("ping time must be positive, got %d", p.Time)
	}
	if p.Value < 0 {
		return fmt.Errorf("ping value cannot be negative, got %d", p.Value)
	}
	return nil
}

// =============================================================================
// ENHANCED MESSAGE PARSING TYPES
// =============================================================================

// RPCMessageParsedContent represents a union type for parsed message content.
type RPCMessageParsedContent interface {
	GetType() string
}

// RPCAPIMessageParsedContentOriginalMatch represents original match data.
type RPCAPIMessageParsedContentOriginalMatch struct {
	Content string `json:"0"`
	Index   int    `json:"index"`
}

// RPCAPIBaseMessageParsedContentText represents base parsed text content.
type RPCAPIBaseMessageParsedContentText struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// GetType implements RPCMessageParsedContent.
func (t *RPCAPIBaseMessageParsedContentText) GetType() string {
	return t.Type
}

// RPCAPIMessageParsedContentText represents parsed text content.
type RPCAPIMessageParsedContentText struct {
	RPCAPIBaseMessageParsedContentText
	OriginalMatch RPCAPIMessageParsedContentOriginalMatch `json:"originalMatch"`
}

// RPCAPIMessageParsedContentMention represents parsed mention content.
type RPCAPIMessageParsedContentMention struct {
	Type         string                             `json:"type"`
	UserID       Snowflake                          `json:"userId"`
	ChannelID    Snowflake                          `json:"channelId"`
	GuildID      Snowflake                          `json:"guildId"`
	ParsedUserID Snowflake                          `json:"parsedUserId"`
	Content      RPCAPIBaseMessageParsedContentText `json:"content"`
}

// GetType implements RPCMessageParsedContent.
func (m *RPCAPIMessageParsedContentMention) GetType() string {
	return m.Type
}

// =============================================================================
// HELPER FUNCTIONS
// =============================================================================

// NewRPCCommandPayload creates a new RPC command payload with validation.
func NewRPCCommandPayload(cmd RPCCommands, args any, nonce string) (*RPCCommandPayload, error) {
	payload := &RPCCommandPayload{
		Cmd:   cmd,
		Args:  args,
		Nonce: nonce,
	}

	if err := payload.Validate(); err != nil {
		return nil, fmt.Errorf("invalid command payload: %w", err)
	}

	return payload, nil
}

// NewRPCEventPayload creates a new RPC event payload with validation.
func NewRPCEventPayload(cmd RPCCommands, evt RPCEvents, data any, nonce *string) (*RPCEventPayload, error) {
	payload := &RPCEventPayload{
		Cmd:   cmd,
		Data:  data,
		Evt:   evt,
		Nonce: nonce,
	}

	if err := payload.Validate(); err != nil {
		return nil, fmt.Errorf("invalid event payload: %w", err)
	}

	return payload, nil
}

// NewRPCErrorPayload creates a new RPC error payload with validation.
func NewRPCErrorPayload(cmd RPCCommands, code RPCErrorCodes, message string, nonce string) (*RPCErrorPayload, error) {
	payload := &RPCErrorPayload{
		Cmd: cmd,
		Data: RPCErrorData{
			Code:    code,
			Message: message,
		},
		Evt:   nil,
		Nonce: nonce,
	}

	if err := payload.Validate(); err != nil {
		return nil, fmt.Errorf("invalid error payload: %w", err)
	}

	return payload, nil
}
