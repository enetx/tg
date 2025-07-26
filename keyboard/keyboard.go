package keyboard

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

// Keyboard represents any type that can produce a gotgbot.ReplyMarkup.
type Keyboard interface{ Markup() gotgbot.ReplyMarkup }

// Compile-time interface checks
var (
	_ Keyboard = (*InlineKeyboard)(nil)
	_ Keyboard = (*ReplyKeyboard)(nil)
)

// Inline creates a new InlineKeyboard builder.
// Optionally accepts one of:
//   - *InlineKeyboard (clones its markup)
//   - *gotgbot.InlineKeyboardMarkup (imports raw markup)
func Inline(from ...any) *InlineKeyboard {
	ik := new(InlineKeyboard)
	if len(from) == 0 || from[0] == nil {
		return ik
	}

	switch v := from[0].(type) {
	case *InlineKeyboard:
		return ik.fromKeyboard(v)
	case *gotgbot.InlineKeyboardMarkup:
		return ik.fromMarkup(v)
	case gotgbot.InlineKeyboardMarkup:
		return ik.fromMarkup(&v)
	default:
		return ik
	}
}

// Reply creates a new ReplyKeyboard builder for creating custom reply keyboards.
func Reply() *ReplyKeyboard { return new(ReplyKeyboard) }
