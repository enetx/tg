package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestMap(t *testing.T) {
	block := rich.Map(40.7128, -74.0060)
	if block == nil {
		t.Fatal("Expected Map to create a BlockMap builder")
	}

	built, ok := block.Build().(gotgbot.InputRichBlockMap)
	if !ok {
		t.Fatal("Expected result to be InputRichBlockMap")
	}

	if built.Location.Latitude != 40.7128 || built.Location.Longitude != -74.0060 {
		t.Error("Expected Location to be set correctly")
	}

	if built.GetType() != gotgbot.InputRichBlockTypeMap {
		t.Error("Expected block type to be map")
	}
}

func TestBlockMap_Accuracy(t *testing.T) {
	block := rich.Map(1, 2)

	result := block.Accuracy(50)
	if result != block {
		t.Error("Expected Accuracy to return same BlockMap instance")
	}

	if result.Build().(gotgbot.InputRichBlockMap).Location.HorizontalAccuracy != 50 {
		t.Error("Expected HorizontalAccuracy to be set correctly")
	}
}

func TestBlockMap_Zoom(t *testing.T) {
	block := rich.Map(1, 2)

	result := block.Zoom(12)
	if result != block {
		t.Error("Expected Zoom to return same BlockMap instance")
	}

	if result.Build().(gotgbot.InputRichBlockMap).Zoom != 12 {
		t.Error("Expected Zoom to be set correctly")
	}
}

func TestBlockMap_Size(t *testing.T) {
	block := rich.Map(1, 2)

	result := block.Size(640, 480)
	if result != block {
		t.Error("Expected Size to return same BlockMap instance")
	}

	built := result.Build().(gotgbot.InputRichBlockMap)
	if built.Width != 640 || built.Height != 480 {
		t.Error("Expected Width and Height to be set correctly")
	}
}

func TestBlockMap_Caption(t *testing.T) {
	block := rich.Map(1, 2)

	result := block.Caption(rich.NewCaption(testText(testCaption)))
	if result != block {
		t.Error("Expected Caption to return same BlockMap instance")
	}

	built := result.Build().(gotgbot.InputRichBlockMap)
	if built.Caption == nil {
		t.Fatal("Expected Caption to be set")
	}

	if string(built.Caption.Text.(gotgbot.RichTextString)) != testCaption.Std() {
		t.Error("Expected Caption text to be set correctly")
	}
}
