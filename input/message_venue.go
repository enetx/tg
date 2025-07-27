package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// MessageVenue represents an input venue message content builder.
type MessageVenue struct {
	input *gotgbot.InputVenueMessageContent
}

// NewMessageVenue creates a new MessageVenue builder with the required fields.
func NewMessageVenue(latitude, longitude float64, title, address String) *MessageVenue {
	return &MessageVenue{
		input: &gotgbot.InputVenueMessageContent{
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title.Std(),
			Address:   address.Std(),
		},
	}
}

// FoursquareID sets the Foursquare identifier of the venue.
func (mv *MessageVenue) FoursquareID(id String) *MessageVenue {
	mv.input.FoursquareId = id.Std()
	return mv
}

// FoursquareType sets the Foursquare type of the venue.
func (mv *MessageVenue) FoursquareType(venueType String) *MessageVenue {
	mv.input.FoursquareType = venueType.Std()
	return mv
}

// GooglePlaceID sets the Google Places identifier of the venue.
func (mv *MessageVenue) GooglePlaceID(id String) *MessageVenue {
	mv.input.GooglePlaceId = id.Std()
	return mv
}

// GooglePlaceType sets the Google Places type of the venue.
func (mv *MessageVenue) GooglePlaceType(venueType String) *MessageVenue {
	mv.input.GooglePlaceType = venueType.Std()
	return mv
}

// Build creates the gotgbot.InputVenueMessageContent.
func (mv *MessageVenue) Build() gotgbot.InputMessageContent {
	return *mv.input
}