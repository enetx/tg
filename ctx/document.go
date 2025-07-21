package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Document struct {
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

// After schedules the document to be sent after the specified duration.
func (d *Document) After(duration time.Duration) *Document {
	d.after = Some(duration)
	return d
}

// DeleteAfter schedules the document message to be deleted after the specified duration.
func (d *Document) DeleteAfter(duration time.Duration) *Document {
	d.deleteAfter = Some(duration)
	return d
}

// Caption sets the caption text for the document.
func (d *Document) Caption(caption String) *Document {
	d.opts.Caption = caption.Std()
	return d
}

// HTML sets the caption parse mode to HTML.
func (d *Document) HTML() *Document {
	d.opts.ParseMode = "HTML"
	return d
}

// Markdown sets the caption parse mode to MarkdownV2.
func (d *Document) Markdown() *Document {
	d.opts.ParseMode = "MarkdownV2"
	return d
}

// Silent disables notification for the document message.
func (d *Document) Silent() *Document {
	d.opts.DisableNotification = true
	return d
}

// Protect enables content protection for the document message.
func (d *Document) Protect() *Document {
	d.opts.ProtectContent = true
	return d
}

// Markup sets the reply markup keyboard for the document message.
func (d *Document) Markup(kb keyboard.KeyboardBuilder) *Document {
	d.opts.ReplyMarkup = kb.Markup()
	return d
}

// Thumbnail sets a custom thumbnail for the document.
func (d *Document) Thumbnail(file String) *Document {
	d.thumb = NewFile(file)

	reader := d.thumb.Open()
	if reader.IsErr() {
		d.err = reader.Err()
		return d
	}

	d.opts.Thumbnail = gotgbot.InputFileByReader(d.thumb.Name().Std(), reader.Ok().Std())
	return d
}

// ReplyTo sets the message ID to reply to.
func (d *Document) ReplyTo(messageID int64) *Document {
	d.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return d
}

// Timeout sets the request timeout duration.
func (d *Document) Timeout(duration time.Duration) *Document {
	d.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return d
}

// Business sets the business connection ID for the document message.
func (d *Document) Business(id String) *Document {
	d.opts.BusinessConnectionId = id.Std()
	return d
}

// Thread sets the message thread ID for the document message.
func (d *Document) Thread(id int64) *Document {
	d.opts.MessageThreadId = id
	return d
}

// DisableContentTypeDetection disables automatic content type detection for the document.
func (d *Document) DisableContentTypeDetection() *Document {
	d.opts.DisableContentTypeDetection = true
	return d
}

// To sets the target chat ID for the document message.
func (d *Document) To(chatID int64) *Document {
	d.chatID = Some(chatID)
	return d
}

// Send sends the document message to Telegram and returns the result.
func (d *Document) Send() Result[*gotgbot.Message] {
	if d.err != nil {
		return Err[*gotgbot.Message](d.err)
	}

	if d.file != nil {
		defer d.file.Close()
	}

	if d.thumb != nil {
		defer d.thumb.Close()
	}

	return d.ctx.timers(d.after, d.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := d.chatID.UnwrapOr(d.ctx.EffectiveChat.Id)
		return ResultOf(d.ctx.Bot.Raw().SendDocument(chatID, d.doc, d.opts))
	})
}
