// Package discord provides Discord API types for Go.
//
// This package contains core Discord API types and utilities for building
// Discord bots and applications in Go. It provides comprehensive type
// definitions for the Discord API v10.
package v10

import (
	"regexp"
)

// This file contains library-wide constants, metadata, and core types.

// DiscordAPIVersion is the Discord API version this library targets.
const DiscordAPIVersion = "10"

// Common Discord constants
const (
	// MaxMessageLength is the maximum length of a Discord message.
	MaxMessageLength = 2000

	// MaxEmbedLength is the maximum total length of all embed content.
	MaxEmbedLength = 6000

	// MaxEmbedTitleLength is the maximum length of an embed title.
	MaxEmbedTitleLength = 256

	// MaxEmbedDescriptionLength is the maximum length of an embed description.
	MaxEmbedDescriptionLength = 4096

	// MaxEmbedFieldNameLength is the maximum length of an embed field name.
	MaxEmbedFieldNameLength = 256

	// MaxEmbedFieldValueLength is the maximum length of an embed field value.
	MaxEmbedFieldValueLength = 1024

	// MaxEmbedFooterTextLength is the maximum length of an embed footer text.
	MaxEmbedFooterTextLength = 2048

	// MaxEmbedAuthorNameLength is the maximum length of an embed author name.
	MaxEmbedAuthorNameLength = 256

	// MaxEmbedFields is the maximum number of fields in an embed.
	MaxEmbedFields = 25

	// MaxEmbedsPerMessage is the maximum number of embeds per message.
	MaxEmbedsPerMessage = 10

	// MaxUsernameLength is the maximum length of a username.
	MaxUsernameLength = 32

	// MinUsernameLength is the minimum length of a username.
	MinUsernameLength = 2

	// MaxNicknameLength is the maximum length of a nickname.
	MaxNicknameLength = 32

	// MaxGuildNameLength is the maximum length of a guild name.
	MaxGuildNameLength = 100

	// MinGuildNameLength is the minimum length of a guild name.
	MinGuildNameLength = 2

	// MaxChannelNameLength is the maximum length of a channel name.
	MaxChannelNameLength = 100

	// MinChannelNameLength is the minimum length of a channel name.
	MinChannelNameLength = 1

	// MaxRoleNameLength is the maximum length of a role name.
	MaxRoleNameLength = 100

	// MinRoleNameLength is the minimum length of a role name.
	MinRoleNameLength = 1

	// MaxReasonLength is the maximum length of an audit log reason.
	MaxReasonLength = 512

	// MaxApplicationCommandNameLength is the maximum length of an application command name.
	MaxApplicationCommandNameLength = 32

	// MinApplicationCommandNameLength is the minimum length of an application command name.
	MinApplicationCommandNameLength = 1

	// MaxApplicationCommandDescriptionLength is the maximum length of an application command description.
	MaxApplicationCommandDescriptionLength = 100

	// MinApplicationCommandDescriptionLength is the minimum length of an application command description.
	MinApplicationCommandDescriptionLength = 1

	// MaxSlashCommandOptions is the maximum number of options for a slash command.
	MaxSlashCommandOptions = 25

	// MaxSlashCommandChoices is the maximum number of choices for a slash command option.
	MaxSlashCommandChoices = 25

	// MaxButtonsPerActionRow is the maximum number of buttons per action row.
	MaxButtonsPerActionRow = 5

	// MaxActionRowsPerMessage is the maximum number of action rows per message.
	MaxActionRowsPerMessage = 5

	// MaxSelectMenuOptions is the maximum number of options in a select menu.
	MaxSelectMenuOptions = 25

	// MaxFileSize is the maximum file size for uploads (8MB).
	MaxFileSize = 8 * 1024 * 1024

	// MaxFileSizeNitro is the maximum file size for Nitro users (50MB).
	MaxFileSizeNitro = 50 * 1024 * 1024

	// MaxGuildMembers is the default maximum number of guild members.
	MaxGuildMembers = 250000

	// MaxGuildMembersWithMoreMembers is the maximum number of guild members with MORE_MEMBERS feature.
	MaxGuildMembersWithMoreMembers = 500000

	// RateLimitGlobal is the global rate limit bucket.
	RateLimitGlobal = "global"

	// DiscordEpoch is the Discord epoch timestamp (January 1, 2015).
	DiscordEpoch = 1420070400000
)

