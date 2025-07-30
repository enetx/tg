package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetChatMemberCount represents a request to get the chat member count.
type GetChatMemberCount struct {
	ctx    *Context
	opts   *gotgbot.GetChatMemberCountOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID for this request.
func (gcm *GetChatMemberCount) ChatID(id int64) *GetChatMemberCount {
	gcm.chatID = g.Some(id)
	return gcm
}

// Timeout sets a custom timeout for this request.
func (gcm *GetChatMemberCount) Timeout(duration time.Duration) *GetChatMemberCount {
	if gcm.opts.RequestOpts == nil {
		gcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcm.opts.RequestOpts.Timeout = duration

	return gcm
}

// APIURL sets a custom API URL for this request.
func (gcm *GetChatMemberCount) APIURL(url g.String) *GetChatMemberCount {
	if gcm.opts.RequestOpts == nil {
		gcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcm.opts.RequestOpts.APIURL = url.Std()

	return gcm
}

// Send executes the GetChatMemberCount request.
func (gcm *GetChatMemberCount) Send() g.Result[g.Int] {
	chatID := gcm.chatID.UnwrapOr(gcm.ctx.EffectiveChat.Id)
	count, err := gcm.ctx.Bot.Raw().GetChatMemberCount(chatID, gcm.opts)

	return g.ResultOf(g.Int(count), err)
}
