package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// CachedAudio represents an inline query result cached audio builder.
type CachedAudio struct {
	inline *gotgbot.InlineQueryResultCachedAudio
}

// NewCachedAudio creates a new CachedAudio builder with the required fields.
func NewCachedAudio(id, audioFileID g.String) *CachedAudio {
	return &CachedAudio{
		inline: &gotgbot.InlineQueryResultCachedAudio{
			Id:          id.Std(),
			AudioFileId: audioFileID.Std(),
		},
	}
}

// Caption sets the caption for the audio.
func (c *CachedAudio) Caption(caption g.String) *CachedAudio {
	c.inline.Caption = caption.Std()
	return c
}

// HTML sets parse mode to HTML for the caption.
func (c *CachedAudio) HTML() *CachedAudio {
	c.inline.ParseMode = "HTML"
	return c
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (c *CachedAudio) Markdown() *CachedAudio {
	c.inline.ParseMode = "MarkdownV2"
	return c
}

// CaptionEntities sets the message entities for the caption.
func (c *CachedAudio) CaptionEntities(e entities.Entities) *CachedAudio {
	c.inline.CaptionEntities = e.Std()
	return c
}

// Markup sets the inline keyboard attached to the message.
func (c *CachedAudio) Markup(kb keyboard.Keyboard) *CachedAudio {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			c.inline.ReplyMarkup = &ikm
		}
	}

	return c
}

// InputMessageContent sets the content of the message to be sent instead of the audio.
func (c *CachedAudio) InputMessageContent(message input.MessageContent) *CachedAudio {
	c.inline.InputMessageContent = message.Build()
	return c
}

// Build creates the gotgbot.InlineQueryResultCachedAudio.
func (c *CachedAudio) Build() gotgbot.InlineQueryResult {
	return *c.inline
}
