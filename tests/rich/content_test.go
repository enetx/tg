package rich_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/rich"
)

func TestNewContent(t *testing.T) {
	content := rich.NewContent(rich.New().HTML(g.String("<b>bold</b>")))
	if content == nil {
		t.Fatal("Expected NewContent to create a Content builder")
	}

	built, ok := content.Build().(gotgbot.InputRichMessageContent)
	if !ok {
		t.Fatal("Expected result to be InputRichMessageContent")
	}

	if built.RichMessage.Html != "<b>bold</b>" {
		t.Error("Expected the rich message to be set correctly")
	}
}

func TestNewContent_Nil(t *testing.T) {
	built, ok := rich.NewContent(nil).Build().(gotgbot.InputRichMessageContent)
	if !ok {
		t.Fatal("Expected result to be InputRichMessageContent")
	}

	if built.RichMessage.Html != "" {
		t.Error("Expected a nil message to leave the content empty")
	}
}

func TestContent_Message(t *testing.T) {
	content := rich.NewContent(nil)

	result := content.Message(rich.New().Markdown(g.String("*bold*")))
	if result != content {
		t.Error("Expected Message to return same Content instance")
	}

	if result.Message(nil) != content {
		t.Error("Expected Message to return same Content instance for nil input")
	}

	built := result.Build().(gotgbot.InputRichMessageContent)
	if built.RichMessage.Markdown != "*bold*" {
		t.Error("Expected Message to keep the previous value when passed nil")
	}
}
