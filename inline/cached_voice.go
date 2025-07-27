package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// CachedVoice represents an inline query result cached voice builder.
type CachedVoice struct {
	inline *gotgbot.InlineQueryResultCachedVoice
}

// NewCachedVoice creates a new CachedVoice builder with the required fields.
func NewCachedVoice(id, voiceFileID, title g.String) *CachedVoice {
	return &CachedVoice{
		inline: &gotgbot.InlineQueryResultCachedVoice{
			Id:          id.Std(),
			VoiceFileId: voiceFileID.Std(),
			Title:       title.Std(),
		},
	}
}

// Caption sets the caption for the voice.
func (c *CachedVoice) Caption(caption g.String) *CachedVoice {
	c.inline.Caption = caption.Std()
	return c
}

// HTML sets parse mode to HTML for the caption.
func (c *CachedVoice) HTML() *CachedVoice {
	c.inline.ParseMode = "HTML"
	return c
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (c *CachedVoice) Markdown() *CachedVoice {
	c.inline.ParseMode = "MarkdownV2"
	return c
}

// CaptionEntities sets the message entities for the caption.
func (c *CachedVoice) CaptionEntities(e entities.Entities) *CachedVoice {
	c.inline.CaptionEntities = e.Std()
	return c
}

// Markup sets the inline keyboard attached to the message.
func (c *CachedVoice) Markup(kb keyboard.Keyboard) *CachedVoice {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			c.inline.ReplyMarkup = &ikm
		}
	}

	return c
}

// InputMessageContent sets the content of the message to be sent instead of the voice.
func (c *CachedVoice) InputMessageContent(message input.MessageContent) *CachedVoice {
	c.inline.InputMessageContent = message.Build()
	return c
}

// Build creates the gotgbot.InlineQueryResultCachedVoice.
func (c *CachedVoice) Build() gotgbot.InlineQueryResult {
	return *c.inline
}
