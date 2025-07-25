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
func (acq *AnswerCallbackQuery) URL(url String) *AnswerCallbackQuery {
	acq.opts.Url = url.Std()
	return acq
}

// Alert displays the answer as an alert instead of a notification.
func (acq *AnswerCallbackQuery) Alert() *AnswerCallbackQuery {
	acq.opts.ShowAlert = true
	return acq
}

// CacheTime sets the maximum amount of time the result may be cached on Telegram's servers.
func (acq *AnswerCallbackQuery) CacheFor(duration time.Duration) *AnswerCallbackQuery {
	acq.opts.CacheTime = int64(duration.Seconds())
	return acq
}

// Timeout sets a custom timeout for this request.
func (acq *AnswerCallbackQuery) Timeout(duration time.Duration) *AnswerCallbackQuery {
	if acq.opts.RequestOpts == nil {
		acq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	acq.opts.RequestOpts.Timeout = duration

	return acq
}

// APIURL sets a custom API URL for this request.
func (acq *AnswerCallbackQuery) APIURL(url String) *AnswerCallbackQuery {
	if acq.opts.RequestOpts == nil {
		acq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	acq.opts.RequestOpts.APIURL = url.Std()

	return acq
}

// Send sends the callback query answer and returns the result.
func (acq *AnswerCallbackQuery) Send() Result[bool] {
	acq.opts.Text = acq.text.Std()
	return ResultOf(acq.ctx.Update.CallbackQuery.Answer(acq.ctx.Bot.Raw(), acq.opts))
}
