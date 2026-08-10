package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestAudio(t *testing.T) {
	block := rich.Audio(testAudio())
	if block == nil {
		t.Fatal("Expected Audio to create a BlockAudio builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockAudio)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockAudio")
	}

	if built.Audio.Media == nil {
		t.Error("Expected Audio media to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeAudio {
		t.Error("Expected block type to be audio")
	}
}

func TestBlockAudio_Audio(t *testing.T) {
	block := rich.Audio(nil)

	result := block.Audio(testAudio())
	if result != block {
		t.Error("Expected Audio to return same BlockAudio instance")
	}

	if result.Audio(nil) != block {
		t.Error("Expected Audio to return same BlockAudio instance for nil input")
	}

	if result.Build().(gotgbot.InputRichBlockAudio).Audio.Media == nil {
		t.Error("Expected Audio to keep the previous value when passed nil")
	}
}

func TestBlockAudio_Caption(t *testing.T) {
	block := rich.Audio(testAudio())

	result := block.Caption(rich.NewCaption(testText(testCaption)))
	if result != block {
		t.Error("Expected Caption to return same BlockAudio instance")
	}

	built := result.Build().(gotgbot.InputRichBlockAudio)
	if built.Caption == nil {
		t.Fatal("Expected Caption to be set")
	}

	if string(built.Caption.Text.(gotgbot.RichTextString)) != testCaption.Std() {
		t.Error("Expected Caption text to be set correctly")
	}
}
