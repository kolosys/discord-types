// Package discord provides core Discord API types for Go.
//
// This package contains fundamental types used throughout the Discord API,
// including Snowflake IDs, permissions, and formatting patterns.
package discord

import (
	"regexp"
	"strconv"
	"time"
)

// Snowflake represents a Discord Snowflake ID.
//
// Snowflakes are unique IDs used throughout Discord's API.
// They are 64-bit integers represented as strings in JSON.
//
// See: https://discord.com/developers/docs/reference#snowflakes
type Snowflake string

// String returns the string representation of the Snowflake.
func (s Snowflake) String() string {
	return string(s)
}

// Int64 converts the Snowflake to an int64.
// Returns an error if the Snowflake is not a valid integer.
func (s Snowflake) Int64() (int64, error) {
	return strconv.ParseInt(string(s), 10, 64)
}

// Time extracts the timestamp from the Snowflake.
// Returns the time the Snowflake was created.
func (s Snowflake) Time() (time.Time, error) {
	id, err := s.Int64()
	if err != nil {
		return time.Time{}, err
	}

	// Discord Epoch (January 1, 2015)
	const discordEpoch = 1420070400000

	// Extract timestamp from Snowflake
	timestamp := (id >> 22) + discordEpoch

	return time.Unix(timestamp/1000, (timestamp%1000)*1000000), nil
}

// IsValid checks if the Snowflake is a valid Discord ID.
func (s Snowflake) IsValid() bool {
	if len(s) < 17 || len(s) > 20 {
		return false
	}

	_, err := strconv.ParseUint(string(s), 10, 64)
	return err == nil
}

// Permissions represents Discord permissions as a string.
//
// Permissions in Discord are represented as a bitfield stored as a string.
//
// See: https://discord.com/developers/docs/topics/permissions
type Permissions string

// String returns the string representation of the permissions.
func (p Permissions) String() string {
	return string(p)
}

// Int64 converts the permissions to an int64 for bitwise operations.
func (p Permissions) Int64() (int64, error) {
	return strconv.ParseInt(string(p), 10, 64)
}

// Has checks if the permissions include the specified permission bit.
func (p Permissions) Has(permission int64) (bool, error) {
	perms, err := p.Int64()
	if err != nil {
		return false, err
	}
	return perms&permission != 0, nil
}

// FormattingPatterns contains regular expressions for parsing Discord message formatting.
//
// See: https://discord.com/developers/docs/reference#message-formatting-formats
var FormattingPatterns = struct {
	// User matches a user mention, strictly without a nickname.
	// The id group property is present on the match result.
	User *regexp.Regexp

	// UserWithNickname matches a user mention, strictly with a nickname.
	// The id group property is present on the match result.
	//
	// Deprecated: Passing ! in user mentions is no longer necessary/supported,
	// and future message contents won't have it.
	UserWithNickname *regexp.Regexp

	// UserWithOptionalNickname matches a user mention, with or without a nickname.
	// The id group property is present on the match result.
	//
	// Deprecated: Passing ! in user mentions is no longer necessary/supported,
	// and future message contents won't have it.
	UserWithOptionalNickname *regexp.Regexp

	// Channel matches a channel mention.
	// The id group property is present on the match result.
	Channel *regexp.Regexp

	// Role matches a role mention.
	// The id group property is present on the match result.
	Role *regexp.Regexp

	// SlashCommand matches a slash command mention.
	// The fullName (possibly including name, subcommandOrGroup and subcommand)
	// and id group properties are present on the match result.
	SlashCommand *regexp.Regexp

	// Emoji matches a custom emoji, either static or animated.
	// The animated, name and id group properties are present on the match result.
	Emoji *regexp.Regexp

	// AnimatedEmoji matches strictly an animated custom emoji.
	// The animated, name and id group properties are present on the match result.
	AnimatedEmoji *regexp.Regexp

	// StaticEmoji matches strictly a static custom emoji.
	// The name and id group properties are present on the match result.
	StaticEmoji *regexp.Regexp

	// Timestamp matches a timestamp, either default or custom styled.
	// The timestamp and style group properties are present on the match result.
	Timestamp *regexp.Regexp

	// DefaultStyledTimestamp matches strictly default styled timestamps.
	// The timestamp group property is present on the match result.
	DefaultStyledTimestamp *regexp.Regexp

	// StyledTimestamp matches strictly custom styled timestamps.
	// The timestamp and style group properties are present on the match result.
	StyledTimestamp *regexp.Regexp

	// GuildNavigation matches a guild navigation mention.
	// The type group property is present on the match result.
	GuildNavigation *regexp.Regexp

	// LinkedRole matches a linked role mention.
	// The id group property is present on the match result.
	LinkedRole *regexp.Regexp
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

// NewSnowflake creates a new Snowflake from a string.
func NewSnowflake(s string) Snowflake {
	return Snowflake(s)
}

// NewSnowflakeFromInt64 creates a new Snowflake from an int64.
func NewSnowflakeFromInt64(i int64) Snowflake {
	return Snowflake(strconv.FormatInt(i, 10))
}

// NewPermissions creates new Permissions from a string.
func NewPermissions(s string) Permissions {
	return Permissions(s)
}

// NewPermissionsFromInt64 creates new Permissions from an int64.
func NewPermissionsFromInt64(i int64) Permissions {
	return Permissions(strconv.FormatInt(i, 10))
}
