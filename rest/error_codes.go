// Package rest provides Discord REST API types and utilities.
//
// This file contains comprehensive REST error codes from Discord API.
package rest

// RESTJSONErrorCode represents Discord API JSON error codes
type RESTJSONErrorCode int

// Discord API JSON Error Codes
// @see https://discord.com/developers/docs/topics/opcodes-and-status-codes#json-json-error-codes
const (
	GeneralError RESTJSONErrorCode = 0

	// 10xxx: Unknown/Invalid Resources
	UnknownAccount                        RESTJSONErrorCode = 10001
	UnknownApplication                    RESTJSONErrorCode = 10002
	UnknownChannel                        RESTJSONErrorCode = 10003
	UnknownGuild                          RESTJSONErrorCode = 10004
	UnknownIntegration                    RESTJSONErrorCode = 10005
	UnknownInvite                         RESTJSONErrorCode = 10006
	UnknownMember                         RESTJSONErrorCode = 10007
	UnknownMessage                        RESTJSONErrorCode = 10008
	UnknownPermissionOverwrite            RESTJSONErrorCode = 10009
	UnknownProvider                       RESTJSONErrorCode = 10010
	UnknownRole                           RESTJSONErrorCode = 10011
	UnknownToken                          RESTJSONErrorCode = 10012
	UnknownUser                           RESTJSONErrorCode = 10013
	UnknownEmoji                          RESTJSONErrorCode = 10014
	UnknownWebhook                        RESTJSONErrorCode = 10015
	UnknownWebhookService                 RESTJSONErrorCode = 10016
	UnknownSession                        RESTJSONErrorCode = 10020
	UnknownAsset                          RESTJSONErrorCode = 10021
	UnknownBan                            RESTJSONErrorCode = 10026
	UnknownSKU                            RESTJSONErrorCode = 10027
	UnknownStoreListing                   RESTJSONErrorCode = 10028
	UnknownEntitlement                    RESTJSONErrorCode = 10029
	UnknownBuild                          RESTJSONErrorCode = 10030
	UnknownLobby                          RESTJSONErrorCode = 10031
	UnknownBranch                         RESTJSONErrorCode = 10032
	UnknownStoreDirectoryLayout           RESTJSONErrorCode = 10033
	UnknownRedistributable                RESTJSONErrorCode = 10036
	UnknownGiftCode                       RESTJSONErrorCode = 10038
	UnknownStream                         RESTJSONErrorCode = 10049
	UnknownPremiumServerSubscribeCooldown RESTJSONErrorCode = 10050
	UnknownGuildTemplate                  RESTJSONErrorCode = 10057
	UnknownDiscoverableServerCategory     RESTJSONErrorCode = 10059
	UnknownSticker                        RESTJSONErrorCode = 10060
	UnknownStickerPack                    RESTJSONErrorCode = 10061
	UnknownInteraction                    RESTJSONErrorCode = 10062
	UnknownApplicationCommand             RESTJSONErrorCode = 10063
	UnknownVoiceState                     RESTJSONErrorCode = 10065
	UnknownApplicationCommandPermissions  RESTJSONErrorCode = 10066
	UnknownStageInstance                  RESTJSONErrorCode = 10067
	UnknownGuildMemberVerificationForm    RESTJSONErrorCode = 10068
	UnknownGuildWelcomeScreen             RESTJSONErrorCode = 10069
	UnknownGuildScheduledEvent            RESTJSONErrorCode = 10070
	UnknownGuildScheduledEventUser        RESTJSONErrorCode = 10071
	UnknownTag                            RESTJSONErrorCode = 10087
	UnknownSound                          RESTJSONErrorCode = 10097

	// 20xxx: Bot/Authorization Issues
	BotsCannotUseThisEndpoint                                                 RESTJSONErrorCode = 20001
	OnlyBotsCanUseThisEndpoint                                                RESTJSONErrorCode = 20002
	ExplicitContentCannotBeSentToTheDesiredRecipient                          RESTJSONErrorCode = 20009
	NotAuthorizedToPerformThisActionOnThisApplication                         RESTJSONErrorCode = 20012
	ActionCannotBePerformedDueToSlowmodeRateLimit                             RESTJSONErrorCode = 20016
	TheMazeIsntMeantForYou                                                    RESTJSONErrorCode = 20017
	OnlyTheOwnerOfThisAccountCanPerformThisAction                             RESTJSONErrorCode = 20018
	AnnouncementEditLimitExceeded                                             RESTJSONErrorCode = 20022
	UnderMinimumAge                                                           RESTJSONErrorCode = 20024
	ChannelSendRateLimit                                                      RESTJSONErrorCode = 20028
	ServerSendRateLimit                                                       RESTJSONErrorCode = 20029
	StageTopicServerNameServerDescriptionOrChannelNamesContainDisallowedWords RESTJSONErrorCode = 20031
	GuildPremiumSubscriptionLevelTooLow                                       RESTJSONErrorCode = 20035

	// 30xxx: Maximum Limits
	MaximumNumberOfGuildsReached                              RESTJSONErrorCode = 30001
	MaximumNumberOfFriendsReached                             RESTJSONErrorCode = 30002
	MaximumNumberOfPinsReachedForTheChannel                   RESTJSONErrorCode = 30003
	MaximumNumberOfRecipientsReached                          RESTJSONErrorCode = 30004
	MaximumNumberOfGuildRolesReached                          RESTJSONErrorCode = 30005
	MaximumNumberOfWebhooksReached                            RESTJSONErrorCode = 30007
	MaximumNumberOfEmojisReached                              RESTJSONErrorCode = 30008
	MaximumNumberOfReactionsReached                           RESTJSONErrorCode = 30010
	MaximumNumberOfGroupDMsReached                            RESTJSONErrorCode = 30011
	MaximumNumberOfGuildChannelsReached                       RESTJSONErrorCode = 30013
	MaximumNumberOfAttachmentsInAMessageReached               RESTJSONErrorCode = 30015
	MaximumNumberOfInvitesReached                             RESTJSONErrorCode = 30016
	MaximumNumberOfAnimatedEmojisReached                      RESTJSONErrorCode = 30018
	MaximumNumberOfServerMembersReached                       RESTJSONErrorCode = 30019
	MaximumNumberOfServerCategoriesReached                    RESTJSONErrorCode = 30030
	GuildAlreadyHasATemplate                                  RESTJSONErrorCode = 30031
	MaximumNumberOfApplicationCommandsReached                 RESTJSONErrorCode = 30032
	MaximumNumberOfThreadParticipantsReached                  RESTJSONErrorCode = 30033
	MaximumNumberOfDailyApplicationCommandCreatesReached      RESTJSONErrorCode = 30034
	MaximumNumberOfBansForNonGuildMembersHaveBeenExceeded     RESTJSONErrorCode = 30035
	MaximumNumberOfBansFetchesHasBeenReached                  RESTJSONErrorCode = 30037
	MaximumNumberOfUncompletedGuildScheduledEventsReached     RESTJSONErrorCode = 30038
	MaximumNumberOfStickersReached                            RESTJSONErrorCode = 30039
	MaximumNumberOfPruneRequestsHasBeenReached                RESTJSONErrorCode = 30040
	MaximumNumberOfGuildWidgetSettingsUpdatesHasBeenReached   RESTJSONErrorCode = 30042
	MaximumNumberOfEditsToMessagesOlderThanOneHourReached     RESTJSONErrorCode = 30046
	MaximumNumberOfPinnedThreadsInAForumChannelHasBeenReached RESTJSONErrorCode = 30047
	MaximumNumberOfTagsInAForumChannelHasBeenReached          RESTJSONErrorCode = 30048
	BitrateIsTooHighForChannelOfThisType                      RESTJSONErrorCode = 30052
	MaximumNumberOfPremiumEmojisReached                       RESTJSONErrorCode = 30056
	MaximumNumberOfWebhooksPerGuildReached                    RESTJSONErrorCode = 30058
	MaximumNumberOfChannelPermissionOverwritesReached         RESTJSONErrorCode = 30060
	TheChannelsForThisGuildAreTooLarge                        RESTJSONErrorCode = 30061
	MaximumNumberOfSoundboardSoundsReached                    RESTJSONErrorCode = 30065

	// 40xxx: Unauthorized/Missing Access
	Unauthorized                                       RESTJSONErrorCode = 40001
	VerifyYourAccount                                  RESTJSONErrorCode = 40002
	OpeningDirectMessagesTooFast                       RESTJSONErrorCode = 40003
	SendMessagesHasBeenTemporarilyDisabled             RESTJSONErrorCode = 40004
	RequestEntityTooLarge                              RESTJSONErrorCode = 40005
	FeatureTemporarilyDisabledServerSide               RESTJSONErrorCode = 40006
	UserBannedFromThisGuild                            RESTJSONErrorCode = 40007
	ConnectionHasBeenRevoked                           RESTJSONErrorCode = 40012
	TargetUserIsNotConnectedToVoice                    RESTJSONErrorCode = 40032
	ThisMessageHasAlreadyBeenCrossposted               RESTJSONErrorCode = 40033
	AnApplicationCommandWithThatNameAlreadyExists      RESTJSONErrorCode = 40041
	ApplicationInteractionFailedToSend                 RESTJSONErrorCode = 40043
	CannotSendAMessageInAForumChannel                  RESTJSONErrorCode = 40058
	InteractionHasAlreadyBeenAcknowledged              RESTJSONErrorCode = 40060
	TagNamesMustBeUnique                               RESTJSONErrorCode = 40061
	ServiceResourceIsBeingRateLimited                  RESTJSONErrorCode = 40062
	ThereAreNoTagsAvailableThatCanBeSetByNonModerators RESTJSONErrorCode = 40066
	TagRequiredToCreateAForumPostInThisChannel         RESTJSONErrorCode = 40067
	EntitlementAlreadyGranted                          RESTJSONErrorCode = 40074

	// 50xxx: Missing Permissions/Invalid Request Data
	MissingAccess                                                                 RESTJSONErrorCode = 50001
	InvalidAccountType                                                            RESTJSONErrorCode = 50002
	CannotExecuteActionOnADMChannel                                               RESTJSONErrorCode = 50003
	GuildWidgetDisabled                                                           RESTJSONErrorCode = 50004
	CannotEditAMessageAuthoredByAnotherUser                                       RESTJSONErrorCode = 50005
	CannotSendAnEmptyMessage                                                      RESTJSONErrorCode = 50006
	CannotSendMessagesToThisUser                                                  RESTJSONErrorCode = 50007
	CannotSendMessagesInANonTextChannel                                           RESTJSONErrorCode = 50008
	ChannelVerificationLevelIsTooHighForYouToGainAccess                           RESTJSONErrorCode = 50009
	OAuth2ApplicationDoesNotHaveABot                                              RESTJSONErrorCode = 50010
	OAuth2ApplicationLimitReached                                                 RESTJSONErrorCode = 50011
	InvalidOAuth2State                                                            RESTJSONErrorCode = 50012
	YouLackPermissionsToPerformThatAction                                         RESTJSONErrorCode = 50013
	InvalidAuthenticationToken                                                    RESTJSONErrorCode = 50014
	NoteTooLong                                                                   RESTJSONErrorCode = 50015
	ProvidedTooFewOrTooManyMessagesToDelete                                       RESTJSONErrorCode = 50016
	InvalidMFALevel                                                               RESTJSONErrorCode = 50017
	AMessageCanOnlyBePinnedToTheChannelItWasSentIn                                RESTJSONErrorCode = 50019
	InviteCodeWasInvalidOrTaken                                                   RESTJSONErrorCode = 50020
	CannotExecuteActionOnASystemMessage                                           RESTJSONErrorCode = 50021
	CannotExecuteActionOnThisChannelType                                          RESTJSONErrorCode = 50024
	InvalidOAuth2AccessTokenProvided                                              RESTJSONErrorCode = 50025
	MissingRequiredOAuth2Scope                                                    RESTJSONErrorCode = 50026
	InvalidWebhookTokenProvided                                                   RESTJSONErrorCode = 50027
	InvalidRole                                                                   RESTJSONErrorCode = 50028
	InvalidRecipients                                                             RESTJSONErrorCode = 50033
	AMessageProvidedWasTooOldToBulkDelete                                         RESTJSONErrorCode = 50034
	InvalidFormBodyOrContentType                                                  RESTJSONErrorCode = 50035
	AnInviteWasAcceptedToAGuildTheApplicationsBotIsNotIn                          RESTJSONErrorCode = 50036
	InvalidActivityAction                                                         RESTJSONErrorCode = 50039
	InvalidAPIVersionProvided                                                     RESTJSONErrorCode = 50041
	FileUploadedExceedsTheMaximumSize                                             RESTJSONErrorCode = 50045
	InvalidFileUploaded                                                           RESTJSONErrorCode = 50046
	CannotSelfRedeemThisGift                                                      RESTJSONErrorCode = 50054
	InvalidGuild                                                                  RESTJSONErrorCode = 50055
	InvalidMessageType                                                            RESTJSONErrorCode = 50068
	PaymentSourceRequiredToRedeemGift                                             RESTJSONErrorCode = 50070
	CannotModifyASystemWebhook                                                    RESTJSONErrorCode = 50073
	CannotDeleteAChannelRequiredForCommunityGuilds                                RESTJSONErrorCode = 50074
	CannotEditStickersWithinAMessage                                              RESTJSONErrorCode = 50080
	InvalidStickerSent                                                            RESTJSONErrorCode = 50081
	TriedToPerformAnOperationOnAnArchivedThread                                   RESTJSONErrorCode = 50083
	InvalidThreadNotificationSettings                                             RESTJSONErrorCode = 50084
	BeforeValueIsEarlierThanTheThreadCreationDate                                 RESTJSONErrorCode = 50085
	CommunityServerChannelsMustBeTextChannels                                     RESTJSONErrorCode = 50086
	TheEntityTypeOfTheEventIsDifferentFromTheEntityYouAreTryingToStartTheEventFor RESTJSONErrorCode = 50091
	ServerNotAvailableInYourLocation                                              RESTJSONErrorCode = 50095
	ServerNeedsMonetizationEnabledInOrderToPerformThisAction                      RESTJSONErrorCode = 50097
	ServerNeedsMoreBoostsToPerformThisAction                                      RESTJSONErrorCode = 50101
	RequestBodyContainsInvalidJSON                                                RESTJSONErrorCode = 50109
	OwnershipCannotBeMovedToABotUser                                              RESTJSONErrorCode = 50132
	FailedToResizeAssetBelowTheMaximumSize                                        RESTJSONErrorCode = 50138
	CannotMixSubscriptionAndNonSubscriptionRolesForAnEmoji                        RESTJSONErrorCode = 50144
	CannotConvertBetweenPremiumEmojiAndNormalEmoji                                RESTJSONErrorCode = 50145
	UploadedFileNotFound                                                          RESTJSONErrorCode = 50146
	VoiceMessagesDoNotSupportAdditionalContent                                    RESTJSONErrorCode = 50159
	VoiceMessagesDoNotSupportAdditionalEmbeds                                     RESTJSONErrorCode = 50160
	VoiceMessagesDoNotSupportAdditionalComponents                                 RESTJSONErrorCode = 50161
	CannotEditVoiceMessages                                                       RESTJSONErrorCode = 50162
	YouCannotSendVoiceMessagesInThisChannel                                       RESTJSONErrorCode = 50173
	TheUserAccountMustFirstBeVerified                                             RESTJSONErrorCode = 50178
	YouDoNotHavePermissionToSendThisSticker                                       RESTJSONErrorCode = 50600
	YouCannotUseAnApplicationCommandInASelfDeprecatedChannel                      RESTJSONErrorCode = 70003
	YouCannotUseApplicationCommandsInASelfDeprecatedChannel                       RESTJSONErrorCode = 70004

	// 60xxx: Two-Factor Authentication Required
	TwoFactorAuthenticationIsRequiredForThisOperation RESTJSONErrorCode = 60003

	// 80xxx: Feature Availability
	NoUsersWithDiscordTagExist  RESTJSONErrorCode = 80004
	ReactionWasBlocked          RESTJSONErrorCode = 90001
	UserCannotUseBurstReactions RESTJSONErrorCode = 90002

	// 110xxx: AutoMod
	APIResourceIsCurrentlyOverloaded RESTJSONErrorCode = 110001

	// 130xxx: Stage Channels
	TheStageIsAlreadyOpen RESTJSONErrorCode = 130000

	// 150xxx: Thread/Forum Limits
	CannotReplyWithoutPermissionToReadMessageHistory RESTJSONErrorCode = 160002
	ThreadHasAlreadyBeenCreatedForThisMessage        RESTJSONErrorCode = 160004
	ThreadIsLocked                                   RESTJSONErrorCode = 160005
	MaximumNumberOfActiveThreadsReached              RESTJSONErrorCode = 160006
	MaximumNumberOfActiveAnnouncementThreadsReached  RESTJSONErrorCode = 160007

	// 170xxx: Webhook Issues
	WebhooksCanOnlyCreateThreadsInForumChannels RESTJSONErrorCode = 170001
	WebhookServicesCannotBeUsedInForumChannels  RESTJSONErrorCode = 170002

	// 180xxx: Scheduled Events
	CannotDeleteAGuildSubscriptionIntegration RESTJSONErrorCode = 180000

	// 200xxx: Announcements/Monetization
	YouCannotSendAMessageUsingAnInactiveWebhook RESTJSONErrorCode = 200001

	// 220xxx: Message/Content Issues
	MessageWasBlockedByAutomaticModeration                            RESTJSONErrorCode = 220001
	TitleWasBlockedByAutomaticModeration                              RESTJSONErrorCode = 220002
	WebhooksPostedToForumChannelsMustHaveAThreadNameOrThreadID        RESTJSONErrorCode = 220003
	WebhooksPostedToForumChannelsCannotHaveBothAThreadNameAndThreadID RESTJSONErrorCode = 220004
	WebhooksCanOnlyCreateThreadsInForumOrMediaChannels                RESTJSONErrorCode = 220005

	// 230xxx: Policy/Harmful Links
	HarmfulLinksDetected            RESTJSONErrorCode = 230001
	ApplicationCommandNameIsInvalid RESTJSONErrorCode = 230002
)

