package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewCachedPhoto(t *testing.T) {
	cached := inline.NewCachedPhoto(testID, testFileID)

	if cached == nil {
		t.Error("Expected CachedPhoto to be created")
	}

	built := cached.Build()
	if built == nil {
		t.Error("Expected CachedPhoto to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultCachedPhoto); ok {
		if result.GetType() != "photo" {
			t.Error("Expected type to be 'photo'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedPhoto")
	}
}

func TestCachedPhoto_Title(t *testing.T) {
	cached := inline.NewCachedPhoto(testID, testFileID)

	result := cached.Title(testTitle)
	if result == nil {
		t.Error("Expected Title method to return CachedPhoto")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedPhoto); ok {
		if v.Title != testTitle.Std() {
			t.Error("Expected Title to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedPhoto")
	}
}

func TestCachedPhoto_Description(t *testing.T) {
	cached := inline.NewCachedPhoto(testID, testFileID)

	result := cached.Description(testDescription)
	if result == nil {
		t.Error("Expected Description method to return CachedPhoto")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedPhoto); ok {
		if v.Description != testDescription.Std() {
			t.Error("Expected Description to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedPhoto")
	}
}

func TestCachedPhoto_Caption(t *testing.T) {
	cached := inline.NewCachedPhoto(testID, testFileID)

	result := cached.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return CachedPhoto")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedPhoto); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedPhoto")
	}
}

func TestCachedPhoto_HTML(t *testing.T) {
	cached := inline.NewCachedPhoto(testID, testFileID)

	result := cached.HTML()
	if result == nil {
		t.Error("Expected HTML method to return CachedPhoto")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedPhoto); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedPhoto")
	}
}

func TestCachedPhoto_Markdown(t *testing.T) {
	cached := inline.NewCachedPhoto(testID, testFileID)

	result := cached.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return CachedPhoto")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedPhoto); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedPhoto")
	}
}

func TestCachedPhoto_CaptionEntities(t *testing.T) {
	cached := inline.NewCachedPhoto(testID, testFileID)
	entities := createTestEntities()

	result := cached.CaptionEntities(entities)
	if result == nil {
		t.Error("Expected CaptionEntities method to return CachedPhoto")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedPhoto); ok {
		if len(v.CaptionEntities) == 0 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedPhoto")
	}
}

func TestCachedPhoto_ShowCaptionAboveMedia(t *testing.T) {
	cached := inline.NewCachedPhoto(testID, testFileID)

	result := cached.ShowCaptionAboveMedia()
	if result == nil {
		t.Error("Expected ShowCaptionAboveMedia method to return CachedPhoto")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedPhoto); ok {
		if !v.ShowCaptionAboveMedia {
			t.Error("Expected ShowCaptionAboveMedia to be true")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedPhoto")
	}
}

func TestCachedPhoto_InputMessageContent(t *testing.T) {
	cached := inline.NewCachedPhoto(testID, testFileID)
	messageContent := createTestMessageContent()

	result := cached.InputMessageContent(messageContent)
	if result == nil {
		t.Error("Expected InputMessageContent method to return CachedPhoto")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedPhoto); ok {
		if v.InputMessageContent == nil {
			t.Error("Expected InputMessageContent to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedPhoto")
	}
}

func TestCachedPhoto_Markup(t *testing.T) {
	cached := inline.NewCachedPhoto(testID, testFileID)
	keyboard := createTestKeyboard()

	result := cached.Markup(keyboard)
	if result == nil {
		t.Error("Expected Markup method to return CachedPhoto")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedPhoto); ok {
		if v.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedPhoto")
	}
}

func TestCachedPhoto_MethodChaining(t *testing.T) {
	result := inline.NewCachedPhoto(testID, testFileID).
		Title(testTitle).
		Description(testDescription).
		Caption(testCaption).
		HTML().
		Markup(createTestKeyboard())

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained CachedPhoto to build correctly")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
