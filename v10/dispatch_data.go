package v10

// Reference: https://discord.com/developers/docs/events/gateway#dispatch-events

// =============================================================================
// AUTO MODERATION EVENTS
// =============================================================================

type ReceiveAutoModerationActionExecutionData struct {
	GuildID              Snowflake            `json:"guild_id"`
	Action               AutoModerationAction `json:"action"`
	RuleID               Snowflake            `json:"rule_id"`
	RuleTriggerType      int                  `json:"rule_trigger_type"`
	UserID               Snowflake            `json:"user_id"`
	ChannelID            *Snowflake           `json:"channel_id,omitempty"`
	MessageID            *Snowflake           `json:"message_id,omitempty"`
	AlertSystemMessageID *Snowflake           `json:"alert_system_message_id,omitempty"`
	Content              string               `json:"content"`
	MatchedKeyword       *string              `json:"matched_keyword"`
	MatchedContent       *string              `json:"matched_content"`
}

// =============================================================================
// APPLICATION COMMAND EVENTS
// =============================================================================

type ReceiveApplicationCommandPermissionsUpdateData struct {
	ID            Snowflake `json:"id"`
	ApplicationID Snowflake `json:"application_id"`
	GuildID       Snowflake `json:"guild_id"`
	Permissions   []any     `json:"permissions"` // ApplicationCommandPermission array
}

// =============================================================================
// CHANNEL EVENTS
// =============================================================================

type ReceiveChannelData struct {
	// Channel data with guild_id required
	// Type would exclude thread channel types
	GuildID Snowflake `json:"guild_id"`
}

type ReceiveChannelPinsData struct {
	GuildID          *Snowflake `json:"guild_id,omitempty"`
	ChannelID        Snowflake  `json:"channel_id"`
	LastPinTimestamp *string    `json:"last_pin_timestamp"`
}

// =============================================================================
// GUILD EVENTS
// =============================================================================

type ReceiveGuildBanData struct {
	GuildID Snowflake `json:"guild_id"`
	User    User      `json:"user"`
}

type ReceiveGuildEmojisUpdateData struct {
	GuildID Snowflake `json:"guild_id"`
	Emojis  []Emoji   `json:"emojis"`
}

type ReceiveGuildStickersUpdateData struct {
	GuildID  Snowflake `json:"guild_id"`
	Stickers []Sticker `json:"stickers"`
}

type ReceiveGuildIntegrationsUpdateData struct {
	GuildID Snowflake `json:"guild_id"`
}

type ReceiveGuildMemberAddData struct {
	GuildMember
	GuildID Snowflake `json:"guild_id"`
}

type ReceiveGuildMemberRemoveData struct {
	GuildID Snowflake `json:"guild_id"`
	User    User      `json:"user"`
}

type ReceiveGuildMemberUpdateData struct {
	GuildID                    Snowflake   `json:"guild_id"`
	Roles                      []Snowflake `json:"roles"`
	User                       User        `json:"user"`
	Nick                       *string     `json:"nick"`
	Avatar                     *string     `json:"avatar"`
	JoinedAt                   *string     `json:"joined_at"`
	PremiumSince               *string     `json:"premium_since"`
	Deaf                       *bool       `json:"deaf,omitempty"`
	Mute                       *bool       `json:"mute,omitempty"`
	Pending                    *bool       `json:"pending,omitempty"`
	CommunicationDisabledUntil *string     `json:"communication_disabled_until"`
}

type ReceiveGuildMembersChunkData struct {
	GuildID    Snowflake     `json:"guild_id"`
	Members    []GuildMember `json:"members"`
	ChunkIndex int           `json:"chunk_index"`
	ChunkCount int           `json:"chunk_count"`
	NotFound   []any         `json:"not_found,omitempty"`
	Presences  []any         `json:"presences,omitempty"` // Presence Update objects
	Nonce      *string       `json:"nonce,omitempty"`
}

type ReceiveGuildRoleCreateData struct {
	GuildID Snowflake `json:"guild_id"`
	Role    Role      `json:"role"`
}

type ReceiveGuildRoleUpdateData struct {
	GuildID Snowflake `json:"guild_id"`
	Role    Role      `json:"role"`
}

type ReceiveGuildRoleDeleteData struct {
	GuildID Snowflake `json:"guild_id"`
	RoleID  Snowflake `json:"role_id"`
}

type ReceiveGuildScheduledEventUserAddData struct {
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	UserID                Snowflake `json:"user_id"`
	GuildID               Snowflake `json:"guild_id"`
}

