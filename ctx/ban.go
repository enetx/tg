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
func (c *Ban) ChatID(id int64) *Ban {
	c.chatID = Some(id)
	return c
}

// RevokeMessages revokes all messages sent by the user when banning.
func (c *Ban) RevokeMessages() *Ban {
	c.opts.RevokeMessages = true
	return c
}

// Until sets the ban expiration date.
func (c *Ban) Until(date time.Time) *Ban {
	c.opts.UntilDate = date.Unix()
	return c
}

// For sets the ban duration from now.
func (c *Ban) For(duration time.Duration) *Ban {
	return c.Until(time.Now().Add(duration))
}

// Timeout sets a custom timeout for this request.
func (c *Ban) Timeout(duration time.Duration) *Ban {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Ban) APIURL(url String) *Ban {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send executes the ban action and returns the result.
func (c *Ban) Send() Result[bool] {
	chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
	return ResultOf(c.ctx.Bot.Raw().BanChatMember(chatID, c.userID, c.opts))
}
