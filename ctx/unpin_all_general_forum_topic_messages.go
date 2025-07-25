package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// UnpinAllGeneralForumTopicMessages represents a request to unpin all messages in the general forum topic.
type UnpinAllGeneralForumTopicMessages struct {
	ctx    *Context
	opts   *gotgbot.UnpinAllGeneralForumTopicMessagesOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID.
func (uagftm *UnpinAllGeneralForumTopicMessages) ChatID(chatID int64) *UnpinAllGeneralForumTopicMessages {
	uagftm.chatID = Some(chatID)
	return uagftm
}

// Timeout sets a custom timeout for this request.
func (uagftm *UnpinAllGeneralForumTopicMessages) Timeout(duration time.Duration) *UnpinAllGeneralForumTopicMessages {
	if uagftm.opts.RequestOpts == nil {
		uagftm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	uagftm.opts.RequestOpts.Timeout = duration

	return uagftm
}

// APIURL sets a custom API URL for this request.
func (uagftm *UnpinAllGeneralForumTopicMessages) APIURL(url String) *UnpinAllGeneralForumTopicMessages {
	if uagftm.opts.RequestOpts == nil {
		uagftm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	uagftm.opts.RequestOpts.APIURL = url.Std()

	return uagftm
}

// Send unpins all messages in the general forum topic.
func (uagftm *UnpinAllGeneralForumTopicMessages) Send() Result[bool] {
	return ResultOf(uagftm.ctx.Bot.Raw().UnpinAllGeneralForumTopicMessages(
		uagftm.chatID.UnwrapOr(uagftm.ctx.EffectiveChat.Id),
		uagftm.opts,
	))
}