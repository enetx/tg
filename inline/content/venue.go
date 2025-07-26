package content

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// VenueBuilder represents an input venue message content builder.
type VenueBuilder struct {
	input *gotgbot.InputVenueMessageContent
}

// NewVenueBuilder creates a new VenueBuilder builder.
func NewVenueBuilder(latitude, longitude float64, title, address String) *VenueBuilder {
	return &VenueBuilder{
		input: &gotgbot.InputVenueMessageContent{
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title.Std(),
			Address:   address.Std(),
		},
	}
}

// FoursquareID sets the Foursquare identifier of the venue.
func (v *VenueBuilder) FoursquareID(id String) *VenueBuilder {
	v.input.FoursquareId = id.Std()
	return v
}

// FoursquareType sets the Foursquare type of the venue.
func (v *VenueBuilder) FoursquareType(venueType String) *VenueBuilder {
	v.input.FoursquareType = venueType.Std()
	return v
}

// GooglePlaceID sets the Google Places identifier of the venue.
func (v *VenueBuilder) GooglePlaceID(id String) *VenueBuilder {
	v.input.GooglePlaceId = id.Std()
	return v
}

// GooglePlaceType sets the Google Places type of the venue.
func (v *VenueBuilder) GooglePlaceType(placeType String) *VenueBuilder {
	v.input.GooglePlaceType = placeType.Std()
	return v
}

// Build creates the gotgbot.InputVenueMessageContent.
func (v *VenueBuilder) Build() gotgbot.InputMessageContent {
	return *v.input
}