package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/rich"
)

func TestBlockButtons(t *testing.T) {
	first := rich.NewButton(testText("First")).Callback("first")
	second := rich.NewButton(testText("Second")).Disabled()
	row := rich.Buttons(nil, first).Button(nil, second).Align(rich.AlignCenter)
	built := row.Build().(gotgbot.InputRichBlockButtons)
	if built.GetType() != "buttons" || built.Align != "center" || len(built.Buttons) != 2 {
		t.Fatalf("Unexpected button block: %+v", built)
	}
	if built.Buttons[0].CallbackData != "first" || built.Buttons[1].Disabled == nil {
		t.Error("Expected callback and disabled buttons in order")
	}
	first.Callback("updated")
	if row.Build().(gotgbot.InputRichBlockButtons).Buttons[0].CallbackData != "updated" {
		t.Error("Build must use the current button values")
	}
	if built.Buttons[0].CallbackData != "first" {
		t.Error("Build must return an independent snapshot")
	}
}
