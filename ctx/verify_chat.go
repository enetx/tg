package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// VerifyChat represents a request to verify a chat.
type VerifyChat struct {
	ctx    *Context
	chatID int64
	opts   *gotgbot.VerifyChatOpts
}

// CustomDescription for the verification; 0-70 characters.
// Must be empty if the organization isn't allowed to provide a custom verification description.
func (vc *VerifyChat) CustomDescription(description g.String) *VerifyChat {
	vc.opts.CustomDescription = description.Std()
	return vc
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
func (vc *VerifyChat) APIURL(url g.String) *VerifyChat {
	if vc.opts.RequestOpts == nil {
		vc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	vc.opts.RequestOpts.APIURL = url.Std()

	return vc
}

// Send verifies the chat.
func (vc *VerifyChat) Send() g.Result[bool] {
	return g.ResultOf(vc.ctx.Bot.Raw().VerifyChat(vc.chatID, vc.opts))
}
