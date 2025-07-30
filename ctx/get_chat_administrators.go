package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetChatAdministrators represents a request to get chat administrators.
type GetChatAdministrators struct {
	ctx    *Context
	opts   *gotgbot.GetChatAdministratorsOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID for this request.
func (gca *GetChatAdministrators) ChatID(id int64) *GetChatAdministrators {
	gca.chatID = g.Some(id)
	return gca
}

// Timeout sets a custom timeout for this request.
func (gca *GetChatAdministrators) Timeout(duration time.Duration) *GetChatAdministrators {
	if gca.opts.RequestOpts == nil {
		gca.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gca.opts.RequestOpts.Timeout = duration

	return gca
}

// APIURL sets a custom API URL for this request.
func (gca *GetChatAdministrators) APIURL(url g.String) *GetChatAdministrators {
	if gca.opts.RequestOpts == nil {
		gca.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gca.opts.RequestOpts.APIURL = url.Std()

	return gca
}

// Send executes the GetChatAdministrators request.
func (gca *GetChatAdministrators) Send() g.Result[g.Slice[gotgbot.ChatMember]] {
	chatID := gca.chatID.UnwrapOr(gca.ctx.EffectiveChat.Id)
	members, err := gca.ctx.Bot.Raw().GetChatAdministrators(chatID, gca.opts)

	return g.ResultOf(g.Slice[gotgbot.ChatMember](members), err)
}
