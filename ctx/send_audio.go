package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

type SendAudio struct {
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

// CaptionEntities sets custom entities for the audio caption.
func (sa *SendAudio) CaptionEntities(e *entities.Entities) *SendAudio {
	sa.opts.CaptionEntities = e.Std()
	return sa
}

// After schedules the audio to be sent after the specified duration.
func (sa *SendAudio) After(duration time.Duration) *SendAudio {
	sa.after = Some(duration)
	return sa
}

// DeleteAfter schedules the audio message to be deleted after the specified duration.
func (sa *SendAudio) DeleteAfter(duration time.Duration) *SendAudio {
	sa.deleteAfter = Some(duration)
	return sa
}

// Caption sets the caption text for the audio.
func (sa *SendAudio) Caption(caption String) *SendAudio {
	sa.opts.Caption = caption.Std()
	return sa
}

// HTML sets the caption parse mode to HTML.
func (sa *SendAudio) HTML() *SendAudio {
	sa.opts.ParseMode = "HTML"
	return sa
}

// Markdown sets the caption parse mode to MarkdownV2.
func (sa *SendAudio) Markdown() *SendAudio {
	sa.opts.ParseMode = "MarkdownV2"
	return sa
}

// Silent disables notification for the audio message.
func (sa *SendAudio) Silent() *SendAudio {
	sa.opts.DisableNotification = true
	return sa
}

// Protect enables content protection for the audio message.
func (sa *SendAudio) Protect() *SendAudio {
	sa.opts.ProtectContent = true
	return sa
}

// Markup sets the reply markup keyboard for the audio message.
func (sa *SendAudio) Markup(kb keyboard.KeyboardBuilder) *SendAudio {
	sa.opts.ReplyMarkup = kb.Markup()
	return sa
}

// Thumbnail sets a custom thumbnail for the audio.
func (sa *SendAudio) Thumbnail(file String) *SendAudio {
	sa.thumb = NewFile(file)

	reader := sa.thumb.Open()
	if reader.IsErr() {
		sa.err = reader.Err()
		return sa
	}

	sa.opts.Thumbnail = gotgbot.InputFileByReader(sa.thumb.Name().Std(), reader.Ok().Std())
	return sa
}

// ReplyTo sets the message ID to reply to.
func (sa *SendAudio) ReplyTo(messageID int64) *SendAudio {
	sa.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return sa
}

// Timeout sets a custom timeout for this request.
func (sa *SendAudio) Timeout(duration time.Duration) *SendAudio {
	if sa.opts.RequestOpts == nil {
		sa.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sa.opts.RequestOpts.Timeout = duration

	return sa
}

// APIURL sets a custom API URL for this request.
func (sa *SendAudio) APIURL(url String) *SendAudio {
	if sa.opts.RequestOpts == nil {
		sa.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sa.opts.RequestOpts.APIURL = url.Std()

	return sa
}

// Business sets the business connection ID for the audio message.
func (sa *SendAudio) Business(id String) *SendAudio {
	sa.opts.BusinessConnectionId = id.Std()
	return sa
}

// Thread sets the message thread ID for the audio message.
func (sa *SendAudio) Thread(id int64) *SendAudio {
	sa.opts.MessageThreadId = id
	return sa
}

// Duration sets the audio duration in seconds.
func (sa *SendAudio) Duration(duration time.Duration) *SendAudio {
	sa.opts.Duration = int64(duration.Seconds())
	return sa
}

// Performer sets the audio performer/artist name.
func (sa *SendAudio) Performer(artist String) *SendAudio {
	sa.opts.Performer = artist.Std()
	return sa
}

// Title sets the audio track title.
func (sa *SendAudio) Title(title String) *SendAudio {
	sa.opts.Title = title.Std()
	return sa
}

// To sets the target chat ID for the audio message.
func (sa *SendAudio) To(chatID int64) *SendAudio {
	sa.chatID = Some(chatID)
	return sa
}

// Send sends the audio message to Telegram and returns the result.
func (sa *SendAudio) Send() Result[*gotgbot.Message] {
	if sa.err != nil {
		return Err[*gotgbot.Message](sa.err)
	}

	if sa.file != nil {
		defer sa.file.Close()
	}

	if sa.thumb != nil {
		defer sa.thumb.Close()
	}

	return sa.ctx.timers(sa.after, sa.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := sa.chatID.UnwrapOr(sa.ctx.EffectiveChat.Id)
		return ResultOf(sa.ctx.Bot.Raw().SendAudio(chatID, sa.doc, sa.opts))
	})
}
