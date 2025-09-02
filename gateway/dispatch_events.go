package gateway

import "github.com/kolosys/discord-types/discord"

// This file contains comprehensive dispatch event types and data structures
// matching the TypeScript discord-api-types library

// =============================================================================
// AUTO MODERATION EVENTS
// =============================================================================

// AutoModerationRuleCreateDispatch represents an auto moderation rule create dispatch event.
type AutoModerationRuleCreateDispatch struct {
	Op GatewayOpcode                        `json:"op"`
	T  string                               `json:"t"`
	S  int                                  `json:"s"`
	D  AutoModerationRuleCreateDispatchData `json:"d"`
}

func (e AutoModerationRuleCreateDispatch) isReceivePayload() {}

type AutoModerationRuleCreateDispatchData struct {
	// This would contain the full auto moderation rule data
	// Defined in payloads/auto_moderation.go
}

// AutoModerationRuleUpdateDispatch represents an auto moderation rule update dispatch event.
type AutoModerationRuleUpdateDispatch struct {
	Op GatewayOpcode                        `json:"op"`
	T  string                               `json:"t"`
	S  int                                  `json:"s"`
	D  AutoModerationRuleUpdateDispatchData `json:"d"`
}

func (e AutoModerationRuleUpdateDispatch) isReceivePayload() {}

type AutoModerationRuleUpdateDispatchData struct {
	// Auto moderation rule data
}

// AutoModerationRuleDeleteDispatch represents an auto moderation rule delete dispatch event.
type AutoModerationRuleDeleteDispatch struct {
	Op GatewayOpcode                        `json:"op"`
	T  string                               `json:"t"`
	S  int                                  `json:"s"`
	D  AutoModerationRuleDeleteDispatchData `json:"d"`
}

func (e AutoModerationRuleDeleteDispatch) isReceivePayload() {}

type AutoModerationRuleDeleteDispatchData struct {
	// Auto moderation rule data
}

// AutoModerationActionExecutionDispatch represents an auto moderation action execution dispatch event.
type AutoModerationActionExecutionDispatch struct {
	Op GatewayOpcode                             `json:"op"`
	T  string                                    `json:"t"`
	S  int                                       `json:"s"`
	D  AutoModerationActionExecutionDispatchData `json:"d"`
}

func (e AutoModerationActionExecutionDispatch) isReceivePayload() {}

type AutoModerationActionExecutionDispatchData struct {
	GuildID              discord.Snowflake  `json:"guild_id"`
	Action               interface{}        `json:"action"` // AutoModerationAction
	RuleID               discord.Snowflake  `json:"rule_id"`
	RuleTriggerType      int                `json:"rule_trigger_type"`
	UserID               discord.Snowflake  `json:"user_id"`
	ChannelID            *discord.Snowflake `json:"channel_id,omitempty"`
	MessageID            *discord.Snowflake `json:"message_id,omitempty"`
	AlertSystemMessageID *discord.Snowflake `json:"alert_system_message_id,omitempty"`
	Content              string             `json:"content"`
	MatchedKeyword       *string            `json:"matched_keyword"`
	MatchedContent       *string            `json:"matched_content"`
}

// =============================================================================
// APPLICATION COMMAND EVENTS
// =============================================================================

// ApplicationCommandPermissionsUpdateDispatch represents application command permissions update dispatch event.
type ApplicationCommandPermissionsUpdateDispatch struct {
	Op GatewayOpcode                                   `json:"op"`
	T  string                                          `json:"t"`
	S  int                                             `json:"s"`
	D  ApplicationCommandPermissionsUpdateDispatchData `json:"d"`
}

func (e ApplicationCommandPermissionsUpdateDispatch) isReceivePayload() {}

type ApplicationCommandPermissionsUpdateDispatchData struct {
	ID            discord.Snowflake `json:"id"`
	ApplicationID discord.Snowflake `json:"application_id"`
	GuildID       discord.Snowflake `json:"guild_id"`
	Permissions   []interface{}     `json:"permissions"` // ApplicationCommandPermission array
}

