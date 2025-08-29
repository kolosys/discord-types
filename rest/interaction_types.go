// Package rest provides Discord REST API types and utilities.
//
// This file contains comprehensive interaction-related REST API request and response types.
package rest

import (
	"github.com/kolosys/discord-types/discord"
	"github.com/kolosys/discord-types/payloads"
)

// ====================
// Interaction Types - Complete Implementation
// ====================

// GetApplicationCommandsQuery represents query parameters for GET /applications/{application.id}/commands
type GetApplicationCommandsQuery struct {
	// Whether to include full localization dictionaries instead of the localized fields
	WithLocalizations *bool `json:"with_localizations,omitempty"`
}

// GetApplicationCommandsResult represents the response from GET /applications/{application.id}/commands
type GetApplicationCommandsResult = []ApplicationCommand

// PostApplicationCommandJSONBody represents the request body for POST /applications/{application.id}/commands
type PostApplicationCommandJSONBody struct {
	// Name of command, 1-32 characters
	Name string `json:"name"`
	// Localization dictionary for the name field
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	// 1-100 character description for CHAT_INPUT commands
	Description *string `json:"description,omitempty"`
	// Localization dictionary for the description field
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	// Parameters for the command, max 25
	Options []ApplicationCommandOption `json:"options,omitempty"`
	// Set of permissions represented as a bit set
	DefaultMemberPermissions *discord.Permissions `json:"default_member_permissions,omitempty"`
	// Indicates whether the command is available in DMs with the app
	DMPermission *bool `json:"dm_permission,omitempty"`
	// Type of command (default 1)
	Type *payloads.ApplicationCommandType `json:"type,omitempty"`
	// Indicates whether the command is age-restricted
	NSFW *bool `json:"nsfw,omitempty"`
	// Installation contexts where the command is available
	IntegrationTypes []payloads.ApplicationIntegrationType `json:"integration_types,omitempty"`
	// Interaction context(s) where the command can be used
	Contexts []payloads.InteractionContextType `json:"contexts,omitempty"`
}

// PostApplicationCommandResult represents the response from POST /applications/{application.id}/commands
type PostApplicationCommandResult = ApplicationCommand

// GetApplicationCommandResult represents the response from GET /applications/{application.id}/commands/{command.id}
type GetApplicationCommandResult = ApplicationCommand

// PatchApplicationCommandJSONBody represents the request body for PATCH /applications/{application.id}/commands/{command.id}
type PatchApplicationCommandJSONBody struct {
	// Name of command, 1-32 characters
	Name *string `json:"name,omitempty"`
	// Localization dictionary for the name field
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	// 1-100 character description for CHAT_INPUT commands
	Description *string `json:"description,omitempty"`
	// Localization dictionary for the description field
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	// Parameters for the command, max 25
	Options []ApplicationCommandOption `json:"options,omitempty"`
	// Set of permissions represented as a bit set
	DefaultMemberPermissions *discord.Permissions `json:"default_member_permissions,omitempty"`
	// Indicates whether the command is available in DMs with the app
	DMPermission *bool `json:"dm_permission,omitempty"`
	// Indicates whether the command is age-restricted
	NSFW *bool `json:"nsfw,omitempty"`
	// Installation contexts where the command is available
	IntegrationTypes []payloads.ApplicationIntegrationType `json:"integration_types,omitempty"`
	// Interaction context(s) where the command can be used
	Contexts []payloads.InteractionContextType `json:"contexts,omitempty"`
}

// PatchApplicationCommandResult represents the response from PATCH /applications/{application.id}/commands/{command.id}
type PatchApplicationCommandResult = ApplicationCommand

// DeleteApplicationCommandResult represents the response from DELETE /applications/{application.id}/commands/{command.id}
type DeleteApplicationCommandResult struct{} // Empty response

// PutApplicationCommandsJSONBody represents the request body for PUT /applications/{application.id}/commands
type PutApplicationCommandsJSONBody = []PostApplicationCommandJSONBody

// PutApplicationCommandsResult represents the response from PUT /applications/{application.id}/commands
type PutApplicationCommandsResult = []ApplicationCommand

