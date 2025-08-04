package effects_test

import (
	"testing"

	"github.com/enetx/tg/types/effects"
)

func TestEffectType_String(t *testing.T) {
	tests := []struct {
		name     string
		effect   effects.EffectType
		expected string
	}{
		{
			name:     "Fire effect",
			effect:   effects.Fire,
			expected: "5104841245755180586",
		},
		{
			name:     "ThumbsUp effect",
			effect:   effects.ThumbsUp,
			expected: "5107584321108051014",
		},
		{
			name:     "ThumbsDown effect",
			effect:   effects.ThumbsDown,
			expected: "5104858069142078462",
		},
		{
			name:     "Heart effect",
			effect:   effects.Heart,
			expected: "️5044134455711629726",
		},
		{
			name:     "Celebration effect",
			effect:   effects.Celebration,
			expected: "5046509860389126442",
		},
		{
			name:     "Poop effect",
			effect:   effects.Poop,
			expected: "5046589136895476101",
		},
		{
			name:     "Unknown effect",
			effect:   effects.EffectType(999),
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.effect.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestEffectType_AllEffectsHaveUniqueStrings(t *testing.T) {
	effects := []effects.EffectType{
		effects.Fire,
		effects.ThumbsUp,
		effects.ThumbsDown,
		effects.Heart,
		effects.Celebration,
		effects.Poop,
	}

	seen := make(map[string]bool)

	for _, effect := range effects {
		str := effect.String()
		if seen[str] {
			t.Errorf("Duplicate effect string found: %s", str)
		}
		seen[str] = true
	}
}

func TestEffectType_ValidEffectsNotUnknown(t *testing.T) {
	effects := []effects.EffectType{
		effects.Fire,
		effects.ThumbsUp,
		effects.ThumbsDown,
		effects.Heart,
		effects.Celebration,
		effects.Poop,
	}

	for _, effect := range effects {
		str := effect.String()
		if str == "unknown" {
			t.Errorf("Valid effect %d should not return 'unknown'", effect)
		}
		if str == "" {
			t.Errorf("Valid effect %d should not return empty string", effect)
		}
	}
}

func TestEffectType_Constants(t *testing.T) {
	// Test that constants have expected values
	if effects.Fire != 0 {
		t.Errorf("Expected Fire to be 0, got %d", effects.Fire)
	}
	if effects.ThumbsUp != 1 {
		t.Errorf("Expected ThumbsUp to be 1, got %d", effects.ThumbsUp)
	}
	if effects.ThumbsDown != 2 {
		t.Errorf("Expected ThumbsDown to be 2, got %d", effects.ThumbsDown)
	}
	if effects.Heart != 3 {
		t.Errorf("Expected Heart to be 3, got %d", effects.Heart)
	}
	if effects.Celebration != 4 {
		t.Errorf("Expected Celebration to be 4, got %d", effects.Celebration)
	}
	if effects.Poop != 5 {
		t.Errorf("Expected Poop to be 5, got %d", effects.Poop)
	}
}
