package v10

// ApplicationCommand represents an application command
// Reference: https://discord.com/developers/docs/interactions/application-commands#application-command-object
type ApplicationCommand struct {
	ID                       Snowflake                    `json:"id"`
	Type                     ApplicationCommandType       `json:"type"`
	ApplicationID            Snowflake                    `json:"application_id"`
	GuildID                  *Snowflake                   `json:"guild_id,omitempty"`
	Name                     string                       `json:"name"`
	NameLocalizations        LocalizationMap              `json:"name_localizations,omitempty"`
	Description              string                       `json:"description"`
	DescriptionLocalizations LocalizationMap              `json:"description_localizations,omitempty"`
	Options                  []ApplicationCommandOption   `json:"options"` // Can only be set for application commands of type CHAT_INPUT.
	DefaultMemberPermissions *Permissions                 `json:"default_member_permissions,omitempty"`
	DMPermission             *bool                        `json:"dm_permission,omitempty"`
	DefaultPermission        *bool                        `json:"default_permission,omitempty"`
	NSFW                     *bool                        `json:"nsfw,omitempty"`
	IntegrationTypes         []ApplicationIntegrationType `json:"integration_types,omitempty"`
	Contexts                 []ApplicationCommandContexts `json:"contexts,omitempty"`
	Handler                  *string                      `json:"handler,omitempty"` // Can only be set for application commands of type PRIMARY_ENTRY_POINT for applications with the EMBEDDED flag (i.e. applications that have an Activity).
}

// ApplicationCommandOption represents an option for an application command
// Reference: https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-structure
type ApplicationCommandOption struct {
	Type                     ApplicationCommandOptionType     `json:"type"`
	Name                     string                           `json:"name"` // Must be unique within an array of application command options.
	NameLocalizations        LocalizationMap                  `json:"name_localizations,omitempty"`
	Description              string                           `json:"description"`
	DescriptionLocalizations LocalizationMap                  `json:"description_localizations,omitempty"`
	Required                 *bool                            `json:"required,omitempty"`
	Choices                  []ApplicationCommandOptionChoice `json:"choices,omitempty"`
	Options                  []ApplicationCommandOption       `json:"options,omitempty"`
	ChannelTypes             []ChannelType                    `json:"channel_types,omitempty"`
	MinValue                 *float64                         `json:"min_value,omitempty"`
	MaxValue                 *float64                         `json:"max_value,omitempty"`
	MinLength                *int                             `json:"min_length,omitempty"`
	MaxLength                *int                             `json:"max_length,omitempty"`
	Autocomplete             *bool                            `json:"autocomplete,omitempty"` // May not be set to true if choices are present.
}

// ApplicationCommandType represents the type of application command
// Reference: https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = iota + 1
	ApplicationCommandTypeUser
	ApplicationCommandTypeMessage
)

// ApplicationCommandOptionType represents the type of application command option
// Reference: https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-type
type ApplicationCommandOptionType int

const (
	ApplicationCommandOptionTypeSubCommand ApplicationCommandOptionType = iota + 1
	ApplicationCommandOptionTypeSubCommandGroup
	ApplicationCommandOptionTypeString
	ApplicationCommandOptionTypeInteger
	ApplicationCommandOptionTypeBoolean
	ApplicationCommandOptionTypeUser
	ApplicationCommandOptionTypeChannel
	ApplicationCommandOptionTypeRole
	ApplicationCommandOptionTypeMentionable
	ApplicationCommandOptionTypeNumber
	ApplicationCommandOptionTypeAttachment
)

// ApplicationCommandContexts represents the contexts of an application command
// Reference: https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-context-types
type ApplicationCommandContexts int

const (
	ApplicationCommandContextsGuild ApplicationCommandContexts = iota
	ApplicationCommandContextsDM
	ApplicationCommandContextsGuildAndDM
)
