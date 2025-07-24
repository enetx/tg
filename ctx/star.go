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
func (gmsb *GetMyStarBalance) Timeout(duration time.Duration) *GetMyStarBalance {
	if gmsb.opts.RequestOpts == nil {
		gmsb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmsb.opts.RequestOpts.Timeout = duration

	return gmsb
}

// APIURL sets a custom API URL for this request.
func (gmsb *GetMyStarBalance) APIURL(url String) *GetMyStarBalance {
	if gmsb.opts.RequestOpts == nil {
		gmsb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmsb.opts.RequestOpts.APIURL = url.Std()

	return gmsb
}

// Send executes the GetMyStarBalance request.
func (gmsb *GetMyStarBalance) Send() Result[*gotgbot.StarAmount] {
	return ResultOf(gmsb.ctx.Bot.Raw().GetMyStarBalance(gmsb.opts))
}

// GetStarTransactions is a request builder for getting star transaction history.
type GetStarTransactions struct {
	ctx  *Context
	opts *gotgbot.GetStarTransactionsOpts
}

// Offset sets the number of transactions to skip.
func (gsts *GetStarTransactions) Offset(offset int64) *GetStarTransactions {
	gsts.opts.Offset = offset
	return gsts
}

// Limit sets the maximum number of transactions to retrieve (1-100, defaults to 100).
func (gsts *GetStarTransactions) Limit(limit int64) *GetStarTransactions {
	gsts.opts.Limit = limit
	return gsts
}

// Timeout sets a custom timeout for this request.
func (gsts *GetStarTransactions) Timeout(duration time.Duration) *GetStarTransactions {
	if gsts.opts.RequestOpts == nil {
		gsts.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gsts.opts.RequestOpts.Timeout = duration

	return gsts
}

// APIURL sets a custom API URL for this request.
func (gsts *GetStarTransactions) APIURL(url String) *GetStarTransactions {
	if gsts.opts.RequestOpts == nil {
		gsts.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gsts.opts.RequestOpts.APIURL = url.Std()

	return gsts
}

// Send executes the GetStarTransactions request.
func (gsts *GetStarTransactions) Send() Result[*gotgbot.StarTransactions] {
	return ResultOf(gsts.ctx.Bot.Raw().GetStarTransactions(gsts.opts))
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
