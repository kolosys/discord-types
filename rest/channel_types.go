// Package rest provides Discord REST API types and utilities.
//
// This file contains comprehensive channel-related REST API request and response types.
package rest

import (
	"github.com/kolosys/discord-types/discord"
	"github.com/kolosys/discord-types/payloads"
)

// ====================
// Channel Types - Complete Implementation
// ====================

// ChannelPatchOverwrite represents a permission overwrite for channel modification
type ChannelPatchOverwrite struct {
	ID    discord.Snowflake      `json:"id"`
	Type  payloads.OverwriteType `json:"type"`
	Allow *discord.Permissions   `json:"allow,omitempty"`
	Deny  *discord.Permissions   `json:"deny,omitempty"`
}

// GetChannelQuery represents query parameters for GET /channels/{channel.id}
type GetChannelQuery struct {
	// No query parameters for this endpoint
}

// GetChannelResult represents the response from GET /channels/{channel.id}
type GetChannelResult = payloads.GuildTextChannel

// PatchChannelJSONBody represents the request body for PATCH /channels/{channel.id}
type PatchChannelJSONBody struct {
	// 1-100 character channel name - Channel types: all
	Name *string `json:"name,omitempty"`

	// The type of channel; only conversion between text and news is supported
	// Channel types: text, news
	Type *payloads.ChannelType `json:"type,omitempty"`

	// The position of the channel in the left-hand listing
	// Channel types: all excluding threads
	Position *int `json:"position,omitempty"`

	// 0-1024 character channel topic (0-4096 characters for thread-only channels)
	// Channel types: text, news, forum, media
	Topic *string `json:"topic,omitempty"`

	// Whether the channel is nsfw
	// Channel types: text, voice, news, forum, media
	NSFW *bool `json:"nsfw,omitempty"`

	// Amount of seconds a user has to wait before sending another message (0-21600)
	// Channel types: text, newsThread, publicThread, privateThread, forum, media
	RateLimitPerUser *int `json:"rate_limit_per_user,omitempty"`

	// The bitrate (in bits) of the voice channel; 8000 to 96000 (128000 for VIP servers)
	// Channel types: voice
	Bitrate *int `json:"bitrate,omitempty"`

	// The user limit of the voice channel; 0 refers to no limit, 1 to 99 refers to a user limit
	// Channel types: voice
	UserLimit *int `json:"user_limit,omitempty"`

	// Channel or category-specific permissions
	// Channel types: all excluding threads
	PermissionOverwrites []ChannelPatchOverwrite `json:"permission_overwrites,omitempty"`

	// ID of the new parent category for a channel
	// Channel types: text, voice, news, stage, forum, media
	ParentID *discord.Snowflake `json:"parent_id,omitempty"`

	// Voice region id for the voice or stage channel, automatic when set to null
	RTCRegion *string `json:"rtc_region,omitempty"`

	// The camera video quality mode of the voice channel
	VideoQualityMode *payloads.VideoQualityMode `json:"video_quality_mode,omitempty"`

	// Whether the thread should be archived
	// Channel types: newsThread, publicThread, privateThread
	Archived *bool `json:"archived,omitempty"`

	// The amount of time in minutes to wait before automatically archiving the thread
	// Channel types: newsThread, publicThread, privateThread
	AutoArchiveDuration *payloads.ThreadAutoArchiveDuration `json:"auto_archive_duration,omitempty"`

	// Whether the thread should be locked
	// Channel types: newsThread, publicThread, privateThread
	Locked *bool `json:"locked,omitempty"`

	// Whether non-moderators can add other non-moderators to a thread
	// Channel types: privateThread
	Invitable *bool `json:"invitable,omitempty"`

	// The default auto-archive duration for threads in forum channels
	// Channel types: forum, media
	DefaultAutoArchiveDuration *payloads.ThreadAutoArchiveDuration `json:"default_auto_archive_duration,omitempty"`

	// Channel flags combined as a bitfield
	Flags *payloads.ChannelFlags `json:"flags,omitempty"`

	// The set of tags that can be used in a forum channel
	// Channel types: forum, media
	AvailableTags []payloads.ForumTag `json:"available_tags,omitempty"`

	// The emoji to show in the add reaction button on a thread in a forum channel
	// Channel types: forum, media
	DefaultReactionEmoji *payloads.DefaultReaction `json:"default_reaction_emoji,omitempty"`

	// The initial rate_limit_per_user to set on newly created threads in a channel
	// Channel types: forum, media
	DefaultThreadRateLimitPerUser *int `json:"default_thread_rate_limit_per_user,omitempty"`

	// The default sort order type used to order posts in forum channels
	// Channel types: forum, media
	DefaultSortOrder *payloads.SortOrderType `json:"default_sort_order,omitempty"`

	// The default forum layout view used to display posts in forum channels
	// Channel types: forum, media
	DefaultForumLayout *payloads.ForumLayoutType `json:"default_forum_layout,omitempty"`
}

