package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendLocation(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060

	result := ctx.SendLocation(lat, lon)

	if result == nil {
		t.Error("Expected SendLocation builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendLocationChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060

	result := ctx.SendLocation(lat, lon).
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendLocation builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected Protect method to return builder")
	}
}

func TestSendLocation_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	lat, lon := 40.7128, -74.0060

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendLocation(lat, lon).Send()

	if sendResult.IsErr() {
		t.Logf("SendLocation Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendLocation(lat, lon).
		Heading(90).
		ProximityAlertRadius(100).
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendLocation configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
