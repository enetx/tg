package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// SetChatTitle represents a request to set the chat title.
type SetChatTitle struct {
	ctx    *Context
	title  String
	opts   *gotgbot.SetChatTitleOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID for this request.
func (sat *SetChatTitle) ChatID(id int64) *SetChatTitle {
	sat.chatID = Some(id)
	return sat
}

// Timeout sets a custom timeout for this request.
func (sat *SetChatTitle) Timeout(duration time.Duration) *SetChatTitle {
	if sat.opts.RequestOpts == nil {
		sat.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sat.opts.RequestOpts.Timeout = duration

	return sat
}

// APIURL sets a custom API URL for this request.
func (sat *SetChatTitle) APIURL(url String) *SetChatTitle {
	if sat.opts.RequestOpts == nil {
		sat.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sat.opts.RequestOpts.APIURL = url.Std()

	return sat
}

// Send executes the SetChatTitle request.
func (sat *SetChatTitle) Send() Result[bool] {
	chatID := sat.chatID.UnwrapOr(sat.ctx.EffectiveChat.Id)
	return ResultOf(sat.ctx.Bot.Raw().SetChatTitle(chatID, sat.title.Std(), sat.opts))
}
