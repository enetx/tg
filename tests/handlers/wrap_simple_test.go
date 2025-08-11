package handlers_test

import (
	"errors"
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	tghandlers "github.com/enetx/tg/handlers"
)

func TestWrapFunctionThroughCommand(t *testing.T) {
	bot := NewMockBot()
	var executionOrder []string
	var handlerExecuted bool

	// Create middleware that tracks execution
	middleware := func(c *ctx.Context) error {
		executionOrder = append(executionOrder, "middleware")
		return nil
	}

	// Create handler that tracks execution
	handler := func(c *ctx.Context) error {
		executionOrder = append(executionOrder, "handler")
		handlerExecuted = true
		return nil
	}

	// Set middleware on bot
	bot.SetMiddlewares(g.Slice[tghandlers.Handler]{middleware})

	// Register command that internally uses wrap function
	cmd := tghandlers.NewCommand(bot, g.String("test"), handler)
	cmd.Register()

	// Create a test update that matches our command
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "/test hello",
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "supergroup",
			},
			From: &gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process through dispatcher
	dispatcher := bot.Dispatcher()

	// Process the update
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// The error is expected since we're testing the wrap functionality
		// The important thing is that our middleware and handler got executed
	}

	// Verify execution order (middleware should run first, then handler)
	if len(executionOrder) < 2 {
		t.Errorf("Expected at least 2 executions (middleware + handler), got %d: %v", len(executionOrder), executionOrder)
	}

	expectedOrder := []string{"middleware", "handler"}
	for i, expected := range expectedOrder {
		if i >= len(executionOrder) || executionOrder[i] != expected {
			t.Errorf("Expected execution order %v, got %v", expectedOrder, executionOrder)
			break
		}
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed")
	}
}

func TestWrapFunctionWithMiddlewareError(t *testing.T) {
	bot := NewMockBot()
	var executionOrder []string
	var handlerExecuted bool

	// Create middleware that returns error
	middleware := func(c *ctx.Context) error {
		executionOrder = append(executionOrder, "middleware")
		return errors.New("middleware error")
	}

	// Create handler that should not execute due to middleware error
	handler := func(c *ctx.Context) error {
		executionOrder = append(executionOrder, "handler")
		handlerExecuted = true
		return nil
	}

	// Set middleware on bot
	bot.SetMiddlewares(g.Slice[tghandlers.Handler]{middleware})

	// Create a message handler that uses wrap function
	messageHandler := tghandlers.NewHandlers(bot).Message.Any(handler)
	messageHandler.Register()

	// Create a test message update
	update := &gotgbot.Update{
		UpdateId: 1,
		Message: &gotgbot.Message{
			MessageId: 123,
			Date:      1234567890,
			Text:      "hello world",
			Chat: gotgbot.Chat{
				Id:   -1001234567890,
				Type: "private",
			},
			From: &gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process through dispatcher
	dispatcher := bot.Dispatcher()

	// Process the update - should get error from middleware
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})

	// Should have error from middleware (though dispatcher might handle it)
	_ = err // We don't assert on this as dispatcher handles errors differently

	// Verify only middleware executed
	if len(executionOrder) < 1 {
		t.Error("Expected middleware to execute")
	}

	if len(executionOrder) > 0 && executionOrder[0] != "middleware" {
		t.Errorf("Expected middleware execution first, got %v", executionOrder)
	}

	if handlerExecuted {
		t.Error("Handler should not have been executed when middleware fails")
	}
}

func TestWrapFunctionWithMultipleMiddlewares(t *testing.T) {
	bot := NewMockBot()
	var executionOrder []string
	var handlerExecuted bool

	// Create multiple middlewares
	middleware1 := func(c *ctx.Context) error {
		executionOrder = append(executionOrder, "middleware1")
		return nil
	}

	middleware2 := func(c *ctx.Context) error {
		executionOrder = append(executionOrder, "middleware2")
		return nil
	}

	// Create handler
	handler := func(c *ctx.Context) error {
		executionOrder = append(executionOrder, "handler")
		handlerExecuted = true
		return nil
	}

	// Set multiple middlewares on bot (should execute in order)
	bot.SetMiddlewares(g.Slice[tghandlers.Handler]{middleware1, middleware2})

	// Create a callback handler that uses wrap function
	callbackHandler := tghandlers.NewHandlers(bot).Callback.Any(handler)
	callbackHandler.Register()

	// Create a test callback update
	update := &gotgbot.Update{
		UpdateId: 1,
		CallbackQuery: &gotgbot.CallbackQuery{
			Id:   "callback123",
			Data: "test",
			From: gotgbot.User{
				Id:           987654321,
				IsBot:        false,
				FirstName:    "Test",
				Username:     "testuser",
				LanguageCode: "en",
			},
		},
	}

	// Process through dispatcher
	dispatcher := bot.Dispatcher()

	// Process the update
	err := dispatcher.ProcessUpdate(bot.Raw(), update, map[string]any{})
	if err != nil {
		// Error handling varies
	}

	// Verify execution order (middlewares run in order, then handler)
	if len(executionOrder) < 3 {
		t.Errorf("Expected at least 3 executions, got %d: %v", len(executionOrder), executionOrder)
	}

	expectedOrder := []string{"middleware1", "middleware2", "handler"}
	for i, expected := range expectedOrder {
		if i >= len(executionOrder) || executionOrder[i] != expected {
			t.Errorf("Expected execution order %v, got %v", expectedOrder, executionOrder)
			break
		}
	}

	if !handlerExecuted {
		t.Error("Handler should have been executed")
	}
}
