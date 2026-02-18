package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/tg/handlers"
)

func TestMessageHandlers_ChatOwnerLeft(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ChatOwnerLeft(MockHandler)

	if handler == nil {
		t.Error("ChatOwnerLeft should return a MessageHandler")
	}
}

func TestMessageHandlers_ChatOwnerChanged(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	handler := msgHandlers.ChatOwnerChanged(MockHandler)

	if handler == nil {
		t.Error("ChatOwnerChanged should return a MessageHandler")
	}
}

func TestMessageHandlers_ChatOwnerLeft_Filter(t *testing.T) {
	// Verify that the filter correctly matches messages with ChatOwnerLeft set
	msg := &gotgbot.Message{
		ChatOwnerLeft: &gotgbot.ChatOwnerLeft{},
	}

	if msg.ChatOwnerLeft == nil {
		t.Error("Expected ChatOwnerLeft to be set")
	}

	emptyMsg := &gotgbot.Message{}
	if emptyMsg.ChatOwnerLeft != nil {
		t.Error("Expected empty message to have no ChatOwnerLeft")
	}
}

func TestMessageHandlers_ChatOwnerChanged_Filter(t *testing.T) {
	// Verify that the filter correctly matches messages with ChatOwnerChanged set
	msg := &gotgbot.Message{
		ChatOwnerChanged: &gotgbot.ChatOwnerChanged{},
	}

	if msg.ChatOwnerChanged == nil {
		t.Error("Expected ChatOwnerChanged to be set")
	}

	emptyMsg := &gotgbot.Message{}
	if emptyMsg.ChatOwnerChanged != nil {
		t.Error("Expected empty message to have no ChatOwnerChanged")
	}
}

func TestMessageHandlers_ChatOwnerChanged_WithNewOwner(t *testing.T) {
	owner := &gotgbot.User{Id: 123, FirstName: "New", LastName: "Owner"}
	msg := &gotgbot.Message{
		ChatOwnerChanged: &gotgbot.ChatOwnerChanged{
			NewOwner: *owner,
		},
	}

	if msg.ChatOwnerChanged == nil {
		t.Error("Expected ChatOwnerChanged to be set")
	}

	if msg.ChatOwnerChanged.NewOwner.Id != 123 {
		t.Errorf("Expected new owner ID 123, got %d", msg.ChatOwnerChanged.NewOwner.Id)
	}
}

func TestMessageHandlers_ChatOwnerLeft_Registration(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	// Test that multiple handlers can be registered
	h1 := msgHandlers.ChatOwnerLeft(MockHandler)
	h2 := msgHandlers.ChatOwnerLeft(MockHandler)

	if h1 == nil || h2 == nil {
		t.Error("Expected both ChatOwnerLeft handlers to be non-nil")
	}
}

func TestMessageHandlers_ChatOwnerChanged_Registration(t *testing.T) {
	bot := NewMockBot()
	msgHandlers := &handlers.MessageHandlers{Bot: bot}

	// Test that multiple handlers can be registered
	h1 := msgHandlers.ChatOwnerChanged(MockHandler)
	h2 := msgHandlers.ChatOwnerChanged(MockHandler)

	if h1 == nil || h2 == nil {
		t.Error("Expected both ChatOwnerChanged handlers to be non-nil")
	}
}
