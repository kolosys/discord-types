package v10

// AuditLog represents a Discord audit log.
// Reference: https://discord.com/developers/docs/resources/audit-log#audit-log-object-audit-log-structure
type AuditLog struct {
	// ApplicationCommands are the list of application commands found in the audit log.
	ApplicationCommands []ApplicationCommandInteraction `json:"application_commands"` // NOTE: ApplicationCommand type to be defined in future release

	// AuditLogEntries are the list of audit log entries, sorted from most to least recent.
	AuditLogEntries []AuditLogEntry `json:"audit_log_entries"`

	// AutoModerationRules are the list of auto moderation rules found in the audit log.
	AutoModerationRules []AutoModerationRule `json:"auto_moderation_rules"`

	// GuildScheduledEvents are the list of guild scheduled events found in the audit log.
	GuildScheduledEvents []GuildScheduledEvent `json:"guild_scheduled_events"`

	// Integrations are the list of partial integration objects.
	Integrations []Integration `json:"integrations"` // NOTE: Integration type to be defined in future release

	// Threads are the list of threads found in the audit log.
	Threads []ThreadChannel `json:"threads"`

	// Users are the list of users found in the audit log.
	Users []User `json:"users"`

	// Webhooks are the webhooks found in the audit log.
	Webhooks []Webhook `json:"webhooks"`
}

// AuditLogEntry represents an audit log entry.
// Reference: https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-entry-structure
type AuditLogEntry struct {
	// TargetID is the id of the affected entity (webhook, user, role, etc.).
	TargetID *string `json:"target_id"`

	// Changes are the changes made to the target_id.
	Changes []AuditLogChange `json:"changes,omitempty"`

	// UserID is the user or app that made the changes.
	UserID *Snowflake `json:"user_id"`

	// ID is the id of the entry.
	ID Snowflake `json:"id"`

	// ActionType is the type of action that occurred.
	ActionType AuditLogEvent `json:"action_type"`

	// Options are the additional info for certain action types.
	Options *AuditLogEntryInfo `json:"options,omitempty"`

	// Reason is the reason for the change (1-512 characters).
	Reason *string `json:"reason,omitempty"`
}

// AuditLogEvent represents audit log events.
// Reference: https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-events
type AuditLogEvent int

