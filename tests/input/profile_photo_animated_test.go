package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestAnimatedPhoto(t *testing.T) {
	animationURL := g.String("https://example.com/animation.mp4")
	animatedPhoto := input.AnimatedPhoto(animationURL)
	if animatedPhoto == nil {
		t.Error("Expected ProfilePhotoAnimated to be created")
	}
	if !assertProfilePhoto(animatedPhoto) {
		t.Error("ProfilePhotoAnimated should implement ProfilePhoto correctly")
	}
}

func TestAnimatedPhoto_MainFrameTimestamp(t *testing.T) {
	animationURL := g.String("https://example.com/animation.mp4")
	animatedPhoto := input.AnimatedPhoto(animationURL)
	timestamp := 2.5
	result := animatedPhoto.MainFrameTimestamp(timestamp)
	if result == nil {
		t.Error("Expected MainFrameTimestamp method to return ProfilePhotoAnimated")
	}
	if result != animatedPhoto {
		t.Error("Expected MainFrameTimestamp to return same ProfilePhotoAnimated instance")
	}

	built := result.Build()
	if v, ok := built.(gotgbot.InputProfilePhotoAnimated); ok {
		if v.MainFrameTimestamp != timestamp {
			t.Errorf("Expected MainFrameTimestamp to be %f, got %f", timestamp, v.MainFrameTimestamp)
		}
	} else {
		t.Error("Expected result to be InputProfilePhotoAnimated")
	}
}

func TestAnimatedPhoto_Build(t *testing.T) {
	animationURL := g.String("https://example.com/animation.mp4")
	animatedPhoto := input.AnimatedPhoto(animationURL)
	built := animatedPhoto.Build()

	if built == nil {
		t.Error("Expected Build to return non-nil result")
	}

	if v, ok := built.(gotgbot.InputProfilePhotoAnimated); ok {
		if v.Animation != animationURL.Std() {
			t.Errorf("Expected Animation to be %s, got %s", animationURL.Std(), v.Animation)
		}
	} else {
		t.Error("Expected result to be InputProfilePhotoAnimated")
	}
}

func TestAnimatedPhoto_BuildReturnsCorrectType(t *testing.T) {
	animationURL := g.String("https://example.com/animation.mp4")
	animatedPhoto := input.AnimatedPhoto(animationURL)
	built := animatedPhoto.Build()

	// Verify that Build() returns the correct type
	if _, ok := any(built).(gotgbot.InputProfilePhotoAnimated); !ok {
		t.Error("Expected Build() to return gotgbot.InputProfilePhotoAnimated")
	}
}

func TestAnimatedPhoto_MethodChaining(t *testing.T) {
	animationURL := g.String("https://example.com/animation.mp4")
	result := input.AnimatedPhoto(animationURL).
		MainFrameTimestamp(2.5)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	if built == nil {
		t.Error("Expected chained AnimatedPhoto to build correctly")
	}

	if _, ok := built.(gotgbot.InputProfilePhotoAnimated); !ok {
		t.Error("Expected result to be InputProfilePhotoAnimated")
	}

	if !assertProfilePhoto(result) {
		t.Error("Expected result to implement ProfilePhoto interface")
	}
}

func TestAnimatedPhoto_EmptyURL(t *testing.T) {
	emptyURL := g.String("")
	animatedPhoto := input.AnimatedPhoto(emptyURL)
	if animatedPhoto == nil {
		t.Error("Expected ProfilePhotoAnimated to be created with empty URL")
	}

	built := animatedPhoto.Build()
	if v, ok := built.(gotgbot.InputProfilePhotoAnimated); ok {
		if v.Animation != "" {
			t.Errorf("Expected empty animation URL, got %s", v.Animation)
		}
	} else {
		t.Error("Expected result to be InputProfilePhotoAnimated")
	}
}
