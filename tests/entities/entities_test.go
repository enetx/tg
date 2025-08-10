package entities_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	. "github.com/enetx/tg/entities"
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

func TestEntitiesMention(t *testing.T) {
	text := g.String("Hello @user")

	entities := New(text).
		Mention("@user", 123456789)

	result := entities.Std()

	if len(result) != 1 {
		t.Errorf("Expected 1 entity, got %d", len(result))
		return
	}

	if result[0].Type != "text_mention" {
		t.Errorf("Expected entity type 'text_mention', got '%s'", result[0].Type)
	}

	if result[0].User == nil {
		t.Error("Expected User to be set")
		return
	}

	if result[0].User.Id != 123456789 {
		t.Errorf("Expected User ID 123456789, got %d", result[0].User.Id)
	}
}

func TestEntitiesCustomEmoji(t *testing.T) {
	text := g.String("Hello 😀 world")

	entities := New(text).
		CustomEmoji("😀", "5789012345")

	result := entities.Std()

	if len(result) != 1 {
		t.Errorf("Expected 1 entity, got %d", len(result))
		return
	}

	if result[0].Type != "custom_emoji" {
		t.Errorf("Expected entity type 'custom_emoji', got '%s'", result[0].Type)
	}

	if result[0].CustomEmojiId != "5789012345" {
		t.Errorf("Expected CustomEmojiId '5789012345', got '%s'", result[0].CustomEmojiId)
	}
}

func TestEntitiesExpandableBlockquote(t *testing.T) {
	text := g.String("Regular text\nThis is an expandable quote")

	entities := New(text).
		ExpandableBlockquote("This is an expandable quote")

	result := entities.Std()

	if len(result) != 1 {
		t.Errorf("Expected 1 entity, got %d", len(result))
		return
	}

	if result[0].Type != "expandable_blockquote" {
		t.Errorf("Expected entity type 'expandable_blockquote', got '%s'", result[0].Type)
	}
}

func TestEntitiesAdd(t *testing.T) {
	text := g.String("Hello world")

	// Create a custom entity
	customEntity := gotgbot.MessageEntity{
		Type:   "mention",
		Offset: 0,
		Length: 5,
	}

	entities := New(text).
		Add(customEntity)

	result := entities.Std()

	if len(result) != 1 {
		t.Errorf("Expected 1 entity, got %d", len(result))
		return
	}

	if result[0].Type != "mention" {
		t.Errorf("Expected entity type 'mention', got '%s'", result[0].Type)
	}
	if result[0].Offset != 0 {
		t.Errorf("Expected offset 0, got %d", result[0].Offset)
	}
	if result[0].Length != 5 {
		t.Errorf("Expected length 5, got %d", result[0].Length)
	}
}

func TestEntitiesImport(t *testing.T) {
	text := g.String("Hello world test")

	// Create entities to import
	importEntities := []gotgbot.MessageEntity{
		{Type: "bold", Offset: 0, Length: 5},
		{Type: "italic", Offset: 6, Length: 5},
	}

	entities := New(text).
		Import(importEntities)

	result := entities.Std()

	if len(result) != 2 {
		t.Errorf("Expected 2 entities, got %d", len(result))
		return
	}

	if result[0].Type != "bold" {
		t.Errorf("Expected first entity type 'bold', got '%s'", result[0].Type)
	}
	if result[1].Type != "italic" {
		t.Errorf("Expected second entity type 'italic', got '%s'", result[1].Type)
	}

	// Test importing nil slice
	entities2 := New(text).Import(nil)
	result2 := entities2.Std()

	if len(result2) != 0 {
		t.Errorf("Expected 0 entities after importing nil, got %d", len(result2))
	}
}

func TestEntitiesClear(t *testing.T) {
	text := g.String("Hello world test")

	entities := New(text).
		Bold("Hello").
		Italic("world").
		Code("test")

	// Check entities were added
	if entities.Count() != 3 {
		t.Errorf("Expected 3 entities before clear, got %d", entities.Count())
	}

	// Clear entities
	entities.Clear()

	if entities.Count() != 0 {
		t.Errorf("Expected 0 entities after clear, got %d", entities.Count())
	}

	result := entities.Std()
	if len(result) != 0 {
		t.Errorf("Expected empty slice after clear, got %d entities", len(result))
	}
}

func TestEntitiesCount(t *testing.T) {
	text := g.String("Hello world test")

	entities := New(text)

	// Initially should be 0
	if entities.Count() != 0 {
		t.Errorf("Expected 0 entities initially, got %d", entities.Count())
	}

	// Add entities and check count
	entities.Bold("Hello")
	if entities.Count() != 1 {
		t.Errorf("Expected 1 entity after adding Bold, got %d", entities.Count())
	}

	entities.Italic("world")
	if entities.Count() != 2 {
		t.Errorf("Expected 2 entities after adding Italic, got %d", entities.Count())
	}

	entities.Code("test")
	if entities.Count() != 3 {
		t.Errorf("Expected 3 entities after adding Code, got %d", entities.Count())
	}
}
