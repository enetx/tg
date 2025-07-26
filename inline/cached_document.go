package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// CachedDocument represents an inline query result cached document builder.
type CachedDocument struct {
	inline *gotgbot.InlineQueryResultCachedDocument
}

// NewCachedDocument creates a new CachedDocument builder with the required fields.
func NewCachedDocument(id, documentFileID, title g.String) *CachedDocument {
	return &CachedDocument{
		inline: &gotgbot.InlineQueryResultCachedDocument{
			Id:             id.Std(),
			DocumentFileId: documentFileID.Std(),
			Title:          title.Std(),
		},
	}
}

// Description sets the short description of the result.
func (c *CachedDocument) Description(desc g.String) *CachedDocument {
	c.inline.Description = desc.Std()
	return c
}

// Caption sets the caption for the document.
func (c *CachedDocument) Caption(caption g.String) *CachedDocument {
	c.inline.Caption = caption.Std()
	return c
}

// HTML sets parse mode to HTML for the caption.
func (c *CachedDocument) HTML() *CachedDocument {
	c.inline.ParseMode = "HTML"
	return c
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (c *CachedDocument) Markdown() *CachedDocument {
	c.inline.ParseMode = "MarkdownV2"
	return c
}

// CaptionEntities sets the message entities for the caption.
func (c *CachedDocument) CaptionEntities(e entities.Entities) *CachedDocument {
	c.inline.CaptionEntities = e.Std()
	return c
}

// Markup sets the inline keyboard attached to the message.
func (c *CachedDocument) Markup(kb keyboard.Keyboard) *CachedDocument {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			c.inline.ReplyMarkup = &ikm
		}
	}

	return c
}

// InputMessageContent sets the content of the message to be sent instead of the document.
func (c *CachedDocument) InputMessageContent(message content.Content) *CachedDocument {
	c.inline.InputMessageContent = message.Build()
	return c
}

// Build creates the gotgbot.InlineQueryResultCachedDocument.
func (c *CachedDocument) Build() gotgbot.InlineQueryResult {
	return *c.inline
}
