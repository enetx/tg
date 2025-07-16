package ctx

import (
	"errors"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/internal/pkg/ffmpeg"
	"github.com/enetx/tg/keyboard"
)

type Video struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendVideoOpts
	file        *File
	thumb       *File
	removethumb bool
	duration    Float
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

func (v *Video) After(duration time.Duration) *Video {
	v.after = Some(duration)
	return v
}

func (v *Video) DeleteAfter(duration time.Duration) *Video {
	v.deleteAfter = Some(duration)
	return v
}

func (v *Video) Width(width int64) *Video {
	v.opts.Width = width
	return v
}

func (v *Video) Height(height int64) *Video {
	v.opts.Height = height
	return v
}

func (v *Video) Duration(duration int64) *Video {
	v.opts.Duration = duration
	return v
}

func (v *Video) Thumbnail(file String) *Video {
	v.thumb = NewFile(file)

	reader := v.thumb.Open()
	if reader.IsErr() {
		v.err = reader.Err()
		return v
	}

	v.opts.Thumbnail = gotgbot.InputFileByReader(v.thumb.Name().Std(), reader.Ok().Std())
	return v
}

func (v *Video) Spoiler() *Video {
	v.opts.HasSpoiler = true
	return v
}

func (v *Video) Streamable() *Video {
	v.opts.SupportsStreaming = true
	return v
}

func (v *Video) Caption(caption String) *Video {
	v.opts.Caption = caption.Std()
	return v
}

func (v *Video) HTML() *Video {
	v.opts.ParseMode = "HTML"
	return v
}

func (v *Video) Markdown() *Video {
	v.opts.ParseMode = "Markdown"
	return v
}

func (v *Video) Silent() *Video {
	v.opts.DisableNotification = true
	return v
}

func (v *Video) Protect() *Video {
	v.opts.ProtectContent = true
	return v
}

func (v *Video) Markup(kb keyboard.KeyboardBuilder) *Video {
	v.opts.ReplyMarkup = kb.Markup()
	return v
}

func (v *Video) ReplyTo(messageID int64) *Video {
	v.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return v
}

func (v *Video) Timeout(duration time.Duration) *Video {
	v.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return v
}

func (v *Video) ApplyMetadata() *Video {
	if v.file == nil {
		v.err = errors.New("video file is not set")
		return v
	}

	path := v.file.Path()
	if path.IsErr() {
		v.err = path.Err()
		return v
	}

	meta := ffmpeg.GetVideoMetadata(path.Ok())
	if meta.IsErr() {
		v.err = meta.Err()
		return v
	}

	info := meta.Ok()

	v.duration = info.Duration.ToFloat().UnwrapOrDefault()
	v.Width(info.Width)
	v.Height(info.Height)

	if !v.duration.IsZero() {
		v.Duration(int64(v.duration))
	}

	return v
}

func (v *Video) GenerateThumbnail(seek ...String) *Video {
	if v.file == nil {
		v.err = errors.New("video file is not set")
		return v
	}

	path := v.file.Path()
	if path.IsErr() {
		v.err = path.Err()
		return v
	}

	if v.duration.IsZero() {
		v.err = errors.New("duration not set, call ApplyMetadata() first")
		return v
	}

	var seekTime String

	if len(seek) != 0 {
		seekTime = seek[0]
	} else {
		seekTime = v.duration.Div(2).Max(1.0).RoundDecimal(3).String()
	}

	thumb := ffmpeg.GenerateThumbnail(path.Ok(), seekTime)
	if thumb.IsErr() {
		v.err = thumb.Err()
		return v
	}

	v.thumb = thumb.Ok()

	reader := v.thumb.Open()
	if reader.IsErr() {
		v.err = reader.Err()
		return v
	}

	v.removethumb = true

	v.opts.Thumbnail = gotgbot.InputFileByReader(v.thumb.Name().Std(), reader.Ok().Std())
	return v
}

func (v *Video) To(chatID int64) *Video {
	v.chatID = Some(chatID)
	return v
}

func (v *Video) Send() Result[*gotgbot.Message] {
	if v.err != nil {
		return Err[*gotgbot.Message](v.err)
	}

	if v.file != nil {
		defer v.file.Close()
	}

	if v.thumb != nil {
		defer func() {
			v.thumb.Close()
			if v.removethumb {
				v.thumb.Remove()
			}
		}()
	}

	return v.ctx.timers(v.after, v.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := v.chatID.UnwrapOr(v.ctx.EffectiveChat.Id)
		return ResultOf(v.ctx.Bot.Raw().SendVideo(chatID, v.doc, v.opts))
	})
}
