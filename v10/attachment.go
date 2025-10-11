package v10

// Attachment represents a message attachment
// https://discord.com/developers/docs/resources/channel#attachment-object-attachment-structure
type Attachment struct {
	ID           Snowflake        `json:"id"`
	Filename     string           `json:"filename"`
	Description  *string          `json:"description,omitempty"`
	ContentType  *string          `json:"content_type,omitempty"`
	Size         int              `json:"size"`
	URL          string           `json:"url"`
	ProxyURL     string           `json:"proxy_url"`
	Height       *int             `json:"height,omitempty"`
	Width        *int             `json:"width,omitempty"`
	Ephemeral    *bool            `json:"ephemeral,omitempty"`
	DurationSecs *float64         `json:"duration_secs,omitempty"`
	Waveform     *string          `json:"waveform,omitempty"`
	Flags        *AttachmentFlags `json:"flags,omitempty"`
}

// AttachmentFlags represents attachment flags
// https://discord.com/developers/docs/resources/channel#attachment-object-attachment-flags
type AttachmentFlags int

const (
	AttachmentFlagIsRemix AttachmentFlags = 1 << 2
)

// PartialAttachment represents a partial attachment for uploading
// https://discord.com/developers/docs/resources/channel#attachment-object
type PartialAttachment struct {
	ID          Snowflake `json:"id"`
	Filename    *string   `json:"filename,omitempty"`
	Description *string   `json:"description,omitempty"`
}
