package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// DeleteForumTopic represents a request to delete a forum topic.
type DeleteForumTopic struct {
	ctx             *Context
	messageThreadID int64
	opts            *gotgbot.DeleteForumTopicOpts
	chatID          Option[int64]
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
func (dft *DeleteForumTopic) APIURL(url String) *DeleteForumTopic {
	if dft.opts.RequestOpts == nil {
		dft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dft.opts.RequestOpts.APIURL = url.Std()

	return dft
}

// ChatID sets the target chat ID for this request.
func (dft *DeleteForumTopic) ChatID(id int64) *DeleteForumTopic {
	dft.chatID = Some(id)
	return dft
}

// Send executes the DeleteForumTopic request.
func (dft *DeleteForumTopic) Send() Result[bool] {
	chatID := dft.chatID.UnwrapOr(dft.ctx.EffectiveChat.Id)
	return ResultOf(dft.ctx.Bot.Raw().DeleteForumTopic(chatID, dft.messageThreadID, dft.opts))
}
