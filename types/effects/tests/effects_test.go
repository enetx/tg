package effects_test

import (
	"testing"

	. "github.com/enetx/tg/types/effects"
)

func TestEffectType_String(t *testing.T) {
	tests := []struct {
		effect   EffectType
		expected string
		name     string
	}{
		{Fire, "5104841245755180586", "Fire"},
		{ThumbsUp, "5107584321108051014", "ThumbsUp"},
		{ThumbsDown, "5104858069142078462", "ThumbsDown"},
		{Heart, "️5044134455711629726", "Heart"},
		{Celebration, "5046509860389126442", "Celebration"},
		{Poop, "5046589136895476101", "Poop"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.effect.String()
			if result != test.expected {
				t.Errorf("Expected %s effect to return '%s', got '%s'", test.name, test.expected, result)
			}
		})
	}
}

func TestEffectType_Constants(t *testing.T) {
	// Test that constants are defined and have expected values
	if int(Fire) != 0 {
		t.Errorf("Expected Fire to be 0, got %d", int(Fire))
	}
	if int(ThumbsUp) != 1 {
		t.Errorf("Expected ThumbsUp to be 1, got %d", int(ThumbsUp))
	}
	if int(ThumbsDown) != 2 {
		t.Errorf("Expected ThumbsDown to be 2, got %d", int(ThumbsDown))
	}
	if int(Heart) != 3 {
		t.Errorf("Expected Heart to be 3, got %d", int(Heart))
	}
	if int(Celebration) != 4 {
		t.Errorf("Expected Celebration to be 4, got %d", int(Celebration))
	}
	if int(Poop) != 5 {
		t.Errorf("Expected Poop to be 5, got %d", int(Poop))
	}
}

func TestEffectType_UnknownEffect(t *testing.T) {
	// Test unknown effect value returns "unknown"
	unknownEffect := EffectType(999)
	result := unknownEffect.String()
	if result != "unknown" {
		t.Errorf("Expected unknown effect to return 'unknown', got '%s'", result)
	}
}

func TestEffectType_ValidEffectIDs(t *testing.T) {
	// Test that all effect IDs are non-empty strings
	effects := []EffectType{Fire, ThumbsUp, ThumbsDown, Heart, Celebration, Poop}

	for _, effect := range effects {
		id := effect.String()
		if id == "" {
			t.Errorf("Expected effect %d to have non-empty ID", int(effect))
		}
		// Telegram effect IDs should be numeric strings
		if len(id) < 10 {
			t.Errorf("Expected effect %d ID to be at least 10 characters, got %d", int(effect), len(id))
		}
	}
}
