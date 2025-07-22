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
func (c *SendAnimation) CaptionEntities(e *entities.Entities) *SendAnimation {
	c.opts.CaptionEntities = e.Std()
	return c
}

// After schedules the animation to be sent after the specified duration.
func (c *SendAnimation) After(duration time.Duration) *SendAnimation {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the animation message to be deleted after the specified duration.
func (c *SendAnimation) DeleteAfter(duration time.Duration) *SendAnimation {
	c.deleteAfter = Some(duration)
	return c
}

// Caption sets the caption text for the animation.
func (c *SendAnimation) Caption(caption String) *SendAnimation {
	c.opts.Caption = caption.Std()
	return c
}

// HTML sets the caption parse mode to HTML.
func (c *SendAnimation) HTML() *SendAnimation {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the caption parse mode to MarkdownV2.
func (c *SendAnimation) Markdown() *SendAnimation {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the animation message.
func (c *SendAnimation) Silent() *SendAnimation {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the animation message.
func (c *SendAnimation) Protect() *SendAnimation {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the animation message.
func (c *SendAnimation) Markup(kb keyboard.KeyboardBuilder) *SendAnimation {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// Duration sets the animation duration in seconds.
func (c *SendAnimation) Duration(duration int64) *SendAnimation {
	c.opts.Duration = duration
	return c
}

// Width sets the animation width.
func (c *SendAnimation) Width(width int64) *SendAnimation {
	c.opts.Width = width
	return c
}

// Height sets the animation height.
func (c *SendAnimation) Height(height int64) *SendAnimation {
	c.opts.Height = height
	return c
}

// Thumbnail sets a custom thumbnail for the animation.
func (c *SendAnimation) Thumbnail(file String) *SendAnimation {
	c.thumb = NewFile(file)

	reader := c.thumb.Open()
	if reader.IsErr() {
		c.err = reader.Err()
		return c
	}

	c.opts.Thumbnail = gotgbot.InputFileByReader(c.thumb.Name().Std(), reader.Ok().Std())

	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendAnimation) ReplyTo(messageID int64) *SendAnimation {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendAnimation) Timeout(duration time.Duration) *SendAnimation {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendAnimation) APIURL(url String) *SendAnimation {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the animation message.
func (c *SendAnimation) Business(id String) *SendAnimation {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the animation message.
func (c *SendAnimation) Thread(id int64) *SendAnimation {
	c.opts.MessageThreadId = id
	return c
}

// ShowCaptionAboveMedia displays the caption above the animation instead of below.
func (c *SendAnimation) ShowCaptionAboveMedia() *SendAnimation {
	c.opts.ShowCaptionAboveMedia = true
	return c
}

// Spoiler marks the animation as a spoiler.
func (c *SendAnimation) Spoiler() *SendAnimation {
	c.opts.HasSpoiler = true
	return c
}

// To sets the target chat ID for the animation message.
func (c *SendAnimation) To(chatID int64) *SendAnimation {
	c.chatID = Some(chatID)
	return c
}

// Send sends the animation message to Telegram and returns the result.
func (c *SendAnimation) Send() Result[*gotgbot.Message] {
	if c.err != nil {
		return Err[*gotgbot.Message](c.err)
	}

	if c.file != nil {
		defer c.file.Close()
	}

	if c.thumb != nil {
		defer c.thumb.Close()
	}

	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendAnimation(chatID, c.doc, c.opts))
	})
}