// PatchChannelResult represents the response from PATCH /channels/{channel.id}
type PatchChannelResult = payloads.GuildTextChannel

// DeleteChannelQuery represents query parameters for DELETE /channels/{channel.id}
type DeleteChannelQuery struct {
	// No query parameters for this endpoint
}

// DeleteChannelResult represents the response from DELETE /channels/{channel.id}
type DeleteChannelResult = payloads.GuildTextChannel

// GetChannelMessagesQuery represents query parameters for GET /channels/{channel.id}/messages
type GetChannelMessagesQuery struct {
	// Get messages around this message ID
	Around *discord.Snowflake `json:"around,omitempty"`
	// Get messages before this message ID
	Before *discord.Snowflake `json:"before,omitempty"`
	// Get messages after this message ID
	After *discord.Snowflake `json:"after,omitempty"`
	// Max number of messages to return (1-100, default 50)
	Limit *int `json:"limit,omitempty"`
}

// GetChannelMessagesResult represents the response from GET /channels/{channel.id}/messages
type GetChannelMessagesResult = []payloads.Message

// PostChannelMessageJSONBody represents the request body for POST /channels/{channel.id}/messages
type PostChannelMessageJSONBody struct {
	// Message contents (up to 2000 characters)
	Content *string `json:"content,omitempty"`

	// Can be used to verify a message was sent (up to 25 characters)
	Nonce *string `json:"nonce,omitempty"`

	// True if this is a TTS message
	TTS *bool `json:"tts,omitempty"`

	// Up to 10 embed objects
	Embeds []payloads.Embed `json:"embeds,omitempty"`

	// Allowed mentions for the message
	AllowedMentions *payloads.AllowedMentions `json:"allowed_mentions,omitempty"`

	// Include to make your message a reply
	MessageReference *payloads.MessageReference `json:"message_reference,omitempty"`

	// Message components to include (up to 5 action rows)
	Components []payloads.ActionRowComponent `json:"components,omitempty"`

	// IDs of up to 3 stickers in the server to send in the message
	StickerIDs []discord.Snowflake `json:"sticker_ids,omitempty"`

	// Attachment objects with filename and description
	Attachments []payloads.PartialAttachment `json:"attachments,omitempty"`

	// Message flags combined as a bitfield (only SUPPRESS_EMBEDS can be set)
	Flags *payloads.MessageFlags `json:"flags,omitempty"`

	// A poll object
	Poll *payloads.Poll `json:"poll,omitempty"`
}

// PostChannelMessageResult represents the response from POST /channels/{channel.id}/messages
type PostChannelMessageResult = payloads.Message

// GetChannelMessageQuery represents query parameters for GET /channels/{channel.id}/messages/{message.id}
type GetChannelMessageQuery struct {
	// No query parameters for this endpoint
}

// GetChannelMessageResult represents the response from GET /channels/{channel.id}/messages/{message.id}
type GetChannelMessageResult = payloads.Message

// PatchChannelMessageJSONBody represents the request body for PATCH /channels/{channel.id}/messages/{message.id}
type PatchChannelMessageJSONBody struct {
	// Message contents (up to 2000 characters)
	Content *string `json:"content,omitempty"`

	// Up to 10 embed objects
	Embeds []payloads.Embed `json:"embeds,omitempty"`

	// Edit the flags of a message (only SUPPRESS_EMBEDS can be toggled)
	Flags *payloads.MessageFlags `json:"flags,omitempty"`

	// Allowed mentions for the message
	AllowedMentions *payloads.AllowedMentions `json:"allowed_mentions,omitempty"`

	// Message components to include (up to 5 action rows)
	Components []payloads.ActionRowComponent `json:"components,omitempty"`

	// Attachment objects with filename and description
	Attachments []payloads.PartialAttachment `json:"attachments,omitempty"`
}

// PatchChannelMessageResult represents the response from PATCH /channels/{channel.id}/messages/{message.id}
type PatchChannelMessageResult = payloads.Message

