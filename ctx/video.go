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
	files       Slice[*File]
	thumb       *File
	removethumb bool
	duration    Float
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// CaptionEntities sets custom entities for the video caption.
func (sv *SendVideo) CaptionEntities(e *entities.Entities) *SendVideo {
	sv.opts.CaptionEntities = e.Std()
	return sv
}

// After schedules the video to be sent after the specified duration.
func (sv *SendVideo) After(duration time.Duration) *SendVideo {
	sv.after = Some(duration)
	return sv
}

// DeleteAfter schedules the video message to be deleted after the specified duration.
func (sv *SendVideo) DeleteAfter(duration time.Duration) *SendVideo {
	sv.deleteAfter = Some(duration)
	return sv
}

// Width sets the video width.
func (sv *SendVideo) Width(width int64) *SendVideo {
	sv.opts.Width = width
	return sv
}

// Height sets the video height.
func (sv *SendVideo) Height(height int64) *SendVideo {
	sv.opts.Height = height
	return sv
}

// Duration sets the video duration in seconds.
func (sv *SendVideo) Duration(duration int64) *SendVideo {
	sv.opts.Duration = duration
	return sv
}

// Thumbnail sets a custom thumbnail for the video.
func (sv *SendVideo) Thumbnail(file String) *SendVideo {
	thumb := NewFile(file)

	reader := thumb.Open()
	if reader.IsErr() {
		sv.err = reader.Err()
		return sv
	}

	sv.files.Push(thumb)
	sv.opts.Thumbnail = gotgbot.InputFileByReader(thumb.Name().Std(), reader.Ok().Std())

	return sv
}

// Spoiler marks the video as a spoiler.
func (sv *SendVideo) Spoiler() *SendVideo {
	sv.opts.HasSpoiler = true
	return sv
}

// Streamable enables streaming support for the video.
func (sv *SendVideo) Streamable() *SendVideo {
	sv.opts.SupportsStreaming = true
	return sv
}

// Caption sets the caption text for the video.
func (sv *SendVideo) Caption(caption String) *SendVideo {
	sv.opts.Caption = caption.Std()
	return sv
}

// HTML sets the caption parse mode to HTML.
func (sv *SendVideo) HTML() *SendVideo {
	sv.opts.ParseMode = "HTML"
	return sv
}

// Markdown sets the caption parse mode to MarkdownV2.
func (sv *SendVideo) Markdown() *SendVideo {
	sv.opts.ParseMode = "MarkdownV2"
	return sv
}

// Silent disables notification for the video message.
func (sv *SendVideo) Silent() *SendVideo {
	sv.opts.DisableNotification = true
	return sv
}

// Protect enables content protection for the video message.
func (sv *SendVideo) Protect() *SendVideo {
	sv.opts.ProtectContent = true
	return sv
}

// Markup sets the reply markup keyboard for the video message.
func (sv *SendVideo) Markup(kb keyboard.KeyboardBuilder) *SendVideo {
	sv.opts.ReplyMarkup = kb.Markup()
	return sv
}

// ReplyTo sets the message ID to reply to.
func (sv *SendVideo) ReplyTo(messageID int64) *SendVideo {
	sv.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return sv
}

