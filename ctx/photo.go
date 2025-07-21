package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Photo struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendPhotoOpts
	file        *File
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// After schedules the photo to be sent after the specified duration.
func (p *Photo) After(duration time.Duration) *Photo {
	p.after = Some(duration)
	return p
}

// DeleteAfter schedules the photo message to be deleted after the specified duration.
func (p *Photo) DeleteAfter(duration time.Duration) *Photo {
	p.deleteAfter = Some(duration)
	return p
}

// Caption sets the caption text for the photo.
func (p *Photo) Caption(caption String) *Photo {
	p.opts.Caption = caption.Std()
	return p
}

// HTML sets the caption parse mode to HTML.
func (p *Photo) HTML() *Photo {
	p.opts.ParseMode = "HTML"
	return p
}

// Markdown sets the caption parse mode to MarkdownV2.
func (p *Photo) Markdown() *Photo {
	p.opts.ParseMode = "MarkdownV2"
	return p
}

// Silent disables notification for the photo message.
func (p *Photo) Silent() *Photo {
	p.opts.DisableNotification = true
	return p
}

// Protect enables content protection for the photo message.
func (p *Photo) Protect() *Photo {
	p.opts.ProtectContent = true
	return p
}

// Markup sets the reply markup keyboard for the photo message.
func (p *Photo) Markup(kb keyboard.KeyboardBuilder) *Photo {
	p.opts.ReplyMarkup = kb.Markup()
	return p
}

// ReplyTo sets the message ID to reply to.
func (p *Photo) ReplyTo(messageID int64) *Photo {
	p.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return p
}

// Timeout sets the request timeout duration.
func (p *Photo) Timeout(duration time.Duration) *Photo {
	p.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return p
}

// Business sets the business connection ID for the photo message.
func (p *Photo) Business(id String) *Photo {
	p.opts.BusinessConnectionId = id.Std()
	return p
}

// Thread sets the message thread ID for the photo message.
func (p *Photo) Thread(id int64) *Photo {
	p.opts.MessageThreadId = id
	return p
}

// ShowCaptionAboveMedia displays the caption above the photo instead of below.
func (p *Photo) ShowCaptionAboveMedia() *Photo {
	p.opts.ShowCaptionAboveMedia = true
	return p
}

// To sets the target chat ID for the photo message.
func (p *Photo) To(chatID int64) *Photo {
	p.chatID = Some(chatID)
	return p
}

// Send sends the photo message to Telegram and returns the result.
func (p *Photo) Send() Result[*gotgbot.Message] {
	if p.err != nil {
		return Err[*gotgbot.Message](p.err)
	}

	if p.file != nil {
		defer p.file.Close()
	}

	return p.ctx.timers(p.after, p.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := p.chatID.UnwrapOr(p.ctx.EffectiveChat.Id)
		return ResultOf(p.ctx.Bot.Raw().SendPhoto(chatID, p.doc, p.opts))
	})
}
