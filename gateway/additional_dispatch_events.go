package gateway

import "github.com/kolosys/discord-types/discord"

// This file contains additional dispatch event types to complete coverage
// matching the TypeScript discord-api-types library

// =============================================================================
// GUILD SCHEDULED EVENT EVENTS
// =============================================================================

// GuildScheduledEventCreateDispatch represents a guild scheduled event create dispatch event.
type GuildScheduledEventCreateDispatch struct {
	Op Opcodes                               `json:"op"`
	T  string                                `json:"t"`
	S  int                                   `json:"s"`
	D  GuildScheduledEventCreateDispatchData `json:"d"`
}

func (e GuildScheduledEventCreateDispatch) isReceivePayload() {}

type GuildScheduledEventCreateDispatchData struct {
	// GuildScheduledEvent data - defined in payloads/guild_scheduled_event.go
}

// GuildScheduledEventUpdateDispatch represents a guild scheduled event update dispatch event.
type GuildScheduledEventUpdateDispatch struct {
	Op Opcodes                               `json:"op"`
	T  string                                `json:"t"`
	S  int                                   `json:"s"`
	D  GuildScheduledEventUpdateDispatchData `json:"d"`
}

func (e GuildScheduledEventUpdateDispatch) isReceivePayload() {}

type GuildScheduledEventUpdateDispatchData = GuildScheduledEventCreateDispatchData

// GuildScheduledEventDeleteDispatch represents a guild scheduled event delete dispatch event.
type GuildScheduledEventDeleteDispatch struct {
	Op Opcodes                               `json:"op"`
	T  string                                `json:"t"`
	S  int                                   `json:"s"`
	D  GuildScheduledEventDeleteDispatchData `json:"d"`
}

func (e GuildScheduledEventDeleteDispatch) isReceivePayload() {}

type GuildScheduledEventDeleteDispatchData = GuildScheduledEventCreateDispatchData

// GuildScheduledEventUserAddDispatch represents a guild scheduled event user add dispatch event.
type GuildScheduledEventUserAddDispatch struct {
	Op Opcodes                                `json:"op"`
	T  string                                 `json:"t"`
	S  int                                    `json:"s"`
	D  GuildScheduledEventUserAddDispatchData `json:"d"`
}

func (e GuildScheduledEventUserAddDispatch) isReceivePayload() {}

type GuildScheduledEventUserAddDispatchData struct {
	GuildScheduledEventID discord.Snowflake `json:"guild_scheduled_event_id"`
	UserID                discord.Snowflake `json:"user_id"`
	GuildID               discord.Snowflake `json:"guild_id"`
}

// GuildScheduledEventUserRemoveDispatch represents a guild scheduled event user remove dispatch event.
type GuildScheduledEventUserRemoveDispatch struct {
	Op Opcodes                                   `json:"op"`
	T  string                                    `json:"t"`
	S  int                                       `json:"s"`
	D  GuildScheduledEventUserRemoveDispatchData `json:"d"`
}

func (e GuildScheduledEventUserRemoveDispatch) isReceivePayload() {}

type GuildScheduledEventUserRemoveDispatchData = GuildScheduledEventUserAddDispatchData

// =============================================================================
// SOUNDBOARD EVENTS
// =============================================================================

// GuildSoundboardSoundCreateDispatch represents a guild soundboard sound create dispatch event.
type GuildSoundboardSoundCreateDispatch struct {
	Op Opcodes                                `json:"op"`
	T  string                                 `json:"t"`
	S  int                                    `json:"s"`
	D  GuildSoundboardSoundCreateDispatchData `json:"d"`
}

func (e GuildSoundboardSoundCreateDispatch) isReceivePayload() {}

type GuildSoundboardSoundCreateDispatchData struct {
	// SoundboardSound data - defined in payloads/soundboard.go
}

// GuildSoundboardSoundUpdateDispatch represents a guild soundboard sound update dispatch event.
type GuildSoundboardSoundUpdateDispatch struct {
	Op Opcodes                                `json:"op"`
	T  string                                 `json:"t"`
	S  int                                    `json:"s"`
	D  GuildSoundboardSoundUpdateDispatchData `json:"d"`
}

func (e GuildSoundboardSoundUpdateDispatch) isReceivePayload() {}

type GuildSoundboardSoundUpdateDispatchData = GuildSoundboardSoundCreateDispatchData

// GuildSoundboardSoundDeleteDispatch represents a guild soundboard sound delete dispatch event.
type GuildSoundboardSoundDeleteDispatch struct {
	Op Opcodes                                `json:"op"`
	T  string                                 `json:"t"`
	S  int                                    `json:"s"`
	D  GuildSoundboardSoundDeleteDispatchData `json:"d"`
}

func (e GuildSoundboardSoundDeleteDispatch) isReceivePayload() {}

