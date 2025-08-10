package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/handlers"
)

// Mock bot implementation for testing
type mockBot struct{}

func (m *mockBot) Raw() *gotgbot.Bot {
	return &gotgbot.Bot{}
}

func (m *mockBot) Dispatcher() *ext.Dispatcher {
	return &ext.Dispatcher{}
}

func (m *mockBot) Updater() *ext.Updater {
	return &ext.Updater{}
}

func TestNewHandlers(t *testing.T) {
	bot := &mockBot{}
	h := handlers.NewHandlers(bot)

	if h == nil {
		t.Fatal("NewHandlers should not return nil")
	}

	// Test that all handler fields are initialized
	if h.Message == nil {
		t.Error("Message handlers should be initialized")
	}
	if h.Callback == nil {
		t.Error("Callback handlers should be initialized")
	}
	if h.Inline == nil {
		t.Error("Inline handlers should be initialized")
	}
	if h.Poll == nil {
		t.Error("Poll handlers should be initialized")
	}
	if h.PollAnswer == nil {
		t.Error("PollAnswer handlers should be initialized")
	}
	if h.ChatMember == nil {
		t.Error("ChatMember handlers should be initialized")
	}
	if h.MyChatMember == nil {
		t.Error("MyChatMember handlers should be initialized")
	}
	if h.ChatJoinRequest == nil {
		t.Error("ChatJoinRequest handlers should be initialized")
	}
	if h.ChosenInlineResult == nil {
		t.Error("ChosenInlineResult handlers should be initialized")
	}
	if h.Shipping == nil {
		t.Error("Shipping handlers should be initialized")
	}
	if h.PreCheckout == nil {
		t.Error("PreCheckout handlers should be initialized")
	}
	if h.Reaction == nil {
		t.Error("Reaction handlers should be initialized")
	}
	if h.PaidMedia == nil {
		t.Error("PaidMedia handlers should be initialized")
	}
	if h.BusinessConnection == nil {
		t.Error("BusinessConnection handlers should be initialized")
	}
	if h.BusinessMessagesDeleted == nil {
		t.Error("BusinessMessagesDeleted handlers should be initialized")
	}
}

func TestHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	h := handlers.NewHandlers(bot)

	// Test that Any method returns the bot instance
	result := h.Any(MockHandler)

	if result == nil {
		t.Error("Any should return the bot instance")
	}

	if result != bot {
		t.Error("Any should return the same bot instance")
	}
}

func TestHandlers_Any_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	h := handlers.NewHandlers(bot)

	// Test that Any method works with nil handler
	result := h.Any(nil)

	if result == nil {
		t.Error("Any should return the bot instance even with nil handler")
	}

	if result != bot {
		t.Error("Any should return the same bot instance")
	}
}

func TestHandlers_Any_MultipleHandlers(t *testing.T) {
	bot := NewMockBot()
	h := handlers.NewHandlers(bot)

	// Test chaining multiple Any calls
	result1 := h.Any(MockHandler)
	result2 := h.Any(MockHandler)

	if result1 != result2 {
		t.Error("Multiple Any calls should return consistent bot instance")
	}

	if result1 != bot || result2 != bot {
		t.Error("Any should always return the same bot instance")
	}
}