// =============================================================================
// CHANNEL EVENTS
// =============================================================================

// ChannelCreateDispatch represents a channel create dispatch event.
type ChannelCreateDispatch struct {
	Op GatewayOpcode             `json:"op"`
	T  string                    `json:"t"`
	S  int                       `json:"s"`
	D  ChannelCreateDispatchData `json:"d"`
}

func (e ChannelCreateDispatch) isReceivePayload() {}

type ChannelCreateDispatchData struct {
	// Channel data with guild_id required
	// Type would exclude thread channel types
	GuildID discord.Snowflake `json:"guild_id"`
	// ... other channel fields
}

// ChannelUpdateDispatch represents a channel update dispatch event.
type ChannelUpdateDispatch struct {
	Op GatewayOpcode             `json:"op"`
	T  string                    `json:"t"`
	S  int                       `json:"s"`
	D  ChannelUpdateDispatchData `json:"d"`
}

func (e ChannelUpdateDispatch) isReceivePayload() {}

type ChannelUpdateDispatchData = ChannelCreateDispatchData

// ChannelDeleteDispatch represents a channel delete dispatch event.
type ChannelDeleteDispatch struct {
	Op GatewayOpcode             `json:"op"`
	T  string                    `json:"t"`
	S  int                       `json:"s"`
	D  ChannelDeleteDispatchData `json:"d"`
}

func (e ChannelDeleteDispatch) isReceivePayload() {}

type ChannelDeleteDispatchData = ChannelCreateDispatchData

// ChannelPinsUpdateDispatch represents a channel pins update dispatch event.
type ChannelPinsUpdateDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  ChannelPinsUpdateDispatchData `json:"d"`
}

func (e ChannelPinsUpdateDispatch) isReceivePayload() {}

type ChannelPinsUpdateDispatchData struct {
	GuildID          *discord.Snowflake `json:"guild_id,omitempty"`
	ChannelID        discord.Snowflake  `json:"channel_id"`
	LastPinTimestamp *string            `json:"last_pin_timestamp"`
}

// =============================================================================
// ENTITLEMENT EVENTS
// =============================================================================

// EntitlementCreateDispatch represents an entitlement create dispatch event.
type EntitlementCreateDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  EntitlementCreateDispatchData `json:"d"`
}

func (e EntitlementCreateDispatch) isReceivePayload() {}

type EntitlementCreateDispatchData struct {
	// Entitlement data - defined in payloads/monetization.go
}

// EntitlementUpdateDispatch represents an entitlement update dispatch event.
type EntitlementUpdateDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  EntitlementUpdateDispatchData `json:"d"`
}

func (e EntitlementUpdateDispatch) isReceivePayload() {}

type EntitlementUpdateDispatchData = EntitlementCreateDispatchData

// EntitlementDeleteDispatch represents an entitlement delete dispatch event.
type EntitlementDeleteDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  EntitlementDeleteDispatchData `json:"d"`
}

func (e EntitlementDeleteDispatch) isReceivePayload() {}

type EntitlementDeleteDispatchData = EntitlementCreateDispatchData

// =============================================================================
// GUILD AUDIT LOG EVENTS
// =============================================================================

// GuildAuditLogEntryCreateDispatch represents a guild audit log entry create dispatch event.
type GuildAuditLogEntryCreateDispatch struct {
	Op GatewayOpcode                        `json:"op"`
	T  string                               `json:"t"`
	S  int                                  `json:"s"`
	D  GuildAuditLogEntryCreateDispatchData `json:"d"`
}

func (e GuildAuditLogEntryCreateDispatch) isReceivePayload() {}

type GuildAuditLogEntryCreateDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	// ... other audit log entry fields
}

// =============================================================================
// GUILD BAN EVENTS
// =============================================================================

