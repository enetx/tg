package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeleteForumTopic represents a request to delete a forum topic.
type DeleteForumTopic struct {
	ctx             *Context
	messageThreadID int64
	opts            *gotgbot.DeleteForumTopicOpts
	chatID          g.Option[int64]
}

// Timeout sets a custom timeout for this request.
func (dft *DeleteForumTopic) Timeout(duration time.Duration) *DeleteForumTopic {
	if dft.opts.RequestOpts == nil {
		dft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dft.opts.RequestOpts.Timeout = duration

	return dft
}

// APIURL sets a custom API URL for this request.
func (dft *DeleteForumTopic) APIURL(url g.String) *DeleteForumTopic {
	if dft.opts.RequestOpts == nil {
		dft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dft.opts.RequestOpts.APIURL = url.Std()

	return dft
}

// ChatID sets the target chat ID for this request.
func (dft *DeleteForumTopic) ChatID(id int64) *DeleteForumTopic {
	dft.chatID = g.Some(id)
	return dft
}

// Send executes the DeleteForumTopic request.
func (dft *DeleteForumTopic) Send() g.Result[bool] {
	chatID := dft.chatID.UnwrapOr(dft.ctx.EffectiveChat.Id)
	return g.ResultOf(dft.ctx.Bot.Raw().DeleteForumTopic(chatID, dft.messageThreadID, dft.opts))
}
