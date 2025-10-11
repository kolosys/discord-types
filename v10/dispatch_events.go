package v10

// DispatchEvent is a generic base struct for all dispatch events
type DispatchEvent[T any] struct {
	Opcode   Opcode  `json:"op"`
	Type     *string `json:"t,omitempty"`
	Sequence *int    `json:"s,omitempty"`
	Data     *T      `json:"d"`
}

// AutoModerationActionExecutionDispatch represents an auto moderation action execution dispatch event.
type AutoModerationActionExecutionDispatch = DispatchEvent[ReceiveAutoModerationActionExecutionData]

// ApplicationCommandPermissionsUpdateDispatch represents application command permissions update dispatch event.
type ApplicationCommandPermissionsUpdateDispatch = DispatchEvent[ReceiveApplicationCommandPermissionsUpdateData]

// ChannelPinsUpdateDispatch represents a channel pins update dispatch event.
type ChannelPinsUpdateDispatch = DispatchEvent[ReceiveChannelPinsData]

// GuildAuditLogEntryCreateDispatch represents a guild audit log entry create dispatch event.
type GuildAuditLogEntryCreateDispatch = DispatchEvent[AuditLogEntry]

// GuildDeleteDispatch represents a guild delete dispatch event.
type GuildDeleteDispatch = DispatchEvent[UnavailableGuild]

// GuildUpdateDispatch represents a guild update dispatch event.
type GuildUpdateDispatch = DispatchEvent[Guild]

// GuildCreateDispatch represents a guild create dispatch event.
type GuildCreateDispatch = DispatchEvent[Guild]

// GuildBanAddDispatch represents a guild ban add dispatch event.
type GuildBanAddDispatch = DispatchEvent[ReceiveGuildBanData]

// GuildBanRemoveDispatch represents a guild ban remove dispatch event.
type GuildBanRemoveDispatch = DispatchEvent[ReceiveGuildBanData]

// GuildEmojisUpdateDispatch represents a guild emojis update dispatch event.
type GuildEmojisUpdateDispatch = DispatchEvent[ReceiveGuildEmojisUpdateData]

// GuildStickersUpdateDispatch represents a guild stickers update dispatch event.
type GuildStickersUpdateDispatch = DispatchEvent[ReceiveGuildStickersUpdateData]

// GuildIntegrationsUpdateDispatch represents a guild integrations update dispatch event.
type GuildIntegrationsUpdateDispatch = DispatchEvent[ReceiveGuildIntegrationsUpdateData]

// GuildMemberAddDispatch represents a guild member add dispatch event.
type GuildMemberAddDispatch = DispatchEvent[ReceiveGuildMemberAddData]

// GuildMemberRemoveDispatch represents a guild member remove dispatch event.
type GuildMemberRemoveDispatch = DispatchEvent[ReceiveGuildMemberRemoveData]

// GuildMemberUpdateDispatch represents a guild member update dispatch event.
type GuildMemberUpdateDispatch = DispatchEvent[ReceiveGuildMemberUpdateData]

// GuildMembersChunkDispatch represents a guild members chunk dispatch event.
type GuildMembersChunkDispatch = DispatchEvent[ReceiveGuildMembersChunkData]

// GuildRoleCreateDispatch represents a guild role create dispatch event.
type GuildRoleCreateDispatch = DispatchEvent[ReceiveGuildRoleCreateData]

// GuildRoleUpdateDispatch represents a guild role update dispatch event.
type GuildRoleUpdateDispatch = DispatchEvent[ReceiveGuildRoleUpdateData]

// GuildRoleDeleteDispatch represents a guild role delete dispatch event.
type GuildRoleDeleteDispatch = DispatchEvent[ReceiveGuildRoleDeleteData]

// GuildScheduledEventCreateDispatch represents a guild scheduled event create dispatch event.
type GuildScheduledEventCreateDispatch = DispatchEvent[GuildScheduledEvent]

// GuildScheduledEventUpdateDispatch represents a guild scheduled event update dispatch event.
type GuildScheduledEventUpdateDispatch = DispatchEvent[GuildScheduledEvent]

// GuildScheduledEventDeleteDispatch represents a guild scheduled event delete dispatch event.
type GuildScheduledEventDeleteDispatch = DispatchEvent[GuildScheduledEvent]

// GuildScheduledEventUserAddDispatch represents a guild scheduled event user add dispatch event.
type GuildScheduledEventUserAddDispatch = DispatchEvent[ReceiveGuildScheduledEventUserAddData]

// GuildScheduledEventUserRemoveDispatch represents a guild scheduled event user remove dispatch event.
type GuildScheduledEventUserRemoveDispatch = DispatchEvent[ReceiveGuildScheduledEventUserRemoveData]

// IntegrationCreateDispatch represents an integration create dispatch event.
type IntegrationCreateDispatch = DispatchEvent[ReceiveIntegrationCreateData]

// IntegrationUpdateDispatch represents an integration update dispatch event.
type IntegrationUpdateDispatch = DispatchEvent[ReceiveIntegrationUpdateData]