// GetGuildApplicationCommandsQuery represents query parameters for GET /applications/{application.id}/guilds/{guild.id}/commands
type GetGuildApplicationCommandsQuery struct {
	// Whether to include full localization dictionaries instead of the localized fields
	WithLocalizations *bool `json:"with_localizations,omitempty"`
}

// GetGuildApplicationCommandsResult represents the response from GET /applications/{application.id}/guilds/{guild.id}/commands
type GetGuildApplicationCommandsResult = []ApplicationCommand

// PostGuildApplicationCommandJSONBody represents the request body for POST /applications/{application.id}/guilds/{guild.id}/commands
type PostGuildApplicationCommandJSONBody = PostApplicationCommandJSONBody

// PostGuildApplicationCommandResult represents the response from POST /applications/{application.id}/guilds/{guild.id}/commands
type PostGuildApplicationCommandResult = ApplicationCommand

// GetGuildApplicationCommandResult represents the response from GET /applications/{application.id}/guilds/{guild.id}/commands/{command.id}
type GetGuildApplicationCommandResult = ApplicationCommand

// PatchGuildApplicationCommandJSONBody represents the request body for PATCH /applications/{application.id}/guilds/{guild.id}/commands/{command.id}
type PatchGuildApplicationCommandJSONBody = PatchApplicationCommandJSONBody

// PatchGuildApplicationCommandResult represents the response from PATCH /applications/{application.id}/guilds/{guild.id}/commands/{command.id}
type PatchGuildApplicationCommandResult = ApplicationCommand

// DeleteGuildApplicationCommandResult represents the response from DELETE /applications/{application.id}/guilds/{guild.id}/commands/{command.id}
type DeleteGuildApplicationCommandResult struct{} // Empty response

// PutGuildApplicationCommandsJSONBody represents the request body for PUT /applications/{application.id}/guilds/{guild.id}/commands
type PutGuildApplicationCommandsJSONBody = []PostGuildApplicationCommandJSONBody

// PutGuildApplicationCommandsResult represents the response from PUT /applications/{application.id}/guilds/{guild.id}/commands
type PutGuildApplicationCommandsResult = []ApplicationCommand

// GetGuildApplicationCommandPermissionsResult represents the response from GET /applications/{application.id}/guilds/{guild.id}/commands/permissions
type GetGuildApplicationCommandPermissionsResult = []ApplicationCommandGuildPermissions

// GetApplicationCommandPermissionsResult represents the response from GET /applications/{application.id}/guilds/{guild.id}/commands/{command.id}/permissions
type GetApplicationCommandPermissionsResult = ApplicationCommandGuildPermissions

// PutApplicationCommandPermissionsJSONBody represents the request body for PUT /applications/{application.id}/guilds/{guild.id}/commands/{command.id}/permissions
type PutApplicationCommandPermissionsJSONBody struct {
	// Permissions for the command in the guild
	Permissions []ApplicationCommandPermission `json:"permissions"`
}

// PutApplicationCommandPermissionsResult represents the response from PUT /applications/{application.id}/guilds/{guild.id}/commands/{command.id}/permissions
type PutApplicationCommandPermissionsResult = ApplicationCommandGuildPermissions

// ApplicationCommandGuildPermissions represents guild permissions for an application command
type ApplicationCommandGuildPermissions struct {
	// ID of the command or the application ID if this is global permissions
	ID discord.Snowflake `json:"id"`
	// ID of the application the command belongs to
	ApplicationID discord.Snowflake `json:"application_id"`
	// ID of the guild
	GuildID discord.Snowflake `json:"guild_id"`
	// Permissions for the command in the guild, max 100
	Permissions []ApplicationCommandPermission `json:"permissions"`
}

// ApplicationCommandPermission represents permissions for an application command
type ApplicationCommandPermission struct {
	// ID of the role, user, or channel. It can also be a permission constant
	ID discord.Snowflake `json:"id"`
	// role, user, or channel
	Type ApplicationCommandPermissionType `json:"type"`
	// true to allow, false to disallow
	Permission bool `json:"permission"`
}