// GuildBanAddDispatch represents a guild ban add dispatch event.
type GuildBanAddDispatch struct {
	Op GatewayOpcode           `json:"op"`
	T  string                  `json:"t"`
	S  int                     `json:"s"`
	D  GuildBanAddDispatchData `json:"d"`
}

func (e GuildBanAddDispatch) isReceivePayload() {}

type GuildBanAddDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	User    interface{}       `json:"user"` // User
}

// GuildBanRemoveDispatch represents a guild ban remove dispatch event.
type GuildBanRemoveDispatch struct {
	Op GatewayOpcode              `json:"op"`
	T  string                     `json:"t"`
	S  int                        `json:"s"`
	D  GuildBanRemoveDispatchData `json:"d"`
}

func (e GuildBanRemoveDispatch) isReceivePayload() {}

type GuildBanRemoveDispatchData = GuildBanAddDispatchData

// =============================================================================
// GUILD DELETE EVENT
// =============================================================================

// GuildDeleteDispatch represents a guild delete dispatch event.
type GuildDeleteDispatch struct {
	Op GatewayOpcode           `json:"op"`
	T  string                  `json:"t"`
	S  int                     `json:"s"`
	D  GuildDeleteDispatchData `json:"d"`
}

func (e GuildDeleteDispatch) isReceivePayload() {}

type GuildDeleteDispatchData struct {
	ID          discord.Snowflake `json:"id"`
	Unavailable *bool             `json:"unavailable,omitempty"`
}

// =============================================================================
// GUILD UPDATE EVENT
// =============================================================================

// GuildUpdateDispatch represents a guild update dispatch event.
type GuildUpdateDispatch struct {
	Op GatewayOpcode           `json:"op"`
	T  string                  `json:"t"`
	S  int                     `json:"s"`
	D  GuildUpdateDispatchData `json:"d"`
}

func (e GuildUpdateDispatch) isReceivePayload() {}

type GuildUpdateDispatchData struct {
	// Full guild data
}

// =============================================================================
// GUILD EMOJIS UPDATE EVENT
// =============================================================================

// GuildEmojisUpdateDispatch represents a guild emojis update dispatch event.
type GuildEmojisUpdateDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  GuildEmojisUpdateDispatchData `json:"d"`
}

func (e GuildEmojisUpdateDispatch) isReceivePayload() {}

type GuildEmojisUpdateDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	Emojis  []interface{}     `json:"emojis"` // Emoji array
}

// =============================================================================
// GUILD STICKERS UPDATE EVENT
// =============================================================================

// GuildStickersUpdateDispatch represents a guild stickers update dispatch event.
type GuildStickersUpdateDispatch struct {
	Op GatewayOpcode                   `json:"op"`
	T  string                          `json:"t"`
	S  int                             `json:"s"`
	D  GuildStickersUpdateDispatchData `json:"d"`
}

func (e GuildStickersUpdateDispatch) isReceivePayload() {}

type GuildStickersUpdateDispatchData struct {
	GuildID  discord.Snowflake `json:"guild_id"`
	Stickers []interface{}     `json:"stickers"` // Sticker array
}

// =============================================================================
// GUILD INTEGRATIONS UPDATE EVENT
// =============================================================================

// GuildIntegrationsUpdateDispatch represents a guild integrations update dispatch event.
type GuildIntegrationsUpdateDispatch struct {
	Op GatewayOpcode                       `json:"op"`
	T  string                              `json:"t"`
	S  int                                 `json:"s"`
	D  GuildIntegrationsUpdateDispatchData `json:"d"`
}

func (e GuildIntegrationsUpdateDispatch) isReceivePayload() {}

type GuildIntegrationsUpdateDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
}

// =============================================================================
// GUILD MEMBER EVENTS
// =============================================================================

// GuildMemberAddDispatch represents a guild member add dispatch event.
type GuildMemberAddDispatch struct {
	Op GatewayOpcode              `json:"op"`
	T  string                     `json:"t"`
	S  int                        `json:"s"`
	D  GuildMemberAddDispatchData `json:"d"`
}

