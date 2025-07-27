package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// ProfilePhotoStatic represents an input profile photo static builder.
type ProfilePhotoStatic struct {
	input *gotgbot.InputProfilePhotoStatic
}

// NewProfilePhotoStatic creates a new ProfilePhotoStatic builder with the required fields.
func NewProfilePhotoStatic(photo String) *ProfilePhotoStatic {
	return &ProfilePhotoStatic{
		input: &gotgbot.InputProfilePhotoStatic{
			Photo: photo.Std(),
		},
	}
}

// Build creates the gotgbot.InputProfilePhotoStatic.
func (pps *ProfilePhotoStatic) Build() gotgbot.InputProfilePhoto {
	return *pps.input
}