// Common colors for embeds
type Color int

const (
	ColorDefault         = 0x000000
	ColorAqua            = 0x1ABC9C
	ColorGreen           = 0x2ECC71
	ColorBlue            = 0x3498DB
	ColorYellow          = 0xF1C40F
	ColorPurple          = 0x9B59B6
	ColorRed             = 0xE74C3C
	ColorOrange          = 0xE67E22
	ColorGrey            = 0x95A5A6
	ColorNavy            = 0x34495E
	ColorDarkAqua        = 0x11806A
	ColorDarkGreen       = 0x1F8B4C
	ColorDarkBlue        = 0x206694
	ColorDarkPurple      = 0x71368A
	ColorDarkRed         = 0xAD1457
	ColorDarkOrange      = 0xA84300
	ColorDarkGrey        = 0x607D8B
	ColorDarkerGrey      = 0x546E7A
	ColorLightGrey       = 0xBCC0C0
	ColorDarkNavy        = 0x2C3E50
	ColorBlurple         = 0x5865F2
	ColorGreyple         = 0x99AAB5
	ColorDarkButNotBlack = 0x2C2F33
	ColorNotQuiteBlack   = 0x23272A
)

// =============================================================================
// CORE TYPES
// =============================================================================

// FormattingPatterns contains regular expressions for parsing Discord message formatting.
var FormattingPatterns = struct {
	User                     *regexp.Regexp
	UserWithNickname         *regexp.Regexp
	UserWithOptionalNickname *regexp.Regexp
	Channel                  *regexp.Regexp
	Role                     *regexp.Regexp
	SlashCommand             *regexp.Regexp
	Emoji                    *regexp.Regexp
	StaticEmoji              *regexp.Regexp
	AnimatedEmoji            *regexp.Regexp
	Timestamp                *regexp.Regexp
	DefaultStyledTimestamp   *regexp.Regexp
	StyledTimestamp          *regexp.Regexp
	GuildNavigation          *regexp.Regexp
	LinkedRole               *regexp.Regexp
}{
	User:                     regexp.MustCompile(`<@(?P<id>\d{17,20})>`),
	UserWithNickname:         regexp.MustCompile(`<@!(?P<id>\d{17,20})>`),
	UserWithOptionalNickname: regexp.MustCompile(`<@!?(?P<id>\d{17,20})>`),
	Channel:                  regexp.MustCompile(`<#(?P<id>\d{17,20})>`),
	Role:                     regexp.MustCompile(`<@&(?P<id>\d{17,20})>`),
	SlashCommand:             regexp.MustCompile(`<\/(?P<fullName>(?P<name>[-_\p{L}\p{N}\p{Devanagari}\p{Thai}]{1,32})(?: (?P<subcommandOrGroup>[-_\p{L}\p{N}\p{Devanagari}\p{Thai}]{1,32}))?(?: (?P<subcommand>[-_\p{L}\p{N}\p{Devanagari}\p{Thai}]{1,32}))?):(?P<id>\d{17,20})>`),
	Emoji:                    regexp.MustCompile(`<(?P<animated>a)?:(?P<name>\w{2,32}):(?P<id>\d{17,20})>`),
	AnimatedEmoji:            regexp.MustCompile(`<(?P<animated>a):(?P<name>\w{2,32}):(?P<id>\d{17,20})>`),
	StaticEmoji:              regexp.MustCompile(`<:(?P<name>\w{2,32}):(?P<id>\d{17,20})>`),
	Timestamp:                regexp.MustCompile(`<t:(?P<timestamp>-?\d{1,13})(?::(?P<style>[DFRTdft]))?>`),
	DefaultStyledTimestamp:   regexp.MustCompile(`<t:(?P<timestamp>-?\d{1,13})>`),
	StyledTimestamp:          regexp.MustCompile(`<t:(?P<timestamp>-?\d{1,13}):(?P<style>[DFRTdft])>`),
	GuildNavigation:          regexp.MustCompile(`<id:(?P<type>customize|browse|guide|linked-roles)>`),
	LinkedRole:               regexp.MustCompile(`<id:linked-roles:(?P<id>\d{17,20})>`),
}
