package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// PaidMediaVideo represents an input paid media video builder.
type PaidMediaVideo struct {
	input *gotgbot.InputPaidMediaVideo
}

// NewPaidMediaVideo creates a new PaidMediaVideo builder with the required fields.
func NewPaidMediaVideo(media String) *PaidMediaVideo {
	return &PaidMediaVideo{
		input: &gotgbot.InputPaidMediaVideo{
			Media: gotgbot.InputFileByURL(media.Std()),
		},
	}
}

// Thumbnail sets the thumbnail for the video using an InputFile.
// Note: Thumbnails must be uploaded files, not URLs.
func (pmv *PaidMediaVideo) Thumbnail(thumbnail gotgbot.InputFile) *PaidMediaVideo {
	pmv.input.Thumbnail = thumbnail
	return pmv
}

// Width sets the video width.
func (pmv *PaidMediaVideo) Width(width int64) *PaidMediaVideo {
	pmv.input.Width = width
	return pmv
}

// Height sets the video height.
func (pmv *PaidMediaVideo) Height(height int64) *PaidMediaVideo {
	pmv.input.Height = height
	return pmv
}

// Duration sets the video duration in seconds.
func (pmv *PaidMediaVideo) Duration(duration int64) *PaidMediaVideo {
	pmv.input.Duration = duration
	return pmv
}

// SupportsStreaming sets whether the video supports streaming.
func (pmv *PaidMediaVideo) SupportsStreaming() *PaidMediaVideo {
	pmv.input.SupportsStreaming = true
	return pmv
}

// Build creates the gotgbot.InputPaidMediaVideo.
func (pmv *PaidMediaVideo) Build() gotgbot.InputPaidMedia {
	return *pmv.input
}
