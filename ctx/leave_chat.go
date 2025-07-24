package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// LeaveChat represents a request to leave a chat.
type LeaveChat struct {
	ctx    *Context
	opts   *gotgbot.LeaveChatOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID.
func (lc *LeaveChat) ChatID(chatID int64) *LeaveChat {
	lc.chatID = Some(chatID)
	return lc
}

// Timeout sets a custom timeout for this request.
func (lc *LeaveChat) Timeout(duration time.Duration) *LeaveChat {
	if lc.opts.RequestOpts == nil {
		lc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	lc.opts.RequestOpts.Timeout = duration

	return lc
}

// APIURL sets a custom API URL for this request.
func (lc *LeaveChat) APIURL(url String) *LeaveChat {
	if lc.opts.RequestOpts == nil {
		lc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	lc.opts.RequestOpts.APIURL = url.Std()

	return lc
}

// Send leaves the chat and returns the result.
func (lc *LeaveChat) Send() Result[bool] {
	chatID := lc.chatID.UnwrapOr(lc.ctx.EffectiveChat.Id)
	return ResultOf(lc.ctx.Bot.Raw().LeaveChat(chatID, lc.opts))
}
