package preview_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	. "github.com/enetx/tg/preview"
)

func TestNew(t *testing.T) {
	preview := New()

	if preview == nil {
		t.Error("Expected New() to create a Preview instance")
	}

	opts := preview.Std()
	if opts == nil {
		t.Error("Expected preview options to be initialized")
	}

	// Test default values
	if opts.IsDisabled {
		t.Error("Expected IsDisabled to be false by default")
	}

	if opts.ShowAboveText {
		t.Error("Expected ShowAboveText to be false by default")
	}

	if opts.PreferLargeMedia {
		t.Error("Expected PreferLargeMedia to be false by default")
	}

	if opts.PreferSmallMedia {
		t.Error("Expected PreferSmallMedia to be false by default")
	}

	if opts.Url != "" {
		t.Error("Expected Url to be empty by default")
	}
}

func TestPreview_URL(t *testing.T) {
	preview := New()
	testURL := g.String("https://example.com")

	result := preview.URL(testURL)

	if result != preview {
		t.Error("Expected URL() to return Preview instance for chaining")
	}

	opts := preview.Std()
	if opts.Url != testURL.Std() {
		t.Errorf("Expected URL to be %s, got %s", testURL.Std(), opts.Url)
	}
}

func TestPreview_Disable(t *testing.T) {
	preview := New()

	result := preview.Disable()

	if result != preview {
		t.Error("Expected Disable() to return Preview instance for chaining")
	}

	opts := preview.Std()
	if !opts.IsDisabled {
		t.Error("Expected IsDisabled to be true after calling Disable()")
	}
}

func TestPreview_Above(t *testing.T) {
	preview := New()

	result := preview.Above()

	if result != preview {
		t.Error("Expected Above() to return Preview instance for chaining")
	}

	opts := preview.Std()
	if !opts.ShowAboveText {
		t.Error("Expected ShowAboveText to be true after calling Above()")
	}
}

func TestPreview_Large(t *testing.T) {
	preview := New()

	result := preview.Large()

	if result != preview {
		t.Error("Expected Large() to return Preview instance for chaining")
	}

	opts := preview.Std()
	if !opts.PreferLargeMedia {
		t.Error("Expected PreferLargeMedia to be true after calling Large()")
	}
}

func TestPreview_Small(t *testing.T) {
	preview := New()

	result := preview.Small()

	if result != preview {
		t.Error("Expected Small() to return Preview instance for chaining")
	}

	opts := preview.Std()
	if !opts.PreferSmallMedia {
		t.Error("Expected PreferSmallMedia to be true after calling Small()")
	}
}

func TestPreview_Import(t *testing.T) {
	preview := New()

	// Create source options
	src := &gotgbot.LinkPreviewOptions{
		IsDisabled:       true,
		Url:              "https://imported.com",
		ShowAboveText:    true,
		PreferLargeMedia: true,
		PreferSmallMedia: false,
	}

	result := preview.Import(src)

	if result != preview {
		t.Error("Expected Import() to return Preview instance for chaining")
	}

	opts := preview.Std()
	if opts.IsDisabled != src.IsDisabled {
		t.Error("Expected IsDisabled to be imported")
	}

	if opts.Url != src.Url {
		t.Error("Expected Url to be imported")
	}

	if opts.ShowAboveText != src.ShowAboveText {
		t.Error("Expected ShowAboveText to be imported")
	}

	if opts.PreferLargeMedia != src.PreferLargeMedia {
		t.Error("Expected PreferLargeMedia to be imported")
	}

	if opts.PreferSmallMedia != src.PreferSmallMedia {
		t.Error("Expected PreferSmallMedia to be imported")
	}
}

func TestPreview_ImportNil(t *testing.T) {
	preview := New()

	// Set some initial values
	preview.Disable()
	preview.URL(g.String("https://test.com"))

	// Import nil should not change anything
	result := preview.Import(nil)

	if result != preview {
		t.Error("Expected Import(nil) to return Preview instance for chaining")
	}

	opts := preview.Std()
	if !opts.IsDisabled {
		t.Error("Expected IsDisabled to remain true after importing nil")
	}

	if opts.Url != "https://test.com" {
		t.Error("Expected Url to remain unchanged after importing nil")
	}
}

func TestPreview_ChainedOperations(t *testing.T) {
	preview := New()

	// Test chaining multiple operations
	result := preview.
		URL(g.String("https://chained.com")).
		Above().
		Large().
		Disable()

	if result != preview {
		t.Error("Expected chained operations to return Preview instance")
	}

	opts := preview.Std()
	if opts.Url != "https://chained.com" {
		t.Error("Expected URL to be set in chained operations")
	}

	if !opts.ShowAboveText {
		t.Error("Expected ShowAboveText to be true in chained operations")
	}

	if !opts.PreferLargeMedia {
		t.Error("Expected PreferLargeMedia to be true in chained operations")
	}

	if !opts.IsDisabled {
		t.Error("Expected IsDisabled to be true in chained operations")
	}
}

func TestPreview_ConflictingMediaPreferences(t *testing.T) {
	preview := New()

	// Test setting both large and small preferences
	preview.Large().Small()

	opts := preview.Std()
	if !opts.PreferLargeMedia {
		t.Error("Expected PreferLargeMedia to remain true")
	}

	if !opts.PreferSmallMedia {
		t.Error("Expected PreferSmallMedia to be true")
	}

	// Test setting small then large
	preview2 := New()
	preview2.Small().Large()

	opts2 := preview2.Std()
	if !opts2.PreferSmallMedia {
		t.Error("Expected PreferSmallMedia to remain true")
	}

	if !opts2.PreferLargeMedia {
		t.Error("Expected PreferLargeMedia to be true")
	}
}

func TestPreview_MultipleURLChanges(t *testing.T) {
	preview := New()

	// Test changing URL multiple times
	preview.URL(g.String("https://first.com"))

	opts := preview.Std()
	if opts.Url != "https://first.com" {
		t.Error("Expected first URL to be set")
	}

	preview.URL(g.String("https://second.com"))

	opts = preview.Std()
	if opts.Url != "https://second.com" {
		t.Error("Expected URL to be updated to second value")
	}
}

func TestPreview_EmptyURL(t *testing.T) {
	preview := New()

	preview.URL(g.String(""))

	opts := preview.Std()
	if opts.Url != "" {
		t.Error("Expected empty URL to be set")
	}
}

func TestPreview_StdReturnsCorrectType(t *testing.T) {
	preview := New()

	opts := preview.Std()

	if opts == nil {
		t.Error("Expected Std() to return non-nil LinkPreviewOptions")
	}

	// Verify it's the correct type
	var _ *gotgbot.LinkPreviewOptions = opts

	// Verify we can modify the returned options (it should be the actual struct, not a copy)
	opts.IsDisabled = true

	opts2 := preview.Std()
	if !opts2.IsDisabled {
		t.Error("Expected modifications to returned options to persist")
	}
}
