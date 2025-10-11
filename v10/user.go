package v10

// User represents a Discord user.
// Reference: https://discord.com/developers/docs/resources/user#user-object
type User struct {
	// ID is the user's id.
	ID Snowflake `json:"id"`

	// Username is the user's username, not unique across the platform.
	Username string `json:"username"`

	// Discriminator is the user's Discord-tag.
	Discriminator string `json:"discriminator"`

	// GlobalName is the user's display name, if it is set.
	// For bots, this is the application name.
	GlobalName *string `json:"global_name"`

	// Avatar is the user's avatar hash.
	// Reference: https://discord.com/developers/docs/reference#image-formatting
	Avatar *string `json:"avatar"`

	// Bot indicates whether the user belongs to an OAuth2 application.
	Bot *bool `json:"bot,omitempty"`

	// System indicates whether the user is an Official Discord System user
	// (part of the urgent message system).
	System *bool `json:"system,omitempty"`

	// MFAEnabled indicates whether the user has two factor enabled on their account.
	MFAEnabled *bool `json:"mfa_enabled,omitempty"`

	// Banner is the user's banner hash.
	// Reference: https://discord.com/developers/docs/reference#image-formatting
	Banner *string `json:"banner,omitempty"`

	// AccentColor is the user's banner color encoded as an integer representation
	// of hexadecimal color code.
	AccentColor *int `json:"accent_color,omitempty"`

	// Locale is the user's chosen language option.
	Locale *string `json:"locale,omitempty"`

	// Verified indicates whether the email on this account has been verified.
	Verified *bool `json:"verified,omitempty"`

	// Email is the user's email.
	Email *string `json:"email,omitempty"`

	// Flags are the flags on a user's account.
	// Reference: https://discord.com/developers/docs/resources/user#user-object-user-flags
	Flags *UserFlags `json:"flags,omitempty"`

	// PremiumType is the type of Nitro subscription on a user's account.
	// Reference: https://discord.com/developers/docs/resources/user#user-object-premium-types
	PremiumType *UserPremiumType `json:"premium_type,omitempty"`

	// PublicFlags are the public flags on a user's account.
	// Reference: https://discord.com/developers/docs/resources/user#user-object-user-flags
	PublicFlags *UserFlags `json:"public_flags,omitempty"`

	// AvatarDecoration is the user's avatar decoration hash.
	//
	// Reference: https://discord.com/developers/docs/reference#image-formatting
	// Deprecated: Use AvatarDecorationData instead.
	AvatarDecoration *string `json:"avatar_decoration,omitempty"`

	// AvatarDecorationData is the data for the user's avatar decoration.
	// Reference: https://discord.com/developers/docs/resources/user#avatar-decoration-data-object
	AvatarDecorationData *AvatarDecorationData `json:"avatar_decoration_data,omitempty"`

	// Collectibles is the data for the user's collectibles.
	// Reference: https://discord.com/developers/docs/resources/user#collectibles
	Collectibles *Collectibles `json:"collectibles,omitempty"`

	// PrimaryGuild is the user's primary guild.
	// Reference: https://discord.com/developers/docs/resources/user#user-object-user-primary-guild
	PrimaryGuild *UserPrimaryGuild `json:"primary_guild,omitempty"`
}

// UserFlags represents Discord user flags.
// Reference: https://discord.com/developers/docs/resources/user#user-object-user-flags
type UserFlags int

// User flag constants
const (
	UserFlagStaff UserFlags = 1 << iota
	UserFlagPartner
	UserFlagHypesquad
	UserFlagBugHunterLevel1
	_
	_
	UserFlagHypeSquadOnlineHouse1
	UserFlagHypeSquadOnlineHouse2
	UserFlagHypeSquadOnlineHouse3
	UserFlagPremiumEarlySupporter
	UserFlagTeamPseudoUser
	_
	_
	UserFlagHasUnreadUrgentMessages
	UserFlagBugHunterLevel2
	_
	UserFlagVerifiedBot
	UserFlagVerifiedDeveloper
	UserFlagCertifiedModerator
	UserFlagBotHTTPInteractions
	_
	_
	UserFlagActiveDeveloper
)

// Large user flags that require special handling
const (
	// UserFlagQuarantined indicates the user's account has been quarantined based on recent activity.
	// This value would be 1 << 44, but bit shifting above 1 << 30 requires special handling.
	UserFlagQuarantined UserFlags = 17_592_186_044_416

	// UserFlagCollaborator is an unstable user flag.
	// This value would be 1 << 50, but bit shifting above 1 << 30 requires special handling.
	UserFlagCollaborator UserFlags = 1_125_899_906_842_624

	// UserFlagRestrictedCollaborator is an unstable user flag.
	// This value would be 1 << 51, but bit shifting above 1 << 30 requires special handling.
	UserFlagRestrictedCollaborator UserFlags = 2_251_799_813_685_248
)

