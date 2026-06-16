package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetManagedBotToken represents a request to get the token of a managed bot.
type GetManagedBotToken struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.GetManagedBotTokenOpts
}

// Timeout sets a custom timeout for this request.
func (gmbt *GetManagedBotToken) Timeout(duration time.Duration) *GetManagedBotToken {
	if gmbt.opts.RequestOpts == nil {
		gmbt.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmbt.opts.RequestOpts.Timeout = duration

	return gmbt
}

// APIURL sets a custom API URL for this request.
func (gmbt *GetManagedBotToken) APIURL(url g.String) *GetManagedBotToken {
	if gmbt.opts.RequestOpts == nil {
		gmbt.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gmbt.opts.RequestOpts.APIURL = url.Std()

	return gmbt
}

// Send retrieves the managed bot token and returns the result.
func (gmbt *GetManagedBotToken) Send() g.Result[g.String] {
	token, err := gmbt.ctx.Bot.Raw().GetManagedBotToken(gmbt.userID, gmbt.opts)
	return g.ResultOf(g.String(token), err)
}