type GuildSoundboardSoundDeleteDispatchData struct {
	SoundID discord.Snowflake `json:"sound_id"`
	GuildID discord.Snowflake `json:"guild_id"`
}

// GuildSoundboardSoundsUpdateDispatch represents a guild soundboard sounds update dispatch event.
type GuildSoundboardSoundsUpdateDispatch struct {
	Op Opcodes                                 `json:"op"`
	T  string                                  `json:"t"`
	S  int                                     `json:"s"`
	D  GuildSoundboardSoundsUpdateDispatchData `json:"d"`
}

func (e GuildSoundboardSoundsUpdateDispatch) isReceivePayload() {}

type GuildSoundboardSoundsUpdateDispatchData struct {
	SoundboardSounds []interface{}     `json:"soundboard_sounds"` // SoundboardSound array
	GuildID          discord.Snowflake `json:"guild_id"`
}

// SoundboardSoundsDispatch represents a soundboard sounds dispatch event.
type SoundboardSoundsDispatch struct {
	Op Opcodes                      `json:"op"`
	T  string                       `json:"t"`
	S  int                          `json:"s"`
	D  SoundboardSoundsDispatchData `json:"d"`
}

func (e SoundboardSoundsDispatch) isReceivePayload() {}

type SoundboardSoundsDispatchData struct {
	SoundboardSounds []interface{}     `json:"soundboard_sounds"` // SoundboardSound array
	GuildID          discord.Snowflake `json:"guild_id"`
}

// =============================================================================
// STAGE INSTANCE EVENTS
// =============================================================================

// StageInstanceCreateDispatch represents a stage instance create dispatch event.
type StageInstanceCreateDispatch struct {
	Op Opcodes                         `json:"op"`
	T  string                          `json:"t"`
	S  int                             `json:"s"`
	D  StageInstanceCreateDispatchData `json:"d"`
}

func (e StageInstanceCreateDispatch) isReceivePayload() {}

type StageInstanceCreateDispatchData struct {
	// StageInstance data - defined in payloads/stage_instance.go
}

// StageInstanceUpdateDispatch represents a stage instance update dispatch event.
type StageInstanceUpdateDispatch struct {
	Op Opcodes                         `json:"op"`
	T  string                          `json:"t"`
	S  int                             `json:"s"`
	D  StageInstanceUpdateDispatchData `json:"d"`
}

func (e StageInstanceUpdateDispatch) isReceivePayload() {}

type StageInstanceUpdateDispatchData = StageInstanceCreateDispatchData

// StageInstanceDeleteDispatch represents a stage instance delete dispatch event.
type StageInstanceDeleteDispatch struct {
	Op Opcodes                         `json:"op"`
	T  string                          `json:"t"`
	S  int                             `json:"s"`
	D  StageInstanceDeleteDispatchData `json:"d"`
}

func (e StageInstanceDeleteDispatch) isReceivePayload() {}

type StageInstanceDeleteDispatchData = StageInstanceCreateDispatchData

// =============================================================================
// SUBSCRIPTION EVENTS
// =============================================================================

// SubscriptionCreateDispatch represents a subscription create dispatch event.
type SubscriptionCreateDispatch struct {
	Op Opcodes                        `json:"op"`
	T  string                         `json:"t"`
	S  int                            `json:"s"`
	D  SubscriptionCreateDispatchData `json:"d"`
}

func (e SubscriptionCreateDispatch) isReceivePayload() {}

type SubscriptionCreateDispatchData struct {
	// Subscription data - defined in payloads/monetization.go
}

// SubscriptionUpdateDispatch represents a subscription update dispatch event.
type SubscriptionUpdateDispatch struct {
	Op Opcodes                        `json:"op"`
	T  string                         `json:"t"`
	S  int                            `json:"s"`
	D  SubscriptionUpdateDispatchData `json:"d"`
}

func (e SubscriptionUpdateDispatch) isReceivePayload() {}

type SubscriptionUpdateDispatchData = SubscriptionCreateDispatchData

// SubscriptionDeleteDispatch represents a subscription delete dispatch event.
type SubscriptionDeleteDispatch struct {
	Op Opcodes                        `json:"op"`
	T  string                         `json:"t"`
	S  int                            `json:"s"`
	D  SubscriptionDeleteDispatchData `json:"d"`
}

func (e SubscriptionDeleteDispatch) isReceivePayload() {}

type SubscriptionDeleteDispatchData = SubscriptionCreateDispatchData

// =============================================================================
// THREAD EVENTS
// =============================================================================

// ThreadCreateDispatch represents a thread create dispatch event.
type ThreadCreateDispatch struct {
	Op Opcodes                  `json:"op"`
	T  string                   `json:"t"`
	S  int                      `json:"s"`
	D  ThreadCreateDispatchData `json:"d"`
}

