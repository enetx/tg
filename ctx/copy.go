package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/g/ref"
	"github.com/enetx/tg/keyboard"
)

type Copy struct {
	ctx         *Context
	fromChatID  int64
	messageID   int64
	opts        *gotgbot.CopyMessageOpts
	toChatID    Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
}

// After schedules the copy to be sent after the specified duration.
func (c *Copy) After(duration time.Duration) *Copy {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the copied message to be deleted after the specified duration.
func (c *Copy) DeleteAfter(duration time.Duration) *Copy {
	c.deleteAfter = Some(duration)
	return c
}

// Caption sets a new caption for the copied message.
func (c *Copy) Caption(caption String) *Copy {
	c.opts.Caption = ref.Of(caption.Std())
	return c
}

// HTML sets the caption parse mode to HTML.
func (c *Copy) HTML() *Copy {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the caption parse mode to MarkdownV2.
func (c *Copy) Markdown() *Copy {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the copied message.
func (c *Copy) Silent() *Copy {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the copied message.
func (c *Copy) Protect() *Copy {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the copied message.
func (c *Copy) Markup(kb keyboard.KeyboardBuilder) *Copy {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *Copy) ReplyTo(messageID int64) *Copy {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets the request timeout duration.
func (c *Copy) Timeout(duration time.Duration) *Copy {
	c.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return c
}

// To sets the target chat ID for the copied message.
func (c *Copy) To(chatID int64) *Copy {
	c.toChatID = Some(chatID)
	return c
}

// Send copies the message to the target chat and returns the result.
func (c *Copy) Send() Result[*gotgbot.MessageId] {
	if c.after.IsSome() {
		go func() {
			<-time.After(c.after.Some())
			chatID := c.toChatID.UnwrapOr(c.ctx.EffectiveChat.Id)
			msgID, err := c.ctx.Bot.Raw().CopyMessage(chatID, c.fromChatID, c.messageID, c.opts)
			if err == nil && msgID != nil && c.deleteAfter.IsSome() {
				c.ctx.Delete().MessageID(msgID.MessageId).ChatID(chatID).After(c.deleteAfter.Some()).Send()
			}
		}()

		return Ok[*gotgbot.MessageId](nil)
	}

	chatID := c.toChatID.UnwrapOr(c.ctx.EffectiveChat.Id)
	result := ResultOf(c.ctx.Bot.Raw().CopyMessage(chatID, c.fromChatID, c.messageID, c.opts))

	if result.IsOk() && c.deleteAfter.IsSome() {
		c.ctx.Delete().MessageID(result.Ok().MessageId).ChatID(chatID).After(c.deleteAfter.Some()).Send()
	}

	return result
}
