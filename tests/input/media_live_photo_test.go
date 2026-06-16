package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestLivePhoto(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	livePhoto := input.LivePhoto(mediaFile, testThumbnailURL)
	if livePhoto == nil {
		t.Error("Expected LivePhoto to be created")
	}
	if !assertMedia(livePhoto) {
		t.Error("LivePhoto should implement Media correctly")
	}
	if !assertPollMedia(livePhoto) {
		t.Error("LivePhoto should implement PollMedia correctly")
	}
	if !assertPollOptionMedia(livePhoto) {
		t.Error("LivePhoto should implement PollOptionMedia correctly")
	}
}

func TestLivePhoto_Build(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	built := input.LivePhoto(mediaFile, testThumbnailURL).Build()

	if v, ok := built.(gotgbot.InputMediaLivePhoto); ok {
		if v.Photo != testThumbnailURL.Std() {
			t.Errorf("Expected Photo to be %s, got %s", testThumbnailURL.Std(), v.Photo)
		}
	} else {
		t.Error("Expected result to be InputMediaLivePhoto")
	}
}

func TestLivePhoto_Caption(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	livePhoto := input.LivePhoto(mediaFile, testThumbnailURL)
	result := livePhoto.Caption(testCaption)
	if result != livePhoto {
		t.Error("Expected Caption to return same MediaLivePhoto instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaLivePhoto); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaLivePhoto")
	}
}

func TestLivePhoto_HTML(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	built := input.LivePhoto(mediaFile, testThumbnailURL).HTML().Build()
	if v, ok := built.(gotgbot.InputMediaLivePhoto); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InputMediaLivePhoto")
	}
}

func TestLivePhoto_Markdown(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	built := input.LivePhoto(mediaFile, testThumbnailURL).Markdown().Build()
	if v, ok := built.(gotgbot.InputMediaLivePhoto); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InputMediaLivePhoto")
	}
}

func TestLivePhoto_CaptionEntities(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	built := input.LivePhoto(mediaFile, testThumbnailURL).CaptionEntities(createTestEntities()).Build()
	if v, ok := built.(gotgbot.InputMediaLivePhoto); ok {
		if len(v.CaptionEntities) == 0 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaLivePhoto")
	}
}

func TestLivePhoto_ShowCaptionAboveMedia(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	built := input.LivePhoto(mediaFile, testThumbnailURL).ShowCaptionAboveMedia().Build()
	if v, ok := built.(gotgbot.InputMediaLivePhoto); ok {
		if !v.ShowCaptionAboveMedia {
			t.Error("Expected ShowCaptionAboveMedia to be set to true")
		}
	} else {
		t.Error("Expected result to be InputMediaLivePhoto")
	}
}

func TestLivePhoto_Spoiler(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	built := input.LivePhoto(mediaFile, testThumbnailURL).Spoiler().Build()
	if v, ok := built.(gotgbot.InputMediaLivePhoto); ok {
		if !v.HasSpoiler {
			t.Error("Expected HasSpoiler to be set to true")
		}
	} else {
		t.Error("Expected result to be InputMediaLivePhoto")
	}
}

func TestLivePhoto_MethodChaining(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	result := input.LivePhoto(mediaFile, testThumbnailURL).
		Caption(testCaption).
		HTML().
		ShowCaptionAboveMedia().
		Spoiler()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	if _, ok := result.Build().(gotgbot.InputMediaLivePhoto); !ok {
		t.Error("Expected result to be InputMediaLivePhoto")
	}
}
