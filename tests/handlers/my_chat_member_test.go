package handlers_test

import (
	"testing"

	"github.com/enetx/tg/handlers"
)

func TestMyChatMemberHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.Any(MockHandler)

	if result == nil {
		t.Error("Any should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("Any should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_ChatID(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.ChatID(-1001234567890, MockHandler)

	if result == nil {
		t.Error("ChatID should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("ChatID should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.
		Any(MockHandler).
		ChatID(-1001234567890, MockHandler)

	if result != myChatMemberHandlers {
		t.Error("Chained methods should return the same MyChatMemberHandlers instance")
	}
}

func TestMyChatMemberHandlers_ZeroChatID(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.ChatID(0, MockHandler)

	if result == nil {
		t.Error("ChatID with zero ID should work")
	}
}

func TestMyChatMemberHandlers_NegativeChatID(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.ChatID(-987654321, MockHandler)

	if result == nil {
		t.Error("ChatID with negative ID should work")
	}
}

func TestMyChatMemberHandlers_LargeChatID(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	largeID := int64(9223372036854775807) // max int64
	result := myChatMemberHandlers.ChatID(largeID, MockHandler)

	if result == nil {
		t.Error("ChatID with large ID should work")
	}
}

func TestMyChatMemberHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}
