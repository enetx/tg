package handlers_test

import (
	"testing"

	"github.com/enetx/tg/handlers"
)

func TestManagedBot_Any(t *testing.T) {
	bot := NewMockBot()
	managedBot := &handlers.ManagedBot{Bot: bot}

	result := managedBot.Any(MockHandler)

	if result == nil {
		t.Error("Any should return ManagedBot")
	}

	if result != managedBot {
		t.Error("Any should return the same ManagedBot instance for chaining")
	}
}

func TestManagedBot_OwnedByUserID(t *testing.T) {
	bot := NewMockBot()
	managedBot := &handlers.ManagedBot{Bot: bot}

	result := managedBot.OwnedByUserID(987654321, MockHandler)

	if result == nil {
		t.Error("OwnedByUserID should return ManagedBot")
	}

	if result != managedBot {
		t.Error("OwnedByUserID should return the same ManagedBot instance for chaining")
	}
}

func TestManagedBot_AboutBotID(t *testing.T) {
	bot := NewMockBot()
	managedBot := &handlers.ManagedBot{Bot: bot}

	result := managedBot.AboutBotID(7000000001, MockHandler)

	if result == nil {
		t.Error("AboutBotID should return ManagedBot")
	}

	if result != managedBot {
		t.Error("AboutBotID should return the same ManagedBot instance for chaining")
	}
}

func TestManagedBot_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	managedBot := &handlers.ManagedBot{Bot: bot}

	result := managedBot.
		Any(MockHandler).
		OwnedByUserID(111, MockHandler).
		AboutBotID(222, MockHandler)

	if result != managedBot {
		t.Error("Chained methods should return the same ManagedBot instance")
	}
}

func TestManagedBot_ZeroIDs(t *testing.T) {
	bot := NewMockBot()
	managedBot := &handlers.ManagedBot{Bot: bot}

	if managedBot.OwnedByUserID(0, MockHandler) == nil {
		t.Error("OwnedByUserID with zero ID should still register and return ManagedBot")
	}

	if managedBot.AboutBotID(0, MockHandler) == nil {
		t.Error("AboutBotID with zero ID should still register and return ManagedBot")
	}
}

func TestManagedBot_NegativeIDs(t *testing.T) {
	bot := NewMockBot()
	managedBot := &handlers.ManagedBot{Bot: bot}

	if managedBot.OwnedByUserID(-1, MockHandler) == nil {
		t.Error("OwnedByUserID with negative ID should still register")
	}

	if managedBot.AboutBotID(-1, MockHandler) == nil {
		t.Error("AboutBotID with negative ID should still register")
	}
}

func TestManagedBot_LargeIDs(t *testing.T) {
	bot := NewMockBot()
	managedBot := &handlers.ManagedBot{Bot: bot}

	largeID := int64(9223372036854775807)
	if managedBot.OwnedByUserID(largeID, MockHandler) == nil {
		t.Error("OwnedByUserID with max int64 should register")
	}

	if managedBot.AboutBotID(largeID, MockHandler) == nil {
		t.Error("AboutBotID with max int64 should register")
	}
}

func TestManagedBot_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	managedBot := &handlers.ManagedBot{Bot: bot}

	if managedBot.Any(nil) == nil {
		t.Error("Any with nil handler should still return ManagedBot")
	}
}
