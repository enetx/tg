package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
)

// MessageLocation represents an input location message content builder.
type MessageLocation struct {
	input *gotgbot.InputLocationMessageContent
}

// Location creates a new MessageLocation builder with the required fields.
func Location(latitude, longitude float64) *MessageLocation {
	return &MessageLocation{
		input: &gotgbot.InputLocationMessageContent{
			Latitude:  latitude,
			Longitude: longitude,
		},
	}
}

// HorizontalAccuracy sets the horizontal accuracy of the location in meters.
func (ml *MessageLocation) HorizontalAccuracy(accuracy float64) *MessageLocation {
	ml.input.HorizontalAccuracy = accuracy
	return ml
}

// LivePeriod sets the period in seconds for which the location can be updated.
func (ml *MessageLocation) LivePeriod(period int64) *MessageLocation {
	ml.input.LivePeriod = period
	return ml
}

// Heading sets the direction in which the user is moving, in degrees.
func (ml *MessageLocation) Heading(heading int64) *MessageLocation {
	ml.input.Heading = heading
	return ml
}

// ProximityAlertRadius sets the radius for proximity alerts about approaching users, in meters.
func (ml *MessageLocation) ProximityAlertRadius(radius int64) *MessageLocation {
	ml.input.ProximityAlertRadius = radius
	return ml
}

// Build creates the gotgbot.InputLocationMessageContent.
func (ml *MessageLocation) Build() gotgbot.InputMessageContent {
	return *ml.input
}
