package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/file"
)

// MediaAudio represents an input media audio builder.
type MediaAudio struct {
	input *gotgbot.InputMediaAudio
}

// Audio creates a new MediaAudio builder with the required fields.
func Audio(media file.File) *MediaAudio {
	return &MediaAudio{
		input: &gotgbot.InputMediaAudio{
			Media: media.Doc,
		},
	}
}

// Thumbnail sets the thumbnail for the audio using an InputFile.
// Note: Thumbnails must be uploaded files, not URLs.
func (ma *MediaAudio) Thumbnail(thumbnail gotgbot.InputFile) *MediaAudio {
	ma.input.Thumbnail = thumbnail
	return ma
}

// Caption sets the caption for the audio.
func (ma *MediaAudio) Caption(caption g.String) *MediaAudio {
	ma.input.Caption = caption.Std()
	return ma
}

// HTML sets parse mode to HTML for the caption.
func (ma *MediaAudio) HTML() *MediaAudio {
	ma.input.ParseMode = "HTML"
	return ma
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (ma *MediaAudio) Markdown() *MediaAudio {
	ma.input.ParseMode = "MarkdownV2"
	return ma
}

// CaptionEntities sets the message entities for the caption.
func (ma *MediaAudio) CaptionEntities(e entities.Entities) *MediaAudio {
	ma.input.CaptionEntities = e.Std()
	return ma
}

// Duration sets the audio duration in seconds.
func (ma *MediaAudio) Duration(duration int64) *MediaAudio {
	ma.input.Duration = duration
	return ma
}

// Performer sets the performer of the audio.
func (ma *MediaAudio) Performer(performer g.String) *MediaAudio {
	ma.input.Performer = performer.Std()
	return ma
}

// Title sets the title of the audio.
func (ma *MediaAudio) Title(title g.String) *MediaAudio {
	ma.input.Title = title.Std()
	return ma
}

// Build creates the gotgbot.InputMediaAudio.
func (ma *MediaAudio) Build() gotgbot.InputMedia {
	return *ma.input
}
