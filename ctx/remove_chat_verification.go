package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// RemoveChatVerification represents a request to remove chat verification.
type RemoveChatVerification struct {
	ctx    *Context
	chatID int64
	opts   *gotgbot.RemoveChatVerificationOpts
}

// Timeout sets a custom timeout for this request.
func (rcv *RemoveChatVerification) Timeout(duration time.Duration) *RemoveChatVerification {
	if rcv.opts.RequestOpts == nil {
		rcv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rcv.opts.RequestOpts.Timeout = duration

	return rcv
}

// APIURL sets a custom API URL for this request.
func (rcv *RemoveChatVerification) APIURL(url String) *RemoveChatVerification {
	if rcv.opts.RequestOpts == nil {
		rcv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rcv.opts.RequestOpts.APIURL = url.Std()

	return rcv
}

// Send removes chat verification.
func (rcv *RemoveChatVerification) Send() Result[bool] {
	return ResultOf(rcv.ctx.Bot.Raw().RemoveChatVerification(rcv.chatID, rcv.opts))
}
