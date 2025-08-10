package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewCachedDocument(t *testing.T) {
	cachedDoc := inline.NewCachedDocument(testID, testFileID, testTitle)

	if cachedDoc == nil {
		t.Error("Expected CachedDocument to be created")
	}

	built := cachedDoc.Build()
	if built == nil {
		t.Error("Expected CachedDocument to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultCachedDocument); ok {
		if result.GetType() != "document" {
			t.Error("Expected CachedDocument type to be 'document'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedDocument")
	}
}

func TestCachedDocument_Description(t *testing.T) {
	cachedDoc := inline.NewCachedDocument(testID, testFileID, testTitle)

	result := cachedDoc.Description(testDescription)
	if result == nil {
		t.Error("Expected Description method to return CachedDocument")
	}

	built := result.Build()
	if docResult, ok := built.(gotgbot.InlineQueryResultCachedDocument); ok {
		if docResult.Description != testDescription.Std() {
			t.Error("Expected Description to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedDocument")
	}
}

func TestCachedDocument_Caption(t *testing.T) {
	cachedDoc := inline.NewCachedDocument(testID, testFileID, testTitle)

	result := cachedDoc.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return CachedDocument")
	}

	built := result.Build()
	if docResult, ok := built.(gotgbot.InlineQueryResultCachedDocument); ok {
		if docResult.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedDocument")
	}
}

func TestCachedDocument_HTML(t *testing.T) {
	cachedDoc := inline.NewCachedDocument(testID, testFileID, testTitle)

	result := cachedDoc.HTML()
	if result == nil {
		t.Error("Expected HTML method to return CachedDocument")
	}

	built := result.Build()
	if docResult, ok := built.(gotgbot.InlineQueryResultCachedDocument); ok {
		if docResult.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedDocument")
	}
}

func TestCachedDocument_Markdown(t *testing.T) {
	cachedDoc := inline.NewCachedDocument(testID, testFileID, testTitle)

	result := cachedDoc.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return CachedDocument")
	}

	built := result.Build()
	if docResult, ok := built.(gotgbot.InlineQueryResultCachedDocument); ok {
		if docResult.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedDocument")
	}
}

func TestCachedDocument_CaptionEntities(t *testing.T) {
	cachedDoc := inline.NewCachedDocument(testID, testFileID, testTitle)
	entities := createTestEntities()

	result := cachedDoc.CaptionEntities(entities)
	if result == nil {
		t.Error("Expected CaptionEntities method to return CachedDocument")
	}

	built := result.Build()
	if docResult, ok := built.(gotgbot.InlineQueryResultCachedDocument); ok {
		if len(docResult.CaptionEntities) == 0 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedDocument")
	}
}

func TestCachedDocument_Markup(t *testing.T) {
	cachedDoc := inline.NewCachedDocument(testID, testFileID, testTitle)
	keyboard := createTestKeyboard()

	result := cachedDoc.Markup(keyboard)
	if result == nil {
		t.Error("Expected Markup method to return CachedDocument")
	}

	built := result.Build()
	if docResult, ok := built.(gotgbot.InlineQueryResultCachedDocument); ok {
		if docResult.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedDocument")
	}
}

func TestCachedDocument_InputMessageContent(t *testing.T) {
	cachedDoc := inline.NewCachedDocument(testID, testFileID, testTitle)
	messageContent := createTestMessageContent()

	result := cachedDoc.InputMessageContent(messageContent)
	if result == nil {
		t.Error("Expected InputMessageContent method to return CachedDocument")
	}

	built := result.Build()
	if docResult, ok := built.(gotgbot.InlineQueryResultCachedDocument); ok {
		if docResult.InputMessageContent == nil {
			t.Error("Expected InputMessageContent to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedDocument")
	}
}

func TestCachedDocument_MethodChaining(t *testing.T) {
	messageContent := createTestMessageContent()

	result := inline.NewCachedDocument(testID, testFileID, testTitle).
		Description(testDescription).
		Caption(testCaption).
		HTML().
		Markup(createTestKeyboard()).
		InputMessageContent(messageContent)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained CachedDocument to build correctly")
	}

	// Verify the result implements QueryResult interface
	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