// Timeout sets a custom timeout for this request.
func (sv *SendVideo) Timeout(duration time.Duration) *SendVideo {
	if sv.opts.RequestOpts == nil {
		sv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sv.opts.RequestOpts.Timeout = duration

	return sv
}

// APIURL sets a custom API URL for this request.
func (sv *SendVideo) APIURL(url String) *SendVideo {
	if sv.opts.RequestOpts == nil {
		sv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sv.opts.RequestOpts.APIURL = url.Std()

	return sv
}

// Business sets the business connection ID for the video message.
func (sv *SendVideo) Business(id String) *SendVideo {
	sv.opts.BusinessConnectionId = id.Std()
	return sv
}

// Thread sets the message thread ID for the video message.
func (sv *SendVideo) Thread(id int64) *SendVideo {
	sv.opts.MessageThreadId = id
	return sv
}

// ShowCaptionAboveMedia displays the caption above the video instead of below.
func (sv *SendVideo) ShowCaptionAboveMedia() *SendVideo {
	sv.opts.ShowCaptionAboveMedia = true
	return sv
}

// Cover sets a cover image for the video.
func (sv *SendVideo) Cover(filename String) *SendVideo {
	cover := NewFile(filename)

	reader := cover.Open()
	if reader.IsErr() {
		sv.err = reader.Err()
		return sv
	}

	sv.files.Push(cover)
	sv.opts.Cover = gotgbot.InputFileByReader(cover.Name().Std(), reader.Ok().Std())

	return sv
}

// StartTimestamp sets the video start timestamp in seconds.
func (sv *SendVideo) StartTimestamp(seconds int64) *SendVideo {
	sv.opts.StartTimestamp = seconds
	return sv
}

// ApplyMetadata automatically extracts and applies video metadata (duration, width, height).
func (sv *SendVideo) ApplyMetadata() *SendVideo {
	if sv.file == nil {
		sv.err = errors.New("video file is not set")
		return sv
	}

	path := sv.file.Path()
	if path.IsErr() {
		sv.err = path.Err()
		return sv
	}

	meta := ffmpeg.GetVideoMetadata(path.Ok())
	if meta.IsErr() {
		sv.err = meta.Err()
		return sv
	}

	info := meta.Ok()

	sv.duration = info.Duration.ToFloat().UnwrapOrDefault()
	sv.Width(info.Width)
	sv.Height(info.Height)

	if !sv.duration.IsZero() {
		sv.Duration(int64(sv.duration))
	}

	return sv
}

// GenerateThumbnail automatically generates a thumbnail from the video at the specified seek time.
func (sv *SendVideo) GenerateThumbnail(seek ...String) *SendVideo {
	if sv.file == nil {
		sv.err = errors.New("video file is not set")
		return sv
	}

	path := sv.file.Path()
	if path.IsErr() {
		sv.err = path.Err()
		return sv
	}

	if sv.duration.IsZero() {
		sv.err = errors.New("duration not set, call ApplyMetadata() first")
		return sv
	}

	var seekTime String

	if len(seek) != 0 {
		seekTime = seek[0]
	} else {
		seekTime = sv.duration.Div(2).Max(1.0).RoundDecimal(3).String()
	}

	thumb := ffmpeg.GenerateThumbnail(path.Ok(), seekTime)
	if thumb.IsErr() {
		sv.err = thumb.Err()
		return sv
	}

	sv.thumb = thumb.Ok()

	reader := sv.thumb.Open()
	if reader.IsErr() {
		sv.err = reader.Err()
		return sv
	}

	sv.opts.Thumbnail = gotgbot.InputFileByReader(sv.thumb.Name().Std(), reader.Ok().Std())
	return sv
}

// To sets the target chat ID for the video message.
func (sv *SendVideo) To(chatID int64) *SendVideo {
	sv.chatID = Some(chatID)
	return sv
}

// Send sends the video message to Telegram and returns the result.
func (sv *SendVideo) Send() Result[*gotgbot.Message] {
	if sv.err != nil {
		return Err[*gotgbot.Message](sv.err)
	}

	if sv.file != nil {
		defer sv.file.Close()
	}

	if sv.thumb != nil {
		defer func() {
			sv.thumb.Close()
			sv.thumb.Remove()
		}()
	}

	defer func() {
		sv.files.Iter().ForEach(func(file *File) {
			file.Close()
		})
	}()

	return sv.ctx.timers(sv.after, sv.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := sv.chatID.UnwrapOr(sv.ctx.EffectiveChat.Id)
		return ResultOf(sv.ctx.Bot.Raw().SendVideo(chatID, sv.doc, sv.opts))
	})
}
