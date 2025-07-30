package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// HideGeneralForumTopic represents a request to hide the general forum topic.
type HideGeneralForumTopic struct {
	ctx    *Context
	opts   *gotgbot.HideGeneralForumTopicOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID.
func (hgft *HideGeneralForumTopic) ChatID(chatID int64) *HideGeneralForumTopic {
	hgft.chatID = g.Some(chatID)
	return hgft
}

// Timeout sets a custom timeout for this request.
func (hgft *HideGeneralForumTopic) Timeout(duration time.Duration) *HideGeneralForumTopic {
	if hgft.opts.RequestOpts == nil {
		hgft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	hgft.opts.RequestOpts.Timeout = duration

	return hgft
}

// APIURL sets a custom API URL for this request.
func (hgft *HideGeneralForumTopic) APIURL(url g.String) *HideGeneralForumTopic {
	if hgft.opts.RequestOpts == nil {
		hgft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	hgft.opts.RequestOpts.APIURL = url.Std()

	return hgft
}

// Send hides the general forum topic.
func (hgft *HideGeneralForumTopic) Send() g.Result[bool] {
	return g.ResultOf(hgft.ctx.Bot.Raw().HideGeneralForumTopic(
		hgft.chatID.UnwrapOr(hgft.ctx.EffectiveChat.Id),
		hgft.opts,
	))
}
