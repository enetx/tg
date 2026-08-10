package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestLinkMedia(t *testing.T) {
	link := input.LinkMedia(testURL)
	if link == nil {
		t.Error("Expected LinkMedia to be created")
	}

	if !assertPollOptionMedia(link) {
		t.Error("LinkMedia should implement PollOptionMedia correctly")
	}

	built := link.BuildPollOptionMedia()
	if v, ok := built.(gotgbot.InputMediaLink); ok {
		if v.Url != testURL.Std() {
			t.Error("Expected Url to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaLink")
	}
}

func TestLinkMedia_URL(t *testing.T) {
	link := input.LinkMedia(testURL)
	updated := g.String("https://example.com/other")

	result := link.URL(updated)
	if result != link {
		t.Error("Expected URL to return same MediaLink instance")
	}

	built := result.BuildPollOptionMedia()
	if v, ok := built.(gotgbot.InputMediaLink); ok {
		if v.Url != updated.Std() {
			t.Error("Expected URL to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaLink")
	}
}

func TestLinkMedia_Type(t *testing.T) {
	link := input.LinkMedia(testURL)
	if link.BuildPollOptionMedia().GetType() != gotgbot.InputPollOptionMediaTypeLink {
		t.Error("Expected media type to be link")
	}
}
