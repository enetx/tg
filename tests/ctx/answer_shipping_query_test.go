package ctx_test

import (
	"testing"
	"time"

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

	testCtx := ctx.New(bot, rawCtx)

	// Test basic creation
	result := testCtx.AnswerShippingQuery()
	if result == nil {
		t.Error("Expected AnswerShippingQuery builder to be created")
	}

	// Test Ok method
	result = result.Ok()
	if result == nil {
		t.Error("Ok method should return AnswerShippingQuery for chaining")
	}

	// Test Error method
	result = testCtx.AnswerShippingQuery().Error(g.String("Shipping not available"))
	if result == nil {
		t.Error("Error method should return AnswerShippingQuery for chaining")
	}

	// Test AddOption method
	option := gotgbot.ShippingOption{
		Id:     "express",
		Title:  "Express Shipping",
		Prices: []gotgbot.LabeledPrice{{Label: "Express", Amount: 1000}},
	}
	result = testCtx.AnswerShippingQuery().AddOption(option)
	if result == nil {
		t.Error("AddOption method should return AnswerShippingQuery for chaining")
	}

	// Test Options method
	options := g.SliceOf[gotgbot.ShippingOption](option)
	result = testCtx.AnswerShippingQuery().Options(options)
	if result == nil {
		t.Error("Options method should return AnswerShippingQuery for chaining")
	}

	// Test Timeout method
	result = testCtx.AnswerShippingQuery().Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return AnswerShippingQuery for chaining")
	}

	// Test APIURL method
	result = testCtx.AnswerShippingQuery().APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return AnswerShippingQuery for chaining")
	}
}

func TestContext_AnswerShippingQueryChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId: 1,
			ShippingQuery: &gotgbot.ShippingQuery{
				Id:   "shipping_456",
				From: gotgbot.User{Id: 789, FirstName: "TestUser"},
			},
		},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test complete method chaining
	result := testCtx.AnswerShippingQuery().
		Ok().
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return AnswerShippingQuery")
	}

	// Test error flow chaining
	errorResult := testCtx.AnswerShippingQuery().
		Error(g.String("We don't ship to your location")).
		Timeout(20 * time.Second)

	if errorResult == nil {
		t.Error("Error flow chaining should work and return AnswerShippingQuery")
	}
}

func TestShippingOptionBuilder(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId: 1,
			ShippingQuery: &gotgbot.ShippingQuery{
				Id:   "shipping_789",
				From: gotgbot.User{Id: 999, FirstName: "Builder"},
			},
		},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test Option builder creation
	optionBuilder := testCtx.AnswerShippingQuery().Option(g.String("standard"), g.String("Standard Shipping"))
	if optionBuilder == nil {
		t.Error("Option method should return ShippingOptionBuilder")
	}

	// Test Price method
	priceBuilder := optionBuilder.Price(g.String("Base Price"), 500)
	if priceBuilder == nil {
		t.Error("Price method should return ShippingOptionBuilder for chaining")
	}

	// Test multiple prices
	multiPriceBuilder := priceBuilder.Price(g.String("Insurance"), 100).Price(g.String("Handling"), 50)
	if multiPriceBuilder == nil {
		t.Error("Multiple Price calls should work with chaining")
	}

	// Test Done method
	backToQuery := multiPriceBuilder.Done()
	if backToQuery == nil {
		t.Error("Done method should return AnswerShippingQuery")
	}
}

func TestAnswerShippingQuery_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId: 1,
			ShippingQuery: &gotgbot.ShippingQuery{
				Id:   "edge_test",
				From: gotgbot.User{Id: 555, FirstName: "Edge"},
			},
		},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with empty error message
	result := testCtx.AnswerShippingQuery().Error(g.String(""))
	if result == nil {
		t.Error("AnswerShippingQuery should handle empty error message")
	}

	// Test with zero timeout
	result = testCtx.AnswerShippingQuery().Timeout(0 * time.Second)
	if result == nil {
		t.Error("AnswerShippingQuery should handle zero timeout")
	}

	// Test with very long timeout
	result = testCtx.AnswerShippingQuery().Timeout(24 * time.Hour)
	if result == nil {
		t.Error("AnswerShippingQuery should handle very long timeout")
	}

	// Test with empty API URL
	result = testCtx.AnswerShippingQuery().APIURL(g.String(""))
	if result == nil {
		t.Error("AnswerShippingQuery should handle empty API URL")
	}

	// Test with empty shipping option parameters
	optionBuilder := testCtx.AnswerShippingQuery().Option(g.String(""), g.String(""))
	if optionBuilder == nil {
		t.Error("Option should handle empty ID and title")
	}

	// Test with zero and negative prices
	priceBuilder := optionBuilder.Price(g.String("Free"), 0).Price(g.String("Discount"), -100)
	if priceBuilder == nil {
		t.Error("Price should handle zero and negative amounts")
	}
}

