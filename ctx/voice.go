package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

type SendVoice struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendVoiceOpts
	file        *File
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// CaptionEntities sets custom entities for the voice caption.
func (c *SendVoice) CaptionEntities(e *entities.Entities) *SendVoice {
	c.opts.CaptionEntities = e.Std()
	return c
}

// After schedules the voice message to be sent after the specified duration.
func (c *SendVoice) After(duration time.Duration) *SendVoice {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the voice message to be deleted after the specified duration.
func (c *SendVoice) DeleteAfter(duration time.Duration) *SendVoice {
	c.deleteAfter = Some(duration)
	return c
}

// Caption sets the caption text for the voice message.
func (c *SendVoice) Caption(caption String) *SendVoice {
	c.opts.Caption = caption.Std()
	return c
}

// HTML sets the caption parse mode to HTML.
func (c *SendVoice) HTML() *SendVoice {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the caption parse mode to MarkdownV2.
func (c *SendVoice) Markdown() *SendVoice {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the voice message.
func (c *SendVoice) Silent() *SendVoice {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the voice message.
func (c *SendVoice) Protect() *SendVoice {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the voice message.
func (c *SendVoice) Markup(kb keyboard.KeyboardBuilder) *SendVoice {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// Duration sets the voice message duration in seconds.
func (c *SendVoice) Duration(duration int64) *SendVoice {
	c.opts.Duration = duration
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendVoice) ReplyTo(messageID int64) *SendVoice {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendVoice) Timeout(duration time.Duration) *SendVoice {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendVoice) APIURL(url String) *SendVoice {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the voice message.
func (c *SendVoice) Business(id String) *SendVoice {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the voice message.
func (c *SendVoice) Thread(id int64) *SendVoice {
	c.opts.MessageThreadId = id
	return c
}

// To sets the target chat ID for the voice message.
func (c *SendVoice) To(chatID int64) *SendVoice {
	c.chatID = Some(chatID)
	return c
}

// Send sends the voice message to Telegram and returns the result.
func (c *SendVoice) Send() Result[*gotgbot.Message] {
	if c.err != nil {
		return Err[*gotgbot.Message](c.err)
	}

	if c.file != nil {
		defer c.file.Close()
	}

	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendVoice(chatID, c.doc, c.opts))
	})
}
