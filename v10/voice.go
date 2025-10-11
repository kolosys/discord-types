package v10

// VoiceGatewayVersion represents the Voice Gateway version.
const VoiceGatewayVersion = "8"

// VoiceEncryptionMode represents voice encryption modes.
// Reference: https://discord.com/developers/docs/topics/voice-connections#transport-encryption-modes
type VoiceEncryptionMode string

// Voice encryption mode constants
const (
	// VoiceEncryptionModeAeadAes256GcmRtpSize represents AEAD AES256-GCM (RTP Size).
	VoiceEncryptionModeAeadAes256GcmRtpSize VoiceEncryptionMode = "aead_aes256_gcm_rtpsize"

	// VoiceEncryptionModeAeadXChaCha20Poly1305RtpSize represents AEAD XChaCha20 Poly1305 (RTP Size).
	VoiceEncryptionModeAeadXChaCha20Poly1305RtpSize VoiceEncryptionMode = "aead_xchacha20_poly1305_rtpsize"

	// VoiceEncryptionModeXSalsa20Poly1305LiteRtpSize represents XSalsa20 Poly1305 Lite (RTP Size).
	// Deprecated: This encryption mode has been discontinued.
	VoiceEncryptionModeXSalsa20Poly1305LiteRtpSize VoiceEncryptionMode = "xsalsa20_poly1305_lite_rtpsize"

	// VoiceEncryptionModeAeadAes256Gcm represents AEAD AES256-GCM.
	// Deprecated: This encryption mode has been discontinued.
	VoiceEncryptionModeAeadAes256Gcm VoiceEncryptionMode = "aead_aes256_gcm"

	// VoiceEncryptionModeXSalsa20Poly1305 represents XSalsa20 Poly1305.
	// Deprecated: This encryption mode has been discontinued.
	VoiceEncryptionModeXSalsa20Poly1305 VoiceEncryptionMode = "xsalsa20_poly1305"

	// VoiceEncryptionModeXSalsa20Poly1305Suffix represents XSalsa20 Poly1305 Suffix.
	// Deprecated: This encryption mode has been discontinued.
	VoiceEncryptionModeXSalsa20Poly1305Suffix VoiceEncryptionMode = "xsalsa20_poly1305_suffix"

	// VoiceEncryptionModeXSalsa20Poly1305Lite represents XSalsa20 Poly1305 Lite.
	// Deprecated: This encryption mode has been discontinued.
	VoiceEncryptionModeXSalsa20Poly1305Lite VoiceEncryptionMode = "xsalsa20_poly1305_lite"
)

// VoiceSpeakingFlags represents voice speaking flags.
// Reference: https://discord.com/developers/docs/topics/voice-connections#speaking
type VoiceSpeakingFlags int

// Voice speaking flag constants
const (
	// VoiceSpeakingFlagMicrophone represents normal transmission of voice audio.
	VoiceSpeakingFlagMicrophone VoiceSpeakingFlags = 1 << 0

	// VoiceSpeakingFlagSoundshare represents transmission of context audio for video, no speaking indicator.
	VoiceSpeakingFlagSoundshare VoiceSpeakingFlags = 1 << 1

	// VoiceSpeakingFlagPriority represents priority speaker, lowering audio of other speakers.
	VoiceSpeakingFlagPriority VoiceSpeakingFlags = 1 << 2
)

// VoicePayload represents the base structure for all Voice Gateway payloads.
type VoicePayload[T any] struct {
	// Op is the opcode for the payload.
	OpCode VoiceOpcode `json:"op"`

	// D is the event data.
	Data T `json:"d,omitempty"`

	// Seq is the sequence number, used for resuming sessions and heartbeats.
	SeqNumber *int `json:"seq,omitempty"`
}

// VoiceBinaryPayload represents binary payloads for DAVE protocol.
type VoiceBinaryPayload []byte

// VoiceHello represents a Hello payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections#heartbeating
type VoiceHello = VoicePayload[VoiceHelloData]

