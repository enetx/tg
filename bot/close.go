package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// Close represents a request to close the bot instance.
type Close struct {
	bot  *Bot
	opts *gotgbot.CloseOpts
}

// Timeout sets a custom timeout for this request.
func (c *Close) Timeout(duration time.Duration) *Close {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Close) APIURL(url g.String) *Close {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send closes the bot instance and returns the result.
func (c *Close) Send() g.Result[bool] {
	return g.ResultOf(c.bot.raw.Close(c.opts))
}
