// Package voice provides Discord Voice Gateway/WebSocket types for API version 8.
//
// This package contains types for Discord Voice Gateway events, payloads,
// and other voice WebSocket-related functionality.
package voice

import "github.com/kolosys/discord-types/discord"

// VoiceGatewayVersion represents the Voice Gateway version.
const VoiceGatewayVersion = "8"

// =============================================================================
// VOICE OPCODES
// =============================================================================

// VoiceOpcodes represents Voice Gateway opcodes.
//
// See: https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-opcodes
type VoiceOpcodes int

// Voice Gateway opcode constants
const (
	// VoiceOpcodeIdentify begins a voice websocket connection.
	VoiceOpcodeIdentify VoiceOpcodes = iota

	// VoiceOpcodeSelectProtocol selects the voice protocol.
	VoiceOpcodeSelectProtocol

	// VoiceOpcodeReady completes the websocket handshake.
	VoiceOpcodeReady

	// VoiceOpcodeHeartbeat keeps the websocket connection alive.
	VoiceOpcodeHeartbeat

	// VoiceOpcodeSessionDescription describes the session.
	VoiceOpcodeSessionDescription

	// VoiceOpcodeSpeaking indicates which users are speaking.
	VoiceOpcodeSpeaking

	// VoiceOpcodeHeartbeatAck is sent to acknowledge a received client heartbeat.
	VoiceOpcodeHeartbeatAck

	// VoiceOpcodeResume resumes a connection.
	VoiceOpcodeResume

	// VoiceOpcodeHello is the time to wait between sending heartbeats in milliseconds.
	VoiceOpcodeHello

	// VoiceOpcodeResumed acknowledges a successful session resume.
	VoiceOpcodeResumed

	_ // 10 is unused

	// VoiceOpcodeClientsConnect indicates one or more clients have connected to the voice channel.
	VoiceOpcodeClientsConnect

	_ // 12 is unused

	// VoiceOpcodeClientDisconnect indicates a client has disconnected from the voice channel.
	VoiceOpcodeClientDisconnect

	_ // 14-20 are unused

	// VoiceOpcodeDavePrepareTransition indicates a downgrade from the DAVE protocol is upcoming.
	VoiceOpcodeDavePrepareTransition = 21

	// VoiceOpcodeDaveExecuteTransition executes a previously announced protocol transition.
	VoiceOpcodeDaveExecuteTransition

	// VoiceOpcodeDaveTransitionReady acknowledges readiness for previously announced transition.
	VoiceOpcodeDaveTransitionReady

	// VoiceOpcodeDavePrepareEpoch indicates a DAVE protocol version or group change is upcoming.
	VoiceOpcodeDavePrepareEpoch

	// VoiceOpcodeDaveMlsExternalSender provides credential and public key for MLS external sender.
	VoiceOpcodeDaveMlsExternalSender

	// VoiceOpcodeDaveMlsKeyPackage provides MLS Key Package for pending group member.
	VoiceOpcodeDaveMlsKeyPackage

	// VoiceOpcodeDaveMlsProposals provides MLS Proposals to be appended or revoked.
	VoiceOpcodeDaveMlsProposals

	// VoiceOpcodeDaveMlsCommitWelcome provides MLS Commit with optional MLS Welcome messages.
	VoiceOpcodeDaveMlsCommitWelcome

	// VoiceOpcodeDaveMlsAnnounceCommitTransition provides MLS Commit to be processed for upcoming transition.
	VoiceOpcodeDaveMlsAnnounceCommitTransition

	// VoiceOpcodeDaveMlsWelcome provides MLS Welcome to group for upcoming transition.
	VoiceOpcodeDaveMlsWelcome

	// VoiceOpcodeDaveMlsInvalidCommitWelcome flags invalid commit or welcome, request re-add.
	VoiceOpcodeDaveMlsInvalidCommitWelcome
)

// =============================================================================
// VOICE CLOSE CODES
// =============================================================================

// VoiceCloseCodes represents Voice Gateway close event codes.
//
// See: https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-close-event-codes
type VoiceCloseCodes int

