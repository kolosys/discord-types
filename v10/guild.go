package v10

// BaseGuild represents the base structure for all guild types.
type BaseGuild struct {
	// ID is the guild id.
	ID Snowflake `json:"id"`
}

// UnavailableGuild represents a guild that is unavailable due to an outage.
//
// Reference: https://discord.com/developers/docs/resources/guild#unavailable-guild-object
type UnavailableGuild struct {
	BaseGuild
	// Unavailable indicates if this guild is unavailable due to an outage.
	Unavailable bool `json:"unavailable"`
}

// PartialGuild represents a partial guild object.
//
// Reference: https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type PartialGuild struct {
	BaseGuild

	// Name is the guild name (2-100 characters, excluding trailing and leading whitespace).
	Name string `json:"name"`

	// Icon is the icon hash.
	//
	// Reference: https://discord.com/developers/docs/reference#image-formatting
	Icon *string `json:"icon"`

	// Splash is the splash hash.
	//
	// Reference: https://discord.com/developers/docs/reference#image-formatting
	Splash *string `json:"splash"`

	// Banner is the banner hash.
	//
	// Reference: https://discord.com/developers/docs/reference#image-formatting
	Banner *string `json:"banner,omitempty"`

	// Description is the description for the guild.
	Description *string `json:"description,omitempty"`

	// Features are the enabled guild features.
	//
	// Reference: https://discord.com/developers/docs/resources/guild#guild-object-guild-features
	Features []GuildFeature `json:"features,omitempty"`

	// VerificationLevel is the verification level required for the guild.
	//
	// Reference: https://discord.com/developers/docs/resources/guild#guild-object-verification-level
	VerificationLevel *GuildVerificationLevel `json:"verification_level,omitempty"`

	// VanityURLCode is the vanity url code for the guild.
	VanityURLCode *string `json:"vanity_url_code,omitempty"`

	// WelcomeScreen is the welcome screen of a Community guild, shown to new members.
	//
	// Reference: https://discord.com/developers/docs/resources/guild#guild-object-welcome-screen
	WelcomeScreen *GuildWelcomeScreen `json:"welcome_screen,omitempty"`
}

// Guild represents a complete Discord guild.
//
// Reference: https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type Guild struct {
	PartialGuild

	// IconHash is the icon hash, returned when in the template object.
	//
	// Reference: https://discord.com/developers/docs/reference#image-formatting
	IconHash *string `json:"icon_hash,omitempty"`

	// DiscoverySplash is the discovery splash hash; only present for guilds with the "DISCOVERABLE" feature.
	//
	// Reference: https://discord.com/developers/docs/reference#image-formatting
	DiscoverySplash *string `json:"discovery_splash"`

	// Owner indicates if the user is the owner of the guild.
	//
	// This field is only received from the Get Current User Guilds endpoint.
	Owner *bool `json:"owner,omitempty"`

	// OwnerID is the ID of owner.
	OwnerID Snowflake `json:"owner_id"`

	// Permissions are the total permissions for the user in the guild (excludes overrides).
	//
	// This field is only received from the Get Current User Guilds endpoint.
	Permissions *Permissions `json:"permissions,omitempty"`

	// Region is the voice region id for the guild.
	//
	// Deprecated: This field has been deprecated in favor of rtc_region on the channel.
	Region string `json:"region"`

	// AFKChannelID is the ID of afk channel.
	AFKChannelID *Snowflake `json:"afk_channel_id"`

	// AFKTimeout is the afk timeout in seconds.
	AFKTimeout AFKTimeout `json:"afk_timeout"`

	// WidgetEnabled indicates if the guild widget is enabled.
	WidgetEnabled *bool `json:"widget_enabled,omitempty"`

	// WidgetChannelID is the channel id that the widget will generate an invite to, or null if set to no invite.
	WidgetChannelID *Snowflake `json:"widget_channel_id,omitempty"`

	// VerificationLevel is the verification level required for the guild.
	VerificationLevel GuildVerificationLevel `json:"verification_level"`

	// DefaultMessageNotifications is the default message notifications level.
	DefaultMessageNotifications GuildMessageNotification `json:"default_message_notifications"`

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
	ApplicationID *Snowflake `json:"application_id"`

	// SystemChannelID is the id of the channel where guild notices such as welcome messages and boost events are posted.
	SystemChannelID *Snowflake `json:"system_channel_id"`

	// SystemChannelFlags are the system channel flags.
	SystemChannelFlags GuildSystemChannelFlags `json:"system_channel_flags"`

	// RulesChannelID is the id of the channel where Community guilds can display rules and/or guidelines.
	RulesChannelID *Snowflake `json:"rules_channel_id"`

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
	PublicUpdatesChannelID *Snowflake `json:"public_updates_channel_id"`

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
	SafetyAlertsChannelID *Snowflake `json:"safety_alerts_channel_id"`

	// IncidentsData is the incidents data for this guild.
	IncidentsData *IncidentsData `json:"incidents_data"`
}

