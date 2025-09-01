package payloads

import (
	"testing"

	"github.com/kolosys/discord-types/discord"
	"github.com/kolosys/discord-types/utils"
)

func TestPermissionFlags(t *testing.T) {
	// Test that all permission flags are properly defined
	tests := []struct {
		name     string
		flag     int64
		expected int64
	}{
		{"CreateInstantInvite", PermissionFlagsBits.CreateInstantInvite, 1 << 0},
		{"KickMembers", PermissionFlagsBits.KickMembers, 1 << 1},
		{"BanMembers", PermissionFlagsBits.BanMembers, 1 << 2},
		{"Administrator", PermissionFlagsBits.Administrator, 1 << 3},
		{"ManageChannels", PermissionFlagsBits.ManageChannels, 1 << 4},
		{"ManageGuild", PermissionFlagsBits.ManageGuild, 1 << 5},
		{"AddReactions", PermissionFlagsBits.AddReactions, 1 << 6},
		{"ViewAuditLog", PermissionFlagsBits.ViewAuditLog, 1 << 7},
		{"PrioritySpeaker", PermissionFlagsBits.PrioritySpeaker, 1 << 8},
		{"Stream", PermissionFlagsBits.Stream, 1 << 9},
		{"ViewChannel", PermissionFlagsBits.ViewChannel, 1 << 10},
		{"SendMessages", PermissionFlagsBits.SendMessages, 1 << 11},
		{"SendTTSMessages", PermissionFlagsBits.SendTTSMessages, 1 << 12},
		{"ManageMessages", PermissionFlagsBits.ManageMessages, 1 << 13},
		{"EmbedLinks", PermissionFlagsBits.EmbedLinks, 1 << 14},
		{"AttachFiles", PermissionFlagsBits.AttachFiles, 1 << 15},
		{"ReadMessageHistory", PermissionFlagsBits.ReadMessageHistory, 1 << 16},
		{"MentionEveryone", PermissionFlagsBits.MentionEveryone, 1 << 17},
		{"UseExternalEmojis", PermissionFlagsBits.UseExternalEmojis, 1 << 18},
		{"ViewGuildInsights", PermissionFlagsBits.ViewGuildInsights, 1 << 19},
		{"Connect", PermissionFlagsBits.Connect, 1 << 20},
		{"Speak", PermissionFlagsBits.Speak, 1 << 21},
		{"MuteMembers", PermissionFlagsBits.MuteMembers, 1 << 22},
		{"DeafenMembers", PermissionFlagsBits.DeafenMembers, 1 << 23},
		{"MoveMembers", PermissionFlagsBits.MoveMembers, 1 << 24},
		{"UseVAD", PermissionFlagsBits.UseVAD, 1 << 25},
		{"ChangeNickname", PermissionFlagsBits.ChangeNickname, 1 << 26},
		{"ManageNicknames", PermissionFlagsBits.ManageNicknames, 1 << 27},
		{"ManageRoles", PermissionFlagsBits.ManageRoles, 1 << 28},
		{"ManageWebhooks", PermissionFlagsBits.ManageWebhooks, 1 << 29},
		{"ManageGuildExpressions", PermissionFlagsBits.ManageGuildExpressions, 1 << 30},
		{"UseApplicationCommands", PermissionFlagsBits.UseApplicationCommands, 1 << 31},
		{"RequestToSpeak", PermissionFlagsBits.RequestToSpeak, 1 << 32},
		{"ManageEvents", PermissionFlagsBits.ManageEvents, 1 << 33},
		{"ManageThreads", PermissionFlagsBits.ManageThreads, 1 << 34},
		{"CreatePublicThreads", PermissionFlagsBits.CreatePublicThreads, 1 << 35},
		{"CreatePrivateThreads", PermissionFlagsBits.CreatePrivateThreads, 1 << 36},
		{"UseExternalStickers", PermissionFlagsBits.UseExternalStickers, 1 << 37},
		{"SendMessagesInThreads", PermissionFlagsBits.SendMessagesInThreads, 1 << 38},
		{"UseEmbeddedActivities", PermissionFlagsBits.UseEmbeddedActivities, 1 << 39},
		{"ModerateMembers", PermissionFlagsBits.ModerateMembers, 1 << 40},
		{"ViewCreatorMonetizationAnalytics", PermissionFlagsBits.ViewCreatorMonetizationAnalytics, 1 << 41},
		{"UseSoundboard", PermissionFlagsBits.UseSoundboard, 1 << 42},
		{"UseExternalSounds", PermissionFlagsBits.UseExternalSounds, 1 << 45},
		{"SendVoiceMessages", PermissionFlagsBits.SendVoiceMessages, 1 << 46},
		{"SendPolls", PermissionFlagsBits.SendPolls, 1 << 49},
		{"UseExternalApps", PermissionFlagsBits.UseExternalApps, 1 << 50},
		{"PinMessages", PermissionFlagsBits.PinMessages, 1 << 51},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.flag != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.flag, tt.expected)
			}
		})
	}
}