type ReceiveGuildScheduledEventUserRemoveData struct {
	GuildScheduledEventID Snowflake `json:"guild_scheduled_event_id"`
	UserID                Snowflake `json:"user_id"`
	GuildID               Snowflake `json:"guild_id"`
}

// =============================================================================
// INTEGRATION EVENTS
// =============================================================================

type ReceiveIntegrationCreateData struct {
	GuildID Snowflake `json:"guild_id"`
	// Integration object fields would be embedded here
}

type ReceiveIntegrationUpdateData struct {
	GuildID Snowflake `json:"guild_id"`
	// Integration object fields would be embedded here
}

type ReceiveIntegrationDeleteData struct {
	ID            Snowflake  `json:"id"`
	GuildID       Snowflake  `json:"guild_id"`
	ApplicationID *Snowflake `json:"application_id,omitempty"`
}

// =============================================================================
// INTERACTION EVENTS
// =============================================================================

type ReceiveInteractionCreateData struct {
	Interaction
}

// =============================================================================
// INVITE EVENTS
// =============================================================================

type ReceiveInviteCreateData struct {
	ChannelID         Snowflake  `json:"channel_id"`
	Code              string     `json:"code"`
	CreatedAt         string     `json:"created_at"`
	GuildID           *Snowflake `json:"guild_id,omitempty"`
	Inviter           *User      `json:"inviter,omitempty"`
	MaxAge            int        `json:"max_age"`
	MaxUses           int        `json:"max_uses"`
	TargetType        *int       `json:"target_type,omitempty"`
	TargetUser        *User      `json:"target_user,omitempty"`
	TargetApplication *any       `json:"target_application,omitempty"` // Partial Application object
	Temporary         bool       `json:"temporary"`
	Uses              int        `json:"uses"`
}

type ReceiveInviteDeleteData struct {
	ChannelID Snowflake  `json:"channel_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
	Code      string     `json:"code"`
}

// =============================================================================
// MESSAGE EVENTS
// =============================================================================

type ReceiveMessageDeleteData struct {
	ID        Snowflake  `json:"id"`
	ChannelID Snowflake  `json:"channel_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
}

type ReceiveMessageDeleteBulkData struct {
	IDs       []Snowflake `json:"ids"`
	ChannelID Snowflake   `json:"channel_id"`
	GuildID   *Snowflake  `json:"guild_id,omitempty"`
}

type ReceiveMessageReactionAddData struct {
	UserID          Snowflake    `json:"user_id"`
	ChannelID       Snowflake    `json:"channel_id"`
	MessageID       Snowflake    `json:"message_id"`
	GuildID         *Snowflake   `json:"guild_id,omitempty"`
	Member          *GuildMember `json:"member,omitempty"`
	Emoji           Emoji        `json:"emoji"`
	MessageAuthorID *Snowflake   `json:"message_author_id,omitempty"`
}

type ReceiveMessageReactionRemoveData struct {
	UserID    Snowflake  `json:"user_id"`
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
	Emoji     Emoji      `json:"emoji"`
}

type ReceiveMessageReactionRemoveAllData struct {
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
}