// Audit log event constants
const (
	// Guild events
	AuditLogEventGuildUpdate AuditLogEvent = 1

	// Channel events
	AuditLogEventChannelCreate          AuditLogEvent = 10
	AuditLogEventChannelUpdate          AuditLogEvent = 11
	AuditLogEventChannelDelete          AuditLogEvent = 12
	AuditLogEventChannelOverwriteCreate AuditLogEvent = 13
	AuditLogEventChannelOverwriteUpdate AuditLogEvent = 14
	AuditLogEventChannelOverwriteDelete AuditLogEvent = 15

	// Member events
	AuditLogEventMemberKick       AuditLogEvent = 20
	AuditLogEventMemberPrune      AuditLogEvent = 21
	AuditLogEventMemberBanAdd     AuditLogEvent = 22
	AuditLogEventMemberBanRemove  AuditLogEvent = 23
	AuditLogEventMemberUpdate     AuditLogEvent = 24
	AuditLogEventMemberRoleUpdate AuditLogEvent = 25
	AuditLogEventMemberMove       AuditLogEvent = 26
	AuditLogEventMemberDisconnect AuditLogEvent = 27
	AuditLogEventBotAdd           AuditLogEvent = 28

	// Role events
	AuditLogEventRoleCreate AuditLogEvent = 30
	AuditLogEventRoleUpdate AuditLogEvent = 31
	AuditLogEventRoleDelete AuditLogEvent = 32

	// Invite events
	AuditLogEventInviteCreate AuditLogEvent = 40
	AuditLogEventInviteUpdate AuditLogEvent = 41
	AuditLogEventInviteDelete AuditLogEvent = 42

	// Webhook events
	AuditLogEventWebhookCreate AuditLogEvent = 50
	AuditLogEventWebhookUpdate AuditLogEvent = 51
	AuditLogEventWebhookDelete AuditLogEvent = 52

	// Emoji events
	AuditLogEventEmojiCreate AuditLogEvent = 60
	AuditLogEventEmojiUpdate AuditLogEvent = 61
	AuditLogEventEmojiDelete AuditLogEvent = 62

	// Message events
	AuditLogEventMessageDelete     AuditLogEvent = 72
	AuditLogEventMessageBulkDelete AuditLogEvent = 73
	AuditLogEventMessagePin        AuditLogEvent = 74
	AuditLogEventMessageUnpin      AuditLogEvent = 75

	// Integration events
	AuditLogEventIntegrationCreate AuditLogEvent = 80
	AuditLogEventIntegrationUpdate AuditLogEvent = 81
	AuditLogEventIntegrationDelete AuditLogEvent = 82

	// Stage instance events
	AuditLogEventStageInstanceCreate AuditLogEvent = 83
	AuditLogEventStageInstanceUpdate AuditLogEvent = 84
	AuditLogEventStageInstanceDelete AuditLogEvent = 85

	// Sticker events
	AuditLogEventStickerCreate AuditLogEvent = 90
	AuditLogEventStickerUpdate AuditLogEvent = 91
	AuditLogEventStickerDelete AuditLogEvent = 92

	// Guild scheduled event events
	AuditLogEventGuildScheduledEventCreate AuditLogEvent = 100
	AuditLogEventGuildScheduledEventUpdate AuditLogEvent = 101
	AuditLogEventGuildScheduledEventDelete AuditLogEvent = 102

	// Thread events
	AuditLogEventThreadCreate AuditLogEvent = 110
	AuditLogEventThreadUpdate AuditLogEvent = 111
	AuditLogEventThreadDelete AuditLogEvent = 112

	// Application command permission events
	AuditLogEventApplicationCommandPermissionUpdate AuditLogEvent = 121

	// Auto moderation events
	AuditLogEventAutoModerationRuleCreate                AuditLogEvent = 140
	AuditLogEventAutoModerationRuleUpdate                AuditLogEvent = 141
	AuditLogEventAutoModerationRuleDelete                AuditLogEvent = 142
	AuditLogEventAutoModerationBlockMessage              AuditLogEvent = 143
	AuditLogEventAutoModerationFlagToChannel             AuditLogEvent = 144
	AuditLogEventAutoModerationUserCommunicationDisabled AuditLogEvent = 145

	// Creator monetization events
	AuditLogEventCreatorMonetizationRequestCreated AuditLogEvent = 150
	AuditLogEventCreatorMonetizationTermsAccepted  AuditLogEvent = 151

	// Onboarding events
	AuditLogEventOnboardingPromptCreate AuditLogEvent = 163
	AuditLogEventOnboardingPromptUpdate AuditLogEvent = 164
	AuditLogEventOnboardingPromptDelete AuditLogEvent = 165
	AuditLogEventOnboardingCreate       AuditLogEvent = 166
	AuditLogEventOnboardingUpdate       AuditLogEvent = 167

	// Guild home events
	AuditLogEventHomeSettingsCreate AuditLogEvent = 190
	AuditLogEventHomeSettingsUpdate AuditLogEvent = 191
)

// AuditLogChange represents a change in an audit log entry.
// Reference: https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-structure
type AuditLogChange struct {
	// NewValue is the new value of the key.
	NewValue any `json:"new_value,omitempty"`

	// OldValue is the old value of the key.
	OldValue any `json:"old_value,omitempty"`

	// Key is the name of the audit log change key.
	Key string `json:"key"`
}

// AuditLogEntryInfo represents optional audit entry information.
// Reference: https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info
type AuditLogEntryInfo struct {
	// ApplicationID is the id of the app whose permissions were targeted.
	ApplicationID *Snowflake `json:"application_id,omitempty"`

	// AutoModerationRuleName is the name of the Auto Moderation rule that was triggered.
	AutoModerationRuleName *string `json:"auto_moderation_rule_name,omitempty"`

	// AutoModerationRuleTriggerType is the trigger type of the Auto Moderation rule that was triggered.
	AutoModerationRuleTriggerType *string `json:"auto_moderation_rule_trigger_type,omitempty"`

	// ChannelID is the id of the channel in which the entities were targeted.
	ChannelID *Snowflake `json:"channel_id,omitempty"`

	// Count is the number of entities that were targeted.
	Count *string `json:"count,omitempty"`

	// DeleteMemberDays is the number of days after which inactive members were kicked.
	DeleteMemberDays *string `json:"delete_member_days,omitempty"`

	// ID is the id of the overwritten entity.
	ID *Snowflake `json:"id,omitempty"`

	// MembersRemoved is the number of members removed by the prune.
	MembersRemoved *string `json:"members_removed,omitempty"`

	// MessageID is the id of the message that was targeted.
	MessageID *Snowflake `json:"message_id,omitempty"`

	// RoleName is the name of the role if type is "0" (not present if type is "1").
	RoleName *string `json:"role_name,omitempty"`

	// Type is the type of overwritten entity - role ("0") or member ("1").
	Type *string `json:"type,omitempty"`

	// IntegrationAccessedScopes are the scopes attached to the private channel integration.
	IntegrationAccessedScopes *[]string `json:"integration_accessed_scopes,omitempty"`
}
