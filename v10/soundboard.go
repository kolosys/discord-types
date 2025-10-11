package v10

import "fmt"

// SoundboardSound represents a Discord soundboard sound.
// Reference: https://discord.com/developers/docs/resources/soundboard#soundboard-sound-object
type SoundboardSound struct {
	// ID is the unique identifier of this sound.
	ID Snowflake `json:"id"`

	// Name is the name of this sound.
	Name string `json:"name"`

	// Volume is the volume of this sound, from 0 to 1.
	Volume float64 `json:"volume"`

	// EmojiID is the id of this sound's custom emoji.
	EmojiID *Snowflake `json:"emoji_id,omitempty"`

	// EmojiName is the unicode character of this sound's standard emoji.
	EmojiName *string `json:"emoji_name,omitempty"`

	// GuildID is the id of the guild this sound belongs to.
	// Null for default sounds.
	GuildID *Snowflake `json:"guild_id,omitempty"`

	// Available indicates whether this sound can be used.
	// May be false due to loss of Server Boosts for guild sounds.
	Available bool `json:"available"`

	// User is the user who created this sound.
	// Only present for guild sounds.
	User *User `json:"user,omitempty"`
}

// Validate checks if the soundboard sound data is valid.
func (s *SoundboardSound) Validate() error {
	if s.ID == "" {
		return fmt.Errorf("soundboard sound ID cannot be empty")
	}
	if s.Name == "" {
		return fmt.Errorf("soundboard sound name cannot be empty")
	}
	if s.Volume < 0 || s.Volume > 1 {
		return fmt.Errorf("soundboard sound volume must be between 0 and 1, got %f", s.Volume)
	}
	return nil
}

// IsGuildSound returns true if this is a guild-specific sound (not a default sound).
func (s *SoundboardSound) IsGuildSound() bool {
	return s.GuildID != nil
}

// IsDefaultSound returns true if this is a default Discord sound.
func (s *SoundboardSound) IsDefaultSound() bool {
	return s.GuildID == nil
}

// HasEmoji returns true if this sound has an associated emoji.
func (s *SoundboardSound) HasEmoji() bool {
	return s.EmojiID != nil || s.EmojiName != nil
}

// IsCreatedBy returns true if this sound was created by the specified user.
func (s *SoundboardSound) IsCreatedBy(userID Snowflake) bool {
	return s.User != nil && s.User.ID == userID
}