func (e GuildMemberAddDispatch) isReceivePayload() {}

type GuildMemberAddDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	// ... other guild member fields
}

// GuildMemberRemoveDispatch represents a guild member remove dispatch event.
type GuildMemberRemoveDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  GuildMemberRemoveDispatchData `json:"d"`
}

func (e GuildMemberRemoveDispatch) isReceivePayload() {}

type GuildMemberRemoveDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	User    interface{}       `json:"user"` // User
}

// GuildMemberUpdateDispatch represents a guild member update dispatch event.
type GuildMemberUpdateDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  GuildMemberUpdateDispatchData `json:"d"`
}

func (e GuildMemberUpdateDispatch) isReceivePayload() {}

type GuildMemberUpdateDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	// ... other updated guild member fields
}

// GuildMembersChunkDispatch represents a guild members chunk dispatch event.
type GuildMembersChunkDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  GuildMembersChunkDispatchData `json:"d"`
}

func (e GuildMembersChunkDispatch) isReceivePayload() {}

type GuildMembersChunkDispatchData struct {
	GuildID    discord.Snowflake `json:"guild_id"`
	Members    []interface{}     `json:"members"` // GuildMember array
	ChunkIndex int               `json:"chunk_index"`
	ChunkCount int               `json:"chunk_count"`
	NotFound   []interface{}     `json:"not_found,omitempty"`
	Presences  []interface{}     `json:"presences,omitempty"` // PresenceUpdate array
	Nonce      *string           `json:"nonce,omitempty"`
}

// =============================================================================
// GUILD ROLE EVENTS
// =============================================================================

// GuildRoleCreateDispatch represents a guild role create dispatch event.
type GuildRoleCreateDispatch struct {
	Op GatewayOpcode               `json:"op"`
	T  string                      `json:"t"`
	S  int                         `json:"s"`
	D  GuildRoleCreateDispatchData `json:"d"`
}

func (e GuildRoleCreateDispatch) isReceivePayload() {}

type GuildRoleCreateDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	Role    interface{}       `json:"role"` // Role
}

// GuildRoleUpdateDispatch represents a guild role update dispatch event.
type GuildRoleUpdateDispatch struct {
	Op GatewayOpcode               `json:"op"`
	T  string                      `json:"t"`
	S  int                         `json:"s"`
	D  GuildRoleUpdateDispatchData `json:"d"`
}

func (e GuildRoleUpdateDispatch) isReceivePayload() {}

type GuildRoleUpdateDispatchData = GuildRoleCreateDispatchData

// GuildRoleDeleteDispatch represents a guild role delete dispatch event.
type GuildRoleDeleteDispatch struct {
	Op GatewayOpcode               `json:"op"`
	T  string                      `json:"t"`
	S  int                         `json:"s"`
	D  GuildRoleDeleteDispatchData `json:"d"`
}

func (e GuildRoleDeleteDispatch) isReceivePayload() {}

type GuildRoleDeleteDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	RoleID  discord.Snowflake `json:"role_id"`
}

// =============================================================================
// MESSAGE EVENTS
// =============================================================================

// MessageUpdateDispatch represents a message update dispatch event.
type MessageUpdateDispatch struct {
	Op GatewayOpcode             `json:"op"`
	T  string                    `json:"t"`
	S  int                       `json:"s"`
	D  MessageUpdateDispatchData `json:"d"`
}

func (e MessageUpdateDispatch) isReceivePayload() {}

type MessageUpdateDispatchData struct {
	// Partial message data with gateway extra fields
	GuildID  *discord.Snowflake `json:"guild_id,omitempty"`
	Member   *interface{}       `json:"member,omitempty"` // GuildMemberNoUser
	Mentions []interface{}      `json:"mentions"`         // UserWithMember array
	// ... other message fields
}

