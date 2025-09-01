package discord

import (
	"testing"
	"time"
)

func TestSnowflake_String(t *testing.T) {
	tests := []struct {
		name     string
		snowflake Snowflake
		expected string
	}{
		{
			name:     "Valid snowflake",
			snowflake: "123456789012345678",
			expected: "123456789012345678",
		},
		{
			name:     "Empty snowflake",
			snowflake: "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.snowflake.String(); got != tt.expected {
				t.Errorf("Snowflake.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSnowflake_Int64(t *testing.T) {
	tests := []struct {
		name        string
		snowflake   Snowflake
		expected    int64
		expectError bool
	}{
		{
			name:        "Valid snowflake",
			snowflake:   "123456789012345678",
			expected:    123456789012345678,
			expectError: false,
		},
		{
			name:        "Invalid snowflake - non-numeric",
			snowflake:   "12345678901234567a",
			expected:    0,
			expectError: true,
		},
		{
			name:        "Empty snowflake",
			snowflake:   "",
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.snowflake.Int64()
			if (err != nil) != tt.expectError {
				t.Errorf("Snowflake.Int64() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if got != tt.expected {
				t.Errorf("Snowflake.Int64() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSnowflake_Time(t *testing.T) {
	tests := []struct {
		name        string
		snowflake   Snowflake
		expectError bool
	}{
		{
			name:        "Valid snowflake",
			snowflake:   "123456789012345678",
			expectError: false,
		},
		{
			name:        "Invalid snowflake",
			snowflake:   "invalid",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.snowflake.Time()
			if (err != nil) != tt.expectError {
				t.Errorf("Snowflake.Time() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if !tt.expectError && got.IsZero() {
				t.Errorf("Snowflake.Time() returned zero time for valid snowflake")
			}
		})
	}
}

func TestSnowflake_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		snowflake Snowflake
		expected bool
	}{
		{
			name:     "Valid snowflake - 17 digits",
			snowflake: "12345678901234567",
			expected: true,
		},
		{
			name:     "Valid snowflake - 18 digits",
			snowflake: "123456789012345678",
			expected: true,
		},
		{
			name:     "Valid snowflake - 19 digits",
			snowflake: "1234567890123456789",
			expected: true,
		},
		{
			name:     "Valid snowflake - 20 digits",
			snowflake: "12345678901234567890",
			expected: true,
		},
		{
			name:     "Invalid snowflake - too short",
			snowflake: "1234567890123456",
			expected: false,
		},
		{
			name:     "Invalid snowflake - too long",
			snowflake: "123456789012345678901",
			expected: false,
		},
		{
			name:     "Invalid snowflake - non-numeric",
			snowflake: "12345678901234567a",
			expected: false,
		},
		{
			name:     "Empty snowflake",
			snowflake: "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.snowflake.IsValid(); got != tt.expected {
				t.Errorf("Snowflake.IsValid() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNewSnowflake(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Snowflake
	}{
		{
			name:     "Valid string",
			input:    "123456789012345678",
			expected: Snowflake("123456789012345678"),
		},
		{
			name:     "Empty string",
			input:    "",
			expected: Snowflake(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSnowflake(tt.input); got != tt.expected {
				t.Errorf("NewSnowflake() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNewSnowflakeFromInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected Snowflake
	}{
		{
			name:     "Positive number",
			input:    123456789012345678,
			expected: Snowflake("123456789012345678"),
		},
		{
			name:     "Zero",
			input:    0,
			expected: Snowflake("0"),
		},
		{
			name:     "Negative number",
			input:    -123456789012345678,
			expected: Snowflake("-123456789012345678"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSnowflakeFromInt64(tt.input); got != tt.expected {
				t.Errorf("NewSnowflakeFromInt64() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSnowflake_TimestampExtraction(t *testing.T) {
	// Test that we can extract a reasonable timestamp from a snowflake
	snowflake := Snowflake("123456789012345678")
	
	extractedTime, err := snowflake.Time()
	if err != nil {
		t.Fatalf("Failed to extract time from snowflake: %v", err)
	}
	
	// The time should be after Discord's epoch (2015-01-01)
	discordEpoch := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	if extractedTime.Before(discordEpoch) {
		t.Errorf("Extracted time %v is before Discord epoch %v", extractedTime, discordEpoch)
	}
	
	// The time should be reasonable (not in the far future)
	farFuture := time.Now().AddDate(10, 0, 0) // 10 years from now
	if extractedTime.After(farFuture) {
		t.Errorf("Extracted time %v is unreasonably far in the future", extractedTime)
	}
}
