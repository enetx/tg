package rich_test

import (
	"encoding/json"
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/rich"
	"github.com/enetx/tg/richtext"
)

func TestRichButtonActions(t *testing.T) {
	tests := []struct {
		name      string
		configure func(*rich.Button) *rich.Button
		field     string
		value     string
	}{
		{"Callback", func(b *rich.Button) *rich.Button { return b.Callback("data") }, "callback_data", `"data"`},
		{"URL", func(b *rich.Button) *rich.Button { return b.URL("https://example.com") }, "url", `"https://example.com"`},
		{"WebApp", func(b *rich.Button) *rich.Button { return b.WebApp("https://example.com") }, "web_app", `{"url":"https://example.com"}`},
		{"LoginURL", func(b *rich.Button) *rich.Button { return b.LoginURL("https://example.com") }, "login_url", `{"url":"https://example.com"}`},
		{"SwitchInlineQuery", func(b *rich.Button) *rich.Button { return b.SwitchInlineQuery("") }, "switch_inline_query", `""`},
		{"SwitchInlineQueryCurrentChat", func(b *rich.Button) *rich.Button { return b.SwitchInlineQueryCurrentChat("") }, "switch_inline_query_current_chat", `""`},
		{"SwitchInlineQueryChosenChat", func(b *rich.Button) *rich.Button {
			return b.SwitchInlineQueryChosenChat(gotgbot.SwitchInlineQueryChosenChat{Query: "search", AllowUserChats: true})
		}, "switch_inline_query_chosen_chat", `{"query":"search","allow_user_chats":true}`},
		{"CopyText", func(b *rich.Button) *rich.Button { return b.CopyText("copy") }, "copy_text", `{"text":"copy"}`},
		{"Disabled", func(b *rich.Button) *rich.Button { return b.Disabled() }, "disabled", `{}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			button := rich.NewButton(richtext.New("old"))
			if tt.configure(button).Text(richtext.New("Button")).Style(keyboard.ButtonStylePrimary) != button {
				t.Fatal("Expected the same builder for chaining")
			}
			data, err := json.Marshal(button.Std())
			if err != nil {
				t.Fatal(err)
			}
			var fields map[string]json.RawMessage
			if err := json.Unmarshal(data, &fields); err != nil {
				t.Fatal(err)
			}
			if string(fields[tt.field]) != tt.value || string(fields["text"]) != `"Button"` ||
				string(fields["style"]) != `"primary"` || len(fields) != 3 {
				t.Errorf("Unexpected button JSON: %s", data)
			}
		})
	}
}

func TestRichButtonTextNode(t *testing.T) {
	button := rich.NewButton(richtext.New("Click")).Callback("data").Link()
	node := richtext.Button(button.Std())
	data, err := json.Marshal(node.Std())
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"type":"button","button":{"text":"Click","style":"link","callback_data":"data"}}` {
		t.Errorf("Unexpected rich text button JSON: %s", data)
	}
}
