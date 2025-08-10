package handlers_test

import (
	"testing"

	"github.com/enetx/tg/handlers"
)

func TestChatJoinRequestHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	chatJoinRequestHandlers := &handlers.ChatJoinRequestHandlers{Bot: bot}

	result := chatJoinRequestHandlers.Any(MockHandler)

	if result == nil {
		t.Error("Any should return ChatJoinRequestHandlers")
	}

	if result != chatJoinRequestHandlers {
		t.Error("Any should return the same ChatJoinRequestHandlers instance for chaining")
	}
}

func TestChatJoinRequestHandlers_FromUserID(t *testing.T) {
	bot := NewMockBot()
	chatJoinRequestHandlers := &handlers.ChatJoinRequestHandlers{Bot: bot}

	result := chatJoinRequestHandlers.FromUserID(987654321, MockHandler)

	if result == nil {
		t.Error("FromUserID should return ChatJoinRequestHandlers")
	}

	if result != chatJoinRequestHandlers {
		t.Error("FromUserID should return the same ChatJoinRequestHandlers instance for chaining")
	}
}

func TestChatJoinRequestHandlers_ChatID(t *testing.T) {
	bot := NewMockBot()
	chatJoinRequestHandlers := &handlers.ChatJoinRequestHandlers{Bot: bot}

	result := chatJoinRequestHandlers.ChatID(-1001234567890, MockHandler)

	if result == nil {
		t.Error("ChatID should return ChatJoinRequestHandlers")
	}

	if result != chatJoinRequestHandlers {
		t.Error("ChatID should return the same ChatJoinRequestHandlers instance for chaining")
	}
}

func TestChatJoinRequestHandlers_HasInviteLink(t *testing.T) {
	bot := NewMockBot()
	chatJoinRequestHandlers := &handlers.ChatJoinRequestHandlers{Bot: bot}

	result := chatJoinRequestHandlers.HasInviteLink(MockHandler)

	if result == nil {
		t.Error("HasInviteLink should return ChatJoinRequestHandlers")
	}

	if result != chatJoinRequestHandlers {
		t.Error("HasInviteLink should return the same ChatJoinRequestHandlers instance for chaining")
	}
}

func TestChatJoinRequestHandlers_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	chatJoinRequestHandlers := &handlers.ChatJoinRequestHandlers{Bot: bot}

	result := chatJoinRequestHandlers.
		Any(MockHandler).
		FromUserID(123456, MockHandler).
		ChatID(-1001234567890, MockHandler)

	if result != chatJoinRequestHandlers {
		t.Error("Chained methods should return the same ChatJoinRequestHandlers instance")
	}
}

func TestChatJoinRequestHandlers_ZeroUserID(t *testing.T) {
	bot := NewMockBot()
	chatJoinRequestHandlers := &handlers.ChatJoinRequestHandlers{Bot: bot}

	result := chatJoinRequestHandlers.FromUserID(0, MockHandler)

	if result == nil {
		t.Error("FromUserID with zero ID should work")
	}
}

func TestChatJoinRequestHandlers_ZeroChatID(t *testing.T) {
	bot := NewMockBot()
	chatJoinRequestHandlers := &handlers.ChatJoinRequestHandlers{Bot: bot}

	result := chatJoinRequestHandlers.ChatID(0, MockHandler)

	if result == nil {
		t.Error("ChatID with zero ID should work")
	}
}

func TestChatJoinRequestHandlers_NegativeIDs(t *testing.T) {
	bot := NewMockBot()
	chatJoinRequestHandlers := &handlers.ChatJoinRequestHandlers{Bot: bot}

	result1 := chatJoinRequestHandlers.FromUserID(-123456789, MockHandler)
	result2 := chatJoinRequestHandlers.ChatID(-987654321, MockHandler)

	if result1 == nil {
		t.Error("FromUserID with negative ID should work")
	}

	if result2 == nil {
		t.Error("ChatID with negative ID should work")
	}
}

func TestChatJoinRequestHandlers_LargeIDs(t *testing.T) {
	bot := NewMockBot()
	chatJoinRequestHandlers := &handlers.ChatJoinRequestHandlers{Bot: bot}

	largeID := int64(9223372036854775807) // max int64

	result1 := chatJoinRequestHandlers.FromUserID(largeID, MockHandler)
	result2 := chatJoinRequestHandlers.ChatID(largeID, MockHandler)

	if result1 == nil {
		t.Error("FromUserID with large ID should work")
	}

	if result2 == nil {
		t.Error("ChatID with large ID should work")
	}
}

func TestChatJoinRequestHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	chatJoinRequestHandlers := &handlers.ChatJoinRequestHandlers{Bot: bot}

	result := chatJoinRequestHandlers.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}
