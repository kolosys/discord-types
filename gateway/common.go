// Package gateway provides Discord Gateway/WebSocket types.
//
// This package contains types for Discord Gateway events, payloads,
// and other WebSocket-related functionality.
package gateway

// This package is independent of the main discord package to avoid circular imports

// GatewayVersion represents the Gateway version.
const GatewayVersion = "10"

// GatewayURLQuery represents Gateway URL query string parameters.
//
// See: https://discord.com/developers/docs/topics/gateway#connecting-gateway-url-query-string-params
type GatewayURLQuery struct {
	// V is the gateway version.
	V string `json:"v"`

	// Encoding is the encoding type.
	Encoding GatewayEncoding `json:"encoding"`

	// Compress is the compression type.
	Compress *GatewayCompression `json:"compress,omitempty"`
}

// GatewayEncoding represents gateway encoding types.
type GatewayEncoding string

// Gateway encoding constants
const (
	// GatewayEncodingETF represents ETF encoding.
	GatewayEncodingETF GatewayEncoding = "etf"

	// GatewayEncodingJSON represents JSON encoding.
	GatewayEncodingJSON GatewayEncoding = "json"
)

// GatewayCompression represents gateway compression types.
type GatewayCompression string

// Gateway compression constants
const (
	// GatewayCompressionZlibStream represents zlib-stream compression.
	GatewayCompressionZlibStream GatewayCompression = "zlib-stream"
)

// GatewayOpcode represents Gateway opcodes.
//
// See: https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-opcodes
type GatewayOpcode int

// Gateway opcode constants
const (
	// OpcodeDispatch indicates an event was dispatched.
	OpcodeDispatch GatewayOpcode = iota

	// OpcodeHeartbeat is a bidirectional opcode to maintain an active gateway connection.
	// Fired periodically by the client, or fired by the gateway to request an immediate heartbeat from the client.
	OpcodeHeartbeat

	// OpcodeIdentify starts a new session during the initial handshake.
	OpcodeIdentify

	// OpcodePresenceUpdate updates the client's presence.
	OpcodePresenceUpdate

	// OpcodeVoiceStateUpdate is used to join/leave or move between voice channels.
	OpcodeVoiceStateUpdate

	_ // 5 is unused

	// OpcodeResume resumes a previous session that was disconnected.
	OpcodeResume

	// OpcodeReconnect indicates you should attempt to reconnect and resume immediately.
	OpcodeReconnect

	// OpcodeRequestGuildMembers requests information about offline guild members in a large guild.
	OpcodeRequestGuildMembers

	// OpcodeInvalidSession indicates the session has been invalidated.
	// You should reconnect and identify/resume accordingly.
	OpcodeInvalidSession

	// OpcodeHello is sent immediately after connecting, contains the heartbeat_interval to use.
	OpcodeHello

	// OpcodeHeartbeatAck is sent in response to receiving a heartbeat to acknowledge that it has been received.
	OpcodeHeartbeatAck
)

// RequestSoundboardSounds requests information about soundboard sounds in a set of guilds.
const OpcodeRequestSoundboardSounds GatewayOpcode = 31

// CloseCodes represents Gateway close event codes.
//
// See: https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-close-event-codes
type CloseCodes int

// Gateway close code constants
const (
	CloseCodeUnknownError CloseCodes = 4000 + iota
	CloseCodeUnknownOpcode
	CloseCodeDecodeError
	CloseCodeNotAuthenticated
	CloseCodeAuthenticationFailed
	CloseCodeAlreadyAuthenticated
	_
	CloseCodeInvalidSeq
	CloseCodeRateLimited
	CloseCodeSessionTimedOut
	CloseCodeInvalidShard
	CloseCodeShardingRequired
	CloseCodeInvalidAPIVersion
	CloseCodeInvalidIntents
	CloseCodeDisallowedIntents
)

// IntentBits represents Gateway intent bits.
//
// See: https://discord.com/developers/docs/topics/gateway#list-of-intents
type IntentBits int

// Gateway intent bit constants
const (
	IntentGuilds IntentBits = 1 << iota
	IntentGuildMembers
	IntentGuildModeration  // Previously IntentGuildBans
	IntentGuildExpressions // Previously IntentGuildEmojisAndStickers
	IntentGuildIntegrations
	IntentGuildWebhooks
	IntentGuildInvites
	IntentGuildVoiceStates
	IntentGuildPresences
	IntentGuildMessages
	IntentGuildMessageReactions
	IntentGuildMessageTyping
	IntentDirectMessages
	IntentDirectMessageReactions
	IntentDirectMessageTyping
	IntentMessageContent
	IntentGuildScheduledEvents
	_
	_
	_
	IntentAutoModerationConfiguration
	IntentAutoModerationExecution
	_
	_
	IntentGuildMessagePolls
	IntentDirectMessagePolls
)

