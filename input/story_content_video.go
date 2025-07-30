package input

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// StoryContentVideo represents an input story content video builder.
type StoryContentVideo struct {
	input *gotgbot.InputStoryContentVideo
}

// StoryVideo creates a new StoryContentVideo builder with the required fields.
func StoryVideo(video g.String) *StoryContentVideo {
	return &StoryContentVideo{
		input: &gotgbot.InputStoryContentVideo{
			Video: video.Std(),
		},
	}
}

// Duration sets the precise duration of the video in seconds.
func (scv *StoryContentVideo) Duration(duration time.Duration) *StoryContentVideo {
	scv.input.Duration = duration.Seconds()
	return scv
}

// CoverFrameTimestamp sets the timestamp in seconds of the frame to use as static cover.
func (scv *StoryContentVideo) CoverFrameTimestamp(timestamp time.Duration) *StoryContentVideo {
	scv.input.CoverFrameTimestamp = timestamp.Seconds()
	return scv
}

// Animation sets whether the video has no sound (is an animation).
func (scv *StoryContentVideo) Animation() *StoryContentVideo {
	scv.input.IsAnimation = true
	return scv
}

// Build creates the gotgbot.InputStoryContentVideo.
func (scv *StoryContentVideo) Build() gotgbot.InputStoryContent {
	return *scv.input
}
