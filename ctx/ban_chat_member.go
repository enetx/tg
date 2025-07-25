package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type BanChatMember struct {
	ctx    *Context
	opts   *gotgbot.BanChatMemberOpts
	userID int64
	chatID Option[int64]
}

// ChatID sets the target chat ID for the ban action.
func (b *BanChatMember) ChatID(id int64) *BanChatMember {
	b.chatID = Some(id)
	return b
}

// RevokeMessages revokes all messages sent by the user when banning.
func (b *BanChatMember) RevokeMessages() *BanChatMember {
	b.opts.RevokeMessages = true
	return b
}

// Until sets the ban expiration date.
func (b *BanChatMember) Until(date time.Time) *BanChatMember {
	b.opts.UntilDate = date.Unix()
	return b
}

// For sets the ban duration from now.
func (b *BanChatMember) For(duration time.Duration) *BanChatMember {
	return b.Until(time.Now().Add(duration))
}

// Timeout sets a custom timeout for this request.
func (b *BanChatMember) Timeout(duration time.Duration) *BanChatMember {
	if b.opts.RequestOpts == nil {
		b.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	b.opts.RequestOpts.Timeout = duration

	return b
}

// APIURL sets a custom API URL for this request.
func (b *BanChatMember) APIURL(url String) *BanChatMember {
	if b.opts.RequestOpts == nil {
		b.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	b.opts.RequestOpts.APIURL = url.Std()

	return b
}

// Send executes the ban action and returns the result.
func (b *BanChatMember) Send() Result[bool] {
	chatID := b.chatID.UnwrapOr(b.ctx.EffectiveChat.Id)
	return ResultOf(b.ctx.Bot.Raw().BanChatMember(chatID, b.userID, b.opts))
}