func TestAnswerShippingQuery_ComplexShippingOptions(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId: 1,
			ShippingQuery: &gotgbot.ShippingQuery{
				Id:   "complex_shipping",
				From: gotgbot.User{Id: 777, FirstName: "Complex"},
			},
		},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test multiple shipping options
	result := testCtx.AnswerShippingQuery().
		Ok().
		Option(g.String("standard"), g.String("Standard Shipping")).
		Price(g.String("Base Rate"), 500).
		Price(g.String("Insurance"), 100).
		Done().
		Option(g.String("express"), g.String("Express Shipping")).
		Price(g.String("Express Rate"), 1200).
		Price(g.String("Priority Handling"), 200).
		Done().
		Option(g.String("overnight"), g.String("Overnight Delivery")).
		Price(g.String("Overnight Rate"), 2500).
		Done()

	if result == nil {
		t.Error("Complex shipping options should work")
	}

	// Test pre-built options
	preBuiltOptions := g.SliceOf[gotgbot.ShippingOption](
		gotgbot.ShippingOption{
			Id:    "bulk",
			Title: "Bulk Shipping",
			Prices: []gotgbot.LabeledPrice{
				{Label: "Bulk Rate", Amount: 300},
				{Label: "Volume Discount", Amount: -50},
			},
		},
		gotgbot.ShippingOption{
			Id:    "premium",
			Title: "Premium Service",
			Prices: []gotgbot.LabeledPrice{
				{Label: "Premium Rate", Amount: 3000},
				{Label: "White Glove", Amount: 500},
				{Label: "Signature Required", Amount: 100},
			},
		},
	)

	result = testCtx.AnswerShippingQuery().Options(preBuiltOptions)
	if result == nil {
		t.Error("Pre-built shipping options should work")
	}
}

func TestAnswerShippingQuery_NoShippingQuery(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId:      1,
			ShippingQuery: nil, // No shipping query
		},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test that builder is created even without shipping query
	result := testCtx.AnswerShippingQuery()
	if result == nil {
		t.Error("AnswerShippingQuery should be created even without shipping query")
	}

	// The actual error would be caught in Send() method, but builder should work
	okResult := result.Ok()
	if okResult == nil {
		t.Error("Builder methods should work even without shipping query")
	}
}

func TestAnswerShippingQuery_PriceEdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId: 1,
			ShippingQuery: &gotgbot.ShippingQuery{
				Id:   "price_edge_cases",
				From: gotgbot.User{Id: 888, FirstName: "PriceTest"},
			},
		},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test extreme price values
	extremePrices := []struct {
		label  string
		amount int64
	}{
		{"Minimum", 1},
		{"Small", 50},
		{"Medium", 1000},
		{"Large", 100000},
		{"Maximum", 9223372036854775807}, // max int64
		{"Zero", 0},
		{"Negative", -1000},
		{"Large Negative", -9223372036854775808}, // min int64
	}

	optionBuilder := testCtx.AnswerShippingQuery().Option(g.String("extreme_prices"), g.String("Extreme Price Testing"))

	for _, price := range extremePrices {
		optionBuilder = optionBuilder.Price(g.String(price.label), price.amount)
		if optionBuilder == nil {
			t.Errorf("Price should handle extreme value: %s (%d)", price.label, price.amount)
		}
	}

	result := optionBuilder.Done()
	if result == nil {
		t.Error("Done should work after extreme price values")
	}
}
