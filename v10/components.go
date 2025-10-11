package v10

import (
	"encoding/json"
	"fmt"
)

// ActionRowComponent represents an action row component
// https://discord.com/developers/docs/interactions/message-components#component-object
type ActionRowComponent struct {
	Type       ComponentType      `json:"type"`
	Components []MessageComponent `json:"components,omitempty"`
}

func (arc ActionRowComponent) GetType() ComponentType {
	return ComponentTypeActionRow
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

// UnmarshalJSON implements custom JSON unmarshaling for ActionRowComponent
func (arc *ActionRowComponent) UnmarshalJSON(data []byte) error {
	// First unmarshal into a temporary struct to get the basic fields
	type tempActionRow struct {
		Type       ComponentType     `json:"type"`
		Components []json.RawMessage `json:"components,omitempty"`
	}

	var temp tempActionRow
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	arc.Type = temp.Type

	// Now unmarshal each component based on its type
	for _, rawComponent := range temp.Components {
		component, err := UnmarshalMessageComponent(rawComponent)
		if err != nil {
			return err
		}
		arc.Components = append(arc.Components, component)
	}

	return nil
}

// UnmarshalMessageComponent unmarshals a raw JSON message into the appropriate MessageComponent type
func UnmarshalMessageComponent(data []byte) (MessageComponent, error) {
	// First, peek at the type field
	var typeInfo struct {
		Type ComponentType `json:"type"`
	}

	if err := json.Unmarshal(data, &typeInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal component type: %w", err)
	}

	// Create the appropriate component type based on the type field
	switch typeInfo.Type {
	case ComponentTypeButton:
		var component ButtonComponent
		if err := json.Unmarshal(data, &component); err != nil {
			return nil, fmt.Errorf("failed to unmarshal ButtonComponent: %w", err)
		}
		return component, nil

	case ComponentTypeStringSelect, ComponentTypeUserSelect, ComponentTypeRoleSelect, ComponentTypeMentionableSelect, ComponentTypeChannelSelect:
		var component SelectMenuComponent
		if err := json.Unmarshal(data, &component); err != nil {
			return nil, fmt.Errorf("failed to unmarshal SelectMenuComponent: %w", err)
		}
		return component, nil

	case ComponentTypeTextInput:
		var component TextInputComponent
		if err := json.Unmarshal(data, &component); err != nil {
			return nil, fmt.Errorf("failed to unmarshal TextInputComponent: %w", err)
		}
		return component, nil

	case ComponentTypeActionRow:
		var component ActionRowComponent
		if err := json.Unmarshal(data, &component); err != nil {
			return nil, fmt.Errorf("failed to unmarshal ActionRowComponent: %w", err)
		}
		return component, nil

	default:
		return nil, fmt.Errorf("unknown component type: %d", typeInfo.Type)
	}
}

// ComponentTypeChecker provides utility functions for checking component types
type ComponentTypeChecker struct{}

// IsButton checks if a MessageComponent is a ButtonComponent
func (ComponentTypeChecker) IsButton(component MessageComponent) bool {
	return component.GetType() == ComponentTypeButton
}

// IsSelectMenu checks if a MessageComponent is a SelectMenuComponent
func (ComponentTypeChecker) IsSelectMenu(component MessageComponent) bool {
	componentType := component.GetType()
	return componentType == ComponentTypeStringSelect ||
		componentType == ComponentTypeUserSelect ||
		componentType == ComponentTypeRoleSelect ||
		componentType == ComponentTypeMentionableSelect ||
		componentType == ComponentTypeChannelSelect
}

// IsTextInput checks if a MessageComponent is a TextInputComponent
func (ComponentTypeChecker) IsTextInput(component MessageComponent) bool {
	return component.GetType() == ComponentTypeTextInput
}

// IsActionRow checks if a MessageComponent is an ActionRowComponent
func (ComponentTypeChecker) IsActionRow(component MessageComponent) bool {
	return component.GetType() == ComponentTypeActionRow
}

// AsButton safely casts a MessageComponent to ButtonComponent
func (ComponentTypeChecker) AsButton(component MessageComponent) (*ButtonComponent, bool) {
	if button, ok := component.(ButtonComponent); ok {
		return &button, true
	}
	return nil, false
}

// AsSelectMenu safely casts a MessageComponent to SelectMenuComponent
func (ComponentTypeChecker) AsSelectMenu(component MessageComponent) (*SelectMenuComponent, bool) {
	if selectMenu, ok := component.(SelectMenuComponent); ok {
		return &selectMenu, true
	}
	return nil, false
}

// AsTextInput safely casts a MessageComponent to TextInputComponent
func (ComponentTypeChecker) AsTextInput(component MessageComponent) (*TextInputComponent, bool) {
	if textInput, ok := component.(TextInputComponent); ok {
		return &textInput, true
	}
	return nil, false
}

// AsActionRow safely casts a MessageComponent to ActionRowComponent
func (ComponentTypeChecker) AsActionRow(component MessageComponent) (*ActionRowComponent, bool) {
	if actionRow, ok := component.(ActionRowComponent); ok {
		return &actionRow, true
	}
	return nil, false
}

// Global instance for convenience
var ComponentChecker = ComponentTypeChecker{}
