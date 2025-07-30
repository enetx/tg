package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetChatMenuButton represents a request to get the menu button of a chat.
type GetChatMenuButton struct {
	ctx    *Context
	chatID g.Option[*int64]
	opts   *gotgbot.GetChatMenuButtonOpts
}

// ChatID sets the target chat ID.
func (gcmb *GetChatMenuButton) ChatID(chatID int64) *GetChatMenuButton {
	gcmb.chatID = g.Some(&chatID)
	return gcmb
}

// Timeout sets a custom timeout for this request.
func (gcmb *GetChatMenuButton) Timeout(duration time.Duration) *GetChatMenuButton {
	if gcmb.opts.RequestOpts == nil {
		gcmb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcmb.opts.RequestOpts.Timeout = duration

	return gcmb
}

// APIURL sets a custom API URL for this request.
func (gcmb *GetChatMenuButton) APIURL(url g.String) *GetChatMenuButton {
	if gcmb.opts.RequestOpts == nil {
		gcmb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gcmb.opts.RequestOpts.APIURL = url.Std()

	return gcmb
}

// Send gets the chat menu button.
func (gcmb *GetChatMenuButton) Send() g.Result[gotgbot.MenuButton] {
	gcmb.opts.ChatId = gcmb.chatID.UnwrapOrDefault()
	return g.ResultOf(gcmb.ctx.Bot.Raw().GetChatMenuButton(gcmb.opts))
}