func (e ThreadCreateDispatch) isReceivePayload() {}

type ThreadCreateDispatchData struct {
	// ThreadChannel data with additional fields
	NewlyCreated *bool `json:"newly_created,omitempty"`
	// ... other thread fields
}

// ThreadUpdateDispatch represents a thread update dispatch event.
type ThreadUpdateDispatch struct {
	Op Opcodes                  `json:"op"`
	T  string                   `json:"t"`
	S  int                      `json:"s"`
	D  ThreadUpdateDispatchData `json:"d"`
}

func (e ThreadUpdateDispatch) isReceivePayload() {}

type ThreadUpdateDispatchData struct {
	// ThreadChannel data
}

// ThreadDeleteDispatch represents a thread delete dispatch event.
type ThreadDeleteDispatch struct {
	Op Opcodes                  `json:"op"`
	T  string                   `json:"t"`
	S  int                      `json:"s"`
	D  ThreadDeleteDispatchData `json:"d"`
}

func (e ThreadDeleteDispatch) isReceivePayload() {}

type ThreadDeleteDispatchData struct {
	ID       discord.Snowflake `json:"id"`
	GuildID  discord.Snowflake `json:"guild_id"`
	ParentID discord.Snowflake `json:"parent_id"`
	Type     int               `json:"type"` // ChannelType
}

// ThreadListSyncDispatch represents a thread list sync dispatch event.
type ThreadListSyncDispatch struct {
	Op Opcodes                    `json:"op"`
	T  string                     `json:"t"`
	S  int                        `json:"s"`
	D  ThreadListSyncDispatchData `json:"d"`
}

func (e ThreadListSyncDispatch) isReceivePayload() {}

type ThreadListSyncDispatchData struct {
	GuildID    discord.Snowflake   `json:"guild_id"`
	ChannelIDs []discord.Snowflake `json:"channel_ids,omitempty"`
	Threads    []interface{}       `json:"threads"` // ThreadChannel array
	Members    []interface{}       `json:"members"` // ThreadMember array
}

// ThreadMembersUpdateDispatch represents a thread members update dispatch event.
type ThreadMembersUpdateDispatch struct {
	Op Opcodes                         `json:"op"`
	T  string                          `json:"t"`
	S  int                             `json:"s"`
	D  ThreadMembersUpdateDispatchData `json:"d"`
}

func (e ThreadMembersUpdateDispatch) isReceivePayload() {}

type ThreadMembersUpdateDispatchData struct {
	ID               discord.Snowflake   `json:"id"`
	GuildID          discord.Snowflake   `json:"guild_id"`
	MemberCount      int                 `json:"member_count"`
	AddedMembers     []interface{}       `json:"added_members,omitempty"` // ThreadMember array
	RemovedMemberIDs []discord.Snowflake `json:"removed_member_ids,omitempty"`
}

// ThreadMemberUpdateDispatch represents a thread member update dispatch event.
type ThreadMemberUpdateDispatch struct {
	Op Opcodes                        `json:"op"`
	T  string                         `json:"t"`
	S  int                            `json:"s"`
	D  ThreadMemberUpdateDispatchData `json:"d"`
}

func (e ThreadMemberUpdateDispatch) isReceivePayload() {}

type ThreadMemberUpdateDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	// ... other thread member fields
}

// =============================================================================
// TYPING START EVENT
// =============================================================================

// TypingStartDispatch represents a typing start dispatch event.
type TypingStartDispatch struct {
	Op Opcodes                 `json:"op"`
	T  string                  `json:"t"`
	S  int                     `json:"s"`
	D  TypingStartDispatchData `json:"d"`
}

func (e TypingStartDispatch) isReceivePayload() {}

type TypingStartDispatchData struct {
	ChannelID discord.Snowflake  `json:"channel_id"`
	GuildID   *discord.Snowflake `json:"guild_id,omitempty"`
	UserID    discord.Snowflake  `json:"user_id"`
	Timestamp int64              `json:"timestamp"`
	Member    *interface{}       `json:"member,omitempty"` // GuildMember
}

// =============================================================================
// USER UPDATE EVENT
// =============================================================================

// UserUpdateDispatch represents a user update dispatch event.
type UserUpdateDispatch struct {
	Op Opcodes                `json:"op"`
	T  string                 `json:"t"`
	S  int                    `json:"s"`
	D  UserUpdateDispatchData `json:"d"`
}

func (e UserUpdateDispatch) isReceivePayload() {}

type UserUpdateDispatchData struct {
	// User data - defined in payloads/user.go
}

// =============================================================================
// PRESENCE UPDATE EVENT
// =============================================================================

// PresenceUpdateDispatch represents a presence update dispatch event.
type PresenceUpdateDispatch struct {
	Op Opcodes                    `json:"op"`
	T  string                     `json:"t"`
	S  int                        `json:"s"`
	D  PresenceUpdateDispatchData `json:"d"`
}

