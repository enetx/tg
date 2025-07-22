package ctx

import (
	"errors"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/internal/pkg/ffmpeg"
	"github.com/enetx/tg/keyboard"
)

type SendVideo struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendVideoOpts
	file        *File
	thumb       *File
	cover       *File
	removethumb bool
	duration    Float
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// CaptionEntities sets custom entities for the video caption.
func (c *SendVideo) CaptionEntities(e *entities.Entities) *SendVideo {
	c.opts.CaptionEntities = e.Std()
	return c
}

// After schedules the video to be sent after the specified duration.
func (c *SendVideo) After(duration time.Duration) *SendVideo {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the video message to be deleted after the specified duration.
func (c *SendVideo) DeleteAfter(duration time.Duration) *SendVideo {
	c.deleteAfter = Some(duration)
	return c
}

// Width sets the video width.
func (c *SendVideo) Width(width int64) *SendVideo {
	c.opts.Width = width
	return c
}

// Height sets the video height.
func (c *SendVideo) Height(height int64) *SendVideo {
	c.opts.Height = height
	return c
}

// Duration sets the video duration in seconds.
func (c *SendVideo) Duration(duration int64) *SendVideo {
	c.opts.Duration = duration
	return c
}

// Thumbnail sets a custom thumbnail for the video.
func (c *SendVideo) Thumbnail(file String) *SendVideo {
	c.thumb = NewFile(file)

	reader := c.thumb.Open()
	if reader.IsErr() {
		c.err = reader.Err()
		return c
	}

	c.opts.Thumbnail = gotgbot.InputFileByReader(c.thumb.Name().Std(), reader.Ok().Std())
	return c
}

// Spoiler marks the video as a spoiler.
func (c *SendVideo) Spoiler() *SendVideo {
	c.opts.HasSpoiler = true
	return c
}

// Streamable enables streaming support for the video.
func (c *SendVideo) Streamable() *SendVideo {
	c.opts.SupportsStreaming = true
	return c
}

// Caption sets the caption text for the video.
func (c *SendVideo) Caption(caption String) *SendVideo {
	c.opts.Caption = caption.Std()
	return c
}

// HTML sets the caption parse mode to HTML.
func (c *SendVideo) HTML() *SendVideo {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the caption parse mode to MarkdownV2.
func (c *SendVideo) Markdown() *SendVideo {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// Silent disables notification for the video message.
func (c *SendVideo) Silent() *SendVideo {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the video message.
func (c *SendVideo) Protect() *SendVideo {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the video message.
func (c *SendVideo) Markup(kb keyboard.KeyboardBuilder) *SendVideo {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendVideo) ReplyTo(messageID int64) *SendVideo {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendVideo) Timeout(duration time.Duration) *SendVideo {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendVideo) APIURL(url String) *SendVideo {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the video message.
func (c *SendVideo) Business(id String) *SendVideo {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the video message.
func (c *SendVideo) Thread(id int64) *SendVideo {
	c.opts.MessageThreadId = id
	return c
}

// ShowCaptionAboveMedia displays the caption above the video instead of below.
func (c *SendVideo) ShowCaptionAboveMedia() *SendVideo {
	c.opts.ShowCaptionAboveMedia = true
	return c
}

// Cover sets a cover image for the video.
func (c *SendVideo) Cover(filename String) *SendVideo {
	c.cover = NewFile(filename)

	reader := c.cover.Open()
	if reader.IsErr() {
		c.err = reader.Err()
		return c
	}

	c.opts.Cover = gotgbot.InputFileByReader(c.cover.Name().Std(), reader.Ok().Std())
	return c
}

// StartTimestamp sets the video start timestamp in seconds.
func (c *SendVideo) StartTimestamp(seconds int64) *SendVideo {
	c.opts.StartTimestamp = seconds
	return c
}

// ApplyMetadata automatically extracts and applies video metadata (duration, width, height).
func (c *SendVideo) ApplyMetadata() *SendVideo {
	if c.file == nil {
		c.err = errors.New("video file is not set")
		return c
	}

	path := c.file.Path()
	if path.IsErr() {
		c.err = path.Err()
		return c
	}

	meta := ffmpeg.GetVideoMetadata(path.Ok())
	if meta.IsErr() {
		c.err = meta.Err()
		return c
	}

	info := meta.Ok()

	c.duration = info.Duration.ToFloat().UnwrapOrDefault()
	c.Width(info.Width)
	c.Height(info.Height)

	if !c.duration.IsZero() {
		c.Duration(int64(c.duration))
	}

	return c
}

// GenerateThumbnail automatically generates a thumbnail from the video at the specified seek time.
func (c *SendVideo) GenerateThumbnail(seek ...String) *SendVideo {
	if c.file == nil {
		c.err = errors.New("video file is not set")
		return c
	}

	path := c.file.Path()
	if path.IsErr() {
		c.err = path.Err()
		return c
	}

	if c.duration.IsZero() {
		c.err = errors.New("duration not set, call ApplyMetadata() first")
		return c
	}

	var seekTime String

	if len(seek) != 0 {
		seekTime = seek[0]
	} else {
		seekTime = c.duration.Div(2).Max(1.0).RoundDecimal(3).String()
	}

	thumb := ffmpeg.GenerateThumbnail(path.Ok(), seekTime)
	if thumb.IsErr() {
		c.err = thumb.Err()
		return c
	}

	c.thumb = thumb.Ok()

	reader := c.thumb.Open()
	if reader.IsErr() {
		c.err = reader.Err()
		return c
	}

	c.removethumb = true

	c.opts.Thumbnail = gotgbot.InputFileByReader(c.thumb.Name().Std(), reader.Ok().Std())
	return c
}

// To sets the target chat ID for the video message.
func (c *SendVideo) To(chatID int64) *SendVideo {
	c.chatID = Some(chatID)
	return c
}

// Send sends the video message to Telegram and returns the result.
func (c *SendVideo) Send() Result[*gotgbot.Message] {
	if c.err != nil {
		return Err[*gotgbot.Message](c.err)
	}

	if c.file != nil {
		defer c.file.Close()
	}

	if c.thumb != nil {
		defer func() {
			c.thumb.Close()
			if c.removethumb {
				c.thumb.Remove()
			}
		}()
	}

	if c.cover != nil {
		defer c.cover.Close()
	}

	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendVideo(chatID, c.doc, c.opts))
	})
}
