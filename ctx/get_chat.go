package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetChat represents a request to get chat information.
type GetChat struct {
	ctx    *Context
	opts   *gotgbot.GetChatOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID for this request.
func (gc *GetChat) ChatID(id int64) *GetChat {
	gc.chatID = g.Some(id)
	return gc
}

// Timeout sets a custom timeout for this request.
func (gc *GetChat) Timeout(duration time.Duration) *GetChat {
	if gc.opts.RequestOpts == nil {
		gc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gc.opts.RequestOpts.Timeout = duration

	return gc
}

// APIURL sets a custom API URL for this request.
func (gc *GetChat) APIURL(url g.String) *GetChat {
	if gc.opts.RequestOpts == nil {
		gc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gc.opts.RequestOpts.APIURL = url.Std()

	return gc
}

// Send executes the GetChat request and returns full chat information.
func (gc *GetChat) Send() g.Result[*gotgbot.ChatFullInfo] {
	chatID := gc.chatID.UnwrapOr(gc.ctx.EffectiveChat.Id)
	return g.ResultOf(gc.ctx.Bot.Raw().GetChat(chatID, gc.opts))
}
