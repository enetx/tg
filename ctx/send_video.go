package ctx

import (
	"errors"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/internal/pkg/ffmpeg"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/suggested"
	"github.com/enetx/tg/types/effects"
)

type SendVideo struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendVideoOpts
	file        *g.File
	files       g.Slice[*g.File]
	thumb       *g.File
	removethumb bool
	duration    g.Float
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	err         error
}

// CaptionEntities sets custom entities for the video caption.
func (sv *SendVideo) CaptionEntities(e *entities.Entities) *SendVideo {
	sv.opts.CaptionEntities = e.Std()
	return sv
}

// After schedules the video to be sent after the specified duration.
func (sv *SendVideo) After(duration time.Duration) *SendVideo {
	sv.after = g.Some(duration)
	return sv
}

// DeleteAfter schedules the video message to be deleted after the specified duration.
func (sv *SendVideo) DeleteAfter(duration time.Duration) *SendVideo {
	sv.deleteAfter = g.Some(duration)
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
func (sv *SendVideo) Duration(duration time.Duration) *SendVideo {
	sv.opts.Duration = int64(duration.Seconds())
	return sv
}

// Thumbnail sets a custom thumbnail for the video.
func (sv *SendVideo) Thumbnail(file g.String) *SendVideo {
	thumb := g.NewFile(file)

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
func (sv *SendVideo) Caption(caption g.String) *SendVideo {
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

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (sv *SendVideo) AllowPaidBroadcast() *SendVideo {
	sv.opts.AllowPaidBroadcast = true
	return sv
}

// Effect sets a message effect for the message.
func (sv *SendVideo) Effect(effect effects.EffectType) *SendVideo {
	sv.opts.MessageEffectId = effect.String()
	return sv
}

// Markup sets the reply markup keyboard for the video message.
func (sv *SendVideo) Markup(kb keyboard.Keyboard) *SendVideo {
	sv.opts.ReplyMarkup = kb.Markup()
	return sv
}

// Reply sets reply parameters using the reply builder.
func (sv *SendVideo) Reply(params *reply.Parameters) *SendVideo {
	sv.opts.ReplyParameters = params.Std()
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
func (sv *SendVideo) APIURL(url g.String) *SendVideo {
	if sv.opts.RequestOpts == nil {
		sv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sv.opts.RequestOpts.APIURL = url.Std()

	return sv
}

// Business sets the business connection ID for the video message.
func (sv *SendVideo) Business(id g.String) *SendVideo {
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

// SuggestedPost sets suggested post parameters for direct messages chats.
func (sv *SendVideo) SuggestedPost(params *suggested.PostParameters) *SendVideo {
	if params != nil {
		sv.opts.SuggestedPostParameters = params.Std()
	}
	return sv
}

// Cover sets a cover image for the video.
func (sv *SendVideo) Cover(filename g.String) *SendVideo {
	cover := g.NewFile(filename)

	reader := cover.Open()
	if reader.IsErr() {
		sv.err = reader.Err()
		return sv
	}

	sv.files.Push(cover)
	sv.opts.Cover = gotgbot.InputFileByReader(cover.Name().Std(), reader.Ok().Std())

	return sv
}

// StartAt sets the video start timestamp from the beginning.
func (sv *SendVideo) StartAt(offset time.Duration) *SendVideo {
	sv.opts.StartTimestamp = int64(offset.Seconds())
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
		sv.Duration(time.Duration(sv.duration) * time.Second)
	}

	return sv
}

// GenerateThumbnail automatically generates a thumbnail from the video at the specified seek time.
func (sv *SendVideo) GenerateThumbnail(seek ...g.String) *SendVideo {
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

	var seekTime g.String

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
	sv.chatID = g.Some(chatID)
	return sv
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (sv *SendVideo) DirectMessagesTopic(topicID int64) *SendVideo {
	sv.opts.DirectMessagesTopicId = topicID
	return sv
}

// Send sends the video message to Telegram and returns the result.
func (sv *SendVideo) Send() g.Result[*gotgbot.Message] {
	if sv.err != nil {
		return g.Err[*gotgbot.Message](sv.err)
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
		sv.files.Iter().ForEach(func(file *g.File) {
			file.Close()
		})
	}()

	return sv.ctx.timers(sv.after, sv.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sv.chatID.UnwrapOr(sv.ctx.EffectiveChat.Id)
		return g.ResultOf(sv.ctx.Bot.Raw().SendVideo(chatID, sv.doc, sv.opts))
	})
}
