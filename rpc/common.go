// Package rpc provides Discord RPC types and utilities.
//
// This package contains types for Discord's Rich Presence Client (RPC),
// which allows applications to interact with the Discord client.
package rpc

import "github.com/kolosys/discord-types/discord"

// =============================================================================
// RPC ERROR CODES
// =============================================================================

// RPCErrorCodes represents RPC error codes.
//
// See: https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-error-codes
type RPCErrorCodes int

// RPC error code constants
const (
	// RPCErrorCodeUnknownError indicates an unknown error occurred.
	RPCErrorCodeUnknownError RPCErrorCodes = 1000

	// RPCErrorCodeInvalidPayload indicates the sent payload was invalid.
	RPCErrorCodeInvalidPayload RPCErrorCodes = 4000 + iota

	// RPCErrorCodeInvalidCommand indicates invalid command name specified.
	RPCErrorCodeInvalidCommand

	// RPCErrorCodeInvalidGuild indicates invalid guild ID specified.
	RPCErrorCodeInvalidGuild

	// RPCErrorCodeInvalidEvent indicates invalid event name specified.
	RPCErrorCodeInvalidEvent

	// RPCErrorCodeInvalidChannel indicates invalid channel ID specified.
	RPCErrorCodeInvalidChannel

	// RPCErrorCodeInvalidPermissions indicates you lack permissions to access the given resource.
	RPCErrorCodeInvalidPermissions

	// RPCErrorCodeInvalidClientId indicates an invalid OAuth2 application ID was used to authorize or authenticate with.
	RPCErrorCodeInvalidClientId

	// RPCErrorCodeInvalidOrigin indicates an invalid OAuth2 application origin was used to authorize or authenticate with.
	RPCErrorCodeInvalidOrigin

	// RPCErrorCodeInvalidToken indicates an invalid OAuth2 token was used to authorize or authenticate with.
	RPCErrorCodeInvalidToken

	// RPCErrorCodeInvalidUser indicates the specified user ID was invalid.
	RPCErrorCodeInvalidUser

	// RPCErrorCodeInvalidInvite indicates invalid invite.
	RPCErrorCodeInvalidInvite

	// RPCErrorCodeInvalidActivityJoinRequest indicates invalid activity join request.
	RPCErrorCodeInvalidActivityJoinRequest

	// RPCErrorCodeInvalidEntitlement indicates invalid entitlement.
	RPCErrorCodeInvalidEntitlement

	// RPCErrorCodeInvalidGiftCode indicates invalid gift code.
	RPCErrorCodeInvalidGiftCode

	// RPCErrorCodeOAuth2Error indicates a standard OAuth2 error occurred; check the data object for the OAuth2 error details.
	RPCErrorCodeOAuth2Error RPCErrorCodes = 5000

	// RPCErrorCodeSelectChannelTimedOut indicates an asynchronous SELECT_TEXT_CHANNEL/SELECT_VOICE_CHANNEL command timed out.
	RPCErrorCodeSelectChannelTimedOut

	// RPCErrorCodeGetGuildTimedOut indicates an asynchronous GET_GUILD command timed out.
	RPCErrorCodeGetGuildTimedOut

	// RPCErrorCodeSelectVoiceForceRequired indicates you tried to join a user to a voice channel but the user was already in one.
	RPCErrorCodeSelectVoiceForceRequired

	// RPCErrorCodeCaptureShortcutAlreadyListening indicates you tried to capture more than one shortcut key at once.
	RPCErrorCodeCaptureShortcutAlreadyListening

	// RPCErrorCodeInvalidActivitySecret indicates invalid activity secret.
	RPCErrorCodeInvalidActivitySecret

	// RPCErrorCodeNoEligibleActivity indicates no eligible activity.
	RPCErrorCodeNoEligibleActivity

	// RPCErrorCodePurchaseCanceled indicates purchase was canceled.
	RPCErrorCodePurchaseCanceled

	// RPCErrorCodePurchaseError indicates purchase error.
	RPCErrorCodePurchaseError

	// RPCErrorCodeUnauthorizedForAchievement indicates unauthorized for achievement.
	RPCErrorCodeUnauthorizedForAchievement

	// RPCErrorCodeRateLimited indicates rate limited.
	RPCErrorCodeRateLimited
)

// =============================================================================
// RPC CLOSE EVENT CODES
// =============================================================================

// RPCCloseEventCodes represents RPC close event codes.
//
// See: https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-close-event-codes
type RPCCloseEventCodes int