// Reference: https://discord.com/developers/docs/resources/guild#guild-preview-object
type GuildPreview struct {
	// ID is the guild id.
	ID Snowflake `json:"id"`

	// Name is the guild name (2-100 characters).
	Name string `json:"name"`

	// Icon is the icon hash.
	Icon *string `json:"icon"`

	// Splash is the splash hash.
	Splash *string `json:"splash"`

	// DiscoverySplash is the discovery splash hash.
	DiscoverySplash *string `json:"discovery_splash"`

	// Emojis are the custom guild emojis.
	Emojis []Emoji `json:"emojis"`

	// Features are the enabled guild features.
	Features []string `json:"features"`

	// ApproximateMemberCount is the approximate number of members in this guild.
	ApproximateMemberCount *int `json:"approximate_member_count"`

	// ApproximatePresenceCount is the approximate number of online members in this guild.
	ApproximatePresenceCount *int `json:"approximate_presence_count"`

	// Description is the description for the guild.
	Description *string `json:"description"`

	// Stickers are the custom guild stickers.
	Stickers []Sticker `json:"stickers"`
}

// GuildBan represents a guild ban.
//
// Reference: https://discord.com/developers/docs/resources/guild#ban-object
type GuildBan struct {
	// User is the user that was banned.
	User User `json:"user"`

	// Reason is the reason for the ban.
	Reason *string `json:"reason,omitempty"`
}

// AFKTimeout represents valid AFK timeout values in seconds.
type AFKTimeout int

// AFK timeout constants
const (
	AFKTimeout_60   AFKTimeout = 60
	AFKTimeout_300  AFKTimeout = 300
	AFKTimeout_900  AFKTimeout = 900
	AFKTimeout_1800 AFKTimeout = 1800
	AFKTimeout_3600 AFKTimeout = 3600
)

// GuildMessageNotification (GMN) represents default message notification levels.
type GuildMessageNotification int

// Default message notification constants
const (
	GMN_ALL GuildMessageNotification = iota
	GMN_OnlyMentions
)

// GuildExplicitContentFilter (GEC) represents explicit content filter levels.
type GuildExplicitContentFilter int

// Explicit content filter constants
const (
	GEC_FilterDisabled GuildExplicitContentFilter = iota
	GEC_FilterMembersWithoutRoles
	GEC_FilterAllMembers
)

// GuildMFALevel represents MFA levels.
type GuildMFALevel int

// MFA level constants
const (
	GuildMFA_None GuildMFALevel = iota
	GuildMFA_Elevated
)

// GuildNSFWLevel represents NSFW levels.
type GuildNSFWLevel int

// NSFW level constants
const (
	GuildNSFW_Default GuildNSFWLevel = iota
	GuildNSFW_Explicit
	GuildNSFW_Safe
	GuildNSFW_AgeRestricted
)

// GuildVerificationLevel (GVL) represents verification levels.
type GuildVerificationLevel int

// Verification level constants
const (
	GVL_None GuildVerificationLevel = iota
	GVL_Low
	GVL_Medium
	GVL_High
	GVL_VeryHigh
)

// GuildPremiumTier (GP) represents premium tiers.
type GuildPremiumTier int

// Premium tier constants
const (
	GP_TierNone GuildPremiumTier = iota
	GP_Tier1
	GP_Tier2
	GP_Tier3
)

// GuildHubType (GH) represents Student Hub types.
type GuildHubType int

// Hub type constants
const (
	GH_Default GuildHubType = iota
	GH_HighSchool
	GH_College
)

// GuildSystemChannelFlags (GSC) represents system channel flags.
type GuildSystemChannelFlags int

// System channel flag constants
const (
	GSC_SuppressJoinNotifications GuildSystemChannelFlags = 1 << iota
	GSC_SuppressPremiumSubscriptions
	GSC_SuppressGuildReminderNotifications
	GSC_SuppressJoinNotificationReplies
	GSC_SuppressRoleSubscriptionPurchaseNotifications
	GSC_SuppressRoleSubscriptionPurchaseNotificationReplies
)

