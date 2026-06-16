package input

import "github.com/PaulSonOfLars/gotgbot/v2"

// MediaLocation represents an input media location builder.
// It can be attached to a poll question, quiz explanation, or poll option.
type MediaLocation struct {
	input *gotgbot.InputMediaLocation
}

// LocationMedia creates a new MediaLocation builder with the required coordinates.
func LocationMedia(latitude, longitude float64) *MediaLocation {
	return &MediaLocation{
		input: &gotgbot.InputMediaLocation{
			Latitude:  latitude,
			Longitude: longitude,
		},
	}
}

// HorizontalAccuracy sets the radius of uncertainty for the location, measured in meters (0-1500).
func (ml *MediaLocation) HorizontalAccuracy(accuracy float64) *MediaLocation {
	ml.input.HorizontalAccuracy = accuracy
	return ml
}

// BuildPollMedia creates the gotgbot.InputPollMedia for use as poll question or explanation media.
func (ml *MediaLocation) BuildPollMedia() gotgbot.InputPollMedia {
	return *ml.input
}

// BuildPollOptionMedia creates the gotgbot.InputPollOptionMedia for use as poll option media.
func (ml *MediaLocation) BuildPollOptionMedia() gotgbot.InputPollOptionMedia {
	return *ml.input
}
