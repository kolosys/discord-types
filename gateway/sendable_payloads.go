package gateway

import "github.com/kolosys/discord-types/discord"

// This file contains comprehensive sendable payload types
// matching the TypeScript discord-api-types library

// =============================================================================
// RESUME PAYLOAD
// =============================================================================

// ResumeData represents Resume payload data.
//
// See: https://discord.com/developers/docs/topics/gateway-events#resume
type ResumeData struct {
	// Token is the session token.
	Token string `json:"token"`

	// SessionID is the session id.
	SessionID string `json:"session_id"`

	// Seq is the last sequence number received.
	Seq int `json:"seq"`
}

// Resume represents a Resume payload.
type Resume struct {
	Op Opcodes    `json:"op"`
	D  ResumeData `json:"d"`
}

func (p Resume) isSendPayload() {}

// =============================================================================
// REQUEST GUILD MEMBERS PAYLOAD
// =============================================================================

// RequestGuildMembersDataBase represents the base structure for requesting guild members.
type RequestGuildMembersDataBase struct {
	// GuildID is the ID of the guild to get members for.
	GuildID discord.Snowflake `json:"guild_id"`

	// Presences indicates whether we want the presences of the matched members.
	Presences *bool `json:"presences,omitempty"`

	// Nonce is used to identify the Guild Members Chunk response.
	// Nonce can only be up to 32 bytes. If you send an invalid nonce it will be ignored
	// and the reply member_chunk(s) will not have a nonce set.
	Nonce *string `json:"nonce,omitempty"`
}

// RequestGuildMembersDataWithUserIDs represents requesting guild members with specific user IDs.
type RequestGuildMembersDataWithUserIDs struct {
	RequestGuildMembersDataBase

	// UserIDs are the user IDs to fetch.
	UserIDs []discord.Snowflake `json:"user_ids"`
}

// RequestGuildMembersDataWithQuery represents requesting guild members with a query.
type RequestGuildMembersDataWithQuery struct {
	RequestGuildMembersDataBase

	// Query is the string that username starts with, or an empty string to return all members.
	Query string `json:"query"`

	// Limit is the maximum number of members to send matching the query.
	// A limit of 0 can be used with an empty string query to return all members.
	Limit int `json:"limit"`
}

// RequestGuildMembersData represents the data for requesting guild members.
// This is a union type that can be either WithUserIDs or WithQuery.
type RequestGuildMembersData interface {
	isRequestGuildMembersData()
}

func (r RequestGuildMembersDataWithUserIDs) isRequestGuildMembersData() {}
func (r RequestGuildMembersDataWithQuery) isRequestGuildMembersData()   {}

// RequestGuildMembers represents a Request Guild Members payload.
//
// See: https://discord.com/developers/docs/topics/gateway-events#request-guild-members
type RequestGuildMembers struct {
	Op Opcodes                 `json:"op"`
	D  RequestGuildMembersData `json:"d"`
}

func (p RequestGuildMembers) isSendPayload() {}

// =============================================================================
// REQUEST SOUNDBOARD SOUNDS PAYLOAD
// =============================================================================

// RequestSoundboardSoundsData represents Request Soundboard Sounds payload data.
//
// See: https://discord.com/developers/docs/topics/gateway-events#request-soundboard-sounds
type RequestSoundboardSoundsData struct {
	// GuildIDs are the ids of the guilds to get soundboard sounds for.
	GuildIDs []discord.Snowflake `json:"guild_ids"`
}

// RequestSoundboardSounds represents a Request Soundboard Sounds payload.
type RequestSoundboardSounds struct {
	Op Opcodes                     `json:"op"`
	D  RequestSoundboardSoundsData `json:"d"`
}

func (p RequestSoundboardSounds) isSendPayload() {}

// =============================================================================
// VOICE STATE UPDATE PAYLOAD
// =============================================================================

// VoiceStateUpdateData represents Voice State Update payload data.
//
// See: https://discord.com/developers/docs/topics/gateway-events#update-voice-state
type VoiceStateUpdateData struct {
	// GuildID is the ID of the guild.
	GuildID discord.Snowflake `json:"guild_id"`

	// ChannelID is the ID of the voice channel client wants to join (null if disconnecting).
	ChannelID *discord.Snowflake `json:"channel_id"`

	// SelfMute indicates if the client is muted.
	SelfMute bool `json:"self_mute"`

	// SelfDeaf indicates if the client is deafened.
	SelfDeaf bool `json:"self_deaf"`
}

