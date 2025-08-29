package payloads

import (
	"time"

	"github.com/kolosys/discord-types/discord"
)

// GuildMember represents a guild member
// https://discord.com/developers/docs/resources/guild#guild-member-object
type GuildMember struct {
	User                       *User                `json:"user,omitempty"`
	Nick                       *string              `json:"nick,omitempty"`
	Avatar                     *string              `json:"avatar,omitempty"`
	Roles                      []discord.Snowflake  `json:"roles"`
	JoinedAt                   time.Time            `json:"joined_at"`
	PremiumSince               *time.Time           `json:"premium_since,omitempty"`
	Deaf                       *bool                `json:"deaf,omitempty"`
	Mute                       *bool                `json:"mute,omitempty"`
	Flags                      *GuildMemberFlags    `json:"flags,omitempty"`
	Pending                    *bool                `json:"pending,omitempty"`
	Permissions                *discord.Permissions `json:"permissions,omitempty"`
	CommunicationDisabledUntil *time.Time           `json:"communication_disabled_until,omitempty"`
}

// GuildMemberFlags represents guild member flags
type GuildMemberFlags int

const (
	GuildMemberFlagDidRejoin            GuildMemberFlags = 1 << 0
	GuildMemberFlagCompletedOnboarding  GuildMemberFlags = 1 << 1
	GuildMemberFlagBypassesVerification GuildMemberFlags = 1 << 2
	GuildMemberFlagStartedOnboarding    GuildMemberFlags = 1 << 3
)

// BaseGuildMember represents the base guild member structure
type BaseGuildMember struct {
	User                       *User               `json:"user,omitempty"`
	Nick                       *string             `json:"nick,omitempty"`
	Avatar                     *string             `json:"avatar,omitempty"`
	Roles                      []discord.Snowflake `json:"roles"`
	JoinedAt                   time.Time           `json:"joined_at"`
	PremiumSince               *time.Time          `json:"premium_since,omitempty"`
	Deaf                       *bool               `json:"deaf,omitempty"`
	Mute                       *bool               `json:"mute,omitempty"`
	Pending                    *bool               `json:"pending,omitempty"`
	CommunicationDisabledUntil *time.Time          `json:"communication_disabled_until,omitempty"`
	Flags                      *GuildMemberFlags   `json:"flags,omitempty"`
}
