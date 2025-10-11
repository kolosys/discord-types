package v10

// StickerType represents the type of sticker
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-types
type StickerType int

const (
	StickerTypeStandard StickerType = iota + 1
	StickerTypeGuild
)

// StickerFormatType represents the format type of sticker
// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-format-types
type StickerFormatType int

const (
	StickerFormatTypePNG StickerFormatType = iota + 1
	StickerFormatTypeAPNG
	StickerFormatTypeLottie
	StickerFormatTypeGIF
)

// Sticker represents a Discord sticker
// https://discord.com/developers/docs/resources/sticker#sticker-object
type Sticker struct {
	ID          Snowflake         `json:"id"`
	PackID      *Snowflake        `json:"pack_id,omitempty"`
	Name        string            `json:"name"`
	Description *string           `json:"description"`
	Tags        string            `json:"tags"`
	Asset       *string           `json:"asset,omitempty"` // Deprecated in v11
	Type        StickerType       `json:"type"`
	FormatType  StickerFormatType `json:"format_type"`
	Available   *bool             `json:"available,omitempty"`
	GuildID     *Snowflake        `json:"guild_id,omitempty"`
	User        *User             `json:"user,omitempty"`
	SortValue   *int              `json:"sort_value,omitempty"`
}

// StickerItem represents a sticker item (partial sticker)
// https://discord.com/developers/docs/resources/sticker#sticker-item-object
type StickerItem struct {
	ID         Snowflake         `json:"id"`
	Name       string            `json:"name"`
	FormatType StickerFormatType `json:"format_type"`
}

// StickerPack represents a sticker pack
// https://discord.com/developers/docs/resources/sticker#sticker-pack-object
type StickerPack struct {
	ID             Snowflake  `json:"id"`
	Stickers       []Sticker  `json:"stickers"`
	Name           string     `json:"name"`
	SKUID          Snowflake  `json:"sku_id"`
	CoverStickerID *Snowflake `json:"cover_sticker_id,omitempty"`
	Description    string     `json:"description"`
	BannerAssetID  *Snowflake `json:"banner_asset_id,omitempty"`
}
