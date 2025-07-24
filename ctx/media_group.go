package ctx

import (
	"errors"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/internal/pkg/ffmpeg"
	"github.com/enetx/tg/internal/tgfile"
	"github.com/enetx/tg/types/effects"
)

type MediaGroup struct {
	ctx         *Context
	media       Slice[gotgbot.InputMedia]
	files       Slice[*File]
	tempfiles   Slice[*File]
	opts        *gotgbot.SendMediaGroupOpts
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// After schedules the media group to be sent after the specified duration.
func (mg *MediaGroup) After(duration time.Duration) *MediaGroup {
	mg.after = Some(duration)
	return mg
}

// DeleteAfter schedules the media group messages to be deleted after the specified duration.
func (mg *MediaGroup) DeleteAfter(duration time.Duration) *MediaGroup {
	mg.deleteAfter = Some(duration)
	return mg
}

// Photo adds a photo to the media group with optional caption.
func (mg *MediaGroup) Photo(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		mg.err = result.Err()
		return mg
	}

	mg.files.Push(result.Ok().File)

	media := gotgbot.InputMediaPhoto{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	mg.media.Push(media)

	return mg
}

// MediaVideoBuilder represents a builder for media group video configuration.
type MediaVideoBuilder struct {
	parent   *MediaGroup
	video    gotgbot.InputMediaVideo
	file     *File
	thumb    *File
	duration Float
	err      error
}

// Width sets the video width.
func (mvb *MediaVideoBuilder) Width(width int64) *MediaVideoBuilder {
	mvb.video.Width = width
	return mvb
}

// Height sets the video height.
func (mvb *MediaVideoBuilder) Height(height int64) *MediaVideoBuilder {
	mvb.video.Height = height
	return mvb
}

// Duration sets the video duration in seconds.
func (mvb *MediaVideoBuilder) Duration(duration int64) *MediaVideoBuilder {
	mvb.video.Duration = duration
	return mvb
}

// Streamable enables streaming support for the video.
func (mvb *MediaVideoBuilder) Streamable() *MediaVideoBuilder {
	mvb.video.SupportsStreaming = true
	return mvb
}

// Caption sets the video caption.
func (mvb *MediaVideoBuilder) Caption(caption String) *MediaVideoBuilder {
	mvb.video.Caption = caption.Std()
	return mvb
}

// HTML sets the caption parse mode to HTML.
func (mvb *MediaVideoBuilder) HTML() *MediaVideoBuilder {
	mvb.video.ParseMode = "HTML"
	return mvb
}

// Markdown sets the caption parse mode to Markdown.
func (mvb *MediaVideoBuilder) Markdown() *MediaVideoBuilder {
	mvb.video.ParseMode = "MarkdownV2"
	return mvb
}

// ShowCaptionAbove shows the caption above the video.
func (mvb *MediaVideoBuilder) ShowCaptionAbove() *MediaVideoBuilder {
	mvb.video.ShowCaptionAboveMedia = true
	return mvb
}

// HasSpoiler marks the video as a spoiler.
func (mvb *MediaVideoBuilder) HasSpoiler() *MediaVideoBuilder {
	mvb.video.HasSpoiler = true
	return mvb
}

// Thumbnail sets a custom thumbnail for the video.
func (mvb *MediaVideoBuilder) Thumbnail(filename String) *MediaVideoBuilder {
	mvb.thumb = NewFile(filename)

	reader := mvb.thumb.Open()
	if reader.IsErr() {
		mvb.err = reader.Err()
		return mvb
	}

	mvb.video.Thumbnail = gotgbot.InputFileByReader(mvb.thumb.Name().Std(), reader.Ok().Std())
	mvb.parent.files.Push(mvb.thumb)

	return mvb
}

// ApplyMetadata automatically extracts and applies video metadata (duration, width, height).
func (mvb *MediaVideoBuilder) ApplyMetadata() *MediaVideoBuilder {
	if mvb.file == nil {
		mvb.err = Errorf("video file is not set")
		return mvb
	}

	path := mvb.file.Path()
	if path.IsErr() {
		mvb.err = path.Err()
		return mvb
	}

	meta := ffmpeg.GetVideoMetadata(path.Ok())
	if meta.IsErr() {
		mvb.err = meta.Err()
		return mvb
	}

	info := meta.Ok()

	mvb.duration = info.Duration.ToFloat().UnwrapOrDefault()
	mvb.Width(info.Width)
	mvb.Height(info.Height)

	if !mvb.duration.IsZero() {
		mvb.Duration(int64(mvb.duration))
	}

	return mvb
}

// GenerateThumbnail automatically generates a thumbnail from the video at the specified seek time.
func (mvb *MediaVideoBuilder) GenerateThumbnail(seek ...String) *MediaVideoBuilder {
	if mvb.file == nil {
		mvb.err = Errorf("video file is not set")
		return mvb
	}

	path := mvb.file.Path()
	if path.IsErr() {
		mvb.err = path.Err()
		return mvb
	}

	if mvb.duration.IsZero() {
		mvb.err = Errorf("duration not set, call ApplyMetadata() first")
		return mvb
	}

	var seekTime String

	if len(seek) != 0 {
		seekTime = seek[0]
	} else {
		seekTime = mvb.duration.Div(2).Max(1.0).RoundDecimal(3).String()
	}

	thumb := ffmpeg.GenerateThumbnail(path.Ok(), seekTime)
	if thumb.IsErr() {
		mvb.err = thumb.Err()
		return mvb
	}

	mvb.thumb = thumb.Ok()

	reader := mvb.thumb.Open()
	if reader.IsErr() {
		mvb.err = reader.Err()
		return mvb
	}

	mvb.video.Thumbnail = gotgbot.InputFileByReader(mvb.thumb.Name().Std(), reader.Ok().Std())
	mvb.parent.tempfiles.Push(mvb.thumb)

	return mvb
}

// Add completes the video configuration and adds it to the media list.
func (mvb *MediaVideoBuilder) Add() *MediaGroup {
	if mvb.err != nil {
		mvb.parent.err = mvb.err
		return mvb.parent
	}

	mvb.parent.files.Push(mvb.file)
	mvb.parent.media.Push(mvb.video)

	return mvb.parent
}

// Video creates a new media video builder.
func (mg *MediaGroup) Video(filename String) *MediaVideoBuilder {
	if mg.err != nil {
		return &MediaVideoBuilder{parent: mg, err: mg.err}
	}

	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		return &MediaVideoBuilder{parent: mg, err: result.Err()}
	}

	return &MediaVideoBuilder{
		parent: mg,
		video:  gotgbot.InputMediaVideo{Media: result.Ok().Doc},
		file:   result.Ok().File,
	}
}

