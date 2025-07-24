package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type AnswerCallbackQuery struct {
	ctx  *Context
	text String
	opts *gotgbot.AnswerCallbackQueryOpts
}

// URL sets a URL to be opened by the user's client when the button is pressed.
func (a *AnswerCallbackQuery) URL(url String) *AnswerCallbackQuery {
	a.opts.Url = url.Std()
	return a
}

// Alert displays the answer as an alert instead of a notification.
func (a *AnswerCallbackQuery) Alert() *AnswerCallbackQuery {
	a.opts.ShowAlert = true
	return a
}

// CacheTime sets the maximum amount of time the result may be cached on Telegram's servers.
func (a *AnswerCallbackQuery) CacheTime(seconds int64) *AnswerCallbackQuery {
	a.opts.CacheTime = seconds
	return a
}

// Timeout sets a custom timeout for this request.
func (a *AnswerCallbackQuery) Timeout(duration time.Duration) *AnswerCallbackQuery {
	if a.opts.RequestOpts == nil {
		a.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	a.opts.RequestOpts.Timeout = duration

	return a
}

// APIURL sets a custom API URL for this request.
func (a *AnswerCallbackQuery) APIURL(url String) *AnswerCallbackQuery {
	if a.opts.RequestOpts == nil {
		a.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	a.opts.RequestOpts.APIURL = url.Std()

	return a
}

// Send sends the callback query answer and returns the result.
func (a *AnswerCallbackQuery) Send() Result[bool] {
	a.opts.Text = a.text.Std()
	return ResultOf(a.ctx.Update.CallbackQuery.Answer(a.ctx.Bot.Raw(), a.opts))
}
