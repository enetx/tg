package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// GetConnection is a request builder for retrieving business connection info.
type GetConnection struct {
	account *Account
	opts    *gotgbot.GetBusinessConnectionOpts
}

// Timeout sets a custom timeout for this request.
func (g *GetConnection) Timeout(duration time.Duration) *GetConnection {
	if g.opts.RequestOpts == nil {
		g.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	g.opts.RequestOpts.Timeout = duration

	return g
}

// APIURL sets a custom API URL for this request.
func (g *GetConnection) APIURL(url String) *GetConnection {
	if g.opts.RequestOpts == nil {
		g.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	g.opts.RequestOpts.APIURL = url.Std()

	return g
}

// Send executes the Get request.
func (g *GetConnection) Send() Result[*gotgbot.BusinessConnection] {
	return ResultOf(g.account.bot.Raw().GetBusinessConnection(
		g.account.connID.Std(),
		g.opts,
	))
}
