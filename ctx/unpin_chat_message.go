package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// UnpinChatMessage represents a request to unpin a message.
type UnpinChatMessage struct {
	ctx    *Context
	opts   *gotgbot.UnpinChatMessageOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID for this request.
func (ucm *UnpinChatMessage) ChatID(id int64) *UnpinChatMessage {
	ucm.chatID = g.Some(id)
	return ucm
}

// MessageID sets the specific message ID to unpin.
func (ucm *UnpinChatMessage) MessageID(messageID int64) *UnpinChatMessage {
	ucm.opts.MessageId = &messageID
	return ucm
}

// Business sets the business connection ID for the unpin action.
func (ucm *UnpinChatMessage) Business(id g.String) *UnpinChatMessage {
	ucm.opts.BusinessConnectionId = id.Std()
	return ucm
}

// Timeout sets a custom timeout for this request.
func (ucm *UnpinChatMessage) Timeout(duration time.Duration) *UnpinChatMessage {
	if ucm.opts.RequestOpts == nil {
		ucm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ucm.opts.RequestOpts.Timeout = duration

	return ucm
}

// APIURL sets a custom API URL for this request.
func (ucm *UnpinChatMessage) APIURL(url g.String) *UnpinChatMessage {
	if ucm.opts.RequestOpts == nil {
		ucm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ucm.opts.RequestOpts.APIURL = url.Std()

	return ucm
}

// Send executes the UnpinChatMessage request.
func (ucm *UnpinChatMessage) Send() g.Result[bool] {
	chatID := ucm.chatID.UnwrapOr(ucm.ctx.EffectiveChat.Id)
	return g.ResultOf(ucm.ctx.Bot.Raw().UnpinChatMessage(chatID, ucm.opts))
}
