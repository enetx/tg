package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type SendVideoNote struct {
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

// After schedules the video note to be sent after the specified duration.
func (svn *SendVideoNote) After(duration time.Duration) *SendVideoNote {
	svn.after = Some(duration)
	return svn
}

// DeleteAfter schedules the video note message to be deleted after the specified duration.
func (svn *SendVideoNote) DeleteAfter(duration time.Duration) *SendVideoNote {
	svn.deleteAfter = Some(duration)
	return svn
}

// Silent disables notification for the video note message.
func (svn *SendVideoNote) Silent() *SendVideoNote {
	svn.opts.DisableNotification = true
	return svn
}

// Protect enables content protection for the video note message.
func (svn *SendVideoNote) Protect() *SendVideoNote {
	svn.opts.ProtectContent = true
	return svn
}

// Markup sets the reply markup keyboard for the video note message.
func (svn *SendVideoNote) Markup(kb keyboard.KeyboardBuilder) *SendVideoNote {
	svn.opts.ReplyMarkup = kb.Markup()
	return svn
}

// Duration sets the video note duration in seconds.
func (svn *SendVideoNote) Duration(duration int64) *SendVideoNote {
	svn.opts.Duration = duration
	return svn
}

// Length sets the video note diameter (video notes are square).
func (svn *SendVideoNote) Length(length int64) *SendVideoNote {
	svn.opts.Length = length
	return svn
}

// Thumbnail sets a custom thumbnail for the video note.
func (svn *SendVideoNote) Thumbnail(file String) *SendVideoNote {
	svn.thumb = NewFile(file)

	reader := svn.thumb.Open()
	if reader.IsErr() {
		svn.err = reader.Err()
		return svn
	}

	svn.opts.Thumbnail = gotgbot.InputFileByReader(svn.thumb.Name().Std(), reader.Ok().Std())
	return svn
}

// ReplyTo sets the message ID to reply to.
func (svn *SendVideoNote) ReplyTo(messageID int64) *SendVideoNote {
	svn.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return svn
}

// Timeout sets a custom timeout for this request.
func (svn *SendVideoNote) Timeout(duration time.Duration) *SendVideoNote {
	if svn.opts.RequestOpts == nil {
		svn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	svn.opts.RequestOpts.Timeout = duration

	return svn
}

// APIURL sets a custom API URL for this request.
func (svn *SendVideoNote) APIURL(url String) *SendVideoNote {
	if svn.opts.RequestOpts == nil {
		svn.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	svn.opts.RequestOpts.APIURL = url.Std()

	return svn
}

// Business sets the business connection ID for the video note message.
func (svn *SendVideoNote) Business(id String) *SendVideoNote {
	svn.opts.BusinessConnectionId = id.Std()
	return svn
}

// Thread sets the message thread ID for the video note message.
func (svn *SendVideoNote) Thread(id int64) *SendVideoNote {
	svn.opts.MessageThreadId = id
	return svn
}

// To sets the target chat ID for the video note message.
func (svn *SendVideoNote) To(chatID int64) *SendVideoNote {
	svn.chatID = Some(chatID)
	return svn
}

// Send sends the video note message to Telegram and returns the result.
func (svn *SendVideoNote) Send() Result[*gotgbot.Message] {
	if svn.err != nil {
		return Err[*gotgbot.Message](svn.err)
	}

	if svn.file != nil {
		defer svn.file.Close()
	}

	if svn.thumb != nil {
		defer svn.thumb.Close()
	}

	return svn.ctx.timers(svn.after, svn.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := svn.chatID.UnwrapOr(svn.ctx.EffectiveChat.Id)
		return ResultOf(svn.ctx.Bot.Raw().SendVideoNote(chatID, svn.doc, svn.opts))
	})
}
