package handlers_test

import (
	"testing"

	"github.com/enetx/tg/handlers"
	"github.com/enetx/tg/types/chatmember"
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

func TestChatMemberHandlers_StatusChange(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.StatusChange(chatmember.Member, chatmember.Administrator, MockHandler)

	if result == nil {
		t.Error("StatusChange should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("StatusChange should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_Joined(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.Joined(MockHandler)

	if result == nil {
		t.Error("Joined should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("Joined should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_Left(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.Left(MockHandler)

	if result == nil {
		t.Error("Left should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("Left should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_Banned(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.Banned(MockHandler)

	if result == nil {
		t.Error("Banned should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("Banned should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_Unbanned(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.Unbanned(MockHandler)

	if result == nil {
		t.Error("Unbanned should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("Unbanned should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_Restricted(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.Restricted(MockHandler)

	if result == nil {
		t.Error("Restricted should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("Restricted should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_Unrestricted(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.Unrestricted(MockHandler)

	if result == nil {
		t.Error("Unrestricted should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("Unrestricted should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_Promoted(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.Promoted(MockHandler)

	if result == nil {
		t.Error("Promoted should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("Promoted should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_Demoted(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.Demoted(MockHandler)

	if result == nil {
		t.Error("Demoted should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("Demoted should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_UserID(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.UserID(123456789, MockHandler)

	if result == nil {
		t.Error("UserID should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("UserID should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_NewStatus(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.NewStatus(chatmember.Administrator, MockHandler)

	if result == nil {
		t.Error("NewStatus should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("NewStatus should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_OldStatus(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.OldStatus(chatmember.Member, MockHandler)

	if result == nil {
		t.Error("OldStatus should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("OldStatus should return the same ChatMemberHandlers instance for chaining")
	}
}

func TestChatMemberHandlers_HasInviteLink(t *testing.T) {
	bot := NewMockBot()
	chatMemberHandlers := &handlers.ChatMemberHandlers{Bot: bot}

	result := chatMemberHandlers.HasInviteLink(MockHandler)

	if result == nil {
		t.Error("HasInviteLink should return ChatMemberHandlers")
	}

	if result != chatMemberHandlers {
		t.Error("HasInviteLink should return the same ChatMemberHandlers instance for chaining")
	}
}
