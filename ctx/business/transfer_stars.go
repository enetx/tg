package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// TransferStars request builder for sending stars.
type TransferStars struct {
	bot    Bot
	connID g.String
	amount int64
	opts   *gotgbot.TransferBusinessAccountStarsOpts
}

// Timeout sets a custom timeout for this request.
func (t *TransferStars) Timeout(duration time.Duration) *TransferStars {
	if t.opts.RequestOpts == nil {
		t.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	t.opts.RequestOpts.Timeout = duration

	return t
}

// APIURL sets a custom API URL for this request.
func (t *TransferStars) APIURL(url g.String) *TransferStars {
	if t.opts.RequestOpts == nil {
		t.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	t.opts.RequestOpts.APIURL = url.Std()

	return t
}

// Send executes the Transfer request.
func (t *TransferStars) Send() g.Result[bool] {
	return g.ResultOf(t.bot.Raw().TransferBusinessAccountStars(t.connID.Std(), t.amount, t.opts))
}
