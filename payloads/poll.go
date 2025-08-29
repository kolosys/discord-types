package payloads

// Poll represents a Discord poll.
//
// See: https://discord.com/developers/docs/resources/poll#poll-object-poll-object-structure
type Poll struct {
	// Question is the question of the poll.
	Question PollMedia `json:"question"`

	// Answers are each of the answers available in the poll, up to 10.
	Answers []PollAnswer `json:"answers"`

	// Expiry is the time when the poll ends (ISO8601 timestamp).
	Expiry string `json:"expiry"`

	// AllowMultiselect indicates whether a user can select multiple answers.
	AllowMultiselect bool `json:"allow_multiselect"`

	// LayoutType is the layout type of the poll.
	LayoutType PollLayoutType `json:"layout_type"`

	// Results are the results of the poll.
	Results *PollResults `json:"results,omitempty"`
}

// BasePoll represents the base structure for creating a poll.
type BasePoll struct {
	// Question is the question of the poll.
	Question PollMedia `json:"question"`
}

// PollDefaults represents the default values for a poll.
type PollDefaults struct {
	// AllowMultiselect indicates whether a user can select multiple answers.
	// Default: false
	AllowMultiselect bool `json:"allow_multiselect"`

	// LayoutType is the layout type of the poll.
	// Default: PollLayoutTypeDefault
	LayoutType PollLayoutType `json:"layout_type"`
}

// PollLayoutType represents the layout type of a poll.
//
// See: https://discord.com/developers/docs/resources/poll#layout-type
type PollLayoutType int

// Poll layout type constants
const (
	// PollLayoutTypeDefault is the default layout type.
	PollLayoutTypeDefault PollLayoutType = 1
)

// PollMedia represents poll media content.
//
// See: https://discord.com/developers/docs/resources/poll#poll-media-object-poll-media-object-structure
type PollMedia struct {
	// Text is the text of the field.
	// The maximum length is 300 for the question, and 55 for any answer.
	Text *string `json:"text,omitempty"`

	// Emoji is the emoji of the field.
	Emoji *PartialEmoji `json:"emoji,omitempty"`
}

// BasePollAnswer represents the base structure for a poll answer.
type BasePollAnswer struct {
	// PollMedia is the data of the answer.
	PollMedia PollMedia `json:"poll_media"`
}

// PollAnswer represents a poll answer.
//
// See: https://discord.com/developers/docs/resources/poll#poll-answer-object-poll-answer-object-structure
type PollAnswer struct {
	BasePollAnswer

	// AnswerID is the ID of the answer. Starts at 1 for the first answer and goes up sequentially.
	AnswerID int `json:"answer_id"`
}

// PollResults represents poll results.
//
// See: https://discord.com/developers/docs/resources/poll#poll-results-object-poll-results-object-structure
type PollResults struct {
	// IsFinalized indicates whether the votes have been precisely counted.
	IsFinalized bool `json:"is_finalized"`

	// AnswerCounts are the counts for each answer.
	AnswerCounts []PollAnswerCount `json:"answer_counts"`
}

// PollAnswerCount represents the vote count for a poll answer.
//
// See: https://discord.com/developers/docs/resources/poll#poll-results-object-poll-answer-count-object-structure
type PollAnswerCount struct {
	// ID is the answer_id.
	ID int `json:"id"`

	// Count is the number of votes for this answer.
	Count int `json:"count"`

	// MeVoted indicates whether the current user voted for this answer.
	MeVoted bool `json:"me_voted"`
}

// PartialEmoji is defined in emoji.go