func TestChannelTypes(t *testing.T) {
	// Test that all channel types are properly defined
	tests := []struct {
		name        string
		channelType ChannelType
		expected    int
	}{
		{"GuildText", ChannelTypeGuildText, 0},
		{"DM", ChannelTypeDM, 1},
		{"GuildVoice", ChannelTypeGuildVoice, 2},
		{"GroupDM", ChannelTypeGroupDM, 3},
		{"GuildCategory", ChannelTypeGuildCategory, 4},
		{"GuildAnnouncement", ChannelTypeGuildAnnouncement, 5},
		{"AnnouncementThread", ChannelTypeAnnouncementThread, 10},
		{"PublicThread", ChannelTypePublicThread, 10},
		{"PrivateThread", ChannelTypePrivateThread, 10},
		{"GuildStageVoice", ChannelTypeGuildStageVoice, 10},
		{"GuildDirectory", ChannelTypeGuildDirectory, 10},
		{"GuildForum", ChannelTypeGuildForum, 10},
		{"GuildMedia", ChannelTypeGuildMedia, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.channelType) != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, int(tt.channelType), tt.expected)
			}
		})
	}
}

func TestMessageFlags(t *testing.T) {
	// Test that all message flags are properly defined
	tests := []struct {
		name     string
		flag     MessageFlags
		expected MessageFlags
	}{
		{"Crossposted", MessageFlagCrossposted, 1 << 0},
		{"IsCrosspost", MessageFlagIsCrosspost, 1 << 1},
		{"SuppressEmbeds", MessageFlagSuppressEmbeds, 1 << 2},
		{"SourceMessageDeleted", MessageFlagSourceMessageDeleted, 1 << 3},
		{"Urgent", MessageFlagUrgent, 1 << 4},
		{"HasThread", MessageFlagHasThread, 1 << 5},
		{"Ephemeral", MessageFlagEphemeral, 1 << 6},
		{"Loading", MessageFlagLoading, 1 << 7},
		{"FailedToMentionSomeRolesInThread", MessageFlagFailedToMentionSomeRolesInThread, 1 << 8},
		{"SuppressNotifications", MessageFlagSuppressNotifications, 1 << 12},
		{"IsVoiceMessage", MessageFlagIsVoiceMessage, 1 << 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.flag != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.flag, tt.expected)
			}
		})
	}
}

func TestUserFlags(t *testing.T) {
	// Test that all user flags are properly defined
	tests := []struct {
		name     string
		flag     UserFlags
		expected UserFlags
	}{
		{"Staff", UserFlagStaff, 1 << 0},
		{"Partner", UserFlagPartner, 1 << 1},
		{"Hypesquad", UserFlagHypesquad, 1 << 2},
		{"BugHunterLevel1", UserFlagBugHunterLevel1, 1 << 3},
		{"MFASMS", UserFlagMFASMS, 1 << 4},
		{"PremiumPromoDismissed", UserFlagPremiumPromoDismissed, 1 << 5},
		{"HypeSquadOnlineHouse1", UserFlagHypeSquadOnlineHouse1, 1 << 6},
		{"HypeSquadOnlineHouse2", UserFlagHypeSquadOnlineHouse2, 1 << 7},
		{"HypeSquadOnlineHouse3", UserFlagHypeSquadOnlineHouse3, 1 << 8},
		{"PremiumEarlySupporter", UserFlagPremiumEarlySupporter, 1 << 9},
		{"TeamPseudoUser", UserFlagTeamPseudoUser, 1 << 10},
		{"HasUnreadUrgentMessages", UserFlagHasUnreadUrgentMessages, 1 << 13},
		{"BugHunterLevel2", UserFlagBugHunterLevel2, 1 << 14},
		{"VerifiedBot", UserFlagVerifiedBot, 1 << 16},
		{"VerifiedDeveloper", UserFlagVerifiedDeveloper, 1 << 17},
		{"CertifiedModerator", UserFlagCertifiedModerator, 1 << 18},
		{"BotHTTPInteractions", UserFlagBotHTTPInteractions, 1 << 19},
		{"Spammer", UserFlagSpammer, 1 << 20},
		{"DisablePremium", UserFlagDisablePremium, 1 << 21},
		{"ActiveDeveloper", UserFlagActiveDeveloper, 1 << 22},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.flag != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.flag, tt.expected)
			}
		})
	}
}