// GuildFeature (GF) represents guild features.
type GuildFeature string

// Guild feature constants
const (
	GF_AnimatedBanner                        GuildFeature = "ANIMATED_BANNER"
	GF_AnimatedIcon                          GuildFeature = "ANIMATED_ICON"
	GF_ApplicationCommandPermissionsV2       GuildFeature = "APPLICATION_COMMAND_PERMISSIONS_V2"
	GF_AutoModeration                        GuildFeature = "AUTO_MODERATION"
	GF_Banner                                GuildFeature = "BANNER"
	GF_Community                             GuildFeature = "COMMUNITY"
	GF_CreatorMonetizableProvisional         GuildFeature = "CREATOR_MONETIZABLE_PROVISIONAL"
	GF_CreatorStorePage                      GuildFeature = "CREATOR_STORE_PAGE"
	GF_DeveloperSupportServer                GuildFeature = "DEVELOPER_SUPPORT_SERVER"
	GF_Discoverable                          GuildFeature = "DISCOVERABLE"
	GF_Featurable                            GuildFeature = "FEATURABLE"
	GF_HasDirectoryEntry                     GuildFeature = "HAS_DIRECTORY_ENTRY"
	GF_Hub                                   GuildFeature = "HUB"
	GF_InvitesDisabled                       GuildFeature = "INVITES_DISABLED"
	GF_InviteSplash                          GuildFeature = "INVITE_SPLASH"
	GF_LinkedToHub                           GuildFeature = "LINKED_TO_HUB"
	GF_MemberVerificationGateEnabled         GuildFeature = "MEMBER_VERIFICATION_GATE_ENABLED"
	GF_MoreSoundboard                        GuildFeature = "MORE_SOUNDBOARD"
	GF_MonetizationEnabled                   GuildFeature = "MONETIZATION_ENABLED"
	GF_MoreStickers                          GuildFeature = "MORE_STICKERS"
	GF_News                                  GuildFeature = "NEWS"
	GF_Partnered                             GuildFeature = "PARTNERED"
	GF_PreviewEnabled                        GuildFeature = "PREVIEW_ENABLED"
	GF_PrivateThreads                        GuildFeature = "PRIVATE_THREADS"
	GF_RaidAlertsDisabled                    GuildFeature = "RAID_ALERTS_DISABLED"
	GF_RelayEnabled                          GuildFeature = "RELAY_ENABLED"
	GF_RoleIcons                             GuildFeature = "ROLE_ICONS"
	GF_RoleSubscriptionsAvailableForPurchase GuildFeature = "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE"
	GF_RoleSubscriptionsEnabled              GuildFeature = "ROLE_SUBSCRIPTIONS_ENABLED"
	GF_Soundboard                            GuildFeature = "SOUNDBOARD"
	GF_TicketedEventsEnabled                 GuildFeature = "TICKETED_EVENTS_ENABLED"
	GF_VanityURL                             GuildFeature = "VANITY_URL"
	GF_Verified                              GuildFeature = "VERIFIED"
	GF_VIPRegions                            GuildFeature = "VIP_REGIONS"
	GF_WelcomeScreenEnabled                  GuildFeature = "WELCOME_SCREEN_ENABLED"
	GF_GuildTags                             GuildFeature = "GUILD_TAGS"
	GF_EnhancedRoleColors                    GuildFeature = "ENHANCED_ROLE_COLORS"
	GF_GuestsEnabled                         GuildFeature = "GUESTS_ENABLED"
)

// Role represents a Discord role.
// Reference: https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	// ID is the role id.
	ID Snowflake `json:"id"`

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
	Permissions Permissions `json:"permissions"`

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
	BotID *Snowflake `json:"bot_id,omitempty"`

	// IntegrationID is the id of the integration this role belongs to.
	IntegrationID *Snowflake `json:"integration_id,omitempty"`

	// PremiumSubscriberID is the id of this role's subscription sku and level.
	PremiumSubscriberID *Snowflake `json:"premium_subscriber_id,omitempty"`

	// SubscriptionListingID is the id of the subscription listing this role belongs to.
	SubscriptionListingID *Snowflake `json:"subscription_listing_id,omitempty"`

	// AvailableForPurchase is whether this role is available for purchase.
	AvailableForPurchase *bool `json:"available_for_purchase,omitempty"`

	// GuildConnections is whether this role is a guild's linked role.
	GuildConnections *bool `json:"guild_connections,omitempty"`
}

