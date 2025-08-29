// Package rest provides Discord REST API types and utilities.
//
// This package contains route builders, request/response types,
// and other utilities for interacting with the Discord REST API.
package rest

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/kolosys/discord-types/discord"
)

// APIVersion represents the Discord API version.
const APIVersion = "10"

// urlSafeCharacters is a regex pattern for characters that are safe in URLs.
var urlSafeCharacters = regexp.MustCompile(`^[a-zA-Z0-9._~:/?#[\]@!$&'()*+,;=-]+$`)

// RouteBuilder provides methods to build Discord API routes.
type RouteBuilder struct{}

// Routes is the global route builder instance.
var Routes = RouteBuilder{}

// ApplicationRoleConnectionMetadata returns the route for application role connection metadata.
//
// Route for:
// - GET `/applications/{application.id}/role-connections/metadata`
// - PUT `/applications/{application.id}/role-connections/metadata`
func (r RouteBuilder) ApplicationRoleConnectionMetadata(applicationID discord.Snowflake) string {
	return fmt.Sprintf("/applications/%s/role-connections/metadata", escapeRouteParam(applicationID.String()))
}

// GuildAutoModerationRules returns the route for guild auto moderation rules.
//
// Route for:
// - GET  `/guilds/{guild.id}/auto-moderation/rules`
// - POST `/guilds/{guild.id}/auto-moderation/rules`
func (r RouteBuilder) GuildAutoModerationRules(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/auto-moderation/rules", escapeRouteParam(guildID.String()))
}

// GuildAutoModerationRule returns the route for a specific guild auto moderation rule.
//
// Routes for:
// - GET    `/guilds/{guild.id}/auto-moderation/rules/{rule.id}`
// - PATCH  `/guilds/{guild.id}/auto-moderation/rules/{rule.id}`
// - DELETE `/guilds/{guild.id}/auto-moderation/rules/{rule.id}`
func (r RouteBuilder) GuildAutoModerationRule(guildID, ruleID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/auto-moderation/rules/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(ruleID.String()))
}

// GuildAuditLog returns the route for guild audit log.
//
// Route for:
// - GET `/guilds/{guild.id}/audit-logs`
func (r RouteBuilder) GuildAuditLog(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/audit-logs", escapeRouteParam(guildID.String()))
}

// Channel returns the route for a channel.
//
// Route for:
// - GET    `/channels/{channel.id}`
// - PATCH  `/channels/{channel.id}`
// - DELETE `/channels/{channel.id}`
func (r RouteBuilder) Channel(channelID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s", escapeRouteParam(channelID.String()))
}

// ChannelMessages returns the route for channel messages.
//
// Route for:
// - GET  `/channels/{channel.id}/messages`
// - POST `/channels/{channel.id}/messages`
func (r RouteBuilder) ChannelMessages(channelID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/messages", escapeRouteParam(channelID.String()))
}

// ChannelMessage returns the route for a specific channel message.
//
// Route for:
// - GET    `/channels/{channel.id}/messages/{message.id}`
// - PATCH  `/channels/{channel.id}/messages/{message.id}`
// - DELETE `/channels/{channel.id}/messages/{message.id}`
func (r RouteBuilder) ChannelMessage(channelID, messageID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/messages/%s",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(messageID.String()))
}

// ChannelMessageCrosspost returns the route for crossposting a message.
//
// Route for:
// - POST `/channels/{channel.id}/messages/{message.id}/crosspost`
func (r RouteBuilder) ChannelMessageCrosspost(channelID, messageID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/messages/%s/crosspost",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(messageID.String()))
}

// ChannelMessageOwnReaction returns the route for own message reactions.
//
// Route for:
// - PUT    `/channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me`
// - DELETE `/channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me`
//
// Note: You need to URL encode the emoji yourself.
func (r RouteBuilder) ChannelMessageOwnReaction(channelID, messageID discord.Snowflake, emoji string) string {
	return fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(messageID.String()),
		escapeRouteParam(emoji))
}

// ChannelMessageUserReaction returns the route for user message reactions.
//
// Route for:
// - DELETE `/channels/{channel.id}/messages/{message.id}/reactions/{emoji}/{user.id}`
//
// Note: You need to URL encode the emoji yourself.
func (r RouteBuilder) ChannelMessageUserReaction(channelID, messageID discord.Snowflake, emoji string, userID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/%s",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(messageID.String()),
		escapeRouteParam(emoji),
		escapeRouteParam(userID.String()))
}

// ChannelMessageReaction returns the route for message reactions.
//
// Route for:
// - GET    `/channels/{channel.id}/messages/{message.id}/reactions/{emoji}`
// - DELETE `/channels/{channel.id}/messages/{message.id}/reactions/{emoji}`
//
// Note: You need to URL encode the emoji yourself.
func (r RouteBuilder) ChannelMessageReaction(channelID, messageID discord.Snowflake, emoji string) string {
	return fmt.Sprintf("/channels/%s/messages/%s/reactions/%s",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(messageID.String()),
		escapeRouteParam(emoji))
}

// ChannelMessageAllReactions returns the route for deleting all message reactions.
//
// Route for:
// - DELETE `/channels/{channel.id}/messages/{message.id}/reactions`
func (r RouteBuilder) ChannelMessageAllReactions(channelID, messageID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/messages/%s/reactions",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(messageID.String()))
}

// ChannelBulkDelete returns the route for bulk deleting messages.
//
// Route for:
// - POST `/channels/{channel.id}/messages/bulk-delete`
func (r RouteBuilder) ChannelBulkDelete(channelID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/messages/bulk-delete", escapeRouteParam(channelID.String()))
}

// ChannelPermission returns the route for channel permission overwrites.
//
// Route for:
// - PUT    `/channels/{channel.id}/permissions/{overwrite.id}`
// - DELETE `/channels/{channel.id}/permissions/{overwrite.id}`
func (r RouteBuilder) ChannelPermission(channelID, overwriteID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/permissions/%s",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(overwriteID.String()))
}

// ChannelInvites returns the route for channel invites.
//
// Route for:
// - GET  `/channels/{channel.id}/invites`
// - POST `/channels/{channel.id}/invites`
func (r RouteBuilder) ChannelInvites(channelID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/invites", escapeRouteParam(channelID.String()))
}