// Audio adds an audio file to the media group with optional caption.
func (mg *MediaGroup) Audio(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		mg.err = result.Err()
		return mg
	}

	mg.files.Push(result.Ok().File)

	media := gotgbot.InputMediaAudio{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	mg.media.Push(media)

	return mg
}

// Document adds a document to the media group with optional caption.
func (mg *MediaGroup) Document(filename String, caption ...String) *MediaGroup {
	result := tgfile.ProcessFile(filename)
	if result.IsErr() {
		mg.err = result.Err()
		return mg
	}

	mg.files.Push(result.Ok().File)

	media := gotgbot.InputMediaDocument{Media: result.Ok().Doc}

	if len(caption) > 0 {
		media.Caption = caption[0].Std()
	}

	mg.media.Push(media)

	return mg
}

// Silent disables notification for the media group messages.
func (mg *MediaGroup) Silent() *MediaGroup {
	mg.opts.DisableNotification = true
	return mg
}

// Protect enables content protection for the media group messages.
func (mg *MediaGroup) Protect() *MediaGroup {
	mg.opts.ProtectContent = true
	return mg
}

// AllowPaidBroadcast allows the media group to be sent in paid broadcast channels.
func (mg *MediaGroup) AllowPaidBroadcast() *MediaGroup {
	mg.opts.AllowPaidBroadcast = true
	return mg
}

// Thread sets the message thread ID for the media group.
func (mg *MediaGroup) Thread(id int64) *MediaGroup {
	mg.opts.MessageThreadId = id
	return mg
}

// Effect sets a message effect for the media group.
func (mg *MediaGroup) Effect(effect effects.EffectType) *MediaGroup {
	mg.opts.MessageEffectId = effect.String()
	return mg
}

// ReplyTo sets the message ID to reply to.
func (mg *MediaGroup) ReplyTo(messageID int64) *MediaGroup {
	mg.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return mg
}

// Business sets the business connection ID for the media group.
func (mg *MediaGroup) Business(id String) *MediaGroup {
	mg.opts.BusinessConnectionId = id.Std()
	return mg
}

// To sets the target chat ID for the media group.
func (mg *MediaGroup) To(chatID int64) *MediaGroup {
	mg.chatID = Some(chatID)
	return mg
}

// Timeout sets a custom timeout for this request.
func (mg *MediaGroup) Timeout(duration time.Duration) *MediaGroup {
	if mg.opts.RequestOpts == nil {
		mg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	mg.opts.RequestOpts.Timeout = duration

	return mg
}

// APIURL sets a custom API URL for this request.
func (mg *MediaGroup) APIURL(url String) *MediaGroup {
	if mg.opts.RequestOpts == nil {
		mg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	mg.opts.RequestOpts.APIURL = url.Std()

	return mg
}

// Send sends the media group to Telegram and returns the result.
func (mg *MediaGroup) Send() Result[Slice[gotgbot.Message]] {
	if mg.err != nil {
		return Err[Slice[gotgbot.Message]](mg.err)
	}

	if mg.media.Len() == 0 {
		return Err[Slice[gotgbot.Message]](errors.New("no media added to media group"))
	}

	defer func() {
		mg.files.Iter().ForEach(func(file *File) {
			file.Close()
		})

		mg.tempfiles.Iter().ForEach(func(file *File) {
			file.Close()
			file.Remove()
		})
	}()

	chatID := mg.chatID.UnwrapOr(mg.ctx.EffectiveChat.Id)

	return ResultOf[Slice[gotgbot.Message]](mg.ctx.Bot.Raw().SendMediaGroup(chatID, mg.media, mg.opts))
}
