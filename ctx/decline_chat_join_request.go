package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeclineChatJoinRequest represents a request to decline a chat join request.
type DeclineChatJoinRequest struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.DeclineChatJoinRequestOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID.
func (dcjr *DeclineChatJoinRequest) ChatID(chatID int64) *DeclineChatJoinRequest {
	dcjr.chatID = g.Some(chatID)
	return dcjr
}

// Timeout sets a custom timeout for this request.
func (dcjr *DeclineChatJoinRequest) Timeout(duration time.Duration) *DeclineChatJoinRequest {
	if dcjr.opts.RequestOpts == nil {
		dcjr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dcjr.opts.RequestOpts.Timeout = duration

	return dcjr
}

// APIURL sets a custom API URL for this request.
func (dcjr *DeclineChatJoinRequest) APIURL(url g.String) *DeclineChatJoinRequest {
	if dcjr.opts.RequestOpts == nil {
		dcjr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dcjr.opts.RequestOpts.APIURL = url.Std()

	return dcjr
}

// Send declines the chat join request and returns the result.
func (dcjr *DeclineChatJoinRequest) Send() g.Result[bool] {
	chatID := dcjr.chatID.UnwrapOr(dcjr.ctx.EffectiveChat.Id)
	return g.ResultOf(dcjr.ctx.Bot.Raw().DeclineChatJoinRequest(chatID, dcjr.userID, dcjr.opts))
}
