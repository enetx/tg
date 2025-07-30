package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// CloseForumTopic represents a request to close a forum topic.
type CloseForumTopic struct {
	ctx             *Context
	messageThreadID int64
	opts            *gotgbot.CloseForumTopicOpts
	chatID          g.Option[int64]
}

// Timeout sets a custom timeout for this request.
func (cft *CloseForumTopic) Timeout(duration time.Duration) *CloseForumTopic {
	if cft.opts.RequestOpts == nil {
		cft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cft.opts.RequestOpts.Timeout = duration

	return cft
}

// APIURL sets a custom API URL for this request.
func (cft *CloseForumTopic) APIURL(url g.String) *CloseForumTopic {
	if cft.opts.RequestOpts == nil {
		cft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cft.opts.RequestOpts.APIURL = url.Std()

	return cft
}

// ChatID sets the target chat ID for this request.
func (cft *CloseForumTopic) ChatID(id int64) *CloseForumTopic {
	cft.chatID = g.Some(id)
	return cft
}

// Send executes the CloseForumTopic request.
func (cft *CloseForumTopic) Send() g.Result[bool] {
	chatID := cft.chatID.UnwrapOr(cft.ctx.EffectiveChat.Id)
	return g.ResultOf(cft.ctx.Bot.Raw().CloseForumTopic(chatID, cft.messageThreadID, cft.opts))
}
