package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// UnpinAllForumTopicMessages represents a request to unpin all messages in a forum topic.
type UnpinAllForumTopicMessages struct {
	ctx             *Context
	messageThreadID int64
	opts            *gotgbot.UnpinAllForumTopicMessagesOpts
	chatID          g.Option[int64]
}

// ChatID sets the target chat ID.
func (uaftm *UnpinAllForumTopicMessages) ChatID(chatID int64) *UnpinAllForumTopicMessages {
	uaftm.chatID = g.Some(chatID)
	return uaftm
}

// Timeout sets a custom timeout for this request.
func (uaftm *UnpinAllForumTopicMessages) Timeout(duration time.Duration) *UnpinAllForumTopicMessages {
	if uaftm.opts.RequestOpts == nil {
		uaftm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	uaftm.opts.RequestOpts.Timeout = duration

	return uaftm
}

// APIURL sets a custom API URL for this request.
func (uaftm *UnpinAllForumTopicMessages) APIURL(url g.String) *UnpinAllForumTopicMessages {
	if uaftm.opts.RequestOpts == nil {
		uaftm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	uaftm.opts.RequestOpts.APIURL = url.Std()

	return uaftm
}

// Send unpins all messages in the forum topic.
func (uaftm *UnpinAllForumTopicMessages) Send() g.Result[bool] {
	return g.ResultOf(uaftm.ctx.Bot.Raw().UnpinAllForumTopicMessages(
		uaftm.chatID.UnwrapOr(uaftm.ctx.EffectiveChat.Id),
		uaftm.messageThreadID,
		uaftm.opts,
	))
}