// ChannelFollowers returns the route for channel followers.
//
// Route for:
// - POST `/channels/{channel.id}/followers`
func (r RouteBuilder) ChannelFollowers(channelID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/followers", escapeRouteParam(channelID.String()))
}

// ChannelTyping returns the route for channel typing indicator.
//
// Route for:
// - POST `/channels/{channel.id}/typing`
func (r RouteBuilder) ChannelTyping(channelID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/typing", escapeRouteParam(channelID.String()))
}

// ChannelMessagesPins returns the route for channel pinned messages.
//
// Route for:
// - GET `/channels/{channel.id}/messages/pins`
func (r RouteBuilder) ChannelMessagesPins(channelID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/messages/pins", escapeRouteParam(channelID.String()))
}

// ChannelMessagesPin returns the route for pinning/unpinning messages.
//
// Route for:
// - PUT    `/channels/{channel.id}/messages/pins/{message.id}`
// - DELETE `/channels/{channel.id}/messages/pins/{message.id}`
func (r RouteBuilder) ChannelMessagesPin(channelID, messageID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/messages/pins/%s",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(messageID.String()))
}

// ChannelRecipient returns the route for group DM recipients.
//
// Route for:
// - PUT    `/channels/{channel.id}/recipients/{user.id}`
// - DELETE `/channels/{channel.id}/recipients/{user.id}`
func (r RouteBuilder) ChannelRecipient(channelID, userID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/recipients/%s",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(userID.String()))
}

// Guild returns the route for a guild.
//
// Route for:
// - GET    `/guilds/{guild.id}`
// - PATCH  `/guilds/{guild.id}`
// - DELETE `/guilds/{guild.id}` (deprecated)
func (r RouteBuilder) Guild(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s", escapeRouteParam(guildID.String()))
}

// GuildChannels returns the route for guild channels.
//
// Route for:
// - GET   `/guilds/{guild.id}/channels`
// - POST  `/guilds/{guild.id}/channels`
// - PATCH `/guilds/{guild.id}/channels`
func (r RouteBuilder) GuildChannels(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/channels", escapeRouteParam(guildID.String()))
}

// GuildMember returns the route for guild members.
//
// Route for:
// - GET    `/guilds/{guild.id}/members/{user.id}`
// - PUT    `/guilds/{guild.id}/members/{user.id}`
// - PATCH  `/guilds/{guild.id}/members/@me`
// - PATCH  `/guilds/{guild.id}/members/{user.id}`
// - DELETE `/guilds/{guild.id}/members/{user.id}`
func (r RouteBuilder) GuildMember(guildID discord.Snowflake, userID string) string {
	if userID == "" {
		userID = "@me"
	}
	return fmt.Sprintf("/guilds/%s/members/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(userID))
}

// GuildMembers returns the route for guild members.
//
// Route for:
// - GET `/guilds/{guild.id}/members`
func (r RouteBuilder) GuildMembers(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/members", escapeRouteParam(guildID.String()))
}

// GuildMembersSearch returns the route for searching guild members.
//
// Route for:
// - GET `/guilds/{guild.id}/members/search`
func (r RouteBuilder) GuildMembersSearch(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/members/search", escapeRouteParam(guildID.String()))
}

// GuildMemberRole returns the route for guild member roles.
//
// Route for:
// - PUT    `/guilds/{guild.id}/members/{user.id}/roles/{role.id}`
// - DELETE `/guilds/{guild.id}/members/{user.id}/roles/{role.id}`
func (r RouteBuilder) GuildMemberRole(guildID, memberID, roleID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/members/%s/roles/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(memberID.String()),
		escapeRouteParam(roleID.String()))
}

// GuildBans returns the route for guild bans.
//
// Route for:
// - GET `/guilds/{guild.id}/bans`
func (r RouteBuilder) GuildBans(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/bans", escapeRouteParam(guildID.String()))
}

// GuildBan returns the route for a specific guild ban.
//
// Route for:
// - GET    `/guilds/{guild.id}/bans/{user.id}`
// - PUT    `/guilds/{guild.id}/bans/{user.id}`
// - DELETE `/guilds/{guild.id}/bans/{user.id}`
func (r RouteBuilder) GuildBan(guildID, userID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/bans/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(userID.String()))
}

// GuildBulkBan returns the route for bulk banning users.
//
// Route for:
// - POST `/guilds/{guild.id}/bulk-ban`
func (r RouteBuilder) GuildBulkBan(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/bulk-ban", escapeRouteParam(guildID.String()))
}

// GuildRoles returns the route for guild roles.
//
// Route for:
// - GET   `/guilds/{guild.id}/roles`
// - POST  `/guilds/{guild.id}/roles`
// - PATCH `/guilds/{guild.id}/roles`
func (r RouteBuilder) GuildRoles(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/roles", escapeRouteParam(guildID.String()))
}

// GuildRole returns the route for a specific guild role.
//
// Route for:
// - GET    `/guilds/{guild.id}/roles/{role.id}`
// - PATCH  `/guilds/{guild.id}/roles/{role.id}`
// - DELETE `/guilds/{guild.id}/roles/{role.id}`
func (r RouteBuilder) GuildRole(guildID, roleID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/roles/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(roleID.String()))
}

// GuildEmojis returns the route for guild emojis.
//
// Route for:
// - GET  `/guilds/{guild.id}/emojis`
// - POST `/guilds/{guild.id}/emojis`
func (r RouteBuilder) GuildEmojis(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/emojis", escapeRouteParam(guildID.String()))
}

// GuildEmoji returns the route for a specific guild emoji.
//
// Route for:
// - GET    `/guilds/{guild.id}/emojis/{emoji.id}`
// - PATCH  `/guilds/{guild.id}/emojis/{emoji.id}`
// - DELETE `/guilds/{guild.id}/emojis/{emoji.id}`
func (r RouteBuilder) GuildEmoji(guildID, emojiID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/emojis/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(emojiID.String()))
}

// GuildPreview returns the route for guild preview.
//
// Route for:
// - GET `/guilds/{guild.id}/preview`
func (r RouteBuilder) GuildPreview(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/preview", escapeRouteParam(guildID.String()))
}

