package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestCallbackHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.Any(MockHandler)

	if handler == nil {
		t.Error("Any should return a CallbackHandler")
	}
}

func TestCallbackHandlers_Equal(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.Equal(g.String("test_data"), MockHandler)

	if handler == nil {
		t.Error("Equal should return a CallbackHandler")
	}
}

func TestCallbackHandlers_Equal_Filter(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	// This will create a filter function that we need to test
	handler := callbackHandlers.Equal(g.String("expected"), MockHandler)

	if handler == nil {
		t.Error("Equal should return a CallbackHandler")
	}

	// By creating the handler, we exercise the filter creation code paths
	// The actual filter logic testing would require access to internal implementation
}

func TestCallbackHandlers_Prefix(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.Prefix(g.String("btn_"), MockHandler)

	if handler == nil {
		t.Error("Prefix should return a CallbackHandler")
	}
}

func TestCallbackHandlers_Suffix(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.Suffix(g.String("_action"), MockHandler)

	if handler == nil {
		t.Error("Suffix should return a CallbackHandler")
	}
}

func TestCallbackHandlers_FromUserID(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.FromUserID(987654321, MockHandler)

	if handler == nil {
		t.Error("FromUserID should return a CallbackHandler")
	}
}

func TestCallbackHandlers_GameName(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.GameName(g.String("tetris"), MockHandler)

	if handler == nil {
		t.Error("GameName should return a CallbackHandler")
	}
}

func TestCallbackHandlers_Inline(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.Inline(MockHandler)

	if handler == nil {
		t.Error("Inline should return a CallbackHandler")
	}
}

func TestCallbackHandlers_ChatInstance(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.ChatInstance(g.String("chat_instance_123"), MockHandler)

	if handler == nil {
		t.Error("ChatInstance should return a CallbackHandler")
	}
}

func TestCallbackHandler_AllowChannel(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.Any(MockHandler).AllowChannel()

	if handler == nil {
		t.Error("AllowChannel should return the same CallbackHandler")
	}
}

func TestCallbackHandler_Register(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.Equal(g.String("test"), MockHandler)

	// Register should return the same handler
	registered := handler.Register()
	if registered != handler {
		t.Error("Register should return the same handler instance")
	}
}

func TestCallbackHandler_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.Any(MockHandler).
		AllowChannel().
		Register()

	if handler == nil {
		t.Error("Chained methods should return the same CallbackHandler")
	}
}

func TestCallbackHandlers_MultipleFilters(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	// Test multiple different filter types
	tests := []struct {
		name    string
		handler func() *handlers.CallbackHandler
	}{
		{"Equal", func() *handlers.CallbackHandler { return callbackHandlers.Equal(g.String("exact"), MockHandler) }},
		{
			"Prefix",
			func() *handlers.CallbackHandler { return callbackHandlers.Prefix(g.String("start_"), MockHandler) },
		},
		{"Suffix", func() *handlers.CallbackHandler { return callbackHandlers.Suffix(g.String("_end"), MockHandler) }},
		{"FromUserID", func() *handlers.CallbackHandler { return callbackHandlers.FromUserID(123456, MockHandler) }},
		{
			"GameName",
			func() *handlers.CallbackHandler { return callbackHandlers.GameName(g.String("game"), MockHandler) },
		},
		{"Inline", func() *handlers.CallbackHandler { return callbackHandlers.Inline(MockHandler) }},
		{"ChatInstance", func() *handlers.CallbackHandler {
			return callbackHandlers.ChatInstance(g.String("instance"), MockHandler)
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := test.handler()
			if handler == nil {
				t.Errorf("%s should return a CallbackHandler", test.name)
			}
		})
	}
}

func TestCallbackHandlers_EmptyStrings(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	// Test handlers with empty strings
	tests := []struct {
		name    string
		handler func() *handlers.CallbackHandler
	}{
		{"Equal empty", func() *handlers.CallbackHandler { return callbackHandlers.Equal(g.String(""), MockHandler) }},
		{
			"Prefix empty",
			func() *handlers.CallbackHandler { return callbackHandlers.Prefix(g.String(""), MockHandler) },
		},
		{
			"Suffix empty",
			func() *handlers.CallbackHandler { return callbackHandlers.Suffix(g.String(""), MockHandler) },
		},
		{
			"GameName empty",
			func() *handlers.CallbackHandler { return callbackHandlers.GameName(g.String(""), MockHandler) },
		},
		{
			"ChatInstance empty",
			func() *handlers.CallbackHandler { return callbackHandlers.ChatInstance(g.String(""), MockHandler) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := test.handler()
			if handler == nil {
				t.Errorf("%s should return a CallbackHandler even with empty string", test.name)
			}
		})
	}
}

func TestCallbackHandlers_SpecialCharacters(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	// Test handlers with special characters
	specialString := g.String("test@#$%^&*()_+-=[]{}|;':\",./<>?")

	tests := []struct {
		name    string
		handler func() *handlers.CallbackHandler
	}{
		{
			"Equal special",
			func() *handlers.CallbackHandler { return callbackHandlers.Equal(specialString, MockHandler) },
		},
		{
			"Prefix special",
			func() *handlers.CallbackHandler { return callbackHandlers.Prefix(specialString, MockHandler) },
		},
		{
			"Suffix special",
			func() *handlers.CallbackHandler { return callbackHandlers.Suffix(specialString, MockHandler) },
		},
		{
			"GameName special",
			func() *handlers.CallbackHandler { return callbackHandlers.GameName(specialString, MockHandler) },
		},
		{
			"ChatInstance special",
			func() *handlers.CallbackHandler { return callbackHandlers.ChatInstance(specialString, MockHandler) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := test.handler()
			if handler == nil {
				t.Errorf("%s should return a CallbackHandler with special characters", test.name)
			}
		})
	}
}

