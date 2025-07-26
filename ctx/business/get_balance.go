package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// GetStarBalance request builder for star balance.
type GetStarBalance struct {
	bot    Bot
	connID String
	opts   *gotgbot.GetBusinessAccountStarBalanceOpts
}

// Timeout sets a custom timeout for this request.
func (gb *GetStarBalance) Timeout(duration time.Duration) *GetStarBalance {
	if gb.opts.RequestOpts == nil {
		gb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gb.opts.RequestOpts.Timeout = duration

	return gb
}

// APIURL sets a custom API URL for this request.
func (gb *GetStarBalance) APIURL(url String) *GetStarBalance {
	if gb.opts.RequestOpts == nil {
		gb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gb.opts.RequestOpts.APIURL = url.Std()

	return gb
}

// Send executes the GetStarBalance request.
func (gb *GetStarBalance) Send() Result[*gotgbot.StarAmount] {
	return ResultOf(gb.bot.Raw().GetBusinessAccountStarBalance(gb.connID.Std(), gb.opts))
}