// GuildPrune returns the route for guild member pruning.
//
// Route for:
// - GET  `/guilds/{guild.id}/prune`
// - POST `/guilds/{guild.id}/prune`
func (r RouteBuilder) GuildPrune(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/prune", escapeRouteParam(guildID.String()))
}

// GuildVoiceRegions returns the route for guild voice regions.
//
// Route for:
// - GET `/guilds/{guild.id}/regions`
func (r RouteBuilder) GuildVoiceRegions(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/regions", escapeRouteParam(guildID.String()))
}

// GuildInvites returns the route for guild invites.
//
// Route for:
// - GET `/guilds/{guild.id}/invites`
func (r RouteBuilder) GuildInvites(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/invites", escapeRouteParam(guildID.String()))
}

// GuildIntegrations returns the route for guild integrations.
//
// Route for:
// - GET `/guilds/{guild.id}/integrations`
func (r RouteBuilder) GuildIntegrations(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/integrations", escapeRouteParam(guildID.String()))
}

// GuildIntegration returns the route for a specific guild integration.
//
// Route for:
// - DELETE `/guilds/{guild.id}/integrations/{integration.id}`
func (r RouteBuilder) GuildIntegration(guildID, integrationID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/integrations/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(integrationID.String()))
}

// GuildWidgetSettings returns the route for guild widget settings.
//
// Route for:
// - GET   `/guilds/{guild.id}/widget`
// - PATCH `/guilds/{guild.id}/widget`
func (r RouteBuilder) GuildWidgetSettings(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/widget", escapeRouteParam(guildID.String()))
}

// GuildWidgetJSON returns the route for guild widget JSON.
//
// Route for:
// - GET `/guilds/{guild.id}/widget.json`
func (r RouteBuilder) GuildWidgetJSON(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/widget.json", escapeRouteParam(guildID.String()))
}

// GuildVanityURL returns the route for guild vanity URL.
//
// Route for:
// - GET `/guilds/{guild.id}/vanity-url`
func (r RouteBuilder) GuildVanityURL(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/vanity-url", escapeRouteParam(guildID.String()))
}

// GuildWidgetImage returns the route for guild widget image.
//
// Route for:
// - GET `/guilds/{guild.id}/widget.png`
func (r RouteBuilder) GuildWidgetImage(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/widget.png", escapeRouteParam(guildID.String()))
}

// User returns the route for users.
//
// Route for:
// - GET   `/users/@me`
// - GET   `/users/{user.id}`
// - PATCH `/users/@me`
func (r RouteBuilder) User(userID string) string {
	if userID == "" {
		userID = "@me"
	}
	return fmt.Sprintf("/users/%s", escapeRouteParam(userID))
}

// UserGuilds returns the route for user guilds.
//
// Route for:
// - GET `/users/@me/guilds`
func (r RouteBuilder) UserGuilds() string {
	return "/users/@me/guilds"
}

// Gateway returns the route for gateway information.
//
// Route for:
// - GET `/gateway`
func (r RouteBuilder) Gateway() string {
	return "/gateway"
}

// GatewayBot returns the route for gateway bot information.
//
// Route for:
// - GET `/gateway/bot`
func (r RouteBuilder) GatewayBot() string {
	return "/gateway/bot"
}

// Webhook returns the route for webhooks.
//
// Route for:
// - GET    `/webhooks/{webhook.id}`
// - GET    `/webhooks/{webhook.id}/{webhook.token}`
// - PATCH  `/webhooks/{webhook.id}`
// - PATCH  `/webhooks/{webhook.id}/{webhook.token}`
// - DELETE `/webhooks/{webhook.id}`
// - DELETE `/webhooks/{webhook.id}/{webhook.token}`
// - POST   `/webhooks/{webhook.id}/{webhook.token}`
// - POST   `/webhooks/{application.id}/{interaction.token}`
func (r RouteBuilder) Webhook(webhookID discord.Snowflake, webhookToken string) string {
	if webhookToken == "" {
		return fmt.Sprintf("/webhooks/%s", escapeRouteParam(webhookID.String()))
	}
	return fmt.Sprintf("/webhooks/%s/%s",
		escapeRouteParam(webhookID.String()),
		escapeRouteParam(webhookToken))
}

// ApplicationCommands returns the route for application commands.
//
// Route for:
// - GET  `/applications/{application.id}/commands`
// - PUT  `/applications/{application.id}/commands`
// - POST `/applications/{application.id}/commands`
func (r RouteBuilder) ApplicationCommands(applicationID discord.Snowflake) string {
	return fmt.Sprintf("/applications/%s/commands", escapeRouteParam(applicationID.String()))
}

// ApplicationCommand returns the route for a specific application command.
//
// Route for:
// - GET    `/applications/{application.id}/commands/{command.id}`
// - PATCH  `/applications/{application.id}/commands/{command.id}`
// - DELETE `/applications/{application.id}/commands/{command.id}`
func (r RouteBuilder) ApplicationCommand(applicationID, commandID discord.Snowflake) string {
	return fmt.Sprintf("/applications/%s/commands/%s",
		escapeRouteParam(applicationID.String()),
		escapeRouteParam(commandID.String()))
}

// ApplicationGuildCommands returns the route for guild-specific application commands.
//
// Route for:
// - GET  `/applications/{application.id}/guilds/{guild.id}/commands`
// - PUT  `/applications/{application.id}/guilds/{guild.id}/commands`
// - POST `/applications/{application.id}/guilds/{guild.id}/commands`
func (r RouteBuilder) ApplicationGuildCommands(applicationID, guildID discord.Snowflake) string {
	return fmt.Sprintf("/applications/%s/guilds/%s/commands",
		escapeRouteParam(applicationID.String()),
		escapeRouteParam(guildID.String()))
}

// InteractionCallback returns the route for interaction callbacks.
//
// Route for:
// - POST `/interactions/{interaction.id}/{interaction.token}/callback`
func (r RouteBuilder) InteractionCallback(interactionID discord.Snowflake, interactionToken string) string {
	return fmt.Sprintf("/interactions/%s/%s/callback",
		escapeRouteParam(interactionID.String()),
		escapeRouteParam(interactionToken))
}

// UserApplicationRoleConnection returns the route for user application role connections.
//
// Route for:
// - GET `/users/@me/applications/{application.id}/role-connection`
// - PUT `/users/@me/applications/{application.id}/role-connection`
func (r RouteBuilder) UserApplicationRoleConnection(applicationID discord.Snowflake) string {
	return fmt.Sprintf("/users/@me/applications/%s/role-connection", escapeRouteParam(applicationID.String()))
}

