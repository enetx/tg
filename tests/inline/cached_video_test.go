package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewCachedVideo(t *testing.T) {
	cached := inline.NewCachedVideo(testID, testFileID, testTitle)

	if cached == nil {
		t.Error("Expected CachedVideo to be created")
	}

	built := cached.Build()
	if built == nil {
		t.Error("Expected CachedVideo to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultCachedVideo); ok {
		if result.GetType() != "video" {
			t.Error("Expected type to be 'video'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVideo")
	}
}

func TestCachedVideo_Description(t *testing.T) {
	cached := inline.NewCachedVideo(testID, testFileID, testTitle)

	result := cached.Description(testDescription)
	if result == nil {
		t.Error("Expected Description method to return CachedVideo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVideo); ok {
		if v.Description != testDescription.Std() {
			t.Error("Expected Description to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVideo")
	}
}

func TestCachedVideo_Caption(t *testing.T) {
	cached := inline.NewCachedVideo(testID, testFileID, testTitle)

	result := cached.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return CachedVideo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVideo); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVideo")
	}
}

func TestCachedVideo_HTML(t *testing.T) {
	cached := inline.NewCachedVideo(testID, testFileID, testTitle)

	result := cached.HTML()
	if result == nil {
		t.Error("Expected HTML method to return CachedVideo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVideo); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVideo")
	}
}

func TestCachedVideo_Markdown(t *testing.T) {
	cached := inline.NewCachedVideo(testID, testFileID, testTitle)

	result := cached.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return CachedVideo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVideo); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVideo")
	}
}

func TestCachedVideo_CaptionEntities(t *testing.T) {
	cached := inline.NewCachedVideo(testID, testFileID, testTitle)
	entities := createTestEntities()

	result := cached.CaptionEntities(entities)
	if result == nil {
		t.Error("Expected CaptionEntities method to return CachedVideo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVideo); ok {
		if len(v.CaptionEntities) == 0 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVideo")
	}
}

func TestCachedVideo_ShowCaptionAboveMedia(t *testing.T) {
	cached := inline.NewCachedVideo(testID, testFileID, testTitle)

	result := cached.ShowCaptionAboveMedia()
	if result == nil {
		t.Error("Expected ShowCaptionAboveMedia method to return CachedVideo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVideo); ok {
		if !v.ShowCaptionAboveMedia {
			t.Error("Expected ShowCaptionAboveMedia to be true")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVideo")
	}
}

func TestCachedVideo_Markup(t *testing.T) {
	cached := inline.NewCachedVideo(testID, testFileID, testTitle)
	keyboard := createTestKeyboard()

	result := cached.Markup(keyboard)
	if result == nil {
		t.Error("Expected Markup method to return CachedVideo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedVideo); ok {
		if v.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedVideo")
	}
}

func TestCachedVideo_MethodChaining(t *testing.T) {
	messageContent := createTestMessageContent()

	result := inline.NewCachedVideo(testID, testFileID, testTitle).
		Description(testDescription).
		Caption(testCaption).
		HTML().
		ShowCaptionAboveMedia().
		Markup(createTestKeyboard()).
		InputMessageContent(messageContent)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained CachedVideo to build correctly")
	}

	if _, ok := built.(gotgbot.InlineQueryResultCachedVideo); !ok {
		t.Error("Expected result to be InlineQueryResultCachedVideo")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
