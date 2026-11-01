package rich

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/g/ref"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/richtext"
)

// Button represents a rich message button builder.
type Button struct {
	button *gotgbot.RichMessageButton
}

// NewButton creates a new Button builder with the required text.
// The text may contain only plain text, custom emoji and date-time nodes.
func NewButton(text *richtext.Text) *Button {
	return &Button{button: &gotgbot.RichMessageButton{Text: text.Std()}}
}

// Text sets the button's display text.
func (b *Button) Text(text *richtext.Text) *Button {
	b.button.Text = text.Std()
	return b
}

// Callback sets the callback data that will be sent when the button is pressed.
func (b *Button) Callback(callback g.String) *Button {
	b.button.CallbackData = callback.Std()
	return b
}

// URL makes the button open the specified URL when pressed.
func (b *Button) URL(url g.String) *Button {
	b.button.Url = url.Std()
	return b
}

// WebApp makes the button launch a Telegram Web App at the specified URL.
func (b *Button) WebApp(url g.String) *Button {
	b.button.WebApp = &gotgbot.WebAppInfo{Url: url.Std()}
	return b
}

// LoginURL makes the button perform Telegram login via the specified URL.
func (b *Button) LoginURL(url g.String) *Button {
	b.button.LoginUrl = &gotgbot.LoginUrl{Url: url.Std()}
	return b
}

// SwitchInlineQuery makes the button open an inline query in another chat.
func (b *Button) SwitchInlineQuery(query g.String) *Button {
	b.button.SwitchInlineQuery = ref.Of(query.Std())
	return b
}

// SwitchInlineQueryCurrentChat makes the button open an inline query in the current chat.
func (b *Button) SwitchInlineQueryCurrentChat(query g.String) *Button {
	b.button.SwitchInlineQueryCurrentChat = ref.Of(query.Std())
	return b
}

// SwitchInlineQueryChosenChat makes the button open an inline query in a chat of the specified type.
func (b *Button) SwitchInlineQueryChosenChat(query gotgbot.SwitchInlineQueryChosenChat) *Button {
	b.button.SwitchInlineQueryChosenChat = &query
	return b
}

// CopyText makes the button copy the specified text to the clipboard when pressed.
func (b *Button) CopyText(text g.String) *Button {
	b.button.CopyText = &gotgbot.CopyTextButton{Text: text.Std()}
	return b
}

// Disabled makes the button do nothing when pressed.
func (b *Button) Disabled() *Button {
	b.button.Disabled = new(gotgbot.DisabledButton)
	return b
}

// Style sets the visual style of the button.
func (b *Button) Style(style keyboard.ButtonStyle) *Button {
	b.button.Style = style.String()
	return b
}

// Link shows a callback button as a regular link without borders.
func (b *Button) Link() *Button {
	b.button.Style = "link"
	return b
}

// Std returns the underlying gotgbot.RichMessageButton.
func (b *Button) Std() gotgbot.RichMessageButton {
	return *b.button
}
