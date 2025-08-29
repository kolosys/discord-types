package payloads

import "github.com/kolosys/discord-types/discord"

// AutoModerationRule represents a Discord auto moderation rule.
//
// See: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-auto-moderation-rule-structure
type AutoModerationRule struct {
	// ID is the id of this rule.
	ID discord.Snowflake `json:"id"`

	// GuildID is the guild which this rule belongs to.
	GuildID discord.Snowflake `json:"guild_id"`

	// Name is the rule name.
	Name string `json:"name"`

	// CreatorID is the user id who created this rule.
	CreatorID discord.Snowflake `json:"creator_id"`

	// EventType is the rule event type.
	EventType AutoModerationRuleEventType `json:"event_type"`

	// TriggerType is the rule trigger type.
	TriggerType AutoModerationRuleTriggerType `json:"trigger_type"`

	// TriggerMetadata is the rule trigger metadata.
	TriggerMetadata AutoModerationRuleTriggerMetadata `json:"trigger_metadata"`

	// Actions are the actions which will execute when this rule is triggered.
	Actions []AutoModerationAction `json:"actions"`

	// Enabled indicates whether this rule is enabled.
	Enabled bool `json:"enabled"`

	// ExemptRoles are the role ids that shouldn't be affected by this rule (Maximum of 20).
	ExemptRoles []discord.Snowflake `json:"exempt_roles"`

	// ExemptChannels are the channel ids that shouldn't be affected by this rule (Maximum of 50).
	ExemptChannels []discord.Snowflake `json:"exempt_channels"`
}

// AutoModerationRuleEventType represents when a rule should be checked.
//
// See: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-event-types
type AutoModerationRuleEventType int

// Auto moderation rule event type constants
const (
	// AutoModerationRuleEventTypeMessageSend indicates when a member sends or edits a message in the guild.
	AutoModerationRuleEventTypeMessageSend AutoModerationRuleEventType = 1

	// AutoModerationRuleEventTypeMemberUpdate indicates when a member edits their profile.
	AutoModerationRuleEventTypeMemberUpdate AutoModerationRuleEventType = 2
)

// AutoModerationRuleTriggerType represents the type of content which can trigger the rule.
//
// See: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-trigger-types
type AutoModerationRuleTriggerType int

// Auto moderation rule trigger type constants
const (
	// AutoModerationRuleTriggerTypeKeyword indicates check if content contains words from a user defined list of keywords.
	AutoModerationRuleTriggerTypeKeyword AutoModerationRuleTriggerType = iota + 1

	// AutoModerationRuleTriggerTypeSpam indicates check if content represents generic spam.
	AutoModerationRuleTriggerTypeSpam

	// AutoModerationRuleTriggerTypeKeywordPreset indicates check if content contains words from internal pre-defined wordsets.
	AutoModerationRuleTriggerTypeKeywordPreset

	// AutoModerationRuleTriggerTypeMentionSpam indicates check if content contains more unique mentions than allowed.
	AutoModerationRuleTriggerTypeMentionSpam

	// AutoModerationRuleTriggerTypeMemberProfile indicates check if member profile contains words from a user defined list of keywords.
	AutoModerationRuleTriggerTypeMemberProfile
)

// AutoModerationRuleTriggerMetadata represents additional data used to determine whether a rule should be triggered.
//
// See: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-trigger-metadata
type AutoModerationRuleTriggerMetadata struct {
	// KeywordFilter are substrings which will be searched for in content (Maximum of 1000).
	KeywordFilter []string `json:"keyword_filter,omitempty"`

	// RegexPatterns are regular expression patterns which will be matched against content (Maximum of 10).
	RegexPatterns []string `json:"regex_patterns,omitempty"`

	// Presets are the internally pre-defined wordsets which will be searched for in content.
	Presets []AutoModerationRuleKeywordPresetType `json:"presets,omitempty"`

	// AllowList are substrings which should not trigger the rule (Maximum of 100 or 1000).
	AllowList []string `json:"allow_list,omitempty"`

	// MentionTotalLimit is the total number of unique role and user mentions allowed per message (Maximum of 50).
	MentionTotalLimit *int `json:"mention_total_limit,omitempty"`

	// MentionRaidProtectionEnabled indicates whether to automatically detect mention raids.
	MentionRaidProtectionEnabled *bool `json:"mention_raid_protection_enabled,omitempty"`
}

// AutoModerationRuleKeywordPresetType represents the internally pre-defined wordsets.
//
// See: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-rule-object-keyword-preset-types
type AutoModerationRuleKeywordPresetType int

// Auto moderation rule keyword preset type constants
const (
	// AutoModerationRuleKeywordPresetTypeProfanity indicates words that may be considered forms of swearing or cursing.
	AutoModerationRuleKeywordPresetTypeProfanity AutoModerationRuleKeywordPresetType = iota + 1

	// AutoModerationRuleKeywordPresetTypeSexualContent indicates words that refer to sexually explicit behavior or activity.
	AutoModerationRuleKeywordPresetTypeSexualContent

	// AutoModerationRuleKeywordPresetTypeSlurs indicates personal insults or words that may be considered hate speech.
	AutoModerationRuleKeywordPresetTypeSlurs
)

// AutoModerationAction represents an action which will execute whenever a rule is triggered.
//
// See: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object-auto-moderation-action-structure
type AutoModerationAction struct {
	// Type is the type of action.
	Type AutoModerationActionType `json:"type"`

	// Metadata is additional metadata needed during execution for this specific action type.
	Metadata *AutoModerationActionMetadata `json:"metadata,omitempty"`
}

// AutoModerationActionType represents the type of action.
//
// See: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object-action-types
type AutoModerationActionType int

// Auto moderation action type constants
const (
	// AutoModerationActionTypeBlockMessage indicates blocks a member's message and prevents it from being posted.
	// A custom explanation can be specified and shown to members whenever their message is blocked.
	AutoModerationActionTypeBlockMessage AutoModerationActionType = iota + 1

	// AutoModerationActionTypeSendAlertMessage indicates logs user content to a specified channel.
	AutoModerationActionTypeSendAlertMessage

	// AutoModerationActionTypeTimeout indicates timeout user for a specified duration.
	AutoModerationActionTypeTimeout

	// AutoModerationActionTypeBlockMemberInteraction indicates prevents a member from using text, voice, or other interactions.
	AutoModerationActionTypeBlockMemberInteraction
)

// AutoModerationActionMetadata represents additional metadata needed during execution for a specific action type.
//
// See: https://discord.com/developers/docs/resources/auto-moderation#auto-moderation-action-object-action-metadata
type AutoModerationActionMetadata struct {
	// ChannelID is the channel to which user content should be logged.
	ChannelID *discord.Snowflake `json:"channel_id,omitempty"`

	// DurationSeconds is the timeout duration in seconds.
	DurationSeconds *int `json:"duration_seconds,omitempty"`

	// CustomMessage is additional explanation that will be shown to members whenever their message is blocked.
	CustomMessage *string `json:"custom_message,omitempty"`
}
