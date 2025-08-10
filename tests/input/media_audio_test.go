package input_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestAudio(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	audio := input.Audio(mediaFile)
	if audio == nil {
		t.Error("Expected Audio to be created")
	}
	if !assertMedia(audio) {
		t.Error("Audio should implement Media correctly")
	}
}

func TestAudio_Caption(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	audio := input.Audio(mediaFile)
	result := audio.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return MediaAudio")
	}
	if result != audio {
		t.Error("Expected Caption to return same MediaAudio instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAudio); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaAudio")
	}
}

func TestAudio_HTML(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	audio := input.Audio(mediaFile)
	result := audio.HTML()
	if result == nil {
		t.Error("Expected HTML method to return MediaAudio")
	}
	if result != audio {
		t.Error("Expected HTML to return same MediaAudio instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAudio); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InputMediaAudio")
	}
}

func TestAudio_Thumbnail(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	thumbnailFile := file.Input(testThumbnailURL).Ok()
	audio := input.Audio(mediaFile)

	result := audio.Thumbnail(thumbnailFile)
	if result == nil {
		t.Error("Expected Thumbnail method to return MediaAudio")
	}
	if result != audio {
		t.Error("Expected Thumbnail to return same MediaAudio instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAudio); ok {
		if v.Thumbnail == nil {
			t.Error("Expected Thumbnail to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaAudio")
	}
}

func TestAudio_Markdown(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	audio := input.Audio(mediaFile)

	result := audio.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return MediaAudio")
	}
	if result != audio {
		t.Error("Expected Markdown to return same MediaAudio instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAudio); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InputMediaAudio")
	}
}

func TestAudio_CaptionEntities(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	audio := input.Audio(mediaFile)
	entities := createTestEntities()

	result := audio.CaptionEntities(entities)
	if result == nil {
		t.Error("Expected CaptionEntities method to return MediaAudio")
	}
	if result != audio {
		t.Error("Expected CaptionEntities to return same MediaAudio instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAudio); ok {
		if len(v.CaptionEntities) == 0 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaAudio")
	}
}

func TestAudio_Duration(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	audio := input.Audio(mediaFile)
	duration := 180 * time.Second
	result := audio.Duration(duration)
	if result == nil {
		t.Error("Expected Duration method to return MediaAudio")
	}
	if result != audio {
		t.Error("Expected Duration to return same MediaAudio instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAudio); ok {
		if v.Duration != 180 {
			t.Errorf("Expected Duration to be 180 seconds, got %d", v.Duration)
		}
	} else {
		t.Error("Expected result to be InputMediaAudio")
	}
}

func TestAudio_Performer(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	audio := input.Audio(mediaFile)
	performer := g.String("Test Artist")
	result := audio.Performer(performer)
	if result == nil {
		t.Error("Expected Performer method to return MediaAudio")
	}
	if result != audio {
		t.Error("Expected Performer to return same MediaAudio instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAudio); ok {
		if v.Performer != performer.Std() {
			t.Error("Expected Performer to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaAudio")
	}
}

func TestAudio_Title(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	audio := input.Audio(mediaFile)
	result := audio.Title(testTitle)
	if result == nil {
		t.Error("Expected Title method to return MediaAudio")
	}
	if result != audio {
		t.Error("Expected Title to return same MediaAudio instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaAudio); ok {
		if v.Title != testTitle.Std() {
			t.Error("Expected Title to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaAudio")
	}
}

func TestAudio_MethodChaining(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	performer := g.String("Test Artist")
	result := input.Audio(mediaFile).
		Caption(testCaption).
		HTML().
		Duration(180 * time.Second).
		Performer(performer).
		Title(testTitle)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Audio to build correctly")
	}

	if _, ok := built.(gotgbot.InputMediaAudio); !ok {
		t.Error("Expected result to be InputMediaAudio")
	}

	if !assertMedia(result) {
		t.Error("Expected result to implement Media interface")
	}
}
