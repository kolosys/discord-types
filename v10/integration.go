package v10

import "time"

// Integration represents a Discord integration.
// For bot integrations, use BotIntegration instead.
// Reference: https://discord.com/developers/docs/resources/guild#integration-object
type Integration struct {
	ID                Snowflake          `json:"id"`                  // Integration ID
	Name              string             `json:"name"`                // Integration name
	Type              string             `json:"type"`                // Integration type (twitch, youtube, discord, or guild_subscription)
	Enabled           bool               `json:"enabled"`             // Whether the integration is enabled
	Syncing           bool               `json:"syncing"`             // Whether the integration is syncing
	RoleID            Snowflake          `json:"role_id"`             // Role ID that the integration uses for server members
	EnableEmoticons   bool               `json:"enable_emoticons"`    // Whether emoticons are enabled
	ExpireBehavior    int                `json:"expire_behavior"`     // Whether emoticons should be synced for this integration (twitch only currently)
	ExpireGracePeriod int                `json:"expire_grace_period"` // Grace period in days for revoked subscribers
	User              User               `json:"user"`                // User who created the integration
	Account           IntegrationAccount `json:"account"`             // Integration account information
	SyncedAt          time.Time          `json:"synced_at"`           // When this integration was last synced
	SubscriberCount   int                `json:"subscriber_count"`    // How many subscribers this integration has
	Revoked           bool               `json:"revoked"`             // Has this integration been revoked
	Application       Application        `json:"application"`         // The bot/OAuth2 application for discord integrations
	Scopes            []OAuth2Scope      `json:"scopes"`              // The scopes the application has been authorized for
}

type IntegrationAccount struct {
	ID   Snowflake `json:"id"`   // Integration account ID
	Name string    `json:"name"` // Integration account name
}

// BotIntegration represents a Discord bot integration.
// Reference: https://discord.com/developers/docs/resources/guild#bot-integration-object
type BotIntegration struct {
	ID          Snowflake          `json:"id"`          // Integration ID
	Name        string             `json:"name"`        // Integration name
	Type        string             `json:"type"`        // Integration type (twitch, youtube, discord, or guild_subscription)
	User        User               `json:"user"`        // User who created the integration
	Account     IntegrationAccount `json:"account"`     // Integration account information
	Enabled     bool               `json:"enabled"`     // Whether the integration is enabled
	Application Application        `json:"application"` // The bot/OAuth2 application for discord integrations
	Scopes      []OAuth2Scope      `json:"scopes"`      // The scopes the application has been authorized for
}
