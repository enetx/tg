package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
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

// After schedules the audio to be sent after the specified duration.
func (a *Audio) After(duration time.Duration) *Audio {
	a.after = Some(duration)
	return a
}

// DeleteAfter schedules the audio message to be deleted after the specified duration.
func (a *Audio) DeleteAfter(duration time.Duration) *Audio {
	a.deleteAfter = Some(duration)
	return a
}

// Caption sets the caption text for the audio.
func (a *Audio) Caption(caption String) *Audio {
	a.opts.Caption = caption.Std()
	return a
}

// HTML sets the caption parse mode to HTML.
func (a *Audio) HTML() *Audio {
	a.opts.ParseMode = "HTML"
	return a
}

// Markdown sets the caption parse mode to MarkdownV2.
func (a *Audio) Markdown() *Audio {
	a.opts.ParseMode = "MarkdownV2"
	return a
}

// Silent disables notification for the audio message.
func (a *Audio) Silent() *Audio {
	a.opts.DisableNotification = true
	return a
}

// Protect enables content protection for the audio message.
func (a *Audio) Protect() *Audio {
	a.opts.ProtectContent = true
	return a
}

// Markup sets the reply markup keyboard for the audio message.
func (a *Audio) Markup(kb keyboard.KeyboardBuilder) *Audio {
	a.opts.ReplyMarkup = kb.Markup()
	return a
}

// Thumbnail sets a custom thumbnail for the audio.
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

// ReplyTo sets the message ID to reply to.
func (a *Audio) ReplyTo(messageID int64) *Audio {
	a.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return a
}

// Timeout sets the request timeout duration.
func (a *Audio) Timeout(duration time.Duration) *Audio {
	a.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return a
}

// Business sets the business connection ID for the audio message.
func (a *Audio) Business(id String) *Audio {
	a.opts.BusinessConnectionId = id.Std()
	return a
}

// Thread sets the message thread ID for the audio message.
func (a *Audio) Thread(id int64) *Audio {
	a.opts.MessageThreadId = id
	return a
}

// Duration sets the audio duration in seconds.
func (a *Audio) Duration(seconds int64) *Audio {
	a.opts.Duration = seconds
	return a
}

// Performer sets the audio performer/artist name.
func (a *Audio) Performer(artist String) *Audio {
	a.opts.Performer = artist.Std()
	return a
}

// Title sets the audio track title.
func (a *Audio) Title(title String) *Audio {
	a.opts.Title = title.Std()
	return a
}

// To sets the target chat ID for the audio message.
func (a *Audio) To(chatID int64) *Audio {
	a.chatID = Some(chatID)
	return a
}

// Send sends the audio message to Telegram and returns the result.
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
