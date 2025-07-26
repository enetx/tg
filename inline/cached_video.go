package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// CachedVideo represents an inline query result cached video builder.
type CachedVideo struct {
	inline *gotgbot.InlineQueryResultCachedVideo
}

// NewCachedVideo creates a new CachedVideo builder with the required fields.
func NewCachedVideo(id, videoFileID, title g.String) *CachedVideo {
	return &CachedVideo{
		inline: &gotgbot.InlineQueryResultCachedVideo{
			Id:          id.Std(),
			VideoFileId: videoFileID.Std(),
			Title:       title.Std(),
		},
	}
}

// Description sets the short description of the result.
func (c *CachedVideo) Description(desc g.String) *CachedVideo {
	c.inline.Description = desc.Std()
	return c
}

// Caption sets the caption for the video.
func (c *CachedVideo) Caption(caption g.String) *CachedVideo {
	c.inline.Caption = caption.Std()
	return c
}

// HTML sets parse mode to HTML for the caption.
func (c *CachedVideo) HTML() *CachedVideo {
	c.inline.ParseMode = "HTML"
	return c
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (c *CachedVideo) Markdown() *CachedVideo {
	c.inline.ParseMode = "MarkdownV2"
	return c
}

// CaptionEntities sets the message entities for the caption.
func (c *CachedVideo) CaptionEntities(e entities.Entities) *CachedVideo {
	c.inline.CaptionEntities = e.Std()
	return c
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (c *CachedVideo) ShowCaptionAboveMedia() *CachedVideo {
	c.inline.ShowCaptionAboveMedia = true
	return c
}

// Markup sets the inline keyboard attached to the message.
func (c *CachedVideo) Markup(kb keyboard.Keyboard) *CachedVideo {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			c.inline.ReplyMarkup = &ikm
		}
	}

	return c
}

// InputMessageContent sets the content of the message to be sent instead of the video.
func (c *CachedVideo) InputMessageContent(message content.Content) *CachedVideo {
	c.inline.InputMessageContent = message.Build()
	return c
}

// Build creates the gotgbot.InlineQueryResultCachedVideo.
func (c *CachedVideo) Build() gotgbot.InlineQueryResult {
	return *c.inline
}
