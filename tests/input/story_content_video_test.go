package input_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestStoryVideo(t *testing.T) {
	videoURL := g.String("https://example.com/video.mp4")
	storyVideo := input.StoryVideo(videoURL)
	if storyVideo == nil {
		t.Error("Expected StoryContentVideo to be created")
	}
	if !assertStoryContent(storyVideo) {
		t.Error("StoryContentVideo should implement StoryContent correctly")
	}
}

func TestStoryVideo_Duration(t *testing.T) {
	videoURL := g.String("https://example.com/video.mp4")
	storyVideo := input.StoryVideo(videoURL)
	duration := 30 * time.Second
	result := storyVideo.Duration(duration)
	if result == nil {
		t.Error("Expected Duration method to return StoryContentVideo")
	}
	if result != storyVideo {
		t.Error("Expected Duration to return same StoryContentVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputStoryContentVideo); ok {
		if v.Duration != 30.0 {
			t.Errorf("Expected Duration to be 30.0 seconds, got %f", v.Duration)
		}
	} else {
		t.Error("Expected result to be InputStoryContentVideo")
	}
}

func TestStoryVideo_CoverFrameTimestamp(t *testing.T) {
	videoURL := g.String("https://example.com/video.mp4")
	storyVideo := input.StoryVideo(videoURL)
	timestamp := 5 * time.Second
	result := storyVideo.CoverFrameTimestamp(timestamp)
	if result == nil {
		t.Error("Expected CoverFrameTimestamp method to return StoryContentVideo")
	}
	if result != storyVideo {
		t.Error("Expected CoverFrameTimestamp to return same StoryContentVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputStoryContentVideo); ok {
		if v.CoverFrameTimestamp != 5.0 {
			t.Errorf("Expected CoverFrameTimestamp to be 5.0 seconds, got %f", v.CoverFrameTimestamp)
		}
	} else {
		t.Error("Expected result to be InputStoryContentVideo")
	}
}

func TestStoryVideo_Animation(t *testing.T) {
	videoURL := g.String("https://example.com/video.mp4")
	storyVideo := input.StoryVideo(videoURL)
	result := storyVideo.Animation()
	if result == nil {
		t.Error("Expected Animation method to return StoryContentVideo")
	}
	if result != storyVideo {
		t.Error("Expected Animation to return same StoryContentVideo instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputStoryContentVideo); ok {
		if !v.IsAnimation {
			t.Error("Expected IsAnimation to be set to true")
		}
	} else {
		t.Error("Expected result to be InputStoryContentVideo")
	}
}

func TestStoryVideo_Build(t *testing.T) {
	videoURL := g.String("https://example.com/video.mp4")
	storyVideo := input.StoryVideo(videoURL)
	built := storyVideo.Build()

	if built == nil {
		t.Error("Expected Build to return non-nil result")
	}

	if v, ok := built.(gotgbot.InputStoryContentVideo); ok {
		if v.Video != videoURL.Std() {
			t.Errorf("Expected Video to be %s, got %s", videoURL.Std(), v.Video)
		}
	} else {
		t.Error("Expected result to be InputStoryContentVideo")
	}
}

func TestStoryVideo_BuildReturnsCorrectType(t *testing.T) {
	videoURL := g.String("https://example.com/video.mp4")
	storyVideo := input.StoryVideo(videoURL)
	built := storyVideo.Build()

	// Verify that Build() returns the correct type
	if _, ok := interface{}(built).(gotgbot.InputStoryContentVideo); !ok {
		t.Error("Expected Build() to return gotgbot.InputStoryContentVideo")
	}
}

func TestStoryVideo_MethodChaining(t *testing.T) {
	videoURL := g.String("https://example.com/video.mp4")
	result := input.StoryVideo(videoURL).
		Duration(30 * time.Second).
		CoverFrameTimestamp(5 * time.Second).
		Animation()

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained StoryVideo to build correctly")
	}

	if _, ok := built.(gotgbot.InputStoryContentVideo); !ok {
		t.Error("Expected result to be InputStoryContentVideo")
	}

	if !assertStoryContent(result) {
		t.Error("Expected result to implement StoryContent interface")
	}
}

func TestStoryVideo_EmptyURL(t *testing.T) {
	emptyURL := g.String("")
	storyVideo := input.StoryVideo(emptyURL)
	if storyVideo == nil {
		t.Error("Expected StoryContentVideo to be created with empty URL")
	}

	built := storyVideo.Build()
	if v, ok := built.(gotgbot.InputStoryContentVideo); ok {
		if v.Video != "" {
			t.Errorf("Expected empty video URL, got %s", v.Video)
		}
	} else {
		t.Error("Expected result to be InputStoryContentVideo")
	}
}
