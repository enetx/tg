package keyboard

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

// KeyboardBuilder represents any type that can produce a gotgbot.ReplyMarkup.
type KeyboardBuilder interface{ Markup() gotgbot.ReplyMarkup }

// Compile-time interface checks
var (
	_ KeyboardBuilder = (*InlineKeyboard)(nil)
	_ KeyboardBuilder = (*ReplyKeyboard)(nil)
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

// Reply returns a new reply keyboard builder.
func Reply() *ReplyKeyboard { return new(ReplyKeyboard) }
