package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestChosenInlineResultHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.Any(MockHandler)

	if result == nil {
		t.Error("Any should return ChosenInlineResultHandlers")
	}

	if result != chosenInlineResultHandlers {
		t.Error("Any should return the same ChosenInlineResultHandlers instance for chaining")
	}
}

func TestChosenInlineResultHandlers_FromUser(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.FromUser(987654321, MockHandler)

	if result == nil {
		t.Error("FromUser should return ChosenInlineResultHandlers")
	}

	if result != chosenInlineResultHandlers {
		t.Error("FromUser should return the same ChosenInlineResultHandlers instance for chaining")
	}
}

func TestChosenInlineResultHandlers_InlineMessage(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.InlineMessage(g.String("result_123"), MockHandler)

	if result == nil {
		t.Error("InlineMessage should return ChosenInlineResultHandlers")
	}

	if result != chosenInlineResultHandlers {
		t.Error("InlineMessage should return the same ChosenInlineResultHandlers instance for chaining")
	}
}

func TestChosenInlineResultHandlers_Query(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.Query(g.String("search query"), MockHandler)

	if result == nil {
		t.Error("Query should return ChosenInlineResultHandlers")
	}

	if result != chosenInlineResultHandlers {
		t.Error("Query should return the same ChosenInlineResultHandlers instance for chaining")
	}
}

func TestChosenInlineResultHandlers_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.
		Any(MockHandler).
		FromUser(123456, MockHandler).
		InlineMessage(g.String("test_result"), MockHandler).
		Query(g.String("test query"), MockHandler)

	if result != chosenInlineResultHandlers {
		t.Error("Chained methods should return the same ChosenInlineResultHandlers instance")
	}
}

func TestChosenInlineResultHandlers_EmptyInlineMessage(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.InlineMessage(g.String(""), MockHandler)

	if result == nil {
		t.Error("InlineMessage with empty string should work")
	}
}

func TestChosenInlineResultHandlers_EmptyQuery(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.Query(g.String(""), MockHandler)

	if result == nil {
		t.Error("Query with empty string should work")
	}
}

func TestChosenInlineResultHandlers_ZeroUserID(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.FromUser(0, MockHandler)

	if result == nil {
		t.Error("FromUser with zero ID should work")
	}
}

func TestChosenInlineResultHandlers_UnicodeStrings(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	unicodeInlineMessage := g.String("результат_123_🎯")
	unicodeQuery := g.String("поиск 🔍 검색")

	result1 := chosenInlineResultHandlers.InlineMessage(unicodeInlineMessage, MockHandler)
	result2 := chosenInlineResultHandlers.Query(unicodeQuery, MockHandler)

	if result1 == nil {
		t.Error("InlineMessage with Unicode should work")
	}

	if result2 == nil {
		t.Error("Query with Unicode should work")
	}
}

func TestChosenInlineResultHandlers_QueryPrefix(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.QueryPrefix(g.String("start_"), MockHandler)

	if result == nil {
		t.Error("QueryPrefix should return ChosenInlineResultHandlers")
	}

	if result != chosenInlineResultHandlers {
		t.Error("QueryPrefix should return the same ChosenInlineResultHandlers instance for chaining")
	}
}

func TestChosenInlineResultHandlers_QuerySuffix(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.QuerySuffix(g.String("_end"), MockHandler)

	if result == nil {
		t.Error("QuerySuffix should return ChosenInlineResultHandlers")
	}

	if result != chosenInlineResultHandlers {
		t.Error("QuerySuffix should return the same ChosenInlineResultHandlers instance for chaining")
	}
}

func TestChosenInlineResultHandlers_Location(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.Location(MockHandler)

	if result == nil {
		t.Error("Location should return ChosenInlineResultHandlers")
	}

	if result != chosenInlineResultHandlers {
		t.Error("Location should return the same ChosenInlineResultHandlers instance for chaining")
	}
}

func TestChosenInlineResultHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	chosenInlineResultHandlers := &handlers.ChosenInlineResultHandlers{Bot: bot}

	result := chosenInlineResultHandlers.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}
