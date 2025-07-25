package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// ApproveChatJoinRequest represents a request to approve a chat join request.
type ApproveChatJoinRequest struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.ApproveChatJoinRequestOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID.
func (acjr *ApproveChatJoinRequest) ChatID(chatID int64) *ApproveChatJoinRequest {
	acjr.chatID = Some(chatID)
	return acjr
}

// Timeout sets a custom timeout for this request.
func (acjr *ApproveChatJoinRequest) Timeout(duration time.Duration) *ApproveChatJoinRequest {
	if acjr.opts.RequestOpts == nil {
		acjr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	acjr.opts.RequestOpts.Timeout = duration

	return acjr
}

// APIURL sets a custom API URL for this request.
func (acjr *ApproveChatJoinRequest) APIURL(url String) *ApproveChatJoinRequest {
	if acjr.opts.RequestOpts == nil {
		acjr.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	acjr.opts.RequestOpts.APIURL = url.Std()

	return acjr
}

// Send approves the chat join request and returns the result.
func (acjr *ApproveChatJoinRequest) Send() Result[bool] {
	chatID := acjr.chatID.UnwrapOr(acjr.ctx.EffectiveChat.Id)
	return ResultOf(acjr.ctx.Bot.Raw().ApproveChatJoinRequest(chatID, acjr.userID, acjr.opts))
}
