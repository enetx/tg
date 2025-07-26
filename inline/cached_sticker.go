package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// CachedSticker represents an inline query result cached sticker builder.
type CachedSticker struct {
	inline *gotgbot.InlineQueryResultCachedSticker
}

// NewCachedSticker creates a new CachedSticker builder with the required fields.
func NewCachedSticker(id, stickerFileID g.String) *CachedSticker {
	return &CachedSticker{
		inline: &gotgbot.InlineQueryResultCachedSticker{
			Id:            id.Std(),
			StickerFileId: stickerFileID.Std(),
		},
	}
}

// Markup sets the inline keyboard attached to the message.
func (c *CachedSticker) Markup(kb keyboard.Keyboard) *CachedSticker {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			c.inline.ReplyMarkup = &ikm
		}
	}

	return c
}

// InputMessageContent sets the content of the message to be sent instead of the sticker.
func (c *CachedSticker) InputMessageContent(message content.Content) *CachedSticker {
	c.inline.InputMessageContent = message.Build()
	return c
}

// Build creates the gotgbot.InlineQueryResultCachedSticker.
func (c *CachedSticker) Build() gotgbot.InlineQueryResult {
	return *c.inline
}
