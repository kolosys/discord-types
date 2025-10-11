package v10

// Lobby represents a Discord lobby
// Reference: https://discord.com/developers/docs/resources/lobby#lobby-object
type Lobby struct {
	// ID is the ID of this channel (lobby)
	ID Snowflake `json:"id"`

	// ApplicationID is the application that created the lobby
	ApplicationID Snowflake `json:"application_id"`

	// Metadata is a dictionary of string key/value pairs. The max total length is 1000.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Members are the members of the lobby
	Members []LobbyMember `json:"members"`

	// LinkedChannel is the guild channel linked to the lobby
	LinkedChannel *Channel `json:"linked_channel,omitempty"`
}

// LobbyMember represents a member of a lobby
//
// Reference: https://discord.com/developers/docs/resources/lobby#lobby-member-object
type LobbyMember struct {
	// User is the user object for the member
	User User `json:"user"`

	// Metadata is a dictionary of metadata key/value pairs
	Metadata map[string]string `json:"metadata"`

	// Flags are the lobby member flags
	Flags LobbyMemberFlags `json:"flags"`
}

// LobbyMemberFlags represents lobby member flags
type LobbyMemberFlags int

const (
	// LobbyMemberFlagCanLinkLobby indicates the member can link the lobby
	LobbyMemberFlagCanLinkLobby LobbyMemberFlags = 1 << 0
)

// Has checks if the flags contain the specified flag
func (f LobbyMemberFlags) Has(flag LobbyMemberFlags) bool {
	return f&flag != 0
}

// Add adds a flag to the flags
func (f LobbyMemberFlags) Add(flag LobbyMemberFlags) LobbyMemberFlags {
	return f | flag
}

// Remove removes a flag from the flags
func (f LobbyMemberFlags) Remove(flag LobbyMemberFlags) LobbyMemberFlags {
	return f &^ flag
}
