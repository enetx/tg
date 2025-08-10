package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/inline"
)

func TestNewPhoto(t *testing.T) {
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)
	if photo == nil {
		t.Error("Expected Photo to be created")
	}

	built := photo.Build()
	if built == nil {
		t.Error("Expected Photo to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if result.GetType() != "photo" {
			t.Error("Expected Photo type to be 'photo'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}
}

func TestPhoto_Size(t *testing.T) {
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)
	result := photo.Size(800, 600)
	if result == nil {
		t.Error("Expected Size method to return Photo")
	}
	if result != photo {
		t.Error("Expected Size to return same Photo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if v.PhotoWidth != 800 || v.PhotoHeight != 600 {
			t.Error("Expected PhotoWidth and PhotoHeight to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}
}

func TestPhoto_Title(t *testing.T) {
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)
	result := photo.Title(testTitle)
	if result == nil {
		t.Error("Expected Title method to return Photo")
	}
	if result != photo {
		t.Error("Expected Title to return same Photo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if v.Title != testTitle.Std() {
			t.Error("Expected Title to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}
}

func TestPhoto_Description(t *testing.T) {
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)
	result := photo.Description(testDescription)
	if result == nil {
		t.Error("Expected Description method to return Photo")
	}
	if result != photo {
		t.Error("Expected Description to return same Photo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if v.Description != testDescription.Std() {
			t.Error("Expected Description to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}
}

func TestPhoto_Caption(t *testing.T) {
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)
	result := photo.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return Photo")
	}
	if result != photo {
		t.Error("Expected Caption to return same Photo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}
}

func TestPhoto_HTML(t *testing.T) {
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)
	result := photo.HTML()
	if result == nil {
		t.Error("Expected HTML method to return Photo")
	}
	if result != photo {
		t.Error("Expected HTML to return same Photo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}
}

func TestPhoto_Markdown(t *testing.T) {
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)

	result := photo.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return Photo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}
}

func TestPhoto_CaptionEntities(t *testing.T) {
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)
	entities := createTestEntities()

	result := photo.CaptionEntities(entities)
	if result == nil {
		t.Error("Expected CaptionEntities method to return Photo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if len(v.CaptionEntities) == 0 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}
}

func TestPhoto_ShowCaptionAboveMedia(t *testing.T) {
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)

	result := photo.ShowCaptionAboveMedia()
	if result == nil {
		t.Error("Expected ShowCaptionAboveMedia method to return Photo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if !v.ShowCaptionAboveMedia {
			t.Error("Expected ShowCaptionAboveMedia to be true")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}
}

func TestPhoto_InputMessageContent(t *testing.T) {
	photo := inline.NewPhoto(testID, testURL, testThumbnailURL)
	messageContent := createTestMessageContent()

	result := photo.InputMessageContent(messageContent)
	if result == nil {
		t.Error("Expected InputMessageContent method to return Photo")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if v.InputMessageContent == nil {
			t.Error("Expected InputMessageContent to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}
}

func TestPhoto_MethodChaining(t *testing.T) {
	result := inline.NewPhoto(testID, testURL, testThumbnailURL).
		Size(800, 600).
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
		t.Error("Expected chained Photo to build correctly")
	}

	if _, ok := built.(gotgbot.InlineQueryResultPhoto); !ok {
		t.Error("Expected result to be InlineQueryResultPhoto")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
