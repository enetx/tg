package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
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

func TestBusinessMessagesDeleted_Private(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	result := businessMessagesDeleted.Private(MockHandler)

	if result == nil {
		t.Error("Private should return BusinessMessagesDeleted")
	}

	if result != businessMessagesDeleted {
		t.Error("Private should return the same BusinessMessagesDeleted instance for chaining")
	}
}

func TestBusinessMessagesDeleted_Group(t *testing.T) {
	bot := NewMockBot()
	businessMessagesDeleted := &handlers.BusinessMessagesDeleted{Bot: bot}

	result := businessMessagesDeleted.Group(MockHandler)

	if result == nil {
		t.Error("Group should return BusinessMessagesDeleted")
	}

	if result != businessMessagesDeleted {
		t.Error("Group should return the same BusinessMessagesDeleted instance for chaining")
	}
}

func TestBusinessMessagesDeletedHandler_CheckUpdate(t *testing.T) {
	handler := handlers.BusinessMessagesDeletedHandler{}

	// Test with nil DeletedBusinessMessages
	ctx := &ext.Context{
		Update: &gotgbot.Update{
			DeletedBusinessMessages: nil,
		},
	}

	if handler.CheckUpdate(nil, ctx) {
		t.Error("CheckUpdate should return false for nil DeletedBusinessMessages")
	}

	// Test with valid DeletedBusinessMessages and nil filter
	ctx.Update.DeletedBusinessMessages = &gotgbot.BusinessMessagesDeleted{
		BusinessConnectionId: "test_connection",
		Chat: gotgbot.Chat{
			Id:   123,
			Type: "private",
		},
	}

	if !handler.CheckUpdate(nil, ctx) {
		t.Error("CheckUpdate should return true for valid DeletedBusinessMessages with nil filter")
	}

	// Test with filter that returns true
	handler.Filter = func(d *gotgbot.BusinessMessagesDeleted) bool {
		return d.BusinessConnectionId == "test_connection"
	}

	if !handler.CheckUpdate(nil, ctx) {
		t.Error("CheckUpdate should return true when filter matches")
	}

	// Test with filter that returns false
	handler.Filter = func(d *gotgbot.BusinessMessagesDeleted) bool {
		return d.BusinessConnectionId == "different_connection"
	}

	if handler.CheckUpdate(nil, ctx) {
		t.Error("CheckUpdate should return false when filter does not match")
	}
}

func TestBusinessMessagesDeletedHandler_HandleUpdate(t *testing.T) {
	called := false
	handler := handlers.BusinessMessagesDeletedHandler{
		Response: func(b *gotgbot.Bot, ctx *ext.Context) error {
			called = true
			return nil
		},
	}

	ctx := &ext.Context{}

	err := handler.HandleUpdate(nil, ctx)

	if err != nil {
		t.Errorf("HandleUpdate should not return error: %v", err)
	}

	if !called {
		t.Error("HandleUpdate should call the response function")
	}
}

func TestBusinessMessagesDeletedHandler_Name(t *testing.T) {
	handler := handlers.BusinessMessagesDeletedHandler{
		Response: func(b *gotgbot.Bot, ctx *ext.Context) error { return nil },
	}

	name := handler.Name()

	if name == "" {
		t.Error("Name should return non-empty string")
	}

	if len(name) < 10 {
		t.Errorf("Name should be descriptive, got: %s", name)
	}

	// Test that two different handlers have different names
	handler2 := handlers.BusinessMessagesDeletedHandler{
		Response: func(b *gotgbot.Bot, ctx *ext.Context) error { return nil },
	}

	name2 := handler2.Name()

	if name == name2 {
		t.Error("Different handlers should have different names")
	}
}
