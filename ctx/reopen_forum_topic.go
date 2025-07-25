package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// ReopenForumTopic represents a request to reopen a forum topic.
type ReopenForumTopic struct {
	ctx             *Context
	messageThreadID int64
	opts            *gotgbot.ReopenForumTopicOpts
	chatID          Option[int64]
}

// Timeout sets a custom timeout for this request.
func (rft *ReopenForumTopic) Timeout(duration time.Duration) *ReopenForumTopic {
	if rft.opts.RequestOpts == nil {
		rft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rft.opts.RequestOpts.Timeout = duration

	return rft
}

// APIURL sets a custom API URL for this request.
func (rft *ReopenForumTopic) APIURL(url String) *ReopenForumTopic {
	if rft.opts.RequestOpts == nil {
		rft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rft.opts.RequestOpts.APIURL = url.Std()

	return rft
}

// ChatID sets the target chat ID for this request.
func (rft *ReopenForumTopic) ChatID(id int64) *ReopenForumTopic {
	rft.chatID = Some(id)
	return rft
}

// Send executes the ReopenForumTopic request.
func (rft *ReopenForumTopic) Send() Result[bool] {
	chatID := rft.chatID.UnwrapOr(rft.ctx.EffectiveChat.Id)
	return ResultOf(rft.ctx.Bot.Raw().ReopenForumTopic(chatID, rft.messageThreadID, rft.opts))
}
