package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// CachedGif represents an inline query result cached GIF builder.
type CachedGif struct {
	inline *gotgbot.InlineQueryResultCachedGif
}

// NewCachedGif creates a new CachedGif builder with the required fields.
func NewCachedGif(id, gifFileID g.String) *CachedGif {
	return &CachedGif{
		inline: &gotgbot.InlineQueryResultCachedGif{
			Id:        id.Std(),
			GifFileId: gifFileID.Std(),
		},
	}
}

// Title sets the title for the result.
func (c *CachedGif) Title(title g.String) *CachedGif {
	c.inline.Title = title.Std()
	return c
}

// Caption sets the caption for the GIF.
func (c *CachedGif) Caption(caption g.String) *CachedGif {
	c.inline.Caption = caption.Std()
	return c
}

// HTML sets parse mode to HTML for the caption.
func (c *CachedGif) HTML() *CachedGif {
	c.inline.ParseMode = "HTML"
	return c
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (c *CachedGif) Markdown() *CachedGif {
	c.inline.ParseMode = "MarkdownV2"
	return c
}

// CaptionEntities sets the message entities for the caption.
func (c *CachedGif) CaptionEntities(e entities.Entities) *CachedGif {
	c.inline.CaptionEntities = e.Std()
	return c
}

// ShowCaptionAboveMedia sets whether to show the result above the message text.
func (c *CachedGif) ShowCaptionAboveMedia() *CachedGif {
	c.inline.ShowCaptionAboveMedia = true
	return c
}

// Markup sets the inline keyboard attached to the message.
func (c *CachedGif) Markup(kb keyboard.Keyboard) *CachedGif {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			c.inline.ReplyMarkup = &ikm
		}
	}

	return c
}

// InputMessageContent sets the content of the message to be sent instead of the GIF.
func (c *CachedGif) InputMessageContent(message input.MessageContent) *CachedGif {
	c.inline.InputMessageContent = message.Build()
	return c
}

// Build creates the gotgbot.InlineQueryResultCachedGif.
func (c *CachedGif) Build() gotgbot.InlineQueryResult {
	return *c.inline
}
