// Package payloads provides Discord API payload types.
//
// This package contains data structures for Discord objects such as users,
// guilds, channels, messages, and other API entities.
package payloads

// Common types and constants for Discord API payloads

// PermissionFlagsBits represents Discord permission flags as constants.
//
// These flags are used for bitwise permission calculations.
// Each permission is represented as a power of 2.
//
// See: https://discord.com/developers/docs/topics/permissions#permissions-bitwise-permission-flags
var PermissionFlagsBits = struct {
	// CreateInstantInvite allows creation of instant invites.
	// Applies to channel types: Text, Voice, Stage
	CreateInstantInvite int64

	// KickMembers allows kicking members.
	KickMembers int64

	// BanMembers allows banning members.
	BanMembers int64

	// Administrator allows all permissions and bypasses channel permission overwrites.
	Administrator int64

	// ManageChannels allows management and editing of channels.
	// Applies to channel types: Text, Voice, Stage
	ManageChannels int64

	// ManageGuild allows management and editing of the guild.
	ManageGuild int64

	// AddReactions allows for the addition of reactions to messages.
	// Applies to channel types: Text, Voice, Stage
	AddReactions int64

	// ViewAuditLog allows for viewing of audit logs.
	ViewAuditLog int64

	// PrioritySpeaker allows for using priority speaker in a voice channel.
	// Applies to channel types: Voice
	PrioritySpeaker int64

	// Stream allows the user to go live.
	// Applies to channel types: Voice, Stage
	Stream int64

	// ViewChannel allows guild members to view a channel, which includes reading
	// messages in text channels and joining voice channels.
	// Applies to channel types: Text, Voice, Stage
	ViewChannel int64

	// SendMessages allows for sending messages in a channel and creating threads
	// in a forum (does not allow sending messages in threads).
	// Applies to channel types: Text, Voice, Stage
	SendMessages int64

	// SendTTSMessages allows for sending of /tts messages.
	// Applies to channel types: Text, Voice, Stage
	SendTTSMessages int64

	// ManageMessages allows for deletion of other users messages.
	// Applies to channel types: Text, Voice, Stage
	ManageMessages int64

	// EmbedLinks allows links sent by users with this permission to be auto-embedded.
	// Applies to channel types: Text, Voice, Stage
	EmbedLinks int64

	// AttachFiles allows for uploading images and files.
	// Applies to channel types: Text, Voice, Stage
	AttachFiles int64

	// ReadMessageHistory allows for reading of message history.
	// Applies to channel types: Text, Voice, Stage
	ReadMessageHistory int64

	// MentionEveryone allows for using the @everyone tag to notify all users
	// in a channel, and the @here tag to notify all online users in a channel.
	// Applies to channel types: Text, Voice, Stage
	MentionEveryone int64

	// UseExternalEmojis allows the usage of custom emojis from other servers.
	// Applies to channel types: Text, Voice, Stage
	UseExternalEmojis int64

	// ViewGuildInsights allows for viewing guild insights.
	ViewGuildInsights int64

	// Connect allows for joining of a voice channel.
	// Applies to channel types: Voice, Stage
	Connect int64

	// Speak allows for speaking in a voice channel.
	// Applies to channel types: Voice
	Speak int64

	// MuteMembers allows for muting members in a voice channel.
	// Applies to channel types: Voice, Stage
	MuteMembers int64

	// DeafenMembers allows for deafening of members in a voice channel.
	// Applies to channel types: Voice
	DeafenMembers int64

	// MoveMembers allows for moving of members between voice channels.
	// Applies to channel types: Voice, Stage
	MoveMembers int64

	// UseVAD allows for using voice-activity-detection in a voice channel.
	// Applies to channel types: Voice
	UseVAD int64

	// ChangeNickname allows for modification of own nickname.
	ChangeNickname int64

	// ManageNicknames allows for modification of other users nicknames.
	ManageNicknames int64

	// ManageRoles allows management and editing of roles.
	// Applies to channel types: Text, Voice, Stage
	ManageRoles int64

	// ManageWebhooks allows management and editing of webhooks.
	// Applies to channel types: Text, Voice, Stage
	ManageWebhooks int64

	// ManageGuildExpressions allows for editing and deleting emojis, stickers,
	// and soundboard sounds created by all users.
	ManageGuildExpressions int64

	// UseApplicationCommands allows members to use application commands,
	// including slash commands and context menu commands.
	// Applies to channel types: Text, Voice, Stage
	UseApplicationCommands int64

	// RequestToSpeak allows for requesting to speak in stage channels.
	// Applies to channel types: Stage
	RequestToSpeak int64

	// ManageEvents allows for editing and deleting scheduled events created by all users.
	// Applies to channel types: Voice, Stage
	ManageEvents int64

	// ManageThreads allows for deleting and archiving threads, and viewing all private threads.
	// Applies to channel types: Text
	ManageThreads int64

	// CreatePublicThreads allows for creating public and announcement threads.
	// Applies to channel types: Text
	CreatePublicThreads int64

	// CreatePrivateThreads allows for creating private threads.
	// Applies to channel types: Text
	CreatePrivateThreads int64

	// UseExternalStickers allows the usage of custom stickers from other servers.
	// Applies to channel types: Text, Voice, Stage
	UseExternalStickers int64

	// SendMessagesInThreads allows for sending messages in threads.
	// Applies to channel types: Text
	SendMessagesInThreads int64

	// UseEmbeddedActivities allows for using Activities (applications with the
	// Embedded flag) in a voice channel.
	// Applies to channel types: Voice
	UseEmbeddedActivities int64

	// ModerateMembers allows for timing out users to prevent them from sending
	// or reacting to messages in chat and threads, and from speaking in voice
	// and stage channels.
	ModerateMembers int64

	// ViewCreatorMonetizationAnalytics allows for viewing role subscription insights.
	ViewCreatorMonetizationAnalytics int64

	// UseSoundboard allows for using soundboard in a voice channel.
	// Applies to channel types: Voice
	UseSoundboard int64

	// CreateGuildExpressions allows for creating emojis, stickers, and soundboard sounds,
	// and editing and deleting those created by the current user.
	CreateGuildExpressions int64

	// CreateEvents allows for creating scheduled events, and editing and deleting
	// those created by the current user.
	// Applies to channel types: Voice, Stage
	CreateEvents int64

	// UseExternalSounds allows the usage of custom soundboard sounds from other servers.
	// Applies to channel types: Voice
	UseExternalSounds int64

	// SendVoiceMessages allows sending voice messages.
	// Applies to channel types: Text, Voice, Stage
	SendVoiceMessages int64

	// SendPolls allows sending polls.
	// Applies to channel types: Text, Voice, Stage
	SendPolls int64

	// UseExternalApps allows user-installed apps to send public responses.
	// When disabled, users will still be allowed to use their apps but the
	// responses will be ephemeral. This only applies to apps not also installed to the server.
	// Applies to channel types: Text, Voice, Stage
	UseExternalApps int64

	// PinMessages allows pinning and unpinning messages.
	// Applies to channel types: Text
	PinMessages int64
}{
	CreateInstantInvite:              1 << 0,
	KickMembers:                      1 << 1,
	BanMembers:                       1 << 2,
	Administrator:                    1 << 3,
	ManageChannels:                   1 << 4,
	ManageGuild:                      1 << 5,
	AddReactions:                     1 << 6,
	ViewAuditLog:                     1 << 7,
	PrioritySpeaker:                  1 << 8,
	Stream:                           1 << 9,
	ViewChannel:                      1 << 10,
	SendMessages:                     1 << 11,
	SendTTSMessages:                  1 << 12,
	ManageMessages:                   1 << 13,
	EmbedLinks:                       1 << 14,
	AttachFiles:                      1 << 15,
	ReadMessageHistory:               1 << 16,
	MentionEveryone:                  1 << 17,
	UseExternalEmojis:                1 << 18,
	ViewGuildInsights:                1 << 19,
	Connect:                          1 << 20,
	Speak:                            1 << 21,
	MuteMembers:                      1 << 22,
	DeafenMembers:                    1 << 23,
	MoveMembers:                      1 << 24,
	UseVAD:                           1 << 25,
	ChangeNickname:                   1 << 26,
	ManageNicknames:                  1 << 27,
	ManageRoles:                      1 << 28,
	ManageWebhooks:                   1 << 29,
	ManageGuildExpressions:           1 << 30,
	UseApplicationCommands:           1 << 31,
	RequestToSpeak:                   1 << 32,
	ManageEvents:                     1 << 33,
	ManageThreads:                    1 << 34,
	CreatePublicThreads:              1 << 35,
	CreatePrivateThreads:             1 << 36,
	UseExternalStickers:              1 << 37,
	SendMessagesInThreads:            1 << 38,
	UseEmbeddedActivities:            1 << 39,
	ModerateMembers:                  1 << 40,
	ViewCreatorMonetizationAnalytics: 1 << 41,
	UseSoundboard:                    1 << 42,
	CreateGuildExpressions:           1 << 43,
	CreateEvents:                     1 << 44,
	UseExternalSounds:                1 << 45,
	SendVoiceMessages:                1 << 46,
	SendPolls:                        1 << 49,
	UseExternalApps:                  1 << 50,
	PinMessages:                      1 << 51,
}