// Voice Gateway close code constants
const (
	// VoiceCloseCodeUnknownOpcode indicates you sent an invalid opcode.
	VoiceCloseCodeUnknownOpcode VoiceCloseCodes = 4001 + iota

	// VoiceCloseCodeFailedToDecode indicates you sent an invalid payload in your identifying to the Gateway.
	VoiceCloseCodeFailedToDecode

	// VoiceCloseCodeNotAuthenticated indicates you sent a payload before identifying with the Gateway.
	VoiceCloseCodeNotAuthenticated

	// VoiceCloseCodeAuthenticationFailed indicates the token you sent in your identify payload is incorrect.
	VoiceCloseCodeAuthenticationFailed

	// VoiceCloseCodeAlreadyAuthenticated indicates you sent more than one identify payload.
	VoiceCloseCodeAlreadyAuthenticated

	// VoiceCloseCodeSessionNoLongerValid indicates your session is no longer valid.
	VoiceCloseCodeSessionNoLongerValid

	_ // 4007 is unused

	_ // 4008 is unused

	// VoiceCloseCodeSessionTimeout indicates your session has timed out.
	VoiceCloseCodeSessionTimeout

	_ // 4010 is unused

	// VoiceCloseCodeServerNotFound indicates we can't find the server you're trying to connect to.
	VoiceCloseCodeServerNotFound

	// VoiceCloseCodeUnknownProtocol indicates we didn't recognize the protocol you sent.
	VoiceCloseCodeUnknownProtocol

	_ // 4013 is unused

	// VoiceCloseCodeDisconnected indicates either the channel was deleted, you were kicked, or the main gateway session was dropped. Should not reconnect.
	VoiceCloseCodeDisconnected

	// VoiceCloseCodeVoiceServerCrashed indicates the server crashed. Our bad! Try resuming.
	VoiceCloseCodeVoiceServerCrashed

	// VoiceCloseCodeUnknownEncryptionMode indicates we didn't recognize your encryption.
	VoiceCloseCodeUnknownEncryptionMode

	_ // 4017-4019 are unused

	// VoiceCloseCodeBadRequest indicates you sent a malformed request.
	VoiceCloseCodeBadRequest = 4020

	// VoiceCloseCodeRateLimited indicates disconnect due to rate limit exceeded. Should not reconnect.
	VoiceCloseCodeRateLimited

	// VoiceCloseCodeCallTerminated indicates disconnect all clients due to call terminated (channel deleted, voice server changed, etc.). Should not reconnect.
	VoiceCloseCodeCallTerminated
)

// =============================================================================
// VOICE ENCRYPTION MODES
// =============================================================================

// VoiceEncryptionMode represents voice encryption modes.
//
// See: https://discord.com/developers/docs/topics/voice-connections#transport-encryption-modes
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

// =============================================================================
// VOICE SPEAKING FLAGS
// =============================================================================

// VoiceSpeakingFlags represents voice speaking flags.
//
// See: https://discord.com/developers/docs/topics/voice-connections#speaking
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

// =============================================================================
// PAYLOAD INTERFACES
// =============================================================================

// VoiceSendPayload represents payloads that can be sent to the Voice Gateway.
type VoiceSendPayload interface {
	isVoiceSendPayload()
}

// VoiceReceivePayload represents payloads that can be received from the Voice Gateway.
type VoiceReceivePayload interface {
	isVoiceReceivePayload()
}

// VoiceBasePayload represents the base structure for all Voice Gateway payloads.
type VoiceBasePayload struct {
	// Op is the opcode for the payload.
	Op VoiceOpcodes `json:"op"`

	// D is the event data.
	D interface{} `json:"d,omitempty"`

	// Seq is the sequence number, used for resuming sessions and heartbeats.
	Seq *int `json:"seq,omitempty"`
}

// VoiceDataPayload represents a payload with typed data.
type VoiceDataPayload[T any] struct {
	Op VoiceOpcodes `json:"op"`
	D  T            `json:"d"`
}

// VoiceBinaryPayload represents binary payloads for DAVE protocol.
type VoiceBinaryPayload []byte

// =============================================================================
// SERVER PAYLOADS (RECEIVE)
// =============================================================================

// VoiceHello represents a Hello payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections#heartbeating
type VoiceHello struct {
	Op VoiceOpcodes   `json:"op"`
	D  VoiceHelloData `json:"d"`
}

func (p VoiceHello) isVoiceReceivePayload() {}

// VoiceHelloData represents Hello payload data.
//
// See: https://discord.com/developers/docs/topics/voice-connections#heartbeating
type VoiceHelloData struct {
	// V is the voice gateway version.
	//
	// See: https://discord.com/developers/docs/topics/voice-connections#voice-gateway-versioning-gateway-versions
	V int `json:"v"`

	// HeartbeatInterval is the interval (in milliseconds) the client should heartbeat with.
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// VoiceReady represents a Ready payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceReady struct {
	Op VoiceOpcodes   `json:"op"`
	D  VoiceReadyData `json:"d"`
}

