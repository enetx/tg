package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// CreateChatSubscriptionInviteLink represents a request to create a subscription invite link.
type CreateChatSubscriptionInviteLink struct {
	ctx                *Context
	subscriptionPeriod int64
	subscriptionPrice  int64
	opts               *gotgbot.CreateChatSubscriptionInviteLinkOpts
	chatID             g.Option[int64]
}

// ChatID sets the target chat ID.
func (ccsil *CreateChatSubscriptionInviteLink) ChatID(chatID int64) *CreateChatSubscriptionInviteLink {
	ccsil.chatID = g.Some(chatID)
	return ccsil
}

// Name sets the invite link name (0-32 characters).
func (ccsil *CreateChatSubscriptionInviteLink) Name(name g.String) *CreateChatSubscriptionInviteLink {
	ccsil.opts.Name = name.Std()
	return ccsil
}

// Timeout sets a custom timeout for this request.
func (ccsil *CreateChatSubscriptionInviteLink) Timeout(duration time.Duration) *CreateChatSubscriptionInviteLink {
	if ccsil.opts.RequestOpts == nil {
		ccsil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ccsil.opts.RequestOpts.Timeout = duration

	return ccsil
}

// APIURL sets a custom API URL for this request.
func (ccsil *CreateChatSubscriptionInviteLink) APIURL(url g.String) *CreateChatSubscriptionInviteLink {
	if ccsil.opts.RequestOpts == nil {
		ccsil.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ccsil.opts.RequestOpts.APIURL = url.Std()

	return ccsil
}

// Send creates the subscription invite link.
func (ccsil *CreateChatSubscriptionInviteLink) Send() g.Result[*gotgbot.ChatInviteLink] {
	return g.ResultOf(ccsil.ctx.Bot.Raw().CreateChatSubscriptionInviteLink(
		ccsil.chatID.UnwrapOr(ccsil.ctx.EffectiveChat.Id),
		ccsil.subscriptionPeriod,
		ccsil.subscriptionPrice,
		ccsil.opts,
	))
}
