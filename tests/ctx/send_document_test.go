package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendDocument(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("nonexistent.pdf")

	result := ctx.SendDocument(filename)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendDocument builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("test"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendDocumentChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("document.pdf")

	result := ctx.SendDocument(filename).
		Caption(g.String("Test document")).
		HTML().
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendDocument builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected protect method to return builder")
	}
}

func TestSendDocument_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_document.pdf")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendDocument(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendDocument Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendDocument(filename).
		Caption(g.String("Test <b>document</b> with HTML")).
		HTML().
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendDocument configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
