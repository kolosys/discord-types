package utils

import (
	"testing"
	"time"

	"github.com/kolosys/discord-types/discord"
)

func TestStringPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: StringPtr(""),
		},
		{
			name:     "Non-empty string",
			input:    "test",
			expected: StringPtr("test"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringPtr(tt.input)
			if result == nil {
				t.Error("StringPtr returned nil")
				return
			}
			if *result != tt.input {
				t.Errorf("StringPtr() = %v, want %v", *result, tt.input)
			}
		})
	}
}

func TestIntPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected *int
	}{
		{
			name:     "Zero",
			input:    0,
			expected: IntPtr(0),
		},
		{
			name:     "Positive number",
			input:    42,
			expected: IntPtr(42),
		},
		{
			name:     "Negative number",
			input:    -42,
			expected: IntPtr(-42),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntPtr(tt.input)
			if result == nil {
				t.Error("IntPtr returned nil")
				return
			}
			if *result != tt.input {
				t.Errorf("IntPtr() = %v, want %v", *result, tt.input)
			}
		})
	}
}

func TestBoolPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected *bool
	}{
		{
			name:     "True",
			input:    true,
			expected: BoolPtr(true),
		},
		{
			name:     "False",
			input:    false,
			expected: BoolPtr(false),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BoolPtr(tt.input)
			if result == nil {
				t.Error("BoolPtr returned nil")
				return
			}
			if *result != tt.input {
				t.Errorf("BoolPtr() = %v, want %v", *result, tt.input)
			}
		})
	}
}

func TestInt64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected *int64
	}{
		{
			name:     "Zero",
			input:    0,
			expected: Int64Ptr(0),
		},
		{
			name:     "Positive number",
			input:    42,
			expected: Int64Ptr(42),
		},
		{
			name:     "Large number",
			input:    1234567890123456789,
			expected: Int64Ptr(1234567890123456789),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int64Ptr(tt.input)
			if result == nil {
				t.Error("Int64Ptr returned nil")
				return
			}
			if *result != tt.input {
				t.Errorf("Int64Ptr() = %v, want %v", *result, tt.input)
			}
		})
	}
}

func TestFloat64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected *float64
	}{
		{
			name:     "Zero",
			input:    0.0,
			expected: Float64Ptr(0.0),
		},
		{
			name:     "Positive number",
			input:    42.5,
			expected: Float64Ptr(42.5),
		},
		{
			name:     "Negative number",
			input:    -42.5,
			expected: Float64Ptr(-42.5),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float64Ptr(tt.input)
			if result == nil {
				t.Error("Float64Ptr returned nil")
				return
			}
			if *result != tt.input {
				t.Errorf("Float64Ptr() = %v, want %v", *result, tt.input)
			}
		})
	}
}

