package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestPaidLivePhoto(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidLivePhoto := input.PaidLivePhoto(mediaFile, testThumbnailURL)
	if paidLivePhoto == nil {
		t.Error("Expected PaidMediaLivePhoto to be created")
	}
	if !assertPaidMedia(paidLivePhoto) {
		t.Error("PaidMediaLivePhoto should implement PaidMedia correctly")
	}
}

func TestPaidLivePhoto_Build(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	built := input.PaidLivePhoto(mediaFile, testThumbnailURL).Build()

	if v, ok := built.(gotgbot.InputPaidMediaLivePhoto); ok {
		if v.Photo != testThumbnailURL.Std() {
			t.Errorf("Expected Photo to be %s, got %s", testThumbnailURL.Std(), v.Photo)
		}
	} else {
		t.Error("Expected result to be InputPaidMediaLivePhoto")
	}
}
