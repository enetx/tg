package input_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestPaidVideo(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidVideo := input.PaidVideo(mediaFile)
	if paidVideo == nil {
		t.Error("Expected PaidMediaVideo to be created")
	}
	if !assertPaidMedia(paidVideo) {
		t.Error("PaidMediaVideo should implement PaidMedia correctly")
	}
}

func TestPaidVideo_Cover(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidVideo := input.PaidVideo(mediaFile)
	cover := g.String("https://example.com/cover.jpg")
	result := paidVideo.Cover(cover)
	if result == nil {
		t.Error("Expected Cover method to return PaidMediaVideo")
	}
	if result != paidVideo {
		t.Error("Expected Cover to return same PaidMediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputPaidMediaVideo); ok {
		if v.Cover != cover.Std() {
			t.Error("Expected Cover to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputPaidMediaVideo")
	}
}

func TestPaidVideo_Thumbnail(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	thumbnailFile := file.Input(testThumbnailURL).Ok()
	paidVideo := input.PaidVideo(mediaFile)

	result := paidVideo.Thumbnail(thumbnailFile)
	if result == nil {
		t.Error("Expected Thumbnail method to return PaidMediaVideo")
	}
	if result != paidVideo {
		t.Error("Expected Thumbnail to return same PaidMediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputPaidMediaVideo); ok {
		if v.Thumbnail == nil {
			t.Error("Expected Thumbnail to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputPaidMediaVideo")
	}
}

func TestPaidVideo_Width(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidVideo := input.PaidVideo(mediaFile)
	width := int64(1920)
	result := paidVideo.Width(width)
	if result == nil {
		t.Error("Expected Width method to return PaidMediaVideo")
	}
	if result != paidVideo {
		t.Error("Expected Width to return same PaidMediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputPaidMediaVideo); ok {
		if v.Width != width {
			t.Errorf("Expected Width to be %d, got %d", width, v.Width)
		}
	} else {
		t.Error("Expected result to be InputPaidMediaVideo")
	}
}

func TestPaidVideo_Height(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidVideo := input.PaidVideo(mediaFile)
	height := int64(1080)
	result := paidVideo.Height(height)
	if result == nil {
		t.Error("Expected Height method to return PaidMediaVideo")
	}
	if result != paidVideo {
		t.Error("Expected Height to return same PaidMediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputPaidMediaVideo); ok {
		if v.Height != height {
			t.Errorf("Expected Height to be %d, got %d", height, v.Height)
		}
	} else {
		t.Error("Expected result to be InputPaidMediaVideo")
	}
}

func TestPaidVideo_Duration(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidVideo := input.PaidVideo(mediaFile)
	duration := 120 * time.Second
	result := paidVideo.Duration(duration)
	if result == nil {
		t.Error("Expected Duration method to return PaidMediaVideo")
	}
	if result != paidVideo {
		t.Error("Expected Duration to return same PaidMediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputPaidMediaVideo); ok {
		if v.Duration != 120 {
			t.Errorf("Expected Duration to be 120 seconds, got %d", v.Duration)
		}
	} else {
		t.Error("Expected result to be InputPaidMediaVideo")
	}
}

func TestPaidVideo_StartAt(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidVideo := input.PaidVideo(mediaFile)
	offset := 30 * time.Second
	result := paidVideo.StartAt(offset)
	if result == nil {
		t.Error("Expected StartAt method to return PaidMediaVideo")
	}
	if result != paidVideo {
		t.Error("Expected StartAt to return same PaidMediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputPaidMediaVideo); ok {
		if v.StartTimestamp != 30 {
			t.Errorf("Expected StartTimestamp to be 30 seconds, got %d", v.StartTimestamp)
		}
	} else {
		t.Error("Expected result to be InputPaidMediaVideo")
	}
}

func TestPaidVideo_Streamable(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	paidVideo := input.PaidVideo(mediaFile)
	result := paidVideo.Streamable()
	if result == nil {
		t.Error("Expected Streamable method to return PaidMediaVideo")
	}
	if result != paidVideo {
		t.Error("Expected Streamable to return same PaidMediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputPaidMediaVideo); ok {
		if !v.SupportsStreaming {
			t.Error("Expected SupportsStreaming to be set to true")
		}
	} else {
		t.Error("Expected result to be InputPaidMediaVideo")
	}
}

func TestPaidVideo_MethodChaining(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	cover := g.String("https://example.com/cover.jpg")
	result := input.PaidVideo(mediaFile).
		Cover(cover).
		Width(1920).
		Height(1080).
		Duration(120 * time.Second).
		StartAt(30 * time.Second).
		Streamable()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained PaidVideo to build correctly")
	}

	if _, ok := built.(gotgbot.InputPaidMediaVideo); !ok {
		t.Error("Expected result to be InputPaidMediaVideo")
	}

	if !assertPaidMedia(result) {
		t.Error("Expected result to implement PaidMedia interface")
	}
}
