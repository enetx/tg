package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestShippingHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	shippingHandlers := &handlers.ShippingHandlers{Bot: bot}

	result := shippingHandlers.Any(MockHandler)

	if result == nil {
		t.Error("Any should return ShippingHandlers")
	}

	if result != shippingHandlers {
		t.Error("Any should return the same ShippingHandlers instance for chaining")
	}
}

func TestShippingHandlers_FromUserID(t *testing.T) {
	bot := NewMockBot()
	shippingHandlers := &handlers.ShippingHandlers{Bot: bot}

	result := shippingHandlers.FromUserID(987654321, MockHandler)

	if result == nil {
		t.Error("FromUserID should return ShippingHandlers")
	}

	if result != shippingHandlers {
		t.Error("FromUserID should return the same ShippingHandlers instance for chaining")
	}
}

func TestShippingHandlers_InvoicePayload(t *testing.T) {
	bot := NewMockBot()
	shippingHandlers := &handlers.ShippingHandlers{Bot: bot}

	result := shippingHandlers.InvoicePayload(g.String("order_123"), MockHandler)

	if result == nil {
		t.Error("InvoicePayload should return ShippingHandlers")
	}

	if result != shippingHandlers {
		t.Error("InvoicePayload should return the same ShippingHandlers instance for chaining")
	}
}

func TestShippingHandlers_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	shippingHandlers := &handlers.ShippingHandlers{Bot: bot}

	result := shippingHandlers.
		Any(MockHandler).
		FromUserID(123456, MockHandler).
		InvoicePayload(g.String("test_payload"), MockHandler)

	if result != shippingHandlers {
		t.Error("Chained methods should return the same ShippingHandlers instance")
	}
}

func TestShippingHandlers_EmptyPayload(t *testing.T) {
	bot := NewMockBot()
	shippingHandlers := &handlers.ShippingHandlers{Bot: bot}

	result := shippingHandlers.InvoicePayload(g.String(""), MockHandler)

	if result == nil {
		t.Error("InvoicePayload with empty string should work")
	}
}

func TestShippingHandlers_ZeroUserID(t *testing.T) {
	bot := NewMockBot()
	shippingHandlers := &handlers.ShippingHandlers{Bot: bot}

	result := shippingHandlers.FromUserID(0, MockHandler)

	if result == nil {
		t.Error("FromUserID with zero ID should work")
	}
}

func TestShippingHandlers_SpecialCharacterPayload(t *testing.T) {
	bot := NewMockBot()
	shippingHandlers := &handlers.ShippingHandlers{Bot: bot}

	specialPayload := g.String("order@#$%^&*()_+-=[]{}|;':\",./<>?123")
	result := shippingHandlers.InvoicePayload(specialPayload, MockHandler)

	if result == nil {
		t.Error("InvoicePayload with special characters should work")
	}
}

func TestShippingHandlers_UnicodePayload(t *testing.T) {
	bot := NewMockBot()
	shippingHandlers := &handlers.ShippingHandlers{Bot: bot}

	unicodePayload := g.String("заказ_123_🛒_주문")
	result := shippingHandlers.InvoicePayload(unicodePayload, MockHandler)

	if result == nil {
		t.Error("InvoicePayload with Unicode characters should work")
	}
}

func TestShippingHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	shippingHandlers := &handlers.ShippingHandlers{Bot: bot}

	result := shippingHandlers.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}

func TestShippingHandlers_HasPayloadPrefix(t *testing.T) {
	bot := NewMockBot()
	shippingHandlers := &handlers.ShippingHandlers{Bot: bot}

	result := shippingHandlers.HasPayloadPrefix(g.String("order_"), MockHandler)

	if result == nil {
		t.Error("HasPayloadPrefix should return ShippingHandlers")
	}

	if result != shippingHandlers {
		t.Error("HasPayloadPrefix should return the same ShippingHandlers instance for chaining")
	}
}
