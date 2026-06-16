package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestVenueMedia(t *testing.T) {
	venue := input.VenueMedia(testLatitude, testLongitude, testTitle, testAddress)
	if venue == nil {
		t.Error("Expected VenueMedia to be created")
	}
	if !assertPollMedia(venue) {
		t.Error("VenueMedia should implement PollMedia correctly")
	}
	if !assertPollOptionMedia(venue) {
		t.Error("VenueMedia should implement PollOptionMedia correctly")
	}
}

func TestVenueMedia_BuildPollMedia(t *testing.T) {
	built := input.VenueMedia(testLatitude, testLongitude, testTitle, testAddress).BuildPollMedia()
	if v, ok := built.(gotgbot.InputMediaVenue); ok {
		if v.Title != testTitle.Std() || v.Address != testAddress.Std() {
			t.Error("Expected title and address to be set correctly")
		}
		if v.Latitude != testLatitude || v.Longitude != testLongitude {
			t.Error("Expected coordinates to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaVenue")
	}
}

func TestVenueMedia_Foursquare(t *testing.T) {
	venue := input.VenueMedia(testLatitude, testLongitude, testTitle, testAddress)
	result := venue.Foursquare(g.String("fsq-id"), g.String("fsq-type"))
	if result != venue {
		t.Error("Expected Foursquare to return same MediaVenue instance")
	}

	built := result.BuildPollMedia()
	if v, ok := built.(gotgbot.InputMediaVenue); ok {
		if v.FoursquareId != "fsq-id" || v.FoursquareType != "fsq-type" {
			t.Error("Expected Foursquare fields to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaVenue")
	}
}

func TestVenueMedia_GooglePlace(t *testing.T) {
	venue := input.VenueMedia(testLatitude, testLongitude, testTitle, testAddress)
	result := venue.GooglePlace(g.String("gp-id"), g.String("gp-type"))
	if result != venue {
		t.Error("Expected GooglePlace to return same MediaVenue instance")
	}

	built := result.BuildPollOptionMedia()
	if v, ok := built.(gotgbot.InputMediaVenue); ok {
		if v.GooglePlaceId != "gp-id" || v.GooglePlaceType != "gp-type" {
			t.Error("Expected GooglePlace fields to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaVenue")
	}
}
