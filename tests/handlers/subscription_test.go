package handlers_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/handlers"
)

func TestSubscription_Any(t *testing.T) {
	bot := NewMockBot()
	subscription := &handlers.Subscription{Bot: bot}

	result := subscription.Any(MockHandler)

	if result == nil {
		t.Error("Any should return Subscription")
	}

	if result != subscription {
		t.Error("Any should return the same Subscription instance for chaining")
	}
}

func TestSubscription_States(t *testing.T) {
	bot := NewMockBot()
	subscription := &handlers.Subscription{Bot: bot}

	if subscription.Active(MockHandler) != subscription {
		t.Error("Active should return the same Subscription instance for chaining")
	}

	if subscription.Canceled(MockHandler) != subscription {
		t.Error("Canceled should return the same Subscription instance for chaining")
	}

	if subscription.Failed(MockHandler) != subscription {
		t.Error("Failed should return the same Subscription instance for chaining")
	}
}

func TestSubscription_UserID(t *testing.T) {
	bot := NewMockBot()
	subscription := &handlers.Subscription{Bot: bot}

	if subscription.UserID(123, MockHandler) != subscription {
		t.Error("UserID should return the same Subscription instance for chaining")
	}
}

func TestSubscription_Payload(t *testing.T) {
	bot := NewMockBot()
	subscription := &handlers.Subscription{Bot: bot}

	if subscription.Payload(g.String("payload_123"), MockHandler) != subscription {
		t.Error("Payload should return the same Subscription instance for chaining")
	}
}

func TestSubscription_ChainedMethods(t *testing.T) {
	bot := NewMockBot()
	subscription := &handlers.Subscription{Bot: bot}

	result := subscription.
		Any(MockHandler).
		Active(MockHandler).
		Canceled(MockHandler).
		Failed(MockHandler).
		UserID(123, MockHandler).
		Payload(g.String("payload_123"), MockHandler)

	if result != subscription {
		t.Error("Chained methods should return the same Subscription instance")
	}
}

func TestSubscriptionHandler_CheckUpdate(t *testing.T) {
	handler := handlers.SubscriptionHandler{
		Response: func(b *gotgbot.Bot, ctx *ext.Context) error { return nil },
	}

	if handler.CheckUpdate(nil, &ext.Context{Update: &gotgbot.Update{}}) {
		t.Error("CheckUpdate should return false without a subscription update")
	}

	ctx := &ext.Context{
		Update: &gotgbot.Update{
			Subscription: &gotgbot.BotSubscriptionUpdated{
				User:           gotgbot.User{Id: 123},
				InvoicePayload: "payload_123",
				State:          "active",
			},
		},
	}

	if !handler.CheckUpdate(nil, ctx) {
		t.Error("CheckUpdate should return true without a filter")
	}

	handler.Filter = func(s *gotgbot.BotSubscriptionUpdated) bool { return s.State == "active" }
	if !handler.CheckUpdate(nil, ctx) {
		t.Error("CheckUpdate should return true when the filter matches")
	}

	handler.Filter = func(s *gotgbot.BotSubscriptionUpdated) bool { return s.State == "canceled" }
	if handler.CheckUpdate(nil, ctx) {
		t.Error("CheckUpdate should return false when the filter does not match")
	}
}

func TestSubscriptionHandler_HandleUpdate(t *testing.T) {
	called := false
	handler := handlers.SubscriptionHandler{
		Response: func(b *gotgbot.Bot, ctx *ext.Context) error {
			called = true
			return nil
		},
	}

	if err := handler.HandleUpdate(nil, &ext.Context{}); err != nil {
		t.Errorf("HandleUpdate should not return error: %v", err)
	}

	if !called {
		t.Error("HandleUpdate should call the response function")
	}
}

func TestSubscriptionHandler_Name(t *testing.T) {
	handler := handlers.SubscriptionHandler{
		Response: func(b *gotgbot.Bot, ctx *ext.Context) error { return nil },
	}

	name := handler.Name()
	if name == "" {
		t.Error("Name should return non-empty string")
	}

	handler2 := handlers.SubscriptionHandler{
		Response: func(b *gotgbot.Bot, ctx *ext.Context) error { return nil },
	}

	if name == handler2.Name() {
		t.Error("Different handlers should have different names")
	}
}
