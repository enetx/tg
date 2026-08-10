package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestAnimation(t *testing.T) {
	block := rich.Animation(testAnimation())
	if block == nil {
		t.Fatal("Expected Animation to create a BlockAnimation builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockAnimation)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockAnimation")
	}

	if built.Animation.Media == nil {
		t.Error("Expected Animation media to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeAnimation {
		t.Error("Expected block type to be animation")
	}
}

func TestBlockAnimation_Animation(t *testing.T) {
	block := rich.Animation(nil)

	result := block.Animation(testAnimation())
	if result != block {
		t.Error("Expected Animation to return same BlockAnimation instance")
	}

	if result.Animation(nil) != block {
		t.Error("Expected Animation to return same BlockAnimation instance for nil input")
	}

	if result.Build().(gotgbot.InputRichBlockAnimation).Animation.Media == nil {
		t.Error("Expected Animation to keep the previous value when passed nil")
	}
}

func TestBlockAnimation_Caption(t *testing.T) {
	block := rich.Animation(testAnimation())

	result := block.Caption(rich.NewCaption(testText(testCaption)))
	if result != block {
		t.Error("Expected Caption to return same BlockAnimation instance")
	}

	built := result.Build().(gotgbot.InputRichBlockAnimation)
	if built.Caption == nil {
		t.Fatal("Expected Caption to be set")
	}

	if string(built.Caption.Text.(gotgbot.RichTextString)) != testCaption.Std() {
		t.Error("Expected Caption text to be set correctly")
	}
}
