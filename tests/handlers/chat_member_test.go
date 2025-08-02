package handlers_test

import (
	"testing"

	"github.com/enetx/tg/handlers"
)

func TestChatMemberHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.Any(MockHandler)

	if result == nil {
		t.Error("Any should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("Any should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_FromUserID(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.FromUserID(987654321, MockHandler)

	if result == nil {
		t.Error("FromUserID should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("FromUserID should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_ChatID(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.ChatID(-1001234567890, MockHandler)

	if result == nil {
		t.Error("ChatID should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("ChatID should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.
		Any(MockHandler).
		FromUserID(123456, MockHandler).
		ChatID(-1001234567890, MockHandler)

	if result != chatMemberHandlers {
		t.Error("Chained methods should return the same ChatMemberHandlers instance")
	}
}

func TestChatMemberHandlers_ZeroUserID(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.FromUserID(0, MockHandler)

	if result == nil {
		t.Error("FromUserID with zero ID should work")
	}
}

func TestChatMemberHandlers_ZeroChatID(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.ChatID(0, MockHandler)

	if result == nil {
		t.Error("ChatID with zero ID should work")
	}
}

func TestChatMemberHandlers_NegativeIDs(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result1 := chatMemberHandlers.FromUserID(-123456789, MockHandler)
	result2 := chatMemberHandlers.ChatID(-987654321, MockHandler)

	if result1 == nil {
		t.Error("FromUserID with negative ID should work")
	}

	if result2 == nil {
		t.Error("ChatID with negative ID should work")
	}
}

func TestChatMemberHandlers_LargeIDs(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	largeID := int64(9223372036854775807) // max int64

	result1 := chatMemberHandlers.FromUserID(largeID, MockHandler)
	result2 := chatMemberHandlers.ChatID(largeID, MockHandler)

	if result1 == nil {
		t.Error("FromUserID with large ID should work")
	}

	if result2 == nil {
		t.Error("ChatID with large ID should work")
	}
}

func TestChatMemberHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}
