package v10

// PartialEmoji represents a partial emoji object
// https://discord.com/developers/docs/resources/emoji#emoji-object
type PartialEmoji struct {
	ID       *Snowflake `json:"id"`
	Name     *string    `json:"name"`
	Animated *bool      `json:"animated,omitempty"`
}

// Emoji represents a Discord emoji
// https://discord.com/developers/docs/resources/emoji#emoji-object-emoji-structure
type Emoji struct {
	PartialEmoji
	Roles         []Snowflake `json:"roles,omitempty"`
	User          *User       `json:"user,omitempty"`
	RequireColons *bool       `json:"require_colons,omitempty"`
	Managed       *bool       `json:"managed,omitempty"`
	Available     *bool       `json:"available,omitempty"`
}

// ApplicationEmoji represents an application-owned emoji
// https://discord.com/developers/docs/resources/emoji#emoji-object-applicationowned-emoji
type ApplicationEmoji struct {
	ID            Snowflake   `json:"id"`
	Name          string      `json:"name"`
	Animated      bool        `json:"animated"`
	User          User        `json:"user"`
	Roles         []Snowflake `json:"roles"`          // Always empty
	RequireColons bool        `json:"require_colons"` // Always true
	Managed       bool        `json:"managed"`        // Always false
	Available     bool        `json:"available"`      // Always true
}
