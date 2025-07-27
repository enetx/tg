package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// PaidMediaPhoto represents an input paid media photo builder.
type PaidMediaPhoto struct {
	input *gotgbot.InputPaidMediaPhoto
}

// NewPaidMediaPhoto creates a new PaidMediaPhoto builder with the required fields.
func NewPaidMediaPhoto(media String) *PaidMediaPhoto {
	return &PaidMediaPhoto{
		input: &gotgbot.InputPaidMediaPhoto{
			Media: gotgbot.InputFileByURL(media.Std()),
		},
	}
}

// Build creates the gotgbot.InputPaidMediaPhoto.
func (pmp *PaidMediaPhoto) Build() gotgbot.InputPaidMedia {
	return *pmp.input
}