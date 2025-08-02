package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestPaidPhoto(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidPhoto := input.PaidPhoto(mediaFile)
	if paidPhoto == nil {
		t.Error("Expected PaidMediaPhoto to be created")
	}
	if !assertPaidMedia(paidPhoto) {
		t.Error("PaidMediaPhoto should implement PaidMedia correctly")
	}
}

func TestPaidPhoto_Build(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidPhoto := input.PaidPhoto(mediaFile)
	built := paidPhoto.Build()

	if built == nil {
		t.Error("Expected Build to return non-nil result")
	}

	if _, ok := built.(gotgbot.InputPaidMediaPhoto); !ok {
		t.Error("Expected result to be InputPaidMediaPhoto")
	}
}

func TestPaidPhoto_BuildReturnsCorrectType(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidPhoto := input.PaidPhoto(mediaFile)
	built := paidPhoto.Build()

	// Verify that Build() returns the correct type
	if _, ok := interface{}(built).(gotgbot.InputPaidMediaPhoto); !ok {
		t.Error("Expected Build() to return gotgbot.InputPaidMediaPhoto")
	}
}

func TestPaidPhoto_MultipleBuildsSameResult(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidPhoto := input.PaidPhoto(mediaFile)

	// Build multiple times to ensure consistency
	built1 := paidPhoto.Build()
	built2 := paidPhoto.Build()

	if built1 == nil || built2 == nil {
		t.Error("Expected builds to return non-nil results")
	}

	// Both should be the same type
	_, ok1 := built1.(gotgbot.InputPaidMediaPhoto)
	_, ok2 := built2.(gotgbot.InputPaidMediaPhoto)

	if !ok1 || !ok2 {
		t.Error("Expected both builds to return InputPaidMediaPhoto")
	}
}
