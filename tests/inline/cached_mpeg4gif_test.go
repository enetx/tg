package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewCachedMpeg4Gif(t *testing.T) {
	cached := inline.NewCachedMpeg4Gif(testID, testFileID)

	if cached == nil {
		t.Error("Expected CachedMpeg4Gif to be created")
	}

	built := cached.Build()
	if built == nil {
		t.Error("Expected CachedMpeg4Gif to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultCachedMpeg4Gif); ok {
		if result.GetType() != "mpeg4_gif" {
			t.Error("Expected type to be 'mpeg4_gif'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedMpeg4Gif")
	}
}

func TestCachedMpeg4Gif_Title(t *testing.T) {
	cached := inline.NewCachedMpeg4Gif(testID, testFileID)

	result := cached.Title(testTitle)
	if result == nil {
		t.Error("Expected Title method to return CachedMpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedMpeg4Gif); ok {
		if v.Title != testTitle.Std() {
			t.Error("Expected Title to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedMpeg4Gif")
	}
}

func TestCachedMpeg4Gif_Caption(t *testing.T) {
	cached := inline.NewCachedMpeg4Gif(testID, testFileID)

	result := cached.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return CachedMpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedMpeg4Gif); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedMpeg4Gif")
	}
}

func TestCachedMpeg4Gif_HTML(t *testing.T) {
	cached := inline.NewCachedMpeg4Gif(testID, testFileID)

	result := cached.HTML()
	if result == nil {
		t.Error("Expected HTML method to return CachedMpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedMpeg4Gif); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedMpeg4Gif")
	}
}

func TestCachedMpeg4Gif_ShowCaptionAboveMedia(t *testing.T) {
	cached := inline.NewCachedMpeg4Gif(testID, testFileID)

	result := cached.ShowCaptionAboveMedia()
	if result == nil {
		t.Error("Expected ShowCaptionAboveMedia method to return CachedMpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedMpeg4Gif); ok {
		if !v.ShowCaptionAboveMedia {
			t.Error("Expected ShowCaptionAboveMedia to be true")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedMpeg4Gif")
	}
}

func TestCachedMpeg4Gif_Markup(t *testing.T) {
	cached := inline.NewCachedMpeg4Gif(testID, testFileID)
	keyboard := createTestKeyboard()

	result := cached.Markup(keyboard)
	if result == nil {
		t.Error("Expected Markup method to return CachedMpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultCachedMpeg4Gif); ok {
		if v.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultCachedMpeg4Gif")
	}
}

func TestCachedMpeg4Gif_MethodChaining(t *testing.T) {
	messageContent := createTestMessageContent()

	result := inline.NewCachedMpeg4Gif(testID, testFileID).
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
		t.Error("Expected chained CachedMpeg4Gif to build correctly")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