func (p VoiceReady) isVoiceReceivePayload() {}

// VoiceReadyData represents Ready payload data.
//
// See: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceReadyData struct {
	// SSRC is the SSRC identifier.
	SSRC int `json:"ssrc"`

	// IP is the UDP IP.
	IP string `json:"ip"`

	// Port is the UDP port.
	Port int `json:"port"`

	// Modes are the supported encryption modes.
	//
	// See: https://discord.com/developers/docs/topics/voice-connections#transport-encryption-modes
	Modes []VoiceEncryptionMode `json:"modes"`
}

// VoiceHeartbeatAck represents a Heartbeat ACK payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections#heartbeating
type VoiceHeartbeatAck struct {
	Op VoiceOpcodes          `json:"op"`
	D  VoiceHeartbeatAckData `json:"d"`
}

func (p VoiceHeartbeatAck) isVoiceReceivePayload() {}

// VoiceHeartbeatAckData represents Heartbeat ACK payload data.
//
// See: https://discord.com/developers/docs/topics/voice-connections#heartbeating
type VoiceHeartbeatAckData struct {
	// T is the integer nonce.
	T int `json:"t"`
}

// VoiceSessionDescription represents a Session Description payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections#transport-encryption-and-sending-voice
type VoiceSessionDescription struct {
	Op VoiceOpcodes                `json:"op"`
	D  VoiceSessionDescriptionData `json:"d"`
}

func (p VoiceSessionDescription) isVoiceReceivePayload() {}

// VoiceSessionDescriptionData represents Session Description payload data.
//
// See: https://discord.com/developers/docs/topics/voice-connections#transport-encryption-and-sending-voice
type VoiceSessionDescriptionData struct {
	// Mode is the selected mode.
	Mode VoiceEncryptionMode `json:"mode"`

	// SecretKey is the secret key.
	SecretKey []int `json:"secret_key"`

	// DaveProtocolVersion is the selected DAVE protocol version.
	//
	// See: https://daveprotocol.com/#select_protocol_ack-4
	DaveProtocolVersion int `json:"dave_protocol_version"`
}

// VoiceResumed represents a Resumed payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections#resuming-voice-connection
type VoiceResumed struct {
	Op VoiceOpcodes `json:"op"`
	D  interface{}  `json:"d"` // null
}

func (p VoiceResumed) isVoiceReceivePayload() {}

// VoiceSpeaking represents a Speaking payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections#speaking
type VoiceSpeaking struct {
	Op VoiceOpcodes      `json:"op"`
	D  VoiceSpeakingData `json:"d"`
}

func (p VoiceSpeaking) isVoiceReceivePayload() {}

// VoiceSpeakingData represents Speaking payload data.
//
// See: https://discord.com/developers/docs/topics/voice-connections#speaking
type VoiceSpeakingData struct {
	// Speaking is the speaking mode flags.
	Speaking VoiceSpeakingFlags `json:"speaking"`

	// SSRC is the SSRC identifier.
	SSRC int `json:"ssrc"`

	// UserID is the user id.
	UserID discord.Snowflake `json:"user_id"`
}

// VoiceClientsConnect represents a Clients Connect payload.
//
// See: https://daveprotocol.com/#clients_connect-11
type VoiceClientsConnect struct {
	Op VoiceOpcodes            `json:"op"`
	D  VoiceClientsConnectData `json:"d"`
}

func (p VoiceClientsConnect) isVoiceReceivePayload() {}

// VoiceClientsConnectData represents Clients Connect payload data.
//
// See: https://daveprotocol.com/#clients_connect-11
type VoiceClientsConnectData struct {
	// UserIDs are the connected user ids.
	UserIDs []discord.Snowflake `json:"user_ids"`
}

// VoiceClientDisconnect represents a Client Disconnect payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections
type VoiceClientDisconnect struct {
	Op VoiceOpcodes              `json:"op"`
	D  VoiceClientDisconnectData `json:"d"`
}

func (p VoiceClientDisconnect) isVoiceReceivePayload() {}

// VoiceClientDisconnectData represents Client Disconnect payload data.
//
// See: https://discord.com/developers/docs/topics/voice-connections
type VoiceClientDisconnectData struct {
	// UserID is the disconnected user id.
	UserID discord.Snowflake `json:"user_id"`
}

