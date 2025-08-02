package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestInlineQueryHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.Any(MockHandler)

	if result == nil {
		t.Error("Any should return InlineQueryHandlers")
	}

	// Should return self for chaining
	if result != inlineHandlers {
		t.Error("Any should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_FromUser(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.FromUser(987654321, MockHandler)

	if result == nil {
		t.Error("FromUser should return InlineQueryHandlers")
	}

	if result != inlineHandlers {
		t.Error("FromUser should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_Query(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.Query(g.String("test query"), MockHandler)

	if result == nil {
		t.Error("Query should return InlineQueryHandlers")
	}

	if result != inlineHandlers {
		t.Error("Query should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_QueryPrefix(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.QueryPrefix(g.String("search "), MockHandler)

	if result == nil {
		t.Error("QueryPrefix should return InlineQueryHandlers")
	}

	if result != inlineHandlers {
		t.Error("QueryPrefix should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_QuerySuffix(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.QuerySuffix(g.String(" help"), MockHandler)

	if result == nil {
		t.Error("QuerySuffix should return InlineQueryHandlers")
	}

	if result != inlineHandlers {
		t.Error("QuerySuffix should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_Location(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.Location(MockHandler)

	if result == nil {
		t.Error("Location should return InlineQueryHandlers")
	}

	if result != inlineHandlers {
		t.Error("Location should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_Sender(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.Sender(MockHandler)

	if result == nil {
		t.Error("Sender should return InlineQueryHandlers")
	}

	if result != inlineHandlers {
		t.Error("Sender should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_Private(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.Private(MockHandler)

	if result == nil {
		t.Error("Private should return InlineQueryHandlers")
	}

	if result != inlineHandlers {
		t.Error("Private should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_Group(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.Group(MockHandler)

	if result == nil {
		t.Error("Group should return InlineQueryHandlers")
	}

	if result != inlineHandlers {
		t.Error("Group should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_Supergroup(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.Supergroup(MockHandler)

	if result == nil {
		t.Error("Supergroup should return InlineQueryHandlers")
	}

	if result != inlineHandlers {
		t.Error("Supergroup should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_Channel(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.Channel(MockHandler)

	if result == nil {
		t.Error("Channel should return InlineQueryHandlers")
	}

	if result != inlineHandlers {
		t.Error("Channel should return the same InlineQueryHandlers instance for chaining")
	}
}

func TestInlineQueryHandlers_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	// Test method chaining
	result := inlineHandlers.
		Any(MockHandler).
		FromUser(12345, MockHandler).
		Query(g.String("test"), MockHandler).
		QueryPrefix(g.String("search"), MockHandler).
		Private(MockHandler)

	if result != inlineHandlers {
		t.Error("Chained methods should return the same InlineQueryHandlers instance")
	}
}

func TestInlineQueryHandlers_EmptyQuery(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.Query(g.String(""), MockHandler)

	if result == nil {
		t.Error("Query with empty string should work")
	}
}

func TestInlineQueryHandlers_EmptyPrefix(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.QueryPrefix(g.String(""), MockHandler)

	if result == nil {
		t.Error("QueryPrefix with empty string should work")
	}
}

func TestInlineQueryHandlers_EmptySuffix(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.QuerySuffix(g.String(""), MockHandler)

	if result == nil {
		t.Error("QuerySuffix with empty string should work")
	}
}

func TestInlineQueryHandlers_UnicodeQuery(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	unicodeQuery := g.String("🔍 поиск 검색 検索")
	result := inlineHandlers.Query(unicodeQuery, MockHandler)

	if result == nil {
		t.Error("Query with Unicode characters should work")
	}
}

func TestInlineQueryHandlers_UnicodePrefix(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	unicodePrefix := g.String("🔍 ")
	result := inlineHandlers.QueryPrefix(unicodePrefix, MockHandler)

	if result == nil {
		t.Error("QueryPrefix with Unicode characters should work")
	}
}

func TestInlineQueryHandlers_UnicodeSuffix(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	unicodeSuffix := g.String(" 🎯")
	result := inlineHandlers.QuerySuffix(unicodeSuffix, MockHandler)

	if result == nil {
		t.Error("QuerySuffix with Unicode characters should work")
	}
}

func TestInlineQueryHandlers_LongQuery(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	longQuery := g.String("this is a very long query string that exceeds normal length expectations and should still be handled properly by the inline query handler system without any issues")
	result := inlineHandlers.Query(longQuery, MockHandler)

	if result == nil {
		t.Error("Query with long string should work")
	}
}

func TestInlineQueryHandlers_SpecialCharacters(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	specialQuery := g.String("test@#$%^&*()_+-=[]{}|;':\",./<>?")

	tests := []struct {
		name string
		test func() *handlers.InlineQueryHandlers
	}{
		{"Query", func() *handlers.InlineQueryHandlers { return inlineHandlers.Query(specialQuery, MockHandler) }},
		{"QueryPrefix", func() *handlers.InlineQueryHandlers { return inlineHandlers.QueryPrefix(specialQuery, MockHandler) }},
		{"QuerySuffix", func() *handlers.InlineQueryHandlers { return inlineHandlers.QuerySuffix(specialQuery, MockHandler) }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.test()
			if result == nil {
				t.Errorf("%s with special characters should work", test.name)
			}
		})
	}
}

func TestInlineQueryHandlers_ZeroUserID(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.FromUser(0, MockHandler)

	if result == nil {
		t.Error("FromUser with zero ID should work")
	}
}

func TestInlineQueryHandlers_NegativeUserID(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	result := inlineHandlers.FromUser(-123456789, MockHandler)

	if result == nil {
		t.Error("FromUser with negative ID should work")
	}
}

func TestInlineQueryHandlers_LargeUserID(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	largeID := int64(9223372036854775807) // max int64
	result := inlineHandlers.FromUser(largeID, MockHandler)

	if result == nil {
		t.Error("FromUser with large ID should work")
	}
}

func TestInlineQueryHandlers_AllChatTypes(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	// Test all chat type handlers
	chatTypeTests := []struct {
		name string
		test func() *handlers.InlineQueryHandlers
	}{
		{"Sender", func() *handlers.InlineQueryHandlers { return inlineHandlers.Sender(MockHandler) }},
		{"Private", func() *handlers.InlineQueryHandlers { return inlineHandlers.Private(MockHandler) }},
		{"Group", func() *handlers.InlineQueryHandlers { return inlineHandlers.Group(MockHandler) }},
		{"Supergroup", func() *handlers.InlineQueryHandlers { return inlineHandlers.Supergroup(MockHandler) }},
		{"Channel", func() *handlers.InlineQueryHandlers { return inlineHandlers.Channel(MockHandler) }},
	}

	for _, test := range chatTypeTests {
		t.Run(test.name, func(t *testing.T) {
			result := test.test()
			if result == nil {
				t.Errorf("%s chat type handler should work", test.name)
			}
		})
	}
}

func TestInlineQueryHandlers_MultipleHandlers(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	// Register multiple handlers for different conditions
	inlineHandlers.
		Any(MockHandler).
		FromUser(123, MockHandler).
		FromUser(456, MockHandler).
		Query(g.String("help"), MockHandler).
		QueryPrefix(g.String("search"), MockHandler).
		QuerySuffix(g.String("?"), MockHandler).
		Location(MockHandler).
		Private(MockHandler).
		Group(MockHandler)

	// Should not cause any issues
}

func TestInlineQueryHandlers_WhitespaceQueries(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	whitespaceTests := []struct {
		name  string
		query g.String
	}{
		{"Single space", g.String(" ")},
		{"Multiple spaces", g.String("   ")},
		{"Tab", g.String("\t")},
		{"Newline", g.String("\n")},
		{"Mixed whitespace", g.String(" \t\n ")},
	}

	for _, test := range whitespaceTests {
		t.Run(test.name, func(t *testing.T) {
			result := inlineHandlers.Query(test.query, MockHandler)
			if result == nil {
				t.Errorf("Query with %s should work", test.name)
			}
		})
	}
}

func TestInlineQueryHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	// Test with nil handler - should not panic
	result := inlineHandlers.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}

func TestInlineQueryHandlers_ComplexChaining(t *testing.T) {
	bot := NewMockBot()
	inlineHandlers := &handlers.InlineQueryHandlers{Bot: bot}

	// Test complex method chaining with various conditions
	result := inlineHandlers.
		Any(MockHandler).
		FromUser(111, MockHandler).
		FromUser(222, MockHandler).
		Query(g.String("exact"), MockHandler).
		QueryPrefix(g.String("cmd_"), MockHandler).
		QuerySuffix(g.String("_end"), MockHandler).
		Location(MockHandler).
		Private(MockHandler).
		Group(MockHandler).
		Supergroup(MockHandler).
		Channel(MockHandler).
		Sender(MockHandler)

	if result != inlineHandlers {
		t.Error("Complex chaining should return the same InlineQueryHandlers instance")
	}
}
