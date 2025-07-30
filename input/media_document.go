package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/file"
)

// MediaDocument represents an input media document builder.
type MediaDocument struct {
	input *gotgbot.InputMediaDocument
}

// Document creates a new MediaDocument builder with the required fields.
func Document(media file.InputFile) *MediaDocument {
	return &MediaDocument{
		input: &gotgbot.InputMediaDocument{
			Media: media.Doc,
		},
	}
}

// Thumbnail sets the thumbnail for the document using an File.
func (md *MediaDocument) Thumbnail(thumbnail file.InputFile) *MediaDocument {
	md.input.Thumbnail = thumbnail.Doc.(gotgbot.InputFile)
	return md
}

// Caption sets the caption for the document.
func (md *MediaDocument) Caption(caption g.String) *MediaDocument {
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
