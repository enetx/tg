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
func (c *SendVideoNote) After(duration time.Duration) *SendVideoNote {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the video note message to be deleted after the specified duration.
func (c *SendVideoNote) DeleteAfter(duration time.Duration) *SendVideoNote {
	c.deleteAfter = Some(duration)
	return c
}

// Silent disables notification for the video note message.
func (c *SendVideoNote) Silent() *SendVideoNote {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the video note message.
func (c *SendVideoNote) Protect() *SendVideoNote {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the video note message.
func (c *SendVideoNote) Markup(kb keyboard.KeyboardBuilder) *SendVideoNote {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// Duration sets the video note duration in seconds.
func (c *SendVideoNote) Duration(duration int64) *SendVideoNote {
	c.opts.Duration = duration
	return c
}

// Length sets the video note diameter (video notes are square).
func (c *SendVideoNote) Length(length int64) *SendVideoNote {
	c.opts.Length = length
	return c
}

// Thumbnail sets a custom thumbnail for the video note.
func (c *SendVideoNote) Thumbnail(file String) *SendVideoNote {
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
func (c *SendVideoNote) ReplyTo(messageID int64) *SendVideoNote {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendVideoNote) Timeout(duration time.Duration) *SendVideoNote {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendVideoNote) APIURL(url String) *SendVideoNote {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the video note message.
func (c *SendVideoNote) Business(id String) *SendVideoNote {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the video note message.
func (c *SendVideoNote) Thread(id int64) *SendVideoNote {
	c.opts.MessageThreadId = id
	return c
}

// To sets the target chat ID for the video note message.
func (c *SendVideoNote) To(chatID int64) *SendVideoNote {
	c.chatID = Some(chatID)
	return c
}

// Send sends the video note message to Telegram and returns the result.
func (c *SendVideoNote) Send() Result[*gotgbot.Message] {
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
		return ResultOf(c.ctx.Bot.Raw().SendVideoNote(chatID, c.doc, c.opts))
	})
}