func TestApplicationFlags(t *testing.T) {
	// Test that all application flags are properly defined
	tests := []struct {
		name     string
		flag     ApplicationFlags
		expected ApplicationFlags
	}{
		{"GatewayPresence", ApplicationFlagGatewayPresence, 1 << 12},
		{"GatewayPresenceLimited", ApplicationFlagGatewayPresenceLimited, 1 << 13},
		{"GatewayGuildMembers", ApplicationFlagGatewayGuildMembers, 1 << 14},
		{"GatewayGuildMembersLimited", ApplicationFlagGatewayGuildMembersLimited, 1 << 15},
		{"VerificationPendingGuildLimit", ApplicationFlagVerificationPendingGuildLimit, 1 << 16},
		{"Embedded", ApplicationFlagEmbedded, 1 << 17},
		{"GatewayMessageContent", ApplicationFlagGatewayMessageContent, 1 << 18},
		{"GatewayMessageContentLimited", ApplicationFlagGatewayMessageContentLimited, 1 << 19},
		{"ApplicationCommandBadge", ApplicationFlagApplicationCommandBadge, 1 << 23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.flag != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.flag, tt.expected)
			}
		})
	}
}

func TestGuildSystemChannelFlags(t *testing.T) {
	// Test that all guild system channel flags are properly defined
	tests := []struct {
		name     string
		flag     GuildSystemChannelFlags
		expected GuildSystemChannelFlags
	}{
		{"SuppressJoinNotifications", GuildSystemChannelFlagsSuppressJoinNotifications, 1 << 0},
		{"SuppressPremiumSubscriptions", GuildSystemChannelFlagsSuppressPremiumSubscriptions, 1 << 1},
		{"SuppressGuildReminderNotifications", GuildSystemChannelFlagsSuppressGuildReminderNotifications, 1 << 2},
		{"SuppressJoinNotificationReplies", GuildSystemChannelFlagsSuppressJoinNotificationReplies, 1 << 3},
		{"SuppressRoleSubscriptionPurchaseNotifications", GuildSystemChannelFlagsSuppressRoleSubscriptionPurchaseNotifications, 1 << 4},
		{"SuppressRoleSubscriptionPurchaseNotificationReplies", GuildSystemChannelFlagsSuppressRoleSubscriptionPurchaseNotificationReplies, 1 << 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.flag != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.flag, tt.expected)
			}
		})
	}
}

func TestGuildVerificationLevel(t *testing.T) {
	// Test that all guild verification levels are properly defined
	tests := []struct {
		name     string
		level    GuildVerificationLevel
		expected int
	}{
		{"None", GuildVerificationLevelNone, 0},
		{"Low", GuildVerificationLevelLow, 1},
		{"Medium", GuildVerificationLevelMedium, 2},
		{"High", GuildVerificationLevelHigh, 3},
		{"VeryHigh", GuildVerificationLevelVeryHigh, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.level) != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, int(tt.level), tt.expected)
			}
		})
	}
}

