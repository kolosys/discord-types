package payloads

import "github.com/kolosys/discord-types/discord"

// GuildScheduledEvent represents a Discord guild scheduled event.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object
type GuildScheduledEvent struct {
	// ID is the id of the guild event.
	ID discord.Snowflake `json:"id"`

	// GuildID is the guild id which the scheduled event belongs to.
	GuildID discord.Snowflake `json:"guild_id"`

	// ChannelID is the channel id in which the scheduled event will be hosted, or null if entity type is EXTERNAL.
	ChannelID *discord.Snowflake `json:"channel_id"`

	// CreatorID is the id of the user that created the scheduled event.
	CreatorID *discord.Snowflake `json:"creator_id,omitempty"`

	// Name is the name of the scheduled event.
	Name string `json:"name"`

	// Description is the description of the scheduled event.
	Description *string `json:"description,omitempty"`

	// ScheduledStartTime is the time the scheduled event will start.
	ScheduledStartTime string `json:"scheduled_start_time"`

	// ScheduledEndTime is the time at which the guild event will end, or null if the event does not have a scheduled time to end.
	ScheduledEndTime *string `json:"scheduled_end_time"`

	// PrivacyLevel is the privacy level of the scheduled event.
	PrivacyLevel GuildScheduledEventPrivacyLevel `json:"privacy_level"`

	// Status is the status of the scheduled event.
	Status GuildScheduledEventStatus `json:"status"`

	// EntityType is the type of hosting entity associated with the scheduled event.
	EntityType GuildScheduledEventEntityType `json:"entity_type"`

	// EntityID is the id of the hosting entity associated with the scheduled event.
	EntityID *discord.Snowflake `json:"entity_id"`

	// EntityMetadata is the entity metadata for the scheduled event.
	EntityMetadata *GuildScheduledEventEntityMetadata `json:"entity_metadata"`

	// Creator is the user that created the scheduled event.
	Creator *User `json:"creator,omitempty"`

	// UserCount is the number of users subscribed to the scheduled event.
	UserCount *int `json:"user_count,omitempty"`

	// Image is the cover image of the scheduled event.
	Image *string `json:"image,omitempty"`

	// RecurrenceRule is the definition for how often this event should recur.
	RecurrenceRule *GuildScheduledEventRecurrenceRule `json:"recurrence_rule"`
}

// GuildScheduledEventPrivacyLevel represents the privacy level of a guild scheduled event.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-privacy-level
type GuildScheduledEventPrivacyLevel int

// Guild scheduled event privacy level constants
const (
	// GuildScheduledEventPrivacyLevelGuildOnly indicates the scheduled event is only accessible to guild members.
	GuildScheduledEventPrivacyLevelGuildOnly GuildScheduledEventPrivacyLevel = 2
)

// GuildScheduledEventStatus represents the status of a guild scheduled event.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-status
type GuildScheduledEventStatus int

// Guild scheduled event status constants
const (
	// GuildScheduledEventStatusScheduled indicates the event is scheduled.
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = iota + 1

	// GuildScheduledEventStatusActive indicates the event is active.
	GuildScheduledEventStatusActive

	// GuildScheduledEventStatusCompleted indicates the event is completed.
	GuildScheduledEventStatusCompleted

	// GuildScheduledEventStatusCanceled indicates the event is canceled.
	GuildScheduledEventStatusCanceled
)

// GuildScheduledEventEntityType represents the entity type of a guild scheduled event.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-entity-types
type GuildScheduledEventEntityType int

// Guild scheduled event entity type constants
const (
	// GuildScheduledEventEntityTypeStageInstance indicates the event is a stage instance.
	GuildScheduledEventEntityTypeStageInstance GuildScheduledEventEntityType = iota + 1

	// GuildScheduledEventEntityTypeVoice indicates the event is a voice channel.
	GuildScheduledEventEntityTypeVoice

	// GuildScheduledEventEntityTypeExternal indicates the event is external.
	GuildScheduledEventEntityTypeExternal
)

// GuildScheduledEventEntityMetadata represents the entity metadata for a guild scheduled event.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-entity-metadata
type GuildScheduledEventEntityMetadata struct {
	// Location is the location of the event. Required for events with entity_type EXTERNAL.
	Location *string `json:"location,omitempty"`
}

// GuildScheduledEventRecurrenceRule represents the recurrence rule for a guild scheduled event.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-recurrence-rule-object-guild-scheduled-event-recurrence-rule-structure
type GuildScheduledEventRecurrenceRule struct {
	// Start is the starting time of the recurrence interval.
	Start string `json:"start"`

	// End is the ending time of the recurrence interval.
	End *string `json:"end"`

	// Frequency is how often the event occurs.
	Frequency GuildScheduledEventRecurrenceRuleFrequency `json:"frequency"`

	// Interval is the spacing between the events, defined by frequency.
	// For example, frequency of Weekly and an interval of 2 would be "every-other week".
	Interval int `json:"interval"`

	// ByWeekday is the set of specific days within a week for the event to recur on.
	ByWeekday []GuildScheduledEventRecurrenceRuleWeekday `json:"by_weekday,omitempty"`

	// ByNWeekday is the set of specific days within an n-week interval for the event to recur on.
	ByNWeekday []GuildScheduledEventRecurrenceRuleNWeekday `json:"by_n_weekday,omitempty"`

	// ByMonth is the set of specific months for the event to recur on.
	ByMonth []GuildScheduledEventRecurrenceRuleMonth `json:"by_month,omitempty"`

	// ByMonthDay is the set of specific dates within a month for the event to recur on.
	ByMonthDay []int `json:"by_month_day,omitempty"`

	// ByYearDay is the set of days within a year for the event to recur on.
	ByYearDay []int `json:"by_year_day,omitempty"`

	// Count is the total amount of times that the event is allowed to recur before stopping.
	Count *int `json:"count,omitempty"`
}

