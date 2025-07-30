package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// UnbanChatSenderChat represents a request to unban a sender chat in a chat.
type UnbanChatSenderChat struct {
	ctx          *Context
	senderChatID int64
	opts         *gotgbot.UnbanChatSenderChatOpts
	chatID       g.Option[int64]
}

// ChatID sets the target chat ID.
func (ucsc *UnbanChatSenderChat) ChatID(chatID int64) *UnbanChatSenderChat {
	ucsc.chatID = g.Some(chatID)
	return ucsc
}

// Timeout sets a custom timeout for this request.
func (ucsc *UnbanChatSenderChat) Timeout(duration time.Duration) *UnbanChatSenderChat {
	if ucsc.opts.RequestOpts == nil {
		ucsc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ucsc.opts.RequestOpts.Timeout = duration

	return ucsc
}

// APIURL sets a custom API URL for this request.
func (ucsc *UnbanChatSenderChat) APIURL(url g.String) *UnbanChatSenderChat {
	if ucsc.opts.RequestOpts == nil {
		ucsc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ucsc.opts.RequestOpts.APIURL = url.Std()

	return ucsc
}

// Send unbans the sender chat from the target chat.
func (ucsc *UnbanChatSenderChat) Send() g.Result[bool] {
	return g.ResultOf(ucsc.ctx.Bot.Raw().UnbanChatSenderChat(
		ucsc.chatID.UnwrapOr(ucsc.ctx.EffectiveChat.Id),
		ucsc.senderChatID,
		ucsc.opts,
	))
}
