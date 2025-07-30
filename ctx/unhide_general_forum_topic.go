package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// UnhideGeneralForumTopic represents a request to unhide the general forum topic.
type UnhideGeneralForumTopic struct {
	ctx    *Context
	opts   *gotgbot.UnhideGeneralForumTopicOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID.
func (ugft *UnhideGeneralForumTopic) ChatID(chatID int64) *UnhideGeneralForumTopic {
	ugft.chatID = g.Some(chatID)
	return ugft
}

// Timeout sets a custom timeout for this request.
func (ugft *UnhideGeneralForumTopic) Timeout(duration time.Duration) *UnhideGeneralForumTopic {
	if ugft.opts.RequestOpts == nil {
		ugft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ugft.opts.RequestOpts.Timeout = duration

	return ugft
}

// APIURL sets a custom API URL for this request.
func (ugft *UnhideGeneralForumTopic) APIURL(url g.String) *UnhideGeneralForumTopic {
	if ugft.opts.RequestOpts == nil {
		ugft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ugft.opts.RequestOpts.APIURL = url.Std()

	return ugft
}

// Send unhides the general forum topic.
func (ugft *UnhideGeneralForumTopic) Send() g.Result[bool] {
	return g.ResultOf(ugft.ctx.Bot.Raw().UnhideGeneralForumTopic(
		ugft.chatID.UnwrapOr(ugft.ctx.EffectiveChat.Id),
		ugft.opts,
	))
}
