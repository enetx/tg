package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type SendContact struct {
	ctx         *Context
	phoneNumber String
	firstName   String
	opts        *gotgbot.SendContactOpts
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
}

// After schedules the contact to be sent after the specified duration.
func (c *SendContact) After(duration time.Duration) *SendContact {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the contact message to be deleted after the specified duration.
func (c *SendContact) DeleteAfter(duration time.Duration) *SendContact {
	c.deleteAfter = Some(duration)
	return c
}

// Silent disables notification for the contact message.
func (c *SendContact) Silent() *SendContact {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the contact message.
func (c *SendContact) Protect() *SendContact {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the contact message.
func (c *SendContact) Markup(kb keyboard.KeyboardBuilder) *SendContact {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// LastName sets the contact's last name.
func (c *SendContact) LastName(lastName String) *SendContact {
	c.opts.LastName = lastName.Std()
	return c
}

// VCard sets additional contact information in vCard format.
func (c *SendContact) VCard(vcard String) *SendContact {
	c.opts.Vcard = vcard.Std()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendContact) ReplyTo(messageID int64) *SendContact {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendContact) Timeout(duration time.Duration) *SendContact {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendContact) APIURL(url String) *SendContact {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the contact message.
func (c *SendContact) Business(id String) *SendContact {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the contact message.
func (c *SendContact) Thread(id int64) *SendContact {
	c.opts.MessageThreadId = id
	return c
}

// To sets the target chat ID for the contact message.
func (c *SendContact) To(chatID int64) *SendContact {
	c.chatID = Some(chatID)
	return c
}

// Send sends the contact message to Telegram and returns the result.
func (c *SendContact) Send() Result[*gotgbot.Message] {
	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendContact(chatID, c.phoneNumber.Std(), c.firstName.Std(), c.opts))
	})
}
