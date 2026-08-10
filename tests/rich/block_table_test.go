package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestTable(t *testing.T) {
	block := rich.Table()
	if block == nil {
		t.Fatal("Expected Table to create a BlockTable builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockTable)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockTable")
	}

	if built.Cells != nil {
		t.Error("Expected an empty table to carry no cells")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeTable {
		t.Error("Expected block type to be table")
	}
}

func TestBlockTable_Row(t *testing.T) {
	block := rich.Table()

	result := block.
		Row(rich.Cell(testText("h1")).Header(), rich.Cell(testText("h2")).Header()).
		Row(rich.Cell(testText("v1")), nil)

	if result != block {
		t.Error("Expected Row to return same BlockTable instance")
	}

	built := result.Build().(gotgbot.InputRichBlockTable)
	if len(built.Cells) != 2 {
		t.Fatal("Expected two rows to be added")
	}

	if len(built.Cells[0]) != 2 || len(built.Cells[1]) != 1 {
		t.Error("Expected Row to skip nil cells")
	}

	if !built.Cells[0][0].IsHeader {
		t.Error("Expected the first row to keep header cells")
	}
}

func TestBlockTable_Bordered(t *testing.T) {
	block := rich.Table()

	result := block.Bordered()
	if result != block {
		t.Error("Expected Bordered to return same BlockTable instance")
	}

	if !result.Build().(gotgbot.InputRichBlockTable).IsBordered {
		t.Error("Expected IsBordered to be set correctly")
	}
}

func TestBlockTable_Striped(t *testing.T) {
	block := rich.Table()

	result := block.Striped()
	if result != block {
		t.Error("Expected Striped to return same BlockTable instance")
	}

	if !result.Build().(gotgbot.InputRichBlockTable).IsStriped {
		t.Error("Expected IsStriped to be set correctly")
	}
}

func TestBlockTable_Caption(t *testing.T) {
	block := rich.Table()

	result := block.Caption(testText(testCaption))
	if result != block {
		t.Error("Expected Caption to return same BlockTable instance")
	}

	built := result.Build().(gotgbot.InputRichBlockTable)
	if string(built.Caption.(gotgbot.RichTextString)) != testCaption.Std() {
		t.Error("Expected Caption to be set correctly")
	}
}
