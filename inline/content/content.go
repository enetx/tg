package content

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// Content represents any type that can build Content.
type Content interface {
	Build() gotgbot.InputMessageContent
}

// Text creates a new TextBuilder.
func Text(messageText String) *TextBuilder {
	return NewTextContent(messageText)
}

// Location creates a new LocationBuilder.
func Location(latitude, longitude float64) *LocationBuilder {
	return NewLocationBuilder(latitude, longitude)
}

// Venue creates a new VenueBuilder.
func Venue(latitude, longitude float64, title, address String) *VenueBuilder {
	return NewVenueBuilder(latitude, longitude, title, address)
}

// Contact creates a new ContactBuilder.
func Contact(phoneNumber, firstName String) *ContactBuilder {
	return NewContactBuilder(phoneNumber, firstName)
}

// Compile-time interface checks
var (
	_ Content = (*TextBuilder)(nil)
	_ Content = (*LocationBuilder)(nil)
	_ Content = (*VenueBuilder)(nil)
	_ Content = (*ContactBuilder)(nil)
)