// VoiceDavePrepareTransition represents a DAVE Prepare Transition payload.
//
// See: https://daveprotocol.com/#dave_protocol_prepare_transition-21
type VoiceDavePrepareTransition struct {
	Op VoiceOpcodes                   `json:"op"`
	D  VoiceDavePrepareTransitionData `json:"d"`
}

func (p VoiceDavePrepareTransition) isVoiceReceivePayload() {}

// VoiceDavePrepareTransitionData represents DAVE Prepare Transition payload data.
//
// See: https://daveprotocol.com/#dave_protocol_prepare_transition-21
type VoiceDavePrepareTransitionData struct {
	// ProtocolVersion is the protocol version.
	ProtocolVersion int `json:"protocol_version"`

	// TransitionID is the transition id.
	TransitionID int `json:"transition_id"`
}

// VoiceDaveExecuteTransition represents a DAVE Execute Transition payload.
//
// See: https://daveprotocol.com/#dave_protocol_execute_transition-22
type VoiceDaveExecuteTransition struct {
	Op VoiceOpcodes                   `json:"op"`
	D  VoiceDaveExecuteTransitionData `json:"d"`
}

func (p VoiceDaveExecuteTransition) isVoiceReceivePayload() {}

// VoiceDaveExecuteTransitionData represents DAVE Execute Transition payload data.
//
// See: https://daveprotocol.com/#dave_protocol_execute_transition-22
type VoiceDaveExecuteTransitionData struct {
	// TransitionID is the transition id.
	TransitionID int `json:"transition_id"`
}

// DAVE MLS payloads using binary data
type VoiceDaveMlsExternalSender VoiceBinaryPayload           // See: https://daveprotocol.com/#dave_mls_external_sender_package-25
type VoiceDaveMlsProposals VoiceBinaryPayload                // See: https://daveprotocol.com/#dave_mls_proposals-27
type VoiceDaveMlsAnnounceCommitTransition VoiceBinaryPayload // See: https://daveprotocol.com/#dave_mls_announce_commit_transition-29
type VoiceDaveMlsWelcome VoiceBinaryPayload                  // See: https://daveprotocol.com/#dave_mls_welcome-30

// =============================================================================
// SENDABLE PAYLOADS (SEND)
// =============================================================================

// VoiceIdentify represents an Identify payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceIdentify struct {
	Op VoiceOpcodes      `json:"op"`
	D  VoiceIdentifyData `json:"d"`
}

func (p VoiceIdentify) isVoiceSendPayload() {}

// VoiceIdentifyData represents Identify payload data.
//
// See: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceIdentifyData struct {
	// ServerID is the id of the server to connect to.
	ServerID discord.Snowflake `json:"server_id"`

	// UserID is the id of the user to connect as.
	UserID discord.Snowflake `json:"user_id"`

	// SessionID is the voice state session id.
	SessionID string `json:"session_id"`

	// Token is the voice connection token.
	Token string `json:"token"`

	// MaxDaveProtocolVersion is the maximum DAVE protocol version supported.
	MaxDaveProtocolVersion *int `json:"max_dave_protocol_version,omitempty"`
}

// VoiceHeartbeat represents a Heartbeat payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceHeartbeat struct {
	Op VoiceOpcodes       `json:"op"`
	D  VoiceHeartbeatData `json:"d"`
}

func (p VoiceHeartbeat) isVoiceSendPayload() {}

// VoiceHeartbeatData represents Heartbeat payload data.
//
// See: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection
type VoiceHeartbeatData struct {
	// T is the integer nonce.
	T int `json:"t"`

	// SeqAck is the last sequence number received.
	SeqAck int `json:"seq_ack"`
}

// VoiceSelectProtocol represents a Select Protocol payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection
type VoiceSelectProtocol struct {
	Op VoiceOpcodes            `json:"op"`
	D  VoiceSelectProtocolData `json:"d"`
}

func (p VoiceSelectProtocol) isVoiceSendPayload() {}

// VoiceSelectProtocolData represents Select Protocol payload data.
//
// See: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection
type VoiceSelectProtocolData struct {
	// Protocol is the voice protocol.
	Protocol string `json:"protocol"`

	// Data is the data associated with the protocol.
	Data VoiceUDPProtocolData `json:"data"`
}

// VoiceUDPProtocolData represents UDP protocol data.
//
// See: https://discord.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection
type VoiceUDPProtocolData struct {
	// Address is the external address.
	Address string `json:"address"`

	// Port is the external UDP port.
	Port int `json:"port"`

	// Mode is the selected mode.
	//
	// See: https://discord.com/developers/docs/topics/voice-connections#transport-encryption-modes
	Mode VoiceEncryptionMode `json:"mode"`
}

