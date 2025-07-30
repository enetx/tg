package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// ReopenGeneralForumTopic represents a request to reopen the general forum topic.
type ReopenGeneralForumTopic struct {
	ctx    *Context
	opts   *gotgbot.ReopenGeneralForumTopicOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID.
func (rgft *ReopenGeneralForumTopic) ChatID(chatID int64) *ReopenGeneralForumTopic {
	rgft.chatID = g.Some(chatID)
	return rgft
}

// Timeout sets a custom timeout for this request.
func (rgft *ReopenGeneralForumTopic) Timeout(duration time.Duration) *ReopenGeneralForumTopic {
	if rgft.opts.RequestOpts == nil {
		rgft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rgft.opts.RequestOpts.Timeout = duration

	return rgft
}

// APIURL sets a custom API URL for this request.
func (rgft *ReopenGeneralForumTopic) APIURL(url g.String) *ReopenGeneralForumTopic {
	if rgft.opts.RequestOpts == nil {
		rgft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rgft.opts.RequestOpts.APIURL = url.Std()

	return rgft
}

// Send reopens the general forum topic.
func (rgft *ReopenGeneralForumTopic) Send() g.Result[bool] {
	return g.ResultOf(rgft.ctx.Bot.Raw().ReopenGeneralForumTopic(
		rgft.chatID.UnwrapOr(rgft.ctx.EffectiveChat.Id),
		rgft.opts,
	))
}
