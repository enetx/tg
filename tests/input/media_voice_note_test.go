package input_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestVoiceNote(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	voiceNote := input.VoiceNote(mediaFile)
	if voiceNote == nil {
		t.Error("Expected VoiceNote to be created")
	}

	built := voiceNote.Build()
	if built.Media == nil {
		t.Error("Expected Media to be set correctly")
	}
}

func TestVoiceNote_Caption(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	voiceNote := input.VoiceNote(mediaFile)

	result := voiceNote.Caption(testCaption)
	if result != voiceNote {
		t.Error("Expected Caption to return same MediaVoiceNote instance")
	}

	if result.Build().Caption != testCaption.Std() {
		t.Error("Expected Caption to be set correctly")
	}
}

func TestVoiceNote_ParseModes(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	voiceNote := input.VoiceNote(mediaFile)

	if voiceNote.HTML().Build().ParseMode != gotgbot.ParseModeHTML {
		t.Error("Expected HTML parse mode to be set correctly")
	}

	if voiceNote.Markdown().Build().ParseMode != gotgbot.ParseModeMarkdownV2 {
		t.Error("Expected Markdown parse mode to be set correctly")
	}
}

func TestVoiceNote_CaptionEntities(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	voiceNote := input.VoiceNote(mediaFile)

	result := voiceNote.CaptionEntities(createTestEntities())
	if result != voiceNote {
		t.Error("Expected CaptionEntities to return same MediaVoiceNote instance")
	}

	if len(result.Build().CaptionEntities) == 0 {
		t.Error("Expected CaptionEntities to be set correctly")
	}
}

func TestVoiceNote_Duration(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	voiceNote := input.VoiceNote(mediaFile)

	result := voiceNote.Duration(90 * time.Second)
	if result != voiceNote {
		t.Error("Expected Duration to return same MediaVoiceNote instance")
	}

	if result.Build().Duration != 90 {
		t.Error("Expected Duration to be set correctly")
	}
}

func TestVoiceNote_Type(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	if input.VoiceNote(mediaFile).Build().GetType() != "voice_note" {
		t.Error("Expected media type to be voice_note")
	}
}
