package payloads

// Shared types used across multiple payload files

// OAuth2Scope is defined in oauth2.go

// ApplicationIntegrationType represents the integration type for applications
type ApplicationIntegrationType int

const (
	ApplicationIntegrationTypeGuildInstall ApplicationIntegrationType = iota
	ApplicationIntegrationTypeUserInstall
)

// ApplicationWebhookEventType represents the type of application webhook event
type ApplicationWebhookEventType string

const (
	ApplicationWebhookEventTypeApplicationAuthorized   ApplicationWebhookEventType = "APPLICATION_AUTHORIZED"
	ApplicationWebhookEventTypeApplicationDeauthorized ApplicationWebhookEventType = "APPLICATION_DEAUTHORIZED"
	ApplicationWebhookEventTypeEntitlementCreate       ApplicationWebhookEventType = "ENTITLEMENT_CREATE"
	ApplicationWebhookEventTypeQuestUserEnrollment     ApplicationWebhookEventType = "QUEST_USER_ENROLLMENT"
)