// Deprecated intent aliases for backwards compatibility
const (
	IntentGuildBans              = IntentGuildModeration
	IntentGuildEmojisAndStickers = IntentGuildExpressions
)

// DispatchEvents represents Gateway dispatch event types.
//
// See: https://discord.com/developers/docs/topics/gateway-events#receive-events
type DispatchEvents string

// Dispatch event constants
const (
	EventApplicationCommandPermissionsUpdate DispatchEvents = "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
	EventAutoModerationActionExecution       DispatchEvents = "AUTO_MODERATION_ACTION_EXECUTION"
	EventAutoModerationRuleCreate            DispatchEvents = "AUTO_MODERATION_RULE_CREATE"
	EventAutoModerationRuleDelete            DispatchEvents = "AUTO_MODERATION_RULE_DELETE"
	EventAutoModerationRuleUpdate            DispatchEvents = "AUTO_MODERATION_RULE_UPDATE"
	EventChannelCreate                       DispatchEvents = "CHANNEL_CREATE"
	EventChannelDelete                       DispatchEvents = "CHANNEL_DELETE"
	EventChannelPinsUpdate                   DispatchEvents = "CHANNEL_PINS_UPDATE"
	EventChannelUpdate                       DispatchEvents = "CHANNEL_UPDATE"
	EventEntitlementCreate                   DispatchEvents = "ENTITLEMENT_CREATE"
	EventEntitlementDelete                   DispatchEvents = "ENTITLEMENT_DELETE"
	EventEntitlementUpdate                   DispatchEvents = "ENTITLEMENT_UPDATE"
	EventGuildAuditLogEntryCreate            DispatchEvents = "GUILD_AUDIT_LOG_ENTRY_CREATE"
	EventGuildBanAdd                         DispatchEvents = "GUILD_BAN_ADD"
	EventGuildBanRemove                      DispatchEvents = "GUILD_BAN_REMOVE"
	EventGuildCreate                         DispatchEvents = "GUILD_CREATE"
	EventGuildDelete                         DispatchEvents = "GUILD_DELETE"
	EventGuildEmojisUpdate                   DispatchEvents = "GUILD_EMOJIS_UPDATE"
	EventGuildIntegrationsUpdate             DispatchEvents = "GUILD_INTEGRATIONS_UPDATE"
	EventGuildMemberAdd                      DispatchEvents = "GUILD_MEMBER_ADD"
	EventGuildMemberRemove                   DispatchEvents = "GUILD_MEMBER_REMOVE"
	EventGuildMembersChunk                   DispatchEvents = "GUILD_MEMBERS_CHUNK"
	EventGuildMemberUpdate                   DispatchEvents = "GUILD_MEMBER_UPDATE"
	EventGuildRoleCreate                     DispatchEvents = "GUILD_ROLE_CREATE"
	EventGuildRoleDelete                     DispatchEvents = "GUILD_ROLE_DELETE"
	EventGuildRoleUpdate                     DispatchEvents = "GUILD_ROLE_UPDATE"
	EventGuildScheduledEventCreate           DispatchEvents = "GUILD_SCHEDULED_EVENT_CREATE"
	EventGuildScheduledEventDelete           DispatchEvents = "GUILD_SCHEDULED_EVENT_DELETE"
	EventGuildScheduledEventUpdate           DispatchEvents = "GUILD_SCHEDULED_EVENT_UPDATE"
	EventGuildScheduledEventUserAdd          DispatchEvents = "GUILD_SCHEDULED_EVENT_USER_ADD"
	EventGuildScheduledEventUserRemove       DispatchEvents = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
	EventGuildSoundboardSoundCreate          DispatchEvents = "GUILD_SOUNDBOARD_SOUND_CREATE"
	EventGuildSoundboardSoundDelete          DispatchEvents = "GUILD_SOUNDBOARD_SOUND_DELETE"
	EventGuildSoundboardSoundsUpdate         DispatchEvents = "GUILD_SOUNDBOARD_SOUNDS_UPDATE"
	EventGuildSoundboardSoundUpdate          DispatchEvents = "GUILD_SOUNDBOARD_SOUND_UPDATE"
	EventSoundboardSounds                    DispatchEvents = "SOUNDBOARD_SOUNDS"
	EventGuildStickersUpdate                 DispatchEvents = "GUILD_STICKERS_UPDATE"
	EventGuildUpdate                         DispatchEvents = "GUILD_UPDATE"
	EventIntegrationCreate                   DispatchEvents = "INTEGRATION_CREATE"
	EventIntegrationDelete                   DispatchEvents = "INTEGRATION_DELETE"
	EventIntegrationUpdate                   DispatchEvents = "INTEGRATION_UPDATE"
	EventInteractionCreate                   DispatchEvents = "INTERACTION_CREATE"
	EventInviteCreate                        DispatchEvents = "INVITE_CREATE"
	EventInviteDelete                        DispatchEvents = "INVITE_DELETE"
	EventMessageCreate                       DispatchEvents = "MESSAGE_CREATE"
	EventMessageDelete                       DispatchEvents = "MESSAGE_DELETE"
	EventMessageDeleteBulk                   DispatchEvents = "MESSAGE_DELETE_BULK"
	EventMessagePollVoteAdd                  DispatchEvents = "MESSAGE_POLL_VOTE_ADD"
	EventMessagePollVoteRemove               DispatchEvents = "MESSAGE_POLL_VOTE_REMOVE"
	EventMessageReactionAdd                  DispatchEvents = "MESSAGE_REACTION_ADD"
	EventMessageReactionRemove               DispatchEvents = "MESSAGE_REACTION_REMOVE"
	EventMessageReactionRemoveAll            DispatchEvents = "MESSAGE_REACTION_REMOVE_ALL"
	EventMessageReactionRemoveEmoji          DispatchEvents = "MESSAGE_REACTION_REMOVE_EMOJI"
	EventMessageUpdate                       DispatchEvents = "MESSAGE_UPDATE"
	EventPresenceUpdate                      DispatchEvents = "PRESENCE_UPDATE"
	EventReady                               DispatchEvents = "READY"
	EventResumed                             DispatchEvents = "RESUMED"
	EventStageInstanceCreate                 DispatchEvents = "STAGE_INSTANCE_CREATE"
	EventStageInstanceDelete                 DispatchEvents = "STAGE_INSTANCE_DELETE"
	EventStageInstanceUpdate                 DispatchEvents = "STAGE_INSTANCE_UPDATE"
	EventSubscriptionCreate                  DispatchEvents = "SUBSCRIPTION_CREATE"
	EventSubscriptionDelete                  DispatchEvents = "SUBSCRIPTION_DELETE"
	EventSubscriptionUpdate                  DispatchEvents = "SUBSCRIPTION_UPDATE"
	EventThreadCreate                        DispatchEvents = "THREAD_CREATE"
	EventThreadDelete                        DispatchEvents = "THREAD_DELETE"
	EventThreadListSync                      DispatchEvents = "THREAD_LIST_SYNC"
	EventThreadMembersUpdate                 DispatchEvents = "THREAD_MEMBERS_UPDATE"
	EventThreadMemberUpdate                  DispatchEvents = "THREAD_MEMBER_UPDATE"
	EventThreadUpdate                        DispatchEvents = "THREAD_UPDATE"
	EventTypingStart                         DispatchEvents = "TYPING_START"
	EventUserUpdate                          DispatchEvents = "USER_UPDATE"
	EventVoiceChannelEffectSend              DispatchEvents = "VOICE_CHANNEL_EFFECT_SEND"
	EventVoiceServerUpdate                   DispatchEvents = "VOICE_SERVER_UPDATE"
	EventVoiceStateUpdate                    DispatchEvents = "VOICE_STATE_UPDATE"
	EventWebhooksUpdate                      DispatchEvents = "WEBHOOKS_UPDATE"
)

