package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// MediaPhoto represents an input media photo builder.
type MediaPhoto struct {
	input *gotgbot.InputMediaPhoto
}

// NewMediaPhoto creates a new MediaPhoto builder with the required fields.
func NewMediaPhoto(media String) *MediaPhoto {
	return &MediaPhoto{
		input: &gotgbot.InputMediaPhoto{
			Media: gotgbot.InputFileByURL(media.Std()),
		},
	}
}

// Caption sets the caption for the photo.
func (mp *MediaPhoto) Caption(caption String) *MediaPhoto {
	mp.input.Caption = caption.Std()
	return mp
}

// HTML sets parse mode to HTML for the caption.
func (mp *MediaPhoto) HTML() *MediaPhoto {
	mp.input.ParseMode = "HTML"
	return mp
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (mp *MediaPhoto) Markdown() *MediaPhoto {
	mp.input.ParseMode = "MarkdownV2"
	return mp
}

// CaptionEntities sets the message entities for the caption.
func (mp *MediaPhoto) CaptionEntities(e entities.Entities) *MediaPhoto {
	mp.input.CaptionEntities = e.Std()
	return mp
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (mp *MediaPhoto) ShowCaptionAboveMedia() *MediaPhoto {
	mp.input.ShowCaptionAboveMedia = true
	return mp
}

// HasSpoiler sets whether the photo has a spoiler.
func (mp *MediaPhoto) HasSpoiler() *MediaPhoto {
	mp.input.HasSpoiler = true
	return mp
}

// Build creates the gotgbot.InputMediaPhoto.
func (mp *MediaPhoto) Build() gotgbot.InputMedia {
	return *mp.input
}