// VoiceHelloData represents Hello payload data.
// Reference: https://discord.com/developers/docs/topics/voice-connections#heartbeating
type VoiceHelloData struct {
	// V is the voice gateway version.
	// Reference: https://discord.com/developers/docs/topics/voice-connections#voice-gateway-versioning-gateway-versions
	Version int `json:"v"`

	// HeartbeatInterval is the interval (in milliseconds) the client should heartbeat with.
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// VoiceReady represents a Ready payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceReady = VoicePayload[VoiceReadyData]

// VoiceReadyData represents Ready payload data.
// Reference: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceReadyData struct {
	// SSRC is the SSRC identifier.
	SSRC int `json:"ssrc"`

	// IP is the UDP IP.
	IP string `json:"ip"`

	// Port is the UDP port.
	Port int `json:"port"`

	// Modes are the supported encryption modes.
	// Reference: https://discord.com/developers/docs/topics/voice-connections#transport-encryption-modes
	Modes []VoiceEncryptionMode `json:"modes"`
}

// VoiceHeartbeatAck represents a Heartbeat ACK payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections#heartbeating
type VoiceHeartbeatAck = VoicePayload[VoiceHeartbeatAckData]

// VoiceHeartbeatAckData represents Heartbeat ACK payload data.
// Reference: https://discord.com/developers/docs/topics/voice-connections#heartbeating
type VoiceHeartbeatAckData struct {
	// T is the integer nonce.
	T int `json:"t"`
}

// VoiceSessionDescription represents a Session Description payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections#transport-encryption-and-sending-voice
type VoiceSessionDescription = VoicePayload[VoiceSessionDescriptionData]

// VoiceSessionDescriptionData represents Session Description payload data.
// Reference: https://discord.com/developers/docs/topics/voice-connections#transport-encryption-and-sending-voice
type VoiceSessionDescriptionData struct {
	// Mode is the selected mode.
	Mode VoiceEncryptionMode `json:"mode"`

	// SecretKey is the secret key.
	SecretKey []int `json:"secret_key"`

	// DaveProtocolVersion is the selected DAVE protocol version.
	// Reference: https://daveprotocol.com/#select_protocol_ack-4
	DaveProtocolVersion int `json:"dave_protocol_version"`
}

// VoiceResumed represents a Resumed payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections#resuming-voice-connection
type VoiceResumed = VoicePayload[any]

// VoiceSpeaking represents a Speaking payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections#speaking
type VoiceSpeaking = VoicePayload[VoiceSpeakingData]

// VoiceSpeakingData represents Speaking payload data.
// Reference: https://discord.com/developers/docs/topics/voice-connections#speaking
type VoiceSpeakingData struct {
	// Speaking is the speaking mode flags.
	Speaking VoiceSpeakingFlags `json:"speaking"`

	// SSRC is the SSRC identifier.
	SSRC int `json:"ssrc"`

	// UserID is the user id.
	UserID Snowflake `json:"user_id"`
}

// VoiceClientsConnect represents a Clients Connect payload.
// Reference: https://daveprotocol.com/#clients_connect-11
type VoiceClientsConnect = VoicePayload[VoiceClientsConnectData]

// VoiceClientsConnectData represents Clients Connect payload data.
// Reference: https://daveprotocol.com/#clients_connect-11
type VoiceClientsConnectData struct {
	// UserIDs are the connected user ids.
	UserIDs []Snowflake `json:"user_ids"`
}

// VoiceClientDisconnect represents a Client Disconnect payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections
type VoiceClientDisconnect = VoicePayload[VoiceClientDisconnectData]

// VoiceClientDisconnectData represents Client Disconnect payload data.
// Reference: https://discord.com/developers/docs/topics/voice-connections
type VoiceClientDisconnectData struct {
	// UserID is the disconnected user id.
	UserID Snowflake `json:"user_id"`
}

// VoiceDavePrepareTransition represents a DAVE Prepare Transition payload.
// Reference: https://daveprotocol.com/#dave_protocol_prepare_transition-21
type VoiceDavePrepareTransition = VoicePayload[VoiceDavePrepareTransitionData]

// VoiceDavePrepareTransitionData represents DAVE Prepare Transition payload data.
// Reference: https://daveprotocol.com/#dave_protocol_prepare_transition-21
type VoiceDavePrepareTransitionData struct {
	// ProtocolVersion is the protocol version.
	ProtocolVersion int `json:"protocol_version"`

	// TransitionID is the transition id.
	TransitionID int `json:"transition_id"`
}

// VoiceDaveExecuteTransition represents a DAVE Execute Transition payload.
// Reference: https://daveprotocol.com/#dave_protocol_execute_transition-22
type VoiceDaveExecuteTransition = VoicePayload[VoiceDaveExecuteTransitionData]

