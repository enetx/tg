package inline_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewVoice(t *testing.T) {
	voice := inline.NewVoice(testID, testURL, testTitle)
	if voice == nil {
		t.Error("Expected Voice to be created")
	}

	built := voice.Build()
	if built == nil {
		t.Error("Expected Voice to build correctly")
	}

	if built.GetType() != "voice" {
		t.Error("Expected Voice type to be 'voice'")
	}
}

func TestVoice_Caption(t *testing.T) {
	voice := inline.NewVoice(testID, testURL, testTitle)
	result := voice.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return Voice")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVoice); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVoice")
	}
}

func TestVoice_Duration(t *testing.T) {
	voice := inline.NewVoice(testID, testURL, testTitle)
	result := voice.Duration(30 * time.Second)
	if result == nil {
		t.Error("Expected Duration method to return Voice")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVoice); ok {
		if v.VoiceDuration != 30 {
			t.Error("Expected VoiceDuration to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVoice")
	}
}

func TestVoice_HTML(t *testing.T) {
	voice := inline.NewVoice(testID, testURL, testTitle)
	result := voice.HTML()
	if result == nil {
		t.Error("Expected HTML method to return Voice")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVoice); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVoice")
	}
}

func TestVoice_Markdown(t *testing.T) {
	voice := inline.NewVoice(testID, testURL, testTitle)
	result := voice.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return Voice")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVoice); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVoice")
	}
}

func TestVoice_Markup(t *testing.T) {
	voice := inline.NewVoice(testID, testURL, testTitle)
	testKeyboard := createTestKeyboard()
	result := voice.Markup(testKeyboard)
	if result == nil {
		t.Error("Expected Markup method to return Voice")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVoice); ok {
		if v.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVoice")
	}
}

func TestVoice_InputMessageContent(t *testing.T) {
	voice := inline.NewVoice(testID, testURL, testTitle)
	messageContent := createTestMessageContent()
	result := voice.InputMessageContent(messageContent)
	if result == nil {
		t.Error("Expected InputMessageContent method to return Voice")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVoice); ok {
		if v.InputMessageContent == nil {
			t.Error("Expected InputMessageContent to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVoice")
	}
}

func TestVoice_MethodChaining(t *testing.T) {
	msg := createTestMessageContent()
	result := inline.NewVoice(testID, testURL, testTitle).
		Caption(testCaption).
		Duration(30 * time.Second).
		HTML().
		Markup(createTestKeyboard()).
		InputMessageContent(msg)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Voice to build correctly")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
