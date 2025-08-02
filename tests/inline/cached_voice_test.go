package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewCachedVoice(t *testing.T) {
	cached := inline.NewCachedVoice(testID, testFileID, testTitle)

	if cached == nil {
		t.Error("Expected CachedVoice to be created")
	}

	built := cached.Build()
	if built == nil {
		t.Error("Expected CachedVoice to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultCachedVoice); ok {
		if result.GetType() != "voice" {
			t.Error("Expected type to be 'voice'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVoice")
	}
}

func TestCachedVoice_Caption(t *testing.T) {
	cached := inline.NewCachedVoice(testID, testFileID, testTitle)

	result := cached.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return CachedVoice")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVoice); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVoice")
	}
}

func TestCachedVoice_HTML(t *testing.T) {
	cached := inline.NewCachedVoice(testID, testFileID, testTitle)

	result := cached.HTML()
	if result == nil {
		t.Error("Expected HTML method to return CachedVoice")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVoice); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVoice")
	}
}

func TestCachedVoice_Markdown(t *testing.T) {
	cached := inline.NewCachedVoice(testID, testFileID, testTitle)

	result := cached.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return CachedVoice")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVoice); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVoice")
	}
}

func TestCachedVoice_Markup(t *testing.T) {
	cached := inline.NewCachedVoice(testID, testFileID, testTitle)
	keyboard := createTestKeyboard()

	result := cached.Markup(keyboard)
	if result == nil {
		t.Error("Expected Markup method to return CachedVoice")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVoice); ok {
		if v.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVoice")
	}
}

func TestCachedVoice_MethodChaining(t *testing.T) {
	messageContent := createTestMessageContent()

	result := inline.NewCachedVoice(testID, testFileID, testTitle).
		Caption(testCaption).
		HTML().
		Markup(createTestKeyboard()).
		InputMessageContent(messageContent)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained CachedVoice to build correctly")
	}

	if _, ok := built.(gotgbot.InlineQueryResultCachedVoice); !ok {
		t.Error("Expected result to be InlineQueryResultCachedVoice")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
