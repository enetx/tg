package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
	"github.com/enetx/tg/types/effects"
)

type SendMessage struct {
	ctx         *Context
	text        String
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	opts        *gotgbot.SendMessageOpts
}

// Entities sets special entities in the message text using Entities builder.
func (c *SendMessage) Entities(e *entities.Entities) *SendMessage {
	c.opts.Entities = e.Std()
	return c
}

// After schedules the message to be sent after the specified duration.
func (c *SendMessage) After(duration time.Duration) *SendMessage {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the message to be deleted after the specified duration.
func (c *SendMessage) DeleteAfter(duration time.Duration) *SendMessage {
	c.deleteAfter = Some(duration)
	return c
}

// To sets the target chat ID for the message.
func (c *SendMessage) To(chatID int64) *SendMessage {
	c.chatID = Some(chatID)
	return c
}

// HTML sets the message parse mode to HTML.
func (c *SendMessage) HTML() *SendMessage {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the message parse mode to MarkdownV2.
func (c *SendMessage) Markdown() *SendMessage {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the message.
func (c *SendMessage) Silent() *SendMessage {
	c.opts.DisableNotification = true
	return c
}

// Effect sets a message effect for the message.
func (c *SendMessage) Effect(effect effects.EffectType) *SendMessage {
	c.opts.MessageEffectId = effect.String()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendMessage) ReplyTo(messageID int64) *SendMessage {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Markup sets the reply markup keyboard for the message.
func (c *SendMessage) Markup(kb keyboard.KeyboardBuilder) *SendMessage {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (c *SendMessage) AllowPaidBroadcast() *SendMessage {
	c.opts.AllowPaidBroadcast = true
	return c
}

// Thread sets the message thread ID for the message.
func (c *SendMessage) Thread(id int64) *SendMessage {
	c.opts.MessageThreadId = id
	return c
}

// ForceReply forces users to reply to the message.
func (c *SendMessage) ForceReply() *SendMessage {
	c.opts.ReplyMarkup = gotgbot.ForceReply{ForceReply: true}
	return c
}

// RemoveKeyboard removes the custom keyboard.
func (c *SendMessage) RemoveKeyboard() *SendMessage {
	c.opts.ReplyMarkup = gotgbot.ReplyKeyboardRemove{RemoveKeyboard: true}
	return c
}

// Preview sets link preview options for the message.
func (c *SendMessage) Preview(p *preview.Preview) *SendMessage {
	c.opts.LinkPreviewOptions = p.Std()
	return c
}

// Business sets the business connection ID for the message.
func (c *SendMessage) Business(id String) *SendMessage {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Protect enables content protection for the message.
func (c *SendMessage) Protect() *SendMessage {
	c.opts.ProtectContent = true
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendMessage) Timeout(duration time.Duration) *SendMessage {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendMessage) APIURL(url String) *SendMessage {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send sends the message to Telegram and returns the result.
func (c *SendMessage) Send() Result[*gotgbot.Message] {
	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendMessage(chatID, c.text.Std(), c.opts))
	})
}
