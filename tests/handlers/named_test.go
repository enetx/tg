package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

// namedHandler is a copy of the unexported type from handlers package for testing
type namedHandler struct {
	name string
	ext.Handler
}

// Name returns the name of the handler
func (n namedHandler) Name() string {
	return n.name
}

func TestNamedHandler(t *testing.T) {
	// Create a base handler (using Message handler as example)
	baseHandler := handlers.NewMessage(nil, func(b *gotgbot.Bot, ctx *ext.Context) error {
		return nil
	})

	// Create named handler
	name := "test_handler"
	handler := namedHandler{
		name:    name,
		Handler: baseHandler,
	}

	// Test Name method
	if handler.Name() != name {
		t.Errorf("Expected handler name %s, got %s", name, handler.Name())
	}

	// Test that it still implements Handler interface
	var _ ext.Handler = handler
}

func TestNamedHandlerWithEmptyName(t *testing.T) {
	// Test named handler with empty name
	baseHandler := handlers.NewMessage(nil, func(b *gotgbot.Bot, ctx *ext.Context) error {
		return nil
	})

	handler := namedHandler{
		name:    "",
		Handler: baseHandler,
	}

	if handler.Name() != "" {
		t.Errorf("Expected empty name, got %s", handler.Name())
	}
}

func TestNamedHandlerWithLongName(t *testing.T) {
	// Test named handler with long name
	longName := "this_is_a_very_long_handler_name_that_contains_many_characters_and_underscores"

	baseHandler := handlers.NewMessage(nil, func(b *gotgbot.Bot, ctx *ext.Context) error {
		return nil
	})

	handler := namedHandler{
		name:    longName,
		Handler: baseHandler,
	}

	if handler.Name() != longName {
		t.Errorf("Expected long name %s, got %s", longName, handler.Name())
	}
}

func TestNamedHandlerWithSpecialCharacters(t *testing.T) {
	// Test named handler with special characters in name
	specialName := "handler-with.special_chars@123"

	baseHandler := handlers.NewMessage(nil, func(b *gotgbot.Bot, ctx *ext.Context) error {
		return nil
	})

	handler := namedHandler{
		name:    specialName,
		Handler: baseHandler,
	}

	if handler.Name() != specialName {
		t.Errorf("Expected special name %s, got %s", specialName, handler.Name())
	}
}

func TestNamedHandlerImplementsNamedHandlerInterface(t *testing.T) {
	// Test that named handler implements the expected interface
	baseHandler := handlers.NewMessage(nil, func(b *gotgbot.Bot, ctx *ext.Context) error {
		return nil
	})

	handler := namedHandler{
		name:    "test",
		Handler: baseHandler,
	}

	// Check if it implements the Name() method
	namer, ok := interface{}(handler).(interface{ Name() string })
	if !ok {
		t.Error("Named handler should implement Name() string method")
	}

	if namer.Name() != "test" {
		t.Errorf("Expected name 'test', got %s", namer.Name())
	}
}

func TestNamedHandlerEmbedding(t *testing.T) {
	// Test that the embedded handler functionality is preserved
	executed := false

	baseHandler := &handlers.Message{
		Response: func(b *gotgbot.Bot, ctx *ext.Context) error {
			executed = true
			return nil
		},
	}

	handler := namedHandler{
		name:    "embedded_test",
		Handler: baseHandler,
	}

	// Create test update and context
	update := CreateTestUpdate()
	bot := &gotgbot.Bot{}
	extCtx := &ext.Context{Update: update}

	// Test that the embedded handler's methods are accessible
	if !handler.CheckUpdate(bot, extCtx) {
		t.Error("Named handler should delegate CheckUpdate to embedded handler")
	}

	err := handler.HandleUpdate(bot, extCtx)
	if err != nil {
		t.Errorf("Named handler should delegate HandleUpdate to embedded handler, got error: %v", err)
	}

	if !executed {
		t.Error("Embedded handler's Response should have been executed")
	}
}

func TestMultipleNamedHandlers(t *testing.T) {
	// Test multiple named handlers with different names
	names := []string{"handler1", "handler2", "handler3"}
	var testHandlers []namedHandler

	for _, name := range names {
		baseHandler := handlers.NewMessage(nil, func(b *gotgbot.Bot, ctx *ext.Context) error {
			return nil
		})

		handler := namedHandler{
			name:    name,
			Handler: baseHandler,
		}

		testHandlers = append(testHandlers, handler)
	}

	// Verify each handler has the correct name
	for i, handler := range testHandlers {
		if handler.Name() != names[i] {
			t.Errorf("Handler %d: expected name %s, got %s", i, names[i], handler.Name())
		}
	}
}

func TestNamedHandlerWithNilBaseHandler(t *testing.T) {
	// Test named handler with nil base handler
	handler := namedHandler{
		name:    "nil_base",
		Handler: nil,
	}

	// Name method should still work
	if handler.Name() != "nil_base" {
		t.Errorf("Expected name 'nil_base', got %s", handler.Name())
	}

	// But using the handler methods should panic (expected behavior)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when calling methods on nil embedded handler")
		}
	}()

	update := CreateTestUpdate()
	bot := &gotgbot.Bot{}
	extCtx := &ext.Context{Update: update}
	_ = handler.CheckUpdate(bot, extCtx)
}

func TestNamedHandlerNameImmutability(t *testing.T) {
	// Test that the name doesn't change after creation
	originalName := "immutable_name"

	baseHandler := handlers.NewMessage(nil, func(b *gotgbot.Bot, ctx *ext.Context) error {
		return nil
	})

	handler := namedHandler{
		name:    originalName,
		Handler: baseHandler,
	}

	// Get name multiple times
	name1 := handler.Name()
	name2 := handler.Name()
	name3 := handler.Name()

	if name1 != originalName || name2 != originalName || name3 != originalName {
		t.Error("Handler name should be immutable")
	}

	if name1 != name2 || name2 != name3 {
		t.Error("Handler name should be consistent across calls")
	}
}
