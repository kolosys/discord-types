package payloads

// ActionRowComponent represents an action row component
// https://discord.com/developers/docs/interactions/message-components#component-object
type ActionRowComponent struct {
	Type       ComponentType      `json:"type"`
	Components []MessageComponent `json:"components,omitempty"`
}

// MessageComponent represents a message component interface
type MessageComponent interface {
	GetType() ComponentType
}

// ButtonComponent represents a button component
type ButtonComponent struct {
	Type     ComponentType `json:"type"`
	Style    ButtonStyle   `json:"style"`
	Label    *string       `json:"label,omitempty"`
	Emoji    *PartialEmoji `json:"emoji,omitempty"`
	CustomID *string       `json:"custom_id,omitempty"`
	URL      *string       `json:"url,omitempty"`
	Disabled *bool         `json:"disabled,omitempty"`
}

func (b ButtonComponent) GetType() ComponentType {
	return ComponentTypeButton
}

// ButtonStyle represents button styles
type ButtonStyle int

const (
	ButtonStylePrimary ButtonStyle = iota + 1
	ButtonStyleSecondary
	ButtonStyleSuccess
	ButtonStyleDanger
	ButtonStyleLink
)

// SelectMenuComponent represents a select menu component
type SelectMenuComponent struct {
	Type         ComponentType  `json:"type"`
	CustomID     string         `json:"custom_id"`
	Options      []SelectOption `json:"options,omitempty"`
	ChannelTypes []ChannelType  `json:"channel_types,omitempty"`
	Placeholder  *string        `json:"placeholder,omitempty"`
	MinValues    *int           `json:"min_values,omitempty"`
	MaxValues    *int           `json:"max_values,omitempty"`
	Disabled     *bool          `json:"disabled,omitempty"`
}

func (s SelectMenuComponent) GetType() ComponentType {
	return s.Type
}

// SelectOption represents a select option
type SelectOption struct {
	Label       string        `json:"label"`
	Value       string        `json:"value"`
	Description *string       `json:"description,omitempty"`
	Emoji       *PartialEmoji `json:"emoji,omitempty"`
	Default     *bool         `json:"default,omitempty"`
}

// TextInputComponent represents a text input component
type TextInputComponent struct {
	Type        ComponentType  `json:"type"`
	CustomID    string         `json:"custom_id"`
	Style       TextInputStyle `json:"style"`
	Label       string         `json:"label"`
	MinLength   *int           `json:"min_length,omitempty"`
	MaxLength   *int           `json:"max_length,omitempty"`
	Required    *bool          `json:"required,omitempty"`
	Value       *string        `json:"value,omitempty"`
	Placeholder *string        `json:"placeholder,omitempty"`
}

func (t TextInputComponent) GetType() ComponentType {
	return ComponentTypeTextInput
}

// TextInputStyle represents text input styles
type TextInputStyle int

const (
	TextInputStyleShort TextInputStyle = iota + 1
	TextInputStyleParagraph
)