// MessageDeleteDispatch represents a message delete dispatch event.
type MessageDeleteDispatch struct {
	Op GatewayOpcode             `json:"op"`
	T  string                    `json:"t"`
	S  int                       `json:"s"`
	D  MessageDeleteDispatchData `json:"d"`
}

func (e MessageDeleteDispatch) isReceivePayload() {}

type MessageDeleteDispatchData struct {
	ID        discord.Snowflake  `json:"id"`
	ChannelID discord.Snowflake  `json:"channel_id"`
	GuildID   *discord.Snowflake `json:"guild_id,omitempty"`
}

// MessageDeleteBulkDispatch represents a message delete bulk dispatch event.
type MessageDeleteBulkDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  MessageDeleteBulkDispatchData `json:"d"`
}

func (e MessageDeleteBulkDispatch) isReceivePayload() {}

type MessageDeleteBulkDispatchData struct {
	IDs       []discord.Snowflake `json:"ids"`
	ChannelID discord.Snowflake   `json:"channel_id"`
	GuildID   *discord.Snowflake  `json:"guild_id,omitempty"`
}

// =============================================================================
// POLL EVENTS
// =============================================================================

// MessagePollVoteAddDispatch represents a message poll vote add dispatch event.
type MessagePollVoteAddDispatch struct {
	Op GatewayOpcode                  `json:"op"`
	T  string                         `json:"t"`
	S  int                            `json:"s"`
	D  MessagePollVoteAddDispatchData `json:"d"`
}

func (e MessagePollVoteAddDispatch) isReceivePayload() {}

type MessagePollVoteAddDispatchData struct {
	UserID    discord.Snowflake  `json:"user_id"`
	ChannelID discord.Snowflake  `json:"channel_id"`
	MessageID discord.Snowflake  `json:"message_id"`
	GuildID   *discord.Snowflake `json:"guild_id,omitempty"`
	AnswerID  int                `json:"answer_id"`
}

// MessagePollVoteRemoveDispatch represents a message poll vote remove dispatch event.
type MessagePollVoteRemoveDispatch struct {
	Op GatewayOpcode                     `json:"op"`
	T  string                            `json:"t"`
	S  int                               `json:"s"`
	D  MessagePollVoteRemoveDispatchData `json:"d"`
}

func (e MessagePollVoteRemoveDispatch) isReceivePayload() {}

type MessagePollVoteRemoveDispatchData = MessagePollVoteAddDispatchData

// =============================================================================
// REACTION EVENTS
// =============================================================================

// MessageReactionAddDispatch represents a message reaction add dispatch event.
type MessageReactionAddDispatch struct {
	Op GatewayOpcode                  `json:"op"`
	T  string                         `json:"t"`
	S  int                            `json:"s"`
	D  MessageReactionAddDispatchData `json:"d"`
}

func (e MessageReactionAddDispatch) isReceivePayload() {}

type MessageReactionAddDispatchData struct {
	UserID          discord.Snowflake  `json:"user_id"`
	ChannelID       discord.Snowflake  `json:"channel_id"`
	MessageID       discord.Snowflake  `json:"message_id"`
	GuildID         *discord.Snowflake `json:"guild_id,omitempty"`
	Member          *interface{}       `json:"member,omitempty"` // GuildMember
	Emoji           interface{}        `json:"emoji"`            // PartialEmoji
	MessageAuthorID *discord.Snowflake `json:"message_author_id,omitempty"`
	Burst           bool               `json:"burst"`
	BurstColors     []string           `json:"burst_colors,omitempty"`
	Type            int                `json:"type"` // ReactionType
}

// MessageReactionRemoveDispatch represents a message reaction remove dispatch event.
type MessageReactionRemoveDispatch struct {
	Op GatewayOpcode                     `json:"op"`
	T  string                            `json:"t"`
	S  int                               `json:"s"`
	D  MessageReactionRemoveDispatchData `json:"d"`
}

func (e MessageReactionRemoveDispatch) isReceivePayload() {}

