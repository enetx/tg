package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type Answer struct {
	ctx  *Context
	text String
	opts *gotgbot.AnswerCallbackQueryOpts
}

// URL sets a URL to be opened by the user's client when the button is pressed.
func (a *Answer) URL(url String) *Answer {
	a.opts.Url = url.Std()
	return a
}

// Alert displays the answer as an alert instead of a notification.
func (a *Answer) Alert() *Answer {
	a.opts.ShowAlert = true
	return a
}

// CacheTime sets the maximum amount of time the result may be cached on Telegram's servers.
func (a *Answer) CacheTime(seconds int64) *Answer {
	a.opts.CacheTime = seconds
	return a
}

// Timeout sets a custom timeout for this request.
func (a *Answer) Timeout(duration time.Duration) *Answer {
	if a.opts.RequestOpts == nil {
		a.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	a.opts.RequestOpts.Timeout = duration

	return a
}

// APIURL sets a custom API URL for this request.
func (a *Answer) APIURL(url String) *Answer {
	if a.opts.RequestOpts == nil {
		a.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	a.opts.RequestOpts.APIURL = url.Std()

	return a
}

// Send sends the callback query answer and returns the result.
func (a *Answer) Send() Result[bool] {
	a.opts.Text = a.text.Std()
	return ResultOf(a.ctx.Update.CallbackQuery.Answer(a.ctx.Bot.Raw(), a.opts))
}
