package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestDocument(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	document := input.Document(mediaFile)
	if document == nil {
		t.Error("Expected Document to be created")
	}
	if !assertMedia(document) {
		t.Error("Document should implement Media correctly")
	}
}

func TestDocument_Caption(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	document := input.Document(mediaFile)
	result := document.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return MediaDocument")
	}
	if result != document {
		t.Error("Expected Caption to return same MediaDocument instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaDocument); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaDocument")
	}
}

func TestDocument_HTML(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	document := input.Document(mediaFile)
	result := document.HTML()
	if result == nil {
		t.Error("Expected HTML method to return MediaDocument")
	}
	if result != document {
		t.Error("Expected HTML to return same MediaDocument instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaDocument); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InputMediaDocument")
	}
}

func TestDocument_Markdown(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	document := input.Document(mediaFile)
	result := document.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return MediaDocument")
	}
	if result != document {
		t.Error("Expected Markdown to return same MediaDocument instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaDocument); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InputMediaDocument")
	}
}

func TestDocument_Thumbnail(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	thumbnailFile := file.Input(testThumbnailURL).Ok()
	document := input.Document(mediaFile)

	result := document.Thumbnail(thumbnailFile)
	if result == nil {
		t.Error("Expected Thumbnail method to return MediaDocument")
	}
	if result != document {
		t.Error("Expected Thumbnail to return same MediaDocument instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaDocument); ok {
		if v.Thumbnail == nil {
			t.Error("Expected Thumbnail to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaDocument")
	}
}

func TestDocument_CaptionEntities(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	document := input.Document(mediaFile)
	entities := createTestEntities()

	result := document.CaptionEntities(entities)
	if result == nil {
		t.Error("Expected CaptionEntities method to return MediaDocument")
	}
	if result != document {
		t.Error("Expected CaptionEntities to return same MediaDocument instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaDocument); ok {
		if len(v.CaptionEntities) == 0 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaDocument")
	}
}

func TestDocument_DisableContentTypeDetection(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	document := input.Document(mediaFile)
	result := document.DisableContentTypeDetection()
	if result == nil {
		t.Error("Expected DisableContentTypeDetection method to return MediaDocument")
	}
	if result != document {
		t.Error("Expected DisableContentTypeDetection to return same MediaDocument instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputMediaDocument); ok {
		if !v.DisableContentTypeDetection {
			t.Error("Expected DisableContentTypeDetection to be set to true")
		}
	} else {
		t.Error("Expected result to be InputMediaDocument")
	}
}

func TestDocument_MethodChaining(t *testing.T) {
	mediaFile := file.Input(testURL).Ok()
	result := input.Document(mediaFile).
		Caption(testCaption).
		HTML().
		DisableContentTypeDetection()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Document to build correctly")
	}

	if _, ok := built.(gotgbot.InputMediaDocument); !ok {
		t.Error("Expected result to be InputMediaDocument")
	}

	if !assertMedia(result) {
		t.Error("Expected result to implement Media interface")
	}
}