// UserGuildMember returns the route for current user's guild member.
//
// Route for:
// - GET `/users/@me/guilds/{guild.id}/member`
func (r RouteBuilder) UserGuildMember(guildID discord.Snowflake) string {
	return fmt.Sprintf("/users/@me/guilds/%s/member", escapeRouteParam(guildID.String()))
}

// UserGuild returns the route for leaving a guild.
//
// Route for:
// - DELETE `/users/@me/guilds/{guild.id}`
func (r RouteBuilder) UserGuild(guildID discord.Snowflake) string {
	return fmt.Sprintf("/users/@me/guilds/%s", escapeRouteParam(guildID.String()))
}

// UserChannels returns the route for user DM channels.
//
// Route for:
// - POST `/users/@me/channels`
func (r RouteBuilder) UserChannels() string {
	return "/users/@me/channels"
}

// UserConnections returns the route for user connections.
//
// Route for:
// - GET `/users/@me/connections`
func (r RouteBuilder) UserConnections() string {
	return "/users/@me/connections"
}

// VoiceRegions returns the route for voice regions.
//
// Route for:
// - GET `/voice/regions`
func (r RouteBuilder) VoiceRegions() string {
	return "/voice/regions"
}

// Invite returns the route for invites.
//
// Route for:
// - GET    `/invites/{invite.code}`
// - DELETE `/invites/{invite.code}`
func (r RouteBuilder) Invite(code string) string {
	return fmt.Sprintf("/invites/%s", escapeRouteParam(code))
}

// Template returns the route for templates.
//
// Route for:
// - GET  `/guilds/templates/{template.code}`
// - POST `/guilds/templates/{template.code}`
func (r RouteBuilder) Template(code string) string {
	return fmt.Sprintf("/guilds/templates/%s", escapeRouteParam(code))
}

// GuildTemplates returns the route for guild templates.
//
// Route for:
// - GET  `/guilds/{guild.id}/templates`
// - POST `/guilds/{guild.id}/templates`
func (r RouteBuilder) GuildTemplates(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/templates", escapeRouteParam(guildID.String()))
}

// GuildTemplate returns the route for a specific guild template.
//
// Route for:
// - PUT    `/guilds/{guild.id}/templates/{template.code}`
// - PATCH  `/guilds/{guild.id}/templates/{template.code}`
// - DELETE `/guilds/{guild.id}/templates/{template.code}`
func (r RouteBuilder) GuildTemplate(guildID discord.Snowflake, code string) string {
	return fmt.Sprintf("/guilds/%s/templates/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(code))
}

// ChannelWebhooks returns the route for channel webhooks.
//
// Route for:
// - GET  `/channels/{channel.id}/webhooks`
// - POST `/channels/{channel.id}/webhooks`
func (r RouteBuilder) ChannelWebhooks(channelID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/webhooks", escapeRouteParam(channelID.String()))
}

// GuildWebhooks returns the route for guild webhooks.
//
// Route for:
// - GET `/guilds/{guild.id}/webhooks`
func (r RouteBuilder) GuildWebhooks(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/webhooks", escapeRouteParam(guildID.String()))
}

// WebhookMessage returns the route for webhook messages.
//
// Route for:
// - GET    `/webhooks/{webhook.id}/{webhook.token}/messages/@original`
// - GET    `/webhooks/{webhook.id}/{webhook.token}/messages/{message.id}`
// - PATCH  `/webhooks/{webhook.id}/{webhook.token}/messages/@original`
// - PATCH  `/webhooks/{webhook.id}/{webhook.token}/messages/{message.id}`
// - DELETE `/webhooks/{webhook.id}/{webhook.token}/messages/@original`
// - DELETE `/webhooks/{webhook.id}/{webhook.token}/messages/{message.id}`
func (r RouteBuilder) WebhookMessage(webhookID discord.Snowflake, webhookToken string, messageID string) string {
	if messageID == "" {
		messageID = "@original"
	}
	return fmt.Sprintf("/webhooks/%s/%s/messages/%s",
		escapeRouteParam(webhookID.String()),
		escapeRouteParam(webhookToken),
		escapeRouteParam(messageID))
}

// WebhookPlatform returns the route for platform-specific webhooks.
//
// Route for:
// - POST `/webhooks/{webhook.id}/{webhook.token}/github`
// - POST `/webhooks/{webhook.id}/{webhook.token}/slack`
func (r RouteBuilder) WebhookPlatform(webhookID discord.Snowflake, webhookToken, platform string) string {
	return fmt.Sprintf("/webhooks/%s/%s/%s",
		escapeRouteParam(webhookID.String()),
		escapeRouteParam(webhookToken),
		escapeRouteParam(platform))
}

// OAuth2CurrentApplication returns the route for current OAuth2 application.
//
// Route for:
// - GET `/oauth2/applications/@me`
func (r RouteBuilder) OAuth2CurrentApplication() string {
	return "/oauth2/applications/@me"
}

// OAuth2CurrentAuthorization returns the route for current OAuth2 authorization.
//
// Route for:
// - GET `/oauth2/@me`
func (r RouteBuilder) OAuth2CurrentAuthorization() string {
	return "/oauth2/@me"
}

// OAuth2Authorization returns the route for OAuth2 authorization.
//
// Route for:
// - GET `/oauth2/authorize`
func (r RouteBuilder) OAuth2Authorization() string {
	return "/oauth2/authorize"
}

// OAuth2TokenExchange returns the route for OAuth2 token exchange.
//
// Route for:
// - POST `/oauth2/token`
func (r RouteBuilder) OAuth2TokenExchange() string {
	return "/oauth2/token"
}

// OAuth2TokenRevocation returns the route for OAuth2 token revocation.
//
// Route for:
// - POST `/oauth2/token/revoke`
func (r RouteBuilder) OAuth2TokenRevocation() string {
	return "/oauth2/token/revoke"
}

// Threads returns the route for creating threads.
//
// Route for:
// - POST `/channels/{channel.id}/threads`
// - POST `/channels/{channel.id}/messages/{message.id}/threads`
func (r RouteBuilder) Threads(parentID discord.Snowflake, messageID *discord.Snowflake) string {
	if messageID != nil {
		return fmt.Sprintf("/channels/%s/messages/%s/threads",
			escapeRouteParam(parentID.String()),
			escapeRouteParam(messageID.String()))
	}
	return fmt.Sprintf("/channels/%s/threads", escapeRouteParam(parentID.String()))
}

// GuildActiveThreads returns the route for guild active threads.
//
// Route for:
// - GET `/guilds/{guild.id}/threads/active`
func (r RouteBuilder) GuildActiveThreads(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/threads/active", escapeRouteParam(guildID.String()))
}

// ChannelThreads returns the route for channel archived threads.
//
// Route for:
// - GET `/channels/{channel.id}/threads/archived/public`
// - GET `/channels/{channel.id}/threads/archived/private`
func (r RouteBuilder) ChannelThreads(channelID discord.Snowflake, archivedStatus string) string {
	return fmt.Sprintf("/channels/%s/threads/archived/%s",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(archivedStatus))
}

// ChannelJoinedArchivedThreads returns the route for user's joined archived threads.
//
// Route for:
// - GET `/channels/{channel.id}/users/@me/threads/archived/private`
func (r RouteBuilder) ChannelJoinedArchivedThreads(channelID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/users/@me/threads/archived/private", escapeRouteParam(channelID.String()))
}

// ThreadMembers returns the route for thread members.
//
// Route for:
// - GET    `/channels/{thread.id}/thread-members`
// - GET    `/channels/{thread.id}/thread-members/{user.id}`
// - PUT    `/channels/{thread.id}/thread-members/@me`
// - PUT    `/channels/{thread.id}/thread-members/{user.id}`
// - DELETE `/channels/{thread.id}/thread-members/@me`
// - DELETE `/channels/{thread.id}/thread-members/{user.id}`
func (r RouteBuilder) ThreadMembers(threadID discord.Snowflake, userID *string) string {
	if userID != nil {
		return fmt.Sprintf("/channels/%s/thread-members/%s",
			escapeRouteParam(threadID.String()),
			escapeRouteParam(*userID))
	}
	return fmt.Sprintf("/channels/%s/thread-members", escapeRouteParam(threadID.String()))
}

// PollAnswerVoters returns the route for poll answer voters.
//
// Route for:
// - GET `/channels/{channel.id}/polls/{message.id}/answers/{answer_id}`
func (r RouteBuilder) PollAnswerVoters(channelID, messageID discord.Snowflake, answerID int) string {
	return fmt.Sprintf("/channels/%s/polls/%s/answers/%d",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(messageID.String()),
		answerID)
}

// ExpirePoll returns the route for expiring a poll.
//
// Route for:
// - POST `/channels/{channel.id}/polls/{message.id}/expire`
func (r RouteBuilder) ExpirePoll(channelID, messageID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/polls/%s/expire",
		escapeRouteParam(channelID.String()),
		escapeRouteParam(messageID.String()))
}

