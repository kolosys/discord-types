package v10

// OAuth2Scope represents Discord OAuth2 scopes.
// Reference: https://discord.com/developers/docs/topics/oauth2#shared-resources-oauth2-scopes
type OAuth2Scope string

// OAuth2 scope constants
const (
	// OAuth2ScopeActivitiesRead allows your app to fetch data from a user's "Now Playing/Recently Played" list.
	OAuth2ScopeActivitiesRead OAuth2Scope = "activities.read"

	// OAuth2ScopeActivitiesWrite allows your app to update a user's activity.
	OAuth2ScopeActivitiesWrite OAuth2Scope = "activities.write"

	// OAuth2ScopeApplicationsBuildsRead allows your app to read build data for a user's applications.
	OAuth2ScopeApplicationsBuildsRead OAuth2Scope = "applications.builds.read"

	// OAuth2ScopeApplicationsBuildsUpload allows your app to upload/update builds for a user's applications.
	OAuth2ScopeApplicationsBuildsUpload OAuth2Scope = "applications.builds.upload"

	// OAuth2ScopeApplicationsCommands allows your app to add commands to a guild.
	OAuth2ScopeApplicationsCommands OAuth2Scope = "applications.commands"

	// OAuth2ScopeApplicationsCommandsUpdate allows your app to update its commands.
	OAuth2ScopeApplicationsCommandsUpdate OAuth2Scope = "applications.commands.update"

	// OAuth2ScopeApplicationsCommandsPermissionsUpdate allows your app to update permissions for its commands in a guild a user has permissions to.
	OAuth2ScopeApplicationsCommandsPermissionsUpdate OAuth2Scope = "applications.commands.permissions.update"

	// OAuth2ScopeApplicationsEntitlements allows your app to read entitlements for a user's applications.
	OAuth2ScopeApplicationsEntitlements OAuth2Scope = "applications.entitlements"

	// OAuth2ScopeApplicationsStoreUpdate allows your app to read and update store data (SKUs, store listings, achievements, etc.) for a user's applications.
	OAuth2ScopeApplicationsStoreUpdate OAuth2Scope = "applications.store.update"

	// OAuth2ScopeBot allows your app to be added to guilds as a bot.
	OAuth2ScopeBot OAuth2Scope = "bot"

	// OAuth2ScopeConnections allows your app to see information about the user's connections to other services.
	OAuth2ScopeConnections OAuth2Scope = "connections"

	// OAuth2ScopeDMChannelsRead allows your app to see information about the user's DM and group DM channels.
	OAuth2ScopeDMChannelsRead OAuth2Scope = "dm_channels.read"

	// OAuth2ScopeEmail allows your app to get the user's email address.
	OAuth2ScopeEmail OAuth2Scope = "email"

	// OAuth2ScopeGDMJoin allows your app to join users to a group dm.
	OAuth2ScopeGDMJoin OAuth2Scope = "gdm.join"

	// OAuth2ScopeGuilds allows your app to see the list of guilds the user is in.
	OAuth2ScopeGuilds OAuth2Scope = "guilds"

	// OAuth2ScopeGuildsJoin allows your app to add the user to a guild as a member.
	OAuth2ScopeGuildsJoin OAuth2Scope = "guilds.join"

	// OAuth2ScopeGuildsMembersRead allows your app to read guild members.
	OAuth2ScopeGuildsMembersRead OAuth2Scope = "guilds.members.read"

	// OAuth2ScopeIdentify allows your app to get basic user information.
	OAuth2ScopeIdentify OAuth2Scope = "identify"

	// OAuth2ScopeMessagesRead allows your app to read messages in channels.
	OAuth2ScopeMessagesRead OAuth2Scope = "messages.read"

	// OAuth2ScopeRelationshipsRead allows your app to see the user's relationships.
	OAuth2ScopeRelationshipsRead OAuth2Scope = "relationships.read"

	// OAuth2ScopeRoleConnectionsWrite allows your app to update a user's connection and metadata for the app.
	OAuth2ScopeRoleConnectionsWrite OAuth2Scope = "role_connections.write"

	// OAuth2ScopeRPC allows for RPC/SDK access.
	OAuth2ScopeRPC OAuth2Scope = "rpc"

	// OAuth2ScopeRPCActivitiesWrite allows your app to control a user's local Discord client.
	OAuth2ScopeRPCActivitiesWrite OAuth2Scope = "rpc.activities.write"

	// OAuth2ScopeRPCNotificationsRead allows your app to receive notifications pushed out to the user.
	OAuth2ScopeRPCNotificationsRead OAuth2Scope = "rpc.notifications.read"

	// OAuth2ScopeRPCVoiceRead allows your app to read a user's voice settings and listen for voice events.
	OAuth2ScopeRPCVoiceRead OAuth2Scope = "rpc.voice.read"

	// OAuth2ScopeRPCVoiceWrite allows your app to update a user's voice settings.
	OAuth2ScopeRPCVoiceWrite OAuth2Scope = "rpc.voice.write"

	// OAuth2ScopeVoice allows your app to connect to voice on user's behalf and see all the voice members.
	OAuth2ScopeVoice OAuth2Scope = "voice"

	// OAuth2ScopeWebhookIncoming allows your app to receive incoming webhooks.
	OAuth2ScopeWebhookIncoming OAuth2Scope = "webhook.incoming"
)
