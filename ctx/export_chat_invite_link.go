package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// ExportChatInviteLink represents a request to export a chat invite link.
type ExportChatInviteLink struct {
	ctx    *Context
	opts   *gotgbot.ExportChatInviteLinkOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID.
func (ecil *ExportChatInviteLink) ChatID(chatID int64) *ExportChatInviteLink {
	ecil.chatID = Some(chatID)
	return ecil
}

// Timeout sets a custom timeout for this request.
func (ecil *ExportChatInviteLink) Timeout(duration time.Duration) *ExportChatInviteLink {
	if ecil.opts.RequestOpts == nil {
		ecil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ecil.opts.RequestOpts.Timeout = duration

	return ecil
}

// APIURL sets a custom API URL for this request.
func (ecil *ExportChatInviteLink) APIURL(url String) *ExportChatInviteLink {
	if ecil.opts.RequestOpts == nil {
		ecil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ecil.opts.RequestOpts.APIURL = url.Std()

	return ecil
}

// Send exports the chat invite link and returns the result.
func (ecil *ExportChatInviteLink) Send() Result[String] {
	chatID := ecil.chatID.UnwrapOr(ecil.ctx.EffectiveChat.Id)
	link, err := ecil.ctx.Bot.Raw().ExportChatInviteLink(chatID, ecil.opts)

	return ResultOf(String(link), err)
}
