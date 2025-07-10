package tg

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type BanChatMemberOpts struct {
	until          time.Time
	revokeMessages bool
}

func NewBanChatMemberOpts() *BanChatMemberOpts { return new(BanChatMemberOpts) }

func (o *BanChatMemberOpts) RevokeMessages() *BanChatMemberOpts {
	o.revokeMessages = true
	return o
}

// opts := tg.NewBanChatMemberOpts().Until(time.Now().Add(24 * time.Hour)) // 1 day
func (o *BanChatMemberOpts) Until(date time.Time) *BanChatMemberOpts {
	o.until = date
	return o
}

// opts := tg.NewBanChatMemberOpts().For(7 * 24 * time.Hour) // 7 days
func (o *BanChatMemberOpts) For(duration time.Duration) *BanChatMemberOpts {
	return o.Until(time.Now().Add(duration))
}

func (o *BanChatMemberOpts) std() *gotgbot.BanChatMemberOpts {
	if o == nil {
		return nil
	}

	return &gotgbot.BanChatMemberOpts{
		UntilDate:      o.until.Unix(),
		RevokeMessages: o.revokeMessages,
	}
}

func (c *Context) Ban(userID int64, opts ...*BanChatMemberOpts) Result[bool] {
	var opt *BanChatMemberOpts
	if len(opts) > 0 {
		opt = opts[0]
	}

	return ResultOf(c.Bot.Raw.BanChatMember(c.EffectiveChat.Id, userID, opt.std()))
}

func (c *Context) Unban(userID Int) Result[bool] {
	return ResultOf(c.Bot.Raw.UnbanChatMember(c.EffectiveChat.Id, userID.Int64(),
		&gotgbot.UnbanChatMemberOpts{
			OnlyIfBanned: true,
		},
	))
}
