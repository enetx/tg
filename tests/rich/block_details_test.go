package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestDetails(t *testing.T) {
	block := rich.Details(testText("Summary"))
	if block == nil {
		t.Fatal("Expected Details to create a BlockDetails builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockDetails)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockDetails")
	}

	if string(built.Summary.(gotgbot.RichTextString)) != "Summary" {
		t.Error("Expected Summary to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeDetails {
		t.Error("Expected block type to be details")
	}
}

func TestBlockDetails_Summary(t *testing.T) {
	block := rich.Details(testText("old"))

	result := block.Summary(testText("new"))
	if result != block {
		t.Error("Expected Summary to return same BlockDetails instance")
	}

	built := result.Build().(gotgbot.InputRichBlockDetails)
	if string(built.Summary.(gotgbot.RichTextString)) != "new" {
		t.Error("Expected Summary to be set correctly")
	}
}

func TestBlockDetails_Block(t *testing.T) {
	block := rich.Details(testText("Summary"))

	result := block.Block(rich.Paragraph(testText("a")), nil)
	if result != block {
		t.Error("Expected Block to return same BlockDetails instance")
	}

	built := result.Build().(gotgbot.InputRichBlockDetails)
	if len(built.Blocks) != 1 {
		t.Error("Expected Block to skip nil blocks")
	}
}

func TestBlockDetails_Open(t *testing.T) {
	block := rich.Details(testText("Summary"))

	result := block.Open()
	if result != block {
		t.Error("Expected Open to return same BlockDetails instance")
	}

	built := result.Build().(gotgbot.InputRichBlockDetails)
	if !built.IsOpen {
		t.Error("Expected IsOpen to be set correctly")
	}
}
