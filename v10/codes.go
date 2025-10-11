package v10

// =============================================================================
// GATEWAY OPCODES
// =============================================================================

// Gateway Opcodes
// Reference: https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-opcodes
type Opcode int

const (
	// OpDispatch - An event was dispatched.
	OpDispatch Opcode = iota
	// OpHeartbeat - Fired periodically by the client to keep the connection alive.
	OpHeartbeat
	// OpIdentify - Starts a new session during the initial handshake.
	OpIdentify
	// OpPresenceUpdate - Update the client's presence.
	OpPresenceUpdate
	// OpVoiceStateUpdate - Used to join/leave or move between voice channels.
	OpVoiceStateUpdate
	_ // 5 is unused
	// OpResume - Resume a previous session that was disconnected.
	OpResume
	// OpReconnect - You should attempt to reconnect and resume immediately.
	OpReconnect
	// OpRequestGuildMembers - Request information about offline guild members in a large guild.
	OpRequestGuildMembers
	// OpInvalidSession - The session has been invalidated. You should reconnect and identify/resume accordingly.
	OpInvalidSession
	// OpHello - Sent immediately after connecting, contains the heartbeat_interval to use.
	OpHello
	// OpHeartbeatACK - Sent in response to receiving a heartbeat to acknowledge that it has been received.
	OpHeartbeatACK
)

// Gateway Close Event Codes
// Reference: https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-close-event-codes
type CloseCode int

const (
	// CloseUnknownError - We're not sure what went wrong. Try reconnecting?
	CloseUnknownError CloseCode = 4000
	// CloseUnknownOpcode - You sent an invalid Gateway opcode or an invalid payload for an opcode.
	CloseUnknownOpcode CloseCode = 4001
	// CloseDecodeError - You sent an invalid payload to Discord.
	CloseDecodeError = 4002
	// CloseNotAuthenticated - You sent us a payload prior to identifying.
	CloseNotAuthenticated = 4003
	// CloseAuthenticationFailed - The account token sent with your identify payload is incorrect.
	CloseAuthenticationFailed = 4004
	// CloseAlreadyAuthenticated - You sent more than one identify payload.
	CloseAlreadyAuthenticated = 4005
	_                         // 4006 is unused
	// CloseInvalidSeq - The sequence sent when resuming the session was invalid.
	CloseInvalidSeq = 4007
	// CloseRateLimited - Woah nelly! You're sending payloads to us too quickly. Slow it down!
	CloseRateLimited = 4008
	// CloseSessionTimedOut - Your session timed out. Reconnect and start a new one.
	CloseSessionTimedOut = 4009
	// CloseInvalidShard - You sent us an invalid shard when identifying.
	CloseInvalidShard = 4010
	// CloseShardingRequired - The session would have handled too many guilds - you are required to shard your connection in order to connect.
	CloseShardingRequired = 4011
	// CloseInvalidAPIVersion - You sent an invalid version for the gateway.
	CloseInvalidAPIVersion = 4012
	// CloseInvalidIntents - You sent an invalid intent for a Gateway Intent.
	CloseInvalidIntents = 4013
	// CloseDisallowedIntents - You sent a disallowed intent for a Gateway Intent.
	CloseDisallowedIntents = 4014
)

// =============================================================================
// VOICE OPCODES
// =============================================================================

// VoiceOpcode represents Voice Gateway opcodes.
// Reference: https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-opcodes
type VoiceOpcode int

