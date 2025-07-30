package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetConnection is a request builder for retrieving business connection info.
type GetConnection struct {
	account *Account
	opts    *gotgbot.GetBusinessConnectionOpts
}

// Timeout sets a custom timeout for this request.
func (gc *GetConnection) Timeout(duration time.Duration) *GetConnection {
	if gc.opts.RequestOpts == nil {
		gc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gc.opts.RequestOpts.Timeout = duration

	return gc
}

// APIURL sets a custom API URL for this request.
func (gc *GetConnection) APIURL(url g.String) *GetConnection {
	if gc.opts.RequestOpts == nil {
		gc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gc.opts.RequestOpts.APIURL = url.Std()

	return gc
}

// Send executes the Get request.
func (gc *GetConnection) Send() g.Result[*gotgbot.BusinessConnection] {
	return g.ResultOf(gc.account.bot.Raw().GetBusinessConnection(
		gc.account.connID.Std(),
		gc.opts,
	))
}
