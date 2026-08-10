package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestVoiceNote(t *testing.T) {
	block := rich.VoiceNote(testVoiceNote())
	if block == nil {
		t.Fatal("Expected VoiceNote to create a BlockVoiceNote builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockVoiceNote)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockVoiceNote")
	}

	if built.VoiceNote.Media == nil {
		t.Error("Expected VoiceNote media to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeVoiceNote {
		t.Error("Expected block type to be voice_note")
	}
}

func TestBlockVoiceNote_VoiceNote(t *testing.T) {
	block := rich.VoiceNote(nil)

	result := block.VoiceNote(testVoiceNote())
	if result != block {
		t.Error("Expected VoiceNote to return same BlockVoiceNote instance")
	}

	if result.VoiceNote(nil) != block {
		t.Error("Expected VoiceNote to return same BlockVoiceNote instance for nil input")
	}

	if result.Build().(gotgbot.InputRichBlockVoiceNote).VoiceNote.Media == nil {
		t.Error("Expected VoiceNote to keep the previous value when passed nil")
	}
}

func TestBlockVoiceNote_Caption(t *testing.T) {
	block := rich.VoiceNote(testVoiceNote())

	result := block.Caption(rich.NewCaption(testText(testCaption)))
	if result != block {
		t.Error("Expected Caption to return same BlockVoiceNote instance")
	}

	built := result.Build().(gotgbot.InputRichBlockVoiceNote)
	if built.Caption == nil {
		t.Fatal("Expected Caption to be set")
	}

	if string(built.Caption.Text.(gotgbot.RichTextString)) != testCaption.Std() {
		t.Error("Expected Caption text to be set correctly")
	}
}
