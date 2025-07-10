package keyboard

import (
	"unsafe"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type ReplyKeyboard struct {
	rows Slice[Slice[gotgbot.KeyboardButton]]
}

func (rb *ReplyKeyboard) Row() *ReplyKeyboard {
	rb.rows.Push([]gotgbot.KeyboardButton{})
	return rb
}

func (rb *ReplyKeyboard) addToLastRow(btn gotgbot.KeyboardButton) *ReplyKeyboard {
	if rb.rows.Empty() {
		rb.rows.Push([]gotgbot.KeyboardButton{btn})
	} else {
		rb.rows[len(rb.rows)-1].Push(btn)
	}

	return rb
}

// Basic text button
func (rb *ReplyKeyboard) Text(text String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{Text: text.Std()})
}

// Request contact button
func (rb *ReplyKeyboard) Contact(text String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:           text.Std(),
		RequestContact: true,
	})
}

// Request location button
func (rb *ReplyKeyboard) Location(text String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:            text.Std(),
		RequestLocation: true,
	})
}

// Launch WebApp
func (rb *ReplyKeyboard) WebApp(text, url String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:   text.Std(),
		WebApp: &gotgbot.WebAppInfo{Url: url.Std()},
	})
}

// Request a poll from user
func (rb *ReplyKeyboard) Poll(text String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:        text.Std(),
		RequestPoll: new(gotgbot.KeyboardButtonPollType),
	})
}

// Request chat
func (rb *ReplyKeyboard) Chat(text String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:        text.Std(),
		RequestChat: new(gotgbot.KeyboardButtonRequestChat),
	})
}

// Request users
func (rb *ReplyKeyboard) Users(text String) *ReplyKeyboard {
	return rb.addToLastRow(gotgbot.KeyboardButton{
		Text:         text.Std(),
		RequestUsers: new(gotgbot.KeyboardButtonRequestUsers),
	})
}

// Final builder
func (rb *ReplyKeyboard) Markup() gotgbot.ReplyMarkup {
	keyboard := *(*([][]gotgbot.KeyboardButton))(unsafe.Pointer(&rb.rows))
	return gotgbot.ReplyKeyboardMarkup{
		Keyboard:       keyboard,
		ResizeKeyboard: true,
	}
}
