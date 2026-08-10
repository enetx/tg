package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestDivider(t *testing.T) {
	block := rich.Divider()
	if block == nil {
		t.Fatal("Expected Divider to create a BlockDivider builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockDivider)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockDivider")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeDivider {
		t.Error("Expected block type to be divider")
	}
}