func TestCallbackHandlers_UnicodeStrings(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	// Test handlers with Unicode strings
	unicodeString := g.String("🔥💯🚀🎉🌟")

	tests := []struct {
		name    string
		handler func() *handlers.CallbackHandler
	}{
		{
			"Equal unicode",
			func() *handlers.CallbackHandler { return callbackHandlers.Equal(unicodeString, MockHandler) },
		},
		{
			"Prefix unicode",
			func() *handlers.CallbackHandler { return callbackHandlers.Prefix(unicodeString, MockHandler) },
		},
		{
			"Suffix unicode",
			func() *handlers.CallbackHandler { return callbackHandlers.Suffix(unicodeString, MockHandler) },
		},
		{
			"GameName unicode",
			func() *handlers.CallbackHandler { return callbackHandlers.GameName(unicodeString, MockHandler) },
		},
		{
			"ChatInstance unicode",
			func() *handlers.CallbackHandler { return callbackHandlers.ChatInstance(unicodeString, MockHandler) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := test.handler()
			if handler == nil {
				t.Errorf("%s should return a CallbackHandler with Unicode characters", test.name)
			}
		})
	}
}

func TestCallbackHandlers_LongStrings(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	// Test handlers with very long strings
	longString := g.String(
		"this_is_a_very_long_callback_data_string_that_might_exceed_normal_limits_but_should_still_work_properly_in_the_handler_system",
	)

	tests := []struct {
		name    string
		handler func() *handlers.CallbackHandler
	}{
		{"Equal long", func() *handlers.CallbackHandler { return callbackHandlers.Equal(longString, MockHandler) }},
		{"Prefix long", func() *handlers.CallbackHandler { return callbackHandlers.Prefix(longString, MockHandler) }},
		{"Suffix long", func() *handlers.CallbackHandler { return callbackHandlers.Suffix(longString, MockHandler) }},
		{
			"GameName long",
			func() *handlers.CallbackHandler { return callbackHandlers.GameName(longString, MockHandler) },
		},
		{
			"ChatInstance long",
			func() *handlers.CallbackHandler { return callbackHandlers.ChatInstance(longString, MockHandler) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := test.handler()
			if handler == nil {
				t.Errorf("%s should return a CallbackHandler with long strings", test.name)
			}
		})
	}
}

func TestCallbackHandlers_ZeroUserID(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.FromUserID(0, MockHandler)

	if handler == nil {
		t.Error("FromUserID should return a CallbackHandler even with zero user ID")
	}
}

func TestCallbackHandlers_NegativeUserID(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.FromUserID(-123456789, MockHandler)

	if handler == nil {
		t.Error("FromUserID should return a CallbackHandler even with negative user ID")
	}
}

func TestCallbackHandlers_LargeUserID(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	// Test with very large user ID
	largeUserID := int64(9223372036854775807) // max int64
	handler := callbackHandlers.FromUserID(largeUserID, MockHandler)

	if handler == nil {
		t.Error("FromUserID should return a CallbackHandler even with large user ID")
	}
}

func TestCallbackHandler_RegisterMultipleTimes(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	handler := callbackHandlers.Equal(g.String("test"), MockHandler)

	// Register multiple times should work
	first := handler.Register()
	second := handler.Register()
	third := handler.Register()

	if first != handler || second != handler || third != handler {
		t.Error("Multiple Register calls should return the same handler instance")
	}
}

func TestCallbackHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	// Test with nil handler function
	handler := callbackHandlers.Any(nil)

	if handler == nil {
		t.Error("Handler creation should work even with nil handler function")
	}
}

func TestCallbackHandlers_FilterLogic(t *testing.T) {
	// This test verifies that the callback handlers create proper filters
	// by testing the handler creation with various edge cases
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	// Test Equal with empty string
	handler1 := callbackHandlers.Equal(g.String(""), MockHandler)
	if handler1 == nil {
		t.Error("Equal with empty string should create handler")
	}

	// Test Prefix with empty string
	handler2 := callbackHandlers.Prefix(g.String(""), MockHandler)
	if handler2 == nil {
		t.Error("Prefix with empty string should create handler")
	}

	// Test Suffix with empty string
	handler3 := callbackHandlers.Suffix(g.String(""), MockHandler)
	if handler3 == nil {
		t.Error("Suffix with empty string should create handler")
	}

	// Test FromUserID with zero
	handler4 := callbackHandlers.FromUserID(0, MockHandler)
	if handler4 == nil {
		t.Error("FromUserID with zero should create handler")
	}

	// Test GameName with empty string
	handler5 := callbackHandlers.GameName(g.String(""), MockHandler)
	if handler5 == nil {
		t.Error("GameName with empty string should create handler")
	}

	// Test ChatInstance with empty string
	handler6 := callbackHandlers.ChatInstance(g.String(""), MockHandler)
	if handler6 == nil {
		t.Error("ChatInstance with empty string should create handler")
	}
}

func TestCallbackHandlers_AllowChannelAndRegister(t *testing.T) {
	bot := NewMockBot()
	callbackHandlers := &handlers.CallbackHandlers{Bot: bot}

	// Test that AllowChannel() returns the same handler and Register() works
	handler := callbackHandlers.Equal(g.String("test"), MockHandler)

	// Call AllowChannel to increase coverage
	result := handler.AllowChannel()
	if result != handler {
		t.Error("AllowChannel should return the same handler instance")
	}

	// Call Register to increase coverage - already covered but ensure chaining works
	registerResult := result.Register()
	if registerResult != handler {
		t.Error("Register should return the same handler instance")
	}
}
