package payloads

import "github.com/kolosys/discord-types/discord"

// BaseGuild represents the base structure for all guild types.
type BaseGuild struct {
	// ID is the guild id.
	ID discord.Snowflake `json:"id"`
}

// UnavailableGuild represents a guild that is unavailable due to an outage.
//
// See: https://discord.com/developers/docs/resources/guild#unavailable-guild-object
type UnavailableGuild struct {
	BaseGuild
	// Unavailable indicates if this guild is unavailable due to an outage.
	Unavailable bool `json:"unavailable"`
}

// PartialGuild represents a partial guild object.
//
// See: https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type PartialGuild struct {
	BaseGuild

	// Name is the guild name (2-100 characters, excluding trailing and leading whitespace).
	Name string `json:"name"`

	// Icon is the icon hash.
	//
	// See: https://discord.com/developers/docs/reference#image-formatting
	Icon *string `json:"icon"`

	// Splash is the splash hash.
	//
	// See: https://discord.com/developers/docs/reference#image-formatting
	Splash *string `json:"splash"`

	// Banner is the banner hash.
	//
	// See: https://discord.com/developers/docs/reference#image-formatting
	Banner *string `json:"banner,omitempty"`

	// Description is the description for the guild.
	Description *string `json:"description,omitempty"`

	// Features are the enabled guild features.
	//
	// See: https://discord.com/developers/docs/resources/guild#guild-object-guild-features
	Features []GuildFeature `json:"features,omitempty"`

	// VerificationLevel is the verification level required for the guild.
	//
	// See: https://discord.com/developers/docs/resources/guild#guild-object-verification-level
	VerificationLevel *GuildVerificationLevel `json:"verification_level,omitempty"`

	// VanityURLCode is the vanity url code for the guild.
	VanityURLCode *string `json:"vanity_url_code,omitempty"`

	// WelcomeScreen is the welcome screen of a Community guild, shown to new members.
	// Returned in the invite object.
	WelcomeScreen *GuildWelcomeScreen `json:"welcome_screen,omitempty"`
}

