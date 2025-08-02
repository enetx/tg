package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestPollAnswerHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	result := pollAnswerHandlers.Any(MockHandler)

	if result == nil {
		t.Error("Any should return PollAnswerHandlers")
	}

	if result != pollAnswerHandlers {
		t.Error("Any should return the same PollAnswerHandlers instance for chaining")
	}
}

func TestPollAnswerHandlers_ID(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	result := pollAnswerHandlers.ID(g.String("poll_123"), MockHandler)

	if result == nil {
		t.Error("ID should return PollAnswerHandlers")
	}

	if result != pollAnswerHandlers {
		t.Error("ID should return the same PollAnswerHandlers instance for chaining")
	}
}

func TestPollAnswerHandlers_FromUserID(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	result := pollAnswerHandlers.FromUserID(987654321, MockHandler)

	if result == nil {
		t.Error("FromUserID should return PollAnswerHandlers")
	}

	if result != pollAnswerHandlers {
		t.Error("FromUserID should return the same PollAnswerHandlers instance for chaining")
	}
}

func TestPollAnswerHandlers_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	result := pollAnswerHandlers.
		Any(MockHandler).
		ID(g.String("test_poll"), MockHandler).
		FromUserID(123456, MockHandler)

	if result != pollAnswerHandlers {
		t.Error("Chained methods should return the same PollAnswerHandlers instance")
	}
}

func TestPollAnswerHandlers_EmptyID(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	result := pollAnswerHandlers.ID(g.String(""), MockHandler)

	if result == nil {
		t.Error("ID with empty string should work")
	}
}

func TestPollAnswerHandlers_ZeroUserID(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	result := pollAnswerHandlers.FromUserID(0, MockHandler)

	if result == nil {
		t.Error("FromUserID with zero ID should work")
	}
}

func TestPollAnswerHandlers_NegativeUserID(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	result := pollAnswerHandlers.FromUserID(-123456789, MockHandler)

	if result == nil {
		t.Error("FromUserID with negative ID should work")
	}
}

func TestPollAnswerHandlers_LargeUserID(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	largeID := int64(9223372036854775807) // max int64
	result := pollAnswerHandlers.FromUserID(largeID, MockHandler)

	if result == nil {
		t.Error("FromUserID with large ID should work")
	}
}

func TestPollAnswerHandlers_SpecialCharacterID(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	specialID := g.String("poll@#$%^&*()_+-=[]{}|;':\",./<>?123")
	result := pollAnswerHandlers.ID(specialID, MockHandler)

	if result == nil {
		t.Error("ID with special characters should work")
	}
}

func TestPollAnswerHandlers_UnicodeID(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	unicodeID := g.String("опрос_123_🗳️_투표")
	result := pollAnswerHandlers.ID(unicodeID, MockHandler)

	if result == nil {
		t.Error("ID with Unicode characters should work")
	}
}

func TestPollAnswerHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	pollAnswerHandlers := &handlers.PollAnswerHandlers{Bot: bot}

	result := pollAnswerHandlers.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}
