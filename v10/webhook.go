package v10

import (
	"time"
)

// Webhook represents a Discord webhook
// https://discord.com/developers/docs/resources/webhook#webhook-object
type Webhook struct {
	ID            Snowflake       `json:"id"`                       // The id of the webhook
	Type          WebhookType     `json:"type"`                     // The type of webhook
	GuildID       *Snowflake      `json:"guild_id,omitempty"`       // The guild id this webhook is for, if any
	ChannelID     Snowflake       `json:"channel_id"`               // The channel id this webhook is for, if any
	User          *User           `json:"user,omitempty"`           // The user this webhook was created by (not returned when getting a webhook with its token)
	Name          *string         `json:"name"`                     // The default name of the webhook
	Avatar        *string         `json:"avatar"`                   // The default user avatar hash of the webhook
	Token         *string         `json:"token,omitempty"`          // The secure token of the webhook (returned for Incoming Webhooks)
	ApplicationID *Snowflake      `json:"application_id"`           // The bot/OAuth2 application that created this webhook
	SourceGuild   *PartialGuild   `json:"source_guild,omitempty"`   // Will be absent if the webhook creator has since lost access to the guild where the followed channel resides
	SourceChannel *PartialChannel `json:"source_channel,omitempty"` // Will be absent if the webhook creator has since lost access to the guild where the followed channel resides
	URL           *string         `json:"url,omitempty"`            // The url used for executing the webhook (returned by the webhooks OAuth2 flow)
}

// WebhookType represents the type of webhook
// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-types
type WebhookType int

const (
	WebhookTypeIncoming WebhookType = iota + 1
	WebhookTypeChannelFollower
	WebhookTypeApplication
)

// WebhookEventBase represents the base structure for webhook events
// https://discord.com/developers/docs/events/webhook-events#webhook-event-payloads
type WebhookEventPayload[T any] struct {
	Version       int                    `json:"version"` // Always 1
	ApplicationID Snowflake              `json:"application_id"`
	Type          ApplicationWebhookType `json:"type"`
	Event         WebhookEventBody[T]    `json:"event,omitempty"`
}

// ApplicationWebhookType represents the type of application webhook
// https://discord.com/developers/docs/events/webhook-events#webhook-types
type ApplicationWebhookType int

const (
	ApplicationWebhookTypePing ApplicationWebhookType = iota
	ApplicationWebhookTypeEvent
)

// WebhookEventBody represents the body of a webhook event
type WebhookEventBody[T any] struct {
	Type      ApplicationWebhookEventType `json:"type"`
	Timestamp time.Time                   `json:"timestamp"`
	Data      T                           `json:"data,omitempty"`
}

// WebhookApplicationAuthorizedEvent represents an application authorized event
type WebhookApplicationAuthorizedEvent = WebhookEventBody[WebhookEventApplicationAuthorizedData]

// WebhookEventApplicationAuthorizedData represents the data for application authorized events
type WebhookEventApplicationAuthorizedData struct {
	IntegrationType *ApplicationIntegrationType `json:"integration_type,omitempty"`
	User            User                        `json:"user"`
	Scopes          []OAuth2Scope               `json:"scopes"`
	Guild           any                         `json:"guild,omitempty"` // Would be *Guild but avoiding circular dependency
}

// WebhookApplicationDeauthorizedEvent represents an application deauthorized event
type WebhookApplicationDeauthorizedEvent = WebhookEventBody[WebhookEventApplicationDeauthorizedData]

// WebhookEventApplicationDeauthorizedData represents the data for application deauthorized events
type WebhookEventApplicationDeauthorizedData struct {
	User User `json:"user"`
}

// WebhookEventEntitlementCreate represents an entitlement create event
type WebhookEntitlementCreateEvent = WebhookEventBody[WebhookEventEntitlementCreateData]

// WebhookEventEntitlementCreateData represents the data for entitlement create events
type WebhookEventEntitlementCreateData struct {
	Entitlement Entitlement `json:"entitlement"`
}
