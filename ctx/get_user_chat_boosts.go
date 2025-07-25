package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// GetUserChatBoosts represents a request to get user chat boosts.
type GetUserChatBoosts struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.GetUserChatBoostsOpts
	chatID Option[int64]
}

// ChatID sets the target chat ID.
func (gucb *GetUserChatBoosts) ChatID(chatID int64) *GetUserChatBoosts {
	gucb.chatID = Some(chatID)
	return gucb
}

// Timeout sets a custom timeout for this request.
func (gucb *GetUserChatBoosts) Timeout(duration time.Duration) *GetUserChatBoosts {
	if gucb.opts.RequestOpts == nil {
		gucb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gucb.opts.RequestOpts.Timeout = duration

	return gucb
}

// APIURL sets a custom API URL for this request.
func (gucb *GetUserChatBoosts) APIURL(url String) *GetUserChatBoosts {
	if gucb.opts.RequestOpts == nil {
		gucb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gucb.opts.RequestOpts.APIURL = url.Std()

	return gucb
}

// Send gets the user chat boosts.
func (gucb *GetUserChatBoosts) Send() Result[*gotgbot.UserChatBoosts] {
	return ResultOf(gucb.ctx.Bot.Raw().GetUserChatBoosts(
		gucb.chatID.UnwrapOr(gucb.ctx.EffectiveChat.Id),
		gucb.userID,
		gucb.opts,
	))
}