// LocalizationMap represents a mapping of locales to localized strings.
//
// See: https://discord.com/developers/docs/reference#locales
type LocalizationMap map[Locale]*string

// Locale represents a Discord locale.
type Locale string

// Discord supported locales
const (
	LocaleDanish        Locale = "da"
	LocaleGerman        Locale = "de"
	LocaleEnglishGB     Locale = "en-GB"
	LocaleEnglishUS     Locale = "en-US"
	LocaleSpanish       Locale = "es-ES"
	LocaleFrench        Locale = "fr"
	LocaleCroatian      Locale = "hr"
	LocaleItalian       Locale = "it"
	LocaleLithuanian    Locale = "lt"
	LocaleHungarian     Locale = "hu"
	LocaleDutch         Locale = "nl"
	LocaleNorwegian     Locale = "no"
	LocalePolish        Locale = "pl"
	LocalePortugueseBR  Locale = "pt-BR"
	LocaleRomanian      Locale = "ro"
	LocaleFinnish       Locale = "fi"
	LocaleSwedish       Locale = "sv-SE"
	LocaleVietnamese    Locale = "vi"
	LocaleTurkish       Locale = "tr"
	LocaleCzech         Locale = "cs"
	LocaleGreek         Locale = "el"
	LocaleBulgarian     Locale = "bg"
	LocaleRussian       Locale = "ru"
	LocaleUkrainian     Locale = "uk"
	LocaleHindi         Locale = "hi"
	LocaleThai          Locale = "th"
	LocaleChineseChina  Locale = "zh-CN"
	LocaleJapanese      Locale = "ja"
	LocaleChineseTaiwan Locale = "zh-TW"
	LocaleKorean        Locale = "ko"
)

// RESTError represents a Discord REST API error response.
//
// See: https://discord.com/developers/docs/topics/opcodes-and-status-codes#json
type RESTError struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Errors  *RESTErrorData `json:"errors,omitempty"`
}

// RESTErrorFieldInformation represents field-specific error information.
type RESTErrorFieldInformation struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// RESTErrorGroupWrapper represents a grouped error wrapper.
type RESTErrorGroupWrapper struct {
	Errors []RESTErrorData `json:"_errors"`
}

// RESTErrorData represents error data which can be a field info, group wrapper,
// string, or a map of nested errors.
type RESTErrorData interface{}

// RESTRateLimit represents a Discord rate limit response.
//
// See: https://discord.com/developers/docs/topics/rate-limits#exceeding-a-rate-limit-rate-limit-response-structure
type RESTRateLimit struct {
	// Code is an error code for some limits.
	Code *int `json:"code,omitempty"`

	// Global indicates if you are being globally rate limited or not.
	Global bool `json:"global"`

	// Message saying you are being rate limited.
	Message string `json:"message"`

	// RetryAfter is the number of seconds to wait before submitting another request.
	RetryAfter float64 `json:"retry_after"`
}
