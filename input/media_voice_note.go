package input

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/file"
)

// MediaVoiceNote represents an input media voice note builder.
// It can be attached to a rich message voice note block only.
type MediaVoiceNote struct {
	input *gotgbot.InputMediaVoiceNote
}

// VoiceNote creates a new MediaVoiceNote builder with the required voice message file.
func VoiceNote(media file.InputFile) *MediaVoiceNote {
	return &MediaVoiceNote{
		input: &gotgbot.InputMediaVoiceNote{
			Media: media.Doc,
		},
	}
}

// Caption sets the caption for the voice message.
func (mvn *MediaVoiceNote) Caption(caption g.String) *MediaVoiceNote {
	mvn.input.Caption = caption.Std()
	return mvn
}

// HTML sets parse mode to HTML for the caption.
func (mvn *MediaVoiceNote) HTML() *MediaVoiceNote {
	mvn.input.ParseMode = gotgbot.ParseModeHTML
	return mvn
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (mvn *MediaVoiceNote) Markdown() *MediaVoiceNote {
	mvn.input.ParseMode = gotgbot.ParseModeMarkdownV2
	return mvn
}

// CaptionEntities sets the message entities for the caption.
func (mvn *MediaVoiceNote) CaptionEntities(e entities.Entities) *MediaVoiceNote {
	mvn.input.CaptionEntities = e.Std()
	return mvn
}

// Duration sets the voice message duration in seconds.
func (mvn *MediaVoiceNote) Duration(duration time.Duration) *MediaVoiceNote {
	mvn.input.Duration = int64(duration.Seconds())
	return mvn
}

// Build creates the gotgbot.InputMediaVoiceNote.
func (mvn *MediaVoiceNote) Build() gotgbot.InputMediaVoiceNote {
	return *mvn.input
}
