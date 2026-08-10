package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestParagraph(t *testing.T) {
	block := rich.Paragraph(testText("Hello"))
	if block == nil {
		t.Fatal("Expected Paragraph to create a BlockParagraph builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockParagraph)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockParagraph")
	}

	if string(built.Text.(gotgbot.RichTextString)) != "Hello" {
		t.Error("Expected Text to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeParagraph {
		t.Error("Expected block type to be paragraph")
	}
}

func TestBlockParagraph_Text(t *testing.T) {
	block := rich.Paragraph(testText("old"))

	result := block.Text(testText("new"))
	if result != block {
		t.Error("Expected Text to return same BlockParagraph instance")
	}

	built := result.Build().(gotgbot.InputRichBlockParagraph)
	if string(built.Text.(gotgbot.RichTextString)) != "new" {
		t.Error("Expected Text to be set correctly")
	}
}
