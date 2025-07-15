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

func (u *Unban) ChatID(id int64) *Unban {
	u.chatID = Some(id)
	return u
}

func (u *Unban) OnlyIfBanned() *Unban {
	u.opts.OnlyIfBanned = true
	return u
}

func (u *Unban) Timeout(duration time.Duration) *Unban {
	u.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return u
}

func (u *Unban) Send() Result[bool] {
	chatID := u.chatID.UnwrapOr(u.ctx.EffectiveChat.Id)
	return ResultOf(u.ctx.Bot.Raw().UnbanChatMember(chatID, u.userID, u.opts))
}
