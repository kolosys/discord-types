package payloads

import "github.com/kolosys/discord-types/discord"

// StageInstance represents a Discord stage instance.
//
// See: https://discord.com/developers/docs/resources/stage-instance#stage-instance-object
type StageInstance struct {
	// ID is the id of this Stage instance.
	ID discord.Snowflake `json:"id"`

	// GuildID is the guild id of the associated Stage channel.
	GuildID discord.Snowflake `json:"guild_id"`

	// ChannelID is the id of the associated Stage channel.
	ChannelID discord.Snowflake `json:"channel_id"`

	// Topic is the topic of the Stage instance (1-120 characters).
	Topic string `json:"topic"`

	// PrivacyLevel is the privacy level of the Stage instance.
	PrivacyLevel StageInstancePrivacyLevel `json:"privacy_level"`

	// DiscoverableDisabled indicates whether or not Stage Discovery is disabled.
	DiscoverableDisabled bool `json:"discoverable_disabled"`

	// GuildScheduledEventID is the id of the scheduled event for this Stage instance.
	GuildScheduledEventID *discord.Snowflake `json:"guild_scheduled_event_id,omitempty"`
}

// StageInstancePrivacyLevel represents the privacy level of a Stage instance.
//
// See: https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-privacy-level
type StageInstancePrivacyLevel int

// Stage instance privacy level constants
const (
	// Deprecated: The Stage instance is visible publicly.
	StageInstancePrivacyLevelPublic StageInstancePrivacyLevel = iota + 1

	// StageInstancePrivacyLevelGuildOnly indicates the Stage instance is visible to only guild members.
	StageInstancePrivacyLevelGuildOnly
)
