package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// EditChatSubscriptionInviteLink represents a request to edit a subscription invite link.
type EditChatSubscriptionInviteLink struct {
	ctx        *Context
	inviteLink g.String
	opts       *gotgbot.EditChatSubscriptionInviteLinkOpts
	chatID     g.Option[int64]
}

// ChatID sets the target chat ID.
func (ecsil *EditChatSubscriptionInviteLink) ChatID(chatID int64) *EditChatSubscriptionInviteLink {
	ecsil.chatID = g.Some(chatID)
	return ecsil
}

// Name sets the invite link name (0-32 characters).
func (ecsil *EditChatSubscriptionInviteLink) Name(name g.String) *EditChatSubscriptionInviteLink {
	ecsil.opts.Name = name.Std()
	return ecsil
}

// Timeout sets a custom timeout for this request.
func (ecsil *EditChatSubscriptionInviteLink) Timeout(duration time.Duration) *EditChatSubscriptionInviteLink {
	if ecsil.opts.RequestOpts == nil {
		ecsil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ecsil.opts.RequestOpts.Timeout = duration

	return ecsil
}

// APIURL sets a custom API URL for this request.
func (ecsil *EditChatSubscriptionInviteLink) APIURL(url g.String) *EditChatSubscriptionInviteLink {
	if ecsil.opts.RequestOpts == nil {
		ecsil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ecsil.opts.RequestOpts.APIURL = url.Std()

	return ecsil
}

// Send edits the subscription invite link.
func (ecsil *EditChatSubscriptionInviteLink) Send() g.Result[*gotgbot.ChatInviteLink] {
	return g.ResultOf(ecsil.ctx.Bot.Raw().EditChatSubscriptionInviteLink(
		ecsil.chatID.UnwrapOr(ecsil.ctx.EffectiveChat.Id),
		ecsil.inviteLink.Std(),
		ecsil.opts,
	))
}
