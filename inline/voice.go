package inline

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// Voice represents an inline query result voice builder.
type Voice struct {
	inline *gotgbot.InlineQueryResultVoice
}

// NewVoice creates a new Voice builder with the required fields.
func NewVoice(id, voiceURL, title g.String) *Voice {
	return &Voice{
		inline: &gotgbot.InlineQueryResultVoice{
			Id:       id.Std(),
			VoiceUrl: voiceURL.Std(),
			Title:    title.Std(),
		},
	}
}

// Caption sets the caption for the voice.
func (v *Voice) Caption(caption g.String) *Voice {
	v.inline.Caption = caption.Std()
	return v
}

// HTML sets parse mode to HTML for the caption.
func (v *Voice) HTML() *Voice {
	v.inline.ParseMode = "HTML"
	return v
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (v *Voice) Markdown() *Voice {
	v.inline.ParseMode = "MarkdownV2"
	return v
}

// CaptionEntities sets the message entities for the caption.
func (v *Voice) CaptionEntities(e entities.Entities) *Voice {
	v.inline.CaptionEntities = e.Std()
	return v
}

// Duration sets the voice recording duration.
func (v *Voice) Duration(duration time.Duration) *Voice {
	v.inline.VoiceDuration = int64(duration.Seconds())
	return v
}

// Markup sets the inline keyboard attached to the message.
func (v *Voice) Markup(kb keyboard.Keyboard) *Voice {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			v.inline.ReplyMarkup = &ikm
		}
	}

	return v
}

// InputMessageContent sets the content of the message to be sent instead of the voice.
func (v *Voice) InputMessageContent(message content.Content) *Voice {
	v.inline.InputMessageContent = message.Build()
	return v
}

// Build creates the gotgbot.InlineQueryResultVoice.
func (v *Voice) Build() gotgbot.InlineQueryResult {
	return *v.inline
}
