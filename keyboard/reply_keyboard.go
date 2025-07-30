package keyboard

import (
	"unsafe"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// ReplyKeyboard helps build Telegram reply keyboard markup using a fluent API.
type ReplyKeyboard struct {
	rows g.Slice[g.Slice[gotgbot.KeyboardButton]]
}

// Row starts a new row for subsequent buttons.
func (rb *ReplyKeyboard) Row() *ReplyKeyboard {
	rb.rows.Push([]gotgbot.KeyboardButton{})
	return rb
}

// addToLastRow adds a button to the last row, creating a new row if needed.
func (rb *ReplyKeyboard) addToLastRow(btn gotgbot.KeyboardButton) *ReplyKeyboard {
	if rb.rows.Empty() {
		rb.rows.Push([]gotgbot.KeyboardButton{btn})
	} else {
		rb.rows[len(rb.rows)-1].Push(btn)
	}

	return rb
}

// Text adds a basic text button to the current row.
func (rb *ReplyKeyboard) Text(text g.String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{Text: text.Std()})
}

// Contact adds a button that requests the user's contact information.
func (rb *ReplyKeyboard) Contact(text g.String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:           text.Std(),
		RequestContact: true,
	})
}

// Location adds a button that requests the user's location.
func (rb *ReplyKeyboard) Location(text g.String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:            text.Std(),
		RequestLocation: true,
	})
}

// WebApp adds a button that launches a Telegram Web App at the specified URL.
func (rb *ReplyKeyboard) WebApp(text, url g.String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:   text.Std(),
		WebApp: &gotgbot.WebAppInfo{Url: url.Std()},
	})
}

// Poll adds a button that requests the user to create a poll.
func (rb *ReplyKeyboard) Poll(text g.String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:        text.Std(),
		RequestPoll: new(gotgbot.KeyboardButtonPollType),
	})
}

// Chat adds a button that requests the user to select a chat.
func (rb *ReplyKeyboard) Chat(text g.String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:        text.Std(),
		RequestChat: new(gotgbot.KeyboardButtonRequestChat),
	})
}

// Users adds a button that requests the user to select users.
func (rb *ReplyKeyboard) Users(text g.String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:         text.Std(),
		RequestUsers: new(gotgbot.KeyboardButtonRequestUsers),
	})
}

// Markup returns the final ReplyKeyboardMarkup structure.
func (rb *ReplyKeyboard) Markup() gotgbot.ReplyMarkup {
	keyboard := *(*([][]gotgbot.KeyboardButton))(unsafe.Pointer(&rb.rows))
	return gotgbot.ReplyKeyboardMarkup{
		Keyboard:       keyboard,
		ResizeKeyboard: true,
	}
}
