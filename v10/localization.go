package v10

// LocalizationMap represents a mapping of locales to localized strings.
// Reference: https://discord.com/developers/docs/reference#locales
type LocalizationMap map[Locale]*string

// Locale represents a Discord locale.
type Locale string

// Discord supported locales
const (
	LocaleDanish        Locale = "da"
	LocaleGerman        Locale = "de"
	LocaleEnglishGB     Locale = "en-GB"
	LocaleEnglishUS     Locale = "en-US"
	LocaleSpanish       Locale = "es-ES"
	LocaleFrench        Locale = "fr"
	LocaleCroatian      Locale = "hr"
	LocaleItalian       Locale = "it"
	LocaleLithuanian    Locale = "lt"
	LocaleHungarian     Locale = "hu"
	LocaleDutch         Locale = "nl"
	LocaleNorwegian     Locale = "no"
	LocalePolish        Locale = "pl"
	LocalePortugueseBR  Locale = "pt-BR"
	LocaleRomanian      Locale = "ro"
	LocaleFinnish       Locale = "fi"
	LocaleSwedish       Locale = "sv-SE"
	LocaleVietnamese    Locale = "vi"
	LocaleTurkish       Locale = "tr"
	LocaleCzech         Locale = "cs"
	LocaleGreek         Locale = "el"
	LocaleBulgarian     Locale = "bg"
	LocaleRussian       Locale = "ru"
	LocaleUkrainian     Locale = "uk"
	LocaleHindi         Locale = "hi"
	LocaleThai          Locale = "th"
	LocaleChineseChina  Locale = "zh-CN"
	LocaleJapanese      Locale = "ja"
	LocaleChineseTaiwan Locale = "zh-TW"
	LocaleKorean        Locale = "ko"
)
