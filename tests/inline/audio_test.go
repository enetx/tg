package inline_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/inline"
)

func TestNewAudio(t *testing.T) {
	audio := inline.NewAudio(testID, testURL, testTitle)

	if audio == nil {
		t.Error("Expected Audio to be created")
	}

	// Test that it builds correctly
	built := audio.Build()
	if built == nil {
		t.Error("Expected Audio to build correctly")
	}

	// Type assertion to access specific audio fields
	if audioResult, ok := built.(gotgbot.InlineQueryResultAudio); ok {
		if audioResult.GetType() != "audio" {
			t.Error("Expected Audio type to be 'audio'")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultAudio")
	}
}

func TestAudio_Caption(t *testing.T) {
	audio := inline.NewAudio(testID, testURL, testTitle)

	result := audio.Caption(testCaption)
	if result == nil {
		t.Error("Expected Caption method to return Audio")
	}

	// Test method chaining returns same instance
	if result != audio {
		t.Error("Expected Caption method to return same Audio instance")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultAudio); ok {
		if audioResult.Caption != testCaption.Std() {
			t.Error("Expected Caption to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultAudio")
	}
}

func TestAudio_HTML(t *testing.T) {
	audio := inline.NewAudio(testID, testURL, testTitle)

	result := audio.HTML()
	if result == nil {
		t.Error("Expected HTML method to return Audio")
	}

	// Test method chaining returns same instance
	if result != audio {
		t.Error("Expected HTML method to return same Audio instance")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultAudio); ok {
		if audioResult.ParseMode != "HTML" {
			t.Error("Expected ParseMode to be set to HTML")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultAudio")
	}
}

func TestAudio_Performer(t *testing.T) {
	audio := inline.NewAudio(testID, testURL, testTitle)
	performer := g.String("Test Artist")

	result := audio.Performer(performer)
	if result == nil {
		t.Error("Expected Performer method to return Audio")
	}

	// Test method chaining returns same instance
	if result != audio {
		t.Error("Expected Performer method to return same Audio instance")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultAudio); ok {
		if audioResult.Performer != performer.Std() {
			t.Error("Expected Performer to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultAudio")
	}
}

func TestAudio_Duration(t *testing.T) {
	audio := inline.NewAudio(testID, testURL, testTitle)
	duration := 180 * time.Second

	result := audio.Duration(duration)
	if result == nil {
		t.Error("Expected Duration method to return Audio")
	}

	// Test method chaining returns same instance
	if result != audio {
		t.Error("Expected Duration method to return same Audio instance")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultAudio); ok {
		if audioResult.AudioDuration != int64(duration.Seconds()) {
			t.Error("Expected Duration to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultAudio")
	}
}

func TestAudio_Markup(t *testing.T) {
	audio := inline.NewAudio(testID, testURL, testTitle)
	keyboard := createTestKeyboard()

	result := audio.Markup(keyboard)
	if result == nil {
		t.Error("Expected Markup method to return Audio")
	}

	// Test method chaining returns same instance
	if result != audio {
		t.Error("Expected Markup method to return same Audio instance")
	}

	built := result.Build()
	if audioResult, ok := built.(gotgbot.InlineQueryResultAudio); ok {
		if audioResult.ReplyMarkup == nil {
			t.Error("Expected ReplyMarkup to be set correctly")
		}
	} else {
		t.Error("Expected result to be InlineQueryResultAudio")
	}
}

func TestAudio_MethodChaining(t *testing.T) {
	performer := g.String("Test Artist")
	duration := 180 * time.Second

	result := inline.NewAudio(testID, testURL, testTitle).
		Caption(testCaption).
		HTML().
		Performer(performer).
		Duration(duration).
		Markup(createTestKeyboard())

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained Audio to build correctly")
	}

	// Verify the result implements QueryResult interface
	if !assertQueryResult(result) {
		t.Error("Expected result to implement QueryResult interface")
	}
}
