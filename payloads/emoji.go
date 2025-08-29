package payloads

import "github.com/kolosys/discord-types/discord"

// PartialEmoji represents a partial emoji object
// Not documented but mentioned in Discord API
type PartialEmoji struct {
	ID       *discord.Snowflake `json:"id"`
	Name     *string            `json:"name"`
	Animated *bool              `json:"animated,omitempty"`
}

// Emoji represents a Discord emoji
// https://discord.com/developers/docs/resources/emoji#emoji-object-emoji-structure
type Emoji struct {
	PartialEmoji
	Roles         []discord.Snowflake `json:"roles,omitempty"`
	User          *User               `json:"user,omitempty"`
	RequireColons *bool               `json:"require_colons,omitempty"`
	Managed       *bool               `json:"managed,omitempty"`
	Available     *bool               `json:"available,omitempty"`
}

// ApplicationEmoji represents an application-owned emoji
// https://discord.com/developers/docs/resources/emoji#emoji-object-applicationowned-emoji
type ApplicationEmoji struct {
	ID            discord.Snowflake   `json:"id"`
	Name          string              `json:"name"`
	Animated      bool                `json:"animated"`
	User          User                `json:"user"`
	Roles         []discord.Snowflake `json:"roles"`          // Always empty
	RequireColons bool                `json:"require_colons"` // Always true
	Managed       bool                `json:"managed"`        // Always false
	Available     bool                `json:"available"`      // Always true
}

// ReactionType represents the type of reaction
type ReactionType int

const (
	ReactionTypeNormal ReactionType = iota
	ReactionTypeBurst
)

// Reaction represents a message reaction
type Reaction struct {
	Count        int                  `json:"count"`
	CountDetails ReactionCountDetails `json:"count_details"`
	Me           bool                 `json:"me"`
	MeBurst      bool                 `json:"me_burst"`
	Emoji        PartialEmoji         `json:"emoji"`
	BurstColors  []string             `json:"burst_colors,omitempty"`
}

// ReactionCountDetails provides detailed reaction count information
type ReactionCountDetails struct {
	Burst  int `json:"burst"`
	Normal int `json:"normal"`
}