// RPC close event code constants
const (
	// RPCCloseEventCodeCloseNormal indicates normal closure.
	RPCCloseEventCodeCloseNormal RPCCloseEventCodes = 1000

	// RPCCloseEventCodeCloseUnsupported indicates unsupported protocol.
	RPCCloseEventCodeCloseUnsupported RPCCloseEventCodes = 1003

	// RPCCloseEventCodeCloseAbnormal indicates abnormal closure.
	RPCCloseEventCodeCloseAbnormal RPCCloseEventCodes = 1006

	// RPCCloseEventCodeInvalidClientId indicates invalid client ID.
	RPCCloseEventCodeInvalidClientId RPCCloseEventCodes = 4000

	// RPCCloseEventCodeInvalidOrigin indicates invalid origin.
	RPCCloseEventCodeInvalidOrigin

	// RPCCloseEventCodeRateLimited indicates rate limited.
	RPCCloseEventCodeRateLimited

	// RPCCloseEventCodeTokenRevoked indicates token revoked.
	RPCCloseEventCodeTokenRevoked

	// RPCCloseEventCodeInvalidVersion indicates invalid version.
	RPCCloseEventCodeInvalidVersion

	// RPCCloseEventCodeInvalidEncoding indicates invalid encoding.
	RPCCloseEventCodeInvalidEncoding
)

// =============================================================================
// RELATIONSHIP TYPES
// =============================================================================

// RelationshipType represents relationship types.
type RelationshipType int

// Relationship type constants
const (
	// RelationshipTypeNone indicates no relationship.
	RelationshipTypeNone RelationshipType = iota

	// RelationshipTypeFriend indicates friend relationship.
	RelationshipTypeFriend

	// RelationshipTypeBlocked indicates blocked relationship.
	RelationshipTypeBlocked

	// RelationshipTypePendingIncoming indicates pending incoming friend request.
	RelationshipTypePendingIncoming

	// RelationshipTypePendingOutgoing indicates pending outgoing friend request.
	RelationshipTypePendingOutgoing

	// RelationshipTypeImplicit indicates implicit relationship.
	RelationshipTypeImplicit
)

// =============================================================================
// VOICE CONNECTION STATES
// =============================================================================

// VoiceConnectionStates represents voice connection states.
type VoiceConnectionStates string

// Voice connection state constants
const (
	// VoiceConnectionStateDisconnected indicates disconnected state.
	VoiceConnectionStateDisconnected VoiceConnectionStates = "DISCONNECTED"

	// VoiceConnectionStateAwaitingEndpoint indicates awaiting endpoint state.
	VoiceConnectionStateAwaitingEndpoint VoiceConnectionStates = "AWAITING_ENDPOINT"

	// VoiceConnectionStateAuthenticating indicates authenticating state.
	VoiceConnectionStateAuthenticating VoiceConnectionStates = "AUTHENTICATING"

	// VoiceConnectionStateConnecting indicates connecting state.
	VoiceConnectionStateConnecting VoiceConnectionStates = "CONNECTING"

	// VoiceConnectionStateConnected indicates connected state.
	VoiceConnectionStateConnected VoiceConnectionStates = "CONNECTED"

	// VoiceConnectionStateVoiceDisconnected indicates voice disconnected state.
	VoiceConnectionStateVoiceDisconnected VoiceConnectionStates = "VOICE_DISCONNECTED"

	// VoiceConnectionStateVoiceConnecting indicates voice connecting state.
	VoiceConnectionStateVoiceConnecting VoiceConnectionStates = "VOICE_CONNECTING"

	// VoiceConnectionStateVoiceConnected indicates voice connected state.
	VoiceConnectionStateVoiceConnected VoiceConnectionStates = "VOICE_CONNECTED"

	// VoiceConnectionStateNoRoute indicates no route state.
	VoiceConnectionStateNoRoute VoiceConnectionStates = "NO_ROUTE"

	// VoiceConnectionStateIceChecking indicates ICE checking state.
	VoiceConnectionStateIceChecking VoiceConnectionStates = "ICE_CHECKING"
)

// =============================================================================
// VOICE SETTINGS
// =============================================================================

// RPCVoiceSettingsMode represents voice settings mode.
type RPCVoiceSettingsMode string

// Voice settings mode constants
const (
	// RPCVoiceSettingsModeVoiceActivity indicates voice activity mode.
	RPCVoiceSettingsModeVoiceActivity RPCVoiceSettingsMode = "VOICE_ACTIVITY"

	// RPCVoiceSettingsModePushToTalk indicates push to talk mode.
	RPCVoiceSettingsModePushToTalk RPCVoiceSettingsMode = "PUSH_TO_TALK"
)

// RPCVoiceSettingsInput represents voice settings input.
type RPCVoiceSettingsInput struct {
	// DeviceID is the input device ID.
	DeviceID string `json:"device_id"`

	// Volume is the input volume (0-100).
	Volume float64 `json:"volume"`

	// AvailableDevices are the available input devices.
	AvailableDevices []RPCDevice `json:"available_devices"`
}

