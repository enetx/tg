package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
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
func (gmsb *GetMyStarBalance) APIURL(url g.String) *GetMyStarBalance {
	if gmsb.opts.RequestOpts == nil {
		gmsb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmsb.opts.RequestOpts.APIURL = url.Std()

	return gmsb
}

// Send executes the GetMyStarBalance request.
func (gmsb *GetMyStarBalance) Send() g.Result[*gotgbot.StarAmount] {
	return g.ResultOf(gmsb.ctx.Bot.Raw().GetMyStarBalance(gmsb.opts))
}
