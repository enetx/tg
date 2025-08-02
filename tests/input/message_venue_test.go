package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestVenue(t *testing.T) {
	venue := input.Venue(testLatitude, testLongitude, testTitle, testAddress)
	if venue == nil {
		t.Error("Expected MessageVenue to be created")
	}
	if !assertMessageContent(venue) {
		t.Error("MessageVenue should implement MessageContent correctly")
	}
}

func TestVenue_FoursquareID(t *testing.T) {
	venue := input.Venue(testLatitude, testLongitude, testTitle, testAddress)
	foursquareID := g.String("4bcad4e3f964a520e40c21e3")
	result := venue.FoursquareID(foursquareID)
	if result == nil {
		t.Error("Expected FoursquareID method to return MessageVenue")
	}
	if result != venue {
		t.Error("Expected FoursquareID to return same MessageVenue instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputVenueMessageContent); ok {
		if v.FoursquareId != foursquareID.Std() {
			t.Error("Expected FoursquareId to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputVenueMessageContent")
	}
}

func TestVenue_FoursquareType(t *testing.T) {
	venue := input.Venue(testLatitude, testLongitude, testTitle, testAddress)
	foursquareType := g.String("restaurant")
	result := venue.FoursquareType(foursquareType)
	if result == nil {
		t.Error("Expected FoursquareType method to return MessageVenue")
	}
	if result != venue {
		t.Error("Expected FoursquareType to return same MessageVenue instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputVenueMessageContent); ok {
		if v.FoursquareType != foursquareType.Std() {
			t.Error("Expected FoursquareType to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputVenueMessageContent")
	}
}

func TestVenue_GooglePlaceID(t *testing.T) {
	venue := input.Venue(testLatitude, testLongitude, testTitle, testAddress)
	googlePlaceID := g.String("ChIJN1t_tDeuEmsRUsoyG83frY4")
	result := venue.GooglePlaceID(googlePlaceID)
	if result == nil {
		t.Error("Expected GooglePlaceID method to return MessageVenue")
	}
	if result != venue {
		t.Error("Expected GooglePlaceID to return same MessageVenue instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputVenueMessageContent); ok {
		if v.GooglePlaceId != googlePlaceID.Std() {
			t.Error("Expected GooglePlaceId to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputVenueMessageContent")
	}
}

func TestVenue_GooglePlaceType(t *testing.T) {
	venue := input.Venue(testLatitude, testLongitude, testTitle, testAddress)
	googlePlaceType := g.String("restaurant")
	result := venue.GooglePlaceType(googlePlaceType)
	if result == nil {
		t.Error("Expected GooglePlaceType method to return MessageVenue")
	}
	if result != venue {
		t.Error("Expected GooglePlaceType to return same MessageVenue instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputVenueMessageContent); ok {
		if v.GooglePlaceType != googlePlaceType.Std() {
			t.Error("Expected GooglePlaceType to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputVenueMessageContent")
	}
}

func TestVenue_Build(t *testing.T) {
	venue := input.Venue(testLatitude, testLongitude, testTitle, testAddress)
	built := venue.Build()

	if v, ok := built.(gotgbot.InputVenueMessageContent); ok {
		if v.Latitude != testLatitude {
			t.Errorf("Expected Latitude to be %f, got %f", testLatitude, v.Latitude)
		}
		if v.Longitude != testLongitude {
			t.Errorf("Expected Longitude to be %f, got %f", testLongitude, v.Longitude)
		}
		if v.Title != testTitle.Std() {
			t.Error("Expected Title to be set correctly")
		}
		if v.Address != testAddress.Std() {
			t.Error("Expected Address to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputVenueMessageContent")
	}
}

func TestVenue_MethodChaining(t *testing.T) {
	foursquareID := g.String("4bcad4e3f964a520e40c21e3")
	foursquareType := g.String("restaurant")
	googlePlaceID := g.String("ChIJN1t_tDeuEmsRUsoyG83frY4")
	googlePlaceType := g.String("restaurant")

	result := input.Venue(testLatitude, testLongitude, testTitle, testAddress).
		FoursquareID(foursquareID).
		FoursquareType(foursquareType).
		GooglePlaceID(googlePlaceID).
		GooglePlaceType(googlePlaceType)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Venue to build correctly")
	}

	if _, ok := built.(gotgbot.InputVenueMessageContent); !ok {
		t.Error("Expected result to be InputVenueMessageContent")
	}

	if !assertMessageContent(result) {
		t.Error("Expected result to implement MessageContent interface")
	}
}
