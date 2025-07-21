package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type VideoNote struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendVideoNoteOpts
	file        *File
	thumb       *File
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

func (vn *VideoNote) After(duration time.Duration) *VideoNote {
	vn.after = Some(duration)
	return vn
}

func (vn *VideoNote) DeleteAfter(duration time.Duration) *VideoNote {
	vn.deleteAfter = Some(duration)
	return vn
}

func (vn *VideoNote) Silent() *VideoNote {
	vn.opts.DisableNotification = true
	return vn
}

func (vn *VideoNote) Protect() *VideoNote {
	vn.opts.ProtectContent = true
	return vn
}

func (vn *VideoNote) Markup(kb keyboard.KeyboardBuilder) *VideoNote {
	vn.opts.ReplyMarkup = kb.Markup()
	return vn
}

func (vn *VideoNote) Duration(duration int64) *VideoNote {
	vn.opts.Duration = duration
	return vn
}

func (vn *VideoNote) Length(length int64) *VideoNote {
	vn.opts.Length = length
	return vn
}

func (vn *VideoNote) Thumbnail(file String) *VideoNote {
	vn.thumb = NewFile(file)

	reader := vn.thumb.Open()
	if reader.IsErr() {
		vn.err = reader.Err()
		return vn
	}

	vn.opts.Thumbnail = gotgbot.InputFileByReader(vn.thumb.Name().Std(), reader.Ok().Std())
	return vn
}

func (vn *VideoNote) ReplyTo(messageID int64) *VideoNote {
	vn.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return vn
}

func (vn *VideoNote) Timeout(duration time.Duration) *VideoNote {
	vn.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return vn
}

func (vn *VideoNote) Business(id String) *VideoNote {
	vn.opts.BusinessConnectionId = id.Std()
	return vn
}

func (vn *VideoNote) Thread(id int64) *VideoNote {
	vn.opts.MessageThreadId = id
	return vn
}

func (vn *VideoNote) To(chatID int64) *VideoNote {
	vn.chatID = Some(chatID)
	return vn
}

func (vn *VideoNote) Send() Result[*gotgbot.Message] {
	if vn.err != nil {
		return Err[*gotgbot.Message](vn.err)
	}

	if vn.file != nil {
		defer vn.file.Close()
	}

	if vn.thumb != nil {
		defer vn.thumb.Close()
	}

	return vn.ctx.timers(vn.after, vn.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := vn.chatID.UnwrapOr(vn.ctx.EffectiveChat.Id)
		return ResultOf(vn.ctx.Bot.Raw().SendVideoNote(chatID, vn.doc, vn.opts))
	})
}
