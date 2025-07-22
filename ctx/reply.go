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

type Reply struct {
	ctx         *Context
	text        String
	opts        *gotgbot.SendMessageOpts
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
}

// Entities sets custom entities for the reply text.
func (c *Reply) Entities(e *entities.Entities) *Reply {
	c.opts.Entities = e.Std()
	return c
}

// After schedules the reply to be sent after the specified duration.
func (c *Reply) After(duration time.Duration) *Reply {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the reply message to be deleted after the specified duration.
func (c *Reply) DeleteAfter(duration time.Duration) *Reply {
	c.deleteAfter = Some(duration)
	return c
}

// HTML sets the reply parse mode to HTML.
func (c *Reply) HTML() *Reply {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the reply parse mode to MarkdownV2.
func (c *Reply) Markdown() *Reply {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the reply message.
func (c *Reply) Silent() *Reply {
	c.opts.DisableNotification = true
	return c
}

// Effect sets a message effect for the reply.
func (c *Reply) Effect(effect effects.EffectType) *Reply {
	c.opts.MessageEffectId = effect.String()
	return c
}

// ReplyTo sets a different message ID to reply to.
func (c *Reply) ReplyTo(id int64) *Reply {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: id}
	return c
}

// Markup sets the reply markup keyboard for the reply message.
func (c *Reply) Markup(kb keyboard.KeyboardBuilder) *Reply {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// AllowPaidBroadcast allows the reply to be sent in paid broadcast channels.
func (c *Reply) AllowPaidBroadcast() *Reply {
	c.opts.AllowPaidBroadcast = true
	return c
}

// Thread sets the message thread ID for the reply.
func (c *Reply) Thread(id int64) *Reply {
	c.opts.MessageThreadId = id
	return c
}

// ForceReply forces users to reply to this message.
func (c *Reply) ForceReply() *Reply {
	c.opts.ReplyMarkup = gotgbot.ForceReply{ForceReply: true}
	return c
}

// RemoveKeyboard removes the custom keyboard.
func (c *Reply) RemoveKeyboard() *Reply {
	c.opts.ReplyMarkup = gotgbot.ReplyKeyboardRemove{RemoveKeyboard: true}
	return c
}

// Preview sets link preview options for the reply.
func (c *Reply) Preview(preview *preview.Preview) *Reply {
	c.opts.LinkPreviewOptions = preview.Std()
	return c
}

// Business sets the business connection ID for the reply.
func (c *Reply) Business(id String) *Reply {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Protect enables content protection for the reply message.
func (c *Reply) Protect() *Reply {
	c.opts.ProtectContent = true
	return c
}

// Timeout sets a custom timeout for this request.
func (c *Reply) Timeout(duration time.Duration) *Reply {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Reply) APIURL(url String) *Reply {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send sends the reply message and returns the result.
func (c *Reply) Send() Result[*gotgbot.Message] {
	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		return ResultOf(c.ctx.EffectiveMessage.Reply(c.ctx.Bot.Raw(), c.text.Std(), c.opts))
	})
}