// VoiceResume represents a Resume payload.
//
// See: https://discord.com/developers/docs/topics/voice-connections#resuming-voice-connection
type VoiceResume struct {
	Op VoiceOpcodes    `json:"op"`
	D  VoiceResumeData `json:"d"`
}

func (p VoiceResume) isVoiceSendPayload() {}

// VoiceResumeData represents Resume payload data.
//
// See: https://discord.com/developers/docs/topics/voice-connections#resuming-voice-connection
type VoiceResumeData struct {
	// ServerID is the id of the server to connect to.
	ServerID discord.Snowflake `json:"server_id"`

	// SessionID is the voice state session id.
	SessionID string `json:"session_id"`

	// Token is the voice connection token.
	Token string `json:"token"`

	// SeqAck is the last received sequence number.
	SeqAck int `json:"seq_ack"`
}

// VoiceSpeakingSend represents a Speaking payload (sendable).
//
// See: https://discord.com/developers/docs/topics/voice-connections#speaking
type VoiceSpeakingSend struct {
	Op VoiceOpcodes          `json:"op"`
	D  VoiceSpeakingSendData `json:"d"`
}

func (p VoiceSpeakingSend) isVoiceSendPayload() {}

// VoiceSpeakingSendData represents Speaking payload data (sendable).
//
// See: https://discord.com/developers/docs/topics/voice-connections#speaking
type VoiceSpeakingSendData struct {
	// Speaking is the speaking mode flags.
	Speaking VoiceSpeakingFlags `json:"speaking"`

	// Delay is the voice delay.
	Delay int `json:"delay"`

	// SSRC is the SSRC identifier.
	SSRC int `json:"ssrc"`
}

// VoiceDaveTransitionReady represents a DAVE Transition Ready payload.
//
// See: https://daveprotocol.com/#dave_protocol_ready_for_transition-23
type VoiceDaveTransitionReady struct {
	Op VoiceOpcodes                 `json:"op"`
	D  VoiceDaveTransitionReadyData `json:"d"`
}

func (p VoiceDaveTransitionReady) isVoiceSendPayload() {}

// VoiceDaveTransitionReadyData represents DAVE Transition Ready payload data.
//
// See: https://daveprotocol.com/#dave_protocol_ready_for_transition-23
type VoiceDaveTransitionReadyData struct {
	// TransitionID is the transition id.
	TransitionID int `json:"transition_id"`
}

// VoiceDavePrepareEpoch represents a DAVE Prepare Epoch payload.
//
// See: https://daveprotocol.com/#dave_protocol_prepare_epoch-24
type VoiceDavePrepareEpoch struct {
	Op VoiceOpcodes              `json:"op"`
	D  VoiceDavePrepareEpochData `json:"d"`
}

func (p VoiceDavePrepareEpoch) isVoiceReceivePayload() {}

// VoiceDavePrepareEpochData represents DAVE Prepare Epoch payload data.
//
// See: https://daveprotocol.com/#dave_protocol_prepare_epoch-24
type VoiceDavePrepareEpochData struct {
	// ProtocolVersion is the protocol version.
	ProtocolVersion int `json:"protocol_version"`

	// Epoch is the epoch id.
	Epoch int `json:"epoch"`
}

// VoiceDaveMlsInvalidCommitWelcome represents a DAVE MLS Invalid Commit Welcome payload.
//
// See: https://daveprotocol.com/#dave_mls_invalid_commit_welcome-31
type VoiceDaveMlsInvalidCommitWelcome struct {
	Op VoiceOpcodes                         `json:"op"`
	D  VoiceDaveMlsInvalidCommitWelcomeData `json:"d"`
}

func (p VoiceDaveMlsInvalidCommitWelcome) isVoiceSendPayload() {}

// VoiceDaveMlsInvalidCommitWelcomeData represents DAVE MLS Invalid Commit Welcome payload data.
//
// See: https://daveprotocol.com/#dave_mls_invalid_commit_welcome-31
type VoiceDaveMlsInvalidCommitWelcomeData struct {
	// TransitionID is the transition id.
	TransitionID int `json:"transition_id"`
}

// DAVE MLS sendable payloads using binary data
type VoiceDaveMlsKeyPackage VoiceBinaryPayload    // See: https://daveprotocol.com/#dave_mls_key_package-26
type VoiceDaveMlsCommitWelcome VoiceBinaryPayload // See: https://daveprotocol.com/#dave_mls_commit_welcome-28