// DeleteChannelMessageQuery represents query parameters for DELETE /channels/{channel.id}/messages/{message.id}
type DeleteChannelMessageQuery struct {
	// No query parameters for this endpoint
}

// PostChannelMessageCrosspostResult represents the response from POST /channels/{channel.id}/messages/{message.id}/crosspost
type PostChannelMessageCrosspostResult = payloads.Message

// PutChannelMessageReactionQuery represents query parameters for PUT /channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me
type PutChannelMessageReactionQuery struct {
	// No query parameters for this endpoint
}

// DeleteChannelMessageReactionQuery represents query parameters for DELETE /channels/{channel.id}/messages/{message.id}/reactions/{emoji}/@me
type DeleteChannelMessageReactionQuery struct {
	// No query parameters for this endpoint
}

// DeleteChannelMessageUserReactionQuery represents query parameters for DELETE /channels/{channel.id}/messages/{message.id}/reactions/{emoji}/{user.id}
type DeleteChannelMessageUserReactionQuery struct {
	// No query parameters for this endpoint
}

// GetChannelMessageReactionsQuery represents query parameters for GET /channels/{channel.id}/messages/{message.id}/reactions/{emoji}
type GetChannelMessageReactionsQuery struct {
	// Get users after this user ID
	After *discord.Snowflake `json:"after,omitempty"`
	// Max number of users to return (1-100, default 25)
	Limit *int `json:"limit,omitempty"`
}

// GetChannelMessageReactionsResult represents the response from GET /channels/{channel.id}/messages/{message.id}/reactions/{emoji}
type GetChannelMessageReactionsResult = []payloads.User

// DeleteChannelMessageAllReactionsQuery represents query parameters for DELETE /channels/{channel.id}/messages/{message.id}/reactions
type DeleteChannelMessageAllReactionsQuery struct {
	// No query parameters for this endpoint
}

// DeleteChannelMessageEmojiReactionQuery represents query parameters for DELETE /channels/{channel.id}/messages/{message.id}/reactions/{emoji}
type DeleteChannelMessageEmojiReactionQuery struct {
	// No query parameters for this endpoint
}

// PostChannelMessagesBulkDeleteJSONBody represents the request body for POST /channels/{channel.id}/messages/bulk-delete
type PostChannelMessagesBulkDeleteJSONBody struct {
	// An array of message ids to delete (2-100)
	Messages []discord.Snowflake `json:"messages"`
}

// PutChannelPermissionJSONBody represents the request body for PUT /channels/{channel.id}/permissions/{overwrite.id}
type PutChannelPermissionJSONBody struct {
	// Permission bit set
	Allow *discord.Permissions `json:"allow,omitempty"`
	// Permission bit set
	Deny *discord.Permissions `json:"deny,omitempty"`
	// 0 for a role or 1 for a member
	Type payloads.OverwriteType `json:"type"`
}

// DeleteChannelPermissionQuery represents query parameters for DELETE /channels/{channel.id}/permissions/{overwrite.id}
type DeleteChannelPermissionQuery struct {
	// No query parameters for this endpoint
}

// GetChannelInvitesResult represents the response from GET /channels/{channel.id}/invites
type GetChannelInvitesResult = []Invite

// PostChannelInviteJSONBody represents the request body for POST /channels/{channel.id}/invites
type PostChannelInviteJSONBody struct {
	// Duration of invite in seconds before expiry, or 0 for never (default 86400)
	MaxAge *int `json:"max_age,omitempty"`
	// Max number of uses or 0 for unlimited (default 0)
	MaxUses *int `json:"max_uses,omitempty"`
	// Whether this invite only grants temporary membership (default false)
	Temporary *bool `json:"temporary,omitempty"`
	// If true, don't try to reuse a similar invite (default false)
	Unique *bool `json:"unique,omitempty"`
	// The type of target for this voice channel invite
	TargetType *InviteTargetType `json:"target_type,omitempty"`
	// The id of the user whose stream to display for this invite
	TargetUserID *discord.Snowflake `json:"target_user_id,omitempty"`
	// The id of the embedded application to open for this invite
	TargetApplicationID *discord.Snowflake `json:"target_application_id,omitempty"`
}

// PostChannelInviteResult represents the response from POST /channels/{channel.id}/invite
type PostChannelInviteResult = Invite

// PostChannelFollowersJSONBody represents the request body for POST /channels/{channel.id}/followers
type PostChannelFollowersJSONBody struct {
	// ID of target channel
	WebhookChannelID discord.Snowflake `json:"webhook_channel_id"`
}

