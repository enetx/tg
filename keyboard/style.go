package keyboard

import "github.com/PaulSonOfLars/gotgbot/v2"

// ButtonStyle represents the visual style of an inline keyboard button.
//
// See: https://core.telegram.org/bots/api#inlinekeyboardbutton
type ButtonStyle int

const (
	// ButtonStyleDefault is the default button style (no explicit style).
	ButtonStyleDefault ButtonStyle = iota
	// ButtonStyleDanger applies a red/danger style to the button.
	ButtonStyleDanger
	// ButtonStyleSuccess applies a green/success style to the button.
	ButtonStyleSuccess
	// ButtonStylePrimary applies a primary/accent style to the button.
	ButtonStylePrimary
)

// String returns the Telegram Bot API string value for the button style.
func (s ButtonStyle) String() string {
	switch s {
	case ButtonStyleDanger:
		return gotgbot.KeyboardButtonStyleDanger
	case ButtonStyleSuccess:
		return gotgbot.KeyboardButtonStyleSuccess
	case ButtonStylePrimary:
		return gotgbot.KeyboardButtonStylePrimary
	default:
		return ""
	}
}
