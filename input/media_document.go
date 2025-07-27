package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// MediaDocument represents an input media document builder.
type MediaDocument struct {
	input *gotgbot.InputMediaDocument
}

// NewMediaDocument creates a new MediaDocument builder with the required fields.
func NewMediaDocument(media String) *MediaDocument {
	return &MediaDocument{
		input: &gotgbot.InputMediaDocument{
			Media: gotgbot.InputFileByURL(media.Std()),
		},
	}
}

// Thumbnail sets the thumbnail for the document using an InputFile.
// Note: Thumbnails must be uploaded files, not URLs.
func (md *MediaDocument) Thumbnail(thumbnail gotgbot.InputFile) *MediaDocument {
	md.input.Thumbnail = thumbnail
	return md
}

// Caption sets the caption for the document.
func (md *MediaDocument) Caption(caption String) *MediaDocument {
	md.input.Caption = caption.Std()
	return md
}

// HTML sets parse mode to HTML for the caption.
func (md *MediaDocument) HTML() *MediaDocument {
	md.input.ParseMode = "HTML"
	return md
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (md *MediaDocument) Markdown() *MediaDocument {
	md.input.ParseMode = "MarkdownV2"
	return md
}

// CaptionEntities sets the message entities for the caption.
func (md *MediaDocument) CaptionEntities(e entities.Entities) *MediaDocument {
	md.input.CaptionEntities = e.Std()
	return md
}

// DisableContentTypeDetection disables automatic server-side content type detection.
func (md *MediaDocument) DisableContentTypeDetection() *MediaDocument {
	md.input.DisableContentTypeDetection = true
	return md
}

// Build creates the gotgbot.InputMediaDocument.
func (md *MediaDocument) Build() gotgbot.InputMedia {
	return *md.input
}
