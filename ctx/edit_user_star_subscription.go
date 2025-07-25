package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// EditUserStarSubscription is a request builder for editing user star subscriptions.
type EditUserStarSubscription struct {
	ctx                     *Context
	userID                  int64
	telegramPaymentChargeID String
	isCanceled              bool
	opts                    *gotgbot.EditUserStarSubscriptionOpts
}

// Timeout sets a custom timeout for this request.
func (c *EditUserStarSubscription) Timeout(duration time.Duration) *EditUserStarSubscription {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *EditUserStarSubscription) APIURL(url String) *EditUserStarSubscription {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send executes the EditUserStarSubscription request.
func (c *EditUserStarSubscription) Send() Result[bool] {
	return ResultOf(c.ctx.Bot.Raw().EditUserStarSubscription(
		c.userID,
		c.telegramPaymentChargeID.Std(),
		c.isCanceled,
		c.opts,
	))
}
