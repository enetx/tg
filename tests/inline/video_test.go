package inline_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

func TestNewVideo(t *testing.T) {
	video := inline.NewVideo(testID, testURL, g.String("video/mp4"), testThumbnailURL, testTitle)
	if video == nil {
		t.Error("Expected Video to be created")
	}

	built := video.Build()
	if built == nil {
		t.Error("Expected Video to build correctly")
	}

	if built.GetType() != "video" {
		t.Error("Expected Video type to be 'video'")
	}
}

func TestVideo_Caption(t *testing.T) {
	video := inline.NewVideo(testID, testURL, g.String("video/mp4"), testThumbnailURL, testTitle)
	result := video.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return Video")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVideo); ok {
		if v.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVideo")
	}
}

func TestVideo_Description(t *testing.T) {
	video := inline.NewVideo(testID, testURL, g.String("video/mp4"), testThumbnailURL, testTitle)
	result := video.Description(testDescription)
	if result == nil {
		t.Error("Expected Description method to return Video")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVideo); ok {
		if v.Description != testDescription.Std() {
			t.Error("Expected Description to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVideo")
	}
}

func TestVideo_Duration(t *testing.T) {
	video := inline.NewVideo(testID, testURL, g.String("video/mp4"), testThumbnailURL, testTitle)
	result := video.Duration(120 * time.Second)
	if result == nil {
		t.Error("Expected Duration method to return Video")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVideo); ok {
		if v.VideoDuration != 120 {
			t.Error("Expected VideoDuration to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVideo")
	}
}

func TestVideo_Size(t *testing.T) {
	video := inline.NewVideo(testID, testURL, g.String("video/mp4"), testThumbnailURL, testTitle)
	result := video.Size(1920, 1080)
	if result == nil {
		t.Error("Expected Size method to return Video")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVideo); ok {
		if v.VideoWidth != 1920 || v.VideoHeight != 1080 {
			t.Error("Expected VideoWidth and VideoHeight to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVideo")
	}
}

func TestVideo_HTML(t *testing.T) {
	video := inline.NewVideo(testID, testURL, g.String("video/mp4"), testThumbnailURL, testTitle)
	result := video.HTML()
	if result == nil {
		t.Error("Expected HTML method to return Video")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVideo); ok {
		if v.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVideo")
	}
}

func TestVideo_Markdown(t *testing.T) {
	video := inline.NewVideo(testID, testURL, g.String("video/mp4"), testThumbnailURL, testTitle)

	result := video.Markdown()
	if result == nil {
		t.Error("Expected Markdown method to return Video")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVideo); ok {
		if v.ParseMode != "MarkdownV2" {
			t.Error("Expected ParseMode to be set to MarkdownV2")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVideo")
	}
}

func TestVideo_CaptionEntities(t *testing.T) {
	video := inline.NewVideo(testID, testURL, g.String("video/mp4"), testThumbnailURL, testTitle)
	entities := createTestEntities()

	result := video.CaptionEntities(entities)
	if result == nil {
		t.Error("Expected CaptionEntities method to return Video")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVideo); ok {
		if len(v.CaptionEntities) == 0 {
			t.Error("Expected CaptionEntities to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVideo")
	}
}

func TestVideo_ShowCaptionAboveMedia(t *testing.T) {
	video := inline.NewVideo(testID, testURL, g.String("video/mp4"), testThumbnailURL, testTitle)
	result := video.ShowCaptionAboveMedia()
	if result == nil {
		t.Error("Expected ShowCaptionAboveMedia method to return Video")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InlineQueryResultVideo); ok {
		if !v.ShowCaptionAboveMedia {
			t.Error("Expected ShowCaptionAboveMedia to be set to true")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultVideo")
	}
}

func TestVideo_MethodChaining(t *testing.T) {
	msg := createTestMessageContent()
	result := inline.NewVideo(testID, testURL, g.String("video/mp4"), testThumbnailURL, testTitle).
		Caption(testCaption).
		Description(testDescription).
		Duration(120*time.Second).
		Size(1920, 1080).
		HTML().
		ShowCaptionAboveMedia().
		Markup(createTestKeyboard()).
		InputMessageContent(msg)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Video to build correctly")
	}

	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
