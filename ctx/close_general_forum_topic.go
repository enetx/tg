package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// CloseGeneralForumTopic represents a request to close the general forum topic.
type CloseGeneralForumTopic struct {
	ctx    *Context
	opts   *gotgbot.CloseGeneralForumTopicOpts
	chatID Option[int64]
}

// Timeout sets a custom timeout for this request.
func (cgft *CloseGeneralForumTopic) Timeout(duration time.Duration) *CloseGeneralForumTopic {
	if cgft.opts.RequestOpts == nil {
		cgft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cgft.opts.RequestOpts.Timeout = duration

	return cgft
}

// APIURL sets a custom API URL for this request.
func (cgft *CloseGeneralForumTopic) APIURL(url String) *CloseGeneralForumTopic {
	if cgft.opts.RequestOpts == nil {
		cgft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cgft.opts.RequestOpts.APIURL = url.Std()

	return cgft
}

// ChatID sets the target chat ID for this request.
func (cgft *CloseGeneralForumTopic) ChatID(id int64) *CloseGeneralForumTopic {
	cgft.chatID = Some(id)
	return cgft
}

// Send executes the CloseGeneralForumTopic request.
func (cgft *CloseGeneralForumTopic) Send() Result[bool] {
	chatID := cgft.chatID.UnwrapOr(cgft.ctx.EffectiveChat.Id)
	return ResultOf(cgft.ctx.Bot.Raw().CloseGeneralForumTopic(chatID, cgft.opts))
}
