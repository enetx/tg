package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// MediaVideo represents an input media video builder.
type MediaVideo struct {
	input *gotgbot.InputMediaVideo
}

// NewMediaVideo creates a new MediaVideo builder with the required fields.
func NewMediaVideo(media String) *MediaVideo {
	return &MediaVideo{
		input: &gotgbot.InputMediaVideo{
			Media: gotgbot.InputFileByURL(media.Std()),
		},
	}
}

// Thumbnail sets the thumbnail for the video using an InputFile.
// Note: Thumbnails must be uploaded files, not URLs.
func (mv *MediaVideo) Thumbnail(thumbnail gotgbot.InputFile) *MediaVideo {
	mv.input.Thumbnail = thumbnail
	return mv
}

// Caption sets the caption for the video.
func (mv *MediaVideo) Caption(caption String) *MediaVideo {
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
func (mv *MediaVideo) Duration(duration int64) *MediaVideo {
	mv.input.Duration = duration
	return mv
}

// SupportsStreaming sets whether the video supports streaming.
func (mv *MediaVideo) SupportsStreaming() *MediaVideo {
	mv.input.SupportsStreaming = true
	return mv
}

// HasSpoiler sets whether the video has a spoiler.
func (mv *MediaVideo) HasSpoiler() *MediaVideo {
	mv.input.HasSpoiler = true
	return mv
}

// Build creates the gotgbot.InputMediaVideo.
func (mv *MediaVideo) Build() gotgbot.InputMedia {
	return *mv.input
}