// IntegrationDeleteDispatch represents an integration delete dispatch event.
type IntegrationDeleteDispatch = DispatchEvent[ReceiveIntegrationDeleteData]

// InteractionCreateDispatch represents an interaction create dispatch event.
type InteractionCreateDispatch = DispatchEvent[ReceiveInteractionCreateData]

// InviteCreateDispatch represents an invite create dispatch event.
type InviteCreateDispatch = DispatchEvent[ReceiveInviteCreateData]

// InviteDeleteDispatch represents an invite delete dispatch event.
type InviteDeleteDispatch = DispatchEvent[ReceiveInviteDeleteData]

// MessageCreateDispatch represents a message create dispatch event.
type MessageCreateDispatch = DispatchEvent[Message]

// MessageUpdateDispatch represents a message update dispatch event.
type MessageUpdateDispatch = DispatchEvent[Message]

// MessageDeleteDispatch represents a message delete dispatch event.
type MessageDeleteDispatch = DispatchEvent[ReceiveMessageDeleteData]

// MessageDeleteBulkDispatch represents a message delete bulk dispatch event.
type MessageDeleteBulkDispatch = DispatchEvent[ReceiveMessageDeleteBulkData]

// MessageReactionAddDispatch represents a message reaction add dispatch event.
type MessageReactionAddDispatch = DispatchEvent[ReceiveMessageReactionAddData]

// MessageReactionRemoveDispatch represents a message reaction remove dispatch event.
type MessageReactionRemoveDispatch = DispatchEvent[ReceiveMessageReactionRemoveData]

// MessageReactionRemoveAllDispatch represents a message reaction remove all dispatch event.
type MessageReactionRemoveAllDispatch = DispatchEvent[ReceiveMessageReactionRemoveAllData]

// MessageReactionRemoveEmojiDispatch represents a message reaction remove emoji dispatch event.
type MessageReactionRemoveEmojiDispatch = DispatchEvent[ReceiveMessageReactionRemoveEmojiData]

// PresenceUpdateDispatch represents a presence update dispatch event.
type PresenceUpdateDispatch = DispatchEvent[ReceivePresenceUpdateData]

// ReadyDispatch represents a ready dispatch event.
type ReadyDispatch = DispatchEvent[ReceiveReadyData]

// ResumedDispatch represents a resumed dispatch event.
type ResumedDispatch = DispatchEvent[ReceiveResumedData]

// StageInstanceCreateDispatch represents a stage instance create dispatch event.
type StageInstanceCreateDispatch = DispatchEvent[Stage]

// StageInstanceUpdateDispatch represents a stage instance update dispatch event.
type StageInstanceUpdateDispatch = DispatchEvent[Stage]

// StageInstanceDeleteDispatch represents a stage instance delete dispatch event.
type StageInstanceDeleteDispatch = DispatchEvent[Stage]

// ThreadCreateDispatch represents a thread create dispatch event.
type ThreadCreateDispatch = DispatchEvent[ThreadChannel]

// ThreadUpdateDispatch represents a thread update dispatch event.
type ThreadUpdateDispatch = DispatchEvent[ThreadChannel]

// ThreadDeleteDispatch represents a thread delete dispatch event.
type ThreadDeleteDispatch = DispatchEvent[ThreadChannel]

// ThreadListSyncDispatch represents a thread list sync dispatch event.
type ThreadListSyncDispatch = DispatchEvent[ReceiveThreadListSyncData]

// ThreadMemberUpdateDispatch represents a thread member update dispatch event.
type ThreadMemberUpdateDispatch = DispatchEvent[ReceiveThreadMemberUpdateData]

// ThreadMembersUpdateDispatch represents a thread members update dispatch event.
type ThreadMembersUpdateDispatch = DispatchEvent[ReceiveThreadMembersUpdateData]

// TypingStartDispatch represents a typing start dispatch event.
type TypingStartDispatch = DispatchEvent[ReceiveTypingStartData]

// UserUpdateDispatch represents a user update dispatch event.
type UserUpdateDispatch = DispatchEvent[ReceiveUserUpdateData]

// VoiceStateUpdateDispatch represents a voice state update dispatch event.
type VoiceStateUpdateDispatch = DispatchEvent[ReceiveVoiceStateUpdateData]

// VoiceServerUpdateDispatch represents a voice server update dispatch event.
type VoiceServerUpdateDispatch = DispatchEvent[ReceiveVoiceServerUpdateData]

// WebhooksUpdateDispatch represents a webhooks update dispatch event.
type WebhooksUpdateDispatch = DispatchEvent[ReceiveWebhooksUpdateData]

// AutoModerationRuleCreateDispatch represents an auto moderation rule create dispatch event.
type AutoModerationRuleCreateDispatch = DispatchEvent[AutoModerationRule]

// AutoModerationRuleUpdateDispatch represents an auto moderation rule update dispatch event.
type AutoModerationRuleUpdateDispatch = DispatchEvent[AutoModerationRule]