// ApplicationCommandPermissionType represents the type of permission
type ApplicationCommandPermissionType int

const (
	ApplicationCommandPermissionTypeRole    ApplicationCommandPermissionType = 1
	ApplicationCommandPermissionTypeUser    ApplicationCommandPermissionType = 2
	ApplicationCommandPermissionTypeChannel ApplicationCommandPermissionType = 3
)

// PostInteractionCallbackJSONBody represents the request body for POST /interactions/{interaction.id}/{interaction.token}/callback
type PostInteractionCallbackJSONBody struct {
	// The type of response
	Type payloads.InteractionResponseType `json:"type"`
	// An optional response message
	Data *InteractionCallbackData `json:"data,omitempty"`
}

// InteractionCallbackData represents the data field in an interaction callback
type InteractionCallbackData struct {
	// Is the response TTS
	TTS *bool `json:"tts,omitempty"`
	// Message content
	Content *string `json:"content,omitempty"`
	// Up to 10 embed objects
	Embeds []payloads.Embed `json:"embeds,omitempty"`
	// Allowed mentions object
	AllowedMentions *payloads.AllowedMentions `json:"allowed_mentions,omitempty"`
	// Message flags combined as a bitfield (only SUPPRESS_EMBEDS and EPHEMERAL can be set)
	Flags *payloads.MessageFlags `json:"flags,omitempty"`
	// Message components to include (up to 5 action rows)
	Components []payloads.ActionRowComponent `json:"components,omitempty"`
	// Attachment objects with filename and description
	Attachments []payloads.PartialAttachment `json:"attachments,omitempty"`
	// Poll object
	Poll *payloads.Poll `json:"poll,omitempty"`
	// Autocomplete choices (max 25 choices)
	Choices []ApplicationCommandOptionChoice `json:"choices,omitempty"`
	// For components, a developer-defined identifier for the component, max 100 characters
	CustomID *string `json:"custom_id,omitempty"`
	// For components, one of text, paragraph
	Style *payloads.TextInputStyle `json:"style,omitempty"`
	// For components, the label for this component, max 45 characters
	Label *string `json:"label,omitempty"`
	// For components, the minimum allowed length of the text, minimum 0, maximum 4000
	MinLength *int `json:"min_length,omitempty"`
	// For components, the maximum allowed length of the text, minimum 1, maximum 4000
	MaxLength *int `json:"max_length,omitempty"`
	// For components, whether this component is required to be filled, default true
	Required *bool `json:"required,omitempty"`
	// For components, a pre-filled value for this component, max 4000 characters
	Value *string `json:"value,omitempty"`
	// For components, placeholder text if nothing is selected, max 150 characters
	Placeholder *string `json:"placeholder,omitempty"`
	// For components, the minimum number of items that must be chosen, default 1, minimum 0, maximum 25
	MinValues *int `json:"min_values,omitempty"`
	// For components, the maximum number of items that can be chosen, default 1, maximum 25
	MaxValues *int `json:"max_values,omitempty"`
	// For components, choice options for select menus, max 25
	Options []payloads.SelectOption `json:"options,omitempty"`
	// For components, list of channel types to include in the channel select component
	ChannelTypes []payloads.ChannelType `json:"channel_types,omitempty"`
	// For components, list of default roles for role select component
	DefaultRoles []discord.Snowflake `json:"default_roles,omitempty"`
	// For components, list of default users for user select component
	DefaultUsers []discord.Snowflake `json:"default_users,omitempty"`
	// For components, list of default channels for channel select component
	DefaultChannels []discord.Snowflake `json:"default_channels,omitempty"`
}

// ApplicationCommandOptionChoice represents a choice for a command option
type ApplicationCommandOptionChoice struct {
	// 1-100 character choice name
	Name string `json:"name"`
	// Localization dictionary for the name field
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	// Value for the choice, up to 100 characters if string
	Value interface{} `json:"value"`
}

// PostInteractionCallbackResult represents the response from POST /interactions/{interaction.id}/{interaction.token}/callback
type PostInteractionCallbackResult struct{} // Empty response

