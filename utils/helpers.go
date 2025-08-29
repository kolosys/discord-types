// Package utils provides utility functions and helpers for Discord API types.
package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kolosys/discord-types/discord"
)

// StringPtr returns a pointer to the given string.
// This is useful for setting optional string fields in structs.
func StringPtr(s string) *string {
	return &s
}

// IntPtr returns a pointer to the given int.
// This is useful for setting optional int fields in structs.
func IntPtr(i int) *int {
	return &i
}

// BoolPtr returns a pointer to the given bool.
// This is useful for setting optional bool fields in structs.
func BoolPtr(b bool) *bool {
	return &b
}

// Int64Ptr returns a pointer to the given int64.
// This is useful for setting optional int64 fields in structs.
func Int64Ptr(i int64) *int64 {
	return &i
}

// Float64Ptr returns a pointer to the given float64.
// This is useful for setting optional float64 fields in structs.
func Float64Ptr(f float64) *float64 {
	return &f
}

// SnowflakePtr returns a pointer to the given Snowflake.
// This is useful for setting optional Snowflake fields in structs.
func SnowflakePtr(s discord.Snowflake) *discord.Snowflake {
	return &s
}

// PermissionsPtr returns a pointer to the given Permissions.
// This is useful for setting optional Permissions fields in structs.
func PermissionsPtr(p discord.Permissions) *discord.Permissions {
	return &p
}

// DerefString safely dereferences a string pointer, returning empty string if nil.
func DerefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// DerefInt safely dereferences an int pointer, returning 0 if nil.
func DerefInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// DerefBool safely dereferences a bool pointer, returning false if nil.
func DerefBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// DerefSnowflake safely dereferences a Snowflake pointer, returning empty Snowflake if nil.
func DerefSnowflake(s *discord.Snowflake) discord.Snowflake {
	if s == nil {
		return ""
	}
	return *s
}

// SnowflakeFromTime creates a Snowflake from a time.Time.
// This generates a Snowflake-like ID based on the timestamp.
// Note: This will not include worker/process IDs and increment, so it's only
// suitable for approximate time-based sorting, not as actual Discord IDs.
func SnowflakeFromTime(t time.Time) discord.Snowflake {
	// Discord Epoch (January 1, 2015)
	const discordEpoch = 1420070400000

	// Convert to milliseconds since Discord epoch
	timestamp := t.UnixMilli() - discordEpoch

	// Shift to make room for worker/process/increment bits (22 bits)
	id := timestamp << 22

	return discord.NewSnowflakeFromInt64(id)
}

// TimeFromSnowflake extracts the timestamp from a Snowflake.
func TimeFromSnowflake(s discord.Snowflake) (time.Time, error) {
	return s.Time()
}

// ValidateSnowflake checks if a string is a valid Snowflake format.
func ValidateSnowflake(s string) bool {
	snowflake := discord.Snowflake(s)
	return snowflake.IsValid()
}

// CalculatePermissions calculates permissions based on base permissions and overwrites.
// This is a simplified version - full permission calculation involves roles, channel overwrites, etc.
func CalculatePermissions(basePermissions discord.Permissions, overwrites ...discord.Permissions) (discord.Permissions, error) {
	base, err := basePermissions.Int64()
	if err != nil {
		return "", err
	}

	result := base

	for _, overwrite := range overwrites {
		over, err := overwrite.Int64()
		if err != nil {
			return "", err
		}
		result |= over
	}

	return discord.NewPermissionsFromInt64(result), nil
}

// HasPermission checks if the given permissions include a specific permission bit.
func HasPermission(permissions discord.Permissions, permission int64) (bool, error) {
	return permissions.Has(permission)
}

// FormatMention creates a user mention string from a user ID.
func FormatMention(userID discord.Snowflake) string {
	return fmt.Sprintf("<@%s>", userID)
}

// FormatChannelMention creates a channel mention string from a channel ID.
func FormatChannelMention(channelID discord.Snowflake) string {
	return fmt.Sprintf("<#%s>", channelID)
}

// FormatRoleMention creates a role mention string from a role ID.
func FormatRoleMention(roleID discord.Snowflake) string {
	return fmt.Sprintf("<@&%s>", roleID)
}

