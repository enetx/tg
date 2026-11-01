package keyboard_test

import (
	"encoding/json"
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/keyboard"
)

func TestInlineKeyboardDisabledButton(t *testing.T) {
	kb := keyboard.Inline().
		Button(keyboard.NewButton().Text("Unavailable").Disabled()).
		Disabled("Coming soon").
		ForceReply()
	markup := kb.Markup().(gotgbot.InlineKeyboardMarkup)
	if len(markup.InlineKeyboard) != 1 || len(markup.InlineKeyboard[0]) != 2 {
		t.Fatal("Expected both disabled buttons in the row")
	}
	for _, button := range markup.InlineKeyboard[0] {
		if button.Disabled == nil || button.CallbackData != "" {
			t.Error("Expected a disabled button without callback data")
		}
	}
	data, err := json.Marshal(markup)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"inline_keyboard":[[{"text":"Unavailable","disabled":{}},{"text":"Coming soon","disabled":{}}]],"force_reply":true}` {
		t.Errorf("Unexpected inline keyboard JSON: %s", data)
	}
}

func TestInlineKeyboardForceReplyImport(t *testing.T) {
	original := keyboard.Inline().Text("Button", "data").ForceReply()
	markup := original.Markup().(gotgbot.InlineKeyboardMarkup)
	for _, source := range []any{original, markup, &markup} {
		cloned := keyboard.Inline(source).Markup().(gotgbot.InlineKeyboardMarkup)
		if !cloned.ForceReply || len(cloned.InlineKeyboard) != 1 {
			t.Error("Expected imported markup to retain ForceReply and buttons")
		}
	}
}

func TestReplyKeyboardForceReply(t *testing.T) {
	data, err := json.Marshal(keyboard.Reply().Text("Reply").ForceReply().Markup())
	if err != nil {
		t.Fatal(err)
	}
	var markup gotgbot.ReplyKeyboardMarkup
	if err := json.Unmarshal(data, &markup); err != nil {
		t.Fatal(err)
	}
	if !markup.ForceReply || !markup.ResizeKeyboard || markup.Keyboard[0][0].Text != "Reply" {
		t.Errorf("Unexpected reply keyboard JSON: %s", data)
	}
}