// PartialUser represents a partial user object with minimal fields.
type PartialUser struct {
	// ID is the user's id.
	ID Snowflake `json:"id"`

	// Username is the user's username.
	Username string `json:"username"`

	// Discriminator is the user's Discord-tag.
	Discriminator string `json:"discriminator"`

	// Avatar is the user's avatar hash.
	Avatar *string `json:"avatar"`
}

// UserPremiumType represents the type of Nitro subscription.
// Reference: https://discord.com/developers/docs/resources/user#user-object-premium-types
type UserPremiumType int

// Premium type constants
const (
	UserPremiumTypeNone UserPremiumType = iota
	UserPremiumTypeNitroClassic
	UserPremiumTypeNitro
	UserPremiumTypeNitroBasic
)

// String implements fmt.Stringer for debugging
func (t UserPremiumType) String() string {
	switch t {
	case UserPremiumTypeNone:
		return "None"
	case UserPremiumTypeNitroClassic:
		return "Nitro Classic"
	case UserPremiumTypeNitro:
		return "Nitro"
	case UserPremiumTypeNitroBasic:
		return "Nitro Basic"
	default:
		return "Unknown"
	}
}

// Connection represents a user's connected account.
// Reference: https://discord.com/developers/docs/resources/user#connection-object
type Connection struct {
	// ID is the id of the connection account.
	ID string `json:"id"`

	// Name is the username of the connection account.
	Name string `json:"name"`

	// Type is the service of the connection.
	Type ConnectionService `json:"type"`

	// Revoked indicates whether the connection is revoked.
	Revoked *bool `json:"revoked,omitempty"`

	// Integrations is an array of partial server integrations.
	Integrations []PartialGuildIntegration `json:"integrations,omitempty"`

	// Verified indicates whether the connection is verified.
	Verified bool `json:"verified"`

	// FriendSync indicates whether friend sync is enabled for this connection.
	FriendSync bool `json:"friend_sync"`

	// ShowActivity indicates whether activities related to this connection
	// will be shown in presence updates.
	ShowActivity bool `json:"show_activity"`

	// TwoWayLink indicates whether this connection supports console voice transfer.
	TwoWayLink bool `json:"two_way_link"`

	// Visibility is the visibility of this connection.
	Visibility ConnectionVisibility `json:"visibility"`
}

// ConnectionService represents the type of connected service.
type ConnectionService string

// Connection service constants
const (
	ConnectionServiceAmazonMusic        ConnectionService = "amazon-music"
	ConnectionServiceBattleNet          ConnectionService = "battlenet"
	ConnectionServiceBluesky            ConnectionService = "bluesky"
	ConnectionServiceBungieNet          ConnectionService = "bungie"
	ConnectionServiceCrunchyroll        ConnectionService = "crunchyroll"
	ConnectionServiceDomain             ConnectionService = "domain"
	ConnectionServiceEBay               ConnectionService = "ebay"
	ConnectionServiceEpicGames          ConnectionService = "epicgames"
	ConnectionServiceFacebook           ConnectionService = "facebook"
	ConnectionServiceGitHub             ConnectionService = "github"
	ConnectionServiceInstagram          ConnectionService = "instagram"
	ConnectionServiceLeagueOfLegends    ConnectionService = "leagueoflegends"
	ConnectionServiceMastodon           ConnectionService = "mastodon"
	ConnectionServicePayPal             ConnectionService = "paypal"
	ConnectionServicePlayStationNetwork ConnectionService = "playstation"
	ConnectionServiceReddit             ConnectionService = "reddit"
	ConnectionServiceRiotGames          ConnectionService = "riotgames"
	ConnectionServiceRoblox             ConnectionService = "roblox"
	ConnectionServiceSpotify            ConnectionService = "spotify"
	ConnectionServiceSkype              ConnectionService = "skype"
	ConnectionServiceSteam              ConnectionService = "steam"
	ConnectionServiceTikTok             ConnectionService = "tiktok"
	ConnectionServiceTwitch             ConnectionService = "twitch"
	ConnectionServiceX                  ConnectionService = "twitter"
	ConnectionServiceTwitter            ConnectionService = ConnectionServiceX // Deprecated: Use ConnectionServiceX
	ConnectionServiceXbox               ConnectionService = "xbox"
	ConnectionServiceYouTube            ConnectionService = "youtube"
)

