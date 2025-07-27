package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// CachedMpeg4Gif represents an inline query result cached MPEG4 GIF builder.
type CachedMpeg4Gif struct {
	inline *gotgbot.InlineQueryResultCachedMpeg4Gif
}

// NewCachedMpeg4Gif creates a new CachedMpeg4Gif builder with the required fields.
func NewCachedMpeg4Gif(id, mpeg4FileID g.String) *CachedMpeg4Gif {
	return &CachedMpeg4Gif{
		inline: &gotgbot.InlineQueryResultCachedMpeg4Gif{
			Id:          id.Std(),
			Mpeg4FileId: mpeg4FileID.Std(),
		},
	}
}

// Title sets the title for the result.
func (c *CachedMpeg4Gif) Title(title g.String) *CachedMpeg4Gif {
	c.inline.Title = title.Std()
	return c
}

// Caption sets the caption for the MPEG4 GIF.
func (c *CachedMpeg4Gif) Caption(caption g.String) *CachedMpeg4Gif {
	c.inline.Caption = caption.Std()
	return c
}

// HTML sets parse mode to HTML for the caption.
func (c *CachedMpeg4Gif) HTML() *CachedMpeg4Gif {
	c.inline.ParseMode = "HTML"
	return c
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (c *CachedMpeg4Gif) Markdown() *CachedMpeg4Gif {
	c.inline.ParseMode = "MarkdownV2"
	return c
}

// CaptionEntities sets the message entities for the caption.
func (c *CachedMpeg4Gif) CaptionEntities(e entities.Entities) *CachedMpeg4Gif {
	c.inline.CaptionEntities = e.Std()
	return c
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (c *CachedMpeg4Gif) ShowCaptionAboveMedia() *CachedMpeg4Gif {
	c.inline.ShowCaptionAboveMedia = true
	return c
}

// Markup sets the inline keyboard attached to the message.
func (c *CachedMpeg4Gif) Markup(kb keyboard.Keyboard) *CachedMpeg4Gif {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			c.inline.ReplyMarkup = &ikm
		}
	}
	return c
}

// InputMessageContent sets the content of the message to be sent instead of the MPEG4 GIF.
func (c *CachedMpeg4Gif) InputMessageContent(message input.MessageContent) *CachedMpeg4Gif {
	c.inline.InputMessageContent = message.Build()
	return c
}

// Build creates the gotgbot.InlineQueryResultCachedMpeg4Gif.
func (c *CachedMpeg4Gif) Build() gotgbot.InlineQueryResult {
	return *c.inline
}
