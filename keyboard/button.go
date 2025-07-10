package keyboard

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/g/ref"
)

type Button struct {
	raw      *gotgbot.InlineKeyboardButton
	on       String
	off      String
	isToggle bool
	isActive bool
	deleted  bool
	parent   *InlineKeyboard

	Get buttonGetter
}

type buttonGetter struct{ b *Button }

func NewButton(ikb ...*gotgbot.InlineKeyboardButton) *Button {
	raw := new(gotgbot.InlineKeyboardButton)
	if len(ikb) > 0 && ikb[0] != nil {
		raw = ikb[0]
	}

	b := &Button{raw: raw}
	b.Get = buttonGetter{b}

	return b
}

func (b *Button) attach(kb *InlineKeyboard) *Button {
	b.parent = kb
	return b
}

func (b *Button) Text(text String) *Button {
	b.raw.Text = text.Std()
	return b
}

func (b *Button) Callback(callback String) *Button {
	b.raw.CallbackData = callback.Std()
	return b
}

func (b *Button) URL(url String) *Button {
	b.raw.Url = url.Std()
	return b
}

func (b *Button) WebApp(url String) *Button {
	b.raw.WebApp = &gotgbot.WebAppInfo{Url: url.Std()}
	return b
}

func (b *Button) LoginURL(url String) *Button {
	b.raw.LoginUrl = &gotgbot.LoginUrl{Url: url.Std()}
	return b
}

func (b *Button) CopyText(toCopy String) *Button {
	b.raw.CopyText = &gotgbot.CopyTextButton{Text: toCopy.Std()}
	return b
}

func (b *Button) Pay() *Button {
	b.raw.Pay = true
	return b
}

func (b *Button) Game() *Button {
	b.raw.CallbackGame = new(gotgbot.CallbackGame)
	return b
}

func (b *Button) SwitchInlineQuery(query String) *Button {
	b.raw.SwitchInlineQuery = ref.Of(query.Std())
	return b
}

func (b *Button) SwitchInlineQueryCurrentChat(query String) *Button {
	b.raw.SwitchInlineQueryCurrentChat = ref.Of(query.Std())
	return b
}

func (b *Button) Delete() {
	b.deleted = true
}

// Toggle-related methods
func (b *Button) On(text String) *Button {
	b.isToggle = true
	b.on = text
	return b
}

func (b *Button) Off(text String) *Button {
	b.isToggle = true
	b.off = text

	return b
}

func (b *Button) SetActive(active bool) *Button {
	b.isToggle = true
	b.isActive = active

	if b.parent != nil {
		b.parent.update(b)
	}

	return b
}

func (b *Button) Flip() *Button {
	b.isActive = !b.isActive
	if b.parent != nil {
		b.parent.update(b)
	}

	return b
}

func (b *Button) build() gotgbot.InlineKeyboardButton {
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
func (g buttonGetter) Callback() String {
	return String(g.b.raw.CallbackData)
}

func (g buttonGetter) Text() String {
	return String(g.b.raw.Text)
}

func (g buttonGetter) URL() String {
	return String(g.b.raw.Url)
}

func (g buttonGetter) IsToggle() bool {
	return g.b.isToggle
}

func (g buttonGetter) IsActive() bool {
	return g.b.isActive
}
