package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewCachedSticker(t *testing.T) {
	cached := inline.NewCachedSticker(testID, testFileID)

	if cached == nil {
		t.Error("Expected CachedSticker to be created")
	}

	built := cached.Build()
	if built == nil {
		t.Error("Expected CachedSticker to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultCachedSticker); ok {
		if result.GetType() != "sticker" {
			t.Error("Expected type to be 'sticker'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedSticker")
	}
}

func TestCachedSticker_Markup(t *testing.T) {
	cached := inline.NewCachedSticker(testID, testFileID)
	keyboard := createTestKeyboard()

	result := cached.Markup(keyboard)
	if result == nil {
		t.Error("Expected Markup method to return CachedSticker")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedSticker); ok {
		if v.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedSticker")
	}
}

func TestCachedSticker_InputMessageContent(t *testing.T) {
	cached := inline.NewCachedSticker(testID, testFileID)
	messageContent := createTestMessageContent()

	result := cached.InputMessageContent(messageContent)
	if result == nil {
		t.Error("Expected InputMessageContent method to return CachedSticker")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedSticker); ok {
		if v.InputMessageContent == nil {
			t.Error("Expected InputMessageContent to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedSticker")
	}
}

func TestCachedSticker_MethodChaining(t *testing.T) {
	messageContent := createTestMessageContent()

	result := inline.NewCachedSticker(testID, testFileID).
		Markup(createTestKeyboard()).
		InputMessageContent(messageContent)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained CachedSticker to build correctly")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
