package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

type SendDocument struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendDocumentOpts
	file        *File
	thumb       *File
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// CaptionEntities sets custom entities for the document caption.
func (c *SendDocument) CaptionEntities(e *entities.Entities) *SendDocument {
	c.opts.CaptionEntities = e.Std()
	return c
}

// After schedules the document to be sent after the specified duration.
func (c *SendDocument) After(duration time.Duration) *SendDocument {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the document message to be deleted after the specified duration.
func (c *SendDocument) DeleteAfter(duration time.Duration) *SendDocument {
	c.deleteAfter = Some(duration)
	return c
}

// Caption sets the caption text for the document.
func (c *SendDocument) Caption(caption String) *SendDocument {
	c.opts.Caption = caption.Std()
	return c
}

// HTML sets the caption parse mode to HTML.
func (c *SendDocument) HTML() *SendDocument {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the caption parse mode to MarkdownV2.
func (c *SendDocument) Markdown() *SendDocument {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the document message.
func (c *SendDocument) Silent() *SendDocument {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the document message.
func (c *SendDocument) Protect() *SendDocument {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the document message.
func (c *SendDocument) Markup(kb keyboard.KeyboardBuilder) *SendDocument {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// Thumbnail sets a custom thumbnail for the document.
func (c *SendDocument) Thumbnail(file String) *SendDocument {
	c.thumb = NewFile(file)

	reader := c.thumb.Open()
	if reader.IsErr() {
		c.err = reader.Err()
		return c
	}

	c.opts.Thumbnail = gotgbot.InputFileByReader(c.thumb.Name().Std(), reader.Ok().Std())
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendDocument) ReplyTo(messageID int64) *SendDocument {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendDocument) Timeout(duration time.Duration) *SendDocument {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendDocument) APIURL(url String) *SendDocument {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the document message.
func (c *SendDocument) Business(id String) *SendDocument {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the document message.
func (c *SendDocument) Thread(id int64) *SendDocument {
	c.opts.MessageThreadId = id
	return c
}

// DisableContentTypeDetection disables automatic content type detection for the document.
func (c *SendDocument) DisableContentTypeDetection() *SendDocument {
	c.opts.DisableContentTypeDetection = true
	return c
}

// To sets the target chat ID for the document message.
func (c *SendDocument) To(chatID int64) *SendDocument {
	c.chatID = Some(chatID)
	return c
}

// Send sends the document message to Telegram and returns the result.
func (c *SendDocument) Send() Result[*gotgbot.Message] {
	if c.err != nil {
		return Err[*gotgbot.Message](c.err)
	}

	if c.file != nil {
		defer c.file.Close()
	}

	if c.thumb != nil {
		defer c.thumb.Close()
	}

	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendDocument(chatID, c.doc, c.opts))
	})
}
