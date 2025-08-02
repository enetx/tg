package inline_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

func TestNewDocument(t *testing.T) {
	doc := inline.NewDocument(testID, testTitle, testURL, g.String("application/pdf"))

	if doc == nil {
		t.Error("Expected Document to be created")
	}

	built := doc.Build()
	if built == nil {
		t.Error("Expected Document to build correctly")
	}

	if result, ok := built.(gotgbot.InlineQueryResultDocument); ok {
		if result.GetType() != "document" {
			t.Error("Expected type to be 'document'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultDocument")
	}
}

func TestDocument_Caption(t *testing.T) {
	doc := inline.NewDocument(testID, testTitle, testURL, g.String("application/pdf"))

	result := doc.Caption(testCaption)
	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultDocument); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultDocument")
	}
}

func TestDocument_Description(t *testing.T) {
	doc := inline.NewDocument(testID, testTitle, testURL, g.String("application/pdf"))

	result := doc.Description(testDescription)
	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultDocument); ok {
		if v.Description != testDescription.Std() {
			t.Error("Expected Description to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultDocument")
	}
}

func TestDocument_ThumbnailURL(t *testing.T) {
	doc := inline.NewDocument(testID, testTitle, testURL, g.String("application/pdf"))

	result := doc.ThumbnailURL(testThumbnailURL)
	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultDocument); ok {
		if v.ThumbnailUrl != testThumbnailURL.Std() {
			t.Error("Expected ThumbnailUrl to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultDocument")
	}
}

func TestDocument_ThumbnailSize(t *testing.T) {
	doc := inline.NewDocument(testID, testTitle, testURL, g.String("application/pdf"))

	result := doc.ThumbnailSize(150, 150)
	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultDocument); ok {
		if v.ThumbnailWidth != 150 || v.ThumbnailHeight != 150 {
			t.Error("Expected ThumbnailWidth and ThumbnailHeight to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultDocument")
	}
}

func TestDocument_HTML(t *testing.T) {
	doc := inline.NewDocument(testID, testTitle, testURL, g.String("application/pdf"))

	result := doc.HTML()
	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultDocument); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultDocument")
	}
}

func TestDocument_MethodChaining(t *testing.T) {
	result := inline.NewDocument(testID, testTitle, testURL, g.String("application/pdf")).
		Caption(testCaption).
		Description(testDescription).
		ThumbnailURL(testThumbnailURL).
		ThumbnailSize(150, 150).
		HTML().
		Markup(createTestKeyboard()).
		InputMessageContent(createTestMessageContent())

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Document to build correctly")
	}

	if _, ok := built.(gotgbot.InlineQueryResultDocument); !ok {
		t.Error("Expected result to be InlineQueryResultDocument")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
