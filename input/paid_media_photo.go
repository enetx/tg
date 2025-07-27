package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/file"
)

// PaidMediaPhoto represents an input paid media photo builder.
type PaidMediaPhoto struct {
	input *gotgbot.InputPaidMediaPhoto
}

// PaidPhoto creates a new PaidMediaPhoto builder with the required fields.
func PaidPhoto(media file.File) *PaidMediaPhoto {
	return &PaidMediaPhoto{
		input: &gotgbot.InputPaidMediaPhoto{
			Media: media.Doc,
		},
	}
}

// Build creates the gotgbot.InputPaidMediaPhoto.
func (pmp *PaidMediaPhoto) Build() gotgbot.InputPaidMedia {
	return *pmp.input
}
