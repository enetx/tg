package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// EditGeneralForumTopic represents a request to edit the general forum topic.
type EditGeneralForumTopic struct {
	ctx    *Context
	name   g.String
	opts   *gotgbot.EditGeneralForumTopicOpts
	chatID g.Option[int64]
}

// Timeout sets a custom timeout for this request.
func (egft *EditGeneralForumTopic) Timeout(duration time.Duration) *EditGeneralForumTopic {
	if egft.opts.RequestOpts == nil {
		egft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	egft.opts.RequestOpts.Timeout = duration

	return egft
}

// APIURL sets a custom API URL for this request.
func (egft *EditGeneralForumTopic) APIURL(url g.String) *EditGeneralForumTopic {
	if egft.opts.RequestOpts == nil {
		egft.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	egft.opts.RequestOpts.APIURL = url.Std()

	return egft
}

// ChatID sets the target chat ID for this request.
func (egft *EditGeneralForumTopic) ChatID(id int64) *EditGeneralForumTopic {
	egft.chatID = g.Some(id)
	return egft
}

// Send executes the EditGeneralForumTopic request.
func (egft *EditGeneralForumTopic) Send() g.Result[bool] {
	chatID := egft.chatID.UnwrapOr(egft.ctx.EffectiveChat.Id)
	return g.ResultOf(egft.ctx.Bot.Raw().EditGeneralForumTopic(chatID, egft.name.Std(), egft.opts))
}
