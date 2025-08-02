package inline_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewGif(t *testing.T) {
	gif := inline.NewGif(testID, testURL, testThumbnailURL)

	if gif == nil {
		t.Error("Expected Gif to be created")
	}

	built := gif.Build()
	if built == nil {
		t.Error("Expected Gif to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultGif); ok {
		if result.GetType() != "gif" {
			t.Error("Expected type to be 'gif'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultGif")
	}
}

func TestGif_Caption(t *testing.T) {
	gif := inline.NewGif(testID, testURL, testThumbnailURL)

	result := gif.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultGif); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultGif")
	}
}

func TestGif_Size(t *testing.T) {
	gif := inline.NewGif(testID, testURL, testThumbnailURL)

	result := gif.Size(320, 240)
	if result == nil {
		t.Error("Expected Size method to return Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultGif); ok {
		if v.GifWidth != 320 || v.GifHeight != 240 {
			t.Error("Expected GifWidth and GifHeight to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultGif")
	}
}

func TestGif_Duration(t *testing.T) {
	gif := inline.NewGif(testID, testURL, testThumbnailURL)

	result := gif.Duration(5 * time.Second)
	if result == nil {
		t.Error("Expected Duration method to return Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultGif); ok {
		if v.GifDuration != 5 {
			t.Error("Expected GifDuration to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultGif")
	}
}

func TestGif_Title(t *testing.T) {
	gif := inline.NewGif(testID, testURL, testThumbnailURL)

	result := gif.Title(testTitle)
	if result == nil {
		t.Error("Expected Title method to return Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultGif); ok {
		if v.Title != testTitle.Std() {
			t.Error("Expected Title to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultGif")
	}
}

func TestGif_HTML(t *testing.T) {
	gif := inline.NewGif(testID, testURL, testThumbnailURL)

	result := gif.HTML()
	if result == nil {
		t.Error("Expected HTML method to return Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultGif); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultGif")
	}
}

func TestGif_MethodChaining(t *testing.T) {
	messageContent := createTestMessageContent()

	result := inline.NewGif(testID, testURL, testThumbnailURL).
		Caption(testCaption).
		Size(320, 240).
		Duration(5 * time.Second).
		Title(testTitle).
		HTML().
		ShowCaptionAboveMedia().
		Markup(createTestKeyboard()).
		InputMessageContent(messageContent)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Gif to build correctly")
	}

	if _, ok := built.(gotgbot.InlineQueryResultGif); !ok {
		t.Error("Expected result to be InlineQueryResultGif")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
