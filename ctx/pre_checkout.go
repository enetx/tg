package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type PreCheckout struct {
	ctx  *Context
	ok   bool
	err  String
	opts *gotgbot.AnswerPreCheckoutQueryOpts
}

// Ok marks the pre-checkout query as successful.
func (c *PreCheckout) Ok() *PreCheckout {
	c.ok = true
	return c
}

// Error marks the pre-checkout query as failed with the specified error message.
func (c *PreCheckout) Error(text String) *PreCheckout {
	c.ok = false
	c.err = text

	return c
}

// Timeout sets a custom timeout for this request.
func (c *PreCheckout) Timeout(duration time.Duration) *PreCheckout {
	if c.opts == nil {
		c.opts = new(gotgbot.AnswerPreCheckoutQueryOpts)
	}

	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *PreCheckout) APIURL(url String) *PreCheckout {
	if c.opts == nil {
		c.opts = new(gotgbot.AnswerPreCheckoutQueryOpts)
	}

	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send answers the pre-checkout query and returns the result.
func (c *PreCheckout) Send() Result[bool] {
	query := c.ctx.Update.PreCheckoutQuery
	if query == nil {
		return Err[bool](Errorf("no precheckout query"))
	}

	if !c.ok {
		c.opts = &gotgbot.AnswerPreCheckoutQueryOpts{ErrorMessage: c.err.Std()}
	}

	return ResultOf(query.Answer(c.ctx.Bot.Raw(), c.ok, c.opts))
}
