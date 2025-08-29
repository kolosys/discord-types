package payloads

import "github.com/kolosys/discord-types/discord"

// VoiceState represents a Discord voice state.
//
// See: https://discord.com/developers/docs/resources/voice#voice-state-object
type VoiceState struct {
	// GuildID is the guild id this voice state is for.
	GuildID *discord.Snowflake `json:"guild_id,omitempty"`

	// ChannelID is the channel id this user is connected to.
	ChannelID *discord.Snowflake `json:"channel_id"`

	// UserID is the user id this voice state is for.
	UserID discord.Snowflake `json:"user_id"`

	// Member is the guild member this voice state is for.
	Member *GuildMember `json:"member,omitempty"`

	// SessionID is the session id for this voice state.
	SessionID string `json:"session_id"`

	// Deaf indicates whether this user is deafened by the server.
	Deaf bool `json:"deaf"`

	// Mute indicates whether this user is muted by the server.
	Mute bool `json:"mute"`

	// SelfDeaf indicates whether this user is locally deafened.
	SelfDeaf bool `json:"self_deaf"`

	// SelfMute indicates whether this user is locally muted.
	SelfMute bool `json:"self_mute"`

	// SelfStream indicates whether this user is streaming using "Go Live".
	SelfStream *bool `json:"self_stream,omitempty"`

	// SelfVideo indicates whether this user's camera is enabled.
	SelfVideo bool `json:"self_video"`

	// Suppress indicates whether this user's permission to speak is denied.
	Suppress bool `json:"suppress"`

	// RequestToSpeakTimestamp is the time at which the user requested to speak.
	RequestToSpeakTimestamp *string `json:"request_to_speak_timestamp"`
}

// VoiceRegion represents a Discord voice region.
//
// See: https://discord.com/developers/docs/resources/voice#voice-region-object
type VoiceRegion struct {
	// ID is the unique ID for the region.
	ID string `json:"id"`

	// Name is the name of the region.
	Name string `json:"name"`

	// Optimal indicates whether this is the closest server to the current user's client.
	Optimal bool `json:"optimal"`

	// Deprecated indicates whether this is a deprecated voice region (avoid switching to these).
	Deprecated bool `json:"deprecated"`

	// Custom indicates whether this is a custom voice region (used for events/etc).
	Custom bool `json:"custom"`
}
