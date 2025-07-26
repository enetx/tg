package inline

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline/content"
	"github.com/enetx/tg/keyboard"
)

// Venue represents an inline query result venue builder.
type Venue struct {
	inline *gotgbot.InlineQueryResultVenue
}

// NewVenue creates a new Venue builder with the required fields.
func NewVenue(id g.String, latitude, longitude float64, title, address g.String) *Venue {
	return &Venue{
		inline: &gotgbot.InlineQueryResultVenue{
			Id:        id.Std(),
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title.Std(),
			Address:   address.Std(),
		},
	}
}

// FoursquareID sets the Foursquare identifier of the venue.
func (v *Venue) FoursquareID(id g.String) *Venue {
	v.inline.FoursquareId = id.Std()
	return v
}

// FoursquareType sets the Foursquare type of the venue.
func (v *Venue) FoursquareType(venueType g.String) *Venue {
	v.inline.FoursquareType = venueType.Std()
	return v
}

// GooglePlaceID sets the Google Places identifier of the venue.
func (v *Venue) GooglePlaceID(id g.String) *Venue {
	v.inline.GooglePlaceId = id.Std()
	return v
}

// GooglePlaceType sets the Google Places type of the venue.
func (v *Venue) GooglePlaceType(venueType g.String) *Venue {
	v.inline.GooglePlaceType = venueType.Std()
	return v
}

// Markup sets the inline keyboard attached to the message.
func (v *Venue) Markup(kb keyboard.Keyboard) *Venue {
	if markup := kb.Markup(); markup != nil {
		if ikm, ok := markup.(gotgbot.InlineKeyboardMarkup); ok {
			v.inline.ReplyMarkup = &ikm
		}
	}

	return v
}

// ThumbnailURL sets the URL of the thumbnail for the result.
func (v *Venue) ThumbnailURL(url g.String) *Venue {
	v.inline.ThumbnailUrl = url.Std()
	return v
}

// ThumbnailSize sets the thumbnail width and height.
func (v *Venue) ThumbnailSize(width, height int64) *Venue {
	v.inline.ThumbnailWidth = width
	v.inline.ThumbnailHeight = height

	return v
}

// InputMessageContent sets the content of the message to be sent instead of the venue.
func (v *Venue) InputMessageContent(message content.Content) *Venue {
	v.inline.InputMessageContent = message.Build()
	return v
}

// Build creates the gotgbot.InlineQueryResultVenue.
func (v *Venue) Build() gotgbot.InlineQueryResult {
	return *v.inline
}
