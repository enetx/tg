package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// BanChatSenderChat represents a request to ban a sender chat in a chat.
type BanChatSenderChat struct {
	ctx          *Context
	senderChatID int64
	opts         *gotgbot.BanChatSenderChatOpts
	chatID       g.Option[int64]
}

// ChatID sets the target chat ID.
func (bcsc *BanChatSenderChat) ChatID(chatID int64) *BanChatSenderChat {
	bcsc.chatID = g.Some(chatID)
	return bcsc
}

// Timeout sets a custom timeout for this request.
func (bcsc *BanChatSenderChat) Timeout(duration time.Duration) *BanChatSenderChat {
	if bcsc.opts.RequestOpts == nil {
		bcsc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	bcsc.opts.RequestOpts.Timeout = duration

	return bcsc
}

// APIURL sets a custom API URL for this request.
func (bcsc *BanChatSenderChat) APIURL(url g.String) *BanChatSenderChat {
	if bcsc.opts.RequestOpts == nil {
		bcsc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	bcsc.opts.RequestOpts.APIURL = url.Std()

	return bcsc
}

// Send bans the sender chat from the target chat.
func (bcsc *BanChatSenderChat) Send() g.Result[bool] {
	return g.ResultOf(bcsc.ctx.Bot.Raw().BanChatSenderChat(
		bcsc.chatID.UnwrapOr(bcsc.ctx.EffectiveChat.Id),
		bcsc.senderChatID,
		bcsc.opts,
	))
}
