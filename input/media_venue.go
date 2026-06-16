package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// MediaVenue represents an input media venue builder.
// It can be attached to a poll question, quiz explanation, or poll option.
type MediaVenue struct {
	input *gotgbot.InputMediaVenue
}

// VenueMedia creates a new MediaVenue builder with the required fields.
func VenueMedia(latitude, longitude float64, title, address g.String) *MediaVenue {
	return &MediaVenue{
		input: &gotgbot.InputMediaVenue{
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title.Std(),
			Address:   address.Std(),
		},
	}
}

// Foursquare sets the Foursquare identifier and type of the venue.
func (mv *MediaVenue) Foursquare(id, kind g.String) *MediaVenue {
	mv.input.FoursquareId = id.Std()
	mv.input.FoursquareType = kind.Std()
	return mv
}

// GooglePlace sets the Google Places identifier and type of the venue.
func (mv *MediaVenue) GooglePlace(id, kind g.String) *MediaVenue {
	mv.input.GooglePlaceId = id.Std()
	mv.input.GooglePlaceType = kind.Std()
	return mv
}

// BuildPollMedia creates the gotgbot.InputPollMedia for use as poll question or explanation media.
func (mv *MediaVenue) BuildPollMedia() gotgbot.InputPollMedia {
	return *mv.input
}

// BuildPollOptionMedia creates the gotgbot.InputPollOptionMedia for use as poll option media.
func (mv *MediaVenue) BuildPollOptionMedia() gotgbot.InputPollOptionMedia {
	return *mv.input
}
