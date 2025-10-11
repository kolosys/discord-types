package v10

import (
	"strconv"
	"time"
)

// Snowflake represents a Discord Snowflake ID.
//
// Snowflakes are unique IDs used throughout Discord's API.
// They are 64-bit integers represented as strings in JSON.
//
// Reference: https://discord.com/developers/docs/reference#snowflakes
type Snowflake string

// String returns the string representation of the Snowflake.
func (s Snowflake) String() string {
	return string(s)
}

// Int64 converts the Snowflake to an int64.
func (s Snowflake) Int64() (int64, error) {
	return strconv.ParseInt(string(s), 10, 64)
}

// NewSnowflake creates a new Snowflake from a string.
func NewSnowflake(s string) Snowflake {
	return Snowflake(s)
}

// NewSnowflakeFromInt64 creates a new Snowflake from an int64.
func NewSnowflakeFromInt64(i int64) Snowflake {
	return Snowflake(strconv.FormatInt(i, 10))
}

// Time extracts the timestamp from the Snowflake.
func (s Snowflake) Time() (time.Time, error) {
	id, err := s.Int64()
	if err != nil {
		return time.Time{}, err
	}

	// Extract timestamp from Snowflake
	timestamp := (id >> 22) + DiscordEpoch

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
