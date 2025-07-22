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
func (c *Unban) ChatID(id int64) *Unban {
	c.chatID = Some(id)
	return c
}

// OnlyIfBanned only unbans the user if they are currently banned.
func (c *Unban) OnlyIfBanned() *Unban {
	c.opts.OnlyIfBanned = true
	return c
}

// Timeout sets a custom timeout for this request.
func (c *Unban) Timeout(duration time.Duration) *Unban {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Unban) APIURL(url String) *Unban {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send executes the unban action and returns the result.
func (c *Unban) Send() Result[bool] {
	chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
	return ResultOf(c.ctx.Bot.Raw().UnbanChatMember(chatID, c.userID, c.opts))
}
