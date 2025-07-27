package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// ProfilePhotoAnimated represents an input profile photo animated builder.
type ProfilePhotoAnimated struct {
	input *gotgbot.InputProfilePhotoAnimated
}

// NewProfilePhotoAnimated creates a new ProfilePhotoAnimated builder with the required fields.
func NewProfilePhotoAnimated(animation String) *ProfilePhotoAnimated {
	return &ProfilePhotoAnimated{
		input: &gotgbot.InputProfilePhotoAnimated{
			Animation: animation.Std(),
		},
	}
}

// MainFrameTimestamp sets the timestamp in seconds of the frame to use as the static profile photo.
func (ppa *ProfilePhotoAnimated) MainFrameTimestamp(timestamp float64) *ProfilePhotoAnimated {
	ppa.input.MainFrameTimestamp = timestamp
	return ppa
}

// Build creates the gotgbot.InputProfilePhotoAnimated.
func (ppa *ProfilePhotoAnimated) Build() gotgbot.InputProfilePhoto {
	return *ppa.input
}