func TestFormatMention(t *testing.T) {
	tests := []struct {
		name     string
		userID   discord.Snowflake
		expected string
	}{
		{
			name:     "Valid user ID",
			userID:   "123456789012345678",
			expected: "<@123456789012345678>",
		},
		{
			name:     "Another valid user ID",
			userID:   "987654321098765432",
			expected: "<@987654321098765432>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatMention(tt.userID)
			if result != tt.expected {
				t.Errorf("FormatMention() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatChannelMention(t *testing.T) {
	tests := []struct {
		name      string
		channelID discord.Snowflake
		expected  string
	}{
		{
			name:      "Valid channel ID",
			channelID: "123456789012345678",
			expected:  "<#123456789012345678>",
		},
		{
			name:      "Another valid channel ID",
			channelID: "987654321098765432",
			expected:  "<#987654321098765432>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatChannelMention(tt.channelID)
			if result != tt.expected {
				t.Errorf("FormatChannelMention() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatRoleMention(t *testing.T) {
	tests := []struct {
		name     string
		roleID   discord.Snowflake
		expected string
	}{
		{
			name:     "Valid role ID",
			roleID:   "123456789012345678",
			expected: "<@&123456789012345678>",
		},
		{
			name:     "Another valid role ID",
			roleID:   "987654321098765432",
			expected: "<@&987654321098765432>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatRoleMention(tt.roleID)
			if result != tt.expected {
				t.Errorf("FormatRoleMention() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatTimestamp(t *testing.T) {
	tests := []struct {
		name     string
		time     time.Time
		style    string
		expected string
	}{
		{
			name:     "Default style",
			time:     time.Unix(1234567890, 0),
			style:    "",
			expected: "<t:1234567890>",
		},
		{
			name:     "Full date style",
			time:     time.Unix(1234567890, 0),
			style:    "F",
			expected: "<t:1234567890:F>",
		},
		{
			name:     "Relative time style",
			time:     time.Unix(1234567890, 0),
			style:    "R",
			expected: "<t:1234567890:R>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatTimestamp(tt.time, tt.style)
			if result != tt.expected {
				t.Errorf("FormatTimestamp() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestParseMention(t *testing.T) {
	tests := []struct {
		name        string
		mention     string
		expectedID  discord.Snowflake
		expectedType string
		expectError bool
	}{
		{
			name:        "Valid user mention",
			mention:     "<@123456789012345678>",
			expectedID:  "123456789012345678",
			expectedType: "user",
			expectError: false,
		},
		{
			name:        "Valid channel mention",
			mention:     "<#123456789012345678>",
			expectedID:  "123456789012345678",
			expectedType: "channel",
			expectError: false,
		},
		{
			name:        "Valid role mention",
			mention:     "<@&123456789012345678>",
			expectedID:  "123456789012345678",
			expectedType: "role",
			expectError: false,
		},
		{
			name:        "Invalid mention",
			mention:     "invalid",
			expectedID:  "",
			expectedType: "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, mentionType, err := ParseMention(tt.mention)
			if (err != nil) != tt.expectError {
				t.Errorf("ParseMention() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError {
				if id != tt.expectedID {
					t.Errorf("ParseMention() id = %v, want %v", id, tt.expectedID)
				}
				if mentionType != tt.expectedType {
					t.Errorf("ParseMention() type = %v, want %v", mentionType, tt.expectedType)
				}
			}
		})
	}
}

func TestBuildCDNURL(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		hash     string
		format   CDNImageFormat
		size     CDNImageSize
		expected string
	}{
		{
			name:     "Default format and size",
			endpoint: "/test",
			hash:     "test_hash",
			format:   CDNImageFormatPNG,
			size:     0,
			expected: "https://cdn.discordapp.com/test.png",
		},
		{
			name:     "With size",
			endpoint: "/test",
			hash:     "test_hash",
			format:   CDNImageFormatPNG,
			size:     CDNImageSize1024,
			expected: "https://cdn.discordapp.com/test.png?size=1024",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildCDNURL(tt.endpoint, tt.hash, tt.format, tt.size)
			if result != tt.expected {
				t.Errorf("BuildCDNURL() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBuildUserAvatarURL(t *testing.T) {
	tests := []struct {
		name       string
		userID     discord.Snowflake
		avatarHash string
		format     CDNImageFormat
		size       CDNImageSize
		expected   string
	}{
		{
			name:       "Valid avatar URL",
			userID:     "123456789012345678",
			avatarHash: "test_hash",
			format:     CDNImageFormatPNG,
			size:       CDNImageSize1024,
			expected:   "https://cdn.discordapp.com/avatars/123456789012345678/test_hash.png?size=1024",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildUserAvatarURL(tt.userID, tt.avatarHash, tt.format, tt.size)
			if result != tt.expected {
				t.Errorf("BuildUserAvatarURL() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBuildGuildIconURL(t *testing.T) {
	tests := []struct {
		name      string
		guildID   discord.Snowflake
		iconHash  string
		format    CDNImageFormat
		size      CDNImageSize
		expected  string
	}{
		{
			name:      "Valid guild icon URL",
			guildID:   "123456789012345678",
			iconHash:  "test_hash",
			format:    CDNImageFormatPNG,
			size:      CDNImageSize1024,
			expected:  "https://cdn.discordapp.com/icons/123456789012345678/test_hash.png?size=1024",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildGuildIconURL(tt.guildID, tt.iconHash, tt.format, tt.size)
			if result != tt.expected {
				t.Errorf("BuildGuildIconURL() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestBuildDefaultUserAvatarURL(t *testing.T) {
	tests := []struct {
		name          string
		discriminator string
		expected      string
	}{
		{
			name:          "Discriminator 0",
			discriminator: "0000",
			expected:      "https://cdn.discordapp.com/embed/avatars/0.png",
		},
		{
			name:          "Discriminator 1",
			discriminator: "0001",
			expected:      "https://cdn.discordapp.com/embed/avatars/1.png",
		},
		{
			name:          "Discriminator 4",
			discriminator: "0004",
			expected:      "https://cdn.discordapp.com/embed/avatars/4.png",
		},
		{
			name:          "Discriminator 5",
			discriminator: "0005",
			expected:      "https://cdn.discordapp.com/embed/avatars/0.png",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildDefaultUserAvatarURL(tt.discriminator)
			if result != tt.expected {
				t.Errorf("BuildDefaultUserAvatarURL() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestColorFromHex(t *testing.T) {
	tests := []struct {
		name        string
		hex         string
		expected    int
		expectError bool
	}{
		{
			name:        "Valid hex color",
			hex:         "#FF0000",
			expected:    16711680, // Red
			expectError: false,
		},
		{
			name:        "Valid hex color without #",
			hex:         "00FF00",
			expected:    65280, // Green
			expectError: false,
		},
		{
			name:        "Invalid hex color",
			hex:         "invalid",
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ColorFromHex(tt.hex)
			if (err != nil) != tt.expectError {
				t.Errorf("ColorFromHex() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError && result != tt.expected {
				t.Errorf("ColorFromHex() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestColorToHex(t *testing.T) {
	tests := []struct {
		name     string
		color    int
		expected string
	}{
		{
			name:     "Red color",
			color:    16711680,
			expected: "#FF0000",
		},
		{
			name:     "Green color",
			color:    65280,
			expected: "#00FF00",
		},
		{
			name:     "Blue color",
			color:    255,
			expected: "#0000FF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ColorToHex(tt.color)
			if result != tt.expected {
				t.Errorf("ColorToHex() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestValidateSnowflake(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid snowflake",
			input:    "123456789012345678",
			expected: true,
		},
		{
			name:     "Invalid snowflake - too short",
			input:    "1234567890123456",
			expected: false,
		},
		{
			name:     "Invalid snowflake - non-numeric",
			input:    "12345678901234567a",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateSnowflake(tt.input)
			if result != tt.expected {
				t.Errorf("ValidateSnowflake() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHasPermission(t *testing.T) {
	tests := []struct {
		name        string
		permissions discord.Permissions
		permission  int64
		expected    bool
		expectError bool
	}{
		{
			name:        "Has permission",
			permissions: "1",
			permission:  1,
			expected:    true,
			expectError: false,
		},
		{
			name:        "Does not have permission",
			permissions: "2",
			permission:  1,
			expected:    false,
			expectError: false,
		},
		{
			name:        "Invalid permissions string",
			permissions: "invalid",
			permission:  1,
			expected:    false,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := HasPermission(tt.permissions, tt.permission)
			if (err != nil) != tt.expectError {
				t.Errorf("HasPermission() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError && result != tt.expected {
				t.Errorf("HasPermission() = %v, want %v", result, tt.expected)
			}
		})
	}
}