// Sticker returns the route for stickers.
//
// Route for:
// - GET `/stickers/{sticker.id}`
func (r RouteBuilder) Sticker(stickerID discord.Snowflake) string {
	return fmt.Sprintf("/stickers/%s", escapeRouteParam(stickerID.String()))
}

// StickerPacks returns the route for sticker packs.
//
// Route for:
// - GET `/sticker-packs`
func (r RouteBuilder) StickerPacks() string {
	return "/sticker-packs"
}

// StickerPack returns the route for a specific sticker pack.
//
// Route for:
// - GET `/sticker-packs/{pack.id}`
func (r RouteBuilder) StickerPack(packID discord.Snowflake) string {
	return fmt.Sprintf("/sticker-packs/%s", escapeRouteParam(packID.String()))
}

// GuildStickers returns the route for guild stickers.
//
// Route for:
// - GET  `/guilds/{guild.id}/stickers`
// - POST `/guilds/{guild.id}/stickers`
func (r RouteBuilder) GuildStickers(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/stickers", escapeRouteParam(guildID.String()))
}

// GuildSticker returns the route for a specific guild sticker.
//
// Route for:
// - GET    `/guilds/{guild.id}/stickers/{sticker.id}`
// - PATCH  `/guilds/{guild.id}/stickers/{sticker.id}`
// - DELETE `/guilds/{guild.id}/stickers/{sticker.id}`
func (r RouteBuilder) GuildSticker(guildID, stickerID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/stickers/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(stickerID.String()))
}

// GuildScheduledEvents returns the route for guild scheduled events.
//
// Route for:
// - GET  `/guilds/{guild.id}/scheduled-events`
// - POST `/guilds/{guild.id}/scheduled-events`
func (r RouteBuilder) GuildScheduledEvents(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/scheduled-events", escapeRouteParam(guildID.String()))
}

// GuildScheduledEvent returns the route for a specific guild scheduled event.
//
// Route for:
// - GET    `/guilds/{guild.id}/scheduled-events/{event.id}`
// - PATCH  `/guilds/{guild.id}/scheduled-events/{event.id}`
// - DELETE `/guilds/{guild.id}/scheduled-events/{event.id}`
func (r RouteBuilder) GuildScheduledEvent(guildID, eventID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/scheduled-events/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(eventID.String()))
}

// GuildScheduledEventUsers returns the route for guild scheduled event users.
//
// Route for:
// - GET `/guilds/{guild.id}/scheduled-events/{event.id}/users`
func (r RouteBuilder) GuildScheduledEventUsers(guildID, eventID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/scheduled-events/%s/users",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(eventID.String()))
}

// SendSoundboardSound returns the route for sending soundboard sounds.
//
// Route for:
// - POST `/channels/{channel.id}/send-soundboard-sound`
func (r RouteBuilder) SendSoundboardSound(channelID discord.Snowflake) string {
	return fmt.Sprintf("/channels/%s/send-soundboard-sound", escapeRouteParam(channelID.String()))
}

// SoundboardDefaultSounds returns the route for default soundboard sounds.
//
// Route for:
// - GET `/soundboard-default-sounds`
func (r RouteBuilder) SoundboardDefaultSounds() string {
	return "/soundboard-default-sounds"
}

