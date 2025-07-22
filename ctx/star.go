package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// GetMyStarBalance is a request builder for getting bot's star balance.
type GetMyStarBalance struct {
	ctx  *Context
	opts *gotgbot.GetMyStarBalanceOpts
}

// Timeout sets a custom timeout for this request.
func (c *GetMyStarBalance) Timeout(duration time.Duration) *GetMyStarBalance {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *GetMyStarBalance) APIURL(url String) *GetMyStarBalance {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send executes the GetMyStarBalance request.
func (c *GetMyStarBalance) Send() Result[*gotgbot.StarAmount] {
	return ResultOf(c.ctx.Bot.Raw().GetMyStarBalance(c.opts))
}

// GetStarTransactions is a request builder for getting star transaction history.
type GetStarTransactions struct {
	ctx  *Context
	opts *gotgbot.GetStarTransactionsOpts
}

// Offset sets the number of transactions to skip.
func (c *GetStarTransactions) Offset(offset int64) *GetStarTransactions {
	c.opts.Offset = offset
	return c
}

// Limit sets the maximum number of transactions to retrieve (1-100, defaults to 100).
func (c *GetStarTransactions) Limit(limit int64) *GetStarTransactions {
	c.opts.Limit = limit
	return c
}

// Timeout sets a custom timeout for this request.
func (c *GetStarTransactions) Timeout(duration time.Duration) *GetStarTransactions {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *GetStarTransactions) APIURL(url String) *GetStarTransactions {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send executes the GetStarTransactions request.
func (c *GetStarTransactions) Send() Result[*gotgbot.StarTransactions] {
	return ResultOf(c.ctx.Bot.Raw().GetStarTransactions(c.opts))
}

// EditUserStarSubscription creates an EditUserStarSubscription request builder.
func (ctx *Context) EditUserStarSubscription(
	userID int64,
	telegramPaymentChargeID String,
	isCanceled bool,
) *EditUserStarSubscription {
	return &EditUserStarSubscription{
		ctx:                     ctx,
		userID:                  userID,
		telegramPaymentChargeID: telegramPaymentChargeID,
		isCanceled:              isCanceled,
		opts:                    new(gotgbot.EditUserStarSubscriptionOpts),
	}
}

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
