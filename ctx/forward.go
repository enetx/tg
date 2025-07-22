package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type Forward struct {
	ctx         *Context
	fromChatID  int64
	messageID   int64
	opts        *gotgbot.ForwardMessageOpts
	toChatID    Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
}

// After schedules the forward to be sent after the specified duration.
func (c *Forward) After(duration time.Duration) *Forward {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the forwarded message to be deleted after the specified duration.
func (c *Forward) DeleteAfter(duration time.Duration) *Forward {
	c.deleteAfter = Some(duration)
	return c
}

// Silent disables notification for the forwarded message.
func (c *Forward) Silent() *Forward {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the forwarded message.
func (c *Forward) Protect() *Forward {
	c.opts.ProtectContent = true
	return c
}

// Timeout sets a custom timeout for this request.
func (c *Forward) Timeout(duration time.Duration) *Forward {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Forward) APIURL(url String) *Forward {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// To sets the target chat ID for the forwarded message.
func (c *Forward) To(chatID int64) *Forward {
	c.toChatID = Some(chatID)
	return c
}

// Send forwards the message to the target chat and returns the result.
func (c *Forward) Send() Result[*gotgbot.Message] {
	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.toChatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().ForwardMessage(chatID, c.fromChatID, c.messageID, c.opts))
	})
}
