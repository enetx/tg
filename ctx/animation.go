package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

type SendAnimation struct {
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

// CaptionEntities sets custom entities for the animation caption.
func (sa *SendAnimation) CaptionEntities(e *entities.Entities) *SendAnimation {
	sa.opts.CaptionEntities = e.Std()
	return sa
}

// After schedules the animation to be sent after the specified duration.
func (sa *SendAnimation) After(duration time.Duration) *SendAnimation {
	sa.after = Some(duration)
	return sa
}

// DeleteAfter schedules the animation message to be deleted after the specified duration.
func (sa *SendAnimation) DeleteAfter(duration time.Duration) *SendAnimation {
	sa.deleteAfter = Some(duration)
	return sa
}

// Caption sets the caption text for the animation.
func (sa *SendAnimation) Caption(caption String) *SendAnimation {
	sa.opts.Caption = caption.Std()
	return sa
}

// HTML sets the caption parse mode to HTML.
func (sa *SendAnimation) HTML() *SendAnimation {
	sa.opts.ParseMode = "HTML"
	return sa
}

// Markdown sets the caption parse mode to MarkdownV2.
func (sa *SendAnimation) Markdown() *SendAnimation {
	sa.opts.ParseMode = "MarkdownV2"
	return sa
}

// Silent disables notification for the animation message.
func (sa *SendAnimation) Silent() *SendAnimation {
	sa.opts.DisableNotification = true
	return sa
}

// Protect enables content protection for the animation message.
func (sa *SendAnimation) Protect() *SendAnimation {
	sa.opts.ProtectContent = true
	return sa
}

// Markup sets the reply markup keyboard for the animation message.
func (sa *SendAnimation) Markup(kb keyboard.KeyboardBuilder) *SendAnimation {
	sa.opts.ReplyMarkup = kb.Markup()
	return sa
}

// Duration sets the animation duration in seconds.
func (sa *SendAnimation) Duration(duration int64) *SendAnimation {
	sa.opts.Duration = duration
	return sa
}

// Width sets the animation width.
func (sa *SendAnimation) Width(width int64) *SendAnimation {
	sa.opts.Width = width
	return sa
}

// Height sets the animation height.
func (sa *SendAnimation) Height(height int64) *SendAnimation {
	sa.opts.Height = height
	return sa
}

// Thumbnail sets a custom thumbnail for the animation.
func (sa *SendAnimation) Thumbnail(file String) *SendAnimation {
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
func (sa *SendAnimation) ReplyTo(messageID int64) *SendAnimation {
	sa.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return sa
}

// Timeout sets a custom timeout for this request.
func (sa *SendAnimation) Timeout(duration time.Duration) *SendAnimation {
	if sa.opts.RequestOpts == nil {
		sa.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sa.opts.RequestOpts.Timeout = duration

	return sa
}

// APIURL sets a custom API URL for this request.
func (sa *SendAnimation) APIURL(url String) *SendAnimation {
	if sa.opts.RequestOpts == nil {
		sa.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sa.opts.RequestOpts.APIURL = url.Std()

	return sa
}

// Business sets the business connection ID for the animation message.
func (sa *SendAnimation) Business(id String) *SendAnimation {
	sa.opts.BusinessConnectionId = id.Std()
	return sa
}

// Thread sets the message thread ID for the animation message.
func (sa *SendAnimation) Thread(id int64) *SendAnimation {
	sa.opts.MessageThreadId = id
	return sa
}

// ShowCaptionAboveMedia displays the caption above the animation instead of below.
func (sa *SendAnimation) ShowCaptionAboveMedia() *SendAnimation {
	sa.opts.ShowCaptionAboveMedia = true
	return sa
}

// Spoiler marks the animation as a spoiler.
func (sa *SendAnimation) Spoiler() *SendAnimation {
	sa.opts.HasSpoiler = true
	return sa
}

// To sets the target chat ID for the animation message.
func (sa *SendAnimation) To(chatID int64) *SendAnimation {
	sa.chatID = Some(chatID)
	return sa
}

// Send sends the animation message to Telegram and returns the result.
func (sa *SendAnimation) Send() Result[*gotgbot.Message] {
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
		return ResultOf(sa.ctx.Bot.Raw().SendAnimation(chatID, sa.doc, sa.opts))
	})
}
