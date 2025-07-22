package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/roles"
)

type Promote struct {
	ctx    *Context
	opts   *gotgbot.PromoteChatMemberOpts
	roles  bool
	userID int64
	chatID Option[int64]
}

// ChatID sets the target chat ID for the promote action.
func (c *Promote) ChatID(id int64) *Promote {
	c.chatID = Some(id)
	return c
}

// Roles sets the administrator roles to grant to the user.
func (c *Promote) Roles(r ...roles.Role) *Promote {
	c.opts = roles.Roles(r...)
	c.roles = true

	return c
}

// Timeout sets a custom timeout for this request.
func (c *Promote) Timeout(duration time.Duration) *Promote {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Promote) APIURL(url String) *Promote {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send promotes the user to administrator and returns the result.
func (c *Promote) Send() Result[bool] {
	if !c.roles {
		return Err[bool](Errorf("roles are required"))
	}

	chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
	return ResultOf(c.ctx.Bot.Raw().PromoteChatMember(chatID, c.userID, c.opts))
}
