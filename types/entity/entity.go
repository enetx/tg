// Package entity defines the entry (message-entity) types used by Telegram.
package entity

import "github.com/PaulSonOfLars/gotgbot/v2"

// EntityType enumerates all supported Telegram message-entity types.
//
// See: https://core.telegram.org/bots/api#messageentity
type EntityType int

const (
	// Plain-text entities
	Mention     EntityType = iota // @username
	Hashtag                       // #hashtag or #hashtag@chatusername
	Cashtag                       // $USD or $USD@chatusername
	BotCommand                    // /start@jobs_bot
	URL                           // https://telegram.org
	Email                         // do-not-reply@telegram.org
	PhoneNumber                   // +1-212-555-0123

	// Text-style entities
	Bold          // <b>bold</b>
	Italic        // <i>italic</i>
	Underline     // <u>underline</u>
	Strikethrough // <s>strikethrough</s>
	Spoiler       // <span class="tg-spoiler">spoiler</span>

	// Quote / block entities
	Blockquote           // > quoted text
	ExpandableBlockquote // collapsed-by-default quotation

	// Code entities
	Code // `monowidth string`
	Pre  // ```monowidth block```

	// Link entities
	TextLink    // clickable text URLs
	TextMention // users without usernames

	// Special
	CustomEmoji // inline custom emoji stickers
	DateTime    // date and time values rendered in the user's locale
)

// String returns the canonical string representation used by Telegram Bot API.
func (e EntityType) String() string {
	switch e {
	case Mention:
		return gotgbot.MessageEntityTypeMention
	case Hashtag:
		return gotgbot.MessageEntityTypeHashtag
	case Cashtag:
		return gotgbot.MessageEntityTypeCashtag
	case BotCommand:
		return gotgbot.MessageEntityTypeBotCommand
	case URL:
		return gotgbot.MessageEntityTypeUrl
	case Email:
		return gotgbot.MessageEntityTypeEmail
	case PhoneNumber:
		return gotgbot.MessageEntityTypePhoneNumber
	case Bold:
		return gotgbot.MessageEntityTypeBold
	case Italic:
		return gotgbot.MessageEntityTypeItalic
	case Underline:
		return gotgbot.MessageEntityTypeUnderline
	case Strikethrough:
		return gotgbot.MessageEntityTypeStrikethrough
	case Spoiler:
		return gotgbot.MessageEntityTypeSpoiler
	case Blockquote:
		return gotgbot.MessageEntityTypeBlockquote
	case ExpandableBlockquote:
		return gotgbot.MessageEntityTypeExpandableBlockquote
	case Code:
		return gotgbot.MessageEntityTypeCode
	case Pre:
		return gotgbot.MessageEntityTypePre
	case TextLink:
		return gotgbot.MessageEntityTypeTextLink
	case TextMention:
		return gotgbot.MessageEntityTypeTextMention
	case CustomEmoji:
		return gotgbot.MessageEntityTypeCustomEmoji
	case DateTime:
		return gotgbot.MessageEntityTypeDateTime
	default:
		return "unknown"
	}
}
