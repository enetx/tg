package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
)

func TestStaticPhoto(t *testing.T) {
	photoURL := g.String("https://example.com/photo.jpg")
	staticPhoto := input.StaticPhoto(photoURL)
	if staticPhoto == nil {
		t.Error("Expected ProfilePhotoStatic to be created")
	}
	if !assertProfilePhoto(staticPhoto) {
		t.Error("ProfilePhotoStatic should implement ProfilePhoto correctly")
	}
}

func TestStaticPhoto_Build(t *testing.T) {
	photoURL := g.String("https://example.com/photo.jpg")
	staticPhoto := input.StaticPhoto(photoURL)
	built := staticPhoto.Build()

	if built == nil {
		t.Error("Expected Build to return non-nil result")
	}

	if v, ok := built.(gotgbot.InputProfilePhotoStatic); ok {
		if v.Photo != photoURL.Std() {
			t.Errorf("Expected Photo to be %s, got %s", photoURL.Std(), v.Photo)
		}
	} else {
		t.Error("Expected result to be InputProfilePhotoStatic")
	}
}

func TestStaticPhoto_BuildReturnsCorrectType(t *testing.T) {
	photoURL := g.String("https://example.com/photo.jpg")
	staticPhoto := input.StaticPhoto(photoURL)
	built := staticPhoto.Build()

	// Verify that Build() returns the correct type
	if _, ok := interface{}(built).(gotgbot.InputProfilePhotoStatic); !ok {
		t.Error("Expected Build() to return gotgbot.InputProfilePhotoStatic")
	}
}

func TestStaticPhoto_EmptyURL(t *testing.T) {
	emptyURL := g.String("")
	staticPhoto := input.StaticPhoto(emptyURL)
	if staticPhoto == nil {
		t.Error("Expected ProfilePhotoStatic to be created with empty URL")
	}

	built := staticPhoto.Build()
	if v, ok := built.(gotgbot.InputProfilePhotoStatic); ok {
		if v.Photo != "" {
			t.Errorf("Expected empty photo URL, got %s", v.Photo)
		}
	} else {
		t.Error("Expected result to be InputProfilePhotoStatic")
	}
}

func TestStaticPhoto_MultipleBuildsSameResult(t *testing.T) {
	photoURL := g.String("https://example.com/photo.jpg")
	staticPhoto := input.StaticPhoto(photoURL)

	// Build multiple times to ensure consistency
	built1 := staticPhoto.Build()
	built2 := staticPhoto.Build()

	if built1 == nil || built2 == nil {
		t.Error("Expected builds to return non-nil results")
	}

	// Both should be the same type and have same content
	v1, ok1 := built1.(gotgbot.InputProfilePhotoStatic)
	v2, ok2 := built2.(gotgbot.InputProfilePhotoStatic)

	if !ok1 || !ok2 {
		t.Error("Expected both builds to return InputProfilePhotoStatic")
	}

	if v1.Photo != v2.Photo {
		t.Error("Expected multiple builds to return consistent results")
	}
}
