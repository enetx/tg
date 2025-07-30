// Package entity defines the entry (message-entity) types used by Telegram.
package entity

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
)

// g.String returns the canonical string representation used by Telegram Bot API.
func (e EntityType) String() string {
	switch e {
	case Mention:
		return "mention"
	case Hashtag:
		return "hashtag"
	case Cashtag:
		return "cashtag"
	case BotCommand:
		return "bot_command"
	case URL:
		return "url"
	case Email:
		return "email"
	case PhoneNumber:
		return "phone_number"
	case Bold:
		return "bold"
	case Italic:
		return "italic"
	case Underline:
		return "underline"
	case Strikethrough:
		return "strikethrough"
	case Spoiler:
		return "spoiler"
	case Blockquote:
		return "blockquote"
	case ExpandableBlockquote:
		return "expandable_blockquote"
	case Code:
		return "code"
	case Pre:
		return "pre"
	case TextLink:
		return "text_link"
	case TextMention:
		return "text_mention"
	case CustomEmoji:
		return "custom_emoji"
	default:
		return "unknown"
	}
}
