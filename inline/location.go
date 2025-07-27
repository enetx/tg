package inline

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// Location represents an inline query result location builder.
type Location struct {
	inline *gotgbot.InlineQueryResultLocation
}

// NewLocation creates a new Location builder with the required fields.
func NewLocation(id g.String, latitude, longitude float64, title g.String) *Location {
	return &Location{
		inline: &gotgbot.InlineQueryResultLocation{
			Id:        id.Std(),
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title.Std(),
		},
	}
}

// HorizontalAccuracy sets the radius of uncertainty for the location.
func (l *Location) HorizontalAccuracy(accuracy float64) *Location {
	l.inline.HorizontalAccuracy = accuracy
	return l
}

// LiveFor sets the period in seconds during which the location can be updated.
func (l *Location) LiveFor(duration time.Duration) *Location {
	l.inline.LivePeriod = int64(duration.Seconds())
	return l
}

// Heading sets the direction in which the user is moving.
func (l *Location) Heading(heading int64) *Location {
	l.inline.Heading = heading
	return l
}

// ProximityAlertRadius sets the maximum distance for proximity alerts.
func (l *Location) ProximityAlertRadius(radius int64) *Location {
	l.inline.ProximityAlertRadius = radius
	return l
}

// Markup sets the inline keyboard attached to the message.
func (l *Location) Markup(kb keyboard.Keyboard) *Location {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			l.inline.ReplyMarkup = &ikm
		}
	}

	return l
}

// ThumbnailURL sets the URL of the thumbnail for the result.
func (l *Location) ThumbnailURL(url g.String) *Location {
	l.inline.ThumbnailUrl = url.Std()
	return l
}

// ThumbnailSize sets the thumbnail width and height.
func (l *Location) ThumbnailSize(width, height int64) *Location {
	l.inline.ThumbnailWidth = width
	l.inline.ThumbnailHeight = height

	return l
}

// InputMessageContent sets the content of the message to be sent instead of the location.
func (l *Location) InputMessageContent(message input.MessageContent) *Location {
	l.inline.InputMessageContent = message.Build()
	return l
}

// Build creates the gotgbot.InlineQueryResultLocation.
func (l *Location) Build() gotgbot.InlineQueryResult {
	return *l.inline
}
