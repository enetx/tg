package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/rich"
)

func TestAnchor(t *testing.T) {
	block := rich.Anchor(g.String("section-1"))
	if block == nil {
		t.Fatal("Expected Anchor to create a BlockAnchor builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockAnchor)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockAnchor")
	}

	if built.Name != "section-1" {
		t.Error("Expected Name to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeAnchor {
		t.Error("Expected block type to be anchor")
	}
}

func TestBlockAnchor_Name(t *testing.T) {
	block := rich.Anchor(g.String("old"))

	result := block.Name(g.String("new"))
	if result != block {
		t.Error("Expected Name to return same BlockAnchor instance")
	}

	built := result.Build().(gotgbot.InputRichBlockAnchor)
	if built.Name != "new" {
		t.Error("Expected Name to be set correctly")
	}
}