// String returns a human-readable description of the error code
func (e RESTJSONErrorCode) String() string {
	switch e {
	case GeneralError:
		return "General Error"
	case UnknownAccount:
		return "Unknown Account"
	case UnknownApplication:
		return "Unknown Application"
	case UnknownChannel:
		return "Unknown Channel"
	case UnknownGuild:
		return "Unknown Guild"
	case UnknownIntegration:
		return "Unknown Integration"
	case UnknownInvite:
		return "Unknown Invite"
	case UnknownMember:
		return "Unknown Member"
	case UnknownMessage:
		return "Unknown Message"
	case UnknownPermissionOverwrite:
		return "Unknown Permission Overwrite"
	case UnknownProvider:
		return "Unknown Provider"
	case UnknownRole:
		return "Unknown Role"
	case UnknownToken:
		return "Unknown Token"
	case UnknownUser:
		return "Unknown User"
	case UnknownEmoji:
		return "Unknown Emoji"
	case UnknownWebhook:
		return "Unknown Webhook"
	case BotsCannotUseThisEndpoint:
		return "Bots cannot use this endpoint"
	case OnlyBotsCanUseThisEndpoint:
		return "Only bots can use this endpoint"
	case ExplicitContentCannotBeSentToTheDesiredRecipient:
		return "Explicit content cannot be sent to the desired recipient(s)"
	case Unauthorized:
		return "Unauthorized. Provide a valid token and try again"
	case MissingAccess:
		return "Missing Access"
	case InvalidAccountType:
		return "Invalid account type"
	case CannotExecuteActionOnADMChannel:
		return "Cannot execute action on a DM channel"
	case GuildWidgetDisabled:
		return "Guild widget disabled"
	case CannotEditAMessageAuthoredByAnotherUser:
		return "Cannot edit a message authored by another user"
	case CannotSendAnEmptyMessage:
		return "Cannot send an empty message"
	case CannotSendMessagesToThisUser:
		return "Cannot send messages to this user"
	case CannotSendMessagesInANonTextChannel:
		return "Cannot send messages in a voice channel"
	case ChannelVerificationLevelIsTooHighForYouToGainAccess:
		return "Channel verification level is too high for you to gain access"
	case OAuth2ApplicationDoesNotHaveABot:
		return "OAuth2 application does not have a bot"
	case OAuth2ApplicationLimitReached:
		return "OAuth2 application limit reached"
	case InvalidOAuth2State:
		return "Invalid OAuth2 state"
	case YouLackPermissionsToPerformThatAction:
		return "You lack permissions to perform that action"
	case InvalidAuthenticationToken:
		return "Invalid authentication token provided"
	case NoteTooLong:
		return "Note was too long"
	case ProvidedTooFewOrTooManyMessagesToDelete:
		return "Provided too few or too many messages to delete. Must provide at least 2 and fewer than 100 messages to delete"
	case MaximumNumberOfGuildsReached:
		return "Maximum number of guilds reached (100)"
	case MaximumNumberOfFriendsReached:
		return "Maximum number of friends reached (1000)"
	case MaximumNumberOfPinsReachedForTheChannel:
		return "Maximum number of pins reached for the channel (50)"
	case MaximumNumberOfRecipientsReached:
		return "Maximum number of recipients reached (10)"
	case MaximumNumberOfGuildRolesReached:
		return "Maximum number of guild roles reached (250)"
	case MaximumNumberOfWebhooksReached:
		return "Maximum number of webhooks reached (10)"
	case MaximumNumberOfEmojisReached:
		return "Maximum number of emojis reached"
	case MaximumNumberOfReactionsReached:
		return "Maximum number of reactions reached (20)"
	case MaximumNumberOfGroupDMsReached:
		return "Maximum number of group DMs reached (10)"
	case MaximumNumberOfGuildChannelsReached:
		return "Maximum number of guild channels reached (500)"
	case MaximumNumberOfAttachmentsInAMessageReached:
		return "Maximum number of attachments in a message reached (10)"
	case MaximumNumberOfInvitesReached:
		return "Maximum number of invites reached (1000)"
	case TwoFactorAuthenticationIsRequiredForThisOperation:
		return "Two factor authentication is required for this operation"
	default:
		return "Unknown Error"
	}
}
