package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// CreateChatInviteLink represents a request to create a new chat invite link.
type CreateChatInviteLink struct {
	ctx    *Context
	opts   *gotgbot.CreateChatInviteLinkOpts
	chatID g.Option[int64]
}

// ChatID sets the target chat ID.
func (ccil *CreateChatInviteLink) ChatID(chatID int64) *CreateChatInviteLink {
	ccil.chatID = g.Some(chatID)
	return ccil
}

// Name sets the invite link name.
func (ccil *CreateChatInviteLink) Name(name g.String) *CreateChatInviteLink {
	ccil.opts.Name = name.Std()
	return ccil
}

// ExpiresAt sets the invite link to expire at the specified time.
func (ccil *CreateChatInviteLink) ExpiresAt(t time.Time) *CreateChatInviteLink {
	ccil.opts.ExpireDate = t.Unix()
	return ccil
}

// ExpiresIn sets the invite link to expire after the given duration.
func (ccil *CreateChatInviteLink) ExpiresIn(duration time.Duration) *CreateChatInviteLink {
	ccil.opts.ExpireDate = time.Now().Add(duration).Unix()
	return ccil
}

// MemberLimit sets the maximum number of users that can be members simultaneously.
func (ccil *CreateChatInviteLink) MemberLimit(limit int64) *CreateChatInviteLink {
	ccil.opts.MemberLimit = limit
	return ccil
}

// CreatesJoinRequest sets whether users joining via this link need to be approved.
func (ccil *CreateChatInviteLink) CreatesJoinRequest() *CreateChatInviteLink {
	ccil.opts.CreatesJoinRequest = true
	return ccil
}

// Timeout sets a custom timeout for this request.
func (ccil *CreateChatInviteLink) Timeout(duration time.Duration) *CreateChatInviteLink {
	if ccil.opts.RequestOpts == nil {
		ccil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ccil.opts.RequestOpts.Timeout = duration

	return ccil
}

// APIURL sets a custom API URL for this request.
func (ccil *CreateChatInviteLink) APIURL(url g.String) *CreateChatInviteLink {
	if ccil.opts.RequestOpts == nil {
		ccil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ccil.opts.RequestOpts.APIURL = url.Std()

	return ccil
}

// Send creates the chat invite link and returns the result.
func (ccil *CreateChatInviteLink) Send() g.Result[*gotgbot.ChatInviteLink] {
	chatID := ccil.chatID.UnwrapOr(ccil.ctx.EffectiveChat.Id)
	return g.ResultOf(ccil.ctx.Bot.Raw().CreateChatInviteLink(chatID, ccil.opts))
}
