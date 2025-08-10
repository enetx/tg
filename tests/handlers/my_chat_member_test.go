package handlers_test

import (
	"testing"

	"github.com/enetx/tg/handlers"
	"github.com/enetx/tg/types/chatmember"
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

func TestMyChatMemberHandlers_StatusChange(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.StatusChange(chatmember.Member, chatmember.Administrator, MockHandler)

	if result == nil {
		t.Error("StatusChange should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("StatusChange should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_Joined(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.Joined(MockHandler)

	if result == nil {
		t.Error("Joined should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("Joined should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_Left(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.Left(MockHandler)

	if result == nil {
		t.Error("Left should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("Left should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_Banned(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.Banned(MockHandler)

	if result == nil {
		t.Error("Banned should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("Banned should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_Unbanned(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.Unbanned(MockHandler)

	if result == nil {
		t.Error("Unbanned should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("Unbanned should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_Restricted(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.Restricted(MockHandler)

	if result == nil {
		t.Error("Restricted should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("Restricted should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_Unrestricted(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.Unrestricted(MockHandler)

	if result == nil {
		t.Error("Unrestricted should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("Unrestricted should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_Promoted(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.Promoted(MockHandler)

	if result == nil {
		t.Error("Promoted should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("Promoted should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_Demoted(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.Demoted(MockHandler)

	if result == nil {
		t.Error("Demoted should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("Demoted should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_UserID(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.UserID(123456789, MockHandler)

	if result == nil {
		t.Error("UserID should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("UserID should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_FromUserID(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.FromUserID(987654321, MockHandler)

	if result == nil {
		t.Error("FromUserID should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("FromUserID should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_NewStatus(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.NewStatus(chatmember.Administrator, MockHandler)

	if result == nil {
		t.Error("NewStatus should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("NewStatus should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_OldStatus(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.OldStatus(chatmember.Member, MockHandler)

	if result == nil {
		t.Error("OldStatus should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("OldStatus should return the same MyChatMemberHandlers instance for chaining")
	}
}

func TestMyChatMemberHandlers_HasInviteLink(t *testing.T) {
	bot := NewMockBot()
	myChatMemberHandlers := &handlers.MyChatMemberHandlers{Bot: bot}

	result := myChatMemberHandlers.HasInviteLink(MockHandler)

	if result == nil {
		t.Error("HasInviteLink should return MyChatMemberHandlers")
	}

	if result != myChatMemberHandlers {
		t.Error("HasInviteLink should return the same MyChatMemberHandlers instance for chaining")
	}
}