type ReceiveMessageReactionRemoveEmojiData struct {
	ChannelID Snowflake  `json:"channel_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
	MessageID Snowflake  `json:"message_id"`
	Emoji     Emoji      `json:"emoji"`
}

// =============================================================================
// PRESENCE EVENTS
// =============================================================================

type ReceivePresenceUpdateData struct {
	User         User      `json:"user"`
	GuildID      Snowflake `json:"guild_id"`
	Status       string    `json:"status"`
	Activities   []any     `json:"activities"`    // Activity objects
	ClientStatus any       `json:"client_status"` // Client Status object
}

// =============================================================================
// READY EVENT
// =============================================================================

type ReceiveReadyData struct {
	V                int     `json:"v"`
	User             User    `json:"user"`
	Guilds           []any   `json:"guilds"` // Unavailable Guild objects
	SessionID        string  `json:"session_id"`
	ResumeGatewayURL string  `json:"resume_gateway_url"`
	Shard            *[2]int `json:"shard,omitempty"`
	Application      any     `json:"application"` // Partial Application object
}

// =============================================================================
// RESUME EVENT
// =============================================================================

type ReceiveResumedData struct {
	// Resumed event has no event data - just an acknowledgment
}

// =============================================================================
// THREAD EVENTS
// =============================================================================

type ReceiveThreadListSyncData struct {
	GuildID    Snowflake   `json:"guild_id"`
	ChannelIDs []Snowflake `json:"channel_ids,omitempty"`
	Threads    []any       `json:"threads"` // Channel objects
	Members    []any       `json:"members"` // Thread Member objects
}

type ReceiveThreadMemberUpdateData struct {
	ID            Snowflake `json:"id"`
	UserID        Snowflake `json:"user_id"`
	JoinTimestamp string    `json:"join_timestamp"`
	Flags         int       `json:"flags"`
	GuildID       Snowflake `json:"guild_id"`
}

type ReceiveThreadMembersUpdateData struct {
	ID               Snowflake   `json:"id"`
	GuildID          Snowflake   `json:"guild_id"`
	MemberCount      int         `json:"member_count"`
	AddedMembers     []any       `json:"added_members,omitempty"` // Thread Member objects
	RemovedMemberIDs []Snowflake `json:"removed_member_ids,omitempty"`
}

// =============================================================================
// TYPING EVENTS
// =============================================================================

type ReceiveTypingStartData struct {
	ChannelID Snowflake    `json:"channel_id"`
	GuildID   *Snowflake   `json:"guild_id,omitempty"`
	UserID    Snowflake    `json:"user_id"`
	Timestamp int          `json:"timestamp"`
	Member    *GuildMember `json:"member,omitempty"`
}

// =============================================================================
// USER EVENTS
// =============================================================================

type ReceiveUserUpdateData struct {
	// User object - same as User type
	User
}

// =============================================================================
// VOICE EVENTS
// =============================================================================

type ReceiveVoiceStateUpdateData struct {
	GuildID                 *Snowflake   `json:"guild_id,omitempty"`
	ChannelID               *Snowflake   `json:"channel_id"`
	UserID                  Snowflake    `json:"user_id"`
	Member                  *GuildMember `json:"member,omitempty"`
	SessionID               string       `json:"session_id"`
	Deaf                    bool         `json:"deaf"`
	Mute                    bool         `json:"mute"`
	SelfDeaf                bool         `json:"self_deaf"`
	SelfMute                bool         `json:"self_mute"`
	SelfStream              *bool        `json:"self_stream,omitempty"`
	SelfVideo               bool         `json:"self_video"`
	Suppress                bool         `json:"suppress"`
	RequestToSpeakTimestamp *string      `json:"request_to_speak_timestamp"`
}

type ReceiveVoiceServerUpdateData struct {
	Token    string    `json:"token"`
	GuildID  Snowflake `json:"guild_id"`
	Endpoint *string   `json:"endpoint"`
}

// =============================================================================
// WEBHOOK EVENTS
// =============================================================================

type ReceiveWebhooksUpdateData struct {
	GuildID   Snowflake `json:"guild_id"`
	ChannelID Snowflake `json:"channel_id"`
}

// =============================================================================
// SOUNDBOARD EVENTS
// =============================================================================

type ReceiveGuildSoundboardSoundCreateData struct {
	SoundboardSound
}

type ReceiveGuildSoundboardSoundUpdateData struct {
	SoundboardSound
}

type ReceiveGuildSoundboardSoundDeleteData struct {
	SoundboardSound
}

type ReceiveGuildSoundboardSoundsUpdateData struct {
	GuildID          Snowflake          `json:"guild_id"`
	SoundboardSounds []*SoundboardSound `json:"soundboard_sounds"`
}

type ReceiveSoundboardSoundsData struct {
	SoundboardSounds []*SoundboardSound `json:"soundboard_sounds"`
}

// =============================================================================
// SUBSCRIPTION EVENTS
// =============================================================================

type ReceiveSubscriptionCreateData struct {
	Subscription
}

type ReceiveSubscriptionUpdateData struct {
	Subscription
}

type ReceiveSubscriptionDeleteData struct {
	Subscription
}

// =============================================================================
// VOICE CHANNEL EFFECT EVENTS
// =============================================================================

type ReceiveVoiceChannelEffectSendData struct {
	ChannelID     Snowflake `json:"channel_id"`
	GuildID       Snowflake `json:"guild_id"`
	UserID        Snowflake `json:"user_id"`
	Emoji         *Emoji    `json:"emoji,omitempty"`
	AnimationType *int      `json:"animation_type,omitempty"`
	AnimationID   *int      `json:"animation_id,omitempty"`
	SoundID       *string   `json:"sound_id,omitempty"`
	SoundVolume   *float64  `json:"sound_volume,omitempty"`
}

// =============================================================================
// POLL EVENTS
// =============================================================================

type ReceiveMessagePollVoteAddData struct {
	UserID    Snowflake  `json:"user_id"`
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
	AnswerID  int        `json:"answer_id"`
}

type ReceiveMessagePollVoteRemoveData struct {
	UserID    Snowflake  `json:"user_id"`
	ChannelID Snowflake  `json:"channel_id"`
	MessageID Snowflake  `json:"message_id"`
	GuildID   *Snowflake `json:"guild_id,omitempty"`
	AnswerID  int        `json:"answer_id"`
}

// =============================================================================
// GATEWAY CONTROL EVENTS
// =============================================================================

type ReceiveHelloData struct {
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type ReceiveReconnectData struct {
	// Reconnect event has no data payload
}

type ReceiveInvalidSessionData struct {
	// Contains a boolean indicating if the session may be resumable
	Resumable bool `json:"d"`
}

type ReceiveRateLimitedData struct {
	// Rate limited event data structure (if any)
}

// =============================================================================
// SEND EVENTS (Client to Server)
// =============================================================================

type SendIdentifyData struct {
	Token          string                       `json:"token"`
	Properties     IdentifyConnectionProperties `json:"properties"`
	Compress       *bool                        `json:"compress,omitempty"`
	LargeThreshold *int                         `json:"large_threshold,omitempty"`
	Shard          *[2]int                      `json:"shard,omitempty"`
	Presence       *PresenceUpdate              `json:"presence,omitempty"`
	Intents        int                          `json:"intents"`
}

type IdentifyConnectionProperties struct {
	OS      string `json:"os"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
}