// VoiceDaveExecuteTransitionData represents DAVE Execute Transition payload data.
// Reference: https://daveprotocol.com/#dave_protocol_execute_transition-22
type VoiceDaveExecuteTransitionData struct {
	// TransitionID is the transition id.
	TransitionID int `json:"transition_id"`
}

// DAVE MLS payloads using binary data
type VoiceDaveMlsExternalSender VoiceBinaryPayload           // Reference: https://daveprotocol.com/#dave_mls_external_sender_package-25
type VoiceDaveMlsProposals VoiceBinaryPayload                // Reference: https://daveprotocol.com/#dave_mls_proposals-27
type VoiceDaveMlsAnnounceCommitTransition VoiceBinaryPayload // Reference: https://daveprotocol.com/#dave_mls_announce_commit_transition-29
type VoiceDaveMlsWelcome VoiceBinaryPayload                  // Reference: https://daveprotocol.com/#dave_mls_welcome-30

// VoiceIdentify represents an Identify payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceIdentify = VoicePayload[VoiceIdentifyData]

// VoiceIdentifyData represents Identify payload data.
// Reference: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceIdentifyData struct {
	// ServerID is the id of the server to connect to.
	ServerID Snowflake `json:"server_id"`

	// UserID is the id of the user to connect as.
	UserID Snowflake `json:"user_id"`

	// SessionID is the voice state session id.
	SessionID string `json:"session_id"`

	// Token is the voice connection token.
	Token string `json:"token"`

	// MaxDaveProtocolVersion is the maximum DAVE protocol version supported.
	MaxDaveProtocolVersion *int `json:"max_dave_protocol_version,omitempty"`
}

// VoiceHeartbeat represents a Heartbeat payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceHeartbeat = VoicePayload[VoiceHeartbeatData]

// VoiceHeartbeatData represents Heartbeat payload data.
// Reference: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceHeartbeatData struct {
	// T is the integer nonce.
	T int `json:"t"`

	// SeqAck is the last sequence number received.
	SeqAck int `json:"seq_ack"`
}

// VoiceSelectProtocol represents a Select Protocol payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection
type VoiceSelectProtocol = VoicePayload[VoiceSelectProtocolData]

// VoiceSelectProtocolData represents Select Protocol payload data.
// Reference: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection
type VoiceSelectProtocolData struct {
	// Protocol is the voice protocol.
	Protocol string `json:"protocol"`

	// Data is the data associated with the protocol.
	Data VoiceUDPProtocolData `json:"data"`
}

// VoiceUDPProtocolData represents UDP protocol data.
// Reference: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection
type VoiceUDPProtocolData struct {
	// Address is the external address.
	Address string `json:"address"`

	// Port is the external UDP port.
	Port int `json:"port"`

	// Mode is the selected mode.
	//
	// Reference: https://discord.com/developers/docs/topics/voice-connections#transport-encryption-modes
	Mode VoiceEncryptionMode `json:"mode"`
}

// VoiceResume represents a Resume payload.
// Reference: https://discord.com/developers/docs/topics/voice-connections#resuming-voice-connection
type VoiceResume = VoicePayload[VoiceResumeData]

// VoiceResumeData represents Resume payload data.
// Reference: https://discord.com/developers/docs/topics/voice-connections#resuming-voice-connection
type VoiceResumeData struct {
	// ServerID is the id of the server to connect to.
	ServerID Snowflake `json:"server_id"`

	// SessionID is the voice state session id.
	SessionID string `json:"session_id"`

	// Token is the voice connection token.
	Token string `json:"token"`

	// SeqAck is the last received sequence number.
	SeqAck int `json:"seq_ack"`
}

// VoiceSpeakingSend represents a Speaking payload (sendable).
// Reference: https://discord.com/developers/docs/topics/voice-connections#speaking
type VoiceSpeakingSend = VoicePayload[VoiceSpeakingSendData]

// VoiceSpeakingSendData represents Speaking payload data (sendable).
// Reference: https://discord.com/developers/docs/topics/voice-connections#speaking
type VoiceSpeakingSendData struct {
	// Speaking is the speaking mode flags.
	Speaking VoiceSpeakingFlags `json:"speaking"`

	// Delay is the voice delay.
	Delay int `json:"delay"`

	// SSRC is the SSRC identifier.
	SSRC int `json:"ssrc"`
}

