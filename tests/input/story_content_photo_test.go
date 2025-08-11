package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestStoryPhoto(t *testing.T) {
	photoURL := g.String("https://example.com/photo.jpg")
	storyPhoto := input.StoryPhoto(photoURL)
	if storyPhoto == nil {
		t.Error("Expected StoryContentPhoto to be created")
	}
	if !assertStoryContent(storyPhoto) {
		t.Error("StoryContentPhoto should implement StoryContent correctly")
	}
}

func TestStoryPhoto_Build(t *testing.T) {
	photoURL := g.String("https://example.com/photo.jpg")
	storyPhoto := input.StoryPhoto(photoURL)
	built := storyPhoto.Build()

	if built == nil {
		t.Error("Expected Build to return non-nil result")
	}

	if v, ok := built.(gotgbot.InputStoryContentPhoto); ok {
		if v.Photo != photoURL.Std() {
			t.Errorf("Expected Photo to be %s, got %s", photoURL.Std(), v.Photo)
		}
	} else {
		t.Error("Expected result to be InputStoryContentPhoto")
	}
}

func TestStoryPhoto_BuildReturnsCorrectType(t *testing.T) {
	photoURL := g.String("https://example.com/photo.jpg")
	storyPhoto := input.StoryPhoto(photoURL)
	built := storyPhoto.Build()

	// Verify that Build() returns the correct type
	if _, ok := any(built).(gotgbot.InputStoryContentPhoto); !ok {
		t.Error("Expected Build() to return gotgbot.InputStoryContentPhoto")
	}
}

func TestStoryPhoto_EmptyURL(t *testing.T) {
	emptyURL := g.String("")
	storyPhoto := input.StoryPhoto(emptyURL)
	if storyPhoto == nil {
		t.Error("Expected StoryContentPhoto to be created with empty URL")
	}

	built := storyPhoto.Build()
	if v, ok := built.(gotgbot.InputStoryContentPhoto); ok {
		if v.Photo != "" {
			t.Errorf("Expected empty photo URL, got %s", v.Photo)
		}
	} else {
		t.Error("Expected result to be InputStoryContentPhoto")
	}
}

func TestStoryPhoto_MultipleBuildsSameResult(t *testing.T) {
	photoURL := g.String("https://example.com/photo.jpg")
	storyPhoto := input.StoryPhoto(photoURL)

	// Build multiple times to ensure consistency
	built1 := storyPhoto.Build()
	built2 := storyPhoto.Build()

	if built1 == nil || built2 == nil {
		t.Error("Expected builds to return non-nil results")
	}

	// Both should be the same type and have same content
	v1, ok1 := built1.(gotgbot.InputStoryContentPhoto)
	v2, ok2 := built2.(gotgbot.InputStoryContentPhoto)

	if !ok1 || !ok2 {
		t.Error("Expected both builds to return InputStoryContentPhoto")
	}

	if v1.Photo != v2.Photo {
		t.Error("Expected multiple builds to return consistent results")
	}
}
