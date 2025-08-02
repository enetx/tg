package input_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestVideo(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	video := input.Video(mediaFile)
	if video == nil {
		t.Error("Expected Video to be created")
	}
	if !assertMedia(video) {
		t.Error("Video should implement Media correctly")
	}
}

func TestVideo_Caption(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	video := input.Video(mediaFile)
	result := video.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return MediaVideo")
	}
	if result != video {
		t.Error("Expected Caption to return same MediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaVideo); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaVideo")
	}
}

func TestVideo_Size(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	video := input.Video(mediaFile)
	result := video.Size(1920, 1080)
	if result == nil {
		t.Error("Expected Size method to return MediaVideo")
	}
	if result != video {
		t.Error("Expected Size to return same MediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaVideo); ok {
		if v.Width != 1920 || v.Height != 1080 {
			t.Error("Expected Width and Height to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaVideo")
	}
}

func TestVideo_Duration(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	video := input.Video(mediaFile)
	duration := 120 * time.Second
	result := video.Duration(duration)
	if result == nil {
		t.Error("Expected Duration method to return MediaVideo")
	}
	if result != video {
		t.Error("Expected Duration to return same MediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaVideo); ok {
		if v.Duration != 120 {
			t.Errorf("Expected Duration to be 120 seconds, got %d", v.Duration)
		}
	} else {
		t.Error("Expected result to be InputMediaVideo")
	}
}

func TestVideo_StartAt(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	video := input.Video(mediaFile)
	offset := 30 * time.Second
	result := video.StartAt(offset)
	if result == nil {
		t.Error("Expected StartAt method to return MediaVideo")
	}
	if result != video {
		t.Error("Expected StartAt to return same MediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaVideo); ok {
		if v.StartTimestamp != 30 {
			t.Errorf("Expected StartTimestamp to be 30 seconds, got %d", v.StartTimestamp)
		}
	} else {
		t.Error("Expected result to be InputMediaVideo")
	}
}

func TestVideo_Streamable(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	video := input.Video(mediaFile)
	result := video.Streamable()
	if result == nil {
		t.Error("Expected Streamable method to return MediaVideo")
	}
	if result != video {
		t.Error("Expected Streamable to return same MediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaVideo); ok {
		if !v.SupportsStreaming {
			t.Error("Expected SupportsStreaming to be set to true")
		}
	} else {
		t.Error("Expected result to be InputMediaVideo")
	}
}

func TestVideo_Cover(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	video := input.Video(mediaFile)
	coverURL := g.String("https://example.com/cover.jpg")
	result := video.Cover(coverURL)
	if result == nil {
		t.Error("Expected Cover method to return MediaVideo")
	}
	if result != video {
		t.Error("Expected Cover to return same MediaVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaVideo); ok {
		if v.Cover != coverURL.Std() {
			t.Error("Expected Cover to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaVideo")
	}
}

func TestVideo_MethodChaining(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	coverURL := g.String("https://example.com/cover.jpg")
	result := input.Video(mediaFile).
		Caption(testCaption).
		HTML().
		Size(1920, 1080).
		Duration(120 * time.Second).
		StartAt(30 * time.Second).
		Cover(coverURL).
		Streamable().
		Spoiler()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Video to build correctly")
	}

	if _, ok := built.(gotgbot.InputMediaVideo); !ok {
		t.Error("Expected result to be InputMediaVideo")
	}

	if !assertMedia(result) {
		t.Error("Expected result to implement Media interface")
	}
}
