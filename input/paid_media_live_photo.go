package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
)

// PaidMediaLivePhoto represents an input paid media live photo builder.
// A live photo consists of a short video (up to 10 seconds, 10 MB) and a static cover photo.
type PaidMediaLivePhoto struct {
	input *gotgbot.InputPaidMediaLivePhoto
}

// PaidLivePhoto creates a new PaidMediaLivePhoto builder. media is the short video and
// photo is the static cover photo (a file_id, URL, or "attach://<name>" reference).
func PaidLivePhoto(media file.InputFile, photo g.String) *PaidMediaLivePhoto {
	return &PaidMediaLivePhoto{
		input: &gotgbot.InputPaidMediaLivePhoto{
			Media: media.Doc,
			Photo: photo.Std(),
		},
	}
}

// Build creates the gotgbot.InputPaidMediaLivePhoto.
func (pmlp *PaidMediaLivePhoto) Build() gotgbot.InputPaidMedia {
	return *pmlp.input
}