func (e PresenceUpdateDispatch) isReceivePayload() {}

type PresenceUpdateDispatchData struct {
	// Presence update data
	User         interface{}       `json:"user"` // PartialUser
	GuildID      discord.Snowflake `json:"guild_id"`
	Status       string            `json:"status"`
	Activities   []interface{}     `json:"activities"`    // Activity array
	ClientStatus interface{}       `json:"client_status"` // ClientStatus
}

// =============================================================================
// INTEGRATION EVENTS
// =============================================================================

// IntegrationCreateDispatch represents an integration create dispatch event.
type IntegrationCreateDispatch struct {
	Op Opcodes                       `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  IntegrationCreateDispatchData `json:"d"`
}

func (e IntegrationCreateDispatch) isReceivePayload() {}

type IntegrationCreateDispatchData struct {
	GuildID discord.Snowflake `json:"guild_id"`
	// ... other integration fields
}

// IntegrationUpdateDispatch represents an integration update dispatch event.
type IntegrationUpdateDispatch struct {
	Op Opcodes                       `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  IntegrationUpdateDispatchData `json:"d"`
}

func (e IntegrationUpdateDispatch) isReceivePayload() {}

type IntegrationUpdateDispatchData = IntegrationCreateDispatchData

// IntegrationDeleteDispatch represents an integration delete dispatch event.
type IntegrationDeleteDispatch struct {
	Op Opcodes                       `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  IntegrationDeleteDispatchData `json:"d"`
}

func (e IntegrationDeleteDispatch) isReceivePayload() {}

type IntegrationDeleteDispatchData struct {
	ID            discord.Snowflake  `json:"id"`
	GuildID       discord.Snowflake  `json:"guild_id"`
	ApplicationID *discord.Snowflake `json:"application_id,omitempty"`
}

// =============================================================================
// INTERACTION CREATE EVENT
// =============================================================================

// InteractionCreateDispatch represents an interaction create dispatch event.
type InteractionCreateDispatch struct {
	Op Opcodes                       `json:"op"`
	T  string                        `json:"t"`
	S  int                           `json:"s"`
	D  InteractionCreateDispatchData `json:"d"`
}

func (e InteractionCreateDispatch) isReceivePayload() {}

type InteractionCreateDispatchData struct {
	// Interaction data - defined in payloads/interactions.go
}

// =============================================================================
// INVITE EVENTS
// =============================================================================

// InviteCreateDispatch represents an invite create dispatch event.
type InviteCreateDispatch struct {
	Op Opcodes                  `json:"op"`
	T  string                   `json:"t"`
	S  int                      `json:"s"`
	D  InviteCreateDispatchData `json:"d"`
}

func (e InviteCreateDispatch) isReceivePayload() {}

type InviteCreateDispatchData struct {
	ChannelID         discord.Snowflake  `json:"channel_id"`
	Code              string             `json:"code"`
	CreatedAt         int64              `json:"created_at"`
	GuildID           *discord.Snowflake `json:"guild_id,omitempty"`
	Inviter           *interface{}       `json:"inviter,omitempty"` // User
	MaxAge            int                `json:"max_age"`
	MaxUses           int                `json:"max_uses"`
	TargetType        *int               `json:"target_type,omitempty"`
	TargetUser        *interface{}       `json:"target_user,omitempty"`        // User
	TargetApplication *interface{}       `json:"target_application,omitempty"` // PartialApplication
	Temporary         bool               `json:"temporary"`
	Uses              int                `json:"uses"`
}

// InviteDeleteDispatch represents an invite delete dispatch event.
type InviteDeleteDispatch struct {
	Op Opcodes                  `json:"op"`
	T  string                   `json:"t"`
	S  int                      `json:"s"`
	D  InviteDeleteDispatchData `json:"d"`
}

func (e InviteDeleteDispatch) isReceivePayload() {}

type InviteDeleteDispatchData struct {
	ChannelID discord.Snowflake  `json:"channel_id"`
	GuildID   *discord.Snowflake `json:"guild_id,omitempty"`
	Code      string             `json:"code"`
}

// =============================================================================
// WEBHOOK UPDATE EVENT
// =============================================================================

// WebhooksUpdateDispatch represents a webhooks update dispatch event.
type WebhooksUpdateDispatch struct {
	Op Opcodes                    `json:"op"`
	T  string                     `json:"t"`
	S  int                        `json:"s"`
	D  WebhooksUpdateDispatchData `json:"d"`
}

func (e WebhooksUpdateDispatch) isReceivePayload() {}

type WebhooksUpdateDispatchData struct {
	GuildID   discord.Snowflake `json:"guild_id"`
	ChannelID discord.Snowflake `json:"channel_id"`
}
