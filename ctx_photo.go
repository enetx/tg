package tg

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
	deleteAfter Option[time.Duration]
	err         error
}

func (p *Photo) DeleteAfter(duration time.Duration) *Photo {
	p.deleteAfter = Some(duration)
	return p
}

func (p *Photo) Caption(caption String) *Photo {
	p.opts.Caption = caption.Std()
	return p
}

func (p *Photo) HTML() *Photo {
	p.opts.ParseMode = "HTML"
	return p
}

func (p *Photo) Markdown() *Photo {
	p.opts.ParseMode = "Markdown"
	return p
}

func (p *Photo) Silent() *Photo {
	p.opts.DisableNotification = true
	return p
}

func (p *Photo) Protect() *Photo {
	p.opts.ProtectContent = true
	return p
}

func (p *Photo) Markup(kb keyboard.KeyboardBuilder) *Photo {
	p.opts.ReplyMarkup = kb.Markup()
	return p
}

func (p *Photo) ReplyTo(messageID int64) *Photo {
	p.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return p
}

func (p *Photo) Timeout(duration time.Duration) *Photo {
	p.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return p
}

func (p *Photo) To(chatID int64) *Photo {
	p.chatID = Some(chatID)
	return p
}

func (p *Photo) Send() Result[*gotgbot.Message] {
	if p.err != nil {
		return Err[*gotgbot.Message](p.err)
	}

	if p.file != nil {
		defer p.file.Close()
	}

	chatID := p.chatID.UnwrapOr(p.ctx.EffectiveChat.Id)
	msg := ResultOf(p.ctx.Bot.Raw.SendPhoto(chatID, p.doc, p.opts))

	if msg.IsOk() && p.deleteAfter.IsSome() {
		p.ctx.Delete().MessageID(msg.Ok().MessageId).After(p.deleteAfter.Some()).Send()
	}

	return msg
}
