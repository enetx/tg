package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Voice struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendVoiceOpts
	file        *File
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

func (v *Voice) After(duration time.Duration) *Voice {
	v.after = Some(duration)
	return v
}

func (v *Voice) DeleteAfter(duration time.Duration) *Voice {
	v.deleteAfter = Some(duration)
	return v
}

func (v *Voice) Caption(caption String) *Voice {
	v.opts.Caption = caption.Std()
	return v
}

func (v *Voice) HTML() *Voice {
	v.opts.ParseMode = "HTML"
	return v
}

func (v *Voice) Markdown() *Voice {
	v.opts.ParseMode = "MarkdownV2"
	return v
}

func (v *Voice) Silent() *Voice {
	v.opts.DisableNotification = true
	return v
}

func (v *Voice) Protect() *Voice {
	v.opts.ProtectContent = true
	return v
}

func (v *Voice) Markup(kb keyboard.KeyboardBuilder) *Voice {
	v.opts.ReplyMarkup = kb.Markup()
	return v
}

func (v *Voice) Duration(duration int64) *Voice {
	v.opts.Duration = duration
	return v
}

func (v *Voice) ReplyTo(messageID int64) *Voice {
	v.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return v
}

func (v *Voice) Timeout(duration time.Duration) *Voice {
	v.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return v
}

func (v *Voice) Business(id String) *Voice {
	v.opts.BusinessConnectionId = id.Std()
	return v
}

func (v *Voice) Thread(id int64) *Voice {
	v.opts.MessageThreadId = id
	return v
}

func (v *Voice) To(chatID int64) *Voice {
	v.chatID = Some(chatID)
	return v
}

func (v *Voice) Send() Result[*gotgbot.Message] {
	if v.err != nil {
		return Err[*gotgbot.Message](v.err)
	}

	if v.file != nil {
		defer v.file.Close()
	}

	return v.ctx.timers(v.after, v.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := v.chatID.UnwrapOr(v.ctx.EffectiveChat.Id)
		return ResultOf(v.ctx.Bot.Raw().SendVoice(chatID, v.doc, v.opts))
	})
}
