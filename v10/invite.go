package v10

type Invite struct {
	// Type is the type of invite.
	Type InviteType `json:"type"`

	// Code is the invite code (unique ID).
	Code string `json:"code"`

	// Guild is the guild this invite is for.
	Guild *Guild `json:"guild,omitempty"`

	// Channel is the channel this invite is for.
	Channel *PartialChannel `json:"channel,omitempty"`

	// Inviter is the user who created the invite.
	Inviter *User `json:"inviter,omitempty"`

	// TargetType is the type of target for this voice channel invite.
	TargetType *InviteTargetType `json:"target_type,omitempty"`

	// TargetUser is the user whose stream to display for this voice channel stream invite.
	TargetUser *User `json:"target_user,omitempty"`

	// TargetApplication is the embedded application to open for this voice channel embedded application invite.
	TargetApplication *PartialApplication `json:"target_application,omitempty"`

	// ApproximatePresenceCount is the approximate count of online members, returned from the GET /invites/<code> endpoint when with_count is true.
	ApproximatePresenceCount *int `json:"approximate_presence_count,omitempty"`

	// ApproximateMemberCount is the approximate count of total members, returned from the GET /invites/<code> endpoint when with_count is true.
	ApproximateMemberCount *int `json:"approximate_member_count,omitempty"`

	// ExpiresAt is the expiration date of this invite.
	ExpiresAt *string `json:"expires_at,omitempty"`

	// GuildScheduledEvent is the guild scheduled event data, only included if guild_scheduled_event_id contains a valid guild scheduled event id.
	GuildScheduledEvent *GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`

	// Flags are guild invite flags for guild invites.
	Flags *InviteFlag `json:"flags,omitempty"`
}

type InviteType int

const (
	InviteTypeGuild InviteType = iota
	InviteTypeDM
	InviteTypeFriend
)

type InviteTargetType int

const (
	InviteTargetTypeStream InviteTargetType = iota + 1
	InviteTargetTypeEmbeddedApplication
)

type InviteFlag int

const (
	InviteFlagGuest = 1 << 0 // this invite is a guest invite for a voice channel
)
