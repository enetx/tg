package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// RevokeChatInviteLink represents a request to revoke a chat invite link.
type RevokeChatInviteLink struct {
	ctx        *Context
	inviteLink String
	opts       *gotgbot.RevokeChatInviteLinkOpts
	chatID     Option[int64]
}

// ChatID sets the target chat ID.
func (rcil *RevokeChatInviteLink) ChatID(chatID int64) *RevokeChatInviteLink {
	rcil.chatID = Some(chatID)
	return rcil
}

// Timeout sets a custom timeout for this request.
func (rcil *RevokeChatInviteLink) Timeout(duration time.Duration) *RevokeChatInviteLink {
	if rcil.opts.RequestOpts == nil {
		rcil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rcil.opts.RequestOpts.Timeout = duration

	return rcil
}

// APIURL sets a custom API URL for this request.
func (rcil *RevokeChatInviteLink) APIURL(url String) *RevokeChatInviteLink {
	if rcil.opts.RequestOpts == nil {
		rcil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rcil.opts.RequestOpts.APIURL = url.Std()

	return rcil
}

// Send revokes the chat invite link and returns the result.
func (rcil *RevokeChatInviteLink) Send() Result[*gotgbot.ChatInviteLink] {
	chatID := rcil.chatID.UnwrapOr(rcil.ctx.EffectiveChat.Id)
	return ResultOf(rcil.ctx.Bot.Raw().RevokeChatInviteLink(chatID, rcil.inviteLink.Std(), rcil.opts))
}
