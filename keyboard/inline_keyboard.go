package keyboard

import (
	"unsafe"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/g/ref"
)

// InlineKeyboard helps build Telegram inline keyboard markup using a fluent API.
type InlineKeyboard struct {
	rows g.Slice[g.Slice[gotgbot.InlineKeyboardButton]]
}

// Row starts a new row for subsequent buttons.
func (b *InlineKeyboard) Row() *InlineKeyboard {
	b.rows.Push([]gotgbot.InlineKeyboardButton{})
	return b
}

// addToLastRow adds a button to the last row, creating a new row if needed.
func (b *InlineKeyboard) addToLastRow(btn gotgbot.InlineKeyboardButton) *InlineKeyboard {
	if b.rows.Empty() {
		b.rows.Push([]gotgbot.InlineKeyboardButton{btn})
	} else {
		b.rows[b.rows.Len()-1].Push(btn)
	}

	return b
}

// Button adds a Button to the keyboard, updating existing buttons with matching callback data.
func (b *InlineKeyboard) Button(btn *Button) *InlineKeyboard {
	if btn == nil || btn.raw == nil {
		return b
	}

	if btn.raw.Pay {
		btn.attach(b)
		return b.addToLastRow(btn.Build())
	}

	cb := btn.raw.CallbackData
	if cb == "" {
		return b
	}

	for i, row := range b.rows {
		for j := range row {
			if b.rows[i][j].CallbackData == cb {
				b.rows[i][j] = btn.Build()
				return b
			}
		}
	}

	btn.attach(b)
	return b.addToLastRow(btn.Build())
}

// update refreshes an existing button in the keyboard based on its callback data.
func (b *InlineKeyboard) update(btn *Button) *InlineKeyboard {
	if btn == nil || btn.raw == nil || btn.raw.CallbackData == "" {
		return b
	}

	cb := btn.raw.CallbackData

	for i, row := range b.rows {
		for j := range row {
			if b.rows[i][j].CallbackData == cb {
				b.rows[i][j] = btn.Build()
				return b
			}
		}
	}

	return b.addToLastRow(btn.Build())
}

// Text adds a text button with callback data to the current row.
func (b *InlineKeyboard) Text(text, callback g.String) *InlineKeyboard {
	return b.addToLastRow(gotgbot.InlineKeyboardButton{
		Text:         text.Std(),
		CallbackData: callback.Std(),
	})
}

// URL adds a button that opens a given URL.
func (b *InlineKeyboard) URL(text, url g.String) *InlineKeyboard {
	return b.addToLastRow(gotgbot.InlineKeyboardButton{
		Text: text.Std(),
		Url:  url.Std(),
	})
}

// WebApp adds a button that opens a Telegram Web App.
func (b *InlineKeyboard) WebApp(text, url g.String) *InlineKeyboard {
	return b.addToLastRow(gotgbot.InlineKeyboardButton{
		Text:   text.Std(),
		WebApp: &gotgbot.WebAppInfo{Url: url.Std()},
	})
}

// LoginURL adds a button for Telegram login via an external URL.
func (b *InlineKeyboard) LoginURL(text, url g.String) *InlineKeyboard {
	return b.addToLastRow(gotgbot.InlineKeyboardButton{
		Text:     text.Std(),
		LoginUrl: &gotgbot.LoginUrl{Url: url.Std()},
	})
}

// CopyText adds a button that copies a predefined text to the clipboard.
func (b *InlineKeyboard) CopyText(text, toCopy g.String) *InlineKeyboard {
	return b.addToLastRow(gotgbot.InlineKeyboardButton{
		Text:     text.Std(),
		CopyText: &gotgbot.CopyTextButton{Text: toCopy.Std()},
	})
}

// Pay adds a payment button. Must be used in invoices and be the first button.
func (b *InlineKeyboard) Pay(text g.String) *InlineKeyboard {
	return b.addToLastRow(gotgbot.InlineKeyboardButton{
		Text: text.Std(),
		Pay:  true,
	})
}

// Game adds a game launch button. Must be the first button in the first row.
func (b *InlineKeyboard) Game(text g.String) *InlineKeyboard {
	return b.addToLastRow(gotgbot.InlineKeyboardButton{
		Text:         text.Std(),
		CallbackGame: new(gotgbot.CallbackGame),
	})
}

// SwitchInlineQuery adds a button that opens inline query in another chat.
func (b *InlineKeyboard) SwitchInlineQuery(text, query g.String) *InlineKeyboard {
	return b.addToLastRow(gotgbot.InlineKeyboardButton{
		Text:              text.Std(),
		SwitchInlineQuery: ref.Of(query.Std()),
	})
}

// SwitchInlineQueryCurrentChat adds a button that triggers inline query in the current chat.
func (b *InlineKeyboard) SwitchInlineQueryCurrentChat(text, query g.String) *InlineKeyboard {
	return b.addToLastRow(gotgbot.InlineKeyboardButton{
		Text:                         text.Std(),
		SwitchInlineQueryCurrentChat: ref.Of(query.Std()),
	})
}

// Markup returns the final InlineKeyboardMarkup structure.
func (b *InlineKeyboard) Markup() gotgbot.ReplyMarkup {
	keyboard := *(*([][]gotgbot.InlineKeyboardButton))(unsafe.Pointer(&b.rows))
	return gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyboard}
}

// fromMarkup directly copies the raw InlineKeyboardMarkup into the builder.
// Uses unsafe.Pointer for zero-allocation conversion.
func (b *InlineKeyboard) fromMarkup(markup gotgbot.ReplyMarkup) *InlineKeyboard {
	switch m := markup.(type) {
	case *gotgbot.InlineKeyboardMarkup:
		b.rows = *(*g.Slice[g.Slice[gotgbot.InlineKeyboardButton]])(unsafe.Pointer(&m.InlineKeyboard))
	case gotgbot.InlineKeyboardMarkup:
		b.rows = *(*g.Slice[g.Slice[gotgbot.InlineKeyboardButton]])(unsafe.Pointer(&m.InlineKeyboard))
	}

	return b
}

// fromKeyboard copies another InlineKeyboard into this builder by reusing its markup.
func (b *InlineKeyboard) fromKeyboard(from *InlineKeyboard) *InlineKeyboard {
	if from == nil {
		return b
	}

	switch markup := from.Markup().(type) {
	case *gotgbot.InlineKeyboardMarkup:
		return b.fromMarkup(markup)
	case gotgbot.InlineKeyboardMarkup:
		return b.fromMarkup(&markup)
	default:
		return b
	}
}

// Edit applies a handler function to each button in the keyboard, allowing for batch modifications.
func (b *InlineKeyboard) Edit(handler func(btn *Button)) *InlineKeyboard {
	if handler == nil {
		return b
	}

	var rows g.Slice[g.Slice[gotgbot.InlineKeyboardButton]]

	for _, row := range b.rows {
		var nrow g.Slice[gotgbot.InlineKeyboardButton]

		for j := range row {
			btn := NewButton(&row[j])
			handler(btn)

			if !btn.deleted {
				nrow.Push(btn.Build())
			}
		}

		if nrow.NotEmpty() {
			rows.Push(nrow)
		}
	}

	b.rows = rows

	return b
}