// VoiceDaveTransitionReady represents a DAVE Transition Ready payload.
// Reference: https://daveprotocol.com/#dave_protocol_ready_for_transition-23
type VoiceDaveTransitionReady = VoicePayload[VoiceDaveTransitionReadyData]

// VoiceDaveTransitionReadyData represents DAVE Transition Ready payload data.
// Reference: https://daveprotocol.com/#dave_protocol_ready_for_transition-23
type VoiceDaveTransitionReadyData struct {
	// TransitionID is the transition id.
	TransitionID int `json:"transition_id"`
}

// VoiceDavePrepareEpoch represents a DAVE Prepare Epoch payload.
// Reference: https://daveprotocol.com/#dave_protocol_prepare_epoch-24
type VoiceDavePrepareEpoch = VoicePayload[VoiceDavePrepareEpochData]

// VoiceDavePrepareEpochData represents DAVE Prepare Epoch payload data.
// Reference: https://daveprotocol.com/#dave_protocol_prepare_epoch-24
type VoiceDavePrepareEpochData struct {
	// ProtocolVersion is the protocol version.
	ProtocolVersion int `json:"protocol_version"`

	// Epoch is the epoch id.
	Epoch int `json:"epoch"`
}

// VoiceDaveMlsInvalidCommitWelcome represents a DAVE MLS Invalid Commit Welcome payload.
// Reference: https://daveprotocol.com/#dave_mls_invalid_commit_welcome-31
type VoiceDaveMlsInvalidCommitWelcome = VoicePayload[VoiceDaveMlsInvalidCommitWelcomeData]

// VoiceDaveMlsInvalidCommitWelcomeData represents DAVE MLS Invalid Commit Welcome payload data.
// Reference: https://daveprotocol.com/#dave_mls_invalid_commit_welcome-31
type VoiceDaveMlsInvalidCommitWelcomeData struct {
	// TransitionID is the transition id.
	TransitionID int `json:"transition_id"`
}

// DAVE MLS sendable payloads using binary data
type VoiceDaveMlsKeyPackage VoiceBinaryPayload    // Reference: https://daveprotocol.com/#dave_mls_key_package-26
type VoiceDaveMlsCommitWelcome VoiceBinaryPayload // Reference: https://daveprotocol.com/#dave_mls_commit_welcome-28

// VoiceState represents a Discord voice state.
// Reference: https://discord.com/developers/docs/resources/voice#voice-state-object
type VoiceState struct {
	// GuildID is the guild id this voice state is for.
	GuildID *Snowflake `json:"guild_id,omitempty"`

	// ChannelID is the channel id this user is connected to.
	ChannelID *Snowflake `json:"channel_id"`

	// UserID is the user id this voice state is for.
	UserID Snowflake `json:"user_id"`

	// Member is the guild member this voice state is for.
	Member *GuildMember `json:"member,omitempty"`

	// SessionID is the session id for this voice state.
	SessionID string `json:"session_id"`

	// Deaf indicates whether this user is deafened by the server.
	Deaf bool `json:"deaf"`

	// Mute indicates whether this user is muted by the server.
	Mute bool `json:"mute"`

	// SelfDeaf indicates whether this user is locally deafened.
	SelfDeaf bool `json:"self_deaf"`

	// SelfMute indicates whether this user is locally muted.
	SelfMute bool `json:"self_mute"`

	// SelfStream indicates whether this user is streaming using "Go Live".
	SelfStream *bool `json:"self_stream,omitempty"`

	// SelfVideo indicates whether this user's camera is enabled.
	SelfVideo bool `json:"self_video"`

	// Suppress indicates whether this user's permission to speak is denied.
	Suppress bool `json:"suppress"`

	// RequestToSpeakTimestamp is the time at which the user requested to speak.
	RequestToSpeakTimestamp *string `json:"request_to_speak_timestamp"`
}

// VoiceRegion represents a Discord voice region.
// Reference: https://discord.com/developers/docs/resources/voice#voice-region-object
type VoiceRegion struct {
	// ID is the unique ID for the region.
	ID string `json:"id"`

	// Name is the name of the region.
	Name string `json:"name"`

	// Optimal indicates whether this is the closest server to the current user's client.
	Optimal bool `json:"optimal"`

	// Deprecated indicates whether this is a deprecated voice region (avoid switching to these).
	Deprecated bool `json:"deprecated"`

	// Custom indicates whether this is a custom voice region (used for events/etc).
	Custom bool `json:"custom"`
}
