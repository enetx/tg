package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/types/permissions"
)

type Restrict struct {
	ctx             *Context
	opts            *gotgbot.RestrictChatMemberOpts
	permissions     *gotgbot.ChatPermissions
	autoPermissions bool
	userID          int64
	chatID          Option[int64]
}

// ChatID sets the target chat ID for the restrict action.
func (c *Restrict) ChatID(id int64) *Restrict {
	c.chatID = Some(id)
	return c
}

// Until sets the restriction expiration time.
func (c *Restrict) Until(t time.Time) *Restrict {
	c.opts.UntilDate = t.Unix()
	return c
}

// For sets the restriction duration from now.
func (c *Restrict) For(d time.Duration) *Restrict {
	return c.Until(time.Now().Add(d))
}

// AutoPermissions uses chat default permissions instead of independent permissions.
func (c *Restrict) AutoPermissions() *Restrict {
	c.autoPermissions = true
	return c
}

// Permissions sets the allowed permissions for the restricted user.
func (c *Restrict) Permissions(perms ...permissions.Permission) *Restrict {
	c.permissions = permissions.Permissions(perms...)
	return c
}

// Timeout sets a custom timeout for this request.
func (c *Restrict) Timeout(duration time.Duration) *Restrict {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Restrict) APIURL(url String) *Restrict {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send restricts the user's permissions and returns the result.
func (c *Restrict) Send() Result[bool] {
	if c.permissions == nil {
		return Err[bool](Errorf("permissions are required"))
	}

	chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
	c.opts.UseIndependentChatPermissions = !c.autoPermissions

	return ResultOf(c.ctx.Bot.Raw().RestrictChatMember(chatID, c.userID, *c.permissions, c.opts))
}