// FormatTimestamp creates a Discord timestamp string from a time.Time.
// The style parameter determines how the timestamp is displayed.
// Valid styles: t, T, d, D, f, F, R
func FormatTimestamp(t time.Time, style string) string {
	timestamp := t.Unix()
	if style == "" {
		return fmt.Sprintf("<t:%d>", timestamp)
	}
	return fmt.Sprintf("<t:%d:%s>", timestamp, style)
}

// ParseMention attempts to parse a mention string and return the ID.
// Returns the ID and the type of mention (user, channel, role, etc.).
func ParseMention(mention string) (discord.Snowflake, string, error) {
	// Check user mention
	if matches := discord.FormattingPatterns.User.FindStringSubmatch(mention); len(matches) > 1 {
		return discord.Snowflake(matches[1]), "user", nil
	}

	// Check channel mention
	if matches := discord.FormattingPatterns.Channel.FindStringSubmatch(mention); len(matches) > 1 {
		return discord.Snowflake(matches[1]), "channel", nil
	}

	// Check role mention
	if matches := discord.FormattingPatterns.Role.FindStringSubmatch(mention); len(matches) > 1 {
		return discord.Snowflake(matches[1]), "role", nil
	}

	return "", "", fmt.Errorf("invalid mention format: %s", mention)
}

// CDNImageSize represents valid CDN image sizes.
type CDNImageSize int

// Valid CDN image sizes
const (
	CDNImageSize16   CDNImageSize = 16
	CDNImageSize32   CDNImageSize = 32
	CDNImageSize64   CDNImageSize = 64
	CDNImageSize128  CDNImageSize = 128
	CDNImageSize256  CDNImageSize = 256
	CDNImageSize512  CDNImageSize = 512
	CDNImageSize1024 CDNImageSize = 1024
	CDNImageSize2048 CDNImageSize = 2048
	CDNImageSize4096 CDNImageSize = 4096
)

// CDNImageFormat represents valid CDN image formats.
type CDNImageFormat string

// Valid CDN image formats
const (
	CDNImageFormatPNG  CDNImageFormat = "png"
	CDNImageFormatJPEG CDNImageFormat = "jpeg"
	CDNImageFormatWebP CDNImageFormat = "webp"
	CDNImageFormatGIF  CDNImageFormat = "gif"
)

// BuildCDNURL builds a CDN URL for Discord assets.
func BuildCDNURL(endpoint, hash string, format CDNImageFormat, size CDNImageSize) string {
	baseURL := "https://cdn.discordapp.com"
	url := fmt.Sprintf("%s%s.%s", baseURL, endpoint, format)

	if size > 0 {
		url += fmt.Sprintf("?size=%d", size)
	}

	return url
}

// BuildUserAvatarURL builds a CDN URL for a user's avatar.
func BuildUserAvatarURL(userID discord.Snowflake, avatarHash string, format CDNImageFormat, size CDNImageSize) string {
	endpoint := fmt.Sprintf("/avatars/%s/%s", userID, avatarHash)
	return BuildCDNURL(endpoint, avatarHash, format, size)
}

// BuildGuildIconURL builds a CDN URL for a guild's icon.
func BuildGuildIconURL(guildID discord.Snowflake, iconHash string, format CDNImageFormat, size CDNImageSize) string {
	endpoint := fmt.Sprintf("/icons/%s/%s", guildID, iconHash)
	return BuildCDNURL(endpoint, iconHash, format, size)
}

// BuildDefaultUserAvatarURL builds a CDN URL for a default user avatar.
func BuildDefaultUserAvatarURL(discriminator string) string {
	// For users on the new username system, use (user.id >> 22) % 6
	// For users on the legacy username system, use discriminator % 5
	discrim, err := strconv.Atoi(discriminator)
	if err != nil {
		discrim = 0
	}

	index := discrim % 5
	return fmt.Sprintf("https://cdn.discordapp.com/embed/avatars/%d.png", index)
}

// ColorFromHex converts a hex color string to an integer.
func ColorFromHex(hex string) (int, error) {
	// Remove # prefix if present
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	// Parse hex string to integer
	color, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid hex color: %s", hex)
	}

	return int(color), nil
}

// ColorToHex converts an integer color to a hex string.
func ColorToHex(color int) string {
	return fmt.Sprintf("#%06X", color)
}