// GuildScheduledEventRecurrenceRuleFrequency represents how often a guild scheduled event occurs.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-recurrence-rule-object-guild-scheduled-event-recurrence-rule-frequency
type GuildScheduledEventRecurrenceRuleFrequency int

// Guild scheduled event recurrence rule frequency constants
const (
	// GuildScheduledEventRecurrenceRuleFrequencyYearly indicates the event occurs yearly.
	GuildScheduledEventRecurrenceRuleFrequencyYearly GuildScheduledEventRecurrenceRuleFrequency = iota

	// GuildScheduledEventRecurrenceRuleFrequencyMonthly indicates the event occurs monthly.
	GuildScheduledEventRecurrenceRuleFrequencyMonthly

	// GuildScheduledEventRecurrenceRuleFrequencyWeekly indicates the event occurs weekly.
	GuildScheduledEventRecurrenceRuleFrequencyWeekly

	// GuildScheduledEventRecurrenceRuleFrequencyDaily indicates the event occurs daily.
	GuildScheduledEventRecurrenceRuleFrequencyDaily
)

// GuildScheduledEventRecurrenceRuleWeekday represents a weekday for guild scheduled event recurrence.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-recurrence-rule-object-guild-scheduled-event-recurrence-rule-weekday
type GuildScheduledEventRecurrenceRuleWeekday int

// Guild scheduled event recurrence rule weekday constants
const (
	// GuildScheduledEventRecurrenceRuleWeekdayMonday represents Monday.
	GuildScheduledEventRecurrenceRuleWeekdayMonday GuildScheduledEventRecurrenceRuleWeekday = iota

	// GuildScheduledEventRecurrenceRuleWeekdayTuesday represents Tuesday.
	GuildScheduledEventRecurrenceRuleWeekdayTuesday

	// GuildScheduledEventRecurrenceRuleWeekdayWednesday represents Wednesday.
	GuildScheduledEventRecurrenceRuleWeekdayWednesday

	// GuildScheduledEventRecurrenceRuleWeekdayThursday represents Thursday.
	GuildScheduledEventRecurrenceRuleWeekdayThursday

	// GuildScheduledEventRecurrenceRuleWeekdayFriday represents Friday.
	GuildScheduledEventRecurrenceRuleWeekdayFriday

	// GuildScheduledEventRecurrenceRuleWeekdaySaturday represents Saturday.
	GuildScheduledEventRecurrenceRuleWeekdaySaturday

	// GuildScheduledEventRecurrenceRuleWeekdaySunday represents Sunday.
	GuildScheduledEventRecurrenceRuleWeekdaySunday
)

// GuildScheduledEventRecurrenceRuleNWeekday represents an n-weekday for guild scheduled event recurrence.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-recurrence-rule-object-guild-scheduled-event-recurrence-rule-n-weekday
type GuildScheduledEventRecurrenceRuleNWeekday struct {
	// N represents which week of the month the weekday occurs in.
	N int `json:"n"`

	// Day represents the day of the week.
	Day GuildScheduledEventRecurrenceRuleWeekday `json:"day"`
}

// GuildScheduledEventRecurrenceRuleMonth represents a month for guild scheduled event recurrence.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-recurrence-rule-object-guild-scheduled-event-recurrence-rule-month
type GuildScheduledEventRecurrenceRuleMonth int

// Guild scheduled event recurrence rule month constants
const (
	// GuildScheduledEventRecurrenceRuleMonthJanuary represents January.
	GuildScheduledEventRecurrenceRuleMonthJanuary GuildScheduledEventRecurrenceRuleMonth = iota + 1

	// GuildScheduledEventRecurrenceRuleMonthFebruary represents February.
	GuildScheduledEventRecurrenceRuleMonthFebruary

	// GuildScheduledEventRecurrenceRuleMonthMarch represents March.
	GuildScheduledEventRecurrenceRuleMonthMarch

	// GuildScheduledEventRecurrenceRuleMonthApril represents April.
	GuildScheduledEventRecurrenceRuleMonthApril

	// GuildScheduledEventRecurrenceRuleMonthMay represents May.
	GuildScheduledEventRecurrenceRuleMonthMay

	// GuildScheduledEventRecurrenceRuleMonthJune represents June.
	GuildScheduledEventRecurrenceRuleMonthJune

	// GuildScheduledEventRecurrenceRuleMonthJuly represents July.
	GuildScheduledEventRecurrenceRuleMonthJuly

	// GuildScheduledEventRecurrenceRuleMonthAugust represents August.
	GuildScheduledEventRecurrenceRuleMonthAugust

	// GuildScheduledEventRecurrenceRuleMonthSeptember represents September.
	GuildScheduledEventRecurrenceRuleMonthSeptember

	// GuildScheduledEventRecurrenceRuleMonthOctober represents October.
	GuildScheduledEventRecurrenceRuleMonthOctober

	// GuildScheduledEventRecurrenceRuleMonthNovember represents November.
	GuildScheduledEventRecurrenceRuleMonthNovember

	// GuildScheduledEventRecurrenceRuleMonthDecember represents December.
	GuildScheduledEventRecurrenceRuleMonthDecember
)

// GuildScheduledEventUser represents a user in a guild scheduled event.
//
// See: https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-user-object
type GuildScheduledEventUser struct {
	// GuildScheduledEventID is the scheduled event id which the user subscribed to.
	GuildScheduledEventID discord.Snowflake `json:"guild_scheduled_event_id"`

	// User is the user which subscribed to an event.
	User User `json:"user"`

	// Member is the guild member data for this user for the guild which this event belongs to, if any.
	Member *GuildMember `json:"member,omitempty"`
}
