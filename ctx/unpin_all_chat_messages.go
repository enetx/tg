package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// UnpinAllChatMessages represents a request to unpin all messages.
type UnpinAllChatMessages struct {
	ctx    *Context
	opts   *gotgbot.UnpinAllChatMessagesOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID for this request.
func (uacm *UnpinAllChatMessages) ChatID(id int64) *UnpinAllChatMessages {
	uacm.chatID = Some(id)
	return uacm
}

// Timeout sets a custom timeout for this request.
func (uacm *UnpinAllChatMessages) Timeout(duration time.Duration) *UnpinAllChatMessages {
	if uacm.opts.RequestOpts == nil {
		uacm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	uacm.opts.RequestOpts.Timeout = duration

	return uacm
}

// APIURL sets a custom API URL for this request.
func (uacm *UnpinAllChatMessages) APIURL(url String) *UnpinAllChatMessages {
	if uacm.opts.RequestOpts == nil {
		uacm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	uacm.opts.RequestOpts.APIURL = url.Std()

	return uacm
}

// Send executes the UnpinAllChatMessages request.
func (uacm *UnpinAllChatMessages) Send() Result[bool] {
	chatID := uacm.chatID.UnwrapOr(uacm.ctx.EffectiveChat.Id)
	return ResultOf(uacm.ctx.Bot.Raw().UnpinAllChatMessages(chatID, uacm.opts))
}