// Guild represents a complete Discord guild.
//
// See: https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type Guild struct {
	PartialGuild

	// IconHash is the icon hash, returned when in the template object.
	//
	// See: https://discord.com/developers/docs/reference#image-formatting
	IconHash *string `json:"icon_hash,omitempty"`

	// DiscoverySplash is the discovery splash hash; only present for guilds with the "DISCOVERABLE" feature.
	//
	// See: https://discord.com/developers/docs/reference#image-formatting
	DiscoverySplash *string `json:"discovery_splash"`

	// Owner indicates if the user is the owner of the guild.
	// This field is only received from the Get Current User Guilds endpoint.
	Owner *bool `json:"owner,omitempty"`

	// OwnerID is the ID of owner.
	OwnerID discord.Snowflake `json:"owner_id"`

	// Permissions are the total permissions for the user in the guild (excludes overrides).
	// This field is only received from the Get Current User Guilds endpoint.
	Permissions *discord.Permissions `json:"permissions,omitempty"`

	// Region is the voice region id for the guild.
	// Deprecated: This field has been deprecated in favor of rtc_region on the channel.
	Region string `json:"region"`

	// AFKChannelID is the ID of afk channel.
	AFKChannelID *discord.Snowflake `json:"afk_channel_id"`

	// AFKTimeout is the afk timeout in seconds.
	AFKTimeout AFKTimeout `json:"afk_timeout"`

	// WidgetEnabled indicates if the guild widget is enabled.
	WidgetEnabled *bool `json:"widget_enabled,omitempty"`

	// WidgetChannelID is the channel id that the widget will generate an invite to, or null if set to no invite.
	WidgetChannelID *discord.Snowflake `json:"widget_channel_id,omitempty"`

	// VerificationLevel is the verification level required for the guild.
	VerificationLevel GuildVerificationLevel `json:"verification_level"`

	// DefaultMessageNotifications is the default message notifications level.
	DefaultMessageNotifications GuildDefaultMessageNotifications `json:"default_message_notifications"`

	// ExplicitContentFilter is the explicit content filter level.
	ExplicitContentFilter GuildExplicitContentFilter `json:"explicit_content_filter"`

	// Roles are the roles in the guild.
	Roles []Role `json:"roles"`

	// Emojis are the custom guild emojis.
	Emojis []Emoji `json:"emojis"`

	// Features are the enabled guild features.
	Features []GuildFeature `json:"features"`

	// MFALevel is the required MFA level for the guild.
	MFALevel GuildMFALevel `json:"mfa_level"`

	// ApplicationID is the application id of the guild creator if it is bot-created.
	ApplicationID *discord.Snowflake `json:"application_id"`

	// SystemChannelID is the id of the channel where guild notices such as welcome messages and boost events are posted.
	SystemChannelID *discord.Snowflake `json:"system_channel_id"`

	// SystemChannelFlags are the system channel flags.
	SystemChannelFlags GuildSystemChannelFlags `json:"system_channel_flags"`

	// RulesChannelID is the id of the channel where Community guilds can display rules and/or guidelines.
	RulesChannelID *discord.Snowflake `json:"rules_channel_id"`

	// MaxPresences is the maximum number of presences for the guild (null is always returned, apart from the largest of guilds).
	MaxPresences *int `json:"max_presences,omitempty"`

	// MaxMembers is the maximum number of members for the guild.
	MaxMembers *int `json:"max_members,omitempty"`

	// VanityURLCode is the vanity url code for the guild.
	VanityURLCode *string `json:"vanity_url_code"`

	// Description is the description for the guild.
	Description *string `json:"description"`

	// Banner is the banner hash.
	Banner *string `json:"banner"`

	// PremiumTier is the premium tier (Server Boost level).
	PremiumTier GuildPremiumTier `json:"premium_tier"`

	// PremiumSubscriptionCount is the number of boosts this guild currently has.
	PremiumSubscriptionCount *int `json:"premium_subscription_count,omitempty"`

	// PreferredLocale is the preferred locale of a Community guild; used in guild discovery and notices from Discord.
	PreferredLocale Locale `json:"preferred_locale"`

	// PublicUpdatesChannelID is the id of the channel where admins and moderators of Community guilds receive notices from Discord.
	PublicUpdatesChannelID *discord.Snowflake `json:"public_updates_channel_id"`

	// MaxVideoChannelUsers is the maximum amount of users in a video channel.
	MaxVideoChannelUsers *int `json:"max_video_channel_users,omitempty"`

	// MaxStageVideoChannelUsers is the maximum amount of users in a stage video channel.
	MaxStageVideoChannelUsers *int `json:"max_stage_video_channel_users,omitempty"`

	// ApproximateMemberCount is the approximate number of members in this guild.
	ApproximateMemberCount *int `json:"approximate_member_count,omitempty"`

	// ApproximatePresenceCount is the approximate number of non-offline members in this guild.
	ApproximatePresenceCount *int `json:"approximate_presence_count,omitempty"`

	// NSFWLevel is the nsfw level of the guild.
	NSFWLevel GuildNSFWLevel `json:"nsfw_level"`

	// Stickers are the custom guild stickers.
	Stickers []Sticker `json:"stickers"`

	// PremiumProgressBarEnabled indicates whether the guild has the boost progress bar enabled.
	PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`

	// HubType is the type of Student Hub the guild is.
	HubType *GuildHubType `json:"hub_type"`

	// SafetyAlertsChannelID is the id of the channel where admins and moderators of Community guilds receive safety alerts from Discord.
	SafetyAlertsChannelID *discord.Snowflake `json:"safety_alerts_channel_id"`

	// IncidentsData is the incidents data for this guild.
	IncidentsData *IncidentsData `json:"incidents_data"`
}

// AFKTimeout represents valid AFK timeout values in seconds.
type AFKTimeout int

// AFK timeout constants
const (
	AFKTimeout60   AFKTimeout = 60
	AFKTimeout300  AFKTimeout = 300
	AFKTimeout900  AFKTimeout = 900
	AFKTimeout1800 AFKTimeout = 1800
	AFKTimeout3600 AFKTimeout = 3600
)

// GuildDefaultMessageNotifications represents default message notification levels.
type GuildDefaultMessageNotifications int

// Default message notification constants
const (
	GuildDefaultMessageNotificationsAllMessages GuildDefaultMessageNotifications = iota
	GuildDefaultMessageNotificationsOnlyMentions
)

// GuildExplicitContentFilter represents explicit content filter levels.
type GuildExplicitContentFilter int

// Explicit content filter constants
const (
	GuildExplicitContentFilterDisabled GuildExplicitContentFilter = iota
	GuildExplicitContentFilterMembersWithoutRoles
	GuildExplicitContentFilterAllMembers
)

// GuildMFALevel represents MFA levels.
type GuildMFALevel int

// MFA level constants
const (
	GuildMFALevelNone GuildMFALevel = iota
	GuildMFALevelElevated
)

// GuildNSFWLevel represents NSFW levels.
type GuildNSFWLevel int

// NSFW level constants
const (
	GuildNSFWLevelDefault GuildNSFWLevel = iota
	GuildNSFWLevelExplicit
	GuildNSFWLevelSafe
	GuildNSFWLevelAgeRestricted
)

// GuildVerificationLevel represents verification levels.
type GuildVerificationLevel int

// Verification level constants
const (
	GuildVerificationLevelNone GuildVerificationLevel = iota
	GuildVerificationLevelLow
	GuildVerificationLevelMedium
	GuildVerificationLevelHigh
	GuildVerificationLevelVeryHigh
)

// GuildPremiumTier represents premium tiers.
type GuildPremiumTier int

// Premium tier constants
const (
	GuildPremiumTierNone GuildPremiumTier = iota
	GuildPremiumTierTier1
	GuildPremiumTierTier2
	GuildPremiumTierTier3
)

// GuildHubType represents Student Hub types.
type GuildHubType int

// Hub type constants
const (
	GuildHubTypeDefault GuildHubType = iota
	GuildHubTypeHighSchool
	GuildHubTypeCollege
)

// GuildSystemChannelFlags represents system channel flags.
type GuildSystemChannelFlags int

// System channel flag constants
const (
	GuildSystemChannelFlagsSuppressJoinNotifications GuildSystemChannelFlags = 1 << iota
	GuildSystemChannelFlagsSuppressPremiumSubscriptions
	GuildSystemChannelFlagsSuppressGuildReminderNotifications
	GuildSystemChannelFlagsSuppressJoinNotificationReplies
	GuildSystemChannelFlagsSuppressRoleSubscriptionPurchaseNotifications
	GuildSystemChannelFlagsSuppressRoleSubscriptionPurchaseNotificationReplies
)

// GuildFeature represents guild features.
type GuildFeature string

// Guild feature constants
const (
	GuildFeatureAnimatedBanner                        GuildFeature = "ANIMATED_BANNER"
	GuildFeatureAnimatedIcon                          GuildFeature = "ANIMATED_ICON"
	GuildFeatureApplicationCommandPermissionsV2       GuildFeature = "APPLICATION_COMMAND_PERMISSIONS_V2"
	GuildFeatureAutoModeration                        GuildFeature = "AUTO_MODERATION"
	GuildFeatureBanner                                GuildFeature = "BANNER"
	GuildFeatureCommunity                             GuildFeature = "COMMUNITY"
	GuildFeatureCreatorMonetizableProvisional         GuildFeature = "CREATOR_MONETIZABLE_PROVISIONAL"
	GuildFeatureCreatorStorePage                      GuildFeature = "CREATOR_STORE_PAGE"
	GuildFeatureDeveloperSupportServer                GuildFeature = "DEVELOPER_SUPPORT_SERVER"
	GuildFeatureDiscoverable                          GuildFeature = "DISCOVERABLE"
	GuildFeatureFeaturable                            GuildFeature = "FEATURABLE"
	GuildFeatureHasDirectoryEntry                     GuildFeature = "HAS_DIRECTORY_ENTRY"
	GuildFeatureHub                                   GuildFeature = "HUB"
	GuildFeatureInvitesDisabled                       GuildFeature = "INVITES_DISABLED"
	GuildFeatureInviteSplash                          GuildFeature = "INVITE_SPLASH"
	GuildFeatureLinkedToHub                           GuildFeature = "LINKED_TO_HUB"
	GuildFeatureMemberVerificationGateEnabled         GuildFeature = "MEMBER_VERIFICATION_GATE_ENABLED"
	GuildFeatureMoreSoundboard                        GuildFeature = "MORE_SOUNDBOARD"
	GuildFeatureMonetizationEnabled                   GuildFeature = "MONETIZATION_ENABLED"
	GuildFeatureMoreStickers                          GuildFeature = "MORE_STICKERS"
	GuildFeatureNews                                  GuildFeature = "NEWS"
	GuildFeaturePartnered                             GuildFeature = "PARTNERED"
	GuildFeaturePreviewEnabled                        GuildFeature = "PREVIEW_ENABLED"
	GuildFeaturePrivateThreads                        GuildFeature = "PRIVATE_THREADS"
	GuildFeatureRaidAlertsDisabled                    GuildFeature = "RAID_ALERTS_DISABLED"
	GuildFeatureRelayEnabled                          GuildFeature = "RELAY_ENABLED"
	GuildFeatureRoleIcons                             GuildFeature = "ROLE_ICONS"
	GuildFeatureRoleSubscriptionsAvailableForPurchase GuildFeature = "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE"
	GuildFeatureRoleSubscriptionsEnabled              GuildFeature = "ROLE_SUBSCRIPTIONS_ENABLED"
	GuildFeatureSoundboard                            GuildFeature = "SOUNDBOARD"
	GuildFeatureTicketedEventsEnabled                 GuildFeature = "TICKETED_EVENTS_ENABLED"
	GuildFeatureVanityURL                             GuildFeature = "VANITY_URL"
	GuildFeatureVerified                              GuildFeature = "VERIFIED"
	GuildFeatureVIPRegions                            GuildFeature = "VIP_REGIONS"
	GuildFeatureWelcomeScreenEnabled                  GuildFeature = "WELCOME_SCREEN_ENABLED"
	GuildFeatureGuildTags                             GuildFeature = "GUILD_TAGS"
	GuildFeatureEnhancedRoleColors                    GuildFeature = "ENHANCED_ROLE_COLORS"
	GuildFeatureGuestsEnabled                         GuildFeature = "GUESTS_ENABLED"
)

// Role represents a Discord role.
//
// See: https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	// ID is the role id.
	ID discord.Snowflake `json:"id"`

	// Name is the role name.
	Name string `json:"name"`

	// Color is the integer representation of hexadecimal color code.
	Color int `json:"color"`

	// Hoist is whether this role is pinned in the user listing.
	Hoist bool `json:"hoist"`

	// Icon is the role icon hash.
	Icon *string `json:"icon,omitempty"`

	// UnicodeEmoji is the role unicode emoji.
	UnicodeEmoji *string `json:"unicode_emoji,omitempty"`

	// Position is the position of this role.
	Position int `json:"position"`

	// Permissions is the permission bit set.
	Permissions discord.Permissions `json:"permissions"`

	// Managed is whether this role is managed by an integration.
	Managed bool `json:"managed"`

	// Mentionable is whether this role is mentionable.
	Mentionable bool `json:"mentionable"`

	// Tags are the tags this role has.
	Tags *RoleTags `json:"tags,omitempty"`
}

// RoleTags represents role tags.
type RoleTags struct {
	// BotID is the id of the bot this role belongs to.
	BotID *discord.Snowflake `json:"bot_id,omitempty"`

	// IntegrationID is the id of the integration this role belongs to.
	IntegrationID *discord.Snowflake `json:"integration_id,omitempty"`

	// PremiumSubscriberID is the id of this role's subscription sku and level.
	PremiumSubscriberID *discord.Snowflake `json:"premium_subscriber_id,omitempty"`

	// SubscriptionListingID is the id of the subscription listing this role belongs to.
	SubscriptionListingID *discord.Snowflake `json:"subscription_listing_id,omitempty"`

	// AvailableForPurchase is whether this role is available for purchase.
	AvailableForPurchase *bool `json:"available_for_purchase,omitempty"`

	// GuildConnections is whether this role is a guild's linked role.
	GuildConnections *bool `json:"guild_connections,omitempty"`
}

// GuildWelcomeScreen represents a guild welcome screen.
//
// See: https://discord.com/developers/docs/resources/guild#welcome-screen-object
type GuildWelcomeScreen struct {
	// Description is the server description shown in the welcome screen.
	Description *string `json:"description,omitempty"`

	// WelcomeChannels are the channels shown in the welcome screen.
	WelcomeChannels []GuildWelcomeScreenChannel `json:"welcome_channels"`
}

// GuildWelcomeScreenChannel represents a welcome screen channel.
type GuildWelcomeScreenChannel struct {
	// ChannelID is the channel's id.
	ChannelID discord.Snowflake `json:"channel_id"`

	// Description is the description shown for the channel.
	Description string `json:"description"`

	// EmojiID is the emoji id, if the emoji is custom.
	EmojiID *discord.Snowflake `json:"emoji_id,omitempty"`

	// EmojiName is the emoji name if custom, the unicode character if standard,
	// or null if no emoji is set.
	EmojiName *string `json:"emoji_name,omitempty"`
}

// IncidentsData represents guild incident data.
type IncidentsData struct {
	// Incidents is the list of incidents.
	Incidents []GuildIncident `json:"incidents"`
}

// GuildIncident represents a guild incident.
type GuildIncident struct {
	// ID is the incident id.
	ID string `json:"id"`

	// Type is the incident type.
	Type string `json:"type"`

	// Status is the incident status.
	Status string `json:"status"`

	// CreatedAt is when the incident was created.
	CreatedAt string `json:"created_at"`

	// UpdatedAt is when the incident was last updated.
	UpdatedAt string `json:"updated_at"`

	// ResolvedAt is when the incident was resolved.
	ResolvedAt *string `json:"resolved_at,omitempty"`

	// Shortlink is the incident shortlink.
	Shortlink string `json:"shortlink"`

	// Title is the incident title.
	Title string `json:"title"`

	// Description is the incident description.
	Description string `json:"description"`

	// Updates are the incident updates.
	Updates []GuildIncidentUpdate `json:"updates"`
}

// GuildIncidentUpdate represents a guild incident update.
type GuildIncidentUpdate struct {
	// ID is the update id.
	ID string `json:"id"`

	// Status is the update status.
	Status string `json:"status"`

	// Body is the update body.
	Body string `json:"body"`

	// CreatedAt is when the update was created.
	CreatedAt string `json:"created_at"`

	// UpdatedAt is when the update was last updated.
	UpdatedAt string `json:"updated_at"`

	// DisplayAt is when the update should be displayed.
	DisplayAt string `json:"display_at"`

	// AffectedComponents are the affected components.
	AffectedComponents []GuildIncidentAffectedComponent `json:"affected_components"`
}

// GuildIncidentAffectedComponent represents an affected component in a guild incident.
type GuildIncidentAffectedComponent struct {
	// Code is the component code.
	Code string `json:"code"`

	// Name is the component name.
	Name string `json:"name"`

	// OldStatus is the old status.
	OldStatus string `json:"old_status"`

	// NewStatus is the new status.
	NewStatus string `json:"new_status"`
}
