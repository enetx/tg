package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/rich"
)

func TestPre(t *testing.T) {
	block := rich.Pre(testText("fmt.Println()"))
	if block == nil {
		t.Fatal("Expected Pre to create a BlockPre builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockPreformatted)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockPreformatted")
	}

	if string(built.Text.(gotgbot.RichTextString)) != "fmt.Println()" {
		t.Error("Expected Text to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypePre {
		t.Error("Expected block type to be pre")
	}
}

func TestBlockPre_Text(t *testing.T) {
	block := rich.Pre(testText("old"))

	result := block.Text(testText("new"))
	if result != block {
		t.Error("Expected Text to return same BlockPre instance")
	}

	built := result.Build().(gotgbot.InputRichBlockPreformatted)
	if string(built.Text.(gotgbot.RichTextString)) != "new" {
		t.Error("Expected Text to be set correctly")
	}
}

func TestBlockPre_Language(t *testing.T) {
	block := rich.Pre(testText("fmt.Println()"))

	result := block.Language(g.String("go"))
	if result != block {
		t.Error("Expected Language to return same BlockPre instance")
	}

	built := result.Build().(gotgbot.InputRichBlockPreformatted)
	if built.Language != "go" {
		t.Error("Expected Language to be set correctly")
	}
}
