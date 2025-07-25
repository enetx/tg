package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

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