// VoiceStateUpdate represents a Voice State Update payload.
type VoiceStateUpdate struct {
	Op Opcodes              `json:"op"`
	D  VoiceStateUpdateData `json:"d"`
}

func (p VoiceStateUpdate) isSendPayload() {}

// =============================================================================
// UPDATE PRESENCE PAYLOAD (Enhanced)
// =============================================================================

// UpdatePresenceData represents Update Presence payload data.
//
// See: https://discord.com/developers/docs/topics/gateway-events#update-presence
type UpdatePresenceData struct {
	// Since is the unix time (in milliseconds) of when the client went idle, or null if the client is not idle.
	Since *int64 `json:"since"`

	// Activities are the user's activities.
	Activities []ActivityUpdateData `json:"activities"`

	// Status is the user's new status.
	Status PresenceUpdateStatus `json:"status"`

	// AFK indicates whether or not the client is afk.
	AFK bool `json:"afk"`
}

// UpdatePresence represents an Update Presence payload.
type UpdatePresence struct {
	Op Opcodes            `json:"op"`
	D  UpdatePresenceData `json:"d"`
}

func (p UpdatePresence) isSendPayload() {}

// =============================================================================
// SEND PAYLOAD UNION TYPE
// =============================================================================

// GatewaySendPayload represents all possible sendable payloads.
type GatewaySendPayload interface {
	SendPayload
	isGatewaySendPayload()
}

// Implement the interface for all sendable payloads
func (p Heartbeat) isGatewaySendPayload()               {}
func (p Identify) isGatewaySendPayload()                {}
func (p RequestGuildMembers) isGatewaySendPayload()     {}
func (p RequestSoundboardSounds) isGatewaySendPayload() {}
func (p Resume) isGatewaySendPayload()                  {}
func (p UpdatePresence) isGatewaySendPayload()          {}
func (p VoiceStateUpdate) isGatewaySendPayload()        {}

// =============================================================================
// RECEIVE PAYLOAD UNION TYPE
// =============================================================================

// GatewayReceivePayload represents all possible receivable payloads.
type GatewayReceivePayload interface {
	ReceivePayload
	isGatewayReceivePayload()
}

// Implement the interface for all receivable payloads
func (p Hello) isGatewayReceivePayload()                 {}
func (p HeartbeatAck) isGatewayReceivePayload()          {}
func (p InvalidSession) isGatewayReceivePayload()        {}
func (p Reconnect) isGatewayReceivePayload()             {}
func (p ReadyDispatch) isGatewayReceivePayload()         {}
func (p MessageCreateDispatch) isGatewayReceivePayload() {}
func (p GuildCreateDispatch) isGatewayReceivePayload()   {}

// Add all the dispatch events from dispatch_events.go
func (p AutoModerationRuleCreateDispatch) isGatewayReceivePayload()            {}
func (p AutoModerationRuleUpdateDispatch) isGatewayReceivePayload()            {}
func (p AutoModerationRuleDeleteDispatch) isGatewayReceivePayload()            {}
func (p AutoModerationActionExecutionDispatch) isGatewayReceivePayload()       {}
func (p ApplicationCommandPermissionsUpdateDispatch) isGatewayReceivePayload() {}
func (p ChannelCreateDispatch) isGatewayReceivePayload()                       {}
func (p ChannelUpdateDispatch) isGatewayReceivePayload()                       {}
func (p ChannelDeleteDispatch) isGatewayReceivePayload()                       {}
func (p ChannelPinsUpdateDispatch) isGatewayReceivePayload()                   {}
func (p EntitlementCreateDispatch) isGatewayReceivePayload()                   {}
func (p EntitlementUpdateDispatch) isGatewayReceivePayload()                   {}
func (p EntitlementDeleteDispatch) isGatewayReceivePayload()                   {}
func (p GuildAuditLogEntryCreateDispatch) isGatewayReceivePayload()            {}
func (p GuildBanAddDispatch) isGatewayReceivePayload()                         {}
func (p GuildBanRemoveDispatch) isGatewayReceivePayload()                      {}
func (p GuildDeleteDispatch) isGatewayReceivePayload()                         {}
func (p GuildUpdateDispatch) isGatewayReceivePayload()                         {}
func (p GuildEmojisUpdateDispatch) isGatewayReceivePayload()                   {}
func (p GuildStickersUpdateDispatch) isGatewayReceivePayload()                 {}
func (p GuildIntegrationsUpdateDispatch) isGatewayReceivePayload()             {}
func (p GuildMemberAddDispatch) isGatewayReceivePayload()                      {}
func (p GuildMemberRemoveDispatch) isGatewayReceivePayload()                   {}
func (p GuildMemberUpdateDispatch) isGatewayReceivePayload()                   {}
func (p GuildMembersChunkDispatch) isGatewayReceivePayload()                   {}
func (p GuildRoleCreateDispatch) isGatewayReceivePayload()                     {}
func (p GuildRoleUpdateDispatch) isGatewayReceivePayload()                     {}
func (p GuildRoleDeleteDispatch) isGatewayReceivePayload()                     {}
func (p MessageUpdateDispatch) isGatewayReceivePayload()                       {}
func (p MessageDeleteDispatch) isGatewayReceivePayload()                       {}
func (p MessageDeleteBulkDispatch) isGatewayReceivePayload()                   {}
func (p MessagePollVoteAddDispatch) isGatewayReceivePayload()                  {}
func (p MessagePollVoteRemoveDispatch) isGatewayReceivePayload()               {}
func (p MessageReactionAddDispatch) isGatewayReceivePayload()                  {}
func (p MessageReactionRemoveDispatch) isGatewayReceivePayload()               {}
func (p MessageReactionRemoveAllDispatch) isGatewayReceivePayload()            {}
func (p MessageReactionRemoveEmojiDispatch) isGatewayReceivePayload()          {}
func (p VoiceServerUpdateDispatch) isGatewayReceivePayload()                   {}
func (p VoiceStateUpdateDispatch) isGatewayReceivePayload()                    {}
func (p VoiceChannelEffectSendDispatch) isGatewayReceivePayload()              {}