// GuildWelcomeScreen represents a guild welcome screen.
// Reference: https://discord.com/developers/docs/resources/guild#welcome-screen-object
type GuildWelcomeScreen struct {
	// Description is the server description shown in the welcome screen.
	Description *string `json:"description,omitempty"`

	// WelcomeChannels are the channels shown in the welcome screen.
	WelcomeChannels []GuildWelcomeScreenChannel `json:"welcome_channels"`
}

// GuildWelcomeScreenChannel represents a welcome screen channel.
type GuildWelcomeScreenChannel struct {
	// ChannelID is the channel's id.
	ChannelID Snowflake `json:"channel_id"`

	// Description is the description shown for the channel.
	Description string `json:"description"`

	// EmojiID is the emoji id, if the emoji is custom.
	EmojiID *Snowflake `json:"emoji_id,omitempty"`

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

type GuildWidget struct {
	ID            Snowflake        `json:"id"`             // guild id
	Name          string           `json:"name"`           // guild name (2-100 characters)
	InstantInvite *string          `json:"instant_invite"` // instant invite for the guilds specified widget invite channel
	Channels      []PartialChannel `json:"channels"`       // voice and stage channels which are accessible by @everyone
	Members       []PartialUser    `json:"members"`        // special widget user objects that includes users presence (Limit 100)
	PresenceCount int              `json:"presence_count"` // number of online members in this guild
}

type GuildWidgetSettings struct {
	Enabled   bool       `json:"enabled"`    // whether the widget is enabled
	ChannelID *Snowflake `json:"channel_id"` // channel id for the widget, if set
}

// GuildOnboarding represents the onboarding configuration for a guild.
//
// Reference: https://discord.com/developers/docs/resources/guild#guild-onboarding-object
type GuildOnboarding struct {
	GuildID           Snowflake               `json:"guild_id"`            // ID of the guild this onboarding is part of
	Prompts           []GuildOnboardingPrompt `json:"prompts"`             // Prompts shown during onboarding and in customize community
	DefaultChannelIDs []Snowflake             `json:"default_channel_ids"` // Channel IDs that members are opted into automatically
	Enabled           bool                    `json:"enabled"`             // Whether onboarding is enabled in the guild
	Mode              GuildOnboardingMode     `json:"mode"`                // Current mode of onboarding
}

// GuildOnboardingPrompt represents a prompt in the onboarding process.
//
// Reference: https://discord.com/developers/docs/resources/guild#guild-onboarding-prompt-object
type GuildOnboardingPrompt struct {
	ID           Snowflake                     `json:"id"`            // ID of the prompt
	Type         GuildOnboardingPromptType     `json:"type"`          // Type of prompt
	Options      []GuildOnboardingPromptOption `json:"options"`       // Options available within this prompt
	Title        string                        `json:"title"`         // Title of the prompt
	SingleSelect bool                          `json:"single_select"` // Whether users are limited to selecting one option for the prompt
	Required     bool                          `json:"required"`      // Whether the prompt is required before a user completes onboarding
	InOnboarding bool                          `json:"in_onboarding"` // Whether the prompt is present in onboarding flow. If false, only appears in Channels & Roles tab
}

// GuildOnboardingPromptOption represents an option for a prompt in onboarding.
//
// Reference: https://discord.com/developers/docs/resources/guild#guild-onboarding-prompt-option-object
type GuildOnboardingPromptOption struct {
	ID          Snowflake   `json:"id"`          // ID of the prompt option
	ChannelIDs  []Snowflake `json:"channel_ids"` // Channel IDs a user is added to when the option is selected
	RoleIDs     []Snowflake `json:"role_ids"`    // Role IDs assigned to a user when the option is selected
	Emoji       *Emoji      `json:"emoji"`       // Emoji of the option
	Title       string      `json:"title"`       // Title of the option
	Description string      `json:"description"` // Description of the option
}

// GuildOnboardingMode represents the onboarding mode for a guild.
type GuildOnboardingMode int

const (
	GuildOnboardingModeDefault  GuildOnboardingMode = 0 // Default onboarding mode
	GuildOnboardingModeAdvanced GuildOnboardingMode = 1 // Advanced onboarding mode
)

// GuildOnboardingPromptType represents the type of onboarding prompt.
type GuildOnboardingPromptType int

const (
	GuildOnboardingPromptTypeMultipleChoice GuildOnboardingPromptType = 0 // Multiple choice prompt
	GuildOnboardingPromptTypeDropdown       GuildOnboardingPromptType = 1 // Dropdown prompt
)
