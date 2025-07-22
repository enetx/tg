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
func (c *Answer) URL(url String) *Answer {
	c.opts.Url = url.Std()
	return c
}

// Alert displays the answer as an alert instead of a notification.
func (c *Answer) Alert() *Answer {
	c.opts.ShowAlert = true
	return c
}

// CacheTime sets the maximum amount of time the result may be cached on Telegram's servers.
func (c *Answer) CacheTime(seconds int64) *Answer {
	c.opts.CacheTime = seconds
	return c
}

// Timeout sets a custom timeout for this request.
func (c *Answer) Timeout(duration time.Duration) *Answer {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Answer) APIURL(url String) *Answer {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send sends the callback query answer and returns the result.
func (c *Answer) Send() Result[bool] {
	c.opts.Text = c.text.Std()
	return ResultOf(c.ctx.Update.CallbackQuery.Answer(c.ctx.Bot.Raw(), c.opts))
}