// FollowedChannel represents a followed channel
type FollowedChannel struct {
	ChannelID discord.Snowflake `json:"channel_id"`
	WebhookID discord.Snowflake `json:"webhook_id"`
}

// PostChannelFollowersResult represents the response from POST /channels/{channel.id}/followers
type PostChannelFollowersResult = FollowedChannel

// PostChannelTypingQuery represents query parameters for POST /channels/{channel.id}/typing
type PostChannelTypingQuery struct {
	// No query parameters for this endpoint
}

// GetChannelPinsResult represents the response from GET /channels/{channel.id}/pins
type GetChannelPinsResult = []payloads.Message

// PutChannelPinQuery represents query parameters for PUT /channels/{channel.id}/pins/{message.id}
type PutChannelPinQuery struct {
	// No query parameters for this endpoint
}

// DeleteChannelPinQuery represents query parameters for DELETE /channels/{channel.id}/pins/{message.id}
type DeleteChannelPinQuery struct {
	// No query parameters for this endpoint
}

// PutChannelRecipientQuery represents query parameters for PUT /channels/{channel.id}/recipients/{user.id}
type PutChannelRecipientQuery struct {
	// No query parameters for this endpoint
}

// DeleteChannelRecipientQuery represents query parameters for DELETE /channels/{channel.id}/recipients/{user.id}
type DeleteChannelRecipientQuery struct {
	// No query parameters for this endpoint
}

// PostChannelThreadsJSONBody represents the request body for POST /channels/{channel.id}/threads
type PostChannelThreadsJSONBody struct {
	// 1-100 character thread name
	Name string `json:"name"`
	// Duration in minutes to automatically archive the thread after recent activity
	AutoArchiveDuration *payloads.ThreadAutoArchiveDuration `json:"auto_archive_duration,omitempty"`
	// The type of thread to create
	Type *payloads.ChannelType `json:"type,omitempty"`
	// Whether non-moderators can add other non-moderators to the thread
	Invitable *bool `json:"invitable,omitempty"`
	// Amount of seconds a user has to wait before sending another message (0-21600)
	RateLimitPerUser *int `json:"rate_limit_per_user,omitempty"`
}

// PostChannelThreadsResult represents the response from POST /channels/{channel.id}/threads
type PostChannelThreadsResult = payloads.ThreadChannel

// PostChannelMessageThreadsJSONBody represents the request body for POST /channels/{channel.id}/messages/{message.id}/threads
type PostChannelMessageThreadsJSONBody struct {
	// 1-100 character thread name
	Name string `json:"name"`
	// Duration in minutes to automatically archive the thread after recent activity
	AutoArchiveDuration *payloads.ThreadAutoArchiveDuration `json:"auto_archive_duration,omitempty"`
	// Amount of seconds a user has to wait before sending another message (0-21600)
	RateLimitPerUser *int `json:"rate_limit_per_user,omitempty"`
}

// PostChannelMessageThreadsResult represents the response from POST /channels/{channel.id}/messages/{message.id}/threads
type PostChannelMessageThreadsResult = payloads.ThreadChannel

// GetChannelThreadsQuery represents query parameters for GET /channels/{channel.id}/threads/archived/{type}
type GetChannelThreadsQuery struct {
	// Returns threads before this timestamp (ISO8601 timestamp)
	Before *string `json:"before,omitempty"`
	// Optional maximum number of threads to return (default 100)
	Limit *int `json:"limit,omitempty"`
}

// ThreadList represents a list of threads
type ThreadList struct {
	Threads []payloads.ThreadChannel `json:"threads"`
	Members []payloads.ThreadMember  `json:"members"`
	HasMore bool                     `json:"has_more"`
}

// GetChannelThreadsResult represents the response from GET /channels/{channel.id}/threads/archived/{type}
type GetChannelThreadsResult = ThreadList

// GetChannelJoinedPrivateArchivedThreadsQuery represents query parameters for GET /channels/{channel.id}/users/@me/threads/archived/private
type GetChannelJoinedPrivateArchivedThreadsQuery struct {
	// Returns threads before this id
	Before *discord.Snowflake `json:"before,omitempty"`
	// Optional maximum number of threads to return (default 100)
	Limit *int `json:"limit,omitempty"`
}

// GetChannelJoinedPrivateArchivedThreadsResult represents the response from GET /channels/{channel.id}/users/@me/threads/archived/private
type GetChannelJoinedPrivateArchivedThreadsResult = ThreadList
