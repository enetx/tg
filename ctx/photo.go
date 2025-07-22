package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

type SendPhoto struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendPhotoOpts
	file        *File
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// CaptionEntities sets custom entities for the photo caption.
func (c *SendPhoto) CaptionEntities(e *entities.Entities) *SendPhoto {
	c.opts.CaptionEntities = e.Std()
	return c
}

// After schedules the photo to be sent after the specified duration.
func (c *SendPhoto) After(duration time.Duration) *SendPhoto {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the photo message to be deleted after the specified duration.
func (c *SendPhoto) DeleteAfter(duration time.Duration) *SendPhoto {
	c.deleteAfter = Some(duration)
	return c
}

// Caption sets the caption text for the photo.
func (c *SendPhoto) Caption(caption String) *SendPhoto {
	c.opts.Caption = caption.Std()
	return c
}

// HTML sets the caption parse mode to HTML.
func (c *SendPhoto) HTML() *SendPhoto {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the caption parse mode to MarkdownV2.
func (c *SendPhoto) Markdown() *SendPhoto {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the photo message.
func (c *SendPhoto) Silent() *SendPhoto {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the photo message.
func (c *SendPhoto) Protect() *SendPhoto {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the photo message.
func (c *SendPhoto) Markup(kb keyboard.KeyboardBuilder) *SendPhoto {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendPhoto) ReplyTo(messageID int64) *SendPhoto {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendPhoto) Timeout(duration time.Duration) *SendPhoto {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendPhoto) APIURL(url String) *SendPhoto {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the photo message.
func (c *SendPhoto) Business(id String) *SendPhoto {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the photo message.
func (c *SendPhoto) Thread(id int64) *SendPhoto {
	c.opts.MessageThreadId = id
	return c
}

// ShowCaptionAboveMedia displays the caption above the photo instead of below.
func (c *SendPhoto) ShowCaptionAboveMedia() *SendPhoto {
	c.opts.ShowCaptionAboveMedia = true
	return c
}

// To sets the target chat ID for the photo message.
func (c *SendPhoto) To(chatID int64) *SendPhoto {
	c.chatID = Some(chatID)
	return c
}

// Send sends the photo message to Telegram and returns the result.
func (c *SendPhoto) Send() Result[*gotgbot.Message] {
	if c.err != nil {
		return Err[*gotgbot.Message](c.err)
	}

	if c.file != nil {
		defer c.file.Close()
	}

	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendPhoto(chatID, c.doc, c.opts))
	})
}
