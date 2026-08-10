package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestThinking(t *testing.T) {
	block := rich.Thinking(testText("Thinking..."))
	if block == nil {
		t.Fatal("Expected Thinking to create a BlockThinking builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockThinking)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockThinking")
	}

	if string(built.Text.(gotgbot.RichTextString)) != "Thinking..." {
		t.Error("Expected Text to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeThinking {
		t.Error("Expected block type to be thinking")
	}
}

func TestBlockThinking_Text(t *testing.T) {
	block := rich.Thinking(testText("old"))

	result := block.Text(testText("new"))
	if result != block {
		t.Error("Expected Text to return same BlockThinking instance")
	}

	built := result.Build().(gotgbot.InputRichBlockThinking)
	if string(built.Text.(gotgbot.RichTextString)) != "new" {
		t.Error("Expected Text to be set correctly")
	}
}