// SendPayload represents payloads that can be sent to the Gateway.
type SendPayload[DATA any] struct {
	Op GatewayOpcode `json:"op"`
	D  DATA
}

func (p SendPayload[DATA]) isSendPayload() {}

// ReceivePayload represents payloads that can be received from the Gateway.
type ReceivePayload[DATA any] struct {
	Op GatewayOpcode `json:"op"`
	D  DATA
}

func (p ReceivePayload[DATA]) isReceivePayload() {}

// BasePayload represents the base structure for all Gateway payloads.
type BasePayload struct {
	// Op is the opcode for the payload.
	Op GatewayOpcode `json:"op"`

	// D is the event data.
	D interface{} `json:"d,omitempty"`

	// S is the sequence number, used for resuming sessions and heartbeats.
	S *int `json:"s,omitempty"`

	// T is the event name for this payload.
	T *string `json:"t,omitempty"`
}

// NonDispatchPayload represents a non-dispatch payload.
type NonDispatchPayload struct {
	BasePayload
}

func (p NonDispatchPayload) isReceivePayload() {}
func (p NonDispatchPayload) isSendPayload()    {}

// DispatchPayload represents a dispatch payload.
type DispatchPayload struct {
	BasePayload
	T string      `json:"t"`
	D interface{} `json:"d"`
}

func (p DispatchPayload) isReceivePayload() {}

