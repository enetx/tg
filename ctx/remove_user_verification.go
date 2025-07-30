package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// RemoveUserVerification represents a request to remove user verification.
type RemoveUserVerification struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.RemoveUserVerificationOpts
}

// Timeout sets a custom timeout for this request.
func (ruv *RemoveUserVerification) Timeout(duration time.Duration) *RemoveUserVerification {
	if ruv.opts.RequestOpts == nil {
		ruv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ruv.opts.RequestOpts.Timeout = duration

	return ruv
}

// APIURL sets a custom API URL for this request.
func (ruv *RemoveUserVerification) APIURL(url g.String) *RemoveUserVerification {
	if ruv.opts.RequestOpts == nil {
		ruv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ruv.opts.RequestOpts.APIURL = url.Std()

	return ruv
}

// Send removes user verification.
func (ruv *RemoveUserVerification) Send() g.Result[bool] {
	return g.ResultOf(ruv.ctx.Bot.Raw().RemoveUserVerification(ruv.userID, ruv.opts))
}
