package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

type UnbanChatMember struct {
	ctx    *Context
	opts   *gotgbot.UnbanChatMemberOpts
	userID int64
	chatID g.Option[int64]
}

// ChatID sets the target chat ID for the unban action.
func (u *UnbanChatMember) ChatID(id int64) *UnbanChatMember {
	u.chatID = g.Some(id)
	return u
}

// OnlyIfBanned only unbans the user if they are currently banned.
func (u *UnbanChatMember) OnlyIfBanned() *UnbanChatMember {
	u.opts.OnlyIfBanned = true
	return u
}

// Timeout sets a custom timeout for this request.
func (u *UnbanChatMember) Timeout(duration time.Duration) *UnbanChatMember {
	if u.opts.RequestOpts == nil {
		u.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	u.opts.RequestOpts.Timeout = duration

	return u
}

// APIURL sets a custom API URL for this request.
func (u *UnbanChatMember) APIURL(url g.String) *UnbanChatMember {
	if u.opts.RequestOpts == nil {
		u.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	u.opts.RequestOpts.APIURL = url.Std()

	return u
}

// Send executes the unban action and returns the result.
func (u *UnbanChatMember) Send() g.Result[bool] {
	chatID := u.chatID.UnwrapOr(u.ctx.EffectiveChat.Id)
	return g.ResultOf(u.ctx.Bot.Raw().UnbanChatMember(chatID, u.userID, u.opts))
}