// ConnectionVisibility represents the visibility of a connection.
type ConnectionVisibility int

// Connection visibility constants
const (
	// ConnectionVisibilityNone makes the connection invisible to everyone except the user themselves.
	ConnectionVisibilityNone ConnectionVisibility = iota

	// ConnectionVisibilityEveryone makes the connection visible to everyone.
	ConnectionVisibilityEveryone
)

// ApplicationRoleConnection represents an application role connection.
// Reference: https://discord.com/developers/docs/resources/user#application-role-connection-object-application-role-connection-structure
type ApplicationRoleConnection struct {
	// PlatformName is the vanity name of the platform a bot has connected (max 50 characters).
	PlatformName *string `json:"platform_name"`

	// PlatformUsername is the username on the platform a bot has connected (max 100 characters).
	PlatformUsername *string `json:"platform_username"`

	// Metadata is an object mapping application role connection metadata keys
	// to their string-ified value (max 100 characters) for the user on the
	// platform a bot has connected.
	Metadata map[string]any `json:"metadata"`
}

// AvatarDecorationData represents avatar decoration data.
// Reference: https://discord.com/developers/docs/resources/user#avatar-decoration-data-object
type AvatarDecorationData struct {
	// Asset is the avatar decoration hash.
	//
	// Reference: https://discord.com/developers/docs/reference#image-formatting
	Asset string `json:"asset"`

	// SkuID is the id of the avatar decoration's SKU.
	SkuID Snowflake `json:"sku_id"`
}

// Collectibles represents the collectibles the user has, excluding Avatar Decorations and Profile Effects.
// Reference: https://discord.com/developers/docs/resources/user#collectibles
type Collectibles struct {
	// Nameplate is object mapping of NameplateData.
	Nameplate *NameplateData `json:"nameplate,omitempty"`
}

// NameplateData represents nameplate data.
// Reference: https://discord.com/developers/docs/resources/user#nameplate
type NameplateData struct {
	// SkuID is the ID of the nameplate SKU.
	SkuID Snowflake `json:"sku_id"`

	// Asset is the path to the nameplate asset.
	Asset string `json:"asset"`

	// Label is the label of this nameplate. Currently unused.
	Label string `json:"label"`

	// Palette is the background color of the nameplate.
	Palette NameplatePalette `json:"palette"`
}

// NameplatePalette represents the background color of a nameplate.
type NameplatePalette string

// Nameplate palette constants
const (
	NameplatePaletteBerry     NameplatePalette = "berry"
	NameplatePaletteBubbleGum NameplatePalette = "bubble_gum"
	NameplatePaletteClover    NameplatePalette = "clover"
	NameplatePaletteCobalt    NameplatePalette = "cobalt"
	NameplatePaletteCrimson   NameplatePalette = "crimson"
	NameplatePaletteForest    NameplatePalette = "forest"
	NameplatePaletteLemon     NameplatePalette = "lemon"
	NameplatePaletteSky       NameplatePalette = "sky"
	NameplatePaletteTeal      NameplatePalette = "teal"
	NameplatePaletteViolet    NameplatePalette = "violet"
	NameplatePaletteWhite     NameplatePalette = "white"
)

// UserPrimaryGuild represents a user's primary guild.
// Reference: https://discord.com/developers/docs/resources/user#user-object-user-primary-guild
type UserPrimaryGuild struct {
	// IdentityGuildID is the id of the user's primary guild.
	IdentityGuildID *Snowflake `json:"identity_guild_id"`

	// IdentityEnabled indicates whether the user is displaying the primary guild's server tag.
	// This can be null if the system clears the identity, e.g. because the server no longer supports tags.
	IdentityEnabled *bool `json:"identity_enabled"`

	// Tag is the text of the user's server tag. Limited to 4 characters.
	Tag *string `json:"tag"`

	// Badge is the server tag badge hash.
	//
	// Reference: https://discord.com/developers/docs/reference#image-formatting
	Badge *string `json:"badge"`
}

// PartialGuildIntegration represents a partial guild integration for connections.
// Reference: https://discord.com/developers/docs/resources/user#connection-object
type PartialGuildIntegration struct {
	// ID is the integration id.
	ID Snowflake `json:"id"`

	// Name is the integration name.
	Name string `json:"name"`

	// Type is the integration type.
	Type string `json:"type"`
}
