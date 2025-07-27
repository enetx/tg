package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// StoryContentPhoto represents an input story content photo builder.
type StoryContentPhoto struct {
	input *gotgbot.InputStoryContentPhoto
}

// NewStoryContentPhoto creates a new StoryContentPhoto builder with the required fields.
func NewStoryContentPhoto(photo String) *StoryContentPhoto {
	return &StoryContentPhoto{
		input: &gotgbot.InputStoryContentPhoto{
			Photo: photo.Std(),
		},
	}
}

// Build creates the gotgbot.InputStoryContentPhoto.
func (scp *StoryContentPhoto) Build() gotgbot.InputStoryContent {
	return *scp.input
}