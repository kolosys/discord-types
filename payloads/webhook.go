package payloads

import (
	"time"

	"github.com/kolosys/discord-types/discord"
)

// WebhookType represents the type of webhook
// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-types
type WebhookType int

const (
	WebhookTypeIncoming WebhookType = iota + 1
	WebhookTypeChannelFollower
	WebhookTypeApplication
)

// Webhook represents a Discord webhook
// https://discord.com/developers/docs/resources/webhook#webhook-object
type Webhook struct {
	ID            discord.Snowflake     `json:"id"`
	Type          WebhookType           `json:"type"`
	GuildID       *discord.Snowflake    `json:"guild_id,omitempty"`
	ChannelID     discord.Snowflake     `json:"channel_id"`
	User          *User                 `json:"user,omitempty"`
	Name          *string               `json:"name"`
	Avatar        *string               `json:"avatar"`
	Token         *string               `json:"token,omitempty"`
	ApplicationID *discord.Snowflake    `json:"application_id"`
	SourceGuild   *WebhookSourceGuild   `json:"source_guild,omitempty"`
	SourceChannel *WebhookSourceChannel `json:"source_channel,omitempty"`
	URL           *string               `json:"url,omitempty"`
}

// WebhookSourceGuild represents the source guild for channel follower webhooks
type WebhookSourceGuild struct {
	ID   discord.Snowflake `json:"id"`
	Name string            `json:"name"`
	Icon *string           `json:"icon"`
}

// ApplicationWebhookType represents the type of application webhook
// https://discord.com/developers/docs/events/webhook-events#webhook-types
type ApplicationWebhookType int

const (
	ApplicationWebhookTypePing ApplicationWebhookType = iota
	ApplicationWebhookTypeEvent
)

// ApplicationWebhookEventType is defined in shared.go

// WebhookEvent represents a webhook event
// https://discord.com/developers/docs/events/webhook-events#webhook-event-payloads
type WebhookEvent interface {
	GetType() ApplicationWebhookType
	GetApplicationID() discord.Snowflake
	GetVersion() int
}

// WebhookEventBase represents the base structure for webhook events
type WebhookEventBase struct {
	Version       int                    `json:"version"` // Always 1
	ApplicationID discord.Snowflake      `json:"application_id"`
	Type          ApplicationWebhookType `json:"type"`
}

func (w WebhookEventBase) GetType() ApplicationWebhookType {
	return w.Type
}

func (w WebhookEventBase) GetApplicationID() discord.Snowflake {
	return w.ApplicationID
}

func (w WebhookEventBase) GetVersion() int {
	return w.Version
}

// WebhookPingEvent represents a ping webhook event
type WebhookPingEvent struct {
	WebhookEventBase
}

// WebhookEventEvent represents an event webhook event
type WebhookEventEvent struct {
	WebhookEventBase
	Event WebhookEventBody `json:"event"`
}

// WebhookEventBody represents the body of a webhook event
type WebhookEventBody interface {
	GetEventType() ApplicationWebhookEventType
	GetTimestamp() time.Time
}

// WebhookEventEventBase represents the base structure for webhook event bodies
type WebhookEventEventBase struct {
	Type      ApplicationWebhookEventType `json:"type"`
	Timestamp time.Time                   `json:"timestamp"`
}

func (w WebhookEventEventBase) GetEventType() ApplicationWebhookEventType {
	return w.Type
}

func (w WebhookEventEventBase) GetTimestamp() time.Time {
	return w.Timestamp
}

// WebhookEventApplicationAuthorized represents an application authorized event
type WebhookEventApplicationAuthorized struct {
	WebhookEventEventBase
	Data WebhookEventApplicationAuthorizedData `json:"data"`
}

// WebhookEventApplicationAuthorizedData represents the data for application authorized events
type WebhookEventApplicationAuthorizedData struct {
	IntegrationType *ApplicationIntegrationType `json:"integration_type,omitempty"`
	User            User                        `json:"user"`
	Scopes          []OAuth2Scope               `json:"scopes"`
	Guild           interface{}                 `json:"guild,omitempty"` // Would be *Guild but avoiding circular dependency
}

// WebhookEventApplicationDeauthorized represents an application deauthorized event
type WebhookEventApplicationDeauthorized struct {
	WebhookEventEventBase
	Data WebhookEventApplicationDeauthorizedData `json:"data"`
}

// WebhookEventApplicationDeauthorizedData represents the data for application deauthorized events
type WebhookEventApplicationDeauthorizedData struct {
	User User `json:"user"`
}

// WebhookEventEntitlementCreate represents an entitlement create event
type WebhookEventEntitlementCreate struct {
	WebhookEventEventBase
	Data Entitlement `json:"data"`
}

// WebhookEventQuestUserEnrollment represents a quest user enrollment event
type WebhookEventQuestUserEnrollment struct {
	WebhookEventEventBase
	Data interface{} `json:"data"` // Currently no data
}

// Forward declarations for types defined in other files
// Guild is defined in guild.go but causes circular dependency, using interface{} for now
// OAuth2Scope is defined in shared.go
