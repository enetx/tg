package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

func TestNewVenue(t *testing.T) {
	venue := inline.NewVenue(testID, 40.7128, -74.0060, testTitle, g.String("New York, NY"))
	if venue == nil {
		t.Error("Expected Venue to be created")
	}

	built := venue.Build()
	if built == nil {
		t.Error("Expected Venue to build correctly")
	}

	if built.GetType() != "venue" {
		t.Error("Expected Venue type to be 'venue'")
	}
}

func TestVenue_FoursquareID(t *testing.T) {
	venue := inline.NewVenue(testID, 40.7128, -74.0060, testTitle, g.String("New York, NY"))
	id := g.String("fs-id-123")

	result := venue.FoursquareID(id)
	if result == nil {
		t.Error("Expected FoursquareID method to return Venue")
	}

	built := result.Build()
	if venueResult, ok := built.(gotgbot.InlineQueryResultVenue); ok {
		if venueResult.FoursquareId != id.Std() {
			t.Error("Expected FoursquareId to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVenue")
	}
}

func TestVenue_FoursquareType(t *testing.T) {
	venue := inline.NewVenue(testID, 40.7128, -74.0060, testTitle, g.String("New York, NY"))
	typ := g.String("arts/landmark")

	result := venue.FoursquareType(typ)
	if result == nil {
		t.Error("Expected FoursquareType method to return Venue")
	}

	built := result.Build()
	if venueResult, ok := built.(gotgbot.InlineQueryResultVenue); ok {
		if venueResult.FoursquareType != typ.Std() {
			t.Error("Expected FoursquareType to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVenue")
	}
}

func TestVenue_GooglePlaceID(t *testing.T) {
	venue := inline.NewVenue(testID, 40.7128, -74.0060, testTitle, g.String("New York, NY"))
	id := g.String("gplace-123")

	result := venue.GooglePlaceID(id)
	if result == nil {
		t.Error("Expected GooglePlaceID method to return Venue")
	}

	built := result.Build()
	if venueResult, ok := built.(gotgbot.InlineQueryResultVenue); ok {
		if venueResult.GooglePlaceId != id.Std() {
			t.Error("Expected GooglePlaceId to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVenue")
	}
}

func TestVenue_GooglePlaceType(t *testing.T) {
	venue := inline.NewVenue(testID, 40.7128, -74.0060, testTitle, g.String("New York, NY"))
	typ := g.String("establishment")

	result := venue.GooglePlaceType(typ)
	if result == nil {
		t.Error("Expected GooglePlaceType method to return Venue")
	}

	built := result.Build()
	if venueResult, ok := built.(gotgbot.InlineQueryResultVenue); ok {
		if venueResult.GooglePlaceType != typ.Std() {
			t.Error("Expected GooglePlaceType to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVenue")
	}
}

func TestVenue_ThumbnailURL(t *testing.T) {
	venue := inline.NewVenue(testID, 40.7128, -74.0060, testTitle, g.String("New York, NY"))

	result := venue.ThumbnailURL(testThumbnailURL)
	if result == nil {
		t.Error("Expected ThumbnailURL method to return Venue")
	}

	built := result.Build()
	if venueResult, ok := built.(gotgbot.InlineQueryResultVenue); ok {
		if venueResult.ThumbnailUrl != testThumbnailURL.Std() {
			t.Error("Expected ThumbnailUrl to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVenue")
	}
}

func TestVenue_ThumbnailSize(t *testing.T) {
	venue := inline.NewVenue(testID, 40.7128, -74.0060, testTitle, g.String("New York, NY"))

	result := venue.ThumbnailSize(150, 150)
	if result == nil {
		t.Error("Expected ThumbnailSize method to return Venue")
	}

	built := result.Build()
	if venueResult, ok := built.(gotgbot.InlineQueryResultVenue); ok {
		if venueResult.ThumbnailWidth != 150 {
			t.Error("Expected ThumbnailWidth to be set correctly")
		}
		if venueResult.ThumbnailHeight != 150 {
			t.Error("Expected ThumbnailHeight to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVenue")
	}
}

func TestVenue_Markup(t *testing.T) {
	venue := inline.NewVenue(testID, 40.7128, -74.0060, testTitle, g.String("New York, NY"))
	keyboard := createTestKeyboard()

	result := venue.Markup(keyboard)
	if result == nil {
		t.Error("Expected Markup method to return Venue")
	}

	built := result.Build()
	if venueResult, ok := built.(gotgbot.InlineQueryResultVenue); ok {
		if venueResult.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVenue")
	}
}

func TestVenue_InputMessageContent(t *testing.T) {
	venue := inline.NewVenue(testID, 40.7128, -74.0060, testTitle, g.String("New York, NY"))
	messageContent := createTestMessageContent()

	result := venue.InputMessageContent(messageContent)
	if result == nil {
		t.Error("Expected InputMessageContent method to return Venue")
	}

	built := result.Build()
	if venueResult, ok := built.(gotgbot.InlineQueryResultVenue); ok {
		if venueResult.InputMessageContent == nil {
			t.Error("Expected InputMessageContent to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVenue")
	}
}

func TestVenue_MethodChaining(t *testing.T) {
	foursquareID := g.String("fs-id-123")
	googleID := g.String("gplace-123")
	msg := createTestMessageContent()

	result := inline.NewVenue(testID, 40.7128, -74.0060, testTitle, g.String("New York, NY")).
		FoursquareID(foursquareID).
		FoursquareType(g.String("arts/landmark")).
		GooglePlaceID(googleID).
		GooglePlaceType(g.String("establishment")).
		ThumbnailURL(testThumbnailURL).
		ThumbnailSize(150, 150).
		Markup(createTestKeyboard()).
		InputMessageContent(msg)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Venue to build correctly")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
