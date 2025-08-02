package handlers_test

import (
	"testing"

	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestPaidMediaHandlers_Any(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	result := paidMediaHandlers.Any(MockHandler)

	if result == nil {
		t.Error("Any should return PaidMediaHandlers")
	}

	if result != paidMediaHandlers {
		t.Error("Any should return the same PaidMediaHandlers instance for chaining")
	}
}

func TestPaidMediaHandlers_FromUserID(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	result := paidMediaHandlers.FromUserID(987654321, MockHandler)

	if result == nil {
		t.Error("FromUserID should return PaidMediaHandlers")
	}

	if result != paidMediaHandlers {
		t.Error("FromUserID should return the same PaidMediaHandlers instance for chaining")
	}
}

func TestPaidMediaHandlers_Payload(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	result := paidMediaHandlers.Payload(g.String("test_payload"), MockHandler)

	if result == nil {
		t.Error("Payload should return PaidMediaHandlers")
	}

	if result != paidMediaHandlers {
		t.Error("Payload should return the same PaidMediaHandlers instance for chaining")
	}
}

func TestPaidMediaHandlers_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	result := paidMediaHandlers.
		Any(MockHandler).
		FromUserID(123456, MockHandler).
		Payload(g.String("test_payload"), MockHandler)

	if result != paidMediaHandlers {
		t.Error("Chained methods should return the same PaidMediaHandlers instance")
	}
}

func TestPaidMediaHandlers_ZeroPayload(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	result := paidMediaHandlers.Payload(g.String(""), MockHandler)

	if result == nil {
		t.Error("Payload with zero should work")
	}
}

func TestPaidMediaHandlers_ZeroUserID(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	result := paidMediaHandlers.FromUserID(0, MockHandler)

	if result == nil {
		t.Error("FromUserID with zero ID should work")
	}
}

func TestPaidMediaHandlers_NegativeUserID(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	result := paidMediaHandlers.FromUserID(-123456789, MockHandler)

	if result == nil {
		t.Error("FromUserID with negative ID should work")
	}
}

func TestPaidMediaHandlers_LargeUserID(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	largeID := int64(9223372036854775807) // max int64
	result := paidMediaHandlers.FromUserID(largeID, MockHandler)

	if result == nil {
		t.Error("FromUserID with large ID should work")
	}
}

func TestPaidMediaHandlers_LargePayload(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	largePayload := g.String("very_long_payload_string_that_exceeds_normal_expectations_123456789")
	result := paidMediaHandlers.Payload(largePayload, MockHandler)

	if result == nil {
		t.Error("Payload with large count should work")
	}
}

func TestPaidMediaHandlers_VariousPayloads(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	payloads := []string{"basic", "premium", "enterprise", "custom_123", "special_offer"}

	for _, payload := range payloads {
		t.Run(payload, func(t *testing.T) {
			result := paidMediaHandlers.Payload(g.String(payload), MockHandler)
			if result == nil {
				t.Errorf("Payload %s should work", payload)
			}
		})
	}
}

func TestPaidMediaHandlers_PayloadPrefix(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	result := paidMediaHandlers.PayloadPrefix(g.String("prefix_"), MockHandler)

	if result == nil {
		t.Error("PayloadPrefix should return PaidMediaHandlers")
	}

	if result != paidMediaHandlers {
		t.Error("PayloadPrefix should return the same PaidMediaHandlers instance for chaining")
	}
}

func TestPaidMediaHandlers_WithNilHandler(t *testing.T) {
	bot := NewMockBot()
	paidMediaHandlers := &handlers.PaidMediaHandlers{Bot: bot}

	result := paidMediaHandlers.Any(nil)

	if result == nil {
		t.Error("Handler registration with nil handler should work")
	}
}
