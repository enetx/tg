package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/input"
)

func TestLocationMedia(t *testing.T) {
	location := input.LocationMedia(testLatitude, testLongitude)
	if location == nil {
		t.Error("Expected LocationMedia to be created")
	}
	if !assertPollMedia(location) {
		t.Error("LocationMedia should implement PollMedia correctly")
	}
	if !assertPollOptionMedia(location) {
		t.Error("LocationMedia should implement PollOptionMedia correctly")
	}
}

func TestLocationMedia_BuildPollMedia(t *testing.T) {
	built := input.LocationMedia(testLatitude, testLongitude).BuildPollMedia()
	if v, ok := built.(gotgbot.InputMediaLocation); ok {
		if v.Latitude != testLatitude || v.Longitude != testLongitude {
			t.Error("Expected coordinates to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaLocation")
	}
}

func TestLocationMedia_HorizontalAccuracy(t *testing.T) {
	location := input.LocationMedia(testLatitude, testLongitude)
	result := location.HorizontalAccuracy(50.0)
	if result != location {
		t.Error("Expected HorizontalAccuracy to return same MediaLocation instance")
	}

	built := result.BuildPollOptionMedia()
	if v, ok := built.(gotgbot.InputMediaLocation); ok {
		if v.HorizontalAccuracy != 50.0 {
			t.Error("Expected HorizontalAccuracy to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaLocation")
	}
}
