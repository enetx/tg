package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
	"github.com/enetx/tg/types/poll"
)

func TestPollHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	result := pollHandlers.Any(MockHandler)

	if result == nil {
		t.Error("Any should return PollHandlers")
	}

	if result != pollHandlers {
		t.Error("Any should return the same PollHandlers instance for chaining")
	}
}

func TestPollHandlers_ID(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	result := pollHandlers.ID(g.String("poll_123"), MockHandler)

	if result == nil {
		t.Error("ID should return PollHandlers")
	}

	if result != pollHandlers {
		t.Error("ID should return the same PollHandlers instance for chaining")
	}
}

func TestPollHandlers_Type(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	result := pollHandlers.Type(poll.Regular, MockHandler)

	if result == nil {
		t.Error("Type should return PollHandlers")
	}

	if result != pollHandlers {
		t.Error("Type should return the same PollHandlers instance for chaining")
	}
}

func TestPollHandlers_Regular(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	result := pollHandlers.Regular(MockHandler)

	if result == nil {
		t.Error("Regular should return PollHandlers")
	}

	if result != pollHandlers {
		t.Error("Regular should return the same PollHandlers instance for chaining")
	}
}

func TestPollHandlers_Quiz(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	result := pollHandlers.Quiz(MockHandler)

	if result == nil {
		t.Error("Quiz should return PollHandlers")
	}

	if result != pollHandlers {
		t.Error("Quiz should return the same PollHandlers instance for chaining")
	}
}

func TestPollHandlers_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	// Test method chaining
	result := pollHandlers.
		Any(MockHandler).
		ID(g.String("test_poll"), MockHandler).
		Regular(MockHandler).
		Quiz(MockHandler)

	if result != pollHandlers {
		t.Error("Chained methods should return the same PollHandlers instance")
	}
}

func TestPollHandlers_EmptyID(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	result := pollHandlers.ID(g.String(""), MockHandler)

	if result == nil {
		t.Error("ID with empty string should work")
	}
}

func TestPollHandlers_LongID(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	longID := g.String("very_long_poll_id_that_exceeds_normal_expectations_and_contains_many_characters_123456789")
	result := pollHandlers.ID(longID, MockHandler)

	if result == nil {
		t.Error("ID with long string should work")
	}
}

func TestPollHandlers_SpecialCharacterID(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	specialID := g.String("poll@#$%^&*()_+-=[]{}|;':\",./<>?123")
	result := pollHandlers.ID(specialID, MockHandler)

	if result == nil {
		t.Error("ID with special characters should work")
	}
}

func TestPollHandlers_UnicodeID(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	unicodeID := g.String("опрос_123_🗳️_투표")
	result := pollHandlers.ID(unicodeID, MockHandler)

	if result == nil {
		t.Error("ID with Unicode characters should work")
	}
}

func TestPollHandlers_AllPollTypes(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	// Test all poll types
	pollTypeTests := []struct {
		name string
		test func() *handlers.PollHandlers
	}{
		{"Regular", func() *handlers.PollHandlers { return pollHandlers.Regular(MockHandler) }},
		{"Quiz", func() *handlers.PollHandlers { return pollHandlers.Quiz(MockHandler) }},
		{"Type Regular", func() *handlers.PollHandlers { return pollHandlers.Type(poll.Regular, MockHandler) }},
		{"Type Quiz", func() *handlers.PollHandlers { return pollHandlers.Type(poll.Quiz, MockHandler) }},
	}

	for _, test := range pollTypeTests {
		t.Run(test.name, func(t *testing.T) {
			result := test.test()
			if result == nil {
				t.Errorf("%s poll type handler should work", test.name)
			}
		})
	}
}

func TestPollHandlers_MultipleHandlers(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	// Register multiple handlers for different conditions
	pollHandlers.
		Any(MockHandler).
		ID(g.String("poll1"), MockHandler).
		ID(g.String("poll2"), MockHandler).
		Regular(MockHandler).
		Quiz(MockHandler).
		Type(poll.Regular, MockHandler).
		Type(poll.Quiz, MockHandler)

	// Should not cause any issues
}

func TestPollHandlers_NumericID(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	numericID := g.String("123456789")
	result := pollHandlers.ID(numericID, MockHandler)

	if result == nil {
		t.Error("ID with numeric string should work")
	}
}

func TestPollHandlers_WhitespaceID(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	whitespaceTests := []struct {
		name string
		id   g.String
	}{
		{"Single space", g.String(" ")},
		{"Multiple spaces", g.String("   ")},
		{"Tab", g.String("\t")},
		{"Newline", g.String("\n")},
		{"Mixed whitespace", g.String(" \t\n ")},
	}

	for _, test := range whitespaceTests {
		t.Run(test.name, func(t *testing.T) {
			result := pollHandlers.ID(test.id, MockHandler)
			if result == nil {
				t.Errorf("ID with %s should work", test.name)
			}
		})
	}
}

func TestPollHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	// Test with nil handler - should not panic
	result := pollHandlers.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}

func TestPollHandlers_ComplexChaining(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	// Test complex method chaining
	result := pollHandlers.
		Any(MockHandler).
		ID(g.String("complex_poll_1"), MockHandler).
		ID(g.String("complex_poll_2"), MockHandler).
		Regular(MockHandler).
		Quiz(MockHandler).
		Type(poll.Regular, MockHandler).
		Type(poll.Quiz, MockHandler).
		Any(MockHandler)

	if result != pollHandlers {
		t.Error("Complex chaining should return the same PollHandlers instance")
	}
}

func TestPollHandlers_RepeatedMethodCalls(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	// Test repeated calls to the same method
	result := pollHandlers.
		Regular(MockHandler).
		Regular(MockHandler).
		Regular(MockHandler)

	if result != pollHandlers {
		t.Error("Repeated method calls should return the same PollHandlers instance")
	}
}

func TestPollHandlers_MixedIDFormats(t *testing.T) {
	bot := NewMockBot()
	pollHandlers := &handlers.PollHandlers{Bot: bot}

	// Test various ID formats
	idFormats := []g.String{
		g.String("poll_123"),
		g.String("POLL_ABC"),
		g.String("poll-with-dashes"),
		g.String("poll.with.dots"),
		g.String("poll:with:colons"),
		g.String("poll|with|pipes"),
		g.String("poll_with_underscores"),
		g.String("pollWithCamelCase"),
		g.String("PollWithPascalCase"),
	}

	for _, id := range idFormats {
		t.Run(id.Std(), func(t *testing.T) {
			result := pollHandlers.ID(id, MockHandler)
			if result == nil {
				t.Errorf("ID format %s should work", id.Std())
			}
		})
	}
}
