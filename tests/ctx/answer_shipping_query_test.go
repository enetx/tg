package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_AnswerShippingQuery(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId: 1,
			ShippingQuery: &gotgbot.ShippingQuery{
				Id:   "shipping_123",
				From: gotgbot.User{Id: 456, FirstName: "Test"},
			},
		},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.AnswerShippingQuery()

	if result == nil {
		t.Error("Expected AnswerShippingQuery builder to be created")
	}

	// Test method chaining
	okResult := result.Ok()
	if okResult == nil {
		t.Error("Expected Ok method to return builder")
	}

	errorResult := result.Error(g.String("Shipping not available"))
	if errorResult == nil {
		t.Error("Expected Error method to return builder")
	}

	// Test Option builder
	optionBuilder := result.Option(g.String("standard"), g.String("Standard Shipping"))
	if optionBuilder == nil {
		t.Error("Expected Option method to return builder")
	}

	priceBuilder := optionBuilder.Price(g.String("Shipping"), 500)
	if priceBuilder == nil {
		t.Error("Expected Price method to return builder")
	}

	backToQuery := priceBuilder.Done()
	if backToQuery == nil {
		t.Error("Expected Done method to return AnswerShippingQuery")
	}
}
