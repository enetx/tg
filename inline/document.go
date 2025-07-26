package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// Document represents an inline query result document builder.
type Document struct {
	inline *gotgbot.InlineQueryResultDocument
}

// NewDocument creates a new Document builder with the required fields.
func NewDocument(id, title, documentURL, mimeType g.String) *Document {
	return &Document{
		inline: &gotgbot.InlineQueryResultDocument{
			Id:          id.Std(),
			Title:       title.Std(),
			DocumentUrl: documentURL.Std(),
			MimeType:    mimeType.Std(),
		},
	}
}

// Caption sets the caption for the document.
func (d *Document) Caption(caption g.String) *Document {
	d.inline.Caption = caption.Std()
	return d
}

// HTML sets parse mode to HTML for the caption.
func (d *Document) HTML() *Document {
	d.inline.ParseMode = "HTML"
	return d
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (d *Document) Markdown() *Document {
	d.inline.ParseMode = "MarkdownV2"
	return d
}

// CaptionEntities sets the message entities for the caption.
func (d *Document) CaptionEntities(e entities.Entities) *Document {
	d.inline.CaptionEntities = e.Std()
	return d
}

// Description sets the short description of the result.
func (d *Document) Description(desc g.String) *Document {
	d.inline.Description = desc.Std()
	return d
}

// Markup sets the inline keyboard attached to the message.
func (d *Document) Markup(kb keyboard.Keyboard) *Document {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			d.inline.ReplyMarkup = &ikm
		}
	}

	return d
}

// ThumbnailURL sets the URL of the thumbnail for the result.
func (d *Document) ThumbnailURL(url g.String) *Document {
	d.inline.ThumbnailUrl = url.Std()
	return d
}

// ThumbnailSize sets the thumbnail width and height.
func (d *Document) ThumbnailSize(width, height int64) *Document {
	d.inline.ThumbnailWidth = width
	d.inline.ThumbnailHeight = height

	return d
}

// InputMessageContent sets the content of the message to be sent instead of the document.
func (d *Document) InputMessageContent(message content.Content) *Document {
	d.inline.InputMessageContent = message.Build()
	return d
}

// Build creates the gotgbot.InlineQueryResultDocument.
func (d *Document) Build() gotgbot.InlineQueryResult {
	return *d.inline
}