// AutoModerationRuleDeleteDispatch represents an auto moderation rule delete dispatch event.
type AutoModerationRuleDeleteDispatch = DispatchEvent[AutoModerationRule]

// ChannelCreateDispatch represents a channel create dispatch event.
type ChannelCreateDispatch = DispatchEvent[any]

// ChannelUpdateDispatch represents a channel update dispatch event.
type ChannelUpdateDispatch = DispatchEvent[any]

// ChannelDeleteDispatch represents a channel delete dispatch event.
type ChannelDeleteDispatch = DispatchEvent[any]

// EntitlementCreateDispatch represents an entitlement create dispatch event.
type EntitlementCreateDispatch = DispatchEvent[Entitlement]

// EntitlementUpdateDispatch represents an entitlement update dispatch event.
type EntitlementUpdateDispatch = DispatchEvent[Entitlement]

// EntitlementDeleteDispatch represents an entitlement delete dispatch event.
type EntitlementDeleteDispatch = DispatchEvent[Entitlement]

// GuildSoundboardSoundCreateDispatch represents a guild soundboard sound create dispatch event.
type GuildSoundboardSoundCreateDispatch = DispatchEvent[ReceiveGuildSoundboardSoundCreateData]

// GuildSoundboardSoundUpdateDispatch represents a guild soundboard sound update dispatch event.
type GuildSoundboardSoundUpdateDispatch = DispatchEvent[ReceiveGuildSoundboardSoundUpdateData]

// GuildSoundboardSoundDeleteDispatch represents a guild soundboard sound delete dispatch event.
type GuildSoundboardSoundDeleteDispatch = DispatchEvent[ReceiveGuildSoundboardSoundDeleteData]

// GuildSoundboardSoundsUpdateDispatch represents a guild soundboard sounds update dispatch event.
type GuildSoundboardSoundsUpdateDispatch = DispatchEvent[ReceiveGuildSoundboardSoundsUpdateData]

// SoundboardSoundsDispatch represents a soundboard sounds dispatch event.
type SoundboardSoundsDispatch = DispatchEvent[ReceiveSoundboardSoundsData]

// SubscriptionCreateDispatch represents a subscription create dispatch event.
type SubscriptionCreateDispatch = DispatchEvent[ReceiveSubscriptionCreateData]

// SubscriptionUpdateDispatch represents a subscription update dispatch event.
type SubscriptionUpdateDispatch = DispatchEvent[ReceiveSubscriptionUpdateData]

// SubscriptionDeleteDispatch represents a subscription delete dispatch event.
type SubscriptionDeleteDispatch = DispatchEvent[ReceiveSubscriptionDeleteData]

// VoiceChannelEffectSendDispatch represents a voice channel effect send dispatch event.
type VoiceChannelEffectSendDispatch = DispatchEvent[ReceiveVoiceChannelEffectSendData]

// MessagePollVoteAddDispatch represents a message poll vote add dispatch event.
type MessagePollVoteAddDispatch = DispatchEvent[ReceiveMessagePollVoteAddData]

// MessagePollVoteRemoveDispatch represents a message poll vote remove dispatch event.
type MessagePollVoteRemoveDispatch = DispatchEvent[ReceiveMessagePollVoteRemoveData]

// HelloDispatch represents a hello dispatch event.
type HelloDispatch = DispatchEvent[ReceiveHelloData]

// ReconnectDispatch represents a reconnect dispatch event.
type ReconnectDispatch = DispatchEvent[ReceiveReconnectData]

// InvalidSessionDispatch represents an invalid session dispatch event.
type InvalidSessionDispatch = DispatchEvent[ReceiveInvalidSessionData]

// RateLimitedDispatch represents a rate limited dispatch event.
type RateLimitedDispatch = DispatchEvent[ReceiveRateLimitedData]

// =============================================================================
// SEND EVENTS (Client to Server)
// =============================================================================

// IdentifyDispatch represents an identify dispatch event (client to server).
type IdentifyDispatch = DispatchEvent[SendIdentifyData]

// ResumeDispatch represents a resume dispatch event (client to server).
type ResumeDispatch = DispatchEvent[SendResumeData]

// HeartbeatDispatch represents a heartbeat dispatch event (client to server).
type HeartbeatDispatch = DispatchEvent[SendHeartbeatData]

// RequestGuildMembersDispatch represents a request guild members dispatch event (client to server).
type RequestGuildMembersDispatch = DispatchEvent[SendRequestGuildMembersData]

// RequestSoundboardSoundsDispatch represents a request soundboard sounds dispatch event (client to server).
type RequestSoundboardSoundsDispatch = DispatchEvent[SendRequestSoundboardSoundsData]

// UpdateVoiceStateDispatch represents an update voice state dispatch event (client to server).
type UpdateVoiceStateDispatch = DispatchEvent[SendUpdateVoiceStateData]

// UpdatePresenceDispatch represents an update presence dispatch event (client to server).
type UpdatePresenceDispatch = DispatchEvent[SendUpdatePresenceData]
