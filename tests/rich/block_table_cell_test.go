package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestCell(t *testing.T) {
	cell := rich.Cell(testText("value"))
	if cell == nil {
		t.Fatal("Expected Cell to create a BlockTableCell builder")
	}

	built := cell.Std()
	if string(built.Text.(gotgbot.RichTextString)) != "value" {
		t.Error("Expected Text to be set correctly")
	}

	if built.Align != "left" || built.Valign != "top" {
		t.Error("Expected Cell to default to left/top alignment")
	}
}

func TestBlockTableCell_Text(t *testing.T) {
	cell := rich.Cell(testText("old"))

	result := cell.Text(testText("new"))
	if result != cell {
		t.Error("Expected Text to return same BlockTableCell instance")
	}

	if string(result.Std().Text.(gotgbot.RichTextString)) != "new" {
		t.Error("Expected Text to be set correctly")
	}
}

func TestBlockTableCell_Header(t *testing.T) {
	cell := rich.Cell(testText("value"))

	result := cell.Header()
	if result != cell {
		t.Error("Expected Header to return same BlockTableCell instance")
	}

	if !result.Std().IsHeader {
		t.Error("Expected IsHeader to be set correctly")
	}
}

func TestBlockTableCell_Spans(t *testing.T) {
	cell := rich.Cell(testText("value"))

	if cell.Colspan(2) != cell || cell.Rowspan(3) != cell {
		t.Error("Expected span setters to return same BlockTableCell instance")
	}

	built := cell.Std()
	if built.Colspan != 2 || built.Rowspan != 3 {
		t.Error("Expected Colspan and Rowspan to be set correctly")
	}
}

func TestBlockTableCell_Alignment(t *testing.T) {
	cell := rich.Cell(testText("value"))

	if cell.Align(rich.AlignCenter) != cell || cell.Valign(rich.ValignBottom) != cell {
		t.Error("Expected alignment setters to return same BlockTableCell instance")
	}

	built := cell.Std()
	if built.Align != "center" || built.Valign != "bottom" {
		t.Error("Expected Align and Valign to be set correctly")
	}
}