// InteractionCallbackWithResponseResult represents the response with callback details
type InteractionCallbackWithResponseResult struct {
	// The interaction object associated with the interaction
	Interaction InteractionCallbackObject `json:"interaction"`
	// The resource that was created by the interaction response
	Resource *InteractionCallbackResourceObject `json:"resource,omitempty"`
}

// InteractionCallbackObject represents the interaction callback object
type InteractionCallbackObject struct {
	// ID of the interaction
	ID discord.Snowflake `json:"id"`
	// Interaction type
	Type payloads.InteractionType `json:"type"`
	// Instance ID of the Activity if one was launched or joined
	ActivityInstanceID *string `json:"activity_instance_id,omitempty"`
	// ID of the message that was created by the interaction
	ResponseMessageID *discord.Snowflake `json:"response_message_id,omitempty"`
	// Whether or not the message is in a loading state
	ResponseMessageLoading *bool `json:"response_message_loading,omitempty"`
	// Whether or not the response message was ephemeral
	ResponseMessageEphemeral *bool `json:"response_message_ephemeral,omitempty"`
}

// InteractionCallbackResourceObject represents the callback resource object
type InteractionCallbackResourceObject struct {
	// Interaction callback type
	Type payloads.InteractionResponseType `json:"type"`
	// Represents the Activity launched by this interaction
	ActivityInstance *InteractionCallbackActivityInstanceResource `json:"activity_instance,omitempty"`
	// Message created by the interaction
	Message *payloads.Message `json:"message,omitempty"`
}

// InteractionCallbackActivityInstanceResource represents an activity instance resource
type InteractionCallbackActivityInstanceResource struct {
	// Instance ID of the Activity if one was launched or joined
	ID string `json:"id"`
}

// GetInteractionOriginalResponseResult represents the response from GET /webhooks/{application.id}/{interaction.token}/messages/@original
type GetInteractionOriginalResponseResult = payloads.Message

// PatchInteractionOriginalResponseJSONBody represents the request body for PATCH /webhooks/{application.id}/{interaction.token}/messages/@original
type PatchInteractionOriginalResponseJSONBody struct {
	// Message contents (up to 2000 characters)
	Content *string `json:"content,omitempty"`
	// Up to 10 embed objects
	Embeds []payloads.Embed `json:"embeds,omitempty"`
	// Allowed mentions for the message
	AllowedMentions *payloads.AllowedMentions `json:"allowed_mentions,omitempty"`
	// Message components to include (up to 5 action rows)
	Components []payloads.ActionRowComponent `json:"components,omitempty"`
	// Attachment objects with filename and description
	Attachments []payloads.PartialAttachment `json:"attachments,omitempty"`
}

// PatchInteractionOriginalResponseResult represents the response from PATCH /webhooks/{application.id}/{interaction.token}/messages/@original
type PatchInteractionOriginalResponseResult = payloads.Message

// DeleteInteractionOriginalResponseResult represents the response from DELETE /webhooks/{application.id}/{interaction.token}/messages/@original
type DeleteInteractionOriginalResponseResult struct{} // Empty response

// PostInteractionFollowupJSONBody represents the request body for POST /webhooks/{application.id}/{interaction.token}
type PostInteractionFollowupJSONBody = InteractionCallbackData

// PostInteractionFollowupResult represents the response from POST /webhooks/{application.id}/{interaction.token}
type PostInteractionFollowupResult = payloads.Message

// GetInteractionFollowupResult represents the response from GET /webhooks/{application.id}/{interaction.token}/messages/{message.id}
type GetInteractionFollowupResult = payloads.Message

// PatchInteractionFollowupJSONBody represents the request body for PATCH /webhooks/{application.id}/{interaction.token}/messages/{message.id}
type PatchInteractionFollowupJSONBody = PatchInteractionOriginalResponseJSONBody

// PatchInteractionFollowupResult represents the response from PATCH /webhooks/{application.id}/{interaction.token}/messages/{message.id}
type PatchInteractionFollowupResult = payloads.Message

// DeleteInteractionFollowupResult represents the response from DELETE /webhooks/{application.id}/{interaction.token}/messages/{message.id}
type DeleteInteractionFollowupResult struct{} // Empty response