// GuildSoundboardSounds returns the route for guild soundboard sounds.
//
// Route for:
// - GET  `/guilds/{guild.id}/soundboard-sounds`
// - POST `/guilds/{guild.id}/soundboard-sounds`
func (r RouteBuilder) GuildSoundboardSounds(guildID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/soundboard-sounds", escapeRouteParam(guildID.String()))
}

// GuildSoundboardSound returns the route for a specific guild soundboard sound.
//
// Route for:
// - GET    `/guilds/{guild.id}/soundboard-sounds/{sound.id}`
// - PATCH  `/guilds/{guild.id}/soundboard-sounds/{sound.id}`
// - DELETE `/guilds/{guild.id}/soundboard-sounds/{sound.id}`
func (r RouteBuilder) GuildSoundboardSound(guildID, soundID discord.Snowflake) string {
	return fmt.Sprintf("/guilds/%s/soundboard-sounds/%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(soundID.String()))
}

// StageInstances returns the route for stage instances.
//
// Route for:
// - POST `/stage-instances`
func (r RouteBuilder) StageInstances() string {
	return "/stage-instances"
}

// StageInstance returns the route for a specific stage instance.
//
// Route for:
// - GET    `/stage-instances/{channel.id}`
// - PATCH  `/stage-instances/{channel.id}`
// - DELETE `/stage-instances/{channel.id}`
func (r RouteBuilder) StageInstance(channelID discord.Snowflake) string {
	return fmt.Sprintf("/stage-instances/%s", escapeRouteParam(channelID.String()))
}

// CurrentApplication returns the route for current application.
//
// Route for:
// - GET   `/applications/@me`
// - PATCH `/applications/@me`
func (r RouteBuilder) CurrentApplication() string {
	return "/applications/@me"
}

// Entitlements returns the route for application entitlements.
//
// Route for:
// - GET  `/applications/{application.id}/entitlements`
// - POST `/applications/{application.id}/entitlements`
func (r RouteBuilder) Entitlements(applicationID discord.Snowflake) string {
	return fmt.Sprintf("/applications/%s/entitlements", escapeRouteParam(applicationID.String()))
}

// Entitlement returns the route for a specific entitlement.
//
// Route for:
// - GET    `/applications/{application.id}/entitlements/{entitlement.id}`
// - DELETE `/applications/{application.id}/entitlements/{entitlement.id}`
func (r RouteBuilder) Entitlement(applicationID, entitlementID discord.Snowflake) string {
	return fmt.Sprintf("/applications/%s/entitlements/%s",
		escapeRouteParam(applicationID.String()),
		escapeRouteParam(entitlementID.String()))
}

// SKUs returns the route for application SKUs.
//
// Route for:
// - GET `/applications/{application.id}/skus`
func (r RouteBuilder) SKUs(applicationID discord.Snowflake) string {
	return fmt.Sprintf("/applications/%s/skus", escapeRouteParam(applicationID.String()))
}

// ConsumeEntitlement returns the route for consuming an entitlement.
//
// Route for:
// - POST `/applications/{application.id}/entitlements/{entitlement.id}/consume`
func (r RouteBuilder) ConsumeEntitlement(applicationID, entitlementID discord.Snowflake) string {
	return fmt.Sprintf("/applications/%s/entitlements/%s/consume",
		escapeRouteParam(applicationID.String()),
		escapeRouteParam(entitlementID.String()))
}

// ApplicationEmojis returns the route for application emojis.
//
// Route for:
// - GET  `/applications/{application.id}/emojis`
// - POST `/applications/{application.id}/emojis`
func (r RouteBuilder) ApplicationEmojis(applicationID discord.Snowflake) string {
	return fmt.Sprintf("/applications/%s/emojis", escapeRouteParam(applicationID.String()))
}

// ApplicationEmoji returns the route for a specific application emoji.
//
// Route for:
// - GET    `/applications/{application.id}/emojis/{emoji.id}`
// - PATCH  `/applications/{application.id}/emojis/{emoji.id}`
// - DELETE `/applications/{application.id}/emojis/{emoji.id}`
func (r RouteBuilder) ApplicationEmoji(applicationID, emojiID discord.Snowflake) string {
	return fmt.Sprintf("/applications/%s/emojis/%s",
		escapeRouteParam(applicationID.String()),
		escapeRouteParam(emojiID.String()))
}

// escapeRouteParam safely escapes a route parameter for URL usage.
// If the parameter already contains safe characters, it returns as-is.
// Otherwise, it URL-encodes the parameter.
func escapeRouteParam(param string) string {
	if param == "" {
		return param
	}

	// If the parameter already contains only URL-safe characters, return as-is
	if urlSafeCharacters.MatchString(param) {
		return param
	}

	// Otherwise, URL-encode it
	return url.QueryEscape(param)
}

// ImageFormat represents supported image formats for CDN assets.
type ImageFormat string

const (
	ImageFormatJPEG   ImageFormat = "jpeg"
	ImageFormatPNG    ImageFormat = "png"
	ImageFormatWebP   ImageFormat = "webp"
	ImageFormatGIF    ImageFormat = "gif"
	ImageFormatLottie ImageFormat = "json"
)

// ImageSize represents valid image sizes for CDN assets.
// Image size can be any power of two between 16 and 4096.
type ImageSize int

const (
	ImageSize16   ImageSize = 16
	ImageSize32   ImageSize = 32
	ImageSize64   ImageSize = 64
	ImageSize128  ImageSize = 128
	ImageSize256  ImageSize = 256
	ImageSize512  ImageSize = 512
	ImageSize1024 ImageSize = 1024
	ImageSize2048 ImageSize = 2048
	ImageSize4096 ImageSize = 4096
)

// DefaultUserAvatarAssets represents the valid default user avatar indices.
type DefaultUserAvatarAssets int

const (
	DefaultUserAvatar0 DefaultUserAvatarAssets = 0
	DefaultUserAvatar1 DefaultUserAvatarAssets = 1
	DefaultUserAvatar2 DefaultUserAvatarAssets = 2
	DefaultUserAvatar3 DefaultUserAvatarAssets = 3
	DefaultUserAvatar4 DefaultUserAvatarAssets = 4
	DefaultUserAvatar5 DefaultUserAvatarAssets = 5
)

// CDNQuery represents query parameters for CDN requests.
type CDNQuery struct {
	// Size is the returned image size. Must be a power of two between 16 and 4096.
	Size *ImageSize `url:"size,omitempty"`
}

// CDNRouteBuilder provides methods to build Discord CDN routes.
type CDNRouteBuilder struct{}

// CDNRoutes is the global CDN route builder instance.
var CDNRoutes = CDNRouteBuilder{}

// Emoji returns the CDN route for emojis.
//
// Route for:
// - GET `/emojis/{emoji.id}.{png|jpeg|webp|gif}`
//
// As this route supports GIFs, the hash will begin with `a_` if it is available in GIF format.
// This route supports the extensions: PNG, JPEG, WebP, GIF
func (r CDNRouteBuilder) Emoji(emojiID discord.Snowflake, format ImageFormat) string {
	return fmt.Sprintf("/emojis/%s.%s", escapeRouteParam(emojiID.String()), format)
}

// GuildIcon returns the CDN route for guild icons.
//
// Route for:
// - GET `/icons/{guild.id}/{guild.icon}.{png|jpeg|webp|gif}`
//
// As this route supports GIFs, the hash will begin with `a_` if it is available in GIF format.
// This route supports the extensions: PNG, JPEG, WebP, GIF
func (r CDNRouteBuilder) GuildIcon(guildID discord.Snowflake, guildIcon string, format ImageFormat) string {
	return fmt.Sprintf("/icons/%s/%s.%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(guildIcon),
		format)
}

// GuildSplash returns the CDN route for guild splash images.
//
// Route for:
// - GET `/splashes/{guild.id}/{guild.splash}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) GuildSplash(guildID discord.Snowflake, guildSplash string, format ImageFormat) string {
	return fmt.Sprintf("/splashes/%s/%s.%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(guildSplash),
		format)
}

