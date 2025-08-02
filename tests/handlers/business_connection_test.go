package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestBusinessConnection_Any(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.Any(MockHandler)

	if result == nil {
		t.Error("Any should return BusinessConnection")
	}

	if result != businessConnection {
		t.Error("Any should return the same BusinessConnection instance for chaining")
	}
}

func TestBusinessConnection_ConnectionID(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.ConnectionID(g.String("business_123"), MockHandler)

	if result == nil {
		t.Error("ConnectionID should return BusinessConnection")
	}

	if result != businessConnection {
		t.Error("ConnectionID should return the same BusinessConnection instance for chaining")
	}
}

func TestBusinessConnection_FromUser(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.FromUser(987654321, MockHandler)

	if result == nil {
		t.Error("FromUser should return BusinessConnection")
	}

	if result != businessConnection {
		t.Error("FromUser should return the same BusinessConnection instance for chaining")
	}
}

func TestBusinessConnection_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.
		Any(MockHandler).
		ConnectionID(g.String("test_business"), MockHandler).
		FromUser(123456, MockHandler)

	if result != businessConnection {
		t.Error("Chained methods should return the same BusinessConnection instance")
	}
}

func TestBusinessConnection_EmptyID(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.ConnectionID(g.String(""), MockHandler)

	if result == nil {
		t.Error("ConnectionID with empty string should work")
	}
}

func TestBusinessConnection_ZeroUserID(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.FromUser(0, MockHandler)

	if result == nil {
		t.Error("FromUser with zero ID should work")
	}
}

func TestBusinessConnection_NegativeUserID(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.FromUser(-123456789, MockHandler)

	if result == nil {
		t.Error("FromUser with negative ID should work")
	}
}

func TestBusinessConnection_LargeUserID(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	largeID := int64(9223372036854775807) // max int64
	result := businessConnection.FromUser(largeID, MockHandler)

	if result == nil {
		t.Error("FromUser with large ID should work")
	}
}

func TestBusinessConnection_SpecialCharacterID(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	specialID := g.String("business@#$%^&*()_+-=[]{}|;':\",./<>?123")
	result := businessConnection.ConnectionID(specialID, MockHandler)

	if result == nil {
		t.Error("ConnectionID with special characters should work")
	}
}

func TestBusinessConnection_UnicodeID(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	unicodeID := g.String("бизнес_123_🏢_비즈니스")
	result := businessConnection.ConnectionID(unicodeID, MockHandler)

	if result == nil {
		t.Error("ConnectionID with Unicode characters should work")
	}
}

func TestBusinessConnection_LongID(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	longID := g.String(
		"very_long_business_connection_id_that_exceeds_normal_expectations_and_contains_many_characters_123456789",
	)
	result := businessConnection.ConnectionID(longID, MockHandler)

	if result == nil {
		t.Error("ConnectionID with long string should work")
	}
}

func TestBusinessConnection_Enabled(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.Enabled(MockHandler)

	if result == nil {
		t.Error("Enabled should return BusinessConnection")
	}

	if result != businessConnection {
		t.Error("Enabled should return the same BusinessConnection instance for chaining")
	}
}

func TestBusinessConnection_Disabled(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.Disabled(MockHandler)

	if result == nil {
		t.Error("Disabled should return BusinessConnection")
	}

	if result != businessConnection {
		t.Error("Disabled should return the same BusinessConnection instance for chaining")
	}
}

func TestBusinessConnection_FromUsername(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.FromUsername(g.String("testuser"), MockHandler)

	if result == nil {
		t.Error("FromUsername should return BusinessConnection")
	}

	if result != businessConnection {
		t.Error("FromUsername should return the same BusinessConnection instance for chaining")
	}
}

func TestBusinessConnection_CanReply(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.CanReply(MockHandler)

	if result == nil {
		t.Error("CanReply should return BusinessConnection")
	}

	if result != businessConnection {
		t.Error("CanReply should return the same BusinessConnection instance for chaining")
	}
}

func TestBusinessConnection_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	businessConnection := &handlers.BusinessConnection{Bot: bot}

	result := businessConnection.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}
