package inline_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewLocation(t *testing.T) {
	latitude := 40.7128
	longitude := -74.0060
	location := inline.NewLocation(testID, latitude, longitude, testTitle)

	if location == nil {
		t.Error("Expected Location to be created")
	}

	built := location.Build()
	if built == nil {
		t.Error("Expected Location to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultLocation); ok {
		if result.GetType() != "location" {
			t.Error("Expected type to be 'location'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultLocation")
	}
}

func TestLocation_HorizontalAccuracy(t *testing.T) {
	location := inline.NewLocation(testID, 40.7128, -74.0060, testTitle)

	result := location.HorizontalAccuracy(50.0)
	if result == nil {
		t.Error("Expected HorizontalAccuracy method to return Location")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultLocation); ok {
		if v.HorizontalAccuracy != 50.0 {
			t.Error("Expected HorizontalAccuracy to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultLocation")
	}
}

func TestLocation_LiveFor(t *testing.T) {
	location := inline.NewLocation(testID, 40.7128, -74.0060, testTitle)

	result := location.LiveFor(600 * time.Second)
	if result == nil {
		t.Error("Expected LiveFor method to return Location")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultLocation); ok {
		if v.LivePeriod != 600 {
			t.Error("Expected LivePeriod to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultLocation")
	}
}

func TestLocation_Heading(t *testing.T) {
	location := inline.NewLocation(testID, 40.7128, -74.0060, testTitle)

	result := location.Heading(90)
	if result == nil {
		t.Error("Expected Heading method to return Location")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultLocation); ok {
		if v.Heading != 90 {
			t.Error("Expected Heading to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultLocation")
	}
}

func TestLocation_ProximityAlertRadius(t *testing.T) {
	location := inline.NewLocation(testID, 40.7128, -74.0060, testTitle)

	result := location.ProximityAlertRadius(100)
	if result == nil {
		t.Error("Expected ProximityAlertRadius method to return Location")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultLocation); ok {
		if v.ProximityAlertRadius != 100 {
			t.Error("Expected ProximityAlertRadius to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultLocation")
	}
}

func TestLocation_ThumbnailURL(t *testing.T) {
	location := inline.NewLocation(testID, 40.7128, -74.0060, testTitle)

	result := location.ThumbnailURL(testThumbnailURL)
	if result == nil {
		t.Error("Expected ThumbnailURL method to return Location")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultLocation); ok {
		if v.ThumbnailUrl != testThumbnailURL.Std() {
			t.Error("Expected ThumbnailUrl to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultLocation")
	}
}

func TestLocation_MethodChaining(t *testing.T) {
	messageContent := createTestMessageContent()

	result := inline.NewLocation(testID, 40.7128, -74.0060, testTitle).
		HorizontalAccuracy(50.0).
		LiveFor(600*time.Second).
		Heading(90).
		ProximityAlertRadius(100).
		ThumbnailURL(testThumbnailURL).
		ThumbnailSize(150, 150).
		Markup(createTestKeyboard()).
		InputMessageContent(messageContent)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Location to build correctly")
	}

	if _, ok := built.(gotgbot.InlineQueryResultLocation); !ok {
		t.Error("Expected result to be InlineQueryResultLocation")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
