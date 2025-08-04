package keyboard

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/g/ref"
)

// Button represents an inline keyboard button with support for toggle functionality and fluent configuration.
type Button struct {
	raw      *gotgbot.InlineKeyboardButton
	on       g.String
	off      g.String
	isToggle bool
	isActive bool
	deleted  bool
	parent   *InlineKeyboard

	Get buttonGetter
}

type buttonGetter struct{ b *Button }

// NewButton creates a new Button instance, optionally initializing it with an existing InlineKeyboardButton.
func NewButton(ikb ...*gotgbot.InlineKeyboardButton) *Button {
	raw := new(gotgbot.InlineKeyboardButton)
	if len(ikb) > 0 && ikb[0] != nil {
		raw = ikb[0]
	}

	b := &Button{raw: raw}
	b.Get = buttonGetter{b}

	return b
}

// attach connects the button to its parent keyboard for state synchronization.
func (b *Button) attach(kb *InlineKeyboard) *Button {
	b.parent = kb
	return b
}

// Text sets the button's display text.
func (b *Button) Text(text g.String) *Button {
	b.raw.Text = text.Std()
	return b
}

// Callback sets the callback data that will be sent when the button is pressed.
func (b *Button) Callback(callback g.String) *Button {
	b.raw.CallbackData = callback.Std()
	return b
}

// URL makes the button open the specified URL when pressed.
func (b *Button) URL(url g.String) *Button {
	b.raw.Url = url.Std()
	return b
}

// WebApp makes the button launch a Telegram Web App at the specified URL.
func (b *Button) WebApp(url g.String) *Button {
	b.raw.WebApp = &gotgbot.WebAppInfo{Url: url.Std()}
	return b
}

// LoginURL makes the button perform Telegram login via the specified URL.
func (b *Button) LoginURL(url g.String) *Button {
	b.raw.LoginUrl = &gotgbot.LoginUrl{Url: url.Std()}
	return b
}

// CopyText makes the button copy the specified text to the user's clipboard when pressed.
func (b *Button) CopyText(toCopy g.String) *Button {
	b.raw.CopyText = &gotgbot.CopyTextButton{Text: toCopy.Std()}
	return b
}

// Pay makes the button trigger a payment flow when pressed.
func (b *Button) Pay() *Button {
	b.raw.Pay = true
	return b
}

// Game makes the button launch a Telegram game when pressed.
func (b *Button) Game() *Button {
	b.raw.CallbackGame = new(gotgbot.CallbackGame)
	return b
}

// SwitchInlineQuery makes the button open an inline query with the specified query string in another chat.
func (b *Button) SwitchInlineQuery(query g.String) *Button {
	b.raw.SwitchInlineQuery = ref.Of(query.Std())
	return b
}

// SwitchInlineQueryCurrentChat makes the button open an inline query with the specified query string in the current chat.
func (b *Button) SwitchInlineQueryCurrentChat(query g.String) *Button {
	b.raw.SwitchInlineQueryCurrentChat = ref.Of(query.Std())
	return b
}

// Delete marks the button for deletion from the keyboard.
func (b *Button) Delete() {
	b.deleted = true
}

// Toggle-related methods

// On sets the text to display when the toggle button is in the active state.
func (b *Button) On(text g.String) *Button {
	b.isToggle = true
	b.on = text
	return b
}

// Off sets the text to display when the toggle button is in the inactive state.
func (b *Button) Off(text g.String) *Button {
	b.isToggle = true
	b.off = text

	return b
}

// SetActive sets the toggle button's active state and updates its parent keyboard if attached.
func (b *Button) SetActive(active bool) *Button {
	b.isToggle = true
	b.isActive = active

	if b.parent != nil {
		b.parent.update(b)
	}

	return b
}

// Flip toggles the button's active state and updates its parent keyboard if attached.
func (b *Button) Flip() *Button {
	b.isActive = !b.isActive
	if b.parent != nil {
		b.parent.update(b)
	}

	return b
}

// Build constructs the final gotgbot.InlineKeyboardButton with toggle state applied.
func (b *Button) Build() gotgbot.InlineKeyboardButton {
	btn := *b.raw

	if b.isToggle {
		if b.isActive {
			btn.Text = b.on.Std()
		} else {
			btn.Text = b.off.Std()
		}
	}

	return btn
}

// Getters
// Callback returns the button's callback data.
func (bg buttonGetter) Callback() g.String {
	return g.String(bg.b.raw.CallbackData)
}

// Text returns the button's display text.
func (bg buttonGetter) Text() g.String {
	return g.String(bg.b.raw.Text)
}

// URL returns the button's URL.
func (bg buttonGetter) URL() g.String {
	return g.String(bg.b.raw.Url)
}

// IsToggle returns true if the button is configured as a toggle button.
func (bg buttonGetter) IsToggle() bool {
	return bg.b.isToggle
}

// IsActive returns the toggle button's current active state.
func (bg buttonGetter) IsActive() bool {
	return bg.b.isActive
}
