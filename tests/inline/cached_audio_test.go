package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

func TestNewCachedAudio(t *testing.T) {
	cachedAudio := inline.NewCachedAudio(testID, testFileID)

	if cachedAudio == nil {
		t.Error("Expected CachedAudio to be created")
	}

	built := cachedAudio.Build()
	if built == nil {
		t.Error("Expected CachedAudio to build correctly")
	}

	if audioResult, ok := built.(gotgbot.InlineQueryResultCachedAudio); ok {
		if audioResult.GetType() != "audio" {
			t.Error("Expected CachedAudio type to be 'audio'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedAudio")
	}
}

func TestCachedAudio_Caption(t *testing.T) {
	cachedAudio := inline.NewCachedAudio(testID, testFileID)

	result := cachedAudio.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return CachedAudio")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultCachedAudio); ok {
		if audioResult.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedAudio")
	}
}

func TestCachedAudio_HTML(t *testing.T) {
	cachedAudio := inline.NewCachedAudio(testID, testFileID)

	result := cachedAudio.HTML()
	if result == nil {
		t.Error("Expected HTML method to return CachedAudio")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultCachedAudio); ok {
		if audioResult.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedAudio")
	}
}

func TestCachedAudio_Markdown(t *testing.T) {
	cachedAudio := inline.NewCachedAudio(testID, testFileID)

	result := cachedAudio.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return CachedAudio")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultCachedAudio); ok {
		if audioResult.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedAudio")
	}
}

func TestCachedAudio_Markup(t *testing.T) {
	cachedAudio := inline.NewCachedAudio(testID, testFileID)
	keyboard := createTestKeyboard()

	result := cachedAudio.Markup(keyboard)
	if result == nil {
		t.Error("Expected Markup method to return CachedAudio")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultCachedAudio); ok {
		if audioResult.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedAudio")
	}
}

func TestCachedAudio_InputMessageContent(t *testing.T) {
	cachedAudio := inline.NewCachedAudio(testID, testFileID)
	messageContent := createTestMessageContent()

	result := cachedAudio.InputMessageContent(messageContent)
	if result == nil {
		t.Error("Expected InputMessageContent method to return CachedAudio")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultCachedAudio); ok {
		if audioResult.InputMessageContent == nil {
			t.Error("Expected InputMessageContent to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedAudio")
	}
}

func TestCachedAudio_CaptionEntities(t *testing.T) {
	cachedAudio := inline.NewCachedAudio(testID, testFileID)
	testText := g.String("Bold text")
	entities := testEntities(testText)

	result := cachedAudio.CaptionEntities(entities)
	if result == nil {
		t.Error("Expected CaptionEntities method to return CachedAudio")
	}

	// Test method chaining returns same instance
	if result != cachedAudio {
		t.Error("Expected CaptionEntities method to return same CachedAudio instance")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultCachedAudio); ok {
		if len(audioResult.CaptionEntities) != 1 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
		if audioResult.CaptionEntities[0].Type != "bold" {
			t.Error("Expected first entity to be bold")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedAudio")
	}
}

func TestCachedAudio_MethodChaining(t *testing.T) {
	messageContent := createTestMessageContent()

	result := inline.NewCachedAudio(testID, testFileID).
		Caption(testCaption).
		HTML().
		Markup(createTestKeyboard()).
		InputMessageContent(messageContent)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained CachedAudio to build correctly")
	}

	// Verify the result implements QueryResult interface
	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
