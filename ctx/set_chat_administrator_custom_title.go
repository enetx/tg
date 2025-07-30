package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetChatAdministratorCustomTitle represents a request to set a custom title for an administrator.
type SetChatAdministratorCustomTitle struct {
	ctx         *Context
	userID      int64
	customTitle g.String
	opts        *gotgbot.SetChatAdministratorCustomTitleOpts
	chatID      g.Option[int64]
}

// ChatID sets the target chat ID for this request.
func (scact *SetChatAdministratorCustomTitle) ChatID(id int64) *SetChatAdministratorCustomTitle {
	scact.chatID = g.Some(id)
	return scact
}

// Timeout sets a custom timeout for this request.
func (scact *SetChatAdministratorCustomTitle) Timeout(duration time.Duration) *SetChatAdministratorCustomTitle {
	if scact.opts.RequestOpts == nil {
		scact.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scact.opts.RequestOpts.Timeout = duration

	return scact
}

// APIURL sets a custom API URL for this request.
func (scact *SetChatAdministratorCustomTitle) APIURL(url g.String) *SetChatAdministratorCustomTitle {
	if scact.opts.RequestOpts == nil {
		scact.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	scact.opts.RequestOpts.APIURL = url.Std()

	return scact
}

// Send executes the SetChatAdministratorCustomTitle request.
func (scact *SetChatAdministratorCustomTitle) Send() g.Result[bool] {
	chatID := scact.chatID.UnwrapOr(scact.ctx.EffectiveChat.Id)
	return g.ResultOf(
		scact.ctx.Bot.Raw().SetChatAdministratorCustomTitle(chatID, scact.userID, scact.customTitle.Std(), scact.opts),
	)
}