// GuildDiscoverySplash returns the CDN route for guild discovery splash images.
//
// Route for:
// - GET `/discovery-splashes/{guild.id}/{guild.discovery_splash}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) GuildDiscoverySplash(guildID discord.Snowflake, guildDiscoverySplash string, format ImageFormat) string {
	return fmt.Sprintf("/discovery-splashes/%s/%s.%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(guildDiscoverySplash),
		format)
}

// GuildBanner returns the CDN route for guild banners.
//
// Route for:
// - GET `/banners/{guild.id}/{guild.banner}.{png|jpeg|webp|gif}`
//
// As this route supports GIFs, the hash will begin with `a_` if it is available in GIF format.
// This route supports the extensions: PNG, JPEG, WebP, GIF
func (r CDNRouteBuilder) GuildBanner(guildID discord.Snowflake, guildBanner string, format ImageFormat) string {
	return fmt.Sprintf("/banners/%s/%s.%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(guildBanner),
		format)
}

// UserBanner returns the CDN route for user banners.
//
// Route for:
// - GET `/banners/{user.id}/{user.banner}.{png|jpeg|webp|gif}`
//
// As this route supports GIFs, the hash will begin with `a_` if it is available in GIF format.
// This route supports the extensions: PNG, JPEG, WebP, GIF
func (r CDNRouteBuilder) UserBanner(userID discord.Snowflake, userBanner string, format ImageFormat) string {
	return fmt.Sprintf("/banners/%s/%s.%s",
		escapeRouteParam(userID.String()),
		escapeRouteParam(userBanner),
		format)
}

// DefaultUserAvatar returns the CDN route for default user avatars.
//
// Route for:
// - GET `/embed/avatars/{index}.png`
//
// The value for `index` parameter depends on whether the user is migrated to the new username system.
// For users on the new username system, `index` will be `(user.id >> 22) % 6`.
// For users on the legacy username system, `index` will be `user.discriminator % 5`.
//
// This route supports the extension: PNG
func (r CDNRouteBuilder) DefaultUserAvatar(index DefaultUserAvatarAssets) string {
	return fmt.Sprintf("/embed/avatars/%d.png", index)
}

// UserAvatar returns the CDN route for user avatars.
//
// Route for:
// - GET `/avatars/{user.id}/{user.avatar}.{png|jpeg|webp|gif}`
//
// As this route supports GIFs, the hash will begin with `a_` if it is available in GIF format.
// This route supports the extensions: PNG, JPEG, WebP, GIF
func (r CDNRouteBuilder) UserAvatar(userID discord.Snowflake, userAvatar string, format ImageFormat) string {
	return fmt.Sprintf("/avatars/%s/%s.%s",
		escapeRouteParam(userID.String()),
		escapeRouteParam(userAvatar),
		format)
}

// GuildMemberAvatar returns the CDN route for guild member avatars.
//
// Route for:
// - GET `/guilds/{guild.id}/users/{user.id}/avatars/{guild_member.avatar}.{png|jpeg|webp|gif}`
//
// As this route supports GIFs, the hash will begin with `a_` if it is available in GIF format.
// This route supports the extensions: PNG, JPEG, WebP, GIF
func (r CDNRouteBuilder) GuildMemberAvatar(guildID, userID discord.Snowflake, memberAvatar string, format ImageFormat) string {
	return fmt.Sprintf("/guilds/%s/users/%s/avatars/%s.%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(userID.String()),
		escapeRouteParam(memberAvatar),
		format)
}

// AvatarDecoration returns the CDN route for avatar decorations.
//
// Route for:
// - GET `/avatar-decoration-presets/{avatar_decoration_data_asset}.png`
//
// This route supports the extension: PNG
func (r CDNRouteBuilder) AvatarDecoration(avatarDecorationDataAsset string) string {
	return fmt.Sprintf("/avatar-decoration-presets/%s.png", escapeRouteParam(avatarDecorationDataAsset))
}

// ApplicationIcon returns the CDN route for application icons.
//
// Route for:
// - GET `/app-icons/{application.id}/{application.icon}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) ApplicationIcon(applicationID discord.Snowflake, applicationIcon string, format ImageFormat) string {
	return fmt.Sprintf("/app-icons/%s/%s.%s",
		escapeRouteParam(applicationID.String()),
		escapeRouteParam(applicationIcon),
		format)
}

// ApplicationCover returns the CDN route for application cover images.
//
// Route for:
// - GET `/app-icons/{application.id}/{application.cover_image}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) ApplicationCover(applicationID discord.Snowflake, applicationCoverImage string, format ImageFormat) string {
	return fmt.Sprintf("/app-icons/%s/%s.%s",
		escapeRouteParam(applicationID.String()),
		escapeRouteParam(applicationCoverImage),
		format)
}

