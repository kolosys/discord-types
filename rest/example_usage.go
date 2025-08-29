// Package rest provides Discord REST API types and utilities.
//
// This file demonstrates usage examples for the REST API enhancements.
// This is for documentation purposes only.
package rest

import (
	"fmt"

	"github.com/kolosys/discord-types/discord"
)

// ExampleRESTUsage demonstrates how to use the REST API enhancements.
func ExampleRESTUsage() {
	// Example Snowflake IDs
	guildID := discord.Snowflake("123456789012345678")
	channelID := discord.Snowflake("987654321098765432")
	userID := discord.Snowflake("111222333444555666")
	messageID := discord.Snowflake("777888999000111222")

	// === API Route Examples ===

	// Basic routes
	fmt.Println("=== API Routes ===")
	fmt.Println("Guild:", Routes.Guild(guildID))
	fmt.Println("Channel:", Routes.Channel(channelID))
	fmt.Println("User:", Routes.User("@me"))

	// Channel routes with advanced features
	fmt.Println("\n=== Channel Routes ===")
	fmt.Println("Messages:", Routes.ChannelMessages(channelID))
	fmt.Println("Bulk Delete:", Routes.ChannelBulkDelete(channelID))
	fmt.Println("Typing:", Routes.ChannelTyping(channelID))
	fmt.Println("Pins:", Routes.ChannelMessagesPins(channelID))
	fmt.Println("Invites:", Routes.ChannelInvites(channelID))
	fmt.Println("Webhooks:", Routes.ChannelWebhooks(channelID))

	// Guild routes with comprehensive features
	fmt.Println("\n=== Guild Routes ===")
	fmt.Println("Members:", Routes.GuildMembers(guildID))
	fmt.Println("Member Search:", Routes.GuildMembersSearch(guildID))
	fmt.Println("Roles:", Routes.GuildRoles(guildID))
	fmt.Println("Bans:", Routes.GuildBans(guildID))
	fmt.Println("Bulk Ban:", Routes.GuildBulkBan(guildID))
	fmt.Println("Emojis:", Routes.GuildEmojis(guildID))
	fmt.Println("Stickers:", Routes.GuildStickers(guildID))
	fmt.Println("Templates:", Routes.GuildTemplates(guildID))
	fmt.Println("Scheduled Events:", Routes.GuildScheduledEvents(guildID))
	fmt.Println("Soundboard Sounds:", Routes.GuildSoundboardSounds(guildID))

	// Thread routes
	fmt.Println("\n=== Thread Routes ===")
	fmt.Println("Create Thread:", Routes.Threads(channelID, nil))
	fmt.Println("Create from Message:", Routes.Threads(channelID, &messageID))
	fmt.Println("Active Threads:", Routes.GuildActiveThreads(guildID))
	fmt.Println("Archived Threads:", Routes.ChannelThreads(channelID, "public"))
	fmt.Println("Thread Members:", Routes.ThreadMembers(channelID, nil))

	// Poll routes
	fmt.Println("\n=== Poll Routes ===")
	fmt.Println("Poll Voters:", Routes.PollAnswerVoters(channelID, messageID, 1))
	fmt.Println("Expire Poll:", Routes.ExpirePoll(channelID, messageID))

	// OAuth2 routes
	fmt.Println("\n=== OAuth2 Routes ===")
	fmt.Println("Authorization:", Routes.OAuth2Authorization())
	fmt.Println("Token Exchange:", Routes.OAuth2TokenExchange())
	fmt.Println("Token Revocation:", Routes.OAuth2TokenRevocation())

	// === CDN Route Examples ===
	fmt.Println("\n=== CDN Routes ===")

	// User assets
	avatarHash := "a1b2c3d4e5f6"
	fmt.Println("User Avatar:", CDNRoutes.UserAvatar(userID, avatarHash, ImageFormatPNG))
	fmt.Println("User Banner:", CDNRoutes.UserBanner(userID, avatarHash, ImageFormatWebP))
	fmt.Println("Default Avatar:", CDNRoutes.DefaultUserAvatar(DefaultUserAvatar0))

	// Guild assets
	iconHash := "guild_icon_hash"
	fmt.Println("Guild Icon:", CDNRoutes.GuildIcon(guildID, iconHash, ImageFormatPNG))
	fmt.Println("Guild Banner:", CDNRoutes.GuildBanner(guildID, iconHash, ImageFormatGIF))
	fmt.Println("Guild Splash:", CDNRoutes.GuildSplash(guildID, iconHash, ImageFormatJPEG))

	// Application assets
	appID := discord.Snowflake("123456789")
	fmt.Println("App Icon:", CDNRoutes.ApplicationIcon(appID, iconHash, ImageFormatPNG))
	fmt.Println("App Asset:", CDNRoutes.ApplicationAsset(appID, "asset123", ImageFormatWebP))

	// Emoji and stickers
	emojiID := discord.Snowflake("emoji123")
	stickerID := discord.Snowflake("sticker456")
	fmt.Println("Emoji:", CDNRoutes.Emoji(emojiID, ImageFormatPNG))
	fmt.Println("Sticker:", CDNRoutes.Sticker(stickerID, ImageFormatLottie))

	// === Query Parameter Examples ===
	fmt.Println("\n=== Query Parameters ===")

	// Message query with pagination
	messageQuery := GetMessagesQuery{
		Limit:  NewLimit(50),
		Before: &messageID,
	}
	messagesURL := BuildAPIURL(Routes.ChannelMessages(channelID), messageQuery)
	fmt.Println("Messages with query:", messagesURL)

	// Guild members with search
	memberSearch := SearchGuildMembersQuery{
		Query: "username",
		Limit: NewLimit(10),
	}
	searchURL := BuildAPIURL(Routes.GuildMembersSearch(guildID), memberSearch)
	fmt.Println("Member search:", searchURL)

	// CDN with size parameter
	cdnQuery := CDNQuery{Size: Size512}
	avatarURL := BuildCDNURL(CDNRoutes.UserAvatar(userID, avatarHash, ImageFormatPNG), cdnQuery)
	fmt.Println("Avatar with size:", avatarURL)

	// === Full URL Examples ===
	fmt.Println("\n=== Complete URLs ===")
	fmt.Println("API Base:", RouteBases.API)
	fmt.Println("CDN Base:", RouteBases.CDN)
	fmt.Println("OAuth2 Auth URL:", OAuth2Routes.AuthorizationURL)
	fmt.Println("OAuth2 Token URL:", OAuth2Routes.TokenURL)

	// === Helper Functions ===
	fmt.Println("\n=== Helper Functions ===")

	// Default avatar calculation
	discriminator := 1234
	defaultIndex := GetDefaultAvatarIndex(userID, &discriminator)
	fmt.Printf("Default avatar index (legacy): %d\n", defaultIndex)

	// New username system (no discriminator)
	newDefaultIndex := GetDefaultAvatarIndex(userID, nil)
	fmt.Printf("Default avatar index (new): %d\n", newDefaultIndex)

	// Image size helpers
	fmt.Println("Common sizes available:", *Size16, *Size64, *Size256, *Size1024)
}

