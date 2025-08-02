package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/input"
)

func TestLocation(t *testing.T) {
	location := input.Location(testLatitude, testLongitude)
	if location == nil {
		t.Error("Expected MessageLocation to be created")
	}
	if !assertMessageContent(location) {
		t.Error("MessageLocation should implement MessageContent correctly")
	}
}

func TestLocation_HorizontalAccuracy(t *testing.T) {
	location := input.Location(testLatitude, testLongitude)
	accuracy := 10.5
	result := location.HorizontalAccuracy(accuracy)
	if result == nil {
		t.Error("Expected HorizontalAccuracy method to return MessageLocation")
	}
	if result != location {
		t.Error("Expected HorizontalAccuracy to return same MessageLocation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputLocationMessageContent); ok {
		if v.HorizontalAccuracy != accuracy {
			t.Errorf("Expected HorizontalAccuracy to be %f, got %f", accuracy, v.HorizontalAccuracy)
		}
	} else {
		t.Error("Expected result to be InputLocationMessageContent")
	}
}

func TestLocation_LivePeriod(t *testing.T) {
	location := input.Location(testLatitude, testLongitude)
	period := int64(300)
	result := location.LivePeriod(period)
	if result == nil {
		t.Error("Expected LivePeriod method to return MessageLocation")
	}
	if result != location {
		t.Error("Expected LivePeriod to return same MessageLocation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputLocationMessageContent); ok {
		if v.LivePeriod != period {
			t.Errorf("Expected LivePeriod to be %d, got %d", period, v.LivePeriod)
		}
	} else {
		t.Error("Expected result to be InputLocationMessageContent")
	}
}

func TestLocation_Heading(t *testing.T) {
	location := input.Location(testLatitude, testLongitude)
	heading := int64(90)
	result := location.Heading(heading)
	if result == nil {
		t.Error("Expected Heading method to return MessageLocation")
	}
	if result != location {
		t.Error("Expected Heading to return same MessageLocation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputLocationMessageContent); ok {
		if v.Heading != heading {
			t.Errorf("Expected Heading to be %d, got %d", heading, v.Heading)
		}
	} else {
		t.Error("Expected result to be InputLocationMessageContent")
	}
}

func TestLocation_ProximityAlertRadius(t *testing.T) {
	location := input.Location(testLatitude, testLongitude)
	radius := int64(100)
	result := location.ProximityAlertRadius(radius)
	if result == nil {
		t.Error("Expected ProximityAlertRadius method to return MessageLocation")
	}
	if result != location {
		t.Error("Expected ProximityAlertRadius to return same MessageLocation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputLocationMessageContent); ok {
		if v.ProximityAlertRadius != radius {
			t.Errorf("Expected ProximityAlertRadius to be %d, got %d", radius, v.ProximityAlertRadius)
		}
	} else {
		t.Error("Expected result to be InputLocationMessageContent")
	}
}

func TestLocation_Build(t *testing.T) {
	location := input.Location(testLatitude, testLongitude)
	built := location.Build()

	if v, ok := built.(gotgbot.InputLocationMessageContent); ok {
		if v.Latitude != testLatitude {
			t.Errorf("Expected Latitude to be %f, got %f", testLatitude, v.Latitude)
		}
		if v.Longitude != testLongitude {
			t.Errorf("Expected Longitude to be %f, got %f", testLongitude, v.Longitude)
		}
	} else {
		t.Error("Expected result to be InputLocationMessageContent")
	}
}

func TestLocation_MethodChaining(t *testing.T) {
	result := input.Location(testLatitude, testLongitude).
		HorizontalAccuracy(10.5).
		LivePeriod(300).
		Heading(90).
		ProximityAlertRadius(100)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Location to build correctly")
	}

	if _, ok := built.(gotgbot.InputLocationMessageContent); !ok {
		t.Error("Expected result to be InputLocationMessageContent")
	}

	if !assertMessageContent(result) {
		t.Error("Expected result to implement MessageContent interface")
	}
}
