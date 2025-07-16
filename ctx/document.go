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
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

func (d *Document) After(duration time.Duration) *Document {
	d.after = Some(duration)
	return d
}

func (d *Document) DeleteAfter(duration time.Duration) *Document {
	d.deleteAfter = Some(duration)
	return d
}

func (d *Document) Caption(caption String) *Document {
	d.opts.Caption = caption.Std()
	return d
}

func (d *Document) HTML() *Document {
	d.opts.ParseMode = "HTML"
	return d
}

func (d *Document) Markdown() *Document {
	d.opts.ParseMode = "Markdown"
	return d
}

func (d *Document) Silent() *Document {
	d.opts.DisableNotification = true
	return d
}

func (d *Document) Protect() *Document {
	d.opts.ProtectContent = true
	return d
}

func (d *Document) Markup(kb keyboard.KeyboardBuilder) *Document {
	d.opts.ReplyMarkup = kb.Markup()
	return d
}

func (d *Document) Thumbnail(file gotgbot.InputFile) *Document {
	d.opts.Thumbnail = file
	return d
}

func (d *Document) ReplyTo(messageID int64) *Document {
	d.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return d
}

func (d *Document) Timeout(duration time.Duration) *Document {
	d.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return d
}

func (d *Document) To(chatID int64) *Document {
	d.chatID = Some(chatID)
	return d
}

func (d *Document) Send() Result[*gotgbot.Message] {
	if d.err != nil {
		return Err[*gotgbot.Message](d.err)
	}

	if d.file != nil {
		defer d.file.Close()
	}

	return d.ctx.timers(d.after, d.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := d.chatID.UnwrapOr(d.ctx.EffectiveChat.Id)
		return ResultOf(d.ctx.Bot.Raw().SendDocument(chatID, d.doc, d.opts))
	})
}
