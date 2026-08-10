package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestFooter(t *testing.T) {
	block := rich.Footer(testText("Footer"))
	if block == nil {
		t.Fatal("Expected Footer to create a BlockFooter builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockFooter)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockFooter")
	}

	if string(built.Text.(gotgbot.RichTextString)) != "Footer" {
		t.Error("Expected Text to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeFooter {
		t.Error("Expected block type to be footer")
	}
}

func TestBlockFooter_Text(t *testing.T) {
	block := rich.Footer(testText("old"))

	result := block.Text(testText("new"))
	if result != block {
		t.Error("Expected Text to return same BlockFooter instance")
	}

	built := result.Build().(gotgbot.InputRichBlockFooter)
	if string(built.Text.(gotgbot.RichTextString)) != "new" {
		t.Error("Expected Text to be set correctly")
	}
}
