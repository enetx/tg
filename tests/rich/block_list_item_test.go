package rich_test

import (
	"testing"

	"github.com/enetx/tg/rich"
)

func TestListItem(t *testing.T) {
	item := rich.ListItem(rich.Paragraph(testText("Item")), nil)
	if item == nil {
		t.Fatal("Expected ListItem to create a BlockListItem builder")
	}

	if len(item.Std().Blocks) != 1 {
		t.Error("Expected ListItem to skip nil blocks")
	}
}

func TestBlockListItem_Block(t *testing.T) {
	item := rich.ListItem()

	result := item.Block(rich.Paragraph(testText("a")), rich.Divider())
	if result != item {
		t.Error("Expected Block to return same BlockListItem instance")
	}

	if len(result.Std().Blocks) != 2 {
		t.Error("Expected Block to append blocks")
	}
}

func TestBlockListItem_Checkbox(t *testing.T) {
	item := rich.ListItem()

	result := item.Checkbox()
	if result != item {
		t.Error("Expected Checkbox to return same BlockListItem instance")
	}

	built := result.Std()
	if !built.HasCheckbox || built.IsChecked {
		t.Error("Expected Checkbox to add an unchecked checkbox")
	}
}

func TestBlockListItem_Checked(t *testing.T) {
	item := rich.ListItem()

	result := item.Checked()
	if result != item {
		t.Error("Expected Checked to return same BlockListItem instance")
	}

	built := result.Std()
	if !built.HasCheckbox || !built.IsChecked {
		t.Error("Expected Checked to add a checked checkbox")
	}
}

func TestBlockListItem_Value(t *testing.T) {
	item := rich.ListItem()

	result := item.Value(3)
	if result != item {
		t.Error("Expected Value to return same BlockListItem instance")
	}

	if result.Std().Value != 3 {
		t.Error("Expected Value to be set correctly")
	}
}

func TestBlockListItem_Label(t *testing.T) {
	item := rich.ListItem()

	result := item.Label(rich.LabelUpperRoman)
	if result != item {
		t.Error("Expected Label to return same BlockListItem instance")
	}

	if result.Std().Type != "I" {
		t.Error("Expected Label to be set correctly")
	}
}

func TestBlockListItem_Empty(t *testing.T) {
	if rich.ListItem().Std().Blocks != nil {
		t.Error("Expected an empty item to carry no blocks")
	}
}