// RPCVoiceSettingsOutput represents voice settings output.
type RPCVoiceSettingsOutput struct {
	// DeviceID is the output device ID.
	DeviceID string `json:"device_id"`

	// Volume is the output volume (0-100).
	Volume float64 `json:"volume"`

	// AvailableDevices are the available output devices.
	AvailableDevices []RPCDevice `json:"available_devices"`
}

// =============================================================================
// DEVICE TYPES
// =============================================================================

// RPCDevice represents an RPC device.
type RPCDevice struct {
	// ID is the device ID.
	ID string `json:"id"`

	// Name is the device name.
	Name string `json:"name"`
}

// RPCDeviceVendor represents a device vendor.
type RPCDeviceVendor struct {
	// Name is the name of the vendor.
	Name string `json:"name"`

	// URL is the URL for the vendor.
	URL string `json:"url"`
}

// RPCDeviceModel represents a device model.
type RPCDeviceModel struct {
	// Name is the name of the model.
	Name string `json:"name"`

	// URL is the URL for the model.
	URL string `json:"url"`
}

// RPCCertifiedDevice represents a certified device.
type RPCCertifiedDevice struct {
	// Type is the device type.
	Type string `json:"type"`

	// ID is the device ID.
	ID string `json:"id"`

	// Vendor is the device vendor.
	Vendor RPCDeviceVendor `json:"vendor"`

	// Model is the device model.
	Model RPCDeviceModel `json:"model"`

	// Related are related devices.
	Related []string `json:"related"`

	// EchoCancellation indicates if echo cancellation is supported.
	EchoCancellation *bool `json:"echo_cancellation,omitempty"`

	// NoiseSuppression indicates if noise suppression is supported.
	NoiseSuppression *bool `json:"noise_suppression,omitempty"`

	// AutomaticGainControl indicates if automatic gain control is supported.
	AutomaticGainControl *bool `json:"automatic_gain_control,omitempty"`

	// HardwareMute indicates if hardware mute is supported.
	HardwareMute *bool `json:"hardware_mute,omitempty"`
}

// =============================================================================
// OAUTH2 APPLICATION
// =============================================================================

// RPCOAuth2Application represents an OAuth2 application.
//
// See: https://discord.com/developers/docs/topics/rpc#authenticate-oauth2-application-structure
type RPCOAuth2Application struct {
	// Description is the application description.
	Description string `json:"description"`

	// Icon is the hash of the icon.
	Icon string `json:"icon"`

	// ID is the application client id.
	ID discord.Snowflake `json:"id"`

	// RPCOrigins is the array of RPC origin urls.
	RPCOrigins []string `json:"rpc_origins"`

	// Name is the application name.
	Name string `json:"name"`
}

// =============================================================================
// VOICE CONNECTION STATUS
// =============================================================================

// RPCVoiceConnectionStatusPing represents ping data for voice connection status.
type RPCVoiceConnectionStatusPing struct {
	// Time is the time the ping was sent.
	Time int64 `json:"time"`

	// Value is the latency of the ping in milliseconds.
	Value int `json:"value"`
}

// =============================================================================
// MESSAGE PARSING TYPES
// =============================================================================

// RPCAPIMessageParsedContentOriginalMatch represents original match data.
type RPCAPIMessageParsedContentOriginalMatch struct {
	// Content is the matched content.
	Content string `json:"0"`

	// Index is the match index.
	Index int `json:"index"`
}

// RPCAPIBaseMessageParsedContentText represents base parsed text content.
type RPCAPIBaseMessageParsedContentText struct {
	// Type is the content type.
	Type string `json:"type"`

	// Content is the text content.
	Content string `json:"content"`
}

// RPCAPIMessageParsedContentText represents parsed text content.
type RPCAPIMessageParsedContentText struct {
	RPCAPIBaseMessageParsedContentText

	// OriginalMatch is the original match data.
	OriginalMatch RPCAPIMessageParsedContentOriginalMatch `json:"originalMatch"`
}

// RPCAPIMessageParsedContentMention represents parsed mention content.
type RPCAPIMessageParsedContentMention struct {
	// Type is the content type.
	Type string `json:"type"`

	// UserID is the user ID.
	UserID discord.Snowflake `json:"userId"`

	// ChannelID is the channel ID.
	ChannelID discord.Snowflake `json:"channelId"`

	// GuildID is the guild ID.
	GuildID discord.Snowflake `json:"guildId"`

	// ParsedUserID is same as UserID.
	ParsedUserID discord.Snowflake `json:"parsedUserId"`

	// Content is the base text content.
	Content RPCAPIBaseMessageParsedContentText `json:"content"`
}
