package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// MediaAnimation represents an input media animation builder.
type MediaAnimation struct {
	input *gotgbot.InputMediaAnimation
	file  *File
}

// NewMediaAnimation creates a new MediaAnimation builder with the required fields.
func NewMediaAnimation(media String) *MediaAnimation {
	return &MediaAnimation{
		input: &gotgbot.InputMediaAnimation{
			Media: gotgbot.InputFileByURL(media.Std()),
		},
	}
}

// Thumbnail sets the thumbnail for the animation using an InputFile.
// Note: Thumbnails must be uploaded files, not URLs.
func (ma *MediaAnimation) Thumbnail(thumbnail gotgbot.InputFile) *MediaAnimation {
	ma.input.Thumbnail = thumbnail
	return ma
}

// Caption sets the caption for the animation.
func (ma *MediaAnimation) Caption(caption String) *MediaAnimation {
	ma.input.Caption = caption.Std()
	return ma
}

// HTML sets parse mode to HTML for the caption.
func (ma *MediaAnimation) HTML() *MediaAnimation {
	ma.input.ParseMode = "HTML"
	return ma
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (ma *MediaAnimation) Markdown() *MediaAnimation {
	ma.input.ParseMode = "MarkdownV2"
	return ma
}

// CaptionEntities sets the message entities for the caption.
func (ma *MediaAnimation) CaptionEntities(e entities.Entities) *MediaAnimation {
	ma.input.CaptionEntities = e.Std()
	return ma
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (ma *MediaAnimation) ShowCaptionAboveMedia() *MediaAnimation {
	ma.input.ShowCaptionAboveMedia = true
	return ma
}

// Size sets the animation width and height.
func (ma *MediaAnimation) Size(width, height int64) *MediaAnimation {
	ma.input.Width = width
	ma.input.Height = height

	return ma
}

// Duration sets the animation duration in seconds.
func (ma *MediaAnimation) Duration(duration int64) *MediaAnimation {
	ma.input.Duration = duration
	return ma
}

// HasSpoiler sets whether the animation has a spoiler.
func (ma *MediaAnimation) HasSpoiler() *MediaAnimation {
	ma.input.HasSpoiler = true
	return ma
}

// Build creates the gotgbot.InputMediaAnimation.
func (ma *MediaAnimation) Build() gotgbot.InputMedia {
	return *ma.input
}
