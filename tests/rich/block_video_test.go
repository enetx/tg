package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestVideo(t *testing.T) {
	block := rich.Video(testVideo())
	if block == nil {
		t.Fatal("Expected Video to create a BlockVideo builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockVideo)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockVideo")
	}

	if built.Video.Media == nil {
		t.Error("Expected Video media to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeVideo {
		t.Error("Expected block type to be video")
	}
}

func TestBlockVideo_Video(t *testing.T) {
	block := rich.Video(nil)

	result := block.Video(testVideo())
	if result != block {
		t.Error("Expected Video to return same BlockVideo instance")
	}

	if result.Video(nil) != block {
		t.Error("Expected Video to return same BlockVideo instance for nil input")
	}

	if result.Build().(gotgbot.InputRichBlockVideo).Video.Media == nil {
		t.Error("Expected Video to keep the previous value when passed nil")
	}
}

func TestBlockVideo_Caption(t *testing.T) {
	block := rich.Video(testVideo())

	result := block.Caption(rich.NewCaption(testText(testCaption)))
	if result != block {
		t.Error("Expected Caption to return same BlockVideo instance")
	}

	built := result.Build().(gotgbot.InputRichBlockVideo)
	if built.Caption == nil {
		t.Fatal("Expected Caption to be set")
	}

	if string(built.Caption.Text.(gotgbot.RichTextString)) != testCaption.Std() {
		t.Error("Expected Caption text to be set correctly")
	}
}
