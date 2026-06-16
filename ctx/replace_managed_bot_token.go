package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// ReplaceManagedBotToken represents a request to revoke the current token of a managed bot
// and generate a new one.
type ReplaceManagedBotToken struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.ReplaceManagedBotTokenOpts
}

// Timeout sets a custom timeout for this request.
func (rmbt *ReplaceManagedBotToken) Timeout(duration time.Duration) *ReplaceManagedBotToken {
	if rmbt.opts.RequestOpts == nil {
		rmbt.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rmbt.opts.RequestOpts.Timeout = duration

	return rmbt
}

// APIURL sets a custom API URL for this request.
func (rmbt *ReplaceManagedBotToken) APIURL(url g.String) *ReplaceManagedBotToken {
	if rmbt.opts.RequestOpts == nil {
		rmbt.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rmbt.opts.RequestOpts.APIURL = url.Std()

	return rmbt
}

// Send replaces the managed bot token and returns the new token.
func (rmbt *ReplaceManagedBotToken) Send() g.Result[g.String] {
	token, err := rmbt.ctx.Bot.Raw().ReplaceManagedBotToken(rmbt.userID, rmbt.opts)
	return g.ResultOf(g.String(token), err)
}
