package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SavePreparedKeyboardButton represents a request to store a keyboard button that can
// be used by a user within a Mini App. The button must be of the type request_users,
// request_chat, or request_managed_bot.
type SavePreparedKeyboardButton struct {
	ctx    *Context
	userID int64
	button gotgbot.KeyboardButton
	opts   *gotgbot.SavePreparedKeyboardButtonOpts
}

// Text sets the text of the button. If only text, icon_custom_emoji_id, and style are used,
// it will be sent as a message when the button is pressed.
func (spkb *SavePreparedKeyboardButton) Text(text g.String) *SavePreparedKeyboardButton {
	spkb.button.Text = text.Std()
	return spkb
}

// RequestUsers configures the button to request the user to select users.
func (spkb *SavePreparedKeyboardButton) RequestUsers(requestID int32) *SavePreparedKeyboardButton {
	spkb.button.RequestUsers = &gotgbot.KeyboardButtonRequestUsers{RequestId: requestID}
	return spkb
}

// RequestChat configures the button to request the user to select a chat.
// Pass channel=true to request a channel chat, false to request a group/supergroup.
func (spkb *SavePreparedKeyboardButton) RequestChat(requestID int32, channel bool) *SavePreparedKeyboardButton {
	spkb.button.RequestChat = &gotgbot.KeyboardButtonRequestChat{
		RequestId:     requestID,
		ChatIsChannel: channel,
	}
	return spkb
}

// RequestManagedBot configures the button to request creation of a managed bot.
func (spkb *SavePreparedKeyboardButton) RequestManagedBot(requestID int32) *SavePreparedKeyboardButton {
	spkb.button.RequestManagedBot = &gotgbot.KeyboardButtonRequestManagedBot{RequestId: requestID}
	return spkb
}

// SuggestedName sets a suggested name for the managed bot.
// Only meaningful when RequestManagedBot has been configured.
func (spkb *SavePreparedKeyboardButton) SuggestedName(name g.String) *SavePreparedKeyboardButton {
	if spkb.button.RequestManagedBot == nil {
		spkb.button.RequestManagedBot = new(gotgbot.KeyboardButtonRequestManagedBot)
	}

	spkb.button.RequestManagedBot.SuggestedName = name.Std()

	return spkb
}

// SuggestedUsername sets a suggested username for the managed bot.
// Only meaningful when RequestManagedBot has been configured.
func (spkb *SavePreparedKeyboardButton) SuggestedUsername(username g.String) *SavePreparedKeyboardButton {
	if spkb.button.RequestManagedBot == nil {
		spkb.button.RequestManagedBot = new(gotgbot.KeyboardButtonRequestManagedBot)
	}

	spkb.button.RequestManagedBot.SuggestedUsername = username.Std()

	return spkb
}

// Timeout sets a custom timeout for this request.
func (spkb *SavePreparedKeyboardButton) Timeout(duration time.Duration) *SavePreparedKeyboardButton {
	if spkb.opts.RequestOpts == nil {
		spkb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	spkb.opts.RequestOpts.Timeout = duration

	return spkb
}

// APIURL sets a custom API URL for this request.
func (spkb *SavePreparedKeyboardButton) APIURL(url g.String) *SavePreparedKeyboardButton {
	if spkb.opts.RequestOpts == nil {
		spkb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	spkb.opts.RequestOpts.APIURL = url.Std()

	return spkb
}

// Send stores the prepared keyboard button and returns the result.
func (spkb *SavePreparedKeyboardButton) Send() g.Result[*gotgbot.PreparedKeyboardButton] {
	return g.ResultOf(spkb.ctx.Bot.Raw().SavePreparedKeyboardButton(spkb.userID, spkb.button, spkb.opts))
}