type SendResumeData struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       int    `json:"seq"`
}

type SendHeartbeatData = int

type SendRequestGuildMembersData struct {
	GuildID   Snowflake   `json:"guild_id"`
	Query     *string     `json:"query,omitempty"`
	Limit     int         `json:"limit"`
	Presences *bool       `json:"presences,omitempty"`
	UserIDs   []Snowflake `json:"user_ids,omitempty"`
	Nonce     *string     `json:"nonce,omitempty"`
}

type SendRequestSoundboardSoundsData struct {
	GuildIDs []Snowflake `json:"guild_ids"`
}

type SendUpdateVoiceStateData struct {
	GuildID   Snowflake  `json:"guild_id"`
	ChannelID *Snowflake `json:"channel_id"`
	SelfMute  bool       `json:"self_mute"`
	SelfDeaf  bool       `json:"self_deaf"`
}

type SendUpdatePresenceData struct {
	Since      *int               `json:"since"`
	Activities []PresenceActivity `json:"activities"`
	Status     string             `json:"status"`
	AFK        bool               `json:"afk"`
}

type PresenceActivity struct {
	Name          string              `json:"name"`
	Type          int                 `json:"type"`
	URL           *string             `json:"url,omitempty"`
	CreatedAt     int64               `json:"created_at"`
	Timestamps    *ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationID *Snowflake          `json:"application_id,omitempty"`
	Details       *string             `json:"details,omitempty"`
	State         *string             `json:"state,omitempty"`
	Emoji         *ActivityEmoji      `json:"emoji,omitempty"`
	Party         *ActivityParty      `json:"party,omitempty"`
	Assets        *ActivityAssets     `json:"assets,omitempty"`
	Secrets       *ActivitySecrets    `json:"secrets,omitempty"`
	Instance      *bool               `json:"instance,omitempty"`
	Flags         *int                `json:"flags,omitempty"`
	Buttons       []ActivityButton    `json:"buttons,omitempty"`
}

type ActivityTimestamps struct {
	Start *int64 `json:"start,omitempty"`
	End   *int64 `json:"end,omitempty"`
}

type ActivityEmoji struct {
	Name     string     `json:"name"`
	ID       *Snowflake `json:"id,omitempty"`
	Animated *bool      `json:"animated,omitempty"`
}

type ActivityParty struct {
	ID   *string `json:"id,omitempty"`
	Size *[2]int `json:"size,omitempty"`
}

type ActivityAssets struct {
	LargeImage *string `json:"large_image,omitempty"`
	LargeText  *string `json:"large_text,omitempty"`
	SmallImage *string `json:"small_image,omitempty"`
	SmallText  *string `json:"small_text,omitempty"`
}

type ActivitySecrets struct {
	Join     *string `json:"join,omitempty"`
	Spectate *string `json:"spectate,omitempty"`
	Match    *string `json:"match,omitempty"`
}

type ActivityButton struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

type PresenceUpdate struct {
	Since      *int               `json:"since"`
	Activities []PresenceActivity `json:"activities"`
	Status     string             `json:"status"`
	AFK        bool               `json:"afk"`
}
