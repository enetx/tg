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
func (c *SendAudio) CaptionEntities(e *entities.Entities) *SendAudio {
	c.opts.CaptionEntities = e.Std()
	return c
}

// After schedules the audio to be sent after the specified duration.
func (c *SendAudio) After(duration time.Duration) *SendAudio {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the audio message to be deleted after the specified duration.
func (c *SendAudio) DeleteAfter(duration time.Duration) *SendAudio {
	c.deleteAfter = Some(duration)
	return c
}

// Caption sets the caption text for the audio.
func (c *SendAudio) Caption(caption String) *SendAudio {
	c.opts.Caption = caption.Std()
	return c
}

// HTML sets the caption parse mode to HTML.
func (c *SendAudio) HTML() *SendAudio {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the caption parse mode to MarkdownV2.
func (c *SendAudio) Markdown() *SendAudio {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the audio message.
func (c *SendAudio) Silent() *SendAudio {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the audio message.
func (c *SendAudio) Protect() *SendAudio {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the audio message.
func (c *SendAudio) Markup(kb keyboard.KeyboardBuilder) *SendAudio {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// Thumbnail sets a custom thumbnail for the audio.
func (c *SendAudio) Thumbnail(file String) *SendAudio {
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
func (c *SendAudio) ReplyTo(messageID int64) *SendAudio {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendAudio) Timeout(duration time.Duration) *SendAudio {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendAudio) APIURL(url String) *SendAudio {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the audio message.
func (c *SendAudio) Business(id String) *SendAudio {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the audio message.
func (c *SendAudio) Thread(id int64) *SendAudio {
	c.opts.MessageThreadId = id
	return c
}

// Duration sets the audio duration in seconds.
func (c *SendAudio) Duration(seconds int64) *SendAudio {
	c.opts.Duration = seconds
	return c
}

// Performer sets the audio performer/artist name.
func (c *SendAudio) Performer(artist String) *SendAudio {
	c.opts.Performer = artist.Std()
	return c
}

// Title sets the audio track title.
func (c *SendAudio) Title(title String) *SendAudio {
	c.opts.Title = title.Std()
	return c
}

// To sets the target chat ID for the audio message.
func (c *SendAudio) To(chatID int64) *SendAudio {
	c.chatID = Some(chatID)
	return c
}

// Send sends the audio message to Telegram and returns the result.
func (c *SendAudio) Send() Result[*gotgbot.Message] {
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
		return ResultOf(c.ctx.Bot.Raw().SendAudio(chatID, c.doc, c.opts))
	})
}
