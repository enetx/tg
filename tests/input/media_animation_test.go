package input_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestAnimation(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	animation := input.Animation(mediaFile)
	if animation == nil {
		t.Error("Expected Animation to be created")
	}
	if !assertMedia(animation) {
		t.Error("Animation should implement Media correctly")
	}
}

func TestAnimation_Caption(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	animation := input.Animation(mediaFile)
	result := animation.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return MediaAnimation")
	}
	if result != animation {
		t.Error("Expected Caption to return same MediaAnimation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAnimation); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaAnimation")
	}
}

func TestAnimation_HTML(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	animation := input.Animation(mediaFile)
	result := animation.HTML()
	if result == nil {
		t.Error("Expected HTML method to return MediaAnimation")
	}
	if result != animation {
		t.Error("Expected HTML to return same MediaAnimation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAnimation); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InputMediaAnimation")
	}
}

func TestAnimation_Markdown(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	animation := input.Animation(mediaFile)
	result := animation.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return MediaAnimation")
	}
	if result != animation {
		t.Error("Expected Markdown to return same MediaAnimation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAnimation); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InputMediaAnimation")
	}
}

func TestAnimation_Thumbnail(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	thumbnailFile := file.Input(testThumbnailURL).Ok()
	animation := input.Animation(mediaFile)

	result := animation.Thumbnail(thumbnailFile)
	if result == nil {
		t.Error("Expected Thumbnail method to return MediaAnimation")
	}
	if result != animation {
		t.Error("Expected Thumbnail to return same MediaAnimation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAnimation); ok {
		if v.Thumbnail == nil {
			t.Error("Expected Thumbnail to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaAnimation")
	}
}

func TestAnimation_CaptionEntities(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	animation := input.Animation(mediaFile)
	entities := createTestEntities()

	result := animation.CaptionEntities(entities)
	if result == nil {
		t.Error("Expected CaptionEntities method to return MediaAnimation")
	}
	if result != animation {
		t.Error("Expected CaptionEntities to return same MediaAnimation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAnimation); ok {
		if len(v.CaptionEntities) == 0 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaAnimation")
	}
}

func TestAnimation_ShowCaptionAboveMedia(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	animation := input.Animation(mediaFile)
	result := animation.ShowCaptionAboveMedia()
	if result == nil {
		t.Error("Expected ShowCaptionAboveMedia method to return MediaAnimation")
	}
	if result != animation {
		t.Error("Expected ShowCaptionAboveMedia to return same MediaAnimation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAnimation); ok {
		if !v.ShowCaptionAboveMedia {
			t.Error("Expected ShowCaptionAboveMedia to be set to true")
		}
	} else {
		t.Error("Expected result to be InputMediaAnimation")
	}
}

func TestAnimation_Size(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	animation := input.Animation(mediaFile)
	result := animation.Size(800, 600)
	if result == nil {
		t.Error("Expected Size method to return MediaAnimation")
	}
	if result != animation {
		t.Error("Expected Size to return same MediaAnimation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAnimation); ok {
		if v.Width != 800 || v.Height != 600 {
			t.Error("Expected Width and Height to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaAnimation")
	}
}

func TestAnimation_Duration(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	animation := input.Animation(mediaFile)
	duration := 30 * time.Second
	result := animation.Duration(duration)
	if result == nil {
		t.Error("Expected Duration method to return MediaAnimation")
	}
	if result != animation {
		t.Error("Expected Duration to return same MediaAnimation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAnimation); ok {
		if v.Duration != 30 {
			t.Errorf("Expected Duration to be 30 seconds, got %d", v.Duration)
		}
	} else {
		t.Error("Expected result to be InputMediaAnimation")
	}
}

func TestAnimation_Spoiler(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	animation := input.Animation(mediaFile)
	result := animation.Spoiler()
	if result == nil {
		t.Error("Expected Spoiler method to return MediaAnimation")
	}
	if result != animation {
		t.Error("Expected Spoiler to return same MediaAnimation instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAnimation); ok {
		if !v.HasSpoiler {
			t.Error("Expected HasSpoiler to be set to true")
		}
	} else {
		t.Error("Expected result to be InputMediaAnimation")
	}
}

func TestAnimation_MethodChaining(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	result := input.Animation(mediaFile).
		Caption(testCaption).
		HTML().
		Size(800, 600).
		Duration(30 * time.Second).
		ShowCaptionAboveMedia().
		Spoiler()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Animation to build correctly")
	}

	if _, ok := built.(gotgbot.InputMediaAnimation); !ok {
		t.Error("Expected result to be InputMediaAnimation")
	}

	if !assertMedia(result) {
		t.Error("Expected result to implement Media interface")
	}
}
