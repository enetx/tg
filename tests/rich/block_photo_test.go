package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestPhoto(t *testing.T) {
	block := rich.Photo(testPhoto())
	if block == nil {
		t.Fatal("Expected Photo to create a BlockPhoto builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockPhoto)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockPhoto")
	}

	if built.Photo.Media == nil {
		t.Error("Expected Photo media to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypePhoto {
		t.Error("Expected block type to be photo")
	}
}

func TestBlockPhoto_Photo(t *testing.T) {
	block := rich.Photo(nil)

	result := block.Photo(testPhoto())
	if result != block {
		t.Error("Expected Photo to return same BlockPhoto instance")
	}

	if result.Photo(nil) != block {
		t.Error("Expected Photo to return same BlockPhoto instance for nil input")
	}

	if result.Build().(gotgbot.InputRichBlockPhoto).Photo.Media == nil {
		t.Error("Expected Photo to keep the previous value when passed nil")
	}
}

func TestBlockPhoto_Caption(t *testing.T) {
	block := rich.Photo(testPhoto())

	result := block.Caption(rich.NewCaption(testText(testCaption)).Credit(testText(testCredit)))
	if result != block {
		t.Error("Expected Caption to return same BlockPhoto instance")
	}

	built := result.Build().(gotgbot.InputRichBlockPhoto)
	if built.Caption == nil {
		t.Fatal("Expected Caption to be set")
	}

	if string(built.Caption.Credit.(gotgbot.RichTextString)) != testCredit.Std() {
		t.Error("Expected Caption credit to be set correctly")
	}
}
