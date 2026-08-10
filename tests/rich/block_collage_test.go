package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestCollage(t *testing.T) {
	block := rich.Collage(rich.Photo(testPhoto()), nil)
	if block == nil {
		t.Fatal("Expected Collage to create a BlockCollage builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockCollage)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockCollage")
	}

	if len(built.Blocks) != 1 {
		t.Error("Expected Collage to skip nil blocks")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeCollage {
		t.Error("Expected block type to be collage")
	}
}

func TestBlockCollage_Block(t *testing.T) {
	block := rich.Collage()

	result := block.Block(rich.Photo(testPhoto()), rich.Video(testVideo()))
	if result != block {
		t.Error("Expected Block to return same BlockCollage instance")
	}

	if len(result.Build().(gotgbot.InputRichBlockCollage).Blocks) != 2 {
		t.Error("Expected Block to append blocks")
	}
}

func TestBlockCollage_Caption(t *testing.T) {
	block := rich.Collage()

	result := block.Caption(rich.NewCaption(testText(testCaption)))
	if result != block {
		t.Error("Expected Caption to return same BlockCollage instance")
	}

	built := result.Build().(gotgbot.InputRichBlockCollage)
	if built.Caption == nil {
		t.Fatal("Expected Caption to be set")
	}

	if string(built.Caption.Text.(gotgbot.RichTextString)) != testCaption.Std() {
		t.Error("Expected Caption text to be set correctly")
	}
}

func TestBlockCollage_Empty(t *testing.T) {
	if rich.Collage().Build().(gotgbot.InputRichBlockCollage).Blocks != nil {
		t.Error("Expected an empty collage to carry no blocks")
	}
}
