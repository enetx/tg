package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/internal/pkg/ffmpeg"
	"github.com/enetx/tg/internal/tgfile"
	"github.com/enetx/tg/keyboard"
)

// SendPaidMedia represents a request to send paid media content.
type SendPaidMedia struct {
	ctx       *Context
	opts      *gotgbot.SendPaidMediaOpts
	chatID    Option[int64]
	starCount int64
	media     Slice[gotgbot.InputPaidMedia]
	files     Slice[*File]
	tempfiles Slice[*File]
	err       error
}

// To sets the target chat ID for sending paid media.
func (spm *SendPaidMedia) To(chatID int64) *SendPaidMedia {
	spm.chatID = Some(chatID)
	return spm
}

// Photo adds a paid photo to the media list.
func (spm *SendPaidMedia) Photo(filename String) *SendPaidMedia {
	if spm.err != nil {
		return spm
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		spm.err = result.Err()
		return spm
	}

	spm.files.Push(result.Ok().File)
	spm.media.Push(gotgbot.InputPaidMediaPhoto{Media: result.Ok().Doc})

	return spm
}

// PaidVideoBuilder represents a builder for paid video configuration.
type PaidVideoBuilder struct {
	parent   *SendPaidMedia
	video    gotgbot.InputPaidMediaVideo
	file     *File
	thumb    *File
	duration Float
	err      error
}

// Width sets the video width.
func (pvb *PaidVideoBuilder) Width(width int64) *PaidVideoBuilder {
	pvb.video.Width = width
	return pvb
}

// Height sets the video height.
func (pvb *PaidVideoBuilder) Height(height int64) *PaidVideoBuilder {
	pvb.video.Height = height
	return pvb
}

// Duration sets the video duration in seconds.
func (pvb *PaidVideoBuilder) Duration(duration int64) *PaidVideoBuilder {
	pvb.video.Duration = duration
	return pvb
}

// Streamable enables streaming support for the video.
func (pvb *PaidVideoBuilder) Streamable() *PaidVideoBuilder {
	pvb.video.SupportsStreaming = true
	return pvb
}

// StartTimestamp sets the start timestamp for the video.
func (pvb *PaidVideoBuilder) StartTimestamp(seconds int64) *PaidVideoBuilder {
	pvb.video.StartTimestamp = seconds
	return pvb
}

// Cover sets a cover image for the video.
func (pvb *PaidVideoBuilder) Cover(filename String) *PaidVideoBuilder {
	pvb.video.Cover = filename.Std()
	return pvb
}

// Thumbnail sets a custom thumbnail for the video.
func (pvb *PaidVideoBuilder) Thumbnail(filename String) *PaidVideoBuilder {
	pvb.thumb = NewFile(filename)

	reader := pvb.thumb.Open()
	if reader.IsErr() {
		pvb.err = reader.Err()
		return pvb
	}

	pvb.video.Thumbnail = gotgbot.InputFileByReader(pvb.thumb.Name().Std(), reader.Ok().Std())
	pvb.parent.files.Push(pvb.thumb)

	return pvb
}

// ApplyMetadata automatically extracts and applies video metadata (duration, width, height).
func (pvb *PaidVideoBuilder) ApplyMetadata() *PaidVideoBuilder {
	if pvb.file == nil {
		pvb.err = Errorf("video file is not set")
		return pvb
	}

	path := pvb.file.Path()
	if path.IsErr() {
		pvb.err = path.Err()
		return pvb
	}

	meta := ffmpeg.GetVideoMetadata(path.Ok())
	if meta.IsErr() {
		pvb.err = meta.Err()
		return pvb
	}

	info := meta.Ok()

	pvb.duration = info.Duration.ToFloat().UnwrapOrDefault()
	pvb.Width(info.Width)
	pvb.Height(info.Height)

	if !pvb.duration.IsZero() {
		pvb.Duration(int64(pvb.duration))
	}

	return pvb
}

// GenerateThumbnail automatically generates a thumbnail from the video at the specified seek time.
func (pvb *PaidVideoBuilder) GenerateThumbnail(seek ...String) *PaidVideoBuilder {
	if pvb.file == nil {
		pvb.err = Errorf("video file is not set")
		return pvb
	}

	path := pvb.file.Path()
	if path.IsErr() {
		pvb.err = path.Err()
		return pvb
	}

	if pvb.duration.IsZero() {
		pvb.err = Errorf("duration not set, call ApplyMetadata() first")
		return pvb
	}

	var seekTime String

	if len(seek) != 0 {
		seekTime = seek[0]
	} else {
		seekTime = pvb.duration.Div(2).Max(1.0).RoundDecimal(3).String()
	}

	thumb := ffmpeg.GenerateThumbnail(path.Ok(), seekTime)
	if thumb.IsErr() {
		pvb.err = thumb.Err()
		return pvb
	}

	pvb.thumb = thumb.Ok()

	reader := pvb.thumb.Open()
	if reader.IsErr() {
		pvb.err = reader.Err()
		return pvb
	}

	pvb.video.Thumbnail = gotgbot.InputFileByReader(pvb.thumb.Name().Std(), reader.Ok().Std())
	pvb.parent.tempfiles.Push(pvb.thumb)

	return pvb
}

