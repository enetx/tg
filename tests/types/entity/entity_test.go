package entity_test

import (
	"testing"

	"github.com/enetx/tg/types/entity"
)

func TestEntityType_String(t *testing.T) {
	tests := []struct {
		name       string
		entityType entity.EntityType
		expected   string
	}{
		// Plain-text entities
		{"Mention", entity.Mention, "mention"},
		{"Hashtag", entity.Hashtag, "hashtag"},
		{"Cashtag", entity.Cashtag, "cashtag"},
		{"BotCommand", entity.BotCommand, "bot_command"},
		{"URL", entity.URL, "url"},
		{"Email", entity.Email, "email"},
		{"PhoneNumber", entity.PhoneNumber, "phone_number"},

		// Text-style entities
		{"Bold", entity.Bold, "bold"},
		{"Italic", entity.Italic, "italic"},
		{"Underline", entity.Underline, "underline"},
		{"Strikethrough", entity.Strikethrough, "strikethrough"},
		{"Spoiler", entity.Spoiler, "spoiler"},

		// Quote / block entities
		{"Blockquote", entity.Blockquote, "blockquote"},
		{"ExpandableBlockquote", entity.ExpandableBlockquote, "expandable_blockquote"},

		// Code entities
		{"Code", entity.Code, "code"},
		{"Pre", entity.Pre, "pre"},

		// Link entities
		{"TextLink", entity.TextLink, "text_link"},
		{"TextMention", entity.TextMention, "text_mention"},

		// Special
		{"CustomEmoji", entity.CustomEmoji, "custom_emoji"},

		// Unknown
		{"Unknown", entity.EntityType(999), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.entityType.String()
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestEntityType_Constants(t *testing.T) {
	// Test a few key constants to ensure they have expected sequential values
	if entity.Mention != 0 {
		t.Errorf("Expected Mention to be 0, got %d", int(entity.Mention))
	}
	if entity.Hashtag != 1 {
		t.Errorf("Expected Hashtag to be 1, got %d", int(entity.Hashtag))
	}
	if entity.Bold != 7 {
		t.Errorf("Expected Bold to be 7, got %d", int(entity.Bold))
	}
	if entity.CustomEmoji != 18 {
		t.Errorf("Expected CustomEmoji to be 18, got %d", int(entity.CustomEmoji))
	}
}

func TestEntityType_Coverage(t *testing.T) {
	// Test all entity types to ensure complete coverage
	allTypes := []entity.EntityType{
		entity.Mention, entity.Hashtag, entity.Cashtag, entity.BotCommand,
		entity.URL, entity.Email, entity.PhoneNumber,
		entity.Bold, entity.Italic, entity.Underline, entity.Strikethrough, entity.Spoiler,
		entity.Blockquote, entity.ExpandableBlockquote,
		entity.Code, entity.Pre,
		entity.TextLink, entity.TextMention,
		entity.CustomEmoji,
	}

	for _, entityType := range allTypes {
		result := entityType.String()
		if result == "unknown" {
			t.Errorf("EntityType %d returned 'unknown', expected specific string", int(entityType))
		}
	}
}
