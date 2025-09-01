package discord

import (
	"testing"
)

func TestFormattingPatterns_User(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid user mention",
			input:    "<@123456789012345678>",
			expected: true,
		},
		{
			name:     "Invalid user mention - too short",
			input:    "<@1234567890123456>",
			expected: false,
		},
		{
			name:     "Invalid user mention - non-numeric",
			input:    "<@12345678901234567a>",
			expected: false,
		},
		{
			name:     "Invalid user mention - missing brackets",
			input:    "@123456789012345678",
			expected: false,
		},
		{
			name:     "Invalid user mention - wrong format",
			input:    "<@!123456789012345678>",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.User.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("User pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_UserWithNickname(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid user mention with nickname",
			input:    "<@!123456789012345678>",
			expected: true,
		},
		{
			name:     "Invalid user mention with nickname - missing !",
			input:    "<@123456789012345678>",
			expected: false,
		},
		{
			name:     "Invalid user mention with nickname - too short",
			input:    "<@!1234567890123456>",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.UserWithNickname.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("UserWithNickname pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_Channel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid channel mention",
			input:    "<#123456789012345678>",
			expected: true,
		},
		{
			name:     "Invalid channel mention - too short",
			input:    "<#1234567890123456>",
			expected: false,
		},
		{
			name:     "Invalid channel mention - non-numeric",
			input:    "<#12345678901234567a>",
			expected: false,
		},
		{
			name:     "Invalid channel mention - missing brackets",
			input:    "#123456789012345678",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.Channel.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("Channel pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_Role(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid role mention",
			input:    "<@&123456789012345678>",
			expected: true,
		},
		{
			name:     "Invalid role mention - too short",
			input:    "<@&1234567890123456>",
			expected: false,
		},
		{
			name:     "Invalid role mention - non-numeric",
			input:    "<@&12345678901234567a>",
			expected: false,
		},
		{
			name:     "Invalid role mention - missing brackets",
			input:    "@&123456789012345678",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.Role.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("Role pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_Emoji(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid static emoji",
			input:    "<:smile:123456789012345678>",
			expected: true,
		},
		{
			name:     "Valid animated emoji",
			input:    "<a:wave:123456789012345678>",
			expected: true,
		},
		{
			name:     "Invalid emoji - too short ID",
			input:    "<:smile:1234567890123456>",
			expected: false,
		},
		{
			name:     "Invalid emoji - non-numeric ID",
			input:    "<:smile:12345678901234567a>",
			expected: false,
		},
		{
			name:     "Invalid emoji - missing brackets",
			input:    ":smile:123456789012345678",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.Emoji.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("Emoji pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_AnimatedEmoji(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid animated emoji",
			input:    "<a:wave:123456789012345678>",
			expected: true,
		},
		{
			name:     "Invalid animated emoji - static emoji",
			input:    "<:smile:123456789012345678>",
			expected: false,
		},
		{
			name:     "Invalid animated emoji - missing a",
			input:    "<:wave:123456789012345678>",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.AnimatedEmoji.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("AnimatedEmoji pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_StaticEmoji(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid static emoji",
			input:    "<:smile:123456789012345678>",
			expected: true,
		},
		{
			name:     "Invalid static emoji - animated emoji",
			input:    "<a:wave:123456789012345678>",
			expected: false,
		},
		{
			name:     "Invalid static emoji - missing colon",
			input:    "<smile:123456789012345678>",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.StaticEmoji.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("StaticEmoji pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_Timestamp(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid default timestamp",
			input:    "<t:1234567890>",
			expected: true,
		},
		{
			name:     "Valid styled timestamp",
			input:    "<t:1234567890:F>",
			expected: true,
		},
		{
			name:     "Valid negative timestamp",
			input:    "<t:-1234567890>",
			expected: true,
		},
		{
			name:     "Invalid timestamp - missing brackets",
			input:    "t:1234567890",
			expected: false,
		},
		{
			name:     "Invalid timestamp - non-numeric",
			input:    "<t:1234567890a>",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.Timestamp.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("Timestamp pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_DefaultStyledTimestamp(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid default styled timestamp",
			input:    "<t:1234567890>",
			expected: true,
		},
		{
			name:     "Invalid default styled timestamp - has style",
			input:    "<t:1234567890:F>",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.DefaultStyledTimestamp.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("DefaultStyledTimestamp pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_StyledTimestamp(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid styled timestamp",
			input:    "<t:1234567890:F>",
			expected: true,
		},
		{
			name:     "Valid styled timestamp with different style",
			input:    "<t:1234567890:d>",
			expected: true,
		},
		{
			name:     "Invalid styled timestamp - no style",
			input:    "<t:1234567890>",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.StyledTimestamp.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("StyledTimestamp pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_SlashCommand(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid slash command",
			input:    "</ping:123456789012345678>",
			expected: true,
		},
		{
			name:     "Valid slash command with subcommand",
			input:    "</config set:123456789012345678>",
			expected: true,
		},
		{
			name:     "Valid slash command with subcommand group",
			input:    "</config admin ban:123456789012345678>",
			expected: true,
		},
		{
			name:     "Invalid slash command - missing brackets",
			input:    "/ping:123456789012345678",
			expected: false,
		},
		{
			name:     "Invalid slash command - too short ID",
			input:    "</ping:1234567890123456>",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched := FormattingPatterns.SlashCommand.MatchString(tt.input)
			if matched != tt.expected {
				t.Errorf("SlashCommand pattern match = %v, want %v for input %s", matched, tt.expected, tt.input)
			}
		})
	}
}

func TestFormattingPatterns_ExtractID(t *testing.T) {
	// Test that we can extract IDs from matches
	userMention := "<@123456789012345678>"
	matches := FormattingPatterns.User.FindStringSubmatch(userMention)
	if len(matches) < 2 {
		t.Fatal("Expected to find ID in user mention")
	}

	extractedID := matches[1]
	expectedID := "123456789012345678"
	if extractedID != expectedID {
		t.Errorf("Extracted ID = %s, want %s", extractedID, expectedID)
	}
}
