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

func (p *Photo) After(duration time.Duration) *Photo {
	p.after = Some(duration)
	return p
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

	return p.ctx.timers(p.after, p.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := p.chatID.UnwrapOr(p.ctx.EffectiveChat.Id)
		return ResultOf(p.ctx.Bot.Raw().SendPhoto(chatID, p.doc, p.opts))
	})
}
