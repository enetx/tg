package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"

	"github.com/enetx/tg/keyboard"

	. "github.com/enetx/g"
)

type Audio struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendAudioOpts
	file        *File
	thumb       *File
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

func (a *Audio) After(duration time.Duration) *Audio {
	a.after = Some(duration)
	return a
}

func (a *Audio) DeleteAfter(duration time.Duration) *Audio {
	a.deleteAfter = Some(duration)
	return a
}

func (a *Audio) Caption(caption String) *Audio {
	a.opts.Caption = caption.Std()
	return a
}

func (a *Audio) HTML() *Audio {
	a.opts.ParseMode = "HTML"
	return a
}

func (a *Audio) Markdown() *Audio {
	a.opts.ParseMode = "Markdown"
	return a
}

func (a *Audio) Silent() *Audio {
	a.opts.DisableNotification = true
	return a
}

func (a *Audio) Protect() *Audio {
	a.opts.ProtectContent = true
	return a
}

func (a *Audio) Markup(kb keyboard.KeyboardBuilder) *Audio {
	a.opts.ReplyMarkup = kb.Markup()
	return a
}

func (a *Audio) Thumbnail(file String) *Audio {
	a.thumb = NewFile(file)

	reader := a.thumb.Open()
	if reader.IsErr() {
		a.err = reader.Err()
		return a
	}

	a.opts.Thumbnail = gotgbot.InputFileByReader(a.thumb.Name().Std(), reader.Ok().Std())
	return a
}

func (a *Audio) ReplyTo(messageID int64) *Audio {
	a.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return a
}

func (a *Audio) Timeout(duration time.Duration) *Audio {
	a.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return a
}

func (a *Audio) To(chatID int64) *Audio {
	a.chatID = Some(chatID)
	return a
}

func (a *Audio) Send() Result[*gotgbot.Message] {
	if a.err != nil {
		return Err[*gotgbot.Message](a.err)
	}

	if a.file != nil {
		defer a.file.Close()
	}

	if a.thumb != nil {
		defer a.thumb.Close()
	}

	return a.ctx.timers(a.after, a.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := a.chatID.UnwrapOr(a.ctx.EffectiveChat.Id)
		return ResultOf(a.ctx.Bot.Raw().SendAudio(chatID, a.doc, a.opts))
	})
}
