package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"

	. "github.com/enetx/g"
)

type Ban struct {
	ctx    *Context
	opts   *gotgbot.BanChatMemberOpts
	userID int64
	chatID Option[int64]
}

func (b *Ban) ChatID(id int64) *Ban {
	b.chatID = Some(id)
	return b
}

func (b *Ban) RevokeMessages() *Ban {
	b.opts.RevokeMessages = true
	return b
}

func (b *Ban) Until(date time.Time) *Ban {
	b.opts.UntilDate = date.Unix()
	return b
}

func (b *Ban) For(duration time.Duration) *Ban {
	return b.Until(time.Now().Add(duration))
}

func (b *Ban) Timeout(duration time.Duration) *Ban {
	b.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return b
}

func (b *Ban) Send() Result[bool] {
	chatID := b.chatID.UnwrapOr(b.ctx.EffectiveChat.Id)
	return ResultOf(b.ctx.Bot.Raw().BanChatMember(chatID, b.userID, b.opts))
}
