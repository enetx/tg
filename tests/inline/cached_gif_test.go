package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewCachedGif(t *testing.T) {
	cachedGif := inline.NewCachedGif(testID, testFileID)

	if cachedGif == nil {
		t.Error("Expected CachedGif to be created")
	}

	built := cachedGif.Build()
	if built == nil {
		t.Error("Expected CachedGif to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultCachedGif); ok {
		if result.GetType() != "gif" {
			t.Error("Expected CachedGif type to be 'gif'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedGif")
	}
}

func TestCachedGif_Title(t *testing.T) {
	cachedGif := inline.NewCachedGif(testID, testFileID)

	result := cachedGif.Title(testTitle)
	if result == nil {
		t.Error("Expected Title method to return CachedGif")
	}

	built := result.Build()
	if gifResult, ok := built.(gotgbot.InlineQueryResultCachedGif); ok {
		if gifResult.Title != testTitle.Std() {
			t.Error("Expected Title to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedGif")
	}
}

func TestCachedGif_Caption(t *testing.T) {
	cachedGif := inline.NewCachedGif(testID, testFileID)

	result := cachedGif.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return CachedGif")
	}

	built := result.Build()
	if gifResult, ok := built.(gotgbot.InlineQueryResultCachedGif); ok {
		if gifResult.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedGif")
	}
}

func TestCachedGif_HTML(t *testing.T) {
	cachedGif := inline.NewCachedGif(testID, testFileID)

	result := cachedGif.HTML()
	if result == nil {
		t.Error("Expected HTML method to return CachedGif")
	}

	built := result.Build()
	if gifResult, ok := built.(gotgbot.InlineQueryResultCachedGif); ok {
		if gifResult.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedGif")
	}
}

func TestCachedGif_ShowCaptionAboveMedia(t *testing.T) {
	cachedGif := inline.NewCachedGif(testID, testFileID)

	result := cachedGif.ShowCaptionAboveMedia()
	if result == nil {
		t.Error("Expected ShowCaptionAboveMedia method to return CachedGif")
	}

	built := result.Build()
	if gifResult, ok := built.(gotgbot.InlineQueryResultCachedGif); ok {
		if !gifResult.ShowCaptionAboveMedia {
			t.Error("Expected ShowCaptionAboveMedia to be set to true")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedGif")
	}
}

func TestCachedGif_Markup(t *testing.T) {
	cachedGif := inline.NewCachedGif(testID, testFileID)
	keyboard := createTestKeyboard()

	result := cachedGif.Markup(keyboard)
	if result == nil {
		t.Error("Expected Markup method to return CachedGif")
	}

	built := result.Build()
	if gifResult, ok := built.(gotgbot.InlineQueryResultCachedGif); ok {
		if gifResult.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedGif")
	}
}

func TestCachedGif_InputMessageContent(t *testing.T) {
	cachedGif := inline.NewCachedGif(testID, testFileID)
	messageContent := createTestMessageContent()

	result := cachedGif.InputMessageContent(messageContent)
	if result == nil {
		t.Error("Expected InputMessageContent method to return CachedGif")
	}

	built := result.Build()
	if gifResult, ok := built.(gotgbot.InlineQueryResultCachedGif); ok {
		if gifResult.InputMessageContent == nil {
			t.Error("Expected InputMessageContent to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedGif")
	}
}

func TestCachedGif_MethodChaining(t *testing.T) {
	messageContent := createTestMessageContent()

	result := inline.NewCachedGif(testID, testFileID).
		Title(testTitle).
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
		t.Error("Expected chained CachedGif to build correctly")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
