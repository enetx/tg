package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/preview"
)

func TestText(t *testing.T) {
	text := input.Text(testText)
	if text == nil {
		t.Error("Expected MessageText to be created")
	}
	if !assertMessageContent(text) {
		t.Error("MessageText should implement MessageContent correctly")
	}
}

func TestText_HTML(t *testing.T) {
	text := input.Text(testText)
	result := text.HTML()
	if result == nil {
		t.Error("Expected HTML method to return MessageText")
	}
	if result != text {
		t.Error("Expected HTML to return same MessageText instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputTextMessageContent); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InputTextMessageContent")
	}
}

func TestText_Markdown(t *testing.T) {
	text := input.Text(testText)
	result := text.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return MessageText")
	}
	if result != text {
		t.Error("Expected Markdown to return same MessageText instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputTextMessageContent); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InputTextMessageContent")
	}
}

func TestText_Entities(t *testing.T) {
	text := input.Text(testText)
	entities := createTestEntities()
	result := text.Entities(entities)
	if result == nil {
		t.Error("Expected Entities method to return MessageText")
	}
	if result != text {
		t.Error("Expected Entities to return same MessageText instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputTextMessageContent); ok {
		if len(v.Entities) == 0 {
			t.Error("Expected Entities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputTextMessageContent")
	}
}

func TestText_Preview(t *testing.T) {
	text := input.Text(testText)
	previewOpts := preview.New().URL(g.String("https://example.com"))
	result := text.Preview(previewOpts)
	if result == nil {
		t.Error("Expected Preview method to return MessageText")
	}
	if result != text {
		t.Error("Expected Preview to return same MessageText instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputTextMessageContent); ok {
		if v.LinkPreviewOptions == nil {
			t.Error("Expected LinkPreviewOptions to be set")
		}
	} else {
		t.Error("Expected result to be InputTextMessageContent")
	}
}

func TestText_Build(t *testing.T) {
	text := input.Text(testText)
	built := text.Build()

	if v, ok := built.(gotgbot.InputTextMessageContent); ok {
		if v.MessageText != testText.Std() {
			t.Errorf("Expected MessageText to be %s, got %s", testText.Std(), v.MessageText)
		}
	} else {
		t.Error("Expected result to be InputTextMessageContent")
	}
}

func TestText_MethodChaining(t *testing.T) {
	result := input.Text(testText).
		HTML()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Text to build correctly")
	}

	if _, ok := built.(gotgbot.InputTextMessageContent); !ok {
		t.Error("Expected result to be InputTextMessageContent")
	}

	if !assertMessageContent(result) {
		t.Error("Expected result to implement MessageContent interface")
	}
}
