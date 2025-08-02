package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_CreateInvoiceLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Product")
	description := g.String("Test product description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	result := ctx.CreateInvoiceLink(title, description, payload, currency)

	if result == nil {
		t.Error("Expected CreateInvoiceLink builder to be created")
	}

	// Test method chaining
	withPrice := result.Price(g.String("Product"), 1000)
	if withPrice == nil {
		t.Error("Expected Price method to return builder")
	}

	withProvider := withPrice.ProviderToken(g.String("provider_token"))
	if withProvider == nil {
		t.Error("Expected ProviderToken method to return builder")
	}

	needName := withProvider.NeedName()
	if needName == nil {
		t.Error("Expected NeedName method to return builder")
	}

	flexible := needName.Flexible()
	if flexible == nil {
		t.Error("Expected Flexible method to return builder")
	}
}
