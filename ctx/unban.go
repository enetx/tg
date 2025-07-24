package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type Unban struct {
	ctx    *Context
	opts   *gotgbot.UnbanChatMemberOpts
	userID int64
	chatID Option[int64]
}

// ChatID sets the target chat ID for the unban action.
func (u *Unban) ChatID(id int64) *Unban {
	u.chatID = Some(id)
	return u
}

// OnlyIfBanned only unbans the user if they are currently banned.
func (u *Unban) OnlyIfBanned() *Unban {
	u.opts.OnlyIfBanned = true
	return u
}

// Timeout sets a custom timeout for this request.
func (u *Unban) Timeout(duration time.Duration) *Unban {
	if u.opts.RequestOpts == nil {
		u.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	u.opts.RequestOpts.Timeout = duration

	return u
}

// APIURL sets a custom API URL for this request.
func (u *Unban) APIURL(url String) *Unban {
	if u.opts.RequestOpts == nil {
		u.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	u.opts.RequestOpts.APIURL = url.Std()

	return u
}

// Send executes the unban action and returns the result.
func (u *Unban) Send() Result[bool] {
	chatID := u.chatID.UnwrapOr(u.ctx.EffectiveChat.Id)
	return ResultOf(u.ctx.Bot.Raw().UnbanChatMember(chatID, u.userID, u.opts))
}
