package v10

// Stage represents a Discord stage instance.
// Reference: https://discord.com/developers/docs/resources/stage-instance#stage-instance-object
type Stage struct {
	// ID is the id of this Stage instance.
	ID Snowflake `json:"id"`

	// GuildID is the guild id of the associated Stage channel.
	GuildID Snowflake `json:"guild_id"`

	// ChannelID is the id of the associated Stage channel.
	ChannelID Snowflake `json:"channel_id"`

	// Topic is the topic of the Stage instance (1-120 characters).
	Topic string `json:"topic"`

	// PrivacyLevel is the privacy level of the Stage instance.
	PrivacyLevel StagePrivacyLevel `json:"privacy_level"`

	// DiscoverableDisabled indicates whether or not Stage Discovery is disabled.
	DiscoverableDisabled bool `json:"discoverable_disabled"`

	// GuildScheduledEventID is the id of the scheduled event for this Stage instance.
	GuildScheduledEventID *Snowflake `json:"guild_scheduled_event_id,omitempty"`
}

// StagePrivacyLevel represents the privacy level of a Stage instance.
// Reference: https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-privacy-level
type StagePrivacyLevel int

// Stage instance privacy level constants
const (
	// Deprecated: The Stage instance is visible publicly.
	StagePrivacyLevelPublic StagePrivacyLevel = iota + 1

	// StagePrivacyLevelGuildOnly indicates the Stage instance is visible to only guild members.
	StagePrivacyLevelGuildOnly
)
