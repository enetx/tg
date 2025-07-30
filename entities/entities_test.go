package entities

import (
	"testing"

	"github.com/enetx/g"
)

func TestEntitiesBasicFormatting(t *testing.T) {
	text := g.String("Hello world test")

	entities := New(text).
		Bold("Hello").
		Italic("world").
		Code("test")

	result := entities.Std()

	if len(result) != 3 {
		t.Errorf("Expected 3 entities, got %d", len(result))
	}

	// Test bold entity
	if result[0].Type != "bold" {
		t.Errorf("Expected first entity type 'bold', got '%s'", result[0].Type)
	}
	if result[0].Offset != 0 || result[0].Length != 5 {
		t.Errorf("Expected bold entity at offset 0 with length 5, got offset %d length %d",
			result[0].Offset, result[0].Length)
	}

	// Test italic entity
	if result[1].Type != "italic" {
		t.Errorf("Expected second entity type 'italic', got '%s'", result[1].Type)
	}
	if result[1].Offset != 6 || result[1].Length != 5 {
		t.Errorf("Expected italic entity at offset 6 with length 5, got offset %d length %d",
			result[1].Offset, result[1].Length)
	}

	// Test code entity
	if result[2].Type != "code" {
		t.Errorf("Expected third entity type 'code', got '%s'", result[2].Type)
	}
	if result[2].Offset != 12 || result[2].Length != 4 {
		t.Errorf("Expected code entity at offset 12 with length 4, got offset %d length %d",
			result[2].Offset, result[2].Length)
	}
}

func TestEntitiesURL(t *testing.T) {
	text := g.String("Click here to visit")

	entities := New(text).
		URL("here", "https://example.com")

	result := entities.Std()

	if len(result) != 1 {
		t.Errorf("Expected 1 entity, got %d", len(result))
	}

	if result[0].Type != "text_link" {
		t.Errorf("Expected entity type 'text_link', got '%s'", result[0].Type)
	}

	if result[0].Url != "https://example.com" {
		t.Errorf("Expected URL 'https://example.com', got '%s'", result[0].Url)
	}
}

func TestEntitiesPreFormatted(t *testing.T) {
	text := g.String("Check this code: func main() {}")
	codeText := g.String("func main() {}")

	entities := New(text).
		Pre(codeText, "go")

	result := entities.Std()

	if len(result) != 1 {
		t.Errorf("Expected 1 entity, got %d", len(result))
	}

	if result[0].Type != "pre" {
		t.Errorf("Expected entity type 'pre', got '%s'", result[0].Type)
	}

	if result[0].Language != "go" {
		t.Errorf("Expected language 'go', got '%s'", result[0].Language)
	}
}

func TestEntitiesSpoiler(t *testing.T) {
	text := g.String("This is a spoiler text")

	entities := New(text).
		Spoiler("spoiler")

	result := entities.Std()

	if len(result) != 1 {
		t.Errorf("Expected 1 entity, got %d", len(result))
	}

	if result[0].Type != "spoiler" {
		t.Errorf("Expected entity type 'spoiler', got '%s'", result[0].Type)
	}
}

func TestEntitiesBlockquote(t *testing.T) {
	text := g.String("Regular text\nThis is a quote")

	entities := New(text).
		Blockquote("This is a quote")

	result := entities.Std()

	if len(result) != 1 {
		t.Errorf("Expected 1 entity, got %d", len(result))
	}

	if result[0].Type != "blockquote" {
		t.Errorf("Expected entity type 'blockquote', got '%s'", result[0].Type)
	}
}

func TestEntitiesMultipleFormatting(t *testing.T) {
	text := g.String("Bold italic underline strikethrough")

	entities := New(text).
		Bold("Bold").
		Italic("italic").
		Underline("underline").
		Strikethrough("strikethrough")

	result := entities.Std()

	if len(result) != 4 {
		t.Errorf("Expected 4 entities, got %d", len(result))
	}

	expectedTypes := []string{"bold", "italic", "underline", "strikethrough"}
	for i, expectedType := range expectedTypes {
		if result[i].Type != expectedType {
			t.Errorf("Expected entity %d type '%s', got '%s'", i, expectedType, result[i].Type)
		}
	}
}

func TestEntitiesEmpty(t *testing.T) {
	text := g.String("No entities here")

	entities := New(text)
	result := entities.Std()

	if len(result) != 0 {
		t.Errorf("Expected 0 entities for empty formatter, got %d", len(result))
	}
}

func TestEntitiesNonExistentText(t *testing.T) {
	text := g.String("Hello world")

	entities := New(text).
		Bold("nonexistent")

	result := entities.Std()

	// Should not add entity for text that doesn't exist
	if len(result) != 0 {
		t.Errorf("Expected 0 entities for non-existent text, got %d", len(result))
	}
}
