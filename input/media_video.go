package input

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/file"
)

// MediaVideo represents an input media video builder.
type MediaVideo struct {
	input *gotgbot.InputMediaVideo
}

// Video creates a new MediaVideo builder with the required fields.
func Video(media file.InputFile) *MediaVideo {
	return &MediaVideo{
		input: &gotgbot.InputMediaVideo{
			Media: media.Doc,
		},
	}
}

// Cover sets a cover image for the video.
func (mv *MediaVideo) Cover(cover g.String) *MediaVideo {
	mv.input.Cover = cover.Std()
	return mv
}

// Thumbnail sets the thumbnail for the video using an InputFile.
func (mv *MediaVideo) Thumbnail(thumbnail file.InputFile) *MediaVideo {
	mv.input.Thumbnail = thumbnail.Doc.(gotgbot.InputFile)
	return mv
}

// Caption sets the caption for the video.
func (mv *MediaVideo) Caption(caption g.String) *MediaVideo {
	mv.input.Caption = caption.Std()
	return mv
}

// HTML sets parse mode to HTML for the caption.
func (mv *MediaVideo) HTML() *MediaVideo {
	mv.input.ParseMode = "HTML"
	return mv
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (mv *MediaVideo) Markdown() *MediaVideo {
	mv.input.ParseMode = "MarkdownV2"
	return mv
}

// CaptionEntities sets the message entities for the caption.
func (mv *MediaVideo) CaptionEntities(e entities.Entities) *MediaVideo {
	mv.input.CaptionEntities = e.Std()
	return mv
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (mv *MediaVideo) ShowCaptionAboveMedia() *MediaVideo {
	mv.input.ShowCaptionAboveMedia = true
	return mv
}

// Size sets the video width and height.
func (mv *MediaVideo) Size(width, height int64) *MediaVideo {
	mv.input.Width = width
	mv.input.Height = height
	return mv
}

// Duration sets the video duration in seconds.
func (mv *MediaVideo) Duration(duration time.Duration) *MediaVideo {
	mv.input.Duration = int64(duration.Seconds())
	return mv
}

// StartAt sets the video start timestamp from the beginning.
func (mv *MediaVideo) StartAt(offset time.Duration) *MediaVideo {
	mv.input.StartTimestamp = int64(offset.Seconds())
	return mv
}

// Streamable enables streaming support for the video.
func (mv *MediaVideo) Streamable() *MediaVideo {
	mv.input.SupportsStreaming = true
	return mv
}

// Spoiler sets whether the video has a spoiler.
func (mv *MediaVideo) Spoiler() *MediaVideo {
	mv.input.HasSpoiler = true
	return mv
}

// Build creates the gotgbot.InputMediaVideo.
func (mv *MediaVideo) Build() gotgbot.InputMedia {
	return *mv.input
}
