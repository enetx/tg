package input

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
)

// PaidMediaVideo represents an input paid media video builder.
type PaidMediaVideo struct {
	input *gotgbot.InputPaidMediaVideo
}

// PaidVideo creates a new PaidMediaVideo builder with the required fields.
func PaidVideo(media file.InputFile) *PaidMediaVideo {
	return &PaidMediaVideo{
		input: &gotgbot.InputPaidMediaVideo{
			Media: media.Doc,
		},
	}
}

// Cover sets a cover image for the video.
func (pmv *PaidMediaVideo) Cover(cover g.String) *PaidMediaVideo {
	pmv.input.Cover = cover.Std()
	return pmv
}

// Thumbnail sets the thumbnail for the video using an InputFile.
func (pmv *PaidMediaVideo) Thumbnail(thumbnail file.InputFile) *PaidMediaVideo {
	pmv.input.Thumbnail = thumbnail.Doc.(gotgbot.InputFile)
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
func (pmv *PaidMediaVideo) Duration(duration time.Duration) *PaidMediaVideo {
	pmv.input.Duration = int64(duration.Seconds())
	return pmv
}

// StartAt sets the video start timestamp from the beginning.
func (pmv *PaidMediaVideo) StartAt(offset time.Duration) *PaidMediaVideo {
	pmv.input.StartTimestamp = int64(offset.Seconds())
	return pmv
}

// Streamable sets whether the video supports streaming.
func (pmv *PaidMediaVideo) Streamable() *PaidMediaVideo {
	pmv.input.SupportsStreaming = true
	return pmv
}

// Build creates the gotgbot.InputPaidMediaVideo.
func (pmv *PaidMediaVideo) Build() gotgbot.InputPaidMedia {
	return *pmv.input
}