func TestGuildPremiumTier(t *testing.T) {
	// Test that all guild premium tiers are properly defined
	tests := []struct {
		name     string
		tier     GuildPremiumTier
		expected int
	}{
		{"None", GuildPremiumTierNone, 0},
		{"Tier1", GuildPremiumTierTier1, 1},
		{"Tier2", GuildPremiumTierTier2, 2},
		{"Tier3", GuildPremiumTierTier3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.tier) != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, int(tt.tier), tt.expected)
			}
		})
	}
}

func TestGuildNSFWLevel(t *testing.T) {
	// Test that all guild NSFW levels are properly defined
	tests := []struct {
		name     string
		level    GuildNSFWLevel
		expected int
	}{
		{"Default", GuildNSFWLevelDefault, 0},
		{"Explicit", GuildNSFWLevelExplicit, 1},
		{"Safe", GuildNSFWLevelSafe, 2},
		{"AgeRestricted", GuildNSFWLevelAgeRestricted, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.level) != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, int(tt.level), tt.expected)
			}
		})
	}
}

func TestGuildHubType(t *testing.T) {
	// Test that all guild hub types are properly defined
	tests := []struct {
		name     string
		hubType  GuildHubType
		expected int
	}{
		{"Default", GuildHubTypeDefault, 0},
		{"HighSchool", GuildHubTypeHighSchool, 1},
		{"College", GuildHubTypeCollege, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int(tt.hubType) != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, int(tt.hubType), tt.expected)
			}
		})
	}
}

func TestRoleTags(t *testing.T) {
	// Test RoleTags struct
	botID := discord.NewSnowflake("123456789012345678")
	integrationID := discord.NewSnowflake("987654321098765432")
	availableForPurchase := true
	guildConnections := true

	roleTags := RoleTags{
		BotID:                &botID,
		IntegrationID:        &integrationID,
		AvailableForPurchase: &availableForPurchase,
		GuildConnections:     &guildConnections,
	}

	// Test that fields are properly set
	if roleTags.BotID == nil || *roleTags.BotID != "123456789012345678" {
		t.Errorf("BotID = %v, want %s", roleTags.BotID, "123456789012345678")
	}

	if roleTags.IntegrationID == nil || *roleTags.IntegrationID != "987654321098765432" {
		t.Errorf("IntegrationID = %v, want %s", roleTags.IntegrationID, "987654321098765432")
	}

	if roleTags.AvailableForPurchase == nil || !*roleTags.AvailableForPurchase {
		t.Error("AvailableForPurchase should be true")
	}

	if roleTags.GuildConnections == nil || !*roleTags.GuildConnections {
		t.Error("GuildConnections should be true")
	}
}

func TestRole(t *testing.T) {
	// Test Role struct
	role := Role{
		ID:           discord.NewSnowflake("123456789012345678"),
		Name:         "Test Role",
		Color:        16711680, // Red
		Hoist:        true,
		Icon:         utils.StringPtr("icon_hash"),
		UnicodeEmoji: utils.StringPtr("😀"),
		Position:     1,
		Permissions:  discord.NewPermissions("123456789"),
		Managed:      false,
		Mentionable:  true,
		Tags:         &RoleTags{},
	}

	// Test that fields are properly set
	if role.ID != "123456789012345678" {
		t.Errorf("ID = %s, want %s", role.ID, "123456789012345678")
	}

	if role.Name != "Test Role" {
		t.Errorf("Name = %s, want %s", role.Name, "Test Role")
	}

	if role.Color != 16711680 {
		t.Errorf("Color = %d, want %d", role.Color, 16711680)
	}

	if !role.Hoist {
		t.Error("Hoist should be true")
	}

	if role.Icon == nil || *role.Icon != "icon_hash" {
		t.Error("Icon should be 'icon_hash'")
	}

	if role.UnicodeEmoji == nil || *role.UnicodeEmoji != "😀" {
		t.Error("UnicodeEmoji should be '😀'")
	}

	if role.Position != 1 {
		t.Errorf("Position = %d, want %d", role.Position, 1)
	}

	if role.Permissions != "123456789" {
		t.Errorf("Permissions = %s, want %s", role.Permissions, "123456789")
	}

	if role.Managed {
		t.Error("Managed should be false")
	}

	if !role.Mentionable {
		t.Error("Mentionable should be true")
	}

	if role.Tags == nil {
		t.Error("Tags should not be nil")
	}
}