type MessageReactionRemoveDispatchData struct {
	UserID    discord.Snowflake  `json:"user_id"`
	ChannelID discord.Snowflake  `json:"channel_id"`
	MessageID discord.Snowflake  `json:"message_id"`
	GuildID   *discord.Snowflake `json:"guild_id,omitempty"`
	Emoji     interface{}        `json:"emoji"` // PartialEmoji
	Burst     bool               `json:"burst"`
	Type      int                `json:"type"` // ReactionType
}

// MessageReactionRemoveAllDispatch represents a message reaction remove all dispatch event.
type MessageReactionRemoveAllDispatch struct {
	Op GatewayOpcode                        `json:"op"`
	T  string                               `json:"t"`
	S  int                                  `json:"s"`
	D  MessageReactionRemoveAllDispatchData `json:"d"`
}

func (e MessageReactionRemoveAllDispatch) isReceivePayload() {}

type MessageReactionRemoveAllDispatchData struct {
	ChannelID discord.Snowflake  `json:"channel_id"`
	MessageID discord.Snowflake  `json:"message_id"`
	GuildID   *discord.Snowflake `json:"guild_id,omitempty"`
}

// MessageReactionRemoveEmojiDispatch represents a message reaction remove emoji dispatch event.
type MessageReactionRemoveEmojiDispatch struct {
	Op GatewayOpcode                          `json:"op"`
	T  string                                 `json:"t"`
	S  int                                    `json:"s"`
	D  MessageReactionRemoveEmojiDispatchData `json:"d"`
}

func (e MessageReactionRemoveEmojiDispatch) isReceivePayload() {}

type MessageReactionRemoveEmojiDispatchData struct {
	ChannelID discord.Snowflake  `json:"channel_id"`
	MessageID discord.Snowflake  `json:"message_id"`
	GuildID   *discord.Snowflake `json:"guild_id,omitempty"`
	Emoji     interface{}        `json:"emoji"` // PartialEmoji
}

// =============================================================================
// VOICE EVENTS
// =============================================================================

// VoiceServerUpdateDispatch represents a voice server update dispatch event.
type VoiceServerUpdateDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  VoiceServerUpdateDispatchData `json:"d"`
}

func (e VoiceServerUpdateDispatch) isReceivePayload() {}

type VoiceServerUpdateDispatchData struct {
	Token    string            `json:"token"`
	GuildID  discord.Snowflake `json:"guild_id"`
	Endpoint *string           `json:"endpoint"`
}

// VoiceStateUpdateDispatch represents a voice state update dispatch event.
type VoiceStateUpdateDispatch struct {
	Op GatewayOpcode                 `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  VoiceStateUpdateDispatchData2 `json:"d"`
}

func (e VoiceStateUpdateDispatch) isReceivePayload() {}

type VoiceStateUpdateDispatchData2 struct {
	// Voice state data with member included
	Member *interface{} `json:"member,omitempty"` // VoiceStateMember
	// ... other voice state fields
}

// VoiceChannelEffectSendDispatch represents a voice channel effect send dispatch event.
type VoiceChannelEffectSendDispatch struct {
	Op GatewayOpcode                      `json:"op"`
	T  string                             `json:"t"`
	S  int                                `json:"s"`
	D  VoiceChannelEffectSendDispatchData `json:"d"`
}

func (e VoiceChannelEffectSendDispatch) isReceivePayload() {}

type VoiceChannelEffectSendDispatchData struct {
	ChannelID     discord.Snowflake `json:"channel_id"`
	GuildID       discord.Snowflake `json:"guild_id"`
	UserID        discord.Snowflake `json:"user_id"`
	Emoji         *interface{}      `json:"emoji,omitempty"` // PartialEmoji
	AnimationType *int              `json:"animation_type,omitempty"`
	AnimationID   *int              `json:"animation_id,omitempty"`
	SoundID       interface{}       `json:"sound_id,omitempty"` // Snowflake or number
	SoundVolume   *float64          `json:"sound_volume,omitempty"`
}

// Add more dispatch events as needed...
