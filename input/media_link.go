package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// MediaLink represents an input media link builder.
// It can be attached to a poll option only.
type MediaLink struct {
	input *gotgbot.InputMediaLink
}

// LinkMedia creates a new MediaLink builder with the required HTTP URL.
func LinkMedia(url g.String) *MediaLink {
	return &MediaLink{
		input: &gotgbot.InputMediaLink{
			Url: url.Std(),
		},
	}
}

// URL sets the HTTP URL of the link.
func (ml *MediaLink) URL(url g.String) *MediaLink {
	ml.input.Url = url.Std()
	return ml
}

// BuildPollOptionMedia creates the gotgbot.InputPollOptionMedia for use as poll option media.
func (ml *MediaLink) BuildPollOptionMedia() gotgbot.InputPollOptionMedia {
	return *ml.input
}
