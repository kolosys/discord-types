package payloads

import "github.com/kolosys/discord-types/discord"

// Application represents a Discord application
// https://discord.com/developers/docs/resources/application#application-object
type Application struct {
	ID                                discord.Snowflake                     `json:"id"`
	Name                              string                                `json:"name"`
	Icon                              *string                               `json:"icon"`
	Description                       string                                `json:"description"`
	RPCOrigins                        []string                              `json:"rpc_origins,omitempty"`
	BotPublic                         bool                                  `json:"bot_public"`
	BotRequireCodeGrant               bool                                  `json:"bot_require_code_grant"`
	Bot                               *User                                 `json:"bot,omitempty"`
	TermsOfServiceURL                 *string                               `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyURL                  *string                               `json:"privacy_policy_url,omitempty"`
	Owner                             *User                                 `json:"owner,omitempty"`
	Summary                           string                                `json:"summary"` // Deprecated in v11, always empty string
	VerifyKey                         string                                `json:"verify_key"`
	Team                              *Team                                 `json:"team"`
	GuildID                           *discord.Snowflake                    `json:"guild_id,omitempty"`
	Guild                             *PartialGuild                         `json:"guild,omitempty"`
	PrimarySKUID                      *discord.Snowflake                    `json:"primary_sku_id,omitempty"`
	Slug                              *string                               `json:"slug,omitempty"`
	CoverImage                        *string                               `json:"cover_image,omitempty"`
	Flags                             ApplicationFlags                      `json:"flags"`
	ApproximateGuildCount             *int                                  `json:"approximate_guild_count,omitempty"`
	ApproximateUserInstallCount       *int                                  `json:"approximate_user_install_count,omitempty"`
	ApproximateUserAuthorizationCount *int                                  `json:"approximate_user_authorization_count,omitempty"`
	RedirectURIs                      []string                              `json:"redirect_uris,omitempty"`
	InteractionsEndpointURL           *string                               `json:"interactions_endpoint_url,omitempty"`
	RoleConnectionsVerificationURL    *string                               `json:"role_connections_verification_url,omitempty"`
	Tags                              []string                              `json:"tags,omitempty"` // Up to 5 tags, max 20 chars each
	InstallParams                     *ApplicationInstallParams             `json:"install_params,omitempty"`
	IntegrationTypesConfig            *ApplicationIntegrationTypesConfigMap `json:"integration_types_config,omitempty"`
	CustomInstallURL                  *string                               `json:"custom_install_url,omitempty"`
	EventWebhooksURL                  *string                               `json:"event_webhooks_url,omitempty"`
	EventWebhooksStatus               ApplicationWebhookEventStatus         `json:"event_webhooks_status"`
	EventWebhooksTypes                []ApplicationWebhookEventType         `json:"event_webhooks_types,omitempty"`
}

// ApplicationInstallParams represents installation parameters for an application
type ApplicationInstallParams struct {
	Scopes      []OAuth2Scope       `json:"scopes"`
	Permissions discord.Permissions `json:"permissions"`
}

// ApplicationIntegrationTypeConfiguration represents configuration for an integration type
type ApplicationIntegrationTypeConfiguration struct {
	OAuth2InstallParams *ApplicationInstallParams `json:"oauth2_install_params,omitempty"`
}

// ApplicationIntegrationTypesConfigMap maps integration types to their configurations
type ApplicationIntegrationTypesConfigMap map[ApplicationIntegrationType]*ApplicationIntegrationTypeConfiguration

// ApplicationFlags represents the application flags bitfield
// https://discord.com/developers/docs/resources/application#application-object-application-flags
type ApplicationFlags int

const (
	ApplicationFlagEmbeddedReleased              ApplicationFlags = 1 << 1
	ApplicationFlagManagedEmoji                  ApplicationFlags = 1 << 2
	ApplicationFlagEmbeddedIAP                   ApplicationFlags = 1 << 3
	ApplicationFlagGroupDMCreate                 ApplicationFlags = 1 << 4
	ApplicationFlagAutoModerationRuleCreateBadge ApplicationFlags = 1 << 6
	ApplicationFlagRPCHasConnected               ApplicationFlags = 1 << 11
	ApplicationFlagGatewayPresence               ApplicationFlags = 1 << 12
	ApplicationFlagGatewayPresenceLimited        ApplicationFlags = 1 << 13
	ApplicationFlagGatewayGuildMembers           ApplicationFlags = 1 << 14
	ApplicationFlagGatewayGuildMembersLimited    ApplicationFlags = 1 << 15
	ApplicationFlagVerificationPendingGuildLimit ApplicationFlags = 1 << 16
	ApplicationFlagEmbedded                      ApplicationFlags = 1 << 17
	ApplicationFlagGatewayMessageContent         ApplicationFlags = 1 << 18
	ApplicationFlagGatewayMessageContentLimited  ApplicationFlags = 1 << 19
	ApplicationFlagEmbeddedFirstParty            ApplicationFlags = 1 << 20
	ApplicationFlagApplicationCommandBadge       ApplicationFlags = 1 << 23
)

// ApplicationRoleConnectionMetadata represents metadata for application role connections
// https://discord.com/developers/docs/resources/application-role-connection-metadata#application-role-connection-metadata-object-application-role-connection-metadata-structure
type ApplicationRoleConnectionMetadata struct {
	Type                     ApplicationRoleConnectionMetadataType `json:"type"`
	Key                      string                                `json:"key"`
	Name                     string                                `json:"name"`
	NameLocalizations        LocalizationMap                       `json:"name_localizations,omitempty"`
	Description              string                                `json:"description"`
	DescriptionLocalizations LocalizationMap                       `json:"description_localizations,omitempty"`
}

// ApplicationRoleConnectionMetadataType represents the type of role connection metadata
// https://discord.com/developers/docs/resources/application-role-connection-metadata#application-role-connection-metadata-object-application-role-connection-metadata-type
type ApplicationRoleConnectionMetadataType int

const (
	ApplicationRoleConnectionMetadataTypeIntegerLessThanOrEqual ApplicationRoleConnectionMetadataType = iota + 1
	ApplicationRoleConnectionMetadataTypeIntegerGreaterThanOrEqual
	ApplicationRoleConnectionMetadataTypeIntegerEqual
	ApplicationRoleConnectionMetadataTypeIntegerNotEqual
	ApplicationRoleConnectionMetadataTypeDatetimeLessThanOrEqual
	ApplicationRoleConnectionMetadataTypeDatetimeGreaterThanOrEqual
	ApplicationRoleConnectionMetadataTypeBooleanEqual
	ApplicationRoleConnectionMetadataTypeBooleanNotEqual
)

// ApplicationWebhookEventStatus represents the status of application webhook events
// https://discord.com/developers/docs/resources/application#application-object-application-event-webhook-status
type ApplicationWebhookEventStatus int

const (
	ApplicationWebhookEventStatusDisabled ApplicationWebhookEventStatus = iota + 1
	ApplicationWebhookEventStatusEnabled
	ApplicationWebhookEventStatusDisabledByDiscord
)

// Forward declarations for types that will be defined in other files
// Team is defined in teams.go