// ApplicationAsset returns the CDN route for application assets.
//
// Route for:
// - GET `/app-assets/{application.id}/{application.asset_id}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) ApplicationAsset(applicationID discord.Snowflake, applicationAssetID string, format ImageFormat) string {
	return fmt.Sprintf("/app-assets/%s/%s.%s",
		escapeRouteParam(applicationID.String()),
		escapeRouteParam(applicationAssetID),
		format)
}

// AchievementIcon returns the CDN route for achievement icons.
//
// Route for:
// - GET `/app-assets/{application.id}/achievements/{achievement.id}/icons/{achievement.icon}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) AchievementIcon(applicationID, achievementID discord.Snowflake, achievementIconHash string, format ImageFormat) string {
	return fmt.Sprintf("/app-assets/%s/achievements/%s/icons/%s.%s",
		escapeRouteParam(applicationID.String()),
		escapeRouteParam(achievementID.String()),
		escapeRouteParam(achievementIconHash),
		format)
}

// StickerPackBanner returns the CDN route for sticker pack banners.
//
// Route for:
// - GET `/app-assets/710982414301790216/store/{sticker_pack.banner.asset_id}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) StickerPackBanner(stickerPackBannerAssetID discord.Snowflake, format ImageFormat) string {
	return fmt.Sprintf("/app-assets/710982414301790216/store/%s.%s",
		escapeRouteParam(stickerPackBannerAssetID.String()),
		format)
}

// StorePageAsset returns the CDN route for store page assets.
//
// Route for:
// - GET `/app-assets/{application.id}/store/{asset.id}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) StorePageAsset(applicationID discord.Snowflake, assetID string, format ImageFormat) string {
	return fmt.Sprintf("/app-assets/%s/store/%s.%s",
		escapeRouteParam(applicationID.String()),
		escapeRouteParam(assetID),
		format)
}

// TeamIcon returns the CDN route for team icons.
//
// Route for:
// - GET `/team-icons/{team.id}/{team.icon}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) TeamIcon(teamID discord.Snowflake, teamIcon string, format ImageFormat) string {
	return fmt.Sprintf("/team-icons/%s/%s.%s",
		escapeRouteParam(teamID.String()),
		escapeRouteParam(teamIcon),
		format)
}

// Sticker returns the CDN route for stickers.
//
// Route for:
// - GET `/stickers/{sticker.id}.{png|json}`
//
// This route supports the extensions: PNG, Lottie, GIF
func (r CDNRouteBuilder) Sticker(stickerID discord.Snowflake, format ImageFormat) string {
	return fmt.Sprintf("/stickers/%s.%s", escapeRouteParam(stickerID.String()), format)
}

// RoleIcon returns the CDN route for role icons.
//
// Route for:
// - GET `/role-icons/{role.id}/{role.icon}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) RoleIcon(roleID discord.Snowflake, roleIcon string, format ImageFormat) string {
	return fmt.Sprintf("/role-icons/%s/%s.%s",
		escapeRouteParam(roleID.String()),
		escapeRouteParam(roleIcon),
		format)
}

// GuildScheduledEventCover returns the CDN route for guild scheduled event covers.
//
// Route for:
// - GET `/guild-events/{guild_scheduled_event.id}/{guild_scheduled_event.image}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) GuildScheduledEventCover(guildScheduledEventID discord.Snowflake, guildScheduledEventCoverImage string, format ImageFormat) string {
	return fmt.Sprintf("/guild-events/%s/%s.%s",
		escapeRouteParam(guildScheduledEventID.String()),
		escapeRouteParam(guildScheduledEventCoverImage),
		format)
}

// GuildMemberBanner returns the CDN route for guild member banners.
//
// Route for:
// - GET `/guilds/{guild.id}/users/{user.id}/banners/{guild_member.banner}.{png|jpeg|webp|gif}`
//
// This route supports the extensions: PNG, JPEG, WebP, GIF
func (r CDNRouteBuilder) GuildMemberBanner(guildID, userID discord.Snowflake, guildMemberBanner string, format ImageFormat) string {
	return fmt.Sprintf("/guilds/%s/users/%s/banners/%s.%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(userID.String()),
		escapeRouteParam(guildMemberBanner),
		format)
}

// SoundboardSound returns the CDN route for soundboard sounds.
//
// Route for:
// - GET `/soundboard-sounds/{sound.id}`
func (r CDNRouteBuilder) SoundboardSound(soundID discord.Snowflake) string {
	return fmt.Sprintf("/soundboard-sounds/%s", escapeRouteParam(soundID.String()))
}

// GuildTagBadge returns the CDN route for guild tag badges.
//
// Route for:
// - GET `/guild-tag-badges/{guild.id}/{badge}.{png|jpeg|webp}`
//
// This route supports the extensions: PNG, JPEG, WebP
func (r CDNRouteBuilder) GuildTagBadge(guildID discord.Snowflake, guildTagBadge string, format ImageFormat) string {
	return fmt.Sprintf("/guild-tag-badges/%s/%s.%s",
		escapeRouteParam(guildID.String()),
		escapeRouteParam(guildTagBadge),
		format)
}

// RouteBases contains base URLs for different Discord services.
var RouteBases = struct {
	API            string
	CDN            string
	Media          string
	Invite         string
	Template       string
	Gift           string
	ScheduledEvent string
}{
	API:            fmt.Sprintf("https://discord.com/api/v%s", APIVersion),
	CDN:            "https://cdn.discordapp.com",
	Media:          "https://media.discordapp.net",
	Invite:         "https://discord.gg",
	Template:       "https://discord.new",
	Gift:           "https://discord.gift",
	ScheduledEvent: "https://discord.com/events",
}

// OAuth2Routes contains OAuth2-specific route helpers.
var OAuth2Routes = struct {
	AuthorizationURL   string
	TokenURL           string
	TokenRevocationURL string
}{
	AuthorizationURL:   RouteBases.API + Routes.OAuth2Authorization(),
	TokenURL:           RouteBases.API + Routes.OAuth2TokenExchange(),
	TokenRevocationURL: RouteBases.API + Routes.OAuth2TokenRevocation(),
}
