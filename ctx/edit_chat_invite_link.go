package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// EditChatInviteLink represents a request to edit an existing chat invite link.
type EditChatInviteLink struct {
	ctx        *Context
	inviteLink String
	opts       *gotgbot.EditChatInviteLinkOpts
	chatID     Option[int64]
}

// ChatID sets the target chat ID.
func (ecil *EditChatInviteLink) ChatID(chatID int64) *EditChatInviteLink {
	ecil.chatID = Some(chatID)
	return ecil
}

// Name sets the invite link name.
func (ecil *EditChatInviteLink) Name(name String) *EditChatInviteLink {
	ecil.opts.Name = name.Std()
	return ecil
}

// ExpiresAt sets the invite link to expire at the specified time.
func (ecil *EditChatInviteLink) ExpiresAt(t time.Time) *EditChatInviteLink {
	ecil.opts.ExpireDate = t.Unix()
	return ecil
}

// ExpiresIn sets the invite link to expire after the given duration.
func (ecil *EditChatInviteLink) ExpiresIn(duration time.Duration) *EditChatInviteLink {
	ecil.opts.ExpireDate = time.Now().Add(duration).Unix()
	return ecil
}

// MemberLimit sets the maximum number of users that can be members simultaneously.
func (ecil *EditChatInviteLink) MemberLimit(limit int64) *EditChatInviteLink {
	ecil.opts.MemberLimit = limit
	return ecil
}

// CreatesJoinRequest sets whether users joining via this link need to be approved.
func (ecil *EditChatInviteLink) CreatesJoinRequest() *EditChatInviteLink {
	ecil.opts.CreatesJoinRequest = true
	return ecil
}

// Timeout sets a custom timeout for this request.
func (ecil *EditChatInviteLink) Timeout(duration time.Duration) *EditChatInviteLink {
	if ecil.opts.RequestOpts == nil {
		ecil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ecil.opts.RequestOpts.Timeout = duration

	return ecil
}

// APIURL sets a custom API URL for this request.
func (ecil *EditChatInviteLink) APIURL(url String) *EditChatInviteLink {
	if ecil.opts.RequestOpts == nil {
		ecil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ecil.opts.RequestOpts.APIURL = url.Std()

	return ecil
}

// Send edits the chat invite link and returns the result.
func (ecil *EditChatInviteLink) Send() Result[*gotgbot.ChatInviteLink] {
	chatID := ecil.chatID.UnwrapOr(ecil.ctx.EffectiveChat.Id)
	return ResultOf(ecil.ctx.Bot.Raw().EditChatInviteLink(chatID, ecil.inviteLink.Std(), ecil.opts))
}
