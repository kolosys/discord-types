package payloads

import "github.com/kolosys/discord-types/discord"

// SoundboardSound represents a Discord soundboard sound.
//
// See: https://discord.com/developers/docs/resources/soundboard#soundboard-sound-object
type SoundboardSound struct {
	// Name is the name of this sound.
	Name string `json:"name"`

	// SoundID is the id of this sound.
	SoundID discord.Snowflake `json:"sound_id"`

	// Volume is the volume of this sound, from 0 to 1.
	Volume float64 `json:"volume"`

	// EmojiID is the id of this sound's custom emoji.
	EmojiID *discord.Snowflake `json:"emoji_id"`

	// EmojiName is the unicode character of this sound's standard emoji.
	EmojiName *string `json:"emoji_name"`

	// GuildID is the id of the guild that this sound is in.
	GuildID *discord.Snowflake `json:"guild_id,omitempty"`

	// Available indicates whether this sound can be used (for guild sounds).
	// May be false due to loss of Server Boosts.
	Available bool `json:"available"`

	// User is the user who created this sound.
	User *User `json:"user,omitempty"`
}
