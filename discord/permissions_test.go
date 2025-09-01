package discord

import (
	"testing"
)

func TestPermissions_String(t *testing.T) {
	tests := []struct {
		name        string
		permissions Permissions
		expected    string
	}{
		{
			name:        "Valid permissions",
			permissions: "123456789",
			expected:    "123456789",
		},
		{
			name:        "Empty permissions",
			permissions: "",
			expected:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.permissions.String(); got != tt.expected {
				t.Errorf("Permissions.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPermissions_Int64(t *testing.T) {
	tests := []struct {
		name        string
		permissions Permissions
		expected    int64
		expectError bool
	}{
		{
			name:        "Valid permissions",
			permissions: "123456789",
			expected:    123456789,
			expectError: false,
		},
		{
			name:        "Invalid permissions - non-numeric",
			permissions: "12345678a",
			expected:    0,
			expectError: true,
		},
		{
			name:        "Empty permissions",
			permissions: "",
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.permissions.Int64()
			if (err != nil) != tt.expectError {
				t.Errorf("Permissions.Int64() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if got != tt.expected {
				t.Errorf("Permissions.Int64() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPermissions_Has(t *testing.T) {
	tests := []struct {
		name        string
		permissions Permissions
		permission  int64
		expected    bool
		expectError bool
	}{
		{
			name:        "Has permission - single bit",
			permissions: "1",
			permission:  1,
			expected:    true,
			expectError: false,
		},
		{
			name:        "Does not have permission - single bit",
			permissions: "2",
			permission:  1,
			expected:    false,
			expectError: false,
		},
		{
			name:        "Has multiple permissions",
			permissions: "3", // 1 + 2
			permission:  1,
			expected:    true,
			expectError: false,
		},
		{
			name:        "Has multiple permissions - second bit",
			permissions: "3", // 1 + 2
			permission:  2,
			expected:    true,
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
			got, err := tt.permissions.Has(tt.permission)
			if (err != nil) != tt.expectError {
				t.Errorf("Permissions.Has() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if got != tt.expected {
				t.Errorf("Permissions.Has() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNewPermissions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Permissions
	}{
		{
			name:     "Valid string",
			input:    "123456789",
			expected: Permissions("123456789"),
		},
		{
			name:     "Empty string",
			input:    "",
			expected: Permissions(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPermissions(tt.input); got != tt.expected {
				t.Errorf("NewPermissions() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNewPermissionsFromInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected Permissions
	}{
		{
			name:     "Positive number",
			input:    123456789,
			expected: Permissions("123456789"),
		},
		{
			name:     "Zero",
			input:    0,
			expected: Permissions("0"),
		},
		{
			name:     "Negative number",
			input:    -123456789,
			expected: Permissions("-123456789"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPermissionsFromInt64(tt.input); got != tt.expected {
				t.Errorf("NewPermissionsFromInt64() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPermissions_BitwiseOperations(t *testing.T) {
	// Test common Discord permission combinations
	permissions := Permissions("2147483647") // All permissions (32 bits)

	// Test some common Discord permissions
	commonPermissions := []int64{
		1 << 0,  // CreateInstantInvite
		1 << 1,  // KickMembers
		1 << 2,  // BanMembers
		1 << 3,  // Administrator
		1 << 4,  // ManageChannels
		1 << 5,  // ManageGuild
		1 << 6,  // AddReactions
		1 << 7,  // ViewAuditLog
		1 << 8,  // PrioritySpeaker
		1 << 9,  // Stream
		1 << 10, // ViewChannel
		1 << 11, // SendMessages
		1 << 12, // SendTTSMessages
		1 << 13, // ManageMessages
		1 << 14, // EmbedLinks
		1 << 15, // AttachFiles
	}

	for _, perm := range commonPermissions {
		has, err := permissions.Has(perm)
		if err != nil {
			t.Errorf("Failed to check permission %d: %v", perm, err)
			continue
		}
		if !has {
			t.Errorf("Expected to have permission %d, but didn't", perm)
		}
	}
}

func TestPermissions_ZeroPermissions(t *testing.T) {
	permissions := Permissions("0")

	// Should not have any permissions
	has, err := permissions.Has(1)
	if err != nil {
		t.Fatalf("Failed to check zero permissions: %v", err)
	}
	if has {
		t.Error("Zero permissions should not have any permission bits set")
	}
}
