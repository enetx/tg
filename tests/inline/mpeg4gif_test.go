package inline_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

func TestNewMpeg4Gif(t *testing.T) {
	mpeg4Gif := inline.NewMpeg4Gif(testID, testURL, testThumbnailURL)

	if mpeg4Gif == nil {
		t.Error("Expected Mpeg4Gif to be created")
	}

	built := mpeg4Gif.Build()
	if built == nil {
		t.Error("Expected Mpeg4Gif to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultMpeg4Gif); ok {
		if result.GetType() != "mpeg4_gif" {
			t.Error("Expected type to be 'mpeg4_gif'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultMpeg4Gif")
	}
}

func TestMpeg4Gif_Caption(t *testing.T) {
	mpeg4Gif := inline.NewMpeg4Gif(testID, testURL, testThumbnailURL)

	result := mpeg4Gif.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return Mpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultMpeg4Gif); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultMpeg4Gif")
	}
}

func TestMpeg4Gif_Size(t *testing.T) {
	mpeg4Gif := inline.NewMpeg4Gif(testID, testURL, testThumbnailURL)

	result := mpeg4Gif.Size(320, 240)
	if result == nil {
		t.Error("Expected Size method to return Mpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultMpeg4Gif); ok {
		if v.Mpeg4Width != 320 || v.Mpeg4Height != 240 {
			t.Error("Expected Mpeg4Width and Mpeg4Height to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultMpeg4Gif")
	}
}

func TestMpeg4Gif_Duration(t *testing.T) {
	mpeg4Gif := inline.NewMpeg4Gif(testID, testURL, testThumbnailURL)

	result := mpeg4Gif.Duration(5 * time.Second)
	if result == nil {
		t.Error("Expected Duration method to return Mpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultMpeg4Gif); ok {
		if v.Mpeg4Duration != 5 {
			t.Error("Expected Mpeg4Duration to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultMpeg4Gif")
	}
}

func TestMpeg4Gif_Title(t *testing.T) {
	mpeg4Gif := inline.NewMpeg4Gif(testID, testURL, testThumbnailURL)

	result := mpeg4Gif.Title(testTitle)
	if result == nil {
		t.Error("Expected Title method to return Mpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultMpeg4Gif); ok {
		if v.Title != testTitle.Std() {
			t.Error("Expected Title to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultMpeg4Gif")
	}
}

func TestMpeg4Gif_ThumbnailMimeType(t *testing.T) {
	mpeg4Gif := inline.NewMpeg4Gif(testID, testURL, testThumbnailURL)
	mimeType := g.String("image/jpeg")

	result := mpeg4Gif.ThumbnailMimeType(mimeType)
	if result == nil {
		t.Error("Expected ThumbnailMimeType method to return Mpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultMpeg4Gif); ok {
		if v.ThumbnailMimeType != mimeType.Std() {
			t.Error("Expected ThumbnailMimeType to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultMpeg4Gif")
	}
}

func TestMpeg4Gif_HTML(t *testing.T) {
	mpeg4Gif := inline.NewMpeg4Gif(testID, testURL, testThumbnailURL)

	result := mpeg4Gif.HTML()
	if result == nil {
		t.Error("Expected HTML method to return Mpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultMpeg4Gif); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultMpeg4Gif")
	}
}

func TestMpeg4Gif_Markdown(t *testing.T) {
	mpeg4Gif := inline.NewMpeg4Gif(testID, testURL, testThumbnailURL)

	result := mpeg4Gif.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return Mpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultMpeg4Gif); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultMpeg4Gif")
	}
}

func TestMpeg4Gif_CaptionEntities(t *testing.T) {
	mpeg4Gif := inline.NewMpeg4Gif(testID, testURL, testThumbnailURL)
	entities := createTestEntities()

	result := mpeg4Gif.CaptionEntities(entities)
	if result == nil {
		t.Error("Expected CaptionEntities method to return Mpeg4Gif")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultMpeg4Gif); ok {
		if len(v.CaptionEntities) == 0 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultMpeg4Gif")
	}
}

func TestMpeg4Gif_MethodChaining(t *testing.T) {
	mimeType := g.String("image/jpeg")
	messageContent := createTestMessageContent()

	result := inline.NewMpeg4Gif(testID, testURL, testThumbnailURL).
		Caption(testCaption).
		Size(320, 240).
		Duration(5 * time.Second).
		Title(testTitle).
		ThumbnailMimeType(mimeType).
		HTML().
		ShowCaptionAboveMedia().
		Markup(createTestKeyboard()).
		InputMessageContent(messageContent)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Mpeg4Gif to build correctly")
	}

	if _, ok := built.(gotgbot.InlineQueryResultMpeg4Gif); !ok {
		t.Error("Expected result to be InlineQueryResultMpeg4Gif")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