// ExampleAdvancedUsage shows more complex scenarios.
func ExampleAdvancedUsage() {
	guildID := discord.Snowflake("123456789012345678")
	channelID := discord.Snowflake("987654321098765432")

	fmt.Println("=== Advanced Usage Examples ===")

	// Building complex audit log queries
	auditQuery := GetAuditLogQuery{
		UserID:     NewSnowflake(discord.Snowflake("admin_user_id")),
		ActionType: NewLimit(1), // GUILD_UPDATE
		Limit:      NewLimit(25),
	}
	auditURL := BuildAPIURL(Routes.GuildAuditLog(guildID), auditQuery)
	fmt.Println("Audit log query:", auditURL)

	// Thread management with comprehensive queries
	threadQuery := GetThreadsQuery{
		Before: NewString("2024-01-01T00:00:00Z"),
		Limit:  NewLimit(100),
	}
	threadsURL := BuildAPIURL(Routes.ChannelThreads(channelID, "private"), threadQuery)
	fmt.Println("Private threads:", threadsURL)

	// Entitlements with multiple filters
	entitlementQuery := GetEntitlementsQuery{
		UserID:       NewSnowflake(discord.Snowflake("user123")),
		SKUIDS:       []discord.Snowflake{"sku1", "sku2", "sku3"},
		ExcludeEnded: NewBool(true),
		Limit:        NewLimit(50),
	}
	entitlementsURL := BuildAPIURL(Routes.Entitlements(discord.Snowflake("app123")), entitlementQuery)
	fmt.Println("Entitlements query:", entitlementsURL)

	// CDN URLs with different formats and sizes for responsive images
	userID := discord.Snowflake("user123")
	avatarHash := "abc123"

	sizes := []*ImageSize{Size64, Size128, Size256, Size512}
	formats := []ImageFormat{ImageFormatPNG, ImageFormatWebP, ImageFormatJPEG}

	fmt.Println("\n=== Responsive Image Examples ===")
	for _, size := range sizes {
		for _, format := range formats {
			cdnQuery := CDNQuery{Size: size}
			url := BuildCDNURL(CDNRoutes.UserAvatar(userID, avatarHash, format), cdnQuery)
			fmt.Printf("Avatar %dx%d %s: %s\n", *size, *size, format, url)
		}
	}
}

// ExampleErrorHandling shows proper error handling patterns.
func ExampleErrorHandling() {
	fmt.Println("=== Error Handling Examples ===")

	// Invalid snowflake handling
	invalidID := discord.Snowflake("not_a_number")
	defaultIndex := GetDefaultAvatarIndex(invalidID, nil)
	fmt.Printf("Invalid ID fallback: %d\n", defaultIndex)

	// Validate snowflakes before use
	userID := discord.Snowflake("123456789012345678")
	if !userID.IsValid() {
		fmt.Println("Invalid user ID")
		return
	}

	// Safe time extraction
	if timestamp, err := userID.Time(); err == nil {
		fmt.Printf("User created at: %s\n", timestamp.Format("2006-01-02 15:04:05"))
	} else {
		fmt.Printf("Could not extract timestamp: %v\n", err)
	}
}
