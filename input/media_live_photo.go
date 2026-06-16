package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/file"
)

// MediaLivePhoto represents an input media live photo builder.
// A live photo consists of a short video (up to 10 seconds, 10 MB) and a static cover photo.
// It can be used in media groups as well as in poll questions, explanations, and options.
type MediaLivePhoto struct {
	input *gotgbot.InputMediaLivePhoto
}

// LivePhoto creates a new MediaLivePhoto builder. media is the short video and photo
// is the static cover photo (a file_id, URL, or "attach://<name>" reference).
func LivePhoto(media file.InputFile, photo g.String) *MediaLivePhoto {
	return &MediaLivePhoto{
		input: &gotgbot.InputMediaLivePhoto{
			Media: media.Doc,
			Photo: photo.Std(),
		},
	}
}

// Caption sets the caption for the live photo.
func (mlp *MediaLivePhoto) Caption(caption g.String) *MediaLivePhoto {
	mlp.input.Caption = caption.Std()
	return mlp
}

// HTML sets parse mode to HTML for the caption.
func (mlp *MediaLivePhoto) HTML() *MediaLivePhoto {
	mlp.input.ParseMode = "HTML"
	return mlp
}

// Markdown sets parse mode to MarkdownV2 for the caption.
func (mlp *MediaLivePhoto) Markdown() *MediaLivePhoto {
	mlp.input.ParseMode = "MarkdownV2"
	return mlp
}

// CaptionEntities sets the message entities for the caption.
func (mlp *MediaLivePhoto) CaptionEntities(e entities.Entities) *MediaLivePhoto {
	mlp.input.CaptionEntities = e.Std()
	return mlp
}

// ShowCaptionAboveMedia sets whether to show the caption above the media.
func (mlp *MediaLivePhoto) ShowCaptionAboveMedia() *MediaLivePhoto {
	mlp.input.ShowCaptionAboveMedia = true
	return mlp
}

// Spoiler sets whether the live photo has a spoiler.
func (mlp *MediaLivePhoto) Spoiler() *MediaLivePhoto {
	mlp.input.HasSpoiler = true
	return mlp
}

// Build creates the gotgbot.InputMediaLivePhoto for use in a media group.
func (mlp *MediaLivePhoto) Build() gotgbot.InputMedia {
	return *mlp.input
}

// BuildPollMedia creates the gotgbot.InputPollMedia for use as poll question or explanation media.
func (mlp *MediaLivePhoto) BuildPollMedia() gotgbot.InputPollMedia {
	return *mlp.input
}

// BuildPollOptionMedia creates the gotgbot.InputPollOptionMedia for use as poll option media.
func (mlp *MediaLivePhoto) BuildPollOptionMedia() gotgbot.InputPollOptionMedia {
	return *mlp.input
}
