package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Contact struct {
	ctx         *Context
	phoneNumber String
	firstName   String
	opts        *gotgbot.SendContactOpts
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
}

// After schedules the contact to be sent after the specified duration.
func (c *Contact) After(duration time.Duration) *Contact {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the contact message to be deleted after the specified duration.
func (c *Contact) DeleteAfter(duration time.Duration) *Contact {
	c.deleteAfter = Some(duration)
	return c
}

// Silent disables notification for the contact message.
func (c *Contact) Silent() *Contact {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the contact message.
func (c *Contact) Protect() *Contact {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the contact message.
func (c *Contact) Markup(kb keyboard.KeyboardBuilder) *Contact {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// LastName sets the contact's last name.
func (c *Contact) LastName(lastName String) *Contact {
	c.opts.LastName = lastName.Std()
	return c
}

// VCard sets additional contact information in vCard format.
func (c *Contact) VCard(vcard String) *Contact {
	c.opts.Vcard = vcard.Std()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *Contact) ReplyTo(messageID int64) *Contact {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets the request timeout duration.
func (c *Contact) Timeout(duration time.Duration) *Contact {
	c.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return c
}

// Business sets the business connection ID for the contact message.
func (c *Contact) Business(id String) *Contact {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the contact message.
func (c *Contact) Thread(id int64) *Contact {
	c.opts.MessageThreadId = id
	return c
}

// To sets the target chat ID for the contact message.
func (c *Contact) To(chatID int64) *Contact {
	c.chatID = Some(chatID)
	return c
}

// Send sends the contact message to Telegram and returns the result.
func (c *Contact) Send() Result[*gotgbot.Message] {
	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendContact(chatID, c.phoneNumber.Std(), c.firstName.Std(), c.opts))
	})
}