// HelloData represents Hello payload data.
//
// See: https://discord.com/developers/docs/topics/gateway-events#hello
type HelloData struct {
	// HeartbeatInterval is the interval (in milliseconds) the client should heartbeat with.
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// Hello represents a Hello payload.
type Hello struct {
	Op GatewayOpcode `json:"op"`
	D  HelloData     `json:"d"`
}

func (p Hello) isReceivePayload() {}

// Heartbeat represents a Heartbeat payload.
//
// See: https://discord.com/developers/docs/topics/gateway#sending-heartbeats
type Heartbeat struct {
	Op GatewayOpcode `json:"op"`
	D  *int          `json:"d"` // Sequence number
}

func (p Heartbeat) isSendPayload() {}

// HeartbeatAck represents a Heartbeat ACK payload.
type HeartbeatAck struct {
	Op GatewayOpcode `json:"op"`
	D  *int          `json:"d"`
}

func (p HeartbeatAck) isReceivePayload() {}

// InvalidSessionData represents the data for an Invalid Session payload.
type InvalidSessionData bool

// InvalidSession represents an Invalid Session payload.
//
// See: https://discord.com/developers/docs/topics/gateway-events#invalid-session
type InvalidSession struct {
	Op GatewayOpcode      `json:"op"`
	D  InvalidSessionData `json:"d"`
}

func (p InvalidSession) isReceivePayload() {}

// Reconnect represents a Reconnect payload.
//
// See: https://discord.com/developers/docs/topics/gateway-events#reconnect
type Reconnect struct {
	Op GatewayOpcode `json:"op"`
}

func (p Reconnect) isReceivePayload() {}

// IdentifyData represents Identify payload data.
//
// See: https://discord.com/developers/docs/topics/gateway-events#identify
type IdentifyData struct {
	// Token is the authentication token.
	Token string `json:"token"`

	// Properties are the connection properties.
	Properties IdentifyProperties `json:"properties"`

	// Compress indicates whether this connection supports compression of packets.
	Compress *bool `json:"compress,omitempty"`

	// LargeThreshold is the value between 50 and 250, total number of members where the gateway will stop sending
	// offline members in the guild member list.
	LargeThreshold *int `json:"large_threshold,omitempty"`

	// Shard is used for Guild Sharding.
	Shard *[2]int `json:"shard,omitempty"`

	// Presence is the presence structure for initial presence information.
	Presence *PresenceUpdateData `json:"presence,omitempty"`

	// Intents are the Gateway Intents you wish to receive.
	Intents int `json:"intents"`
}

// IdentifyProperties represents connection properties for Identify.
//
// See: https://discord.com/developers/docs/topics/gateway-events#identify-identify-connection-properties
type IdentifyProperties struct {
	// OS is your operating system.
	OS string `json:"os"`

	// Browser is your library name.
	Browser string `json:"browser"`

	// Device is your library name.
	Device string `json:"device"`
}

// Identify represents an Identify payload.
type Identify struct {
	Op GatewayOpcode `json:"op"`
	D  IdentifyData  `json:"d"`
}

func (p Identify) isSendPayload() {}

// PresenceUpdateData represents presence update data.
type PresenceUpdateData struct {
	// Since is the unix time (in milliseconds) of when the client went idle, or null if the client is not idle.
	Since *int64 `json:"since"`

	// Activities are the user's activities.
	Activities []ActivityUpdateData `json:"activities"`

	// Status is the user's new status.
	Status PresenceUpdateStatus `json:"status"`

	// AFK indicates whether or not the client is afk.
	AFK bool `json:"afk"`
}

// ActivityUpdateData represents activity update data.
type ActivityUpdateData struct {
	// Name is the activity's name.
	Name string `json:"name"`

	// Type is the activity type.
	Type ActivityType `json:"type"`

	// URL is the stream URL, is validated when type is 1.
	URL *string `json:"url,omitempty"`

	// State is the user's current party status.
	State *string `json:"state,omitempty"`
}

// ActivityType represents activity types.
type ActivityType int

// Activity type constants
const (
	ActivityTypeGame ActivityType = iota
	ActivityTypeStreaming
	ActivityTypeListening
	ActivityTypeWatching
	ActivityTypeCustom
	ActivityTypeCompeting
)

// PresenceUpdateStatus represents presence update status.
type PresenceUpdateStatus string

// Presence status constants
const (
	PresenceStatusOnline    PresenceUpdateStatus = "online"
	PresenceStatusDND       PresenceUpdateStatus = "dnd"
	PresenceStatusIdle      PresenceUpdateStatus = "idle"
	PresenceStatusInvisible PresenceUpdateStatus = "invisible"
	PresenceStatusOffline   PresenceUpdateStatus = "offline"
)

// PresenceUpdate represents a Presence Update payload.
type PresenceUpdate struct {
	Op GatewayOpcode      `json:"op"`
	D  PresenceUpdateData `json:"d"`
}

func (p PresenceUpdate) isSendPayload() {}
