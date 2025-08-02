package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestPreCheckoutHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	result := preCheckoutHandlers.Any(MockHandler)

	if result == nil {
		t.Error("Any should return PreCheckoutHandlers")
	}

	if result != preCheckoutHandlers {
		t.Error("Any should return the same PreCheckoutHandlers instance for chaining")
	}
}

func TestPreCheckoutHandlers_FromUserID(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	result := preCheckoutHandlers.FromUserID(987654321, MockHandler)

	if result == nil {
		t.Error("FromUserID should return PreCheckoutHandlers")
	}

	if result != preCheckoutHandlers {
		t.Error("FromUserID should return the same PreCheckoutHandlers instance for chaining")
	}
}

func TestPreCheckoutHandlers_HasPayloadPrefix(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	result := preCheckoutHandlers.HasPayloadPrefix(g.String("order_123"), MockHandler)

	if result == nil {
		t.Error("HasPayloadPrefix should return PreCheckoutHandlers")
	}

	if result != preCheckoutHandlers {
		t.Error("HasPayloadPrefix should return the same PreCheckoutHandlers instance for chaining")
	}
}

func TestPreCheckoutHandlers_Currency(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	result := preCheckoutHandlers.Currency(g.String("USD"), MockHandler)

	if result == nil {
		t.Error("Currency should return PreCheckoutHandlers")
	}

	if result != preCheckoutHandlers {
		t.Error("Currency should return the same PreCheckoutHandlers instance for chaining")
	}
}

func TestPreCheckoutHandlers_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	result := preCheckoutHandlers.
		Any(MockHandler).
		FromUserID(123456, MockHandler).
		HasPayloadPrefix(g.String("test_payload"), MockHandler).
		Currency(g.String("EUR"), MockHandler)

	if result != preCheckoutHandlers {
		t.Error("Chained methods should return the same PreCheckoutHandlers instance")
	}
}

func TestPreCheckoutHandlers_EmptyPayload(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	result := preCheckoutHandlers.HasPayloadPrefix(g.String(""), MockHandler)

	if result == nil {
		t.Error("HasPayloadPrefix with empty string should work")
	}
}

func TestPreCheckoutHandlers_EmptyCurrency(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	result := preCheckoutHandlers.Currency(g.String(""), MockHandler)

	if result == nil {
		t.Error("Currency with empty string should work")
	}
}

func TestPreCheckoutHandlers_ZeroUserID(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	result := preCheckoutHandlers.FromUserID(0, MockHandler)

	if result == nil {
		t.Error("FromUserID with zero ID should work")
	}
}

func TestPreCheckoutHandlers_VariousCurrencies(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	currencies := []string{"USD", "EUR", "GBP", "JPY", "CAD", "AUD", "CHF", "CNY", "RUB", "XTR"}

	for _, currency := range currencies {
		t.Run(currency, func(t *testing.T) {
			result := preCheckoutHandlers.Currency(g.String(currency), MockHandler)
			if result == nil {
				t.Errorf("Currency %s should work", currency)
			}
		})
	}
}

func TestPreCheckoutHandlers_SpecialCharacterPayload(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	specialPayload := g.String("order@#$%^&*()_+-=[]{}|;':\",./<>?123")
	result := preCheckoutHandlers.HasPayloadPrefix(specialPayload, MockHandler)

	if result == nil {
		t.Error("HasPayloadPrefix with special characters should work")
	}
}

func TestPreCheckoutHandlers_UnicodePayload(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	unicodePayload := g.String("заказ_123_💰_주문")
	result := preCheckoutHandlers.HasPayloadPrefix(unicodePayload, MockHandler)

	if result == nil {
		t.Error("HasPayloadPrefix with Unicode characters should work")
	}
}

func TestPreCheckoutHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	preCheckoutHandlers := &handlers.PreCheckoutHandlers{Bot: bot}

	result := preCheckoutHandlers.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}
