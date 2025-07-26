package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// CachedPhoto represents an inline query result cached photo builder.
type CachedPhoto struct {
	inline *gotgbot.InlineQueryResultCachedPhoto
}

// NewCachedPhoto creates a new CachedPhoto builder with the required fields.
func NewCachedPhoto(id, photoFileID g.String) *CachedPhoto {
	return &CachedPhoto{
		inline: &gotgbot.InlineQueryResultCachedPhoto{
			Id:          id.Std(),
			PhotoFileId: photoFileID.Std(),
		},
	}
}

// Title sets the title for the result.
func (c *CachedPhoto) Title(title g.String) *CachedPhoto {
	c.inline.Title = title.Std()
	return c
}

// Description sets the short description of the result.
func (c *CachedPhoto) Description(desc g.String) *CachedPhoto {
	c.inline.Description = desc.Std()
	return c
}

// Caption sets the caption for the photo.
func (c *CachedPhoto) Caption(caption g.String) *CachedPhoto {
	c.inline.Caption = caption.Std()
	return c
}

// HTML sets parse mode to HTML for the caption.
func (c *CachedPhoto) HTML() *CachedPhoto {
	c.inline.ParseMode = "HTML"
	return c
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (c *CachedPhoto) Markdown() *CachedPhoto {
	c.inline.ParseMode = "MarkdownV2"
	return c
}

// CaptionEntities sets the message entities for the caption.
func (c *CachedPhoto) CaptionEntities(e entities.Entities) *CachedPhoto {
	c.inline.CaptionEntities = e.Std()
	return c
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (c *CachedPhoto) ShowCaptionAboveMedia() *CachedPhoto {
	c.inline.ShowCaptionAboveMedia = true
	return c
}

// Markup sets the inline keyboard attached to the message.
func (c *CachedPhoto) Markup(kb keyboard.Keyboard) *CachedPhoto {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			c.inline.ReplyMarkup = &ikm
		}
	}

	return c
}

// InputMessageContent sets the content of the message to be sent instead of the photo.
func (c *CachedPhoto) InputMessageContent(message content.Content) *CachedPhoto {
	c.inline.InputMessageContent = message.Build()
	return c
}

// Build creates the gotgbot.InlineQueryResultCachedPhoto.
func (c *CachedPhoto) Build() gotgbot.InlineQueryResult {
	return *c.inline
}
