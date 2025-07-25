package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// VerifyChat represents a request to verify a chat.
type VerifyChat struct {
	ctx    *Context
	chatID int64
	opts   *gotgbot.VerifyChatOpts
}

// Timeout sets a custom timeout for this request.
func (vc *VerifyChat) Timeout(duration time.Duration) *VerifyChat {
	if vc.opts.RequestOpts == nil {
		vc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	vc.opts.RequestOpts.Timeout = duration

	return vc
}

// APIURL sets a custom API URL for this request.
func (vc *VerifyChat) APIURL(url String) *VerifyChat {
	if vc.opts.RequestOpts == nil {
		vc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	vc.opts.RequestOpts.APIURL = url.Std()

	return vc
}

// Send verifies the chat.
func (vc *VerifyChat) Send() Result[bool] {
	return ResultOf(vc.ctx.Bot.Raw().VerifyChat(vc.chatID, vc.opts))
}