// Add implementations for additional dispatch events
func (p GuildScheduledEventCreateDispatch) isGatewayReceivePayload()     {}
func (p GuildScheduledEventUpdateDispatch) isGatewayReceivePayload()     {}
func (p GuildScheduledEventDeleteDispatch) isGatewayReceivePayload()     {}
func (p GuildScheduledEventUserAddDispatch) isGatewayReceivePayload()    {}
func (p GuildScheduledEventUserRemoveDispatch) isGatewayReceivePayload() {}
func (p GuildSoundboardSoundCreateDispatch) isGatewayReceivePayload()    {}
func (p GuildSoundboardSoundUpdateDispatch) isGatewayReceivePayload()    {}
func (p GuildSoundboardSoundDeleteDispatch) isGatewayReceivePayload()    {}
func (p GuildSoundboardSoundsUpdateDispatch) isGatewayReceivePayload()   {}
func (p SoundboardSoundsDispatch) isGatewayReceivePayload()              {}
func (p StageInstanceCreateDispatch) isGatewayReceivePayload()           {}
func (p StageInstanceUpdateDispatch) isGatewayReceivePayload()           {}
func (p StageInstanceDeleteDispatch) isGatewayReceivePayload()           {}
func (p SubscriptionCreateDispatch) isGatewayReceivePayload()            {}
func (p SubscriptionUpdateDispatch) isGatewayReceivePayload()            {}
func (p SubscriptionDeleteDispatch) isGatewayReceivePayload()            {}
func (p ThreadCreateDispatch) isGatewayReceivePayload()                  {}
func (p ThreadUpdateDispatch) isGatewayReceivePayload()                  {}
func (p ThreadDeleteDispatch) isGatewayReceivePayload()                  {}
func (p ThreadListSyncDispatch) isGatewayReceivePayload()                {}
func (p ThreadMembersUpdateDispatch) isGatewayReceivePayload()           {}
func (p ThreadMemberUpdateDispatch) isGatewayReceivePayload()            {}
func (p TypingStartDispatch) isGatewayReceivePayload()                   {}
func (p UserUpdateDispatch) isGatewayReceivePayload()                    {}
func (p PresenceUpdateDispatch) isGatewayReceivePayload()                {}
func (p IntegrationCreateDispatch) isGatewayReceivePayload()             {}
func (p IntegrationUpdateDispatch) isGatewayReceivePayload()             {}
func (p IntegrationDeleteDispatch) isGatewayReceivePayload()             {}
func (p InteractionCreateDispatch) isGatewayReceivePayload()             {}
func (p InviteCreateDispatch) isGatewayReceivePayload()                  {}
func (p InviteDeleteDispatch) isGatewayReceivePayload()                  {}
func (p WebhooksUpdateDispatch) isGatewayReceivePayload()                {}
