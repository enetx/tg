package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestBusinessMessagesDeleted_Any(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	result := businessMessagesDeleted.Any(MockHandler)

	if result == nil {
		t.Error("Any should return BusinessMessagesDeleted")
	}

	if result != businessMessagesDeleted {
		t.Error("Any should return the same BusinessMessagesDeleted instance for chaining")
	}
}

func TestBusinessMessagesDeleted_ConnectionID(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	result := businessMessagesDeleted.ConnectionID(g.String("business_123"), MockHandler)

	if result == nil {
		t.Error("ConnectionID should return BusinessMessagesDeleted")
	}

	if result != businessMessagesDeleted {
		t.Error("ConnectionID should return the same BusinessMessagesDeleted instance for chaining")
	}
}

func TestBusinessMessagesDeleted_ChatID(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	result := businessMessagesDeleted.ChatID(-1001234567890, MockHandler)

	if result == nil {
		t.Error("ChatID should return BusinessMessagesDeleted")
	}

	if result != businessMessagesDeleted {
		t.Error("ChatID should return the same BusinessMessagesDeleted instance for chaining")
	}
}

func TestBusinessMessagesDeleted_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	result := businessMessagesDeleted.
		Any(MockHandler).
		ConnectionID(g.String("test_business"), MockHandler).
		ChatID(-1001234567890, MockHandler)

	if result != businessMessagesDeleted {
		t.Error("Chained methods should return the same BusinessMessagesDeleted instance")
	}
}

func TestBusinessMessagesDeleted_EmptyConnectionID(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	result := businessMessagesDeleted.ConnectionID(g.String(""), MockHandler)

	if result == nil {
		t.Error("ConnectionID with empty string should work")
	}
}

func TestBusinessMessagesDeleted_ZeroChatID(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	result := businessMessagesDeleted.ChatID(0, MockHandler)

	if result == nil {
		t.Error("ChatID with zero ID should work")
	}
}

func TestBusinessMessagesDeleted_NegativeChatID(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	result := businessMessagesDeleted.ChatID(-987654321, MockHandler)

	if result == nil {
		t.Error("ChatID with negative ID should work")
	}
}

func TestBusinessMessagesDeleted_LargeChatID(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	largeID := int64(9223372036854775807) // max int64
	result := businessMessagesDeleted.ChatID(largeID, MockHandler)

	if result == nil {
		t.Error("ChatID with large ID should work")
	}
}

func TestBusinessMessagesDeleted_SpecialCharacterConnectionID(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	specialID := g.String("business@#$%^&*()_+-=[]{}|;':\",./<>?123")
	result := businessMessagesDeleted.ConnectionID(specialID, MockHandler)

	if result == nil {
		t.Error("ConnectionID with special characters should work")
	}
}

func TestBusinessMessagesDeleted_UnicodeConnectionID(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	unicodeID := g.String("бизнес_123_🗑️_비즈니스")
	result := businessMessagesDeleted.ConnectionID(unicodeID, MockHandler)

	if result == nil {
		t.Error("ConnectionID with Unicode characters should work")
	}
}

func TestBusinessMessagesDeleted_LongConnectionID(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	longID := g.String(
		"very_long_business_connection_id_for_deleted_messages_that_exceeds_normal_expectations_and_contains_many_characters_123456789",
	)
	result := businessMessagesDeleted.ConnectionID(longID, MockHandler)

	if result == nil {
		t.Error("ConnectionID with long string should work")
	}
}

func TestBusinessMessagesDeleted_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	result := businessMessagesDeleted.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}
