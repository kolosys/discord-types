package gateway

import "github.com/kolosys/discord-types/discord"

// GatewayPayload is a generic base struct for all gateway payloads
type GatewayPayload[T any] struct {
	Op GatewayOpcode `json:"op"`
	D  T             `json:"d"`
}

// ResumeData represents Resume payload data.
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
type Resume = GatewayPayload[ResumeData]

// RequestGuildMembersData represents the data for requesting guild members.
type RequestGuildMembersData struct {
	// GuildID is the ID of the guild to get members for.
	GuildID discord.Snowflake `json:"guild_id"`

	// Presences indicates whether we want the presences of the matched members.
	Presences *bool `json:"presences,omitempty"`

	// Nonce is used to identify the Guild Members Chunk response.
	// Nonce can only be up to 32 bytes. If you send an invalid nonce it will be ignored
	// and the reply member_chunk(s) will not have a nonce set.
	Nonce *string `json:"nonce,omitempty"`

	// UserIDs are the user IDs to fetch (mutually exclusive with Query).
	UserIDs []discord.Snowflake `json:"user_ids,omitempty"`

	// Query is the string that username starts with, or an empty string to return all members.
	// Mutually exclusive with UserIDs.
	Query string `json:"query,omitempty"`

	// Limit is the maximum number of members to send matching the query.
	// A limit of 0 can be used with an empty string query to return all members.
	Limit int `json:"limit,omitempty"`
}

// RequestGuildMembers represents a Request Guild Members payload.
// See: https://discord.com/developers/docs/topics/gateway-events#request-guild-members
type RequestGuildMembers = GatewayPayload[RequestGuildMembersData]

// RequestSoundboardSoundsData represents Request Soundboard Sounds payload data.
// See: https://discord.com/developers/docs/topics/gateway-events#request-soundboard-sounds
type RequestSoundboardSoundsData struct {
	// GuildIDs are the ids of the guilds to get soundboard sounds for.
	GuildIDs []discord.Snowflake `json:"guild_ids"`
}

// RequestSoundboardSounds represents a Request Soundboard Sounds payload.
type RequestSoundboardSounds = GatewayPayload[RequestSoundboardSoundsData]

// VoiceStateUpdateData represents Voice State Update payload data.
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
type VoiceStateUpdate = GatewayPayload[VoiceStateUpdateData]

// UpdatePresenceData represents Update Presence payload data.
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
type UpdatePresence = GatewayPayload[UpdatePresenceData]

// GatewaySendPayload represents all possible sendable payloads.
type GatewaySendPayload interface {
	SendPayload[any]
}

// GatewayReceivePayload represents all possible receivable payloads.
type GatewayReceivePayload interface {
	ReceivePayload[any]
}
