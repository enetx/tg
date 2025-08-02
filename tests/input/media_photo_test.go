package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestPhoto(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	photo := input.Photo(mediaFile)
	if photo == nil {
		t.Error("Expected Photo to be created")
	}
	if !assertMedia(photo) {
		t.Error("Photo should implement Media correctly")
	}
}

func TestPhoto_Caption(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	photo := input.Photo(mediaFile)
	result := photo.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return MediaPhoto")
	}
	if result != photo {
		t.Error("Expected Caption to return same MediaPhoto instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaPhoto); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaPhoto")
	}
}

func TestPhoto_HTML(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	photo := input.Photo(mediaFile)
	result := photo.HTML()
	if result == nil {
		t.Error("Expected HTML method to return MediaPhoto")
	}
	if result != photo {
		t.Error("Expected HTML to return same MediaPhoto instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaPhoto); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InputMediaPhoto")
	}
}

func TestPhoto_Markdown(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	photo := input.Photo(mediaFile)
	result := photo.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return MediaPhoto")
	}
	if result != photo {
		t.Error("Expected Markdown to return same MediaPhoto instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaPhoto); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InputMediaPhoto")
	}
}

func TestPhoto_ShowCaptionAboveMedia(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	photo := input.Photo(mediaFile)
	result := photo.ShowCaptionAboveMedia()
	if result == nil {
		t.Error("Expected ShowCaptionAboveMedia method to return MediaPhoto")
	}
	if result != photo {
		t.Error("Expected ShowCaptionAboveMedia to return same MediaPhoto instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaPhoto); ok {
		if !v.ShowCaptionAboveMedia {
			t.Error("Expected ShowCaptionAboveMedia to be set to true")
		}
	} else {
		t.Error("Expected result to be InputMediaPhoto")
	}
}

func TestPhoto_Spoiler(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	photo := input.Photo(mediaFile)
	result := photo.Spoiler()
	if result == nil {
		t.Error("Expected Spoiler method to return MediaPhoto")
	}
	if result != photo {
		t.Error("Expected Spoiler to return same MediaPhoto instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaPhoto); ok {
		if !v.HasSpoiler {
			t.Error("Expected HasSpoiler to be set to true")
		}
	} else {
		t.Error("Expected result to be InputMediaPhoto")
	}
}

func TestPhoto_MethodChaining(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	result := input.Photo(mediaFile).
		Caption(testCaption).
		HTML().
		ShowCaptionAboveMedia().
		Spoiler()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Photo to build correctly")
	}

	if _, ok := built.(gotgbot.InputMediaPhoto); !ok {
		t.Error("Expected result to be InputMediaPhoto")
	}

	if !assertMedia(result) {
		t.Error("Expected result to implement Media interface")
	}
}
