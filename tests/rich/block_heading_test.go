package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestHeading(t *testing.T) {
	block := rich.Heading(testText("Title"))
	if block == nil {
		t.Fatal("Expected Heading to create a BlockHeading builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockSectionHeading)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockSectionHeading")
	}

	if string(built.Text.(gotgbot.RichTextString)) != "Title" {
		t.Error("Expected Text to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeHeading {
		t.Error("Expected block type to be heading")
	}
}

func TestBlockHeading_Text(t *testing.T) {
	block := rich.Heading(testText("old"))

	result := block.Text(testText("new"))
	if result != block {
		t.Error("Expected Text to return same BlockHeading instance")
	}

	built := result.Build().(gotgbot.InputRichBlockSectionHeading)
	if string(built.Text.(gotgbot.RichTextString)) != "new" {
		t.Error("Expected Text to be set correctly")
	}
}

func TestBlockHeading_Size(t *testing.T) {
	block := rich.Heading(testText("Title"))

	result := block.Size(2)
	if result != block {
		t.Error("Expected Size to return same BlockHeading instance")
	}

	built := result.Build().(gotgbot.InputRichBlockSectionHeading)
	if built.Size != 2 {
		t.Error("Expected Size to be set correctly")
	}
}
