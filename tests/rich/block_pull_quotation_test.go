package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestPullQuotation(t *testing.T) {
	block := rich.PullQuotation(testText("Quoted"))
	if block == nil {
		t.Fatal("Expected PullQuotation to create a BlockPullQuotation builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockPullQuotation)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockPullQuotation")
	}

	if string(built.Text.(gotgbot.RichTextString)) != "Quoted" {
		t.Error("Expected Text to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypePullquote {
		t.Error("Expected block type to be pullquote")
	}
}

func TestBlockPullQuotation_Text(t *testing.T) {
	block := rich.PullQuotation(testText("old"))

	result := block.Text(testText("new"))
	if result != block {
		t.Error("Expected Text to return same BlockPullQuotation instance")
	}

	built := result.Build().(gotgbot.InputRichBlockPullQuotation)
	if string(built.Text.(gotgbot.RichTextString)) != "new" {
		t.Error("Expected Text to be set correctly")
	}
}

func TestBlockPullQuotation_Credit(t *testing.T) {
	block := rich.PullQuotation(testText("Quoted"))

	result := block.Credit(testText(testCredit))
	if result != block {
		t.Error("Expected Credit to return same BlockPullQuotation instance")
	}

	built := result.Build().(gotgbot.InputRichBlockPullQuotation)
	if string(built.Credit.(gotgbot.RichTextString)) != testCredit.Std() {
		t.Error("Expected Credit to be set correctly")
	}
}
