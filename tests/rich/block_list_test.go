package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestList(t *testing.T) {
	block := rich.List(rich.ListItem(rich.Paragraph(testText("Item"))), nil)
	if block == nil {
		t.Fatal("Expected List to create a BlockList builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockList)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockList")
	}

	if len(built.Items) != 1 {
		t.Error("Expected List to skip nil items")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeList {
		t.Error("Expected block type to be list")
	}
}

func TestBlockList_Item(t *testing.T) {
	block := rich.List()

	result := block.Item(rich.ListItem(), rich.ListItem())
	if result != block {
		t.Error("Expected Item to return same BlockList instance")
	}

	built := result.Build().(gotgbot.InputRichBlockList)
	if len(built.Items) != 2 {
		t.Error("Expected Item to append items")
	}
}

func TestBlockList_Empty(t *testing.T) {
	built := rich.List().Build().(gotgbot.InputRichBlockList)
	if built.Items != nil {
		t.Error("Expected an empty list to carry no items")
	}
}
