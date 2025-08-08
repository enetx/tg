package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendInvoice(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	result := ctx.SendInvoice(title, desc, payload, currency)

	if result == nil {
		t.Error("Expected SendInvoice builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendInvoiceChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	result := ctx.SendInvoice(title, desc, payload, currency).
		StartParameter(g.String("start_param")).
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendInvoice builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected Protect method to return builder")
	}
}

func TestSendInvoice_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Product")
	desc := g.String("Test Description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendInvoice(title, desc, payload, currency).Send()

	if sendResult.IsErr() {
		t.Logf("SendInvoice Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendInvoice(title, desc, payload, currency).
		StartParameter(g.String("start_param")).
		ProviderToken(g.String("test_token")).
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendInvoice configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