// Add completes the video configuration and adds it to the media list.
func (pvb *PaidVideoBuilder) Add() *SendPaidMedia {
	if pvb.err != nil {
		pvb.parent.err = pvb.err
		return pvb.parent
	}

	pvb.parent.files.Push(pvb.file)
	pvb.parent.media.Push(pvb.video)

	return pvb.parent
}

// Video creates a new paid video builder.
func (spm *SendPaidMedia) Video(filename String) *PaidVideoBuilder {
	if spm.err != nil {
		return &PaidVideoBuilder{parent: spm, err: spm.err}
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		return &PaidVideoBuilder{parent: spm, err: result.Err()}
	}

	return &PaidVideoBuilder{
		parent: spm,
		video:  gotgbot.InputPaidMediaVideo{Media: result.Ok().Doc},
		file:   result.Ok().File,
	}
}

// Business sets the business connection ID for the paid media.
func (spm *SendPaidMedia) Business(businessConnectionID String) *SendPaidMedia {
	spm.opts.BusinessConnectionId = businessConnectionID.Std()
	return spm
}

// Payload sets the bot-defined payload for internal processing.
func (spm *SendPaidMedia) Payload(payload String) *SendPaidMedia {
	spm.opts.Payload = payload.Std()
	return spm
}

// Caption sets the media caption.
func (spm *SendPaidMedia) Caption(caption String) *SendPaidMedia {
	spm.opts.Caption = caption.Std()
	return spm
}

// HTML sets the caption parse mode to HTML.
func (spm *SendPaidMedia) HTML() *SendPaidMedia {
	spm.opts.ParseMode = "HTML"
	return spm
}

// Markdown sets the caption parse mode to Markdown.
func (spm *SendPaidMedia) Markdown() *SendPaidMedia {
	spm.opts.ParseMode = "MarkdownV2"
	return spm
}

// ShowCaptionAbove shows the caption above the media.
func (spm *SendPaidMedia) ShowCaptionAbove() *SendPaidMedia {
	spm.opts.ShowCaptionAboveMedia = true
	return spm
}

// Silent sends the message silently (no notification sound).
func (spm *SendPaidMedia) Silent() *SendPaidMedia {
	spm.opts.DisableNotification = true
	return spm
}

// Protect protects the media from forwarding and saving.
func (spm *SendPaidMedia) Protect() *SendPaidMedia {
	spm.opts.ProtectContent = true
	return spm
}

// AllowPaidBroadcast allows paid broadcast for high-speed delivery.
func (spm *SendPaidMedia) AllowPaidBroadcast() *SendPaidMedia {
	spm.opts.AllowPaidBroadcast = true
	return spm
}

// ReplyTo sets the message ID to reply to.
func (spm *SendPaidMedia) ReplyTo(messageID int64) *SendPaidMedia {
	spm.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return spm
}

// Markup sets the reply markup keyboard.
func (spm *SendPaidMedia) Markup(kb keyboard.KeyboardBuilder) *SendPaidMedia {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		spm.opts.ReplyMarkup = markup
	}

	return spm
}

// Timeout sets a custom timeout for this request.
func (spm *SendPaidMedia) Timeout(duration time.Duration) *SendPaidMedia {
	if spm.opts.RequestOpts == nil {
		spm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	spm.opts.RequestOpts.Timeout = duration

	return spm
}

// APIURL sets a custom API URL for this request.
func (spm *SendPaidMedia) APIURL(url String) *SendPaidMedia {
	if spm.opts.RequestOpts == nil {
		spm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	spm.opts.RequestOpts.APIURL = url.Std()

	return spm
}

// Send sends the paid media and returns the message.
func (spm *SendPaidMedia) Send() Result[*gotgbot.Message] {
	if spm.err != nil {
		return Err[*gotgbot.Message](spm.err)
	}

	if spm.media.Empty() {
		return Err[*gotgbot.Message](Errorf("no paid media specified"))
	}

	if spm.media.Len() > 10 {
		return Err[*gotgbot.Message](Errorf("too many media items: {} (maximum 10)", spm.media.Len()))
	}

	if spm.starCount < 1 || spm.starCount > 10000 {
		return Err[*gotgbot.Message](Errorf("star count must be between 1-10000, got {}", spm.starCount))
	}

	defer func() {
		spm.files.Iter().ForEach(func(file *File) {
			file.Close()
		})

		spm.tempfiles.Iter().ForEach(func(file *File) {
			file.Close()
			file.Remove()
		})
	}()

	chatID := spm.chatID.UnwrapOr(spm.ctx.EffectiveChat.Id)

	return ResultOf(spm.ctx.Bot.Raw().SendPaidMedia(chatID, spm.starCount, spm.media, spm.opts))
}