// Voice Gateway opcode constants
const (
	// VoiceOpcodeIdentify begins a voice websocket connection.
	VoiceOpcodeIdentify VoiceOpcode = iota
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
	_ // 14 are unused
	_ // 15 is unused
	_ // 16 is unused
	_ // 17 is unused
	_ // 18 is unused
	_ // 19 is unused
	_ // 20 is unused
	// VoiceOpcodeDavePrepareTransition indicates a downgrade from the DAVE protocol is upcoming.
	VoiceOpcodeDavePrepareTransition
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

// VoiceCloseCode represents Voice Gateway close event codes.
// Reference: https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-close-event-codes
type VoiceCloseCode int

// Voice Gateway close code constants
const (
	// VoiceCloseCodeUnknownOpcode indicates you sent an invalid opcode.
	VoiceCloseCodeUnknownOpcode VoiceCloseCode = 4001 + iota

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
// HTTP RESPONSE CODES
// =============================================================================

// HTTP Response Codes
// Reference: https://discord.com/developers/docs/topics/opcodes-and-status-codes#http-http-response-codes
type HTTPCode int

const (
	// HTTP 200 - The request completed successfully.
	HTTPOk HTTPCode = 200
	// HTTP 201 - The entity was created successfully.
	HTTPCreated HTTPCode = 201
	// HTTP 204 - The request completed successfully but returned no content.
	HTTPNoContent HTTPCode = 204
	// HTTP 304 - The entity was not modified (no action was taken).
	HTTPNotModified HTTPCode = 304
	// HTTP 400 - The request was improperly formatted, or the server couldn't understand it.
	HTTPBadRequest HTTPCode = 400
	// HTTP 401 - The Authorization header was missing or invalid.
	HTTPUnauthorized HTTPCode = 401
	// HTTP 403 - The Authorization token you passed did not have permission to the resource.
	HTTPForbidden HTTPCode = 403
	// HTTP 404 - The resource at the location specified doesn't exist.
	HTTPNotFound HTTPCode = 404
	// HTTP 405 - The HTTP method used is not valid for the location specified.
	HTTPMethodNotAllowed HTTPCode = 405
	// HTTP 429 - You are being rate limited, see Rate Limits.
	HTTPTooManyRequests HTTPCode = 429
	// HTTP 502 - There was not a gateway available to process your request. Wait a bit and retry.
	HTTPBadGateway HTTPCode = 502
)

// =============================================================================
// JSON ERROR CODES
// =============================================================================

// JSON Error Codes
// Reference: https://discord.com/developers/docs/topics/opcodes-and-status-codes#json-json-error-codes
const (
	// General error (such as a malformed request body, amongst other things)
	JSONErrorGeneral = 0

	// Unknown account
	JSONErrorUnknownAccount = 10001
	// Unknown application
	JSONErrorUnknownApplication = 10002
	// Unknown channel
	JSONErrorUnknownChannel = 10003
	// Unknown guild
	JSONErrorUnknownGuild = 10004
	// Unknown integration
	JSONErrorUnknownIntegration = 10005
	// Unknown invite
	JSONErrorUnknownInvite = 10006
	// Unknown member
	JSONErrorUnknownMember = 10007
	// Unknown message
	JSONErrorUnknownMessage = 10008
	// Unknown permission overwrite
	JSONErrorUnknownPermissionOverwrite = 10009
	// Unknown provider
	JSONErrorUnknownProvider = 10010
	// Unknown role
	JSONErrorUnknownRole = 10011
	// Unknown token
	JSONErrorUnknownToken = 10012
	// Unknown user
	JSONErrorUnknownUser = 10013
	// Unknown emoji
	JSONErrorUnknownEmoji = 10014
	// Unknown webhook
	JSONErrorUnknownWebhook = 10015
	// Unknown webhook service
	JSONErrorUnknownWebhookService = 10016
	// Unknown session
	JSONErrorUnknownSession = 10020
	// Unknown ban
	JSONErrorUnknownBan = 10026
	// Unknown SKU
	JSONErrorUnknownSKU = 10027
	// Unknown Store Listing
	JSONErrorUnknownStoreListing = 10028
	// Unknown entitlement
	JSONErrorUnknownEntitlement = 10029
	// Unknown build
	JSONErrorUnknownBuild = 10030
	// Unknown lobby
	JSONErrorUnknownLobby = 10031
	// Unknown branch
	JSONErrorUnknownBranch = 10032
	// Unknown store directory layout
	JSONErrorUnknownStoreDirectoryLayout = 10033
	// Unknown redistributable
	JSONErrorUnknownRedistributable = 10036
	// Unknown gift code
	JSONErrorUnknownGiftCode = 10038
	// Unknown stream
	JSONErrorUnknownStream = 10049
	// Unknown premium server subscribe cooldown
	JSONErrorUnknownPremiumServerSubscribeCooldown = 10050
	// Unknown guild template
	JSONErrorUnknownGuildTemplate = 10057
	// Unknown discoverable server category
	JSONErrorUnknownDiscoverableServerCategory = 10059
	// Unknown sticker
	JSONErrorUnknownSticker = 10060
	// Unknown interaction
	JSONErrorUnknownInteraction = 10062
	// Unknown application command
	JSONErrorUnknownApplicationCommand = 10063
	// Unknown voice state
	JSONErrorUnknownVoiceState = 10065
	// Unknown application command permissions
	JSONErrorUnknownApplicationCommandPermissions = 10066
	// Unknown Stage Instance
	JSONErrorUnknownStageInstance = 10067
	// Unknown Guild Member Verification Form
	JSONErrorUnknownGuildMemberVerificationForm = 10068
	// Unknown Guild Welcome Screen
	JSONErrorUnknownGuildWelcomeScreen = 10069
	// Unknown Guild Scheduled Event
	JSONErrorUnknownGuildScheduledEvent = 10070
	// Unknown Guild Scheduled Event User
	JSONErrorUnknownGuildScheduledEventUser = 10071
	// Unknown Tag
	JSONErrorUnknownTag = 10087

	// Bots cannot use this endpoint
	JSONErrorBotsCannotUse = 20001
	// Only bots can use this endpoint
	JSONErrorOnlyBotsCanUse = 20002
	// Explicit content cannot be sent to the recipient(s)
	JSONErrorExplicitContentCannotBeSent = 20009
	// You are not authorized to perform this action on this application
	JSONErrorNotAuthorizedForApplication = 20012
	// This action cannot be performed due to slowmode rate limit
	JSONErrorSlowmodeRateLimit = 20016
	// The owner of this account can only send messages to accounts with premium
	JSONErrorAccountOwnerOnlyPremium = 20018
	// This message cannot be edited due to announcement rate limits
	JSONErrorAnnouncementRateLimit = 20022
	// Under minimum age
	JSONErrorUnderMinimumAge = 20024
	// The channel you are writing has hit the write rate limit
	JSONErrorChannelWriteRateLimit = 20028
	// The write action you are performing on the server has hit the write rate limit
	JSONErrorServerWriteRateLimit = 20029
	// Your Stage topic, server name, server description, or channel names contain words that are not allowed
	JSONErrorDisallowedWords = 20031
	// Guild premium subscription level too low
	JSONErrorGuildPremiumTooLow = 20035

	// Maximum number of guilds reached (100)
	JSONErrorMaxGuilds = 30001
	// Maximum number of friends reached (1000)
	JSONErrorMaxFriends = 30002
	// Maximum number of pins reached for the channel (50)
	JSONErrorMaxPins = 30003
	// Maximum number of recipients reached (10)
	JSONErrorMaxRecipients = 30004
	// Maximum number of guild roles reached (250)
	JSONErrorMaxGuildRoles = 30005
	// Maximum number of webhooks reached (10)
	JSONErrorMaxWebhooks = 30007
	// Maximum number of emojis reached
	JSONErrorMaxEmojis = 30008
	// Maximum number of reactions reached (20)
	JSONErrorMaxReactions = 30010
	// Maximum number of guild channels reached (500)
	JSONErrorMaxGuildChannels = 30013
	// Maximum number of attachments in a message reached (10)
	JSONErrorMaxAttachments = 30015
	// Maximum number of invites reached (1000)
	JSONErrorMaxInvites = 30016
	// Maximum number of animated emojis reached
	JSONErrorMaxAnimatedEmojis = 30018
	// Maximum number of server members reached
	JSONErrorMaxServerMembers = 30019
	// Maximum number of server categories has been reached (5)
	JSONErrorMaxServerCategories = 30030
	// Guild already has a template
	JSONErrorGuildAlreadyHasTemplate = 30031
	// Maximum number of application commands reached
	JSONErrorMaxApplicationCommands = 30032
	// Maximum number of thread participants has been reached (1000)
	JSONErrorMaxThreadParticipants = 30033
	// Maximum number of daily application command creates has been reached (200)
	JSONErrorMaxDailyApplicationCommandCreates = 30034
	// Maximum number of bans for non-guild members have been exceeded
	JSONErrorMaxBansForNonGuildMembers = 30035
	// Maximum number of bans fetches has been reached
	JSONErrorMaxBanFetches = 30037
	// Maximum number of uncompleted guild scheduled events reached (100)
	JSONErrorMaxUncompletedGuildScheduledEvents = 30038
	// Maximum number of stickers reached
	JSONErrorMaxStickers = 30039
	// Maximum number of prune requests has been reached. Try again later
	JSONErrorMaxPruneRequests = 30040
	// Maximum number of guild widget settings updates has been reached. Try again later
	JSONErrorMaxGuildWidgetSettingsUpdates = 30041
	// Maximum number of edits to messages older than 1 hour reached. Try again later
	JSONErrorMaxEditsToOldMessages = 30042
	// Maximum number of pinned threads in a forum channel has been reached
	JSONErrorMaxPinnedThreadsInForum = 30043
	// Maximum number of tags in a forum channel has been reached
	JSONErrorMaxTagsInForum = 30044
	// Bitrate is too high for channel of this type
	JSONErrorBitrateTooHigh = 30045
	// Maximum number of premium emojis reached (25)
	JSONErrorMaxPremiumEmojis = 30046
	// Maximum number of webhooks per guild reached (1000)
	JSONErrorMaxWebhooksPerGuild = 30047
	// Maximum number of channel permission overwrites for a channel has been reached (100)
	JSONErrorMaxChannelPermissionOverwrites = 30048
	// The channels for this guild are too large
	JSONErrorChannelsTooLarge = 30049

	// Unauthorized. Provide a valid token and try again
	JSONErrorUnauthorized = 40001
	// You need to verify your account in order to perform this action
	JSONErrorNeedVerification = 40002
	// You are opening direct messages too fast
	JSONErrorOpeningDMsTooFast = 40003
	// Send messages has been temporarily disabled
	JSONErrorSendMessagesDisabled = 40004
	// Request entity too large. Try sending something smaller in size
	JSONErrorRequestEntityTooLarge = 40005
	// This feature has been temporarily disabled server-side
	JSONErrorFeatureTemporarilyDisabled = 40006
	// The user is banned from this guild
	JSONErrorUserBannedFromGuild = 40007
	// Connection has been revoked
	JSONErrorConnectionRevoked = 40012
	// Target user is not connected to voice
	JSONErrorTargetUserNotConnectedToVoice = 40032
	// This message has already been crossposted
	JSONErrorMessageAlreadyCrossposted = 40033
	// An application command with that name already exists
	JSONErrorApplicationCommandNameExists = 40041
	// Application interaction failed to send
	JSONErrorApplicationInteractionFailedToSend = 40043
	// Cannot send a message in a forum channel
	JSONErrorCannotSendMessageInForumChannel = 40058
	// Interaction has already been acknowledged
	JSONErrorInteractionAlreadyAcknowledged = 40060
	// Tag names must be unique
	JSONErrorTagNamesMustBeUnique = 40061
	// Service resource is being rate limited
	JSONErrorServiceResourceRateLimited = 40062
	// There are no tags available that can be set by non-moderators
	JSONErrorNoTagsAvailableForNonModerators = 40066
	// A tag is required to create a forum post in this channel
	JSONErrorTagRequiredToCreateForumPost = 40067

	// Missing access
	JSONErrorMissingAccess = 50001
	// Invalid account type
	JSONErrorInvalidAccountType = 50002
	// Cannot execute action on a DM channel
	JSONErrorCannotExecuteOnDM = 50003
	// Guild widget disabled
	JSONErrorGuildWidgetDisabled = 50004
	// Cannot edit a message authored by another user
	JSONErrorCannotEditOtherUserMessage = 50005
	// Cannot send an empty message
	JSONErrorCannotSendEmptyMessage = 50006
	// Cannot send messages to this user
	JSONErrorCannotSendMessagesToUser = 50007
	// Cannot send messages in a voice channel
	JSONErrorCannotSendMessagesInVoiceChannel = 50008
	// Channel verification level is too high for you to gain access
	JSONErrorChannelVerificationTooHigh = 50009
	// OAuth2 application does not have a bot
	JSONErrorOAuth2ApplicationNoBot = 50010
	// OAuth2 application limit reached
	JSONErrorOAuth2ApplicationLimitReached = 50011
	// Invalid OAuth2 state
	JSONErrorInvalidOAuth2State = 50012
	// You lack permissions to perform that action
	JSONErrorLackPermissions = 50013
	// Invalid authentication token provided
	JSONErrorInvalidAuthToken = 50014
	// Note was too long
	JSONErrorNoteTooLong = 50015
	// Provided too few or too many messages to delete. Must provide at least 2 and fewer than 100 messages to delete
	JSONErrorInvalidBulkDeleteCount = 50016
	// Invalid MFA Level
	JSONErrorInvalidMFALevel = 50017
	// A message can only be pinned to the channel it was sent in
	JSONErrorMessageCanOnlyBePinnedInSameChannel = 50019
	// Invite code was either invalid or taken
	JSONErrorInvalidOrTakenInviteCode = 50020
	// Cannot execute action on a system message
	JSONErrorCannotExecuteOnSystemMessage = 50021
	// Cannot execute action on this channel type
	JSONErrorCannotExecuteOnChannelType = 50024
	// Invalid OAuth2 access token provided
	JSONErrorInvalidOAuth2AccessToken = 50025
	// Missing required OAuth2 scope
	JSONErrorMissingRequiredOAuth2Scope = 50026
	// Invalid webhook token provided
	JSONErrorInvalidWebhookToken = 50027
	// Invalid role
	JSONErrorInvalidRole = 50028
	// Invalid Recipient(s)
	JSONErrorInvalidRecipients = 50033
	// A message provided was too old to bulk delete
	JSONErrorMessageTooOldToBulkDelete = 50034
	// Invalid form body (returned for both application/json and multipart/form-data bodies), or invalid Content-Type provided
	JSONErrorInvalidFormBody = 50035
	// An invite was accepted to a guild the application's bot is not in
	JSONErrorInviteAcceptedToGuildBotNotIn = 50036
	// Invalid Activity Action
	JSONErrorInvalidActivityAction = 50039
	// Invalid API version provided
	JSONErrorInvalidAPIVersion = 50041
	// File uploaded exceeds the maximum size
	JSONErrorFileUploadExceedsMaxSize = 50045
	// Invalid file uploaded
	JSONErrorInvalidFileUploaded = 50046
	// Cannot self-redeem this gift
	JSONErrorCannotSelfRedeemGift = 50054
	// Invalid Guild
	JSONErrorInvalidGuild = 50055
	// Invalid message type
	JSONErrorInvalidMessageType = 50068
	// Payment source required to redeem gift
	JSONErrorPaymentSourceRequiredToRedeemGift = 50070
	// Cannot modify a system webhook
	JSONErrorCannotModifySystemWebhook = 50073
	// Cannot delete a channel required for Community guilds
	JSONErrorCannotDeleteRequiredCommunityChannel = 50074
	// Cannot edit stickers within a message
	JSONErrorCannotEditStickersWithinMessage = 50080
	// Invalid sticker sent
	JSONErrorInvalidStickerSent = 50081
	// Tried to perform an operation on an archived thread, such as editing a message or adding a user to the thread
	JSONErrorOperationOnArchivedThread = 50083
	// Invalid thread notification settings
	JSONErrorInvalidThreadNotificationSettings = 50084
	// before value is earlier than the thread creation date
	JSONErrorBeforeValueEarlierThanThreadCreation = 50085
	// Community server channels must be text channels
	JSONErrorCommunityServerChannelsMustBeText = 50086
	// The entity type of the event is different from the entity you are trying to start the event for
	JSONErrorEventEntityTypeDifferent = 50091
	// This server is not available in your location
	JSONErrorServerNotAvailableInLocation = 50095
	// This server needs monetization enabled in order to perform this action
	JSONErrorServerNeedsMonetizationEnabled = 50097
	// This server needs more boosts to perform this action
	JSONErrorServerNeedsMoreBoosts = 50101
	// The request body contains invalid JSON.
	JSONErrorInvalidJSON = 50109
	// Owner cannot be pending member
	JSONErrorOwnerCannotBePendingMember = 50131
	// Ownership cannot be transferred to a bot user
	JSONErrorOwnershipCannotBeTransferredToBot = 50132
	// Failed to resize asset below the maximum size: 262144
	JSONErrorFailedToResizeAsset = 50138
	// Cannot mix subscription and non subscription roles for an emoji
	JSONErrorCannotMixSubscriptionAndNonSubscriptionRoles = 50144
	// Cannot convert between premium emoji and normal emoji
	JSONErrorCannotConvertBetweenPremiumAndNormalEmoji = 50145
	// Uploaded file not found.
	JSONErrorUploadedFileNotFound = 50146
	// Voice messages do not support additional content.
	JSONErrorVoiceMessagesNoAdditionalContent = 50159
	// Voice messages must have a single audio attachment.
	JSONErrorVoiceMessagesMustHaveSingleAudioAttachment = 50160
	// Voice messages must have supporting metadata.
	JSONErrorVoiceMessagesMustHaveSupportingMetadata = 50161
	// Voice messages cannot be edited.
	JSONErrorVoiceMessagesCannotBeEdited = 50162
	// You cannot send voice messages in this channel.
	JSONErrorCannotSendVoiceMessagesInChannel = 50163
	// The user account must first be verified
	JSONErrorUserAccountMustBeVerified = 50164
	// You do not have permission to send this sticker.
	JSONErrorNoPermissionToSendSticker = 50165

	// Two factor is required for this operation
	JSONErrorTwoFactorRequired = 60003

	// No users with DiscordTag exist
	JSONErrorNoUsersWithDiscordTag = 80004

	// Reaction was blocked
	JSONErrorReactionBlocked = 90001

	// User cannot use burst reactions
	JSONErrorUserCannotUseBurstReactions = 90002

	// Application not yet available. Try again later
	JSONErrorApplicationNotYetAvailable = 110001

	// API resource is currently overloaded. Try again a little later
	JSONErrorAPIResourceOverloaded = 130000

	// The Stage is already open
	JSONErrorStageAlreadyOpen = 150006

	// Cannot reply without permission to read message history
	JSONErrorCannotReplyWithoutPermissionToReadMessageHistory = 160002

	// A thread has already been created for this message
	JSONErrorThreadAlreadyCreatedForMessage = 160004

	// Thread is locked
	JSONErrorThreadLocked = 160005

	// Maximum number of active threads reached
	JSONErrorMaxActiveThreads = 160006

	// Maximum number of active announcement threads reached
	JSONErrorMaxActiveAnnouncementThreads = 160007

	// Invalid JSON for uploaded Lottie file
	JSONErrorInvalidJSONForLottieFile = 170001

	// Uploaded Lotties cannot contain rasterized images such as PNG or JPEG
	JSONErrorLottiesCannotContainRasterizedImages = 170002

	// Sticker maximum framerate exceeded
	JSONErrorStickerMaxFramerateExceeded = 170003

	// Sticker frame count exceeds maximum of 1000 frames
	JSONErrorStickerMaxFrameCountExceeded = 170004

	// Lottie animation maximum dimensions exceeded
	JSONErrorLottieMaxDimensionsExceeded = 170005

	// Sticker frame rate is either too small or too large
	JSONErrorStickerFrameRateInvalid = 170006

	// Sticker animation duration exceeds maximum of 5 seconds
	JSONErrorStickerMaxDurationExceeded = 170007

	// Cannot update a finished event
	JSONErrorCannotUpdateFinishedEvent = 180000

	// Failed to create stage needed for stage event
	JSONErrorFailedToCreateStageForEvent = 180002

	// Message was blocked by automatic moderation
	JSONErrorMessageBlockedByAutoModeration = 200000

	// Title was blocked by automatic moderation
	JSONErrorTitleBlockedByAutoModeration = 200001

	// Webhooks posted to forum channels must have a thread_name or thread_id
	JSONErrorWebhookForumMustHaveThreadNameOrID = 220001

	// Webhooks posted to forum channels cannot have both a thread_name and thread_id
	JSONErrorWebhookForumCannotHaveBothThreadNameAndID = 220002

	// Webhooks can only create threads in forum channels
	JSONErrorWebhookCanOnlyCreateThreadsInForum = 220003

	// Webhook services cannot be used in forum channels
	JSONErrorWebhookServicesCannotBeUsedInForum = 220004

	// Message blocked by harmful links filter
	JSONErrorMessageBlockedByHarmfulLinksFilter = 240000
)
