package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Animation struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendAnimationOpts
	file        *File
	thumb       *File
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// After schedules the animation to be sent after the specified duration.
func (a *Animation) After(duration time.Duration) *Animation {
	a.after = Some(duration)
	return a
}

// DeleteAfter schedules the animation message to be deleted after the specified duration.
func (a *Animation) DeleteAfter(duration time.Duration) *Animation {
	a.deleteAfter = Some(duration)
	return a
}

// Caption sets the caption text for the animation.
func (a *Animation) Caption(caption String) *Animation {
	a.opts.Caption = caption.Std()
	return a
}

// HTML sets the caption parse mode to HTML.
func (a *Animation) HTML() *Animation {
	a.opts.ParseMode = "HTML"
	return a
}

// Markdown sets the caption parse mode to MarkdownV2.
func (a *Animation) Markdown() *Animation {
	a.opts.ParseMode = "MarkdownV2"
	return a
}

// Silent disables notification for the animation message.
func (a *Animation) Silent() *Animation {
	a.opts.DisableNotification = true
	return a
}

// Protect enables content protection for the animation message.
func (a *Animation) Protect() *Animation {
	a.opts.ProtectContent = true
	return a
}

// Markup sets the reply markup keyboard for the animation message.
func (a *Animation) Markup(kb keyboard.KeyboardBuilder) *Animation {
	a.opts.ReplyMarkup = kb.Markup()
	return a
}

// Duration sets the animation duration in seconds.
func (a *Animation) Duration(duration int64) *Animation {
	a.opts.Duration = duration
	return a
}

// Width sets the animation width.
func (a *Animation) Width(width int64) *Animation {
	a.opts.Width = width
	return a
}

// Height sets the animation height.
func (a *Animation) Height(height int64) *Animation {
	a.opts.Height = height
	return a
}

// Thumbnail sets a custom thumbnail for the animation.
func (a *Animation) Thumbnail(file String) *Animation {
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
func (a *Animation) ReplyTo(messageID int64) *Animation {
	a.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return a
}

// Timeout sets the request timeout duration.
func (a *Animation) Timeout(duration time.Duration) *Animation {
	a.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return a
}

// Business sets the business connection ID for the animation message.
func (a *Animation) Business(id String) *Animation {
	a.opts.BusinessConnectionId = id.Std()
	return a
}

// Thread sets the message thread ID for the animation message.
func (a *Animation) Thread(id int64) *Animation {
	a.opts.MessageThreadId = id
	return a
}

// ShowCaptionAboveMedia displays the caption above the animation instead of below.
func (a *Animation) ShowCaptionAboveMedia() *Animation {
	a.opts.ShowCaptionAboveMedia = true
	return a
}

// Spoiler marks the animation as a spoiler.
func (a *Animation) Spoiler() *Animation {
	a.opts.HasSpoiler = true
	return a
}

// To sets the target chat ID for the animation message.
func (a *Animation) To(chatID int64) *Animation {
	a.chatID = Some(chatID)
	return a
}

// Send sends the animation message to Telegram and returns the result.
func (a *Animation) Send() Result[*gotgbot.Message] {
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
		return ResultOf(a.ctx.Bot.Raw().SendAnimation(chatID, a.doc, a.opts))
	})
}
