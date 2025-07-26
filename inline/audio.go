package inline

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// Audio represents an inline query result audio builder.
type Audio struct {
	inline *gotgbot.InlineQueryResultAudio
}

// NewAudio creates a new Audio builder with the required fields.
func NewAudio(id, audioURL, title g.String) *Audio {
	return &Audio{
		inline: &gotgbot.InlineQueryResultAudio{
			Id:       id.Std(),
			AudioUrl: audioURL.Std(),
			Title:    title.Std(),
		},
	}
}

// Caption sets the caption for the audio.
func (a *Audio) Caption(caption g.String) *Audio {
	a.inline.Caption = caption.Std()
	return a
}

// HTML sets parse mode to HTML for the caption.
func (a *Audio) HTML() *Audio {
	a.inline.ParseMode = "HTML"
	return a
}

func (a *Audio) Markdown() *Audio {
	a.inline.ParseMode = "MarkdownV2"
	return a
}

// CaptionEntities sets the message entities for the caption.
func (a *Audio) CaptionEntities(e entities.Entities) *Audio {
	a.inline.CaptionEntities = e.Std()
	return a
}

// Performer sets the performer of the audio.
func (a *Audio) Performer(performer g.String) *Audio {
	a.inline.Performer = performer.Std()
	return a
}

// Duration sets the audio duration in seconds.
func (a *Audio) Duration(duration time.Duration) *Audio {
	a.inline.AudioDuration = int64(duration.Seconds())
	return a
}

// Markup sets the inline keyboard attached to the message.
func (a *Audio) Markup(kb keyboard.Keyboard) *Audio {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			a.inline.ReplyMarkup = &ikm
		}
	}

	return a
}

// InputMessageContent sets the content of the message to be sent instead of the audio.
func (a *Audio) InputMessageContent(message content.Content) *Audio {
	a.inline.InputMessageContent = message.Build()
	return a
}

// Build creates the gotgbot.InlineQueryResultAudio.
func (a *Audio) Build() gotgbot.InlineQueryResult {
	return *a.inline
}
