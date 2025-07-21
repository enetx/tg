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

// ChatID sets the target chat ID for the ban action.
func (b *Ban) ChatID(id int64) *Ban {
	b.chatID = Some(id)
	return b
}

// RevokeMessages revokes all messages sent by the user when banning.
func (b *Ban) RevokeMessages() *Ban {
	b.opts.RevokeMessages = true
	return b
}

// Until sets the ban expiration date.
func (b *Ban) Until(date time.Time) *Ban {
	b.opts.UntilDate = date.Unix()
	return b
}

// For sets the ban duration from now.
func (b *Ban) For(duration time.Duration) *Ban {
	return b.Until(time.Now().Add(duration))
}

// Timeout sets the request timeout duration.
func (b *Ban) Timeout(duration time.Duration) *Ban {
	b.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return b
}

// Send executes the ban action and returns the result.
func (b *Ban) Send() Result[bool] {
	chatID := b.chatID.UnwrapOr(b.ctx.EffectiveChat.Id)
	return ResultOf(b.ctx.Bot.Raw().BanChatMember(chatID, b.userID, b.opts))
}
