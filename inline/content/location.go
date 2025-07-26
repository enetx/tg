package content

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

// LocationBuilder represents an input location message content builder.
type LocationBuilder struct {
	input *gotgbot.InputLocationMessageContent
}

// NewLocationBuilder creates a new LocationBuilder builder.
func NewLocationBuilder(latitude, longitude float64) *LocationBuilder {
	return &LocationBuilder{
		input: &gotgbot.InputLocationMessageContent{
			Latitude:  latitude,
			Longitude: longitude,
		},
	}
}

// HorizontalAccuracy sets the radius of uncertainty for the location in meters (0-1500).
func (l *LocationBuilder) HorizontalAccuracy(accuracy float64) *LocationBuilder {
	l.input.HorizontalAccuracy = accuracy
	return l
}

// LiveFor sets the period in seconds for which the location will be updated.
func (l *LocationBuilder) LiveFor(duration time.Duration) *LocationBuilder {
	l.input.LivePeriod = int64(duration.Seconds())
	return l
}

// Heading sets the direction in which user is moving, in degrees (1-360).
func (l *LocationBuilder) Heading(heading int64) *LocationBuilder {
	l.input.Heading = heading
	return l
}

// ProximityAlertRadius sets the maximum distance for proximity alerts about approaching another chat member, in meters (1-100000).
func (l *LocationBuilder) ProximityAlertRadius(radius int64) *LocationBuilder {
	l.input.ProximityAlertRadius = radius
	return l
}

// Build creates the gotgbot.InputLocationMessageContent.
func (l *LocationBuilder) Build() gotgbot.InputMessageContent {
	return *l.input
}
