package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestSlideshow(t *testing.T) {
	block := rich.Slideshow(rich.Photo(testPhoto()), nil)
	if block == nil {
		t.Fatal("Expected Slideshow to create a BlockSlideshow builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockSlideshow)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockSlideshow")
	}

	if len(built.Blocks) != 1 {
		t.Error("Expected Slideshow to skip nil blocks")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeSlideshow {
		t.Error("Expected block type to be slideshow")
	}
}

func TestBlockSlideshow_Block(t *testing.T) {
	block := rich.Slideshow()

	result := block.Block(rich.Photo(testPhoto()), rich.Video(testVideo()))
	if result != block {
		t.Error("Expected Block to return same BlockSlideshow instance")
	}

	if len(result.Build().(gotgbot.InputRichBlockSlideshow).Blocks) != 2 {
		t.Error("Expected Block to append blocks")
	}
}

func TestBlockSlideshow_Caption(t *testing.T) {
	block := rich.Slideshow()

	result := block.Caption(rich.NewCaption(testText(testCaption)))
	if result != block {
		t.Error("Expected Caption to return same BlockSlideshow instance")
	}

	built := result.Build().(gotgbot.InputRichBlockSlideshow)
	if built.Caption == nil {
		t.Fatal("Expected Caption to be set")
	}

	if string(built.Caption.Text.(gotgbot.RichTextString)) != testCaption.Std() {
		t.Error("Expected Caption text to be set correctly")
	}
}

func TestBlockSlideshow_Empty(t *testing.T) {
	if rich.Slideshow().Build().(gotgbot.InputRichBlockSlideshow).Blocks != nil {
		t.Error("Expected an empty slideshow to carry no blocks")
	}
}
